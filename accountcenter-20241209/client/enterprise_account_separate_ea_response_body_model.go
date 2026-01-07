// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountSeparateEaResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseAccountSeparateEaResponseBody
  GetCode() *string 
  SetMessage(v string) *EnterpriseAccountSeparateEaResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnterpriseAccountSeparateEaResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseAccountSeparateEaResponseBody
  GetSuccess() *bool 
}

type EnterpriseAccountSeparateEaResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountSeparateEaResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountSeparateEaResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountSeparateEaResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseAccountSeparateEaResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseAccountSeparateEaResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseAccountSeparateEaResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseAccountSeparateEaResponseBody) SetCode(v string) *EnterpriseAccountSeparateEaResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseAccountSeparateEaResponseBody) SetMessage(v string) *EnterpriseAccountSeparateEaResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseAccountSeparateEaResponseBody) SetRequestId(v string) *EnterpriseAccountSeparateEaResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseAccountSeparateEaResponseBody) SetSuccess(v bool) *EnterpriseAccountSeparateEaResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseAccountSeparateEaResponseBody) Validate() error {
  return dara.Validate(s)
}

