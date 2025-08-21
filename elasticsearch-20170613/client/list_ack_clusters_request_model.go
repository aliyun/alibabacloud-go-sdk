// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAckClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPage(v int32) *ListAckClustersRequest
	GetPage() *int32
	SetSize(v int32) *ListAckClustersRequest
	GetSize() *int32
	SetVpcId(v string) *ListAckClustersRequest
	GetVpcId() *string
}

type ListAckClustersRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// 3
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the ACK clusters belong.
	//
	// example:
	//
	// vpc-bp12nu14urf0upaf4****
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListAckClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAckClustersRequest) GoString() string {
	return s.String()
}

func (s *ListAckClustersRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListAckClustersRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListAckClustersRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListAckClustersRequest) SetPage(v int32) *ListAckClustersRequest {
	s.Page = &v
	return s
}

func (s *ListAckClustersRequest) SetSize(v int32) *ListAckClustersRequest {
	s.Size = &v
	return s
}

func (s *ListAckClustersRequest) SetVpcId(v string) *ListAckClustersRequest {
	s.VpcId = &v
	return s
}

func (s *ListAckClustersRequest) Validate() error {
	return dara.Validate(s)
}
