// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGrantRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssociatedCcnId(v string) *DescribeGrantRulesRequest
	GetAssociatedCcnId() *string
	SetOwnerAccount(v string) *DescribeGrantRulesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeGrantRulesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeGrantRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGrantRulesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeGrantRulesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeGrantRulesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeGrantRulesRequest
	GetResourceOwnerId() *int64
}

type DescribeGrantRulesRequest struct {
	// The ID of the CCN instance.
	//
	// example:
	//
	// ccn-n2935s1mnwv8i*****
	AssociatedCcnId *string `json:"AssociatedCcnId,omitempty" xml:"AssociatedCcnId,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: **10**. Maximum value: **50**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the CCN instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeGrantRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesRequest) GetAssociatedCcnId() *string {
	return s.AssociatedCcnId
}

func (s *DescribeGrantRulesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeGrantRulesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeGrantRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGrantRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGrantRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGrantRulesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeGrantRulesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeGrantRulesRequest) SetAssociatedCcnId(v string) *DescribeGrantRulesRequest {
	s.AssociatedCcnId = &v
	return s
}

func (s *DescribeGrantRulesRequest) SetOwnerAccount(v string) *DescribeGrantRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeGrantRulesRequest) SetOwnerId(v int64) *DescribeGrantRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeGrantRulesRequest) SetPageNumber(v int32) *DescribeGrantRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGrantRulesRequest) SetPageSize(v int32) *DescribeGrantRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGrantRulesRequest) SetRegionId(v string) *DescribeGrantRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGrantRulesRequest) SetResourceOwnerAccount(v string) *DescribeGrantRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeGrantRulesRequest) SetResourceOwnerId(v int64) *DescribeGrantRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeGrantRulesRequest) Validate() error {
	return dara.Validate(s)
}
