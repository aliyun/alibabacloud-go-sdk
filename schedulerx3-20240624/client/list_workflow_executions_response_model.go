// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowExecutionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkflowExecutionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkflowExecutionsResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkflowExecutionsResponseBody) *ListWorkflowExecutionsResponse
	GetBody() *ListWorkflowExecutionsResponseBody
}

type ListWorkflowExecutionsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkflowExecutionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkflowExecutionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowExecutionsResponse) GoString() string {
	return s.String()
}

func (s *ListWorkflowExecutionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkflowExecutionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkflowExecutionsResponse) GetBody() *ListWorkflowExecutionsResponseBody {
	return s.Body
}

func (s *ListWorkflowExecutionsResponse) SetHeaders(v map[string]*string) *ListWorkflowExecutionsResponse {
	s.Headers = v
	return s
}

func (s *ListWorkflowExecutionsResponse) SetStatusCode(v int32) *ListWorkflowExecutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkflowExecutionsResponse) SetBody(v *ListWorkflowExecutionsResponseBody) *ListWorkflowExecutionsResponse {
	s.Body = v
	return s
}

func (s *ListWorkflowExecutionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
