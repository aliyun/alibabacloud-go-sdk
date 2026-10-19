// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkflowTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWorkflowTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWorkflowTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWorkflowTemplateResponseBody) *DeleteWorkflowTemplateResponse
	GetBody() *DeleteWorkflowTemplateResponseBody
}

type DeleteWorkflowTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkflowTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkflowTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkflowTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWorkflowTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWorkflowTemplateResponse) GetBody() *DeleteWorkflowTemplateResponseBody {
	return s.Body
}

func (s *DeleteWorkflowTemplateResponse) SetHeaders(v map[string]*string) *DeleteWorkflowTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkflowTemplateResponse) SetStatusCode(v int32) *DeleteWorkflowTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkflowTemplateResponse) SetBody(v *DeleteWorkflowTemplateResponseBody) *DeleteWorkflowTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteWorkflowTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
