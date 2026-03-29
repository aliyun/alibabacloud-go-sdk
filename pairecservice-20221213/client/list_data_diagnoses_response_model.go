// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataDiagnosesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataDiagnosesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataDiagnosesResponse
	GetStatusCode() *int32
	SetBody(v *ListDataDiagnosesResponseBody) *ListDataDiagnosesResponse
	GetBody() *ListDataDiagnosesResponseBody
}

type ListDataDiagnosesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataDiagnosesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataDiagnosesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataDiagnosesResponse) GoString() string {
	return s.String()
}

func (s *ListDataDiagnosesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataDiagnosesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataDiagnosesResponse) GetBody() *ListDataDiagnosesResponseBody {
	return s.Body
}

func (s *ListDataDiagnosesResponse) SetHeaders(v map[string]*string) *ListDataDiagnosesResponse {
	s.Headers = v
	return s
}

func (s *ListDataDiagnosesResponse) SetStatusCode(v int32) *ListDataDiagnosesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataDiagnosesResponse) SetBody(v *ListDataDiagnosesResponseBody) *ListDataDiagnosesResponse {
	s.Body = v
	return s
}

func (s *ListDataDiagnosesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
