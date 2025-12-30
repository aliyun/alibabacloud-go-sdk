// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPatternQueryEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatBIPatternQueryEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatBIPatternQueryEntriesResponse
	GetStatusCode() *int32
	SetBody(v *ChatBIPatternQueryEntriesResponseBody) *ChatBIPatternQueryEntriesResponse
	GetBody() *ChatBIPatternQueryEntriesResponseBody
}

type ChatBIPatternQueryEntriesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatBIPatternQueryEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatBIPatternQueryEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPatternQueryEntriesResponse) GoString() string {
	return s.String()
}

func (s *ChatBIPatternQueryEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatBIPatternQueryEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatBIPatternQueryEntriesResponse) GetBody() *ChatBIPatternQueryEntriesResponseBody {
	return s.Body
}

func (s *ChatBIPatternQueryEntriesResponse) SetHeaders(v map[string]*string) *ChatBIPatternQueryEntriesResponse {
	s.Headers = v
	return s
}

func (s *ChatBIPatternQueryEntriesResponse) SetStatusCode(v int32) *ChatBIPatternQueryEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatBIPatternQueryEntriesResponse) SetBody(v *ChatBIPatternQueryEntriesResponseBody) *ChatBIPatternQueryEntriesResponse {
	s.Body = v
	return s
}

func (s *ChatBIPatternQueryEntriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
