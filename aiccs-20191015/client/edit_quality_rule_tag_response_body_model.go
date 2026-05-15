// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditQualityRuleTagResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EditQualityRuleTagResponseBody
  GetCode() *string 
  SetData(v string) *EditQualityRuleTagResponseBody
  GetData() *string 
  SetMessage(v string) *EditQualityRuleTagResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EditQualityRuleTagResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EditQualityRuleTagResponseBody
  GetSuccess() *bool 
}

type EditQualityRuleTagResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EditQualityRuleTagResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EditQualityRuleTagResponseBody) GoString() string {
  return s.String()
}

func (s *EditQualityRuleTagResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EditQualityRuleTagResponseBody) GetData() *string  {
  return s.Data
}

func (s *EditQualityRuleTagResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EditQualityRuleTagResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EditQualityRuleTagResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EditQualityRuleTagResponseBody) SetCode(v string) *EditQualityRuleTagResponseBody {
  s.Code = &v
  return s
}

func (s *EditQualityRuleTagResponseBody) SetData(v string) *EditQualityRuleTagResponseBody {
  s.Data = &v
  return s
}

func (s *EditQualityRuleTagResponseBody) SetMessage(v string) *EditQualityRuleTagResponseBody {
  s.Message = &v
  return s
}

func (s *EditQualityRuleTagResponseBody) SetRequestId(v string) *EditQualityRuleTagResponseBody {
  s.RequestId = &v
  return s
}

func (s *EditQualityRuleTagResponseBody) SetSuccess(v bool) *EditQualityRuleTagResponseBody {
  s.Success = &v
  return s
}

func (s *EditQualityRuleTagResponseBody) Validate() error {
  return dara.Validate(s)
}

