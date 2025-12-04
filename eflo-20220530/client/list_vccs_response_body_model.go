// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVccsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListVccsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *ListVccsResponseBody
	GetCode() *int32
	SetContent(v *ListVccsResponseBodyContent) *ListVccsResponseBody
	GetContent() *ListVccsResponseBodyContent
	SetMessage(v string) *ListVccsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListVccsResponseBody
	GetRequestId() *string
}

type ListVccsResponseBody struct {
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
	Content *ListVccsResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
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
	// 28451248-7038-5184-B5D3-80F104654BE8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVccsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVccsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVccsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListVccsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListVccsResponseBody) GetContent() *ListVccsResponseBodyContent {
	return s.Content
}

func (s *ListVccsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVccsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVccsResponseBody) SetAccessDeniedDetail(v string) *ListVccsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListVccsResponseBody) SetCode(v int32) *ListVccsResponseBody {
	s.Code = &v
	return s
}

func (s *ListVccsResponseBody) SetContent(v *ListVccsResponseBodyContent) *ListVccsResponseBody {
	s.Content = v
	return s
}

func (s *ListVccsResponseBody) SetMessage(v string) *ListVccsResponseBody {
	s.Message = &v
	return s
}

func (s *ListVccsResponseBody) SetRequestId(v string) *ListVccsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVccsResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVccsResponseBodyContent struct {
	// Lingjun Connection Information List
	Data []*ListVccsResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 0
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListVccsResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListVccsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListVccsResponseBodyContent) GetData() []*ListVccsResponseBodyContentData {
	return s.Data
}

func (s *ListVccsResponseBodyContent) GetTotal() *int64 {
	return s.Total
}

func (s *ListVccsResponseBodyContent) SetData(v []*ListVccsResponseBodyContentData) *ListVccsResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListVccsResponseBodyContent) SetTotal(v int64) *ListVccsResponseBodyContent {
	s.Total = &v
	return s
}

func (s *ListVccsResponseBodyContent) Validate() error {
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

type ListVccsResponseBodyContentData struct {
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
	// The bandwidth of the port.
	//
	// example:
	//
	// 1000
	BandwidthStr *string `json:"BandwidthStr,omitempty" xml:"BandwidthStr,omitempty"`
	// bgp as number
	//
	// example:
	//
	// bgpAsn
	BgpAsn *string `json:"BgpAsn,omitempty" xml:"BgpAsn,omitempty"`
	// bgp network segment
	//
	// example:
	//
	// 172.16.128.0/24
	BgpCidr *string `json:"BgpCidr,omitempty" xml:"BgpCidr,omitempty"`
	// The ID of the CEN instance; [What is the CEN?](https://help.aliyun.com/document_detail/181681.html)
	//
	// You can call the [DescribeCens](https://help.aliyun.com/document_detail/468215.htm) to query the information of CEN instances under the current Alibaba Cloud account.
	//
	// example:
	//
	// cen-w15qot0pfvs83pkckj
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// Account to which cen belongs
	//
	// example:
	//
	// 1238685214107736
	CenOwnerId *string `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
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
	// 1678273219000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Current process node
	//
	// example:
	//
	// test-xxxx-node-x
	CurrentNode *string `json:"CurrentNode,omitempty" xml:"CurrentNode,omitempty"`
	// List of bound Lingjun HUB information
	ErInfos []*ListVccsResponseBodyContentDataErInfos `json:"ErInfos,omitempty" xml:"ErInfos,omitempty" type:"Repeated"`
	// The time when the application expired.
	//
	// example:
	//
	// 1678273219000
	ExpirationDate *string `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	// The time when the cluster was updated.
	//
	// example:
	//
	// 1678273219000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
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
	// some message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The port type of the Express Connect circuit. Valid values:
	//
	// 	- **100GBase-LR**: 100,000 megabytes of single-mode optical port (10 km)
	//
	// example:
	//
	// 100GBase-LR
	PortType *string `json:"PortType,omitempty" xml:"PortType,omitempty"`
	// Process progress; value returns 0 to 1; not started is null
	//
	// example:
	//
	// 1
	Rate *float64 `json:"Rate,omitempty" xml:"Rate,omitempty"`
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
	// The compute specification.
	//
	// example:
	//
	// Large
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The state of the rule.
	//
	// example:
	//
	// Init
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags.
	Tags []*ListVccsResponseBodyContentDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The job ID.
	//
	// example:
	//
	// task-cd544092-ed0a-49e9-83eb-e8c94770dccf
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1620939556166279
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The ID of the Lingjun connection instance.
	//
	// example:
	//
	// vcc-cn-zvp2w222001
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
	// vpc-f8ziirfl9k25h2qn7y4f8
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Lingjun network segment information (applicable to the scene where the old version of Lingjun connection is directly bound to Lingjun network segment)
	VpdBaseInfo *ListVccsResponseBodyContentDataVpdBaseInfo `json:"VpdBaseInfo,omitempty" xml:"VpdBaseInfo,omitempty" type:"Struct"`
	// Lingjun CIDR block instance ID
	//
	// example:
	//
	// vpd-eoiy88ju
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListVccsResponseBodyContentData) String() string {
	return dara.Prettify(s)
}

func (s ListVccsResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListVccsResponseBodyContentData) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *ListVccsResponseBodyContentData) GetBandwidthStr() *string {
	return s.BandwidthStr
}

