// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEmailVerifiedRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEmail(v string) *EmailVerifiedRequest
  GetEmail() *string 
  SetLang(v string) *EmailVerifiedRequest
  GetLang() *string 
  SetUserClientIp(v string) *EmailVerifiedRequest
  GetUserClientIp() *string 
}

type EmailVerifiedRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // abc@aliyun.com
  Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
  // example:
  // 
  // en
  Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
  // example:
  // 
  // 127.0.0.1
  UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s EmailVerifiedRequest) String() string {
  return dara.Prettify(s)
}

func (s EmailVerifiedRequest) GoString() string {
  return s.String()
}

func (s *EmailVerifiedRequest) GetEmail() *string  {
  return s.Email
}

func (s *EmailVerifiedRequest) GetLang() *string  {
  return s.Lang
}

func (s *EmailVerifiedRequest) GetUserClientIp() *string  {
  return s.UserClientIp
}

func (s *EmailVerifiedRequest) SetEmail(v string) *EmailVerifiedRequest {
  s.Email = &v
  return s
}

func (s *EmailVerifiedRequest) SetLang(v string) *EmailVerifiedRequest {
  s.Lang = &v
  return s
}

func (s *EmailVerifiedRequest) SetUserClientIp(v string) *EmailVerifiedRequest {
  s.UserClientIp = &v
  return s
}

func (s *EmailVerifiedRequest) Validate() error {
  return dara.Validate(s)
}

