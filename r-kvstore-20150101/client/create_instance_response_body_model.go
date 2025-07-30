// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int64) *CreateInstanceResponseBody
	GetBandwidth() *int64
	SetCapacity(v int64) *CreateInstanceResponseBody
	GetCapacity() *int64
	SetChargeType(v string) *CreateInstanceResponseBody
	GetChargeType() *string
	SetConfig(v string) *CreateInstanceResponseBody
	GetConfig() *string
	SetConnectionDomain(v string) *CreateInstanceResponseBody
	GetConnectionDomain() *string
	SetConnections(v int64) *CreateInstanceResponseBody
	GetConnections() *int64
	SetEndTime(v string) *CreateInstanceResponseBody
	GetEndTime() *string
	SetInstanceId(v string) *CreateInstanceResponseBody
	GetInstanceId() *string
	SetInstanceName(v string) *CreateInstanceResponseBody
	GetInstanceName() *string
	SetInstanceStatus(v string) *CreateInstanceResponseBody
	GetInstanceStatus() *string
	SetNetworkType(v string) *CreateInstanceResponseBody
	GetNetworkType() *string
	SetNodeType(v string) *CreateInstanceResponseBody
	GetNodeType() *string
	SetOrderId(v int64) *CreateInstanceResponseBody
	GetOrderId() *int64
	SetPort(v int32) *CreateInstanceResponseBody
	GetPort() *int32
	SetPrivateIpAddr(v string) *CreateInstanceResponseBody
	GetPrivateIpAddr() *string
	SetQPS(v int64) *CreateInstanceResponseBody
	GetQPS() *int64
	SetRegionId(v string) *CreateInstanceResponseBody
	GetRegionId() *string
	SetRequestId(v string) *CreateInstanceResponseBody
	GetRequestId() *string
	SetUserName(v string) *CreateInstanceResponseBody
	GetUserName() *string
	SetVSwitchId(v string) *CreateInstanceResponseBody
	GetVSwitchId() *string
	SetVpcId(v string) *CreateInstanceResponseBody
	GetVpcId() *string
	SetZoneId(v string) *CreateInstanceResponseBody
	GetZoneId() *string
}

