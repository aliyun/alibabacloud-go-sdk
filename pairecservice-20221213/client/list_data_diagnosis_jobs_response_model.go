// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataDiagnosisJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataDiagnosisJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataDiagnosisJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataDiagnosisJobsResponseBody) *ListDataDiagnosisJobsResponse
	GetBody() *ListDataDiagnosisJobsResponseBody
}

type ListDataDiagnosisJobsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataDiagnosisJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataDiagnosisJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataDiagnosisJobsResponse) GoString() string {
	return s.String()
}

func (s *ListDataDiagnosisJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataDiagnosisJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataDiagnosisJobsResponse) GetBody() *ListDataDiagnosisJobsResponseBody {
	return s.Body
}

func (s *ListDataDiagnosisJobsResponse) SetHeaders(v map[string]*string) *ListDataDiagnosisJobsResponse {
	s.Headers = v
	return s
}

func (s *ListDataDiagnosisJobsResponse) SetStatusCode(v int32) *ListDataDiagnosisJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataDiagnosisJobsResponse) SetBody(v *ListDataDiagnosisJobsResponseBody) *ListDataDiagnosisJobsResponse {
	s.Body = v
	return s
}

func (s *ListDataDiagnosisJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
