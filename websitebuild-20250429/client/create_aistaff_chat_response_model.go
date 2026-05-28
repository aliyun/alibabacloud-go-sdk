// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAIStaffChatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAIStaffChatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAIStaffChatResponse
	GetStatusCode() *int32
	SetBody(v *CreateAIStaffChatResponseBody) *CreateAIStaffChatResponse
	GetBody() *CreateAIStaffChatResponseBody
}

type CreateAIStaffChatResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAIStaffChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAIStaffChatResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAIStaffChatResponse) GoString() string {
	return s.String()
}

func (s *CreateAIStaffChatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAIStaffChatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAIStaffChatResponse) GetBody() *CreateAIStaffChatResponseBody {
	return s.Body
}

func (s *CreateAIStaffChatResponse) SetHeaders(v map[string]*string) *CreateAIStaffChatResponse {
	s.Headers = v
	return s
}

func (s *CreateAIStaffChatResponse) SetStatusCode(v int32) *CreateAIStaffChatResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAIStaffChatResponse) SetBody(v *CreateAIStaffChatResponseBody) *CreateAIStaffChatResponse {
	s.Body = v
	return s
}

func (s *CreateAIStaffChatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
