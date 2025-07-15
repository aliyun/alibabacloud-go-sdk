// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageToGroupUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendMessageToGroupUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendMessageToGroupUsersResponse
	GetStatusCode() *int32
	SetBody(v *SendMessageToGroupUsersResponseBody) *SendMessageToGroupUsersResponse
	GetBody() *SendMessageToGroupUsersResponseBody
}

type SendMessageToGroupUsersResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendMessageToGroupUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendMessageToGroupUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s SendMessageToGroupUsersResponse) GoString() string {
	return s.String()
}

func (s *SendMessageToGroupUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendMessageToGroupUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendMessageToGroupUsersResponse) GetBody() *SendMessageToGroupUsersResponseBody {
	return s.Body
}

func (s *SendMessageToGroupUsersResponse) SetHeaders(v map[string]*string) *SendMessageToGroupUsersResponse {
	s.Headers = v
	return s
}

func (s *SendMessageToGroupUsersResponse) SetStatusCode(v int32) *SendMessageToGroupUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *SendMessageToGroupUsersResponse) SetBody(v *SendMessageToGroupUsersResponseBody) *SendMessageToGroupUsersResponse {
	s.Body = v
	return s
}

func (s *SendMessageToGroupUsersResponse) Validate() error {
	return dara.Validate(s)
}
