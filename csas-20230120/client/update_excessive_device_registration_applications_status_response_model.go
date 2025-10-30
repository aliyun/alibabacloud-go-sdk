// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExcessiveDeviceRegistrationApplicationsStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse
	GetBody() *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody
}

type UpdateExcessiveDeviceRegistrationApplicationsStatusResponse struct {
	Headers    map[string]*string                                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExcessiveDeviceRegistrationApplicationsStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateExcessiveDeviceRegistrationApplicationsStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse) GetBody() *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody {
	return s.Body
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse) SetHeaders(v map[string]*string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse) SetStatusCode(v int32) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse) SetBody(v *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
