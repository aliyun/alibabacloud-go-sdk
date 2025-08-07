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
	// Specifies whether to enable storage compression. Set the value to **ON**.
	//
	// example:
	//
	// ON
	CompressStorage *string `json:"CompressStorage,omitempty" xml:"CompressStorage,omitempty"`
	// The cluster ID.
	//
	// >  You can call the DescribeDBClusters operation to query information about all PolarDB clusters that are deployed in a specified region, such as cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The list of nodes for the drill.
	//
	// >  You can specify only one node for a node-level disaster recovery drill. For a primary zone-level disaster recovery drill, you can either choose not to specify this parameter or specify all nodes.
	//
	// example:
	//
	// pi-rwxxx
	DBNodeCrashList *string `json:"DBNodeCrashList,omitempty" xml:"DBNodeCrashList,omitempty"`
	// The method used to replicate data across zones. Valid values:
	//
	// 	- **AsyncSync**: the asynchronous mode.
	//
	// 	- **SemiSync**: the semi-synchronous mode.
	//
	// example:
	//
	// AsynSync
	DataSyncMode *string `json:"DataSyncMode,omitempty" xml:"DataSyncMode,omitempty"`
	// The fault injection method. Valid values:
	//
	// 	- 0: `Crash SQL`-based fault injection.
	//
	// Enumerated values:
	//
	// 	- CrashSQLInjection: CrashSQLInjection.
	//
	// example:
	//
	// 0
	FaultInjectionType *string `json:"FaultInjectionType,omitempty" xml:"FaultInjectionType,omitempty"`
	// The level of the disaster recovery drill. Valid values:
	//
	// 	- `0` or `FaultInjection`: The primary zone level.
	//
	// 	- `1`: The node level.
	//
	// >
	//
	// 	- In **primary zone-level disaster recovery drill*	- scenarios, all compute nodes in the primary zone are unavailable. Data loss occurs during failovers in the scenarios.
	//
	// 	- In **node-level disaster recovery drill*	- scenarios, you can specify only one compute node for the disaster recovery drill. You can use the `DBNodeCrashList` parameter to specify the name of the compute node that you want to use for the drill.
	//
	// Enumerated values:
	//
	// 	- FaultInjectToPrimaryAz
	//
	// 	- FaultInjectToDbNode
	//
	// 	- FaultInjection
	//
	// 	- 0
	//
	// 	- 1
	//
	// example:
	//
	// 0
	FaultSimulateMode *string `json:"FaultSimulateMode,omitempty" xml:"FaultSimulateMode,omitempty"`
	// Specifies whether to enable automatic IMCI-based query acceleration. IMCI is short for In-Memory Column Index. Valid values:
	//
	// 	- `ON`: enables automatic IMCI-based query acceleration.
	//
	// 	- `OFF`: disables automatic IMCI-based query acceleration.
	//
	// >
	//
	// 	- This parameter is supported only for PolarDB for MySQL clusters.
	//
	// 	- For information about the cluster version limits, see [Automatic IMCI-based query acceleration](https://help.aliyun.com/document_detail/2854119.html).
	//
	// example:
	//
	// OFF
	ImciAutoIndex        *string `json:"ImciAutoIndex,omitempty" xml:"ImciAutoIndex,omitempty"`
	ModifyRowCompression *string `json:"ModifyRowCompression,omitempty" xml:"ModifyRowCompression,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to enable cross-zone automatic switchover. Valid values:
	//
	// 	- **ON**: enables cross-zone automatic switchover.
	//
	// 	- **OFF**: disables cross-zone automatic switchover.
	//
	// example:
	//
	// ON
	StandbyHAMode *string `json:"StandbyHAMode,omitempty" xml:"StandbyHAMode,omitempty"`
	// Specifies whether to enable automatic storage scaling for the Standard Edition cluster. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// example:
	//
	// Enable
	StorageAutoScale *string `json:"StorageAutoScale,omitempty" xml:"StorageAutoScale,omitempty"`
	// The maximum storage capacity of the cluster of Standard Edition in automatic scaling. Unit: GB.
	//
	// >  The maximum value of this parameter is 32000.
	//
	// example:
	//
	// 800
	StorageUpperBound *int64  `json:"StorageUpperBound,omitempty" xml:"StorageUpperBound,omitempty"`
	TableMeta         *string `json:"TableMeta,omitempty" xml:"TableMeta,omitempty"`
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
