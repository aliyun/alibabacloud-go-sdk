// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTempModifyDBNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *TempModifyDBNodeResponseBody
	GetDBClusterId() *string
	SetDBNodeIds(v []*string) *TempModifyDBNodeResponseBody
	GetDBNodeIds() []*string
	SetOrderId(v string) *TempModifyDBNodeResponseBody
	GetOrderId() *string
	SetRequestId(v string) *TempModifyDBNodeResponseBody
	GetRequestId() *string
}

type TempModifyDBNodeResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// pc-xxxxxxxxxxxxxxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The details of the nodes.
	DBNodeIds []*string `json:"DBNodeIds,omitempty" xml:"DBNodeIds,omitempty" type:"Repeated"`
	// The ID of the order.
	//
	// example:
	//
	// 2035638*******
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 69A85BAF-1089-4CDF-A82F-0A140F******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TempModifyDBNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TempModifyDBNodeResponseBody) GoString() string {
	return s.String()
}

func (s *TempModifyDBNodeResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *TempModifyDBNodeResponseBody) GetDBNodeIds() []*string {
	return s.DBNodeIds
}

func (s *TempModifyDBNodeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *TempModifyDBNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TempModifyDBNodeResponseBody) SetDBClusterId(v string) *TempModifyDBNodeResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *TempModifyDBNodeResponseBody) SetDBNodeIds(v []*string) *TempModifyDBNodeResponseBody {
	s.DBNodeIds = v
	return s
}

func (s *TempModifyDBNodeResponseBody) SetOrderId(v string) *TempModifyDBNodeResponseBody {
	s.OrderId = &v
	return s
}

func (s *TempModifyDBNodeResponseBody) SetRequestId(v string) *TempModifyDBNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *TempModifyDBNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
