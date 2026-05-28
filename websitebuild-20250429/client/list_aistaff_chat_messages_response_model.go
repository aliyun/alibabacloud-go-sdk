// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIStaffChatMessagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAIStaffChatMessagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAIStaffChatMessagesResponse
	GetStatusCode() *int32
	SetBody(v *ListAIStaffChatMessagesResponseBody) *ListAIStaffChatMessagesResponse
	GetBody() *ListAIStaffChatMessagesResponseBody
}

type ListAIStaffChatMessagesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAIStaffChatMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAIStaffChatMessagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAIStaffChatMessagesResponse) GoString() string {
	return s.String()
}

func (s *ListAIStaffChatMessagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAIStaffChatMessagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAIStaffChatMessagesResponse) GetBody() *ListAIStaffChatMessagesResponseBody {
	return s.Body
}

func (s *ListAIStaffChatMessagesResponse) SetHeaders(v map[string]*string) *ListAIStaffChatMessagesResponse {
	s.Headers = v
	return s
}

func (s *ListAIStaffChatMessagesResponse) SetStatusCode(v int32) *ListAIStaffChatMessagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAIStaffChatMessagesResponse) SetBody(v *ListAIStaffChatMessagesResponseBody) *ListAIStaffChatMessagesResponse {
	s.Body = v
	return s
}

func (s *ListAIStaffChatMessagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
