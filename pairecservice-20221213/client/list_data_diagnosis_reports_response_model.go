// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataDiagnosisReportsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataDiagnosisReportsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataDiagnosisReportsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataDiagnosisReportsResponseBody) *ListDataDiagnosisReportsResponse
	GetBody() *ListDataDiagnosisReportsResponseBody
}

type ListDataDiagnosisReportsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataDiagnosisReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataDiagnosisReportsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataDiagnosisReportsResponse) GoString() string {
	return s.String()
}

func (s *ListDataDiagnosisReportsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataDiagnosisReportsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataDiagnosisReportsResponse) GetBody() *ListDataDiagnosisReportsResponseBody {
	return s.Body
}

func (s *ListDataDiagnosisReportsResponse) SetHeaders(v map[string]*string) *ListDataDiagnosisReportsResponse {
	s.Headers = v
	return s
}

func (s *ListDataDiagnosisReportsResponse) SetStatusCode(v int32) *ListDataDiagnosisReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataDiagnosisReportsResponse) SetBody(v *ListDataDiagnosisReportsResponseBody) *ListDataDiagnosisReportsResponse {
	s.Body = v
	return s
}

func (s *ListDataDiagnosisReportsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
