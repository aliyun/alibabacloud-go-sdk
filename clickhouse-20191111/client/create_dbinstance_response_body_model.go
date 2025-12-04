// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateDBInstanceResponseBody
	GetDBClusterId() *string
	SetOrderId(v string) *CreateDBInstanceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateDBInstanceResponseBody
	GetRequestId() *string
}

type CreateDBInstanceResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// cc-bp1qx68m06981****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 21137950671****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateDBInstanceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBInstanceResponseBody) SetDBClusterId(v string) *CreateDBInstanceResponseBody {
	s.DBClusterId = &v
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
