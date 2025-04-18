package controller

import (
	"github.com/cybozu-go/mantle/test/util"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("util.getCephClusterIDFromPVC", func() {
	var namespace string
	storageClassName := util.GetUniqueName("sc-")

	It("should create namespace", func() {
		namespace = resMgr.CreateNamespace()
	})

	DescribeTable("matrix test",
		func(ctx SpecContext, sc *storagev1.StorageClass, pvc *corev1.PersistentVolumeClaim, expectedClusterID string) {
			// create resources
			if sc != nil {
				err := k8sClient.Create(ctx, sc)
				Expect(err).NotTo(HaveOccurred())
			}
			pvc.Namespace = namespace
			err := k8sClient.Create(ctx, pvc)
			Expect(err).NotTo(HaveOccurred())

			// test main
			clusterID, err := getCephClusterIDFromPVC(ctx, k8sClient, pvc)
			Expect(err).NotTo(HaveOccurred())
			Expect(clusterID).To(Equal(expectedClusterID))

			// cleanup
			if sc != nil {
				err := k8sClient.Delete(ctx, sc)
				Expect(err).NotTo(HaveOccurred())
			}
			err = k8sClient.Delete(ctx, pvc)
			Expect(err).NotTo(HaveOccurred())
		},
		Entry("can get clusterID",
			&storagev1.StorageClass{
				ObjectMeta: metav1.ObjectMeta{
					Name: storageClassName,
				},
				Provisioner: "test-ceph-cluster.rbd.csi.ceph.com",
				Parameters: map[string]string{
					"clusterID": "test-ceph-cluster",
				},
			},
			&corev1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{
					Name: util.GetUniqueName("pvc-"),
				},
				Spec: corev1.PersistentVolumeClaimSpec{
					AccessModes: []corev1.PersistentVolumeAccessMode{corev1.ReadWriteOnce},
					Resources: corev1.VolumeResourceRequirements{
						Requests: corev1.ResourceList{
							corev1.ResourceStorage: *resource.NewQuantity(1, resource.BinarySI),
						},
					},
					StorageClassName: &storageClassName,
				},
			},
			"test-ceph-cluster"),
		Entry("no storage class",
			nil,
			&corev1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-pvc",
				},
				Spec: corev1.PersistentVolumeClaimSpec{
					AccessModes: []corev1.PersistentVolumeAccessMode{corev1.ReadWriteOnce},
					Resources: corev1.VolumeResourceRequirements{
						Requests: corev1.ResourceList{
							corev1.ResourceStorage: *resource.NewQuantity(1, resource.BinarySI),
						},
					},
				},
			},
			""),
		Entry("storage class is not managed by RBD",
			&storagev1.StorageClass{
				ObjectMeta: metav1.ObjectMeta{
					Name: storageClassName,
				},
				Provisioner: "topolvm.cybozu.com",
			},
			&corev1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{
					Name: util.GetUniqueName("pvc-"),
				},
				Spec: corev1.PersistentVolumeClaimSpec{
					AccessModes: []corev1.PersistentVolumeAccessMode{corev1.ReadWriteOnce},
					Resources: corev1.VolumeResourceRequirements{
						Requests: corev1.ResourceList{
							corev1.ResourceStorage: *resource.NewQuantity(1, resource.BinarySI),
						},
					},
					StorageClassName: &storageClassName,
				},
			},
			""),
		Entry("clusterID is not set",
			&storagev1.StorageClass{
				ObjectMeta: metav1.ObjectMeta{
					Name: storageClassName,
				},
				Provisioner: "test-ceph-cluster.rbd.csi.ceph.com",
			},
			&corev1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{
					Name: util.GetUniqueName("pvc-"),
				},
				Spec: corev1.PersistentVolumeClaimSpec{
					AccessModes: []corev1.PersistentVolumeAccessMode{corev1.ReadWriteOnce},
					Resources: corev1.VolumeResourceRequirements{
						Requests: corev1.ResourceList{
							corev1.ResourceStorage: *resource.NewQuantity(1, resource.BinarySI),
						},
					},
					StorageClassName: &storageClassName,
				},
			},
			""),
	)
})
