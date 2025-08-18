// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *ListNodesResponseBody
	GetCount() *int32
	SetNodeModels(v []*ListNodesResponseBodyNodeModels) *ListNodesResponseBody
	GetNodeModels() []*ListNodesResponseBodyNodeModels
	SetPerPageSize(v int32) *ListNodesResponseBody
	GetPerPageSize() *int32
	SetRequestId(v string) *ListNodesResponseBody
	GetRequestId() *string
	SetToPage(v int32) *ListNodesResponseBody
	GetToPage() *int32
}

type ListNodesResponseBody struct {
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The resource nodes.
	NodeModels []*ListNodesResponseBodyNodeModels `json:"NodeModels,omitempty" xml:"NodeModels,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PerPageSize *int32 `json:"PerPageSize,omitempty" xml:"PerPageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	ToPage *int32 `json:"ToPage,omitempty" xml:"ToPage,omitempty"`
}

func (s ListNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListNodesResponseBody) GetNodeModels() []*ListNodesResponseBodyNodeModels {
	return s.NodeModels
}

func (s *ListNodesResponseBody) GetPerPageSize() *int32 {
	return s.PerPageSize
}

func (s *ListNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNodesResponseBody) GetToPage() *int32 {
	return s.ToPage
}

func (s *ListNodesResponseBody) SetCount(v int32) *ListNodesResponseBody {
	s.Count = &v
	return s
}

func (s *ListNodesResponseBody) SetNodeModels(v []*ListNodesResponseBodyNodeModels) *ListNodesResponseBody {
	s.NodeModels = v
	return s
}

func (s *ListNodesResponseBody) SetPerPageSize(v int32) *ListNodesResponseBody {
	s.PerPageSize = &v
	return s
}

func (s *ListNodesResponseBody) SetRequestId(v string) *ListNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodesResponseBody) SetToPage(v int32) *ListNodesResponseBody {
	s.ToPage = &v
	return s
}

func (s *ListNodesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListNodesResponseBodyNodeModels struct {
	// The billing method of the resource node.
	//
	// >  This parameter is returned only if the ChargeResourceMode parameter of the delivery group to which the resource node belongs is set to Node.
	//
	// Valid values:
	//
	// 	- PostPaid: pay-as-you-go
	//
	// 	- Prepaid: subscription
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The ID of the resource node.
	//
	// >  This parameter is returned only if the ChargeResourceMode parameter of the delivery group to which the resource node belongs is set to Node.
	//
	// example:
	//
	// i-bp13********
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ListNodesResponseBodyNodeModels) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyNodeModels) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyNodeModels) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListNodesResponseBodyNodeModels) GetNodeId() *string {
	return s.NodeId
}

func (s *ListNodesResponseBodyNodeModels) SetChargeType(v string) *ListNodesResponseBodyNodeModels {
	s.ChargeType = &v
	return s
}

func (s *ListNodesResponseBodyNodeModels) SetNodeId(v string) *ListNodesResponseBodyNodeModels {
	s.NodeId = &v
	return s
}

func (s *ListNodesResponseBodyNodeModels) Validate() error {
	return dara.Validate(s)
}
