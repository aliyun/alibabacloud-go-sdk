// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataDiagnosisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataDiagnosisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataDiagnosisResponse
	GetStatusCode() *int32
	SetBody(v *GetDataDiagnosisResponseBody) *GetDataDiagnosisResponse
	GetBody() *GetDataDiagnosisResponseBody
}

type GetDataDiagnosisResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataDiagnosisResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *GetDataDiagnosisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataDiagnosisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataDiagnosisResponse) GetBody() *GetDataDiagnosisResponseBody {
	return s.Body
}

func (s *GetDataDiagnosisResponse) SetHeaders(v map[string]*string) *GetDataDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *GetDataDiagnosisResponse) SetStatusCode(v int32) *GetDataDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataDiagnosisResponse) SetBody(v *GetDataDiagnosisResponseBody) *GetDataDiagnosisResponse {
	s.Body = v
	return s
}

func (s *GetDataDiagnosisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
