// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRepoBuildRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBuildRuleId(v string) *UpdateRepoBuildRuleResponseBody
	GetBuildRuleId() *string
	SetCode(v string) *UpdateRepoBuildRuleResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *UpdateRepoBuildRuleResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *UpdateRepoBuildRuleResponseBody
	GetRequestId() *string
}

type UpdateRepoBuildRuleResponseBody struct {
	// The ID of the building rule.
	//
	// example:
	//
	// crbr-ly77w5i3t31f****
	BuildRuleId *string `json:"BuildRuleId,omitempty" xml:"BuildRuleId,omitempty"`
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
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BC648259-91A7-4502-BED3-EDF64361FA83
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRepoBuildRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRepoBuildRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRepoBuildRuleResponseBody) GetBuildRuleId() *string {
	return s.BuildRuleId
}

func (s *UpdateRepoBuildRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateRepoBuildRuleResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *UpdateRepoBuildRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRepoBuildRuleResponseBody) SetBuildRuleId(v string) *UpdateRepoBuildRuleResponseBody {
	s.BuildRuleId = &v
	return s
}

func (s *UpdateRepoBuildRuleResponseBody) SetCode(v string) *UpdateRepoBuildRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRepoBuildRuleResponseBody) SetIsSuccess(v bool) *UpdateRepoBuildRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateRepoBuildRuleResponseBody) SetRequestId(v string) *UpdateRepoBuildRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRepoBuildRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
