// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNewNetDeviceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *NewNetDeviceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *NewNetDeviceInfoResponse
	GetStatusCode() *int32
	SetBody(v *NewNetDeviceInfoResponseBody) *NewNetDeviceInfoResponse
	GetBody() *NewNetDeviceInfoResponseBody
}

type NewNetDeviceInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *NewNetDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s NewNetDeviceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s NewNetDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *NewNetDeviceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *NewNetDeviceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *NewNetDeviceInfoResponse) GetBody() *NewNetDeviceInfoResponseBody {
	return s.Body
}

func (s *NewNetDeviceInfoResponse) SetHeaders(v map[string]*string) *NewNetDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *NewNetDeviceInfoResponse) SetStatusCode(v int32) *NewNetDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *NewNetDeviceInfoResponse) SetBody(v *NewNetDeviceInfoResponseBody) *NewNetDeviceInfoResponse {
	s.Body = v
	return s
}

func (s *NewNetDeviceInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
