// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEipAddresses(v []*string) *DescribeInstancesRequest
	GetEipAddresses() []*string
	SetEnsRegionId(v string) *DescribeInstancesRequest
	GetEnsRegionId() *string
	SetEnsRegionIds(v string) *DescribeInstancesRequest
	GetEnsRegionIds() *string
	SetEnsServiceId(v string) *DescribeInstancesRequest
	GetEnsServiceId() *string
	SetImageId(v string) *DescribeInstancesRequest
	GetImageId() *string
	SetInstanceId(v string) *DescribeInstancesRequest
	GetInstanceId() *string
	SetInstanceIds(v string) *DescribeInstancesRequest
	GetInstanceIds() *string
	SetInstanceName(v string) *DescribeInstancesRequest
	GetInstanceName() *string
	SetInstanceResourceType(v string) *DescribeInstancesRequest
	GetInstanceResourceType() *string
	SetInstanceType(v string) *DescribeInstancesRequest
	GetInstanceType() *string
	SetIntranetIp(v string) *DescribeInstancesRequest
	GetIntranetIp() *string
	SetNetworkId(v string) *DescribeInstancesRequest
	GetNetworkId() *string
	SetOrderByParams(v string) *DescribeInstancesRequest
	GetOrderByParams() *string
	SetPageNumber(v int32) *DescribeInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v string) *DescribeInstancesRequest
	GetPageSize() *string
	SetSearchKey(v string) *DescribeInstancesRequest
	GetSearchKey() *string
	SetSecurityGroupId(v string) *DescribeInstancesRequest
	GetSecurityGroupId() *string
	SetServiceStatus(v []*string) *DescribeInstancesRequest
	GetServiceStatus() []*string
	SetStatus(v string) *DescribeInstancesRequest
	GetStatus() *string
	SetTags(v []*DescribeInstancesRequestTags) *DescribeInstancesRequest
	GetTags() []*DescribeInstancesRequestTags
	SetVSwitchId(v string) *DescribeInstancesRequest
	GetVSwitchId() *string
	SetVpdId(v string) *DescribeInstancesRequest
	GetVpdId() *string
	SetVpdVSwitchId(v string) *DescribeInstancesRequest
	GetVpdVSwitchId() *string
}

type DescribeInstancesRequest struct {
	// The list of EIP addresses used to filter instances.
	EipAddresses []*string `json:"EipAddresses,omitempty" xml:"EipAddresses,omitempty" type:"Repeated"`
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
	ServiceStatus []*string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty" type:"Repeated"`
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
	Tags []*DescribeInstancesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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

func (s DescribeInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) GetEipAddresses() []*string {
	return s.EipAddresses
}

func (s *DescribeInstancesRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeInstancesRequest) GetEnsRegionIds() *string {
	return s.EnsRegionIds
}

func (s *DescribeInstancesRequest) GetEnsServiceId() *string {
	return s.EnsServiceId
}

func (s *DescribeInstancesRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstancesRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeInstancesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstancesRequest) GetInstanceResourceType() *string {
	return s.InstanceResourceType
}

func (s *DescribeInstancesRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstancesRequest) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeInstancesRequest) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeInstancesRequest) GetOrderByParams() *string {
	return s.OrderByParams
}

func (s *DescribeInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstancesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeInstancesRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *DescribeInstancesRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeInstancesRequest) GetServiceStatus() []*string {
	return s.ServiceStatus
}

func (s *DescribeInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstancesRequest) GetTags() []*DescribeInstancesRequestTags {
	return s.Tags
}

func (s *DescribeInstancesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeInstancesRequest) GetVpdId() *string {
	return s.VpdId
}

func (s *DescribeInstancesRequest) GetVpdVSwitchId() *string {
	return s.VpdVSwitchId
}

func (s *DescribeInstancesRequest) SetEipAddresses(v []*string) *DescribeInstancesRequest {
	s.EipAddresses = v
	return s
}

func (s *DescribeInstancesRequest) SetEnsRegionId(v string) *DescribeInstancesRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeInstancesRequest) SetEnsRegionIds(v string) *DescribeInstancesRequest {
	s.EnsRegionIds = &v
	return s
}

func (s *DescribeInstancesRequest) SetEnsServiceId(v string) *DescribeInstancesRequest {
	s.EnsServiceId = &v
	return s
}

func (s *DescribeInstancesRequest) SetImageId(v string) *DescribeInstancesRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceId(v string) *DescribeInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceIds(v string) *DescribeInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceName(v string) *DescribeInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceResourceType(v string) *DescribeInstancesRequest {
	s.InstanceResourceType = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceType(v string) *DescribeInstancesRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstancesRequest) SetIntranetIp(v string) *DescribeInstancesRequest {
	s.IntranetIp = &v
	return s
}

func (s *DescribeInstancesRequest) SetNetworkId(v string) *DescribeInstancesRequest {
	s.NetworkId = &v
	return s
}

func (s *DescribeInstancesRequest) SetOrderByParams(v string) *DescribeInstancesRequest {
	s.OrderByParams = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageNumber(v int32) *DescribeInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageSize(v string) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetSearchKey(v string) *DescribeInstancesRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeInstancesRequest) SetSecurityGroupId(v string) *DescribeInstancesRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeInstancesRequest) SetServiceStatus(v []*string) *DescribeInstancesRequest {
	s.ServiceStatus = v
	return s
}

func (s *DescribeInstancesRequest) SetStatus(v string) *DescribeInstancesRequest {
	s.Status = &v
	return s
}

func (s *DescribeInstancesRequest) SetTags(v []*DescribeInstancesRequestTags) *DescribeInstancesRequest {
	s.Tags = v
	return s
}

func (s *DescribeInstancesRequest) SetVSwitchId(v string) *DescribeInstancesRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeInstancesRequest) SetVpdId(v string) *DescribeInstancesRequest {
	s.VpdId = &v
	return s
}

func (s *DescribeInstancesRequest) SetVpdVSwitchId(v string) *DescribeInstancesRequest {
	s.VpdVSwitchId = &v
	return s
}

func (s *DescribeInstancesRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesRequestTags struct {
	// The tag key of the instance. Valid values of N: 1 to 20.
	//
	// example:
	//
	// tag
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the instance. Valid values of N: 1 to 20.
	//
	// example:
	//
	// 2
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesRequestTags) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequestTags) GetKey() *string {
	return s.Key
}

func (s *DescribeInstancesRequestTags) GetValue() *string {
	return s.Value
}

func (s *DescribeInstancesRequestTags) SetKey(v string) *DescribeInstancesRequestTags {
	s.Key = &v
	return s
}

func (s *DescribeInstancesRequestTags) SetValue(v string) *DescribeInstancesRequestTags {
	s.Value = &v
	return s
}

func (s *DescribeInstancesRequestTags) Validate() error {
	return dara.Validate(s)
}
