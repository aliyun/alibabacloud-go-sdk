// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkflowTemplateVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWorkflowTemplateVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWorkflowTemplateVersionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWorkflowTemplateVersionResponseBody) *DeleteWorkflowTemplateVersionResponse
	GetBody() *DeleteWorkflowTemplateVersionResponseBody
}

type DeleteWorkflowTemplateVersionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkflowTemplateVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkflowTemplateVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkflowTemplateVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowTemplateVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWorkflowTemplateVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWorkflowTemplateVersionResponse) GetBody() *DeleteWorkflowTemplateVersionResponseBody {
	return s.Body
}

func (s *DeleteWorkflowTemplateVersionResponse) SetHeaders(v map[string]*string) *DeleteWorkflowTemplateVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkflowTemplateVersionResponse) SetStatusCode(v int32) *DeleteWorkflowTemplateVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionResponse) SetBody(v *DeleteWorkflowTemplateVersionResponseBody) *DeleteWorkflowTemplateVersionResponse {
	s.Body = v
	return s
}

func (s *DeleteWorkflowTemplateVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
