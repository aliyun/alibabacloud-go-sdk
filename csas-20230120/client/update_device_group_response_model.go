// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeviceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDeviceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDeviceGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDeviceGroupResponseBody) *UpdateDeviceGroupResponse
	GetBody() *UpdateDeviceGroupResponseBody
}

type UpdateDeviceGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeviceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeviceGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeviceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDeviceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDeviceGroupResponse) GetBody() *UpdateDeviceGroupResponseBody {
	return s.Body
}

func (s *UpdateDeviceGroupResponse) SetHeaders(v map[string]*string) *UpdateDeviceGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeviceGroupResponse) SetStatusCode(v int32) *UpdateDeviceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDeviceGroupResponse) SetBody(v *UpdateDeviceGroupResponseBody) *UpdateDeviceGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateDeviceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
