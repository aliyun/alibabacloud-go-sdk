// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVccResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetVccResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *GetVccResponseBody
	GetCode() *int32
	SetContent(v *GetVccResponseBodyContent) *GetVccResponseBody
	GetContent() *GetVccResponseBodyContent
	SetMessage(v string) *GetVccResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetVccResponseBody
	GetRequestId() *string
}

type GetVccResponseBody struct {
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
	Content *GetVccResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// CAD09E47-B651-5206-B2DC-3AB78C8EB446
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVccResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVccResponseBody) GoString() string {
	return s.String()
}

func (s *GetVccResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetVccResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetVccResponseBody) GetContent() *GetVccResponseBodyContent {
	return s.Content
}

func (s *GetVccResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVccResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVccResponseBody) SetAccessDeniedDetail(v string) *GetVccResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetVccResponseBody) SetCode(v int32) *GetVccResponseBody {
	s.Code = &v
	return s
}

func (s *GetVccResponseBody) SetContent(v *GetVccResponseBodyContent) *GetVccResponseBody {
	s.Content = v
	return s
}

func (s *GetVccResponseBody) SetMessage(v string) *GetVccResponseBody {
	s.Message = &v
	return s
}

func (s *GetVccResponseBody) SetRequestId(v string) *GetVccResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVccResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVccResponseBodyContent struct {
	// Express Connect circuit access point ID:
	//
	// 	- **ap-cn-wulanchabu-jn-ts-A**: Ulanqab-Jining-A
	//
	// 	- **ap-cn-heyuan-yc-ts-SA127**: Heyuan-Yuancheng-A
	//
	// example:
	//
	// ap-cn-wulanchabu-jn-ts-A
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// Alibaba Cloud route information list
	AliyunRouterInfo []*GetVccResponseBodyContentAliyunRouterInfo `json:"AliyunRouterInfo,omitempty" xml:"AliyunRouterInfo,omitempty" type:"Repeated"`
	// Whether Lingjun HUB has been bound to a network instance
	//
	// 	- **true**: Bound
	//
	// 	- **false**: unbound
	//
	// example:
	//
	// true
	AttachErStatus *bool `json:"AttachErStatus,omitempty" xml:"AttachErStatus,omitempty"`
	// bandwidth
	//
	// example:
	//
	// 20
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The bandwidth of the port.
	//
	// example:
	//
	// 1G
	BandwidthStr *string `json:"BandwidthStr,omitempty" xml:"BandwidthStr,omitempty"`
	// BGP AS number
	//
	// example:
	//
	// 45644
	BgpAsn *string `json:"BgpAsn,omitempty" xml:"BgpAsn,omitempty"`
	// BGP CIDR block
	//
	// example:
	//
	// 10.4.0.0/24
	BgpCidr *string `json:"BgpCidr,omitempty" xml:"BgpCidr,omitempty"`
	// The ID of the CEN instance; [What is the CEN?](https://help.aliyun.com/document_detail/181681.html)
	//
	// You can call the [DescribeCens](https://help.aliyun.com/document_detail/468215.htm) to query the information of CEN instances under the current Alibaba Cloud account.
	//
	// example:
	//
	// cen-m2iskbojlvda5w65fp
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// Account to which the CEN belongs
	//
	// example:
	//
	// 1620939556166279
	CenOwnerId *string `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	// Lingjun Network Routing Information List
	CisRouterInfo []*GetVccResponseBodyContentCisRouterInfo `json:"CisRouterInfo,omitempty" xml:"CisRouterInfo,omitempty" type:"Repeated"`
	// Commodity code
	//
	// example:
	//
	// bccluster_cloudconnectionpre_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The connection mode. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **CENTR**
	//
	// example:
	//
	// CENTR
	ConnectionType *string `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1648085472000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Current Node
	//
	// example:
	//
	// task-xxx-node-x
	CurrentNode *string `json:"CurrentNode,omitempty" xml:"CurrentNode,omitempty"`
	// Cycle
	//
	// example:
	//
	// 1
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// List of bound Lingjun HUB information
	ErInfos []*GetVccResponseBodyContentErInfos `json:"ErInfos,omitempty" xml:"ErInfos,omitempty" type:"Repeated"`
	// The time when the application expired.
	//
	// example:
	//
	// 1678379917000
	ExpirationDate *string `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	// The time when the agent was last modified.
	//
	// example:
	//
	// 1648085472000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The billing method for network usage.
	//
	// 	- **PayByTraffic**: pay-by-traffic
	//
	// 	- **PayByBandwidth**: pay-by-bandwidth
	//
	// example:
	//
	// PayByBandwidth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The connectivity provider of the Express Connect circuit. Valid values:
	//
	// 	- **CO**: other connectivity providers in the Chinese mainland
	//
	// example:
	//
	// CO
	LineOperator *string `json:"LineOperator,omitempty" xml:"LineOperator,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// You don\\"t have the permission of this operation, action=eflo:GetVcc, arn=acs:eflo:cn-heyuan:1263399219805497:vcc/vcc-cn-fhh3yxjwe01, resourceGroup=null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **PREPAY**: subscription
	//
	// 	- **POSTPAY**: pay-as-you-go
	//
	// example:
	//
	// PrePay
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The port type of the Express Connect circuit. Valid values:
	//
	// 	- **100GBase-LR**: 100,000 megabytes of single-mode optical port (10 km)
	//
	// example:
	//
	// 100GBase-LR
	PortType *string `json:"PortType,omitempty" xml:"PortType,omitempty"`
	// The billing cycle. Valid values:
	//
	// 	- **Month**: Billed on a monthly basis
	//
	// 	- **Year**: Billed on an annual basis
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of your Alibaba Cloud resource group.
	//
	// For more information about resource groups, see [Resource groups](https://help.aliyun.com/document_detail/94475.htm).
	//
	// example:
	//
	// rg-aek2l4sq6l7unhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specification; value:
	//
	// 	- **Large**: Large
	//
	// example:
	//
	// Large
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The status of the cache reserve instance. Valid values:
	//
	// 	- **Available**: Normal.
	//
	// 	- **Not Available**: Not available.
	//
	// 	- **Executing**: The task is being executed.
	//
	// 	- **Deleting**: The account is being deleted
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags.
	Tags []*GetVccResponseBodyContentTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1620939556166279
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The ID of the vSwitch. [Virtual Private Cloud VSwitch](https://help.aliyun.com/document_detail/100380.html).
	//
	// You can call the [DescribeVSwitches](https://help.aliyun.com/document_detail/35748.html) operation to query created vSwitches.
	//
	// example:
	//
	// vsw-uf6u8473r84e6n1n19he5
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Information list of border routers
	VbrInfos []*GetVccResponseBodyContentVbrInfos `json:"VbrInfos,omitempty" xml:"VbrInfos,omitempty" type:"Repeated"`
	// The ID of the Lingjun connection instance.
	//
	// example:
	//
	// vcc-cn-cqf2xh40101
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
	// The name of the Lingjun connection instance.
	//
	// example:
	//
	// vcc-heyuan-backup
	VccName *string `json:"VccName,omitempty" xml:"VccName,omitempty"`
	// Virtual Private Cloud IDs; [What is Virtual Private Cloud](https://help.aliyun.com/document_detail/34217.html)
	//
	// You can call the [DescribeVpcs](https://help.aliyun.com/document_detail/35739.html#demo-0) operation to query the specified VPC.
	//
	// example:
	//
	// vpc-j6ctp4n75306phv5tmpsm
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Lingjun network segment information (applicable to the scene where the old version of Lingjun connection is directly bound to Lingjun network segment)
	VpdBaseInfo *GetVccResponseBodyContentVpdBaseInfo `json:"VpdBaseInfo,omitempty" xml:"VpdBaseInfo,omitempty" type:"Struct"`
	// Lingjun CIDR block instance ID
	//
	// example:
	//
	// vpd-d3isyds4
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetVccResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s GetVccResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetVccResponseBodyContent) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *GetVccResponseBodyContent) GetAliyunRouterInfo() []*GetVccResponseBodyContentAliyunRouterInfo {
	return s.AliyunRouterInfo
}

