// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirtualMFADevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVirtualMFADevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVirtualMFADevicesResponse
	GetStatusCode() *int32
	SetBody(v *ListVirtualMFADevicesResponseBody) *ListVirtualMFADevicesResponse
	GetBody() *ListVirtualMFADevicesResponseBody
}

type ListVirtualMFADevicesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVirtualMFADevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVirtualMFADevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualMFADevicesResponse) GoString() string {
	return s.String()
}

func (s *ListVirtualMFADevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVirtualMFADevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVirtualMFADevicesResponse) GetBody() *ListVirtualMFADevicesResponseBody {
	return s.Body
}

func (s *ListVirtualMFADevicesResponse) SetHeaders(v map[string]*string) *ListVirtualMFADevicesResponse {
	s.Headers = v
	return s
}

func (s *ListVirtualMFADevicesResponse) SetStatusCode(v int32) *ListVirtualMFADevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVirtualMFADevicesResponse) SetBody(v *ListVirtualMFADevicesResponseBody) *ListVirtualMFADevicesResponse {
	s.Body = v
	return s
}

func (s *ListVirtualMFADevicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
