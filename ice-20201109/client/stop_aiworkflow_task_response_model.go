// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAIWorkflowTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopAIWorkflowTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopAIWorkflowTaskResponse
	GetStatusCode() *int32
	SetBody(v *StopAIWorkflowTaskResponseBody) *StopAIWorkflowTaskResponse
	GetBody() *StopAIWorkflowTaskResponseBody
}

type StopAIWorkflowTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopAIWorkflowTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopAIWorkflowTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StopAIWorkflowTaskResponse) GoString() string {
	return s.String()
}

func (s *StopAIWorkflowTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopAIWorkflowTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopAIWorkflowTaskResponse) GetBody() *StopAIWorkflowTaskResponseBody {
	return s.Body
}

func (s *StopAIWorkflowTaskResponse) SetHeaders(v map[string]*string) *StopAIWorkflowTaskResponse {
	s.Headers = v
	return s
}

func (s *StopAIWorkflowTaskResponse) SetStatusCode(v int32) *StopAIWorkflowTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopAIWorkflowTaskResponse) SetBody(v *StopAIWorkflowTaskResponseBody) *StopAIWorkflowTaskResponse {
	s.Body = v
	return s
}

func (s *StopAIWorkflowTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
