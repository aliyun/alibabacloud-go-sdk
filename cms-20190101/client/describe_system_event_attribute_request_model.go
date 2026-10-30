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
	// The end time.
	//
	// The value is a UNIX timestamp that represents the number of milliseconds that have elapsed since January 1, 1970.
	//
	// example:
	//
	// 1552221584949
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The event type.
	//
	// >For more information about valid values for the event type, see [DescribeSystemEventMetaList](https://help.aliyun.com/document_detail/114972.html).
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
	// The event level. Valid values:
	//
	// - CRITICAL: critical.
	//
	// - WARN: warning.
	//
	// - INFO: information.
	//
	// example:
	//
	// CRITICAL
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The event name.
	//
	// >For more information about valid values for the event name, see [DescribeSystemEventMetaList](https://help.aliyun.com/document_detail/114972.html).
	//
	// example:
	//
	// BucketIngressBandwidth
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number.
	//
	// Valid values: 1 to 100000000.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The abbreviation of the product name.
	//
	// >For more information about valid values for the product name abbreviation, see [DescribeSystemEventMetaList](https://help.aliyun.com/document_detail/114972.html).
	//
	// example:
	//
	// oss
	Product  *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The keywords used to search event content. Valid values:
	//
	// - If the event content you want to search for includes both a and b, you can search for `a and b`.
	//
	// - If the event content you want to search for includes either a or b, you can search for `a or b`.
	//
	// example:
	//
	// cms
	SearchKeywords *string `json:"SearchKeywords,omitempty" xml:"SearchKeywords,omitempty"`
	// The start time.
	//
	// The value is a UNIX timestamp that represents the number of milliseconds that have elapsed since January 1, 1970.
	//
	// example:
	//
	// 1552199984949
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The event status.
	//
	// >For more information about valid values for the event status, see [DescribeSystemEventMetaList](https://help.aliyun.com/document_detail/114972.html).
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
