// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDeviceResponse
	GetStatusCode() *int32
	SetBody(v *CreateDeviceResponseBody) *CreateDeviceResponse
	GetBody() *CreateDeviceResponseBody
}

type CreateDeviceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDeviceResponse) GoString() string {
	return s.String()
}

func (s *CreateDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDeviceResponse) GetBody() *CreateDeviceResponseBody {
	return s.Body
}

func (s *CreateDeviceResponse) SetHeaders(v map[string]*string) *CreateDeviceResponse {
	s.Headers = v
	return s
}

func (s *CreateDeviceResponse) SetStatusCode(v int32) *CreateDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeviceResponse) SetBody(v *CreateDeviceResponseBody) *CreateDeviceResponse {
	s.Body = v
	return s
}

func (s *CreateDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
