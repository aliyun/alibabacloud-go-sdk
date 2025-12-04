// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpdsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListVpdsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *ListVpdsResponseBody
	GetCode() *int32
	SetContent(v *ListVpdsResponseBodyContent) *ListVpdsResponseBody
	GetContent() *ListVpdsResponseBodyContent
	SetMessage(v string) *ListVpdsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListVpdsResponseBody
	GetRequestId() *string
}

type ListVpdsResponseBody struct {
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
	Content *ListVpdsResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9C50C9CD-E799-54DA-BA7A-1FAF3DF80857
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVpdsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVpdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpdsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListVpdsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListVpdsResponseBody) GetContent() *ListVpdsResponseBodyContent {
	return s.Content
}

func (s *ListVpdsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVpdsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVpdsResponseBody) SetAccessDeniedDetail(v string) *ListVpdsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListVpdsResponseBody) SetCode(v int32) *ListVpdsResponseBody {
	s.Code = &v
	return s
}

func (s *ListVpdsResponseBody) SetContent(v *ListVpdsResponseBodyContent) *ListVpdsResponseBody {
	s.Content = v
	return s
}

func (s *ListVpdsResponseBody) SetMessage(v string) *ListVpdsResponseBody {
	s.Message = &v
	return s
}

func (s *ListVpdsResponseBody) SetRequestId(v string) *ListVpdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpdsResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVpdsResponseBodyContent struct {
	// The returned data.
	Data []*ListVpdsResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries that are returned.
	//
	// example:
	//
	// 0
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListVpdsResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListVpdsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListVpdsResponseBodyContent) GetData() []*ListVpdsResponseBodyContentData {
	return s.Data
}

func (s *ListVpdsResponseBodyContent) GetTotal() *int64 {
	return s.Total
}

func (s *ListVpdsResponseBodyContent) SetData(v []*ListVpdsResponseBodyContentData) *ListVpdsResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListVpdsResponseBodyContent) SetTotal(v int64) *ListVpdsResponseBodyContent {
	s.Total = &v
	return s
}

