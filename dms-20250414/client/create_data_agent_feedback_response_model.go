// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataAgentFeedbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataAgentFeedbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataAgentFeedbackResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataAgentFeedbackResponseBody) *CreateDataAgentFeedbackResponse
	GetBody() *CreateDataAgentFeedbackResponseBody
}

type CreateDataAgentFeedbackResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataAgentFeedbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataAgentFeedbackResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentFeedbackResponse) GoString() string {
	return s.String()
}

func (s *CreateDataAgentFeedbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataAgentFeedbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataAgentFeedbackResponse) GetBody() *CreateDataAgentFeedbackResponseBody {
	return s.Body
}

func (s *CreateDataAgentFeedbackResponse) SetHeaders(v map[string]*string) *CreateDataAgentFeedbackResponse {
	s.Headers = v
	return s
}

func (s *CreateDataAgentFeedbackResponse) SetStatusCode(v int32) *CreateDataAgentFeedbackResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataAgentFeedbackResponse) SetBody(v *CreateDataAgentFeedbackResponseBody) *CreateDataAgentFeedbackResponse {
	s.Body = v
	return s
}

func (s *CreateDataAgentFeedbackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
