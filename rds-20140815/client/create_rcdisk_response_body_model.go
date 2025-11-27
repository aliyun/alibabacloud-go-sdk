// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRCDiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiskId(v string) *CreateRCDiskResponseBody
	GetDiskId() *string
	SetOrderId(v string) *CreateRCDiskResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateRCDiskResponseBody
	GetRequestId() *string
}

type CreateRCDiskResponseBody struct {
	// The cloud disk ID.
	//
	// example:
	//
	// rcd-2zegrjtnkp6dqbe1egca
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 100789370230206
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F2911788-25E8-42E5-A3A3-1B38D263F01E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRCDiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRCDiskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRCDiskResponseBody) GetDiskId() *string {
	return s.DiskId
}

func (s *CreateRCDiskResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateRCDiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRCDiskResponseBody) SetDiskId(v string) *CreateRCDiskResponseBody {
	s.DiskId = &v
	return s
}

func (s *CreateRCDiskResponseBody) SetOrderId(v string) *CreateRCDiskResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateRCDiskResponseBody) SetRequestId(v string) *CreateRCDiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRCDiskResponseBody) Validate() error {
	return dara.Validate(s)
}
