// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowTemplateVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkflowTemplateVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkflowTemplateVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkflowTemplateVersionsResponseBody) *ListWorkflowTemplateVersionsResponse
	GetBody() *ListWorkflowTemplateVersionsResponseBody
}

type ListWorkflowTemplateVersionsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkflowTemplateVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkflowTemplateVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowTemplateVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListWorkflowTemplateVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkflowTemplateVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkflowTemplateVersionsResponse) GetBody() *ListWorkflowTemplateVersionsResponseBody {
	return s.Body
}

func (s *ListWorkflowTemplateVersionsResponse) SetHeaders(v map[string]*string) *ListWorkflowTemplateVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListWorkflowTemplateVersionsResponse) SetStatusCode(v int32) *ListWorkflowTemplateVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkflowTemplateVersionsResponse) SetBody(v *ListWorkflowTemplateVersionsResponseBody) *ListWorkflowTemplateVersionsResponse {
	s.Body = v
	return s
}

func (s *ListWorkflowTemplateVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
