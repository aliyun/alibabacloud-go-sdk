// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeploymentSetConstraints interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultValue(v string) *DeploymentSetConstraints
	GetDefaultValue() *string
	SetEnableState(v string) *DeploymentSetConstraints
	GetEnableState() *string
	SetReplacementStrategy(v *ReplacementStrategy) *DeploymentSetConstraints
	GetReplacementStrategy() *ReplacementStrategy
	SetValues(v []*string) *DeploymentSetConstraints
	GetValues() []*string
}

type DeploymentSetConstraints struct {
	// 默认值。
	//
	// example:
	//
	// CLUSTER
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// example:
	//
	// 是否启用部署集限制策略
	EnableState *string `json:"EnableState,omitempty" xml:"EnableState,omitempty"`
	// 替换策略。
	ReplacementStrategy *ReplacementStrategy `json:"ReplacementStrategy,omitempty" xml:"ReplacementStrategy,omitempty"`
	// 枚举值。
	//
	// example:
	//
	// ["CLUSTER","NODE_GROUP","NONE"]
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DeploymentSetConstraints) String() string {
	return dara.Prettify(s)
}

func (s DeploymentSetConstraints) GoString() string {
	return s.String()
}

func (s *DeploymentSetConstraints) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *DeploymentSetConstraints) GetEnableState() *string {
	return s.EnableState
}

func (s *DeploymentSetConstraints) GetReplacementStrategy() *ReplacementStrategy {
	return s.ReplacementStrategy
}

func (s *DeploymentSetConstraints) GetValues() []*string {
	return s.Values
}

func (s *DeploymentSetConstraints) SetDefaultValue(v string) *DeploymentSetConstraints {
	s.DefaultValue = &v
	return s
}

func (s *DeploymentSetConstraints) SetEnableState(v string) *DeploymentSetConstraints {
	s.EnableState = &v
	return s
}

func (s *DeploymentSetConstraints) SetReplacementStrategy(v *ReplacementStrategy) *DeploymentSetConstraints {
	s.ReplacementStrategy = v
	return s
}

func (s *DeploymentSetConstraints) SetValues(v []*string) *DeploymentSetConstraints {
	s.Values = v
	return s
}

func (s *DeploymentSetConstraints) Validate() error {
	return dara.Validate(s)
}
