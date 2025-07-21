// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportDeviceOtaInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReportDeviceOtaInfoResponseBody
	GetCode() *string
	SetMessage(v string) *ReportDeviceOtaInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReportDeviceOtaInfoResponseBody
	GetRequestId() *string
}

type ReportDeviceOtaInfoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportDeviceOtaInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReportDeviceOtaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ReportDeviceOtaInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReportDeviceOtaInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReportDeviceOtaInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReportDeviceOtaInfoResponseBody) SetCode(v string) *ReportDeviceOtaInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ReportDeviceOtaInfoResponseBody) SetMessage(v string) *ReportDeviceOtaInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ReportDeviceOtaInfoResponseBody) SetRequestId(v string) *ReportDeviceOtaInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportDeviceOtaInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
