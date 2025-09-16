// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePublicUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePublicUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePublicUrlResponse
	GetStatusCode() *int32
	SetBody(v *DeletePublicUrlResponseBody) *DeletePublicUrlResponse
	GetBody() *DeletePublicUrlResponseBody
}

type DeletePublicUrlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePublicUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePublicUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePublicUrlResponse) GoString() string {
	return s.String()
}

func (s *DeletePublicUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePublicUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePublicUrlResponse) GetBody() *DeletePublicUrlResponseBody {
	return s.Body
}

func (s *DeletePublicUrlResponse) SetHeaders(v map[string]*string) *DeletePublicUrlResponse {
	s.Headers = v
	return s
}

func (s *DeletePublicUrlResponse) SetStatusCode(v int32) *DeletePublicUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePublicUrlResponse) SetBody(v *DeletePublicUrlResponseBody) *DeletePublicUrlResponse {
	s.Body = v
	return s
}

func (s *DeletePublicUrlResponse) Validate() error {
	return dara.Validate(s)
}
