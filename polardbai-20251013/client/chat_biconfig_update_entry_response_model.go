// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIConfigUpdateEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatBIConfigUpdateEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatBIConfigUpdateEntryResponse
	GetStatusCode() *int32
	SetBody(v *ChatBIConfigUpdateEntryResponseBody) *ChatBIConfigUpdateEntryResponse
	GetBody() *ChatBIConfigUpdateEntryResponseBody
}

type ChatBIConfigUpdateEntryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatBIConfigUpdateEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatBIConfigUpdateEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatBIConfigUpdateEntryResponse) GoString() string {
	return s.String()
}

func (s *ChatBIConfigUpdateEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatBIConfigUpdateEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatBIConfigUpdateEntryResponse) GetBody() *ChatBIConfigUpdateEntryResponseBody {
	return s.Body
}

func (s *ChatBIConfigUpdateEntryResponse) SetHeaders(v map[string]*string) *ChatBIConfigUpdateEntryResponse {
	s.Headers = v
	return s
}

func (s *ChatBIConfigUpdateEntryResponse) SetStatusCode(v int32) *ChatBIConfigUpdateEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatBIConfigUpdateEntryResponse) SetBody(v *ChatBIConfigUpdateEntryResponseBody) *ChatBIConfigUpdateEntryResponse {
	s.Body = v
	return s
}

func (s *ChatBIConfigUpdateEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
