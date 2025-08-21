// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTongyiChatHistorysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTongyiChatHistorysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTongyiChatHistorysResponse
	GetStatusCode() *int32
	SetBody(v *ListTongyiChatHistorysResponseBody) *ListTongyiChatHistorysResponse
	GetBody() *ListTongyiChatHistorysResponseBody
}

type ListTongyiChatHistorysResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTongyiChatHistorysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTongyiChatHistorysResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTongyiChatHistorysResponse) GoString() string {
	return s.String()
}

func (s *ListTongyiChatHistorysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTongyiChatHistorysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTongyiChatHistorysResponse) GetBody() *ListTongyiChatHistorysResponseBody {
	return s.Body
}

func (s *ListTongyiChatHistorysResponse) SetHeaders(v map[string]*string) *ListTongyiChatHistorysResponse {
	s.Headers = v
	return s
}

func (s *ListTongyiChatHistorysResponse) SetStatusCode(v int32) *ListTongyiChatHistorysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTongyiChatHistorysResponse) SetBody(v *ListTongyiChatHistorysResponseBody) *ListTongyiChatHistorysResponse {
	s.Body = v
	return s
}

func (s *ListTongyiChatHistorysResponse) Validate() error {
	return dara.Validate(s)
}
