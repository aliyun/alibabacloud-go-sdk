// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountUpdateSessionExpireTimeResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseAccountUpdateSessionExpireTimeResponseBody
  GetCode() *string 
  SetData(v bool) *EnterpriseAccountUpdateSessionExpireTimeResponseBody
  GetData() *bool 
  SetMessage(v string) *EnterpriseAccountUpdateSessionExpireTimeResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnterpriseAccountUpdateSessionExpireTimeResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseAccountUpdateSessionExpireTimeResponseBody
  GetSuccess() *bool 
}

type EnterpriseAccountUpdateSessionExpireTimeResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountUpdateSessionExpireTimeResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountUpdateSessionExpireTimeResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountUpdateSessionExpireTimeResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseAccountUpdateSessionExpireTimeResponseBody) GetData() *bool  {
  return s.Data
}

func (s *EnterpriseAccountUpdateSessionExpireTimeResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseAccountUpdateSessionExpireTimeResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseAccountUpdateSessionExpireTimeResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseAccountUpdateSessionExpireTimeResponseBody) SetCode(v string) *EnterpriseAccountUpdateSessionExpireTimeResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseAccountUpdateSessionExpireTimeResponseBody) SetData(v bool) *EnterpriseAccountUpdateSessionExpireTimeResponseBody {
  s.Data = &v
  return s
}

func (s *EnterpriseAccountUpdateSessionExpireTimeResponseBody) SetMessage(v string) *EnterpriseAccountUpdateSessionExpireTimeResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseAccountUpdateSessionExpireTimeResponseBody) SetRequestId(v string) *EnterpriseAccountUpdateSessionExpireTimeResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseAccountUpdateSessionExpireTimeResponseBody) SetSuccess(v bool) *EnterpriseAccountUpdateSessionExpireTimeResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseAccountUpdateSessionExpireTimeResponseBody) Validate() error {
  return dara.Validate(s)
}

