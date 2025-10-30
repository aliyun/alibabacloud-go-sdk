// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserDevicesSharingStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUserDevicesSharingStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUserDevicesSharingStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUserDevicesSharingStatusResponseBody) *UpdateUserDevicesSharingStatusResponse
	GetBody() *UpdateUserDevicesSharingStatusResponseBody
}

type UpdateUserDevicesSharingStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserDevicesSharingStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserDevicesSharingStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserDevicesSharingStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserDevicesSharingStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUserDevicesSharingStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUserDevicesSharingStatusResponse) GetBody() *UpdateUserDevicesSharingStatusResponseBody {
	return s.Body
}

func (s *UpdateUserDevicesSharingStatusResponse) SetHeaders(v map[string]*string) *UpdateUserDevicesSharingStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponse) SetStatusCode(v int32) *UpdateUserDevicesSharingStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponse) SetBody(v *UpdateUserDevicesSharingStatusResponseBody) *UpdateUserDevicesSharingStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
