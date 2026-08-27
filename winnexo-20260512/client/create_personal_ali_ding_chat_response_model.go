// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalAliDingChatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePersonalAliDingChatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePersonalAliDingChatResponse
	GetStatusCode() *int32
	SetBody(v *CreatePersonalAliDingChatResponseBody) *CreatePersonalAliDingChatResponse
	GetBody() *CreatePersonalAliDingChatResponseBody
}

type CreatePersonalAliDingChatResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePersonalAliDingChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePersonalAliDingChatResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalAliDingChatResponse) GoString() string {
	return s.String()
}

func (s *CreatePersonalAliDingChatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePersonalAliDingChatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePersonalAliDingChatResponse) GetBody() *CreatePersonalAliDingChatResponseBody {
	return s.Body
}

func (s *CreatePersonalAliDingChatResponse) SetHeaders(v map[string]*string) *CreatePersonalAliDingChatResponse {
	s.Headers = v
	return s
}

func (s *CreatePersonalAliDingChatResponse) SetStatusCode(v int32) *CreatePersonalAliDingChatResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePersonalAliDingChatResponse) SetBody(v *CreatePersonalAliDingChatResponseBody) *CreatePersonalAliDingChatResponse {
	s.Body = v
	return s
}

func (s *CreatePersonalAliDingChatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
