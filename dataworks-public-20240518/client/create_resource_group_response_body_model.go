// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateResourceGroupResponseBody
	GetRequestId() *string
	SetResourceGroupOrder(v *CreateResourceGroupResponseBodyResourceGroupOrder) *CreateResourceGroupResponseBody
	GetResourceGroupOrder() *CreateResourceGroupResponseBodyResourceGroupOrder
	SetSuccess(v bool) *CreateResourceGroupResponseBody
	GetSuccess() *bool
}

type CreateResourceGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the order that is used to create the serverless resource group.
	ResourceGroupOrder *CreateResourceGroupResponseBodyResourceGroupOrder `json:"ResourceGroupOrder,omitempty" xml:"ResourceGroupOrder,omitempty" type:"Struct"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateResourceGroupResponseBody) GetResourceGroupOrder() *CreateResourceGroupResponseBodyResourceGroupOrder {
	return s.ResourceGroupOrder
}

func (s *CreateResourceGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateResourceGroupResponseBody) SetRequestId(v string) *CreateResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceGroupResponseBody) SetResourceGroupOrder(v *CreateResourceGroupResponseBodyResourceGroupOrder) *CreateResourceGroupResponseBody {
	s.ResourceGroupOrder = v
	return s
}

func (s *CreateResourceGroupResponseBody) SetSuccess(v bool) *CreateResourceGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateResourceGroupResponseBodyResourceGroupOrder struct {
	// The ID of the serverless resource group.
	//
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the order that is used to create the serverless resource group.
	//
	// example:
	//
	// 2391982058XXXXX
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The instance ID of the order that is used to create the serverless resource group.
	//
	// example:
	//
	// c442b330-3b10-4584-959e-736e4edXXXXX
	OrderInstanceId *string `json:"OrderInstanceId,omitempty" xml:"OrderInstanceId,omitempty"`
}

func (s CreateResourceGroupResponseBodyResourceGroupOrder) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceGroupResponseBodyResourceGroupOrder) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupResponseBodyResourceGroupOrder) GetId() *string {
	return s.Id
}

func (s *CreateResourceGroupResponseBodyResourceGroupOrder) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateResourceGroupResponseBodyResourceGroupOrder) GetOrderInstanceId() *string {
	return s.OrderInstanceId
}

func (s *CreateResourceGroupResponseBodyResourceGroupOrder) SetId(v string) *CreateResourceGroupResponseBodyResourceGroupOrder {
	s.Id = &v
	return s
}

func (s *CreateResourceGroupResponseBodyResourceGroupOrder) SetOrderId(v int64) *CreateResourceGroupResponseBodyResourceGroupOrder {
	s.OrderId = &v
	return s
}

func (s *CreateResourceGroupResponseBodyResourceGroupOrder) SetOrderInstanceId(v string) *CreateResourceGroupResponseBodyResourceGroupOrder {
	s.OrderInstanceId = &v
	return s
}

func (s *CreateResourceGroupResponseBodyResourceGroupOrder) Validate() error {
	return dara.Validate(s)
}
