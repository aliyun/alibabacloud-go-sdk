// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOnlineTestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOnlineTestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOnlineTestResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOnlineTestResponseBody) *DeleteOnlineTestResponse
	GetBody() *DeleteOnlineTestResponseBody
}

type DeleteOnlineTestResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOnlineTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOnlineTestResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOnlineTestResponse) GoString() string {
	return s.String()
}

func (s *DeleteOnlineTestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOnlineTestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOnlineTestResponse) GetBody() *DeleteOnlineTestResponseBody {
	return s.Body
}

func (s *DeleteOnlineTestResponse) SetHeaders(v map[string]*string) *DeleteOnlineTestResponse {
	s.Headers = v
	return s
}

func (s *DeleteOnlineTestResponse) SetStatusCode(v int32) *DeleteOnlineTestResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOnlineTestResponse) SetBody(v *DeleteOnlineTestResponseBody) *DeleteOnlineTestResponse {
	s.Body = v
	return s
}

func (s *DeleteOnlineTestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
