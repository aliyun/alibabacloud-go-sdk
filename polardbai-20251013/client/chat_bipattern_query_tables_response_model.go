// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPatternQueryTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatBIPatternQueryTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatBIPatternQueryTablesResponse
	GetStatusCode() *int32
	SetBody(v *ChatBIPatternQueryTablesResponseBody) *ChatBIPatternQueryTablesResponse
	GetBody() *ChatBIPatternQueryTablesResponseBody
}

type ChatBIPatternQueryTablesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatBIPatternQueryTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatBIPatternQueryTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPatternQueryTablesResponse) GoString() string {
	return s.String()
}

func (s *ChatBIPatternQueryTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatBIPatternQueryTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatBIPatternQueryTablesResponse) GetBody() *ChatBIPatternQueryTablesResponseBody {
	return s.Body
}

func (s *ChatBIPatternQueryTablesResponse) SetHeaders(v map[string]*string) *ChatBIPatternQueryTablesResponse {
	s.Headers = v
	return s
}

func (s *ChatBIPatternQueryTablesResponse) SetStatusCode(v int32) *ChatBIPatternQueryTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatBIPatternQueryTablesResponse) SetBody(v *ChatBIPatternQueryTablesResponseBody) *ChatBIPatternQueryTablesResponse {
	s.Body = v
	return s
}

func (s *ChatBIPatternQueryTablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
