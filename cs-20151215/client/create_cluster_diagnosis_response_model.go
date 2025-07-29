// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterDiagnosisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateClusterDiagnosisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateClusterDiagnosisResponse
	GetStatusCode() *int32
	SetBody(v *CreateClusterDiagnosisResponseBody) *CreateClusterDiagnosisResponse
	GetBody() *CreateClusterDiagnosisResponseBody
}

type CreateClusterDiagnosisResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClusterDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateClusterDiagnosisResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterDiagnosisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateClusterDiagnosisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateClusterDiagnosisResponse) GetBody() *CreateClusterDiagnosisResponseBody {
	return s.Body
}

func (s *CreateClusterDiagnosisResponse) SetHeaders(v map[string]*string) *CreateClusterDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *CreateClusterDiagnosisResponse) SetStatusCode(v int32) *CreateClusterDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClusterDiagnosisResponse) SetBody(v *CreateClusterDiagnosisResponseBody) *CreateClusterDiagnosisResponse {
	s.Body = v
	return s
}

func (s *CreateClusterDiagnosisResponse) Validate() error {
	return dara.Validate(s)
}
