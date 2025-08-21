// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
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
}

type DescribeInstancesRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-beijing-cmcc
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The IDs of the regions. The value is a JSON array that consists of up to 100 IDs. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// ["cn-suzhou-telecom","cn-chengdu-telecom"]
	EnsRegionIds *string `json:"EnsRegionIds,omitempty" xml:"EnsRegionIds,omitempty"`
	// The ID of the edge service. You can use the ID to query information about the instances that are created in the edge service.
	//
	// example:
	//
	// ens-20190730202316s****
	EnsServiceId *string `json:"EnsServiceId,omitempty" xml:"EnsServiceId,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// centos_6_08_64_20G_alibase_****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-5iqczfxps7csjrxeca****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IDs of the instances. The value is a JSON array that consists of up to 100 IDs. Separate IDs with commas (,).
	//
	// example:
	//
	// ["i-5iqczfxps7csjrxeca07****", "i-5iqczfxps7csjrxeca07****"]]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// TestName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The condition that you want to use to filter instances by category. Valid values:
	//
	// 	- EnsInstance: ENS instances that you purchase.
	//
	// 	- EnsService: ENS instances that belong to edge services.
	//
	// 	- BuildMachine: ENS instances that are configured with image builders.
	//
	// 	- EnsPostPaidInstance: Pay-as-you-go ENS instances that you purchase.
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
	// The internal IP address of the instance.
	//
	// example:
	//
	// 47.100.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The ID of the network.
	//
	// example:
	//
	// n-2zeuphj08tt7q3brd****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The method that you want to use to sort instances. The value of this parameter is in the JSON format.
	//
	// You can sort instances by name, expiration time, node ID, or creation time. You can specify one or more methods.
	//
	// example:
	//
	// {"InstanceNameSort":"asc","ExpireTimeSort":"asc","CreationTimeSort":"desc"}}
	OrderByParams *string `json:"OrderByParams,omitempty" xml:"OrderByParams,omitempty"`
	// The page number. Pages start from page **1**.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. The maximum value is **100**.
	//
	// Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keyword that you use to query the logs of the service. You can specify the values of parameters such as **ip**, **InstanceName**, and **InstanceId*	- as the keyword.
	//
	// example:
	//
	// Joshua
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-5kyicq2kfcapxrdds6tar7jqb
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The status of the service. Valid values.
	ServiceStatus []*string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty" type:"Repeated"`
	// The status of the instance. Valid values:
	//
	// 	- Running
	//
	// 	- Stopped
	//
	// 	- Expired
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags that are added to the resource. This operation does not return tag information. You can call this operation in combination with the tag-related operations.
	Tags []*DescribeInstancesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-2zeh0r1pabwtg6wcs****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
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

func (s *DescribeInstancesRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesRequestTags struct {
	// The key of the tag that are to add to the instance. Valid values: 1 to 20.
	//
	// example:
	//
	// tag
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the instance. Valid values: 1 to 20.
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
