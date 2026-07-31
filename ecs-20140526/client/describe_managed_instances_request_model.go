// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeManagedInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivationId(v string) *DescribeManagedInstancesRequest
	GetActivationId() *string
	SetConnected(v string) *DescribeManagedInstancesRequest
	GetConnected() *string
	SetInstanceId(v []*string) *DescribeManagedInstancesRequest
	GetInstanceId() []*string
	SetInstanceIp(v string) *DescribeManagedInstancesRequest
	GetInstanceIp() *string
	SetInstanceName(v string) *DescribeManagedInstancesRequest
	GetInstanceName() *string
	SetMachineId(v string) *DescribeManagedInstancesRequest
	GetMachineId() *string
	SetMaxResults(v int32) *DescribeManagedInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeManagedInstancesRequest
	GetNextToken() *string
	SetOsType(v string) *DescribeManagedInstancesRequest
	GetOsType() *string
	SetOwnerAccount(v string) *DescribeManagedInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeManagedInstancesRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *DescribeManagedInstancesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeManagedInstancesRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeManagedInstancesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeManagedInstancesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeManagedInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeManagedInstancesRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeManagedInstancesRequestTag) *DescribeManagedInstancesRequest
	GetTag() []*DescribeManagedInstancesRequestTag
}

type DescribeManagedInstancesRequest struct {
	// The activation code ID.
	//
	// example:
	//
	// 4ECEEE12-56F1-4FBC-9AB1-890F7494****
	ActivationId *string `json:"ActivationId,omitempty" xml:"ActivationId,omitempty"`
	// Specifies whether the managed instance is connected.
	//
	// true: The managed instance is connected. You can manage the managed instance by using Cloud Assistant.
	//
	// false: The managed instance is not connected. The server may be shut down or Cloud Assistant Agent may not be properly installed.
	//
	// example:
	//
	// true
	Connected *string `json:"Connected,omitempty" xml:"Connected,omitempty"`
	// The ID of the managed instance. Valid values of N: 1 to 50.
	//
	// example:
	//
	// mi-hz018jrc1o0****
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// The internal IP address or public IP address of the managed instance.
	//
	// example:
	//
	// ``192.168.**.**``
	InstanceIp *string `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty"`
	// The name of the managed instance.
	//
	// example:
	//
	// my-webapp-server
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The value of the MachineId parameter specified when registering the managed instance. A maximum of 36 characters are allowed.
	//
	// Example registration script:
	//
	// ```
	//
	// aliyun-service --register \\
	//
	//   --RegionId=ap-southeast-1 \\
	//
	//   --ActivationId=xxxxxxxxxxx \\
	//
	//   --ActivationCode=xxxxxxxxx \\
	//
	//   --MachineId=xxxxxx \\ # Optional parameter that specifies the unique identifier of the machine
	//
	//   --ForceResue
	//
	// ```
	//
	// - If MachineId and ForceResult are specified during registration, Cloud Assistant generates a fixed managed instance ID for this MachineId.
	//
	// - If MachineId is not explicitly specified, Cloud Assistant automatically generates a MachineId value based on the hardware information of the machine.
	//
	// - Recommendation: Explicitly specify MachineId and ForceResult to mark the mapping between managed instances and on-premises machines.
	//
	// example:
	//
	// GOG4X8312A0188
	MachineId *string `json:"MachineId,omitempty" xml:"MachineId,omitempty"`
	// The maximum number of entries per page for a paging query.
	//
	// Maximum value: 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this parameter to the NextToken value returned in the previous API call.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The operating system type of the managed instance. Valid values:
	//
	// - windows
	//
	// - linux
	//
	// - FreeBSD
	//
	// example:
	//
	// windows
	OsType       *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// > This parameter is about to be deprecated. Use NextToken and MaxResults to complete paging query operations.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// > This parameter is about to be deprecated. Use NextToken and MaxResults to complete paging query operations.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. Currently supported regions: China (Qingdao), China (Beijing), China (Zhangjiakou), China (Hohhot), China (Ulanqab), China (Hangzhou), China (Shanghai), China (Shenzhen), China (Heyuan), China (Guangzhou), China (Chengdu), Hong Kong (China), Singapore, Japan (Tokyo), US (Silicon Valley), and US (Virginia).
	//
	// You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query region IDs and other information.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the managed instance belongs.
	//
	// example:
	//
	// rg-123******
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	Tag []*DescribeManagedInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeManagedInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeManagedInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeManagedInstancesRequest) GetActivationId() *string {
	return s.ActivationId
}

