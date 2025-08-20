// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBuildRecordByRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBuildRecordId(v string) *CreateBuildRecordByRuleResponseBody
	GetBuildRecordId() *string
	SetCode(v string) *CreateBuildRecordByRuleResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *CreateBuildRecordByRuleResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *CreateBuildRecordByRuleResponseBody
	GetRequestId() *string
}

type CreateBuildRecordByRuleResponseBody struct {
	// The ID of the image building record.
	//
	// example:
	//
	// 0A311FC5-B8C6-4332-80E4-539EB73****
	BuildRecordId *string `json:"BuildRecordId,omitempty" xml:"BuildRecordId,omitempty"`
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
	// B01B8857-A16E-40E9-A37E-764F15776FAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBuildRecordByRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBuildRecordByRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBuildRecordByRuleResponseBody) GetBuildRecordId() *string {
	return s.BuildRecordId
}

func (s *CreateBuildRecordByRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateBuildRecordByRuleResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *CreateBuildRecordByRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBuildRecordByRuleResponseBody) SetBuildRecordId(v string) *CreateBuildRecordByRuleResponseBody {
	s.BuildRecordId = &v
	return s
}

func (s *CreateBuildRecordByRuleResponseBody) SetCode(v string) *CreateBuildRecordByRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBuildRecordByRuleResponseBody) SetIsSuccess(v bool) *CreateBuildRecordByRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateBuildRecordByRuleResponseBody) SetRequestId(v string) *CreateBuildRecordByRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBuildRecordByRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
