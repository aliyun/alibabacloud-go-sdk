// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendByFilterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendByFilterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendByFilterResponse
	GetStatusCode() *int32
	SetBody(v *SendByFilterResponseBody) *SendByFilterResponse
	GetBody() *SendByFilterResponseBody
}

type SendByFilterResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendByFilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendByFilterResponse) String() string {
	return dara.Prettify(s)
}

func (s SendByFilterResponse) GoString() string {
	return s.String()
}

func (s *SendByFilterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendByFilterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendByFilterResponse) GetBody() *SendByFilterResponseBody {
	return s.Body
}

func (s *SendByFilterResponse) SetHeaders(v map[string]*string) *SendByFilterResponse {
	s.Headers = v
	return s
}

func (s *SendByFilterResponse) SetStatusCode(v int32) *SendByFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *SendByFilterResponse) SetBody(v *SendByFilterResponseBody) *SendByFilterResponse {
	s.Body = v
	return s
}

func (s *SendByFilterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
