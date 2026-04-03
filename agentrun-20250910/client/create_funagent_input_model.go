// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFunagentInput interface {
	dara.Model
	String() string
	GoString() string
	SetAdminName(v string) *CreateFunagentInput
	GetAdminName() *string
	SetAdminSecret(v string) *CreateFunagentInput
	GetAdminSecret() *string
	SetCpu(v float32) *CreateFunagentInput
	GetCpu() *float32
	SetDbConnections(v int32) *CreateFunagentInput
	GetDbConnections() *int32
	SetDbHost(v string) *CreateFunagentInput
	GetDbHost() *string
	SetDbInstanceId(v string) *CreateFunagentInput
	GetDbInstanceId() *string
	SetDbName(v string) *CreateFunagentInput
	GetDbName() *string
	SetDbPassword(v string) *CreateFunagentInput
	GetDbPassword() *string
	SetDbPort(v int32) *CreateFunagentInput
	GetDbPort() *int32
	SetDbType(v string) *CreateFunagentInput
	GetDbType() *string
	SetDbUsername(v string) *CreateFunagentInput
	GetDbUsername() *string
	SetDescription(v string) *CreateFunagentInput
	GetDescription() *string
	SetFunAgentName(v string) *CreateFunagentInput
	GetFunAgentName() *string
	SetMemory(v int32) *CreateFunagentInput
	GetMemory() *int32
	SetRegionId(v string) *CreateFunagentInput
	GetRegionId() *string
	SetReplicas(v int32) *CreateFunagentInput
	GetReplicas() *int32
	SetSecurityGroupId(v string) *CreateFunagentInput
	GetSecurityGroupId() *string
	SetVpcId(v string) *CreateFunagentInput
	GetVpcId() *string
	SetVswitchIds(v string) *CreateFunagentInput
	GetVswitchIds() *string
}

type CreateFunagentInput struct {
	// `string`，必填
	//
	// This parameter is required.
	AdminName *string `json:"adminName,omitempty" xml:"adminName,omitempty"`
	// `string`，必填
	//
	// This parameter is required.
	AdminSecret *string `json:"adminSecret,omitempty" xml:"adminSecret,omitempty"`
	// `float64`，必填
	//
	// This parameter is required.
	Cpu *float32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// `int32`，必填
	//
	// This parameter is required.
	DbConnections *int32 `json:"dbConnections,omitempty" xml:"dbConnections,omitempty"`
	// `string`，必填
	//
	// This parameter is required.
	DbHost *string `json:"dbHost,omitempty" xml:"dbHost,omitempty"`
	// `string`，必填
	DbInstanceId *string `json:"dbInstanceId,omitempty" xml:"dbInstanceId,omitempty"`
	// `string`，必填
	//
	// This parameter is required.
	DbName *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
	// `string`，必填
	//
	// This parameter is required.
	DbPassword *string `json:"dbPassword,omitempty" xml:"dbPassword,omitempty"`
	// `int32`，必填
	//
	// This parameter is required.
	DbPort *int32 `json:"dbPort,omitempty" xml:"dbPort,omitempty"`
	// `string`，必填
	//
	// This parameter is required.
	DbType *string `json:"dbType,omitempty" xml:"dbType,omitempty"`
	// `string`，必填
	//
	// This parameter is required.
	DbUsername *string `json:"dbUsername,omitempty" xml:"dbUsername,omitempty"`
	// `string`，必填
	//
	// This parameter is required.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// `string`，必填
	//
	// This parameter is required.
	FunAgentName *string `json:"funAgentName,omitempty" xml:"funAgentName,omitempty"`
	// `int32`，必填
	//
	// This parameter is required.
	Memory *int32 `json:"memory,omitempty" xml:"memory,omitempty"`
	// 可选； `omitempty`
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// `int32`，必填
	//
	// This parameter is required.
	Replicas *int32 `json:"replicas,omitempty" xml:"replicas,omitempty"`
	// `string`，必填
	//
	// This parameter is required.
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// `string`，必填
	//
	// This parameter is required.
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// `string`，必填；JSON 数组 [{vswitchId, zoneId}]（core 解码为 `*string` 但校验非空）
	//
	// This parameter is required.
	VswitchIds *string `json:"vswitchIds,omitempty" xml:"vswitchIds,omitempty"`
}

