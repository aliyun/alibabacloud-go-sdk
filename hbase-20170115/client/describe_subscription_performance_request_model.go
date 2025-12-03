// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubscriptionPerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeSubscriptionPerformanceRequest
	GetEndTime() *string
	SetKey(v string) *DescribeSubscriptionPerformanceRequest
	GetKey() *string
	SetOwnerId(v int64) *DescribeSubscriptionPerformanceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSubscriptionPerformanceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSubscriptionPerformanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSubscriptionPerformanceRequest
	GetResourceOwnerId() *int64
	SetSourceInstanceId(v string) *DescribeSubscriptionPerformanceRequest
	GetSourceInstanceId() *string
	SetStartTime(v string) *DescribeSubscriptionPerformanceRequest
	GetStartTime() *string
	SetSubscriptionId(v string) *DescribeSubscriptionPerformanceRequest
	GetSubscriptionId() *string
}

type DescribeSubscriptionPerformanceRequest struct {
	// This parameter is required.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	Key                  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	// This parameter is required.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	SubscriptionId *string `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
}

func (s DescribeSubscriptionPerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionPerformanceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSubscriptionPerformanceRequest) GetKey() *string {
	return s.Key
}

func (s *DescribeSubscriptionPerformanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSubscriptionPerformanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSubscriptionPerformanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSubscriptionPerformanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSubscriptionPerformanceRequest) GetSourceInstanceId() *string {
	return s.SourceInstanceId
}

func (s *DescribeSubscriptionPerformanceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSubscriptionPerformanceRequest) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *DescribeSubscriptionPerformanceRequest) SetEndTime(v string) *DescribeSubscriptionPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSubscriptionPerformanceRequest) SetKey(v string) *DescribeSubscriptionPerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeSubscriptionPerformanceRequest) SetOwnerId(v int64) *DescribeSubscriptionPerformanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSubscriptionPerformanceRequest) SetRegionId(v string) *DescribeSubscriptionPerformanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSubscriptionPerformanceRequest) SetResourceOwnerAccount(v string) *DescribeSubscriptionPerformanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSubscriptionPerformanceRequest) SetResourceOwnerId(v int64) *DescribeSubscriptionPerformanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSubscriptionPerformanceRequest) SetSourceInstanceId(v string) *DescribeSubscriptionPerformanceRequest {
	s.SourceInstanceId = &v
	return s
}

func (s *DescribeSubscriptionPerformanceRequest) SetStartTime(v string) *DescribeSubscriptionPerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSubscriptionPerformanceRequest) SetSubscriptionId(v string) *DescribeSubscriptionPerformanceRequest {
	s.SubscriptionId = &v
	return s
}

func (s *DescribeSubscriptionPerformanceRequest) Validate() error {
	return dara.Validate(s)
}
