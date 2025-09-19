// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotPresetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHoneypotPresetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHoneypotPresetResponse
	GetStatusCode() *int32
	SetBody(v *ListHoneypotPresetResponseBody) *ListHoneypotPresetResponse
	GetBody() *ListHoneypotPresetResponseBody
}

type ListHoneypotPresetResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHoneypotPresetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHoneypotPresetResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotPresetResponse) GoString() string {
	return s.String()
}

func (s *ListHoneypotPresetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHoneypotPresetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHoneypotPresetResponse) GetBody() *ListHoneypotPresetResponseBody {
	return s.Body
}

func (s *ListHoneypotPresetResponse) SetHeaders(v map[string]*string) *ListHoneypotPresetResponse {
	s.Headers = v
	return s
}

func (s *ListHoneypotPresetResponse) SetStatusCode(v int32) *ListHoneypotPresetResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHoneypotPresetResponse) SetBody(v *ListHoneypotPresetResponseBody) *ListHoneypotPresetResponse {
	s.Body = v
	return s
}

func (s *ListHoneypotPresetResponse) Validate() error {
	return dara.Validate(s)
}
