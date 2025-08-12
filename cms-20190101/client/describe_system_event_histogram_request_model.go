// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSystemEventHistogramRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeSystemEventHistogramRequest
	GetEndTime() *string
	SetEventType(v string) *DescribeSystemEventHistogramRequest
	GetEventType() *string
	SetGroupId(v string) *DescribeSystemEventHistogramRequest
	GetGroupId() *string
	SetLevel(v string) *DescribeSystemEventHistogramRequest
	GetLevel() *string
	SetName(v string) *DescribeSystemEventHistogramRequest
	GetName() *string
	SetProduct(v string) *DescribeSystemEventHistogramRequest
	GetProduct() *string
	SetRegionId(v string) *DescribeSystemEventHistogramRequest
	GetRegionId() *string
	SetSearchKeywords(v string) *DescribeSystemEventHistogramRequest
	GetSearchKeywords() *string
	SetStartTime(v string) *DescribeSystemEventHistogramRequest
	GetStartTime() *string
	SetStatus(v string) *DescribeSystemEventHistogramRequest
	GetStatus() *string
}

type DescribeSystemEventHistogramRequest struct {
	// The end time.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1552220485596
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The event type.
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
	// 12345
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The level of the event. Valid values:
	//
	// 	- CRITICAL
	//
	// 	- WARN
	//
	// 	- INFO
	//
	// example:
	//
	// CRITICAL
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The event name.
	//
	// >  You can call the [DescribeSystemEventMetaList](https://help.aliyun.com/document_detail/114972.html) operation to query the names of system events.
	//
	// example:
	//
	// BucketIngressBandwidth
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The abbreviation of the service name.
	//
	// >  You can call the [DescribeSystemEventMetaList](https://help.aliyun.com/document_detail/114972.html) operation to query the abbreviations of service names.
	//
	// example:
	//
	// OSS
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
	// The start time.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1552209685596
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The event status.
	//
	// >  You can call the [DescribeSystemEventMetaList](https://help.aliyun.com/document_detail/114972.html) operation to query the status of system events.
	//
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSystemEventHistogramRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemEventHistogramRequest) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventHistogramRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSystemEventHistogramRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeSystemEventHistogramRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeSystemEventHistogramRequest) GetLevel() *string {
	return s.Level
}

func (s *DescribeSystemEventHistogramRequest) GetName() *string {
	return s.Name
}

func (s *DescribeSystemEventHistogramRequest) GetProduct() *string {
	return s.Product
}

func (s *DescribeSystemEventHistogramRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSystemEventHistogramRequest) GetSearchKeywords() *string {
	return s.SearchKeywords
}

func (s *DescribeSystemEventHistogramRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSystemEventHistogramRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeSystemEventHistogramRequest) SetEndTime(v string) *DescribeSystemEventHistogramRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSystemEventHistogramRequest) SetEventType(v string) *DescribeSystemEventHistogramRequest {
	s.EventType = &v
	return s
}

func (s *DescribeSystemEventHistogramRequest) SetGroupId(v string) *DescribeSystemEventHistogramRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeSystemEventHistogramRequest) SetLevel(v string) *DescribeSystemEventHistogramRequest {
	s.Level = &v
	return s
}

func (s *DescribeSystemEventHistogramRequest) SetName(v string) *DescribeSystemEventHistogramRequest {
	s.Name = &v
	return s
}

func (s *DescribeSystemEventHistogramRequest) SetProduct(v string) *DescribeSystemEventHistogramRequest {
	s.Product = &v
	return s
}

func (s *DescribeSystemEventHistogramRequest) SetRegionId(v string) *DescribeSystemEventHistogramRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSystemEventHistogramRequest) SetSearchKeywords(v string) *DescribeSystemEventHistogramRequest {
	s.SearchKeywords = &v
	return s
}

func (s *DescribeSystemEventHistogramRequest) SetStartTime(v string) *DescribeSystemEventHistogramRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSystemEventHistogramRequest) SetStatus(v string) *DescribeSystemEventHistogramRequest {
	s.Status = &v
	return s
}

func (s *DescribeSystemEventHistogramRequest) Validate() error {
	return dara.Validate(s)
}
