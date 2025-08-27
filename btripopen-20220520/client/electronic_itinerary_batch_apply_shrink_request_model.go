// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iElectronicItineraryBatchApplyShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApplyItineraryListShrink(v string) *ElectronicItineraryBatchApplyShrinkRequest
  GetApplyItineraryListShrink() *string 
  SetCanReprint(v bool) *ElectronicItineraryBatchApplyShrinkRequest
  GetCanReprint() *bool 
}

type ElectronicItineraryBatchApplyShrinkRequest struct {
  // This parameter is required.
  ApplyItineraryListShrink *string `json:"apply_itinerary_list,omitempty" xml:"apply_itinerary_list,omitempty"`
  // example:
  // 
  // true
  CanReprint *bool `json:"can_reprint,omitempty" xml:"can_reprint,omitempty"`
}

func (s ElectronicItineraryBatchApplyShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ElectronicItineraryBatchApplyShrinkRequest) GoString() string {
  return s.String()
}

func (s *ElectronicItineraryBatchApplyShrinkRequest) GetApplyItineraryListShrink() *string  {
  return s.ApplyItineraryListShrink
}

func (s *ElectronicItineraryBatchApplyShrinkRequest) GetCanReprint() *bool  {
  return s.CanReprint
}

func (s *ElectronicItineraryBatchApplyShrinkRequest) SetApplyItineraryListShrink(v string) *ElectronicItineraryBatchApplyShrinkRequest {
  s.ApplyItineraryListShrink = &v
  return s
}

func (s *ElectronicItineraryBatchApplyShrinkRequest) SetCanReprint(v bool) *ElectronicItineraryBatchApplyShrinkRequest {
  s.CanReprint = &v
  return s
}

func (s *ElectronicItineraryBatchApplyShrinkRequest) Validate() error {
  return dara.Validate(s)
}

