// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFeedbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FeedbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FeedbackResponse
	GetStatusCode() *int32
	SetBody(v *FeedbackResponseBody) *FeedbackResponse
	GetBody() *FeedbackResponseBody
}

type FeedbackResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FeedbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FeedbackResponse) String() string {
	return dara.Prettify(s)
}

func (s FeedbackResponse) GoString() string {
	return s.String()
}

func (s *FeedbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FeedbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FeedbackResponse) GetBody() *FeedbackResponseBody {
	return s.Body
}

func (s *FeedbackResponse) SetHeaders(v map[string]*string) *FeedbackResponse {
	s.Headers = v
	return s
}

func (s *FeedbackResponse) SetStatusCode(v int32) *FeedbackResponse {
	s.StatusCode = &v
	return s
}

func (s *FeedbackResponse) SetBody(v *FeedbackResponseBody) *FeedbackResponse {
	s.Body = v
	return s
}

func (s *FeedbackResponse) Validate() error {
	return dara.Validate(s)
}
