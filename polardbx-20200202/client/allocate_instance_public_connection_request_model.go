// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateInstancePublicConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionStringPrefix(v string) *AllocateInstancePublicConnectionRequest
	GetConnectionStringPrefix() *string
	SetDBInstanceName(v string) *AllocateInstancePublicConnectionRequest
	GetDBInstanceName() *string
	SetOwnerAccount(v string) *AllocateInstancePublicConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AllocateInstancePublicConnectionRequest
	GetOwnerId() *int64
	SetPort(v string) *AllocateInstancePublicConnectionRequest
	GetPort() *string
	SetRegionId(v string) *AllocateInstancePublicConnectionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AllocateInstancePublicConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AllocateInstancePublicConnectionRequest
	GetResourceOwnerId() *int64
}

type AllocateInstancePublicConnectionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-**************
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-**************
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AllocateInstancePublicConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateInstancePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocateInstancePublicConnectionRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *AllocateInstancePublicConnectionRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *AllocateInstancePublicConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AllocateInstancePublicConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AllocateInstancePublicConnectionRequest) GetPort() *string {
	return s.Port
}

func (s *AllocateInstancePublicConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AllocateInstancePublicConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AllocateInstancePublicConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AllocateInstancePublicConnectionRequest) SetConnectionStringPrefix(v string) *AllocateInstancePublicConnectionRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetDBInstanceName(v string) *AllocateInstancePublicConnectionRequest {
	s.DBInstanceName = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetOwnerAccount(v string) *AllocateInstancePublicConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetOwnerId(v int64) *AllocateInstancePublicConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetPort(v string) *AllocateInstancePublicConnectionRequest {
	s.Port = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetRegionId(v string) *AllocateInstancePublicConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetResourceOwnerAccount(v string) *AllocateInstancePublicConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetResourceOwnerId(v int64) *AllocateInstancePublicConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) Validate() error {
	return dara.Validate(s)
}
