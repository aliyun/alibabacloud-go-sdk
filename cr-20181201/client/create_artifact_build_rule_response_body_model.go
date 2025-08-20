// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateArtifactBuildRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBuildRuleId(v string) *CreateArtifactBuildRuleResponseBody
	GetBuildRuleId() *string
	SetCode(v string) *CreateArtifactBuildRuleResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *CreateArtifactBuildRuleResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *CreateArtifactBuildRuleResponseBody
	GetRequestId() *string
}

type CreateArtifactBuildRuleResponseBody struct {
	// The ID of the accelerated image building rule.
	//
	// example:
	//
	// crabr-7dfa5qye5****
	BuildRuleId *string `json:"BuildRuleId,omitempty" xml:"BuildRuleId,omitempty"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C8E90AB5-0A96-5D12-9E59-11EE46360642
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateArtifactBuildRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateArtifactBuildRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateArtifactBuildRuleResponseBody) GetBuildRuleId() *string {
	return s.BuildRuleId
}

func (s *CreateArtifactBuildRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateArtifactBuildRuleResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *CreateArtifactBuildRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateArtifactBuildRuleResponseBody) SetBuildRuleId(v string) *CreateArtifactBuildRuleResponseBody {
	s.BuildRuleId = &v
	return s
}

func (s *CreateArtifactBuildRuleResponseBody) SetCode(v string) *CreateArtifactBuildRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateArtifactBuildRuleResponseBody) SetIsSuccess(v bool) *CreateArtifactBuildRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateArtifactBuildRuleResponseBody) SetRequestId(v string) *CreateArtifactBuildRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateArtifactBuildRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
