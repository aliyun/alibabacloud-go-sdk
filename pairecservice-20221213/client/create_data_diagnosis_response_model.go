// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataDiagnosisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataDiagnosisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataDiagnosisResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataDiagnosisResponseBody) *CreateDataDiagnosisResponse
	GetBody() *CreateDataDiagnosisResponseBody
}

type CreateDataDiagnosisResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataDiagnosisResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *CreateDataDiagnosisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataDiagnosisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataDiagnosisResponse) GetBody() *CreateDataDiagnosisResponseBody {
	return s.Body
}

func (s *CreateDataDiagnosisResponse) SetHeaders(v map[string]*string) *CreateDataDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *CreateDataDiagnosisResponse) SetStatusCode(v int32) *CreateDataDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataDiagnosisResponse) SetBody(v *CreateDataDiagnosisResponseBody) *CreateDataDiagnosisResponse {
	s.Body = v
	return s
}

func (s *CreateDataDiagnosisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
