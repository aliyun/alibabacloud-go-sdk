// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetDeviceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNetDeviceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNetDeviceInfoResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNetDeviceInfoResponseBody) *UpdateNetDeviceInfoResponse
	GetBody() *UpdateNetDeviceInfoResponseBody
}

type UpdateNetDeviceInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNetDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNetDeviceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateNetDeviceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNetDeviceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNetDeviceInfoResponse) GetBody() *UpdateNetDeviceInfoResponseBody {
	return s.Body
}

func (s *UpdateNetDeviceInfoResponse) SetHeaders(v map[string]*string) *UpdateNetDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateNetDeviceInfoResponse) SetStatusCode(v int32) *UpdateNetDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNetDeviceInfoResponse) SetBody(v *UpdateNetDeviceInfoResponseBody) *UpdateNetDeviceInfoResponse {
	s.Body = v
	return s
}

func (s *UpdateNetDeviceInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
