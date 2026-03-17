// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBindableSmartAccessGatewaysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCcnId(v string) *DescribeBindableSmartAccessGatewaysRequest
	GetCcnId() *string
	SetCrossAccount(v bool) *DescribeBindableSmartAccessGatewaysRequest
	GetCrossAccount() *bool
	SetName(v string) *DescribeBindableSmartAccessGatewaysRequest
	GetName() *string
	SetOwnerAccount(v string) *DescribeBindableSmartAccessGatewaysRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeBindableSmartAccessGatewaysRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeBindableSmartAccessGatewaysRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBindableSmartAccessGatewaysRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeBindableSmartAccessGatewaysRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeBindableSmartAccessGatewaysRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeBindableSmartAccessGatewaysRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeBindableSmartAccessGatewaysRequest
	GetSmartAGId() *string
}

type DescribeBindableSmartAccessGatewaysRequest struct {
	// The ID of the CCN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccn-fu75a6m4clv7ai****
	CcnId *string `json:"CcnId,omitempty" xml:"CcnId,omitempty"`
	// Specifies whether to query only the SAG instances that belong to another Alibaba Cloud account. Valid values:
	//
	// 	- **false*	- (default): no
	//
	// 	- **true**: yes
	//
	// example:
	//
	// false
	CrossAccount *bool `json:"CrossAccount,omitempty" xml:"CrossAccount,omitempty"`
	// The name of the SAG instance.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// sdggd111
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**. Maximum value: **50**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
	//
	// example:
	//
	// sag-00uc4vgxch1zsu****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s DescribeBindableSmartAccessGatewaysRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBindableSmartAccessGatewaysRequest) GoString() string {
	return s.String()
}

func (s *DescribeBindableSmartAccessGatewaysRequest) GetCcnId() *string {
	return s.CcnId
}

func (s *DescribeBindableSmartAccessGatewaysRequest) GetCrossAccount() *bool {
	return s.CrossAccount
}

func (s *DescribeBindableSmartAccessGatewaysRequest) GetName() *string {
	return s.Name
}

func (s *DescribeBindableSmartAccessGatewaysRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeBindableSmartAccessGatewaysRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeBindableSmartAccessGatewaysRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBindableSmartAccessGatewaysRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBindableSmartAccessGatewaysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBindableSmartAccessGatewaysRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeBindableSmartAccessGatewaysRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeBindableSmartAccessGatewaysRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeBindableSmartAccessGatewaysRequest) SetCcnId(v string) *DescribeBindableSmartAccessGatewaysRequest {
	s.CcnId = &v
	return s
}

func (s *DescribeBindableSmartAccessGatewaysRequest) SetCrossAccount(v bool) *DescribeBindableSmartAccessGatewaysRequest {
	s.CrossAccount = &v
	return s
}

func (s *DescribeBindableSmartAccessGatewaysRequest) SetName(v string) *DescribeBindableSmartAccessGatewaysRequest {
	s.Name = &v
	return s
}

func (s *DescribeBindableSmartAccessGatewaysRequest) SetOwnerAccount(v string) *DescribeBindableSmartAccessGatewaysRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBindableSmartAccessGatewaysRequest) SetOwnerId(v int64) *DescribeBindableSmartAccessGatewaysRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBindableSmartAccessGatewaysRequest) SetPageNumber(v int32) *DescribeBindableSmartAccessGatewaysRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBindableSmartAccessGatewaysRequest) SetPageSize(v int32) *DescribeBindableSmartAccessGatewaysRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBindableSmartAccessGatewaysRequest) SetRegionId(v string) *DescribeBindableSmartAccessGatewaysRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBindableSmartAccessGatewaysRequest) SetResourceOwnerAccount(v string) *DescribeBindableSmartAccessGatewaysRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBindableSmartAccessGatewaysRequest) SetResourceOwnerId(v int64) *DescribeBindableSmartAccessGatewaysRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBindableSmartAccessGatewaysRequest) SetSmartAGId(v string) *DescribeBindableSmartAccessGatewaysRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeBindableSmartAccessGatewaysRequest) Validate() error {
	return dara.Validate(s)
}
