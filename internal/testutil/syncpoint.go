package testutil

import (
	"sync/atomic"
)

type SyncPoint struct {
	syncCh atomic.Value // of type chan chan struct{}
}

var nilSyncCh interface{} = (chan struct{})(nil)

func (sp *SyncPoint) Block() (synced <-chan struct{}) {
	syncCh := make(chan struct{})
	sp.syncCh.Store(syncCh)
	return syncCh
}

func (sp *SyncPoint) Sync() {
	syncChIface := sp.syncCh.Load()
	if syncChIface == nil || syncChIface == nilSyncCh {
		return
	}
	sp.syncCh.Store(nilSyncCh)
	syncCh := syncChIface.(chan struct{})
	syncCh <- struct{}{}
}
