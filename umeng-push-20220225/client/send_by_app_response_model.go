// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendByAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendByAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendByAppResponse
	GetStatusCode() *int32
	SetBody(v *SendByAppResponseBody) *SendByAppResponse
	GetBody() *SendByAppResponseBody
}

type SendByAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendByAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendByAppResponse) String() string {
	return dara.Prettify(s)
}

func (s SendByAppResponse) GoString() string {
	return s.String()
}

func (s *SendByAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendByAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendByAppResponse) GetBody() *SendByAppResponseBody {
	return s.Body
}

func (s *SendByAppResponse) SetHeaders(v map[string]*string) *SendByAppResponse {
	s.Headers = v
	return s
}

func (s *SendByAppResponse) SetStatusCode(v int32) *SendByAppResponse {
	s.StatusCode = &v
	return s
}

func (s *SendByAppResponse) SetBody(v *SendByAppResponseBody) *SendByAppResponse {
	s.Body = v
	return s
}

func (s *SendByAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
