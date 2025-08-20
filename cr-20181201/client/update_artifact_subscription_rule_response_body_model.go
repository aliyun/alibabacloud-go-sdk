// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateArtifactSubscriptionRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateArtifactSubscriptionRuleResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *UpdateArtifactSubscriptionRuleResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *UpdateArtifactSubscriptionRuleResponseBody
	GetRequestId() *string
}

type UpdateArtifactSubscriptionRuleResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the API request is successful. Valid values:
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
}

func (s UpdateArtifactSubscriptionRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateArtifactSubscriptionRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateArtifactSubscriptionRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateArtifactSubscriptionRuleResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *UpdateArtifactSubscriptionRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateArtifactSubscriptionRuleResponseBody) SetCode(v string) *UpdateArtifactSubscriptionRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleResponseBody) SetIsSuccess(v bool) *UpdateArtifactSubscriptionRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleResponseBody) SetRequestId(v string) *UpdateArtifactSubscriptionRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
