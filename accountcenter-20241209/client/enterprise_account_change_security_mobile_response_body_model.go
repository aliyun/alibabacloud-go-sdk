// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountChangeSecurityMobileResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseAccountChangeSecurityMobileResponseBody
  GetCode() *string 
  SetData(v bool) *EnterpriseAccountChangeSecurityMobileResponseBody
  GetData() *bool 
  SetMessage(v string) *EnterpriseAccountChangeSecurityMobileResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnterpriseAccountChangeSecurityMobileResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseAccountChangeSecurityMobileResponseBody
  GetSuccess() *bool 
}

type EnterpriseAccountChangeSecurityMobileResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountChangeSecurityMobileResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountChangeSecurityMobileResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountChangeSecurityMobileResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseAccountChangeSecurityMobileResponseBody) GetData() *bool  {
  return s.Data
}

func (s *EnterpriseAccountChangeSecurityMobileResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseAccountChangeSecurityMobileResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseAccountChangeSecurityMobileResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseAccountChangeSecurityMobileResponseBody) SetCode(v string) *EnterpriseAccountChangeSecurityMobileResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityMobileResponseBody) SetData(v bool) *EnterpriseAccountChangeSecurityMobileResponseBody {
  s.Data = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityMobileResponseBody) SetMessage(v string) *EnterpriseAccountChangeSecurityMobileResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityMobileResponseBody) SetRequestId(v string) *EnterpriseAccountChangeSecurityMobileResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityMobileResponseBody) SetSuccess(v bool) *EnterpriseAccountChangeSecurityMobileResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityMobileResponseBody) Validate() error {
  return dara.Validate(s)
}

