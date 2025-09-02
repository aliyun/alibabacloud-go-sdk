// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDIAlarmRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDIAlarmRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDIAlarmRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetDIAlarmRuleResponseBody) *GetDIAlarmRuleResponse
	GetBody() *GetDIAlarmRuleResponseBody
}

type GetDIAlarmRuleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDIAlarmRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDIAlarmRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDIAlarmRuleResponse) GoString() string {
	return s.String()
}

func (s *GetDIAlarmRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDIAlarmRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDIAlarmRuleResponse) GetBody() *GetDIAlarmRuleResponseBody {
	return s.Body
}

func (s *GetDIAlarmRuleResponse) SetHeaders(v map[string]*string) *GetDIAlarmRuleResponse {
	s.Headers = v
	return s
}

func (s *GetDIAlarmRuleResponse) SetStatusCode(v int32) *GetDIAlarmRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDIAlarmRuleResponse) SetBody(v *GetDIAlarmRuleResponseBody) *GetDIAlarmRuleResponse {
	s.Body = v
	return s
}

func (s *GetDIAlarmRuleResponse) Validate() error {
	return dara.Validate(s)
}
