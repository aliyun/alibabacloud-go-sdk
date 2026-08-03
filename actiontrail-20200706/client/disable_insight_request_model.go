// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableInsightRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInsightType(v string) *DisableInsightRequest
	GetInsightType() *string
}

type DisableInsightRequest struct {
	// The type of the Insights event. Valid values:
	//
	// - IpInsight: IP request events.
	//
	// - ApiCallRateInsight: High-risk API call events.
	//
	// - ApiErrorRateInsight: API error events.
	//
	// - AkInsight: AccessKey pair call events.
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

func (s DisableInsightRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableInsightRequest) GoString() string {
	return s.String()
}

func (s *DisableInsightRequest) GetInsightType() *string {
	return s.InsightType
}

func (s *DisableInsightRequest) SetInsightType(v string) *DisableInsightRequest {
	s.InsightType = &v
	return s
}

func (s *DisableInsightRequest) Validate() error {
	return dara.Validate(s)
}
