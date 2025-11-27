// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReadOnlyDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionString(v string) *CreateReadOnlyDBInstanceResponseBody
	GetConnectionString() *string
	SetDBInstanceId(v string) *CreateReadOnlyDBInstanceResponseBody
	GetDBInstanceId() *string
	SetOrderId(v string) *CreateReadOnlyDBInstanceResponseBody
	GetOrderId() *string
	SetPort(v string) *CreateReadOnlyDBInstanceResponseBody
	GetPort() *string
	SetRequestId(v string) *CreateReadOnlyDBInstanceResponseBody
	GetRequestId() *string
}

type CreateReadOnlyDBInstanceResponseBody struct {
	// The internal endpoint that is used to connect to the read-only instance.
	//
	// example:
	//
	// rr-****.mysql.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The ID of the read-only instance.
	//
	// example:
	//
	// rr-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 10078937****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The internal port number that is used to connect to the read-only instance.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1E43AAE0-BEE8-43DA-860D-EAF2AA0724DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateReadOnlyDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateReadOnlyDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateReadOnlyDBInstanceResponseBody) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *CreateReadOnlyDBInstanceResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateReadOnlyDBInstanceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateReadOnlyDBInstanceResponseBody) GetPort() *string {
	return s.Port
}

func (s *CreateReadOnlyDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateReadOnlyDBInstanceResponseBody) SetConnectionString(v string) *CreateReadOnlyDBInstanceResponseBody {
	s.ConnectionString = &v
	return s
}

func (s *CreateReadOnlyDBInstanceResponseBody) SetDBInstanceId(v string) *CreateReadOnlyDBInstanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *CreateReadOnlyDBInstanceResponseBody) SetOrderId(v string) *CreateReadOnlyDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateReadOnlyDBInstanceResponseBody) SetPort(v string) *CreateReadOnlyDBInstanceResponseBody {
	s.Port = &v
	return s
}

func (s *CreateReadOnlyDBInstanceResponseBody) SetRequestId(v string) *CreateReadOnlyDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateReadOnlyDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
