// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHistoryEventsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventLevelsShrink(v string) *DescribeHistoryEventsShrinkRequest
	GetEventLevelsShrink() *string
	SetEventStatusShrink(v string) *DescribeHistoryEventsShrinkRequest
	GetEventStatusShrink() *string
	SetEventTypesShrink(v string) *DescribeHistoryEventsShrinkRequest
	GetEventTypesShrink() *string
	SetPageNumber(v int32) *DescribeHistoryEventsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHistoryEventsShrinkRequest
	GetPageSize() *int32
	SetResourceIdsShrink(v string) *DescribeHistoryEventsShrinkRequest
	GetResourceIdsShrink() *string
}

type DescribeHistoryEventsShrinkRequest struct {
	// The levels of the event-triggered alerts.
	EventLevelsShrink *string `json:"EventLevels,omitempty" xml:"EventLevels,omitempty"`
	// Event status list.
	EventStatusShrink *string `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
	// The list of event types.
	//
	// This parameter is required.
	EventTypesShrink *string `json:"EventTypes,omitempty" xml:"EventTypes,omitempty"`
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
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceIdsShrink *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
}

func (s DescribeHistoryEventsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryEventsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeHistoryEventsShrinkRequest) GetEventLevelsShrink() *string {
	return s.EventLevelsShrink
}

func (s *DescribeHistoryEventsShrinkRequest) GetEventStatusShrink() *string {
	return s.EventStatusShrink
}

func (s *DescribeHistoryEventsShrinkRequest) GetEventTypesShrink() *string {
	return s.EventTypesShrink
}

func (s *DescribeHistoryEventsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHistoryEventsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHistoryEventsShrinkRequest) GetResourceIdsShrink() *string {
	return s.ResourceIdsShrink
}

func (s *DescribeHistoryEventsShrinkRequest) SetEventLevelsShrink(v string) *DescribeHistoryEventsShrinkRequest {
	s.EventLevelsShrink = &v
	return s
}

func (s *DescribeHistoryEventsShrinkRequest) SetEventStatusShrink(v string) *DescribeHistoryEventsShrinkRequest {
	s.EventStatusShrink = &v
	return s
}

func (s *DescribeHistoryEventsShrinkRequest) SetEventTypesShrink(v string) *DescribeHistoryEventsShrinkRequest {
	s.EventTypesShrink = &v
	return s
}

func (s *DescribeHistoryEventsShrinkRequest) SetPageNumber(v int32) *DescribeHistoryEventsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHistoryEventsShrinkRequest) SetPageSize(v int32) *DescribeHistoryEventsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHistoryEventsShrinkRequest) SetResourceIdsShrink(v string) *DescribeHistoryEventsShrinkRequest {
	s.ResourceIdsShrink = &v
	return s
}

func (s *DescribeHistoryEventsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
