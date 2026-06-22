// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKopilotListConversationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *KopilotListConversationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *KopilotListConversationsResponse
	GetStatusCode() *int32
	SetBody(v *KopilotListConversationsResponseBody) *KopilotListConversationsResponse
	GetBody() *KopilotListConversationsResponseBody
}

type KopilotListConversationsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KopilotListConversationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s KopilotListConversationsResponse) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationsResponse) GoString() string {
	return s.String()
}

func (s *KopilotListConversationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *KopilotListConversationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *KopilotListConversationsResponse) GetBody() *KopilotListConversationsResponseBody {
	return s.Body
}

func (s *KopilotListConversationsResponse) SetHeaders(v map[string]*string) *KopilotListConversationsResponse {
	s.Headers = v
	return s
}

func (s *KopilotListConversationsResponse) SetStatusCode(v int32) *KopilotListConversationsResponse {
	s.StatusCode = &v
	return s
}

func (s *KopilotListConversationsResponse) SetBody(v *KopilotListConversationsResponseBody) *KopilotListConversationsResponse {
	s.Body = v
	return s
}

func (s *KopilotListConversationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
