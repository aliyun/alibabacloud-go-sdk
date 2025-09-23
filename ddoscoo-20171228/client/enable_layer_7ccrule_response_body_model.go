// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableLayer7CCRuleResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableLayer7CCRuleResponseBody
  GetRequestId() *string 
}

type EnableLayer7CCRuleResponseBody struct {
  // example:
  // 
  // CF33B4C3-196E-4015-AADD-5CAD00057B80
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableLayer7CCRuleResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableLayer7CCRuleResponseBody) GoString() string {
  return s.String()
}

func (s *EnableLayer7CCRuleResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableLayer7CCRuleResponseBody) SetRequestId(v string) *EnableLayer7CCRuleResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableLayer7CCRuleResponseBody) Validate() error {
  return dara.Validate(s)
}

