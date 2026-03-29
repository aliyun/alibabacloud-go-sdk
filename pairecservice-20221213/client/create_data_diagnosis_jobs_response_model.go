// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataDiagnosisJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataDiagnosisJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataDiagnosisJobsResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataDiagnosisJobsResponseBody) *CreateDataDiagnosisJobsResponse
	GetBody() *CreateDataDiagnosisJobsResponseBody
}

type CreateDataDiagnosisJobsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataDiagnosisJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataDiagnosisJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataDiagnosisJobsResponse) GoString() string {
	return s.String()
}

func (s *CreateDataDiagnosisJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataDiagnosisJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataDiagnosisJobsResponse) GetBody() *CreateDataDiagnosisJobsResponseBody {
	return s.Body
}

func (s *CreateDataDiagnosisJobsResponse) SetHeaders(v map[string]*string) *CreateDataDiagnosisJobsResponse {
	s.Headers = v
	return s
}

func (s *CreateDataDiagnosisJobsResponse) SetStatusCode(v int32) *CreateDataDiagnosisJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataDiagnosisJobsResponse) SetBody(v *CreateDataDiagnosisJobsResponseBody) *CreateDataDiagnosisJobsResponse {
	s.Body = v
	return s
}

func (s *CreateDataDiagnosisJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
