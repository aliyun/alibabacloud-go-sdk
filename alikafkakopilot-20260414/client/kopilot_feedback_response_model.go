// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKopilotFeedbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *KopilotFeedbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *KopilotFeedbackResponse
	GetStatusCode() *int32
	SetBody(v *KopilotFeedbackResponseBody) *KopilotFeedbackResponse
	GetBody() *KopilotFeedbackResponseBody
}

type KopilotFeedbackResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KopilotFeedbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s KopilotFeedbackResponse) String() string {
	return dara.Prettify(s)
}

func (s KopilotFeedbackResponse) GoString() string {
	return s.String()
}

func (s *KopilotFeedbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *KopilotFeedbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *KopilotFeedbackResponse) GetBody() *KopilotFeedbackResponseBody {
	return s.Body
}

func (s *KopilotFeedbackResponse) SetHeaders(v map[string]*string) *KopilotFeedbackResponse {
	s.Headers = v
	return s
}

func (s *KopilotFeedbackResponse) SetStatusCode(v int32) *KopilotFeedbackResponse {
	s.StatusCode = &v
	return s
}

func (s *KopilotFeedbackResponse) SetBody(v *KopilotFeedbackResponseBody) *KopilotFeedbackResponse {
	s.Body = v
	return s
}

func (s *KopilotFeedbackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
