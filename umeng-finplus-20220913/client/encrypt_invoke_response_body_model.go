// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEncryptInvokeResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EncryptInvokeResponseBody
  GetCode() *string 
  SetData(v *EncryptInvokeResponseBodyData) *EncryptInvokeResponseBody
  GetData() *EncryptInvokeResponseBodyData 
  SetMsg(v string) *EncryptInvokeResponseBody
  GetMsg() *string 
  SetRequestId(v string) *EncryptInvokeResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EncryptInvokeResponseBody
  GetSuccess() *bool 
}

type EncryptInvokeResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *EncryptInvokeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EncryptInvokeResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EncryptInvokeResponseBody) GoString() string {
  return s.String()
}

func (s *EncryptInvokeResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EncryptInvokeResponseBody) GetData() *EncryptInvokeResponseBodyData  {
  return s.Data
}

func (s *EncryptInvokeResponseBody) GetMsg() *string  {
  return s.Msg
}

func (s *EncryptInvokeResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EncryptInvokeResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EncryptInvokeResponseBody) SetCode(v string) *EncryptInvokeResponseBody {
  s.Code = &v
  return s
}

func (s *EncryptInvokeResponseBody) SetData(v *EncryptInvokeResponseBodyData) *EncryptInvokeResponseBody {
  s.Data = v
  return s
}

func (s *EncryptInvokeResponseBody) SetMsg(v string) *EncryptInvokeResponseBody {
  s.Msg = &v
  return s
}

func (s *EncryptInvokeResponseBody) SetRequestId(v string) *EncryptInvokeResponseBody {
  s.RequestId = &v
  return s
}

func (s *EncryptInvokeResponseBody) SetSuccess(v bool) *EncryptInvokeResponseBody {
  s.Success = &v
  return s
}

func (s *EncryptInvokeResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EncryptInvokeResponseBodyData struct {
  EncryptData *string `json:"encryptData,omitempty" xml:"encryptData,omitempty"`
  EncryptKey *string `json:"encryptKey,omitempty" xml:"encryptKey,omitempty"`
  Sign *string `json:"sign,omitempty" xml:"sign,omitempty"`
}

func (s EncryptInvokeResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EncryptInvokeResponseBodyData) GoString() string {
  return s.String()
}

func (s *EncryptInvokeResponseBodyData) GetEncryptData() *string  {
  return s.EncryptData
}

func (s *EncryptInvokeResponseBodyData) GetEncryptKey() *string  {
  return s.EncryptKey
}

func (s *EncryptInvokeResponseBodyData) GetSign() *string  {
  return s.Sign
}

func (s *EncryptInvokeResponseBodyData) SetEncryptData(v string) *EncryptInvokeResponseBodyData {
  s.EncryptData = &v
  return s
}

func (s *EncryptInvokeResponseBodyData) SetEncryptKey(v string) *EncryptInvokeResponseBodyData {
  s.EncryptKey = &v
  return s
}

func (s *EncryptInvokeResponseBodyData) SetSign(v string) *EncryptInvokeResponseBodyData {
  s.Sign = &v
  return s
}

func (s *EncryptInvokeResponseBodyData) Validate() error {
  return dara.Validate(s)
}

