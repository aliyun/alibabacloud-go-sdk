// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIConfigCreateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatBIConfigCreateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatBIConfigCreateResponse
	GetStatusCode() *int32
	SetBody(v *ChatBIConfigCreateResponseBody) *ChatBIConfigCreateResponse
	GetBody() *ChatBIConfigCreateResponseBody
}

type ChatBIConfigCreateResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatBIConfigCreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatBIConfigCreateResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatBIConfigCreateResponse) GoString() string {
	return s.String()
}

func (s *ChatBIConfigCreateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatBIConfigCreateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatBIConfigCreateResponse) GetBody() *ChatBIConfigCreateResponseBody {
	return s.Body
}

func (s *ChatBIConfigCreateResponse) SetHeaders(v map[string]*string) *ChatBIConfigCreateResponse {
	s.Headers = v
	return s
}

func (s *ChatBIConfigCreateResponse) SetStatusCode(v int32) *ChatBIConfigCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatBIConfigCreateResponse) SetBody(v *ChatBIConfigCreateResponseBody) *ChatBIConfigCreateResponse {
	s.Body = v
	return s
}

func (s *ChatBIConfigCreateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