func (s *GetVccResponseBodyContent) GetAttachErStatus() *bool {
	return s.AttachErStatus
}

func (s *GetVccResponseBodyContent) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *GetVccResponseBodyContent) GetBandwidthStr() *string {
	return s.BandwidthStr
}

func (s *GetVccResponseBodyContent) GetBgpAsn() *string {
	return s.BgpAsn
}

func (s *GetVccResponseBodyContent) GetBgpCidr() *string {
	return s.BgpCidr
}

func (s *GetVccResponseBodyContent) GetCenId() *string {
	return s.CenId
}

func (s *GetVccResponseBodyContent) GetCenOwnerId() *string {
	return s.CenOwnerId
}

func (s *GetVccResponseBodyContent) GetCisRouterInfo() []*GetVccResponseBodyContentCisRouterInfo {
	return s.CisRouterInfo
}

func (s *GetVccResponseBodyContent) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetVccResponseBodyContent) GetConnectionType() *string {
	return s.ConnectionType
}

func (s *GetVccResponseBodyContent) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetVccResponseBodyContent) GetCurrentNode() *string {
	return s.CurrentNode
}

func (s *GetVccResponseBodyContent) GetDuration() *string {
	return s.Duration
}

func (s *GetVccResponseBodyContent) GetErInfos() []*GetVccResponseBodyContentErInfos {
	return s.ErInfos
}

func (s *GetVccResponseBodyContent) GetExpirationDate() *string {
	return s.ExpirationDate
}

func (s *GetVccResponseBodyContent) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetVccResponseBodyContent) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *GetVccResponseBodyContent) GetLineOperator() *string {
	return s.LineOperator
}

func (s *GetVccResponseBodyContent) GetMessage() *string {
	return s.Message
}

func (s *GetVccResponseBodyContent) GetPayType() *string {
	return s.PayType
}

