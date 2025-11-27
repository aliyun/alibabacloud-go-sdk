// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceSecurityGroupRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceSecurityGroupRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceSecurityGroupRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceSecurityGroupRuleResponseBody) *ModifyDBInstanceSecurityGroupRuleResponse
	GetBody() *ModifyDBInstanceSecurityGroupRuleResponseBody
}

type ModifyDBInstanceSecurityGroupRuleResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceSecurityGroupRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceSecurityGroupRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceSecurityGroupRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSecurityGroupRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceSecurityGroupRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceSecurityGroupRuleResponse) GetBody() *ModifyDBInstanceSecurityGroupRuleResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceSecurityGroupRuleResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceSecurityGroupRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceSecurityGroupRuleResponse) SetStatusCode(v int32) *ModifyDBInstanceSecurityGroupRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceSecurityGroupRuleResponse) SetBody(v *ModifyDBInstanceSecurityGroupRuleResponseBody) *ModifyDBInstanceSecurityGroupRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceSecurityGroupRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
