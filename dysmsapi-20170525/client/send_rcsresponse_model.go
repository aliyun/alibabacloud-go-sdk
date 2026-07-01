// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendRCSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendRCSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendRCSResponse
	GetStatusCode() *int32
	SetBody(v *SendRCSResponseBody) *SendRCSResponse
	GetBody() *SendRCSResponseBody
}

type SendRCSResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendRCSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendRCSResponse) String() string {
	return dara.Prettify(s)
}

func (s SendRCSResponse) GoString() string {
	return s.String()
}

func (s *SendRCSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendRCSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendRCSResponse) GetBody() *SendRCSResponseBody {
	return s.Body
}

func (s *SendRCSResponse) SetHeaders(v map[string]*string) *SendRCSResponse {
	s.Headers = v
	return s
}

func (s *SendRCSResponse) SetStatusCode(v int32) *SendRCSResponse {
	s.StatusCode = &v
	return s
}

func (s *SendRCSResponse) SetBody(v *SendRCSResponseBody) *SendRCSResponse {
	s.Body = v
	return s
}

func (s *SendRCSResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
