// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDIAlarmRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDIAlarmRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDIAlarmRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateDIAlarmRuleResponseBody) *CreateDIAlarmRuleResponse
	GetBody() *CreateDIAlarmRuleResponseBody
}

type CreateDIAlarmRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDIAlarmRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDIAlarmRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDIAlarmRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateDIAlarmRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDIAlarmRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDIAlarmRuleResponse) GetBody() *CreateDIAlarmRuleResponseBody {
	return s.Body
}

func (s *CreateDIAlarmRuleResponse) SetHeaders(v map[string]*string) *CreateDIAlarmRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateDIAlarmRuleResponse) SetStatusCode(v int32) *CreateDIAlarmRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDIAlarmRuleResponse) SetBody(v *CreateDIAlarmRuleResponseBody) *CreateDIAlarmRuleResponse {
	s.Body = v
	return s
}

func (s *CreateDIAlarmRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
