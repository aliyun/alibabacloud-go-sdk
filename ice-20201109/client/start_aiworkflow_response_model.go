// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAIWorkflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartAIWorkflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartAIWorkflowResponse
	GetStatusCode() *int32
	SetBody(v *StartAIWorkflowResponseBody) *StartAIWorkflowResponse
	GetBody() *StartAIWorkflowResponseBody
}

type StartAIWorkflowResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartAIWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartAIWorkflowResponse) String() string {
	return dara.Prettify(s)
}

func (s StartAIWorkflowResponse) GoString() string {
	return s.String()
}

func (s *StartAIWorkflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartAIWorkflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartAIWorkflowResponse) GetBody() *StartAIWorkflowResponseBody {
	return s.Body
}

func (s *StartAIWorkflowResponse) SetHeaders(v map[string]*string) *StartAIWorkflowResponse {
	s.Headers = v
	return s
}

func (s *StartAIWorkflowResponse) SetStatusCode(v int32) *StartAIWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *StartAIWorkflowResponse) SetBody(v *StartAIWorkflowResponseBody) *StartAIWorkflowResponse {
	s.Body = v
	return s
}

func (s *StartAIWorkflowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
