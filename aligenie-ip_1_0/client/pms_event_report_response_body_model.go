// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPmsEventReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *PmsEventReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *PmsEventReportResponseBody
	GetRequestId() *string
	SetResult(v bool) *PmsEventReportResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *PmsEventReportResponseBody
	GetStatusCode() *int32
}

type PmsEventReportResponseBody struct {
	// example:
	//
	// success
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s PmsEventReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PmsEventReportResponseBody) GoString() string {
	return s.String()
}

func (s *PmsEventReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PmsEventReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PmsEventReportResponseBody) GetResult() *bool {
	return s.Result
}

func (s *PmsEventReportResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PmsEventReportResponseBody) SetMessage(v string) *PmsEventReportResponseBody {
	s.Message = &v
	return s
}

func (s *PmsEventReportResponseBody) SetRequestId(v string) *PmsEventReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *PmsEventReportResponseBody) SetResult(v bool) *PmsEventReportResponseBody {
	s.Result = &v
	return s
}

func (s *PmsEventReportResponseBody) SetStatusCode(v int32) *PmsEventReportResponseBody {
	s.StatusCode = &v
	return s
}

func (s *PmsEventReportResponseBody) Validate() error {
	return dara.Validate(s)
}
