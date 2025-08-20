// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepoTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateRepoTagResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *CreateRepoTagResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *CreateRepoTagResponseBody
	GetRequestId() *string
}

type CreateRepoTagResponseBody struct {
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
	// C4C7DD0C-C9D6-437A-A7EE-8BY*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRepoTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRepoTagResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepoTagResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateRepoTagResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *CreateRepoTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRepoTagResponseBody) SetCode(v string) *CreateRepoTagResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRepoTagResponseBody) SetIsSuccess(v bool) *CreateRepoTagResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateRepoTagResponseBody) SetRequestId(v string) *CreateRepoTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRepoTagResponseBody) Validate() error {
	return dara.Validate(s)
}
