// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateInstanceNetworkResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateInstanceNetworkResponseBody
	GetSuccess() *bool
}

type UpdateInstanceNetworkResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// DSSDF-SEWE-*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInstanceNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceNetworkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateInstanceNetworkResponseBody) SetRequestId(v string) *UpdateInstanceNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceNetworkResponseBody) SetSuccess(v bool) *UpdateInstanceNetworkResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateInstanceNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}
