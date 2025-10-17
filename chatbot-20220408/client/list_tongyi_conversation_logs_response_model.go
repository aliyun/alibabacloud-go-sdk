// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTongyiConversationLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTongyiConversationLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTongyiConversationLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListTongyiConversationLogsResponseBody) *ListTongyiConversationLogsResponse
	GetBody() *ListTongyiConversationLogsResponseBody
}

type ListTongyiConversationLogsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTongyiConversationLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTongyiConversationLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTongyiConversationLogsResponse) GoString() string {
	return s.String()
}

func (s *ListTongyiConversationLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTongyiConversationLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTongyiConversationLogsResponse) GetBody() *ListTongyiConversationLogsResponseBody {
	return s.Body
}

func (s *ListTongyiConversationLogsResponse) SetHeaders(v map[string]*string) *ListTongyiConversationLogsResponse {
	s.Headers = v
	return s
}

func (s *ListTongyiConversationLogsResponse) SetStatusCode(v int32) *ListTongyiConversationLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTongyiConversationLogsResponse) SetBody(v *ListTongyiConversationLogsResponseBody) *ListTongyiConversationLogsResponse {
	s.Body = v
	return s
}

func (s *ListTongyiConversationLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
