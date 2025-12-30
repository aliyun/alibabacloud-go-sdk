// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPatternUpdateEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatBIPatternUpdateEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatBIPatternUpdateEntryResponse
	GetStatusCode() *int32
	SetBody(v *ChatBIPatternUpdateEntryResponseBody) *ChatBIPatternUpdateEntryResponse
	GetBody() *ChatBIPatternUpdateEntryResponseBody
}

type ChatBIPatternUpdateEntryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatBIPatternUpdateEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatBIPatternUpdateEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPatternUpdateEntryResponse) GoString() string {
	return s.String()
}

func (s *ChatBIPatternUpdateEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatBIPatternUpdateEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatBIPatternUpdateEntryResponse) GetBody() *ChatBIPatternUpdateEntryResponseBody {
	return s.Body
}

func (s *ChatBIPatternUpdateEntryResponse) SetHeaders(v map[string]*string) *ChatBIPatternUpdateEntryResponse {
	s.Headers = v
	return s
}

func (s *ChatBIPatternUpdateEntryResponse) SetStatusCode(v int32) *ChatBIPatternUpdateEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatBIPatternUpdateEntryResponse) SetBody(v *ChatBIPatternUpdateEntryResponseBody) *ChatBIPatternUpdateEntryResponse {
	s.Body = v
	return s
}

func (s *ChatBIPatternUpdateEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
