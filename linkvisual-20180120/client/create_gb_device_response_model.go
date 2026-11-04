// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGbDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGbDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGbDeviceResponse
	GetStatusCode() *int32
	SetBody(v *CreateGbDeviceResponseBody) *CreateGbDeviceResponse
	GetBody() *CreateGbDeviceResponseBody
}

type CreateGbDeviceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGbDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGbDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGbDeviceResponse) GoString() string {
	return s.String()
}

func (s *CreateGbDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGbDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGbDeviceResponse) GetBody() *CreateGbDeviceResponseBody {
	return s.Body
}

func (s *CreateGbDeviceResponse) SetHeaders(v map[string]*string) *CreateGbDeviceResponse {
	s.Headers = v
	return s
}

func (s *CreateGbDeviceResponse) SetStatusCode(v int32) *CreateGbDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGbDeviceResponse) SetBody(v *CreateGbDeviceResponseBody) *CreateGbDeviceResponse {
	s.Body = v
	return s
}

func (s *CreateGbDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
