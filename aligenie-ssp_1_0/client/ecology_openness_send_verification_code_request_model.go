// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEcologyOpennessSendVerificationCodeRequest interface {
  dara.Model
  String() string
  GoString() string
  SetPhoneNumber(v string) *EcologyOpennessSendVerificationCodeRequest
  GetPhoneNumber() *string 
  SetRegion(v string) *EcologyOpennessSendVerificationCodeRequest
  GetRegion() *string 
  SetSessionId(v string) *EcologyOpennessSendVerificationCodeRequest
  GetSessionId() *string 
}

type EcologyOpennessSendVerificationCodeRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // 18612345678
  PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // +86
  Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // dbe2eb4458302b9246c6da17fbc95f4b
  SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s EcologyOpennessSendVerificationCodeRequest) String() string {
  return dara.Prettify(s)
}

func (s EcologyOpennessSendVerificationCodeRequest) GoString() string {
  return s.String()
}

func (s *EcologyOpennessSendVerificationCodeRequest) GetPhoneNumber() *string  {
  return s.PhoneNumber
}

func (s *EcologyOpennessSendVerificationCodeRequest) GetRegion() *string  {
  return s.Region
}

func (s *EcologyOpennessSendVerificationCodeRequest) GetSessionId() *string  {
  return s.SessionId
}

func (s *EcologyOpennessSendVerificationCodeRequest) SetPhoneNumber(v string) *EcologyOpennessSendVerificationCodeRequest {
  s.PhoneNumber = &v
  return s
}

func (s *EcologyOpennessSendVerificationCodeRequest) SetRegion(v string) *EcologyOpennessSendVerificationCodeRequest {
  s.Region = &v
  return s
}

func (s *EcologyOpennessSendVerificationCodeRequest) SetSessionId(v string) *EcologyOpennessSendVerificationCodeRequest {
  s.SessionId = &v
  return s
}

func (s *EcologyOpennessSendVerificationCodeRequest) Validate() error {
  return dara.Validate(s)
}

