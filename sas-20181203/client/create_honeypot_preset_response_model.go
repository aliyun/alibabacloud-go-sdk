// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHoneypotPresetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHoneypotPresetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHoneypotPresetResponse
	GetStatusCode() *int32
	SetBody(v *CreateHoneypotPresetResponseBody) *CreateHoneypotPresetResponse
	GetBody() *CreateHoneypotPresetResponseBody
}

type CreateHoneypotPresetResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHoneypotPresetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHoneypotPresetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHoneypotPresetResponse) GoString() string {
	return s.String()
}

func (s *CreateHoneypotPresetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHoneypotPresetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHoneypotPresetResponse) GetBody() *CreateHoneypotPresetResponseBody {
	return s.Body
}

func (s *CreateHoneypotPresetResponse) SetHeaders(v map[string]*string) *CreateHoneypotPresetResponse {
	s.Headers = v
	return s
}

func (s *CreateHoneypotPresetResponse) SetStatusCode(v int32) *CreateHoneypotPresetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHoneypotPresetResponse) SetBody(v *CreateHoneypotPresetResponseBody) *CreateHoneypotPresetResponse {
	s.Body = v
	return s
}

func (s *CreateHoneypotPresetResponse) Validate() error {
	return dara.Validate(s)
}
