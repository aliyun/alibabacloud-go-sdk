// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataDiagnosisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataDiagnosisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataDiagnosisResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataDiagnosisResponseBody) *DeleteDataDiagnosisResponse
	GetBody() *DeleteDataDiagnosisResponseBody
}

type DeleteDataDiagnosisResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataDiagnosisResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataDiagnosisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataDiagnosisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataDiagnosisResponse) GetBody() *DeleteDataDiagnosisResponseBody {
	return s.Body
}

func (s *DeleteDataDiagnosisResponse) SetHeaders(v map[string]*string) *DeleteDataDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataDiagnosisResponse) SetStatusCode(v int32) *DeleteDataDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataDiagnosisResponse) SetBody(v *DeleteDataDiagnosisResponseBody) *DeleteDataDiagnosisResponse {
	s.Body = v
	return s
}

func (s *DeleteDataDiagnosisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
