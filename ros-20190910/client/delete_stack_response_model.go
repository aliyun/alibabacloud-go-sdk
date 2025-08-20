// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStackResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStackResponseBody) *DeleteStackResponse
	GetBody() *DeleteStackResponseBody
}

type DeleteStackResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStackResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStackResponse) GoString() string {
	return s.String()
}

func (s *DeleteStackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStackResponse) GetBody() *DeleteStackResponseBody {
	return s.Body
}

func (s *DeleteStackResponse) SetHeaders(v map[string]*string) *DeleteStackResponse {
	s.Headers = v
	return s
}

func (s *DeleteStackResponse) SetStatusCode(v int32) *DeleteStackResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStackResponse) SetBody(v *DeleteStackResponseBody) *DeleteStackResponse {
	s.Body = v
	return s
}

func (s *DeleteStackResponse) Validate() error {
	return dara.Validate(s)
}
