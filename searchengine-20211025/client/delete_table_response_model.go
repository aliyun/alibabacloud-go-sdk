// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTableResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTableResponseBody) *DeleteTableResponse
	GetBody() *DeleteTableResponseBody
}

type DeleteTableResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTableResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTableResponse) GoString() string {
	return s.String()
}

func (s *DeleteTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTableResponse) GetBody() *DeleteTableResponseBody {
	return s.Body
}

func (s *DeleteTableResponse) SetHeaders(v map[string]*string) *DeleteTableResponse {
	s.Headers = v
	return s
}

func (s *DeleteTableResponse) SetStatusCode(v int32) *DeleteTableResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTableResponse) SetBody(v *DeleteTableResponseBody) *DeleteTableResponse {
	s.Body = v
	return s
}

func (s *DeleteTableResponse) Validate() error {
	return dara.Validate(s)
}
