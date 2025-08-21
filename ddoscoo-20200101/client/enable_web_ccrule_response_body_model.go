// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableWebCCRuleResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableWebCCRuleResponseBody
  GetRequestId() *string 
}

type EnableWebCCRuleResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableWebCCRuleResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableWebCCRuleResponseBody) GoString() string {
  return s.String()
}

func (s *EnableWebCCRuleResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableWebCCRuleResponseBody) SetRequestId(v string) *EnableWebCCRuleResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableWebCCRuleResponseBody) Validate() error {
  return dara.Validate(s)
}

