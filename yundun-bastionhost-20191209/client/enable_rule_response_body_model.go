// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableRuleResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableRuleResponseBody
  GetRequestId() *string 
}

type EnableRuleResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 4F6C075F-FC86-476E-943B-097BD4E12948
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableRuleResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableRuleResponseBody) GoString() string {
  return s.String()
}

func (s *EnableRuleResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableRuleResponseBody) SetRequestId(v string) *EnableRuleResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableRuleResponseBody) Validate() error {
  return dara.Validate(s)
}

