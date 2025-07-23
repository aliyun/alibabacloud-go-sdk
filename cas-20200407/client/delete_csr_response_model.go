// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCsrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCsrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCsrResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCsrResponseBody) *DeleteCsrResponse
	GetBody() *DeleteCsrResponseBody
}

type DeleteCsrResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCsrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCsrResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCsrResponse) GoString() string {
	return s.String()
}

func (s *DeleteCsrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCsrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCsrResponse) GetBody() *DeleteCsrResponseBody {
	return s.Body
}

func (s *DeleteCsrResponse) SetHeaders(v map[string]*string) *DeleteCsrResponse {
	s.Headers = v
	return s
}

func (s *DeleteCsrResponse) SetStatusCode(v int32) *DeleteCsrResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCsrResponse) SetBody(v *DeleteCsrResponseBody) *DeleteCsrResponse {
	s.Body = v
	return s
}

func (s *DeleteCsrResponse) Validate() error {
	return dara.Validate(s)
}
