// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceSecurityGroupRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDBInstanceSecurityGroupRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDBInstanceSecurityGroupRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateDBInstanceSecurityGroupRuleResponseBody) *CreateDBInstanceSecurityGroupRuleResponse
	GetBody() *CreateDBInstanceSecurityGroupRuleResponseBody
}

type CreateDBInstanceSecurityGroupRuleResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDBInstanceSecurityGroupRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDBInstanceSecurityGroupRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceSecurityGroupRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceSecurityGroupRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDBInstanceSecurityGroupRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDBInstanceSecurityGroupRuleResponse) GetBody() *CreateDBInstanceSecurityGroupRuleResponseBody {
	return s.Body
}

func (s *CreateDBInstanceSecurityGroupRuleResponse) SetHeaders(v map[string]*string) *CreateDBInstanceSecurityGroupRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateDBInstanceSecurityGroupRuleResponse) SetStatusCode(v int32) *CreateDBInstanceSecurityGroupRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBInstanceSecurityGroupRuleResponse) SetBody(v *CreateDBInstanceSecurityGroupRuleResponseBody) *CreateDBInstanceSecurityGroupRuleResponse {
	s.Body = v
	return s
}

func (s *CreateDBInstanceSecurityGroupRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
