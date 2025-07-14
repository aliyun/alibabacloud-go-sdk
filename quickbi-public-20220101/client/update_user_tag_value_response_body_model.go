// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserTagValueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateUserTagValueResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateUserTagValueResponseBody
	GetResult() *bool
	SetSuccess(v bool) *UpdateUserTagValueResponseBody
	GetSuccess() *bool
}

type UpdateUserTagValueResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 46e5374665ba4b679ee22e2a29270
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns the result of modifying the user tag. Possible values:
	//
	// - true: Operation succeeded
	//
	// - false: Operation failed
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Value range:
	//
	// - true: The request was successful - false: The request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateUserTagValueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserTagValueResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserTagValueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserTagValueResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateUserTagValueResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateUserTagValueResponseBody) SetRequestId(v string) *UpdateUserTagValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserTagValueResponseBody) SetResult(v bool) *UpdateUserTagValueResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateUserTagValueResponseBody) SetSuccess(v bool) *UpdateUserTagValueResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateUserTagValueResponseBody) Validate() error {
	return dara.Validate(s)
}
