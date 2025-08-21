// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSystemLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeSystemLogRequest
	GetEndTime() *int64
	SetEntityObject(v string) *DescribeSystemLogRequest
	GetEntityObject() *string
	SetEntityType(v int32) *DescribeSystemLogRequest
	GetEntityType() *int32
	SetPageNumber(v int32) *DescribeSystemLogRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSystemLogRequest
	GetPageSize() *int32
	SetStartTime(v int64) *DescribeSystemLogRequest
	GetStartTime() *int64
}

type DescribeSystemLogRequest struct {
	// The end of the time range to query. The bills of the burstable clean bandwidth that are issued before this point in time are queried. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1640966400000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IP address of the instance.
	//
	// > You can call the [DescribeInstanceDetails](https://help.aliyun.com/document_detail/91490.html) operation to query the IP addresses of all instances.
	//
	// example:
	//
	// 203.107.XX.XX
	EntityObject *string `json:"EntityObject,omitempty" xml:"EntityObject,omitempty"`
	// The type of the system log. Set the value to **20**, which indicates the billing logs for the burstable clean bandwidth.
	//
	// > You must specify this parameter. Otherwise, the call fails.
	//
	// example:
	//
	// 20
	EntityType *int32 `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The number of the page to return.
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
	// The beginning of the time range to query. The bills of the burstable clean bandwidth that are issued after this point in time are queried. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1609430400000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSystemLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeSystemLogRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSystemLogRequest) GetEntityObject() *string {
	return s.EntityObject
}

func (s *DescribeSystemLogRequest) GetEntityType() *int32 {
	return s.EntityType
}

func (s *DescribeSystemLogRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSystemLogRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSystemLogRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSystemLogRequest) SetEndTime(v int64) *DescribeSystemLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSystemLogRequest) SetEntityObject(v string) *DescribeSystemLogRequest {
	s.EntityObject = &v
	return s
}

func (s *DescribeSystemLogRequest) SetEntityType(v int32) *DescribeSystemLogRequest {
	s.EntityType = &v
	return s
}

func (s *DescribeSystemLogRequest) SetPageNumber(v int32) *DescribeSystemLogRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSystemLogRequest) SetPageSize(v int32) *DescribeSystemLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSystemLogRequest) SetStartTime(v int64) *DescribeSystemLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSystemLogRequest) Validate() error {
	return dara.Validate(s)
}
