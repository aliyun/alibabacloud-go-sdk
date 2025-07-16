// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreateDBInstanceResponseBody
	GetDBInstanceName() *string
	SetOrderId(v string) *CreateDBInstanceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateDBInstanceResponseBody
	GetRequestId() *string
}

type CreateDBInstanceResponseBody struct {
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
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateDBInstanceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBInstanceResponseBody) SetDBInstanceName(v string) *CreateDBInstanceResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetOrderId(v string) *CreateDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetRequestId(v string) *CreateDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
