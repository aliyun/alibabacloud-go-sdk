// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleItemListSubItemPositionMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetJourneyIndex(v int32) *ModuleItemListSubItemPositionMapValue
	GetJourneyIndex() *int32
	SetSegmentIndex(v int32) *ModuleItemListSubItemPositionMapValue
	GetSegmentIndex() *int32
}

type ModuleItemListSubItemPositionMapValue struct {
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s ModuleItemListSubItemPositionMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleItemListSubItemPositionMapValue) GoString() string {
	return s.String()
}

func (s *ModuleItemListSubItemPositionMapValue) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *ModuleItemListSubItemPositionMapValue) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *ModuleItemListSubItemPositionMapValue) SetJourneyIndex(v int32) *ModuleItemListSubItemPositionMapValue {
	s.JourneyIndex = &v
	return s
}

func (s *ModuleItemListSubItemPositionMapValue) SetSegmentIndex(v int32) *ModuleItemListSubItemPositionMapValue {
	s.SegmentIndex = &v
	return s
}

func (s *ModuleItemListSubItemPositionMapValue) Validate() error {
	return dara.Validate(s)
}
