// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteZnodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteZnodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteZnodeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteZnodeResponseBody) *DeleteZnodeResponse
	GetBody() *DeleteZnodeResponseBody
}

type DeleteZnodeResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteZnodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteZnodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteZnodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteZnodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteZnodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteZnodeResponse) GetBody() *DeleteZnodeResponseBody {
	return s.Body
}

func (s *DeleteZnodeResponse) SetHeaders(v map[string]*string) *DeleteZnodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteZnodeResponse) SetStatusCode(v int32) *DeleteZnodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteZnodeResponse) SetBody(v *DeleteZnodeResponseBody) *DeleteZnodeResponse {
	s.Body = v
	return s
}

func (s *DeleteZnodeResponse) Validate() error {
	return dara.Validate(s)
}
