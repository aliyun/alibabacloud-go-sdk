// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSaseUserTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteSaseUserTagResponseBody
	GetCode() *int32
	SetRequestId(v string) *DeleteSaseUserTagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSaseUserTagResponseBody
	GetSuccess() *bool
}

type DeleteSaseUserTagResponseBody struct {
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
	// 54A4055A-343D-583E-9EAC-D12231148A68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSaseUserTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSaseUserTagResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSaseUserTagResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteSaseUserTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSaseUserTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSaseUserTagResponseBody) SetCode(v int32) *DeleteSaseUserTagResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSaseUserTagResponseBody) SetRequestId(v string) *DeleteSaseUserTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSaseUserTagResponseBody) SetSuccess(v bool) *DeleteSaseUserTagResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSaseUserTagResponseBody) Validate() error {
	return dara.Validate(s)
}
