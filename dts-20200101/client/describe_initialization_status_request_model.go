// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInitializationStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DescribeInitializationStatusRequest
	GetAccountId() *string
	SetOwnerId(v string) *DescribeInitializationStatusRequest
	GetOwnerId() *string
	SetPageNum(v int32) *DescribeInitializationStatusRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeInitializationStatusRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeInitializationStatusRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeInitializationStatusRequest
	GetResourceGroupId() *string
	SetSynchronizationJobId(v string) *DescribeInitializationStatusRequest
	GetSynchronizationJobId() *string
}

type DescribeInitializationStatusRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	OwnerId   *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than **0*	- and does not exceed the maximum value of the Integer data type. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page. Valid values: **1*	- to **100**. Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the data synchronization instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the data synchronization instance. You can call the [DescribeSynchronizationJobs](https://help.aliyun.com/document_detail/49454.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtsi76118o3w92****
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
}

func (s DescribeInitializationStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInitializationStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeInitializationStatusRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeInitializationStatusRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeInitializationStatusRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeInitializationStatusRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInitializationStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInitializationStatusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInitializationStatusRequest) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *DescribeInitializationStatusRequest) SetAccountId(v string) *DescribeInitializationStatusRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeInitializationStatusRequest) SetOwnerId(v string) *DescribeInitializationStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInitializationStatusRequest) SetPageNum(v int32) *DescribeInitializationStatusRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeInitializationStatusRequest) SetPageSize(v int32) *DescribeInitializationStatusRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInitializationStatusRequest) SetRegionId(v string) *DescribeInitializationStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInitializationStatusRequest) SetResourceGroupId(v string) *DescribeInitializationStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInitializationStatusRequest) SetSynchronizationJobId(v string) *DescribeInitializationStatusRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *DescribeInitializationStatusRequest) Validate() error {
	return dara.Validate(s)
}
