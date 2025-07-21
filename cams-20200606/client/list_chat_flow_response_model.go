// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListChatFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListChatFlowResponse
	GetStatusCode() *int32
	SetBody(v *ListChatFlowResponseBody) *ListChatFlowResponse
	GetBody() *ListChatFlowResponseBody
}

type ListChatFlowResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChatFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChatFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s ListChatFlowResponse) GoString() string {
	return s.String()
}

func (s *ListChatFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListChatFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListChatFlowResponse) GetBody() *ListChatFlowResponseBody {
	return s.Body
}

func (s *ListChatFlowResponse) SetHeaders(v map[string]*string) *ListChatFlowResponse {
	s.Headers = v
	return s
}

func (s *ListChatFlowResponse) SetStatusCode(v int32) *ListChatFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChatFlowResponse) SetBody(v *ListChatFlowResponseBody) *ListChatFlowResponse {
	s.Body = v
	return s
}

func (s *ListChatFlowResponse) Validate() error {
	return dara.Validate(s)
}
