// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateArtifactLifecycleRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateArtifactLifecycleRuleResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *CreateArtifactLifecycleRuleResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *CreateArtifactLifecycleRuleResponseBody
	GetRequestId() *string
	SetRuleId(v string) *CreateArtifactLifecycleRuleResponseBody
	GetRuleId() *string
}

type CreateArtifactLifecycleRuleResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// True
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0AA66379-B880-5123-9F6A-96BB25D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// cralr-b6thg027zmk1****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateArtifactLifecycleRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateArtifactLifecycleRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateArtifactLifecycleRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateArtifactLifecycleRuleResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *CreateArtifactLifecycleRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateArtifactLifecycleRuleResponseBody) GetRuleId() *string {
	return s.RuleId
}

func (s *CreateArtifactLifecycleRuleResponseBody) SetCode(v string) *CreateArtifactLifecycleRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateArtifactLifecycleRuleResponseBody) SetIsSuccess(v bool) *CreateArtifactLifecycleRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateArtifactLifecycleRuleResponseBody) SetRequestId(v string) *CreateArtifactLifecycleRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateArtifactLifecycleRuleResponseBody) SetRuleId(v string) *CreateArtifactLifecycleRuleResponseBody {
	s.RuleId = &v
	return s
}

func (s *CreateArtifactLifecycleRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
