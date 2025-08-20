// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateArtifactLifecycleRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateArtifactLifecycleRuleResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *UpdateArtifactLifecycleRuleResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *UpdateArtifactLifecycleRuleResponseBody
	GetRequestId() *string
}

type UpdateArtifactLifecycleRuleResponseBody struct {
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
	// BF92FC2E-455F-5600-A276-D2150A59****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateArtifactLifecycleRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateArtifactLifecycleRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateArtifactLifecycleRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateArtifactLifecycleRuleResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *UpdateArtifactLifecycleRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateArtifactLifecycleRuleResponseBody) SetCode(v string) *UpdateArtifactLifecycleRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleResponseBody) SetIsSuccess(v bool) *UpdateArtifactLifecycleRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleResponseBody) SetRequestId(v string) *UpdateArtifactLifecycleRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
