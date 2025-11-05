// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditNewBuyStatusResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EditNewBuyStatusResponseBody
  GetCode() *string 
  SetData(v string) *EditNewBuyStatusResponseBody
  GetData() *string 
  SetMessage(v string) *EditNewBuyStatusResponseBody
  GetMessage() *string 
  SetMsg(v string) *EditNewBuyStatusResponseBody
  GetMsg() *string 
  SetRequestId(v string) *EditNewBuyStatusResponseBody
  GetRequestId() *string 
}

type EditNewBuyStatusResponseBody struct {
  // Status Code</br>
  // 
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // Success or not</br>
  // 
  // example:
  // 
  // true
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // Message</br>
  // 
  // example:
  // 
  // success
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // Message</br>
  // 
  // example:
  // 
  // success
  Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
  // Request ID</br>
  // 
  // example:
  // 
  // 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EditNewBuyStatusResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EditNewBuyStatusResponseBody) GoString() string {
  return s.String()
}

func (s *EditNewBuyStatusResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EditNewBuyStatusResponseBody) GetData() *string  {
  return s.Data
}

func (s *EditNewBuyStatusResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EditNewBuyStatusResponseBody) GetMsg() *string  {
  return s.Msg
}

func (s *EditNewBuyStatusResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EditNewBuyStatusResponseBody) SetCode(v string) *EditNewBuyStatusResponseBody {
  s.Code = &v
  return s
}

func (s *EditNewBuyStatusResponseBody) SetData(v string) *EditNewBuyStatusResponseBody {
  s.Data = &v
  return s
}

func (s *EditNewBuyStatusResponseBody) SetMessage(v string) *EditNewBuyStatusResponseBody {
  s.Message = &v
  return s
}

func (s *EditNewBuyStatusResponseBody) SetMsg(v string) *EditNewBuyStatusResponseBody {
  s.Msg = &v
  return s
}

func (s *EditNewBuyStatusResponseBody) SetRequestId(v string) *EditNewBuyStatusResponseBody {
  s.RequestId = &v
  return s
}

func (s *EditNewBuyStatusResponseBody) Validate() error {
  return dara.Validate(s)
}

