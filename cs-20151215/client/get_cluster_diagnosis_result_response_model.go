// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterDiagnosisResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClusterDiagnosisResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClusterDiagnosisResultResponse
	GetStatusCode() *int32
	SetBody(v *GetClusterDiagnosisResultResponseBody) *GetClusterDiagnosisResultResponse
	GetBody() *GetClusterDiagnosisResultResponseBody
}

type GetClusterDiagnosisResultResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClusterDiagnosisResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterDiagnosisResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClusterDiagnosisResultResponse) GoString() string {
	return s.String()
}

func (s *GetClusterDiagnosisResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClusterDiagnosisResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClusterDiagnosisResultResponse) GetBody() *GetClusterDiagnosisResultResponseBody {
	return s.Body
}

func (s *GetClusterDiagnosisResultResponse) SetHeaders(v map[string]*string) *GetClusterDiagnosisResultResponse {
	s.Headers = v
	return s
}

func (s *GetClusterDiagnosisResultResponse) SetStatusCode(v int32) *GetClusterDiagnosisResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterDiagnosisResultResponse) SetBody(v *GetClusterDiagnosisResultResponseBody) *GetClusterDiagnosisResultResponse {
	s.Body = v
	return s
}

func (s *GetClusterDiagnosisResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
