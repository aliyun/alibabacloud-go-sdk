// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowDefinitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkflowDefinitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkflowDefinitionResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkflowDefinitionResponseBody) *GetWorkflowDefinitionResponse
	GetBody() *GetWorkflowDefinitionResponseBody
}

type GetWorkflowDefinitionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkflowDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkflowDefinitionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowDefinitionResponse) GoString() string {
	return s.String()
}

func (s *GetWorkflowDefinitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkflowDefinitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkflowDefinitionResponse) GetBody() *GetWorkflowDefinitionResponseBody {
	return s.Body
}

func (s *GetWorkflowDefinitionResponse) SetHeaders(v map[string]*string) *GetWorkflowDefinitionResponse {
	s.Headers = v
	return s
}

func (s *GetWorkflowDefinitionResponse) SetStatusCode(v int32) *GetWorkflowDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkflowDefinitionResponse) SetBody(v *GetWorkflowDefinitionResponseBody) *GetWorkflowDefinitionResponse {
	s.Body = v
	return s
}

func (s *GetWorkflowDefinitionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
