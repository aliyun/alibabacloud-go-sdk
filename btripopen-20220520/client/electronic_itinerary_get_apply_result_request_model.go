// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iElectronicItineraryGetApplyResultRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBatchApplyNo(v string) *ElectronicItineraryGetApplyResultRequest
  GetBatchApplyNo() *string 
}

type ElectronicItineraryGetApplyResultRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // 0c7d5944608943479a9a4d40c1eec618
  BatchApplyNo *string `json:"batch_apply_no,omitempty" xml:"batch_apply_no,omitempty"`
}

func (s ElectronicItineraryGetApplyResultRequest) String() string {
  return dara.Prettify(s)
}

func (s ElectronicItineraryGetApplyResultRequest) GoString() string {
  return s.String()
}

func (s *ElectronicItineraryGetApplyResultRequest) GetBatchApplyNo() *string  {
  return s.BatchApplyNo
}

func (s *ElectronicItineraryGetApplyResultRequest) SetBatchApplyNo(v string) *ElectronicItineraryGetApplyResultRequest {
  s.BatchApplyNo = &v
  return s
}

func (s *ElectronicItineraryGetApplyResultRequest) Validate() error {
  return dara.Validate(s)
}

