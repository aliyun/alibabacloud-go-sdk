// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBClusterResponseBody
	GetDBClusterId() *string
	SetOrderId(v string) *ModifyDBClusterResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyDBClusterResponseBody
	GetRequestId() *string
}

type ModifyDBClusterResponseBody struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// example:
	//
	// amv-bp1r053byu48p****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 2035629****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0D6BF3E2-41D8-57F6-9A62-A13A70377952
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyDBClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterResponseBody) SetDBClusterId(v string) *ModifyDBClusterResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterResponseBody) SetOrderId(v string) *ModifyDBClusterResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDBClusterResponseBody) SetRequestId(v string) *ModifyDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
