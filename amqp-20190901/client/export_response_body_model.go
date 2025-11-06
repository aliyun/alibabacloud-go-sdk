// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *ExportResponseBody
  GetCode() *int32 
  SetData(v bool) *ExportResponseBody
  GetData() *bool 
  SetMessage(v string) *ExportResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExportResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExportResponseBody
  GetSuccess() *bool 
}

type ExportResponseBody struct {
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExportResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportResponseBody) GoString() string {
  return s.String()
}

func (s *ExportResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *ExportResponseBody) GetData() *bool  {
  return s.Data
}

func (s *ExportResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExportResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExportResponseBody) SetCode(v int32) *ExportResponseBody {
  s.Code = &v
  return s
}

func (s *ExportResponseBody) SetData(v bool) *ExportResponseBody {
  s.Data = &v
  return s
}

func (s *ExportResponseBody) SetMessage(v string) *ExportResponseBody {
  s.Message = &v
  return s
}

func (s *ExportResponseBody) SetRequestId(v string) *ExportResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportResponseBody) SetSuccess(v bool) *ExportResponseBody {
  s.Success = &v
  return s
}

func (s *ExportResponseBody) Validate() error {
  return dara.Validate(s)
}

