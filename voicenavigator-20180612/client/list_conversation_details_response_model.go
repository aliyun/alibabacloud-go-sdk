// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConversationDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConversationDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConversationDetailsResponse
	GetStatusCode() *int32
	SetBody(v *ListConversationDetailsResponseBody) *ListConversationDetailsResponse
	GetBody() *ListConversationDetailsResponseBody
}

type ListConversationDetailsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConversationDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConversationDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConversationDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListConversationDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConversationDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConversationDetailsResponse) GetBody() *ListConversationDetailsResponseBody {
	return s.Body
}

func (s *ListConversationDetailsResponse) SetHeaders(v map[string]*string) *ListConversationDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListConversationDetailsResponse) SetStatusCode(v int32) *ListConversationDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConversationDetailsResponse) SetBody(v *ListConversationDetailsResponseBody) *ListConversationDetailsResponse {
	s.Body = v
	return s
}

func (s *ListConversationDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
