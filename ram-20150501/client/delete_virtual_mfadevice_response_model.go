// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVirtualMFADeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVirtualMFADeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVirtualMFADeviceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVirtualMFADeviceResponseBody) *DeleteVirtualMFADeviceResponse
	GetBody() *DeleteVirtualMFADeviceResponseBody
}

type DeleteVirtualMFADeviceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVirtualMFADeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVirtualMFADeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVirtualMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteVirtualMFADeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVirtualMFADeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVirtualMFADeviceResponse) GetBody() *DeleteVirtualMFADeviceResponseBody {
	return s.Body
}

func (s *DeleteVirtualMFADeviceResponse) SetHeaders(v map[string]*string) *DeleteVirtualMFADeviceResponse {
	s.Headers = v
	return s
}

func (s *DeleteVirtualMFADeviceResponse) SetStatusCode(v int32) *DeleteVirtualMFADeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVirtualMFADeviceResponse) SetBody(v *DeleteVirtualMFADeviceResponseBody) *DeleteVirtualMFADeviceResponse {
	s.Body = v
	return s
}

func (s *DeleteVirtualMFADeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
