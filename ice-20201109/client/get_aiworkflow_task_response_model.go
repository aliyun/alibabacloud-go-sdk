// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIWorkflowTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAIWorkflowTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAIWorkflowTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetAIWorkflowTaskResponseBody) *GetAIWorkflowTaskResponse
	GetBody() *GetAIWorkflowTaskResponseBody
}

type GetAIWorkflowTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAIWorkflowTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAIWorkflowTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAIWorkflowTaskResponse) GoString() string {
	return s.String()
}

func (s *GetAIWorkflowTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAIWorkflowTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAIWorkflowTaskResponse) GetBody() *GetAIWorkflowTaskResponseBody {
	return s.Body
}

func (s *GetAIWorkflowTaskResponse) SetHeaders(v map[string]*string) *GetAIWorkflowTaskResponse {
	s.Headers = v
	return s
}

func (s *GetAIWorkflowTaskResponse) SetStatusCode(v int32) *GetAIWorkflowTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAIWorkflowTaskResponse) SetBody(v *GetAIWorkflowTaskResponseBody) *GetAIWorkflowTaskResponse {
	s.Body = v
	return s
}

func (s *GetAIWorkflowTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
