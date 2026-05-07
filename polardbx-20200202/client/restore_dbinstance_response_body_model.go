// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *RestoreDBInstanceResponseBody
	GetDBInstanceName() *string
	SetOrderId(v string) *RestoreDBInstanceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *RestoreDBInstanceResponseBody
	GetRequestId() *string
}

type RestoreDBInstanceResponseBody struct {
	// example:
	//
	// pxc-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// 12345
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 9B2F3840-XXXX-XXXX-XXXX-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestoreDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestoreDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreDBInstanceResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *RestoreDBInstanceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *RestoreDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestoreDBInstanceResponseBody) SetDBInstanceName(v string) *RestoreDBInstanceResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *RestoreDBInstanceResponseBody) SetOrderId(v string) *RestoreDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *RestoreDBInstanceResponseBody) SetRequestId(v string) *RestoreDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestoreDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
