// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrustDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTrustDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTrustDevicesResponse
	GetStatusCode() *int32
	SetBody(v *ListTrustDevicesResponseBody) *ListTrustDevicesResponse
	GetBody() *ListTrustDevicesResponseBody
}

type ListTrustDevicesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrustDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrustDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTrustDevicesResponse) GoString() string {
	return s.String()
}

func (s *ListTrustDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTrustDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTrustDevicesResponse) GetBody() *ListTrustDevicesResponseBody {
	return s.Body
}

func (s *ListTrustDevicesResponse) SetHeaders(v map[string]*string) *ListTrustDevicesResponse {
	s.Headers = v
	return s
}

func (s *ListTrustDevicesResponse) SetStatusCode(v int32) *ListTrustDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrustDevicesResponse) SetBody(v *ListTrustDevicesResponseBody) *ListTrustDevicesResponse {
	s.Body = v
	return s
}

func (s *ListTrustDevicesResponse) Validate() error {
	return dara.Validate(s)
}
