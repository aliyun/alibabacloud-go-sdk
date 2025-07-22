// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAutoLiveStreamRuleResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableAutoLiveStreamRuleResponseBody
  GetRequestId() *string 
}

type EnableAutoLiveStreamRuleResponseBody struct {
  // example:
  // 
  // 760bad53276431c499e30dc36f6b26be
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableAutoLiveStreamRuleResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableAutoLiveStreamRuleResponseBody) GoString() string {
  return s.String()
}

func (s *EnableAutoLiveStreamRuleResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableAutoLiveStreamRuleResponseBody) SetRequestId(v string) *EnableAutoLiveStreamRuleResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableAutoLiveStreamRuleResponseBody) Validate() error {
  return dara.Validate(s)
}