func (s *GetVccResponseBodyContent) GetPortType() *string {
	return s.PortType
}

func (s *GetVccResponseBodyContent) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *GetVccResponseBodyContent) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVccResponseBodyContent) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetVccResponseBodyContent) GetSpec() *string {
	return s.Spec
}

func (s *GetVccResponseBodyContent) GetStatus() *string {
	return s.Status
}

func (s *GetVccResponseBodyContent) GetTags() []*GetVccResponseBodyContentTags {
	return s.Tags
}

func (s *GetVccResponseBodyContent) GetTenantId() *string {
	return s.TenantId
}

func (s *GetVccResponseBodyContent) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetVccResponseBodyContent) GetVbrInfos() []*GetVccResponseBodyContentVbrInfos {
	return s.VbrInfos
}

func (s *GetVccResponseBodyContent) GetVccId() *string {
	return s.VccId
}

func (s *GetVccResponseBodyContent) GetVccName() *string {
	return s.VccName
}

func (s *GetVccResponseBodyContent) GetVpcId() *string {
	return s.VpcId
}

func (s *GetVccResponseBodyContent) GetVpdBaseInfo() *GetVccResponseBodyContentVpdBaseInfo {
	return s.VpdBaseInfo
}

func (s *GetVccResponseBodyContent) GetVpdId() *string {
	return s.VpdId
}

func (s *GetVccResponseBodyContent) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetVccResponseBodyContent) SetAccessPointId(v string) *GetVccResponseBodyContent {
	s.AccessPointId = &v
	return s
}

func (s *GetVccResponseBodyContent) SetAliyunRouterInfo(v []*GetVccResponseBodyContentAliyunRouterInfo) *GetVccResponseBodyContent {
	s.AliyunRouterInfo = v
	return s
}

func (s *GetVccResponseBodyContent) SetAttachErStatus(v bool) *GetVccResponseBodyContent {
	s.AttachErStatus = &v
	return s
}

func (s *GetVccResponseBodyContent) SetBandwidth(v int32) *GetVccResponseBodyContent {
	s.Bandwidth = &v
	return s
}

func (s *GetVccResponseBodyContent) SetBandwidthStr(v string) *GetVccResponseBodyContent {
	s.BandwidthStr = &v
	return s
}

func (s *GetVccResponseBodyContent) SetBgpAsn(v string) *GetVccResponseBodyContent {
	s.BgpAsn = &v
	return s
}

func (s *GetVccResponseBodyContent) SetBgpCidr(v string) *GetVccResponseBodyContent {
	s.BgpCidr = &v
	return s
}

func (s *GetVccResponseBodyContent) SetCenId(v string) *GetVccResponseBodyContent {
	s.CenId = &v
	return s
}

func (s *GetVccResponseBodyContent) SetCenOwnerId(v string) *GetVccResponseBodyContent {
	s.CenOwnerId = &v
	return s
}

func (s *GetVccResponseBodyContent) SetCisRouterInfo(v []*GetVccResponseBodyContentCisRouterInfo) *GetVccResponseBodyContent {
	s.CisRouterInfo = v
	return s
}

func (s *GetVccResponseBodyContent) SetCommodityCode(v string) *GetVccResponseBodyContent {
	s.CommodityCode = &v
	return s
}

func (s *GetVccResponseBodyContent) SetConnectionType(v string) *GetVccResponseBodyContent {
	s.ConnectionType = &v
	return s
}

func (s *GetVccResponseBodyContent) SetCreateTime(v string) *GetVccResponseBodyContent {
	s.CreateTime = &v
	return s
}

func (s *GetVccResponseBodyContent) SetCurrentNode(v string) *GetVccResponseBodyContent {
	s.CurrentNode = &v
	return s
}

func (s *GetVccResponseBodyContent) SetDuration(v string) *GetVccResponseBodyContent {
	s.Duration = &v
	return s
}

func (s *GetVccResponseBodyContent) SetErInfos(v []*GetVccResponseBodyContentErInfos) *GetVccResponseBodyContent {
	s.ErInfos = v
	return s
}

func (s *GetVccResponseBodyContent) SetExpirationDate(v string) *GetVccResponseBodyContent {
	s.ExpirationDate = &v
	return s
}

func (s *GetVccResponseBodyContent) SetGmtModified(v string) *GetVccResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *GetVccResponseBodyContent) SetInternetChargeType(v string) *GetVccResponseBodyContent {
	s.InternetChargeType = &v
	return s
}

func (s *GetVccResponseBodyContent) SetLineOperator(v string) *GetVccResponseBodyContent {
	s.LineOperator = &v
	return s
}

func (s *GetVccResponseBodyContent) SetMessage(v string) *GetVccResponseBodyContent {
	s.Message = &v
	return s
}

func (s *GetVccResponseBodyContent) SetPayType(v string) *GetVccResponseBodyContent {
	s.PayType = &v
	return s
}

func (s *GetVccResponseBodyContent) SetPortType(v string) *GetVccResponseBodyContent {
	s.PortType = &v
	return s
}

