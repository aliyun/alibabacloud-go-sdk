// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameWorkflowDefinitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenameWorkflowDefinitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenameWorkflowDefinitionResponse
	GetStatusCode() *int32
	SetBody(v *RenameWorkflowDefinitionResponseBody) *RenameWorkflowDefinitionResponse
	GetBody() *RenameWorkflowDefinitionResponseBody
}

type RenameWorkflowDefinitionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenameWorkflowDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenameWorkflowDefinitionResponse) String() string {
	return dara.Prettify(s)
}

func (s RenameWorkflowDefinitionResponse) GoString() string {
	return s.String()
}

func (s *RenameWorkflowDefinitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenameWorkflowDefinitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenameWorkflowDefinitionResponse) GetBody() *RenameWorkflowDefinitionResponseBody {
	return s.Body
}

func (s *RenameWorkflowDefinitionResponse) SetHeaders(v map[string]*string) *RenameWorkflowDefinitionResponse {
	s.Headers = v
	return s
}

func (s *RenameWorkflowDefinitionResponse) SetStatusCode(v int32) *RenameWorkflowDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameWorkflowDefinitionResponse) SetBody(v *RenameWorkflowDefinitionResponseBody) *RenameWorkflowDefinitionResponse {
	s.Body = v
	return s
}

func (s *RenameWorkflowDefinitionResponse) Validate() error {
	return dara.Validate(s)
}
