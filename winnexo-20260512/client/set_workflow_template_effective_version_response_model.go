// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWorkflowTemplateEffectiveVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetWorkflowTemplateEffectiveVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetWorkflowTemplateEffectiveVersionResponse
	GetStatusCode() *int32
	SetBody(v *SetWorkflowTemplateEffectiveVersionResponseBody) *SetWorkflowTemplateEffectiveVersionResponse
	GetBody() *SetWorkflowTemplateEffectiveVersionResponseBody
}

type SetWorkflowTemplateEffectiveVersionResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetWorkflowTemplateEffectiveVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetWorkflowTemplateEffectiveVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s SetWorkflowTemplateEffectiveVersionResponse) GoString() string {
	return s.String()
}

func (s *SetWorkflowTemplateEffectiveVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetWorkflowTemplateEffectiveVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetWorkflowTemplateEffectiveVersionResponse) GetBody() *SetWorkflowTemplateEffectiveVersionResponseBody {
	return s.Body
}

func (s *SetWorkflowTemplateEffectiveVersionResponse) SetHeaders(v map[string]*string) *SetWorkflowTemplateEffectiveVersionResponse {
	s.Headers = v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionResponse) SetStatusCode(v int32) *SetWorkflowTemplateEffectiveVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionResponse) SetBody(v *SetWorkflowTemplateEffectiveVersionResponseBody) *SetWorkflowTemplateEffectiveVersionResponse {
	s.Body = v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
