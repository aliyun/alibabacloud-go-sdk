// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEcologyOpennessAuthenticateRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEncodeKey(v string) *EcologyOpennessAuthenticateRequest
  GetEncodeKey() *string 
  SetEncodeType(v string) *EcologyOpennessAuthenticateRequest
  GetEncodeType() *string 
  SetLoginStateAccessToken(v string) *EcologyOpennessAuthenticateRequest
  GetLoginStateAccessToken() *string 
}

type EcologyOpennessAuthenticateRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // 12*****
  EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // PROJECT_ID
  EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // d15aa*****ee
  LoginStateAccessToken *string `json:"LoginStateAccessToken,omitempty" xml:"LoginStateAccessToken,omitempty"`
}

func (s EcologyOpennessAuthenticateRequest) String() string {
  return dara.Prettify(s)
}

func (s EcologyOpennessAuthenticateRequest) GoString() string {
  return s.String()
}

func (s *EcologyOpennessAuthenticateRequest) GetEncodeKey() *string  {
  return s.EncodeKey
}

func (s *EcologyOpennessAuthenticateRequest) GetEncodeType() *string  {
  return s.EncodeType
}

func (s *EcologyOpennessAuthenticateRequest) GetLoginStateAccessToken() *string  {
  return s.LoginStateAccessToken
}

func (s *EcologyOpennessAuthenticateRequest) SetEncodeKey(v string) *EcologyOpennessAuthenticateRequest {
  s.EncodeKey = &v
  return s
}

func (s *EcologyOpennessAuthenticateRequest) SetEncodeType(v string) *EcologyOpennessAuthenticateRequest {
  s.EncodeType = &v
  return s
}

func (s *EcologyOpennessAuthenticateRequest) SetLoginStateAccessToken(v string) *EcologyOpennessAuthenticateRequest {
  s.LoginStateAccessToken = &v
  return s
}

func (s *EcologyOpennessAuthenticateRequest) Validate() error {
  return dara.Validate(s)
}

