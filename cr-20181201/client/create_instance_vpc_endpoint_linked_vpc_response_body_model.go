// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceVpcEndpointLinkedVpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateInstanceVpcEndpointLinkedVpcResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *CreateInstanceVpcEndpointLinkedVpcResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *CreateInstanceVpcEndpointLinkedVpcResponseBody
	GetRequestId() *string
}

type CreateInstanceVpcEndpointLinkedVpcResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
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
	// D4978DCC-ECBD-40B0-A714-EE6959B22C77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceVpcEndpointLinkedVpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceVpcEndpointLinkedVpcResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceVpcEndpointLinkedVpcResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateInstanceVpcEndpointLinkedVpcResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *CreateInstanceVpcEndpointLinkedVpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceVpcEndpointLinkedVpcResponseBody) SetCode(v string) *CreateInstanceVpcEndpointLinkedVpcResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceVpcEndpointLinkedVpcResponseBody) SetIsSuccess(v bool) *CreateInstanceVpcEndpointLinkedVpcResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateInstanceVpcEndpointLinkedVpcResponseBody) SetRequestId(v string) *CreateInstanceVpcEndpointLinkedVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceVpcEndpointLinkedVpcResponseBody) Validate() error {
	return dara.Validate(s)
}
