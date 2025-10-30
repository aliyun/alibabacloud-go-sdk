// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserDeviceResponse
	GetStatusCode() *int32
	SetBody(v *GetUserDeviceResponseBody) *GetUserDeviceResponse
	GetBody() *GetUserDeviceResponseBody
}

type GetUserDeviceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserDeviceResponse) GoString() string {
	return s.String()
}

func (s *GetUserDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserDeviceResponse) GetBody() *GetUserDeviceResponseBody {
	return s.Body
}

func (s *GetUserDeviceResponse) SetHeaders(v map[string]*string) *GetUserDeviceResponse {
	s.Headers = v
	return s
}

func (s *GetUserDeviceResponse) SetStatusCode(v int32) *GetUserDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserDeviceResponse) SetBody(v *GetUserDeviceResponseBody) *GetUserDeviceResponse {
	s.Body = v
	return s
}

func (s *GetUserDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
