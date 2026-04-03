// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFunagent interface {
	dara.Model
	String() string
	GoString() string
	SetAdminName(v string) *Funagent
	GetAdminName() *string
	SetAdminSecret(v string) *Funagent
	GetAdminSecret() *string
	SetCpu(v float32) *Funagent
	GetCpu() *float32
	SetCreatedAt(v string) *Funagent
	GetCreatedAt() *string
	SetDbConnections(v int32) *Funagent
	GetDbConnections() *int32
	SetDbHost(v string) *Funagent
	GetDbHost() *string
	SetDbInstanceId(v string) *Funagent
	GetDbInstanceId() *string
	SetDbName(v string) *Funagent
	GetDbName() *string
	SetDbPassword(v string) *Funagent
	GetDbPassword() *string
	SetDbPort(v int32) *Funagent
	GetDbPort() *int32
	SetDbType(v string) *Funagent
	GetDbType() *string
	SetDbUsername(v string) *Funagent
	GetDbUsername() *string
	SetDescription(v string) *Funagent
	GetDescription() *string
	SetEndpoint(v string) *Funagent
	GetEndpoint() *string
	SetFunagentId(v string) *Funagent
	GetFunagentId() *string
	SetFunagentName(v string) *Funagent
	GetFunagentName() *string
	SetImageUrl(v string) *Funagent
	GetImageUrl() *string
	SetMemory(v int32) *Funagent
	GetMemory() *int32
	SetRegionId(v string) *Funagent
	GetRegionId() *string
	SetReplicas(v int32) *Funagent
	GetReplicas() *int32
	SetSecurityGroupId(v string) *Funagent
	GetSecurityGroupId() *string
	SetStatus(v string) *Funagent
	GetStatus() *string
	SetTenantId(v string) *Funagent
	GetTenantId() *string
	SetUpdatedAt(v string) *Funagent
	GetUpdatedAt() *string
	SetVersion(v string) *Funagent
	GetVersion() *string
	SetVpcId(v string) *Funagent
	GetVpcId() *string
	SetVswitchIds(v string) *Funagent
	GetVswitchIds() *string
}

