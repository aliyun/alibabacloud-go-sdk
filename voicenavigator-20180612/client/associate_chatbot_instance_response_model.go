// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateChatbotInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateChatbotInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateChatbotInstanceResponse
	GetStatusCode() *int32
	SetBody(v *AssociateChatbotInstanceResponseBody) *AssociateChatbotInstanceResponse
	GetBody() *AssociateChatbotInstanceResponseBody
}

type AssociateChatbotInstanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateChatbotInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateChatbotInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateChatbotInstanceResponse) GoString() string {
	return s.String()
}

func (s *AssociateChatbotInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateChatbotInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateChatbotInstanceResponse) GetBody() *AssociateChatbotInstanceResponseBody {
	return s.Body
}

func (s *AssociateChatbotInstanceResponse) SetHeaders(v map[string]*string) *AssociateChatbotInstanceResponse {
	s.Headers = v
	return s
}

func (s *AssociateChatbotInstanceResponse) SetStatusCode(v int32) *AssociateChatbotInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateChatbotInstanceResponse) SetBody(v *AssociateChatbotInstanceResponseBody) *AssociateChatbotInstanceResponse {
	s.Body = v
	return s
}

func (s *AssociateChatbotInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
