// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowExecutionDAGResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkflowExecutionDAGResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkflowExecutionDAGResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkflowExecutionDAGResponseBody) *GetWorkflowExecutionDAGResponse
	GetBody() *GetWorkflowExecutionDAGResponseBody
}

type GetWorkflowExecutionDAGResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkflowExecutionDAGResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkflowExecutionDAGResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowExecutionDAGResponse) GoString() string {
	return s.String()
}

func (s *GetWorkflowExecutionDAGResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkflowExecutionDAGResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkflowExecutionDAGResponse) GetBody() *GetWorkflowExecutionDAGResponseBody {
	return s.Body
}

func (s *GetWorkflowExecutionDAGResponse) SetHeaders(v map[string]*string) *GetWorkflowExecutionDAGResponse {
	s.Headers = v
	return s
}

func (s *GetWorkflowExecutionDAGResponse) SetStatusCode(v int32) *GetWorkflowExecutionDAGResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkflowExecutionDAGResponse) SetBody(v *GetWorkflowExecutionDAGResponseBody) *GetWorkflowExecutionDAGResponse {
	s.Body = v
	return s
}

func (s *GetWorkflowExecutionDAGResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
