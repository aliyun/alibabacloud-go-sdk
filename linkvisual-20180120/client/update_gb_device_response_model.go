// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGbDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGbDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGbDeviceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGbDeviceResponseBody) *UpdateGbDeviceResponse
	GetBody() *UpdateGbDeviceResponseBody
}

type UpdateGbDeviceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGbDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGbDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGbDeviceResponse) GoString() string {
	return s.String()
}

func (s *UpdateGbDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGbDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGbDeviceResponse) GetBody() *UpdateGbDeviceResponseBody {
	return s.Body
}

func (s *UpdateGbDeviceResponse) SetHeaders(v map[string]*string) *UpdateGbDeviceResponse {
	s.Headers = v
	return s
}

func (s *UpdateGbDeviceResponse) SetStatusCode(v int32) *UpdateGbDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGbDeviceResponse) SetBody(v *UpdateGbDeviceResponseBody) *UpdateGbDeviceResponse {
	s.Body = v
	return s
}

func (s *UpdateGbDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
