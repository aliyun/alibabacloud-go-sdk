// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateMajorProtectionBlackIpRequest struct {
	// 防护对象1domain 	描述信息。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpiredTime *int64  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IpList      *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	TemplateId  *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateMajorProtectionBlackIpRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMajorProtectionBlackIpRequest) GoString() string {
	return s.String()
}

func (s *CreateMajorProtectionBlackIpRequest) SetDescription(v string) *CreateMajorProtectionBlackIpRequest {
	s.Description = &v
	return s
}

func (s *CreateMajorProtectionBlackIpRequest) SetExpiredTime(v int64) *CreateMajorProtectionBlackIpRequest {
	s.ExpiredTime = &v
	return s
}

func (s *CreateMajorProtectionBlackIpRequest) SetInstanceId(v string) *CreateMajorProtectionBlackIpRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateMajorProtectionBlackIpRequest) SetIpList(v string) *CreateMajorProtectionBlackIpRequest {
	s.IpList = &v
	return s
}

func (s *CreateMajorProtectionBlackIpRequest) SetRuleId(v int64) *CreateMajorProtectionBlackIpRequest {
	s.RuleId = &v
	return s
}

func (s *CreateMajorProtectionBlackIpRequest) SetTemplateId(v int64) *CreateMajorProtectionBlackIpRequest {
	s.TemplateId = &v
	return s
}

type CreateMajorProtectionBlackIpResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMajorProtectionBlackIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMajorProtectionBlackIpResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMajorProtectionBlackIpResponseBody) SetRequestId(v string) *CreateMajorProtectionBlackIpResponseBody {
	s.RequestId = &v
	return s
}

type CreateMajorProtectionBlackIpResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMajorProtectionBlackIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMajorProtectionBlackIpResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMajorProtectionBlackIpResponse) GoString() string {
	return s.String()
}

func (s *CreateMajorProtectionBlackIpResponse) SetHeaders(v map[string]*string) *CreateMajorProtectionBlackIpResponse {
	s.Headers = v
	return s
}

func (s *CreateMajorProtectionBlackIpResponse) SetStatusCode(v int32) *CreateMajorProtectionBlackIpResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMajorProtectionBlackIpResponse) SetBody(v *CreateMajorProtectionBlackIpResponseBody) *CreateMajorProtectionBlackIpResponse {
	s.Body = v
	return s
}

type DescribeInstanceRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRequest) SetRegionId(v string) *DescribeInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceRequest) SetResourceGroupId(v string) *DescribeInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeInstanceResponseBody struct {
	// 实例详情
	Details *DescribeInstanceResponseBodyDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Struct"`
	// 套餐
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// RegionId
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) SetDetails(v *DescribeInstanceResponseBodyDetails) *DescribeInstanceResponseBody {
	s.Details = v
	return s
}

