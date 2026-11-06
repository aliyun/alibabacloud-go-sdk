// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowNodeConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWorkflowNodeConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWorkflowNodeConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWorkflowNodeConfigurationResponseBody) *UpdateWorkflowNodeConfigurationResponse
	GetBody() *UpdateWorkflowNodeConfigurationResponseBody
}

type UpdateWorkflowNodeConfigurationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkflowNodeConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkflowNodeConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowNodeConfigurationResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowNodeConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWorkflowNodeConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWorkflowNodeConfigurationResponse) GetBody() *UpdateWorkflowNodeConfigurationResponseBody {
	return s.Body
}

func (s *UpdateWorkflowNodeConfigurationResponse) SetHeaders(v map[string]*string) *UpdateWorkflowNodeConfigurationResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponse) SetStatusCode(v int32) *UpdateWorkflowNodeConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponse) SetBody(v *UpdateWorkflowNodeConfigurationResponseBody) *UpdateWorkflowNodeConfigurationResponse {
	s.Body = v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
