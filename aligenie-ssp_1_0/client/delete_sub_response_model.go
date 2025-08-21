// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSubResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSubResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSubResponseBody) *DeleteSubResponse
	GetBody() *DeleteSubResponseBody
}

type DeleteSubResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSubResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSubResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubResponse) GoString() string {
	return s.String()
}

func (s *DeleteSubResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSubResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSubResponse) GetBody() *DeleteSubResponseBody {
	return s.Body
}

func (s *DeleteSubResponse) SetHeaders(v map[string]*string) *DeleteSubResponse {
	s.Headers = v
	return s
}

func (s *DeleteSubResponse) SetStatusCode(v int32) *DeleteSubResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSubResponse) SetBody(v *DeleteSubResponseBody) *DeleteSubResponse {
	s.Body = v
	return s
}

func (s *DeleteSubResponse) Validate() error {
	return dara.Validate(s)
}
