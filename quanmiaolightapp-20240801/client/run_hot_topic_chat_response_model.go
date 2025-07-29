// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunHotTopicChatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunHotTopicChatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunHotTopicChatResponse
	GetStatusCode() *int32
	SetBody(v *RunHotTopicChatResponseBody) *RunHotTopicChatResponse
	GetBody() *RunHotTopicChatResponseBody
}

type RunHotTopicChatResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunHotTopicChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunHotTopicChatResponse) String() string {
	return dara.Prettify(s)
}

func (s RunHotTopicChatResponse) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunHotTopicChatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunHotTopicChatResponse) GetBody() *RunHotTopicChatResponseBody {
	return s.Body
}

func (s *RunHotTopicChatResponse) SetHeaders(v map[string]*string) *RunHotTopicChatResponse {
	s.Headers = v
	return s
}

func (s *RunHotTopicChatResponse) SetStatusCode(v int32) *RunHotTopicChatResponse {
	s.StatusCode = &v
	return s
}

func (s *RunHotTopicChatResponse) SetBody(v *RunHotTopicChatResponseBody) *RunHotTopicChatResponse {
	s.Body = v
	return s
}

func (s *RunHotTopicChatResponse) Validate() error {
	return dara.Validate(s)
}
