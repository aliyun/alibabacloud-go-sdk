// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iElectronicItineraryBatchApplyRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApplyItineraryList(v []*ElectronicItineraryBatchApplyRequestApplyItineraryList) *ElectronicItineraryBatchApplyRequest
  GetApplyItineraryList() []*ElectronicItineraryBatchApplyRequestApplyItineraryList 
  SetCanReprint(v bool) *ElectronicItineraryBatchApplyRequest
  GetCanReprint() *bool 
}

type ElectronicItineraryBatchApplyRequest struct {
  // This parameter is required.
  ApplyItineraryList []*ElectronicItineraryBatchApplyRequestApplyItineraryList `json:"apply_itinerary_list,omitempty" xml:"apply_itinerary_list,omitempty" type:"Repeated"`
  // example:
  // 
  // true
  CanReprint *bool `json:"can_reprint,omitempty" xml:"can_reprint,omitempty"`
}

func (s ElectronicItineraryBatchApplyRequest) String() string {
  return dara.Prettify(s)
}

func (s ElectronicItineraryBatchApplyRequest) GoString() string {
  return s.String()
}

func (s *ElectronicItineraryBatchApplyRequest) GetApplyItineraryList() []*ElectronicItineraryBatchApplyRequestApplyItineraryList  {
  return s.ApplyItineraryList
}

func (s *ElectronicItineraryBatchApplyRequest) GetCanReprint() *bool  {
  return s.CanReprint
}

func (s *ElectronicItineraryBatchApplyRequest) SetApplyItineraryList(v []*ElectronicItineraryBatchApplyRequestApplyItineraryList) *ElectronicItineraryBatchApplyRequest {
  s.ApplyItineraryList = v
  return s
}

func (s *ElectronicItineraryBatchApplyRequest) SetCanReprint(v bool) *ElectronicItineraryBatchApplyRequest {
  s.CanReprint = &v
  return s
}

func (s *ElectronicItineraryBatchApplyRequest) Validate() error {
  return dara.Validate(s)
}

type ElectronicItineraryBatchApplyRequestApplyItineraryList struct {
  PurchaserName *string `json:"purchaser_name,omitempty" xml:"purchaser_name,omitempty"`
  // example:
  // 
  // tax3213132131
  PurchaserTaxNo *string `json:"purchaser_tax_no,omitempty" xml:"purchaser_tax_no,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 1
  PurchaserType *int32 `json:"purchaser_type,omitempty" xml:"purchaser_type,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 781-2205431917
  TicketNo *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
}

func (s ElectronicItineraryBatchApplyRequestApplyItineraryList) String() string {
  return dara.Prettify(s)
}

func (s ElectronicItineraryBatchApplyRequestApplyItineraryList) GoString() string {
  return s.String()
}

func (s *ElectronicItineraryBatchApplyRequestApplyItineraryList) GetPurchaserName() *string  {
  return s.PurchaserName
}

func (s *ElectronicItineraryBatchApplyRequestApplyItineraryList) GetPurchaserTaxNo() *string  {
  return s.PurchaserTaxNo
}

func (s *ElectronicItineraryBatchApplyRequestApplyItineraryList) GetPurchaserType() *int32  {
  return s.PurchaserType
}

func (s *ElectronicItineraryBatchApplyRequestApplyItineraryList) GetTicketNo() *string  {
  return s.TicketNo
}

func (s *ElectronicItineraryBatchApplyRequestApplyItineraryList) SetPurchaserName(v string) *ElectronicItineraryBatchApplyRequestApplyItineraryList {
  s.PurchaserName = &v
  return s
}

func (s *ElectronicItineraryBatchApplyRequestApplyItineraryList) SetPurchaserTaxNo(v string) *ElectronicItineraryBatchApplyRequestApplyItineraryList {
  s.PurchaserTaxNo = &v
  return s
}

func (s *ElectronicItineraryBatchApplyRequestApplyItineraryList) SetPurchaserType(v int32) *ElectronicItineraryBatchApplyRequestApplyItineraryList {
  s.PurchaserType = &v
  return s
}

func (s *ElectronicItineraryBatchApplyRequestApplyItineraryList) SetTicketNo(v string) *ElectronicItineraryBatchApplyRequestApplyItineraryList {
  s.TicketNo = &v
  return s
}

func (s *ElectronicItineraryBatchApplyRequestApplyItineraryList) Validate() error {
  return dara.Validate(s)
}

