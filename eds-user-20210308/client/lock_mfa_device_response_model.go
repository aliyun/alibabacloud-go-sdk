// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLockMfaDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LockMfaDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LockMfaDeviceResponse
	GetStatusCode() *int32
	SetBody(v *LockMfaDeviceResponseBody) *LockMfaDeviceResponse
	GetBody() *LockMfaDeviceResponseBody
}

type LockMfaDeviceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LockMfaDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LockMfaDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s LockMfaDeviceResponse) GoString() string {
	return s.String()
}

func (s *LockMfaDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LockMfaDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LockMfaDeviceResponse) GetBody() *LockMfaDeviceResponseBody {
	return s.Body
}

func (s *LockMfaDeviceResponse) SetHeaders(v map[string]*string) *LockMfaDeviceResponse {
	s.Headers = v
	return s
}

func (s *LockMfaDeviceResponse) SetStatusCode(v int32) *LockMfaDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *LockMfaDeviceResponse) SetBody(v *LockMfaDeviceResponseBody) *LockMfaDeviceResponse {
	s.Body = v
	return s
}

func (s *LockMfaDeviceResponse) Validate() error {
	return dara.Validate(s)
}