type CreateInstanceResponseBody struct {
	// The maximum bandwidth of the instance. Unit: MB/s.
	//
	// example:
	//
	// 32
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The storage capacity of the instance. Unit: MB.
	//
	// example:
	//
	// 16384
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **PrePaid**: subscription
	//
	// 	- **PostPaid**: pay-as-you-go
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The configurations of the instance.
	//
	// example:
	//
	// {\\"EvictionPolicy\\":\\"volatile-lru\\",\\"hash-max-ziplist-entries\\":512,\\"zset-max-ziplist-entries\\":128,\\"zset-max-ziplist-value\\":64,\\"set-max-intset-entries\\":512,\\"hash-max-ziplist-value\\":64}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The internal endpoint of the instance.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****.redis.rds.aliyuncs.com
	ConnectionDomain *string `json:"ConnectionDomain,omitempty" xml:"ConnectionDomain,omitempty"`
	// The maximum number of connections supported by the instance.
	//
	// example:
	//
	// 10000
	Connections *int64 `json:"Connections,omitempty" xml:"Connections,omitempty"`
	// The time when the subscription expires. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-01-18T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The GUID of the instance.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// apitest
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The state of the instance. The return value is Creating.
	//
	// example:
	//
	// Creating
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- **CLASSIC**: classic network
	//
	// 	- **VPC**: VPC
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The node type. Valid values:
	//
	// 	- **STAND_ALONE**: standalone
	//
	// 	- **MASTER_SLAVE**: master-replica
	//
	// example:
	//
	// MASTER_SLAVE
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The ID of the order.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2084452111111
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The port number that is used to connect to the instance.
	//
	// example:
	//
	// 6379
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The private IP address of the instance.
	//
	// example:
	//
	// 172.16.0.10
	PrivateIpAddr *string `json:"PrivateIpAddr,omitempty" xml:"PrivateIpAddr,omitempty"`
	// The expected maximum queries per second (QPS).
	//
	// example:
	//
	// 100000
	QPS *int64 `json:"QPS,omitempty" xml:"QPS,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hongkong
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5DEA3CC9-F81D-4387-8E97-CEA40F09****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The username that is used to connect to the instance. By default, Tair (Redis OSS-compatible) provides a username that is named after the instance ID.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The ID of the vSwitch to which the instance is connected.
	//
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-bp1nme44gek34slfc****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID of the instance.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *CreateInstanceResponseBody) GetCapacity() *int64 {
	return s.Capacity
}

func (s *CreateInstanceResponseBody) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateInstanceResponseBody) GetConfig() *string {
	return s.Config
}

func (s *CreateInstanceResponseBody) GetConnectionDomain() *string {
	return s.ConnectionDomain
}

func (s *CreateInstanceResponseBody) GetConnections() *int64 {
	return s.Connections
}

func (s *CreateInstanceResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInstanceResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateInstanceResponseBody) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *CreateInstanceResponseBody) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateInstanceResponseBody) GetNodeType() *string {
	return s.NodeType
}

func (s *CreateInstanceResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateInstanceResponseBody) GetPort() *int32 {
	return s.Port
}

func (s *CreateInstanceResponseBody) GetPrivateIpAddr() *string {
	return s.PrivateIpAddr
}

func (s *CreateInstanceResponseBody) GetQPS() *int64 {
	return s.QPS
}

func (s *CreateInstanceResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceResponseBody) GetUserName() *string {
	return s.UserName
}

func (s *CreateInstanceResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateInstanceResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateInstanceResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateInstanceResponseBody) SetBandwidth(v int64) *CreateInstanceResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *CreateInstanceResponseBody) SetCapacity(v int64) *CreateInstanceResponseBody {
	s.Capacity = &v
	return s
}

func (s *CreateInstanceResponseBody) SetChargeType(v string) *CreateInstanceResponseBody {
	s.ChargeType = &v
	return s
}

func (s *CreateInstanceResponseBody) SetConfig(v string) *CreateInstanceResponseBody {
	s.Config = &v
	return s
}

func (s *CreateInstanceResponseBody) SetConnectionDomain(v string) *CreateInstanceResponseBody {
	s.ConnectionDomain = &v
	return s
}

func (s *CreateInstanceResponseBody) SetConnections(v int64) *CreateInstanceResponseBody {
	s.Connections = &v
	return s
}

func (s *CreateInstanceResponseBody) SetEndTime(v string) *CreateInstanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *CreateInstanceResponseBody) SetInstanceId(v string) *CreateInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetInstanceName(v string) *CreateInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceResponseBody) SetInstanceStatus(v string) *CreateInstanceResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *CreateInstanceResponseBody) SetNetworkType(v string) *CreateInstanceResponseBody {
	s.NetworkType = &v
	return s
}

func (s *CreateInstanceResponseBody) SetNodeType(v string) *CreateInstanceResponseBody {
	s.NodeType = &v
	return s
}

func (s *CreateInstanceResponseBody) SetOrderId(v int64) *CreateInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetPort(v int32) *CreateInstanceResponseBody {
	s.Port = &v
	return s
}

func (s *CreateInstanceResponseBody) SetPrivateIpAddr(v string) *CreateInstanceResponseBody {
	s.PrivateIpAddr = &v
	return s
}

func (s *CreateInstanceResponseBody) SetQPS(v int64) *CreateInstanceResponseBody {
	s.QPS = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRegionId(v string) *CreateInstanceResponseBody {
	s.RegionId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetUserName(v string) *CreateInstanceResponseBody {
	s.UserName = &v
	return s
}

func (s *CreateInstanceResponseBody) SetVSwitchId(v string) *CreateInstanceResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetVpcId(v string) *CreateInstanceResponseBody {
	s.VpcId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetZoneId(v string) *CreateInstanceResponseBody {
	s.ZoneId = &v
	return s
}

func (s *CreateInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
