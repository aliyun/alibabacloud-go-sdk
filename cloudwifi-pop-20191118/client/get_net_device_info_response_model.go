// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetDeviceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNetDeviceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNetDeviceInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetNetDeviceInfoResponseBody) *GetNetDeviceInfoResponse
	GetBody() *GetNetDeviceInfoResponseBody
}

type GetNetDeviceInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNetDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNetDeviceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNetDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *GetNetDeviceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNetDeviceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNetDeviceInfoResponse) GetBody() *GetNetDeviceInfoResponseBody {
	return s.Body
}

func (s *GetNetDeviceInfoResponse) SetHeaders(v map[string]*string) *GetNetDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *GetNetDeviceInfoResponse) SetStatusCode(v int32) *GetNetDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNetDeviceInfoResponse) SetBody(v *GetNetDeviceInfoResponseBody) *GetNetDeviceInfoResponse {
	s.Body = v
	return s
}

func (s *GetNetDeviceInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
