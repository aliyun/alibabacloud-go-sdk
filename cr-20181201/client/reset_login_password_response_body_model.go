// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetLoginPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResetLoginPasswordResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *ResetLoginPasswordResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *ResetLoginPasswordResponseBody
	GetRequestId() *string
}

type ResetLoginPasswordResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EB9C5722-51E2-4497-A573-575B0CA5CE0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetLoginPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetLoginPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetLoginPasswordResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResetLoginPasswordResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ResetLoginPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetLoginPasswordResponseBody) SetCode(v string) *ResetLoginPasswordResponseBody {
	s.Code = &v
	return s
}

func (s *ResetLoginPasswordResponseBody) SetIsSuccess(v bool) *ResetLoginPasswordResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ResetLoginPasswordResponseBody) SetRequestId(v string) *ResetLoginPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetLoginPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
