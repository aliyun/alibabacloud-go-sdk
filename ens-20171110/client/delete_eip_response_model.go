// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEipResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEipResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEipResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEipResponseBody) *DeleteEipResponse
	GetBody() *DeleteEipResponseBody
}

type DeleteEipResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEipResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEipResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEipResponse) GoString() string {
	return s.String()
}

func (s *DeleteEipResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEipResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEipResponse) GetBody() *DeleteEipResponseBody {
	return s.Body
}

func (s *DeleteEipResponse) SetHeaders(v map[string]*string) *DeleteEipResponse {
	s.Headers = v
	return s
}

func (s *DeleteEipResponse) SetStatusCode(v int32) *DeleteEipResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEipResponse) SetBody(v *DeleteEipResponseBody) *DeleteEipResponse {
	s.Body = v
	return s
}

func (s *DeleteEipResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
