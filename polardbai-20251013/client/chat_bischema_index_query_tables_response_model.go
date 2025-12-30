// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBISchemaIndexQueryTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatBISchemaIndexQueryTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatBISchemaIndexQueryTablesResponse
	GetStatusCode() *int32
	SetBody(v *ChatBISchemaIndexQueryTablesResponseBody) *ChatBISchemaIndexQueryTablesResponse
	GetBody() *ChatBISchemaIndexQueryTablesResponseBody
}

type ChatBISchemaIndexQueryTablesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatBISchemaIndexQueryTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatBISchemaIndexQueryTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatBISchemaIndexQueryTablesResponse) GoString() string {
	return s.String()
}

func (s *ChatBISchemaIndexQueryTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatBISchemaIndexQueryTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatBISchemaIndexQueryTablesResponse) GetBody() *ChatBISchemaIndexQueryTablesResponseBody {
	return s.Body
}

func (s *ChatBISchemaIndexQueryTablesResponse) SetHeaders(v map[string]*string) *ChatBISchemaIndexQueryTablesResponse {
	s.Headers = v
	return s
}

func (s *ChatBISchemaIndexQueryTablesResponse) SetStatusCode(v int32) *ChatBISchemaIndexQueryTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatBISchemaIndexQueryTablesResponse) SetBody(v *ChatBISchemaIndexQueryTablesResponseBody) *ChatBISchemaIndexQueryTablesResponse {
	s.Body = v
	return s
}

func (s *ChatBISchemaIndexQueryTablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
