// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubnetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetSubnetResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *GetSubnetResponseBody
	GetCode() *int32
	SetContent(v *GetSubnetResponseBodyContent) *GetSubnetResponseBody
	GetContent() *GetSubnetResponseBodyContent
	SetMessage(v string) *GetSubnetResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSubnetResponseBody
	GetRequestId() *string
}

type GetSubnetResponseBody struct {
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
	// The response data.
	Content *GetSubnetResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0901F411-28FA-5B9C-BAEE-7776463FF0DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSubnetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSubnetResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubnetResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetSubnetResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetSubnetResponseBody) GetContent() *GetSubnetResponseBodyContent {
	return s.Content
}

func (s *GetSubnetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSubnetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSubnetResponseBody) SetAccessDeniedDetail(v string) *GetSubnetResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetSubnetResponseBody) SetCode(v int32) *GetSubnetResponseBody {
	s.Code = &v
	return s
}

func (s *GetSubnetResponseBody) SetContent(v *GetSubnetResponseBodyContent) *GetSubnetResponseBody {
	s.Content = v
	return s
}

func (s *GetSubnetResponseBody) SetMessage(v string) *GetSubnetResponseBody {
	s.Message = &v
	return s
}

func (s *GetSubnetResponseBody) SetRequestId(v string) *GetSubnetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubnetResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSubnetResponseBodyContent struct {
	// The number of available IP addresses.
	//
	// example:
	//
	// 1024
	AvailableIps *int32 `json:"AvailableIps,omitempty" xml:"AvailableIps,omitempty"`
	// The CIDR block of the Subnet.
	//
	// 	- The network segment of the subnet must be a proper subset of the network segment of Lingjun to which it belongs, and the mask must be between 16 bits and 29 bits, which can provide 8 to 65536 addresses. For example, the CIDR block of the Lingjun CIDR block is 192.168.0.0/16, and the CIDR blocks of the subnets under the Lingjun CIDR block are 192.168.0.0/17 to 192.168.0.0/29.
	//
	// 	- The first and last three IP addresses of each subnet segment are reserved by the system. For example, the CIDR blocks of the subnet are 192.168.1.0/24,192.168.1.0, 192.168.1.253, 192.168.1.254, and 192.168.1.255.
	//
	// example:
	//
	// 10.10.10.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The time when the activation code was created.
	//
	// example:
	//
	// 1648085472000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the O\\&M task was modified.
	//
	// example:
	//
	// 1678273219000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The number of SLB.
	//
	// example:
	//
	// 0
	LbCount *int64 `json:"LbCount,omitempty" xml:"LbCount,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// test example
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of NCs.
	//
	// example:
	//
	// 16
	NcCount *int32 `json:"NcCount,omitempty" xml:"NcCount,omitempty"`
	// Number of Lingjun network interface controller
	//
	// example:
	//
	// 4
	NetworkInterfaceCount *int32 `json:"NetworkInterfaceCount,omitempty" xml:"NetworkInterfaceCount,omitempty"`
	// The total number of secondary private IP addresses.
	//
	// example:
	//
	// 20
	PrivateIpCount *int64 `json:"PrivateIpCount,omitempty" xml:"PrivateIpCount,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of your Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-aek2l4sq6l7unhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the cache reserve instance.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the Lingjun subnet instance.
	//
	// example:
	//
	// subnet-aj93mko8
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// The name of the Lingjun subnet instance.
	//
	// example:
	//
	// subnet-1
	SubnetName *string `json:"SubnetName,omitempty" xml:"SubnetName,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags.
	Tags []*GetSubnetResponseBodyContentTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The tenant ID.
	//
	// example:
	//
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Lingjun Subnet Usage Type; optional; optional. Valid values:
	//
	// 	- **Empty for common data types**
	//
	// 	- **OOB*	- :OOB type
	//
	// 	- **LB**: LB type
	//
	// example:
	//
	// OOB
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The information about the network segment of Lingjun.
	VpdBaseInfo *GetSubnetResponseBodyContentVpdBaseInfo `json:"VpdBaseInfo,omitempty" xml:"VpdBaseInfo,omitempty" type:"Struct"`
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

func (s GetSubnetResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s GetSubnetResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetSubnetResponseBodyContent) GetAvailableIps() *int32 {
	return s.AvailableIps
}

func (s *GetSubnetResponseBodyContent) GetCidr() *string {
	return s.Cidr
}

func (s *GetSubnetResponseBodyContent) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetSubnetResponseBodyContent) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetSubnetResponseBodyContent) GetLbCount() *int64 {
	return s.LbCount
}

func (s *GetSubnetResponseBodyContent) GetMessage() *string {
	return s.Message
}

func (s *GetSubnetResponseBodyContent) GetNcCount() *int32 {
	return s.NcCount
}

func (s *GetSubnetResponseBodyContent) GetNetworkInterfaceCount() *int32 {
	return s.NetworkInterfaceCount
}

func (s *GetSubnetResponseBodyContent) GetPrivateIpCount() *int64 {
	return s.PrivateIpCount
}

func (s *GetSubnetResponseBodyContent) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSubnetResponseBodyContent) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetSubnetResponseBodyContent) GetStatus() *string {
	return s.Status
}

func (s *GetSubnetResponseBodyContent) GetSubnetId() *string {
	return s.SubnetId
}

func (s *GetSubnetResponseBodyContent) GetSubnetName() *string {
	return s.SubnetName
}

func (s *GetSubnetResponseBodyContent) GetTags() []*GetSubnetResponseBodyContentTags {
	return s.Tags
}

func (s *GetSubnetResponseBodyContent) GetTenantId() *string {
	return s.TenantId
}

func (s *GetSubnetResponseBodyContent) GetType() *string {
	return s.Type
}

func (s *GetSubnetResponseBodyContent) GetVpdBaseInfo() *GetSubnetResponseBodyContentVpdBaseInfo {
	return s.VpdBaseInfo
}

func (s *GetSubnetResponseBodyContent) GetVpdId() *string {
	return s.VpdId
}

func (s *GetSubnetResponseBodyContent) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetSubnetResponseBodyContent) SetAvailableIps(v int32) *GetSubnetResponseBodyContent {
	s.AvailableIps = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetCidr(v string) *GetSubnetResponseBodyContent {
	s.Cidr = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetCreateTime(v string) *GetSubnetResponseBodyContent {
	s.CreateTime = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetGmtModified(v string) *GetSubnetResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetLbCount(v int64) *GetSubnetResponseBodyContent {
	s.LbCount = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetMessage(v string) *GetSubnetResponseBodyContent {
	s.Message = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetNcCount(v int32) *GetSubnetResponseBodyContent {
	s.NcCount = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetNetworkInterfaceCount(v int32) *GetSubnetResponseBodyContent {
	s.NetworkInterfaceCount = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetPrivateIpCount(v int64) *GetSubnetResponseBodyContent {
	s.PrivateIpCount = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetRegionId(v string) *GetSubnetResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetResourceGroupId(v string) *GetSubnetResponseBodyContent {
	s.ResourceGroupId = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetStatus(v string) *GetSubnetResponseBodyContent {
	s.Status = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetSubnetId(v string) *GetSubnetResponseBodyContent {
	s.SubnetId = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetSubnetName(v string) *GetSubnetResponseBodyContent {
	s.SubnetName = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetTags(v []*GetSubnetResponseBodyContentTags) *GetSubnetResponseBodyContent {
	s.Tags = v
	return s
}

func (s *GetSubnetResponseBodyContent) SetTenantId(v string) *GetSubnetResponseBodyContent {
	s.TenantId = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetType(v string) *GetSubnetResponseBodyContent {
	s.Type = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetVpdBaseInfo(v *GetSubnetResponseBodyContentVpdBaseInfo) *GetSubnetResponseBodyContent {
	s.VpdBaseInfo = v
	return s
}

func (s *GetSubnetResponseBodyContent) SetVpdId(v string) *GetSubnetResponseBodyContent {
	s.VpdId = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetZoneId(v string) *GetSubnetResponseBodyContent {
	s.ZoneId = &v
	return s
}

func (s *GetSubnetResponseBodyContent) Validate() error {
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

type GetSubnetResponseBodyContentTags struct {
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

func (s GetSubnetResponseBodyContentTags) String() string {
	return dara.Prettify(s)
}

func (s GetSubnetResponseBodyContentTags) GoString() string {
	return s.String()
}

func (s *GetSubnetResponseBodyContentTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetSubnetResponseBodyContentTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetSubnetResponseBodyContentTags) SetTagKey(v string) *GetSubnetResponseBodyContentTags {
	s.TagKey = &v
	return s
}

func (s *GetSubnetResponseBodyContentTags) SetTagValue(v string) *GetSubnetResponseBodyContentTags {
	s.TagValue = &v
	return s
}

func (s *GetSubnetResponseBodyContentTags) Validate() error {
	return dara.Validate(s)
}

type GetSubnetResponseBodyContentVpdBaseInfo struct {
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
	// The time when the activation code was created.
	//
	// example:
	//
	// 1678273219000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the Lingjun CIDR block.
	//
	// example:
	//
	// vpd-iv2zm1qf
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The name of the Lingjun CIDR block.
	//
	// example:
	//
	// vpd-1
	VpdName *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
}

func (s GetSubnetResponseBodyContentVpdBaseInfo) String() string {
	return dara.Prettify(s)
}

func (s GetSubnetResponseBodyContentVpdBaseInfo) GoString() string {
	return s.String()
}

func (s *GetSubnetResponseBodyContentVpdBaseInfo) GetCidr() *string {
	return s.Cidr
}

func (s *GetSubnetResponseBodyContentVpdBaseInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetSubnetResponseBodyContentVpdBaseInfo) GetVpdId() *string {
	return s.VpdId
}

func (s *GetSubnetResponseBodyContentVpdBaseInfo) GetVpdName() *string {
	return s.VpdName
}

func (s *GetSubnetResponseBodyContentVpdBaseInfo) SetCidr(v string) *GetSubnetResponseBodyContentVpdBaseInfo {
	s.Cidr = &v
	return s
}

func (s *GetSubnetResponseBodyContentVpdBaseInfo) SetCreateTime(v string) *GetSubnetResponseBodyContentVpdBaseInfo {
	s.CreateTime = &v
	return s
}

func (s *GetSubnetResponseBodyContentVpdBaseInfo) SetVpdId(v string) *GetSubnetResponseBodyContentVpdBaseInfo {
	s.VpdId = &v
	return s
}

func (s *GetSubnetResponseBodyContentVpdBaseInfo) SetVpdName(v string) *GetSubnetResponseBodyContentVpdBaseInfo {
	s.VpdName = &v
	return s
}

func (s *GetSubnetResponseBodyContentVpdBaseInfo) Validate() error {
	return dara.Validate(s)
}
