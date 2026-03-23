// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstanceSecurityGroupRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteDBInstanceSecurityGroupRuleResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteDBInstanceSecurityGroupRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDBInstanceSecurityGroupRuleResponseBody
	GetRequestId() *string
}

type DeleteDBInstanceSecurityGroupRuleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBInstanceSecurityGroupRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstanceSecurityGroupRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceSecurityGroupRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteDBInstanceSecurityGroupRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDBInstanceSecurityGroupRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDBInstanceSecurityGroupRuleResponseBody) SetCode(v string) *DeleteDBInstanceSecurityGroupRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDBInstanceSecurityGroupRuleResponseBody) SetMessage(v string) *DeleteDBInstanceSecurityGroupRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDBInstanceSecurityGroupRuleResponseBody) SetRequestId(v string) *DeleteDBInstanceSecurityGroupRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBInstanceSecurityGroupRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
