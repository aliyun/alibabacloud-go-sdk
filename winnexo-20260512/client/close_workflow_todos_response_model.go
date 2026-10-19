// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseWorkflowTodosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseWorkflowTodosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseWorkflowTodosResponse
	GetStatusCode() *int32
	SetBody(v *CloseWorkflowTodosResponseBody) *CloseWorkflowTodosResponse
	GetBody() *CloseWorkflowTodosResponseBody
}

type CloseWorkflowTodosResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseWorkflowTodosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseWorkflowTodosResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseWorkflowTodosResponse) GoString() string {
	return s.String()
}

func (s *CloseWorkflowTodosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseWorkflowTodosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseWorkflowTodosResponse) GetBody() *CloseWorkflowTodosResponseBody {
	return s.Body
}

func (s *CloseWorkflowTodosResponse) SetHeaders(v map[string]*string) *CloseWorkflowTodosResponse {
	s.Headers = v
	return s
}

func (s *CloseWorkflowTodosResponse) SetStatusCode(v int32) *CloseWorkflowTodosResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseWorkflowTodosResponse) SetBody(v *CloseWorkflowTodosResponseBody) *CloseWorkflowTodosResponse {
	s.Body = v
	return s
}

func (s *CloseWorkflowTodosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
