// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeDDoSEventsRequest
	GetEndTime() *int64
	SetInstanceIds(v []*string) *DescribeDDoSEventsRequest
	GetInstanceIds() []*string
	SetPageNumber(v int32) *DescribeDDoSEventsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDDoSEventsRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeDDoSEventsRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribeDDoSEventsRequest
	GetStartTime() *int64
}

type DescribeDDoSEventsRequest struct {
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// > This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// example:
	//
	// 1583683200
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the instance.
	//
	// > You can call the [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) operation to query the IDs of all instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The number of the page to return. For example, to query the returned results on the first page, set the value to **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
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
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// > This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1582992000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDDoSEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDoSEventsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDDoSEventsRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeDDoSEventsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDDoSEventsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDDoSEventsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDDoSEventsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDDoSEventsRequest) SetEndTime(v int64) *DescribeDDoSEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetInstanceIds(v []*string) *DescribeDDoSEventsRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeDDoSEventsRequest) SetPageNumber(v int32) *DescribeDDoSEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetPageSize(v int32) *DescribeDDoSEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetResourceGroupId(v string) *DescribeDDoSEventsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetStartTime(v int64) *DescribeDDoSEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDDoSEventsRequest) Validate() error {
	return dara.Validate(s)
}
