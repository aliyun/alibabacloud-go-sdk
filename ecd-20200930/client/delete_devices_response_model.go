// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDevicesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDevicesResponseBody) *DeleteDevicesResponse
	GetBody() *DeleteDevicesResponseBody
}

type DeleteDevicesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDevicesResponse) GoString() string {
	return s.String()
}

func (s *DeleteDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDevicesResponse) GetBody() *DeleteDevicesResponseBody {
	return s.Body
}

func (s *DeleteDevicesResponse) SetHeaders(v map[string]*string) *DeleteDevicesResponse {
	s.Headers = v
	return s
}

func (s *DeleteDevicesResponse) SetStatusCode(v int32) *DeleteDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDevicesResponse) SetBody(v *DeleteDevicesResponseBody) *DeleteDevicesResponse {
	s.Body = v
	return s
}

func (s *DeleteDevicesResponse) Validate() error {
	return dara.Validate(s)
}
