// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUserResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteUserResponseBody
	GetResult() *bool
	SetSuccess(v bool) *DeleteUserResponseBody
	GetSuccess() *bool
}

type DeleteUserResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// DC4E1E63-B337-44F8-8C22-6F00DF67E2C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns the execution result of the interface. Possible values:
	//
	// - true: Execution succeeded
	//
	// - false: Execution failed
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Possible values:
	//
	// - true: The request was successful - false: The request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteUserResponseBody) SetRequestId(v string) *DeleteUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserResponseBody) SetResult(v bool) *DeleteUserResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteUserResponseBody) SetSuccess(v bool) *DeleteUserResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteUserResponseBody) Validate() error {
	return dara.Validate(s)
}
