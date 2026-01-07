// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountUpdateOperateRiskControlResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseAccountUpdateOperateRiskControlResponseBody
  GetCode() *string 
  SetData(v bool) *EnterpriseAccountUpdateOperateRiskControlResponseBody
  GetData() *bool 
  SetMessage(v string) *EnterpriseAccountUpdateOperateRiskControlResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnterpriseAccountUpdateOperateRiskControlResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseAccountUpdateOperateRiskControlResponseBody
  GetSuccess() *bool 
}

type EnterpriseAccountUpdateOperateRiskControlResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountUpdateOperateRiskControlResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountUpdateOperateRiskControlResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountUpdateOperateRiskControlResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseAccountUpdateOperateRiskControlResponseBody) GetData() *bool  {
  return s.Data
}

func (s *EnterpriseAccountUpdateOperateRiskControlResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseAccountUpdateOperateRiskControlResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseAccountUpdateOperateRiskControlResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseAccountUpdateOperateRiskControlResponseBody) SetCode(v string) *EnterpriseAccountUpdateOperateRiskControlResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlResponseBody) SetData(v bool) *EnterpriseAccountUpdateOperateRiskControlResponseBody {
  s.Data = &v
  return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlResponseBody) SetMessage(v string) *EnterpriseAccountUpdateOperateRiskControlResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlResponseBody) SetRequestId(v string) *EnterpriseAccountUpdateOperateRiskControlResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlResponseBody) SetSuccess(v bool) *EnterpriseAccountUpdateOperateRiskControlResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlResponseBody) Validate() error {
  return dara.Validate(s)
}

