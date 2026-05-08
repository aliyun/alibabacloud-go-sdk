// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContextsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteContextsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteContextsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteContextsResponseBody) *DeleteContextsResponse
	GetBody() *DeleteContextsResponseBody
}

type DeleteContextsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteContextsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteContextsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextsResponse) GoString() string {
	return s.String()
}

func (s *DeleteContextsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteContextsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteContextsResponse) GetBody() *DeleteContextsResponseBody {
	return s.Body
}

func (s *DeleteContextsResponse) SetHeaders(v map[string]*string) *DeleteContextsResponse {
	s.Headers = v
	return s
}

func (s *DeleteContextsResponse) SetStatusCode(v int32) *DeleteContextsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContextsResponse) SetBody(v *DeleteContextsResponseBody) *DeleteContextsResponse {
	s.Body = v
	return s
}

func (s *DeleteContextsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
