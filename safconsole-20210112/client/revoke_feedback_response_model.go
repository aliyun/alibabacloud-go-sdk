// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeFeedbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeFeedbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeFeedbackResponse
	GetStatusCode() *int32
	SetBody(v *RevokeFeedbackResponseBody) *RevokeFeedbackResponse
	GetBody() *RevokeFeedbackResponseBody
}

type RevokeFeedbackResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeFeedbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeFeedbackResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeFeedbackResponse) GoString() string {
	return s.String()
}

func (s *RevokeFeedbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeFeedbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeFeedbackResponse) GetBody() *RevokeFeedbackResponseBody {
	return s.Body
}

func (s *RevokeFeedbackResponse) SetHeaders(v map[string]*string) *RevokeFeedbackResponse {
	s.Headers = v
	return s
}

func (s *RevokeFeedbackResponse) SetStatusCode(v int32) *RevokeFeedbackResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeFeedbackResponse) SetBody(v *RevokeFeedbackResponseBody) *RevokeFeedbackResponse {
	s.Body = v
	return s
}

func (s *RevokeFeedbackResponse) Validate() error {
	return dara.Validate(s)
}
