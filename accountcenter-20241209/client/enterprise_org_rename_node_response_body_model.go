// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseOrgRenameNodeResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseOrgRenameNodeResponseBody
  GetCode() *string 
  SetData(v bool) *EnterpriseOrgRenameNodeResponseBody
  GetData() *bool 
  SetMessage(v string) *EnterpriseOrgRenameNodeResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnterpriseOrgRenameNodeResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseOrgRenameNodeResponseBody
  GetSuccess() *bool 
}

type EnterpriseOrgRenameNodeResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseOrgRenameNodeResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseOrgRenameNodeResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseOrgRenameNodeResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseOrgRenameNodeResponseBody) GetData() *bool  {
  return s.Data
}

func (s *EnterpriseOrgRenameNodeResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseOrgRenameNodeResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseOrgRenameNodeResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseOrgRenameNodeResponseBody) SetCode(v string) *EnterpriseOrgRenameNodeResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseOrgRenameNodeResponseBody) SetData(v bool) *EnterpriseOrgRenameNodeResponseBody {
  s.Data = &v
  return s
}

func (s *EnterpriseOrgRenameNodeResponseBody) SetMessage(v string) *EnterpriseOrgRenameNodeResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseOrgRenameNodeResponseBody) SetRequestId(v string) *EnterpriseOrgRenameNodeResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseOrgRenameNodeResponseBody) SetSuccess(v bool) *EnterpriseOrgRenameNodeResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseOrgRenameNodeResponseBody) Validate() error {
  return dara.Validate(s)
}

