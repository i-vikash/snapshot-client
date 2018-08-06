package main

import (
	"fmt"

	snapcrdv1 "github.com/openebs/external-storage/snapshot/pkg/apis/volumesnapshot/v1"
	snapshotclient "github.com/openebs/external-storage/snapshot/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := snapshotclient.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	volumeSnapshot := &snapcrdv1.VolumeSnapshot{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "snapshot-test",
			Namespace: "default",
		},
		Spec: snapcrdv1.VolumeSnapshotSpec{
			PersistentVolumeClaimName: "demo-vol1-claim",
		},
	}

	createdVolumeSnapshot, err := clientset.VolumesnapshotV1().VolumeSnapshots("default").Create(volumeSnapshot)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Created VolumeSnapshot : %+v", createdVolumeSnapshot)

	snapshots, err := clientset.VolumesnapshotV1().VolumeSnapshots("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("There are %d snapshots in cluster", len(snapshots.Items))
}
