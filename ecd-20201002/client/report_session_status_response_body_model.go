// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportSessionStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReportSessionStatusResponseBody
	GetRequestId() *string
}

type ReportSessionStatusResponseBody struct {
	// example:
	//
	// 0EE5DE20-25F4-5870-9D56-C259A45B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportSessionStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReportSessionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ReportSessionStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReportSessionStatusResponseBody) SetRequestId(v string) *ReportSessionStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportSessionStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
