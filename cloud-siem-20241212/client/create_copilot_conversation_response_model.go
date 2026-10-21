// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCopilotConversationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCopilotConversationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCopilotConversationResponse
	GetStatusCode() *int32
	SetBody(v *CreateCopilotConversationResponseBody) *CreateCopilotConversationResponse
	GetBody() *CreateCopilotConversationResponseBody
}

type CreateCopilotConversationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCopilotConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCopilotConversationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCopilotConversationResponse) GoString() string {
	return s.String()
}

func (s *CreateCopilotConversationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCopilotConversationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCopilotConversationResponse) GetBody() *CreateCopilotConversationResponseBody {
	return s.Body
}

func (s *CreateCopilotConversationResponse) SetHeaders(v map[string]*string) *CreateCopilotConversationResponse {
	s.Headers = v
	return s
}

func (s *CreateCopilotConversationResponse) SetStatusCode(v int32) *CreateCopilotConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCopilotConversationResponse) SetBody(v *CreateCopilotConversationResponseBody) *CreateCopilotConversationResponse {
	s.Body = v
	return s
}

func (s *CreateCopilotConversationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
