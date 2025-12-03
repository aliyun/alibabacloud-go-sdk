// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMultiZoneClusterNodeTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *ModifyMultiZoneClusterNodeTypeResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyMultiZoneClusterNodeTypeResponseBody
	GetRequestId() *string
}

type ModifyMultiZoneClusterNodeTypeResponseBody struct {
	// example:
	//
	// 12341234123****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 06CF7A6F-A81C-431D-BACD-793F24A67C54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyMultiZoneClusterNodeTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMultiZoneClusterNodeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMultiZoneClusterNodeTypeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyMultiZoneClusterNodeTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMultiZoneClusterNodeTypeResponseBody) SetOrderId(v string) *ModifyMultiZoneClusterNodeTypeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyMultiZoneClusterNodeTypeResponseBody) SetRequestId(v string) *ModifyMultiZoneClusterNodeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMultiZoneClusterNodeTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
