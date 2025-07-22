// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddHDMInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEngine(v string) *AddHDMInstanceRequest
	GetEngine() *string
	SetFlushAccount(v string) *AddHDMInstanceRequest
	GetFlushAccount() *string
	SetInstanceAlias(v string) *AddHDMInstanceRequest
	GetInstanceAlias() *string
	SetInstanceArea(v string) *AddHDMInstanceRequest
	GetInstanceArea() *string
	SetInstanceId(v string) *AddHDMInstanceRequest
	GetInstanceId() *string
	SetIp(v string) *AddHDMInstanceRequest
	GetIp() *string
	SetNetworkType(v string) *AddHDMInstanceRequest
	GetNetworkType() *string
	SetPassword(v string) *AddHDMInstanceRequest
	GetPassword() *string
	SetPort(v string) *AddHDMInstanceRequest
	GetPort() *string
	SetRegion(v string) *AddHDMInstanceRequest
	GetRegion() *string
	SetUsername(v string) *AddHDMInstanceRequest
	GetUsername() *string
	SetVpcId(v string) *AddHDMInstanceRequest
	GetVpcId() *string
	SetContext(v string) *AddHDMInstanceRequest
	GetContext() *string
}

type AddHDMInstanceRequest struct {
	// The database engine. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **PostgreSQL**
	//
	// 	- **SQLServer**
	//
	// 	- **PolarDBMySQL**
	//
	// 	- **PolarDBPostgreSQL**
	//
	// 	- **Redis**
	//
	// 	- **MongoDB**
	//
	// 	- **PolarDBOracle**
	//
	// 	- **PolarDBX**
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	FlushAccount *string `json:"FlushAccount,omitempty" xml:"FlushAccount,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// yuecq--test****
	InstanceAlias *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	// The type of the instance on which the database is deployed. Valid values:
	//
	// 	- **RDS**: an Alibaba Cloud database instance.
	//
	// 	- **ECS**: an Elastic Compute Service (ECS) instance on which a self-managed database is deployed.
	//
	// 	- **IDC**: a self-managed database instance that is not deployed on Alibaba Cloud.
	//
	// >  IDC refers to your data center.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS
	InstanceArea *string `json:"InstanceArea,omitempty" xml:"InstanceArea,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The endpoint that is used to access the instance over internal networks.
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****.mysql.rds.aliyuncs.com
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The network type of the instance.
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The password for the username.
	//
	// example:
	//
	// 122****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The port that is used to access the instance over internal networks.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The username that is used to log on to the database.
	//
	// example:
	//
	// test****
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-m5e666n89m2bx8jar****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	Context *string `json:"__context,omitempty" xml:"__context,omitempty"`
}

func (s AddHDMInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddHDMInstanceRequest) GoString() string {
	return s.String()
}

func (s *AddHDMInstanceRequest) GetEngine() *string {
	return s.Engine
}

func (s *AddHDMInstanceRequest) GetFlushAccount() *string {
	return s.FlushAccount
}

func (s *AddHDMInstanceRequest) GetInstanceAlias() *string {
	return s.InstanceAlias
}

func (s *AddHDMInstanceRequest) GetInstanceArea() *string {
	return s.InstanceArea
}

func (s *AddHDMInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddHDMInstanceRequest) GetIp() *string {
	return s.Ip
}

func (s *AddHDMInstanceRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *AddHDMInstanceRequest) GetPassword() *string {
	return s.Password
}

func (s *AddHDMInstanceRequest) GetPort() *string {
	return s.Port
}

func (s *AddHDMInstanceRequest) GetRegion() *string {
	return s.Region
}

func (s *AddHDMInstanceRequest) GetUsername() *string {
	return s.Username
}

func (s *AddHDMInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *AddHDMInstanceRequest) GetContext() *string {
	return s.Context
}

func (s *AddHDMInstanceRequest) SetEngine(v string) *AddHDMInstanceRequest {
	s.Engine = &v
	return s
}

func (s *AddHDMInstanceRequest) SetFlushAccount(v string) *AddHDMInstanceRequest {
	s.FlushAccount = &v
	return s
}

func (s *AddHDMInstanceRequest) SetInstanceAlias(v string) *AddHDMInstanceRequest {
	s.InstanceAlias = &v
	return s
}

func (s *AddHDMInstanceRequest) SetInstanceArea(v string) *AddHDMInstanceRequest {
	s.InstanceArea = &v
	return s
}

func (s *AddHDMInstanceRequest) SetInstanceId(v string) *AddHDMInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *AddHDMInstanceRequest) SetIp(v string) *AddHDMInstanceRequest {
	s.Ip = &v
	return s
}

func (s *AddHDMInstanceRequest) SetNetworkType(v string) *AddHDMInstanceRequest {
	s.NetworkType = &v
	return s
}

func (s *AddHDMInstanceRequest) SetPassword(v string) *AddHDMInstanceRequest {
	s.Password = &v
	return s
}

func (s *AddHDMInstanceRequest) SetPort(v string) *AddHDMInstanceRequest {
	s.Port = &v
	return s
}

func (s *AddHDMInstanceRequest) SetRegion(v string) *AddHDMInstanceRequest {
	s.Region = &v
	return s
}

func (s *AddHDMInstanceRequest) SetUsername(v string) *AddHDMInstanceRequest {
	s.Username = &v
	return s
}

func (s *AddHDMInstanceRequest) SetVpcId(v string) *AddHDMInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *AddHDMInstanceRequest) SetContext(v string) *AddHDMInstanceRequest {
	s.Context = &v
	return s
}

func (s *AddHDMInstanceRequest) Validate() error {
	return dara.Validate(s)
}
