// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBInstanceDescriptionResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceDescriptionResponseBody struct {
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceDescriptionResponseBody) SetRequestId(v string) *ModifyDBInstanceDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
