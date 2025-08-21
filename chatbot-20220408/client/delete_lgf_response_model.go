// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLgfResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLgfResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLgfResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLgfResponseBody) *DeleteLgfResponse
	GetBody() *DeleteLgfResponseBody
}

type DeleteLgfResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLgfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLgfResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLgfResponse) GoString() string {
	return s.String()
}

func (s *DeleteLgfResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLgfResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLgfResponse) GetBody() *DeleteLgfResponseBody {
	return s.Body
}

func (s *DeleteLgfResponse) SetHeaders(v map[string]*string) *DeleteLgfResponse {
	s.Headers = v
	return s
}

func (s *DeleteLgfResponse) SetStatusCode(v int32) *DeleteLgfResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLgfResponse) SetBody(v *DeleteLgfResponseBody) *DeleteLgfResponse {
	s.Body = v
	return s
}

func (s *DeleteLgfResponse) Validate() error {
	return dara.Validate(s)
}
