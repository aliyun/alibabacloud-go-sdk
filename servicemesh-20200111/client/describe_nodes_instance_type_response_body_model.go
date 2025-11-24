// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNodesInstanceTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceTypes(v []*DescribeNodesInstanceTypeResponseBodyInstanceTypes) *DescribeNodesInstanceTypeResponseBody
	GetInstanceTypes() []*DescribeNodesInstanceTypeResponseBodyInstanceTypes
	SetRequestId(v string) *DescribeNodesInstanceTypeResponseBody
	GetRequestId() *string
}

type DescribeNodesInstanceTypeResponseBody struct {
	// The instance types of the nodes.
	InstanceTypes []*DescribeNodesInstanceTypeResponseBodyInstanceTypes `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// BD65C0AD-D3C6-48D3-8D93-38D2015C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNodesInstanceTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodesInstanceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNodesInstanceTypeResponseBody) GetInstanceTypes() []*DescribeNodesInstanceTypeResponseBodyInstanceTypes {
	return s.InstanceTypes
}

func (s *DescribeNodesInstanceTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNodesInstanceTypeResponseBody) SetInstanceTypes(v []*DescribeNodesInstanceTypeResponseBodyInstanceTypes) *DescribeNodesInstanceTypeResponseBody {
	s.InstanceTypes = v
	return s
}

func (s *DescribeNodesInstanceTypeResponseBody) SetRequestId(v string) *DescribeNodesInstanceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNodesInstanceTypeResponseBody) Validate() error {
	if s.InstanceTypes != nil {
		for _, item := range s.InstanceTypes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNodesInstanceTypeResponseBodyInstanceTypes struct {
	// The label keys of nodes which have MultiBuffer optimization supported. (Enable optimization only if pod being scheduled to node which have a label key specified by this field and its value equals with the value field)
	//
	// example:
	//
	// feature.node.kubernetes.io/mb-feature-enable
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Indicates whether the instance type supports Multi-Buffer acceleration. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	MultiBufferEnabled *bool `json:"MultiBufferEnabled,omitempty" xml:"MultiBufferEnabled,omitempty"`
	// The instance type of the node.
	//
	// example:
	//
	// ecs.g7.xlarge
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The label keys of nodes which have MultiBuffer optimization supported. (Enable optimization only if pod being scheduled to node which have a label key specified by the key field and its value equals with this field)
	//
	// example:
	//
	// true
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeNodesInstanceTypeResponseBodyInstanceTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodesInstanceTypeResponseBodyInstanceTypes) GoString() string {
	return s.String()
}

func (s *DescribeNodesInstanceTypeResponseBodyInstanceTypes) GetKey() *string {
	return s.Key
}

func (s *DescribeNodesInstanceTypeResponseBodyInstanceTypes) GetMultiBufferEnabled() *bool {
	return s.MultiBufferEnabled
}

func (s *DescribeNodesInstanceTypeResponseBodyInstanceTypes) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeNodesInstanceTypeResponseBodyInstanceTypes) GetValue() *string {
	return s.Value
}

func (s *DescribeNodesInstanceTypeResponseBodyInstanceTypes) SetKey(v string) *DescribeNodesInstanceTypeResponseBodyInstanceTypes {
	s.Key = &v
	return s
}

func (s *DescribeNodesInstanceTypeResponseBodyInstanceTypes) SetMultiBufferEnabled(v bool) *DescribeNodesInstanceTypeResponseBodyInstanceTypes {
	s.MultiBufferEnabled = &v
	return s
}

func (s *DescribeNodesInstanceTypeResponseBodyInstanceTypes) SetNodeType(v string) *DescribeNodesInstanceTypeResponseBodyInstanceTypes {
	s.NodeType = &v
	return s
}

func (s *DescribeNodesInstanceTypeResponseBodyInstanceTypes) SetValue(v string) *DescribeNodesInstanceTypeResponseBodyInstanceTypes {
	s.Value = &v
	return s
}

func (s *DescribeNodesInstanceTypeResponseBodyInstanceTypes) Validate() error {
	return dara.Validate(s)
}
