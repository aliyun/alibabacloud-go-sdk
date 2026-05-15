// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditQualityRuleResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EditQualityRuleResponseBody
  GetCode() *string 
  SetMessage(v string) *EditQualityRuleResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EditQualityRuleResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EditQualityRuleResponseBody
  GetSuccess() *bool 
}

type EditQualityRuleResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EditQualityRuleResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EditQualityRuleResponseBody) GoString() string {
  return s.String()
}

func (s *EditQualityRuleResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EditQualityRuleResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EditQualityRuleResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EditQualityRuleResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EditQualityRuleResponseBody) SetCode(v string) *EditQualityRuleResponseBody {
  s.Code = &v
  return s
}

func (s *EditQualityRuleResponseBody) SetMessage(v string) *EditQualityRuleResponseBody {
  s.Message = &v
  return s
}

func (s *EditQualityRuleResponseBody) SetRequestId(v string) *EditQualityRuleResponseBody {
  s.RequestId = &v
  return s
}

func (s *EditQualityRuleResponseBody) SetSuccess(v bool) *EditQualityRuleResponseBody {
  s.Success = &v
  return s
}

func (s *EditQualityRuleResponseBody) Validate() error {
  return dara.Validate(s)
}

