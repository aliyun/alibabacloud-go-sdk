// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPatternIndexQueryTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatBIPatternIndexQueryTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatBIPatternIndexQueryTablesResponse
	GetStatusCode() *int32
	SetBody(v *ChatBIPatternIndexQueryTablesResponseBody) *ChatBIPatternIndexQueryTablesResponse
	GetBody() *ChatBIPatternIndexQueryTablesResponseBody
}

type ChatBIPatternIndexQueryTablesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatBIPatternIndexQueryTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatBIPatternIndexQueryTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPatternIndexQueryTablesResponse) GoString() string {
	return s.String()
}

func (s *ChatBIPatternIndexQueryTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatBIPatternIndexQueryTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatBIPatternIndexQueryTablesResponse) GetBody() *ChatBIPatternIndexQueryTablesResponseBody {
	return s.Body
}

func (s *ChatBIPatternIndexQueryTablesResponse) SetHeaders(v map[string]*string) *ChatBIPatternIndexQueryTablesResponse {
	s.Headers = v
	return s
}

func (s *ChatBIPatternIndexQueryTablesResponse) SetStatusCode(v int32) *ChatBIPatternIndexQueryTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatBIPatternIndexQueryTablesResponse) SetBody(v *ChatBIPatternIndexQueryTablesResponseBody) *ChatBIPatternIndexQueryTablesResponse {
	s.Body = v
	return s
}

func (s *ChatBIPatternIndexQueryTablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
