// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIConfigQueryEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatBIConfigQueryEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatBIConfigQueryEntriesResponse
	GetStatusCode() *int32
	SetBody(v *ChatBIConfigQueryEntriesResponseBody) *ChatBIConfigQueryEntriesResponse
	GetBody() *ChatBIConfigQueryEntriesResponseBody
}

type ChatBIConfigQueryEntriesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatBIConfigQueryEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatBIConfigQueryEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatBIConfigQueryEntriesResponse) GoString() string {
	return s.String()
}

func (s *ChatBIConfigQueryEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatBIConfigQueryEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatBIConfigQueryEntriesResponse) GetBody() *ChatBIConfigQueryEntriesResponseBody {
	return s.Body
}

func (s *ChatBIConfigQueryEntriesResponse) SetHeaders(v map[string]*string) *ChatBIConfigQueryEntriesResponse {
	s.Headers = v
	return s
}

func (s *ChatBIConfigQueryEntriesResponse) SetStatusCode(v int32) *ChatBIConfigQueryEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatBIConfigQueryEntriesResponse) SetBody(v *ChatBIConfigQueryEntriesResponseBody) *ChatBIConfigQueryEntriesResponse {
	s.Body = v
	return s
}

func (s *ChatBIConfigQueryEntriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
