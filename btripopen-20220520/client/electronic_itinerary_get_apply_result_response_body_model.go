// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iElectronicItineraryGetApplyResultResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *ElectronicItineraryGetApplyResultResponseBody
  GetCode() *int32 
  SetMessage(v string) *ElectronicItineraryGetApplyResultResponseBody
  GetMessage() *string 
  SetModule(v *ElectronicItineraryGetApplyResultResponseBodyModule) *ElectronicItineraryGetApplyResultResponseBody
  GetModule() *ElectronicItineraryGetApplyResultResponseBodyModule 
  SetRequestId(v string) *ElectronicItineraryGetApplyResultResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ElectronicItineraryGetApplyResultResponseBody
  GetSuccess() *bool 
  SetTraceId(v string) *ElectronicItineraryGetApplyResultResponseBody
  GetTraceId() *string 
}

type ElectronicItineraryGetApplyResultResponseBody struct {
  // example:
  // 
  // 200
  Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  Module *ElectronicItineraryGetApplyResultResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
  // requestId
  // 
  // example:
  // 
  // B72B39C8-32DE-558D-AD1C-D53F11F6ADFE
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
  // traceId
  // 
  // example:
  // 
  // 21041ce316577904808056433edbb2
  TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s ElectronicItineraryGetApplyResultResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ElectronicItineraryGetApplyResultResponseBody) GoString() string {
  return s.String()
}

func (s *ElectronicItineraryGetApplyResultResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *ElectronicItineraryGetApplyResultResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ElectronicItineraryGetApplyResultResponseBody) GetModule() *ElectronicItineraryGetApplyResultResponseBodyModule  {
  return s.Module
}

func (s *ElectronicItineraryGetApplyResultResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ElectronicItineraryGetApplyResultResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ElectronicItineraryGetApplyResultResponseBody) GetTraceId() *string  {
  return s.TraceId
}

func (s *ElectronicItineraryGetApplyResultResponseBody) SetCode(v int32) *ElectronicItineraryGetApplyResultResponseBody {
  s.Code = &v
  return s
}

func (s *ElectronicItineraryGetApplyResultResponseBody) SetMessage(v string) *ElectronicItineraryGetApplyResultResponseBody {
  s.Message = &v
  return s
}

func (s *ElectronicItineraryGetApplyResultResponseBody) SetModule(v *ElectronicItineraryGetApplyResultResponseBodyModule) *ElectronicItineraryGetApplyResultResponseBody {
  s.Module = v
  return s
}

func (s *ElectronicItineraryGetApplyResultResponseBody) SetRequestId(v string) *ElectronicItineraryGetApplyResultResponseBody {
  s.RequestId = &v
  return s
}

func (s *ElectronicItineraryGetApplyResultResponseBody) SetSuccess(v bool) *ElectronicItineraryGetApplyResultResponseBody {
  s.Success = &v
  return s
}

func (s *ElectronicItineraryGetApplyResultResponseBody) SetTraceId(v string) *ElectronicItineraryGetApplyResultResponseBody {
  s.TraceId = &v
  return s
}

func (s *ElectronicItineraryGetApplyResultResponseBody) Validate() error {
  return dara.Validate(s)
}

type ElectronicItineraryGetApplyResultResponseBodyModule struct {
  ApplyTicketList []*ElectronicItineraryGetApplyResultResponseBodyModuleApplyTicketList `json:"apply_ticket_list,omitempty" xml:"apply_ticket_list,omitempty" type:"Repeated"`
  BatchApplyNo *string `json:"batch_apply_no,omitempty" xml:"batch_apply_no,omitempty"`
}

func (s ElectronicItineraryGetApplyResultResponseBodyModule) String() string {
  return dara.Prettify(s)
}

func (s ElectronicItineraryGetApplyResultResponseBodyModule) GoString() string {
  return s.String()
}

