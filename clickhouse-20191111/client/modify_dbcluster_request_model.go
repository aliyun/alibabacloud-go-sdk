// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterClass(v string) *ModifyDBClusterRequest
	GetDBClusterClass() *string
	SetDBClusterId(v string) *ModifyDBClusterRequest
	GetDBClusterId() *string
	SetDBNodeGroupCount(v string) *ModifyDBClusterRequest
	GetDBNodeGroupCount() *string
	SetDBNodeStorage(v string) *ModifyDBClusterRequest
	GetDBNodeStorage() *string
	SetDbNodeStorageType(v string) *ModifyDBClusterRequest
	GetDbNodeStorageType() *string
	SetDisableWriteWindows(v string) *ModifyDBClusterRequest
	GetDisableWriteWindows() *string
	SetOwnerAccount(v string) *ModifyDBClusterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyDBClusterRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyDBClusterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterRequest
	GetResourceOwnerId() *int64
}

type ModifyDBClusterRequest struct {
	// The specifications of the cluster.
	//
	// 	- Valid values when the cluster is of Single-replica Edition:
	//
	//     	- **S8**
	//
	//     	- **S16**
	//
	//     	- **S32**
	//
	//     	- **S64**
	//
	//     	- **S104**
	//
	// 	- Valid values when the cluster is of Double-replica Edition:
	//
	//     	- **C8**
	//
	//     	- **C16**
	//
	//     	- **C32**
	//
	//     	- **C64**
	//
	//     	- **C104**
	//
	// This parameter is required.
	//
	// example:
	//
	// S4-NEW
	DBClusterClass *string `json:"DBClusterClass,omitempty" xml:"DBClusterClass,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-bp19lo45sy98x****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The number of nodes in the cluster.
	//
	// 	- If the cluster is of Single-replica Edition, the value must be an integer that ranges from 1 to 48.
	//
	// 	- If the cluster is of Double-replica Edition, the value must be an integer that ranges from 1 to 24.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DBNodeGroupCount *string `json:"DBNodeGroupCount,omitempty" xml:"DBNodeGroupCount,omitempty"`
	// The storage capacity of a single node of the cluster. Unit: GB.
	//
	// Valid values: 100 to 32000.
	//
	// >  This value is a multiple of 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	DBNodeStorage *string `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	// The storage type of the cluster. Valid values:
	//
	// 	- **CloudESSD**: The cluster uses an Enterprise SSD (ESSD) of performance level 1 (PL1).
	//
	// 	- **CloudESSD_PL2**: The cluster uses an ESSD of PL2.
	//
	// 	- **CloudESSD_PL3**: The cluster uses an ESSD of PL3.
	//
	// 	- **CloudEfficiency**: The cluster uses an ultra disk.
	//
	// 	- **CloudSSD**: The cluster uses a standard SSD.
	//
	// example:
	//
	// CloudESSD
	DbNodeStorageType *string `json:"DbNodeStorageType,omitempty" xml:"DbNodeStorageType,omitempty"`
	// The time window during which write operations are stopped. Separate the start time and end time with commas (,). Specify the time in the ISO 8601 standard.
	//
	// example:
	//
	// 2024-07-09T20:00:00+08:00,2024-07-09T21:00:00+08:00
	DisableWriteWindows *string `json:"DisableWriteWindows,omitempty" xml:"DisableWriteWindows,omitempty"`
	OwnerAccount        *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterRequest) GetDBClusterClass() *string {
	return s.DBClusterClass
}

func (s *ModifyDBClusterRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterRequest) GetDBNodeGroupCount() *string {
	return s.DBNodeGroupCount
}

func (s *ModifyDBClusterRequest) GetDBNodeStorage() *string {
	return s.DBNodeStorage
}

func (s *ModifyDBClusterRequest) GetDbNodeStorageType() *string {
	return s.DbNodeStorageType
}

func (s *ModifyDBClusterRequest) GetDisableWriteWindows() *string {
	return s.DisableWriteWindows
}

func (s *ModifyDBClusterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBClusterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterRequest) SetDBClusterClass(v string) *ModifyDBClusterRequest {
	s.DBClusterClass = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBClusterId(v string) *ModifyDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBNodeGroupCount(v string) *ModifyDBClusterRequest {
	s.DBNodeGroupCount = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBNodeStorage(v string) *ModifyDBClusterRequest {
	s.DBNodeStorage = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDbNodeStorageType(v string) *ModifyDBClusterRequest {
	s.DbNodeStorageType = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDisableWriteWindows(v string) *ModifyDBClusterRequest {
	s.DisableWriteWindows = &v
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

func (s *ModifyDBClusterRequest) SetRegionId(v string) *ModifyDBClusterRequest {
	s.RegionId = &v
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

func (s *ModifyDBClusterRequest) Validate() error {
	return dara.Validate(s)
}
