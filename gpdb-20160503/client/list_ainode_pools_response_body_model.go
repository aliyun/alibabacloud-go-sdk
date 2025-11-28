// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAINodePoolsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAINodePoolInfos(v []*ListAINodePoolsResponseBodyAINodePoolInfos) *ListAINodePoolsResponseBody
	GetAINodePoolInfos() []*ListAINodePoolsResponseBodyAINodePoolInfos
	SetRequestId(v string) *ListAINodePoolsResponseBody
	GetRequestId() *string
}

type ListAINodePoolsResponseBody struct {
	// Details of the AI node resource pool.
	AINodePoolInfos []*ListAINodePoolsResponseBodyAINodePoolInfos `json:"AINodePoolInfos,omitempty" xml:"AINodePoolInfos,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAINodePoolsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAINodePoolsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAINodePoolsResponseBody) GetAINodePoolInfos() []*ListAINodePoolsResponseBodyAINodePoolInfos {
	return s.AINodePoolInfos
}

func (s *ListAINodePoolsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAINodePoolsResponseBody) SetAINodePoolInfos(v []*ListAINodePoolsResponseBodyAINodePoolInfos) *ListAINodePoolsResponseBody {
	s.AINodePoolInfos = v
	return s
}

func (s *ListAINodePoolsResponseBody) SetRequestId(v string) *ListAINodePoolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAINodePoolsResponseBody) Validate() error {
	if s.AINodePoolInfos != nil {
		for _, item := range s.AINodePoolInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAINodePoolsResponseBodyAINodePoolInfos struct {
	// AI node details.
	AINodeInfos []*ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos `json:"AINodeInfos,omitempty" xml:"AINodeInfos,omitempty" type:"Repeated"`
	// The ID of the resource pool to which the AI node belongs.
	//
	// example:
	//
	// aipool-xxxxxxxxx
	AINodePoolId *string `json:"AINodePoolId,omitempty" xml:"AINodePoolId,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 2
	NodeNum *string `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
}

func (s ListAINodePoolsResponseBodyAINodePoolInfos) String() string {
	return dara.Prettify(s)
}

func (s ListAINodePoolsResponseBodyAINodePoolInfos) GoString() string {
	return s.String()
}

func (s *ListAINodePoolsResponseBodyAINodePoolInfos) GetAINodeInfos() []*ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos {
	return s.AINodeInfos
}

func (s *ListAINodePoolsResponseBodyAINodePoolInfos) GetAINodePoolId() *string {
	return s.AINodePoolId
}

func (s *ListAINodePoolsResponseBodyAINodePoolInfos) GetNodeNum() *string {
	return s.NodeNum
}

func (s *ListAINodePoolsResponseBodyAINodePoolInfos) SetAINodeInfos(v []*ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos) *ListAINodePoolsResponseBodyAINodePoolInfos {
	s.AINodeInfos = v
	return s
}

func (s *ListAINodePoolsResponseBodyAINodePoolInfos) SetAINodePoolId(v string) *ListAINodePoolsResponseBodyAINodePoolInfos {
	s.AINodePoolId = &v
	return s
}

func (s *ListAINodePoolsResponseBodyAINodePoolInfos) SetNodeNum(v string) *ListAINodePoolsResponseBodyAINodePoolInfos {
	s.NodeNum = &v
	return s
}

func (s *ListAINodePoolsResponseBodyAINodePoolInfos) Validate() error {
	if s.AINodeInfos != nil {
		for _, item := range s.AINodeInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos struct {
	// The binding type of the AI node.
	//
	// example:
	//
	// model_serving
	BindObject *string `json:"BindObject,omitempty" xml:"BindObject,omitempty"`
	// The status of the AI node.
	//
	// 	- unbound: The node is not bound.
	//
	// 	- bound: The node is bound.
	BindStatus *string `json:"BindStatus,omitempty" xml:"BindStatus,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-10-09T02:07:15Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The namespace.
	//
	// example:
	//
	// adbpg-ainode
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The name of the AI node.
	//
	// example:
	//
	// ai-xxxxxxxxx
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The AI node specifications. The list of supported specifications is shown below.
	//
	//     ADB.AIMedium.1
	//
	//     ADB.AILarge.1
	//
	//     ADB.AIStandard.2
	//
	//     ADB.AIMedium.2
	//
	//     ADB.AILarge.2
	//
	//     ADB.AIXLarge.2
	//
	//     ADB.AIStandard.6
	//
	//     ADB.AIMedium.6
	//
	//     ADB.AILarge.6
	//
	//     ADB.AIXLarge.6
	//
	//     ADB.AIStandard.3
	//
	//     ADB.AIMedium.3
	//
	//     ADB.AILarge.3
	//
	//     ADB.AIXLarge.3
	//
	//     ADB.AIStandard.4
	//
	//     ADB.AIMedium.4
	//
	//     ADB.AILarge.4
	//
	//     ADB.AIXLarge.4
	//
	//     ADB.AIStandard.5
	//
	//     ADB.AIMedium.5
	//
	//     ADB.AILarge.5
	//
	//     ADB.AIXLarge.5
	//
	//     ADB.AIStandard.8
	//
	//     ADB.AIMedium.8
	//
	//     ADB.AILarge.8
	//
	//     ADB.AIXLarge.8
	//
	//     ADB.AI2XLarge.8
	//
	// example:
	//
	// ADB.AIStandard.1
	NodeSpec *string `json:"NodeSpec,omitempty" xml:"NodeSpec,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2025-06-16T02:04:42Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos) String() string {
	return dara.Prettify(s)
}

func (s ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos) GoString() string {
	return s.String()
}

func (s *ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos) GetBindObject() *string {
	return s.BindObject
}

func (s *ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos) GetBindStatus() *string {
	return s.BindStatus
}

func (s *ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos) GetNamespace() *string {
	return s.Namespace
}

func (s *ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos) GetNodeName() *string {
	return s.NodeName
}

func (s *ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos) GetNodeSpec() *string {
	return s.NodeSpec
}

func (s *ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos) SetBindObject(v string) *ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos {
	s.BindObject = &v
	return s
}

func (s *ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos) SetBindStatus(v string) *ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos {
	s.BindStatus = &v
	return s
}

func (s *ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos) SetCreateTime(v string) *ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos {
	s.CreateTime = &v
	return s
}

func (s *ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos) SetNamespace(v string) *ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos {
	s.Namespace = &v
	return s
}

func (s *ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos) SetNodeName(v string) *ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos {
	s.NodeName = &v
	return s
}

func (s *ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos) SetNodeSpec(v string) *ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos {
	s.NodeSpec = &v
	return s
}

func (s *ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos) SetUpdateTime(v string) *ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos {
	s.UpdateTime = &v
	return s
}

func (s *ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos) Validate() error {
	return dara.Validate(s)
}
