// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHistoryEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventLevels(v []*string) *DescribeHistoryEventsRequest
	GetEventLevels() []*string
	SetEventStatus(v []*string) *DescribeHistoryEventsRequest
	GetEventStatus() []*string
	SetEventTypes(v []*string) *DescribeHistoryEventsRequest
	GetEventTypes() []*string
	SetPageNumber(v int32) *DescribeHistoryEventsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHistoryEventsRequest
	GetPageSize() *int32
	SetResourceIds(v []*string) *DescribeHistoryEventsRequest
	GetResourceIds() []*string
}

type DescribeHistoryEventsRequest struct {
	// The levels of the event-triggered alerts.
	EventLevels []*string `json:"EventLevels,omitempty" xml:"EventLevels,omitempty" type:"Repeated"`
	// Event status list.
	EventStatus []*string `json:"EventStatus,omitempty" xml:"EventStatus,omitempty" type:"Repeated"`
	// The list of event types.
	//
	// This parameter is required.
	EventTypes []*string `json:"EventTypes,omitempty" xml:"EventTypes,omitempty" type:"Repeated"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The IDs of resources.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
}

func (s DescribeHistoryEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeHistoryEventsRequest) GetEventLevels() []*string {
	return s.EventLevels
}

func (s *DescribeHistoryEventsRequest) GetEventStatus() []*string {
	return s.EventStatus
}

func (s *DescribeHistoryEventsRequest) GetEventTypes() []*string {
	return s.EventTypes
}

func (s *DescribeHistoryEventsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHistoryEventsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHistoryEventsRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *DescribeHistoryEventsRequest) SetEventLevels(v []*string) *DescribeHistoryEventsRequest {
	s.EventLevels = v
	return s
}

func (s *DescribeHistoryEventsRequest) SetEventStatus(v []*string) *DescribeHistoryEventsRequest {
	s.EventStatus = v
	return s
}

func (s *DescribeHistoryEventsRequest) SetEventTypes(v []*string) *DescribeHistoryEventsRequest {
	s.EventTypes = v
	return s
}

func (s *DescribeHistoryEventsRequest) SetPageNumber(v int32) *DescribeHistoryEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHistoryEventsRequest) SetPageSize(v int32) *DescribeHistoryEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHistoryEventsRequest) SetResourceIds(v []*string) *DescribeHistoryEventsRequest {
	s.ResourceIds = v
	return s
}

func (s *DescribeHistoryEventsRequest) Validate() error {
	return dara.Validate(s)
}
