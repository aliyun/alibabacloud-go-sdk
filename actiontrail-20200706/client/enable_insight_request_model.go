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
  // 	- IpInsight: Insights event on IP address
  // 
  // 	- ApiCallRateInsight: Insights event on API call rate
  // 
  // 	- ApiErrorRateInsight: Insights event on API error rate
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

