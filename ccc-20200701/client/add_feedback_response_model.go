// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFeedbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddFeedbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddFeedbackResponse
	GetStatusCode() *int32
	SetBody(v *AddFeedbackResponseBody) *AddFeedbackResponse
	GetBody() *AddFeedbackResponseBody
}

type AddFeedbackResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFeedbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFeedbackResponse) String() string {
	return dara.Prettify(s)
}

func (s AddFeedbackResponse) GoString() string {
	return s.String()
}

func (s *AddFeedbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddFeedbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddFeedbackResponse) GetBody() *AddFeedbackResponseBody {
	return s.Body
}

func (s *AddFeedbackResponse) SetHeaders(v map[string]*string) *AddFeedbackResponse {
	s.Headers = v
	return s
}

func (s *AddFeedbackResponse) SetStatusCode(v int32) *AddFeedbackResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFeedbackResponse) SetBody(v *AddFeedbackResponseBody) *AddFeedbackResponse {
	s.Body = v
	return s
}

func (s *AddFeedbackResponse) Validate() error {
	return dara.Validate(s)
}
