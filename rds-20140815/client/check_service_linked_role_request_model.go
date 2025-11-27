// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckServiceLinkedRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *CheckServiceLinkedRoleRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CheckServiceLinkedRoleRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CheckServiceLinkedRoleRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CheckServiceLinkedRoleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckServiceLinkedRoleRequest
	GetResourceOwnerId() *int64
	SetServiceLinkedRole(v string) *CheckServiceLinkedRoleRequest
	GetServiceLinkedRole() *string
}

type CheckServiceLinkedRoleRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// >  You can specify any region for this parameter, which does not affect your query results. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/26243.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The SLR name.
	//
	// >  For more information about the SLRs supported by ApsaraDB RDS, see [Service-linked roles](https://help.aliyun.com/document_detail/342840.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// AliyunServiceRoleForRdsPgsqlOnEcs
	ServiceLinkedRole *string `json:"ServiceLinkedRole,omitempty" xml:"ServiceLinkedRole,omitempty"`
}

func (s CheckServiceLinkedRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckServiceLinkedRoleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckServiceLinkedRoleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CheckServiceLinkedRoleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckServiceLinkedRoleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckServiceLinkedRoleRequest) GetServiceLinkedRole() *string {
	return s.ServiceLinkedRole
}

func (s *CheckServiceLinkedRoleRequest) SetOwnerId(v int64) *CheckServiceLinkedRoleRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckServiceLinkedRoleRequest) SetRegionId(v string) *CheckServiceLinkedRoleRequest {
	s.RegionId = &v
	return s
}

func (s *CheckServiceLinkedRoleRequest) SetResourceGroupId(v string) *CheckServiceLinkedRoleRequest {
	s.ResourceGroupId = &v
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

func (s *CheckServiceLinkedRoleRequest) SetServiceLinkedRole(v string) *CheckServiceLinkedRoleRequest {
	s.ServiceLinkedRole = &v
	return s
}

func (s *CheckServiceLinkedRoleRequest) Validate() error {
	return dara.Validate(s)
}
