// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeDiagnosisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvokeDiagnosisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvokeDiagnosisResponse
	GetStatusCode() *int32
	SetBody(v *InvokeDiagnosisResponseBody) *InvokeDiagnosisResponse
	GetBody() *InvokeDiagnosisResponseBody
}

type InvokeDiagnosisResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvokeDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokeDiagnosisResponse) String() string {
	return dara.Prettify(s)
}

func (s InvokeDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *InvokeDiagnosisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvokeDiagnosisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvokeDiagnosisResponse) GetBody() *InvokeDiagnosisResponseBody {
	return s.Body
}

func (s *InvokeDiagnosisResponse) SetHeaders(v map[string]*string) *InvokeDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *InvokeDiagnosisResponse) SetStatusCode(v int32) *InvokeDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeDiagnosisResponse) SetBody(v *InvokeDiagnosisResponseBody) *InvokeDiagnosisResponse {
	s.Body = v
	return s
}

func (s *InvokeDiagnosisResponse) Validate() error {
	return dara.Validate(s)
}
