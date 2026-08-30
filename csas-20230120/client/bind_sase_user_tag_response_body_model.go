// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindSaseUserTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *BindSaseUserTagResponseBody
	GetCode() *int32
	SetRequestId(v string) *BindSaseUserTagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BindSaseUserTagResponseBody
	GetSuccess() *bool
}

type BindSaseUserTagResponseBody struct {
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
	// BE4FB974-11BC-5453-9BE1-1606A73EACA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindSaseUserTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindSaseUserTagResponseBody) GoString() string {
	return s.String()
}

func (s *BindSaseUserTagResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *BindSaseUserTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindSaseUserTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BindSaseUserTagResponseBody) SetCode(v int32) *BindSaseUserTagResponseBody {
	s.Code = &v
	return s
}

func (s *BindSaseUserTagResponseBody) SetRequestId(v string) *BindSaseUserTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindSaseUserTagResponseBody) SetSuccess(v bool) *BindSaseUserTagResponseBody {
	s.Success = &v
	return s
}

func (s *BindSaseUserTagResponseBody) Validate() error {
	return dara.Validate(s)
}
