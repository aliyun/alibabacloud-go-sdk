// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHoneypotNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHoneypotNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHoneypotNodeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHoneypotNodeResponseBody) *UpdateHoneypotNodeResponse
	GetBody() *UpdateHoneypotNodeResponseBody
}

type UpdateHoneypotNodeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHoneypotNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHoneypotNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHoneypotNodeResponse) GoString() string {
	return s.String()
}

func (s *UpdateHoneypotNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHoneypotNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHoneypotNodeResponse) GetBody() *UpdateHoneypotNodeResponseBody {
	return s.Body
}

func (s *UpdateHoneypotNodeResponse) SetHeaders(v map[string]*string) *UpdateHoneypotNodeResponse {
	s.Headers = v
	return s
}

func (s *UpdateHoneypotNodeResponse) SetStatusCode(v int32) *UpdateHoneypotNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHoneypotNodeResponse) SetBody(v *UpdateHoneypotNodeResponseBody) *UpdateHoneypotNodeResponse {
	s.Body = v
	return s
}

func (s *UpdateHoneypotNodeResponse) Validate() error {
	return dara.Validate(s)
}
