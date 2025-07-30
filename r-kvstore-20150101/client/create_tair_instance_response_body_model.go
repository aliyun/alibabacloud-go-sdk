// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTairInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int64) *CreateTairInstanceResponseBody
	GetBandwidth() *int64
	SetChargeType(v string) *CreateTairInstanceResponseBody
	GetChargeType() *string
	SetConfig(v string) *CreateTairInstanceResponseBody
	GetConfig() *string
	SetConnectionDomain(v string) *CreateTairInstanceResponseBody
	GetConnectionDomain() *string
	SetConnections(v int64) *CreateTairInstanceResponseBody
	GetConnections() *int64
	SetInstanceId(v string) *CreateTairInstanceResponseBody
	GetInstanceId() *string
	SetInstanceName(v string) *CreateTairInstanceResponseBody
	GetInstanceName() *string
	SetInstanceStatus(v string) *CreateTairInstanceResponseBody
	GetInstanceStatus() *string
	SetOrderId(v int64) *CreateTairInstanceResponseBody
	GetOrderId() *int64
	SetPort(v int32) *CreateTairInstanceResponseBody
	GetPort() *int32
	SetQPS(v int64) *CreateTairInstanceResponseBody
	GetQPS() *int64
	SetRegionId(v string) *CreateTairInstanceResponseBody
	GetRegionId() *string
	SetRequestId(v string) *CreateTairInstanceResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateTairInstanceResponseBody
	GetTaskId() *string
	SetZoneId(v string) *CreateTairInstanceResponseBody
	GetZoneId() *string
}

type CreateTairInstanceResponseBody struct {
	// The maximum bandwidth of the instance. Unit: Mbit/s.
	//
	// example:
	//
	// 96
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **PrePaid**: subscription
	//
	// 	- **PostPaid**: pay-as-you-go
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The detailed configurations of the instance. The value is a JSON string. For more information about the parameters, see [Configure parameters](https://help.aliyun.com/document_detail/43885.html).
	//
	// example:
	//
	// {\\"EvictionPolicy\\":\\"volatile-lru\\",\\"hash-max-ziplist-entries\\":512,\\"zset-max-ziplist-entries\\":128,\\"list-max-ziplist-entries\\":512,\\"list-max-ziplist-value\\":64,\\"zset-max-ziplist-value\\":64,\\"set-max-intset-entries\\":512,\\"hash-max-ziplist-value\\":64}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The internal endpoint of the instance.
	//
	// example:
	//
	// r-bp13ac3d047b****.tairpena.rds.aliyuncs.com
	ConnectionDomain *string `json:"ConnectionDomain,omitempty" xml:"ConnectionDomain,omitempty"`
	// The maximum number of connections to the instance.
	//
	// example:
	//
	// 10000
	Connections *int64 `json:"Connections,omitempty" xml:"Connections,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// r-bp13ac3d047b****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// >  This parameter is returned only if the **InstanceName*	- parameter is specified in the request.
	//
	// example:
	//
	// redistest
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The current status of the instance. The value is **Creating**.
	//
	// example:
	//
	// Creating
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 2084452111111
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The service port number of the instance.
	//
	// example:
	//
	// 6379
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The maximum number of read and write operations that can be processed by the instance per second. The value is a theoretical value.
	//
	// example:
	//
	// 100000
	QPS *int64 `json:"QPS,omitempty" xml:"QPS,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 12123216-4B00-4378-BE4B-08005BFC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 10****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateTairInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTairInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTairInstanceResponseBody) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *CreateTairInstanceResponseBody) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateTairInstanceResponseBody) GetConfig() *string {
	return s.Config
}

func (s *CreateTairInstanceResponseBody) GetConnectionDomain() *string {
	return s.ConnectionDomain
}

func (s *CreateTairInstanceResponseBody) GetConnections() *int64 {
	return s.Connections
}

func (s *CreateTairInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateTairInstanceResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateTairInstanceResponseBody) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *CreateTairInstanceResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateTairInstanceResponseBody) GetPort() *int32 {
	return s.Port
}

func (s *CreateTairInstanceResponseBody) GetQPS() *int64 {
	return s.QPS
}

func (s *CreateTairInstanceResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTairInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTairInstanceResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateTairInstanceResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateTairInstanceResponseBody) SetBandwidth(v int64) *CreateTairInstanceResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *CreateTairInstanceResponseBody) SetChargeType(v string) *CreateTairInstanceResponseBody {
	s.ChargeType = &v
	return s
}

func (s *CreateTairInstanceResponseBody) SetConfig(v string) *CreateTairInstanceResponseBody {
	s.Config = &v
	return s
}

func (s *CreateTairInstanceResponseBody) SetConnectionDomain(v string) *CreateTairInstanceResponseBody {
	s.ConnectionDomain = &v
	return s
}

func (s *CreateTairInstanceResponseBody) SetConnections(v int64) *CreateTairInstanceResponseBody {
	s.Connections = &v
	return s
}

func (s *CreateTairInstanceResponseBody) SetInstanceId(v string) *CreateTairInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateTairInstanceResponseBody) SetInstanceName(v string) *CreateTairInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *CreateTairInstanceResponseBody) SetInstanceStatus(v string) *CreateTairInstanceResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *CreateTairInstanceResponseBody) SetOrderId(v int64) *CreateTairInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateTairInstanceResponseBody) SetPort(v int32) *CreateTairInstanceResponseBody {
	s.Port = &v
	return s
}

func (s *CreateTairInstanceResponseBody) SetQPS(v int64) *CreateTairInstanceResponseBody {
	s.QPS = &v
	return s
}

func (s *CreateTairInstanceResponseBody) SetRegionId(v string) *CreateTairInstanceResponseBody {
	s.RegionId = &v
	return s
}

func (s *CreateTairInstanceResponseBody) SetRequestId(v string) *CreateTairInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTairInstanceResponseBody) SetTaskId(v string) *CreateTairInstanceResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateTairInstanceResponseBody) SetZoneId(v string) *CreateTairInstanceResponseBody {
	s.ZoneId = &v
	return s
}

func (s *CreateTairInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
