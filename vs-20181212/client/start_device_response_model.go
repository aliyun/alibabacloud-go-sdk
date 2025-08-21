// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartDeviceResponse
	GetStatusCode() *int32
	SetBody(v *StartDeviceResponseBody) *StartDeviceResponse
	GetBody() *StartDeviceResponseBody
}

type StartDeviceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s StartDeviceResponse) GoString() string {
	return s.String()
}

func (s *StartDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartDeviceResponse) GetBody() *StartDeviceResponseBody {
	return s.Body
}

func (s *StartDeviceResponse) SetHeaders(v map[string]*string) *StartDeviceResponse {
	s.Headers = v
	return s
}

func (s *StartDeviceResponse) SetStatusCode(v int32) *StartDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDeviceResponse) SetBody(v *StartDeviceResponseBody) *StartDeviceResponse {
	s.Body = v
	return s
}

func (s *StartDeviceResponse) Validate() error {
	return dara.Validate(s)
}
