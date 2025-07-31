// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceDiskTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *ModifyDBInstanceDiskTypeResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyDBInstanceDiskTypeResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceDiskTypeResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 21012719476****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D8F1D721-6439-4257-A89C-F1E8E9C9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceDiskTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceDiskTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDiskTypeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyDBInstanceDiskTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceDiskTypeResponseBody) SetOrderId(v string) *ModifyDBInstanceDiskTypeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDBInstanceDiskTypeResponseBody) SetRequestId(v string) *ModifyDBInstanceDiskTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceDiskTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
