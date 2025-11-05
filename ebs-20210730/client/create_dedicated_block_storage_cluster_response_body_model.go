// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDedicatedBlockStorageClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDbscId(v string) *CreateDedicatedBlockStorageClusterResponseBody
	GetDbscId() *string
	SetOrderId(v string) *CreateDedicatedBlockStorageClusterResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateDedicatedBlockStorageClusterResponseBody
	GetRequestId() *string
}

type CreateDedicatedBlockStorageClusterResponseBody struct {
	// The ID of the dedicated block storage cluster.
	//
	// example:
	//
	// dbsc-f8z4d3k4nsgg9okb****
	DbscId *string `json:"DbscId,omitempty" xml:"DbscId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 50155660025****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDedicatedBlockStorageClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDedicatedBlockStorageClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDedicatedBlockStorageClusterResponseBody) GetDbscId() *string {
	return s.DbscId
}

func (s *CreateDedicatedBlockStorageClusterResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateDedicatedBlockStorageClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDedicatedBlockStorageClusterResponseBody) SetDbscId(v string) *CreateDedicatedBlockStorageClusterResponseBody {
	s.DbscId = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterResponseBody) SetOrderId(v string) *CreateDedicatedBlockStorageClusterResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterResponseBody) SetRequestId(v string) *CreateDedicatedBlockStorageClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
