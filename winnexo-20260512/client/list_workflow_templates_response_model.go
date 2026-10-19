// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkflowTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkflowTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkflowTemplatesResponseBody) *ListWorkflowTemplatesResponse
	GetBody() *ListWorkflowTemplatesResponseBody
}

type ListWorkflowTemplatesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkflowTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkflowTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListWorkflowTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkflowTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkflowTemplatesResponse) GetBody() *ListWorkflowTemplatesResponseBody {
	return s.Body
}

func (s *ListWorkflowTemplatesResponse) SetHeaders(v map[string]*string) *ListWorkflowTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListWorkflowTemplatesResponse) SetStatusCode(v int32) *ListWorkflowTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkflowTemplatesResponse) SetBody(v *ListWorkflowTemplatesResponseBody) *ListWorkflowTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListWorkflowTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
