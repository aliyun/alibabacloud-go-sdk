// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepoTagScanTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateRepoTagScanTaskResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *CreateRepoTagScanTaskResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *CreateRepoTagScanTaskResponseBody
	GetRequestId() *string
}

type CreateRepoTagScanTaskResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the API request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BC648259-91A7-4502-BED3-EDF64361FA83
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRepoTagScanTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRepoTagScanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepoTagScanTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateRepoTagScanTaskResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *CreateRepoTagScanTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRepoTagScanTaskResponseBody) SetCode(v string) *CreateRepoTagScanTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRepoTagScanTaskResponseBody) SetIsSuccess(v bool) *CreateRepoTagScanTaskResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateRepoTagScanTaskResponseBody) SetRequestId(v string) *CreateRepoTagScanTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRepoTagScanTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
