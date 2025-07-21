// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeviceBindedEndUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDeviceBindedEndUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDeviceBindedEndUserResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDeviceBindedEndUserResponseBody) *UpdateDeviceBindedEndUserResponse
	GetBody() *UpdateDeviceBindedEndUserResponseBody
}

type UpdateDeviceBindedEndUserResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDeviceBindedEndUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeviceBindedEndUserResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeviceBindedEndUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeviceBindedEndUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDeviceBindedEndUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDeviceBindedEndUserResponse) GetBody() *UpdateDeviceBindedEndUserResponseBody {
	return s.Body
}

func (s *UpdateDeviceBindedEndUserResponse) SetHeaders(v map[string]*string) *UpdateDeviceBindedEndUserResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeviceBindedEndUserResponse) SetStatusCode(v int32) *UpdateDeviceBindedEndUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDeviceBindedEndUserResponse) SetBody(v *UpdateDeviceBindedEndUserResponseBody) *UpdateDeviceBindedEndUserResponse {
	s.Body = v
	return s
}

func (s *UpdateDeviceBindedEndUserResponse) Validate() error {
	return dara.Validate(s)
}
