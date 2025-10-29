// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiagnosisItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDiagnosisItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDiagnosisItemsResponse
	GetStatusCode() *int32
	SetBody(v *ListDiagnosisItemsResponseBody) *ListDiagnosisItemsResponse
	GetBody() *ListDiagnosisItemsResponseBody
}

type ListDiagnosisItemsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDiagnosisItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDiagnosisItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnosisItemsResponse) GoString() string {
	return s.String()
}

func (s *ListDiagnosisItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDiagnosisItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDiagnosisItemsResponse) GetBody() *ListDiagnosisItemsResponseBody {
	return s.Body
}

func (s *ListDiagnosisItemsResponse) SetHeaders(v map[string]*string) *ListDiagnosisItemsResponse {
	s.Headers = v
	return s
}

func (s *ListDiagnosisItemsResponse) SetStatusCode(v int32) *ListDiagnosisItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDiagnosisItemsResponse) SetBody(v *ListDiagnosisItemsResponseBody) *ListDiagnosisItemsResponse {
	s.Body = v
	return s
}

func (s *ListDiagnosisItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
