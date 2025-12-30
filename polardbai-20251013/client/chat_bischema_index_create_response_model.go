// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBISchemaIndexCreateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatBISchemaIndexCreateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatBISchemaIndexCreateResponse
	GetStatusCode() *int32
	SetBody(v *ChatBISchemaIndexCreateResponseBody) *ChatBISchemaIndexCreateResponse
	GetBody() *ChatBISchemaIndexCreateResponseBody
}

type ChatBISchemaIndexCreateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatBISchemaIndexCreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatBISchemaIndexCreateResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatBISchemaIndexCreateResponse) GoString() string {
	return s.String()
}

func (s *ChatBISchemaIndexCreateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatBISchemaIndexCreateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatBISchemaIndexCreateResponse) GetBody() *ChatBISchemaIndexCreateResponseBody {
	return s.Body
}

func (s *ChatBISchemaIndexCreateResponse) SetHeaders(v map[string]*string) *ChatBISchemaIndexCreateResponse {
	s.Headers = v
	return s
}

func (s *ChatBISchemaIndexCreateResponse) SetStatusCode(v int32) *ChatBISchemaIndexCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatBISchemaIndexCreateResponse) SetBody(v *ChatBISchemaIndexCreateResponseBody) *ChatBISchemaIndexCreateResponse {
	s.Body = v
	return s
}

func (s *ChatBISchemaIndexCreateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
