// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceSecurityGroupRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyDBInstanceSecurityGroupRuleResponseBody
	GetCode() *string
	SetMessage(v string) *ModifyDBInstanceSecurityGroupRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyDBInstanceSecurityGroupRuleResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceSecurityGroupRuleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceSecurityGroupRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceSecurityGroupRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSecurityGroupRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyDBInstanceSecurityGroupRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyDBInstanceSecurityGroupRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceSecurityGroupRuleResponseBody) SetCode(v string) *ModifyDBInstanceSecurityGroupRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyDBInstanceSecurityGroupRuleResponseBody) SetMessage(v string) *ModifyDBInstanceSecurityGroupRuleResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyDBInstanceSecurityGroupRuleResponseBody) SetRequestId(v string) *ModifyDBInstanceSecurityGroupRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceSecurityGroupRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
