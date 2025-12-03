// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendByDeviceFileIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendByDeviceFileIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendByDeviceFileIdResponse
	GetStatusCode() *int32
	SetBody(v *SendByDeviceFileIdResponseBody) *SendByDeviceFileIdResponse
	GetBody() *SendByDeviceFileIdResponseBody
}

type SendByDeviceFileIdResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendByDeviceFileIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendByDeviceFileIdResponse) String() string {
	return dara.Prettify(s)
}

func (s SendByDeviceFileIdResponse) GoString() string {
	return s.String()
}

func (s *SendByDeviceFileIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendByDeviceFileIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendByDeviceFileIdResponse) GetBody() *SendByDeviceFileIdResponseBody {
	return s.Body
}

func (s *SendByDeviceFileIdResponse) SetHeaders(v map[string]*string) *SendByDeviceFileIdResponse {
	s.Headers = v
	return s
}

func (s *SendByDeviceFileIdResponse) SetStatusCode(v int32) *SendByDeviceFileIdResponse {
	s.StatusCode = &v
	return s
}

func (s *SendByDeviceFileIdResponse) SetBody(v *SendByDeviceFileIdResponseBody) *SendByDeviceFileIdResponse {
	s.Body = v
	return s
}

func (s *SendByDeviceFileIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
