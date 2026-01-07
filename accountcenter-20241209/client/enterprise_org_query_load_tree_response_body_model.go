// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseOrgQueryLoadTreeResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseOrgQueryLoadTreeResponseBody
  GetCode() *string 
  SetMessage(v string) *EnterpriseOrgQueryLoadTreeResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnterpriseOrgQueryLoadTreeResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseOrgQueryLoadTreeResponseBody
  GetSuccess() *bool 
  SetTreeDto(v string) *EnterpriseOrgQueryLoadTreeResponseBody
  GetTreeDto() *string 
}

type EnterpriseOrgQueryLoadTreeResponseBody struct {
  // example:
  // 
  // successful
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // Successful
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // A93073FC-1E86-58BA-AB83-54DA6BDB4F03
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // True
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
  TreeDto *string `json:"TreeDto,omitempty" xml:"TreeDto,omitempty"`
}

func (s EnterpriseOrgQueryLoadTreeResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseOrgQueryLoadTreeResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseOrgQueryLoadTreeResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseOrgQueryLoadTreeResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseOrgQueryLoadTreeResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseOrgQueryLoadTreeResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseOrgQueryLoadTreeResponseBody) GetTreeDto() *string  {
  return s.TreeDto
}

func (s *EnterpriseOrgQueryLoadTreeResponseBody) SetCode(v string) *EnterpriseOrgQueryLoadTreeResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseOrgQueryLoadTreeResponseBody) SetMessage(v string) *EnterpriseOrgQueryLoadTreeResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseOrgQueryLoadTreeResponseBody) SetRequestId(v string) *EnterpriseOrgQueryLoadTreeResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseOrgQueryLoadTreeResponseBody) SetSuccess(v bool) *EnterpriseOrgQueryLoadTreeResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseOrgQueryLoadTreeResponseBody) SetTreeDto(v string) *EnterpriseOrgQueryLoadTreeResponseBody {
  s.TreeDto = &v
  return s
}

func (s *EnterpriseOrgQueryLoadTreeResponseBody) Validate() error {
  return dara.Validate(s)
}

