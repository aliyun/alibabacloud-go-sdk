// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateArtifactSubscriptionRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateArtifactSubscriptionRuleResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *CreateArtifactSubscriptionRuleResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *CreateArtifactSubscriptionRuleResponseBody
	GetRequestId() *string
	SetRuleId(v string) *CreateArtifactSubscriptionRuleResponseBody
	GetRuleId() *string
}

type CreateArtifactSubscriptionRuleResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicate whether the API request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 02B27D80-FD32-5155-931A-93700779BB9E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the artifact subscription rule.
	//
	// example:
	//
	// crasr-lxdfele7dg4****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateArtifactSubscriptionRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateArtifactSubscriptionRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateArtifactSubscriptionRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateArtifactSubscriptionRuleResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *CreateArtifactSubscriptionRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateArtifactSubscriptionRuleResponseBody) GetRuleId() *string {
	return s.RuleId
}

func (s *CreateArtifactSubscriptionRuleResponseBody) SetCode(v string) *CreateArtifactSubscriptionRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleResponseBody) SetIsSuccess(v bool) *CreateArtifactSubscriptionRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleResponseBody) SetRequestId(v string) *CreateArtifactSubscriptionRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleResponseBody) SetRuleId(v string) *CreateArtifactSubscriptionRuleResponseBody {
	s.RuleId = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
