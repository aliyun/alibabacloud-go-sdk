// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatMessagesTaskStopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatMessagesTaskStopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatMessagesTaskStopResponse
	GetStatusCode() *int32
	SetBody(v *ChatMessagesTaskStopResponseBody) *ChatMessagesTaskStopResponse
	GetBody() *ChatMessagesTaskStopResponseBody
}

type ChatMessagesTaskStopResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatMessagesTaskStopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatMessagesTaskStopResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatMessagesTaskStopResponse) GoString() string {
	return s.String()
}

func (s *ChatMessagesTaskStopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatMessagesTaskStopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatMessagesTaskStopResponse) GetBody() *ChatMessagesTaskStopResponseBody {
	return s.Body
}

func (s *ChatMessagesTaskStopResponse) SetHeaders(v map[string]*string) *ChatMessagesTaskStopResponse {
	s.Headers = v
	return s
}

func (s *ChatMessagesTaskStopResponse) SetStatusCode(v int32) *ChatMessagesTaskStopResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatMessagesTaskStopResponse) SetBody(v *ChatMessagesTaskStopResponseBody) *ChatMessagesTaskStopResponse {
	s.Body = v
	return s
}

func (s *ChatMessagesTaskStopResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
