// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVmcoreDiagnosisTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVmcoreDiagnosisTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVmcoreDiagnosisTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateVmcoreDiagnosisTaskResponseBody) *CreateVmcoreDiagnosisTaskResponse
	GetBody() *CreateVmcoreDiagnosisTaskResponseBody
}

type CreateVmcoreDiagnosisTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVmcoreDiagnosisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVmcoreDiagnosisTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVmcoreDiagnosisTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVmcoreDiagnosisTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVmcoreDiagnosisTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVmcoreDiagnosisTaskResponse) GetBody() *CreateVmcoreDiagnosisTaskResponseBody {
	return s.Body
}

func (s *CreateVmcoreDiagnosisTaskResponse) SetHeaders(v map[string]*string) *CreateVmcoreDiagnosisTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateVmcoreDiagnosisTaskResponse) SetStatusCode(v int32) *CreateVmcoreDiagnosisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVmcoreDiagnosisTaskResponse) SetBody(v *CreateVmcoreDiagnosisTaskResponseBody) *CreateVmcoreDiagnosisTaskResponse {
	s.Body = v
	return s
}

func (s *CreateVmcoreDiagnosisTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
