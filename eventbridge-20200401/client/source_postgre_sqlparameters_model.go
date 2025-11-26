// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSourcePostgreSQLParameters interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseName(v string) *SourcePostgreSQLParameters
	GetDatabaseName() *string
	SetHostName(v string) *SourcePostgreSQLParameters
	GetHostName() *string
	SetNetworkType(v string) *SourcePostgreSQLParameters
	GetNetworkType() *string
	SetPassword(v string) *SourcePostgreSQLParameters
	GetPassword() *string
	SetPort(v int32) *SourcePostgreSQLParameters
	GetPort() *int32
	SetRegionId(v string) *SourcePostgreSQLParameters
	GetRegionId() *string
	SetSchemaName(v string) *SourcePostgreSQLParameters
	GetSchemaName() *string
	SetSecurityGroupId(v string) *SourcePostgreSQLParameters
	GetSecurityGroupId() *string
	SetSnapshotMode(v string) *SourcePostgreSQLParameters
	GetSnapshotMode() *string
	SetTableNames(v string) *SourcePostgreSQLParameters
	GetTableNames() *string
	SetUser(v string) *SourcePostgreSQLParameters
	GetUser() *string
	SetVSwitchIds(v string) *SourcePostgreSQLParameters
	GetVSwitchIds() *string
	SetVpcId(v string) *SourcePostgreSQLParameters
	GetVpcId() *string
}

type SourcePostgreSQLParameters struct {
	DatabaseName    *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	HostName        *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	NetworkType     *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Port            *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SchemaName      *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SnapshotMode    *string `json:"SnapshotMode,omitempty" xml:"SnapshotMode,omitempty"`
	TableNames      *string `json:"TableNames,omitempty" xml:"TableNames,omitempty"`
	User            *string `json:"User,omitempty" xml:"User,omitempty"`
	VSwitchIds      *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s SourcePostgreSQLParameters) String() string {
	return dara.Prettify(s)
}

func (s SourcePostgreSQLParameters) GoString() string {
	return s.String()
}

func (s *SourcePostgreSQLParameters) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *SourcePostgreSQLParameters) GetHostName() *string {
	return s.HostName
}

func (s *SourcePostgreSQLParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *SourcePostgreSQLParameters) GetPassword() *string {
	return s.Password
}

func (s *SourcePostgreSQLParameters) GetPort() *int32 {
	return s.Port
}

func (s *SourcePostgreSQLParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *SourcePostgreSQLParameters) GetSchemaName() *string {
	return s.SchemaName
}

func (s *SourcePostgreSQLParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *SourcePostgreSQLParameters) GetSnapshotMode() *string {
	return s.SnapshotMode
}

func (s *SourcePostgreSQLParameters) GetTableNames() *string {
	return s.TableNames
}

func (s *SourcePostgreSQLParameters) GetUser() *string {
	return s.User
}

func (s *SourcePostgreSQLParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *SourcePostgreSQLParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *SourcePostgreSQLParameters) SetDatabaseName(v string) *SourcePostgreSQLParameters {
	s.DatabaseName = &v
	return s
}

func (s *SourcePostgreSQLParameters) SetHostName(v string) *SourcePostgreSQLParameters {
	s.HostName = &v
	return s
}

func (s *SourcePostgreSQLParameters) SetNetworkType(v string) *SourcePostgreSQLParameters {
	s.NetworkType = &v
	return s
}

func (s *SourcePostgreSQLParameters) SetPassword(v string) *SourcePostgreSQLParameters {
	s.Password = &v
	return s
}

func (s *SourcePostgreSQLParameters) SetPort(v int32) *SourcePostgreSQLParameters {
	s.Port = &v
	return s
}

func (s *SourcePostgreSQLParameters) SetRegionId(v string) *SourcePostgreSQLParameters {
	s.RegionId = &v
	return s
}

func (s *SourcePostgreSQLParameters) SetSchemaName(v string) *SourcePostgreSQLParameters {
	s.SchemaName = &v
	return s
}

func (s *SourcePostgreSQLParameters) SetSecurityGroupId(v string) *SourcePostgreSQLParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *SourcePostgreSQLParameters) SetSnapshotMode(v string) *SourcePostgreSQLParameters {
	s.SnapshotMode = &v
	return s
}

func (s *SourcePostgreSQLParameters) SetTableNames(v string) *SourcePostgreSQLParameters {
	s.TableNames = &v
	return s
}

func (s *SourcePostgreSQLParameters) SetUser(v string) *SourcePostgreSQLParameters {
	s.User = &v
	return s
}

func (s *SourcePostgreSQLParameters) SetVSwitchIds(v string) *SourcePostgreSQLParameters {
	s.VSwitchIds = &v
	return s
}

func (s *SourcePostgreSQLParameters) SetVpcId(v string) *SourcePostgreSQLParameters {
	s.VpcId = &v
	return s
}

func (s *SourcePostgreSQLParameters) Validate() error {
	return dara.Validate(s)
}
