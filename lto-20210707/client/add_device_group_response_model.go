// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDeviceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDeviceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDeviceGroupResponse
	GetStatusCode() *int32
	SetBody(v *AddDeviceGroupResponseBody) *AddDeviceGroupResponse
	GetBody() *AddDeviceGroupResponseBody
}

type AddDeviceGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDeviceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDeviceGroupResponse) GoString() string {
	return s.String()
}

func (s *AddDeviceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDeviceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDeviceGroupResponse) GetBody() *AddDeviceGroupResponseBody {
	return s.Body
}

func (s *AddDeviceGroupResponse) SetHeaders(v map[string]*string) *AddDeviceGroupResponse {
	s.Headers = v
	return s
}

func (s *AddDeviceGroupResponse) SetStatusCode(v int32) *AddDeviceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDeviceGroupResponse) SetBody(v *AddDeviceGroupResponseBody) *AddDeviceGroupResponse {
	s.Body = v
	return s
}

func (s *AddDeviceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
