// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskAssignRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTaskAssignRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTaskAssignRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateTaskAssignRuleResponseBody) *CreateTaskAssignRuleResponse
	GetBody() *CreateTaskAssignRuleResponseBody
}

type CreateTaskAssignRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTaskAssignRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTaskAssignRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskAssignRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateTaskAssignRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTaskAssignRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTaskAssignRuleResponse) GetBody() *CreateTaskAssignRuleResponseBody {
	return s.Body
}

func (s *CreateTaskAssignRuleResponse) SetHeaders(v map[string]*string) *CreateTaskAssignRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateTaskAssignRuleResponse) SetStatusCode(v int32) *CreateTaskAssignRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTaskAssignRuleResponse) SetBody(v *CreateTaskAssignRuleResponseBody) *CreateTaskAssignRuleResponse {
	s.Body = v
	return s
}

func (s *CreateTaskAssignRuleResponse) Validate() error {
	return dara.Validate(s)
}
