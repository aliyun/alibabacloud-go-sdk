// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCDiskSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *ModifyRCDiskSpecResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *ModifyRCDiskSpecResponseBody
	GetRequestId() *string
}

type ModifyRCDiskSpecResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 245053924720608
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRCDiskSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCDiskSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRCDiskSpecResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ModifyRCDiskSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRCDiskSpecResponseBody) SetOrderId(v int64) *ModifyRCDiskSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyRCDiskSpecResponseBody) SetRequestId(v string) *ModifyRCDiskSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRCDiskSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
