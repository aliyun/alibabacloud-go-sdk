// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHoneypotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHoneypotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHoneypotResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHoneypotResponseBody) *DeleteHoneypotResponse
	GetBody() *DeleteHoneypotResponseBody
}

type DeleteHoneypotResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHoneypotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHoneypotResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHoneypotResponse) GoString() string {
	return s.String()
}

func (s *DeleteHoneypotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHoneypotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHoneypotResponse) GetBody() *DeleteHoneypotResponseBody {
	return s.Body
}

func (s *DeleteHoneypotResponse) SetHeaders(v map[string]*string) *DeleteHoneypotResponse {
	s.Headers = v
	return s
}

func (s *DeleteHoneypotResponse) SetStatusCode(v int32) *DeleteHoneypotResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHoneypotResponse) SetBody(v *DeleteHoneypotResponseBody) *DeleteHoneypotResponse {
	s.Body = v
	return s
}

func (s *DeleteHoneypotResponse) Validate() error {
	return dara.Validate(s)
}
