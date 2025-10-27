// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckServiceLinkedRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *CheckServiceLinkedRoleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CheckServiceLinkedRoleRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CheckServiceLinkedRoleRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CheckServiceLinkedRoleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckServiceLinkedRoleRequest
	GetResourceOwnerId() *int64
	SetServiceName(v string) *CheckServiceLinkedRoleRequest
	GetServiceName() *string
}

type CheckServiceLinkedRoleRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
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

func (s CheckServiceLinkedRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CheckServiceLinkedRoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckServiceLinkedRoleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckServiceLinkedRoleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckServiceLinkedRoleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckServiceLinkedRoleRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *CheckServiceLinkedRoleRequest) SetOwnerAccount(v string) *CheckServiceLinkedRoleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckServiceLinkedRoleRequest) SetOwnerId(v int64) *CheckServiceLinkedRoleRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckServiceLinkedRoleRequest) SetRegionId(v string) *CheckServiceLinkedRoleRequest {
	s.RegionId = &v
	return s
}

func (s *CheckServiceLinkedRoleRequest) SetResourceOwnerAccount(v string) *CheckServiceLinkedRoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckServiceLinkedRoleRequest) SetResourceOwnerId(v int64) *CheckServiceLinkedRoleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckServiceLinkedRoleRequest) SetServiceName(v string) *CheckServiceLinkedRoleRequest {
	s.ServiceName = &v
	return s
}

func (s *CheckServiceLinkedRoleRequest) Validate() error {
	return dara.Validate(s)
}
