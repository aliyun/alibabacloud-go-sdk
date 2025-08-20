// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChartRepositoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateChartRepositoryResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *UpdateChartRepositoryResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *UpdateChartRepositoryResponseBody
	GetRequestId() *string
}

type UpdateChartRepositoryResponseBody struct {
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

func (s UpdateChartRepositoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateChartRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateChartRepositoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateChartRepositoryResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *UpdateChartRepositoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateChartRepositoryResponseBody) SetCode(v string) *UpdateChartRepositoryResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateChartRepositoryResponseBody) SetIsSuccess(v bool) *UpdateChartRepositoryResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateChartRepositoryResponseBody) SetRequestId(v string) *UpdateChartRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateChartRepositoryResponseBody) Validate() error {
	return dara.Validate(s)
}
