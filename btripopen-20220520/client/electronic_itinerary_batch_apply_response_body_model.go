// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iElectronicItineraryBatchApplyResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *ElectronicItineraryBatchApplyResponseBody
  GetCode() *int32 
  SetMessage(v string) *ElectronicItineraryBatchApplyResponseBody
  GetMessage() *string 
  SetModule(v *ElectronicItineraryBatchApplyResponseBodyModule) *ElectronicItineraryBatchApplyResponseBody
  GetModule() *ElectronicItineraryBatchApplyResponseBodyModule 
  SetRequestId(v string) *ElectronicItineraryBatchApplyResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ElectronicItineraryBatchApplyResponseBody
  GetSuccess() *bool 
  SetTraceId(v string) *ElectronicItineraryBatchApplyResponseBody
  GetTraceId() *string 
}

type ElectronicItineraryBatchApplyResponseBody struct {
  // example:
  // 
  // 200
  Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  Module *ElectronicItineraryBatchApplyResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
  // requestId
  // 
  // example:
  // 
  // A5009956-1077-52FB-B520-EA8C7E91D722
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

func (s ElectronicItineraryBatchApplyResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ElectronicItineraryBatchApplyResponseBody) GoString() string {
  return s.String()
}

func (s *ElectronicItineraryBatchApplyResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *ElectronicItineraryBatchApplyResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ElectronicItineraryBatchApplyResponseBody) GetModule() *ElectronicItineraryBatchApplyResponseBodyModule  {
  return s.Module
}

func (s *ElectronicItineraryBatchApplyResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ElectronicItineraryBatchApplyResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ElectronicItineraryBatchApplyResponseBody) GetTraceId() *string  {
  return s.TraceId
}

func (s *ElectronicItineraryBatchApplyResponseBody) SetCode(v int32) *ElectronicItineraryBatchApplyResponseBody {
  s.Code = &v
  return s
}

func (s *ElectronicItineraryBatchApplyResponseBody) SetMessage(v string) *ElectronicItineraryBatchApplyResponseBody {
  s.Message = &v
  return s
}

func (s *ElectronicItineraryBatchApplyResponseBody) SetModule(v *ElectronicItineraryBatchApplyResponseBodyModule) *ElectronicItineraryBatchApplyResponseBody {
  s.Module = v
  return s
}

func (s *ElectronicItineraryBatchApplyResponseBody) SetRequestId(v string) *ElectronicItineraryBatchApplyResponseBody {
  s.RequestId = &v
  return s
}

func (s *ElectronicItineraryBatchApplyResponseBody) SetSuccess(v bool) *ElectronicItineraryBatchApplyResponseBody {
  s.Success = &v
  return s
}

func (s *ElectronicItineraryBatchApplyResponseBody) SetTraceId(v string) *ElectronicItineraryBatchApplyResponseBody {
  s.TraceId = &v
  return s
}

func (s *ElectronicItineraryBatchApplyResponseBody) Validate() error {
  return dara.Validate(s)
}

type ElectronicItineraryBatchApplyResponseBodyModule struct {
  ApplyFailedTicketList []*ElectronicItineraryBatchApplyResponseBodyModuleApplyFailedTicketList `json:"apply_failed_ticket_list,omitempty" xml:"apply_failed_ticket_list,omitempty" type:"Repeated"`
  BatchNo *string `json:"batch_no,omitempty" xml:"batch_no,omitempty"`
}

func (s ElectronicItineraryBatchApplyResponseBodyModule) String() string {
  return dara.Prettify(s)
}

func (s ElectronicItineraryBatchApplyResponseBodyModule) GoString() string {
  return s.String()
}

func (s *ElectronicItineraryBatchApplyResponseBodyModule) GetApplyFailedTicketList() []*ElectronicItineraryBatchApplyResponseBodyModuleApplyFailedTicketList  {
  return s.ApplyFailedTicketList
}

func (s *ElectronicItineraryBatchApplyResponseBodyModule) GetBatchNo() *string  {
  return s.BatchNo
}

func (s *ElectronicItineraryBatchApplyResponseBodyModule) SetApplyFailedTicketList(v []*ElectronicItineraryBatchApplyResponseBodyModuleApplyFailedTicketList) *ElectronicItineraryBatchApplyResponseBodyModule {
  s.ApplyFailedTicketList = v
  return s
}

func (s *ElectronicItineraryBatchApplyResponseBodyModule) SetBatchNo(v string) *ElectronicItineraryBatchApplyResponseBodyModule {
  s.BatchNo = &v
  return s
}

func (s *ElectronicItineraryBatchApplyResponseBodyModule) Validate() error {
  return dara.Validate(s)
}

type ElectronicItineraryBatchApplyResponseBodyModuleApplyFailedTicketList struct {
  FailedCode *int32 `json:"failed_code,omitempty" xml:"failed_code,omitempty"`
  FailedReason *string `json:"failed_reason,omitempty" xml:"failed_reason,omitempty"`
  TicketNo *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
}

func (s ElectronicItineraryBatchApplyResponseBodyModuleApplyFailedTicketList) String() string {
  return dara.Prettify(s)
}

func (s ElectronicItineraryBatchApplyResponseBodyModuleApplyFailedTicketList) GoString() string {
  return s.String()
}

func (s *ElectronicItineraryBatchApplyResponseBodyModuleApplyFailedTicketList) GetFailedCode() *int32  {
  return s.FailedCode
}

func (s *ElectronicItineraryBatchApplyResponseBodyModuleApplyFailedTicketList) GetFailedReason() *string  {
  return s.FailedReason
}

func (s *ElectronicItineraryBatchApplyResponseBodyModuleApplyFailedTicketList) GetTicketNo() *string  {
  return s.TicketNo
}

func (s *ElectronicItineraryBatchApplyResponseBodyModuleApplyFailedTicketList) SetFailedCode(v int32) *ElectronicItineraryBatchApplyResponseBodyModuleApplyFailedTicketList {
  s.FailedCode = &v
  return s
}

func (s *ElectronicItineraryBatchApplyResponseBodyModuleApplyFailedTicketList) SetFailedReason(v string) *ElectronicItineraryBatchApplyResponseBodyModuleApplyFailedTicketList {
  s.FailedReason = &v
  return s
}

func (s *ElectronicItineraryBatchApplyResponseBodyModuleApplyFailedTicketList) SetTicketNo(v string) *ElectronicItineraryBatchApplyResponseBodyModuleApplyFailedTicketList {
  s.TicketNo = &v
  return s
}

func (s *ElectronicItineraryBatchApplyResponseBodyModuleApplyFailedTicketList) Validate() error {
  return dara.Validate(s)
}

