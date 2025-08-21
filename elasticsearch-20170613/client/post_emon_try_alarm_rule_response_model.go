// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostEmonTryAlarmRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PostEmonTryAlarmRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PostEmonTryAlarmRuleResponse
	GetStatusCode() *int32
	SetBody(v *PostEmonTryAlarmRuleResponseBody) *PostEmonTryAlarmRuleResponse
	GetBody() *PostEmonTryAlarmRuleResponseBody
}

type PostEmonTryAlarmRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PostEmonTryAlarmRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostEmonTryAlarmRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s PostEmonTryAlarmRuleResponse) GoString() string {
	return s.String()
}

func (s *PostEmonTryAlarmRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PostEmonTryAlarmRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PostEmonTryAlarmRuleResponse) GetBody() *PostEmonTryAlarmRuleResponseBody {
	return s.Body
}

func (s *PostEmonTryAlarmRuleResponse) SetHeaders(v map[string]*string) *PostEmonTryAlarmRuleResponse {
	s.Headers = v
	return s
}

func (s *PostEmonTryAlarmRuleResponse) SetStatusCode(v int32) *PostEmonTryAlarmRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *PostEmonTryAlarmRuleResponse) SetBody(v *PostEmonTryAlarmRuleResponseBody) *PostEmonTryAlarmRuleResponse {
	s.Body = v
	return s
}

func (s *PostEmonTryAlarmRuleResponse) Validate() error {
	return dara.Validate(s)
}
