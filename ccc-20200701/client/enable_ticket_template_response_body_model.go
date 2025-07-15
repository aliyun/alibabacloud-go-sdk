// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableTicketTemplateResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnableTicketTemplateResponseBody
  GetCode() *string 
  SetData(v interface{}) *EnableTicketTemplateResponseBody
  GetData() interface{} 
  SetHttpStatusCode(v int32) *EnableTicketTemplateResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *EnableTicketTemplateResponseBody
  GetMessage() *string 
  SetParams(v []*string) *EnableTicketTemplateResponseBody
  GetParams() []*string 
  SetRequestId(v string) *EnableTicketTemplateResponseBody
  GetRequestId() *string 
}

type EnableTicketTemplateResponseBody struct {
  // example:
  // 
  // OK
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
  // example:
  // 
  // BA03159C-E808-4FF1-B27E-A61B6E888D7F
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableTicketTemplateResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableTicketTemplateResponseBody) GoString() string {
  return s.String()
}

func (s *EnableTicketTemplateResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnableTicketTemplateResponseBody) GetData() interface{}  {
  return s.Data
}

func (s *EnableTicketTemplateResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *EnableTicketTemplateResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableTicketTemplateResponseBody) GetParams() []*string  {
  return s.Params
}

func (s *EnableTicketTemplateResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableTicketTemplateResponseBody) SetCode(v string) *EnableTicketTemplateResponseBody {
  s.Code = &v
  return s
}

func (s *EnableTicketTemplateResponseBody) SetData(v interface{}) *EnableTicketTemplateResponseBody {
  s.Data = v
  return s
}

func (s *EnableTicketTemplateResponseBody) SetHttpStatusCode(v int32) *EnableTicketTemplateResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *EnableTicketTemplateResponseBody) SetMessage(v string) *EnableTicketTemplateResponseBody {
  s.Message = &v
  return s
}

func (s *EnableTicketTemplateResponseBody) SetParams(v []*string) *EnableTicketTemplateResponseBody {
  s.Params = v
  return s
}

func (s *EnableTicketTemplateResponseBody) SetRequestId(v string) *EnableTicketTemplateResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableTicketTemplateResponseBody) Validate() error {
  return dara.Validate(s)
}

