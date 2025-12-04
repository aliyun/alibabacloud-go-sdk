// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubnetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListSubnetsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *ListSubnetsResponseBody
	GetCode() *int32
	SetContent(v *ListSubnetsResponseBodyContent) *ListSubnetsResponseBody
	GetContent() *ListSubnetsResponseBodyContent
	SetMessage(v string) *ListSubnetsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSubnetsResponseBody
	GetRequestId() *string
}

type ListSubnetsResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *ListSubnetsResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 7F9082CC-3D94-560F-A575-8E8EF6CE2CB8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSubnetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSubnetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubnetsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListSubnetsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListSubnetsResponseBody) GetContent() *ListSubnetsResponseBodyContent {
	return s.Content
}

func (s *ListSubnetsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSubnetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSubnetsResponseBody) SetAccessDeniedDetail(v string) *ListSubnetsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListSubnetsResponseBody) SetCode(v int32) *ListSubnetsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSubnetsResponseBody) SetContent(v *ListSubnetsResponseBodyContent) *ListSubnetsResponseBody {
	s.Content = v
	return s
}

func (s *ListSubnetsResponseBody) SetMessage(v string) *ListSubnetsResponseBody {
	s.Message = &v
	return s
}

func (s *ListSubnetsResponseBody) SetRequestId(v string) *ListSubnetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubnetsResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSubnetsResponseBodyContent struct {
	// Lingjun subnet information list
	Data []*ListSubnetsResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSubnetsResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListSubnetsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListSubnetsResponseBodyContent) GetData() []*ListSubnetsResponseBodyContentData {
	return s.Data
}

func (s *ListSubnetsResponseBodyContent) GetTotal() *int64 {
	return s.Total
}

func (s *ListSubnetsResponseBodyContent) SetData(v []*ListSubnetsResponseBodyContentData) *ListSubnetsResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListSubnetsResponseBodyContent) SetTotal(v int64) *ListSubnetsResponseBodyContent {
	s.Total = &v
	return s
}

func (s *ListSubnetsResponseBodyContent) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSubnetsResponseBodyContentData struct {
	// The CIDR block of the Subnet.
	//
	// 	- The network segment of the subnet must be a proper subset of the network segment of Lingjun to which it belongs, and the mask must be between 16 bits and 29 bits, which can provide 8 to 65536 addresses. For example, the CIDR block of the Lingjun CIDR block is 192.168.0.0/16, and the CIDR blocks of the subnets under the Lingjun CIDR block are 192.168.0.0/17 to 192.168.0.0/29.
	//
	// 	- The first and last three IP addresses of each subnet segment are reserved by the system. For example, the CIDR blocks of the subnet are 192.168.1.0/24,192.168.1.0, 192.168.1.253, 192.168.1.254, and 192.168.1.255.
	//
	// example:
	//
	// 172.18.0.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1678273219000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the agent was last modified.
	//
	// example:
	//
	// 1678273219000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Number of NCs
	//
	// example:
	//
	// 16
	NcCount *int32 `json:"NcCount,omitempty" xml:"NcCount,omitempty"`
	// Number of Lingjun network interface controller
	//
	// example:
	//
	// 1
	NetworkInterfaceCount *int32 `json:"NetworkInterfaceCount,omitempty" xml:"NetworkInterfaceCount,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of your Alibaba Cloud resource group.
	//
	// For more information about resource groups, see [Resource groups](https://help.aliyun.com/document_detail/94475.htm?spm=a2c4g.11186623.0.0.29e15d7akXhpuu).
	//
	// example:
	//
	// rg-aek2l4sq6l7unhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the cache reserve instance. Valid values:
	//
	// 	- **Available**: Normal
	//
	// 	- **Not Available**: Unavailable
	//
	// 	- **Executing**: Executing
	//
	// 	- **Deleting**: The node is being deleted.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Lingjun subnet instance ID
	//
	// example:
	//
	// subnet-c6wci55i
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// Lingjun subnet instance name
	//
	// example:
	//
	// yzp-rg-test3
	SubnetName *string `json:"SubnetName,omitempty" xml:"SubnetName,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags.
	Tags []*ListSubnetsResponseBodyContentDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Lingjun Subnet Usage Type; optional; optional. Valid values:
	//
	// 	- **If you do not set this field for a common type**
	//
	// 	- **OOB*	- :OOB type
	//
	// 	- **LB**: LB type
	//
	// example:
	//
	// OOB
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// vpd basic information
	VpdBaseInfo *ListSubnetsResponseBodyContentDataVpdBaseInfo `json:"VpdBaseInfo,omitempty" xml:"VpdBaseInfo,omitempty" type:"Struct"`
	// The ID of the Lingjun CIDR block.
	//
	// example:
	//
	// vpd-ze3na0wf
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListSubnetsResponseBodyContentData) String() string {
	return dara.Prettify(s)
}

