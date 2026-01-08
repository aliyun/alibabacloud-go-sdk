// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatappVerifyCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChatappVerifyCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChatappVerifyCodeResponse
	GetStatusCode() *int32
	SetBody(v *GetChatappVerifyCodeResponseBody) *GetChatappVerifyCodeResponse
	GetBody() *GetChatappVerifyCodeResponseBody
}

type GetChatappVerifyCodeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatappVerifyCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatappVerifyCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChatappVerifyCodeResponse) GoString() string {
	return s.String()
}

func (s *GetChatappVerifyCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChatappVerifyCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChatappVerifyCodeResponse) GetBody() *GetChatappVerifyCodeResponseBody {
	return s.Body
}

func (s *GetChatappVerifyCodeResponse) SetHeaders(v map[string]*string) *GetChatappVerifyCodeResponse {
	s.Headers = v
	return s
}

func (s *GetChatappVerifyCodeResponse) SetStatusCode(v int32) *GetChatappVerifyCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatappVerifyCodeResponse) SetBody(v *GetChatappVerifyCodeResponseBody) *GetChatappVerifyCodeResponse {
	s.Body = v
	return s
}

func (s *GetChatappVerifyCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
