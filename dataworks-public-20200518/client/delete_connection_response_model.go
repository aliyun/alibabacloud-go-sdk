// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteConnectionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteConnectionResponseBody) *DeleteConnectionResponse
	GetBody() *DeleteConnectionResponseBody
}

type DeleteConnectionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteConnectionResponse) GoString() string {
	return s.String()
}

func (s *DeleteConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteConnectionResponse) GetBody() *DeleteConnectionResponseBody {
	return s.Body
}

func (s *DeleteConnectionResponse) SetHeaders(v map[string]*string) *DeleteConnectionResponse {
	s.Headers = v
	return s
}

func (s *DeleteConnectionResponse) SetStatusCode(v int32) *DeleteConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConnectionResponse) SetBody(v *DeleteConnectionResponseBody) *DeleteConnectionResponse {
	s.Body = v
	return s
}

func (s *DeleteConnectionResponse) Validate() error {
	return dara.Validate(s)
}
