// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionString(v string) *CreateDBInstanceResponseBody
	GetConnectionString() *string
	SetDBInstanceId(v string) *CreateDBInstanceResponseBody
	GetDBInstanceId() *string
	SetOrderId(v string) *CreateDBInstanceResponseBody
	GetOrderId() *string
	SetPort(v string) *CreateDBInstanceResponseBody
	GetPort() *string
	SetRequestId(v string) *CreateDBInstanceResponseBody
	GetRequestId() *string
}

type CreateDBInstanceResponseBody struct {
	// This parameter is deprecated and will not return a value.
	//
	// You can use the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/86910.html) interface to view the connection address of the instance.
	//
	// example:
	//
	// gp-bp12ga6v69h86****.gpdb.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Order ID.
	//
	// example:
	//
	// 111111111111
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// This parameter is deprecated and will not return a value.
	//
	// You can use the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/86910.html) interface to view the port number of the instance.
	//
	// example:
	//
	// 3432
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 5414A4E5-4C36-4461-95FC-************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceResponseBody) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *CreateDBInstanceResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDBInstanceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateDBInstanceResponseBody) GetPort() *string {
	return s.Port
}

func (s *CreateDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBInstanceResponseBody) SetConnectionString(v string) *CreateDBInstanceResponseBody {
	s.ConnectionString = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetDBInstanceId(v string) *CreateDBInstanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetOrderId(v string) *CreateDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetPort(v string) *CreateDBInstanceResponseBody {
	s.Port = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetRequestId(v string) *CreateDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
