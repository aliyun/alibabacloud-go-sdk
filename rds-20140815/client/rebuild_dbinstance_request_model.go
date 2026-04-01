// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebuildDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *RebuildDBInstanceRequest
	GetDBInstanceId() *string
	SetDedicatedHostGroupId(v string) *RebuildDBInstanceRequest
	GetDedicatedHostGroupId() *string
	SetDedicatedHostId(v string) *RebuildDBInstanceRequest
	GetDedicatedHostId() *string
	SetOwnerId(v int64) *RebuildDBInstanceRequest
	GetOwnerId() *int64
	SetRebuildNodeType(v string) *RebuildDBInstanceRequest
	GetRebuildNodeType() *string
	SetRegionId(v string) *RebuildDBInstanceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RebuildDBInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RebuildDBInstanceRequest
	GetResourceOwnerId() *int64
}

type RebuildDBInstanceRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The dedicated cluster ID. You can call the DescribeDedicatedHostGroups operation to query the dedicated cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dhg-4nxxxxxxx
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	// The ID of the host on which the system rebuilds the secondary instance.
	//
	// >  If you do not specify this parameter, the system preferentially rebuilds the secondary instance on the original host on which the secondary instance resides. If the remaining storage of the original host is insufficient, the system rebuilds the secondary instance on a host on which the primary instance does not reside. If no suitable hosts are found, the system reports an error that indicates insufficient storage.
	//
	// example:
	//
	// i-bpxxxxxxx
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The role of the secondary instance that you want to rebuild. Valid values:
	//
	// 	- **FOLLOWER**: secondary instance
	//
	// 	- **LOG**: logger instance
	//
	// example:
	//
	// FOLLOWER
	RebuildNodeType *string `json:"RebuildNodeType,omitempty" xml:"RebuildNodeType,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RebuildDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RebuildDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *RebuildDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *RebuildDBInstanceRequest) GetDedicatedHostGroupId() *string {
	return s.DedicatedHostGroupId
}

func (s *RebuildDBInstanceRequest) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *RebuildDBInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RebuildDBInstanceRequest) GetRebuildNodeType() *string {
	return s.RebuildNodeType
}

func (s *RebuildDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RebuildDBInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RebuildDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RebuildDBInstanceRequest) SetDBInstanceId(v string) *RebuildDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *RebuildDBInstanceRequest) SetDedicatedHostGroupId(v string) *RebuildDBInstanceRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *RebuildDBInstanceRequest) SetDedicatedHostId(v string) *RebuildDBInstanceRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *RebuildDBInstanceRequest) SetOwnerId(v int64) *RebuildDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RebuildDBInstanceRequest) SetRebuildNodeType(v string) *RebuildDBInstanceRequest {
	s.RebuildNodeType = &v
	return s
}

func (s *RebuildDBInstanceRequest) SetRegionId(v string) *RebuildDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RebuildDBInstanceRequest) SetResourceOwnerAccount(v string) *RebuildDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RebuildDBInstanceRequest) SetResourceOwnerId(v int64) *RebuildDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RebuildDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
