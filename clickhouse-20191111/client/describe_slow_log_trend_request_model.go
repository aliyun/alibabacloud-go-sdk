// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeSlowLogTrendRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeSlowLogTrendRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *DescribeSlowLogTrendRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSlowLogTrendRequest
	GetOwnerId() *int64
	SetQueryDurationMs(v int32) *DescribeSlowLogTrendRequest
	GetQueryDurationMs() *int32
	SetRegionId(v string) *DescribeSlowLogTrendRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSlowLogTrendRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSlowLogTrendRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeSlowLogTrendRequest
	GetStartTime() *string
}

type DescribeSlowLogTrendRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cc-bp11xxl475ui8****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2022-05-27 16:00:00
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The execution duration of slow SQL queries. Minimum value: **1000**. Default value: **1000**. Unit: milliseconds. Slow SQL queries that last longer than the specified duration are returned in response parameters.
	//
	// example:
	//
	// 1000
	QueryDurationMs *int32 `json:"QueryDurationMs,omitempty" xml:"QueryDurationMs,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2022-05-20 16:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSlowLogTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSlowLogTrendRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSlowLogTrendRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSlowLogTrendRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSlowLogTrendRequest) GetQueryDurationMs() *int32 {
	return s.QueryDurationMs
}

func (s *DescribeSlowLogTrendRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSlowLogTrendRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSlowLogTrendRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSlowLogTrendRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSlowLogTrendRequest) SetDBClusterId(v string) *DescribeSlowLogTrendRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetEndTime(v string) *DescribeSlowLogTrendRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetOwnerAccount(v string) *DescribeSlowLogTrendRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetOwnerId(v int64) *DescribeSlowLogTrendRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetQueryDurationMs(v int32) *DescribeSlowLogTrendRequest {
	s.QueryDurationMs = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetRegionId(v string) *DescribeSlowLogTrendRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetResourceOwnerAccount(v string) *DescribeSlowLogTrendRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetResourceOwnerId(v int64) *DescribeSlowLogTrendRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetStartTime(v string) *DescribeSlowLogTrendRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) Validate() error {
	return dara.Validate(s)
}
