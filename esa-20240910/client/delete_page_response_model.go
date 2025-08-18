// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePageResponse
	GetStatusCode() *int32
	SetBody(v *DeletePageResponseBody) *DeletePageResponse
	GetBody() *DeletePageResponseBody
}

type DeletePageResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePageResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePageResponse) GoString() string {
	return s.String()
}

func (s *DeletePageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePageResponse) GetBody() *DeletePageResponseBody {
	return s.Body
}

func (s *DeletePageResponse) SetHeaders(v map[string]*string) *DeletePageResponse {
	s.Headers = v
	return s
}

func (s *DeletePageResponse) SetStatusCode(v int32) *DeletePageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePageResponse) SetBody(v *DeletePageResponseBody) *DeletePageResponse {
	s.Body = v
	return s
}

func (s *DeletePageResponse) Validate() error {
	return dara.Validate(s)
}
