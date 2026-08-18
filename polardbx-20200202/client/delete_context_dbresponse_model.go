// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContextDBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteContextDBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteContextDBResponse
	GetStatusCode() *int32
	SetBody(v *DeleteContextDBResponseBody) *DeleteContextDBResponse
	GetBody() *DeleteContextDBResponseBody
}

type DeleteContextDBResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteContextDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteContextDBResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextDBResponse) GoString() string {
	return s.String()
}

func (s *DeleteContextDBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteContextDBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteContextDBResponse) GetBody() *DeleteContextDBResponseBody {
	return s.Body
}

func (s *DeleteContextDBResponse) SetHeaders(v map[string]*string) *DeleteContextDBResponse {
	s.Headers = v
	return s
}

func (s *DeleteContextDBResponse) SetStatusCode(v int32) *DeleteContextDBResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContextDBResponse) SetBody(v *DeleteContextDBResponseBody) *DeleteContextDBResponse {
	s.Body = v
	return s
}

func (s *DeleteContextDBResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
