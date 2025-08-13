// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAssociatedResourceRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAssociatedResourceRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAssociatedResourceRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAssociatedResourceRuleResponseBody) *UpdateAssociatedResourceRuleResponse
	GetBody() *UpdateAssociatedResourceRuleResponseBody
}

type UpdateAssociatedResourceRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAssociatedResourceRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAssociatedResourceRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAssociatedResourceRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateAssociatedResourceRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAssociatedResourceRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAssociatedResourceRuleResponse) GetBody() *UpdateAssociatedResourceRuleResponseBody {
	return s.Body
}

func (s *UpdateAssociatedResourceRuleResponse) SetHeaders(v map[string]*string) *UpdateAssociatedResourceRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateAssociatedResourceRuleResponse) SetStatusCode(v int32) *UpdateAssociatedResourceRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAssociatedResourceRuleResponse) SetBody(v *UpdateAssociatedResourceRuleResponseBody) *UpdateAssociatedResourceRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateAssociatedResourceRuleResponse) Validate() error {
	return dara.Validate(s)
}
