// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataDiagnosisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataDiagnosisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataDiagnosisResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataDiagnosisResponseBody) *UpdateDataDiagnosisResponse
	GetBody() *UpdateDataDiagnosisResponseBody
}

type UpdateDataDiagnosisResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataDiagnosisResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataDiagnosisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataDiagnosisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataDiagnosisResponse) GetBody() *UpdateDataDiagnosisResponseBody {
	return s.Body
}

func (s *UpdateDataDiagnosisResponse) SetHeaders(v map[string]*string) *UpdateDataDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataDiagnosisResponse) SetStatusCode(v int32) *UpdateDataDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataDiagnosisResponse) SetBody(v *UpdateDataDiagnosisResponseBody) *UpdateDataDiagnosisResponse {
	s.Body = v
	return s
}

func (s *UpdateDataDiagnosisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
