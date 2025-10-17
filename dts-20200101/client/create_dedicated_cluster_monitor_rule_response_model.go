// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDedicatedClusterMonitorRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDedicatedClusterMonitorRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDedicatedClusterMonitorRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateDedicatedClusterMonitorRuleResponseBody) *CreateDedicatedClusterMonitorRuleResponse
	GetBody() *CreateDedicatedClusterMonitorRuleResponseBody
}

type CreateDedicatedClusterMonitorRuleResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDedicatedClusterMonitorRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDedicatedClusterMonitorRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDedicatedClusterMonitorRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateDedicatedClusterMonitorRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDedicatedClusterMonitorRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDedicatedClusterMonitorRuleResponse) GetBody() *CreateDedicatedClusterMonitorRuleResponseBody {
	return s.Body
}

func (s *CreateDedicatedClusterMonitorRuleResponse) SetHeaders(v map[string]*string) *CreateDedicatedClusterMonitorRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateDedicatedClusterMonitorRuleResponse) SetStatusCode(v int32) *CreateDedicatedClusterMonitorRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDedicatedClusterMonitorRuleResponse) SetBody(v *CreateDedicatedClusterMonitorRuleResponseBody) *CreateDedicatedClusterMonitorRuleResponse {
	s.Body = v
	return s
}

func (s *CreateDedicatedClusterMonitorRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
