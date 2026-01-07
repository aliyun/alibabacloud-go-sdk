// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseContactDeleteResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseContactDeleteResponseBody
  GetCode() *string 
  SetData(v bool) *EnterpriseContactDeleteResponseBody
  GetData() *bool 
  SetMessage(v string) *EnterpriseContactDeleteResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnterpriseContactDeleteResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseContactDeleteResponseBody
  GetSuccess() *bool 
}

type EnterpriseContactDeleteResponseBody struct {
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // true
  Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
  // example:
  // 
  // Success
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // 36E0A548-4BA3-549B-8855-22E3F5C6D47E
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // True
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseContactDeleteResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseContactDeleteResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseContactDeleteResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseContactDeleteResponseBody) GetData() *bool  {
  return s.Data
}

func (s *EnterpriseContactDeleteResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseContactDeleteResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseContactDeleteResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseContactDeleteResponseBody) SetCode(v string) *EnterpriseContactDeleteResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseContactDeleteResponseBody) SetData(v bool) *EnterpriseContactDeleteResponseBody {
  s.Data = &v
  return s
}

func (s *EnterpriseContactDeleteResponseBody) SetMessage(v string) *EnterpriseContactDeleteResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseContactDeleteResponseBody) SetRequestId(v string) *EnterpriseContactDeleteResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseContactDeleteResponseBody) SetSuccess(v bool) *EnterpriseContactDeleteResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseContactDeleteResponseBody) Validate() error {
  return dara.Validate(s)
}

