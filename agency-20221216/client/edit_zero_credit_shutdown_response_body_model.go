// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditZeroCreditShutdownResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EditZeroCreditShutdownResponseBody
  GetCode() *string 
  SetData(v string) *EditZeroCreditShutdownResponseBody
  GetData() *string 
  SetMessage(v string) *EditZeroCreditShutdownResponseBody
  GetMessage() *string 
  SetMsg(v string) *EditZeroCreditShutdownResponseBody
  GetMsg() *string 
  SetRequestId(v string) *EditZeroCreditShutdownResponseBody
  GetRequestId() *string 
}

type EditZeroCreditShutdownResponseBody struct {
  // Success or not</br>
  // 
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // Request ID</br>
  // 
  // example:
  // 
  // true
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // Message</br>
  // 
  // example:
  // 
  // Message</br>
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // NO_STOP
  // 
  // example:
  // 
  // SUCCESS
  Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
  // success
  // 
  // example:
  // 
  // 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EditZeroCreditShutdownResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EditZeroCreditShutdownResponseBody) GoString() string {
  return s.String()
}

func (s *EditZeroCreditShutdownResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EditZeroCreditShutdownResponseBody) GetData() *string  {
  return s.Data
}

func (s *EditZeroCreditShutdownResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EditZeroCreditShutdownResponseBody) GetMsg() *string  {
  return s.Msg
}

func (s *EditZeroCreditShutdownResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EditZeroCreditShutdownResponseBody) SetCode(v string) *EditZeroCreditShutdownResponseBody {
  s.Code = &v
  return s
}

func (s *EditZeroCreditShutdownResponseBody) SetData(v string) *EditZeroCreditShutdownResponseBody {
  s.Data = &v
  return s
}

func (s *EditZeroCreditShutdownResponseBody) SetMessage(v string) *EditZeroCreditShutdownResponseBody {
  s.Message = &v
  return s
}

func (s *EditZeroCreditShutdownResponseBody) SetMsg(v string) *EditZeroCreditShutdownResponseBody {
  s.Msg = &v
  return s
}

func (s *EditZeroCreditShutdownResponseBody) SetRequestId(v string) *EditZeroCreditShutdownResponseBody {
  s.RequestId = &v
  return s
}

func (s *EditZeroCreditShutdownResponseBody) Validate() error {
  return dara.Validate(s)
}

