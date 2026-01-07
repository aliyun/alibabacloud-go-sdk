// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseRoleDeleteBizRoleResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseRoleDeleteBizRoleResponseBody
  GetCode() *string 
  SetMessage(v string) *EnterpriseRoleDeleteBizRoleResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnterpriseRoleDeleteBizRoleResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseRoleDeleteBizRoleResponseBody
  GetSuccess() *bool 
}

type EnterpriseRoleDeleteBizRoleResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseRoleDeleteBizRoleResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseRoleDeleteBizRoleResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseRoleDeleteBizRoleResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseRoleDeleteBizRoleResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseRoleDeleteBizRoleResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseRoleDeleteBizRoleResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseRoleDeleteBizRoleResponseBody) SetCode(v string) *EnterpriseRoleDeleteBizRoleResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseRoleDeleteBizRoleResponseBody) SetMessage(v string) *EnterpriseRoleDeleteBizRoleResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseRoleDeleteBizRoleResponseBody) SetRequestId(v string) *EnterpriseRoleDeleteBizRoleResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseRoleDeleteBizRoleResponseBody) SetSuccess(v bool) *EnterpriseRoleDeleteBizRoleResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseRoleDeleteBizRoleResponseBody) Validate() error {
  return dara.Validate(s)
}

