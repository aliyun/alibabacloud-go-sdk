// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupDingtalkChatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGroupDingtalkChatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGroupDingtalkChatResponse
	GetStatusCode() *int32
	SetBody(v *CreateGroupDingtalkChatResponseBody) *CreateGroupDingtalkChatResponse
	GetBody() *CreateGroupDingtalkChatResponseBody
}

type CreateGroupDingtalkChatResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupDingtalkChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGroupDingtalkChatResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupDingtalkChatResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupDingtalkChatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGroupDingtalkChatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGroupDingtalkChatResponse) GetBody() *CreateGroupDingtalkChatResponseBody {
	return s.Body
}

func (s *CreateGroupDingtalkChatResponse) SetHeaders(v map[string]*string) *CreateGroupDingtalkChatResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupDingtalkChatResponse) SetStatusCode(v int32) *CreateGroupDingtalkChatResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupDingtalkChatResponse) SetBody(v *CreateGroupDingtalkChatResponseBody) *CreateGroupDingtalkChatResponse {
	s.Body = v
	return s
}

func (s *CreateGroupDingtalkChatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
