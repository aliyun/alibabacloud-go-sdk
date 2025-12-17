// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDiagnoseReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDiagnoseReportResponseBody
	GetRequestId() *string
}

type DeleteDiagnoseReportResponseBody struct {
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDiagnoseReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDiagnoseReportResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDiagnoseReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDiagnoseReportResponseBody) SetRequestId(v string) *DeleteDiagnoseReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDiagnoseReportResponseBody) Validate() error {
	return dara.Validate(s)
}