func (s *DescribeManagedInstancesRequest) GetConnected() *string {
	return s.Connected
}

func (s *DescribeManagedInstancesRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *DescribeManagedInstancesRequest) GetInstanceIp() *string {
	return s.InstanceIp
}

func (s *DescribeManagedInstancesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeManagedInstancesRequest) GetMachineId() *string {
	return s.MachineId
}

func (s *DescribeManagedInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeManagedInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeManagedInstancesRequest) GetOsType() *string {
	return s.OsType
}

func (s *DescribeManagedInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeManagedInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeManagedInstancesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeManagedInstancesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeManagedInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeManagedInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeManagedInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeManagedInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeManagedInstancesRequest) GetTag() []*DescribeManagedInstancesRequestTag {
	return s.Tag
}

func (s *DescribeManagedInstancesRequest) SetActivationId(v string) *DescribeManagedInstancesRequest {
	s.ActivationId = &v
	return s
}

func (s *DescribeManagedInstancesRequest) SetConnected(v string) *DescribeManagedInstancesRequest {
	s.Connected = &v
	return s
}

func (s *DescribeManagedInstancesRequest) SetInstanceId(v []*string) *DescribeManagedInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *DescribeManagedInstancesRequest) SetInstanceIp(v string) *DescribeManagedInstancesRequest {
	s.InstanceIp = &v
	return s
}

func (s *DescribeManagedInstancesRequest) SetInstanceName(v string) *DescribeManagedInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeManagedInstancesRequest) SetMachineId(v string) *DescribeManagedInstancesRequest {
	s.MachineId = &v
	return s
}

func (s *DescribeManagedInstancesRequest) SetMaxResults(v int32) *DescribeManagedInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeManagedInstancesRequest) SetNextToken(v string) *DescribeManagedInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeManagedInstancesRequest) SetOsType(v string) *DescribeManagedInstancesRequest {
	s.OsType = &v
	return s
}

func (s *DescribeManagedInstancesRequest) SetOwnerAccount(v string) *DescribeManagedInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeManagedInstancesRequest) SetOwnerId(v int64) *DescribeManagedInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeManagedInstancesRequest) SetPageNumber(v int64) *DescribeManagedInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeManagedInstancesRequest) SetPageSize(v int64) *DescribeManagedInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeManagedInstancesRequest) SetRegionId(v string) *DescribeManagedInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeManagedInstancesRequest) SetResourceGroupId(v string) *DescribeManagedInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeManagedInstancesRequest) SetResourceOwnerAccount(v string) *DescribeManagedInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeManagedInstancesRequest) SetResourceOwnerId(v int64) *DescribeManagedInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeManagedInstancesRequest) SetTag(v []*DescribeManagedInstancesRequestTag) *DescribeManagedInstancesRequest {
	s.Tag = v
	return s
}

func (s *DescribeManagedInstancesRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeManagedInstancesRequestTag struct {
	// The tag key of the managed instance. Valid values of N: 1 to 20. The tag key cannot be an empty string.
	//
	// If you use a single tag to filter resources, the resource count with this tag cannot exceed 1,000. If you use multiple tags to filter resources, the resource count of resources that have all specified tags attached cannot exceed 1,000. If the resource count exceeds 1,000, call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation to query resources.
	//
	// The tag key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the managed instance. Valid values of N: 1 to 20. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeManagedInstancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeManagedInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeManagedInstancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeManagedInstancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeManagedInstancesRequestTag) SetKey(v string) *DescribeManagedInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeManagedInstancesRequestTag) SetValue(v string) *DescribeManagedInstancesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeManagedInstancesRequestTag) Validate() error {
	return dara.Validate(s)
}
