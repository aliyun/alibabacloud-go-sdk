// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestEventSourceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceMySQLParameters(v *TestEventSourceConfigRequestSourceMySQLParameters) *TestEventSourceConfigRequest
	GetSourceMySQLParameters() *TestEventSourceConfigRequestSourceMySQLParameters
}

type TestEventSourceConfigRequest struct {
	// The parameters that are configured if you specify MySQL as the event source.
	SourceMySQLParameters *TestEventSourceConfigRequestSourceMySQLParameters `json:"SourceMySQLParameters,omitempty" xml:"SourceMySQLParameters,omitempty" type:"Struct"`
}

func (s TestEventSourceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s TestEventSourceConfigRequest) GoString() string {
	return s.String()
}

func (s *TestEventSourceConfigRequest) GetSourceMySQLParameters() *TestEventSourceConfigRequestSourceMySQLParameters {
	return s.SourceMySQLParameters
}

func (s *TestEventSourceConfigRequest) SetSourceMySQLParameters(v *TestEventSourceConfigRequestSourceMySQLParameters) *TestEventSourceConfigRequest {
	s.SourceMySQLParameters = v
	return s
}

func (s *TestEventSourceConfigRequest) Validate() error {
	if s.SourceMySQLParameters != nil {
		if err := s.SourceMySQLParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TestEventSourceConfigRequestSourceMySQLParameters struct {
	AllowedCIDRs *string `json:"AllowedCIDRs,omitempty" xml:"AllowedCIDRs,omitempty"`
	// The database name.
	//
	// example:
	//
	// database1
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The endpoint of the database.
	//
	// example:
	//
	// rm-bp1vxxx.mysql.rds.aliyuncs.com
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The network type.
	//
	// Valid values:
	//
	// 	- PrivateNetwork
	//
	// 	- PublicNetwork
	//
	// example:
	//
	// PrivateNetwork
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The password that is used for authentication.
	//
	// example:
	//
	// 1234xxx
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The port that is used to connect to the database.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-xxx
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The table name. The name must be prefixed with the database name. ${DatabaseName}.${TableName}
	//
	// example:
	//
	// database1.table1
	TableNames *string `json:"TableNames,omitempty" xml:"TableNames,omitempty"`
	// The username that is used to log on to the database.
	//
	// example:
	//
	// user***
	User *string `json:"User,omitempty" xml:"User,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1gb7xxx
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-xxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s TestEventSourceConfigRequestSourceMySQLParameters) String() string {
	return dara.Prettify(s)
}

func (s TestEventSourceConfigRequestSourceMySQLParameters) GoString() string {
	return s.String()
}

func (s *TestEventSourceConfigRequestSourceMySQLParameters) GetAllowedCIDRs() *string {
	return s.AllowedCIDRs
}

func (s *TestEventSourceConfigRequestSourceMySQLParameters) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *TestEventSourceConfigRequestSourceMySQLParameters) GetHostName() *string {
	return s.HostName
}

func (s *TestEventSourceConfigRequestSourceMySQLParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *TestEventSourceConfigRequestSourceMySQLParameters) GetPassword() *string {
	return s.Password
}

func (s *TestEventSourceConfigRequestSourceMySQLParameters) GetPort() *int32 {
	return s.Port
}

func (s *TestEventSourceConfigRequestSourceMySQLParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *TestEventSourceConfigRequestSourceMySQLParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *TestEventSourceConfigRequestSourceMySQLParameters) GetTableNames() *string {
	return s.TableNames
}

func (s *TestEventSourceConfigRequestSourceMySQLParameters) GetUser() *string {
	return s.User
}

func (s *TestEventSourceConfigRequestSourceMySQLParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *TestEventSourceConfigRequestSourceMySQLParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *TestEventSourceConfigRequestSourceMySQLParameters) SetAllowedCIDRs(v string) *TestEventSourceConfigRequestSourceMySQLParameters {
	s.AllowedCIDRs = &v
	return s
}

func (s *TestEventSourceConfigRequestSourceMySQLParameters) SetDatabaseName(v string) *TestEventSourceConfigRequestSourceMySQLParameters {
	s.DatabaseName = &v
	return s
}

func (s *TestEventSourceConfigRequestSourceMySQLParameters) SetHostName(v string) *TestEventSourceConfigRequestSourceMySQLParameters {
	s.HostName = &v
	return s
}

func (s *TestEventSourceConfigRequestSourceMySQLParameters) SetNetworkType(v string) *TestEventSourceConfigRequestSourceMySQLParameters {
	s.NetworkType = &v
	return s
}

func (s *TestEventSourceConfigRequestSourceMySQLParameters) SetPassword(v string) *TestEventSourceConfigRequestSourceMySQLParameters {
	s.Password = &v
	return s
}

func (s *TestEventSourceConfigRequestSourceMySQLParameters) SetPort(v int32) *TestEventSourceConfigRequestSourceMySQLParameters {
	s.Port = &v
	return s
}

func (s *TestEventSourceConfigRequestSourceMySQLParameters) SetRegionId(v string) *TestEventSourceConfigRequestSourceMySQLParameters {
	s.RegionId = &v
	return s
}

func (s *TestEventSourceConfigRequestSourceMySQLParameters) SetSecurityGroupId(v string) *TestEventSourceConfigRequestSourceMySQLParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *TestEventSourceConfigRequestSourceMySQLParameters) SetTableNames(v string) *TestEventSourceConfigRequestSourceMySQLParameters {
	s.TableNames = &v
	return s
}

func (s *TestEventSourceConfigRequestSourceMySQLParameters) SetUser(v string) *TestEventSourceConfigRequestSourceMySQLParameters {
	s.User = &v
	return s
}

func (s *TestEventSourceConfigRequestSourceMySQLParameters) SetVSwitchIds(v string) *TestEventSourceConfigRequestSourceMySQLParameters {
	s.VSwitchIds = &v
	return s
}

func (s *TestEventSourceConfigRequestSourceMySQLParameters) SetVpcId(v string) *TestEventSourceConfigRequestSourceMySQLParameters {
	s.VpcId = &v
	return s
}

func (s *TestEventSourceConfigRequestSourceMySQLParameters) Validate() error {
	return dara.Validate(s)
}