func (s *ListVccsResponseBodyContentData) GetBgpAsn() *string {
	return s.BgpAsn
}

func (s *ListVccsResponseBodyContentData) GetBgpCidr() *string {
	return s.BgpCidr
}

func (s *ListVccsResponseBodyContentData) GetCenId() *string {
	return s.CenId
}

func (s *ListVccsResponseBodyContentData) GetCenOwnerId() *string {
	return s.CenOwnerId
}

func (s *ListVccsResponseBodyContentData) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListVccsResponseBodyContentData) GetConnectionType() *string {
	return s.ConnectionType
}

func (s *ListVccsResponseBodyContentData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListVccsResponseBodyContentData) GetCurrentNode() *string {
	return s.CurrentNode
}

func (s *ListVccsResponseBodyContentData) GetErInfos() []*ListVccsResponseBodyContentDataErInfos {
	return s.ErInfos
}

func (s *ListVccsResponseBodyContentData) GetExpirationDate() *string {
	return s.ExpirationDate
}

func (s *ListVccsResponseBodyContentData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListVccsResponseBodyContentData) GetLineOperator() *string {
	return s.LineOperator
}

func (s *ListVccsResponseBodyContentData) GetMessage() *string {
	return s.Message
}

func (s *ListVccsResponseBodyContentData) GetPortType() *string {
	return s.PortType
}

func (s *ListVccsResponseBodyContentData) GetRate() *float64 {
	return s.Rate
}

func (s *ListVccsResponseBodyContentData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVccsResponseBodyContentData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVccsResponseBodyContentData) GetSpec() *string {
	return s.Spec
}

func (s *ListVccsResponseBodyContentData) GetStatus() *string {
	return s.Status
}

func (s *ListVccsResponseBodyContentData) GetTags() []*ListVccsResponseBodyContentDataTags {
	return s.Tags
}

func (s *ListVccsResponseBodyContentData) GetTaskId() *string {
	return s.TaskId
}

func (s *ListVccsResponseBodyContentData) GetTenantId() *string {
	return s.TenantId
}

func (s *ListVccsResponseBodyContentData) GetVccId() *string {
	return s.VccId
}

func (s *ListVccsResponseBodyContentData) GetVccName() *string {
	return s.VccName
}

func (s *ListVccsResponseBodyContentData) GetVpcId() *string {
	return s.VpcId
}

func (s *ListVccsResponseBodyContentData) GetVpdBaseInfo() *ListVccsResponseBodyContentDataVpdBaseInfo {
	return s.VpdBaseInfo
}

func (s *ListVccsResponseBodyContentData) GetVpdId() *string {
	return s.VpdId
}

