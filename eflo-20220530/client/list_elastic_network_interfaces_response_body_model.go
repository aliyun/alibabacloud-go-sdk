// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListElasticNetworkInterfacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListElasticNetworkInterfacesResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *ListElasticNetworkInterfacesResponseBody
	GetCode() *int32
	SetContent(v *ListElasticNetworkInterfacesResponseBodyContent) *ListElasticNetworkInterfacesResponseBody
	GetContent() *ListElasticNetworkInterfacesResponseBodyContent
	SetMessage(v string) *ListElasticNetworkInterfacesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListElasticNetworkInterfacesResponseBody
	GetRequestId() *string
}

type ListElasticNetworkInterfacesResponseBody struct {
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
	Content *ListElasticNetworkInterfacesResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// AC8C713A-A9F4-5984-A5E1-76496DF35153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListElasticNetworkInterfacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListElasticNetworkInterfacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListElasticNetworkInterfacesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListElasticNetworkInterfacesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListElasticNetworkInterfacesResponseBody) GetContent() *ListElasticNetworkInterfacesResponseBodyContent {
	return s.Content
}

func (s *ListElasticNetworkInterfacesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListElasticNetworkInterfacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListElasticNetworkInterfacesResponseBody) SetAccessDeniedDetail(v string) *ListElasticNetworkInterfacesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBody) SetCode(v int32) *ListElasticNetworkInterfacesResponseBody {
	s.Code = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBody) SetContent(v *ListElasticNetworkInterfacesResponseBodyContent) *ListElasticNetworkInterfacesResponseBody {
	s.Content = v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBody) SetMessage(v string) *ListElasticNetworkInterfacesResponseBody {
	s.Message = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBody) SetRequestId(v string) *ListElasticNetworkInterfacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListElasticNetworkInterfacesResponseBodyContent struct {
	// lingjun Elastic Network Interface information list
	Data []*ListElasticNetworkInterfacesResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListElasticNetworkInterfacesResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListElasticNetworkInterfacesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListElasticNetworkInterfacesResponseBodyContent) GetData() []*ListElasticNetworkInterfacesResponseBodyContentData {
	return s.Data
}

func (s *ListElasticNetworkInterfacesResponseBodyContent) GetTotal() *int64 {
	return s.Total
}

func (s *ListElasticNetworkInterfacesResponseBodyContent) SetData(v []*ListElasticNetworkInterfacesResponseBodyContentData) *ListElasticNetworkInterfacesResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContent) SetTotal(v int64) *ListElasticNetworkInterfacesResponseBodyContent {
	s.Total = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContent) Validate() error {
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

type ListElasticNetworkInterfacesResponseBodyContentData struct {
	// The time when the data address was created.
	//
	// example:
	//
	// 1601176751000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance description.
	//
	// example:
	//
	// No description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Lingjun Elastic Network Interface ID
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// vswitch gateway address
	//
	// example:
	//
	// 172.16.****
	Gateway *string `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// The time when the agent was last modified.
	//
	// example:
	//
	// 1640187007000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The IP address of the BE cluster.
	//
	// example:
	//
	// 10.0.0.13
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// mac address
	//
	// example:
	//
	// E0:01:A6:4A:6A:D0
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// vswitch mask bits
	//
	// example:
	//
	// 24
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// e01-cn-uax37m1****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2l4sq6l7unhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-f8z4wr1b41x3qsc9****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags.
	Tags []*ListElasticNetworkInterfacesResponseBodyContentDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// network interface controller type, the default type DEFAULT cannot be manually released
	//
	// Valid value:
	//
	// 	- CUSTOM: custom type.
	//
	// 	- DEFAULT: system type.
	//
	// example:
	//
	// DEFAULT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-uf6u8473r84e9****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-f8ziirfl9k25h2qn7****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListElasticNetworkInterfacesResponseBodyContentData) String() string {
	return dara.Prettify(s)
}

func (s ListElasticNetworkInterfacesResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) GetDescription() *string {
	return s.Description
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) GetElasticNetworkInterfaceId() *string {
	return s.ElasticNetworkInterfaceId
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) GetGateway() *string {
	return s.Gateway
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) GetIp() *string {
	return s.Ip
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) GetMac() *string {
	return s.Mac
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) GetMask() *string {
	return s.Mask
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) GetMessage() *string {
	return s.Message
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) GetNodeId() *string {
	return s.NodeId
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) GetStatus() *string {
	return s.Status
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) GetTags() []*ListElasticNetworkInterfacesResponseBodyContentDataTags {
	return s.Tags
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) GetType() *string {
	return s.Type
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) GetVpcId() *string {
	return s.VpcId
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetCreateTime(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.CreateTime = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetDescription(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.Description = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetElasticNetworkInterfaceId(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetGateway(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.Gateway = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetGmtModified(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.GmtModified = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetIp(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.Ip = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetMac(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.Mac = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetMask(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.Mask = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetMessage(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.Message = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetNodeId(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.NodeId = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetRegionId(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetResourceGroupId(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetSecurityGroupId(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.SecurityGroupId = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetStatus(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.Status = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetTags(v []*ListElasticNetworkInterfacesResponseBodyContentDataTags) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.Tags = v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetType(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.Type = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetVSwitchId(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.VSwitchId = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetVpcId(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.VpcId = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetZoneId(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.ZoneId = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) Validate() error {
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

type ListElasticNetworkInterfacesResponseBodyContentDataTags struct {
	// The tag key.
	//
	// example:
	//
	// key-test
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// Tag value.
	//
	// example:
	//
	// key-value
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListElasticNetworkInterfacesResponseBodyContentDataTags) String() string {
	return dara.Prettify(s)
}

func (s ListElasticNetworkInterfacesResponseBodyContentDataTags) GoString() string {
	return s.String()
}

func (s *ListElasticNetworkInterfacesResponseBodyContentDataTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListElasticNetworkInterfacesResponseBodyContentDataTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListElasticNetworkInterfacesResponseBodyContentDataTags) SetTagKey(v string) *ListElasticNetworkInterfacesResponseBodyContentDataTags {
	s.TagKey = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentDataTags) SetTagValue(v string) *ListElasticNetworkInterfacesResponseBodyContentDataTags {
	s.TagValue = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentDataTags) Validate() error {
	return dara.Validate(s)
}
