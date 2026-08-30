// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSaseUserTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateSaseUserTagResponseBody
	GetCode() *int32
	SetRequestId(v string) *UpdateSaseUserTagResponseBody
	GetRequestId() *string
	SetSuccess(v string) *UpdateSaseUserTagResponseBody
	GetSuccess() *string
}

type UpdateSaseUserTagResponseBody struct {
	// The HTTP status code or POP error code. Valid values:
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
	// 88285ACB-CE66-58A2-9283-0FD6B5E833BB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSaseUserTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSaseUserTagResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSaseUserTagResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateSaseUserTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSaseUserTagResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *UpdateSaseUserTagResponseBody) SetCode(v int32) *UpdateSaseUserTagResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSaseUserTagResponseBody) SetRequestId(v string) *UpdateSaseUserTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSaseUserTagResponseBody) SetSuccess(v string) *UpdateSaseUserTagResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSaseUserTagResponseBody) Validate() error {
	return dara.Validate(s)
}
