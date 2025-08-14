// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateMessageChatTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateMessageChatTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateMessageChatTokenResponse
	GetStatusCode() *int32
	SetBody(v *GenerateMessageChatTokenResponseBody) *GenerateMessageChatTokenResponse
	GetBody() *GenerateMessageChatTokenResponseBody
}

type GenerateMessageChatTokenResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateMessageChatTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateMessageChatTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateMessageChatTokenResponse) GoString() string {
	return s.String()
}

func (s *GenerateMessageChatTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateMessageChatTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateMessageChatTokenResponse) GetBody() *GenerateMessageChatTokenResponseBody {
	return s.Body
}

func (s *GenerateMessageChatTokenResponse) SetHeaders(v map[string]*string) *GenerateMessageChatTokenResponse {
	s.Headers = v
	return s
}

func (s *GenerateMessageChatTokenResponse) SetStatusCode(v int32) *GenerateMessageChatTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateMessageChatTokenResponse) SetBody(v *GenerateMessageChatTokenResponseBody) *GenerateMessageChatTokenResponse {
	s.Body = v
	return s
}

func (s *GenerateMessageChatTokenResponse) Validate() error {
	return dara.Validate(s)
}
