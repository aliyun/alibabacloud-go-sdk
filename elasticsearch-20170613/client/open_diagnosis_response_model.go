// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenDiagnosisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenDiagnosisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenDiagnosisResponse
	GetStatusCode() *int32
	SetBody(v *OpenDiagnosisResponseBody) *OpenDiagnosisResponse
	GetBody() *OpenDiagnosisResponseBody
}

type OpenDiagnosisResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenDiagnosisResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *OpenDiagnosisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenDiagnosisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenDiagnosisResponse) GetBody() *OpenDiagnosisResponseBody {
	return s.Body
}

func (s *OpenDiagnosisResponse) SetHeaders(v map[string]*string) *OpenDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *OpenDiagnosisResponse) SetStatusCode(v int32) *OpenDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenDiagnosisResponse) SetBody(v *OpenDiagnosisResponseBody) *OpenDiagnosisResponse {
	s.Body = v
	return s
}

func (s *OpenDiagnosisResponse) Validate() error {
	return dara.Validate(s)
}
