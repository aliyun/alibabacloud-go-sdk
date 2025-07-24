// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplicationConfigParam interface {
	dara.Model
	String() string
	GoString() string
	SetConfigAction(v string) *ApplicationConfigParam
	GetConfigAction() *string
	SetConfigFileName(v string) *ApplicationConfigParam
	GetConfigFileName() *string
	SetConfigItemDescription(v string) *ApplicationConfigParam
	GetConfigItemDescription() *string
	SetConfigItemKey(v string) *ApplicationConfigParam
	GetConfigItemKey() *string
	SetConfigItemValue(v string) *ApplicationConfigParam
	GetConfigItemValue() *string
	SetConfigScope(v string) *ApplicationConfigParam
	GetConfigScope() *string
	SetEffectiveActions(v string) *ApplicationConfigParam
	GetEffectiveActions() *string
	SetEffectiveType(v string) *ApplicationConfigParam
	GetEffectiveType() *string
	SetNodeGroupId(v string) *ApplicationConfigParam
	GetNodeGroupId() *string
	SetNodeId(v string) *ApplicationConfigParam
	GetNodeId() *string
}

type ApplicationConfigParam struct {
	// example:
	//
	// DELETE
	ConfigAction *string `json:"ConfigAction,omitempty" xml:"ConfigAction,omitempty"`
	// example:
	//
	// hdfs-site.xml
	ConfigFileName *string `json:"ConfigFileName,omitempty" xml:"ConfigFileName,omitempty"`
	// example:
	//
	// namenode checkpoint period
	ConfigItemDescription *string `json:"ConfigItemDescription,omitempty" xml:"ConfigItemDescription,omitempty"`
	// example:
	//
	// dfs.namenode.checkpoint.period
	ConfigItemKey *string `json:"ConfigItemKey,omitempty" xml:"ConfigItemKey,omitempty"`
	// example:
	//
	// 3600s
	ConfigItemValue *string `json:"ConfigItemValue,omitempty" xml:"ConfigItemValue,omitempty"`
	ConfigScope     *string `json:"ConfigScope,omitempty" xml:"ConfigScope,omitempty"`
	// example:
	//
	// RESTART
	EffectiveActions *string `json:"EffectiveActions,omitempty" xml:"EffectiveActions,omitempty"`
	// example:
	//
	// MANUAL
	EffectiveType *string `json:"EffectiveType,omitempty" xml:"EffectiveType,omitempty"`
	// example:
	//
	// G-DE1CF4661E09****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// example:
	//
	// i-bp10h9rezawz1i4o****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ApplicationConfigParam) String() string {
	return dara.Prettify(s)
}

func (s ApplicationConfigParam) GoString() string {
	return s.String()
}

func (s *ApplicationConfigParam) GetConfigAction() *string {
	return s.ConfigAction
}

func (s *ApplicationConfigParam) GetConfigFileName() *string {
	return s.ConfigFileName
}

func (s *ApplicationConfigParam) GetConfigItemDescription() *string {
	return s.ConfigItemDescription
}

func (s *ApplicationConfigParam) GetConfigItemKey() *string {
	return s.ConfigItemKey
}

func (s *ApplicationConfigParam) GetConfigItemValue() *string {
	return s.ConfigItemValue
}

func (s *ApplicationConfigParam) GetConfigScope() *string {
	return s.ConfigScope
}

func (s *ApplicationConfigParam) GetEffectiveActions() *string {
	return s.EffectiveActions
}

func (s *ApplicationConfigParam) GetEffectiveType() *string {
	return s.EffectiveType
}

func (s *ApplicationConfigParam) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ApplicationConfigParam) GetNodeId() *string {
	return s.NodeId
}

func (s *ApplicationConfigParam) SetConfigAction(v string) *ApplicationConfigParam {
	s.ConfigAction = &v
	return s
}

func (s *ApplicationConfigParam) SetConfigFileName(v string) *ApplicationConfigParam {
	s.ConfigFileName = &v
	return s
}

func (s *ApplicationConfigParam) SetConfigItemDescription(v string) *ApplicationConfigParam {
	s.ConfigItemDescription = &v
	return s
}

func (s *ApplicationConfigParam) SetConfigItemKey(v string) *ApplicationConfigParam {
	s.ConfigItemKey = &v
	return s
}

func (s *ApplicationConfigParam) SetConfigItemValue(v string) *ApplicationConfigParam {
	s.ConfigItemValue = &v
	return s
}

func (s *ApplicationConfigParam) SetConfigScope(v string) *ApplicationConfigParam {
	s.ConfigScope = &v
	return s
}

func (s *ApplicationConfigParam) SetEffectiveActions(v string) *ApplicationConfigParam {
	s.EffectiveActions = &v
	return s
}

func (s *ApplicationConfigParam) SetEffectiveType(v string) *ApplicationConfigParam {
	s.EffectiveType = &v
	return s
}

func (s *ApplicationConfigParam) SetNodeGroupId(v string) *ApplicationConfigParam {
	s.NodeGroupId = &v
	return s
}

func (s *ApplicationConfigParam) SetNodeId(v string) *ApplicationConfigParam {
	s.NodeId = &v
	return s
}

func (s *ApplicationConfigParam) Validate() error {
	return dara.Validate(s)
}
