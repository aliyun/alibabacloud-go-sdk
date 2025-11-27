// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceNetInfoForChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceNetInfos(v *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfos) *DescribeDBInstanceNetInfoForChannelResponseBody
	GetDBInstanceNetInfos() *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfos
	SetInstanceNetworkType(v string) *DescribeDBInstanceNetInfoForChannelResponseBody
	GetInstanceNetworkType() *string
	SetRequestId(v string) *DescribeDBInstanceNetInfoForChannelResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceNetInfoForChannelResponseBody struct {
	// The information about the instance connection.
	DBInstanceNetInfos *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfos `json:"DBInstanceNetInfos,omitempty" xml:"DBInstanceNetInfos,omitempty" type:"Struct"`
	// The network type of the instance. Valid values:
	//
	// 	- **VPC**: a virtual private cloud (VPC)
	//
	// 	- **Classic**: classic network
	//
	// example:
	//
	// VPC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 61DF1F28-F409-50C0-B90A-CCE82D44****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceNetInfoForChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoForChannelResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBody) GetDBInstanceNetInfos() *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfos {
	return s.DBInstanceNetInfos
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBody) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBody) SetDBInstanceNetInfos(v *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfos) *DescribeDBInstanceNetInfoForChannelResponseBody {
	s.DBInstanceNetInfos = v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBody) SetInstanceNetworkType(v string) *DescribeDBInstanceNetInfoForChannelResponseBody {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBody) SetRequestId(v string) *DescribeDBInstanceNetInfoForChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBody) Validate() error {
	if s.DBInstanceNetInfos != nil {
		if err := s.DBInstanceNetInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfos struct {
	DBInstanceNetInfo []*DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo `json:"DBInstanceNetInfo,omitempty" xml:"DBInstanceNetInfo,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfos) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfos) GetDBInstanceNetInfo() []*DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	return s.DBInstanceNetInfo
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfos) SetDBInstanceNetInfo(v []*DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfos {
	s.DBInstanceNetInfo = v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfos) Validate() error {
	if s.DBInstanceNetInfo != nil {
		for _, item := range s.DBInstanceNetInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo struct {
	// The availability of the instance. Valid values:
	//
	// 	- **Unavailable**
	//
	// 	- **Available**
	//
	// example:
	//
	// Unavailable
	Availability *string `json:"Availability,omitempty" xml:"Availability,omitempty"`
	// The endpoint of the instance.
	//
	// example:
	//
	// rm-uf6w*****.mysql.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The type of the endpoint. Valid values:
	//
	// 	- **Normal**: a regular endpoint
	//
	// 	- **ReadWriteSplitting**: a read/write splitting endpoint that is assigned after the shared proxy feature is enabled.
	//
	// example:
	//
	// Normal
	ConnectionStringType *string `json:"ConnectionStringType,omitempty" xml:"ConnectionStringType,omitempty"`
	// The information about read weights to implement read/write splitting after the shared proxy feature is enabled.
	DBInstanceWeights *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeights `json:"DBInstanceWeights,omitempty" xml:"DBInstanceWeights,omitempty" type:"Struct"`
	// The policy that is used to assign read weights. This parameter is returned only for a read/write splitting endpoint that is assigned after the shared proxy feature is enabled. Valid values:
	//
	// 	- **Standard**: The system automatically allocates read weights to the instance and its read-only instances based on the specifications of the instances.
	//
	// 	- **Custom**: You must manually allocate read weights to the instance and its read-only instances.
	//
	// example:
	//
	// Standard
	DistributionType *string `json:"DistributionType,omitempty" xml:"DistributionType,omitempty"`
	// The IP address of the instance.
	//
	// example:
	//
	// 172.16.XX.XX
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// The network type of the IP address. Valid values:
	//
	// 	- **Public**: the Internet
	//
	// 	- **Inner**: the classic network
	//
	// 	- **Private**: a virtual private cloud (VPC)
	//
	// example:
	//
	// Inner
	IPType *string `json:"IPType,omitempty" xml:"IPType,omitempty"`
	// The latency threshold that is allowed for read/write splitting of the shared proxy feature. Unit: seconds.
	//
	// >  This parameter is returned only when **ConnectionStringType*	- is set to **ReadWriteSplitting**.
	//
	// example:
	//
	// 12
	MaxDelayTime *string `json:"MaxDelayTime,omitempty" xml:"MaxDelayTime,omitempty"`
	// The port number of the instance.
	//
	// example:
	//
	// 6379
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The details of the IP address whitelist.
	SecurityIPGroups *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroups `json:"SecurityIPGroups,omitempty" xml:"SecurityIPGroups,omitempty" type:"Struct"`
	// An internal parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// Disabled
	Upgradeable *string `json:"Upgradeable,omitempty" xml:"Upgradeable,omitempty"`
	// The VPC ID of the instance.
	//
	// example:
	//
	// vpc-bp1nme44gek34slfc****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the instance.
	//
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The expiration time of the endpoint of the classic network type. Unit: seconds.
	//
	// example:
	//
	// 5183779
	ExpiredTime *string `json:"expiredTime,omitempty" xml:"expiredTime,omitempty"`
}

func (s DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetAvailability() *string {
	return s.Availability
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetConnectionStringType() *string {
	return s.ConnectionStringType
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetDBInstanceWeights() *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeights {
	return s.DBInstanceWeights
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetDistributionType() *string {
	return s.DistributionType
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetIPAddress() *string {
	return s.IPAddress
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetIPType() *string {
	return s.IPType
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetMaxDelayTime() *string {
	return s.MaxDelayTime
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetPort() *string {
	return s.Port
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetSecurityIPGroups() *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroups {
	return s.SecurityIPGroups
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetUpgradeable() *string {
	return s.Upgradeable
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetAvailability(v string) *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.Availability = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetConnectionString(v string) *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetConnectionStringType(v string) *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.ConnectionStringType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetDBInstanceWeights(v *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeights) *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.DBInstanceWeights = v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetDistributionType(v string) *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.DistributionType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetIPAddress(v string) *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.IPAddress = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetIPType(v string) *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.IPType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetMaxDelayTime(v string) *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.MaxDelayTime = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetPort(v string) *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetSecurityIPGroups(v *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroups) *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.SecurityIPGroups = v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetUpgradeable(v string) *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.Upgradeable = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetVPCId(v string) *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.VPCId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetVSwitchId(v string) *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetExpiredTime(v string) *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfo) Validate() error {
	if s.DBInstanceWeights != nil {
		if err := s.DBInstanceWeights.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityIPGroups != nil {
		if err := s.SecurityIPGroups.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeights struct {
	DBInstanceWeight []*DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight `json:"DBInstanceWeight,omitempty" xml:"DBInstanceWeight,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeights) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeights) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeights) GetDBInstanceWeight() []*DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight {
	return s.DBInstanceWeight
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeights) SetDBInstanceWeight(v []*DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight) *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeights {
	s.DBInstanceWeight = v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeights) Validate() error {
	if s.DBInstanceWeight != nil {
		for _, item := range s.DBInstanceWeight {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight struct {
	// The availability of the instance. Valid values:
	//
	// 	- **Unavailable**
	//
	// 	- **Available**
	//
	// example:
	//
	// Unavailable
	Availability *string `json:"Availability,omitempty" xml:"Availability,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5*****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The instance type. Valid values:
	//
	// 	- **Master**: primary instance
	//
	// 	- **Readonly**: read-only instance
	//
	// example:
	//
	// Master
	DBInstanceType *string `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	// The weight of the instance.
	//
	// example:
	//
	// 100
	Weight *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight) GetAvailability() *string {
	return s.Availability
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight) GetDBInstanceType() *string {
	return s.DBInstanceType
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight) GetWeight() *string {
	return s.Weight
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight) SetAvailability(v string) *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight {
	s.Availability = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight) SetDBInstanceId(v string) *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight) SetDBInstanceType(v string) *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight) SetWeight(v string) *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight {
	s.Weight = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroups struct {
	SecurityIPGroup []*DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroupsSecurityIPGroup `json:"securityIPGroup,omitempty" xml:"securityIPGroup,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroups) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroups) GetSecurityIPGroup() []*DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroupsSecurityIPGroup {
	return s.SecurityIPGroup
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroups) SetSecurityIPGroup(v []*DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroupsSecurityIPGroup) *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroups {
	s.SecurityIPGroup = v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroups) Validate() error {
	if s.SecurityIPGroup != nil {
		for _, item := range s.SecurityIPGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroupsSecurityIPGroup struct {
	// The name of the IP address whitelist.
	//
	// example:
	//
	// Default
	SecurityIPGroupName *string `json:"SecurityIPGroupName,omitempty" xml:"SecurityIPGroupName,omitempty"`
	// The IP addresses that is contained in the IP address whitelist.
	//
	// example:
	//
	// 127.0.XX.XX
	SecurityIPs *string `json:"SecurityIPs,omitempty" xml:"SecurityIPs,omitempty"`
}

func (s DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroupsSecurityIPGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroupsSecurityIPGroup) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroupsSecurityIPGroup) GetSecurityIPGroupName() *string {
	return s.SecurityIPGroupName
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroupsSecurityIPGroup) GetSecurityIPs() *string {
	return s.SecurityIPs
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroupsSecurityIPGroup) SetSecurityIPGroupName(v string) *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroupsSecurityIPGroup {
	s.SecurityIPGroupName = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroupsSecurityIPGroup) SetSecurityIPs(v string) *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroupsSecurityIPGroup {
	s.SecurityIPs = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroupsSecurityIPGroup) Validate() error {
	return dara.Validate(s)
}
