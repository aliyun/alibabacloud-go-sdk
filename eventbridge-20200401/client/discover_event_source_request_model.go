// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiscoverEventSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceMySQLParameters(v *DiscoverEventSourceRequestSourceMySQLParameters) *DiscoverEventSourceRequest
	GetSourceMySQLParameters() *DiscoverEventSourceRequestSourceMySQLParameters
}

type DiscoverEventSourceRequest struct {
	SourceMySQLParameters *DiscoverEventSourceRequestSourceMySQLParameters `json:"SourceMySQLParameters,omitempty" xml:"SourceMySQLParameters,omitempty" type:"Struct"`
}

func (s DiscoverEventSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DiscoverEventSourceRequest) GoString() string {
	return s.String()
}

func (s *DiscoverEventSourceRequest) GetSourceMySQLParameters() *DiscoverEventSourceRequestSourceMySQLParameters {
	return s.SourceMySQLParameters
}

func (s *DiscoverEventSourceRequest) SetSourceMySQLParameters(v *DiscoverEventSourceRequestSourceMySQLParameters) *DiscoverEventSourceRequest {
	s.SourceMySQLParameters = v
	return s
}

func (s *DiscoverEventSourceRequest) Validate() error {
	if s.SourceMySQLParameters != nil {
		if err := s.SourceMySQLParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DiscoverEventSourceRequestSourceMySQLParameters struct {
	// example:
	//
	// database1
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// example:
	//
	// rm-xxx.mysql.rds.aliyuncs.com
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// example:
	//
	// 20
	Limit *string `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// example:
	//
	// PrivateNetwork
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// example:
	//
	// 30
	Offset *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// example:
	//
	// 1234xxx
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// sg-bp1ic0vsbwyv176e9inx
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// database1.table1
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// user1
	User *string `json:"User,omitempty" xml:"User,omitempty"`
	// example:
	//
	// vsw-gw824tpaptxtlo256lqub
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// example:
	//
	// vpc-uf6hwiei8u5uil3bfahc1
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DiscoverEventSourceRequestSourceMySQLParameters) String() string {
	return dara.Prettify(s)
}

func (s DiscoverEventSourceRequestSourceMySQLParameters) GoString() string {
	return s.String()
}

func (s *DiscoverEventSourceRequestSourceMySQLParameters) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DiscoverEventSourceRequestSourceMySQLParameters) GetHostName() *string {
	return s.HostName
}

func (s *DiscoverEventSourceRequestSourceMySQLParameters) GetLimit() *string {
	return s.Limit
}

func (s *DiscoverEventSourceRequestSourceMySQLParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DiscoverEventSourceRequestSourceMySQLParameters) GetOffset() *string {
	return s.Offset
}

func (s *DiscoverEventSourceRequestSourceMySQLParameters) GetPassword() *string {
	return s.Password
}

func (s *DiscoverEventSourceRequestSourceMySQLParameters) GetPort() *int32 {
	return s.Port
}

func (s *DiscoverEventSourceRequestSourceMySQLParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *DiscoverEventSourceRequestSourceMySQLParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DiscoverEventSourceRequestSourceMySQLParameters) GetTableName() *string {
	return s.TableName
}

func (s *DiscoverEventSourceRequestSourceMySQLParameters) GetUser() *string {
	return s.User
}

func (s *DiscoverEventSourceRequestSourceMySQLParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *DiscoverEventSourceRequestSourceMySQLParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *DiscoverEventSourceRequestSourceMySQLParameters) SetDatabaseName(v string) *DiscoverEventSourceRequestSourceMySQLParameters {
	s.DatabaseName = &v
	return s
}

func (s *DiscoverEventSourceRequestSourceMySQLParameters) SetHostName(v string) *DiscoverEventSourceRequestSourceMySQLParameters {
	s.HostName = &v
	return s
}

func (s *DiscoverEventSourceRequestSourceMySQLParameters) SetLimit(v string) *DiscoverEventSourceRequestSourceMySQLParameters {
	s.Limit = &v
	return s
}

func (s *DiscoverEventSourceRequestSourceMySQLParameters) SetNetworkType(v string) *DiscoverEventSourceRequestSourceMySQLParameters {
	s.NetworkType = &v
	return s
}

func (s *DiscoverEventSourceRequestSourceMySQLParameters) SetOffset(v string) *DiscoverEventSourceRequestSourceMySQLParameters {
	s.Offset = &v
	return s
}

func (s *DiscoverEventSourceRequestSourceMySQLParameters) SetPassword(v string) *DiscoverEventSourceRequestSourceMySQLParameters {
	s.Password = &v
	return s
}

func (s *DiscoverEventSourceRequestSourceMySQLParameters) SetPort(v int32) *DiscoverEventSourceRequestSourceMySQLParameters {
	s.Port = &v
	return s
}

func (s *DiscoverEventSourceRequestSourceMySQLParameters) SetRegionId(v string) *DiscoverEventSourceRequestSourceMySQLParameters {
	s.RegionId = &v
	return s
}

func (s *DiscoverEventSourceRequestSourceMySQLParameters) SetSecurityGroupId(v string) *DiscoverEventSourceRequestSourceMySQLParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *DiscoverEventSourceRequestSourceMySQLParameters) SetTableName(v string) *DiscoverEventSourceRequestSourceMySQLParameters {
	s.TableName = &v
	return s
}

func (s *DiscoverEventSourceRequestSourceMySQLParameters) SetUser(v string) *DiscoverEventSourceRequestSourceMySQLParameters {
	s.User = &v
	return s
}

func (s *DiscoverEventSourceRequestSourceMySQLParameters) SetVSwitchIds(v string) *DiscoverEventSourceRequestSourceMySQLParameters {
	s.VSwitchIds = &v
	return s
}

func (s *DiscoverEventSourceRequestSourceMySQLParameters) SetVpcId(v string) *DiscoverEventSourceRequestSourceMySQLParameters {
	s.VpcId = &v
	return s
}

func (s *DiscoverEventSourceRequestSourceMySQLParameters) Validate() error {
	return dara.Validate(s)
}
