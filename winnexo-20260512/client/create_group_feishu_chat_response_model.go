// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupFeishuChatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGroupFeishuChatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGroupFeishuChatResponse
	GetStatusCode() *int32
	SetBody(v *CreateGroupFeishuChatResponseBody) *CreateGroupFeishuChatResponse
	GetBody() *CreateGroupFeishuChatResponseBody
}

type CreateGroupFeishuChatResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupFeishuChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGroupFeishuChatResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupFeishuChatResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupFeishuChatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGroupFeishuChatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGroupFeishuChatResponse) GetBody() *CreateGroupFeishuChatResponseBody {
	return s.Body
}

func (s *CreateGroupFeishuChatResponse) SetHeaders(v map[string]*string) *CreateGroupFeishuChatResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupFeishuChatResponse) SetStatusCode(v int32) *CreateGroupFeishuChatResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupFeishuChatResponse) SetBody(v *CreateGroupFeishuChatResponseBody) *CreateGroupFeishuChatResponse {
	s.Body = v
	return s
}

func (s *CreateGroupFeishuChatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
