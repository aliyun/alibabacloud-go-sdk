// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConversationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConversationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConversationsResponse
	GetStatusCode() *int32
	SetBody(v *ListConversationsResponseBody) *ListConversationsResponse
	GetBody() *ListConversationsResponseBody
}

type ListConversationsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConversationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConversationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConversationsResponse) GoString() string {
	return s.String()
}

func (s *ListConversationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConversationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConversationsResponse) GetBody() *ListConversationsResponseBody {
	return s.Body
}

func (s *ListConversationsResponse) SetHeaders(v map[string]*string) *ListConversationsResponse {
	s.Headers = v
	return s
}

func (s *ListConversationsResponse) SetStatusCode(v int32) *ListConversationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConversationsResponse) SetBody(v *ListConversationsResponseBody) *ListConversationsResponse {
	s.Body = v
	return s
}

func (s *ListConversationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
