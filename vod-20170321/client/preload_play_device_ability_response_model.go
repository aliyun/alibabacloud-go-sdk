// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreloadPlayDeviceAbilityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreloadPlayDeviceAbilityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreloadPlayDeviceAbilityResponse
	GetStatusCode() *int32
	SetBody(v *PreloadPlayDeviceAbilityResponseBody) *PreloadPlayDeviceAbilityResponse
	GetBody() *PreloadPlayDeviceAbilityResponseBody
}

type PreloadPlayDeviceAbilityResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreloadPlayDeviceAbilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreloadPlayDeviceAbilityResponse) String() string {
	return dara.Prettify(s)
}

func (s PreloadPlayDeviceAbilityResponse) GoString() string {
	return s.String()
}

func (s *PreloadPlayDeviceAbilityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreloadPlayDeviceAbilityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreloadPlayDeviceAbilityResponse) GetBody() *PreloadPlayDeviceAbilityResponseBody {
	return s.Body
}

func (s *PreloadPlayDeviceAbilityResponse) SetHeaders(v map[string]*string) *PreloadPlayDeviceAbilityResponse {
	s.Headers = v
	return s
}

func (s *PreloadPlayDeviceAbilityResponse) SetStatusCode(v int32) *PreloadPlayDeviceAbilityResponse {
	s.StatusCode = &v
	return s
}

func (s *PreloadPlayDeviceAbilityResponse) SetBody(v *PreloadPlayDeviceAbilityResponseBody) *PreloadPlayDeviceAbilityResponse {
	s.Body = v
	return s
}

func (s *PreloadPlayDeviceAbilityResponse) Validate() error {
	return dara.Validate(s)
}
