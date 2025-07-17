// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiagnoseReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDiagnoseReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDiagnoseReportResponse
	GetStatusCode() *int32
	SetBody(v *CreateDiagnoseReportResponseBody) *CreateDiagnoseReportResponse
	GetBody() *CreateDiagnoseReportResponseBody
}

type CreateDiagnoseReportResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDiagnoseReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDiagnoseReportResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDiagnoseReportResponse) GoString() string {
	return s.String()
}

func (s *CreateDiagnoseReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDiagnoseReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDiagnoseReportResponse) GetBody() *CreateDiagnoseReportResponseBody {
	return s.Body
}

func (s *CreateDiagnoseReportResponse) SetHeaders(v map[string]*string) *CreateDiagnoseReportResponse {
	s.Headers = v
	return s
}

func (s *CreateDiagnoseReportResponse) SetStatusCode(v int32) *CreateDiagnoseReportResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDiagnoseReportResponse) SetBody(v *CreateDiagnoseReportResponseBody) *CreateDiagnoseReportResponse {
	s.Body = v
	return s
}

func (s *CreateDiagnoseReportResponse) Validate() error {
	return dara.Validate(s)
}
