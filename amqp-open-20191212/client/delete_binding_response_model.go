// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBindingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBindingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBindingResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBindingResponseBody) *DeleteBindingResponse
	GetBody() *DeleteBindingResponseBody
}

type DeleteBindingResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBindingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBindingResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBindingResponse) GoString() string {
	return s.String()
}

func (s *DeleteBindingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBindingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBindingResponse) GetBody() *DeleteBindingResponseBody {
	return s.Body
}

func (s *DeleteBindingResponse) SetHeaders(v map[string]*string) *DeleteBindingResponse {
	s.Headers = v
	return s
}

func (s *DeleteBindingResponse) SetStatusCode(v int32) *DeleteBindingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBindingResponse) SetBody(v *DeleteBindingResponseBody) *DeleteBindingResponse {
	s.Body = v
	return s
}

func (s *DeleteBindingResponse) Validate() error {
	return dara.Validate(s)
}
