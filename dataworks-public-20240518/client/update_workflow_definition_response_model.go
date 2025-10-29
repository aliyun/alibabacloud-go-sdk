// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowDefinitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWorkflowDefinitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWorkflowDefinitionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWorkflowDefinitionResponseBody) *UpdateWorkflowDefinitionResponse
	GetBody() *UpdateWorkflowDefinitionResponseBody
}

type UpdateWorkflowDefinitionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkflowDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkflowDefinitionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowDefinitionResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDefinitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWorkflowDefinitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWorkflowDefinitionResponse) GetBody() *UpdateWorkflowDefinitionResponseBody {
	return s.Body
}

func (s *UpdateWorkflowDefinitionResponse) SetHeaders(v map[string]*string) *UpdateWorkflowDefinitionResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkflowDefinitionResponse) SetStatusCode(v int32) *UpdateWorkflowDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkflowDefinitionResponse) SetBody(v *UpdateWorkflowDefinitionResponseBody) *UpdateWorkflowDefinitionResponse {
	s.Body = v
	return s
}

func (s *UpdateWorkflowDefinitionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
