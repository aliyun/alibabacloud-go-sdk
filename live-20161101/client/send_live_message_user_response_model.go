// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendLiveMessageUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendLiveMessageUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendLiveMessageUserResponse
	GetStatusCode() *int32
	SetBody(v *SendLiveMessageUserResponseBody) *SendLiveMessageUserResponse
	GetBody() *SendLiveMessageUserResponseBody
}

type SendLiveMessageUserResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendLiveMessageUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendLiveMessageUserResponse) String() string {
	return dara.Prettify(s)
}

func (s SendLiveMessageUserResponse) GoString() string {
	return s.String()
}

func (s *SendLiveMessageUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendLiveMessageUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendLiveMessageUserResponse) GetBody() *SendLiveMessageUserResponseBody {
	return s.Body
}

func (s *SendLiveMessageUserResponse) SetHeaders(v map[string]*string) *SendLiveMessageUserResponse {
	s.Headers = v
	return s
}

func (s *SendLiveMessageUserResponse) SetStatusCode(v int32) *SendLiveMessageUserResponse {
	s.StatusCode = &v
	return s
}

func (s *SendLiveMessageUserResponse) SetBody(v *SendLiveMessageUserResponseBody) *SendLiveMessageUserResponse {
	s.Body = v
	return s
}

func (s *SendLiveMessageUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
