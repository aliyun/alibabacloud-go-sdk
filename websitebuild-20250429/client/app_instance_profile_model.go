// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppInstanceProfile interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationType(v string) *AppInstanceProfile
	GetApplicationType() *string
	SetApplicationTypeText(v string) *AppInstanceProfile
	GetApplicationTypeText() *string
	SetBizId(v string) *AppInstanceProfile
	GetBizId() *string
	SetCaseAuditStatus(v string) *AppInstanceProfile
	GetCaseAuditStatus() *string
	SetCaseShelfStatus(v string) *AppInstanceProfile
	GetCaseShelfStatus() *string
	SetChatbiTaskId(v string) *AppInstanceProfile
	GetChatbiTaskId() *string
	SetChatbiTaskStatus(v string) *AppInstanceProfile
	GetChatbiTaskStatus() *string
	SetCommodityCode(v string) *AppInstanceProfile
	GetCommodityCode() *string
	SetCustomerService(v string) *AppInstanceProfile
	GetCustomerService() *string
	SetDeployArea(v string) *AppInstanceProfile
	GetDeployArea() *string
	SetIcpbaNo(v string) *AppInstanceProfile
	GetIcpbaNo() *string
	SetIndependentDeployment(v *AppInstanceProfileIndependentDeployment) *AppInstanceProfile
	GetIndependentDeployment() *AppInstanceProfileIndependentDeployment
	SetInstanceId(v string) *AppInstanceProfile
	GetInstanceId() *string
	SetOpenChatBi(v bool) *AppInstanceProfile
	GetOpenChatBi() *bool
	SetOrdTime(v string) *AppInstanceProfile
	GetOrdTime() *string
	SetOrderId(v string) *AppInstanceProfile
	GetOrderId() *string
	SetPayTime(v string) *AppInstanceProfile
	GetPayTime() *string
	SetPreviewUrl(v string) *AppInstanceProfile
	GetPreviewUrl() *string
	SetSeoSite(v string) *AppInstanceProfile
	GetSeoSite() *string
	SetShareRenderWatermark(v string) *AppInstanceProfile
	GetShareRenderWatermark() *string
	SetSiteVersion(v string) *AppInstanceProfile
	GetSiteVersion() *string
	SetSiteVersionText(v string) *AppInstanceProfile
	GetSiteVersionText() *string
	SetSource(v string) *AppInstanceProfile
	GetSource() *string
	SetTemplateEtag(v string) *AppInstanceProfile
	GetTemplateEtag() *string
	SetTemplateId(v string) *AppInstanceProfile
	GetTemplateId() *string
}

