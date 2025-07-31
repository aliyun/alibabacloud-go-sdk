// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCharacterType(v string) *ModifyParametersRequest
	GetCharacterType() *string
	SetDBInstanceId(v string) *ModifyParametersRequest
	GetDBInstanceId() *string
	SetNodeId(v string) *ModifyParametersRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *ModifyParametersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyParametersRequest
	GetOwnerId() *int64
	SetParameters(v string) *ModifyParametersRequest
	GetParameters() *string
	SetRegionId(v string) *ModifyParametersRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyParametersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyParametersRequest
	GetResourceOwnerId() *int64
	SetSwitchMode(v string) *ModifyParametersRequest
	GetSwitchMode() *string
}

type ModifyParametersRequest struct {
	// The role of the instance. Valid values:
	//
	// 	- **db**: a shard node.
	//
	// 	- **cs**: a Configserver node.
	//
	// 	- **mongos**: a mongos node.
	//
	// example:
	//
	// db
	CharacterType *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	// The instance ID.
	//
	// >  If you set this parameter to the ID of a sharded cluster instance, you must also specify the NodeId parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp19f409d75****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the mongos or shard node in the specified sharded cluster instance.
	//
	// >  This parameter is valid only when DBInstanceId is set to the ID of a sharded cluster instance.
	//
	// example:
	//
	// d-bp1b7bb3bbe****
	NodeId       *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The instance parameters that you want to modify and their values. Specify this parameter in a JSON string. Sample format: {"ParameterName1":"ParameterValue1","ParameterName2":"ParameterValue2"}.
	//
	// >  You can call the [DescribeParameterTemplates](https://help.aliyun.com/document_detail/67618.html) operation to query a list of default parameter templates.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"operationProfiling.mode":"all","operationProfiling.slowOpThresholdMs":"200"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SwitchMode           *string `json:"SwitchMode,omitempty" xml:"SwitchMode,omitempty"`
}

func (s ModifyParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyParametersRequest) GoString() string {
	return s.String()
}

func (s *ModifyParametersRequest) GetCharacterType() *string {
	return s.CharacterType
}

func (s *ModifyParametersRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyParametersRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *ModifyParametersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyParametersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyParametersRequest) GetParameters() *string {
	return s.Parameters
}

func (s *ModifyParametersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyParametersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyParametersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyParametersRequest) GetSwitchMode() *string {
	return s.SwitchMode
}

func (s *ModifyParametersRequest) SetCharacterType(v string) *ModifyParametersRequest {
	s.CharacterType = &v
	return s
}

func (s *ModifyParametersRequest) SetDBInstanceId(v string) *ModifyParametersRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyParametersRequest) SetNodeId(v string) *ModifyParametersRequest {
	s.NodeId = &v
	return s
}

func (s *ModifyParametersRequest) SetOwnerAccount(v string) *ModifyParametersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyParametersRequest) SetOwnerId(v int64) *ModifyParametersRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyParametersRequest) SetParameters(v string) *ModifyParametersRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyParametersRequest) SetRegionId(v string) *ModifyParametersRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyParametersRequest) SetResourceOwnerAccount(v string) *ModifyParametersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyParametersRequest) SetResourceOwnerId(v int64) *ModifyParametersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyParametersRequest) SetSwitchMode(v string) *ModifyParametersRequest {
	s.SwitchMode = &v
	return s
}

func (s *ModifyParametersRequest) Validate() error {
	return dara.Validate(s)
}
