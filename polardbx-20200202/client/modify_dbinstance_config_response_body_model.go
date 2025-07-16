// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBInstanceConfigResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceConfigResponseBody struct {
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceConfigResponseBody) SetRequestId(v string) *ModifyDBInstanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
