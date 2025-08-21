// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceStatusInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeviceStatusInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeviceStatusInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetDeviceStatusInfoResponseBody) *GetDeviceStatusInfoResponse
	GetBody() *GetDeviceStatusInfoResponseBody
}

type GetDeviceStatusInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceStatusInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeviceStatusInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceStatusInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeviceStatusInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeviceStatusInfoResponse) GetBody() *GetDeviceStatusInfoResponseBody {
	return s.Body
}

func (s *GetDeviceStatusInfoResponse) SetHeaders(v map[string]*string) *GetDeviceStatusInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceStatusInfoResponse) SetStatusCode(v int32) *GetDeviceStatusInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceStatusInfoResponse) SetBody(v *GetDeviceStatusInfoResponseBody) *GetDeviceStatusInfoResponse {
	s.Body = v
	return s
}

func (s *GetDeviceStatusInfoResponse) Validate() error {
	return dara.Validate(s)
}
