// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateDBClusterResponseBody
	GetDBClusterId() *string
	SetOrderId(v string) *CreateDBClusterResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateDBClusterResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *CreateDBClusterResponseBody
	GetResourceGroupId() *string
}

type CreateDBClusterResponseBody struct {
	// Cluster ID.
	//
	// example:
	//
	// pc-bp1s826a1up******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Order ID.
	//
	// example:
	//
	// 211454967******
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Request ID.
	//
	// example:
	//
	// E56531A4-E552-40BA-9C58-137B80******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-***************
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateDBClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBClusterResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateDBClusterResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateDBClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBClusterResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDBClusterResponseBody) SetDBClusterId(v string) *CreateDBClusterResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBClusterResponseBody) SetOrderId(v string) *CreateDBClusterResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDBClusterResponseBody) SetRequestId(v string) *CreateDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBClusterResponseBody) SetResourceGroupId(v string) *CreateDBClusterResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