func (s *ListVccsResponseBodyContentData) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListVccsResponseBodyContentData) SetAccessPointId(v string) *ListVccsResponseBodyContentData {
	s.AccessPointId = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetBandwidthStr(v string) *ListVccsResponseBodyContentData {
	s.BandwidthStr = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetBgpAsn(v string) *ListVccsResponseBodyContentData {
	s.BgpAsn = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetBgpCidr(v string) *ListVccsResponseBodyContentData {
	s.BgpCidr = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetCenId(v string) *ListVccsResponseBodyContentData {
	s.CenId = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetCenOwnerId(v string) *ListVccsResponseBodyContentData {
	s.CenOwnerId = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetCommodityCode(v string) *ListVccsResponseBodyContentData {
	s.CommodityCode = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetConnectionType(v string) *ListVccsResponseBodyContentData {
	s.ConnectionType = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetCreateTime(v string) *ListVccsResponseBodyContentData {
	s.CreateTime = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetCurrentNode(v string) *ListVccsResponseBodyContentData {
	s.CurrentNode = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetErInfos(v []*ListVccsResponseBodyContentDataErInfos) *ListVccsResponseBodyContentData {
	s.ErInfos = v
	return s
}

func (s *ListVccsResponseBodyContentData) SetExpirationDate(v string) *ListVccsResponseBodyContentData {
	s.ExpirationDate = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetGmtModified(v string) *ListVccsResponseBodyContentData {
	s.GmtModified = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetLineOperator(v string) *ListVccsResponseBodyContentData {
	s.LineOperator = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetMessage(v string) *ListVccsResponseBodyContentData {
	s.Message = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetPortType(v string) *ListVccsResponseBodyContentData {
	s.PortType = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetRate(v float64) *ListVccsResponseBodyContentData {
	s.Rate = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetRegionId(v string) *ListVccsResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetResourceGroupId(v string) *ListVccsResponseBodyContentData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetSpec(v string) *ListVccsResponseBodyContentData {
	s.Spec = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetStatus(v string) *ListVccsResponseBodyContentData {
	s.Status = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetTags(v []*ListVccsResponseBodyContentDataTags) *ListVccsResponseBodyContentData {
	s.Tags = v
	return s
}

func (s *ListVccsResponseBodyContentData) SetTaskId(v string) *ListVccsResponseBodyContentData {
	s.TaskId = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetTenantId(v string) *ListVccsResponseBodyContentData {
	s.TenantId = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetVccId(v string) *ListVccsResponseBodyContentData {
	s.VccId = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetVccName(v string) *ListVccsResponseBodyContentData {
	s.VccName = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetVpcId(v string) *ListVccsResponseBodyContentData {
	s.VpcId = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetVpdBaseInfo(v *ListVccsResponseBodyContentDataVpdBaseInfo) *ListVccsResponseBodyContentData {
	s.VpdBaseInfo = v
	return s
}

func (s *ListVccsResponseBodyContentData) SetVpdId(v string) *ListVccsResponseBodyContentData {
	s.VpdId = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetZoneId(v string) *ListVccsResponseBodyContentData {
	s.ZoneId = &v
	return s
}

func (s *ListVccsResponseBodyContentData) Validate() error {
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
	if s.VpdBaseInfo != nil {
		if err := s.VpdBaseInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVccsResponseBodyContentDataErInfos struct {
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
	// 1678273219000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Description
	//
	// example:
	//
	// test_api_coverage
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Elastic Router ID
	//
	// example:
	//
	// er-a7rqv1rq
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// ER instance name
	//
	// example:
	//
	// er-1
	ErName *string `json:"ErName,omitempty" xml:"ErName,omitempty"`
	// The time when the agent was last modified.
	//
	// example:
	//
	// 1678273219000
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
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ER region information
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

func (s ListVccsResponseBodyContentDataErInfos) String() string {
	return dara.Prettify(s)
}

func (s ListVccsResponseBodyContentDataErInfos) GoString() string {
	return s.String()
}

func (s *ListVccsResponseBodyContentDataErInfos) GetConnections() *int64 {
	return s.Connections
}

func (s *ListVccsResponseBodyContentDataErInfos) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListVccsResponseBodyContentDataErInfos) GetDescription() *string {
	return s.Description
}

func (s *ListVccsResponseBodyContentDataErInfos) GetErId() *string {
	return s.ErId
}

func (s *ListVccsResponseBodyContentDataErInfos) GetErName() *string {
	return s.ErName
}

func (s *ListVccsResponseBodyContentDataErInfos) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListVccsResponseBodyContentDataErInfos) GetMasterZoneId() *string {
	return s.MasterZoneId
}

func (s *ListVccsResponseBodyContentDataErInfos) GetMessage() *string {
	return s.Message
}

func (s *ListVccsResponseBodyContentDataErInfos) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVccsResponseBodyContentDataErInfos) GetRouteMaps() *int64 {
	return s.RouteMaps
}

func (s *ListVccsResponseBodyContentDataErInfos) GetStatus() *string {
	return s.Status
}

func (s *ListVccsResponseBodyContentDataErInfos) GetTenantId() *string {
	return s.TenantId
}

func (s *ListVccsResponseBodyContentDataErInfos) SetConnections(v int64) *ListVccsResponseBodyContentDataErInfos {
	s.Connections = &v
	return s
}

func (s *ListVccsResponseBodyContentDataErInfos) SetCreateTime(v string) *ListVccsResponseBodyContentDataErInfos {
	s.CreateTime = &v
	return s
}

func (s *ListVccsResponseBodyContentDataErInfos) SetDescription(v string) *ListVccsResponseBodyContentDataErInfos {
	s.Description = &v
	return s
}

func (s *ListVccsResponseBodyContentDataErInfos) SetErId(v string) *ListVccsResponseBodyContentDataErInfos {
	s.ErId = &v
	return s
}

func (s *ListVccsResponseBodyContentDataErInfos) SetErName(v string) *ListVccsResponseBodyContentDataErInfos {
	s.ErName = &v
	return s
}

func (s *ListVccsResponseBodyContentDataErInfos) SetGmtModified(v string) *ListVccsResponseBodyContentDataErInfos {
	s.GmtModified = &v
	return s
}

func (s *ListVccsResponseBodyContentDataErInfos) SetMasterZoneId(v string) *ListVccsResponseBodyContentDataErInfos {
	s.MasterZoneId = &v
	return s
}

func (s *ListVccsResponseBodyContentDataErInfos) SetMessage(v string) *ListVccsResponseBodyContentDataErInfos {
	s.Message = &v
	return s
}

func (s *ListVccsResponseBodyContentDataErInfos) SetRegionId(v string) *ListVccsResponseBodyContentDataErInfos {
	s.RegionId = &v
	return s
}

func (s *ListVccsResponseBodyContentDataErInfos) SetRouteMaps(v int64) *ListVccsResponseBodyContentDataErInfos {
	s.RouteMaps = &v
	return s
}

func (s *ListVccsResponseBodyContentDataErInfos) SetStatus(v string) *ListVccsResponseBodyContentDataErInfos {
	s.Status = &v
	return s
}

func (s *ListVccsResponseBodyContentDataErInfos) SetTenantId(v string) *ListVccsResponseBodyContentDataErInfos {
	s.TenantId = &v
	return s
}

func (s *ListVccsResponseBodyContentDataErInfos) Validate() error {
	return dara.Validate(s)
}

type ListVccsResponseBodyContentDataTags struct {
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

func (s ListVccsResponseBodyContentDataTags) String() string {
	return dara.Prettify(s)
}

func (s ListVccsResponseBodyContentDataTags) GoString() string {
	return s.String()
}

func (s *ListVccsResponseBodyContentDataTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListVccsResponseBodyContentDataTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListVccsResponseBodyContentDataTags) SetTagKey(v string) *ListVccsResponseBodyContentDataTags {
	s.TagKey = &v
	return s
}

func (s *ListVccsResponseBodyContentDataTags) SetTagValue(v string) *ListVccsResponseBodyContentDataTags {
	s.TagValue = &v
	return s
}

func (s *ListVccsResponseBodyContentDataTags) Validate() error {
	return dara.Validate(s)
}

type ListVccsResponseBodyContentDataVpdBaseInfo struct {
	// The CIDR block of the VPD.
	//
	// 	- We recommend that you use an RFC private endpoint as the Lingjun CIDR block, such as 10.0.0.0/8,172.16.0.0/12,192.168.0.0/16. In scenarios where the Doringjun CIDR block is connected to each other or where the Lingjun CIDR block is connected to a VPC, make sure that the addresses do not conflict with each other.
	//
	// 	- You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, or 169.254.0.0/16 and their subnets as the primary IPv4 CIDR block of the VPD.
	//
	// example:
	//
	// 10.0.0.0/13
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1668158213000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Lingjun CIDR block instance ID
	//
	// example:
	//
	// vpd-9n7ioqrp
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// Lingjun CIDR block instance name
	//
	// example:
	//
	// yzp-rg-test3
	VpdName *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
}

func (s ListVccsResponseBodyContentDataVpdBaseInfo) String() string {
	return dara.Prettify(s)
}

func (s ListVccsResponseBodyContentDataVpdBaseInfo) GoString() string {
	return s.String()
}

func (s *ListVccsResponseBodyContentDataVpdBaseInfo) GetCidr() *string {
	return s.Cidr
}

func (s *ListVccsResponseBodyContentDataVpdBaseInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListVccsResponseBodyContentDataVpdBaseInfo) GetVpdId() *string {
	return s.VpdId
}

func (s *ListVccsResponseBodyContentDataVpdBaseInfo) GetVpdName() *string {
	return s.VpdName
}

func (s *ListVccsResponseBodyContentDataVpdBaseInfo) SetCidr(v string) *ListVccsResponseBodyContentDataVpdBaseInfo {
	s.Cidr = &v
	return s
}

func (s *ListVccsResponseBodyContentDataVpdBaseInfo) SetCreateTime(v string) *ListVccsResponseBodyContentDataVpdBaseInfo {
	s.CreateTime = &v
	return s
}

func (s *ListVccsResponseBodyContentDataVpdBaseInfo) SetVpdId(v string) *ListVccsResponseBodyContentDataVpdBaseInfo {
	s.VpdId = &v
	return s
}

func (s *ListVccsResponseBodyContentDataVpdBaseInfo) SetVpdName(v string) *ListVccsResponseBodyContentDataVpdBaseInfo {
	s.VpdName = &v
	return s
}

func (s *ListVccsResponseBodyContentDataVpdBaseInfo) Validate() error {
	return dara.Validate(s)
}
