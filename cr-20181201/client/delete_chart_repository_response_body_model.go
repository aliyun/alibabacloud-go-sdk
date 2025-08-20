// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChartRepositoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteChartRepositoryResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *DeleteChartRepositoryResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteChartRepositoryResponseBody
	GetRequestId() *string
}

type DeleteChartRepositoryResponseBody struct {
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
	// 12589EF7-96E2-4554-AAD7-F7209E88CAD3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteChartRepositoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteChartRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChartRepositoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteChartRepositoryResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteChartRepositoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteChartRepositoryResponseBody) SetCode(v string) *DeleteChartRepositoryResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteChartRepositoryResponseBody) SetIsSuccess(v bool) *DeleteChartRepositoryResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteChartRepositoryResponseBody) SetRequestId(v string) *DeleteChartRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteChartRepositoryResponseBody) Validate() error {
	return dara.Validate(s)
}
