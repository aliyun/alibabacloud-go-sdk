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
	// The cluster ID.
	//
	// example:
	//
	// am-bp1u8c0mgfg58****
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
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
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
