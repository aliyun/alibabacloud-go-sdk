// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFunctionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFunctionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteFunctionResponseBody
	GetSuccess() *bool
}

type DeleteFunctionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 88198F19-A36B-52A9-AE44-4518A688XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFunctionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFunctionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFunctionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteFunctionResponseBody) SetRequestId(v string) *DeleteFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFunctionResponseBody) SetSuccess(v bool) *DeleteFunctionResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFunctionResponseBody) Validate() error {
	return dara.Validate(s)
}
