// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetVpdResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *GetVpdResponseBody
	GetCode() *int32
	SetContent(v *GetVpdResponseBodyContent) *GetVpdResponseBody
	GetContent() *GetVpdResponseBodyContent
	SetMessage(v string) *GetVpdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetVpdResponseBody
	GetRequestId() *string
}

type GetVpdResponseBody struct {
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
	// The data returned.
	Content *GetVpdResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
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
	// A88DFED5-24B7-5A3E-87DE-380BF06F3C90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVpdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVpdResponseBody) GoString() string {
	return s.String()
}

func (s *GetVpdResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetVpdResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetVpdResponseBody) GetContent() *GetVpdResponseBodyContent {
	return s.Content
}

func (s *GetVpdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVpdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVpdResponseBody) SetAccessDeniedDetail(v string) *GetVpdResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetVpdResponseBody) SetCode(v int32) *GetVpdResponseBody {
	s.Code = &v
	return s
}

func (s *GetVpdResponseBody) SetContent(v *GetVpdResponseBodyContent) *GetVpdResponseBody {
	s.Content = v
	return s
}

func (s *GetVpdResponseBody) SetMessage(v string) *GetVpdResponseBody {
	s.Message = &v
	return s
}

func (s *GetVpdResponseBody) SetRequestId(v string) *GetVpdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVpdResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVpdResponseBodyContent struct {
	// Whether the Lingjun HUB(ER) has been bound.
	//
	// 	- **true**: ER is bound.
	//
	// 	- **false**: No ER is bound.
	//
	// example:
	//
	// true
	AttachErStatus *bool `json:"AttachErStatus,omitempty" xml:"AttachErStatus,omitempty"`
	// The CIDR block.
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
	// The information of the bound Lingjun HUB(ER).
	ErInfos []*GetVpdResponseBodyContentErInfos `json:"ErInfos,omitempty" xml:"ErInfos,omitempty" type:"Repeated"`
	// The time when the O\\&M task was modified.
	//
	// example:
	//
	// 2023-10-25 15:57:16
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of NCs.
	//
	// example:
	//
	// 16
	NcCount *int32 `json:"NcCount,omitempty" xml:"NcCount,omitempty"`
	// Number of Lingjun network interface controller.
	//
	// example:
	//
	// 1
	NetworkInterfaceCount *int32 `json:"NetworkInterfaceCount,omitempty" xml:"NetworkInterfaceCount,omitempty"`
	// The total number of secondary private IP addresses.
	//
	// example:
	//
	// 10
	PrivateIpCount *int64 `json:"PrivateIpCount,omitempty" xml:"PrivateIpCount,omitempty"`
	// The total quota information.
	//
	// example:
	//
	// 10
	Quota *int32 `json:"Quota,omitempty" xml:"Quota,omitempty"`
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
	// Internal Service CIDR block.
	//
	// example:
	//
	// 169.254.252.0/23
	ServiceCidr *string `json:"ServiceCidr,omitempty" xml:"ServiceCidr,omitempty"`
	// The current state of the instance.
	//
	// Valid value:
	//
	// 	- Not Available: Not Available.
	//
	// 	- Available: Normal: Available: Normal.
	//
	// 	- Deleting: Deleting: Deleting: Deleting.
	//
	// 	- Executing: executing: Executing: executing.
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
	SubnetCount *int64 `json:"SubnetCount,omitempty" xml:"SubnetCount,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags.
	Tags []*GetVpdResponseBodyContentTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The tenant ID.
	//
	// example:
	//
	// 1620939556166279
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The ID of the VPD instance.
	//
	// example:
	//
	// vpd-fuliephf
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The name of the Lingjun CIDR block.
	//
	// example:
	//
	// vpd-1
	VpdName *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
}

func (s GetVpdResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s GetVpdResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetVpdResponseBodyContent) GetAttachErStatus() *bool {
	return s.AttachErStatus
}

func (s *GetVpdResponseBodyContent) GetCidr() *string {
	return s.Cidr
}

func (s *GetVpdResponseBodyContent) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetVpdResponseBodyContent) GetErInfos() []*GetVpdResponseBodyContentErInfos {
	return s.ErInfos
}

func (s *GetVpdResponseBodyContent) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetVpdResponseBodyContent) GetMessage() *string {
	return s.Message
}

func (s *GetVpdResponseBodyContent) GetNcCount() *int32 {
	return s.NcCount
}

func (s *GetVpdResponseBodyContent) GetNetworkInterfaceCount() *int32 {
	return s.NetworkInterfaceCount
}

func (s *GetVpdResponseBodyContent) GetPrivateIpCount() *int64 {
	return s.PrivateIpCount
}

func (s *GetVpdResponseBodyContent) GetQuota() *int32 {
	return s.Quota
}

func (s *GetVpdResponseBodyContent) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVpdResponseBodyContent) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetVpdResponseBodyContent) GetSecondaryCidrBlocks() []*string {
	return s.SecondaryCidrBlocks
}

func (s *GetVpdResponseBodyContent) GetServiceCidr() *string {
	return s.ServiceCidr
}

func (s *GetVpdResponseBodyContent) GetStatus() *string {
	return s.Status
}

func (s *GetVpdResponseBodyContent) GetSubnetCount() *int64 {
	return s.SubnetCount
}

func (s *GetVpdResponseBodyContent) GetTags() []*GetVpdResponseBodyContentTags {
	return s.Tags
}

func (s *GetVpdResponseBodyContent) GetTenantId() *string {
	return s.TenantId
}

func (s *GetVpdResponseBodyContent) GetVpdId() *string {
	return s.VpdId
}

func (s *GetVpdResponseBodyContent) GetVpdName() *string {
	return s.VpdName
}

func (s *GetVpdResponseBodyContent) SetAttachErStatus(v bool) *GetVpdResponseBodyContent {
	s.AttachErStatus = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetCidr(v string) *GetVpdResponseBodyContent {
	s.Cidr = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetCreateTime(v string) *GetVpdResponseBodyContent {
	s.CreateTime = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetErInfos(v []*GetVpdResponseBodyContentErInfos) *GetVpdResponseBodyContent {
	s.ErInfos = v
	return s
}

func (s *GetVpdResponseBodyContent) SetGmtModified(v string) *GetVpdResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetMessage(v string) *GetVpdResponseBodyContent {
	s.Message = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetNcCount(v int32) *GetVpdResponseBodyContent {
	s.NcCount = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetNetworkInterfaceCount(v int32) *GetVpdResponseBodyContent {
	s.NetworkInterfaceCount = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetPrivateIpCount(v int64) *GetVpdResponseBodyContent {
	s.PrivateIpCount = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetQuota(v int32) *GetVpdResponseBodyContent {
	s.Quota = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetRegionId(v string) *GetVpdResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetResourceGroupId(v string) *GetVpdResponseBodyContent {
	s.ResourceGroupId = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetSecondaryCidrBlocks(v []*string) *GetVpdResponseBodyContent {
	s.SecondaryCidrBlocks = v
	return s
}

func (s *GetVpdResponseBodyContent) SetServiceCidr(v string) *GetVpdResponseBodyContent {
	s.ServiceCidr = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetStatus(v string) *GetVpdResponseBodyContent {
	s.Status = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetSubnetCount(v int64) *GetVpdResponseBodyContent {
	s.SubnetCount = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetTags(v []*GetVpdResponseBodyContentTags) *GetVpdResponseBodyContent {
	s.Tags = v
	return s
}

func (s *GetVpdResponseBodyContent) SetTenantId(v string) *GetVpdResponseBodyContent {
	s.TenantId = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetVpdId(v string) *GetVpdResponseBodyContent {
	s.VpdId = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetVpdName(v string) *GetVpdResponseBodyContent {
	s.VpdName = &v
	return s
}

func (s *GetVpdResponseBodyContent) Validate() error {
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

type GetVpdResponseBodyContentErInfos struct {
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
	// 1678273219000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the synchronization task.
	//
	// example:
	//
	// Restore verifying
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the Elastic Router (ER) instance.
	//
	// example:
	//
	// er-a7rqv1rq
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Elastic Router (ER) Instance Name
	//
	// example:
	//
	// er-1
	ErName *string `json:"ErName,omitempty" xml:"ErName,omitempty"`
	// The time when the O\\&M task was modified.
	//
	// example:
	//
	// 1678273219000
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
	// The ID of the region to which the Elastic Router (ER) belongs.
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
	// t464p4fql1bog
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetVpdResponseBodyContentErInfos) String() string {
	return dara.Prettify(s)
}

func (s GetVpdResponseBodyContentErInfos) GoString() string {
	return s.String()
}

func (s *GetVpdResponseBodyContentErInfos) GetConnections() *int64 {
	return s.Connections
}

func (s *GetVpdResponseBodyContentErInfos) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetVpdResponseBodyContentErInfos) GetDescription() *string {
	return s.Description
}

func (s *GetVpdResponseBodyContentErInfos) GetErId() *string {
	return s.ErId
}

func (s *GetVpdResponseBodyContentErInfos) GetErName() *string {
	return s.ErName
}

func (s *GetVpdResponseBodyContentErInfos) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetVpdResponseBodyContentErInfos) GetMasterZoneId() *string {
	return s.MasterZoneId
}

func (s *GetVpdResponseBodyContentErInfos) GetMessage() *string {
	return s.Message
}

func (s *GetVpdResponseBodyContentErInfos) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVpdResponseBodyContentErInfos) GetRouteMaps() *int64 {
	return s.RouteMaps
}

func (s *GetVpdResponseBodyContentErInfos) GetStatus() *string {
	return s.Status
}

func (s *GetVpdResponseBodyContentErInfos) GetTenantId() *string {
	return s.TenantId
}

func (s *GetVpdResponseBodyContentErInfos) SetConnections(v int64) *GetVpdResponseBodyContentErInfos {
	s.Connections = &v
	return s
}

func (s *GetVpdResponseBodyContentErInfos) SetCreateTime(v string) *GetVpdResponseBodyContentErInfos {
	s.CreateTime = &v
	return s
}

func (s *GetVpdResponseBodyContentErInfos) SetDescription(v string) *GetVpdResponseBodyContentErInfos {
	s.Description = &v
	return s
}

func (s *GetVpdResponseBodyContentErInfos) SetErId(v string) *GetVpdResponseBodyContentErInfos {
	s.ErId = &v
	return s
}

func (s *GetVpdResponseBodyContentErInfos) SetErName(v string) *GetVpdResponseBodyContentErInfos {
	s.ErName = &v
	return s
}

func (s *GetVpdResponseBodyContentErInfos) SetGmtModified(v string) *GetVpdResponseBodyContentErInfos {
	s.GmtModified = &v
	return s
}

func (s *GetVpdResponseBodyContentErInfos) SetMasterZoneId(v string) *GetVpdResponseBodyContentErInfos {
	s.MasterZoneId = &v
	return s
}

func (s *GetVpdResponseBodyContentErInfos) SetMessage(v string) *GetVpdResponseBodyContentErInfos {
	s.Message = &v
	return s
}

func (s *GetVpdResponseBodyContentErInfos) SetRegionId(v string) *GetVpdResponseBodyContentErInfos {
	s.RegionId = &v
	return s
}

func (s *GetVpdResponseBodyContentErInfos) SetRouteMaps(v int64) *GetVpdResponseBodyContentErInfos {
	s.RouteMaps = &v
	return s
}

func (s *GetVpdResponseBodyContentErInfos) SetStatus(v string) *GetVpdResponseBodyContentErInfos {
	s.Status = &v
	return s
}

func (s *GetVpdResponseBodyContentErInfos) SetTenantId(v string) *GetVpdResponseBodyContentErInfos {
	s.TenantId = &v
	return s
}

func (s *GetVpdResponseBodyContentErInfos) Validate() error {
	return dara.Validate(s)
}

type GetVpdResponseBodyContentTags struct {
	// The tag key.
	//
	// You cannot specify an empty string as a tag key. It can be up to 64 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// tag-subent-region
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

func (s GetVpdResponseBodyContentTags) String() string {
	return dara.Prettify(s)
}

func (s GetVpdResponseBodyContentTags) GoString() string {
	return s.String()
}

func (s *GetVpdResponseBodyContentTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetVpdResponseBodyContentTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetVpdResponseBodyContentTags) SetTagKey(v string) *GetVpdResponseBodyContentTags {
	s.TagKey = &v
	return s
}

func (s *GetVpdResponseBodyContentTags) SetTagValue(v string) *GetVpdResponseBodyContentTags {
	s.TagValue = &v
	return s
}

func (s *GetVpdResponseBodyContentTags) Validate() error {
	return dara.Validate(s)
}
