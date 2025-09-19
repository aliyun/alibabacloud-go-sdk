// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHoneypotPresetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHoneypotPresetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHoneypotPresetResponse
	GetStatusCode() *int32
	SetBody(v *GetHoneypotPresetResponseBody) *GetHoneypotPresetResponse
	GetBody() *GetHoneypotPresetResponseBody
}

type GetHoneypotPresetResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHoneypotPresetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHoneypotPresetResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotPresetResponse) GoString() string {
	return s.String()
}

func (s *GetHoneypotPresetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHoneypotPresetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHoneypotPresetResponse) GetBody() *GetHoneypotPresetResponseBody {
	return s.Body
}

func (s *GetHoneypotPresetResponse) SetHeaders(v map[string]*string) *GetHoneypotPresetResponse {
	s.Headers = v
	return s
}

func (s *GetHoneypotPresetResponse) SetStatusCode(v int32) *GetHoneypotPresetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHoneypotPresetResponse) SetBody(v *GetHoneypotPresetResponseBody) *GetHoneypotPresetResponse {
	s.Body = v
	return s
}

func (s *GetHoneypotPresetResponse) Validate() error {
	return dara.Validate(s)
}
