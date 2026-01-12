// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLocationServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLocationServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLocationServiceResponse
	GetStatusCode() *int32
	SetBody(v *ListLocationServiceResponseBody) *ListLocationServiceResponse
	GetBody() *ListLocationServiceResponseBody
}

type ListLocationServiceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLocationServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLocationServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLocationServiceResponse) GoString() string {
	return s.String()
}

func (s *ListLocationServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLocationServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLocationServiceResponse) GetBody() *ListLocationServiceResponseBody {
	return s.Body
}

func (s *ListLocationServiceResponse) SetHeaders(v map[string]*string) *ListLocationServiceResponse {
	s.Headers = v
	return s
}

func (s *ListLocationServiceResponse) SetStatusCode(v int32) *ListLocationServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLocationServiceResponse) SetBody(v *ListLocationServiceResponseBody) *ListLocationServiceResponse {
	s.Body = v
	return s
}

func (s *ListLocationServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
