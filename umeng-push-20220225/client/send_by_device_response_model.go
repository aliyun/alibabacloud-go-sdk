// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendByDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendByDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendByDeviceResponse
	GetStatusCode() *int32
	SetBody(v *SendByDeviceResponseBody) *SendByDeviceResponse
	GetBody() *SendByDeviceResponseBody
}

type SendByDeviceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendByDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendByDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s SendByDeviceResponse) GoString() string {
	return s.String()
}

func (s *SendByDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendByDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendByDeviceResponse) GetBody() *SendByDeviceResponseBody {
	return s.Body
}

func (s *SendByDeviceResponse) SetHeaders(v map[string]*string) *SendByDeviceResponse {
	s.Headers = v
	return s
}

func (s *SendByDeviceResponse) SetStatusCode(v int32) *SendByDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *SendByDeviceResponse) SetBody(v *SendByDeviceResponseBody) *SendByDeviceResponse {
	s.Body = v
	return s
}

func (s *SendByDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
