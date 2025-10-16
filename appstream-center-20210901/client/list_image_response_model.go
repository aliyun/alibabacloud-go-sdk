// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListImageResponse
	GetStatusCode() *int32
	SetBody(v *ListImageResponseBody) *ListImageResponse
	GetBody() *ListImageResponseBody
}

type ListImageResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListImageResponse) GoString() string {
	return s.String()
}

func (s *ListImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListImageResponse) GetBody() *ListImageResponseBody {
	return s.Body
}

func (s *ListImageResponse) SetHeaders(v map[string]*string) *ListImageResponse {
	s.Headers = v
	return s
}

func (s *ListImageResponse) SetStatusCode(v int32) *ListImageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImageResponse) SetBody(v *ListImageResponseBody) *ListImageResponse {
	s.Body = v
	return s
}

func (s *ListImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
