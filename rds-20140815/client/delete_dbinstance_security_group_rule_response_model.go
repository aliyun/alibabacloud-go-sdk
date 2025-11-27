// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstanceSecurityGroupRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDBInstanceSecurityGroupRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDBInstanceSecurityGroupRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDBInstanceSecurityGroupRuleResponseBody) *DeleteDBInstanceSecurityGroupRuleResponse
	GetBody() *DeleteDBInstanceSecurityGroupRuleResponseBody
}

type DeleteDBInstanceSecurityGroupRuleResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDBInstanceSecurityGroupRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDBInstanceSecurityGroupRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstanceSecurityGroupRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceSecurityGroupRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDBInstanceSecurityGroupRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDBInstanceSecurityGroupRuleResponse) GetBody() *DeleteDBInstanceSecurityGroupRuleResponseBody {
	return s.Body
}

func (s *DeleteDBInstanceSecurityGroupRuleResponse) SetHeaders(v map[string]*string) *DeleteDBInstanceSecurityGroupRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBInstanceSecurityGroupRuleResponse) SetStatusCode(v int32) *DeleteDBInstanceSecurityGroupRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBInstanceSecurityGroupRuleResponse) SetBody(v *DeleteDBInstanceSecurityGroupRuleResponseBody) *DeleteDBInstanceSecurityGroupRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteDBInstanceSecurityGroupRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
