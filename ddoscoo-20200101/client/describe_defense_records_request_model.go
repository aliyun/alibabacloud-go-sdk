// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeDefenseRecordsRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribeDefenseRecordsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *DescribeDefenseRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDefenseRecordsRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeDefenseRecordsRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribeDefenseRecordsRequest
	GetStartTime() *int64
}

type DescribeDefenseRecordsRequest struct {
	// The end of the time range to query. This value is a UNIX timestamp. Units: miliseconds.
	//
	// > The time must be in the latest 90 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1583683200000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the instance.
	//
	// > You can call the [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) operation to query the IDs of all instances.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return. For example, to query the returned results on the first page, set the value to **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: **50**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The beginning of the time range to query. This value is a UNIX timestamp. Units: miliseconds.
	//
	// > The time must be in the latest 90 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1582992000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDefenseRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRecordsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDefenseRecordsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDefenseRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDefenseRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDefenseRecordsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDefenseRecordsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDefenseRecordsRequest) SetEndTime(v int64) *DescribeDefenseRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDefenseRecordsRequest) SetInstanceId(v string) *DescribeDefenseRecordsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseRecordsRequest) SetPageNumber(v int32) *DescribeDefenseRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDefenseRecordsRequest) SetPageSize(v int32) *DescribeDefenseRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDefenseRecordsRequest) SetResourceGroupId(v string) *DescribeDefenseRecordsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDefenseRecordsRequest) SetStartTime(v int64) *DescribeDefenseRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDefenseRecordsRequest) Validate() error {
	return dara.Validate(s)
}
