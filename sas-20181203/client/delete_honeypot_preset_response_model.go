// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHoneypotPresetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHoneypotPresetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHoneypotPresetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHoneypotPresetResponseBody) *DeleteHoneypotPresetResponse
	GetBody() *DeleteHoneypotPresetResponseBody
}

type DeleteHoneypotPresetResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHoneypotPresetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHoneypotPresetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHoneypotPresetResponse) GoString() string {
	return s.String()
}

func (s *DeleteHoneypotPresetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHoneypotPresetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHoneypotPresetResponse) GetBody() *DeleteHoneypotPresetResponseBody {
	return s.Body
}

func (s *DeleteHoneypotPresetResponse) SetHeaders(v map[string]*string) *DeleteHoneypotPresetResponse {
	s.Headers = v
	return s
}

func (s *DeleteHoneypotPresetResponse) SetStatusCode(v int32) *DeleteHoneypotPresetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHoneypotPresetResponse) SetBody(v *DeleteHoneypotPresetResponseBody) *DeleteHoneypotPresetResponse {
	s.Body = v
	return s
}

func (s *DeleteHoneypotPresetResponse) Validate() error {
	return dara.Validate(s)
}
