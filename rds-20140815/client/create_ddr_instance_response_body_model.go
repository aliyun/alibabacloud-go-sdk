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
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	DBInstanceId     *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OrderId          *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Port             *string `json:"Port,omitempty" xml:"Port,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