func (s *DescribeInstanceResponseBody) SetEdition(v string) *DescribeInstanceResponseBody {
	s.Edition = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetInstanceId(v string) *DescribeInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRegionId(v string) *DescribeInstanceResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRequestId(v string) *DescribeInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceResponseBodyDetails struct {
	AclRuleMaxIpCount *int64 `json:"AclRuleMaxIpCount,omitempty" xml:"AclRuleMaxIpCount,omitempty"`
	// 支持扫描防护
	AntiScan *bool `json:"AntiScan,omitempty" xml:"AntiScan,omitempty"`
	// 扫描防护模板数
	AntiScanTemplateMaxCount *int64 `json:"AntiScanTemplateMaxCount,omitempty" xml:"AntiScanTemplateMaxCount,omitempty"`
	// 最大回源数
	BackendMaxCount *int64 `json:"BackendMaxCount,omitempty" xml:"BackendMaxCount,omitempty"`
	// 基础防护
	BaseWafGroup *bool `json:"BaseWafGroup,omitempty" xml:"BaseWafGroup,omitempty"`
	// 基础防护规则
	BaseWafGroupRuleInTemplateMaxCount *int64 `json:"BaseWafGroupRuleInTemplateMaxCount,omitempty" xml:"BaseWafGroupRuleInTemplateMaxCount,omitempty"`
	// 基础防护规则最大数量
	BaseWafGroupRuleTemplateMaxCount *int64 `json:"BaseWafGroupRuleTemplateMaxCount,omitempty" xml:"BaseWafGroupRuleTemplateMaxCount,omitempty"`
	// 最大可添加CNAME数
	CnameResourceMaxCount *int64 `json:"CnameResourceMaxCount,omitempty" xml:"CnameResourceMaxCount,omitempty"`
	// 支持自定义响应
	CustomResponse *bool `json:"CustomResponse,omitempty" xml:"CustomResponse,omitempty"`
	// 自定义响应模板包含规则数
	CustomResponseRuleInTemplateMaxCount *int64 `json:"CustomResponseRuleInTemplateMaxCount,omitempty" xml:"CustomResponseRuleInTemplateMaxCount,omitempty"`
	// 自定义响应模板数
	CustomResponseTemplateMaxCount *int64 `json:"CustomResponseTemplateMaxCount,omitempty" xml:"CustomResponseTemplateMaxCount,omitempty"`
	// 支持自定义规则
	CustomRule *bool `json:"CustomRule,omitempty" xml:"CustomRule,omitempty"`
	// 包含字符串
	CustomRuleAction *string `json:"CustomRuleAction,omitempty" xml:"CustomRuleAction,omitempty"`
	// 自定义规则匹配条件
	CustomRuleCondition *string `json:"CustomRuleCondition,omitempty" xml:"CustomRuleCondition,omitempty"`
	// 自定义规则模板包含规则数
	CustomRuleInTemplateMaxCount *int64 `json:"CustomRuleInTemplateMaxCount,omitempty" xml:"CustomRuleInTemplateMaxCount,omitempty"`
	// 自定义规则限速对象
	CustomRuleRatelimitor *string `json:"CustomRuleRatelimitor,omitempty" xml:"CustomRuleRatelimitor,omitempty"`
	// 自定义规则模板数
	CustomRuleTemplateMaxCount *int64 `json:"CustomRuleTemplateMaxCount,omitempty" xml:"CustomRuleTemplateMaxCount,omitempty"`
	// 最大防护组数量
	DefenseGroupMaxCount *int64 `json:"DefenseGroupMaxCount,omitempty" xml:"DefenseGroupMaxCount,omitempty"`
	// 一个防护组内最大包含对象数量
	DefenseObjectInGroupMaxCount *int64 `json:"DefenseObjectInGroupMaxCount,omitempty" xml:"DefenseObjectInGroupMaxCount,omitempty"`
	// 一个模板内关联对象的最大数量
	DefenseObjectInTemplateMaxCount *int64 `json:"DefenseObjectInTemplateMaxCount,omitempty" xml:"DefenseObjectInTemplateMaxCount,omitempty"`
	// 最大防护对象数量
	DefenseObjectMaxCount *int64 `json:"DefenseObjectMaxCount,omitempty" xml:"DefenseObjectMaxCount,omitempty"`
	// 独享IP
	ExclusiveIp *bool `json:"ExclusiveIp,omitempty" xml:"ExclusiveIp,omitempty"`
	// Gslb
	Gslb *bool `json:"Gslb,omitempty" xml:"Gslb,omitempty"`
	// HTTP端口可用范围
	HttpPorts *string `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty"`
	// HTTPS端口可用范围
	HttpsPorts *string `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty"`
	// 支持IP黑名单
	IpBlacklist *bool `json:"IpBlacklist,omitempty" xml:"IpBlacklist,omitempty"`
	// IP黑名单规则包含IP数
	IpBlacklistIpInRuleMaxCount *int64 `json:"IpBlacklistIpInRuleMaxCount,omitempty" xml:"IpBlacklistIpInRuleMaxCount,omitempty"`
	// IP黑名单模板包含规则数
	IpBlacklistRuleInTemplateMaxCount *int64 `json:"IpBlacklistRuleInTemplateMaxCount,omitempty" xml:"IpBlacklistRuleInTemplateMaxCount,omitempty"`
	// /黑名单模板数
	IpBlacklistTemplateMaxCount *int64 `json:"IpBlacklistTemplateMaxCount,omitempty" xml:"IpBlacklistTemplateMaxCount,omitempty"`
	// Ipv6
	Ipv6 *bool `json:"Ipv6,omitempty" xml:"Ipv6,omitempty"`
	// 日志服务是否开启
	LogService *bool `json:"LogService,omitempty" xml:"LogService,omitempty"`
	// 是否支持重保
	MajorProtection *bool `json:"MajorProtection,omitempty" xml:"MajorProtection,omitempty"`
	// 重保模板的最大数量
	MajorProtectionTemplateMaxCount *int64 `json:"MajorProtectionTemplateMaxCount,omitempty" xml:"MajorProtectionTemplateMaxCount,omitempty"`
	// 海量IP单次上传文件IP的最大数量
	VastIpBlacklistInFileMaxCount *int64 `json:"VastIpBlacklistInFileMaxCount,omitempty" xml:"VastIpBlacklistInFileMaxCount,omitempty"`
	// 海量IP单次页面操作的最大数量
	VastIpBlacklistInOperationMaxCount *int64 `json:"VastIpBlacklistInOperationMaxCount,omitempty" xml:"VastIpBlacklistInOperationMaxCount,omitempty"`
	// 海量IP的最大数量（单用户）
	VastIpBlacklistMaxCount *int64 `json:"VastIpBlacklistMaxCount,omitempty" xml:"VastIpBlacklistMaxCount,omitempty"`
	// 是否支持白名单
	Whitelist *bool `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
	// 白名单规则匹配条件
	WhitelistLogical *string `json:"WhitelistLogical,omitempty" xml:"WhitelistLogical,omitempty"`
	// 白名单规则匹配条件
	WhitelistRuleCondition *string `json:"WhitelistRuleCondition,omitempty" xml:"WhitelistRuleCondition,omitempty"`
	// 白名单模板包含规则数
	WhitelistRuleInTemplateMaxCount *int64 `json:"WhitelistRuleInTemplateMaxCount,omitempty" xml:"WhitelistRuleInTemplateMaxCount,omitempty"`
	// 白名单模板数
	WhitelistTemplateMaxCount *int64 `json:"WhitelistTemplateMaxCount,omitempty" xml:"WhitelistTemplateMaxCount,omitempty"`
}

func (s DescribeInstanceResponseBodyDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyDetails) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyDetails) SetAclRuleMaxIpCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.AclRuleMaxIpCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetAntiScan(v bool) *DescribeInstanceResponseBodyDetails {
	s.AntiScan = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetAntiScanTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.AntiScanTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBackendMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.BackendMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBaseWafGroup(v bool) *DescribeInstanceResponseBodyDetails {
	s.BaseWafGroup = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBaseWafGroupRuleInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.BaseWafGroupRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBaseWafGroupRuleTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.BaseWafGroupRuleTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCnameResourceMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.CnameResourceMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomResponse(v bool) *DescribeInstanceResponseBodyDetails {
	s.CustomResponse = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomResponseRuleInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.CustomResponseRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomResponseTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.CustomResponseTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomRule(v bool) *DescribeInstanceResponseBodyDetails {
	s.CustomRule = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomRuleAction(v string) *DescribeInstanceResponseBodyDetails {
	s.CustomRuleAction = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomRuleCondition(v string) *DescribeInstanceResponseBodyDetails {
	s.CustomRuleCondition = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomRuleInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.CustomRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomRuleRatelimitor(v string) *DescribeInstanceResponseBodyDetails {
	s.CustomRuleRatelimitor = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomRuleTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.CustomRuleTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetDefenseGroupMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.DefenseGroupMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetDefenseObjectInGroupMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.DefenseObjectInGroupMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetDefenseObjectInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.DefenseObjectInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetDefenseObjectMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.DefenseObjectMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetExclusiveIp(v bool) *DescribeInstanceResponseBodyDetails {
	s.ExclusiveIp = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetGslb(v bool) *DescribeInstanceResponseBodyDetails {
	s.Gslb = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetHttpPorts(v string) *DescribeInstanceResponseBodyDetails {
	s.HttpPorts = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetHttpsPorts(v string) *DescribeInstanceResponseBodyDetails {
	s.HttpsPorts = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetIpBlacklist(v bool) *DescribeInstanceResponseBodyDetails {
	s.IpBlacklist = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetIpBlacklistIpInRuleMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.IpBlacklistIpInRuleMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetIpBlacklistRuleInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.IpBlacklistRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetIpBlacklistTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.IpBlacklistTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetIpv6(v bool) *DescribeInstanceResponseBodyDetails {
	s.Ipv6 = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetLogService(v bool) *DescribeInstanceResponseBodyDetails {
	s.LogService = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetMajorProtection(v bool) *DescribeInstanceResponseBodyDetails {
	s.MajorProtection = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetMajorProtectionTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.MajorProtectionTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetVastIpBlacklistInFileMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.VastIpBlacklistInFileMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetVastIpBlacklistInOperationMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.VastIpBlacklistInOperationMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetVastIpBlacklistMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.VastIpBlacklistMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetWhitelist(v bool) *DescribeInstanceResponseBodyDetails {
	s.Whitelist = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetWhitelistLogical(v string) *DescribeInstanceResponseBodyDetails {
	s.WhitelistLogical = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetWhitelistRuleCondition(v string) *DescribeInstanceResponseBodyDetails {
	s.WhitelistRuleCondition = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetWhitelistRuleInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.WhitelistRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetWhitelistTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.WhitelistTemplateMaxCount = &v
	return s
}

type DescribeInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponse) SetHeaders(v map[string]*string) *DescribeInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceResponse) SetStatusCode(v int32) *DescribeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceResponse) SetBody(v *DescribeInstanceResponseBody) *DescribeInstanceResponse {
	s.Body = v
	return s
}

type DescribeInstanceCompatibleRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeInstanceCompatibleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceCompatibleRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceCompatibleRequest) SetRegionId(v string) *DescribeInstanceCompatibleRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceCompatibleRequest) SetResourceGroupId(v string) *DescribeInstanceCompatibleRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeInstanceCompatibleResponseBody struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// 实例详情
	Details *DescribeInstanceCompatibleResponseBodyDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Struct"`
	// 套餐
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 付费类型
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// RegionId
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceCompatibleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceCompatibleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceCompatibleResponseBody) SetCommodityCode(v string) *DescribeInstanceCompatibleResponseBody {
	s.CommodityCode = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBody) SetDetails(v *DescribeInstanceCompatibleResponseBodyDetails) *DescribeInstanceCompatibleResponseBody {
	s.Details = v
	return s
}

func (s *DescribeInstanceCompatibleResponseBody) SetEdition(v string) *DescribeInstanceCompatibleResponseBody {
	s.Edition = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBody) SetInstanceId(v string) *DescribeInstanceCompatibleResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBody) SetPayType(v string) *DescribeInstanceCompatibleResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBody) SetRegionId(v string) *DescribeInstanceCompatibleResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBody) SetRequestId(v string) *DescribeInstanceCompatibleResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceCompatibleResponseBodyDetails struct {
	// 支持扫描防护
	AntiScan *bool `json:"AntiScan,omitempty" xml:"AntiScan,omitempty"`
	// 扫描防护模板数
	AntiScanTemplateMaxCount *int64 `json:"AntiScanTemplateMaxCount,omitempty" xml:"AntiScanTemplateMaxCount,omitempty"`
	// 最大回源数
	BackendMaxCount *int64 `json:"BackendMaxCount,omitempty" xml:"BackendMaxCount,omitempty"`
	// 基础防护
	BaseWafGroup *bool `json:"BaseWafGroup,omitempty" xml:"BaseWafGroup,omitempty"`
	// 基础防护规则
	BaseWafGroupRuleInTemplateMaxCount *int64 `json:"BaseWafGroupRuleInTemplateMaxCount,omitempty" xml:"BaseWafGroupRuleInTemplateMaxCount,omitempty"`
	// 基础防护规则最大数量
	BaseWafGroupRuleTemplateMaxCount *int64 `json:"BaseWafGroupRuleTemplateMaxCount,omitempty" xml:"BaseWafGroupRuleTemplateMaxCount,omitempty"`
	// 最大可添加CNAME数
	CnameResourceMaxCount *int64 `json:"CnameResourceMaxCount,omitempty" xml:"CnameResourceMaxCount,omitempty"`
	// 支持自定义响应
	CustomResponse *bool `json:"CustomResponse,omitempty" xml:"CustomResponse,omitempty"`
	// 自定义响应模板包含规则数
	CustomResponseRuleInTemplateMaxCount *int64 `json:"CustomResponseRuleInTemplateMaxCount,omitempty" xml:"CustomResponseRuleInTemplateMaxCount,omitempty"`
	// 自定义响应模板数
	CustomResponseTemplateMaxCount *int64 `json:"CustomResponseTemplateMaxCount,omitempty" xml:"CustomResponseTemplateMaxCount,omitempty"`
	// 支持自定义规则
	CustomRule *bool `json:"CustomRule,omitempty" xml:"CustomRule,omitempty"`
	// 包含字符串
	CustomRuleAction *string `json:"CustomRuleAction,omitempty" xml:"CustomRuleAction,omitempty"`
	// 自定义规则匹配条件
	CustomRuleCondition *string `json:"CustomRuleCondition,omitempty" xml:"CustomRuleCondition,omitempty"`
	// 自定义规则模板包含规则数
	CustomRuleInTemplateMaxCount *int64 `json:"CustomRuleInTemplateMaxCount,omitempty" xml:"CustomRuleInTemplateMaxCount,omitempty"`
	// 自定义规则限速对象
	CustomRuleRatelimitor *string `json:"CustomRuleRatelimitor,omitempty" xml:"CustomRuleRatelimitor,omitempty"`
	// 自定义规则模板数
	CustomRuleTemplateMaxCount *int64 `json:"CustomRuleTemplateMaxCount,omitempty" xml:"CustomRuleTemplateMaxCount,omitempty"`
	// 最大防护组数量
	DefenseGroupMaxCount *int64 `json:"DefenseGroupMaxCount,omitempty" xml:"DefenseGroupMaxCount,omitempty"`
	// 一个防护组内最大包含对象数量
	DefenseObjectInGroupMaxCount *int64 `json:"DefenseObjectInGroupMaxCount,omitempty" xml:"DefenseObjectInGroupMaxCount,omitempty"`
	// 一个模板内关联对象的最大数量
	DefenseObjectInTemplateMaxCount *int64 `json:"DefenseObjectInTemplateMaxCount,omitempty" xml:"DefenseObjectInTemplateMaxCount,omitempty"`
	// 最大防护对象数量
	DefenseObjectMaxCount *int64 `json:"DefenseObjectMaxCount,omitempty" xml:"DefenseObjectMaxCount,omitempty"`
	// 独享IP
	ExclusiveIp *bool `json:"ExclusiveIp,omitempty" xml:"ExclusiveIp,omitempty"`
	// Gslb
	Gslb *bool `json:"Gslb,omitempty" xml:"Gslb,omitempty"`
	// 支持IP黑名单
	IpBlacklist *bool `json:"IpBlacklist,omitempty" xml:"IpBlacklist,omitempty"`
	// IP黑名单规则包含IP数
	IpBlacklistIpInRuleMaxCount *int64 `json:"IpBlacklistIpInRuleMaxCount,omitempty" xml:"IpBlacklistIpInRuleMaxCount,omitempty"`
	// IP黑名单模板包含规则数
	IpBlacklistRuleInTemplateMaxCount *int64 `json:"IpBlacklistRuleInTemplateMaxCount,omitempty" xml:"IpBlacklistRuleInTemplateMaxCount,omitempty"`
	// /黑名单模板数
	IpBlacklistTemplateMaxCount *int64 `json:"IpBlacklistTemplateMaxCount,omitempty" xml:"IpBlacklistTemplateMaxCount,omitempty"`
	// Ipv6
	Ipv6 *bool `json:"Ipv6,omitempty" xml:"Ipv6,omitempty"`
	// 日志服务是否开启
	LogService *bool `json:"LogService,omitempty" xml:"LogService,omitempty"`
	// 是否支持白名单
	Whitelist *bool `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
	// 白名单规则匹配条件
	WhitelistLogical *string `json:"WhitelistLogical,omitempty" xml:"WhitelistLogical,omitempty"`
	// 白名单规则匹配条件
	WhitelistRuleCondition *string `json:"WhitelistRuleCondition,omitempty" xml:"WhitelistRuleCondition,omitempty"`
	// 白名单模板包含规则数
	WhitelistRuleInTemplateMaxCount *int64 `json:"WhitelistRuleInTemplateMaxCount,omitempty" xml:"WhitelistRuleInTemplateMaxCount,omitempty"`
	// 白名单模板数
	WhitelistTemplateMaxCount *int64 `json:"WhitelistTemplateMaxCount,omitempty" xml:"WhitelistTemplateMaxCount,omitempty"`
}

func (s DescribeInstanceCompatibleResponseBodyDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceCompatibleResponseBodyDetails) GoString() string {
	return s.String()
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetAntiScan(v bool) *DescribeInstanceCompatibleResponseBodyDetails {
	s.AntiScan = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetAntiScanTemplateMaxCount(v int64) *DescribeInstanceCompatibleResponseBodyDetails {
	s.AntiScanTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetBackendMaxCount(v int64) *DescribeInstanceCompatibleResponseBodyDetails {
	s.BackendMaxCount = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetBaseWafGroup(v bool) *DescribeInstanceCompatibleResponseBodyDetails {
	s.BaseWafGroup = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetBaseWafGroupRuleInTemplateMaxCount(v int64) *DescribeInstanceCompatibleResponseBodyDetails {
	s.BaseWafGroupRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetBaseWafGroupRuleTemplateMaxCount(v int64) *DescribeInstanceCompatibleResponseBodyDetails {
	s.BaseWafGroupRuleTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetCnameResourceMaxCount(v int64) *DescribeInstanceCompatibleResponseBodyDetails {
	s.CnameResourceMaxCount = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetCustomResponse(v bool) *DescribeInstanceCompatibleResponseBodyDetails {
	s.CustomResponse = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetCustomResponseRuleInTemplateMaxCount(v int64) *DescribeInstanceCompatibleResponseBodyDetails {
	s.CustomResponseRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetCustomResponseTemplateMaxCount(v int64) *DescribeInstanceCompatibleResponseBodyDetails {
	s.CustomResponseTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetCustomRule(v bool) *DescribeInstanceCompatibleResponseBodyDetails {
	s.CustomRule = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetCustomRuleAction(v string) *DescribeInstanceCompatibleResponseBodyDetails {
	s.CustomRuleAction = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetCustomRuleCondition(v string) *DescribeInstanceCompatibleResponseBodyDetails {
	s.CustomRuleCondition = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetCustomRuleInTemplateMaxCount(v int64) *DescribeInstanceCompatibleResponseBodyDetails {
	s.CustomRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetCustomRuleRatelimitor(v string) *DescribeInstanceCompatibleResponseBodyDetails {
	s.CustomRuleRatelimitor = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetCustomRuleTemplateMaxCount(v int64) *DescribeInstanceCompatibleResponseBodyDetails {
	s.CustomRuleTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetDefenseGroupMaxCount(v int64) *DescribeInstanceCompatibleResponseBodyDetails {
	s.DefenseGroupMaxCount = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetDefenseObjectInGroupMaxCount(v int64) *DescribeInstanceCompatibleResponseBodyDetails {
	s.DefenseObjectInGroupMaxCount = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetDefenseObjectInTemplateMaxCount(v int64) *DescribeInstanceCompatibleResponseBodyDetails {
	s.DefenseObjectInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetDefenseObjectMaxCount(v int64) *DescribeInstanceCompatibleResponseBodyDetails {
	s.DefenseObjectMaxCount = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetExclusiveIp(v bool) *DescribeInstanceCompatibleResponseBodyDetails {
	s.ExclusiveIp = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetGslb(v bool) *DescribeInstanceCompatibleResponseBodyDetails {
	s.Gslb = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetIpBlacklist(v bool) *DescribeInstanceCompatibleResponseBodyDetails {
	s.IpBlacklist = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetIpBlacklistIpInRuleMaxCount(v int64) *DescribeInstanceCompatibleResponseBodyDetails {
	s.IpBlacklistIpInRuleMaxCount = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetIpBlacklistRuleInTemplateMaxCount(v int64) *DescribeInstanceCompatibleResponseBodyDetails {
	s.IpBlacklistRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetIpBlacklistTemplateMaxCount(v int64) *DescribeInstanceCompatibleResponseBodyDetails {
	s.IpBlacklistTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetIpv6(v bool) *DescribeInstanceCompatibleResponseBodyDetails {
	s.Ipv6 = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetLogService(v bool) *DescribeInstanceCompatibleResponseBodyDetails {
	s.LogService = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetWhitelist(v bool) *DescribeInstanceCompatibleResponseBodyDetails {
	s.Whitelist = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetWhitelistLogical(v string) *DescribeInstanceCompatibleResponseBodyDetails {
	s.WhitelistLogical = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetWhitelistRuleCondition(v string) *DescribeInstanceCompatibleResponseBodyDetails {
	s.WhitelistRuleCondition = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetWhitelistRuleInTemplateMaxCount(v int64) *DescribeInstanceCompatibleResponseBodyDetails {
	s.WhitelistRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceCompatibleResponseBodyDetails) SetWhitelistTemplateMaxCount(v int64) *DescribeInstanceCompatibleResponseBodyDetails {
	s.WhitelistTemplateMaxCount = &v
	return s
}

type DescribeInstanceCompatibleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceCompatibleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceCompatibleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceCompatibleResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceCompatibleResponse) SetHeaders(v map[string]*string) *DescribeInstanceCompatibleResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceCompatibleResponse) SetStatusCode(v int32) *DescribeInstanceCompatibleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceCompatibleResponse) SetBody(v *DescribeInstanceCompatibleResponseBody) *DescribeInstanceCompatibleResponse {
	s.Body = v
	return s
}

type DescribeInstanceExtendRequest struct {
	Edition         *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeInstanceExtendRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceExtendRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceExtendRequest) SetEdition(v string) *DescribeInstanceExtendRequest {
	s.Edition = &v
	return s
}

func (s *DescribeInstanceExtendRequest) SetRegionId(v string) *DescribeInstanceExtendRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceExtendRequest) SetResourceGroupId(v string) *DescribeInstanceExtendRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeInstanceExtendResponseBody struct {
	Instances []*DescribeInstanceExtendResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceExtendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceExtendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceExtendResponseBody) SetInstances(v []*DescribeInstanceExtendResponseBodyInstances) *DescribeInstanceExtendResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstanceExtendResponseBody) SetRequestId(v string) *DescribeInstanceExtendResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceExtendResponseBodyInstances struct {
	ExpireTime *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInstanceExtendResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceExtendResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstanceExtendResponseBodyInstances) SetExpireTime(v int64) *DescribeInstanceExtendResponseBodyInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstanceExtendResponseBodyInstances) SetInstanceId(v string) *DescribeInstanceExtendResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceExtendResponseBodyInstances) SetRegionId(v string) *DescribeInstanceExtendResponseBodyInstances {
	s.RegionId = &v
	return s
}

type DescribeInstanceExtendResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceExtendResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceExtendResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceExtendResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceExtendResponse) SetHeaders(v map[string]*string) *DescribeInstanceExtendResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceExtendResponse) SetStatusCode(v int32) *DescribeInstanceExtendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceExtendResponse) SetBody(v *DescribeInstanceExtendResponseBody) *DescribeInstanceExtendResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-qingdao":            tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-beijing":            tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-chengdu":            tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-huhehaote":          tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-hangzhou":           tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai":           tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-shenzhen":           tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-heyuan":             tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-wulanchabu":         tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-hongkong":           tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-1":        tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3":        tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":        tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":             tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"us-west-1":             tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"us-east-1":             tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"eu-central-1":          tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"me-east-1":             tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":            tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("waf-openapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMajorProtectionBlackIpWithOptions(request *CreateMajorProtectionBlackIpRequest, runtime *util.RuntimeOptions) (_result *CreateMajorProtectionBlackIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExpiredTime)) {
		query["ExpiredTime"] = request.ExpiredTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IpList)) {
		query["IpList"] = request.IpList
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMajorProtectionBlackIp"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMajorProtectionBlackIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMajorProtectionBlackIp(request *CreateMajorProtectionBlackIpRequest) (_result *CreateMajorProtectionBlackIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMajorProtectionBlackIpResponse{}
	_body, _err := client.CreateMajorProtectionBlackIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceWithOptions(request *DescribeInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstance"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstance(request *DescribeInstanceRequest) (_result *DescribeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.DescribeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceCompatibleWithOptions(request *DescribeInstanceCompatibleRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceCompatibleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceCompatible"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceCompatibleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceCompatible(request *DescribeInstanceCompatibleRequest) (_result *DescribeInstanceCompatibleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceCompatibleResponse{}
	_body, _err := client.DescribeInstanceCompatibleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceExtendWithOptions(request *DescribeInstanceExtendRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceExtendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Edition)) {
		query["Edition"] = request.Edition
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceExtend"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceExtendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceExtend(request *DescribeInstanceExtendRequest) (_result *DescribeInstanceExtendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceExtendResponse{}
	_body, _err := client.DescribeInstanceExtendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
