// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowDirectoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkflowDirectoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkflowDirectoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkflowDirectoriesResponseBody) *ListWorkflowDirectoriesResponse
	GetBody() *ListWorkflowDirectoriesResponseBody
}

type ListWorkflowDirectoriesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkflowDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkflowDirectoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *ListWorkflowDirectoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkflowDirectoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkflowDirectoriesResponse) GetBody() *ListWorkflowDirectoriesResponseBody {
	return s.Body
}

func (s *ListWorkflowDirectoriesResponse) SetHeaders(v map[string]*string) *ListWorkflowDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *ListWorkflowDirectoriesResponse) SetStatusCode(v int32) *ListWorkflowDirectoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkflowDirectoriesResponse) SetBody(v *ListWorkflowDirectoriesResponseBody) *ListWorkflowDirectoriesResponse {
	s.Body = v
	return s
}

func (s *ListWorkflowDirectoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