type AppInstanceProfile struct {
	// The application type code. For example, PC_WebSite indicates a PC website.
	//
	// example:
	//
	// PC_WebSite
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The display name of the application type.
	//
	// example:
	//
	// PC Website
	ApplicationTypeText *string `json:"ApplicationTypeText,omitempty" xml:"ApplicationTypeText,omitempty"`
	// The business identifier of the application instance.
	//
	// example:
	//
	// WS_SITE_1
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The audit status.
	//
	// example:
	//
	// 1
	CaseAuditStatus *string `json:"CaseAuditStatus,omitempty" xml:"CaseAuditStatus,omitempty"`
	// The listing status.
	//
	// example:
	//
	// online
	CaseShelfStatus *string `json:"CaseShelfStatus,omitempty" xml:"CaseShelfStatus,omitempty"`
	// The ID of the ChatBI BuildSchemaIndex asynchronous task.
	//
	// example:
	//
	// task-1
	ChatbiTaskId *string `json:"ChatbiTaskId,omitempty" xml:"ChatbiTaskId,omitempty"`
	// The status of the ChatBI BuildSchemaIndex task. Valid values: running, finish, and failed.
	//
	// example:
	//
	// finish
	ChatbiTaskStatus *string `json:"ChatbiTaskStatus,omitempty" xml:"ChatbiTaskStatus,omitempty"`
	// The commodity code used when placing the order.
	//
	// example:
	//
	// domain_website_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The customer service contact information or description.
	//
	// example:
	//
	// support@example.com
	CustomerService *string `json:"CustomerService,omitempty" xml:"CustomerService,omitempty"`
	// The deployment region code of the application.
	//
	// example:
	//
	// HongKong
	DeployArea *string `json:"DeployArea,omitempty" xml:"DeployArea,omitempty"`
	// The ICP filing number bound to the custom domain name. Multiple filing numbers are separated by commas (,).
	//
	// example:
	//
	// 京ICP备12345678号-1
	IcpbaNo *string `json:"IcpbaNo,omitempty" xml:"IcpbaNo,omitempty"`
	// The dedicated server information. This field is computed in real time by DescribeAppInstance and is not populated in list API responses.
	IndependentDeployment *AppInstanceProfileIndependentDeployment `json:"IndependentDeployment,omitempty" xml:"IndependentDeployment,omitempty" type:"Struct"`
	// The ID of the associated Lingxiao instance.
	//
	// example:
	//
	// lx-1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the current user is included in the ChatBI feature canary release whitelist.
	OpenChatBi *bool `json:"OpenChatBi,omitempty" xml:"OpenChatBi,omitempty"`
	// The time when the order was placed, in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2026-09-17 10:30:00
	OrdTime *string `json:"OrdTime,omitempty" xml:"OrdTime,omitempty"`
	// The order ID.
	//
	// example:
	//
	// order-1
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The time when the payment was made, in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2026-09-17 10:35:00
	PayTime *string `json:"PayTime,omitempty" xml:"PayTime,omitempty"`
	// The URL used to access the application preview page.
	//
	// example:
	//
	// https://example.com/preview
	PreviewUrl *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	// The SEO site configuration identifier.
	//
	// example:
	//
	// seo-site-1
	SeoSite *string `json:"SeoSite,omitempty" xml:"SeoSite,omitempty"`
	// The switch for the share rendering watermark. A value of 1 indicates enabled, and a value of 0 indicates disabled. This field is empty if not configured.
	//
	// example:
	//
	// 1
	ShareRenderWatermark *string `json:"ShareRenderWatermark,omitempty" xml:"ShareRenderWatermark,omitempty"`
	// The site version code.
	//
	// example:
	//
	// Max_Edition
	SiteVersion *string `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// The display name of the site version, including the deployment region.
	//
	// example:
	//
	// Flagship Edition [Hong Kong]
	SiteVersionText *string `json:"SiteVersionText,omitempty" xml:"SiteVersionText,omitempty"`
	// The identifier of the creation or synchronization source.
	//
	// example:
	//
	// MARKET_CLOUD_DREAM
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ETag identifier of the current application template.
	//
	// example:
	//
	// etag-1
	TemplateEtag *string `json:"TemplateEtag,omitempty" xml:"TemplateEtag,omitempty"`
	// The template ID used by the current application.
	//
	// example:
	//
	// tpl-1
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s AppInstanceProfile) String() string {
	return dara.Prettify(s)
}

func (s AppInstanceProfile) GoString() string {
	return s.String()
}

func (s *AppInstanceProfile) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *AppInstanceProfile) GetApplicationTypeText() *string {
	return s.ApplicationTypeText
}

func (s *AppInstanceProfile) GetBizId() *string {
	return s.BizId
}

func (s *AppInstanceProfile) GetCaseAuditStatus() *string {
	return s.CaseAuditStatus
}

func (s *AppInstanceProfile) GetCaseShelfStatus() *string {
	return s.CaseShelfStatus
}

func (s *AppInstanceProfile) GetChatbiTaskId() *string {
	return s.ChatbiTaskId
}

func (s *AppInstanceProfile) GetChatbiTaskStatus() *string {
	return s.ChatbiTaskStatus
}

func (s *AppInstanceProfile) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *AppInstanceProfile) GetCustomerService() *string {
	return s.CustomerService
}

func (s *AppInstanceProfile) GetDeployArea() *string {
	return s.DeployArea
}

func (s *AppInstanceProfile) GetIcpbaNo() *string {
	return s.IcpbaNo
}

func (s *AppInstanceProfile) GetIndependentDeployment() *AppInstanceProfileIndependentDeployment {
	return s.IndependentDeployment
}

func (s *AppInstanceProfile) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AppInstanceProfile) GetOpenChatBi() *bool {
	return s.OpenChatBi
}

func (s *AppInstanceProfile) GetOrdTime() *string {
	return s.OrdTime
}

func (s *AppInstanceProfile) GetOrderId() *string {
	return s.OrderId
}

func (s *AppInstanceProfile) GetPayTime() *string {
	return s.PayTime
}

func (s *AppInstanceProfile) GetPreviewUrl() *string {
	return s.PreviewUrl
}

func (s *AppInstanceProfile) GetSeoSite() *string {
	return s.SeoSite
}

func (s *AppInstanceProfile) GetShareRenderWatermark() *string {
	return s.ShareRenderWatermark
}

func (s *AppInstanceProfile) GetSiteVersion() *string {
	return s.SiteVersion
}

func (s *AppInstanceProfile) GetSiteVersionText() *string {
	return s.SiteVersionText
}

func (s *AppInstanceProfile) GetSource() *string {
	return s.Source
}

func (s *AppInstanceProfile) GetTemplateEtag() *string {
	return s.TemplateEtag
}

func (s *AppInstanceProfile) GetTemplateId() *string {
	return s.TemplateId
}

func (s *AppInstanceProfile) SetApplicationType(v string) *AppInstanceProfile {
	s.ApplicationType = &v
	return s
}

func (s *AppInstanceProfile) SetApplicationTypeText(v string) *AppInstanceProfile {
	s.ApplicationTypeText = &v
	return s
}

func (s *AppInstanceProfile) SetBizId(v string) *AppInstanceProfile {
	s.BizId = &v
	return s
}

func (s *AppInstanceProfile) SetCaseAuditStatus(v string) *AppInstanceProfile {
	s.CaseAuditStatus = &v
	return s
}

func (s *AppInstanceProfile) SetCaseShelfStatus(v string) *AppInstanceProfile {
	s.CaseShelfStatus = &v
	return s
}

func (s *AppInstanceProfile) SetChatbiTaskId(v string) *AppInstanceProfile {
	s.ChatbiTaskId = &v
	return s
}

func (s *AppInstanceProfile) SetChatbiTaskStatus(v string) *AppInstanceProfile {
	s.ChatbiTaskStatus = &v
	return s
}

func (s *AppInstanceProfile) SetCommodityCode(v string) *AppInstanceProfile {
	s.CommodityCode = &v
	return s
}

func (s *AppInstanceProfile) SetCustomerService(v string) *AppInstanceProfile {
	s.CustomerService = &v
	return s
}

func (s *AppInstanceProfile) SetDeployArea(v string) *AppInstanceProfile {
	s.DeployArea = &v
	return s
}

func (s *AppInstanceProfile) SetIcpbaNo(v string) *AppInstanceProfile {
	s.IcpbaNo = &v
	return s
}

func (s *AppInstanceProfile) SetIndependentDeployment(v *AppInstanceProfileIndependentDeployment) *AppInstanceProfile {
	s.IndependentDeployment = v
	return s
}

func (s *AppInstanceProfile) SetInstanceId(v string) *AppInstanceProfile {
	s.InstanceId = &v
	return s
}

func (s *AppInstanceProfile) SetOpenChatBi(v bool) *AppInstanceProfile {
	s.OpenChatBi = &v
	return s
}

func (s *AppInstanceProfile) SetOrdTime(v string) *AppInstanceProfile {
	s.OrdTime = &v
	return s
}

func (s *AppInstanceProfile) SetOrderId(v string) *AppInstanceProfile {
	s.OrderId = &v
	return s
}

func (s *AppInstanceProfile) SetPayTime(v string) *AppInstanceProfile {
	s.PayTime = &v
	return s
}

func (s *AppInstanceProfile) SetPreviewUrl(v string) *AppInstanceProfile {
	s.PreviewUrl = &v
	return s
}

func (s *AppInstanceProfile) SetSeoSite(v string) *AppInstanceProfile {
	s.SeoSite = &v
	return s
}

func (s *AppInstanceProfile) SetShareRenderWatermark(v string) *AppInstanceProfile {
	s.ShareRenderWatermark = &v
	return s
}

func (s *AppInstanceProfile) SetSiteVersion(v string) *AppInstanceProfile {
	s.SiteVersion = &v
	return s
}

func (s *AppInstanceProfile) SetSiteVersionText(v string) *AppInstanceProfile {
	s.SiteVersionText = &v
	return s
}

func (s *AppInstanceProfile) SetSource(v string) *AppInstanceProfile {
	s.Source = &v
	return s
}

func (s *AppInstanceProfile) SetTemplateEtag(v string) *AppInstanceProfile {
	s.TemplateEtag = &v
	return s
}

func (s *AppInstanceProfile) SetTemplateId(v string) *AppInstanceProfile {
	s.TemplateId = &v
	return s
}

func (s *AppInstanceProfile) Validate() error {
	if s.IndependentDeployment != nil {
		if err := s.IndependentDeployment.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AppInstanceProfileIndependentDeployment struct {
	// Specifies whether a retry is allowed after the dedicated server fails to start.
	CanRetry *bool `json:"CanRetry,omitempty" xml:"CanRetry,omitempty"`
	// The deployment region code of the dedicated server.
	//
	// example:
	//
	// HongKong
	DeployArea *string `json:"DeployArea,omitempty" xml:"DeployArea,omitempty"`
	// The DNS record type used when binding a domain name to the dedicated IP address.
	//
	// example:
	//
	// A
	DnsRecordType *string `json:"DnsRecordType,omitempty" xml:"DnsRecordType,omitempty"`
	// The eligibility status of the dedicated server. Valid values:
	//
	// - UPGRADE_REQUIRED: A version upgrade is required.
	//
	// - AVAILABLE: The dedicated server can be started.
	//
	// - UNAVAILABLE: The dedicated server is currently unavailable.
	//
	// example:
	//
	// AVAILABLE
	Eligibility *string `json:"Eligibility,omitempty" xml:"Eligibility,omitempty"`
	// The status of the dedicated server. Valid values:
	//
	// - NOT_ENABLED: Not started.
	//
	// - ENABLING: Starting.
	//
	// - ACTIVE: Started.
	//
	// - FAILED: Failed to start.
	//
	// example:
	//
	// ACTIVE
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The dedicated IP address assigned after the dedicated server is started.
	//
	// example:
	//
	// 203.0.113.10
	IndependentIp *string `json:"IndependentIp,omitempty" xml:"IndependentIp,omitempty"`
	// The display specification of the dedicated server. This does not represent the real-time resource usage of the container.
	//
	// example:
	//
	// 2 vCPUs 2 GB
	ServerSpec *string `json:"ServerSpec,omitempty" xml:"ServerSpec,omitempty"`
	// The task ID for starting the dedicated server.
	//
	// example:
	//
	// 17
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AppInstanceProfileIndependentDeployment) String() string {
	return dara.Prettify(s)
}

func (s AppInstanceProfileIndependentDeployment) GoString() string {
	return s.String()
}

func (s *AppInstanceProfileIndependentDeployment) GetCanRetry() *bool {
	return s.CanRetry
}

func (s *AppInstanceProfileIndependentDeployment) GetDeployArea() *string {
	return s.DeployArea
}

func (s *AppInstanceProfileIndependentDeployment) GetDnsRecordType() *string {
	return s.DnsRecordType
}

func (s *AppInstanceProfileIndependentDeployment) GetEligibility() *string {
	return s.Eligibility
}

func (s *AppInstanceProfileIndependentDeployment) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *AppInstanceProfileIndependentDeployment) GetIndependentIp() *string {
	return s.IndependentIp
}

func (s *AppInstanceProfileIndependentDeployment) GetServerSpec() *string {
	return s.ServerSpec
}

func (s *AppInstanceProfileIndependentDeployment) GetTaskId() *int64 {
	return s.TaskId
}

func (s *AppInstanceProfileIndependentDeployment) SetCanRetry(v bool) *AppInstanceProfileIndependentDeployment {
	s.CanRetry = &v
	return s
}

func (s *AppInstanceProfileIndependentDeployment) SetDeployArea(v string) *AppInstanceProfileIndependentDeployment {
	s.DeployArea = &v
	return s
}

func (s *AppInstanceProfileIndependentDeployment) SetDnsRecordType(v string) *AppInstanceProfileIndependentDeployment {
	s.DnsRecordType = &v
	return s
}

func (s *AppInstanceProfileIndependentDeployment) SetEligibility(v string) *AppInstanceProfileIndependentDeployment {
	s.Eligibility = &v
	return s
}

func (s *AppInstanceProfileIndependentDeployment) SetEnableStatus(v string) *AppInstanceProfileIndependentDeployment {
	s.EnableStatus = &v
	return s
}

func (s *AppInstanceProfileIndependentDeployment) SetIndependentIp(v string) *AppInstanceProfileIndependentDeployment {
	s.IndependentIp = &v
	return s
}

func (s *AppInstanceProfileIndependentDeployment) SetServerSpec(v string) *AppInstanceProfileIndependentDeployment {
	s.ServerSpec = &v
	return s
}

func (s *AppInstanceProfileIndependentDeployment) SetTaskId(v int64) *AppInstanceProfileIndependentDeployment {
	s.TaskId = &v
	return s
}

func (s *AppInstanceProfileIndependentDeployment) Validate() error {
	return dara.Validate(s)
}
