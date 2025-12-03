// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEaiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEaiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEaiResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEaiResponseBody) *DeleteEaiResponse
	GetBody() *DeleteEaiResponseBody
}

type DeleteEaiResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEaiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEaiResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEaiResponse) GoString() string {
	return s.String()
}

func (s *DeleteEaiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEaiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEaiResponse) GetBody() *DeleteEaiResponseBody {
	return s.Body
}

func (s *DeleteEaiResponse) SetHeaders(v map[string]*string) *DeleteEaiResponse {
	s.Headers = v
	return s
}

func (s *DeleteEaiResponse) SetStatusCode(v int32) *DeleteEaiResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEaiResponse) SetBody(v *DeleteEaiResponseBody) *DeleteEaiResponse {
	s.Body = v
	return s
}

func (s *DeleteEaiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
