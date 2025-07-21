// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportAppOtaInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReportAppOtaInfoResponseBody
	GetCode() *string
	SetMessage(v string) *ReportAppOtaInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReportAppOtaInfoResponseBody
	GetRequestId() *string
}

type ReportAppOtaInfoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportAppOtaInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReportAppOtaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ReportAppOtaInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReportAppOtaInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReportAppOtaInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReportAppOtaInfoResponseBody) SetCode(v string) *ReportAppOtaInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ReportAppOtaInfoResponseBody) SetMessage(v string) *ReportAppOtaInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ReportAppOtaInfoResponseBody) SetRequestId(v string) *ReportAppOtaInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportAppOtaInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
