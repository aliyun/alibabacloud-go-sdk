// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkflowsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkflowsResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkflowsResponseBody) *ListWorkflowsResponse
	GetBody() *ListWorkflowsResponseBody
}

type ListWorkflowsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkflowsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkflowsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowsResponse) GoString() string {
	return s.String()
}

func (s *ListWorkflowsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkflowsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkflowsResponse) GetBody() *ListWorkflowsResponseBody {
	return s.Body
}

func (s *ListWorkflowsResponse) SetHeaders(v map[string]*string) *ListWorkflowsResponse {
	s.Headers = v
	return s
}

func (s *ListWorkflowsResponse) SetStatusCode(v int32) *ListWorkflowsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkflowsResponse) SetBody(v *ListWorkflowsResponseBody) *ListWorkflowsResponse {
	s.Body = v
	return s
}

func (s *ListWorkflowsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
