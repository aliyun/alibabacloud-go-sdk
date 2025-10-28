// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRequestDiagnosisResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRequestDiagnosisResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRequestDiagnosisResultResponse
	GetStatusCode() *int32
	SetBody(v *GetRequestDiagnosisResultResponseBody) *GetRequestDiagnosisResultResponse
	GetBody() *GetRequestDiagnosisResultResponseBody
}

type GetRequestDiagnosisResultResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRequestDiagnosisResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRequestDiagnosisResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRequestDiagnosisResultResponse) GoString() string {
	return s.String()
}

func (s *GetRequestDiagnosisResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRequestDiagnosisResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRequestDiagnosisResultResponse) GetBody() *GetRequestDiagnosisResultResponseBody {
	return s.Body
}

func (s *GetRequestDiagnosisResultResponse) SetHeaders(v map[string]*string) *GetRequestDiagnosisResultResponse {
	s.Headers = v
	return s
}

func (s *GetRequestDiagnosisResultResponse) SetStatusCode(v int32) *GetRequestDiagnosisResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRequestDiagnosisResultResponse) SetBody(v *GetRequestDiagnosisResultResponseBody) *GetRequestDiagnosisResultResponse {
	s.Body = v
	return s
}

func (s *GetRequestDiagnosisResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
