// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDeviceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableDeviceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableDeviceGroupResponse
	GetStatusCode() *int32
	SetBody(v *DisableDeviceGroupResponseBody) *DisableDeviceGroupResponse
	GetBody() *DisableDeviceGroupResponseBody
}

type DisableDeviceGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableDeviceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableDeviceGroupResponse) GoString() string {
	return s.String()
}

func (s *DisableDeviceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableDeviceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableDeviceGroupResponse) GetBody() *DisableDeviceGroupResponseBody {
	return s.Body
}

func (s *DisableDeviceGroupResponse) SetHeaders(v map[string]*string) *DisableDeviceGroupResponse {
	s.Headers = v
	return s
}

func (s *DisableDeviceGroupResponse) SetStatusCode(v int32) *DisableDeviceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableDeviceGroupResponse) SetBody(v *DisableDeviceGroupResponseBody) *DisableDeviceGroupResponse {
	s.Body = v
	return s
}

func (s *DisableDeviceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
