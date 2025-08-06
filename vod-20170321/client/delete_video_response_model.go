// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVideoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVideoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVideoResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVideoResponseBody) *DeleteVideoResponse
	GetBody() *DeleteVideoResponseBody
}

type DeleteVideoResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVideoResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVideoResponse) GoString() string {
	return s.String()
}

func (s *DeleteVideoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVideoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVideoResponse) GetBody() *DeleteVideoResponseBody {
	return s.Body
}

func (s *DeleteVideoResponse) SetHeaders(v map[string]*string) *DeleteVideoResponse {
	s.Headers = v
	return s
}

func (s *DeleteVideoResponse) SetStatusCode(v int32) *DeleteVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVideoResponse) SetBody(v *DeleteVideoResponseBody) *DeleteVideoResponse {
	s.Body = v
	return s
}

func (s *DeleteVideoResponse) Validate() error {
	return dara.Validate(s)
}
