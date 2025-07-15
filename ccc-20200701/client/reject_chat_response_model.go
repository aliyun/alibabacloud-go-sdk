// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectChatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RejectChatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RejectChatResponse
	GetStatusCode() *int32
	SetBody(v *RejectChatResponseBody) *RejectChatResponse
	GetBody() *RejectChatResponseBody
}

type RejectChatResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RejectChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RejectChatResponse) String() string {
	return dara.Prettify(s)
}

func (s RejectChatResponse) GoString() string {
	return s.String()
}

func (s *RejectChatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RejectChatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RejectChatResponse) GetBody() *RejectChatResponseBody {
	return s.Body
}

func (s *RejectChatResponse) SetHeaders(v map[string]*string) *RejectChatResponse {
	s.Headers = v
	return s
}

func (s *RejectChatResponse) SetStatusCode(v int32) *RejectChatResponse {
	s.StatusCode = &v
	return s
}

func (s *RejectChatResponse) SetBody(v *RejectChatResponseBody) *RejectChatResponse {
	s.Body = v
	return s
}

func (s *RejectChatResponse) Validate() error {
	return dara.Validate(s)
}
