// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceBasicInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeviceBasicInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeviceBasicInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListDeviceBasicInfoResponseBody) *ListDeviceBasicInfoResponse
	GetBody() *ListDeviceBasicInfoResponseBody
}

type ListDeviceBasicInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeviceBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeviceBasicInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceBasicInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeviceBasicInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeviceBasicInfoResponse) GetBody() *ListDeviceBasicInfoResponseBody {
	return s.Body
}

func (s *ListDeviceBasicInfoResponse) SetHeaders(v map[string]*string) *ListDeviceBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceBasicInfoResponse) SetStatusCode(v int32) *ListDeviceBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceBasicInfoResponse) SetBody(v *ListDeviceBasicInfoResponseBody) *ListDeviceBasicInfoResponse {
	s.Body = v
	return s
}

func (s *ListDeviceBasicInfoResponse) Validate() error {
	return dara.Validate(s)
}
