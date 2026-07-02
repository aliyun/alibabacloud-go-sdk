// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTairSkvDdbTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int64) *CreateTairSkvDdbTableResponseBody
	GetBandwidth() *int64
	SetChargeType(v string) *CreateTairSkvDdbTableResponseBody
	GetChargeType() *string
	SetConfig(v string) *CreateTairSkvDdbTableResponseBody
	GetConfig() *string
	SetConnectionDomain(v string) *CreateTairSkvDdbTableResponseBody
	GetConnectionDomain() *string
	SetConnections(v int64) *CreateTairSkvDdbTableResponseBody
	GetConnections() *int64
	SetInstanceId(v string) *CreateTairSkvDdbTableResponseBody
	GetInstanceId() *string
	SetInstanceStatus(v string) *CreateTairSkvDdbTableResponseBody
	GetInstanceStatus() *string
	SetOrderId(v int64) *CreateTairSkvDdbTableResponseBody
	GetOrderId() *int64
	SetPort(v int32) *CreateTairSkvDdbTableResponseBody
	GetPort() *int32
	SetQPS(v int64) *CreateTairSkvDdbTableResponseBody
	GetQPS() *int64
	SetRegionId(v string) *CreateTairSkvDdbTableResponseBody
	GetRegionId() *string
	SetRequestId(v string) *CreateTairSkvDdbTableResponseBody
	GetRequestId() *string
	SetTableName(v string) *CreateTairSkvDdbTableResponseBody
	GetTableName() *string
	SetTaskId(v string) *CreateTairSkvDdbTableResponseBody
	GetTaskId() *string
	SetZoneId(v string) *CreateTairSkvDdbTableResponseBody
	GetZoneId() *string
}

type CreateTairSkvDdbTableResponseBody struct {
	// The bandwidth limit of the instance. Unit: MB/s.
	//
	// example:
	//
	// 96
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The billing method. Valid values:
	//
	// - PrePaid: subscription.
	//
	// - PostPaid: pay-as-you-go.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The detailed configuration of the instance in JSON format. For more information about the parameters, see [Parameter settings](https://help.aliyun.com/document_detail/43885.html).
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
	// r-bp1zxszhcgatnx*
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The current status of the instance. The value is fixed to Creating.
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
	// 6379
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The queries per second (QPS). This value is the theoretical value for the current instance specification.
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
	// 2D9F3768-EDA9-4811-943E-42C8006E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// apitest
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the task flow that is executed for the creation.
	//
	// example:
	//
	// 11111****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The zone ID of the instance.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateTairSkvDdbTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTairSkvDdbTableResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTairSkvDdbTableResponseBody) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *CreateTairSkvDdbTableResponseBody) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateTairSkvDdbTableResponseBody) GetConfig() *string {
	return s.Config
}

func (s *CreateTairSkvDdbTableResponseBody) GetConnectionDomain() *string {
	return s.ConnectionDomain
}

func (s *CreateTairSkvDdbTableResponseBody) GetConnections() *int64 {
	return s.Connections
}

func (s *CreateTairSkvDdbTableResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateTairSkvDdbTableResponseBody) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *CreateTairSkvDdbTableResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateTairSkvDdbTableResponseBody) GetPort() *int32 {
	return s.Port
}

func (s *CreateTairSkvDdbTableResponseBody) GetQPS() *int64 {
	return s.QPS
}

func (s *CreateTairSkvDdbTableResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTairSkvDdbTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTairSkvDdbTableResponseBody) GetTableName() *string {
	return s.TableName
}

func (s *CreateTairSkvDdbTableResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateTairSkvDdbTableResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateTairSkvDdbTableResponseBody) SetBandwidth(v int64) *CreateTairSkvDdbTableResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *CreateTairSkvDdbTableResponseBody) SetChargeType(v string) *CreateTairSkvDdbTableResponseBody {
	s.ChargeType = &v
	return s
}

func (s *CreateTairSkvDdbTableResponseBody) SetConfig(v string) *CreateTairSkvDdbTableResponseBody {
	s.Config = &v
	return s
}

func (s *CreateTairSkvDdbTableResponseBody) SetConnectionDomain(v string) *CreateTairSkvDdbTableResponseBody {
	s.ConnectionDomain = &v
	return s
}

func (s *CreateTairSkvDdbTableResponseBody) SetConnections(v int64) *CreateTairSkvDdbTableResponseBody {
	s.Connections = &v
	return s
}

func (s *CreateTairSkvDdbTableResponseBody) SetInstanceId(v string) *CreateTairSkvDdbTableResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateTairSkvDdbTableResponseBody) SetInstanceStatus(v string) *CreateTairSkvDdbTableResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *CreateTairSkvDdbTableResponseBody) SetOrderId(v int64) *CreateTairSkvDdbTableResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateTairSkvDdbTableResponseBody) SetPort(v int32) *CreateTairSkvDdbTableResponseBody {
	s.Port = &v
	return s
}

func (s *CreateTairSkvDdbTableResponseBody) SetQPS(v int64) *CreateTairSkvDdbTableResponseBody {
	s.QPS = &v
	return s
}

func (s *CreateTairSkvDdbTableResponseBody) SetRegionId(v string) *CreateTairSkvDdbTableResponseBody {
	s.RegionId = &v
	return s
}

func (s *CreateTairSkvDdbTableResponseBody) SetRequestId(v string) *CreateTairSkvDdbTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTairSkvDdbTableResponseBody) SetTableName(v string) *CreateTairSkvDdbTableResponseBody {
	s.TableName = &v
	return s
}

func (s *CreateTairSkvDdbTableResponseBody) SetTaskId(v string) *CreateTairSkvDdbTableResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateTairSkvDdbTableResponseBody) SetZoneId(v string) *CreateTairSkvDdbTableResponseBody {
	s.ZoneId = &v
	return s
}

func (s *CreateTairSkvDdbTableResponseBody) Validate() error {
	return dara.Validate(s)
}
