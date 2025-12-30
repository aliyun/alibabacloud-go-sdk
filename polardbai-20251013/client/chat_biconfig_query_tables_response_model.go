// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIConfigQueryTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatBIConfigQueryTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatBIConfigQueryTablesResponse
	GetStatusCode() *int32
	SetBody(v *ChatBIConfigQueryTablesResponseBody) *ChatBIConfigQueryTablesResponse
	GetBody() *ChatBIConfigQueryTablesResponseBody
}

type ChatBIConfigQueryTablesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatBIConfigQueryTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatBIConfigQueryTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatBIConfigQueryTablesResponse) GoString() string {
	return s.String()
}

func (s *ChatBIConfigQueryTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatBIConfigQueryTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatBIConfigQueryTablesResponse) GetBody() *ChatBIConfigQueryTablesResponseBody {
	return s.Body
}

func (s *ChatBIConfigQueryTablesResponse) SetHeaders(v map[string]*string) *ChatBIConfigQueryTablesResponse {
	s.Headers = v
	return s
}

func (s *ChatBIConfigQueryTablesResponse) SetStatusCode(v int32) *ChatBIConfigQueryTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatBIConfigQueryTablesResponse) SetBody(v *ChatBIConfigQueryTablesResponseBody) *ChatBIConfigQueryTablesResponse {
	s.Body = v
	return s
}

func (s *ChatBIConfigQueryTablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
