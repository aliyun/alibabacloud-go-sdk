// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySchedulerRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySchedulerRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySchedulerRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifySchedulerRuleResponseBody) *ModifySchedulerRuleResponse
	GetBody() *ModifySchedulerRuleResponseBody
}

type ModifySchedulerRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySchedulerRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySchedulerRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySchedulerRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifySchedulerRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySchedulerRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySchedulerRuleResponse) GetBody() *ModifySchedulerRuleResponseBody {
	return s.Body
}

func (s *ModifySchedulerRuleResponse) SetHeaders(v map[string]*string) *ModifySchedulerRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifySchedulerRuleResponse) SetStatusCode(v int32) *ModifySchedulerRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySchedulerRuleResponse) SetBody(v *ModifySchedulerRuleResponseBody) *ModifySchedulerRuleResponse {
	s.Body = v
	return s
}

func (s *ModifySchedulerRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