func (s CreateFunagentInput) String() string {
	return dara.Prettify(s)
}

func (s CreateFunagentInput) GoString() string {
	return s.String()
}

func (s *CreateFunagentInput) GetAdminName() *string {
	return s.AdminName
}

func (s *CreateFunagentInput) GetAdminSecret() *string {
	return s.AdminSecret
}

func (s *CreateFunagentInput) GetCpu() *float32 {
	return s.Cpu
}

func (s *CreateFunagentInput) GetDbConnections() *int32 {
	return s.DbConnections
}

func (s *CreateFunagentInput) GetDbHost() *string {
	return s.DbHost
}

func (s *CreateFunagentInput) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *CreateFunagentInput) GetDbName() *string {
	return s.DbName
}

func (s *CreateFunagentInput) GetDbPassword() *string {
	return s.DbPassword
}

func (s *CreateFunagentInput) GetDbPort() *int32 {
	return s.DbPort
}

func (s *CreateFunagentInput) GetDbType() *string {
	return s.DbType
}

func (s *CreateFunagentInput) GetDbUsername() *string {
	return s.DbUsername
}

func (s *CreateFunagentInput) GetDescription() *string {
	return s.Description
}

func (s *CreateFunagentInput) GetFunAgentName() *string {
	return s.FunAgentName
}

func (s *CreateFunagentInput) GetMemory() *int32 {
	return s.Memory
}

func (s *CreateFunagentInput) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateFunagentInput) GetReplicas() *int32 {
	return s.Replicas
}

func (s *CreateFunagentInput) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateFunagentInput) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateFunagentInput) GetVswitchIds() *string {
	return s.VswitchIds
}

func (s *CreateFunagentInput) SetAdminName(v string) *CreateFunagentInput {
	s.AdminName = &v
	return s
}

func (s *CreateFunagentInput) SetAdminSecret(v string) *CreateFunagentInput {
	s.AdminSecret = &v
	return s
}

func (s *CreateFunagentInput) SetCpu(v float32) *CreateFunagentInput {
	s.Cpu = &v
	return s
}

func (s *CreateFunagentInput) SetDbConnections(v int32) *CreateFunagentInput {
	s.DbConnections = &v
	return s
}

func (s *CreateFunagentInput) SetDbHost(v string) *CreateFunagentInput {
	s.DbHost = &v
	return s
}

func (s *CreateFunagentInput) SetDbInstanceId(v string) *CreateFunagentInput {
	s.DbInstanceId = &v
	return s
}

func (s *CreateFunagentInput) SetDbName(v string) *CreateFunagentInput {
	s.DbName = &v
	return s
}

func (s *CreateFunagentInput) SetDbPassword(v string) *CreateFunagentInput {
	s.DbPassword = &v
	return s
}

func (s *CreateFunagentInput) SetDbPort(v int32) *CreateFunagentInput {
	s.DbPort = &v
	return s
}

func (s *CreateFunagentInput) SetDbType(v string) *CreateFunagentInput {
	s.DbType = &v
	return s
}

func (s *CreateFunagentInput) SetDbUsername(v string) *CreateFunagentInput {
	s.DbUsername = &v
	return s
}

func (s *CreateFunagentInput) SetDescription(v string) *CreateFunagentInput {
	s.Description = &v
	return s
}

func (s *CreateFunagentInput) SetFunAgentName(v string) *CreateFunagentInput {
	s.FunAgentName = &v
	return s
}

func (s *CreateFunagentInput) SetMemory(v int32) *CreateFunagentInput {
	s.Memory = &v
	return s
}

func (s *CreateFunagentInput) SetRegionId(v string) *CreateFunagentInput {
	s.RegionId = &v
	return s
}

func (s *CreateFunagentInput) SetReplicas(v int32) *CreateFunagentInput {
	s.Replicas = &v
	return s
}

func (s *CreateFunagentInput) SetSecurityGroupId(v string) *CreateFunagentInput {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateFunagentInput) SetVpcId(v string) *CreateFunagentInput {
	s.VpcId = &v
	return s
}

func (s *CreateFunagentInput) SetVswitchIds(v string) *CreateFunagentInput {
	s.VswitchIds = &v
	return s
}

func (s *CreateFunagentInput) Validate() error {
	return dara.Validate(s)
}
