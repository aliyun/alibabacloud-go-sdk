// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceBasicInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeviceBasicInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeviceBasicInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetDeviceBasicInfoResponseBody) *GetDeviceBasicInfoResponse
	GetBody() *GetDeviceBasicInfoResponseBody
}

type GetDeviceBasicInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeviceBasicInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceBasicInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeviceBasicInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeviceBasicInfoResponse) GetBody() *GetDeviceBasicInfoResponseBody {
	return s.Body
}

func (s *GetDeviceBasicInfoResponse) SetHeaders(v map[string]*string) *GetDeviceBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceBasicInfoResponse) SetStatusCode(v int32) *GetDeviceBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceBasicInfoResponse) SetBody(v *GetDeviceBasicInfoResponseBody) *GetDeviceBasicInfoResponse {
	s.Body = v
	return s
}

func (s *GetDeviceBasicInfoResponse) Validate() error {
	return dara.Validate(s)
}
