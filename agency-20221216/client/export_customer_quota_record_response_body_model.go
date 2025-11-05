// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportCustomerQuotaRecordResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExportCustomerQuotaRecordResponseBody
  GetCode() *string 
  SetData(v *ExportCustomerQuotaRecordResponseBodyData) *ExportCustomerQuotaRecordResponseBody
  GetData() *ExportCustomerQuotaRecordResponseBodyData 
  SetMsg(v string) *ExportCustomerQuotaRecordResponseBody
  GetMsg() *string 
  SetRequestId(v string) *ExportCustomerQuotaRecordResponseBody
  GetRequestId() *string 
}

type ExportCustomerQuotaRecordResponseBody struct {
  // Code
  // 
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // Data
  Data *ExportCustomerQuotaRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // Description
  // 
  // example:
  // 
  // SUCCESS
  Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
  // ID of the Request
  // 
  // example:
  // 
  // 210bc4b416874189683843905d9f9a
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportCustomerQuotaRecordResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportCustomerQuotaRecordResponseBody) GoString() string {
  return s.String()
}

func (s *ExportCustomerQuotaRecordResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExportCustomerQuotaRecordResponseBody) GetData() *ExportCustomerQuotaRecordResponseBodyData  {
  return s.Data
}

func (s *ExportCustomerQuotaRecordResponseBody) GetMsg() *string  {
  return s.Msg
}

func (s *ExportCustomerQuotaRecordResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportCustomerQuotaRecordResponseBody) SetCode(v string) *ExportCustomerQuotaRecordResponseBody {
  s.Code = &v
  return s
}

func (s *ExportCustomerQuotaRecordResponseBody) SetData(v *ExportCustomerQuotaRecordResponseBodyData) *ExportCustomerQuotaRecordResponseBody {
  s.Data = v
  return s
}

func (s *ExportCustomerQuotaRecordResponseBody) SetMsg(v string) *ExportCustomerQuotaRecordResponseBody {
  s.Msg = &v
  return s
}

func (s *ExportCustomerQuotaRecordResponseBody) SetRequestId(v string) *ExportCustomerQuotaRecordResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportCustomerQuotaRecordResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExportCustomerQuotaRecordResponseBodyData struct {
  // Estimated duration, in minutes.
  // 
  // example:
  // 
  // 1
  Cost *int32 `json:"Cost,omitempty" xml:"Cost,omitempty"`
  // ID of Export task
  // 
  // example:
  // 
  // 1231
  Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ExportCustomerQuotaRecordResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExportCustomerQuotaRecordResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExportCustomerQuotaRecordResponseBodyData) GetCost() *int32  {
  return s.Cost
}

func (s *ExportCustomerQuotaRecordResponseBodyData) GetId() *int64  {
  return s.Id
}

func (s *ExportCustomerQuotaRecordResponseBodyData) SetCost(v int32) *ExportCustomerQuotaRecordResponseBodyData {
  s.Cost = &v
  return s
}

func (s *ExportCustomerQuotaRecordResponseBodyData) SetId(v int64) *ExportCustomerQuotaRecordResponseBodyData {
  s.Id = &v
  return s
}

func (s *ExportCustomerQuotaRecordResponseBodyData) Validate() error {
  return dara.Validate(s)
}

