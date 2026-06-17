// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSystemEventCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeSystemEventCountRequest
	GetEndTime() *string
	SetEventType(v string) *DescribeSystemEventCountRequest
	GetEventType() *string
	SetGroupId(v string) *DescribeSystemEventCountRequest
	GetGroupId() *string
	SetLevel(v string) *DescribeSystemEventCountRequest
	GetLevel() *string
	SetName(v string) *DescribeSystemEventCountRequest
	GetName() *string
	SetProduct(v string) *DescribeSystemEventCountRequest
	GetProduct() *string
	SetRegionId(v string) *DescribeSystemEventCountRequest
	GetRegionId() *string
	SetSearchKeywords(v string) *DescribeSystemEventCountRequest
	GetSearchKeywords() *string
	SetStartTime(v string) *DescribeSystemEventCountRequest
	GetStartTime() *string
	SetStatus(v string) *DescribeSystemEventCountRequest
	GetStatus() *string
}

type DescribeSystemEventCountRequest struct {
	// The timestamp of the end time for the event query. Unit: milliseconds.
	//
	// example:
	//
	// 1635993921000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The event type.
	//
	// Call the DescribeSystemEventMetaList operation to obtain the value of the `EventType` response parameter, which provides the event types for all Alibaba Cloud services under the current Alibaba Cloud account. For more information, see [DescribeSystemEventMetaList](https://help.aliyun.com/document_detail/114972.html).
	//
	// example:
	//
	// StatusNotification
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The application group ID.
	//
	// example:
	//
	// 17285****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The event level. Valid values:
	//
	// - Critical: critical.
	//
	// - Warn: warning.
	//
	// - Info: information.
	//
	// Call the DescribeSystemEventMetaList operation to obtain the value of the `Level` response parameter, which provides the event levels for all Alibaba Cloud services under the current Alibaba Cloud account. For more information, see [DescribeSystemEventMetaList](https://help.aliyun.com/document_detail/114972.html).
	//
	// example:
	//
	// Info
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The event name.
	//
	// Call the DescribeSystemEventMetaList operation to obtain the value of the `Name` response parameter, which provides the event names for all Alibaba Cloud services under the current Alibaba Cloud account. For more information, see [DescribeSystemEventMetaList](https://help.aliyun.com/document_detail/114972.html).
	//
	// example:
	//
	// Instance:StateChange
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the Alibaba Cloud service.
	//
	// Call the DescribeSystemEventMetaList operation to obtain the value of the `Product` response parameter, which provides the Alibaba Cloud service names for all events under the current Alibaba Cloud account. For more information, see [DescribeSystemEventMetaList](https://help.aliyun.com/document_detail/114972.html).
	//
	// example:
	//
	// ECS
	Product  *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The keywords contained in the event content for searching. Valid values:
	//
	// - To search for event content that contains both a and b, search for `a and b`.
	//
	// - To search for event content that contains either a or b, search for `a or b`.
	//
	// example:
	//
	// ECS
	SearchKeywords *string `json:"SearchKeywords,omitempty" xml:"SearchKeywords,omitempty"`
	// The timestamp of the start time for the event query. Unit: milliseconds.
	//
	// example:
	//
	// 1635993541000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The event status.
	//
	// Call the DescribeSystemEventMetaList operation to obtain the value of the `Status` response parameter, which provides the event statuses for all Alibaba Cloud services under the current Alibaba Cloud account. For more information, see [DescribeSystemEventMetaList](https://help.aliyun.com/document_detail/114972.html).
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSystemEventCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemEventCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventCountRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSystemEventCountRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeSystemEventCountRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeSystemEventCountRequest) GetLevel() *string {
	return s.Level
}

func (s *DescribeSystemEventCountRequest) GetName() *string {
	return s.Name
}

func (s *DescribeSystemEventCountRequest) GetProduct() *string {
	return s.Product
}

func (s *DescribeSystemEventCountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSystemEventCountRequest) GetSearchKeywords() *string {
	return s.SearchKeywords
}

func (s *DescribeSystemEventCountRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSystemEventCountRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeSystemEventCountRequest) SetEndTime(v string) *DescribeSystemEventCountRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSystemEventCountRequest) SetEventType(v string) *DescribeSystemEventCountRequest {
	s.EventType = &v
	return s
}

func (s *DescribeSystemEventCountRequest) SetGroupId(v string) *DescribeSystemEventCountRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeSystemEventCountRequest) SetLevel(v string) *DescribeSystemEventCountRequest {
	s.Level = &v
	return s
}

func (s *DescribeSystemEventCountRequest) SetName(v string) *DescribeSystemEventCountRequest {
	s.Name = &v
	return s
}

func (s *DescribeSystemEventCountRequest) SetProduct(v string) *DescribeSystemEventCountRequest {
	s.Product = &v
	return s
}

func (s *DescribeSystemEventCountRequest) SetRegionId(v string) *DescribeSystemEventCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSystemEventCountRequest) SetSearchKeywords(v string) *DescribeSystemEventCountRequest {
	s.SearchKeywords = &v
	return s
}

func (s *DescribeSystemEventCountRequest) SetStartTime(v string) *DescribeSystemEventCountRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSystemEventCountRequest) SetStatus(v string) *DescribeSystemEventCountRequest {
	s.Status = &v
	return s
}

func (s *DescribeSystemEventCountRequest) Validate() error {
	return dara.Validate(s)
}
