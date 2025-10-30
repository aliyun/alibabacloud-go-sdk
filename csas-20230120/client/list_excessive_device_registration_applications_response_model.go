// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExcessiveDeviceRegistrationApplicationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExcessiveDeviceRegistrationApplicationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExcessiveDeviceRegistrationApplicationsResponse
	GetStatusCode() *int32
	SetBody(v *ListExcessiveDeviceRegistrationApplicationsResponseBody) *ListExcessiveDeviceRegistrationApplicationsResponse
	GetBody() *ListExcessiveDeviceRegistrationApplicationsResponseBody
}

type ListExcessiveDeviceRegistrationApplicationsResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExcessiveDeviceRegistrationApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExcessiveDeviceRegistrationApplicationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExcessiveDeviceRegistrationApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponse) GetBody() *ListExcessiveDeviceRegistrationApplicationsResponseBody {
	return s.Body
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponse) SetHeaders(v map[string]*string) *ListExcessiveDeviceRegistrationApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponse) SetStatusCode(v int32) *ListExcessiveDeviceRegistrationApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponse) SetBody(v *ListExcessiveDeviceRegistrationApplicationsResponseBody) *ListExcessiveDeviceRegistrationApplicationsResponse {
	s.Body = v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