func (s *GetVccResponseBodyContent) SetPricingCycle(v string) *GetVccResponseBodyContent {
	s.PricingCycle = &v
	return s
}

func (s *GetVccResponseBodyContent) SetRegionId(v string) *GetVccResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetVccResponseBodyContent) SetResourceGroupId(v string) *GetVccResponseBodyContent {
	s.ResourceGroupId = &v
	return s
}

func (s *GetVccResponseBodyContent) SetSpec(v string) *GetVccResponseBodyContent {
	s.Spec = &v
	return s
}

func (s *GetVccResponseBodyContent) SetStatus(v string) *GetVccResponseBodyContent {
	s.Status = &v
	return s
}

func (s *GetVccResponseBodyContent) SetTags(v []*GetVccResponseBodyContentTags) *GetVccResponseBodyContent {
	s.Tags = v
	return s
}

func (s *GetVccResponseBodyContent) SetTenantId(v string) *GetVccResponseBodyContent {
	s.TenantId = &v
	return s
}

func (s *GetVccResponseBodyContent) SetVSwitchId(v string) *GetVccResponseBodyContent {
	s.VSwitchId = &v
	return s
}

func (s *GetVccResponseBodyContent) SetVbrInfos(v []*GetVccResponseBodyContentVbrInfos) *GetVccResponseBodyContent {
	s.VbrInfos = v
	return s
}

func (s *GetVccResponseBodyContent) SetVccId(v string) *GetVccResponseBodyContent {
	s.VccId = &v
	return s
}

func (s *GetVccResponseBodyContent) SetVccName(v string) *GetVccResponseBodyContent {
	s.VccName = &v
	return s
}

func (s *GetVccResponseBodyContent) SetVpcId(v string) *GetVccResponseBodyContent {
	s.VpcId = &v
	return s
}

func (s *GetVccResponseBodyContent) SetVpdBaseInfo(v *GetVccResponseBodyContentVpdBaseInfo) *GetVccResponseBodyContent {
	s.VpdBaseInfo = v
	return s
}

func (s *GetVccResponseBodyContent) SetVpdId(v string) *GetVccResponseBodyContent {
	s.VpdId = &v
	return s
}

func (s *GetVccResponseBodyContent) SetZoneId(v string) *GetVccResponseBodyContent {
	s.ZoneId = &v
	return s
}

