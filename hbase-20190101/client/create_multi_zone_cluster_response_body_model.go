// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultiZoneClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateMultiZoneClusterResponseBody
	GetClusterId() *string
	SetOrderId(v string) *CreateMultiZoneClusterResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateMultiZoneClusterResponseBody
	GetRequestId() *string
}

type CreateMultiZoneClusterResponseBody struct {
	// example:
	//
	// ld-t4nn71xa0yn56****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 23232453****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 7F68E8F5-0377-4CF8-8B1D-FFFD6F5804D5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMultiZoneClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiZoneClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMultiZoneClusterResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateMultiZoneClusterResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateMultiZoneClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMultiZoneClusterResponseBody) SetClusterId(v string) *CreateMultiZoneClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateMultiZoneClusterResponseBody) SetOrderId(v string) *CreateMultiZoneClusterResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateMultiZoneClusterResponseBody) SetRequestId(v string) *CreateMultiZoneClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMultiZoneClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
