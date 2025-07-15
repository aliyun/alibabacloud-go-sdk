// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockVirtualMFADeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnlockVirtualMFADeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnlockVirtualMFADeviceResponse
	GetStatusCode() *int32
	SetBody(v *UnlockVirtualMFADeviceResponseBody) *UnlockVirtualMFADeviceResponse
	GetBody() *UnlockVirtualMFADeviceResponseBody
}

type UnlockVirtualMFADeviceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnlockVirtualMFADeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnlockVirtualMFADeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s UnlockVirtualMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *UnlockVirtualMFADeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnlockVirtualMFADeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnlockVirtualMFADeviceResponse) GetBody() *UnlockVirtualMFADeviceResponseBody {
	return s.Body
}

func (s *UnlockVirtualMFADeviceResponse) SetHeaders(v map[string]*string) *UnlockVirtualMFADeviceResponse {
	s.Headers = v
	return s
}

func (s *UnlockVirtualMFADeviceResponse) SetStatusCode(v int32) *UnlockVirtualMFADeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *UnlockVirtualMFADeviceResponse) SetBody(v *UnlockVirtualMFADeviceResponseBody) *UnlockVirtualMFADeviceResponse {
	s.Body = v
	return s
}

func (s *UnlockVirtualMFADeviceResponse) Validate() error {
	return dara.Validate(s)
}
