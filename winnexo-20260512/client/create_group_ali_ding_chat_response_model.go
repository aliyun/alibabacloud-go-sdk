// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupAliDingChatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGroupAliDingChatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGroupAliDingChatResponse
	GetStatusCode() *int32
	SetBody(v *CreateGroupAliDingChatResponseBody) *CreateGroupAliDingChatResponse
	GetBody() *CreateGroupAliDingChatResponseBody
}

type CreateGroupAliDingChatResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupAliDingChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGroupAliDingChatResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupAliDingChatResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupAliDingChatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGroupAliDingChatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGroupAliDingChatResponse) GetBody() *CreateGroupAliDingChatResponseBody {
	return s.Body
}

func (s *CreateGroupAliDingChatResponse) SetHeaders(v map[string]*string) *CreateGroupAliDingChatResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupAliDingChatResponse) SetStatusCode(v int32) *CreateGroupAliDingChatResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupAliDingChatResponse) SetBody(v *CreateGroupAliDingChatResponseBody) *CreateGroupAliDingChatResponse {
	s.Body = v
	return s
}

func (s *CreateGroupAliDingChatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