func (s ListSubnetsResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListSubnetsResponseBodyContentData) GetCidr() *string {
	return s.Cidr
}

func (s *ListSubnetsResponseBodyContentData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSubnetsResponseBodyContentData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListSubnetsResponseBodyContentData) GetMessage() *string {
	return s.Message
}

func (s *ListSubnetsResponseBodyContentData) GetNcCount() *int32 {
	return s.NcCount
}

func (s *ListSubnetsResponseBodyContentData) GetNetworkInterfaceCount() *int32 {
	return s.NetworkInterfaceCount
}

func (s *ListSubnetsResponseBodyContentData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSubnetsResponseBodyContentData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListSubnetsResponseBodyContentData) GetStatus() *string {
	return s.Status
}

func (s *ListSubnetsResponseBodyContentData) GetSubnetId() *string {
	return s.SubnetId
}

func (s *ListSubnetsResponseBodyContentData) GetSubnetName() *string {
	return s.SubnetName
}

func (s *ListSubnetsResponseBodyContentData) GetTags() []*ListSubnetsResponseBodyContentDataTags {
	return s.Tags
}

func (s *ListSubnetsResponseBodyContentData) GetTenantId() *string {
	return s.TenantId
}

func (s *ListSubnetsResponseBodyContentData) GetType() *string {
	return s.Type
}

func (s *ListSubnetsResponseBodyContentData) GetVpdBaseInfo() *ListSubnetsResponseBodyContentDataVpdBaseInfo {
	return s.VpdBaseInfo
}

func (s *ListSubnetsResponseBodyContentData) GetVpdId() *string {
	return s.VpdId
}

