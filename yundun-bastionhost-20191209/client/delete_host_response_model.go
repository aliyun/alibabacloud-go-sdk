// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHostResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHostResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHostResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHostResponseBody) *DeleteHostResponse
	GetBody() *DeleteHostResponseBody
}

type DeleteHostResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHostResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHostResponse) GoString() string {
	return s.String()
}

func (s *DeleteHostResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHostResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHostResponse) GetBody() *DeleteHostResponseBody {
	return s.Body
}

func (s *DeleteHostResponse) SetHeaders(v map[string]*string) *DeleteHostResponse {
	s.Headers = v
	return s
}

func (s *DeleteHostResponse) SetStatusCode(v int32) *DeleteHostResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHostResponse) SetBody(v *DeleteHostResponseBody) *DeleteHostResponse {
	s.Body = v
	return s
}

func (s *DeleteHostResponse) Validate() error {
	return dara.Validate(s)
}
