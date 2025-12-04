// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetElasticNetworkInterfaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetElasticNetworkInterfaceResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *GetElasticNetworkInterfaceResponseBody
	GetCode() *int32
	SetContent(v *GetElasticNetworkInterfaceResponseBodyContent) *GetElasticNetworkInterfaceResponseBody
	GetContent() *GetElasticNetworkInterfaceResponseBodyContent
	SetMessage(v string) *GetElasticNetworkInterfaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetElasticNetworkInterfaceResponseBody
	GetRequestId() *string
}

type GetElasticNetworkInterfaceResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *GetElasticNetworkInterfaceResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0901F411-28FA-5B9C-BAEE-7776463FF0DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetElasticNetworkInterfaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetElasticNetworkInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetElasticNetworkInterfaceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetElasticNetworkInterfaceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetElasticNetworkInterfaceResponseBody) GetContent() *GetElasticNetworkInterfaceResponseBodyContent {
	return s.Content
}

func (s *GetElasticNetworkInterfaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetElasticNetworkInterfaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetElasticNetworkInterfaceResponseBody) SetAccessDeniedDetail(v string) *GetElasticNetworkInterfaceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBody) SetCode(v int32) *GetElasticNetworkInterfaceResponseBody {
	s.Code = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBody) SetContent(v *GetElasticNetworkInterfaceResponseBodyContent) *GetElasticNetworkInterfaceResponseBody {
	s.Content = v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBody) SetMessage(v string) *GetElasticNetworkInterfaceResponseBody {
	s.Message = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBody) SetRequestId(v string) *GetElasticNetworkInterfaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetElasticNetworkInterfaceResponseBodyContent struct {
	// The time when the data address was created.
	//
	// example:
	//
	// 2022-01-13 12:51:41
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance description.
	//
	// example:
	//
	// terraform-example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Lingjun Elastic Network Interface ID
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// Whether to enable the jumboFrame capability
	//
	// example:
	//
	// True
	EnableJumboFrame *bool `json:"EnableJumboFrame,omitempty" xml:"EnableJumboFrame,omitempty"`
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
	// 2022-01-13 12:51:41
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Elastic Network Interface IP
	//
	// example:
	//
	// 203.107.****
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// IPV6 address
	Ipv6Addresses []*GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses `json:"Ipv6Addresses,omitempty" xml:"Ipv6Addresses,omitempty" type:"Repeated"`
	// mac address
	//
	// example:
	//
	// 00:22:6D:97:**:**
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
	// Lingjun Node ID
	//
	// example:
	//
	// e01-cn-lbj3aej****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// Secondary private IP address
	PrivateIpAddresses []*GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses `json:"PrivateIpAddresses,omitempty" xml:"PrivateIpAddresses,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 资源组实例ID
	//
	// example:
	//
	// rg-acfmzzka6bnjvbi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-0jl5s4p4laalruk7****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The state of the private gateway.
	//
	// Valid value:
	//
	// 	- Create Failed: the creation failure.
	//
	// 	- Delete Failed: the that failed to be deleted.
	//
	// 	- Executing
	//
	// 	- Available
	//
	// 	- Deleting
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The details of the resource tags.
	Tags []*GetElasticNetworkInterfaceResponseBodyContentTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// NIC Type
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
	// VPC ID
	//
	// example:
	//
	// vpc-j6ctp4n75306****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetElasticNetworkInterfaceResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s GetElasticNetworkInterfaceResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) GetDescription() *string {
	return s.Description
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) GetElasticNetworkInterfaceId() *string {
	return s.ElasticNetworkInterfaceId
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) GetEnableJumboFrame() *bool {
	return s.EnableJumboFrame
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) GetGateway() *string {
	return s.Gateway
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) GetIp() *string {
	return s.Ip
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) GetIpv6Addresses() []*GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses {
	return s.Ipv6Addresses
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) GetMac() *string {
	return s.Mac
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) GetMask() *string {
	return s.Mask
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) GetMessage() *string {
	return s.Message
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) GetNodeId() *string {
	return s.NodeId
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) GetPrivateIpAddresses() []*GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses {
	return s.PrivateIpAddresses
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) GetRegionId() *string {
	return s.RegionId
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) GetStatus() *string {
	return s.Status
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) GetTags() []*GetElasticNetworkInterfaceResponseBodyContentTags {
	return s.Tags
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) GetType() *string {
	return s.Type
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) GetVpcId() *string {
	return s.VpcId
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetCreateTime(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.CreateTime = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetDescription(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.Description = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetElasticNetworkInterfaceId(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetEnableJumboFrame(v bool) *GetElasticNetworkInterfaceResponseBodyContent {
	s.EnableJumboFrame = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetGateway(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.Gateway = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetGmtModified(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetIp(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.Ip = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetIpv6Addresses(v []*GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) *GetElasticNetworkInterfaceResponseBodyContent {
	s.Ipv6Addresses = v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetMac(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.Mac = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetMask(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.Mask = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetMessage(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.Message = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetNodeId(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.NodeId = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetPrivateIpAddresses(v []*GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) *GetElasticNetworkInterfaceResponseBodyContent {
	s.PrivateIpAddresses = v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetRegionId(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetResourceGroupId(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.ResourceGroupId = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetSecurityGroupId(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.SecurityGroupId = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetStatus(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.Status = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetTags(v []*GetElasticNetworkInterfaceResponseBodyContentTags) *GetElasticNetworkInterfaceResponseBodyContent {
	s.Tags = v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetType(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.Type = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetVSwitchId(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.VSwitchId = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetVpcId(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.VpcId = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetZoneId(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.ZoneId = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) Validate() error {
	if s.Ipv6Addresses != nil {
		for _, item := range s.Ipv6Addresses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PrivateIpAddresses != nil {
		for _, item := range s.PrivateIpAddresses {
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

type GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses struct {
	// The instance description.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Lingjun Elastic Network Interface ID
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1585816811000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the cluster was updated.
	//
	// example:
	//
	// 1549012834000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// IPV6 unique identifier
	//
	// example:
	//
	// sip-sg3xabeq
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// IPV6 address
	//
	// example:
	//
	// 2408:4005:3aa:1000:470d:66fb:56a5:****
	Ipv6Address *string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) String() string {
	return dara.Prettify(s)
}

func (s GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) GoString() string {
	return s.String()
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) GetDescription() *string {
	return s.Description
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) GetElasticNetworkInterfaceId() *string {
	return s.ElasticNetworkInterfaceId
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) GetIpName() *string {
	return s.IpName
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) GetIpv6Address() *string {
	return s.Ipv6Address
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) GetMessage() *string {
	return s.Message
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) GetRegionId() *string {
	return s.RegionId
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) GetStatus() *string {
	return s.Status
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) SetDescription(v string) *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses {
	s.Description = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) SetElasticNetworkInterfaceId(v string) *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) SetGmtCreate(v string) *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses {
	s.GmtCreate = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) SetGmtModified(v string) *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses {
	s.GmtModified = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) SetIpName(v string) *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses {
	s.IpName = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) SetIpv6Address(v string) *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses {
	s.Ipv6Address = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) SetMessage(v string) *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses {
	s.Message = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) SetRegionId(v string) *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses {
	s.RegionId = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) SetStatus(v string) *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses {
	s.Status = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) Validate() error {
	return dara.Validate(s)
}

type GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses struct {
	// The instance description.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Lingjun Elastic Network Interface ID
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1672971789000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the cluster was updated.
	//
	// example:
	//
	// 1672971789000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Lingjun Elastic Network Interface Secondary Private IP Unique Identifier
	//
	// example:
	//
	// sip-ywz****
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Lingjun Elastic Network Interface secondary private IP address
	//
	// example:
	//
	// 172.16.****
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) String() string {
	return dara.Prettify(s)
}

func (s GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) GoString() string {
	return s.String()
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) GetDescription() *string {
	return s.Description
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) GetElasticNetworkInterfaceId() *string {
	return s.ElasticNetworkInterfaceId
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) GetIpName() *string {
	return s.IpName
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) GetMessage() *string {
	return s.Message
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) GetRegionId() *string {
	return s.RegionId
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) GetStatus() *string {
	return s.Status
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) SetDescription(v string) *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses {
	s.Description = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) SetElasticNetworkInterfaceId(v string) *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) SetGmtCreate(v string) *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses {
	s.GmtCreate = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) SetGmtModified(v string) *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses {
	s.GmtModified = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) SetIpName(v string) *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses {
	s.IpName = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) SetMessage(v string) *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses {
	s.Message = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) SetPrivateIpAddress(v string) *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses {
	s.PrivateIpAddress = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) SetRegionId(v string) *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses {
	s.RegionId = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) SetStatus(v string) *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses {
	s.Status = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) Validate() error {
	return dara.Validate(s)
}

type GetElasticNetworkInterfaceResponseBodyContentTags struct {
	// The tag key.
	//
	// example:
	//
	// testkey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// testValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetElasticNetworkInterfaceResponseBodyContentTags) String() string {
	return dara.Prettify(s)
}

func (s GetElasticNetworkInterfaceResponseBodyContentTags) GoString() string {
	return s.String()
}

func (s *GetElasticNetworkInterfaceResponseBodyContentTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetElasticNetworkInterfaceResponseBodyContentTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetElasticNetworkInterfaceResponseBodyContentTags) SetTagKey(v string) *GetElasticNetworkInterfaceResponseBodyContentTags {
	s.TagKey = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentTags) SetTagValue(v string) *GetElasticNetworkInterfaceResponseBodyContentTags {
	s.TagValue = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentTags) Validate() error {
	return dara.Validate(s)
}
