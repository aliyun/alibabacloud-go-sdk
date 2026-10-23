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
	CompressStorage *string `json:"CompressStorage,omitempty" xml:"CompressStorage,omitempty"`
	// The connection resource quota of the AgenticDB cluster.
	//
	// example:
	//
	// 100
	ConnectionResourceQuota *int64 `json:"ConnectionResourceQuota,omitempty" xml:"ConnectionResourceQuota,omitempty"`
	// The cluster ID.
	//
	// > You can call the DescribeDBClusters operation to query all clusters in the specified region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The list of drill node instance names.
	//
	// > Node dimension drills support only a single node. For zone dimension drills, you can omit this parameter or pass all nodes.
	//
	// example:
	//
	// pi-rwxxx
	DBNodeCrashList *string `json:"DBNodeCrashList,omitempty" xml:"DBNodeCrashList,omitempty"`
	// The cross-zone data replication mode of the cluster. Valid values:
	//
	// - **AsyncSync**: Asynchronous.
	//
	// - **SemiSync**: Semi-synchronous.
	//
	// example:
	//
	// AsynSync
	DataSyncMode *string `json:"DataSyncMode,omitempty" xml:"DataSyncMode,omitempty"`
	// The fault injection method. Valid values:
	//
	// - 0: Instance fault injection based on `Crash SQL`.
	//
	// example:
	//
	// 0
	FaultInjectionType *string `json:"FaultInjectionType,omitempty" xml:"FaultInjectionType,omitempty"`
	// The fault drill dimension for the cluster. Valid values:
	//
	// - `0` or `FaultInjection`: Primary zone dimension fault drill.
	//
	// - `1`: Node dimension fault drill.
	//
	// > - In the **primary zone dimension fault drill*	- scenario, all compute nodes in the primary zone become unavailable, and the disaster recovery switchover in this scenario is lossy.
	//
	// > - In the **node dimension fault drill*	- scenario, only a single compute node can be drilled. Pass the desired compute node name through `DBNodeCrashList`.
	//
	// example:
	//
	// 0
	FaultSimulateMode *string `json:"FaultSimulateMode,omitempty" xml:"FaultSimulateMode,omitempty"`
	// The automatic IMCI-based query acceleration feature. Valid values:
	//
	// - `ON`: Enabled.
	//
	// - `OFF`: Disabled.
	//
	// > - Only PolarDB for MySQL clusters are supported.
	//
	// > - For cluster version limits, refer to [Automatic IMCI-based query acceleration (AutoIndex)](https://help.aliyun.com/document_detail/2854119.html).
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
	// - **ON**: Enables cross-zone automatic switchover.
	//
	// - **OFF**: Disables cross-zone automatic switchover.
	//
	// example:
	//
	// ON
	StandbyHAMode *string `json:"StandbyHAMode,omitempty" xml:"StandbyHAMode,omitempty"`
	// Specifies whether to enable automatic storage scaling for a Standard Edition cluster. Valid values:
	//
	// - Enable: Enables automatic storage scaling.
	//
	// - Disable: Disables automatic storage scaling.
	//
	// example:
	//
	// Enable
	StorageAutoScale *string `json:"StorageAutoScale,omitempty" xml:"StorageAutoScale,omitempty"`
	// The upper limit for automatic storage scaling of a Standard Edition cluster. Unit: GB.
	//
	// > The maximum value is 32000.
	//
	// example:
	//
	// 800
	StorageUpperBound *int64 `json:"StorageUpperBound,omitempty" xml:"StorageUpperBound,omitempty"`
	// The JSON string of the target database and table information to be restored. The values of the database and table information are all of the string type.
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
	// > You can call the [DescribeMetaList](https://help.aliyun.com/document_detail/194770.html) operation to view the details of the database and table names that support restoration, and fill the supported database and table information into the corresponding positions in the preceding example.
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
