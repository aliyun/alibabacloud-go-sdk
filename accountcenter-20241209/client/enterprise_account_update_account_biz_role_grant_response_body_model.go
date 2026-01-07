// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountUpdateAccountBizRoleGrantResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody
  GetCode() *string 
  SetMessage(v string) *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody
  GetSuccess() *bool 
}

type EnterpriseAccountUpdateAccountBizRoleGrantResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountUpdateAccountBizRoleGrantResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountUpdateAccountBizRoleGrantResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody) SetCode(v string) *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody) SetMessage(v string) *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody) SetRequestId(v string) *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody) SetSuccess(v bool) *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody) Validate() error {
  return dara.Validate(s)
}

