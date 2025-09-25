// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSseChatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SseChatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SseChatResponse
	GetStatusCode() *int32
	SetBody(v *SseChatResponseBody) *SseChatResponse
	GetBody() *SseChatResponseBody
}

type SseChatResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SseChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SseChatResponse) String() string {
	return dara.Prettify(s)
}

func (s SseChatResponse) GoString() string {
	return s.String()
}

func (s *SseChatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SseChatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SseChatResponse) GetBody() *SseChatResponseBody {
	return s.Body
}

func (s *SseChatResponse) SetHeaders(v map[string]*string) *SseChatResponse {
	s.Headers = v
	return s
}

func (s *SseChatResponse) SetStatusCode(v int32) *SseChatResponse {
	s.StatusCode = &v
	return s
}

func (s *SseChatResponse) SetBody(v *SseChatResponseBody) *SseChatResponse {
	s.Body = v
	return s
}

func (s *SseChatResponse) Validate() error {
	return dara.Validate(s)
}
