// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiagnoseReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDiagnoseReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDiagnoseReportResponse
	GetStatusCode() *int32
	SetBody(v *ListDiagnoseReportResponseBody) *ListDiagnoseReportResponse
	GetBody() *ListDiagnoseReportResponseBody
}

type ListDiagnoseReportResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDiagnoseReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDiagnoseReportResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnoseReportResponse) GoString() string {
	return s.String()
}

func (s *ListDiagnoseReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDiagnoseReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDiagnoseReportResponse) GetBody() *ListDiagnoseReportResponseBody {
	return s.Body
}

func (s *ListDiagnoseReportResponse) SetHeaders(v map[string]*string) *ListDiagnoseReportResponse {
	s.Headers = v
	return s
}

func (s *ListDiagnoseReportResponse) SetStatusCode(v int32) *ListDiagnoseReportResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDiagnoseReportResponse) SetBody(v *ListDiagnoseReportResponseBody) *ListDiagnoseReportResponse {
	s.Body = v
	return s
}

func (s *ListDiagnoseReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
