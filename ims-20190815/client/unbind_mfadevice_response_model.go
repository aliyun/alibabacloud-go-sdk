// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindMFADeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindMFADeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindMFADeviceResponse
	GetStatusCode() *int32
	SetBody(v *UnbindMFADeviceResponseBody) *UnbindMFADeviceResponse
	GetBody() *UnbindMFADeviceResponseBody
}

type UnbindMFADeviceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindMFADeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindMFADeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *UnbindMFADeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindMFADeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindMFADeviceResponse) GetBody() *UnbindMFADeviceResponseBody {
	return s.Body
}

func (s *UnbindMFADeviceResponse) SetHeaders(v map[string]*string) *UnbindMFADeviceResponse {
	s.Headers = v
	return s
}

func (s *UnbindMFADeviceResponse) SetStatusCode(v int32) *UnbindMFADeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindMFADeviceResponse) SetBody(v *UnbindMFADeviceResponseBody) *UnbindMFADeviceResponse {
	s.Body = v
	return s
}

func (s *UnbindMFADeviceResponse) Validate() error {
	return dara.Validate(s)
}
