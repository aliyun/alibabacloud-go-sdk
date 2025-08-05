// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSourceMySQLParameters interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseName(v string) *SourceMySQLParameters
	GetDatabaseName() *string
	SetHostName(v string) *SourceMySQLParameters
	GetHostName() *string
	SetIncludeSchemaChanges(v string) *SourceMySQLParameters
	GetIncludeSchemaChanges() *string
	SetNetworkType(v string) *SourceMySQLParameters
	GetNetworkType() *string
	SetPassword(v string) *SourceMySQLParameters
	GetPassword() *string
	SetPort(v int32) *SourceMySQLParameters
	GetPort() *int32
	SetRegionId(v string) *SourceMySQLParameters
	GetRegionId() *string
	SetSecurityGroupId(v string) *SourceMySQLParameters
	GetSecurityGroupId() *string
	SetSnapshotMode(v string) *SourceMySQLParameters
	GetSnapshotMode() *string
	SetTableNames(v string) *SourceMySQLParameters
	GetTableNames() *string
	SetUser(v string) *SourceMySQLParameters
	GetUser() *string
	SetVSwitchIds(v string) *SourceMySQLParameters
	GetVSwitchIds() *string
	SetVpcId(v string) *SourceMySQLParameters
	GetVpcId() *string
}

type SourceMySQLParameters struct {
	DatabaseName         *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	HostName             *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	IncludeSchemaChanges *string `json:"IncludeSchemaChanges,omitempty" xml:"IncludeSchemaChanges,omitempty"`
	NetworkType          *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Password             *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Port                 *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupId      *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SnapshotMode         *string `json:"SnapshotMode,omitempty" xml:"SnapshotMode,omitempty"`
	TableNames           *string `json:"TableNames,omitempty" xml:"TableNames,omitempty"`
	User                 *string `json:"User,omitempty" xml:"User,omitempty"`
	VSwitchIds           *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s SourceMySQLParameters) String() string {
	return dara.Prettify(s)
}

func (s SourceMySQLParameters) GoString() string {
	return s.String()
}

func (s *SourceMySQLParameters) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *SourceMySQLParameters) GetHostName() *string {
	return s.HostName
}

func (s *SourceMySQLParameters) GetIncludeSchemaChanges() *string {
	return s.IncludeSchemaChanges
}

func (s *SourceMySQLParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *SourceMySQLParameters) GetPassword() *string {
	return s.Password
}

func (s *SourceMySQLParameters) GetPort() *int32 {
	return s.Port
}

func (s *SourceMySQLParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *SourceMySQLParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *SourceMySQLParameters) GetSnapshotMode() *string {
	return s.SnapshotMode
}

func (s *SourceMySQLParameters) GetTableNames() *string {
	return s.TableNames
}

func (s *SourceMySQLParameters) GetUser() *string {
	return s.User
}

func (s *SourceMySQLParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *SourceMySQLParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *SourceMySQLParameters) SetDatabaseName(v string) *SourceMySQLParameters {
	s.DatabaseName = &v
	return s
}

func (s *SourceMySQLParameters) SetHostName(v string) *SourceMySQLParameters {
	s.HostName = &v
	return s
}

func (s *SourceMySQLParameters) SetIncludeSchemaChanges(v string) *SourceMySQLParameters {
	s.IncludeSchemaChanges = &v
	return s
}

func (s *SourceMySQLParameters) SetNetworkType(v string) *SourceMySQLParameters {
	s.NetworkType = &v
	return s
}

func (s *SourceMySQLParameters) SetPassword(v string) *SourceMySQLParameters {
	s.Password = &v
	return s
}

func (s *SourceMySQLParameters) SetPort(v int32) *SourceMySQLParameters {
	s.Port = &v
	return s
}

func (s *SourceMySQLParameters) SetRegionId(v string) *SourceMySQLParameters {
	s.RegionId = &v
	return s
}

func (s *SourceMySQLParameters) SetSecurityGroupId(v string) *SourceMySQLParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *SourceMySQLParameters) SetSnapshotMode(v string) *SourceMySQLParameters {
	s.SnapshotMode = &v
	return s
}

func (s *SourceMySQLParameters) SetTableNames(v string) *SourceMySQLParameters {
	s.TableNames = &v
	return s
}

func (s *SourceMySQLParameters) SetUser(v string) *SourceMySQLParameters {
	s.User = &v
	return s
}

func (s *SourceMySQLParameters) SetVSwitchIds(v string) *SourceMySQLParameters {
	s.VSwitchIds = &v
	return s
}

func (s *SourceMySQLParameters) SetVpcId(v string) *SourceMySQLParameters {
	s.VpcId = &v
	return s
}

func (s *SourceMySQLParameters) Validate() error {
	return dara.Validate(s)
}
