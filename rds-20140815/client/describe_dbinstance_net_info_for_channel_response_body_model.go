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
	Availability         *string                                                                                              `json:"Availability,omitempty" xml:"Availability,omitempty"`
	ConnectionString     *string                                                                                              `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	ConnectionStringType *string                                                                                              `json:"ConnectionStringType,omitempty" xml:"ConnectionStringType,omitempty"`
	DBInstanceWeights    *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoDBInstanceWeights `json:"DBInstanceWeights,omitempty" xml:"DBInstanceWeights,omitempty" type:"Struct"`
	DistributionType     *string                                                                                              `json:"DistributionType,omitempty" xml:"DistributionType,omitempty"`
	IPAddress            *string                                                                                              `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	IPType               *string                                                                                              `json:"IPType,omitempty" xml:"IPType,omitempty"`
	MaxDelayTime         *string                                                                                              `json:"MaxDelayTime,omitempty" xml:"MaxDelayTime,omitempty"`
	Port                 *string                                                                                              `json:"Port,omitempty" xml:"Port,omitempty"`
	SecurityIPGroups     *DescribeDBInstanceNetInfoForChannelResponseBodyDBInstanceNetInfosDBInstanceNetInfoSecurityIPGroups  `json:"SecurityIPGroups,omitempty" xml:"SecurityIPGroups,omitempty" type:"Struct"`
	Upgradeable          *string                                                                                              `json:"Upgradeable,omitempty" xml:"Upgradeable,omitempty"`
	VPCId                *string                                                                                              `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId            *string                                                                                              `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ExpiredTime          *string                                                                                              `json:"expiredTime,omitempty" xml:"expiredTime,omitempty"`
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
	Availability   *string `json:"Availability,omitempty" xml:"Availability,omitempty"`
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceType *string `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	Weight         *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
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
	SecurityIPGroupName *string `json:"SecurityIPGroupName,omitempty" xml:"SecurityIPGroupName,omitempty"`
	SecurityIPs         *string `json:"SecurityIPs,omitempty" xml:"SecurityIPs,omitempty"`
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
