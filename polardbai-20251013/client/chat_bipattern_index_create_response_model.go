// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPatternIndexCreateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatBIPatternIndexCreateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatBIPatternIndexCreateResponse
	GetStatusCode() *int32
	SetBody(v *ChatBIPatternIndexCreateResponseBody) *ChatBIPatternIndexCreateResponse
	GetBody() *ChatBIPatternIndexCreateResponseBody
}

type ChatBIPatternIndexCreateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatBIPatternIndexCreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatBIPatternIndexCreateResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPatternIndexCreateResponse) GoString() string {
	return s.String()
}

func (s *ChatBIPatternIndexCreateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatBIPatternIndexCreateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatBIPatternIndexCreateResponse) GetBody() *ChatBIPatternIndexCreateResponseBody {
	return s.Body
}

func (s *ChatBIPatternIndexCreateResponse) SetHeaders(v map[string]*string) *ChatBIPatternIndexCreateResponse {
	s.Headers = v
	return s
}

func (s *ChatBIPatternIndexCreateResponse) SetStatusCode(v int32) *ChatBIPatternIndexCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatBIPatternIndexCreateResponse) SetBody(v *ChatBIPatternIndexCreateResponseBody) *ChatBIPatternIndexCreateResponse {
	s.Body = v
	return s
}

func (s *ChatBIPatternIndexCreateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
