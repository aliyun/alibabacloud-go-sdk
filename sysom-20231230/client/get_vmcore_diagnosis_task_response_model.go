// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVmcoreDiagnosisTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVmcoreDiagnosisTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVmcoreDiagnosisTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetVmcoreDiagnosisTaskResponseBody) *GetVmcoreDiagnosisTaskResponse
	GetBody() *GetVmcoreDiagnosisTaskResponseBody
}

type GetVmcoreDiagnosisTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVmcoreDiagnosisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVmcoreDiagnosisTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVmcoreDiagnosisTaskResponse) GoString() string {
	return s.String()
}

func (s *GetVmcoreDiagnosisTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVmcoreDiagnosisTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVmcoreDiagnosisTaskResponse) GetBody() *GetVmcoreDiagnosisTaskResponseBody {
	return s.Body
}

func (s *GetVmcoreDiagnosisTaskResponse) SetHeaders(v map[string]*string) *GetVmcoreDiagnosisTaskResponse {
	s.Headers = v
	return s
}

func (s *GetVmcoreDiagnosisTaskResponse) SetStatusCode(v int32) *GetVmcoreDiagnosisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVmcoreDiagnosisTaskResponse) SetBody(v *GetVmcoreDiagnosisTaskResponseBody) *GetVmcoreDiagnosisTaskResponse {
	s.Body = v
	return s
}

func (s *GetVmcoreDiagnosisTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