func (s *ListSubnetsResponseBodyContentData) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListSubnetsResponseBodyContentData) SetCidr(v string) *ListSubnetsResponseBodyContentData {
	s.Cidr = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetCreateTime(v string) *ListSubnetsResponseBodyContentData {
	s.CreateTime = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetGmtModified(v string) *ListSubnetsResponseBodyContentData {
	s.GmtModified = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetMessage(v string) *ListSubnetsResponseBodyContentData {
	s.Message = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetNcCount(v int32) *ListSubnetsResponseBodyContentData {
	s.NcCount = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetNetworkInterfaceCount(v int32) *ListSubnetsResponseBodyContentData {
	s.NetworkInterfaceCount = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetRegionId(v string) *ListSubnetsResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetResourceGroupId(v string) *ListSubnetsResponseBodyContentData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetStatus(v string) *ListSubnetsResponseBodyContentData {
	s.Status = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetSubnetId(v string) *ListSubnetsResponseBodyContentData {
	s.SubnetId = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetSubnetName(v string) *ListSubnetsResponseBodyContentData {
	s.SubnetName = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetTags(v []*ListSubnetsResponseBodyContentDataTags) *ListSubnetsResponseBodyContentData {
	s.Tags = v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetTenantId(v string) *ListSubnetsResponseBodyContentData {
	s.TenantId = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetType(v string) *ListSubnetsResponseBodyContentData {
	s.Type = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetVpdBaseInfo(v *ListSubnetsResponseBodyContentDataVpdBaseInfo) *ListSubnetsResponseBodyContentData {
	s.VpdBaseInfo = v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetVpdId(v string) *ListSubnetsResponseBodyContentData {
	s.VpdId = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetZoneId(v string) *ListSubnetsResponseBodyContentData {
	s.ZoneId = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VpdBaseInfo != nil {
		if err := s.VpdBaseInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSubnetsResponseBodyContentDataTags struct {
	// The tag key.
	//
	// You cannot specify an empty string as a tag key. It can be up to 64 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// tag-subnet
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag that is added to the resource.
	//
	// The tag value can be empty or a string of up to 128 characters. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// Each key-value pair must be unique. You can specify values for at most 20 tag keys in each call.
	//
	// example:
	//
	// subnet-group-1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListSubnetsResponseBodyContentDataTags) String() string {
	return dara.Prettify(s)
}

func (s ListSubnetsResponseBodyContentDataTags) GoString() string {
	return s.String()
}

func (s *ListSubnetsResponseBodyContentDataTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListSubnetsResponseBodyContentDataTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListSubnetsResponseBodyContentDataTags) SetTagKey(v string) *ListSubnetsResponseBodyContentDataTags {
	s.TagKey = &v
	return s
}

func (s *ListSubnetsResponseBodyContentDataTags) SetTagValue(v string) *ListSubnetsResponseBodyContentDataTags {
	s.TagValue = &v
	return s
}

func (s *ListSubnetsResponseBodyContentDataTags) Validate() error {
	return dara.Validate(s)
}

type ListSubnetsResponseBodyContentDataVpdBaseInfo struct {
	// The CIDR block of the VPD.
	//
	// 	- We recommend that you use an RFC private endpoint as the Lingjun CIDR block, such as 10.0.0.0/8,172.16.0.0/12,192.168.0.0/16. In scenarios where the Doringjun CIDR block is connected to each other or where the Lingjun CIDR block is connected to a VPC, make sure that the addresses do not conflict with each other.
	//
	// 	- You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, or 169.254.0.0/16 and their subnets as the primary IPv4 CIDR block of the VPD.
	//
	// example:
	//
	// 10.0.0.0/8
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1678273219000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Lingjun CIDR block instance ID
	//
	// example:
	//
	// vpd-d3isyds4
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// Lingjun CIDR block instance name
	//
	// example:
	//
	// yzp-rg-test3
	VpdName *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
}

func (s ListSubnetsResponseBodyContentDataVpdBaseInfo) String() string {
	return dara.Prettify(s)
}

func (s ListSubnetsResponseBodyContentDataVpdBaseInfo) GoString() string {
	return s.String()
}

func (s *ListSubnetsResponseBodyContentDataVpdBaseInfo) GetCidr() *string {
	return s.Cidr
}

func (s *ListSubnetsResponseBodyContentDataVpdBaseInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSubnetsResponseBodyContentDataVpdBaseInfo) GetVpdId() *string {
	return s.VpdId
}

func (s *ListSubnetsResponseBodyContentDataVpdBaseInfo) GetVpdName() *string {
	return s.VpdName
}

func (s *ListSubnetsResponseBodyContentDataVpdBaseInfo) SetCidr(v string) *ListSubnetsResponseBodyContentDataVpdBaseInfo {
	s.Cidr = &v
	return s
}

func (s *ListSubnetsResponseBodyContentDataVpdBaseInfo) SetCreateTime(v string) *ListSubnetsResponseBodyContentDataVpdBaseInfo {
	s.CreateTime = &v
	return s
}

func (s *ListSubnetsResponseBodyContentDataVpdBaseInfo) SetVpdId(v string) *ListSubnetsResponseBodyContentDataVpdBaseInfo {
	s.VpdId = &v
	return s
}

func (s *ListSubnetsResponseBodyContentDataVpdBaseInfo) SetVpdName(v string) *ListSubnetsResponseBodyContentDataVpdBaseInfo {
	s.VpdName = &v
	return s
}

func (s *ListSubnetsResponseBodyContentDataVpdBaseInfo) Validate() error {
	return dara.Validate(s)
}
