// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePADiagnosisTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePADiagnosisTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePADiagnosisTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreatePADiagnosisTaskResponseBody) *CreatePADiagnosisTaskResponse
	GetBody() *CreatePADiagnosisTaskResponseBody
}

type CreatePADiagnosisTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePADiagnosisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePADiagnosisTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePADiagnosisTaskResponse) GoString() string {
	return s.String()
}

func (s *CreatePADiagnosisTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePADiagnosisTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePADiagnosisTaskResponse) GetBody() *CreatePADiagnosisTaskResponseBody {
	return s.Body
}

func (s *CreatePADiagnosisTaskResponse) SetHeaders(v map[string]*string) *CreatePADiagnosisTaskResponse {
	s.Headers = v
	return s
}

func (s *CreatePADiagnosisTaskResponse) SetStatusCode(v int32) *CreatePADiagnosisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePADiagnosisTaskResponse) SetBody(v *CreatePADiagnosisTaskResponseBody) *CreatePADiagnosisTaskResponse {
	s.Body = v
	return s
}

func (s *CreatePADiagnosisTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
