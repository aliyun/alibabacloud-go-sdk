// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BookResponse
	GetStatusCode() *int32
	SetBody(v *BookResponseBody) *BookResponse
	GetBody() *BookResponseBody
}

type BookResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BookResponseBody  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BookResponse) String() string {
	return dara.Prettify(s)
}

func (s BookResponse) GoString() string {
	return s.String()
}

func (s *BookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BookResponse) GetBody() *BookResponseBody {
	return s.Body
}

func (s *BookResponse) SetHeaders(v map[string]*string) *BookResponse {
	s.Headers = v
	return s
}

func (s *BookResponse) SetStatusCode(v int32) *BookResponse {
	s.StatusCode = &v
	return s
}

func (s *BookResponse) SetBody(v *BookResponseBody) *BookResponse {
	s.Body = v
	return s
}

func (s *BookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
