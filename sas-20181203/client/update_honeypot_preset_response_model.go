// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHoneypotPresetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHoneypotPresetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHoneypotPresetResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHoneypotPresetResponseBody) *UpdateHoneypotPresetResponse
	GetBody() *UpdateHoneypotPresetResponseBody
}

type UpdateHoneypotPresetResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHoneypotPresetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHoneypotPresetResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHoneypotPresetResponse) GoString() string {
	return s.String()
}

func (s *UpdateHoneypotPresetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHoneypotPresetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHoneypotPresetResponse) GetBody() *UpdateHoneypotPresetResponseBody {
	return s.Body
}

func (s *UpdateHoneypotPresetResponse) SetHeaders(v map[string]*string) *UpdateHoneypotPresetResponse {
	s.Headers = v
	return s
}

func (s *UpdateHoneypotPresetResponse) SetStatusCode(v int32) *UpdateHoneypotPresetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHoneypotPresetResponse) SetBody(v *UpdateHoneypotPresetResponseBody) *UpdateHoneypotPresetResponse {
	s.Body = v
	return s
}

func (s *UpdateHoneypotPresetResponse) Validate() error {
	return dara.Validate(s)
}
