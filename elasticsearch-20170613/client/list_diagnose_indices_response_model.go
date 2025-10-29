// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiagnoseIndicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDiagnoseIndicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDiagnoseIndicesResponse
	GetStatusCode() *int32
	SetBody(v *ListDiagnoseIndicesResponseBody) *ListDiagnoseIndicesResponse
	GetBody() *ListDiagnoseIndicesResponseBody
}

type ListDiagnoseIndicesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDiagnoseIndicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDiagnoseIndicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnoseIndicesResponse) GoString() string {
	return s.String()
}

func (s *ListDiagnoseIndicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDiagnoseIndicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDiagnoseIndicesResponse) GetBody() *ListDiagnoseIndicesResponseBody {
	return s.Body
}

func (s *ListDiagnoseIndicesResponse) SetHeaders(v map[string]*string) *ListDiagnoseIndicesResponse {
	s.Headers = v
	return s
}

func (s *ListDiagnoseIndicesResponse) SetStatusCode(v int32) *ListDiagnoseIndicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDiagnoseIndicesResponse) SetBody(v *ListDiagnoseIndicesResponseBody) *ListDiagnoseIndicesResponse {
	s.Body = v
	return s
}

func (s *ListDiagnoseIndicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
