// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalFeishuChatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePersonalFeishuChatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePersonalFeishuChatResponse
	GetStatusCode() *int32
	SetBody(v *CreatePersonalFeishuChatResponseBody) *CreatePersonalFeishuChatResponse
	GetBody() *CreatePersonalFeishuChatResponseBody
}

type CreatePersonalFeishuChatResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePersonalFeishuChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePersonalFeishuChatResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalFeishuChatResponse) GoString() string {
	return s.String()
}

func (s *CreatePersonalFeishuChatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePersonalFeishuChatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePersonalFeishuChatResponse) GetBody() *CreatePersonalFeishuChatResponseBody {
	return s.Body
}

func (s *CreatePersonalFeishuChatResponse) SetHeaders(v map[string]*string) *CreatePersonalFeishuChatResponse {
	s.Headers = v
	return s
}

func (s *CreatePersonalFeishuChatResponse) SetStatusCode(v int32) *CreatePersonalFeishuChatResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePersonalFeishuChatResponse) SetBody(v *CreatePersonalFeishuChatResponseBody) *CreatePersonalFeishuChatResponse {
	s.Body = v
	return s
}

func (s *CreatePersonalFeishuChatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
