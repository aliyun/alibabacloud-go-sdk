// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceEndpointAclPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateInstanceEndpointAclPolicyResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *CreateInstanceEndpointAclPolicyResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *CreateInstanceEndpointAclPolicyResponseBody
	GetRequestId() *string
}

type CreateInstanceEndpointAclPolicyResponseBody struct {
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
	// D735C5EC-4206-4F48-A090-307BF56BEB99
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceEndpointAclPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceEndpointAclPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceEndpointAclPolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateInstanceEndpointAclPolicyResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *CreateInstanceEndpointAclPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceEndpointAclPolicyResponseBody) SetCode(v string) *CreateInstanceEndpointAclPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceEndpointAclPolicyResponseBody) SetIsSuccess(v bool) *CreateInstanceEndpointAclPolicyResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateInstanceEndpointAclPolicyResponseBody) SetRequestId(v string) *CreateInstanceEndpointAclPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceEndpointAclPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
