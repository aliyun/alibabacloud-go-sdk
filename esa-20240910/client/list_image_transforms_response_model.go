// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageTransformsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListImageTransformsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListImageTransformsResponse
	GetStatusCode() *int32
	SetBody(v *ListImageTransformsResponseBody) *ListImageTransformsResponse
	GetBody() *ListImageTransformsResponseBody
}

type ListImageTransformsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImageTransformsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImageTransformsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListImageTransformsResponse) GoString() string {
	return s.String()
}

func (s *ListImageTransformsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListImageTransformsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListImageTransformsResponse) GetBody() *ListImageTransformsResponseBody {
	return s.Body
}

func (s *ListImageTransformsResponse) SetHeaders(v map[string]*string) *ListImageTransformsResponse {
	s.Headers = v
	return s
}

func (s *ListImageTransformsResponse) SetStatusCode(v int32) *ListImageTransformsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImageTransformsResponse) SetBody(v *ListImageTransformsResponseBody) *ListImageTransformsResponse {
	s.Body = v
	return s
}

func (s *ListImageTransformsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
