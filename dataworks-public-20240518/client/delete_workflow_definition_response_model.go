// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkflowDefinitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWorkflowDefinitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWorkflowDefinitionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWorkflowDefinitionResponseBody) *DeleteWorkflowDefinitionResponse
	GetBody() *DeleteWorkflowDefinitionResponseBody
}

type DeleteWorkflowDefinitionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkflowDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkflowDefinitionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkflowDefinitionResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowDefinitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWorkflowDefinitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWorkflowDefinitionResponse) GetBody() *DeleteWorkflowDefinitionResponseBody {
	return s.Body
}

func (s *DeleteWorkflowDefinitionResponse) SetHeaders(v map[string]*string) *DeleteWorkflowDefinitionResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkflowDefinitionResponse) SetStatusCode(v int32) *DeleteWorkflowDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkflowDefinitionResponse) SetBody(v *DeleteWorkflowDefinitionResponseBody) *DeleteWorkflowDefinitionResponse {
	s.Body = v
	return s
}

func (s *DeleteWorkflowDefinitionResponse) Validate() error {
	return dara.Validate(s)
}
