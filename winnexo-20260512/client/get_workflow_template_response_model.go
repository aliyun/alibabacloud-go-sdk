// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkflowTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkflowTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkflowTemplateResponseBody) *GetWorkflowTemplateResponse
	GetBody() *GetWorkflowTemplateResponseBody
}

type GetWorkflowTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkflowTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkflowTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetWorkflowTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkflowTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkflowTemplateResponse) GetBody() *GetWorkflowTemplateResponseBody {
	return s.Body
}

func (s *GetWorkflowTemplateResponse) SetHeaders(v map[string]*string) *GetWorkflowTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetWorkflowTemplateResponse) SetStatusCode(v int32) *GetWorkflowTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkflowTemplateResponse) SetBody(v *GetWorkflowTemplateResponseBody) *GetWorkflowTemplateResponse {
	s.Body = v
	return s
}

func (s *GetWorkflowTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
