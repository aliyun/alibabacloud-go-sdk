// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterOperateLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DescribeClusterOperateLogsRequest
	GetAccountId() *string
	SetClientToken(v string) *DescribeClusterOperateLogsRequest
	GetClientToken() *string
	SetDedicatedClusterId(v string) *DescribeClusterOperateLogsRequest
	GetDedicatedClusterId() *string
	SetDtsJobId(v string) *DescribeClusterOperateLogsRequest
	GetDtsJobId() *string
	SetEndTime(v int64) *DescribeClusterOperateLogsRequest
	GetEndTime() *int64
	SetOwnerID(v string) *DescribeClusterOperateLogsRequest
	GetOwnerID() *string
	SetPageNumber(v int32) *DescribeClusterOperateLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeClusterOperateLogsRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeClusterOperateLogsRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribeClusterOperateLogsRequest
	GetStartTime() *int64
}

type DescribeClusterOperateLogsRequest struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Ensures the idempotency of the request. Generate a parameter value from your client to ensure that the value is unique across different requests. **ClientToken*	- supports only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the DTS dedicated cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtsxxxxx
	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" xml:"DedicatedClusterId,omitempty"`
	// The ID of the data migration or data synchronization task.
	//
	// example:
	//
	// k2gm967v16f****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The end timestamp. Unit: milliseconds (ms).
	//
	// example:
	//
	// 1650866995000
	EndTime *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerID *string `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	// The page number. The value must be a positive integer that does not exceed the maximum value of the Integer data type. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of log entries per page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The start timestamp. If this parameter is not specified, data from the last seven days is returned by default. Unit: milliseconds (ms).
	//
	// example:
	//
	// 1650866955000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeClusterOperateLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterOperateLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterOperateLogsRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeClusterOperateLogsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeClusterOperateLogsRequest) GetDedicatedClusterId() *string {
	return s.DedicatedClusterId
}

func (s *DescribeClusterOperateLogsRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeClusterOperateLogsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeClusterOperateLogsRequest) GetOwnerID() *string {
	return s.OwnerID
}

func (s *DescribeClusterOperateLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeClusterOperateLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeClusterOperateLogsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeClusterOperateLogsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeClusterOperateLogsRequest) SetAccountId(v string) *DescribeClusterOperateLogsRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeClusterOperateLogsRequest) SetClientToken(v string) *DescribeClusterOperateLogsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeClusterOperateLogsRequest) SetDedicatedClusterId(v string) *DescribeClusterOperateLogsRequest {
	s.DedicatedClusterId = &v
	return s
}

func (s *DescribeClusterOperateLogsRequest) SetDtsJobId(v string) *DescribeClusterOperateLogsRequest {
	s.DtsJobId = &v
	return s
}

func (s *DescribeClusterOperateLogsRequest) SetEndTime(v int64) *DescribeClusterOperateLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeClusterOperateLogsRequest) SetOwnerID(v string) *DescribeClusterOperateLogsRequest {
	s.OwnerID = &v
	return s
}

func (s *DescribeClusterOperateLogsRequest) SetPageNumber(v int32) *DescribeClusterOperateLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeClusterOperateLogsRequest) SetPageSize(v int32) *DescribeClusterOperateLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeClusterOperateLogsRequest) SetResourceGroupId(v string) *DescribeClusterOperateLogsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeClusterOperateLogsRequest) SetStartTime(v int64) *DescribeClusterOperateLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeClusterOperateLogsRequest) Validate() error {
	return dara.Validate(s)
}
