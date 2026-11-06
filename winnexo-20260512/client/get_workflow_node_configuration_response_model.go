// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowNodeConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkflowNodeConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkflowNodeConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkflowNodeConfigurationResponseBody) *GetWorkflowNodeConfigurationResponse
	GetBody() *GetWorkflowNodeConfigurationResponseBody
}

type GetWorkflowNodeConfigurationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkflowNodeConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkflowNodeConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowNodeConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetWorkflowNodeConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkflowNodeConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkflowNodeConfigurationResponse) GetBody() *GetWorkflowNodeConfigurationResponseBody {
	return s.Body
}

func (s *GetWorkflowNodeConfigurationResponse) SetHeaders(v map[string]*string) *GetWorkflowNodeConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetWorkflowNodeConfigurationResponse) SetStatusCode(v int32) *GetWorkflowNodeConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkflowNodeConfigurationResponse) SetBody(v *GetWorkflowNodeConfigurationResponseBody) *GetWorkflowNodeConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetWorkflowNodeConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
