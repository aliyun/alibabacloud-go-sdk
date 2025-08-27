// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleFlightItemListItemListSubItemPositionMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetJourneyIndex(v int32) *ModuleFlightItemListItemListSubItemPositionMapValue
	GetJourneyIndex() *int32
	SetSegmentIndex(v int32) *ModuleFlightItemListItemListSubItemPositionMapValue
	GetSegmentIndex() *int32
}

type ModuleFlightItemListItemListSubItemPositionMapValue struct {
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s ModuleFlightItemListItemListSubItemPositionMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListItemListSubItemPositionMapValue) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListItemListSubItemPositionMapValue) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *ModuleFlightItemListItemListSubItemPositionMapValue) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *ModuleFlightItemListItemListSubItemPositionMapValue) SetJourneyIndex(v int32) *ModuleFlightItemListItemListSubItemPositionMapValue {
	s.JourneyIndex = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemPositionMapValue) SetSegmentIndex(v int32) *ModuleFlightItemListItemListSubItemPositionMapValue {
	s.SegmentIndex = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemPositionMapValue) Validate() error {
	return dara.Validate(s)
}
