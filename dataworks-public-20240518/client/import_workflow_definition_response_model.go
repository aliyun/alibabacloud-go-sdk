// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportWorkflowDefinitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportWorkflowDefinitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportWorkflowDefinitionResponse
	GetStatusCode() *int32
	SetBody(v *ImportWorkflowDefinitionResponseBody) *ImportWorkflowDefinitionResponse
	GetBody() *ImportWorkflowDefinitionResponseBody
}

type ImportWorkflowDefinitionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportWorkflowDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportWorkflowDefinitionResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportWorkflowDefinitionResponse) GoString() string {
	return s.String()
}

func (s *ImportWorkflowDefinitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportWorkflowDefinitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportWorkflowDefinitionResponse) GetBody() *ImportWorkflowDefinitionResponseBody {
	return s.Body
}

func (s *ImportWorkflowDefinitionResponse) SetHeaders(v map[string]*string) *ImportWorkflowDefinitionResponse {
	s.Headers = v
	return s
}

func (s *ImportWorkflowDefinitionResponse) SetStatusCode(v int32) *ImportWorkflowDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportWorkflowDefinitionResponse) SetBody(v *ImportWorkflowDefinitionResponseBody) *ImportWorkflowDefinitionResponse {
	s.Body = v
	return s
}

func (s *ImportWorkflowDefinitionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
