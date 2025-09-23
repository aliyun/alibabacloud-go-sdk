// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendFeedbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendFeedbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendFeedbackResponse
	GetStatusCode() *int32
	SetBody(v *SendFeedbackResponseBody) *SendFeedbackResponse
	GetBody() *SendFeedbackResponseBody
}

type SendFeedbackResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendFeedbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendFeedbackResponse) String() string {
	return dara.Prettify(s)
}

func (s SendFeedbackResponse) GoString() string {
	return s.String()
}

func (s *SendFeedbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendFeedbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendFeedbackResponse) GetBody() *SendFeedbackResponseBody {
	return s.Body
}

func (s *SendFeedbackResponse) SetHeaders(v map[string]*string) *SendFeedbackResponse {
	s.Headers = v
	return s
}

func (s *SendFeedbackResponse) SetStatusCode(v int32) *SendFeedbackResponse {
	s.StatusCode = &v
	return s
}

func (s *SendFeedbackResponse) SetBody(v *SendFeedbackResponseBody) *SendFeedbackResponse {
	s.Body = v
	return s
}

func (s *SendFeedbackResponse) Validate() error {
	return dara.Validate(s)
}
