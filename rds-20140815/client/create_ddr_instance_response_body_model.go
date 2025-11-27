// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDdrInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionString(v string) *CreateDdrInstanceResponseBody
	GetConnectionString() *string
	SetDBInstanceId(v string) *CreateDdrInstanceResponseBody
	GetDBInstanceId() *string
	SetOrderId(v string) *CreateDdrInstanceResponseBody
	GetOrderId() *string
	SetPort(v string) *CreateDdrInstanceResponseBody
	GetPort() *string
	SetRequestId(v string) *CreateDdrInstanceResponseBody
	GetRequestId() *string
}

type CreateDdrInstanceResponseBody struct {
	// The endpoint that is used to connect to the destination instance.
	//
	// >  The **DBInstanceNetType*	- parameter indicates whether the endpoint is internal or public.
	//
	// example:
	//
	// rm-xxxxx.mysql.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The destination instance ID.
	//
	// example:
	//
	// rm-xxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 2038691xxxxx
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The port number that is used to connect to the destination instance.
	//
	// > **DBInstanceNetType*	- indicates whether the port is internal or public.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E52666CC-330E-418A-8E5B-A19E3FB42D13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDdrInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDdrInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDdrInstanceResponseBody) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *CreateDdrInstanceResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDdrInstanceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateDdrInstanceResponseBody) GetPort() *string {
	return s.Port
}

func (s *CreateDdrInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDdrInstanceResponseBody) SetConnectionString(v string) *CreateDdrInstanceResponseBody {
	s.ConnectionString = &v
	return s
}

func (s *CreateDdrInstanceResponseBody) SetDBInstanceId(v string) *CreateDdrInstanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDdrInstanceResponseBody) SetOrderId(v string) *CreateDdrInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDdrInstanceResponseBody) SetPort(v string) *CreateDdrInstanceResponseBody {
	s.Port = &v
	return s
}

func (s *CreateDdrInstanceResponseBody) SetRequestId(v string) *CreateDdrInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDdrInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
