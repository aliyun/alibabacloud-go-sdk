// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterArchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBClusterArchResponseBody
	GetDBClusterId() *string
	SetOrderId(v string) *ModifyDBClusterArchResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyDBClusterArchResponseBody
	GetRequestId() *string
}

type ModifyDBClusterArchResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 2148126708*****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6A2EE5B4-CC9F-46E1-A747-E43BC9******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterArchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterArchResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterArchResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterArchResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyDBClusterArchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterArchResponseBody) SetDBClusterId(v string) *ModifyDBClusterArchResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterArchResponseBody) SetOrderId(v string) *ModifyDBClusterArchResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDBClusterArchResponseBody) SetRequestId(v string) *ModifyDBClusterArchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterArchResponseBody) Validate() error {
	return dara.Validate(s)
}
