// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepositoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateRepositoryResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *CreateRepositoryResponseBody
	GetIsSuccess() *bool
	SetRepoId(v string) *CreateRepositoryResponseBody
	GetRepoId() *string
	SetRequestId(v string) *CreateRepositoryResponseBody
	GetRequestId() *string
}

type CreateRepositoryResponseBody struct {
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
	// The ID of the image repository.
	//
	// example:
	//
	// crr-xwvi3osiy4ff****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 886FB272-15C3-44FC-AA54-F4ABD5B93A28
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRepositoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepositoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateRepositoryResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *CreateRepositoryResponseBody) GetRepoId() *string {
	return s.RepoId
}

func (s *CreateRepositoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRepositoryResponseBody) SetCode(v string) *CreateRepositoryResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRepositoryResponseBody) SetIsSuccess(v bool) *CreateRepositoryResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateRepositoryResponseBody) SetRepoId(v string) *CreateRepositoryResponseBody {
	s.RepoId = &v
	return s
}

func (s *CreateRepositoryResponseBody) SetRequestId(v string) *CreateRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRepositoryResponseBody) Validate() error {
	return dara.Validate(s)
}
