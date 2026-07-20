// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleGroupItemSubItemPositionMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetJourneyIndex(v int32) *ModuleGroupItemSubItemPositionMapValue
	GetJourneyIndex() *int32
	SetSegmentIndex(v int32) *ModuleGroupItemSubItemPositionMapValue
	GetSegmentIndex() *int32
}

type ModuleGroupItemSubItemPositionMapValue struct {
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s ModuleGroupItemSubItemPositionMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleGroupItemSubItemPositionMapValue) GoString() string {
	return s.String()
}

func (s *ModuleGroupItemSubItemPositionMapValue) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *ModuleGroupItemSubItemPositionMapValue) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *ModuleGroupItemSubItemPositionMapValue) SetJourneyIndex(v int32) *ModuleGroupItemSubItemPositionMapValue {
	s.JourneyIndex = &v
	return s
}

func (s *ModuleGroupItemSubItemPositionMapValue) SetSegmentIndex(v int32) *ModuleGroupItemSubItemPositionMapValue {
	s.SegmentIndex = &v
	return s
}

func (s *ModuleGroupItemSubItemPositionMapValue) Validate() error {
	return dara.Validate(s)
}
