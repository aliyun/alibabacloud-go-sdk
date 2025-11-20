// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteClientResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteClientResponse
	GetStatusCode() *int32
	SetBody(v *DeleteClientResponseBody) *DeleteClientResponse
	GetBody() *DeleteClientResponseBody
}

type DeleteClientResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteClientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteClientResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientResponse) GoString() string {
	return s.String()
}

func (s *DeleteClientResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteClientResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteClientResponse) GetBody() *DeleteClientResponseBody {
	return s.Body
}

func (s *DeleteClientResponse) SetHeaders(v map[string]*string) *DeleteClientResponse {
	s.Headers = v
	return s
}

func (s *DeleteClientResponse) SetStatusCode(v int32) *DeleteClientResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClientResponse) SetBody(v *DeleteClientResponseBody) *DeleteClientResponse {
	s.Body = v
	return s
}

func (s *DeleteClientResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
