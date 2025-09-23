// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceShareResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteResourceShareResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteResourceShareResponse
	GetStatusCode() *int32
	SetBody(v *DeleteResourceShareResponseBody) *DeleteResourceShareResponse
	GetBody() *DeleteResourceShareResponseBody
}

type DeleteResourceShareResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceShareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceShareResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceShareResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceShareResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteResourceShareResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteResourceShareResponse) GetBody() *DeleteResourceShareResponseBody {
	return s.Body
}

func (s *DeleteResourceShareResponse) SetHeaders(v map[string]*string) *DeleteResourceShareResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceShareResponse) SetStatusCode(v int32) *DeleteResourceShareResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceShareResponse) SetBody(v *DeleteResourceShareResponseBody) *DeleteResourceShareResponse {
	s.Body = v
	return s
}

func (s *DeleteResourceShareResponse) Validate() error {
	return dara.Validate(s)
}
