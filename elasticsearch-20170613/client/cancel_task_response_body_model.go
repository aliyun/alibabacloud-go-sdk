// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelTaskResponseBody
	GetRequestId() *string
	SetResult(v bool) *CancelTaskResponseBody
	GetResult() *bool
}

type CancelTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return results:
	//
	// 	- true: the task was cancelled successfully
	//
	// 	- false: the task was cancelled successfully failed
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CancelTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelTaskResponseBody) GetResult() *bool {
	return s.Result
}

func (s *CancelTaskResponseBody) SetRequestId(v string) *CancelTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelTaskResponseBody) SetResult(v bool) *CancelTaskResponseBody {
	s.Result = &v
	return s
}

func (s *CancelTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
