// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStrategyTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *ModifyStrategyTargetRequest
	GetConfig() *string
	SetSourceIp(v string) *ModifyStrategyTargetRequest
	GetSourceIp() *string
	SetTarget(v string) *ModifyStrategyTargetRequest
	GetTarget() *string
	SetType(v string) *ModifyStrategyTargetRequest
	GetType() *string
}

type ModifyStrategyTargetRequest struct {
	// The ID of the baseline check policy. The ID is returned after the policy is created. The value of this parameter is in the JSON format and contains the following field:
	//
	// 	- **strategyId**: the ID of the policy
	//
	// This parameter is required.
	//
	// example:
	//
	// {"strategyId":8070645}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The information about the asset group to which the policy is applied. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **TargetType**: the type of the asset to which the policy is applied. Set the value to **groupId**, which indicates that the policy is applied to an asset group.
	//
	// 	- **BindUuidCount**: the number of servers to which the policy is applied.
	//
	// 	- **Target**: the ID of the asset group.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"Target":"9273980","BindUuidCount":5320,"TargetType":"groupId","Flag":"del"},{"Target":"9677606","TargetType":"groupId","Flag":"del"},{"Target":"10121607","BindUuidCount":7,"TargetType":"groupId","Flag":"add"},{"Target":"10670708","BindUuidCount":2,"TargetType":"groupId","Flag":"del"},{"Target":"11246338","BindUuidCount":6,"TargetType":"groupId","Flag":"del"},{"Target":"11291161","BindUuidCount":13,"TargetType":"groupId","Flag":"del"}]
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The type of the configuration. Set the value to **hc_strategy**.
	//
	// This parameter is required.
	//
	// example:
	//
	// hc_strategy
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyStrategyTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyStrategyTargetRequest) GoString() string {
	return s.String()
}

func (s *ModifyStrategyTargetRequest) GetConfig() *string {
	return s.Config
}

func (s *ModifyStrategyTargetRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyStrategyTargetRequest) GetTarget() *string {
	return s.Target
}

func (s *ModifyStrategyTargetRequest) GetType() *string {
	return s.Type
}

func (s *ModifyStrategyTargetRequest) SetConfig(v string) *ModifyStrategyTargetRequest {
	s.Config = &v
	return s
}

func (s *ModifyStrategyTargetRequest) SetSourceIp(v string) *ModifyStrategyTargetRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyStrategyTargetRequest) SetTarget(v string) *ModifyStrategyTargetRequest {
	s.Target = &v
	return s
}

func (s *ModifyStrategyTargetRequest) SetType(v string) *ModifyStrategyTargetRequest {
	s.Type = &v
	return s
}

func (s *ModifyStrategyTargetRequest) Validate() error {
	return dara.Validate(s)
}
