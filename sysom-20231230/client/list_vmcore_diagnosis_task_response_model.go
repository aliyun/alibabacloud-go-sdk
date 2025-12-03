// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVmcoreDiagnosisTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVmcoreDiagnosisTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVmcoreDiagnosisTaskResponse
	GetStatusCode() *int32
	SetBody(v *ListVmcoreDiagnosisTaskResponseBody) *ListVmcoreDiagnosisTaskResponse
	GetBody() *ListVmcoreDiagnosisTaskResponseBody
}

type ListVmcoreDiagnosisTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVmcoreDiagnosisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVmcoreDiagnosisTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVmcoreDiagnosisTaskResponse) GoString() string {
	return s.String()
}

func (s *ListVmcoreDiagnosisTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVmcoreDiagnosisTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVmcoreDiagnosisTaskResponse) GetBody() *ListVmcoreDiagnosisTaskResponseBody {
	return s.Body
}

func (s *ListVmcoreDiagnosisTaskResponse) SetHeaders(v map[string]*string) *ListVmcoreDiagnosisTaskResponse {
	s.Headers = v
	return s
}

func (s *ListVmcoreDiagnosisTaskResponse) SetStatusCode(v int32) *ListVmcoreDiagnosisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVmcoreDiagnosisTaskResponse) SetBody(v *ListVmcoreDiagnosisTaskResponseBody) *ListVmcoreDiagnosisTaskResponse {
	s.Body = v
	return s
}

func (s *ListVmcoreDiagnosisTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
