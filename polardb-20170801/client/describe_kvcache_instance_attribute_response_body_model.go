// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKVCacheInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionString(v string) *DescribeKVCacheInstanceAttributeResponseBody
	GetConnectionString() *string
	SetCreationTime(v string) *DescribeKVCacheInstanceAttributeResponseBody
	GetCreationTime() *string
	SetExpireTime(v string) *DescribeKVCacheInstanceAttributeResponseBody
	GetExpireTime() *string
	SetExpired(v string) *DescribeKVCacheInstanceAttributeResponseBody
	GetExpired() *string
	SetKVCacheInstanceId(v string) *DescribeKVCacheInstanceAttributeResponseBody
	GetKVCacheInstanceId() *string
	SetKVCacheInstanceStatus(v string) *DescribeKVCacheInstanceAttributeResponseBody
	GetKVCacheInstanceStatus() *string
	SetKVCacheSpace(v string) *DescribeKVCacheInstanceAttributeResponseBody
	GetKVCacheSpace() *string
	SetPayType(v string) *DescribeKVCacheInstanceAttributeResponseBody
	GetPayType() *string
	SetPort(v string) *DescribeKVCacheInstanceAttributeResponseBody
	GetPort() *string
	SetRegionId(v string) *DescribeKVCacheInstanceAttributeResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeKVCacheInstanceAttributeResponseBody
	GetRequestId() *string
	SetVPCId(v string) *DescribeKVCacheInstanceAttributeResponseBody
	GetVPCId() *string
	SetVSwitchId(v string) *DescribeKVCacheInstanceAttributeResponseBody
	GetVSwitchId() *string
	SetZoneId(v string) *DescribeKVCacheInstanceAttributeResponseBody
	GetZoneId() *string
}

type DescribeKVCacheInstanceAttributeResponseBody struct {
	// The endpoint of the instance.
	//
	// example:
	//
	// 172.168.0.1
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The time when the instance was created, in ISO 8601 format.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The time when the instance expires, in ISO 8601 format. An empty string is returned for pay-as-you-go instances.
	//
	// example:
	//
	// 2026-12-31T00:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the instance has expired.
	//
	// example:
	//
	// false
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The KVCache instance ID.
	//
	// example:
	//
	// pkv-123456
	KVCacheInstanceId *string `json:"KVCacheInstanceId,omitempty" xml:"KVCacheInstanceId,omitempty"`
	// The KVCache instance status.
	//
	// example:
	//
	// Running
	KVCacheInstanceStatus *string `json:"KVCacheInstanceStatus,omitempty" xml:"KVCacheInstanceStatus,omitempty"`
	// The storage capacity of the KVCache instance.
	//
	// example:
	//
	// 100
	KVCacheSpace *string `json:"KVCacheSpace,omitempty" xml:"KVCacheSpace,omitempty"`
	// The billing method.
	//
	// example:
	//
	// PREPAID
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The port.
	//
	// example:
	//
	// 5432
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-123
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-123
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-beijing-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeKVCacheInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKVCacheInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) GetExpired() *string {
	return s.Expired
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) GetKVCacheInstanceId() *string {
	return s.KVCacheInstanceId
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) GetKVCacheInstanceStatus() *string {
	return s.KVCacheInstanceStatus
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) GetKVCacheSpace() *string {
	return s.KVCacheSpace
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) GetPort() *string {
	return s.Port
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) SetConnectionString(v string) *DescribeKVCacheInstanceAttributeResponseBody {
	s.ConnectionString = &v
	return s
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) SetCreationTime(v string) *DescribeKVCacheInstanceAttributeResponseBody {
	s.CreationTime = &v
	return s
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) SetExpireTime(v string) *DescribeKVCacheInstanceAttributeResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) SetExpired(v string) *DescribeKVCacheInstanceAttributeResponseBody {
	s.Expired = &v
	return s
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) SetKVCacheInstanceId(v string) *DescribeKVCacheInstanceAttributeResponseBody {
	s.KVCacheInstanceId = &v
	return s
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) SetKVCacheInstanceStatus(v string) *DescribeKVCacheInstanceAttributeResponseBody {
	s.KVCacheInstanceStatus = &v
	return s
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) SetKVCacheSpace(v string) *DescribeKVCacheInstanceAttributeResponseBody {
	s.KVCacheSpace = &v
	return s
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) SetPayType(v string) *DescribeKVCacheInstanceAttributeResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) SetPort(v string) *DescribeKVCacheInstanceAttributeResponseBody {
	s.Port = &v
	return s
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) SetRegionId(v string) *DescribeKVCacheInstanceAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) SetRequestId(v string) *DescribeKVCacheInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) SetVPCId(v string) *DescribeKVCacheInstanceAttributeResponseBody {
	s.VPCId = &v
	return s
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) SetVSwitchId(v string) *DescribeKVCacheInstanceAttributeResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) SetZoneId(v string) *DescribeKVCacheInstanceAttributeResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeKVCacheInstanceAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
