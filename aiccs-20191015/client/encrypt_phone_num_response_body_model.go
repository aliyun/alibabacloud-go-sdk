// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEncryptPhoneNumResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EncryptPhoneNumResponseBody
  GetCode() *string 
  SetData(v string) *EncryptPhoneNumResponseBody
  GetData() *string 
  SetMessage(v string) *EncryptPhoneNumResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EncryptPhoneNumResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EncryptPhoneNumResponseBody
  GetSuccess() *bool 
}

type EncryptPhoneNumResponseBody struct {
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // example:
  // 
  // xxxx
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // EE338D98-9BD3-4413-B165
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EncryptPhoneNumResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EncryptPhoneNumResponseBody) GoString() string {
  return s.String()
}

func (s *EncryptPhoneNumResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EncryptPhoneNumResponseBody) GetData() *string  {
  return s.Data
}

func (s *EncryptPhoneNumResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EncryptPhoneNumResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EncryptPhoneNumResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EncryptPhoneNumResponseBody) SetCode(v string) *EncryptPhoneNumResponseBody {
  s.Code = &v
  return s
}

func (s *EncryptPhoneNumResponseBody) SetData(v string) *EncryptPhoneNumResponseBody {
  s.Data = &v
  return s
}

func (s *EncryptPhoneNumResponseBody) SetMessage(v string) *EncryptPhoneNumResponseBody {
  s.Message = &v
  return s
}

func (s *EncryptPhoneNumResponseBody) SetRequestId(v string) *EncryptPhoneNumResponseBody {
  s.RequestId = &v
  return s
}

func (s *EncryptPhoneNumResponseBody) SetSuccess(v bool) *EncryptPhoneNumResponseBody {
  s.Success = &v
  return s
}

func (s *EncryptPhoneNumResponseBody) Validate() error {
  return dara.Validate(s)
}

