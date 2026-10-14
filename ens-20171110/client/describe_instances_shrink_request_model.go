// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEipAddressesShrink(v string) *DescribeInstancesShrinkRequest
	GetEipAddressesShrink() *string
	SetEnsRegionId(v string) *DescribeInstancesShrinkRequest
	GetEnsRegionId() *string
	SetEnsRegionIds(v string) *DescribeInstancesShrinkRequest
	GetEnsRegionIds() *string
	SetEnsServiceId(v string) *DescribeInstancesShrinkRequest
	GetEnsServiceId() *string
	SetImageId(v string) *DescribeInstancesShrinkRequest
	GetImageId() *string
	SetInstanceId(v string) *DescribeInstancesShrinkRequest
	GetInstanceId() *string
	SetInstanceIds(v string) *DescribeInstancesShrinkRequest
	GetInstanceIds() *string
	SetInstanceName(v string) *DescribeInstancesShrinkRequest
	GetInstanceName() *string
	SetInstanceResourceType(v string) *DescribeInstancesShrinkRequest
	GetInstanceResourceType() *string
	SetInstanceType(v string) *DescribeInstancesShrinkRequest
	GetInstanceType() *string
	SetIntranetIp(v string) *DescribeInstancesShrinkRequest
	GetIntranetIp() *string
	SetNetworkId(v string) *DescribeInstancesShrinkRequest
	GetNetworkId() *string
	SetOrderByParams(v string) *DescribeInstancesShrinkRequest
	GetOrderByParams() *string
	SetPageNumber(v int32) *DescribeInstancesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v string) *DescribeInstancesShrinkRequest
	GetPageSize() *string
	SetSearchKey(v string) *DescribeInstancesShrinkRequest
	GetSearchKey() *string
	SetSecurityGroupId(v string) *DescribeInstancesShrinkRequest
	GetSecurityGroupId() *string
	SetServiceStatusShrink(v string) *DescribeInstancesShrinkRequest
	GetServiceStatusShrink() *string
	SetStatus(v string) *DescribeInstancesShrinkRequest
	GetStatus() *string
	SetTagsShrink(v string) *DescribeInstancesShrinkRequest
	GetTagsShrink() *string
	SetVSwitchId(v string) *DescribeInstancesShrinkRequest
	GetVSwitchId() *string
	SetVpdId(v string) *DescribeInstancesShrinkRequest
	GetVpdId() *string
	SetVpdVSwitchId(v string) *DescribeInstancesShrinkRequest
	GetVpdVSwitchId() *string
}

