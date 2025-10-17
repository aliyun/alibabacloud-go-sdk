// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInclinedNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeInclinedNodesResponseBodyItems) *DescribeInclinedNodesResponseBody
	GetItems() []*DescribeInclinedNodesResponseBodyItems
	SetRequestId(v string) *DescribeInclinedNodesResponseBody
	GetRequestId() *string
}

type DescribeInclinedNodesResponseBody struct {
	// The queried storage nodes.
	Items []*DescribeInclinedNodesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// C0BF6685-0519-543E-90F8-DB8949E4D5F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInclinedNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInclinedNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInclinedNodesResponseBody) GetItems() []*DescribeInclinedNodesResponseBodyItems {
	return s.Items
}

func (s *DescribeInclinedNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInclinedNodesResponseBody) SetItems(v []*DescribeInclinedNodesResponseBodyItems) *DescribeInclinedNodesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeInclinedNodesResponseBody) SetRequestId(v string) *DescribeInclinedNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInclinedNodesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInclinedNodesResponseBodyItems struct {
	// The disk usage of the storage node.
	//
	// example:
	//
	// 90.5
	DiskUsageRatio *string `json:"DiskUsageRatio,omitempty" xml:"DiskUsageRatio,omitempty"`
	// The number of the storage node.
	//
	// example:
	//
	// Node1
	Node *string `json:"Node,omitempty" xml:"Node,omitempty"`
}

func (s DescribeInclinedNodesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeInclinedNodesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeInclinedNodesResponseBodyItems) GetDiskUsageRatio() *string {
	return s.DiskUsageRatio
}

func (s *DescribeInclinedNodesResponseBodyItems) GetNode() *string {
	return s.Node
}

func (s *DescribeInclinedNodesResponseBodyItems) SetDiskUsageRatio(v string) *DescribeInclinedNodesResponseBodyItems {
	s.DiskUsageRatio = &v
	return s
}

func (s *DescribeInclinedNodesResponseBodyItems) SetNode(v string) *DescribeInclinedNodesResponseBodyItems {
	s.Node = &v
	return s
}

func (s *DescribeInclinedNodesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
