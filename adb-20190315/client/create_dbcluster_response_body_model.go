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
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition cluster.
	//
	// example:
	//
	// am-bp1r053byu48p****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 202353278****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the cluster belongs.
	//
	// example:
	//
	// rg-4690g37929****
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
