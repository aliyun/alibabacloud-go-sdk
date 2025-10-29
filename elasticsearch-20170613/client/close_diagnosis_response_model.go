// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseDiagnosisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseDiagnosisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseDiagnosisResponse
	GetStatusCode() *int32
	SetBody(v *CloseDiagnosisResponseBody) *CloseDiagnosisResponse
	GetBody() *CloseDiagnosisResponseBody
}

type CloseDiagnosisResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseDiagnosisResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *CloseDiagnosisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseDiagnosisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseDiagnosisResponse) GetBody() *CloseDiagnosisResponseBody {
	return s.Body
}

func (s *CloseDiagnosisResponse) SetHeaders(v map[string]*string) *CloseDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *CloseDiagnosisResponse) SetStatusCode(v int32) *CloseDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseDiagnosisResponse) SetBody(v *CloseDiagnosisResponseBody) *CloseDiagnosisResponse {
	s.Body = v
	return s
}

func (s *CloseDiagnosisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
