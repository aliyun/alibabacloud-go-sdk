// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateReportResponseBody
	GetRequestId() *string
	SetResult(v string) *CreateReportResponseBody
	GetResult() *string
	SetSuccess(v bool) *CreateReportResponseBody
	GetSuccess() *bool
}

type CreateReportResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateReportResponseBody) GoString() string {
	return s.String()
}

func (s *CreateReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateReportResponseBody) GetResult() *string {
	return s.Result
}

func (s *CreateReportResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateReportResponseBody) SetRequestId(v string) *CreateReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateReportResponseBody) SetResult(v string) *CreateReportResponseBody {
	s.Result = &v
	return s
}

func (s *CreateReportResponseBody) SetSuccess(v bool) *CreateReportResponseBody {
	s.Success = &v
	return s
}

func (s *CreateReportResponseBody) Validate() error {
	return dara.Validate(s)
}
