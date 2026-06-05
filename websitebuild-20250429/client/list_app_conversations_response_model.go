// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppConversationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppConversationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppConversationsResponse
	GetStatusCode() *int32
	SetBody(v *ListAppConversationsResponseBody) *ListAppConversationsResponse
	GetBody() *ListAppConversationsResponseBody
}

type ListAppConversationsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppConversationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppConversationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppConversationsResponse) GoString() string {
	return s.String()
}

func (s *ListAppConversationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppConversationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppConversationsResponse) GetBody() *ListAppConversationsResponseBody {
	return s.Body
}

func (s *ListAppConversationsResponse) SetHeaders(v map[string]*string) *ListAppConversationsResponse {
	s.Headers = v
	return s
}

func (s *ListAppConversationsResponse) SetStatusCode(v int32) *ListAppConversationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppConversationsResponse) SetBody(v *ListAppConversationsResponseBody) *ListAppConversationsResponse {
	s.Body = v
	return s
}

func (s *ListAppConversationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
