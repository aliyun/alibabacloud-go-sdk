// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceLinkedRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *CreateServiceLinkedRoleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateServiceLinkedRoleRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateServiceLinkedRoleRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateServiceLinkedRoleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateServiceLinkedRoleRequest
	GetResourceOwnerId() *int64
	SetServiceName(v string) *CreateServiceLinkedRoleRequest
	GetServiceName() *string
}

type CreateServiceLinkedRoleRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// ads.aliyuncs.com
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s CreateServiceLinkedRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
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

func (s *CreateServiceLinkedRoleRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *CreateServiceLinkedRoleRequest) SetOwnerAccount(v string) *CreateServiceLinkedRoleRequest {
	s.OwnerAccount = &v
	return s
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

func (s *CreateServiceLinkedRoleRequest) SetServiceName(v string) *CreateServiceLinkedRoleRequest {
	s.ServiceName = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) Validate() error {
	return dara.Validate(s)
}
