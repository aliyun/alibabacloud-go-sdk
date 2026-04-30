// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAIDBClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateAIDBClusterResponseBody
	GetDBClusterId() *string
	SetOrderId(v string) *CreateAIDBClusterResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateAIDBClusterResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *CreateAIDBClusterResponseBody
	GetResourceGroupId() *string
}

type CreateAIDBClusterResponseBody struct {
	// example:
	//
	// pc-******************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 2035638*******
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CDB3258F-B5DE-43C4-8935-CBA0CA******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rg-************
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateAIDBClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAIDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAIDBClusterResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateAIDBClusterResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateAIDBClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAIDBClusterResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateAIDBClusterResponseBody) SetDBClusterId(v string) *CreateAIDBClusterResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *CreateAIDBClusterResponseBody) SetOrderId(v string) *CreateAIDBClusterResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateAIDBClusterResponseBody) SetRequestId(v string) *CreateAIDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAIDBClusterResponseBody) SetResourceGroupId(v string) *CreateAIDBClusterResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAIDBClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
