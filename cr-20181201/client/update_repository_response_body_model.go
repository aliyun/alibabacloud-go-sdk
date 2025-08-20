// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRepositoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateRepositoryResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *UpdateRepositoryResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *UpdateRepositoryResponseBody
	GetRequestId() *string
}

type UpdateRepositoryResponseBody struct {
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// example:
	//
	// CC43EC6B-0DD4-40AE-8811-B0519617051A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRepositoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateRepositoryResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *UpdateRepositoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRepositoryResponseBody) SetCode(v string) *UpdateRepositoryResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRepositoryResponseBody) SetIsSuccess(v bool) *UpdateRepositoryResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateRepositoryResponseBody) SetRequestId(v string) *UpdateRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRepositoryResponseBody) Validate() error {
	return dara.Validate(s)
}