func (s *GetVccResponseBodyContent) Validate() error {
	if s.AliyunRouterInfo != nil {
		for _, item := range s.AliyunRouterInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CisRouterInfo != nil {
		for _, item := range s.CisRouterInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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
	if s.VbrInfos != nil {
		for _, item := range s.VbrInfos {
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

type GetVccResponseBodyContentAliyunRouterInfo struct {
	// IPv4 address of Alibaba Cloud-side interconnection
	//
	// example:
	//
	// 169.254.248.30
	LocalGatewayIp *string `json:"LocalGatewayIp,omitempty" xml:"LocalGatewayIp,omitempty"`
	// Masking
	//
	// example:
	//
	// 255.255.255.248
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// Express Connect circuit ID
	//
	// example:
	//
	// pc-0jlof4bphlsnxbdztkvad
	PcId *string `json:"PcId,omitempty" xml:"PcId,omitempty"`
	// Lingjun Side Interconnection IPv4 Address
	//
	// example:
	//
	// 169.254.248.28
	PeerGatewayIp *string `json:"PeerGatewayIp,omitempty" xml:"PeerGatewayIp,omitempty"`
	// The ID of the VBR.
	//
	// example:
	//
	// vbr-2ze4i85p6vb9nwcan5xt0
	VbrId *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
	// VLAN ID of the VBR
	//
	// example:
	//
	// 1042
	VlanId *string `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
}

func (s GetVccResponseBodyContentAliyunRouterInfo) String() string {
	return dara.Prettify(s)
}

func (s GetVccResponseBodyContentAliyunRouterInfo) GoString() string {
	return s.String()
}

func (s *GetVccResponseBodyContentAliyunRouterInfo) GetLocalGatewayIp() *string {
	return s.LocalGatewayIp
}

func (s *GetVccResponseBodyContentAliyunRouterInfo) GetMask() *string {
	return s.Mask
}

func (s *GetVccResponseBodyContentAliyunRouterInfo) GetPcId() *string {
	return s.PcId
}

func (s *GetVccResponseBodyContentAliyunRouterInfo) GetPeerGatewayIp() *string {
	return s.PeerGatewayIp
}

func (s *GetVccResponseBodyContentAliyunRouterInfo) GetVbrId() *string {
	return s.VbrId
}

func (s *GetVccResponseBodyContentAliyunRouterInfo) GetVlanId() *string {
	return s.VlanId
}

func (s *GetVccResponseBodyContentAliyunRouterInfo) SetLocalGatewayIp(v string) *GetVccResponseBodyContentAliyunRouterInfo {
	s.LocalGatewayIp = &v
	return s
}

func (s *GetVccResponseBodyContentAliyunRouterInfo) SetMask(v string) *GetVccResponseBodyContentAliyunRouterInfo {
	s.Mask = &v
	return s
}

func (s *GetVccResponseBodyContentAliyunRouterInfo) SetPcId(v string) *GetVccResponseBodyContentAliyunRouterInfo {
	s.PcId = &v
	return s
}

func (s *GetVccResponseBodyContentAliyunRouterInfo) SetPeerGatewayIp(v string) *GetVccResponseBodyContentAliyunRouterInfo {
	s.PeerGatewayIp = &v
	return s
}

func (s *GetVccResponseBodyContentAliyunRouterInfo) SetVbrId(v string) *GetVccResponseBodyContentAliyunRouterInfo {
	s.VbrId = &v
	return s
}

func (s *GetVccResponseBodyContentAliyunRouterInfo) SetVlanId(v string) *GetVccResponseBodyContentAliyunRouterInfo {
	s.VlanId = &v
	return s
}

func (s *GetVccResponseBodyContentAliyunRouterInfo) Validate() error {
	return dara.Validate(s)
}

type GetVccResponseBodyContentCisRouterInfo struct {
	// Leased Line Information List
	CcInfos []*GetVccResponseBodyContentCisRouterInfoCcInfos `json:"CcInfos,omitempty" xml:"CcInfos,omitempty" type:"Repeated"`
	// The ID of the on-cloud router instance.
	//
	// example:
	//
	// ccr-1ms84am0
	CcrId *string `json:"CcrId,omitempty" xml:"CcrId,omitempty"`
}

func (s GetVccResponseBodyContentCisRouterInfo) String() string {
	return dara.Prettify(s)
}

func (s GetVccResponseBodyContentCisRouterInfo) GoString() string {
	return s.String()
}

func (s *GetVccResponseBodyContentCisRouterInfo) GetCcInfos() []*GetVccResponseBodyContentCisRouterInfoCcInfos {
	return s.CcInfos
}

func (s *GetVccResponseBodyContentCisRouterInfo) GetCcrId() *string {
	return s.CcrId
}

func (s *GetVccResponseBodyContentCisRouterInfo) SetCcInfos(v []*GetVccResponseBodyContentCisRouterInfoCcInfos) *GetVccResponseBodyContentCisRouterInfo {
	s.CcInfos = v
	return s
}

func (s *GetVccResponseBodyContentCisRouterInfo) SetCcrId(v string) *GetVccResponseBodyContentCisRouterInfo {
	s.CcrId = &v
	return s
}

func (s *GetVccResponseBodyContentCisRouterInfo) Validate() error {
	if s.CcInfos != nil {
		for _, item := range s.CcInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVccResponseBodyContentCisRouterInfoCcInfos struct {
	// Leased Line ID
	//
	// example:
	//
	// cc-73aeex5o
	CcId *string `json:"CcId,omitempty" xml:"CcId,omitempty"`
	// Lingjun Side Interconnection IPv4 Address
	//
	// example:
	//
	// 169.254.248.26
	LocalGatewayIp *string `json:"LocalGatewayIp,omitempty" xml:"LocalGatewayIp,omitempty"`
	// Lingjun Side Interconnection IPv4 Address
	//
	// example:
	//
	// 169.254.248.30
	RemoteGatewayIp *string `json:"RemoteGatewayIp,omitempty" xml:"RemoteGatewayIp,omitempty"`
	// The state of the rule.
	//
	// example:
	//
	// Established
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Subnet mask
	//
	// example:
	//
	// 255.255.255.248
	SubnetMask *string `json:"SubnetMask,omitempty" xml:"SubnetMask,omitempty"`
	// Vlan ID of the leased line
	//
	// example:
	//
	// Ethernet1042
	VlanId *string `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
}

func (s GetVccResponseBodyContentCisRouterInfoCcInfos) String() string {
	return dara.Prettify(s)
}

func (s GetVccResponseBodyContentCisRouterInfoCcInfos) GoString() string {
	return s.String()
}

func (s *GetVccResponseBodyContentCisRouterInfoCcInfos) GetCcId() *string {
	return s.CcId
}

func (s *GetVccResponseBodyContentCisRouterInfoCcInfos) GetLocalGatewayIp() *string {
	return s.LocalGatewayIp
}

func (s *GetVccResponseBodyContentCisRouterInfoCcInfos) GetRemoteGatewayIp() *string {
	return s.RemoteGatewayIp
}

func (s *GetVccResponseBodyContentCisRouterInfoCcInfos) GetStatus() *string {
	return s.Status
}

func (s *GetVccResponseBodyContentCisRouterInfoCcInfos) GetSubnetMask() *string {
	return s.SubnetMask
}

func (s *GetVccResponseBodyContentCisRouterInfoCcInfos) GetVlanId() *string {
	return s.VlanId
}

func (s *GetVccResponseBodyContentCisRouterInfoCcInfos) SetCcId(v string) *GetVccResponseBodyContentCisRouterInfoCcInfos {
	s.CcId = &v
	return s
}

func (s *GetVccResponseBodyContentCisRouterInfoCcInfos) SetLocalGatewayIp(v string) *GetVccResponseBodyContentCisRouterInfoCcInfos {
	s.LocalGatewayIp = &v
	return s
}

func (s *GetVccResponseBodyContentCisRouterInfoCcInfos) SetRemoteGatewayIp(v string) *GetVccResponseBodyContentCisRouterInfoCcInfos {
	s.RemoteGatewayIp = &v
	return s
}

func (s *GetVccResponseBodyContentCisRouterInfoCcInfos) SetStatus(v string) *GetVccResponseBodyContentCisRouterInfoCcInfos {
	s.Status = &v
	return s
}

func (s *GetVccResponseBodyContentCisRouterInfoCcInfos) SetSubnetMask(v string) *GetVccResponseBodyContentCisRouterInfoCcInfos {
	s.SubnetMask = &v
	return s
}

func (s *GetVccResponseBodyContentCisRouterInfoCcInfos) SetVlanId(v string) *GetVccResponseBodyContentCisRouterInfoCcInfos {
	s.VlanId = &v
	return s
}

func (s *GetVccResponseBodyContentCisRouterInfoCcInfos) Validate() error {
	return dara.Validate(s)
}

type GetVccResponseBodyContentErInfos struct {
	// Connections
	//
	// example:
	//
	// 2
	Connections *int64 `json:"Connections,omitempty" xml:"Connections,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1678379917000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Description
	//
	// example:
	//
	// this is test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Lingjun HUB ID
	//
	// example:
	//
	// er-p68b0jwn
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Lingjun HUB Instance Name
	//
	// example:
	//
	// er-1
	ErName *string `json:"ErName,omitempty" xml:"ErName,omitempty"`
	// The time when the agent was last modified.
	//
	// example:
	//
	// 1678379917000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Primary Zone
	//
	// example:
	//
	// cn-wulanchabu-b
	MasterZoneId *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// test message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Lingjun HUB Region Information
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Number of routing policy
	//
	// example:
	//
	// 2
	RouteMaps *int64 `json:"RouteMaps,omitempty" xml:"RouteMaps,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1620939556166277
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetVccResponseBodyContentErInfos) String() string {
	return dara.Prettify(s)
}

func (s GetVccResponseBodyContentErInfos) GoString() string {
	return s.String()
}

func (s *GetVccResponseBodyContentErInfos) GetConnections() *int64 {
	return s.Connections
}

func (s *GetVccResponseBodyContentErInfos) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetVccResponseBodyContentErInfos) GetDescription() *string {
	return s.Description
}

func (s *GetVccResponseBodyContentErInfos) GetErId() *string {
	return s.ErId
}

func (s *GetVccResponseBodyContentErInfos) GetErName() *string {
	return s.ErName
}

func (s *GetVccResponseBodyContentErInfos) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetVccResponseBodyContentErInfos) GetMasterZoneId() *string {
	return s.MasterZoneId
}

func (s *GetVccResponseBodyContentErInfos) GetMessage() *string {
	return s.Message
}

func (s *GetVccResponseBodyContentErInfos) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVccResponseBodyContentErInfos) GetRouteMaps() *int64 {
	return s.RouteMaps
}

func (s *GetVccResponseBodyContentErInfos) GetStatus() *string {
	return s.Status
}

func (s *GetVccResponseBodyContentErInfos) GetTenantId() *string {
	return s.TenantId
}

func (s *GetVccResponseBodyContentErInfos) SetConnections(v int64) *GetVccResponseBodyContentErInfos {
	s.Connections = &v
	return s
}

func (s *GetVccResponseBodyContentErInfos) SetCreateTime(v string) *GetVccResponseBodyContentErInfos {
	s.CreateTime = &v
	return s
}

func (s *GetVccResponseBodyContentErInfos) SetDescription(v string) *GetVccResponseBodyContentErInfos {
	s.Description = &v
	return s
}

func (s *GetVccResponseBodyContentErInfos) SetErId(v string) *GetVccResponseBodyContentErInfos {
	s.ErId = &v
	return s
}

func (s *GetVccResponseBodyContentErInfos) SetErName(v string) *GetVccResponseBodyContentErInfos {
	s.ErName = &v
	return s
}

func (s *GetVccResponseBodyContentErInfos) SetGmtModified(v string) *GetVccResponseBodyContentErInfos {
	s.GmtModified = &v
	return s
}

func (s *GetVccResponseBodyContentErInfos) SetMasterZoneId(v string) *GetVccResponseBodyContentErInfos {
	s.MasterZoneId = &v
	return s
}

func (s *GetVccResponseBodyContentErInfos) SetMessage(v string) *GetVccResponseBodyContentErInfos {
	s.Message = &v
	return s
}

func (s *GetVccResponseBodyContentErInfos) SetRegionId(v string) *GetVccResponseBodyContentErInfos {
	s.RegionId = &v
	return s
}

func (s *GetVccResponseBodyContentErInfos) SetRouteMaps(v int64) *GetVccResponseBodyContentErInfos {
	s.RouteMaps = &v
	return s
}

func (s *GetVccResponseBodyContentErInfos) SetStatus(v string) *GetVccResponseBodyContentErInfos {
	s.Status = &v
	return s
}

func (s *GetVccResponseBodyContentErInfos) SetTenantId(v string) *GetVccResponseBodyContentErInfos {
	s.TenantId = &v
	return s
}

func (s *GetVccResponseBodyContentErInfos) Validate() error {
	return dara.Validate(s)
}

type GetVccResponseBodyContentTags struct {
	// The tag key.
	//
	// You cannot specify an empty string as a tag key. It can be up to 64 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// tag-vcc
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag that is added to the resource.
	//
	// The tag value can be empty or a string of up to 128 characters. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// Each key-value pair must be unique. You can specify values for at most 20 tag keys in each call.
	//
	// example:
	//
	// vcc-group-1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetVccResponseBodyContentTags) String() string {
	return dara.Prettify(s)
}

func (s GetVccResponseBodyContentTags) GoString() string {
	return s.String()
}

func (s *GetVccResponseBodyContentTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetVccResponseBodyContentTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetVccResponseBodyContentTags) SetTagKey(v string) *GetVccResponseBodyContentTags {
	s.TagKey = &v
	return s
}

func (s *GetVccResponseBodyContentTags) SetTagValue(v string) *GetVccResponseBodyContentTags {
	s.TagValue = &v
	return s
}

func (s *GetVccResponseBodyContentTags) Validate() error {
	return dara.Validate(s)
}

type GetVccResponseBodyContentVbrInfos struct {
	// CEN ID
	//
	// example:
	//
	// cen-cx0qua8q6cm4z9****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1683250981000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the agent was last modified.
	//
	// example:
	//
	// 1673578603000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The status of the VBR. Valid values:
	//
	// 	- unconfirmed
	//
	// 	- active: The VPN gateway is in a normal state.
	//
	// 	- terminating: The connection is being terminated.
	//
	// 	- terminated: The connection is terminated.
	//
	// 	- recovering: The task is being recovered.
	//
	// 	- deleting: The GDN is being deleted.
	//
	// 	- Available: The service is available.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// BGP neighbor information list
	VbrBgpPeers []*GetVccResponseBodyContentVbrInfosVbrBgpPeers `json:"VbrBgpPeers,omitempty" xml:"VbrBgpPeers,omitempty" type:"Repeated"`
	// The ID of the border router.
	//
	// example:
	//
	// vbr-wz96agu9h3d50z****
	VbrId *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
}

func (s GetVccResponseBodyContentVbrInfos) String() string {
	return dara.Prettify(s)
}

func (s GetVccResponseBodyContentVbrInfos) GoString() string {
	return s.String()
}

func (s *GetVccResponseBodyContentVbrInfos) GetCenId() *string {
	return s.CenId
}

func (s *GetVccResponseBodyContentVbrInfos) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetVccResponseBodyContentVbrInfos) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetVccResponseBodyContentVbrInfos) GetStatus() *string {
	return s.Status
}

