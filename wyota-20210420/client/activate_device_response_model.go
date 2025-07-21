// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ActivateDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ActivateDeviceResponse
	GetStatusCode() *int32
	SetBody(v *ActivateDeviceResponseBody) *ActivateDeviceResponse
	GetBody() *ActivateDeviceResponseBody
}

type ActivateDeviceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActivateDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivateDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s ActivateDeviceResponse) GoString() string {
	return s.String()
}

func (s *ActivateDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ActivateDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ActivateDeviceResponse) GetBody() *ActivateDeviceResponseBody {
	return s.Body
}

func (s *ActivateDeviceResponse) SetHeaders(v map[string]*string) *ActivateDeviceResponse {
	s.Headers = v
	return s
}

func (s *ActivateDeviceResponse) SetStatusCode(v int32) *ActivateDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateDeviceResponse) SetBody(v *ActivateDeviceResponseBody) *ActivateDeviceResponse {
	s.Body = v
	return s
}

func (s *ActivateDeviceResponse) Validate() error {
	return dara.Validate(s)
}
