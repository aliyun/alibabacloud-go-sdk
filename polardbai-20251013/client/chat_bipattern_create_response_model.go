// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPatternCreateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatBIPatternCreateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatBIPatternCreateResponse
	GetStatusCode() *int32
	SetBody(v *ChatBIPatternCreateResponseBody) *ChatBIPatternCreateResponse
	GetBody() *ChatBIPatternCreateResponseBody
}

type ChatBIPatternCreateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatBIPatternCreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatBIPatternCreateResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPatternCreateResponse) GoString() string {
	return s.String()
}

func (s *ChatBIPatternCreateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatBIPatternCreateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatBIPatternCreateResponse) GetBody() *ChatBIPatternCreateResponseBody {
	return s.Body
}

func (s *ChatBIPatternCreateResponse) SetHeaders(v map[string]*string) *ChatBIPatternCreateResponse {
	s.Headers = v
	return s
}

func (s *ChatBIPatternCreateResponse) SetStatusCode(v int32) *ChatBIPatternCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatBIPatternCreateResponse) SetBody(v *ChatBIPatternCreateResponseBody) *ChatBIPatternCreateResponse {
	s.Body = v
	return s
}

func (s *ChatBIPatternCreateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
