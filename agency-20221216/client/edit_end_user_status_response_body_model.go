// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditEndUserStatusResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EditEndUserStatusResponseBody
  GetCode() *string 
  SetData(v string) *EditEndUserStatusResponseBody
  GetData() *string 
  SetMessage(v string) *EditEndUserStatusResponseBody
  GetMessage() *string 
  SetMsg(v string) *EditEndUserStatusResponseBody
  GetMsg() *string 
  SetRequestId(v string) *EditEndUserStatusResponseBody
  GetRequestId() *string 
}

type EditEndUserStatusResponseBody struct {
  // The status code.
  // 
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // Indicates whether the operation is successful.
  // 
  // example:
  // 
  // true
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // The prompt message.
  // 
  // example:
  // 
  // 成功
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The prompt message.
  // 
  // example:
  // 
  // success
  Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EditEndUserStatusResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EditEndUserStatusResponseBody) GoString() string {
  return s.String()
}

func (s *EditEndUserStatusResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EditEndUserStatusResponseBody) GetData() *string  {
  return s.Data
}

func (s *EditEndUserStatusResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EditEndUserStatusResponseBody) GetMsg() *string  {
  return s.Msg
}

func (s *EditEndUserStatusResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EditEndUserStatusResponseBody) SetCode(v string) *EditEndUserStatusResponseBody {
  s.Code = &v
  return s
}

func (s *EditEndUserStatusResponseBody) SetData(v string) *EditEndUserStatusResponseBody {
  s.Data = &v
  return s
}

func (s *EditEndUserStatusResponseBody) SetMessage(v string) *EditEndUserStatusResponseBody {
  s.Message = &v
  return s
}

func (s *EditEndUserStatusResponseBody) SetMsg(v string) *EditEndUserStatusResponseBody {
  s.Msg = &v
  return s
}

func (s *EditEndUserStatusResponseBody) SetRequestId(v string) *EditEndUserStatusResponseBody {
  s.RequestId = &v
  return s
}

func (s *EditEndUserStatusResponseBody) Validate() error {
  return dara.Validate(s)
}

