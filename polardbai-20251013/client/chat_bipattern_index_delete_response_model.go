// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPatternIndexDeleteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatBIPatternIndexDeleteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatBIPatternIndexDeleteResponse
	GetStatusCode() *int32
	SetBody(v *ChatBIPatternIndexDeleteResponseBody) *ChatBIPatternIndexDeleteResponse
	GetBody() *ChatBIPatternIndexDeleteResponseBody
}

type ChatBIPatternIndexDeleteResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatBIPatternIndexDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatBIPatternIndexDeleteResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPatternIndexDeleteResponse) GoString() string {
	return s.String()
}

func (s *ChatBIPatternIndexDeleteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatBIPatternIndexDeleteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatBIPatternIndexDeleteResponse) GetBody() *ChatBIPatternIndexDeleteResponseBody {
	return s.Body
}

func (s *ChatBIPatternIndexDeleteResponse) SetHeaders(v map[string]*string) *ChatBIPatternIndexDeleteResponse {
	s.Headers = v
	return s
}

func (s *ChatBIPatternIndexDeleteResponse) SetStatusCode(v int32) *ChatBIPatternIndexDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatBIPatternIndexDeleteResponse) SetBody(v *ChatBIPatternIndexDeleteResponseBody) *ChatBIPatternIndexDeleteResponse {
	s.Body = v
	return s
}

func (s *ChatBIPatternIndexDeleteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
