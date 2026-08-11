// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseOrgDeleteNodeResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseOrgDeleteNodeResponseBody
  GetCode() *string 
  SetData(v bool) *EnterpriseOrgDeleteNodeResponseBody
  GetData() *bool 
  SetMessage(v string) *EnterpriseOrgDeleteNodeResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnterpriseOrgDeleteNodeResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseOrgDeleteNodeResponseBody
  GetSuccess() *bool 
}

type EnterpriseOrgDeleteNodeResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseOrgDeleteNodeResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseOrgDeleteNodeResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseOrgDeleteNodeResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseOrgDeleteNodeResponseBody) GetData() *bool  {
  return s.Data
}

func (s *EnterpriseOrgDeleteNodeResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseOrgDeleteNodeResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseOrgDeleteNodeResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseOrgDeleteNodeResponseBody) SetCode(v string) *EnterpriseOrgDeleteNodeResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseOrgDeleteNodeResponseBody) SetData(v bool) *EnterpriseOrgDeleteNodeResponseBody {
  s.Data = &v
  return s
}

func (s *EnterpriseOrgDeleteNodeResponseBody) SetMessage(v string) *EnterpriseOrgDeleteNodeResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseOrgDeleteNodeResponseBody) SetRequestId(v string) *EnterpriseOrgDeleteNodeResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseOrgDeleteNodeResponseBody) SetSuccess(v bool) *EnterpriseOrgDeleteNodeResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseOrgDeleteNodeResponseBody) Validate() error {
  return dara.Validate(s)
}

