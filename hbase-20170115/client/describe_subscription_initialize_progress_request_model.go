// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubscriptionInitializeProgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeSubscriptionInitializeProgressRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeSubscriptionInitializeProgressRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSubscriptionInitializeProgressRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSubscriptionInitializeProgressRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSubscriptionInitializeProgressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSubscriptionInitializeProgressRequest
	GetResourceOwnerId() *int64
	SetSubscriptionId(v string) *DescribeSubscriptionInitializeProgressRequest
	GetSubscriptionId() *string
}

type DescribeSubscriptionInitializeProgressRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	SubscriptionId *string `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
}

func (s DescribeSubscriptionInitializeProgressRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionInitializeProgressRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInitializeProgressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSubscriptionInitializeProgressRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSubscriptionInitializeProgressRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSubscriptionInitializeProgressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSubscriptionInitializeProgressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSubscriptionInitializeProgressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSubscriptionInitializeProgressRequest) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *DescribeSubscriptionInitializeProgressRequest) SetOwnerId(v int64) *DescribeSubscriptionInitializeProgressRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressRequest) SetPageNumber(v int32) *DescribeSubscriptionInitializeProgressRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressRequest) SetPageSize(v int32) *DescribeSubscriptionInitializeProgressRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressRequest) SetRegionId(v string) *DescribeSubscriptionInitializeProgressRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressRequest) SetResourceOwnerAccount(v string) *DescribeSubscriptionInitializeProgressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressRequest) SetResourceOwnerId(v int64) *DescribeSubscriptionInitializeProgressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressRequest) SetSubscriptionId(v string) *DescribeSubscriptionInitializeProgressRequest {
	s.SubscriptionId = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressRequest) Validate() error {
	return dara.Validate(s)
}
