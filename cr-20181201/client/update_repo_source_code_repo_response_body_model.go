// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRepoSourceCodeRepoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateRepoSourceCodeRepoResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *UpdateRepoSourceCodeRepoResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *UpdateRepoSourceCodeRepoResponseBody
	GetRequestId() *string
}

type UpdateRepoSourceCodeRepoResponseBody struct {
	// The return value.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F56D589D-AF7F-4900-BA46-62C780AC2C10
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRepoSourceCodeRepoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRepoSourceCodeRepoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRepoSourceCodeRepoResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateRepoSourceCodeRepoResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *UpdateRepoSourceCodeRepoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRepoSourceCodeRepoResponseBody) SetCode(v string) *UpdateRepoSourceCodeRepoResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRepoSourceCodeRepoResponseBody) SetIsSuccess(v bool) *UpdateRepoSourceCodeRepoResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateRepoSourceCodeRepoResponseBody) SetRequestId(v string) *UpdateRepoSourceCodeRepoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRepoSourceCodeRepoResponseBody) Validate() error {
	return dara.Validate(s)
}
