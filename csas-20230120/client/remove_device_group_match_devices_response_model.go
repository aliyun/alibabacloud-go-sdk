// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDeviceGroupMatchDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveDeviceGroupMatchDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveDeviceGroupMatchDevicesResponse
	GetStatusCode() *int32
	SetBody(v *RemoveDeviceGroupMatchDevicesResponseBody) *RemoveDeviceGroupMatchDevicesResponse
	GetBody() *RemoveDeviceGroupMatchDevicesResponseBody
}

type RemoveDeviceGroupMatchDevicesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveDeviceGroupMatchDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveDeviceGroupMatchDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveDeviceGroupMatchDevicesResponse) GoString() string {
	return s.String()
}

func (s *RemoveDeviceGroupMatchDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveDeviceGroupMatchDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveDeviceGroupMatchDevicesResponse) GetBody() *RemoveDeviceGroupMatchDevicesResponseBody {
	return s.Body
}

func (s *RemoveDeviceGroupMatchDevicesResponse) SetHeaders(v map[string]*string) *RemoveDeviceGroupMatchDevicesResponse {
	s.Headers = v
	return s
}

func (s *RemoveDeviceGroupMatchDevicesResponse) SetStatusCode(v int32) *RemoveDeviceGroupMatchDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveDeviceGroupMatchDevicesResponse) SetBody(v *RemoveDeviceGroupMatchDevicesResponseBody) *RemoveDeviceGroupMatchDevicesResponse {
	s.Body = v
	return s
}

func (s *RemoveDeviceGroupMatchDevicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
