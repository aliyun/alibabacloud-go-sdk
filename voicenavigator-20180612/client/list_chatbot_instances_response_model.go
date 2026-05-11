// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatbotInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListChatbotInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListChatbotInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListChatbotInstancesResponseBody) *ListChatbotInstancesResponse
	GetBody() *ListChatbotInstancesResponseBody
}

type ListChatbotInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChatbotInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChatbotInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListChatbotInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListChatbotInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListChatbotInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListChatbotInstancesResponse) GetBody() *ListChatbotInstancesResponseBody {
	return s.Body
}

func (s *ListChatbotInstancesResponse) SetHeaders(v map[string]*string) *ListChatbotInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListChatbotInstancesResponse) SetStatusCode(v int32) *ListChatbotInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChatbotInstancesResponse) SetBody(v *ListChatbotInstancesResponseBody) *ListChatbotInstancesResponse {
	s.Body = v
	return s
}

func (s *ListChatbotInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
