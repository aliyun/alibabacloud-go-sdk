// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceVpcEndpointLinkedVpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteInstanceVpcEndpointLinkedVpcResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *DeleteInstanceVpcEndpointLinkedVpcResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteInstanceVpcEndpointLinkedVpcResponseBody
	GetRequestId() *string
}

type DeleteInstanceVpcEndpointLinkedVpcResponseBody struct {
	// The return value.
	//
	// example:
	//
	// true
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
	// 20FE7A66-0044-4E23-BBEC-C434EADBD7AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceVpcEndpointLinkedVpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceVpcEndpointLinkedVpcResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceVpcEndpointLinkedVpcResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteInstanceVpcEndpointLinkedVpcResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteInstanceVpcEndpointLinkedVpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInstanceVpcEndpointLinkedVpcResponseBody) SetCode(v string) *DeleteInstanceVpcEndpointLinkedVpcResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstanceVpcEndpointLinkedVpcResponseBody) SetIsSuccess(v bool) *DeleteInstanceVpcEndpointLinkedVpcResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteInstanceVpcEndpointLinkedVpcResponseBody) SetRequestId(v string) *DeleteInstanceVpcEndpointLinkedVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceVpcEndpointLinkedVpcResponseBody) Validate() error {
	return dara.Validate(s)
}
