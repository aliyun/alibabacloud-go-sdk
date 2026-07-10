// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleFlightItemListBestPriceItemSubItemPositionMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetJourneyIndex(v int32) *ModuleFlightItemListBestPriceItemSubItemPositionMapValue
	GetJourneyIndex() *int32
	SetSegmentIndex(v int32) *ModuleFlightItemListBestPriceItemSubItemPositionMapValue
	GetSegmentIndex() *int32
}

type ModuleFlightItemListBestPriceItemSubItemPositionMapValue struct {
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s ModuleFlightItemListBestPriceItemSubItemPositionMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListBestPriceItemSubItemPositionMapValue) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListBestPriceItemSubItemPositionMapValue) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *ModuleFlightItemListBestPriceItemSubItemPositionMapValue) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *ModuleFlightItemListBestPriceItemSubItemPositionMapValue) SetJourneyIndex(v int32) *ModuleFlightItemListBestPriceItemSubItemPositionMapValue {
	s.JourneyIndex = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemPositionMapValue) SetSegmentIndex(v int32) *ModuleFlightItemListBestPriceItemSubItemPositionMapValue {
	s.SegmentIndex = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemPositionMapValue) Validate() error {
	return dara.Validate(s)
}