func (s *ElectronicItineraryGetApplyResultResponseBodyModule) GetApplyTicketList() []*ElectronicItineraryGetApplyResultResponseBodyModuleApplyTicketList  {
  return s.ApplyTicketList
}

func (s *ElectronicItineraryGetApplyResultResponseBodyModule) GetBatchApplyNo() *string  {
  return s.BatchApplyNo
}

func (s *ElectronicItineraryGetApplyResultResponseBodyModule) SetApplyTicketList(v []*ElectronicItineraryGetApplyResultResponseBodyModuleApplyTicketList) *ElectronicItineraryGetApplyResultResponseBodyModule {
  s.ApplyTicketList = v
  return s
}

func (s *ElectronicItineraryGetApplyResultResponseBodyModule) SetBatchApplyNo(v string) *ElectronicItineraryGetApplyResultResponseBodyModule {
  s.BatchApplyNo = &v
  return s
}

func (s *ElectronicItineraryGetApplyResultResponseBodyModule) Validate() error {
  return dara.Validate(s)
}

type ElectronicItineraryGetApplyResultResponseBodyModuleApplyTicketList struct {
  // example:
  // 
  // 5001
  FailedCode *int32 `json:"failed_code,omitempty" xml:"failed_code,omitempty"`
  FailedReason *string `json:"failed_reason,omitempty" xml:"failed_reason,omitempty"`
  ItineraryStatus *int32 `json:"itinerary_status,omitempty" xml:"itinerary_status,omitempty"`
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
  // example:
  // 
  // 781-6605285563
  TicketNo *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
}

func (s ElectronicItineraryGetApplyResultResponseBodyModuleApplyTicketList) String() string {
  return dara.Prettify(s)
}

func (s ElectronicItineraryGetApplyResultResponseBodyModuleApplyTicketList) GoString() string {
  return s.String()
}

func (s *ElectronicItineraryGetApplyResultResponseBodyModuleApplyTicketList) GetFailedCode() *int32  {
  return s.FailedCode
}

func (s *ElectronicItineraryGetApplyResultResponseBodyModuleApplyTicketList) GetFailedReason() *string  {
  return s.FailedReason
}

func (s *ElectronicItineraryGetApplyResultResponseBodyModuleApplyTicketList) GetItineraryStatus() *int32  {
  return s.ItineraryStatus
}

func (s *ElectronicItineraryGetApplyResultResponseBodyModuleApplyTicketList) GetRemark() *string  {
  return s.Remark
}

func (s *ElectronicItineraryGetApplyResultResponseBodyModuleApplyTicketList) GetTicketNo() *string  {
  return s.TicketNo
}

func (s *ElectronicItineraryGetApplyResultResponseBodyModuleApplyTicketList) SetFailedCode(v int32) *ElectronicItineraryGetApplyResultResponseBodyModuleApplyTicketList {
  s.FailedCode = &v
  return s
}

func (s *ElectronicItineraryGetApplyResultResponseBodyModuleApplyTicketList) SetFailedReason(v string) *ElectronicItineraryGetApplyResultResponseBodyModuleApplyTicketList {
  s.FailedReason = &v
  return s
}

func (s *ElectronicItineraryGetApplyResultResponseBodyModuleApplyTicketList) SetItineraryStatus(v int32) *ElectronicItineraryGetApplyResultResponseBodyModuleApplyTicketList {
  s.ItineraryStatus = &v
  return s
}

func (s *ElectronicItineraryGetApplyResultResponseBodyModuleApplyTicketList) SetRemark(v string) *ElectronicItineraryGetApplyResultResponseBodyModuleApplyTicketList {
  s.Remark = &v
  return s
}

func (s *ElectronicItineraryGetApplyResultResponseBodyModuleApplyTicketList) SetTicketNo(v string) *ElectronicItineraryGetApplyResultResponseBodyModuleApplyTicketList {
  s.TicketNo = &v
  return s
}

func (s *ElectronicItineraryGetApplyResultResponseBodyModuleApplyTicketList) Validate() error {
  return dara.Validate(s)
}

