// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRequestDiagnosisPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRequestDiagnosisPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRequestDiagnosisPageResponse
	GetStatusCode() *int32
	SetBody(v *GetRequestDiagnosisPageResponseBody) *GetRequestDiagnosisPageResponse
	GetBody() *GetRequestDiagnosisPageResponseBody
}

type GetRequestDiagnosisPageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRequestDiagnosisPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRequestDiagnosisPageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRequestDiagnosisPageResponse) GoString() string {
	return s.String()
}

func (s *GetRequestDiagnosisPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRequestDiagnosisPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRequestDiagnosisPageResponse) GetBody() *GetRequestDiagnosisPageResponseBody {
	return s.Body
}

func (s *GetRequestDiagnosisPageResponse) SetHeaders(v map[string]*string) *GetRequestDiagnosisPageResponse {
	s.Headers = v
	return s
}

func (s *GetRequestDiagnosisPageResponse) SetStatusCode(v int32) *GetRequestDiagnosisPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRequestDiagnosisPageResponse) SetBody(v *GetRequestDiagnosisPageResponseBody) *GetRequestDiagnosisPageResponse {
	s.Body = v
	return s
}

func (s *GetRequestDiagnosisPageResponse) Validate() error {
	return dara.Validate(s)
}
