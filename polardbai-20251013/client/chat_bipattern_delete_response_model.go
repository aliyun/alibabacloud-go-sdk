// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPatternDeleteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatBIPatternDeleteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatBIPatternDeleteResponse
	GetStatusCode() *int32
	SetBody(v *ChatBIPatternDeleteResponseBody) *ChatBIPatternDeleteResponse
	GetBody() *ChatBIPatternDeleteResponseBody
}

type ChatBIPatternDeleteResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatBIPatternDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatBIPatternDeleteResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPatternDeleteResponse) GoString() string {
	return s.String()
}

func (s *ChatBIPatternDeleteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatBIPatternDeleteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatBIPatternDeleteResponse) GetBody() *ChatBIPatternDeleteResponseBody {
	return s.Body
}

func (s *ChatBIPatternDeleteResponse) SetHeaders(v map[string]*string) *ChatBIPatternDeleteResponse {
	s.Headers = v
	return s
}

func (s *ChatBIPatternDeleteResponse) SetStatusCode(v int32) *ChatBIPatternDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatBIPatternDeleteResponse) SetBody(v *ChatBIPatternDeleteResponseBody) *ChatBIPatternDeleteResponse {
	s.Body = v
	return s
}

func (s *ChatBIPatternDeleteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
