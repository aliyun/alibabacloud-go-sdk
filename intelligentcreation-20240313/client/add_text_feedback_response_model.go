// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTextFeedbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddTextFeedbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddTextFeedbackResponse
	GetStatusCode() *int32
	SetBody(v *AddTextFeedbackResponseBody) *AddTextFeedbackResponse
	GetBody() *AddTextFeedbackResponseBody
}

type AddTextFeedbackResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTextFeedbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTextFeedbackResponse) String() string {
	return dara.Prettify(s)
}

func (s AddTextFeedbackResponse) GoString() string {
	return s.String()
}

func (s *AddTextFeedbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddTextFeedbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddTextFeedbackResponse) GetBody() *AddTextFeedbackResponseBody {
	return s.Body
}

func (s *AddTextFeedbackResponse) SetHeaders(v map[string]*string) *AddTextFeedbackResponse {
	s.Headers = v
	return s
}

func (s *AddTextFeedbackResponse) SetStatusCode(v int32) *AddTextFeedbackResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTextFeedbackResponse) SetBody(v *AddTextFeedbackResponseBody) *AddTextFeedbackResponse {
	s.Body = v
	return s
}

func (s *AddTextFeedbackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
