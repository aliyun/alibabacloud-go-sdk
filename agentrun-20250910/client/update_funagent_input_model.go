// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFunagentInput interface {
	dara.Model
	String() string
	GoString() string
	SetAdminName(v string) *UpdateFunagentInput
	GetAdminName() *string
	SetAdminSecret(v string) *UpdateFunagentInput
	GetAdminSecret() *string
	SetCpu(v float32) *UpdateFunagentInput
	GetCpu() *float32
	SetDbConnections(v int32) *UpdateFunagentInput
	GetDbConnections() *int32
	SetDbHost(v string) *UpdateFunagentInput
	GetDbHost() *string
	SetDbInstanceId(v string) *UpdateFunagentInput
	GetDbInstanceId() *string
	SetDbName(v string) *UpdateFunagentInput
	GetDbName() *string
	SetDbPassword(v string) *UpdateFunagentInput
	GetDbPassword() *string
	SetDbPort(v int32) *UpdateFunagentInput
	GetDbPort() *int32
	SetDbType(v string) *UpdateFunagentInput
	GetDbType() *string
	SetDbUsername(v string) *UpdateFunagentInput
	GetDbUsername() *string
	SetDescription(v string) *UpdateFunagentInput
	GetDescription() *string
	SetMemory(v int32) *UpdateFunagentInput
	GetMemory() *int32
	SetReplicas(v int32) *UpdateFunagentInput
	GetReplicas() *int32
	SetVersion(v string) *UpdateFunagentInput
	GetVersion() *string
}

type UpdateFunagentInput struct {
	AdminName *string `json:"adminName,omitempty" xml:"adminName,omitempty"`
	// 敏感字段
	AdminSecret   *string  `json:"adminSecret,omitempty" xml:"adminSecret,omitempty"`
	Cpu           *float32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	DbConnections *int32   `json:"dbConnections,omitempty" xml:"dbConnections,omitempty"`
	DbHost        *string  `json:"dbHost,omitempty" xml:"dbHost,omitempty"`
	DbInstanceId  *string  `json:"dbInstanceId,omitempty" xml:"dbInstanceId,omitempty"`
	DbName        *string  `json:"dbName,omitempty" xml:"dbName,omitempty"`
	// 敏感字段
	DbPassword  *string `json:"dbPassword,omitempty" xml:"dbPassword,omitempty"`
	DbPort      *int32  `json:"dbPort,omitempty" xml:"dbPort,omitempty"`
	DbType      *string `json:"dbType,omitempty" xml:"dbType,omitempty"`
	DbUsername  *string `json:"dbUsername,omitempty" xml:"dbUsername,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Memory      *int32  `json:"memory,omitempty" xml:"memory,omitempty"`
	Replicas    *int32  `json:"replicas,omitempty" xml:"replicas,omitempty"`
	Version     *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UpdateFunagentInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateFunagentInput) GoString() string {
	return s.String()
}

func (s *UpdateFunagentInput) GetAdminName() *string {
	return s.AdminName
}

func (s *UpdateFunagentInput) GetAdminSecret() *string {
	return s.AdminSecret
}

func (s *UpdateFunagentInput) GetCpu() *float32 {
	return s.Cpu
}

func (s *UpdateFunagentInput) GetDbConnections() *int32 {
	return s.DbConnections
}

func (s *UpdateFunagentInput) GetDbHost() *string {
	return s.DbHost
}

func (s *UpdateFunagentInput) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *UpdateFunagentInput) GetDbName() *string {
	return s.DbName
}

func (s *UpdateFunagentInput) GetDbPassword() *string {
	return s.DbPassword
}

func (s *UpdateFunagentInput) GetDbPort() *int32 {
	return s.DbPort
}

func (s *UpdateFunagentInput) GetDbType() *string {
	return s.DbType
}

func (s *UpdateFunagentInput) GetDbUsername() *string {
	return s.DbUsername
}

func (s *UpdateFunagentInput) GetDescription() *string {
	return s.Description
}

func (s *UpdateFunagentInput) GetMemory() *int32 {
	return s.Memory
}

func (s *UpdateFunagentInput) GetReplicas() *int32 {
	return s.Replicas
}

func (s *UpdateFunagentInput) GetVersion() *string {
	return s.Version
}

func (s *UpdateFunagentInput) SetAdminName(v string) *UpdateFunagentInput {
	s.AdminName = &v
	return s
}

func (s *UpdateFunagentInput) SetAdminSecret(v string) *UpdateFunagentInput {
	s.AdminSecret = &v
	return s
}

func (s *UpdateFunagentInput) SetCpu(v float32) *UpdateFunagentInput {
	s.Cpu = &v
	return s
}

func (s *UpdateFunagentInput) SetDbConnections(v int32) *UpdateFunagentInput {
	s.DbConnections = &v
	return s
}

func (s *UpdateFunagentInput) SetDbHost(v string) *UpdateFunagentInput {
	s.DbHost = &v
	return s
}

func (s *UpdateFunagentInput) SetDbInstanceId(v string) *UpdateFunagentInput {
	s.DbInstanceId = &v
	return s
}

func (s *UpdateFunagentInput) SetDbName(v string) *UpdateFunagentInput {
	s.DbName = &v
	return s
}

func (s *UpdateFunagentInput) SetDbPassword(v string) *UpdateFunagentInput {
	s.DbPassword = &v
	return s
}

func (s *UpdateFunagentInput) SetDbPort(v int32) *UpdateFunagentInput {
	s.DbPort = &v
	return s
}

func (s *UpdateFunagentInput) SetDbType(v string) *UpdateFunagentInput {
	s.DbType = &v
	return s
}

func (s *UpdateFunagentInput) SetDbUsername(v string) *UpdateFunagentInput {
	s.DbUsername = &v
	return s
}

func (s *UpdateFunagentInput) SetDescription(v string) *UpdateFunagentInput {
	s.Description = &v
	return s
}

func (s *UpdateFunagentInput) SetMemory(v int32) *UpdateFunagentInput {
	s.Memory = &v
	return s
}

func (s *UpdateFunagentInput) SetReplicas(v int32) *UpdateFunagentInput {
	s.Replicas = &v
	return s
}

func (s *UpdateFunagentInput) SetVersion(v string) *UpdateFunagentInput {
	s.Version = &v
	return s
}

func (s *UpdateFunagentInput) Validate() error {
	return dara.Validate(s)
}
