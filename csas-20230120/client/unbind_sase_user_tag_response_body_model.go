// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindSaseUserTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UnbindSaseUserTagResponseBody
	GetCode() *int32
	SetRequestId(v string) *UnbindSaseUserTagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UnbindSaseUserTagResponseBody
	GetSuccess() *bool
}

type UnbindSaseUserTagResponseBody struct {
	// The API status code or POP error code. Valid values:
	//
	// - **2xx**: Success.
	//
	// - **3xx**: Redirection.
	//
	// - **4xx**: Request error.
	//
	// - **5xx**: Server error.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F4A9C844-1B0A-59E8-966F-4945DFF3C88D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnbindSaseUserTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindSaseUserTagResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindSaseUserTagResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UnbindSaseUserTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindSaseUserTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UnbindSaseUserTagResponseBody) SetCode(v int32) *UnbindSaseUserTagResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindSaseUserTagResponseBody) SetRequestId(v string) *UnbindSaseUserTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindSaseUserTagResponseBody) SetSuccess(v bool) *UnbindSaseUserTagResponseBody {
	s.Success = &v
	return s
}

func (s *UnbindSaseUserTagResponseBody) Validate() error {
	return dara.Validate(s)
}
