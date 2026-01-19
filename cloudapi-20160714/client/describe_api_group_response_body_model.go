// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBasePath(v string) *DescribeApiGroupResponseBody
	GetBasePath() *string
	SetBillingStatus(v string) *DescribeApiGroupResponseBody
	GetBillingStatus() *string
	SetCloudMarketCommodity(v bool) *DescribeApiGroupResponseBody
	GetCloudMarketCommodity() *bool
	SetCmsMonitorGroup(v string) *DescribeApiGroupResponseBody
	GetCmsMonitorGroup() *string
	SetCompatibleFlags(v string) *DescribeApiGroupResponseBody
	GetCompatibleFlags() *string
	SetCreatedTime(v string) *DescribeApiGroupResponseBody
	GetCreatedTime() *string
	SetCustomAppCodeConfig(v string) *DescribeApiGroupResponseBody
	GetCustomAppCodeConfig() *string
	SetCustomDomains(v *DescribeApiGroupResponseBodyCustomDomains) *DescribeApiGroupResponseBody
	GetCustomDomains() *DescribeApiGroupResponseBodyCustomDomains
	SetCustomTraceConfig(v string) *DescribeApiGroupResponseBody
	GetCustomTraceConfig() *string
	SetCustomerConfigs(v string) *DescribeApiGroupResponseBody
	GetCustomerConfigs() *string
	SetDedicatedInstanceType(v string) *DescribeApiGroupResponseBody
	GetDedicatedInstanceType() *string
	SetDefaultDomain(v string) *DescribeApiGroupResponseBody
	GetDefaultDomain() *string
	SetDescription(v string) *DescribeApiGroupResponseBody
	GetDescription() *string
	SetDisableInnerDomain(v bool) *DescribeApiGroupResponseBody
	GetDisableInnerDomain() *bool
	SetGroupId(v string) *DescribeApiGroupResponseBody
	GetGroupId() *string
	SetGroupName(v string) *DescribeApiGroupResponseBody
	GetGroupName() *string
	SetHttpsPolicy(v string) *DescribeApiGroupResponseBody
	GetHttpsPolicy() *string
	SetIllegalStatus(v string) *DescribeApiGroupResponseBody
	GetIllegalStatus() *string
	SetInstanceId(v string) *DescribeApiGroupResponseBody
	GetInstanceId() *string
	SetInstanceType(v string) *DescribeApiGroupResponseBody
	GetInstanceType() *string
	SetIpv6Status(v string) *DescribeApiGroupResponseBody
	GetIpv6Status() *string
	SetMigrationError(v string) *DescribeApiGroupResponseBody
	GetMigrationError() *string
	SetMigrationStatus(v string) *DescribeApiGroupResponseBody
	GetMigrationStatus() *string
	SetModifiedTime(v string) *DescribeApiGroupResponseBody
	GetModifiedTime() *string
	SetPassthroughHeaders(v string) *DescribeApiGroupResponseBody
	GetPassthroughHeaders() *string
	SetRegionId(v string) *DescribeApiGroupResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeApiGroupResponseBody
	GetRequestId() *string
	SetStageItems(v *DescribeApiGroupResponseBodyStageItems) *DescribeApiGroupResponseBody
	GetStageItems() *DescribeApiGroupResponseBodyStageItems
	SetStatus(v string) *DescribeApiGroupResponseBody
	GetStatus() *string
	SetSubDomain(v string) *DescribeApiGroupResponseBody
	GetSubDomain() *string
	SetTrafficLimit(v int32) *DescribeApiGroupResponseBody
	GetTrafficLimit() *int32
	SetUserLogConfig(v string) *DescribeApiGroupResponseBody
	GetUserLogConfig() *string
	SetVpcDomain(v string) *DescribeApiGroupResponseBody
	GetVpcDomain() *string
	SetVpcSlbIntranetDomain(v string) *DescribeApiGroupResponseBody
	GetVpcSlbIntranetDomain() *string
}

