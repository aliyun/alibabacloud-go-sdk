// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceNetInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceNetInfos(v *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) *DescribeDBInstanceNetInfoResponseBody
	GetDBInstanceNetInfos() *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos
	SetInstanceNetworkType(v string) *DescribeDBInstanceNetInfoResponseBody
	GetInstanceNetworkType() *string
	SetRequestId(v string) *DescribeDBInstanceNetInfoResponseBody
	GetRequestId() *string
	SetSecurityIPMode(v string) *DescribeDBInstanceNetInfoResponseBody
	GetSecurityIPMode() *string
}

type DescribeDBInstanceNetInfoResponseBody struct {
	// The information about the endpoints of the instance.
	DBInstanceNetInfos *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos `json:"DBInstanceNetInfos,omitempty" xml:"DBInstanceNetInfos,omitempty" type:"Struct"`
	// The network type of the instance. Valid values:
	//
	// 	- **Classic**: classic network
	//
	// 	- **VPC**: virtual private cloud (VPC)
	//
	// example:
	//
	// VPC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 777C4593-8053-427B-99E2-105593277CAB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The whitelist mode of the instance. Valid values:
	//
	// 	- **normal**: standard whitelist mode
	//
	// 	- **safety**: enhanced whitelist mode
	//
	// example:
	//
	// safety
	SecurityIPMode *string `json:"SecurityIPMode,omitempty" xml:"SecurityIPMode,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBody) GetDBInstanceNetInfos() *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos {
	return s.DBInstanceNetInfos
}

func (s *DescribeDBInstanceNetInfoResponseBody) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *DescribeDBInstanceNetInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceNetInfoResponseBody) GetSecurityIPMode() *string {
	return s.SecurityIPMode
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetDBInstanceNetInfos(v *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) *DescribeDBInstanceNetInfoResponseBody {
	s.DBInstanceNetInfos = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetInstanceNetworkType(v string) *DescribeDBInstanceNetInfoResponseBody {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetRequestId(v string) *DescribeDBInstanceNetInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetSecurityIPMode(v string) *DescribeDBInstanceNetInfoResponseBody {
	s.SecurityIPMode = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBody) Validate() error {
	if s.DBInstanceNetInfos != nil {
		if err := s.DBInstanceNetInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos struct {
	DBInstanceNetInfo []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo `json:"DBInstanceNetInfo,omitempty" xml:"DBInstanceNetInfo,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) GetDBInstanceNetInfo() []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	return s.DBInstanceNetInfo
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) SetDBInstanceNetInfo(v []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos {
	s.DBInstanceNetInfo = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) Validate() error {
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

type DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo struct {
	// The Tabular Data Stream (TDS) port of the instance for which Babelfish is enabled.
	//
	// >  This parameter applies only to ApsaraDB RDS for PostgreSQL instances. For more information about Babelfish for ApsaraDB RDS for PostgreSQL, see [Introduction to Babelfish](https://help.aliyun.com/document_detail/428613.html).
	//
	// example:
	//
	// 1433
	BabelfishPort *string `json:"BabelfishPort,omitempty" xml:"BabelfishPort,omitempty"`
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
	// 	- **ReadWriteSplitting**: a read/write splitting endpoint
	//
	// example:
	//
	// Normal
	ConnectionStringType *string `json:"ConnectionStringType,omitempty" xml:"ConnectionStringType,omitempty"`
	// The information about the instance weight.
	//
	// >  This parameter is returned only when the read/write splitting feature is enabled for the instance.
	DBInstanceWeights *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeights `json:"DBInstanceWeights,omitempty" xml:"DBInstanceWeights,omitempty" type:"Struct"`
	// The policy that is used to assign read weights. This parameter is returned only for a read/write splitting endpoint. Valid values:
	//
	// 	- **Standard**: The system automatically allocates read weights to the instance and its read-only instances based on the specifications of the instances.
	//
	// 	- **Custom**: You must manually allocate read weights to the instance and its read-only instances.
	//
	// example:
	//
	// Standard
	DistributionType *string `json:"DistributionType,omitempty" xml:"DistributionType,omitempty"`
	// The remaining validity period of the instance in the classic network in hybrid access mode. Unit: seconds.
	//
	// example:
	//
	// 1209534
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 192.168.XX.XX
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// The network type.
	//
	// 	- Valid values when the instance resides in the classic network:
	//
	//     	- **Inner**
	//
	//     	- **Public**
	//
	// 	- Valid values when the instance resides in a virtual private cloud (VPC):
	//
	//     	- **Private**
	//
	//     	- **Public**
	//
	// example:
	//
	// Public
	IPType *string `json:"IPType,omitempty" xml:"IPType,omitempty"`
	// The latency threshold. This parameter is returned only for a read/write splitting endpoint. Unit: seconds.
	//
	// >  If the latency on a read-only instance exceeds the specified threshold, ApsaraDB RDS no longer forwards read requests to the read-only instance.
	//
	// example:
	//
	// 12
	MaxDelayTime *string `json:"MaxDelayTime,omitempty" xml:"MaxDelayTime,omitempty"`
	// The PgBouncer port.
	//
	// >  This parameter is returned only when PgBouncer is enabled for the instance that runs PostgreSQL.
	//
	// example:
	//
	// 6432
	PGBouncerPort *string `json:"PGBouncerPort,omitempty" xml:"PGBouncerPort,omitempty"`
	// The port that is used to connect to the instance.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The IP addresses in the whitelist for the instance.
	SecurityIPGroups *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroups `json:"SecurityIPGroups,omitempty" xml:"SecurityIPGroups,omitempty" type:"Struct"`
	// Indicates whether the IP version can be updated. Valid values:
	//
	// 	- **Enable**
	//
	// 	- **Disabled**
	//
	// >  The IP version can be updated from IPv4 to IPv6.
	//
	// example:
	//
	// Disabled
	Upgradeable *string `json:"Upgradeable,omitempty" xml:"Upgradeable,omitempty"`
	// The VPC ID of the instance.
	//
	// example:
	//
	// vpc-uf6f7l4fg90*****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-uf6adz52c2p*****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetBabelfishPort() *string {
	return s.BabelfishPort
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetConnectionStringType() *string {
	return s.ConnectionStringType
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetDBInstanceWeights() *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeights {
	return s.DBInstanceWeights
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetDistributionType() *string {
	return s.DistributionType
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetIPAddress() *string {
	return s.IPAddress
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetIPType() *string {
	return s.IPType
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetMaxDelayTime() *string {
	return s.MaxDelayTime
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetPGBouncerPort() *string {
	return s.PGBouncerPort
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetPort() *string {
	return s.Port
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetSecurityIPGroups() *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroups {
	return s.SecurityIPGroups
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetUpgradeable() *string {
	return s.Upgradeable
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetBabelfishPort(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.BabelfishPort = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetConnectionString(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetConnectionStringType(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.ConnectionStringType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetDBInstanceWeights(v *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeights) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.DBInstanceWeights = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetDistributionType(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.DistributionType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetExpiredTime(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetIPAddress(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.IPAddress = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetIPType(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.IPType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetMaxDelayTime(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.MaxDelayTime = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetPGBouncerPort(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.PGBouncerPort = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetPort(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetSecurityIPGroups(v *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroups) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.SecurityIPGroups = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetUpgradeable(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.Upgradeable = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetVPCId(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.VPCId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetVSwitchId(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) Validate() error {
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

type DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeights struct {
	DBInstanceWeight []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight `json:"DBInstanceWeight,omitempty" xml:"DBInstanceWeight,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeights) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeights) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeights) GetDBInstanceWeight() []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight {
	return s.DBInstanceWeight
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeights) SetDBInstanceWeight(v []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeights {
	s.DBInstanceWeight = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeights) Validate() error {
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

type DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight struct {
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
	// The type of the instance. Valid values:
	//
	// 	- **Master**: primary instance
	//
	// 	- **Readonly**: read-only instance
	//
	// example:
	//
	// Master
	DBInstanceType *string `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	// A deprecated parameter.
	//
	// example:
	//
	// None
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The weight of the instance.
	//
	// example:
	//
	// 100
	Weight *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight) GetAvailability() *string {
	return s.Availability
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight) GetDBInstanceType() *string {
	return s.DBInstanceType
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight) GetRole() *string {
	return s.Role
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight) GetWeight() *string {
	return s.Weight
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight) SetAvailability(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight {
	s.Availability = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight) SetDBInstanceId(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight) SetDBInstanceType(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight) SetRole(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight {
	s.Role = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight) SetWeight(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight {
	s.Weight = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeightsDBInstanceWeight) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroups struct {
	SecurityIPGroup []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroupsSecurityIPGroup `json:"securityIPGroup,omitempty" xml:"securityIPGroup,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroups) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroups) GetSecurityIPGroup() []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroupsSecurityIPGroup {
	return s.SecurityIPGroup
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroups) SetSecurityIPGroup(v []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroupsSecurityIPGroup) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroups {
	s.SecurityIPGroup = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroups) Validate() error {
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

type DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroupsSecurityIPGroup struct {
	// The name of the IP address whitelist.
	//
	// example:
	//
	// Default
	SecurityIPGroupName *string `json:"SecurityIPGroupName,omitempty" xml:"SecurityIPGroupName,omitempty"`
	// The IP address in the whitelist.
	//
	// example:
	//
	// 127.0.XX.XX
	SecurityIPs *string `json:"SecurityIPs,omitempty" xml:"SecurityIPs,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroupsSecurityIPGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroupsSecurityIPGroup) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroupsSecurityIPGroup) GetSecurityIPGroupName() *string {
	return s.SecurityIPGroupName
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroupsSecurityIPGroup) GetSecurityIPs() *string {
	return s.SecurityIPs
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroupsSecurityIPGroup) SetSecurityIPGroupName(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroupsSecurityIPGroup {
	s.SecurityIPGroupName = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroupsSecurityIPGroup) SetSecurityIPs(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroupsSecurityIPGroup {
	s.SecurityIPs = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroupsSecurityIPGroup) Validate() error {
	return dara.Validate(s)
}
