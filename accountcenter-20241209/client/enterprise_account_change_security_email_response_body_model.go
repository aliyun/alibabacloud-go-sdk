// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountChangeSecurityEmailResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseAccountChangeSecurityEmailResponseBody
  GetCode() *string 
  SetData(v bool) *EnterpriseAccountChangeSecurityEmailResponseBody
  GetData() *bool 
  SetMessage(v string) *EnterpriseAccountChangeSecurityEmailResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnterpriseAccountChangeSecurityEmailResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseAccountChangeSecurityEmailResponseBody
  GetSuccess() *bool 
}

type EnterpriseAccountChangeSecurityEmailResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountChangeSecurityEmailResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountChangeSecurityEmailResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountChangeSecurityEmailResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseAccountChangeSecurityEmailResponseBody) GetData() *bool  {
  return s.Data
}

func (s *EnterpriseAccountChangeSecurityEmailResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseAccountChangeSecurityEmailResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseAccountChangeSecurityEmailResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseAccountChangeSecurityEmailResponseBody) SetCode(v string) *EnterpriseAccountChangeSecurityEmailResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityEmailResponseBody) SetData(v bool) *EnterpriseAccountChangeSecurityEmailResponseBody {
  s.Data = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityEmailResponseBody) SetMessage(v string) *EnterpriseAccountChangeSecurityEmailResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityEmailResponseBody) SetRequestId(v string) *EnterpriseAccountChangeSecurityEmailResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityEmailResponseBody) SetSuccess(v bool) *EnterpriseAccountChangeSecurityEmailResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityEmailResponseBody) Validate() error {
  return dara.Validate(s)
}

