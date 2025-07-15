// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCasterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCasterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCasterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCasterResponseBody) *DeleteCasterResponse
	GetBody() *DeleteCasterResponseBody
}

type DeleteCasterResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCasterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCasterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCasterResponse) GoString() string {
	return s.String()
}

func (s *DeleteCasterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCasterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCasterResponse) GetBody() *DeleteCasterResponseBody {
	return s.Body
}

func (s *DeleteCasterResponse) SetHeaders(v map[string]*string) *DeleteCasterResponse {
	s.Headers = v
	return s
}

func (s *DeleteCasterResponse) SetStatusCode(v int32) *DeleteCasterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCasterResponse) SetBody(v *DeleteCasterResponseBody) *DeleteCasterResponse {
	s.Body = v
	return s
}

func (s *DeleteCasterResponse) Validate() error {
	return dara.Validate(s)
}
