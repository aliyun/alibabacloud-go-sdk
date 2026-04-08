// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPADiagnosisTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPADiagnosisTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPADiagnosisTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetPADiagnosisTaskResponseBody) *GetPADiagnosisTaskResponse
	GetBody() *GetPADiagnosisTaskResponseBody
}

type GetPADiagnosisTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPADiagnosisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPADiagnosisTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPADiagnosisTaskResponse) GoString() string {
	return s.String()
}

func (s *GetPADiagnosisTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPADiagnosisTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPADiagnosisTaskResponse) GetBody() *GetPADiagnosisTaskResponseBody {
	return s.Body
}

func (s *GetPADiagnosisTaskResponse) SetHeaders(v map[string]*string) *GetPADiagnosisTaskResponse {
	s.Headers = v
	return s
}

func (s *GetPADiagnosisTaskResponse) SetStatusCode(v int32) *GetPADiagnosisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPADiagnosisTaskResponse) SetBody(v *GetPADiagnosisTaskResponseBody) *GetPADiagnosisTaskResponse {
	s.Body = v
	return s
}

func (s *GetPADiagnosisTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
