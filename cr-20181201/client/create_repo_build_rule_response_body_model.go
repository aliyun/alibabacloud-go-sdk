// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepoBuildRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBuildRuleId(v string) *CreateRepoBuildRuleResponseBody
	GetBuildRuleId() *string
	SetCode(v string) *CreateRepoBuildRuleResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *CreateRepoBuildRuleResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *CreateRepoBuildRuleResponseBody
	GetRequestId() *string
}

type CreateRepoBuildRuleResponseBody struct {
	// The build rule ID.
	//
	// example:
	//
	// crbr-ly77w5i3t31f****
	BuildRuleId *string `json:"BuildRuleId,omitempty" xml:"BuildRuleId,omitempty"`
	// The return code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - `true`: The call is successful.
	//
	// - `false`: The call failed.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4CE1F661-75DD-4EBD-A4AD-057B26834ABB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRepoBuildRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRepoBuildRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepoBuildRuleResponseBody) GetBuildRuleId() *string {
	return s.BuildRuleId
}

func (s *CreateRepoBuildRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateRepoBuildRuleResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *CreateRepoBuildRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRepoBuildRuleResponseBody) SetBuildRuleId(v string) *CreateRepoBuildRuleResponseBody {
	s.BuildRuleId = &v
	return s
}

func (s *CreateRepoBuildRuleResponseBody) SetCode(v string) *CreateRepoBuildRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRepoBuildRuleResponseBody) SetIsSuccess(v bool) *CreateRepoBuildRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateRepoBuildRuleResponseBody) SetRequestId(v string) *CreateRepoBuildRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRepoBuildRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
