// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowDefinitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWorkflowDefinitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWorkflowDefinitionResponse
	GetStatusCode() *int32
	SetBody(v *CreateWorkflowDefinitionResponseBody) *CreateWorkflowDefinitionResponse
	GetBody() *CreateWorkflowDefinitionResponseBody
}

type CreateWorkflowDefinitionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkflowDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkflowDefinitionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowDefinitionResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkflowDefinitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWorkflowDefinitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWorkflowDefinitionResponse) GetBody() *CreateWorkflowDefinitionResponseBody {
	return s.Body
}

func (s *CreateWorkflowDefinitionResponse) SetHeaders(v map[string]*string) *CreateWorkflowDefinitionResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkflowDefinitionResponse) SetStatusCode(v int32) *CreateWorkflowDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkflowDefinitionResponse) SetBody(v *CreateWorkflowDefinitionResponseBody) *CreateWorkflowDefinitionResponse {
	s.Body = v
	return s
}

func (s *CreateWorkflowDefinitionResponse) Validate() error {
	return dara.Validate(s)
}
