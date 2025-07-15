// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClaimChatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClaimChatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClaimChatResponse
	GetStatusCode() *int32
	SetBody(v *ClaimChatResponseBody) *ClaimChatResponse
	GetBody() *ClaimChatResponseBody
}

type ClaimChatResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClaimChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClaimChatResponse) String() string {
	return dara.Prettify(s)
}

func (s ClaimChatResponse) GoString() string {
	return s.String()
}

func (s *ClaimChatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClaimChatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClaimChatResponse) GetBody() *ClaimChatResponseBody {
	return s.Body
}

func (s *ClaimChatResponse) SetHeaders(v map[string]*string) *ClaimChatResponse {
	s.Headers = v
	return s
}

func (s *ClaimChatResponse) SetStatusCode(v int32) *ClaimChatResponse {
	s.StatusCode = &v
	return s
}

func (s *ClaimChatResponse) SetBody(v *ClaimChatResponseBody) *ClaimChatResponse {
	s.Body = v
	return s
}

func (s *ClaimChatResponse) Validate() error {
	return dara.Validate(s)
}
