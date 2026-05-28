// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAIStaffConversationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAIStaffConversationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAIStaffConversationResponse
	GetStatusCode() *int32
	SetBody(v *CreateAIStaffConversationResponseBody) *CreateAIStaffConversationResponse
	GetBody() *CreateAIStaffConversationResponseBody
}

type CreateAIStaffConversationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAIStaffConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAIStaffConversationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAIStaffConversationResponse) GoString() string {
	return s.String()
}

func (s *CreateAIStaffConversationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAIStaffConversationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAIStaffConversationResponse) GetBody() *CreateAIStaffConversationResponseBody {
	return s.Body
}

func (s *CreateAIStaffConversationResponse) SetHeaders(v map[string]*string) *CreateAIStaffConversationResponse {
	s.Headers = v
	return s
}

func (s *CreateAIStaffConversationResponse) SetStatusCode(v int32) *CreateAIStaffConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAIStaffConversationResponse) SetBody(v *CreateAIStaffConversationResponseBody) *CreateAIStaffConversationResponse {
	s.Body = v
	return s
}

func (s *CreateAIStaffConversationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
