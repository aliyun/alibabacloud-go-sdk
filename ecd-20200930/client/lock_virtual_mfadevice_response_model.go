// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLockVirtualMFADeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LockVirtualMFADeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LockVirtualMFADeviceResponse
	GetStatusCode() *int32
	SetBody(v *LockVirtualMFADeviceResponseBody) *LockVirtualMFADeviceResponse
	GetBody() *LockVirtualMFADeviceResponseBody
}

type LockVirtualMFADeviceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LockVirtualMFADeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LockVirtualMFADeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s LockVirtualMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *LockVirtualMFADeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LockVirtualMFADeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LockVirtualMFADeviceResponse) GetBody() *LockVirtualMFADeviceResponseBody {
	return s.Body
}

func (s *LockVirtualMFADeviceResponse) SetHeaders(v map[string]*string) *LockVirtualMFADeviceResponse {
	s.Headers = v
	return s
}

func (s *LockVirtualMFADeviceResponse) SetStatusCode(v int32) *LockVirtualMFADeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *LockVirtualMFADeviceResponse) SetBody(v *LockVirtualMFADeviceResponseBody) *LockVirtualMFADeviceResponse {
	s.Body = v
	return s
}

func (s *LockVirtualMFADeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
