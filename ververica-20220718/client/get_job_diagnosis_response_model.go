// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobDiagnosisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJobDiagnosisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJobDiagnosisResponse
	GetStatusCode() *int32
	SetBody(v *GetJobDiagnosisResponseBody) *GetJobDiagnosisResponse
	GetBody() *GetJobDiagnosisResponseBody
}

type GetJobDiagnosisResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobDiagnosisResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJobDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *GetJobDiagnosisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJobDiagnosisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJobDiagnosisResponse) GetBody() *GetJobDiagnosisResponseBody {
	return s.Body
}

func (s *GetJobDiagnosisResponse) SetHeaders(v map[string]*string) *GetJobDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *GetJobDiagnosisResponse) SetStatusCode(v int32) *GetJobDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobDiagnosisResponse) SetBody(v *GetJobDiagnosisResponseBody) *GetJobDiagnosisResponse {
	s.Body = v
	return s
}

func (s *GetJobDiagnosisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
