// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIConfigDeleteEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatBIConfigDeleteEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatBIConfigDeleteEntryResponse
	GetStatusCode() *int32
	SetBody(v *ChatBIConfigDeleteEntryResponseBody) *ChatBIConfigDeleteEntryResponse
	GetBody() *ChatBIConfigDeleteEntryResponseBody
}

type ChatBIConfigDeleteEntryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatBIConfigDeleteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatBIConfigDeleteEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatBIConfigDeleteEntryResponse) GoString() string {
	return s.String()
}

func (s *ChatBIConfigDeleteEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatBIConfigDeleteEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatBIConfigDeleteEntryResponse) GetBody() *ChatBIConfigDeleteEntryResponseBody {
	return s.Body
}

func (s *ChatBIConfigDeleteEntryResponse) SetHeaders(v map[string]*string) *ChatBIConfigDeleteEntryResponse {
	s.Headers = v
	return s
}

func (s *ChatBIConfigDeleteEntryResponse) SetStatusCode(v int32) *ChatBIConfigDeleteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatBIConfigDeleteEntryResponse) SetBody(v *ChatBIConfigDeleteEntryResponseBody) *ChatBIConfigDeleteEntryResponse {
	s.Body = v
	return s
}

func (s *ChatBIConfigDeleteEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
