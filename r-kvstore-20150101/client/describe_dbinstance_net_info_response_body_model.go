// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceNetInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceNetworkType(v string) *DescribeDBInstanceNetInfoResponseBody
	GetInstanceNetworkType() *string
	SetNetInfoItems(v *DescribeDBInstanceNetInfoResponseBodyNetInfoItems) *DescribeDBInstanceNetInfoResponseBody
	GetNetInfoItems() *DescribeDBInstanceNetInfoResponseBodyNetInfoItems
	SetRequestId(v string) *DescribeDBInstanceNetInfoResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceNetInfoResponseBody struct {
	// The network type. Valid values:
	//
	// 	- **CLASSIC**: The instance runs in a classic network.
	//
	// 	- **VPC**: The instance runs in a virtual private cloud (VPC).
	//
	// example:
	//
	// CLASSIC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The network information about the instance.
	NetInfoItems *DescribeDBInstanceNetInfoResponseBodyNetInfoItems `json:"NetInfoItems,omitempty" xml:"NetInfoItems,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// FC77D4E1-2A7C-4F0B-A4CC-CE0B9C314B9B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBody) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *DescribeDBInstanceNetInfoResponseBody) GetNetInfoItems() *DescribeDBInstanceNetInfoResponseBodyNetInfoItems {
	return s.NetInfoItems
}

func (s *DescribeDBInstanceNetInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetInstanceNetworkType(v string) *DescribeDBInstanceNetInfoResponseBody {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetNetInfoItems(v *DescribeDBInstanceNetInfoResponseBodyNetInfoItems) *DescribeDBInstanceNetInfoResponseBody {
	s.NetInfoItems = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetRequestId(v string) *DescribeDBInstanceNetInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceNetInfoResponseBodyNetInfoItems struct {
	InstanceNetInfo []*DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo `json:"InstanceNetInfo,omitempty" xml:"InstanceNetInfo,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceNetInfoResponseBodyNetInfoItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyNetInfoItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItems) GetInstanceNetInfo() []*DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	return s.InstanceNetInfo
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItems) SetInstanceNetInfo(v []*DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) *DescribeDBInstanceNetInfoResponseBodyNetInfoItems {
	s.InstanceNetInfo = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItems) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo struct {
	// The endpoint of the instance.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****.redis.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- **0**: Internet
	//
	// 	- **1**: classic network
	//
	// 	- **2**: Virtual Private Cloud (VPC)
	//
	// example:
	//
	// 1
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// Indicates whether the address is a private endpoint. Valid values:
	//
	// 	- **0**: The address is not a private endpoint.
	//
	// 	- **1**: The address is a private endpoint.
	//
	// example:
	//
	// 0
	DirectConnection *int32 `json:"DirectConnection,omitempty" xml:"DirectConnection,omitempty"`
	// The expiration time of the classic network endpoint. Unit: seconds.
	//
	// example:
	//
	// 5183779
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 172.16.49.***
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// The network type of the IP address. Valid values:
	//
	// 	- **Public**: Internet
	//
	// 	- **Inner**: classic network
	//
	// 	- **Private**: VPC
	//
	// example:
	//
	// Inner
	IPType *string `json:"IPType,omitempty" xml:"IPType,omitempty"`
	// Indicates whether the address is the endpoint for the secondary zone. Valid values: 1 and 0. A value of 1 indicates that the address is the endpoint for the secondary zone.
	//
	// >  This parameter is returned only after you enable the multi-zone read/write splitting architecture for the instance.
	//
	// example:
	//
	// 1
	IsSlaveProxy *int32 `json:"IsSlaveProxy,omitempty" xml:"IsSlaveProxy,omitempty"`
	// The service port of the instance.
	//
	// example:
	//
	// 6379
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The remaining validity period of the classic network endpoint. Unit: seconds.
	//
	// >  **A value of 0 indicates that the endpoint never expires.
	//
	// example:
	//
	// 0
	Upgradeable *string `json:"Upgradeable,omitempty" xml:"Upgradeable,omitempty"`
	// The ID of the VPC to which the instance belongs.
	//
	// example:
	//
	// vpc-bp1nme44gek34slfc****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// r-bp1ky7j6qc7umk****
	VPCInstanceId *string `json:"VPCInstanceId,omitempty" xml:"VPCInstanceId,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) GetDBInstanceNetType() *string {
	return s.DBInstanceNetType
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) GetDirectConnection() *int32 {
	return s.DirectConnection
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) GetIPAddress() *string {
	return s.IPAddress
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) GetIPType() *string {
	return s.IPType
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) GetIsSlaveProxy() *int32 {
	return s.IsSlaveProxy
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) GetPort() *string {
	return s.Port
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) GetUpgradeable() *string {
	return s.Upgradeable
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) GetVPCInstanceId() *string {
	return s.VPCInstanceId
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetConnectionString(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetDBInstanceNetType(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.DBInstanceNetType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetDirectConnection(v int32) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.DirectConnection = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetExpiredTime(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetIPAddress(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.IPAddress = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetIPType(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.IPType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetIsSlaveProxy(v int32) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.IsSlaveProxy = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetPort(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetUpgradeable(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.Upgradeable = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetVPCId(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.VPCId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetVPCInstanceId(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.VPCInstanceId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetVSwitchId(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) Validate() error {
	return dara.Validate(s)
}
