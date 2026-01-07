// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody
  GetCode() *string 
  SetMessage(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody
  GetMessage() *string 
  SetPass(v bool) *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody
  GetPass() *bool 
  SetRequestId(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody
  GetSuccess() *bool 
}

type EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  Pass *bool `json:"Pass,omitempty" xml:"Pass,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody) GetPass() *bool  {
  return s.Pass
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody) SetCode(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody) SetMessage(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody) SetPass(v bool) *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody {
  s.Pass = &v
  return s
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody) SetRequestId(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody) SetSuccess(v bool) *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody) Validate() error {
  return dara.Validate(s)
}

