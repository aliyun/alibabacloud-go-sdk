// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteErResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteErResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteErResponse
	GetStatusCode() *int32
	SetBody(v *DeleteErResponseBody) *DeleteErResponse
	GetBody() *DeleteErResponseBody
}

type DeleteErResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteErResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteErResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteErResponse) GoString() string {
	return s.String()
}

func (s *DeleteErResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteErResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteErResponse) GetBody() *DeleteErResponseBody {
	return s.Body
}

func (s *DeleteErResponse) SetHeaders(v map[string]*string) *DeleteErResponse {
	s.Headers = v
	return s
}

func (s *DeleteErResponse) SetStatusCode(v int32) *DeleteErResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteErResponse) SetBody(v *DeleteErResponseBody) *DeleteErResponse {
	s.Body = v
	return s
}

func (s *DeleteErResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
