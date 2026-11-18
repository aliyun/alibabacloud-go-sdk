// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairAIGatewayInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *DescribeTairAIGatewayInstancesRequest
	GetChargeType() *string
	SetExpired(v string) *DescribeTairAIGatewayInstancesRequest
	GetExpired() *string
	SetInstanceIds(v string) *DescribeTairAIGatewayInstancesRequest
	GetInstanceIds() *string
	SetInstanceStatus(v string) *DescribeTairAIGatewayInstancesRequest
	GetInstanceStatus() *string
	SetPageNumber(v int32) *DescribeTairAIGatewayInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTairAIGatewayInstancesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeTairAIGatewayInstancesRequest
	GetRegionId() *string
	SetSearchKey(v string) *DescribeTairAIGatewayInstancesRequest
	GetSearchKey() *string
	SetSecurityToken(v string) *DescribeTairAIGatewayInstancesRequest
	GetSecurityToken() *string
	SetVSwitchId(v string) *DescribeTairAIGatewayInstancesRequest
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeTairAIGatewayInstancesRequest
	GetVpcId() *string
	SetZoneId(v string) *DescribeTairAIGatewayInstancesRequest
	GetZoneId() *string
}

type DescribeTairAIGatewayInstancesRequest struct {
	// The billing method. Valid values:
	//
	// - **PostPaid**: pay-as-you-go
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The expiration status of the instance. Valid values:
	//
	// 	- **true**: The instance has expired.
	//
	// 	- **false**: The instance has not expired.
	//
	// example:
	//
	// false
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The IDs of the instances to query.
	//
	// > Separate multiple instance IDs with commas (,). You can specify up to 30 instance IDs in a single request.
	//
	// example:
	//
	// tg-2zezxszhcgatnx****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The instance status. Valid values:
	//
	// 	- **Normal**: The instance is running.
	//
	// 	- **Creating**: The instance is being created.
	//
	// 	- **Changing**: The instance is being modified.
	//
	// 	- **Inactive**: The instance is disabled.
	//
	// 	- **Flushing**: The instance data is being purged.
	//
	// 	- **Released**: The instance is released.
	//
	// 	- **Transforming**: The instance is being transformed.
	//
	// 	- **Migrating**: The instance is being migrated.
	//
	// 	- **BackupRecovering**: The instance is recovering from a backup.
	//
	// 	- **MinorVersionUpgrading**: The minor version of the instance is being upgraded.
	//
	// 	- **NetworkModifying**: The network of the instance is being modified.
	//
	// 	- **SSLModifying**: The SSL configuration of the instance is being modified.
	//
	// 	- **MajorVersionUpgrading**: A major engine version upgrade is in progress. The instance remains accessible.
	//
	// > For more information about instance statuses and impacts, see [Instance statuses and impacts](https://help.aliyun.com/document_detail/200740.html).
	//
	// example:
	//
	// Normal
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The search keyword. Fuzzy search is supported for instance names or instance IDs.
	//
	// example:
	//
	// apitest
	SearchKey     *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The ID of the vSwitch in the VPC. You can call the [DescribeVpcs](https://help.aliyun.com/document_detail/35739.html) operation to query the vSwitch ID.
	//
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-bp1nme44gek34slfc****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-beijing-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeTairAIGatewayInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairAIGatewayInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTairAIGatewayInstancesRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeTairAIGatewayInstancesRequest) GetExpired() *string {
	return s.Expired
}

func (s *DescribeTairAIGatewayInstancesRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeTairAIGatewayInstancesRequest) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeTairAIGatewayInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTairAIGatewayInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTairAIGatewayInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTairAIGatewayInstancesRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *DescribeTairAIGatewayInstancesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeTairAIGatewayInstancesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeTairAIGatewayInstancesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeTairAIGatewayInstancesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeTairAIGatewayInstancesRequest) SetChargeType(v string) *DescribeTairAIGatewayInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesRequest) SetExpired(v string) *DescribeTairAIGatewayInstancesRequest {
	s.Expired = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesRequest) SetInstanceIds(v string) *DescribeTairAIGatewayInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesRequest) SetInstanceStatus(v string) *DescribeTairAIGatewayInstancesRequest {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesRequest) SetPageNumber(v int32) *DescribeTairAIGatewayInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesRequest) SetPageSize(v int32) *DescribeTairAIGatewayInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesRequest) SetRegionId(v string) *DescribeTairAIGatewayInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesRequest) SetSearchKey(v string) *DescribeTairAIGatewayInstancesRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesRequest) SetSecurityToken(v string) *DescribeTairAIGatewayInstancesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesRequest) SetVSwitchId(v string) *DescribeTairAIGatewayInstancesRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesRequest) SetVpcId(v string) *DescribeTairAIGatewayInstancesRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesRequest) SetZoneId(v string) *DescribeTairAIGatewayInstancesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesRequest) Validate() error {
	return dara.Validate(s)
}
