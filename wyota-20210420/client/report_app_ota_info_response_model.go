// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportAppOtaInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReportAppOtaInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReportAppOtaInfoResponse
	GetStatusCode() *int32
	SetBody(v *ReportAppOtaInfoResponseBody) *ReportAppOtaInfoResponse
	GetBody() *ReportAppOtaInfoResponseBody
}

type ReportAppOtaInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportAppOtaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportAppOtaInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ReportAppOtaInfoResponse) GoString() string {
	return s.String()
}

func (s *ReportAppOtaInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReportAppOtaInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReportAppOtaInfoResponse) GetBody() *ReportAppOtaInfoResponseBody {
	return s.Body
}

func (s *ReportAppOtaInfoResponse) SetHeaders(v map[string]*string) *ReportAppOtaInfoResponse {
	s.Headers = v
	return s
}

func (s *ReportAppOtaInfoResponse) SetStatusCode(v int32) *ReportAppOtaInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportAppOtaInfoResponse) SetBody(v *ReportAppOtaInfoResponseBody) *ReportAppOtaInfoResponse {
	s.Body = v
	return s
}

func (s *ReportAppOtaInfoResponse) Validate() error {
	return dara.Validate(s)
}
