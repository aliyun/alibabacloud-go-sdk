// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowDAGResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWorkflowDAGResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWorkflowDAGResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWorkflowDAGResponseBody) *UpdateWorkflowDAGResponse
	GetBody() *UpdateWorkflowDAGResponseBody
}

type UpdateWorkflowDAGResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkflowDAGResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkflowDAGResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowDAGResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDAGResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWorkflowDAGResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWorkflowDAGResponse) GetBody() *UpdateWorkflowDAGResponseBody {
	return s.Body
}

func (s *UpdateWorkflowDAGResponse) SetHeaders(v map[string]*string) *UpdateWorkflowDAGResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkflowDAGResponse) SetStatusCode(v int32) *UpdateWorkflowDAGResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkflowDAGResponse) SetBody(v *UpdateWorkflowDAGResponseBody) *UpdateWorkflowDAGResponse {
	s.Body = v
	return s
}

func (s *UpdateWorkflowDAGResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
