// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountChangeLoginPasswordResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseAccountChangeLoginPasswordResponseBody
  GetCode() *string 
  SetMessage(v string) *EnterpriseAccountChangeLoginPasswordResponseBody
  GetMessage() *string 
  SetPass(v bool) *EnterpriseAccountChangeLoginPasswordResponseBody
  GetPass() *bool 
  SetRequestId(v string) *EnterpriseAccountChangeLoginPasswordResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseAccountChangeLoginPasswordResponseBody
  GetSuccess() *bool 
}

type EnterpriseAccountChangeLoginPasswordResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  Pass *bool `json:"Pass,omitempty" xml:"Pass,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountChangeLoginPasswordResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountChangeLoginPasswordResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountChangeLoginPasswordResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseAccountChangeLoginPasswordResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseAccountChangeLoginPasswordResponseBody) GetPass() *bool  {
  return s.Pass
}

func (s *EnterpriseAccountChangeLoginPasswordResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseAccountChangeLoginPasswordResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseAccountChangeLoginPasswordResponseBody) SetCode(v string) *EnterpriseAccountChangeLoginPasswordResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseAccountChangeLoginPasswordResponseBody) SetMessage(v string) *EnterpriseAccountChangeLoginPasswordResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseAccountChangeLoginPasswordResponseBody) SetPass(v bool) *EnterpriseAccountChangeLoginPasswordResponseBody {
  s.Pass = &v
  return s
}

func (s *EnterpriseAccountChangeLoginPasswordResponseBody) SetRequestId(v string) *EnterpriseAccountChangeLoginPasswordResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseAccountChangeLoginPasswordResponseBody) SetSuccess(v bool) *EnterpriseAccountChangeLoginPasswordResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseAccountChangeLoginPasswordResponseBody) Validate() error {
  return dara.Validate(s)
}

