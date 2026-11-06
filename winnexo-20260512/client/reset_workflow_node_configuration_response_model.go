// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetWorkflowNodeConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetWorkflowNodeConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetWorkflowNodeConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *ResetWorkflowNodeConfigurationResponseBody) *ResetWorkflowNodeConfigurationResponse
	GetBody() *ResetWorkflowNodeConfigurationResponseBody
}

type ResetWorkflowNodeConfigurationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetWorkflowNodeConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetWorkflowNodeConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetWorkflowNodeConfigurationResponse) GoString() string {
	return s.String()
}

func (s *ResetWorkflowNodeConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetWorkflowNodeConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetWorkflowNodeConfigurationResponse) GetBody() *ResetWorkflowNodeConfigurationResponseBody {
	return s.Body
}

func (s *ResetWorkflowNodeConfigurationResponse) SetHeaders(v map[string]*string) *ResetWorkflowNodeConfigurationResponse {
	s.Headers = v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponse) SetStatusCode(v int32) *ResetWorkflowNodeConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponse) SetBody(v *ResetWorkflowNodeConfigurationResponseBody) *ResetWorkflowNodeConfigurationResponse {
	s.Body = v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
