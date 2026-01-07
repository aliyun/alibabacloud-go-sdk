// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountRemoveMfaResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseAccountRemoveMfaResponseBody
  GetCode() *string 
  SetData(v bool) *EnterpriseAccountRemoveMfaResponseBody
  GetData() *bool 
  SetMessage(v string) *EnterpriseAccountRemoveMfaResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnterpriseAccountRemoveMfaResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseAccountRemoveMfaResponseBody
  GetSuccess() *bool 
}

type EnterpriseAccountRemoveMfaResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountRemoveMfaResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountRemoveMfaResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountRemoveMfaResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseAccountRemoveMfaResponseBody) GetData() *bool  {
  return s.Data
}

func (s *EnterpriseAccountRemoveMfaResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseAccountRemoveMfaResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseAccountRemoveMfaResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseAccountRemoveMfaResponseBody) SetCode(v string) *EnterpriseAccountRemoveMfaResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseAccountRemoveMfaResponseBody) SetData(v bool) *EnterpriseAccountRemoveMfaResponseBody {
  s.Data = &v
  return s
}

func (s *EnterpriseAccountRemoveMfaResponseBody) SetMessage(v string) *EnterpriseAccountRemoveMfaResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseAccountRemoveMfaResponseBody) SetRequestId(v string) *EnterpriseAccountRemoveMfaResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseAccountRemoveMfaResponseBody) SetSuccess(v bool) *EnterpriseAccountRemoveMfaResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseAccountRemoveMfaResponseBody) Validate() error {
  return dara.Validate(s)
}

