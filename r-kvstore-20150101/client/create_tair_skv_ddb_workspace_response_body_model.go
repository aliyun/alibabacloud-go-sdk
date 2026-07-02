// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTairSkvDdbWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int64) *CreateTairSkvDdbWorkspaceResponseBody
	GetBandwidth() *int64
	SetChargeType(v string) *CreateTairSkvDdbWorkspaceResponseBody
	GetChargeType() *string
	SetConfig(v string) *CreateTairSkvDdbWorkspaceResponseBody
	GetConfig() *string
	SetConnectionDomain(v string) *CreateTairSkvDdbWorkspaceResponseBody
	GetConnectionDomain() *string
	SetConnections(v int64) *CreateTairSkvDdbWorkspaceResponseBody
	GetConnections() *int64
	SetInstanceId(v string) *CreateTairSkvDdbWorkspaceResponseBody
	GetInstanceId() *string
	SetInstanceName(v string) *CreateTairSkvDdbWorkspaceResponseBody
	GetInstanceName() *string
	SetInstanceStatus(v string) *CreateTairSkvDdbWorkspaceResponseBody
	GetInstanceStatus() *string
	SetOrderId(v int64) *CreateTairSkvDdbWorkspaceResponseBody
	GetOrderId() *int64
	SetPort(v int32) *CreateTairSkvDdbWorkspaceResponseBody
	GetPort() *int32
	SetQPS(v int64) *CreateTairSkvDdbWorkspaceResponseBody
	GetQPS() *int64
	SetRegionId(v string) *CreateTairSkvDdbWorkspaceResponseBody
	GetRegionId() *string
	SetRequestId(v string) *CreateTairSkvDdbWorkspaceResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateTairSkvDdbWorkspaceResponseBody
	GetTaskId() *string
	SetZoneId(v string) *CreateTairSkvDdbWorkspaceResponseBody
	GetZoneId() *string
}

type CreateTairSkvDdbWorkspaceResponseBody struct {
	// The bandwidth limit of the instance. Unit: MB/s.
	//
	// example:
	//
	// 96
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The billing method. Valid values:
	//
	// 	- `PrePaid`: subscription.
	//
	// 	- `PostPaid`: pay-as-you-go.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The detailed configuration of the instance.
	//
	// example:
	//
	// {\\"EvictionPolicy\\":\\"volatile-lru\\",\\"hash-max-ziplist-entries\\":512,\\"zset-max-ziplist-entries\\":128,\\"zset-max-ziplist-value\\":64,\\"set-max-intset-entries\\":512,\\"hash-max-ziplist-value\\":64}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The internal endpoint of the Redis instance.
	//
	// example:
	//
	// r-bp1zxszhcgatnx**.redis.rds.aliyuncs.com
	ConnectionDomain *string `json:"ConnectionDomain,omitempty" xml:"ConnectionDomain,omitempty"`
	// The maximum number of connections supported by the instance.
	//
	// example:
	//
	// 10000
	Connections *int64 `json:"Connections,omitempty" xml:"Connections,omitempty"`
	// The globally unique instance ID.
	//
	// example:
	//
	// r-bp1zxszhcgatnx**
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// apitest
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The current status of the instance. The return value is fixed as Creating.
	//
	// example:
	//
	// Creating
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 20741011111111
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The Redis service port.
	//
	// example:
	//
	// 443
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The queries per second (QPS). This value is the theoretical value for the current instance specification.
	//
	// example:
	//
	// 100000
	QPS *int64 `json:"QPS,omitempty" xml:"QPS,omitempty"`
	// The region in which the instance resides.
	//
	// example:
	//
	// cn-hongkong
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5DEA3CC9-F81D-4387-8E97-CEA40F09****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task flow that is executed for the creation.
	//
	// example:
	//
	// 1111
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The zone ID of the instance.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateTairSkvDdbWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTairSkvDdbWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) GetConfig() *string {
	return s.Config
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) GetConnectionDomain() *string {
	return s.ConnectionDomain
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) GetConnections() *int64 {
	return s.Connections
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) GetPort() *int32 {
	return s.Port
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) GetQPS() *int64 {
	return s.QPS
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) SetBandwidth(v int64) *CreateTairSkvDdbWorkspaceResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) SetChargeType(v string) *CreateTairSkvDdbWorkspaceResponseBody {
	s.ChargeType = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) SetConfig(v string) *CreateTairSkvDdbWorkspaceResponseBody {
	s.Config = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) SetConnectionDomain(v string) *CreateTairSkvDdbWorkspaceResponseBody {
	s.ConnectionDomain = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) SetConnections(v int64) *CreateTairSkvDdbWorkspaceResponseBody {
	s.Connections = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) SetInstanceId(v string) *CreateTairSkvDdbWorkspaceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) SetInstanceName(v string) *CreateTairSkvDdbWorkspaceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) SetInstanceStatus(v string) *CreateTairSkvDdbWorkspaceResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) SetOrderId(v int64) *CreateTairSkvDdbWorkspaceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) SetPort(v int32) *CreateTairSkvDdbWorkspaceResponseBody {
	s.Port = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) SetQPS(v int64) *CreateTairSkvDdbWorkspaceResponseBody {
	s.QPS = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) SetRegionId(v string) *CreateTairSkvDdbWorkspaceResponseBody {
	s.RegionId = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) SetRequestId(v string) *CreateTairSkvDdbWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) SetTaskId(v string) *CreateTairSkvDdbWorkspaceResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) SetZoneId(v string) *CreateTairSkvDdbWorkspaceResponseBody {
	s.ZoneId = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}
