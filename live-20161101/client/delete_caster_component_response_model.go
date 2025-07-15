// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCasterComponentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCasterComponentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCasterComponentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCasterComponentResponseBody) *DeleteCasterComponentResponse
	GetBody() *DeleteCasterComponentResponseBody
}

type DeleteCasterComponentResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCasterComponentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCasterComponentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCasterComponentResponse) GoString() string {
	return s.String()
}

func (s *DeleteCasterComponentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCasterComponentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCasterComponentResponse) GetBody() *DeleteCasterComponentResponseBody {
	return s.Body
}

func (s *DeleteCasterComponentResponse) SetHeaders(v map[string]*string) *DeleteCasterComponentResponse {
	s.Headers = v
	return s
}

func (s *DeleteCasterComponentResponse) SetStatusCode(v int32) *DeleteCasterComponentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCasterComponentResponse) SetBody(v *DeleteCasterComponentResponseBody) *DeleteCasterComponentResponse {
	s.Body = v
	return s
}

func (s *DeleteCasterComponentResponse) Validate() error {
	return dara.Validate(s)
}
