// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iComponentLayout interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationName(v string) *ComponentLayout
	GetApplicationName() *string
	SetComponentName(v string) *ComponentLayout
	GetComponentName() *string
	SetNodeSelector(v *ComponentLayoutNodeSelector) *ComponentLayout
	GetNodeSelector() *ComponentLayoutNodeSelector
}

type ComponentLayout struct {
	// 应用名称。
	//
	// example:
	//
	// HDFS
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// 组件名称。
	//
	// example:
	//
	// DataNode
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// 节点选择器。
	NodeSelector *ComponentLayoutNodeSelector `json:"NodeSelector,omitempty" xml:"NodeSelector,omitempty" type:"Struct"`
}

func (s ComponentLayout) String() string {
	return dara.Prettify(s)
}

func (s ComponentLayout) GoString() string {
	return s.String()
}

func (s *ComponentLayout) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ComponentLayout) GetComponentName() *string {
	return s.ComponentName
}

func (s *ComponentLayout) GetNodeSelector() *ComponentLayoutNodeSelector {
	return s.NodeSelector
}

func (s *ComponentLayout) SetApplicationName(v string) *ComponentLayout {
	s.ApplicationName = &v
	return s
}

func (s *ComponentLayout) SetComponentName(v string) *ComponentLayout {
	s.ComponentName = &v
	return s
}

func (s *ComponentLayout) SetNodeSelector(v *ComponentLayoutNodeSelector) *ComponentLayout {
	s.NodeSelector = v
	return s
}

func (s *ComponentLayout) Validate() error {
	if s.NodeSelector != nil {
		if err := s.NodeSelector.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ComponentLayoutNodeSelector struct {
	// 节点结束编号，包含结束编号。
	//
	// example:
	//
	// 3
	NodeEndIndex *int32 `json:"NodeEndIndex,omitempty" xml:"NodeEndIndex,omitempty"`
	// 节点组ID。
	//
	// example:
	//
	// G-F609686D45D4ABCD
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// 机器组下标编号。
	//
	// example:
	//
	// 1
	NodeGroupIndex *int32 `json:"NodeGroupIndex,omitempty" xml:"NodeGroupIndex,omitempty"`
	// 机器组名。
	//
	// example:
	//
	// master-1
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// SelectType = NODE_GROUP 且 nodeGroupId 不存在时使用
	//
	// example:
	//
	// [null]
	NodeGroupTypes []*string `json:"NodeGroupTypes,omitempty" xml:"NodeGroupTypes,omitempty" type:"Repeated"`
	// 节点名称列表。
	//
	// example:
	//
	// [null]
	NodeNames []*string `json:"NodeNames,omitempty" xml:"NodeNames,omitempty" type:"Repeated"`
	// 节点选择类型。
	//
	// This parameter is required.
	//
	// example:
	//
	// CLUSTER
	NodeSelectType *string `json:"NodeSelectType,omitempty" xml:"NodeSelectType,omitempty"`
	// 节点开始编号，包含开始编号。
	//
	// example:
	//
	// 1
	NodeStartIndex *int32 `json:"NodeStartIndex,omitempty" xml:"NodeStartIndex,omitempty"`
}

func (s ComponentLayoutNodeSelector) String() string {
	return dara.Prettify(s)
}

func (s ComponentLayoutNodeSelector) GoString() string {
	return s.String()
}

func (s *ComponentLayoutNodeSelector) GetNodeEndIndex() *int32 {
	return s.NodeEndIndex
}

func (s *ComponentLayoutNodeSelector) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ComponentLayoutNodeSelector) GetNodeGroupIndex() *int32 {
	return s.NodeGroupIndex
}

func (s *ComponentLayoutNodeSelector) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *ComponentLayoutNodeSelector) GetNodeGroupTypes() []*string {
	return s.NodeGroupTypes
}

func (s *ComponentLayoutNodeSelector) GetNodeNames() []*string {
	return s.NodeNames
}

func (s *ComponentLayoutNodeSelector) GetNodeSelectType() *string {
	return s.NodeSelectType
}

func (s *ComponentLayoutNodeSelector) GetNodeStartIndex() *int32 {
	return s.NodeStartIndex
}

func (s *ComponentLayoutNodeSelector) SetNodeEndIndex(v int32) *ComponentLayoutNodeSelector {
	s.NodeEndIndex = &v
	return s
}

func (s *ComponentLayoutNodeSelector) SetNodeGroupId(v string) *ComponentLayoutNodeSelector {
	s.NodeGroupId = &v
	return s
}

func (s *ComponentLayoutNodeSelector) SetNodeGroupIndex(v int32) *ComponentLayoutNodeSelector {
	s.NodeGroupIndex = &v
	return s
}

func (s *ComponentLayoutNodeSelector) SetNodeGroupName(v string) *ComponentLayoutNodeSelector {
	s.NodeGroupName = &v
	return s
}

func (s *ComponentLayoutNodeSelector) SetNodeGroupTypes(v []*string) *ComponentLayoutNodeSelector {
	s.NodeGroupTypes = v
	return s
}

func (s *ComponentLayoutNodeSelector) SetNodeNames(v []*string) *ComponentLayoutNodeSelector {
	s.NodeNames = v
	return s
}

func (s *ComponentLayoutNodeSelector) SetNodeSelectType(v string) *ComponentLayoutNodeSelector {
	s.NodeSelectType = &v
	return s
}

func (s *ComponentLayoutNodeSelector) SetNodeStartIndex(v int32) *ComponentLayoutNodeSelector {
	s.NodeStartIndex = &v
	return s
}

func (s *ComponentLayoutNodeSelector) Validate() error {
	return dara.Validate(s)
}
