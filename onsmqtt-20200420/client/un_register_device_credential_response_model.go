// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnRegisterDeviceCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnRegisterDeviceCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnRegisterDeviceCredentialResponse
	GetStatusCode() *int32
	SetBody(v *UnRegisterDeviceCredentialResponseBody) *UnRegisterDeviceCredentialResponse
	GetBody() *UnRegisterDeviceCredentialResponseBody
}

type UnRegisterDeviceCredentialResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnRegisterDeviceCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnRegisterDeviceCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s UnRegisterDeviceCredentialResponse) GoString() string {
	return s.String()
}

func (s *UnRegisterDeviceCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnRegisterDeviceCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnRegisterDeviceCredentialResponse) GetBody() *UnRegisterDeviceCredentialResponseBody {
	return s.Body
}

func (s *UnRegisterDeviceCredentialResponse) SetHeaders(v map[string]*string) *UnRegisterDeviceCredentialResponse {
	s.Headers = v
	return s
}

func (s *UnRegisterDeviceCredentialResponse) SetStatusCode(v int32) *UnRegisterDeviceCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *UnRegisterDeviceCredentialResponse) SetBody(v *UnRegisterDeviceCredentialResponseBody) *UnRegisterDeviceCredentialResponse {
	s.Body = v
	return s
}

func (s *UnRegisterDeviceCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
