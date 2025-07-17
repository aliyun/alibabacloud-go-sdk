// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDIAlarmRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDIAlarmRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDIAlarmRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDIAlarmRuleResponseBody) *DeleteDIAlarmRuleResponse
	GetBody() *DeleteDIAlarmRuleResponseBody
}

type DeleteDIAlarmRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDIAlarmRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDIAlarmRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDIAlarmRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteDIAlarmRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDIAlarmRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDIAlarmRuleResponse) GetBody() *DeleteDIAlarmRuleResponseBody {
	return s.Body
}

func (s *DeleteDIAlarmRuleResponse) SetHeaders(v map[string]*string) *DeleteDIAlarmRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteDIAlarmRuleResponse) SetStatusCode(v int32) *DeleteDIAlarmRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDIAlarmRuleResponse) SetBody(v *DeleteDIAlarmRuleResponseBody) *DeleteDIAlarmRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteDIAlarmRuleResponse) Validate() error {
	return dara.Validate(s)
}
