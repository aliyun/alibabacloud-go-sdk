// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatappEmbedSignUpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatappEmbedSignUpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatappEmbedSignUpResponse
	GetStatusCode() *int32
	SetBody(v *ChatappEmbedSignUpResponseBody) *ChatappEmbedSignUpResponse
	GetBody() *ChatappEmbedSignUpResponseBody
}

type ChatappEmbedSignUpResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatappEmbedSignUpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatappEmbedSignUpResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatappEmbedSignUpResponse) GoString() string {
	return s.String()
}

func (s *ChatappEmbedSignUpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatappEmbedSignUpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatappEmbedSignUpResponse) GetBody() *ChatappEmbedSignUpResponseBody {
	return s.Body
}

func (s *ChatappEmbedSignUpResponse) SetHeaders(v map[string]*string) *ChatappEmbedSignUpResponse {
	s.Headers = v
	return s
}

func (s *ChatappEmbedSignUpResponse) SetStatusCode(v int32) *ChatappEmbedSignUpResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatappEmbedSignUpResponse) SetBody(v *ChatappEmbedSignUpResponseBody) *ChatappEmbedSignUpResponse {
	s.Body = v
	return s
}

func (s *ChatappEmbedSignUpResponse) Validate() error {
	return dara.Validate(s)
}
