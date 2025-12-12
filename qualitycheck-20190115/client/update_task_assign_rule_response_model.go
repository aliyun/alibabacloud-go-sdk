// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskAssignRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTaskAssignRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTaskAssignRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTaskAssignRuleResponseBody) *UpdateTaskAssignRuleResponse
	GetBody() *UpdateTaskAssignRuleResponseBody
}

type UpdateTaskAssignRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskAssignRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskAssignRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskAssignRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskAssignRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTaskAssignRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTaskAssignRuleResponse) GetBody() *UpdateTaskAssignRuleResponseBody {
	return s.Body
}

func (s *UpdateTaskAssignRuleResponse) SetHeaders(v map[string]*string) *UpdateTaskAssignRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskAssignRuleResponse) SetStatusCode(v int32) *UpdateTaskAssignRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskAssignRuleResponse) SetBody(v *UpdateTaskAssignRuleResponseBody) *UpdateTaskAssignRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateTaskAssignRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
