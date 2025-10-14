// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSDGResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSDGResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSDGResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSDGResponseBody) *DeleteSDGResponse
	GetBody() *DeleteSDGResponseBody
}

type DeleteSDGResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSDGResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSDGResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSDGResponse) GoString() string {
	return s.String()
}

func (s *DeleteSDGResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSDGResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSDGResponse) GetBody() *DeleteSDGResponseBody {
	return s.Body
}

func (s *DeleteSDGResponse) SetHeaders(v map[string]*string) *DeleteSDGResponse {
	s.Headers = v
	return s
}

func (s *DeleteSDGResponse) SetStatusCode(v int32) *DeleteSDGResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSDGResponse) SetBody(v *DeleteSDGResponseBody) *DeleteSDGResponse {
	s.Body = v
	return s
}

func (s *DeleteSDGResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
