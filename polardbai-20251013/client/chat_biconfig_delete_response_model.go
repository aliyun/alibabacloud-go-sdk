// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIConfigDeleteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatBIConfigDeleteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatBIConfigDeleteResponse
	GetStatusCode() *int32
	SetBody(v *ChatBIConfigDeleteResponseBody) *ChatBIConfigDeleteResponse
	GetBody() *ChatBIConfigDeleteResponseBody
}

type ChatBIConfigDeleteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatBIConfigDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatBIConfigDeleteResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatBIConfigDeleteResponse) GoString() string {
	return s.String()
}

func (s *ChatBIConfigDeleteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatBIConfigDeleteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatBIConfigDeleteResponse) GetBody() *ChatBIConfigDeleteResponseBody {
	return s.Body
}

func (s *ChatBIConfigDeleteResponse) SetHeaders(v map[string]*string) *ChatBIConfigDeleteResponse {
	s.Headers = v
	return s
}

func (s *ChatBIConfigDeleteResponse) SetStatusCode(v int32) *ChatBIConfigDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatBIConfigDeleteResponse) SetBody(v *ChatBIConfigDeleteResponseBody) *ChatBIConfigDeleteResponse {
	s.Body = v
	return s
}

func (s *ChatBIConfigDeleteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
