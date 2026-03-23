// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceSecurityGroupRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDBInstanceSecurityGroupRuleResponseBody
	GetCode() *string
	SetMessage(v string) *CreateDBInstanceSecurityGroupRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDBInstanceSecurityGroupRuleResponseBody
	GetRequestId() *string
}

type CreateDBInstanceSecurityGroupRuleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBInstanceSecurityGroupRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceSecurityGroupRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceSecurityGroupRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDBInstanceSecurityGroupRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDBInstanceSecurityGroupRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBInstanceSecurityGroupRuleResponseBody) SetCode(v string) *CreateDBInstanceSecurityGroupRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDBInstanceSecurityGroupRuleResponseBody) SetMessage(v string) *CreateDBInstanceSecurityGroupRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDBInstanceSecurityGroupRuleResponseBody) SetRequestId(v string) *CreateDBInstanceSecurityGroupRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBInstanceSecurityGroupRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
