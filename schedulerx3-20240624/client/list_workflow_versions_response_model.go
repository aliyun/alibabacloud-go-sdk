// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkflowVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkflowVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkflowVersionsResponseBody) *ListWorkflowVersionsResponse
	GetBody() *ListWorkflowVersionsResponseBody
}

type ListWorkflowVersionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkflowVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkflowVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListWorkflowVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkflowVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkflowVersionsResponse) GetBody() *ListWorkflowVersionsResponseBody {
	return s.Body
}

func (s *ListWorkflowVersionsResponse) SetHeaders(v map[string]*string) *ListWorkflowVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListWorkflowVersionsResponse) SetStatusCode(v int32) *ListWorkflowVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkflowVersionsResponse) SetBody(v *ListWorkflowVersionsResponseBody) *ListWorkflowVersionsResponse {
	s.Body = v
	return s
}

func (s *ListWorkflowVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
