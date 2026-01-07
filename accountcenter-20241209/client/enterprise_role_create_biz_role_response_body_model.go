// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseRoleCreateBizRoleResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetBizRoleCode(v string) *EnterpriseRoleCreateBizRoleResponseBody
  GetBizRoleCode() *string 
  SetCode(v string) *EnterpriseRoleCreateBizRoleResponseBody
  GetCode() *string 
  SetMessage(v string) *EnterpriseRoleCreateBizRoleResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnterpriseRoleCreateBizRoleResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseRoleCreateBizRoleResponseBody
  GetSuccess() *bool 
}

type EnterpriseRoleCreateBizRoleResponseBody struct {
  BizRoleCode *string `json:"BizRoleCode,omitempty" xml:"BizRoleCode,omitempty"`
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseRoleCreateBizRoleResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseRoleCreateBizRoleResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseRoleCreateBizRoleResponseBody) GetBizRoleCode() *string  {
  return s.BizRoleCode
}

func (s *EnterpriseRoleCreateBizRoleResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseRoleCreateBizRoleResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseRoleCreateBizRoleResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseRoleCreateBizRoleResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseRoleCreateBizRoleResponseBody) SetBizRoleCode(v string) *EnterpriseRoleCreateBizRoleResponseBody {
  s.BizRoleCode = &v
  return s
}

func (s *EnterpriseRoleCreateBizRoleResponseBody) SetCode(v string) *EnterpriseRoleCreateBizRoleResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseRoleCreateBizRoleResponseBody) SetMessage(v string) *EnterpriseRoleCreateBizRoleResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseRoleCreateBizRoleResponseBody) SetRequestId(v string) *EnterpriseRoleCreateBizRoleResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseRoleCreateBizRoleResponseBody) SetSuccess(v bool) *EnterpriseRoleCreateBizRoleResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseRoleCreateBizRoleResponseBody) Validate() error {
  return dara.Validate(s)
}

