// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckReadableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckReadableResponseBody
	GetRequestId() *string
	SetResult(v bool) *CheckReadableResponseBody
	GetResult() *bool
	SetSuccess(v bool) *CheckReadableResponseBody
	GetSuccess() *bool
}

type CheckReadableResponseBody struct {
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

func (s CheckReadableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckReadableResponseBody) GoString() string {
	return s.String()
}

func (s *CheckReadableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckReadableResponseBody) GetResult() *bool {
	return s.Result
}

func (s *CheckReadableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckReadableResponseBody) SetRequestId(v string) *CheckReadableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckReadableResponseBody) SetResult(v bool) *CheckReadableResponseBody {
	s.Result = &v
	return s
}

func (s *CheckReadableResponseBody) SetSuccess(v bool) *CheckReadableResponseBody {
	s.Success = &v
	return s
}

func (s *CheckReadableResponseBody) Validate() error {
	return dara.Validate(s)
}
