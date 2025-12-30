// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPatternDeleteEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatBIPatternDeleteEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatBIPatternDeleteEntryResponse
	GetStatusCode() *int32
	SetBody(v *ChatBIPatternDeleteEntryResponseBody) *ChatBIPatternDeleteEntryResponse
	GetBody() *ChatBIPatternDeleteEntryResponseBody
}

type ChatBIPatternDeleteEntryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatBIPatternDeleteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatBIPatternDeleteEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPatternDeleteEntryResponse) GoString() string {
	return s.String()
}

func (s *ChatBIPatternDeleteEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatBIPatternDeleteEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatBIPatternDeleteEntryResponse) GetBody() *ChatBIPatternDeleteEntryResponseBody {
	return s.Body
}

func (s *ChatBIPatternDeleteEntryResponse) SetHeaders(v map[string]*string) *ChatBIPatternDeleteEntryResponse {
	s.Headers = v
	return s
}

func (s *ChatBIPatternDeleteEntryResponse) SetStatusCode(v int32) *ChatBIPatternDeleteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatBIPatternDeleteEntryResponse) SetBody(v *ChatBIPatternDeleteEntryResponseBody) *ChatBIPatternDeleteEntryResponse {
	s.Body = v
	return s
}

func (s *ChatBIPatternDeleteEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
