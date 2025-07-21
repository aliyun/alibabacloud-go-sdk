// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTenantDeviceOtaInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTenantDeviceOtaInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTenantDeviceOtaInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListTenantDeviceOtaInfoResponseBody) *ListTenantDeviceOtaInfoResponse
	GetBody() *ListTenantDeviceOtaInfoResponseBody
}

type ListTenantDeviceOtaInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTenantDeviceOtaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTenantDeviceOtaInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTenantDeviceOtaInfoResponse) GoString() string {
	return s.String()
}

func (s *ListTenantDeviceOtaInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTenantDeviceOtaInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTenantDeviceOtaInfoResponse) GetBody() *ListTenantDeviceOtaInfoResponseBody {
	return s.Body
}

func (s *ListTenantDeviceOtaInfoResponse) SetHeaders(v map[string]*string) *ListTenantDeviceOtaInfoResponse {
	s.Headers = v
	return s
}

func (s *ListTenantDeviceOtaInfoResponse) SetStatusCode(v int32) *ListTenantDeviceOtaInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTenantDeviceOtaInfoResponse) SetBody(v *ListTenantDeviceOtaInfoResponseBody) *ListTenantDeviceOtaInfoResponse {
	s.Body = v
	return s
}

func (s *ListTenantDeviceOtaInfoResponse) Validate() error {
	return dara.Validate(s)
}
