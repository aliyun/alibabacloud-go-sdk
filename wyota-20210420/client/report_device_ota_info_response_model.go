// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportDeviceOtaInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReportDeviceOtaInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReportDeviceOtaInfoResponse
	GetStatusCode() *int32
	SetBody(v *ReportDeviceOtaInfoResponseBody) *ReportDeviceOtaInfoResponse
	GetBody() *ReportDeviceOtaInfoResponseBody
}

type ReportDeviceOtaInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportDeviceOtaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportDeviceOtaInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ReportDeviceOtaInfoResponse) GoString() string {
	return s.String()
}

func (s *ReportDeviceOtaInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReportDeviceOtaInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReportDeviceOtaInfoResponse) GetBody() *ReportDeviceOtaInfoResponseBody {
	return s.Body
}

func (s *ReportDeviceOtaInfoResponse) SetHeaders(v map[string]*string) *ReportDeviceOtaInfoResponse {
	s.Headers = v
	return s
}

func (s *ReportDeviceOtaInfoResponse) SetStatusCode(v int32) *ReportDeviceOtaInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportDeviceOtaInfoResponse) SetBody(v *ReportDeviceOtaInfoResponseBody) *ReportDeviceOtaInfoResponse {
	s.Body = v
	return s
}

func (s *ReportDeviceOtaInfoResponse) Validate() error {
	return dara.Validate(s)
}
