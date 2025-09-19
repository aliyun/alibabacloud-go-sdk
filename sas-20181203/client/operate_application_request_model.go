// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContainerWebDefenseApplicationDTOS(v []*OperateApplicationRequestContainerWebDefenseApplicationDTOS) *OperateApplicationRequest
	GetContainerWebDefenseApplicationDTOS() []*OperateApplicationRequestContainerWebDefenseApplicationDTOS
	SetRuleId(v int64) *OperateApplicationRequest
	GetRuleId() *int64
}

type OperateApplicationRequest struct {
	// The container application that is protected from being tampered with.
	//
	// This parameter is required.
	ContainerWebDefenseApplicationDTOS []*OperateApplicationRequestContainerWebDefenseApplicationDTOS `json:"ContainerWebDefenseApplicationDTOS,omitempty" xml:"ContainerWebDefenseApplicationDTOS,omitempty" type:"Repeated"`
	// The ID of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 300566
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s OperateApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateApplicationRequest) GoString() string {
	return s.String()
}

func (s *OperateApplicationRequest) GetContainerWebDefenseApplicationDTOS() []*OperateApplicationRequestContainerWebDefenseApplicationDTOS {
	return s.ContainerWebDefenseApplicationDTOS
}

func (s *OperateApplicationRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *OperateApplicationRequest) SetContainerWebDefenseApplicationDTOS(v []*OperateApplicationRequestContainerWebDefenseApplicationDTOS) *OperateApplicationRequest {
	s.ContainerWebDefenseApplicationDTOS = v
	return s
}

func (s *OperateApplicationRequest) SetRuleId(v int64) *OperateApplicationRequest {
	s.RuleId = &v
	return s
}

func (s *OperateApplicationRequest) Validate() error {
	return dara.Validate(s)
}

type OperateApplicationRequestContainerWebDefenseApplicationDTOS struct {
	// The ID of the cluster to which the container belongs.
	//
	// >  You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ceb68cc58234141828677e383bd21ff0c
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to apply the configuration to the asset. Valid values:
	//
	// 	- **add**: applied
	//
	// 	- **del**: not applied
	//
	// This parameter is required.
	//
	// example:
	//
	// add
	Flag *string `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// The application ID. If the application is newly added, you do not need to specify this parameter.
	//
	// example:
	//
	// 196
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The value of the application tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// app:app-003
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s OperateApplicationRequestContainerWebDefenseApplicationDTOS) String() string {
	return dara.Prettify(s)
}

func (s OperateApplicationRequestContainerWebDefenseApplicationDTOS) GoString() string {
	return s.String()
}

func (s *OperateApplicationRequestContainerWebDefenseApplicationDTOS) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateApplicationRequestContainerWebDefenseApplicationDTOS) GetFlag() *string {
	return s.Flag
}

func (s *OperateApplicationRequestContainerWebDefenseApplicationDTOS) GetId() *int64 {
	return s.Id
}

func (s *OperateApplicationRequestContainerWebDefenseApplicationDTOS) GetTag() *string {
	return s.Tag
}

func (s *OperateApplicationRequestContainerWebDefenseApplicationDTOS) SetClusterId(v string) *OperateApplicationRequestContainerWebDefenseApplicationDTOS {
	s.ClusterId = &v
	return s
}

func (s *OperateApplicationRequestContainerWebDefenseApplicationDTOS) SetFlag(v string) *OperateApplicationRequestContainerWebDefenseApplicationDTOS {
	s.Flag = &v
	return s
}

func (s *OperateApplicationRequestContainerWebDefenseApplicationDTOS) SetId(v int64) *OperateApplicationRequestContainerWebDefenseApplicationDTOS {
	s.Id = &v
	return s
}

func (s *OperateApplicationRequestContainerWebDefenseApplicationDTOS) SetTag(v string) *OperateApplicationRequestContainerWebDefenseApplicationDTOS {
	s.Tag = &v
	return s
}

func (s *OperateApplicationRequestContainerWebDefenseApplicationDTOS) Validate() error {
	return dara.Validate(s)
}
