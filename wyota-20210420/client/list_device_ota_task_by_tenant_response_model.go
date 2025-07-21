// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceOtaTaskByTenantResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeviceOtaTaskByTenantResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeviceOtaTaskByTenantResponse
	GetStatusCode() *int32
	SetBody(v *ListDeviceOtaTaskByTenantResponseBody) *ListDeviceOtaTaskByTenantResponse
	GetBody() *ListDeviceOtaTaskByTenantResponseBody
}

type ListDeviceOtaTaskByTenantResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeviceOtaTaskByTenantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeviceOtaTaskByTenantResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceOtaTaskByTenantResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceOtaTaskByTenantResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeviceOtaTaskByTenantResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeviceOtaTaskByTenantResponse) GetBody() *ListDeviceOtaTaskByTenantResponseBody {
	return s.Body
}

func (s *ListDeviceOtaTaskByTenantResponse) SetHeaders(v map[string]*string) *ListDeviceOtaTaskByTenantResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponse) SetStatusCode(v int32) *ListDeviceOtaTaskByTenantResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponse) SetBody(v *ListDeviceOtaTaskByTenantResponseBody) *ListDeviceOtaTaskByTenantResponse {
	s.Body = v
	return s
}

func (s *ListDeviceOtaTaskByTenantResponse) Validate() error {
	return dara.Validate(s)
}
