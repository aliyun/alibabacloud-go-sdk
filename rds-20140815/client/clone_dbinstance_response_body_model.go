// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CloneDBInstanceResponseBody
	GetDBInstanceId() *string
	SetOrderId(v string) *CloneDBInstanceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CloneDBInstanceResponseBody
	GetRequestId() *string
}

type CloneDBInstanceResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 100789370****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1E43AAE0-BEE8-43DA-860D-EAF2AA0724DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloneDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloneDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CloneDBInstanceResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CloneDBInstanceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CloneDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloneDBInstanceResponseBody) SetDBInstanceId(v string) *CloneDBInstanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *CloneDBInstanceResponseBody) SetOrderId(v string) *CloneDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CloneDBInstanceResponseBody) SetRequestId(v string) *CloneDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloneDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
