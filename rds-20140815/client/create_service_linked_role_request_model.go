// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceLinkedRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *CreateServiceLinkedRoleRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateServiceLinkedRoleRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateServiceLinkedRoleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateServiceLinkedRoleRequest
	GetResourceOwnerId() *int64
	SetServiceLinkedRole(v string) *CreateServiceLinkedRoleRequest
	GetServiceLinkedRole() *string
}

type CreateServiceLinkedRoleRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call DescribeRegions to query the available regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The service-linked role. Valid values:
	//
	// - **AliyunServiceRoleForRds**: the service-linked role associate with ApsaraDB RDS for MySQL.
	//
	// - **AliyunServiceRoleForRdsPgsqlOnEcs**: the service-linked role associate with ApsaraDB RDS for PostgreSQL.
	//
	// - **AliyunServiceRoleForRDSProxyOnEcs**: the service-linked role associate with the database proxy of ApsaraDB RDS for PostgreSQL.
	//
	// This parameter is required.
	//
	// example:
	//
	// AliyunServiceRoleForRdsPgsqlOnEcs
	ServiceLinkedRole *string `json:"ServiceLinkedRole,omitempty" xml:"ServiceLinkedRole,omitempty"`
}

func (s CreateServiceLinkedRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateServiceLinkedRoleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateServiceLinkedRoleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateServiceLinkedRoleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateServiceLinkedRoleRequest) GetServiceLinkedRole() *string {
	return s.ServiceLinkedRole
}

func (s *CreateServiceLinkedRoleRequest) SetOwnerId(v int64) *CreateServiceLinkedRoleRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) SetRegionId(v string) *CreateServiceLinkedRoleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) SetResourceOwnerAccount(v string) *CreateServiceLinkedRoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) SetResourceOwnerId(v int64) *CreateServiceLinkedRoleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) SetServiceLinkedRole(v string) *CreateServiceLinkedRoleRequest {
	s.ServiceLinkedRole = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) Validate() error {
	return dara.Validate(s)
}
