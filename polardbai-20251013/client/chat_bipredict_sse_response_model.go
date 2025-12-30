// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPredictSseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatBIPredictSseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatBIPredictSseResponse
	GetStatusCode() *int32
	SetBody(v *ChatBIPredictSseResponseBody) *ChatBIPredictSseResponse
	GetBody() *ChatBIPredictSseResponseBody
}

type ChatBIPredictSseResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatBIPredictSseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatBIPredictSseResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPredictSseResponse) GoString() string {
	return s.String()
}

func (s *ChatBIPredictSseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatBIPredictSseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatBIPredictSseResponse) GetBody() *ChatBIPredictSseResponseBody {
	return s.Body
}

func (s *ChatBIPredictSseResponse) SetHeaders(v map[string]*string) *ChatBIPredictSseResponse {
	s.Headers = v
	return s
}

func (s *ChatBIPredictSseResponse) SetStatusCode(v int32) *ChatBIPredictSseResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatBIPredictSseResponse) SetBody(v *ChatBIPredictSseResponseBody) *ChatBIPredictSseResponse {
	s.Body = v
	return s
}

func (s *ChatBIPredictSseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
