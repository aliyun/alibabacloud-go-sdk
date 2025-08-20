// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceEndpointStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateInstanceEndpointStatusResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *UpdateInstanceEndpointStatusResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *UpdateInstanceEndpointStatusResponseBody
	GetRequestId() *string
}

type UpdateInstanceEndpointStatusResponseBody struct {
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
	// 2FC14396-A16A-42BA-AAE4-BB94D956DF09
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceEndpointStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceEndpointStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceEndpointStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateInstanceEndpointStatusResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *UpdateInstanceEndpointStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceEndpointStatusResponseBody) SetCode(v string) *UpdateInstanceEndpointStatusResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceEndpointStatusResponseBody) SetIsSuccess(v bool) *UpdateInstanceEndpointStatusResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateInstanceEndpointStatusResponseBody) SetRequestId(v string) *UpdateInstanceEndpointStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceEndpointStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
