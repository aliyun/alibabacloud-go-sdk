// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountUpdateIpMaskResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseAccountUpdateIpMaskResponseBody
  GetCode() *string 
  SetData(v bool) *EnterpriseAccountUpdateIpMaskResponseBody
  GetData() *bool 
  SetMessage(v string) *EnterpriseAccountUpdateIpMaskResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnterpriseAccountUpdateIpMaskResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseAccountUpdateIpMaskResponseBody
  GetSuccess() *bool 
}

type EnterpriseAccountUpdateIpMaskResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountUpdateIpMaskResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountUpdateIpMaskResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountUpdateIpMaskResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseAccountUpdateIpMaskResponseBody) GetData() *bool  {
  return s.Data
}

func (s *EnterpriseAccountUpdateIpMaskResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseAccountUpdateIpMaskResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseAccountUpdateIpMaskResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseAccountUpdateIpMaskResponseBody) SetCode(v string) *EnterpriseAccountUpdateIpMaskResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseAccountUpdateIpMaskResponseBody) SetData(v bool) *EnterpriseAccountUpdateIpMaskResponseBody {
  s.Data = &v
  return s
}

func (s *EnterpriseAccountUpdateIpMaskResponseBody) SetMessage(v string) *EnterpriseAccountUpdateIpMaskResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseAccountUpdateIpMaskResponseBody) SetRequestId(v string) *EnterpriseAccountUpdateIpMaskResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseAccountUpdateIpMaskResponseBody) SetSuccess(v bool) *EnterpriseAccountUpdateIpMaskResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseAccountUpdateIpMaskResponseBody) Validate() error {
  return dara.Validate(s)
}

