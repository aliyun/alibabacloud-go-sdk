// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportReversedDeductionHistoryResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExportReversedDeductionHistoryResponseBody
  GetCode() *string 
  SetData(v *ExportReversedDeductionHistoryResponseBodyData) *ExportReversedDeductionHistoryResponseBody
  GetData() *ExportReversedDeductionHistoryResponseBodyData 
  SetMessage(v string) *ExportReversedDeductionHistoryResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExportReversedDeductionHistoryResponseBody
  GetRequestId() *string 
}

type ExportReversedDeductionHistoryResponseBody struct {
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *ExportReversedDeductionHistoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // example:
  // 
  // SUCCESS
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // asfasd123
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportReversedDeductionHistoryResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportReversedDeductionHistoryResponseBody) GoString() string {
  return s.String()
}

func (s *ExportReversedDeductionHistoryResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExportReversedDeductionHistoryResponseBody) GetData() *ExportReversedDeductionHistoryResponseBodyData  {
  return s.Data
}

func (s *ExportReversedDeductionHistoryResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExportReversedDeductionHistoryResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportReversedDeductionHistoryResponseBody) SetCode(v string) *ExportReversedDeductionHistoryResponseBody {
  s.Code = &v
  return s
}

func (s *ExportReversedDeductionHistoryResponseBody) SetData(v *ExportReversedDeductionHistoryResponseBodyData) *ExportReversedDeductionHistoryResponseBody {
  s.Data = v
  return s
}

func (s *ExportReversedDeductionHistoryResponseBody) SetMessage(v string) *ExportReversedDeductionHistoryResponseBody {
  s.Message = &v
  return s
}

func (s *ExportReversedDeductionHistoryResponseBody) SetRequestId(v string) *ExportReversedDeductionHistoryResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportReversedDeductionHistoryResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExportReversedDeductionHistoryResponseBodyData struct {
  // example:
  // 
  // 1
  Cost *int32 `json:"Cost,omitempty" xml:"Cost,omitempty"`
  // example:
  // 
  // 2131
  Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ExportReversedDeductionHistoryResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExportReversedDeductionHistoryResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExportReversedDeductionHistoryResponseBodyData) GetCost() *int32  {
  return s.Cost
}

func (s *ExportReversedDeductionHistoryResponseBodyData) GetId() *int64  {
  return s.Id
}

func (s *ExportReversedDeductionHistoryResponseBodyData) SetCost(v int32) *ExportReversedDeductionHistoryResponseBodyData {
  s.Cost = &v
  return s
}

func (s *ExportReversedDeductionHistoryResponseBodyData) SetId(v int64) *ExportReversedDeductionHistoryResponseBodyData {
  s.Id = &v
  return s
}

func (s *ExportReversedDeductionHistoryResponseBodyData) Validate() error {
  return dara.Validate(s)
}

