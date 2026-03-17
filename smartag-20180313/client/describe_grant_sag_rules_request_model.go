// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGrantSagRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeGrantSagRulesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeGrantSagRulesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeGrantSagRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGrantSagRulesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeGrantSagRulesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeGrantSagRulesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeGrantSagRulesRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeGrantSagRulesRequest
	GetSmartAGId() *string
}

type DescribeGrantSagRulesRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the returned page.
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
	// The ID of the region where the SAG instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-hdg*************
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s DescribeGrantSagRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantSagRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGrantSagRulesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeGrantSagRulesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeGrantSagRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGrantSagRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGrantSagRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGrantSagRulesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeGrantSagRulesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeGrantSagRulesRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeGrantSagRulesRequest) SetOwnerAccount(v string) *DescribeGrantSagRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeGrantSagRulesRequest) SetOwnerId(v int64) *DescribeGrantSagRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeGrantSagRulesRequest) SetPageNumber(v int32) *DescribeGrantSagRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGrantSagRulesRequest) SetPageSize(v int32) *DescribeGrantSagRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGrantSagRulesRequest) SetRegionId(v string) *DescribeGrantSagRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGrantSagRulesRequest) SetResourceOwnerAccount(v string) *DescribeGrantSagRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeGrantSagRulesRequest) SetResourceOwnerId(v int64) *DescribeGrantSagRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeGrantSagRulesRequest) SetSmartAGId(v string) *DescribeGrantSagRulesRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeGrantSagRulesRequest) Validate() error {
	return dara.Validate(s)
}
