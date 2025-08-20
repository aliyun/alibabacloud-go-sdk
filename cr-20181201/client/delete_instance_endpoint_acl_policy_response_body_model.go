// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceEndpointAclPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteInstanceEndpointAclPolicyResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *DeleteInstanceEndpointAclPolicyResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteInstanceEndpointAclPolicyResponseBody
	GetRequestId() *string
}

type DeleteInstanceEndpointAclPolicyResponseBody struct {
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
	// BDB1F145-F0FF-44E9-AADF-A678642A7C7D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceEndpointAclPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceEndpointAclPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceEndpointAclPolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteInstanceEndpointAclPolicyResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteInstanceEndpointAclPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInstanceEndpointAclPolicyResponseBody) SetCode(v string) *DeleteInstanceEndpointAclPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstanceEndpointAclPolicyResponseBody) SetIsSuccess(v bool) *DeleteInstanceEndpointAclPolicyResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteInstanceEndpointAclPolicyResponseBody) SetRequestId(v string) *DeleteInstanceEndpointAclPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceEndpointAclPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
