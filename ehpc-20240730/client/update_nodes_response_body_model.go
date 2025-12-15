// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateNodesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateNodesResponseBody
	GetSuccess() *bool
}

type UpdateNodesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNodesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateNodesResponseBody) SetRequestId(v string) *UpdateNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNodesResponseBody) SetSuccess(v bool) *UpdateNodesResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateNodesResponseBody) Validate() error {
	return dara.Validate(s)
}
