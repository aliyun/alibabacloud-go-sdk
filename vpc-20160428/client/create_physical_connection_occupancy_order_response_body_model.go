// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePhysicalConnectionOccupancyOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreatePhysicalConnectionOccupancyOrderResponseBodyData) *CreatePhysicalConnectionOccupancyOrderResponseBody
	GetData() *CreatePhysicalConnectionOccupancyOrderResponseBodyData
	SetRequestId(v string) *CreatePhysicalConnectionOccupancyOrderResponseBody
	GetRequestId() *string
}

type CreatePhysicalConnectionOccupancyOrderResponseBody struct {
	// The details.
	Data *CreatePhysicalConnectionOccupancyOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9B9300FE-11E2-4E3B-949C-BED3B44DD26D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePhysicalConnectionOccupancyOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePhysicalConnectionOccupancyOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePhysicalConnectionOccupancyOrderResponseBody) GetData() *CreatePhysicalConnectionOccupancyOrderResponseBodyData {
	return s.Data
}

func (s *CreatePhysicalConnectionOccupancyOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePhysicalConnectionOccupancyOrderResponseBody) SetData(v *CreatePhysicalConnectionOccupancyOrderResponseBodyData) *CreatePhysicalConnectionOccupancyOrderResponseBody {
	s.Data = v
	return s
}

func (s *CreatePhysicalConnectionOccupancyOrderResponseBody) SetRequestId(v string) *CreatePhysicalConnectionOccupancyOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePhysicalConnectionOccupancyOrderResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreatePhysicalConnectionOccupancyOrderResponseBodyData struct {
	// The ID of the order that is placed.
	//
	// example:
	//
	// 50187055895****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreatePhysicalConnectionOccupancyOrderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreatePhysicalConnectionOccupancyOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePhysicalConnectionOccupancyOrderResponseBodyData) GetOrderId() *string {
	return s.OrderId
}

func (s *CreatePhysicalConnectionOccupancyOrderResponseBodyData) SetOrderId(v string) *CreatePhysicalConnectionOccupancyOrderResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreatePhysicalConnectionOccupancyOrderResponseBodyData) Validate() error {
	return dara.Validate(s)
}
