// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnlockDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnlockDeviceResponse
	GetStatusCode() *int32
	SetBody(v *UnlockDeviceResponseBody) *UnlockDeviceResponse
	GetBody() *UnlockDeviceResponseBody
}

type UnlockDeviceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnlockDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnlockDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s UnlockDeviceResponse) GoString() string {
	return s.String()
}

func (s *UnlockDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnlockDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnlockDeviceResponse) GetBody() *UnlockDeviceResponseBody {
	return s.Body
}

func (s *UnlockDeviceResponse) SetHeaders(v map[string]*string) *UnlockDeviceResponse {
	s.Headers = v
	return s
}

func (s *UnlockDeviceResponse) SetStatusCode(v int32) *UnlockDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *UnlockDeviceResponse) SetBody(v *UnlockDeviceResponseBody) *UnlockDeviceResponse {
	s.Body = v
	return s
}

func (s *UnlockDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
