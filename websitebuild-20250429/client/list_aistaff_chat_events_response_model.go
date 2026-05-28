// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIStaffChatEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAIStaffChatEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAIStaffChatEventsResponse
	GetStatusCode() *int32
	SetBody(v *ListAIStaffChatEventsResponseBody) *ListAIStaffChatEventsResponse
	GetBody() *ListAIStaffChatEventsResponseBody
}

type ListAIStaffChatEventsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAIStaffChatEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAIStaffChatEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAIStaffChatEventsResponse) GoString() string {
	return s.String()
}

func (s *ListAIStaffChatEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAIStaffChatEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAIStaffChatEventsResponse) GetBody() *ListAIStaffChatEventsResponseBody {
	return s.Body
}

func (s *ListAIStaffChatEventsResponse) SetHeaders(v map[string]*string) *ListAIStaffChatEventsResponse {
	s.Headers = v
	return s
}

func (s *ListAIStaffChatEventsResponse) SetStatusCode(v int32) *ListAIStaffChatEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAIStaffChatEventsResponse) SetBody(v *ListAIStaffChatEventsResponseBody) *ListAIStaffChatEventsResponse {
	s.Body = v
	return s
}

func (s *ListAIStaffChatEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
