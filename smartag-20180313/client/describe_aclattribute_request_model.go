// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeACLAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *DescribeACLAttributeRequest
	GetAclId() *string
	SetDirection(v string) *DescribeACLAttributeRequest
	GetDirection() *string
	SetName(v string) *DescribeACLAttributeRequest
	GetName() *string
	SetOrder(v string) *DescribeACLAttributeRequest
	GetOrder() *string
	SetOwnerAccount(v string) *DescribeACLAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeACLAttributeRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeACLAttributeRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeACLAttributeRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeACLAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeACLAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeACLAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribeACLAttributeRequest struct {
	// The ID of the ACL.
	//
	// This parameter is required.
	//
	// example:
	//
	// acl-ohlexqptfhyaq****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The direction of traffic in which the ACL rule is applied. Valid values:
	//
	// 	- **in**: The ACL rule controls inbound network traffic of the on-premises network that is associated with the Smart Access Gateway (SAG) instance.
	//
	// 	- **out**: The ACL rule controls outbound network traffic of the on-premises network that is associated with the SAG instance.
	//
	// example:
	//
	// out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The name of the ACL.
	//
	// The name must be 2 to 100 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// >  This parameter supports fuzzy match.
	//
	// example:
	//
	// doctest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 1255444444
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: **1*	- to **50**.
	//
	// Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the ACL is deployed.
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

func (s DescribeACLAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeACLAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeACLAttributeRequest) GetAclId() *string {
	return s.AclId
}

func (s *DescribeACLAttributeRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeACLAttributeRequest) GetName() *string {
	return s.Name
}

func (s *DescribeACLAttributeRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeACLAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeACLAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeACLAttributeRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeACLAttributeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeACLAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeACLAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeACLAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeACLAttributeRequest) SetAclId(v string) *DescribeACLAttributeRequest {
	s.AclId = &v
	return s
}

func (s *DescribeACLAttributeRequest) SetDirection(v string) *DescribeACLAttributeRequest {
	s.Direction = &v
	return s
}

func (s *DescribeACLAttributeRequest) SetName(v string) *DescribeACLAttributeRequest {
	s.Name = &v
	return s
}

func (s *DescribeACLAttributeRequest) SetOrder(v string) *DescribeACLAttributeRequest {
	s.Order = &v
	return s
}

func (s *DescribeACLAttributeRequest) SetOwnerAccount(v string) *DescribeACLAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeACLAttributeRequest) SetOwnerId(v int64) *DescribeACLAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeACLAttributeRequest) SetPageNumber(v int32) *DescribeACLAttributeRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeACLAttributeRequest) SetPageSize(v int32) *DescribeACLAttributeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeACLAttributeRequest) SetRegionId(v string) *DescribeACLAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeACLAttributeRequest) SetResourceOwnerAccount(v string) *DescribeACLAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeACLAttributeRequest) SetResourceOwnerId(v int64) *DescribeACLAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeACLAttributeRequest) Validate() error {
	return dara.Validate(s)
}
