// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChartReleaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteChartReleaseResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *DeleteChartReleaseResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteChartReleaseResponseBody
	GetRequestId() *string
}

type DeleteChartReleaseResponseBody struct {
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
	// C2D6CE47-6DEF-45F4-A1AC-90F3AFBA751F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteChartReleaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteChartReleaseResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChartReleaseResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteChartReleaseResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteChartReleaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteChartReleaseResponseBody) SetCode(v string) *DeleteChartReleaseResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteChartReleaseResponseBody) SetIsSuccess(v bool) *DeleteChartReleaseResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteChartReleaseResponseBody) SetRequestId(v string) *DeleteChartReleaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteChartReleaseResponseBody) Validate() error {
	return dara.Validate(s)
}
