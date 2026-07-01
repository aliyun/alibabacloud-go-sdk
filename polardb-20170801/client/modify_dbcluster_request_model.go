// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompressStorage(v string) *ModifyDBClusterRequest
	GetCompressStorage() *string
	SetConnectionResourceQuota(v int64) *ModifyDBClusterRequest
	GetConnectionResourceQuota() *int64
	SetDBClusterId(v string) *ModifyDBClusterRequest
	GetDBClusterId() *string
	SetDBNodeCrashList(v string) *ModifyDBClusterRequest
	GetDBNodeCrashList() *string
	SetDataSyncMode(v string) *ModifyDBClusterRequest
	GetDataSyncMode() *string
	SetFaultInjectionType(v string) *ModifyDBClusterRequest
	GetFaultInjectionType() *string
	SetFaultSimulateMode(v string) *ModifyDBClusterRequest
	GetFaultSimulateMode() *string
	SetImciAutoIndex(v string) *ModifyDBClusterRequest
	GetImciAutoIndex() *string
	SetModifyRowCompression(v string) *ModifyDBClusterRequest
	GetModifyRowCompression() *string
	SetOwnerAccount(v string) *ModifyDBClusterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBClusterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterRequest
	GetResourceOwnerId() *int64
	SetStandbyHAMode(v string) *ModifyDBClusterRequest
	GetStandbyHAMode() *string
	SetStorageAutoScale(v string) *ModifyDBClusterRequest
	GetStorageAutoScale() *string
	SetStorageUpperBound(v int64) *ModifyDBClusterRequest
	GetStorageUpperBound() *int64
	SetTableMeta(v string) *ModifyDBClusterRequest
	GetTableMeta() *string
}

type ModifyDBClusterRequest struct {
	// Enables storage compression. Set the value to **ON**.
	//
	// example:
	//
	// ON
	CompressStorage         *string `json:"CompressStorage,omitempty" xml:"CompressStorage,omitempty"`
	ConnectionResourceQuota *int64  `json:"ConnectionResourceQuota,omitempty" xml:"ConnectionResourceQuota,omitempty"`
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/173433.html) operation to query information about all clusters in the specified region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The list of node instance names for the disaster recovery drill.
	//
	// > Node-level drills support only a single node. For zone-level drills, you can leave this parameter empty or specify all nodes.
	//
	// example:
	//
	// pi-rwxxx
	DBNodeCrashList *string `json:"DBNodeCrashList,omitempty" xml:"DBNodeCrashList,omitempty"`
	// The cross-zone data replication mode of the cluster. Valid values:
	//
	// - **AsyncSync**: asynchronous
	//
	// - **SemiSync**: semi-synchronous
	//
	// example:
	//
	// AsynSync
	DataSyncMode *string `json:"DataSyncMode,omitempty" xml:"DataSyncMode,omitempty"`
	// The fault injection method. Valid values:
	//
	// - 0: instance fault injection based on `Crash SQL`
	//
	// example:
	//
	// 0
	FaultInjectionType *string `json:"FaultInjectionType,omitempty" xml:"FaultInjectionType,omitempty"`
	// The dimension of the disaster recovery drill for the cluster. Valid values:
	//
	// - `0` or `FaultInjection`: primary zone-level disaster recovery drill.
	//
	// - `1`: node-level disaster recovery drill.
	//
	// > - In the **primary zone-level disaster recovery drill*	- scenario, all compute nodes in the primary zone become unavailable. The failover in this scenario causes service interruptions.
	//
	// > - In the **node-level disaster recovery drill*	- scenario, only a single compute node is supported for the drill. Specify the desired compute node name by using `DBNodeCrashList`.
	//
	// example:
	//
	// 0
	FaultSimulateMode *string `json:"FaultSimulateMode,omitempty" xml:"FaultSimulateMode,omitempty"`
	// The automatic IMCI-based query acceleration feature. Valid values:
	//
	// - `ON`: enabled.
	//
	// - `OFF`: disabled.
	//
	// > - Only PolarDB for MySQL clusters are supported.
	//
	// > - For cluster version requirements, see [Automatic acceleration (AutoIndex)](https://help.aliyun.com/document_detail/2854119.html).
	//
	// example:
	//
	// OFF
	ImciAutoIndex *string `json:"ImciAutoIndex,omitempty" xml:"ImciAutoIndex,omitempty"`
	// Modifies the row compression settings.
	//
	// example:
	//
	// OFF
	ModifyRowCompression *string `json:"ModifyRowCompression,omitempty" xml:"ModifyRowCompression,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The cross-zone automatic switchover mode of the cluster. Valid values:
	//
	// - **ON**: enables cross-zone automatic switchover.
	//
	// - **OFF**: disables cross-zone automatic switchover.
	//
	// example:
	//
	// ON
	StandbyHAMode *string `json:"StandbyHAMode,omitempty" xml:"StandbyHAMode,omitempty"`
	// Specifies whether to enable automatic storage scaling for the Standard Edition cluster. Valid values:
	//
	// - Enable: enables automatic storage scaling.
	//
	// - Disable: disables automatic storage scaling.
	//
	// example:
	//
	// Enable
	StorageAutoScale *string `json:"StorageAutoScale,omitempty" xml:"StorageAutoScale,omitempty"`
	// The upper limit for automatic storage scaling of the Standard Edition cluster. Unit: GB.
	//
	// > The maximum value is 32000.
	//
	// example:
	//
	// 800
	StorageUpperBound *int64 `json:"StorageUpperBound,omitempty" xml:"StorageUpperBound,omitempty"`
	// The JSON string that contains the information about the databases and tables to be restored. The values of the database and table information are strings.
	//
	// Example: `[
	//
	//    {
	//
	//        "tables":[
	//
	//            {
	//
	//                "name":"testtb",
	//
	//                "type":"table",
	//
	//                "newname":"testtb_restore"
	//
	//            }
	//
	//        ],
	//
	//        "name":"testdb",
	//
	//        "type":"db",
	//
	//        "newname":"testdb_restore"
	//
	//    }
	//
	// ]`.
	//
	// > You can call the [DescribeMetaList](https://help.aliyun.com/document_detail/194770.html) operation to query the names of databases and tables that can be restored, and then specify the information in the corresponding fields in the preceding example.
	//
	// example:
	//
	// [ { "tables":[ { "name":"testtb", "type":"table", "newname":"testtb_restore" } ], "name":"testdb", "type":"db", "newname":"testdb_restore" } ]
	TableMeta *string `json:"TableMeta,omitempty" xml:"TableMeta,omitempty"`
}

