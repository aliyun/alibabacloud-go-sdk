// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAllMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAllMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAllMessageResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAllMessageResponseBody) *DeleteAllMessageResponse
	GetBody() *DeleteAllMessageResponseBody
}

type DeleteAllMessageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAllMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAllMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAllMessageResponse) GoString() string {
	return s.String()
}

func (s *DeleteAllMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAllMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAllMessageResponse) GetBody() *DeleteAllMessageResponseBody {
	return s.Body
}

func (s *DeleteAllMessageResponse) SetHeaders(v map[string]*string) *DeleteAllMessageResponse {
	s.Headers = v
	return s
}

func (s *DeleteAllMessageResponse) SetStatusCode(v int32) *DeleteAllMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAllMessageResponse) SetBody(v *DeleteAllMessageResponseBody) *DeleteAllMessageResponse {
	s.Body = v
	return s
}

func (s *DeleteAllMessageResponse) Validate() error {
	return dara.Validate(s)
}
