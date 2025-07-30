// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockMfaDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnlockMfaDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnlockMfaDeviceResponse
	GetStatusCode() *int32
	SetBody(v *UnlockMfaDeviceResponseBody) *UnlockMfaDeviceResponse
	GetBody() *UnlockMfaDeviceResponseBody
}

type UnlockMfaDeviceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnlockMfaDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnlockMfaDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s UnlockMfaDeviceResponse) GoString() string {
	return s.String()
}

func (s *UnlockMfaDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnlockMfaDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnlockMfaDeviceResponse) GetBody() *UnlockMfaDeviceResponseBody {
	return s.Body
}

func (s *UnlockMfaDeviceResponse) SetHeaders(v map[string]*string) *UnlockMfaDeviceResponse {
	s.Headers = v
	return s
}

func (s *UnlockMfaDeviceResponse) SetStatusCode(v int32) *UnlockMfaDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *UnlockMfaDeviceResponse) SetBody(v *UnlockMfaDeviceResponseBody) *UnlockMfaDeviceResponse {
	s.Body = v
	return s
}

func (s *UnlockMfaDeviceResponse) Validate() error {
	return dara.Validate(s)
}
