// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowDefinitionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkflowDefinitionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkflowDefinitionsResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkflowDefinitionsResponseBody) *ListWorkflowDefinitionsResponse
	GetBody() *ListWorkflowDefinitionsResponseBody
}

type ListWorkflowDefinitionsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkflowDefinitionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkflowDefinitionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowDefinitionsResponse) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkflowDefinitionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkflowDefinitionsResponse) GetBody() *ListWorkflowDefinitionsResponseBody {
	return s.Body
}

func (s *ListWorkflowDefinitionsResponse) SetHeaders(v map[string]*string) *ListWorkflowDefinitionsResponse {
	s.Headers = v
	return s
}

func (s *ListWorkflowDefinitionsResponse) SetStatusCode(v int32) *ListWorkflowDefinitionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkflowDefinitionsResponse) SetBody(v *ListWorkflowDefinitionsResponseBody) *ListWorkflowDefinitionsResponse {
	s.Body = v
	return s
}

func (s *ListWorkflowDefinitionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