type DescribeInstancesShrinkRequest struct {
	// The list of EIP addresses used to filter instances.
	EipAddressesShrink *string `json:"EipAddresses,omitempty" xml:"EipAddresses,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing-cmcc
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The list of edge zone IDs. This parameter is a formatted JSON array that contains up to 100 IDs separated by commas (,).
	//
	// example:
	//
	// ["cn-suzhou-telecom","cn-chengdu-telecom"]
	EnsRegionIds *string `json:"EnsRegionIds,omitempty" xml:"EnsRegionIds,omitempty"`
	// The Edge Node Service ID. You can use this ID to query the instances created under the edge service.
	//
	// example:
	//
	// ens-20190730202316s****
	EnsServiceId *string `json:"EnsServiceId,omitempty" xml:"EnsServiceId,omitempty"`
	// The image ID.
	//
	// example:
	//
	// centos_6_08_64_20G_alibase_****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-5iqczfxps7csjrxeca****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of instance IDs. This parameter is a formatted JSON array that contains up to 100 IDs separated by commas (,).
	//
	// example:
	//
	// ["i-5iqczfxps7csjrxeca07****", "i-5iqczfxps7csjrxeca07****"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The instance name.
	//
	// example:
	//
	// TestName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The filter type. Valid values:
	//
	// example:
	//
	// EnsService
	InstanceResourceType *string `json:"InstanceResourceType,omitempty" xml:"InstanceResourceType,omitempty"`
	// The instance type.
	//
	// example:
	//
	// ens.se1.tiny
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The internal IP address.
	//
	// example:
	//
	// 47.100.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The network ID.
	//
	// example:
	//
	// n-2zeuphj08tt7q3brd****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The sorting parameter in JSON format.
	//
	// example:
	//
	// {
	//
	//     "InstanceNameSort": "asc",
	//
	//     "ExpireTimeSort": "asc",
	//
	//     "CreationTimeSort": "desc",
	//
	//     "EnsRegionId": "desc"
	//
	// }
	OrderByParams *string `json:"OrderByParams,omitempty" xml:"OrderByParams,omitempty"`
	// The page number of the instance status list. The value starts from **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query. Maximum value: **100**.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search keyword. Supports parameters such as **ip**, **InstanceName**, and **InstanceId**.
	//
	// example:
	//
	// Joshua
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-5kyicq2kfcapxrdds6tar7jqb
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The business status of the instance, which does not represent the actual running status of the instance.
	ServiceStatusShrink *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// The instance status. Valid values:
	//
	// - Running: Running.
	//
	// - Stopped: Stopped.
	//
	// - Expired: Expired.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The resource tags. This operation does not return tag information. Use this parameter in combination with the APIs in the tag category.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-2zeh0r1pabwtg6wcs****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPD network ID.
	//
	// example:
	//
	// vpd-5s7xxxxxxx
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The VPD vSwitch ID.
	//
	// example:
	//
	// vsw-5s7xxxxxxx
	VpdVSwitchId *string `json:"VpdVSwitchId,omitempty" xml:"VpdVSwitchId,omitempty"`
}

func (s DescribeInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesShrinkRequest) GetEipAddressesShrink() *string {
	return s.EipAddressesShrink
}

func (s *DescribeInstancesShrinkRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeInstancesShrinkRequest) GetEnsRegionIds() *string {
	return s.EnsRegionIds
}

func (s *DescribeInstancesShrinkRequest) GetEnsServiceId() *string {
	return s.EnsServiceId
}

func (s *DescribeInstancesShrinkRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeInstancesShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstancesShrinkRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeInstancesShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstancesShrinkRequest) GetInstanceResourceType() *string {
	return s.InstanceResourceType
}

func (s *DescribeInstancesShrinkRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstancesShrinkRequest) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeInstancesShrinkRequest) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeInstancesShrinkRequest) GetOrderByParams() *string {
	return s.OrderByParams
}

func (s *DescribeInstancesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstancesShrinkRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeInstancesShrinkRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *DescribeInstancesShrinkRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeInstancesShrinkRequest) GetServiceStatusShrink() *string {
	return s.ServiceStatusShrink
}

func (s *DescribeInstancesShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstancesShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *DescribeInstancesShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeInstancesShrinkRequest) GetVpdId() *string {
	return s.VpdId
}

func (s *DescribeInstancesShrinkRequest) GetVpdVSwitchId() *string {
	return s.VpdVSwitchId
}

func (s *DescribeInstancesShrinkRequest) SetEipAddressesShrink(v string) *DescribeInstancesShrinkRequest {
	s.EipAddressesShrink = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetEnsRegionId(v string) *DescribeInstancesShrinkRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetEnsRegionIds(v string) *DescribeInstancesShrinkRequest {
	s.EnsRegionIds = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetEnsServiceId(v string) *DescribeInstancesShrinkRequest {
	s.EnsServiceId = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetImageId(v string) *DescribeInstancesShrinkRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetInstanceId(v string) *DescribeInstancesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetInstanceIds(v string) *DescribeInstancesShrinkRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetInstanceName(v string) *DescribeInstancesShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetInstanceResourceType(v string) *DescribeInstancesShrinkRequest {
	s.InstanceResourceType = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetInstanceType(v string) *DescribeInstancesShrinkRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetIntranetIp(v string) *DescribeInstancesShrinkRequest {
	s.IntranetIp = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetNetworkId(v string) *DescribeInstancesShrinkRequest {
	s.NetworkId = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetOrderByParams(v string) *DescribeInstancesShrinkRequest {
	s.OrderByParams = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetPageNumber(v int32) *DescribeInstancesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetPageSize(v string) *DescribeInstancesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetSearchKey(v string) *DescribeInstancesShrinkRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetSecurityGroupId(v string) *DescribeInstancesShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetServiceStatusShrink(v string) *DescribeInstancesShrinkRequest {
	s.ServiceStatusShrink = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetStatus(v string) *DescribeInstancesShrinkRequest {
	s.Status = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetTagsShrink(v string) *DescribeInstancesShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetVSwitchId(v string) *DescribeInstancesShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetVpdId(v string) *DescribeInstancesShrinkRequest {
	s.VpdId = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetVpdVSwitchId(v string) *DescribeInstancesShrinkRequest {
	s.VpdVSwitchId = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
