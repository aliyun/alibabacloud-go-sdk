// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDiagnosisResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDiagnosisResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDiagnosisResultResponse
	GetStatusCode() *int32
	SetBody(v *GetDiagnosisResultResponseBody) *GetDiagnosisResultResponse
	GetBody() *GetDiagnosisResultResponseBody
}

type GetDiagnosisResultResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDiagnosisResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDiagnosisResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDiagnosisResultResponse) GoString() string {
	return s.String()
}

func (s *GetDiagnosisResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDiagnosisResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDiagnosisResultResponse) GetBody() *GetDiagnosisResultResponseBody {
	return s.Body
}

func (s *GetDiagnosisResultResponse) SetHeaders(v map[string]*string) *GetDiagnosisResultResponse {
	s.Headers = v
	return s
}

func (s *GetDiagnosisResultResponse) SetStatusCode(v int32) *GetDiagnosisResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDiagnosisResultResponse) SetBody(v *GetDiagnosisResultResponseBody) *GetDiagnosisResultResponse {
	s.Body = v
	return s
}

func (s *GetDiagnosisResultResponse) Validate() error {
	return dara.Validate(s)
}
