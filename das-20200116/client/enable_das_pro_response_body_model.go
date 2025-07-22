// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDasProResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnableDasProResponseBody
  GetCode() *string 
  SetData(v string) *EnableDasProResponseBody
  GetData() *string 
  SetMessage(v string) *EnableDasProResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableDasProResponseBody
  GetRequestId() *string 
  SetSuccess(v string) *EnableDasProResponseBody
  GetSuccess() *string 
  SetSynchro(v string) *EnableDasProResponseBody
  GetSynchro() *string 
}

type EnableDasProResponseBody struct {
  // The HTTP status code returned.
  // 
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The reserved parameter.
  // 
  // example:
  // 
  // None
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // The returned message.
  // 
  // >  If the request was successful, **Successful*	- is returned. If the request failed, an error message that contains information such as an error code is returned.
  // 
  // example:
  // 
  // Successful
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 7172BECE-588A-5961-8126-C216E16B****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful. Valid values:
  // 
  // 	- **true**
  // 
  // 	- **false**
  // 
  // example:
  // 
  // true
  Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
  // The reserved parameter.
  // 
  // example:
  // 
  // None
  Synchro *string `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s EnableDasProResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableDasProResponseBody) GoString() string {
  return s.String()
}

func (s *EnableDasProResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnableDasProResponseBody) GetData() *string  {
  return s.Data
}

func (s *EnableDasProResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableDasProResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableDasProResponseBody) GetSuccess() *string  {
  return s.Success
}

func (s *EnableDasProResponseBody) GetSynchro() *string  {
  return s.Synchro
}

func (s *EnableDasProResponseBody) SetCode(v string) *EnableDasProResponseBody {
  s.Code = &v
  return s
}

func (s *EnableDasProResponseBody) SetData(v string) *EnableDasProResponseBody {
  s.Data = &v
  return s
}

func (s *EnableDasProResponseBody) SetMessage(v string) *EnableDasProResponseBody {
  s.Message = &v
  return s
}

func (s *EnableDasProResponseBody) SetRequestId(v string) *EnableDasProResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableDasProResponseBody) SetSuccess(v string) *EnableDasProResponseBody {
  s.Success = &v
  return s
}

func (s *EnableDasProResponseBody) SetSynchro(v string) *EnableDasProResponseBody {
  s.Synchro = &v
  return s
}

func (s *EnableDasProResponseBody) Validate() error {
  return dara.Validate(s)
}

