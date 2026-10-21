// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCopilotConversationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCopilotConversationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCopilotConversationResponse
	GetStatusCode() *int32
	SetBody(v *GetCopilotConversationResponseBody) *GetCopilotConversationResponse
	GetBody() *GetCopilotConversationResponseBody
}

type GetCopilotConversationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCopilotConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCopilotConversationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCopilotConversationResponse) GoString() string {
	return s.String()
}

func (s *GetCopilotConversationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCopilotConversationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCopilotConversationResponse) GetBody() *GetCopilotConversationResponseBody {
	return s.Body
}

func (s *GetCopilotConversationResponse) SetHeaders(v map[string]*string) *GetCopilotConversationResponse {
	s.Headers = v
	return s
}

func (s *GetCopilotConversationResponse) SetStatusCode(v int32) *GetCopilotConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCopilotConversationResponse) SetBody(v *GetCopilotConversationResponseBody) *GetCopilotConversationResponse {
	s.Body = v
	return s
}

func (s *GetCopilotConversationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
