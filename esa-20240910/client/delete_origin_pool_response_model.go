// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOriginPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOriginPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOriginPoolResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOriginPoolResponseBody) *DeleteOriginPoolResponse
	GetBody() *DeleteOriginPoolResponseBody
}

type DeleteOriginPoolResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOriginPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOriginPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOriginPoolResponse) GoString() string {
	return s.String()
}

func (s *DeleteOriginPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOriginPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOriginPoolResponse) GetBody() *DeleteOriginPoolResponseBody {
	return s.Body
}

func (s *DeleteOriginPoolResponse) SetHeaders(v map[string]*string) *DeleteOriginPoolResponse {
	s.Headers = v
	return s
}

func (s *DeleteOriginPoolResponse) SetStatusCode(v int32) *DeleteOriginPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOriginPoolResponse) SetBody(v *DeleteOriginPoolResponseBody) *DeleteOriginPoolResponse {
	s.Body = v
	return s
}

func (s *DeleteOriginPoolResponse) Validate() error {
	return dara.Validate(s)
}
