// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirtualMFADeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVirtualMFADeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVirtualMFADeviceResponse
	GetStatusCode() *int32
	SetBody(v *CreateVirtualMFADeviceResponseBody) *CreateVirtualMFADeviceResponse
	GetBody() *CreateVirtualMFADeviceResponseBody
}

type CreateVirtualMFADeviceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVirtualMFADeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVirtualMFADeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVirtualMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *CreateVirtualMFADeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVirtualMFADeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVirtualMFADeviceResponse) GetBody() *CreateVirtualMFADeviceResponseBody {
	return s.Body
}

func (s *CreateVirtualMFADeviceResponse) SetHeaders(v map[string]*string) *CreateVirtualMFADeviceResponse {
	s.Headers = v
	return s
}

func (s *CreateVirtualMFADeviceResponse) SetStatusCode(v int32) *CreateVirtualMFADeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVirtualMFADeviceResponse) SetBody(v *CreateVirtualMFADeviceResponseBody) *CreateVirtualMFADeviceResponse {
	s.Body = v
	return s
}

func (s *CreateVirtualMFADeviceResponse) Validate() error {
	return dara.Validate(s)
}
