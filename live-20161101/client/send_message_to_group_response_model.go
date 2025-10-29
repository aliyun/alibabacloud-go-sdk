// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageToGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendMessageToGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendMessageToGroupResponse
	GetStatusCode() *int32
	SetBody(v *SendMessageToGroupResponseBody) *SendMessageToGroupResponse
	GetBody() *SendMessageToGroupResponseBody
}

type SendMessageToGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendMessageToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendMessageToGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s SendMessageToGroupResponse) GoString() string {
	return s.String()
}

func (s *SendMessageToGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendMessageToGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendMessageToGroupResponse) GetBody() *SendMessageToGroupResponseBody {
	return s.Body
}

func (s *SendMessageToGroupResponse) SetHeaders(v map[string]*string) *SendMessageToGroupResponse {
	s.Headers = v
	return s
}

func (s *SendMessageToGroupResponse) SetStatusCode(v int32) *SendMessageToGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *SendMessageToGroupResponse) SetBody(v *SendMessageToGroupResponseBody) *SendMessageToGroupResponse {
	s.Body = v
	return s
}

func (s *SendMessageToGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