func (s ModifyDBClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterRequest) GetCompressStorage() *string {
	return s.CompressStorage
}

func (s *ModifyDBClusterRequest) GetConnectionResourceQuota() *int64 {
	return s.ConnectionResourceQuota
}

func (s *ModifyDBClusterRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterRequest) GetDBNodeCrashList() *string {
	return s.DBNodeCrashList
}

func (s *ModifyDBClusterRequest) GetDataSyncMode() *string {
	return s.DataSyncMode
}

func (s *ModifyDBClusterRequest) GetFaultInjectionType() *string {
	return s.FaultInjectionType
}

func (s *ModifyDBClusterRequest) GetFaultSimulateMode() *string {
	return s.FaultSimulateMode
}

func (s *ModifyDBClusterRequest) GetImciAutoIndex() *string {
	return s.ImciAutoIndex
}

func (s *ModifyDBClusterRequest) GetModifyRowCompression() *string {
	return s.ModifyRowCompression
}

func (s *ModifyDBClusterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterRequest) GetStandbyHAMode() *string {
	return s.StandbyHAMode
}

func (s *ModifyDBClusterRequest) GetStorageAutoScale() *string {
	return s.StorageAutoScale
}

func (s *ModifyDBClusterRequest) GetStorageUpperBound() *int64 {
	return s.StorageUpperBound
}

func (s *ModifyDBClusterRequest) GetTableMeta() *string {
	return s.TableMeta
}

func (s *ModifyDBClusterRequest) SetCompressStorage(v string) *ModifyDBClusterRequest {
	s.CompressStorage = &v
	return s
}

func (s *ModifyDBClusterRequest) SetConnectionResourceQuota(v int64) *ModifyDBClusterRequest {
	s.ConnectionResourceQuota = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBClusterId(v string) *ModifyDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBNodeCrashList(v string) *ModifyDBClusterRequest {
	s.DBNodeCrashList = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDataSyncMode(v string) *ModifyDBClusterRequest {
	s.DataSyncMode = &v
	return s
}

func (s *ModifyDBClusterRequest) SetFaultInjectionType(v string) *ModifyDBClusterRequest {
	s.FaultInjectionType = &v
	return s
}

func (s *ModifyDBClusterRequest) SetFaultSimulateMode(v string) *ModifyDBClusterRequest {
	s.FaultSimulateMode = &v
	return s
}

func (s *ModifyDBClusterRequest) SetImciAutoIndex(v string) *ModifyDBClusterRequest {
	s.ImciAutoIndex = &v
	return s
}

func (s *ModifyDBClusterRequest) SetModifyRowCompression(v string) *ModifyDBClusterRequest {
	s.ModifyRowCompression = &v
	return s
}

func (s *ModifyDBClusterRequest) SetOwnerAccount(v string) *ModifyDBClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterRequest) SetOwnerId(v int64) *ModifyDBClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterRequest) SetResourceOwnerId(v int64) *ModifyDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetStandbyHAMode(v string) *ModifyDBClusterRequest {
	s.StandbyHAMode = &v
	return s
}

func (s *ModifyDBClusterRequest) SetStorageAutoScale(v string) *ModifyDBClusterRequest {
	s.StorageAutoScale = &v
	return s
}

func (s *ModifyDBClusterRequest) SetStorageUpperBound(v int64) *ModifyDBClusterRequest {
	s.StorageUpperBound = &v
	return s
}

func (s *ModifyDBClusterRequest) SetTableMeta(v string) *ModifyDBClusterRequest {
	s.TableMeta = &v
	return s
}

func (s *ModifyDBClusterRequest) Validate() error {
	return dara.Validate(s)
}