func (s *ListVpdsResponseBodyContent) Validate() error {
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

type ListVpdsResponseBodyContentData struct {
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
	// Dependencies.
	Dependence map[string]interface{} `json:"Dependence,omitempty" xml:"Dependence,omitempty"`
	// The information list of the bound Lingjun HUB(ER).
	ErInfos []*ListVpdsResponseBodyContentDataErInfos `json:"ErInfos,omitempty" xml:"ErInfos,omitempty" type:"Repeated"`
	// The time when the O\\&M task was modified.
	//
	// example:
	//
	// 1678273219000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// nc quantity.
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
	// The list of additional CIDR blocks.
	SecondaryCidrBlocks []*string `json:"SecondaryCidrBlocks,omitempty" xml:"SecondaryCidrBlocks,omitempty" type:"Repeated"`
	// The Service CIDR block.
	//
	// example:
	//
	// 169.254.252.0/23
	ServiceCidr *string `json:"ServiceCidr,omitempty" xml:"ServiceCidr,omitempty"`
	// The task status.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of subnets.
	//
	// example:
	//
	// 1
	SubnetCount *int32 `json:"SubnetCount,omitempty" xml:"SubnetCount,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags.
	Tags []*ListVpdsResponseBodyContentDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The tenant ID.
	//
	// example:
	//
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The ID of the VPD instance.
	//
	// example:
	//
	// vpd-lg4dppgi
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The name of the VPD.
	//
	// example:
	//
	// vpd-1
	VpdName *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
}

func (s ListVpdsResponseBodyContentData) String() string {
	return dara.Prettify(s)
}

func (s ListVpdsResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListVpdsResponseBodyContentData) GetCidr() *string {
	return s.Cidr
}

func (s *ListVpdsResponseBodyContentData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListVpdsResponseBodyContentData) GetDependence() map[string]interface{} {
	return s.Dependence
}

func (s *ListVpdsResponseBodyContentData) GetErInfos() []*ListVpdsResponseBodyContentDataErInfos {
	return s.ErInfos
}

func (s *ListVpdsResponseBodyContentData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListVpdsResponseBodyContentData) GetMessage() *string {
	return s.Message
}

func (s *ListVpdsResponseBodyContentData) GetNcCount() *int32 {
	return s.NcCount
}

func (s *ListVpdsResponseBodyContentData) GetNetworkInterfaceCount() *int32 {
	return s.NetworkInterfaceCount
}

func (s *ListVpdsResponseBodyContentData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpdsResponseBodyContentData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVpdsResponseBodyContentData) GetSecondaryCidrBlocks() []*string {
	return s.SecondaryCidrBlocks
}

func (s *ListVpdsResponseBodyContentData) GetServiceCidr() *string {
	return s.ServiceCidr
}

func (s *ListVpdsResponseBodyContentData) GetStatus() *string {
	return s.Status
}

func (s *ListVpdsResponseBodyContentData) GetSubnetCount() *int32 {
	return s.SubnetCount
}

func (s *ListVpdsResponseBodyContentData) GetTags() []*ListVpdsResponseBodyContentDataTags {
	return s.Tags
}

func (s *ListVpdsResponseBodyContentData) GetTenantId() *string {
	return s.TenantId
}

func (s *ListVpdsResponseBodyContentData) GetVpdId() *string {
	return s.VpdId
}

func (s *ListVpdsResponseBodyContentData) GetVpdName() *string {
	return s.VpdName
}

func (s *ListVpdsResponseBodyContentData) SetCidr(v string) *ListVpdsResponseBodyContentData {
	s.Cidr = &v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetCreateTime(v string) *ListVpdsResponseBodyContentData {
	s.CreateTime = &v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetDependence(v map[string]interface{}) *ListVpdsResponseBodyContentData {
	s.Dependence = v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetErInfos(v []*ListVpdsResponseBodyContentDataErInfos) *ListVpdsResponseBodyContentData {
	s.ErInfos = v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetGmtModified(v string) *ListVpdsResponseBodyContentData {
	s.GmtModified = &v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetMessage(v string) *ListVpdsResponseBodyContentData {
	s.Message = &v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetNcCount(v int32) *ListVpdsResponseBodyContentData {
	s.NcCount = &v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetNetworkInterfaceCount(v int32) *ListVpdsResponseBodyContentData {
	s.NetworkInterfaceCount = &v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetRegionId(v string) *ListVpdsResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetResourceGroupId(v string) *ListVpdsResponseBodyContentData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetSecondaryCidrBlocks(v []*string) *ListVpdsResponseBodyContentData {
	s.SecondaryCidrBlocks = v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetServiceCidr(v string) *ListVpdsResponseBodyContentData {
	s.ServiceCidr = &v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetStatus(v string) *ListVpdsResponseBodyContentData {
	s.Status = &v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetSubnetCount(v int32) *ListVpdsResponseBodyContentData {
	s.SubnetCount = &v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetTags(v []*ListVpdsResponseBodyContentDataTags) *ListVpdsResponseBodyContentData {
	s.Tags = v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetTenantId(v string) *ListVpdsResponseBodyContentData {
	s.TenantId = &v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetVpdId(v string) *ListVpdsResponseBodyContentData {
	s.VpdId = &v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetVpdName(v string) *ListVpdsResponseBodyContentData {
	s.VpdName = &v
	return s
}

func (s *ListVpdsResponseBodyContentData) Validate() error {
	if s.ErInfos != nil {
		for _, item := range s.ErInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type ListVpdsResponseBodyContentDataErInfos struct {
	// The number of connections.
	//
	// example:
	//
	// 2
	Connections *int64 `json:"Connections,omitempty" xml:"Connections,omitempty"`
	// The time when the activation code was created.
	//
	// example:
	//
	// 2023-12-26 20:16:36
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the synchronization task.
	//
	// example:
	//
	// No description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the Elastic Router (ER) instance.
	//
	// example:
	//
	// er-63vzm0fw
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The name of the Lingjun HUB(ER) instance.
	//
	// example:
	//
	// er-1
	ErName *string `json:"ErName,omitempty" xml:"ErName,omitempty"`
	// The time when the O\\&M task was modified.
	//
	// example:
	//
	// 2023-12-26 20:16:36
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The primary zone.
	//
	// example:
	//
	// cn-wulanchabu-b
	MasterZoneId *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The supported region.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of routing policy.
	//
	// example:
	//
	// 2
	RouteMaps *int64 `json:"RouteMaps,omitempty" xml:"RouteMaps,omitempty"`
	// The task status.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 1620939556166277
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListVpdsResponseBodyContentDataErInfos) String() string {
	return dara.Prettify(s)
}

func (s ListVpdsResponseBodyContentDataErInfos) GoString() string {
	return s.String()
}

func (s *ListVpdsResponseBodyContentDataErInfos) GetConnections() *int64 {
	return s.Connections
}

func (s *ListVpdsResponseBodyContentDataErInfos) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListVpdsResponseBodyContentDataErInfos) GetDescription() *string {
	return s.Description
}

func (s *ListVpdsResponseBodyContentDataErInfos) GetErId() *string {
	return s.ErId
}

func (s *ListVpdsResponseBodyContentDataErInfos) GetErName() *string {
	return s.ErName
}

func (s *ListVpdsResponseBodyContentDataErInfos) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListVpdsResponseBodyContentDataErInfos) GetMasterZoneId() *string {
	return s.MasterZoneId
}

func (s *ListVpdsResponseBodyContentDataErInfos) GetMessage() *string {
	return s.Message
}

func (s *ListVpdsResponseBodyContentDataErInfos) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpdsResponseBodyContentDataErInfos) GetRouteMaps() *int64 {
	return s.RouteMaps
}

func (s *ListVpdsResponseBodyContentDataErInfos) GetStatus() *string {
	return s.Status
}

func (s *ListVpdsResponseBodyContentDataErInfos) GetTenantId() *string {
	return s.TenantId
}

func (s *ListVpdsResponseBodyContentDataErInfos) SetConnections(v int64) *ListVpdsResponseBodyContentDataErInfos {
	s.Connections = &v
	return s
}

func (s *ListVpdsResponseBodyContentDataErInfos) SetCreateTime(v string) *ListVpdsResponseBodyContentDataErInfos {
	s.CreateTime = &v
	return s
}

func (s *ListVpdsResponseBodyContentDataErInfos) SetDescription(v string) *ListVpdsResponseBodyContentDataErInfos {
	s.Description = &v
	return s
}

func (s *ListVpdsResponseBodyContentDataErInfos) SetErId(v string) *ListVpdsResponseBodyContentDataErInfos {
	s.ErId = &v
	return s
}

func (s *ListVpdsResponseBodyContentDataErInfos) SetErName(v string) *ListVpdsResponseBodyContentDataErInfos {
	s.ErName = &v
	return s
}

func (s *ListVpdsResponseBodyContentDataErInfos) SetGmtModified(v string) *ListVpdsResponseBodyContentDataErInfos {
	s.GmtModified = &v
	return s
}

func (s *ListVpdsResponseBodyContentDataErInfos) SetMasterZoneId(v string) *ListVpdsResponseBodyContentDataErInfos {
	s.MasterZoneId = &v
	return s
}

func (s *ListVpdsResponseBodyContentDataErInfos) SetMessage(v string) *ListVpdsResponseBodyContentDataErInfos {
	s.Message = &v
	return s
}

func (s *ListVpdsResponseBodyContentDataErInfos) SetRegionId(v string) *ListVpdsResponseBodyContentDataErInfos {
	s.RegionId = &v
	return s
}

func (s *ListVpdsResponseBodyContentDataErInfos) SetRouteMaps(v int64) *ListVpdsResponseBodyContentDataErInfos {
	s.RouteMaps = &v
	return s
}

func (s *ListVpdsResponseBodyContentDataErInfos) SetStatus(v string) *ListVpdsResponseBodyContentDataErInfos {
	s.Status = &v
	return s
}

func (s *ListVpdsResponseBodyContentDataErInfos) SetTenantId(v string) *ListVpdsResponseBodyContentDataErInfos {
	s.TenantId = &v
	return s
}

func (s *ListVpdsResponseBodyContentDataErInfos) Validate() error {
	return dara.Validate(s)
}

type ListVpdsResponseBodyContentDataTags struct {
	// The tag key.
	//
	// You cannot specify an empty string as a tag key. It can be up to 64 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// tag-vpd-region
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag that is added to the resource.
	//
	// The tag value can be empty or a string of up to 128 characters. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// Each key-value pair must be unique. You can specify values for at most 20 tag keys in each call.
	//
	// example:
	//
	// cn-wulanchabu
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListVpdsResponseBodyContentDataTags) String() string {
	return dara.Prettify(s)
}

func (s ListVpdsResponseBodyContentDataTags) GoString() string {
	return s.String()
}

func (s *ListVpdsResponseBodyContentDataTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListVpdsResponseBodyContentDataTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListVpdsResponseBodyContentDataTags) SetTagKey(v string) *ListVpdsResponseBodyContentDataTags {
	s.TagKey = &v
	return s
}

func (s *ListVpdsResponseBodyContentDataTags) SetTagValue(v string) *ListVpdsResponseBodyContentDataTags {
	s.TagValue = &v
	return s
}

func (s *ListVpdsResponseBodyContentDataTags) Validate() error {
	return dara.Validate(s)
}
