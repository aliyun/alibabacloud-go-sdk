// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDBResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDBResponseBody) *DeleteDBResponse
	GetBody() *DeleteDBResponseBody
}

type DeleteDBResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDBResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDBResponse) GetBody() *DeleteDBResponseBody {
	return s.Body
}

func (s *DeleteDBResponse) SetHeaders(v map[string]*string) *DeleteDBResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBResponse) SetStatusCode(v int32) *DeleteDBResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBResponse) SetBody(v *DeleteDBResponseBody) *DeleteDBResponse {
	s.Body = v
	return s
}

func (s *DeleteDBResponse) Validate() error {
	return dara.Validate(s)
}
