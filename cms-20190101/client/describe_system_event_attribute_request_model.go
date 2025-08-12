// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSystemEventAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeSystemEventAttributeRequest
	GetEndTime() *string
	SetEventType(v string) *DescribeSystemEventAttributeRequest
	GetEventType() *string
	SetGroupId(v string) *DescribeSystemEventAttributeRequest
	GetGroupId() *string
	SetLevel(v string) *DescribeSystemEventAttributeRequest
	GetLevel() *string
	SetName(v string) *DescribeSystemEventAttributeRequest
	GetName() *string
	SetPageNumber(v int32) *DescribeSystemEventAttributeRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSystemEventAttributeRequest
	GetPageSize() *int32
	SetProduct(v string) *DescribeSystemEventAttributeRequest
	GetProduct() *string
	SetRegionId(v string) *DescribeSystemEventAttributeRequest
	GetRegionId() *string
	SetSearchKeywords(v string) *DescribeSystemEventAttributeRequest
	GetSearchKeywords() *string
	SetStartTime(v string) *DescribeSystemEventAttributeRequest
	GetStartTime() *string
	SetStatus(v string) *DescribeSystemEventAttributeRequest
	GetStatus() *string
}

type DescribeSystemEventAttributeRequest struct {
	// The end of the time range to query.
	//
	// The value must be a UNIX timestamp. It is the number of seconds that have elapsed since 00:00:00 UTC, January 1, 1970.
	//
	// example:
	//
	// 1552221584949
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of the system event.
	//
	// >  You can call the [DescribeSystemEventMetaList](https://help.aliyun.com/document_detail/114972.html) operation to query the types of system events.
	//
	// example:
	//
	// Exception
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The ID of the application group.
	//
	// example:
	//
	// 12346
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The level of the system event. Valid values:
	//
	// 	- CRITICAL: critical
	//
	// 	- WARN: warning
	//
	// 	- INFO: information
	//
	// example:
	//
	// CRITICAL
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The name of the system event.
	//
	// >  You can call the [DescribeSystemEventMetaList](https://help.aliyun.com/document_detail/114972.html) operation to query the names of system events.
	//
	// example:
	//
	// BucketIngressBandwidth
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of the page to return.
	//
	// Valid values: 1 to 100000000.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The abbreviation of the service name.
	//
	// >  You can call the [DescribeSystemEventMetaList](https://help.aliyun.com/document_detail/114972.html) operation to query the abbreviations of service names.
	//
	// example:
	//
	// oss
	Product  *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The keywords that are used to search for the system event. Valid values:
	//
	// 	- If you want to search for the system event whose content contains a and b, set the value to `a and b`.
	//
	// 	- If you want to search for the system event whose content contains a or b, set the value to `a or b`.
	//
	// example:
	//
	// cms
	SearchKeywords *string `json:"SearchKeywords,omitempty" xml:"SearchKeywords,omitempty"`
	// The beginning of the time range to query.
	//
	// The value must be a UNIX timestamp. It is the number of seconds that have elapsed since 00:00:00 UTC, January 1, 1970.
	//
	// example:
	//
	// 1552199984949
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the system event.
	//
	// >  You can call the [DescribeSystemEventMetaList](https://help.aliyun.com/document_detail/114972.html) operation to query the statuses of system events.
	//
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSystemEventAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemEventAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventAttributeRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSystemEventAttributeRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeSystemEventAttributeRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeSystemEventAttributeRequest) GetLevel() *string {
	return s.Level
}

func (s *DescribeSystemEventAttributeRequest) GetName() *string {
	return s.Name
}

func (s *DescribeSystemEventAttributeRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSystemEventAttributeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSystemEventAttributeRequest) GetProduct() *string {
	return s.Product
}

func (s *DescribeSystemEventAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSystemEventAttributeRequest) GetSearchKeywords() *string {
	return s.SearchKeywords
}

func (s *DescribeSystemEventAttributeRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSystemEventAttributeRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeSystemEventAttributeRequest) SetEndTime(v string) *DescribeSystemEventAttributeRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSystemEventAttributeRequest) SetEventType(v string) *DescribeSystemEventAttributeRequest {
	s.EventType = &v
	return s
}

func (s *DescribeSystemEventAttributeRequest) SetGroupId(v string) *DescribeSystemEventAttributeRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeSystemEventAttributeRequest) SetLevel(v string) *DescribeSystemEventAttributeRequest {
	s.Level = &v
	return s
}

func (s *DescribeSystemEventAttributeRequest) SetName(v string) *DescribeSystemEventAttributeRequest {
	s.Name = &v
	return s
}

func (s *DescribeSystemEventAttributeRequest) SetPageNumber(v int32) *DescribeSystemEventAttributeRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSystemEventAttributeRequest) SetPageSize(v int32) *DescribeSystemEventAttributeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSystemEventAttributeRequest) SetProduct(v string) *DescribeSystemEventAttributeRequest {
	s.Product = &v
	return s
}

func (s *DescribeSystemEventAttributeRequest) SetRegionId(v string) *DescribeSystemEventAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSystemEventAttributeRequest) SetSearchKeywords(v string) *DescribeSystemEventAttributeRequest {
	s.SearchKeywords = &v
	return s
}

func (s *DescribeSystemEventAttributeRequest) SetStartTime(v string) *DescribeSystemEventAttributeRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSystemEventAttributeRequest) SetStatus(v string) *DescribeSystemEventAttributeRequest {
	s.Status = &v
	return s
}

func (s *DescribeSystemEventAttributeRequest) Validate() error {
	return dara.Validate(s)
}
