// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVhostResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVhostResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVhostResponse
	GetStatusCode() *int32
	SetBody(v *ListVhostResponseBody) *ListVhostResponse
	GetBody() *ListVhostResponseBody
}

type ListVhostResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVhostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVhostResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVhostResponse) GoString() string {
	return s.String()
}

func (s *ListVhostResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVhostResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVhostResponse) GetBody() *ListVhostResponseBody {
	return s.Body
}

func (s *ListVhostResponse) SetHeaders(v map[string]*string) *ListVhostResponse {
	s.Headers = v
	return s
}

func (s *ListVhostResponse) SetStatusCode(v int32) *ListVhostResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVhostResponse) SetBody(v *ListVhostResponseBody) *ListVhostResponse {
	s.Body = v
	return s
}

func (s *ListVhostResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
