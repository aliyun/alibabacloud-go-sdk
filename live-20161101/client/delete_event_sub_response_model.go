// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventSubResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEventSubResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEventSubResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEventSubResponseBody) *DeleteEventSubResponse
	GetBody() *DeleteEventSubResponseBody
}

type DeleteEventSubResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEventSubResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEventSubResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventSubResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventSubResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEventSubResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEventSubResponse) GetBody() *DeleteEventSubResponseBody {
	return s.Body
}

func (s *DeleteEventSubResponse) SetHeaders(v map[string]*string) *DeleteEventSubResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventSubResponse) SetStatusCode(v int32) *DeleteEventSubResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEventSubResponse) SetBody(v *DeleteEventSubResponseBody) *DeleteEventSubResponse {
	s.Body = v
	return s
}

func (s *DeleteEventSubResponse) Validate() error {
	return dara.Validate(s)
}
