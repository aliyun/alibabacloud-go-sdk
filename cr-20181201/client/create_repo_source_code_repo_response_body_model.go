// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepoSourceCodeRepoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateRepoSourceCodeRepoResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *CreateRepoSourceCodeRepoResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *CreateRepoSourceCodeRepoResponseBody
	GetRequestId() *string
}

type CreateRepoSourceCodeRepoResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
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
	// 4CE1F661-75DD-4EBD-A4AD-057B26834ABB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRepoSourceCodeRepoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRepoSourceCodeRepoResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepoSourceCodeRepoResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateRepoSourceCodeRepoResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *CreateRepoSourceCodeRepoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRepoSourceCodeRepoResponseBody) SetCode(v string) *CreateRepoSourceCodeRepoResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRepoSourceCodeRepoResponseBody) SetIsSuccess(v bool) *CreateRepoSourceCodeRepoResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateRepoSourceCodeRepoResponseBody) SetRequestId(v string) *CreateRepoSourceCodeRepoResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRepoSourceCodeRepoResponseBody) Validate() error {
	return dara.Validate(s)
}
