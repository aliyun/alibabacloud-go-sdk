// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableInsightRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInsightType(v string) *EnableInsightRequest
  GetInsightType() *string 
}

type EnableInsightRequest struct {
  // The type of the Insights event. Valid values:
  // 
  // - IpInsight: IP address request events.
  // 
  // - ApiCallRateInsight: Unusual API call events.
  // 
  // - ApiErrorRateInsight: API error events.
  // 
  // - AkInsight: Unusual AccessKey pair call events.
  // 
  // - PolicyChangeInsight: Permission change events.
  // 
  // - PasswordChangeInsight: Password change events.
  // 
  // - TrailConcealmentInsight: Trail concealment events.
  // 
  // example:
  // 
  // IpInsight
  InsightType *string `json:"InsightType,omitempty" xml:"InsightType,omitempty"`
}

func (s EnableInsightRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableInsightRequest) GoString() string {
  return s.String()
}

func (s *EnableInsightRequest) GetInsightType() *string  {
  return s.InsightType
}

func (s *EnableInsightRequest) SetInsightType(v string) *EnableInsightRequest {
  s.InsightType = &v
  return s
}

func (s *EnableInsightRequest) Validate() error {
  return dara.Validate(s)
}

