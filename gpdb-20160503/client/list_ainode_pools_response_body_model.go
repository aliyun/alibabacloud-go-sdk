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
	AINodePoolInfos []*ListAINodePoolsResponseBodyAINodePoolInfos `json:"AINodePoolInfos,omitempty" xml:"AINodePoolInfos,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type ListAINodePoolsResponseBodyAINodePoolInfos struct {
	AINodeInfos []*ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos `json:"AINodeInfos,omitempty" xml:"AINodeInfos,omitempty" type:"Repeated"`
	// example:
	//
	// aipool-xxxxxxxxx
	AINodePoolId *string `json:"AINodePoolId,omitempty" xml:"AINodePoolId,omitempty"`
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
	return dara.Validate(s)
}

type ListAINodePoolsResponseBodyAINodePoolInfosAINodeInfos struct {
	// example:
	//
	// model_serving
	BindObject *string `json:"BindObject,omitempty" xml:"BindObject,omitempty"`
	BindStatus *string `json:"BindStatus,omitempty" xml:"BindStatus,omitempty"`
	// example:
	//
	// 2024-10-09T02:07:15Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// adbpg-ainode
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// ai-xxxxxxxxx
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// example:
	//
	// ADB.AIStandard.1
	NodeSpec *string `json:"NodeSpec,omitempty" xml:"NodeSpec,omitempty"`
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
