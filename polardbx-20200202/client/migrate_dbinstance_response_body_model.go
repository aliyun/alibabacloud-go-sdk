// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *MigrateDBInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *MigrateDBInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MigrateDBInstanceResponseBody
	GetSuccess() *bool
}

type MigrateDBInstanceResponseBody struct {
	// example:
	//
	// *****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MigrateDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MigrateDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateDBInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MigrateDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MigrateDBInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MigrateDBInstanceResponseBody) SetMessage(v string) *MigrateDBInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *MigrateDBInstanceResponseBody) SetRequestId(v string) *MigrateDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *MigrateDBInstanceResponseBody) SetSuccess(v bool) *MigrateDBInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *MigrateDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
