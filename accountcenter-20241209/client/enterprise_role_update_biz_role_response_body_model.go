// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseRoleUpdateBizRoleResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseRoleUpdateBizRoleResponseBody
  GetCode() *string 
  SetMessage(v string) *EnterpriseRoleUpdateBizRoleResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnterpriseRoleUpdateBizRoleResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseRoleUpdateBizRoleResponseBody
  GetSuccess() *bool 
}

type EnterpriseRoleUpdateBizRoleResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseRoleUpdateBizRoleResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseRoleUpdateBizRoleResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseRoleUpdateBizRoleResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseRoleUpdateBizRoleResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseRoleUpdateBizRoleResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseRoleUpdateBizRoleResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseRoleUpdateBizRoleResponseBody) SetCode(v string) *EnterpriseRoleUpdateBizRoleResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseRoleUpdateBizRoleResponseBody) SetMessage(v string) *EnterpriseRoleUpdateBizRoleResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseRoleUpdateBizRoleResponseBody) SetRequestId(v string) *EnterpriseRoleUpdateBizRoleResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseRoleUpdateBizRoleResponseBody) SetSuccess(v bool) *EnterpriseRoleUpdateBizRoleResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseRoleUpdateBizRoleResponseBody) Validate() error {
  return dara.Validate(s)
}

