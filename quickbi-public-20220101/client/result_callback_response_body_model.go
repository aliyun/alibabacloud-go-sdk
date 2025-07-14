// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResultCallbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResultCallbackResponseBody
	GetRequestId() *string
	SetResult(v bool) *ResultCallbackResponseBody
	GetResult() *bool
	SetSuccess(v bool) *ResultCallbackResponseBody
	GetSuccess() *bool
}

type ResultCallbackResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution result of the interface is returned. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request fails.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:
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

func (s ResultCallbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResultCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *ResultCallbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResultCallbackResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ResultCallbackResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ResultCallbackResponseBody) SetRequestId(v string) *ResultCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResultCallbackResponseBody) SetResult(v bool) *ResultCallbackResponseBody {
	s.Result = &v
	return s
}

func (s *ResultCallbackResponseBody) SetSuccess(v bool) *ResultCallbackResponseBody {
	s.Success = &v
	return s
}

func (s *ResultCallbackResponseBody) Validate() error {
	return dara.Validate(s)
}
