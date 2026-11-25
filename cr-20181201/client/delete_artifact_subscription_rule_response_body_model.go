// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteArtifactSubscriptionRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteArtifactSubscriptionRuleResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *DeleteArtifactSubscriptionRuleResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteArtifactSubscriptionRuleResponseBody
	GetRequestId() *string
}

type DeleteArtifactSubscriptionRuleResponseBody struct {
	// The return code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the API call was successful. Valid values:
	//
	// - `true`: The API call was successful.
	//
	// - `false`: The API call failed.
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

func (s DeleteArtifactSubscriptionRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteArtifactSubscriptionRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteArtifactSubscriptionRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteArtifactSubscriptionRuleResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteArtifactSubscriptionRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteArtifactSubscriptionRuleResponseBody) SetCode(v string) *DeleteArtifactSubscriptionRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteArtifactSubscriptionRuleResponseBody) SetIsSuccess(v bool) *DeleteArtifactSubscriptionRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteArtifactSubscriptionRuleResponseBody) SetRequestId(v string) *DeleteArtifactSubscriptionRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteArtifactSubscriptionRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
