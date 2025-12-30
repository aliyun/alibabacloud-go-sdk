// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBISchemaIndexDeleteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatBISchemaIndexDeleteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatBISchemaIndexDeleteResponse
	GetStatusCode() *int32
	SetBody(v *ChatBISchemaIndexDeleteResponseBody) *ChatBISchemaIndexDeleteResponse
	GetBody() *ChatBISchemaIndexDeleteResponseBody
}

type ChatBISchemaIndexDeleteResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatBISchemaIndexDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatBISchemaIndexDeleteResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatBISchemaIndexDeleteResponse) GoString() string {
	return s.String()
}

func (s *ChatBISchemaIndexDeleteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatBISchemaIndexDeleteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatBISchemaIndexDeleteResponse) GetBody() *ChatBISchemaIndexDeleteResponseBody {
	return s.Body
}

func (s *ChatBISchemaIndexDeleteResponse) SetHeaders(v map[string]*string) *ChatBISchemaIndexDeleteResponse {
	s.Headers = v
	return s
}

func (s *ChatBISchemaIndexDeleteResponse) SetStatusCode(v int32) *ChatBISchemaIndexDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatBISchemaIndexDeleteResponse) SetBody(v *ChatBISchemaIndexDeleteResponseBody) *ChatBISchemaIndexDeleteResponse {
	s.Body = v
	return s
}

func (s *ChatBISchemaIndexDeleteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
