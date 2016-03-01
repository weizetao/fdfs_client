package fdfs_client

import (
	"testing"
)

func TestSyncByFilename(t *testing.T) {

	fdfsClient := FdfsClient{}

	storeServ := StorageServer{}
	storeServ.ipAddr = "10.140.80.93"
	storeServ.port = 23000

	err := fdfsClient.SyncFileNotify(&storeServ, 1000, "/data/slot0/data/09/0B/73/TYoBAFbFa3YAAAF1CXt0cTWw1_o1679991")
	if err != nil {
		t.Errorf("SyncByfilename error %s", err.Error())
	}
}
