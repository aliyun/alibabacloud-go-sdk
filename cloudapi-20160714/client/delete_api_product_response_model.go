// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApiProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApiProductResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApiProductResponseBody) *DeleteApiProductResponse
	GetBody() *DeleteApiProductResponseBody
}

type DeleteApiProductResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApiProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApiProductResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiProductResponse) GoString() string {
	return s.String()
}

func (s *DeleteApiProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApiProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApiProductResponse) GetBody() *DeleteApiProductResponseBody {
	return s.Body
}

func (s *DeleteApiProductResponse) SetHeaders(v map[string]*string) *DeleteApiProductResponse {
	s.Headers = v
	return s
}

func (s *DeleteApiProductResponse) SetStatusCode(v int32) *DeleteApiProductResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApiProductResponse) SetBody(v *DeleteApiProductResponseBody) *DeleteApiProductResponse {
	s.Body = v
	return s
}

func (s *DeleteApiProductResponse) Validate() error {
	return dara.Validate(s)
}
