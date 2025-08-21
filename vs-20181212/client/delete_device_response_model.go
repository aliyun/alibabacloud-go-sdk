// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDeviceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDeviceResponseBody) *DeleteDeviceResponse
	GetBody() *DeleteDeviceResponseBody
}

type DeleteDeviceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDeviceResponse) GetBody() *DeleteDeviceResponseBody {
	return s.Body
}

func (s *DeleteDeviceResponse) SetHeaders(v map[string]*string) *DeleteDeviceResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeviceResponse) SetStatusCode(v int32) *DeleteDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeviceResponse) SetBody(v *DeleteDeviceResponseBody) *DeleteDeviceResponse {
	s.Body = v
	return s
}

func (s *DeleteDeviceResponse) Validate() error {
	return dara.Validate(s)
}
