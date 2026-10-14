// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *CreateDiskResponseBody
	GetInstanceIds() []*string
	SetOrderId(v string) *CreateDiskResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateDiskResponseBody
	GetRequestId() *string
}

type CreateDiskResponseBody struct {
	// The collection of instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The order ID. Multiple IDs are separated by commas (,).
	//
	// example:
	//
	// 21127020370****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7030AB96-57CF-1C68-9FEE-D60E547FD79C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDiskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDiskResponseBody) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *CreateDiskResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateDiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDiskResponseBody) SetInstanceIds(v []*string) *CreateDiskResponseBody {
	s.InstanceIds = v
	return s
}

func (s *CreateDiskResponseBody) SetOrderId(v string) *CreateDiskResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDiskResponseBody) SetRequestId(v string) *CreateDiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDiskResponseBody) Validate() error {
	return dara.Validate(s)
}
