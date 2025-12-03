// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEaiAllResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEaiAllResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEaiAllResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEaiAllResponseBody) *DeleteEaiAllResponse
	GetBody() *DeleteEaiAllResponseBody
}

type DeleteEaiAllResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEaiAllResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEaiAllResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEaiAllResponse) GoString() string {
	return s.String()
}

func (s *DeleteEaiAllResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEaiAllResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEaiAllResponse) GetBody() *DeleteEaiAllResponseBody {
	return s.Body
}

func (s *DeleteEaiAllResponse) SetHeaders(v map[string]*string) *DeleteEaiAllResponse {
	s.Headers = v
	return s
}

func (s *DeleteEaiAllResponse) SetStatusCode(v int32) *DeleteEaiAllResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEaiAllResponse) SetBody(v *DeleteEaiAllResponseBody) *DeleteEaiAllResponse {
	s.Body = v
	return s
}

func (s *DeleteEaiAllResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
