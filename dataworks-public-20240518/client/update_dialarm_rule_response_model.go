// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDIAlarmRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDIAlarmRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDIAlarmRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDIAlarmRuleResponseBody) *UpdateDIAlarmRuleResponse
	GetBody() *UpdateDIAlarmRuleResponseBody
}

type UpdateDIAlarmRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDIAlarmRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDIAlarmRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIAlarmRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateDIAlarmRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDIAlarmRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDIAlarmRuleResponse) GetBody() *UpdateDIAlarmRuleResponseBody {
	return s.Body
}

func (s *UpdateDIAlarmRuleResponse) SetHeaders(v map[string]*string) *UpdateDIAlarmRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateDIAlarmRuleResponse) SetStatusCode(v int32) *UpdateDIAlarmRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDIAlarmRuleResponse) SetBody(v *UpdateDIAlarmRuleResponseBody) *UpdateDIAlarmRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateDIAlarmRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