func (s *GetVccResponseBodyContentVbrInfos) GetVbrBgpPeers() []*GetVccResponseBodyContentVbrInfosVbrBgpPeers {
	return s.VbrBgpPeers
}

func (s *GetVccResponseBodyContentVbrInfos) GetVbrId() *string {
	return s.VbrId
}

func (s *GetVccResponseBodyContentVbrInfos) SetCenId(v string) *GetVccResponseBodyContentVbrInfos {
	s.CenId = &v
	return s
}

func (s *GetVccResponseBodyContentVbrInfos) SetGmtCreate(v string) *GetVccResponseBodyContentVbrInfos {
	s.GmtCreate = &v
	return s
}

func (s *GetVccResponseBodyContentVbrInfos) SetGmtModified(v string) *GetVccResponseBodyContentVbrInfos {
	s.GmtModified = &v
	return s
}

func (s *GetVccResponseBodyContentVbrInfos) SetStatus(v string) *GetVccResponseBodyContentVbrInfos {
	s.Status = &v
	return s
}

func (s *GetVccResponseBodyContentVbrInfos) SetVbrBgpPeers(v []*GetVccResponseBodyContentVbrInfosVbrBgpPeers) *GetVccResponseBodyContentVbrInfos {
	s.VbrBgpPeers = v
	return s
}

