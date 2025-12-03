// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessLogsDownloadAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancerId(v string) *DescribeAccessLogsDownloadAttributeRequest
	GetLoadBalancerId() *string
	SetLogType(v string) *DescribeAccessLogsDownloadAttributeRequest
	GetLogType() *string
	SetOwnerAccount(v string) *DescribeAccessLogsDownloadAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAccessLogsDownloadAttributeRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeAccessLogsDownloadAttributeRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAccessLogsDownloadAttributeRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeAccessLogsDownloadAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAccessLogsDownloadAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAccessLogsDownloadAttributeRequest
	GetResourceOwnerId() *int64
	SetTags(v string) *DescribeAccessLogsDownloadAttributeRequest
	GetTags() *string
}

type DescribeAccessLogsDownloadAttributeRequest struct {
	// The CLB instance ID.
	//
	// example:
	//
	// lb-bp1b6c719dfa08ex*****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The type of access log. Set the value to **layer7**, which specifies Layer 7 access logs.
	//
	// example:
	//
	// layer7
	LogType      *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the CLB instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags that are added to the CLB instance. The tags must be key-value pairs that are contained in a JSON dictionary.
	//
	// You can specify up to 10 tags in each call.
	//
	// example:
	//
	// [{"tagKey":"Key1","tagValue":"Value1"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s DescribeAccessLogsDownloadAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessLogsDownloadAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessLogsDownloadAttributeRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeAccessLogsDownloadAttributeRequest) GetLogType() *string {
	return s.LogType
}

func (s *DescribeAccessLogsDownloadAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAccessLogsDownloadAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAccessLogsDownloadAttributeRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAccessLogsDownloadAttributeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAccessLogsDownloadAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAccessLogsDownloadAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAccessLogsDownloadAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAccessLogsDownloadAttributeRequest) GetTags() *string {
	return s.Tags
}

func (s *DescribeAccessLogsDownloadAttributeRequest) SetLoadBalancerId(v string) *DescribeAccessLogsDownloadAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeRequest) SetLogType(v string) *DescribeAccessLogsDownloadAttributeRequest {
	s.LogType = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeRequest) SetOwnerAccount(v string) *DescribeAccessLogsDownloadAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeRequest) SetOwnerId(v int64) *DescribeAccessLogsDownloadAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeRequest) SetPageNumber(v int32) *DescribeAccessLogsDownloadAttributeRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeRequest) SetPageSize(v int32) *DescribeAccessLogsDownloadAttributeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeRequest) SetRegionId(v string) *DescribeAccessLogsDownloadAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeRequest) SetResourceOwnerAccount(v string) *DescribeAccessLogsDownloadAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeRequest) SetResourceOwnerId(v int64) *DescribeAccessLogsDownloadAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeRequest) SetTags(v string) *DescribeAccessLogsDownloadAttributeRequest {
	s.Tags = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeRequest) Validate() error {
	return dara.Validate(s)
}
