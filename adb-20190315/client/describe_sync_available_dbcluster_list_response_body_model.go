// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSyncAvailableDBClusterListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSyncAvailableDBClusterListResponseBody
	GetRequestId() *string
	SetSyncAvailableDBClusters(v []*DescribeSyncAvailableDBClusterListResponseBodySyncAvailableDBClusters) *DescribeSyncAvailableDBClusterListResponseBody
	GetSyncAvailableDBClusters() []*DescribeSyncAvailableDBClusterListResponseBodySyncAvailableDBClusters
}

type DescribeSyncAvailableDBClusterListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FBD1DD96-AD1D-516C-9D9A-60BA081F66EE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried instances or clusters.
	SyncAvailableDBClusters []*DescribeSyncAvailableDBClusterListResponseBodySyncAvailableDBClusters `json:"SyncAvailableDBClusters,omitempty" xml:"SyncAvailableDBClusters,omitempty" type:"Repeated"`
}

func (s DescribeSyncAvailableDBClusterListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncAvailableDBClusterListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSyncAvailableDBClusterListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSyncAvailableDBClusterListResponseBody) GetSyncAvailableDBClusters() []*DescribeSyncAvailableDBClusterListResponseBodySyncAvailableDBClusters {
	return s.SyncAvailableDBClusters
}

func (s *DescribeSyncAvailableDBClusterListResponseBody) SetRequestId(v string) *DescribeSyncAvailableDBClusterListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSyncAvailableDBClusterListResponseBody) SetSyncAvailableDBClusters(v []*DescribeSyncAvailableDBClusterListResponseBodySyncAvailableDBClusters) *DescribeSyncAvailableDBClusterListResponseBody {
	s.SyncAvailableDBClusters = v
	return s
}

func (s *DescribeSyncAvailableDBClusterListResponseBody) Validate() error {
	if s.SyncAvailableDBClusters != nil {
		for _, item := range s.SyncAvailableDBClusters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSyncAvailableDBClusterListResponseBodySyncAvailableDBClusters struct {
	// The description of the instance or cluster.
	//
	// example:
	//
	// DB1
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The instance or cluster ID.
	//
	// example:
	//
	// rm-bp1l3yh04y7us147n
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database type of the instance or cluster.
	//
	// example:
	//
	// rds
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The storage size.
	//
	// example:
	//
	// 1000
	StorageSize *float32 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// The number of tables.
	//
	// example:
	//
	// 100
	TableNumber *int32 `json:"TableNumber,omitempty" xml:"TableNumber,omitempty"`
}

func (s DescribeSyncAvailableDBClusterListResponseBodySyncAvailableDBClusters) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncAvailableDBClusterListResponseBodySyncAvailableDBClusters) GoString() string {
	return s.String()
}

func (s *DescribeSyncAvailableDBClusterListResponseBodySyncAvailableDBClusters) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *DescribeSyncAvailableDBClusterListResponseBodySyncAvailableDBClusters) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSyncAvailableDBClusterListResponseBodySyncAvailableDBClusters) GetDBType() *string {
	return s.DBType
}

func (s *DescribeSyncAvailableDBClusterListResponseBodySyncAvailableDBClusters) GetStorageSize() *float32 {
	return s.StorageSize
}

func (s *DescribeSyncAvailableDBClusterListResponseBodySyncAvailableDBClusters) GetTableNumber() *int32 {
	return s.TableNumber
}

func (s *DescribeSyncAvailableDBClusterListResponseBodySyncAvailableDBClusters) SetDBClusterDescription(v string) *DescribeSyncAvailableDBClusterListResponseBodySyncAvailableDBClusters {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeSyncAvailableDBClusterListResponseBodySyncAvailableDBClusters) SetDBClusterId(v string) *DescribeSyncAvailableDBClusterListResponseBodySyncAvailableDBClusters {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSyncAvailableDBClusterListResponseBodySyncAvailableDBClusters) SetDBType(v string) *DescribeSyncAvailableDBClusterListResponseBodySyncAvailableDBClusters {
	s.DBType = &v
	return s
}

func (s *DescribeSyncAvailableDBClusterListResponseBodySyncAvailableDBClusters) SetStorageSize(v float32) *DescribeSyncAvailableDBClusterListResponseBodySyncAvailableDBClusters {
	s.StorageSize = &v
	return s
}

func (s *DescribeSyncAvailableDBClusterListResponseBodySyncAvailableDBClusters) SetTableNumber(v int32) *DescribeSyncAvailableDBClusterListResponseBodySyncAvailableDBClusters {
	s.TableNumber = &v
	return s
}

func (s *DescribeSyncAvailableDBClusterListResponseBodySyncAvailableDBClusters) Validate() error {
	return dara.Validate(s)
}
