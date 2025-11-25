// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchSchedulerRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchSchedulerRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchSchedulerRuleResponse
	GetStatusCode() *int32
	SetBody(v *SwitchSchedulerRuleResponseBody) *SwitchSchedulerRuleResponse
	GetBody() *SwitchSchedulerRuleResponseBody
}

type SwitchSchedulerRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchSchedulerRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchSchedulerRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchSchedulerRuleResponse) GoString() string {
	return s.String()
}

func (s *SwitchSchedulerRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchSchedulerRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchSchedulerRuleResponse) GetBody() *SwitchSchedulerRuleResponseBody {
	return s.Body
}

func (s *SwitchSchedulerRuleResponse) SetHeaders(v map[string]*string) *SwitchSchedulerRuleResponse {
	s.Headers = v
	return s
}

func (s *SwitchSchedulerRuleResponse) SetStatusCode(v int32) *SwitchSchedulerRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchSchedulerRuleResponse) SetBody(v *SwitchSchedulerRuleResponseBody) *SwitchSchedulerRuleResponse {
	s.Body = v
	return s
}

func (s *SwitchSchedulerRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
