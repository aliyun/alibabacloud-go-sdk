// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIUpdateTableValidationColumnsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatBIUpdateTableValidationColumnsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatBIUpdateTableValidationColumnsResponse
	GetStatusCode() *int32
	SetBody(v *ChatBIUpdateTableValidationColumnsResponseBody) *ChatBIUpdateTableValidationColumnsResponse
	GetBody() *ChatBIUpdateTableValidationColumnsResponseBody
}

type ChatBIUpdateTableValidationColumnsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatBIUpdateTableValidationColumnsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatBIUpdateTableValidationColumnsResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatBIUpdateTableValidationColumnsResponse) GoString() string {
	return s.String()
}

func (s *ChatBIUpdateTableValidationColumnsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatBIUpdateTableValidationColumnsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatBIUpdateTableValidationColumnsResponse) GetBody() *ChatBIUpdateTableValidationColumnsResponseBody {
	return s.Body
}

func (s *ChatBIUpdateTableValidationColumnsResponse) SetHeaders(v map[string]*string) *ChatBIUpdateTableValidationColumnsResponse {
	s.Headers = v
	return s
}

func (s *ChatBIUpdateTableValidationColumnsResponse) SetStatusCode(v int32) *ChatBIUpdateTableValidationColumnsResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatBIUpdateTableValidationColumnsResponse) SetBody(v *ChatBIUpdateTableValidationColumnsResponseBody) *ChatBIUpdateTableValidationColumnsResponse {
	s.Body = v
	return s
}

func (s *ChatBIUpdateTableValidationColumnsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
