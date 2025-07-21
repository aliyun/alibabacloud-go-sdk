// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatappVerifyAndRegisterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatappVerifyAndRegisterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatappVerifyAndRegisterResponse
	GetStatusCode() *int32
	SetBody(v *ChatappVerifyAndRegisterResponseBody) *ChatappVerifyAndRegisterResponse
	GetBody() *ChatappVerifyAndRegisterResponseBody
}

type ChatappVerifyAndRegisterResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatappVerifyAndRegisterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatappVerifyAndRegisterResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatappVerifyAndRegisterResponse) GoString() string {
	return s.String()
}

func (s *ChatappVerifyAndRegisterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatappVerifyAndRegisterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatappVerifyAndRegisterResponse) GetBody() *ChatappVerifyAndRegisterResponseBody {
	return s.Body
}

func (s *ChatappVerifyAndRegisterResponse) SetHeaders(v map[string]*string) *ChatappVerifyAndRegisterResponse {
	s.Headers = v
	return s
}

func (s *ChatappVerifyAndRegisterResponse) SetStatusCode(v int32) *ChatappVerifyAndRegisterResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatappVerifyAndRegisterResponse) SetBody(v *ChatappVerifyAndRegisterResponseBody) *ChatappVerifyAndRegisterResponse {
	s.Body = v
	return s
}

func (s *ChatappVerifyAndRegisterResponse) Validate() error {
	return dara.Validate(s)
}
