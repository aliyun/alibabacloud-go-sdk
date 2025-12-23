// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWorkflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWorkflowResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWorkflowResponseBody) *UpdateWorkflowResponse
	GetBody() *UpdateWorkflowResponseBody
}

type UpdateWorkflowResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkflowResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWorkflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWorkflowResponse) GetBody() *UpdateWorkflowResponseBody {
	return s.Body
}

func (s *UpdateWorkflowResponse) SetHeaders(v map[string]*string) *UpdateWorkflowResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkflowResponse) SetStatusCode(v int32) *UpdateWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkflowResponse) SetBody(v *UpdateWorkflowResponseBody) *UpdateWorkflowResponse {
	s.Body = v
	return s
}

func (s *UpdateWorkflowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
