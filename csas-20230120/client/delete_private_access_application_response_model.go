// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivateAccessApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePrivateAccessApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePrivateAccessApplicationResponse
	GetStatusCode() *int32
	SetBody(v *DeletePrivateAccessApplicationResponseBody) *DeletePrivateAccessApplicationResponse
	GetBody() *DeletePrivateAccessApplicationResponseBody
}

type DeletePrivateAccessApplicationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePrivateAccessApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePrivateAccessApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivateAccessApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeletePrivateAccessApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePrivateAccessApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePrivateAccessApplicationResponse) GetBody() *DeletePrivateAccessApplicationResponseBody {
	return s.Body
}

func (s *DeletePrivateAccessApplicationResponse) SetHeaders(v map[string]*string) *DeletePrivateAccessApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeletePrivateAccessApplicationResponse) SetStatusCode(v int32) *DeletePrivateAccessApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrivateAccessApplicationResponse) SetBody(v *DeletePrivateAccessApplicationResponseBody) *DeletePrivateAccessApplicationResponse {
	s.Body = v
	return s
}

func (s *DeletePrivateAccessApplicationResponse) Validate() error {
	return dara.Validate(s)
}
