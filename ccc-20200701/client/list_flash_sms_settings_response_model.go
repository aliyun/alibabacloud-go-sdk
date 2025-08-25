// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlashSmsSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFlashSmsSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFlashSmsSettingsResponse
	GetStatusCode() *int32
	SetBody(v *ListFlashSmsSettingsResponseBody) *ListFlashSmsSettingsResponse
	GetBody() *ListFlashSmsSettingsResponseBody
}

type ListFlashSmsSettingsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFlashSmsSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFlashSmsSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFlashSmsSettingsResponse) GoString() string {
	return s.String()
}

func (s *ListFlashSmsSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFlashSmsSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFlashSmsSettingsResponse) GetBody() *ListFlashSmsSettingsResponseBody {
	return s.Body
}

func (s *ListFlashSmsSettingsResponse) SetHeaders(v map[string]*string) *ListFlashSmsSettingsResponse {
	s.Headers = v
	return s
}

func (s *ListFlashSmsSettingsResponse) SetStatusCode(v int32) *ListFlashSmsSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFlashSmsSettingsResponse) SetBody(v *ListFlashSmsSettingsResponseBody) *ListFlashSmsSettingsResponse {
	s.Body = v
	return s
}

func (s *ListFlashSmsSettingsResponse) Validate() error {
	return dara.Validate(s)
}