type DescribeApiGroupResponseBody struct {
	// The root path of the API.
	//
	// example:
	//
	// /qqq
	BasePath *string `json:"BasePath,omitempty" xml:"BasePath,omitempty"`
	// The billing status of the API group.
	//
	// 	- **NORMAL**: The API group is normal.
	//
	// 	- **LOCKED**: The API group is locked due to overdue payments.
	//
	// example:
	//
	// NORMAL
	BillingStatus *string `json:"BillingStatus,omitempty" xml:"BillingStatus,omitempty"`
	// The products on Alibaba Cloud Marketplace.
	//
	// example:
	//
	// false
	CloudMarketCommodity *bool `json:"CloudMarketCommodity,omitempty" xml:"CloudMarketCommodity,omitempty"`
	// The CloudMonitor application group.
	//
	// example:
	//
	// 217008423
	CmsMonitorGroup *string `json:"CmsMonitorGroup,omitempty" xml:"CmsMonitorGroup,omitempty"`
	// The list of associated tags. Separate multiple tags with commas (,).
	//
	// example:
	//
	// depart:dep1
	CompatibleFlags *string `json:"CompatibleFlags,omitempty" xml:"CompatibleFlags,omitempty"`
	// The creation time (UTC) of the API group.
	//
	// example:
	//
	// 2016-08-01T06:53:02Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The custom appcode configuration.
	//
	// example:
	//
	// {"location":"HEADER","name":"myAppCodeHeader"}
	CustomAppCodeConfig *string `json:"CustomAppCodeConfig,omitempty" xml:"CustomAppCodeConfig,omitempty"`
	// The details about the custom domain name.
	CustomDomains *DescribeApiGroupResponseBodyCustomDomains `json:"CustomDomains,omitempty" xml:"CustomDomains,omitempty" type:"Struct"`
	// The custom trace configuration.
	//
	// example:
	//
	// {\\"parameterLocation\\":\\"HEADER\\",\\"parameterName\\":\\"traceId\\"}
	CustomTraceConfig *string `json:"CustomTraceConfig,omitempty" xml:"CustomTraceConfig,omitempty"`
	// The list of custom configuration items.
	//
	// example:
	//
	// removeResponseServerHeader
	CustomerConfigs *string `json:"CustomerConfigs,omitempty" xml:"CustomerConfigs,omitempty"`
	// The type of exclusive instance where the group is located
	//
	// - VPC fusion type exclusive instance: vpc_connect
	//
	// - Traditional type exclusive instance: normal
	//
	// example:
	//
	// normal
	DedicatedInstanceType *string `json:"DedicatedInstanceType,omitempty" xml:"DedicatedInstanceType,omitempty"`
	// The default domain name.
	//
	// example:
	//
	// mkt.api.gaore.com
	DefaultDomain *string `json:"DefaultDomain,omitempty" xml:"DefaultDomain,omitempty"`
	// The description of the API group.
	//
	// example:
	//
	// New weather informations.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether access over the public second-level domain name is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	DisableInnerDomain *bool `json:"DisableInnerDomain,omitempty" xml:"DisableInnerDomain,omitempty"`
	// The ID of the API group. This ID is generated by the system and globally unique.
	//
	// example:
	//
	// 523e8dc7bbe04613b5b1d726c2a7889d
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the group to which the API belongs.
	//
	// example:
	//
	// NewWeather
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The HTTPS policy.
	//
	// example:
	//
	// HTTPS2_TLS1_0
	HttpsPolicy *string `json:"HttpsPolicy,omitempty" xml:"HttpsPolicy,omitempty"`
	// The validity status of the API group. Valid values:
	//
	// 	- **NORMAL**: The API group is normal.
	//
	// 	- **LOCKED**: The API group is locked because it is not valid.
	//
	// example:
	//
	// NORMAL
	IllegalStatus *string `json:"IllegalStatus,omitempty" xml:"IllegalStatus,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// apigateway-cn-v6419k43xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the instance.
	//
	// 	- CLASSIC_SHARED: shared instance that uses the classic network configuration
	//
	// 	- VPC_SHARED: shared instance that uses VPC
	//
	// 	- VPC_DEDICATED: dedicated instance that uses VPC
	//
	// example:
	//
	// VPC_SHARED
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The IPv6 status.
	//
	// example:
	//
	// UNBIND
	Ipv6Status *string `json:"Ipv6Status,omitempty" xml:"Ipv6Status,omitempty"`
	// The reason for the failure of the group migration instance task. When the value of the MigrationStatus parameter is Failed, it is not empty.
	//
	// example:
	//
	// The current instance conflicts with the target instance.
	MigrationError *string `json:"MigrationError,omitempty" xml:"MigrationError,omitempty"`
	// Group migration instance task status
	//
	// - Running
	//
	// - Success
	//
	// - Failed
	//
	// example:
	//
	// Fail
	MigrationStatus *string `json:"MigrationStatus,omitempty" xml:"MigrationStatus,omitempty"`
	// The last modification time (UTC) of the API group.
	//
	// example:
	//
	// 2016-08-01T06:54:32Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// Specifies whether to pass headers.
	//
	// example:
	//
	// eagleeye-rpcid,x-b3-traceid
	PassthroughHeaders *string `json:"PassthroughHeaders,omitempty" xml:"PassthroughHeaders,omitempty"`
	// The region to which the API group belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 03442A3D-3B7D-434C-8A95-A5FEB989B519
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The runtime environment information.
	StageItems *DescribeApiGroupResponseBodyStageItems `json:"StageItems,omitempty" xml:"StageItems,omitempty" type:"Struct"`
	// The status of the API group.
	//
	// 	- **NORMAL**: The API group is normal.
	//
	// 	- **DELETE**: The API group is deleted.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The second-level domain name automatically assigned to the API group.
	//
	// example:
	//
	// 27d50c0f2e54b359919923d908bb015-cn-hangzhou.alicloudapi.com
	SubDomain *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	// The upper QPS limit of the API group. The default value is 500. You can increase the upper limit by submitting an application.
	//
	// example:
	//
	// 500
	TrafficLimit *int32 `json:"TrafficLimit,omitempty" xml:"TrafficLimit,omitempty"`
	// The user log settings.
	//
	// example:
	//
	// {\\"requestBody\\":true,\\"responseBody\\":true,\\"queryString\\":\\"test\\",\\"requestHeaders\\":\\"test\\",\\"responseHeaders\\":\\"test\\",\\"jwtClaims\\":\\"test\\"}
	UserLogConfig *string `json:"UserLogConfig,omitempty" xml:"UserLogConfig,omitempty"`
	// The VPC domain name.
	//
	// example:
	//
	// e4****7151954***acbd9f7****1058a-ap-southeast-1-vpc.alicloudapi.com
	VpcDomain *string `json:"VpcDomain,omitempty" xml:"VpcDomain,omitempty"`
	// The VPC SLB domain name.
	//
	// example:
	//
	// 257e9d450e924d00b976b0ecfb7184c2-cn-beijing-intranet.alicloudapi.com
	VpcSlbIntranetDomain *string `json:"VpcSlbIntranetDomain,omitempty" xml:"VpcSlbIntranetDomain,omitempty"`
}

func (s DescribeApiGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupResponseBody) GetBasePath() *string {
	return s.BasePath
}

func (s *DescribeApiGroupResponseBody) GetBillingStatus() *string {
	return s.BillingStatus
}

func (s *DescribeApiGroupResponseBody) GetCloudMarketCommodity() *bool {
	return s.CloudMarketCommodity
}

func (s *DescribeApiGroupResponseBody) GetCmsMonitorGroup() *string {
	return s.CmsMonitorGroup
}

func (s *DescribeApiGroupResponseBody) GetCompatibleFlags() *string {
	return s.CompatibleFlags
}

func (s *DescribeApiGroupResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeApiGroupResponseBody) GetCustomAppCodeConfig() *string {
	return s.CustomAppCodeConfig
}

func (s *DescribeApiGroupResponseBody) GetCustomDomains() *DescribeApiGroupResponseBodyCustomDomains {
	return s.CustomDomains
}

func (s *DescribeApiGroupResponseBody) GetCustomTraceConfig() *string {
	return s.CustomTraceConfig
}

func (s *DescribeApiGroupResponseBody) GetCustomerConfigs() *string {
	return s.CustomerConfigs
}

func (s *DescribeApiGroupResponseBody) GetDedicatedInstanceType() *string {
	return s.DedicatedInstanceType
}

func (s *DescribeApiGroupResponseBody) GetDefaultDomain() *string {
	return s.DefaultDomain
}

func (s *DescribeApiGroupResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeApiGroupResponseBody) GetDisableInnerDomain() *bool {
	return s.DisableInnerDomain
}

func (s *DescribeApiGroupResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApiGroupResponseBody) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeApiGroupResponseBody) GetHttpsPolicy() *string {
	return s.HttpsPolicy
}

func (s *DescribeApiGroupResponseBody) GetIllegalStatus() *string {
	return s.IllegalStatus
}

func (s *DescribeApiGroupResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApiGroupResponseBody) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeApiGroupResponseBody) GetIpv6Status() *string {
	return s.Ipv6Status
}

func (s *DescribeApiGroupResponseBody) GetMigrationError() *string {
	return s.MigrationError
}

func (s *DescribeApiGroupResponseBody) GetMigrationStatus() *string {
	return s.MigrationStatus
}

func (s *DescribeApiGroupResponseBody) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeApiGroupResponseBody) GetPassthroughHeaders() *string {
	return s.PassthroughHeaders
}

func (s *DescribeApiGroupResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApiGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApiGroupResponseBody) GetStageItems() *DescribeApiGroupResponseBodyStageItems {
	return s.StageItems
}

func (s *DescribeApiGroupResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeApiGroupResponseBody) GetSubDomain() *string {
	return s.SubDomain
}

func (s *DescribeApiGroupResponseBody) GetTrafficLimit() *int32 {
	return s.TrafficLimit
}

func (s *DescribeApiGroupResponseBody) GetUserLogConfig() *string {
	return s.UserLogConfig
}

func (s *DescribeApiGroupResponseBody) GetVpcDomain() *string {
	return s.VpcDomain
}

func (s *DescribeApiGroupResponseBody) GetVpcSlbIntranetDomain() *string {
	return s.VpcSlbIntranetDomain
}

func (s *DescribeApiGroupResponseBody) SetBasePath(v string) *DescribeApiGroupResponseBody {
	s.BasePath = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetBillingStatus(v string) *DescribeApiGroupResponseBody {
	s.BillingStatus = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetCloudMarketCommodity(v bool) *DescribeApiGroupResponseBody {
	s.CloudMarketCommodity = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetCmsMonitorGroup(v string) *DescribeApiGroupResponseBody {
	s.CmsMonitorGroup = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetCompatibleFlags(v string) *DescribeApiGroupResponseBody {
	s.CompatibleFlags = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetCreatedTime(v string) *DescribeApiGroupResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetCustomAppCodeConfig(v string) *DescribeApiGroupResponseBody {
	s.CustomAppCodeConfig = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetCustomDomains(v *DescribeApiGroupResponseBodyCustomDomains) *DescribeApiGroupResponseBody {
	s.CustomDomains = v
	return s
}

func (s *DescribeApiGroupResponseBody) SetCustomTraceConfig(v string) *DescribeApiGroupResponseBody {
	s.CustomTraceConfig = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetCustomerConfigs(v string) *DescribeApiGroupResponseBody {
	s.CustomerConfigs = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetDedicatedInstanceType(v string) *DescribeApiGroupResponseBody {
	s.DedicatedInstanceType = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetDefaultDomain(v string) *DescribeApiGroupResponseBody {
	s.DefaultDomain = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetDescription(v string) *DescribeApiGroupResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetDisableInnerDomain(v bool) *DescribeApiGroupResponseBody {
	s.DisableInnerDomain = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetGroupId(v string) *DescribeApiGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetGroupName(v string) *DescribeApiGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetHttpsPolicy(v string) *DescribeApiGroupResponseBody {
	s.HttpsPolicy = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetIllegalStatus(v string) *DescribeApiGroupResponseBody {
	s.IllegalStatus = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetInstanceId(v string) *DescribeApiGroupResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetInstanceType(v string) *DescribeApiGroupResponseBody {
	s.InstanceType = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetIpv6Status(v string) *DescribeApiGroupResponseBody {
	s.Ipv6Status = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetMigrationError(v string) *DescribeApiGroupResponseBody {
	s.MigrationError = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetMigrationStatus(v string) *DescribeApiGroupResponseBody {
	s.MigrationStatus = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetModifiedTime(v string) *DescribeApiGroupResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetPassthroughHeaders(v string) *DescribeApiGroupResponseBody {
	s.PassthroughHeaders = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetRegionId(v string) *DescribeApiGroupResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetRequestId(v string) *DescribeApiGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetStageItems(v *DescribeApiGroupResponseBodyStageItems) *DescribeApiGroupResponseBody {
	s.StageItems = v
	return s
}

func (s *DescribeApiGroupResponseBody) SetStatus(v string) *DescribeApiGroupResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetSubDomain(v string) *DescribeApiGroupResponseBody {
	s.SubDomain = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetTrafficLimit(v int32) *DescribeApiGroupResponseBody {
	s.TrafficLimit = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetUserLogConfig(v string) *DescribeApiGroupResponseBody {
	s.UserLogConfig = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetVpcDomain(v string) *DescribeApiGroupResponseBody {
	s.VpcDomain = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetVpcSlbIntranetDomain(v string) *DescribeApiGroupResponseBody {
	s.VpcSlbIntranetDomain = &v
	return s
}

func (s *DescribeApiGroupResponseBody) Validate() error {
	if s.CustomDomains != nil {
		if err := s.CustomDomains.Validate(); err != nil {
			return err
		}
	}
	if s.StageItems != nil {
		if err := s.StageItems.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApiGroupResponseBodyCustomDomains struct {
	DomainItem []*DescribeApiGroupResponseBodyCustomDomainsDomainItem `json:"DomainItem,omitempty" xml:"DomainItem,omitempty" type:"Repeated"`
}

func (s DescribeApiGroupResponseBodyCustomDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiGroupResponseBodyCustomDomains) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupResponseBodyCustomDomains) GetDomainItem() []*DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	return s.DomainItem
}

func (s *DescribeApiGroupResponseBodyCustomDomains) SetDomainItem(v []*DescribeApiGroupResponseBodyCustomDomainsDomainItem) *DescribeApiGroupResponseBodyCustomDomains {
	s.DomainItem = v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomains) Validate() error {
	if s.DomainItem != nil {
		for _, item := range s.DomainItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApiGroupResponseBodyCustomDomainsDomainItem struct {
	// The alias of the associated environment.
	//
	// example:
	//
	// TEST1
	BindStageAlias *string `json:"BindStageAlias,omitempty" xml:"BindStageAlias,omitempty"`
	// The environment in which the associated API group runs.
	//
	// example:
	//
	// TEST
	BindStageName *string `json:"BindStageName,omitempty" xml:"BindStageName,omitempty"`
	// The SSL certificate ID, which is automatically generated by the system.
	//
	// example:
	//
	// 6EF60BEC-0242-43AF-BB20-270359FB54A7
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The name of the SSL certificate.
	//
	// example:
	//
	// myCertificate
	CertificateName *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	// The time when the certificate expires.
	//
	// example:
	//
	// 2635123476000
	CertificateValidEnd *int64 `json:"CertificateValidEnd,omitempty" xml:"CertificateValidEnd,omitempty"`
	// The time when the certificate takes effect.
	//
	// example:
	//
	// 1689043476000
	CertificateValidStart    *int64 `json:"CertificateValidStart,omitempty" xml:"CertificateValidStart,omitempty"`
	ClientCertSDnPassThrough *bool  `json:"ClientCertSDnPassThrough,omitempty" xml:"ClientCertSDnPassThrough,omitempty"`
	// The type of the custom domain name.
	//
	// example:
	//
	// intranet
	CustomDomainType *string `json:"CustomDomainType,omitempty" xml:"CustomDomainType,omitempty"`
	// The binding status of the custom domain name. Valid values:
	//
	// 	- **BINDING**: The domain name is bound.
	//
	// 	- **BOUND**: The domain name is not bound.
	//
	// example:
	//
	// BINDING
	DomainBindingStatus *string `json:"DomainBindingStatus,omitempty" xml:"DomainBindingStatus,omitempty"`
	// The domain name resolution status. Valid values:
	//
	// 	- **RESOLVED**
	//
	// 	- **UNRESOLVED**
	//
	// example:
	//
	// RESOLVED
	DomainCNAMEStatus *string `json:"DomainCNAMEStatus,omitempty" xml:"DomainCNAMEStatus,omitempty"`
	// The validity status of the domain name. Valid values:
	//
	// 	- **NORMAL**: The domain name is valid.
	//
	// 	- **ABNORMAL**: The domain name is invalid. This status affects API calls and needs to be rectified as soon as possible.
	//
	// example:
	//
	// ABNORMAL
	DomainLegalStatus *string `json:"DomainLegalStatus,omitempty" xml:"DomainLegalStatus,omitempty"`
	// The domain name.
	//
	// example:
	//
	// api.demo.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Remarks about the domain name, such as the cause of an exception.
	//
	// example:
	//
	// Indicates that the domain name does not have an ICP filing.
	DomainRemark *string `json:"DomainRemark,omitempty" xml:"DomainRemark,omitempty"`
	// The status of the domain that uses the WebSocket feature.
	//
	// example:
	//
	// CLOSE
	DomainWebSocketStatus *string `json:"DomainWebSocketStatus,omitempty" xml:"DomainWebSocketStatus,omitempty"`
	// Indicates whether to redirect HTTP requests to HTTPS.
	//
	// example:
	//
	// false
	IsHttpRedirectToHttps *bool  `json:"IsHttpRedirectToHttps,omitempty" xml:"IsHttpRedirectToHttps,omitempty"`
	SslOcspCacheEnable    *bool  `json:"SslOcspCacheEnable,omitempty" xml:"SslOcspCacheEnable,omitempty"`
	SslOcspEnable         *bool  `json:"SslOcspEnable,omitempty" xml:"SslOcspEnable,omitempty"`
	SslVerifyDepth        *int32 `json:"SslVerifyDepth,omitempty" xml:"SslVerifyDepth,omitempty"`
	// The wildcard domain name mode.
	//
	// example:
	//
	// [\\"{test}.test.com\\"]
	WildcardDomainPatterns *string `json:"WildcardDomainPatterns,omitempty" xml:"WildcardDomainPatterns,omitempty"`
}

func (s DescribeApiGroupResponseBodyCustomDomainsDomainItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiGroupResponseBodyCustomDomainsDomainItem) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) GetBindStageAlias() *string {
	return s.BindStageAlias
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) GetBindStageName() *string {
	return s.BindStageName
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) GetCertificateId() *string {
	return s.CertificateId
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) GetCertificateName() *string {
	return s.CertificateName
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) GetCertificateValidEnd() *int64 {
	return s.CertificateValidEnd
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) GetCertificateValidStart() *int64 {
	return s.CertificateValidStart
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) GetClientCertSDnPassThrough() *bool {
	return s.ClientCertSDnPassThrough
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) GetCustomDomainType() *string {
	return s.CustomDomainType
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) GetDomainBindingStatus() *string {
	return s.DomainBindingStatus
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) GetDomainCNAMEStatus() *string {
	return s.DomainCNAMEStatus
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) GetDomainLegalStatus() *string {
	return s.DomainLegalStatus
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) GetDomainRemark() *string {
	return s.DomainRemark
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) GetDomainWebSocketStatus() *string {
	return s.DomainWebSocketStatus
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) GetIsHttpRedirectToHttps() *bool {
	return s.IsHttpRedirectToHttps
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) GetSslOcspCacheEnable() *bool {
	return s.SslOcspCacheEnable
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) GetSslOcspEnable() *bool {
	return s.SslOcspEnable
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) GetSslVerifyDepth() *int32 {
	return s.SslVerifyDepth
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) GetWildcardDomainPatterns() *string {
	return s.WildcardDomainPatterns
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetBindStageAlias(v string) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.BindStageAlias = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetBindStageName(v string) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.BindStageName = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetCertificateId(v string) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.CertificateId = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetCertificateName(v string) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.CertificateName = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetCertificateValidEnd(v int64) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.CertificateValidEnd = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetCertificateValidStart(v int64) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.CertificateValidStart = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetClientCertSDnPassThrough(v bool) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.ClientCertSDnPassThrough = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetCustomDomainType(v string) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.CustomDomainType = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetDomainBindingStatus(v string) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.DomainBindingStatus = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetDomainCNAMEStatus(v string) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.DomainCNAMEStatus = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetDomainLegalStatus(v string) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.DomainLegalStatus = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetDomainName(v string) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.DomainName = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetDomainRemark(v string) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.DomainRemark = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetDomainWebSocketStatus(v string) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.DomainWebSocketStatus = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetIsHttpRedirectToHttps(v bool) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.IsHttpRedirectToHttps = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetSslOcspCacheEnable(v bool) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.SslOcspCacheEnable = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetSslOcspEnable(v bool) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.SslOcspEnable = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetSslVerifyDepth(v int32) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.SslVerifyDepth = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetWildcardDomainPatterns(v string) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.WildcardDomainPatterns = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) Validate() error {
	return dara.Validate(s)
}

type DescribeApiGroupResponseBodyStageItems struct {
	StageInfo []*DescribeApiGroupResponseBodyStageItemsStageInfo `json:"StageInfo,omitempty" xml:"StageInfo,omitempty" type:"Repeated"`
}

func (s DescribeApiGroupResponseBodyStageItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiGroupResponseBodyStageItems) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupResponseBodyStageItems) GetStageInfo() []*DescribeApiGroupResponseBodyStageItemsStageInfo {
	return s.StageInfo
}

func (s *DescribeApiGroupResponseBodyStageItems) SetStageInfo(v []*DescribeApiGroupResponseBodyStageItemsStageInfo) *DescribeApiGroupResponseBodyStageItems {
	s.StageInfo = v
	return s
}

func (s *DescribeApiGroupResponseBodyStageItems) Validate() error {
	if s.StageInfo != nil {
		for _, item := range s.StageInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApiGroupResponseBodyStageItemsStageInfo struct {
	// The environment description.
	//
	// example:
	//
	// MYTEST
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// 123e8dc7bbe01613b5b1d726c2a7888e
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	// The environment name.
	//
	// example:
	//
	// TEST
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApiGroupResponseBodyStageItemsStageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiGroupResponseBodyStageItemsStageInfo) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupResponseBodyStageItemsStageInfo) GetDescription() *string {
	return s.Description
}

func (s *DescribeApiGroupResponseBodyStageItemsStageInfo) GetStageId() *string {
	return s.StageId
}

func (s *DescribeApiGroupResponseBodyStageItemsStageInfo) GetStageName() *string {
	return s.StageName
}

func (s *DescribeApiGroupResponseBodyStageItemsStageInfo) SetDescription(v string) *DescribeApiGroupResponseBodyStageItemsStageInfo {
	s.Description = &v
	return s
}

func (s *DescribeApiGroupResponseBodyStageItemsStageInfo) SetStageId(v string) *DescribeApiGroupResponseBodyStageItemsStageInfo {
	s.StageId = &v
	return s
}

func (s *DescribeApiGroupResponseBodyStageItemsStageInfo) SetStageName(v string) *DescribeApiGroupResponseBodyStageItemsStageInfo {
	s.StageName = &v
	return s
}

func (s *DescribeApiGroupResponseBodyStageItemsStageInfo) Validate() error {
	return dara.Validate(s)
}
