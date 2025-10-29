// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveWorkflowDefinitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MoveWorkflowDefinitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MoveWorkflowDefinitionResponse
	GetStatusCode() *int32
	SetBody(v *MoveWorkflowDefinitionResponseBody) *MoveWorkflowDefinitionResponse
	GetBody() *MoveWorkflowDefinitionResponseBody
}

type MoveWorkflowDefinitionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveWorkflowDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveWorkflowDefinitionResponse) String() string {
	return dara.Prettify(s)
}

func (s MoveWorkflowDefinitionResponse) GoString() string {
	return s.String()
}

func (s *MoveWorkflowDefinitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MoveWorkflowDefinitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MoveWorkflowDefinitionResponse) GetBody() *MoveWorkflowDefinitionResponseBody {
	return s.Body
}

func (s *MoveWorkflowDefinitionResponse) SetHeaders(v map[string]*string) *MoveWorkflowDefinitionResponse {
	s.Headers = v
	return s
}

func (s *MoveWorkflowDefinitionResponse) SetStatusCode(v int32) *MoveWorkflowDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveWorkflowDefinitionResponse) SetBody(v *MoveWorkflowDefinitionResponseBody) *MoveWorkflowDefinitionResponse {
	s.Body = v
	return s
}

func (s *MoveWorkflowDefinitionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
