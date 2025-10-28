// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMethodsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMethodsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMethodsResponse
	GetStatusCode() *int32
	SetBody(v *ListMethodsResponseBody) *ListMethodsResponse
	GetBody() *ListMethodsResponseBody
}

type ListMethodsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMethodsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMethodsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMethodsResponse) GoString() string {
	return s.String()
}

func (s *ListMethodsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMethodsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMethodsResponse) GetBody() *ListMethodsResponseBody {
	return s.Body
}

func (s *ListMethodsResponse) SetHeaders(v map[string]*string) *ListMethodsResponse {
	s.Headers = v
	return s
}

func (s *ListMethodsResponse) SetStatusCode(v int32) *ListMethodsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMethodsResponse) SetBody(v *ListMethodsResponseBody) *ListMethodsResponse {
	s.Body = v
	return s
}

func (s *ListMethodsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
