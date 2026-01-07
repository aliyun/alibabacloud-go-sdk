// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountUpdateAccountAliasResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseAccountUpdateAccountAliasResponseBody
  GetCode() *string 
  SetData(v bool) *EnterpriseAccountUpdateAccountAliasResponseBody
  GetData() *bool 
  SetMessage(v string) *EnterpriseAccountUpdateAccountAliasResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnterpriseAccountUpdateAccountAliasResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseAccountUpdateAccountAliasResponseBody
  GetSuccess() *bool 
}

type EnterpriseAccountUpdateAccountAliasResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountUpdateAccountAliasResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountUpdateAccountAliasResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountUpdateAccountAliasResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseAccountUpdateAccountAliasResponseBody) GetData() *bool  {
  return s.Data
}

func (s *EnterpriseAccountUpdateAccountAliasResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseAccountUpdateAccountAliasResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseAccountUpdateAccountAliasResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseAccountUpdateAccountAliasResponseBody) SetCode(v string) *EnterpriseAccountUpdateAccountAliasResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseAccountUpdateAccountAliasResponseBody) SetData(v bool) *EnterpriseAccountUpdateAccountAliasResponseBody {
  s.Data = &v
  return s
}

func (s *EnterpriseAccountUpdateAccountAliasResponseBody) SetMessage(v string) *EnterpriseAccountUpdateAccountAliasResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseAccountUpdateAccountAliasResponseBody) SetRequestId(v string) *EnterpriseAccountUpdateAccountAliasResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseAccountUpdateAccountAliasResponseBody) SetSuccess(v bool) *EnterpriseAccountUpdateAccountAliasResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseAccountUpdateAccountAliasResponseBody) Validate() error {
  return dara.Validate(s)
}

