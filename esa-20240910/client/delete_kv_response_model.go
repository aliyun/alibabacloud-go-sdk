// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKvResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteKvResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteKvResponse
	GetStatusCode() *int32
	SetBody(v *DeleteKvResponseBody) *DeleteKvResponse
	GetBody() *DeleteKvResponseBody
}

type DeleteKvResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKvResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteKvResponse) GoString() string {
	return s.String()
}

func (s *DeleteKvResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteKvResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteKvResponse) GetBody() *DeleteKvResponseBody {
	return s.Body
}

func (s *DeleteKvResponse) SetHeaders(v map[string]*string) *DeleteKvResponse {
	s.Headers = v
	return s
}

func (s *DeleteKvResponse) SetStatusCode(v int32) *DeleteKvResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKvResponse) SetBody(v *DeleteKvResponseBody) *DeleteKvResponse {
	s.Body = v
	return s
}

func (s *DeleteKvResponse) Validate() error {
	return dara.Validate(s)
}