func (s *GetVccResponseBodyContentVbrInfos) SetVbrId(v string) *GetVccResponseBodyContentVbrInfos {
	s.VbrId = &v
	return s
}

func (s *GetVccResponseBodyContentVbrInfos) Validate() error {
	if s.VbrBgpPeers != nil {
		for _, item := range s.VbrBgpPeers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVccResponseBodyContentVbrInfosVbrBgpPeers struct {
	// BGP Group ID
	//
	// example:
	//
	// bgpg-2ze2sit2vakrkapvy****
	BgpGroupId *string `json:"BgpGroupId,omitempty" xml:"BgpGroupId,omitempty"`
	// BGP peer ID
	//
	// example:
	//
	// bgp-uf6heugif9enu48rj****
	BgpPeerId *string `json:"BgpPeerId,omitempty" xml:"BgpPeerId,omitempty"`
	// Peer AS No.
	//
	// example:
	//
	// 98765****
	PeerAsn *string `json:"PeerAsn,omitempty" xml:"PeerAsn,omitempty"`
	// BGP peer IP address
	//
	// example:
	//
	// 169.254.****
	PeerIpAddress *string `json:"PeerIpAddress,omitempty" xml:"PeerIpAddress,omitempty"`
	// The status of the BGP peer. Valid values:
	//
	// 	- Pending: pending
	//
	// 	- Available: The route is available.
	//
	// 	- Modifying: being modified
	//
	// 	- Deleting: The IPv4 gateway is being deleted.
	//
	// 	- Deleted
	//
	// 	- Not Available
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetVccResponseBodyContentVbrInfosVbrBgpPeers) String() string {
	return dara.Prettify(s)
}

func (s GetVccResponseBodyContentVbrInfosVbrBgpPeers) GoString() string {
	return s.String()
}

func (s *GetVccResponseBodyContentVbrInfosVbrBgpPeers) GetBgpGroupId() *string {
	return s.BgpGroupId
}

func (s *GetVccResponseBodyContentVbrInfosVbrBgpPeers) GetBgpPeerId() *string {
	return s.BgpPeerId
}

func (s *GetVccResponseBodyContentVbrInfosVbrBgpPeers) GetPeerAsn() *string {
	return s.PeerAsn
}

func (s *GetVccResponseBodyContentVbrInfosVbrBgpPeers) GetPeerIpAddress() *string {
	return s.PeerIpAddress
}

func (s *GetVccResponseBodyContentVbrInfosVbrBgpPeers) GetStatus() *string {
	return s.Status
}

func (s *GetVccResponseBodyContentVbrInfosVbrBgpPeers) SetBgpGroupId(v string) *GetVccResponseBodyContentVbrInfosVbrBgpPeers {
	s.BgpGroupId = &v
	return s
}

func (s *GetVccResponseBodyContentVbrInfosVbrBgpPeers) SetBgpPeerId(v string) *GetVccResponseBodyContentVbrInfosVbrBgpPeers {
	s.BgpPeerId = &v
	return s
}

func (s *GetVccResponseBodyContentVbrInfosVbrBgpPeers) SetPeerAsn(v string) *GetVccResponseBodyContentVbrInfosVbrBgpPeers {
	s.PeerAsn = &v
	return s
}

func (s *GetVccResponseBodyContentVbrInfosVbrBgpPeers) SetPeerIpAddress(v string) *GetVccResponseBodyContentVbrInfosVbrBgpPeers {
	s.PeerIpAddress = &v
	return s
}

func (s *GetVccResponseBodyContentVbrInfosVbrBgpPeers) SetStatus(v string) *GetVccResponseBodyContentVbrInfosVbrBgpPeers {
	s.Status = &v
	return s
}

func (s *GetVccResponseBodyContentVbrInfosVbrBgpPeers) Validate() error {
	return dara.Validate(s)
}

type GetVccResponseBodyContentVpdBaseInfo struct {
	// Network address segment
	//
	// example:
	//
	// 172.18.0.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1678379917000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Lingjun CIDR block instance ID
	//
	// example:
	//
	// vpd-ppdunxzc
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// Lingjun CIDR block instance name
	//
	// example:
	//
	// yzp-rg-test3
	VpdName *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
}

func (s GetVccResponseBodyContentVpdBaseInfo) String() string {
	return dara.Prettify(s)
}

func (s GetVccResponseBodyContentVpdBaseInfo) GoString() string {
	return s.String()
}

func (s *GetVccResponseBodyContentVpdBaseInfo) GetCidr() *string {
	return s.Cidr
}

func (s *GetVccResponseBodyContentVpdBaseInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetVccResponseBodyContentVpdBaseInfo) GetVpdId() *string {
	return s.VpdId
}

func (s *GetVccResponseBodyContentVpdBaseInfo) GetVpdName() *string {
	return s.VpdName
}

func (s *GetVccResponseBodyContentVpdBaseInfo) SetCidr(v string) *GetVccResponseBodyContentVpdBaseInfo {
	s.Cidr = &v
	return s
}

func (s *GetVccResponseBodyContentVpdBaseInfo) SetCreateTime(v string) *GetVccResponseBodyContentVpdBaseInfo {
	s.CreateTime = &v
	return s
}

func (s *GetVccResponseBodyContentVpdBaseInfo) SetVpdId(v string) *GetVccResponseBodyContentVpdBaseInfo {
	s.VpdId = &v
	return s
}

func (s *GetVccResponseBodyContentVpdBaseInfo) SetVpdName(v string) *GetVccResponseBodyContentVpdBaseInfo {
	s.VpdName = &v
	return s
}

func (s *GetVccResponseBodyContentVpdBaseInfo) Validate() error {
	return dara.Validate(s)
}
