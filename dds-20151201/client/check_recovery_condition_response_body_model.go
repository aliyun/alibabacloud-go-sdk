// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckRecoveryConditionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CheckRecoveryConditionResponseBody
	GetDBInstanceName() *string
	SetIsValid(v bool) *CheckRecoveryConditionResponseBody
	GetIsValid() *bool
	SetRequestId(v string) *CheckRecoveryConditionResponseBody
	GetRequestId() *string
}

type CheckRecoveryConditionResponseBody struct {
	// The instance ID
	//
	// example:
	//
	// dds-bp1378****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// Indicates whether the data of the instance can be restored. Valid values:
	//
	// 	- **true**: The data of the instance can be restored.
	//
	// 	- **false**: The data of the instance cannot be restored.
	//
	// example:
	//
	// true
	IsValid *bool `json:"IsValid,omitempty" xml:"IsValid,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D563A3E7-6010-45FE-A0CD-9283414C9657
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckRecoveryConditionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckRecoveryConditionResponseBody) GoString() string {
	return s.String()
}

func (s *CheckRecoveryConditionResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CheckRecoveryConditionResponseBody) GetIsValid() *bool {
	return s.IsValid
}

func (s *CheckRecoveryConditionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckRecoveryConditionResponseBody) SetDBInstanceName(v string) *CheckRecoveryConditionResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *CheckRecoveryConditionResponseBody) SetIsValid(v bool) *CheckRecoveryConditionResponseBody {
	s.IsValid = &v
	return s
}

func (s *CheckRecoveryConditionResponseBody) SetRequestId(v string) *CheckRecoveryConditionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckRecoveryConditionResponseBody) Validate() error {
	return dara.Validate(s)
}
