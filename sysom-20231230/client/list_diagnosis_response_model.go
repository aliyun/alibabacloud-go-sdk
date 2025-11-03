// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiagnosisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDiagnosisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDiagnosisResponse
	GetStatusCode() *int32
	SetBody(v *ListDiagnosisResponseBody) *ListDiagnosisResponse
	GetBody() *ListDiagnosisResponseBody
}

type ListDiagnosisResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDiagnosisResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *ListDiagnosisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDiagnosisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDiagnosisResponse) GetBody() *ListDiagnosisResponseBody {
	return s.Body
}

func (s *ListDiagnosisResponse) SetHeaders(v map[string]*string) *ListDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *ListDiagnosisResponse) SetStatusCode(v int32) *ListDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDiagnosisResponse) SetBody(v *ListDiagnosisResponseBody) *ListDiagnosisResponse {
	s.Body = v
	return s
}

func (s *ListDiagnosisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
