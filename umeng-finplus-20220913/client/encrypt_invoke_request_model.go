// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEncryptInvokeRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClientId(v int64) *EncryptInvokeRequest
  GetClientId() *int64 
  SetData(v string) *EncryptInvokeRequest
  GetData() *string 
  SetEncryptKey(v string) *EncryptInvokeRequest
  GetEncryptKey() *string 
  SetMethodName(v string) *EncryptInvokeRequest
  GetMethodName() *string 
  SetSign(v string) *EncryptInvokeRequest
  GetSign() *string 
}

type EncryptInvokeRequest struct {
  ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
  Data *string `json:"data,omitempty" xml:"data,omitempty"`
  EncryptKey *string `json:"encryptKey,omitempty" xml:"encryptKey,omitempty"`
  MethodName *string `json:"methodName,omitempty" xml:"methodName,omitempty"`
  Sign *string `json:"sign,omitempty" xml:"sign,omitempty"`
}

func (s EncryptInvokeRequest) String() string {
  return dara.Prettify(s)
}

func (s EncryptInvokeRequest) GoString() string {
  return s.String()
}

func (s *EncryptInvokeRequest) GetClientId() *int64  {
  return s.ClientId
}

func (s *EncryptInvokeRequest) GetData() *string  {
  return s.Data
}

func (s *EncryptInvokeRequest) GetEncryptKey() *string  {
  return s.EncryptKey
}

func (s *EncryptInvokeRequest) GetMethodName() *string  {
  return s.MethodName
}

func (s *EncryptInvokeRequest) GetSign() *string  {
  return s.Sign
}

func (s *EncryptInvokeRequest) SetClientId(v int64) *EncryptInvokeRequest {
  s.ClientId = &v
  return s
}

func (s *EncryptInvokeRequest) SetData(v string) *EncryptInvokeRequest {
  s.Data = &v
  return s
}

func (s *EncryptInvokeRequest) SetEncryptKey(v string) *EncryptInvokeRequest {
  s.EncryptKey = &v
  return s
}

func (s *EncryptInvokeRequest) SetMethodName(v string) *EncryptInvokeRequest {
  s.MethodName = &v
  return s
}

func (s *EncryptInvokeRequest) SetSign(v string) *EncryptInvokeRequest {
  s.Sign = &v
  return s
}

func (s *EncryptInvokeRequest) Validate() error {
  return dara.Validate(s)
}

