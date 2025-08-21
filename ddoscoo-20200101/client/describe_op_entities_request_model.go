// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpEntitiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeOpEntitiesRequest
	GetEndTime() *int64
	SetEntityObject(v string) *DescribeOpEntitiesRequest
	GetEntityObject() *string
	SetEntityType(v int32) *DescribeOpEntitiesRequest
	GetEntityType() *int32
	SetPageNumber(v int32) *DescribeOpEntitiesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeOpEntitiesRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeOpEntitiesRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribeOpEntitiesRequest
	GetStartTime() *int64
}

type DescribeOpEntitiesRequest struct {
	// The end of the time range to query. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// > The time must be in the latest 30 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1583683200000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The operation object that you want to query.
	//
	// example:
	//
	// 203.***.***.132
	EntityObject *string `json:"EntityObject,omitempty" xml:"EntityObject,omitempty"`
	// The type of the operation object that you want to query. Valid values:
	//
	// 	- **1**: the IP address of the Anti-DDoS Pro or Anti-DDoS Premium instance
	//
	// 	- **2**: Anti-DDoS plans
	//
	// 	- **3**: ECS instances
	//
	// 	- **4**: all logs
	//
	// example:
	//
	// 1
	EntityType *int32 `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
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
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// > The time must be in the latest 30 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1582992000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeOpEntitiesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpEntitiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeOpEntitiesRequest) GetEntityObject() *string {
	return s.EntityObject
}

func (s *DescribeOpEntitiesRequest) GetEntityType() *int32 {
	return s.EntityType
}

func (s *DescribeOpEntitiesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeOpEntitiesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeOpEntitiesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeOpEntitiesRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeOpEntitiesRequest) SetEndTime(v int64) *DescribeOpEntitiesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetEntityObject(v string) *DescribeOpEntitiesRequest {
	s.EntityObject = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetEntityType(v int32) *DescribeOpEntitiesRequest {
	s.EntityType = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetPageNumber(v int32) *DescribeOpEntitiesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetPageSize(v int32) *DescribeOpEntitiesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetResourceGroupId(v string) *DescribeOpEntitiesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetStartTime(v int64) *DescribeOpEntitiesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOpEntitiesRequest) Validate() error {
	return dara.Validate(s)
}
