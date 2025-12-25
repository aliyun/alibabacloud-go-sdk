// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowDAGVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWorkflowDAGVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWorkflowDAGVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWorkflowDAGVersionResponseBody) *UpdateWorkflowDAGVersionResponse
	GetBody() *UpdateWorkflowDAGVersionResponseBody
}

type UpdateWorkflowDAGVersionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkflowDAGVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkflowDAGVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowDAGVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDAGVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWorkflowDAGVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWorkflowDAGVersionResponse) GetBody() *UpdateWorkflowDAGVersionResponseBody {
	return s.Body
}

func (s *UpdateWorkflowDAGVersionResponse) SetHeaders(v map[string]*string) *UpdateWorkflowDAGVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkflowDAGVersionResponse) SetStatusCode(v int32) *UpdateWorkflowDAGVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkflowDAGVersionResponse) SetBody(v *UpdateWorkflowDAGVersionResponseBody) *UpdateWorkflowDAGVersionResponse {
	s.Body = v
	return s
}

func (s *UpdateWorkflowDAGVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
