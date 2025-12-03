// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubscriptionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeSubscriptionsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeSubscriptionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSubscriptionsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSubscriptionsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSubscriptionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSubscriptionsRequest
	GetResourceOwnerId() *int64
	SetSubscriptionId(v string) *DescribeSubscriptionsRequest
	GetSubscriptionId() *string
}

type DescribeSubscriptionsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SubscriptionId       *string `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
}

func (s DescribeSubscriptionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSubscriptionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSubscriptionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSubscriptionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSubscriptionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSubscriptionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSubscriptionsRequest) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *DescribeSubscriptionsRequest) SetOwnerId(v int64) *DescribeSubscriptionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSubscriptionsRequest) SetPageNumber(v int32) *DescribeSubscriptionsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSubscriptionsRequest) SetPageSize(v int32) *DescribeSubscriptionsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSubscriptionsRequest) SetRegionId(v string) *DescribeSubscriptionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSubscriptionsRequest) SetResourceOwnerAccount(v string) *DescribeSubscriptionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSubscriptionsRequest) SetResourceOwnerId(v int64) *DescribeSubscriptionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSubscriptionsRequest) SetSubscriptionId(v string) *DescribeSubscriptionsRequest {
	s.SubscriptionId = &v
	return s
}

func (s *DescribeSubscriptionsRequest) Validate() error {
	return dara.Validate(s)
}
