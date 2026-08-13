// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopChatMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopChatMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopChatMessageResponse
	GetStatusCode() *int32
	SetBody(v *StopChatMessageResponseBody) *StopChatMessageResponse
	GetBody() *StopChatMessageResponseBody
}

type StopChatMessageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopChatMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopChatMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s StopChatMessageResponse) GoString() string {
	return s.String()
}

func (s *StopChatMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopChatMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopChatMessageResponse) GetBody() *StopChatMessageResponseBody {
	return s.Body
}

func (s *StopChatMessageResponse) SetHeaders(v map[string]*string) *StopChatMessageResponse {
	s.Headers = v
	return s
}

func (s *StopChatMessageResponse) SetStatusCode(v int32) *StopChatMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *StopChatMessageResponse) SetBody(v *StopChatMessageResponseBody) *StopChatMessageResponse {
	s.Body = v
	return s
}

func (s *StopChatMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
