// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSoftwareForUserDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSoftwareForUserDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSoftwareForUserDeviceResponse
	GetStatusCode() *int32
	SetBody(v *ListSoftwareForUserDeviceResponseBody) *ListSoftwareForUserDeviceResponse
	GetBody() *ListSoftwareForUserDeviceResponseBody
}

type ListSoftwareForUserDeviceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSoftwareForUserDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSoftwareForUserDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSoftwareForUserDeviceResponse) GoString() string {
	return s.String()
}

func (s *ListSoftwareForUserDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSoftwareForUserDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSoftwareForUserDeviceResponse) GetBody() *ListSoftwareForUserDeviceResponseBody {
	return s.Body
}

func (s *ListSoftwareForUserDeviceResponse) SetHeaders(v map[string]*string) *ListSoftwareForUserDeviceResponse {
	s.Headers = v
	return s
}

func (s *ListSoftwareForUserDeviceResponse) SetStatusCode(v int32) *ListSoftwareForUserDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSoftwareForUserDeviceResponse) SetBody(v *ListSoftwareForUserDeviceResponseBody) *ListSoftwareForUserDeviceResponse {
	s.Body = v
	return s
}

func (s *ListSoftwareForUserDeviceResponse) Validate() error {
	return dara.Validate(s)
}
