// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetDeviceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNetDeviceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNetDeviceInfoResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNetDeviceInfoResponseBody) *DeleteNetDeviceInfoResponse
	GetBody() *DeleteNetDeviceInfoResponseBody
}

type DeleteNetDeviceInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNetDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNetDeviceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *DeleteNetDeviceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNetDeviceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNetDeviceInfoResponse) GetBody() *DeleteNetDeviceInfoResponseBody {
	return s.Body
}

func (s *DeleteNetDeviceInfoResponse) SetHeaders(v map[string]*string) *DeleteNetDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *DeleteNetDeviceInfoResponse) SetStatusCode(v int32) *DeleteNetDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNetDeviceInfoResponse) SetBody(v *DeleteNetDeviceInfoResponseBody) *DeleteNetDeviceInfoResponse {
	s.Body = v
	return s
}

func (s *DeleteNetDeviceInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
