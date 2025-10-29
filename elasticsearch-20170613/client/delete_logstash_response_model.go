// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogstashResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLogstashResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLogstashResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLogstashResponseBody) *DeleteLogstashResponse
	GetBody() *DeleteLogstashResponseBody
}

type DeleteLogstashResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLogstashResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLogstashResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogstashResponse) GoString() string {
	return s.String()
}

func (s *DeleteLogstashResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLogstashResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLogstashResponse) GetBody() *DeleteLogstashResponseBody {
	return s.Body
}

func (s *DeleteLogstashResponse) SetHeaders(v map[string]*string) *DeleteLogstashResponse {
	s.Headers = v
	return s
}

func (s *DeleteLogstashResponse) SetStatusCode(v int32) *DeleteLogstashResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLogstashResponse) SetBody(v *DeleteLogstashResponseBody) *DeleteLogstashResponse {
	s.Body = v
	return s
}

func (s *DeleteLogstashResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