type Funagent struct {
	AdminName *string `json:"adminName,omitempty" xml:"adminName,omitempty"`
	// 敏感；响应中应脱敏
	AdminSecret *string  `json:"adminSecret,omitempty" xml:"adminSecret,omitempty"`
	Cpu         *float32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// ISO 8601
	CreatedAt     *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	DbConnections *int32  `json:"dbConnections,omitempty" xml:"dbConnections,omitempty"`
	DbHost        *string `json:"dbHost,omitempty" xml:"dbHost,omitempty"`
	DbInstanceId  *string `json:"dbInstanceId,omitempty" xml:"dbInstanceId,omitempty"`
	DbName        *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
	// 敏感；响应中应脱敏
	DbPassword  *string `json:"dbPassword,omitempty" xml:"dbPassword,omitempty"`
	DbPort      *int32  `json:"dbPort,omitempty" xml:"dbPort,omitempty"`
	DbType      *string `json:"dbType,omitempty" xml:"dbType,omitempty"`
	DbUsername  *string `json:"dbUsername,omitempty" xml:"dbUsername,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Endpoint    *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// UUID 字符串
	FunagentId      *string `json:"funagentId,omitempty" xml:"funagentId,omitempty"`
	FunagentName    *string `json:"funagentName,omitempty" xml:"funagentName,omitempty"`
	ImageUrl        *string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
	Memory          *int32  `json:"memory,omitempty" xml:"memory,omitempty"`
	RegionId        *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Replicas        *int32  `json:"replicas,omitempty" xml:"replicas,omitempty"`
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
	TenantId        *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// ISO 8601
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	Version   *string `json:"version,omitempty" xml:"version,omitempty"`
	VpcId     *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// 与存储一致时为 JSON 字符串
	VswitchIds *string `json:"vswitchIds,omitempty" xml:"vswitchIds,omitempty"`
}

func (s Funagent) String() string {
	return dara.Prettify(s)
}

func (s Funagent) GoString() string {
	return s.String()
}

func (s *Funagent) GetAdminName() *string {
	return s.AdminName
}

func (s *Funagent) GetAdminSecret() *string {
	return s.AdminSecret
}

func (s *Funagent) GetCpu() *float32 {
	return s.Cpu
}

func (s *Funagent) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *Funagent) GetDbConnections() *int32 {
	return s.DbConnections
}

func (s *Funagent) GetDbHost() *string {
	return s.DbHost
}

func (s *Funagent) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *Funagent) GetDbName() *string {
	return s.DbName
}

func (s *Funagent) GetDbPassword() *string {
	return s.DbPassword
}

func (s *Funagent) GetDbPort() *int32 {
	return s.DbPort
}

func (s *Funagent) GetDbType() *string {
	return s.DbType
}

func (s *Funagent) GetDbUsername() *string {
	return s.DbUsername
}

func (s *Funagent) GetDescription() *string {
	return s.Description
}

func (s *Funagent) GetEndpoint() *string {
	return s.Endpoint
}

func (s *Funagent) GetFunagentId() *string {
	return s.FunagentId
}

func (s *Funagent) GetFunagentName() *string {
	return s.FunagentName
}

func (s *Funagent) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *Funagent) GetMemory() *int32 {
	return s.Memory
}

func (s *Funagent) GetRegionId() *string {
	return s.RegionId
}

func (s *Funagent) GetReplicas() *int32 {
	return s.Replicas
}

func (s *Funagent) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *Funagent) GetStatus() *string {
	return s.Status
}

func (s *Funagent) GetTenantId() *string {
	return s.TenantId
}

func (s *Funagent) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *Funagent) GetVersion() *string {
	return s.Version
}

func (s *Funagent) GetVpcId() *string {
	return s.VpcId
}

func (s *Funagent) GetVswitchIds() *string {
	return s.VswitchIds
}

func (s *Funagent) SetAdminName(v string) *Funagent {
	s.AdminName = &v
	return s
}

func (s *Funagent) SetAdminSecret(v string) *Funagent {
	s.AdminSecret = &v
	return s
}

func (s *Funagent) SetCpu(v float32) *Funagent {
	s.Cpu = &v
	return s
}

func (s *Funagent) SetCreatedAt(v string) *Funagent {
	s.CreatedAt = &v
	return s
}

func (s *Funagent) SetDbConnections(v int32) *Funagent {
	s.DbConnections = &v
	return s
}

func (s *Funagent) SetDbHost(v string) *Funagent {
	s.DbHost = &v
	return s
}

func (s *Funagent) SetDbInstanceId(v string) *Funagent {
	s.DbInstanceId = &v
	return s
}

func (s *Funagent) SetDbName(v string) *Funagent {
	s.DbName = &v
	return s
}

func (s *Funagent) SetDbPassword(v string) *Funagent {
	s.DbPassword = &v
	return s
}

func (s *Funagent) SetDbPort(v int32) *Funagent {
	s.DbPort = &v
	return s
}

func (s *Funagent) SetDbType(v string) *Funagent {
	s.DbType = &v
	return s
}

func (s *Funagent) SetDbUsername(v string) *Funagent {
	s.DbUsername = &v
	return s
}

func (s *Funagent) SetDescription(v string) *Funagent {
	s.Description = &v
	return s
}

func (s *Funagent) SetEndpoint(v string) *Funagent {
	s.Endpoint = &v
	return s
}

func (s *Funagent) SetFunagentId(v string) *Funagent {
	s.FunagentId = &v
	return s
}

func (s *Funagent) SetFunagentName(v string) *Funagent {
	s.FunagentName = &v
	return s
}

func (s *Funagent) SetImageUrl(v string) *Funagent {
	s.ImageUrl = &v
	return s
}

func (s *Funagent) SetMemory(v int32) *Funagent {
	s.Memory = &v
	return s
}

func (s *Funagent) SetRegionId(v string) *Funagent {
	s.RegionId = &v
	return s
}

func (s *Funagent) SetReplicas(v int32) *Funagent {
	s.Replicas = &v
	return s
}

func (s *Funagent) SetSecurityGroupId(v string) *Funagent {
	s.SecurityGroupId = &v
	return s
}

func (s *Funagent) SetStatus(v string) *Funagent {
	s.Status = &v
	return s
}

func (s *Funagent) SetTenantId(v string) *Funagent {
	s.TenantId = &v
	return s
}

func (s *Funagent) SetUpdatedAt(v string) *Funagent {
	s.UpdatedAt = &v
	return s
}

func (s *Funagent) SetVersion(v string) *Funagent {
	s.Version = &v
	return s
}

func (s *Funagent) SetVpcId(v string) *Funagent {
	s.VpcId = &v
	return s
}

func (s *Funagent) SetVswitchIds(v string) *Funagent {
	s.VswitchIds = &v
	return s
}

func (s *Funagent) Validate() error {
	return dara.Validate(s)
}
