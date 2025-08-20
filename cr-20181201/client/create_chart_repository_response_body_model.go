// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChartRepositoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateChartRepositoryResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *CreateChartRepositoryResponseBody
	GetIsSuccess() *bool
	SetRepoId(v string) *CreateChartRepositoryResponseBody
	GetRepoId() *string
	SetRequestId(v string) *CreateChartRepositoryResponseBody
	GetRequestId() *string
}

type CreateChartRepositoryResponseBody struct {
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
	// The ID of the repository.
	//
	// example:
	//
	// crcr-2micxey5ewj4****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 60390244-A483-491A-B41D-F866C95380A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateChartRepositoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateChartRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChartRepositoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateChartRepositoryResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *CreateChartRepositoryResponseBody) GetRepoId() *string {
	return s.RepoId
}

func (s *CreateChartRepositoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateChartRepositoryResponseBody) SetCode(v string) *CreateChartRepositoryResponseBody {
	s.Code = &v
	return s
}

func (s *CreateChartRepositoryResponseBody) SetIsSuccess(v bool) *CreateChartRepositoryResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateChartRepositoryResponseBody) SetRepoId(v string) *CreateChartRepositoryResponseBody {
	s.RepoId = &v
	return s
}

func (s *CreateChartRepositoryResponseBody) SetRequestId(v string) *CreateChartRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateChartRepositoryResponseBody) Validate() error {
	return dara.Validate(s)
}
