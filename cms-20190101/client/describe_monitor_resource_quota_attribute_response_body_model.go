// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorResourceQuotaAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeMonitorResourceQuotaAttributeResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeMonitorResourceQuotaAttributeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeMonitorResourceQuotaAttributeResponseBody
	GetRequestId() *string
	SetResourceQuota(v *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) *DescribeMonitorResourceQuotaAttributeResponseBody
	GetResourceQuota() *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota
}

type DescribeMonitorResourceQuotaAttributeResponseBody struct {
	// The status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 31BC7201-00F2-47B2-B7B9-6A173076ACE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details about the resource quotas of CloudMonitor.
	ResourceQuota *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota `json:"ResourceQuota,omitempty" xml:"ResourceQuota,omitempty" type:"Struct"`
}

func (s DescribeMonitorResourceQuotaAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBody) GetResourceQuota() *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	return s.ResourceQuota
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBody) SetCode(v string) *DescribeMonitorResourceQuotaAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBody) SetMessage(v string) *DescribeMonitorResourceQuotaAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBody) SetRequestId(v string) *DescribeMonitorResourceQuotaAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBody) SetResourceQuota(v *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) *DescribeMonitorResourceQuotaAttributeResponseBody {
	s.ResourceQuota = v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota struct {
	// The details about the quota of API calls.
	Api *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi `json:"Api,omitempty" xml:"Api,omitempty" type:"Struct"`
	// The details about the quota for custom monitoring.
	CustomMonitor *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor `json:"CustomMonitor,omitempty" xml:"CustomMonitor,omitempty" type:"Struct"`
	// The details about the quota of Hybrid Cloud Monitoring.
	EnterpriseQuota *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEnterpriseQuota `json:"EnterpriseQuota,omitempty" xml:"EnterpriseQuota,omitempty" type:"Struct"`
	// The details about the quota for event monitoring.
	EventMonitor *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor `json:"EventMonitor,omitempty" xml:"EventMonitor,omitempty" type:"Struct"`
	// The time when the resource plan expires.
	//
	// example:
	//
	// 2021-02-28
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the resource plan.
	//
	// example:
	//
	// cms_edition-cn-n6w20rn****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The details about the quota for log monitoring.
	LogMonitor *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor `json:"LogMonitor,omitempty" xml:"LogMonitor,omitempty" type:"Struct"`
	// The details about the quota of alert phone calls.
	Phone *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone `json:"Phone,omitempty" xml:"Phone,omitempty" type:"Struct"`
	// The details about the quota of alert text messages.
	SMS *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS `json:"SMS,omitempty" xml:"SMS,omitempty" type:"Struct"`
	// The quota of browser detection tasks.
	SiteMonitorBrowser *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorBrowser `json:"SiteMonitorBrowser,omitempty" xml:"SiteMonitorBrowser,omitempty" type:"Struct"`
	// The details about the quota of ECS detection points for site monitoring.
	SiteMonitorEcsProbe *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe `json:"SiteMonitorEcsProbe,omitempty" xml:"SiteMonitorEcsProbe,omitempty" type:"Struct"`
	// The quota of mobile detection tasks.
	SiteMonitorMobile *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorMobile `json:"SiteMonitorMobile,omitempty" xml:"SiteMonitorMobile,omitempty" type:"Struct"`
	// The details about the quota of carrier detection points for site monitoring.
	SiteMonitorOperatorProbe *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe `json:"SiteMonitorOperatorProbe,omitempty" xml:"SiteMonitorOperatorProbe,omitempty" type:"Struct"`
	// The quota of site monitoring tasks.
	SiteMonitorTask *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask `json:"SiteMonitorTask,omitempty" xml:"SiteMonitorTask,omitempty" type:"Struct"`
	// The current edition of CloudMonitor. Valid values:
	//
	// 	- free: Free Edition
	//
	// 	- pro: Pro Edition
	//
	// 	- cms_post: pay-as-you-go
	//
	// example:
	//
	// pro
	SuitInfo *string `json:"SuitInfo,omitempty" xml:"SuitInfo,omitempty"`
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) GetApi() *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi {
	return s.Api
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) GetCustomMonitor() *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor {
	return s.CustomMonitor
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) GetEnterpriseQuota() *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEnterpriseQuota {
	return s.EnterpriseQuota
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) GetEventMonitor() *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor {
	return s.EventMonitor
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) GetLogMonitor() *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor {
	return s.LogMonitor
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) GetPhone() *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone {
	return s.Phone
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) GetSMS() *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS {
	return s.SMS
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) GetSiteMonitorBrowser() *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorBrowser {
	return s.SiteMonitorBrowser
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) GetSiteMonitorEcsProbe() *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe {
	return s.SiteMonitorEcsProbe
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) GetSiteMonitorMobile() *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorMobile {
	return s.SiteMonitorMobile
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) GetSiteMonitorOperatorProbe() *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe {
	return s.SiteMonitorOperatorProbe
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) GetSiteMonitorTask() *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask {
	return s.SiteMonitorTask
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) GetSuitInfo() *string {
	return s.SuitInfo
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) SetApi(v *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	s.Api = v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) SetCustomMonitor(v *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	s.CustomMonitor = v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) SetEnterpriseQuota(v *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEnterpriseQuota) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	s.EnterpriseQuota = v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) SetEventMonitor(v *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	s.EventMonitor = v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) SetExpireTime(v string) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	s.ExpireTime = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) SetInstanceId(v string) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	s.InstanceId = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) SetLogMonitor(v *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	s.LogMonitor = v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) SetPhone(v *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	s.Phone = v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) SetSMS(v *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	s.SMS = v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) SetSiteMonitorBrowser(v *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorBrowser) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	s.SiteMonitorBrowser = v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) SetSiteMonitorEcsProbe(v *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	s.SiteMonitorEcsProbe = v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) SetSiteMonitorMobile(v *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorMobile) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	s.SiteMonitorMobile = v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) SetSiteMonitorOperatorProbe(v *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	s.SiteMonitorOperatorProbe = v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) SetSiteMonitorTask(v *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	s.SiteMonitorTask = v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) SetSuitInfo(v string) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	s.SuitInfo = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi struct {
	// The total quota of API calls. Unit: 10,000 calls.
	//
	// example:
	//
	// 500
	QuotaLimit *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	// The quota of API calls in your resource plan. Unit: 10,000 calls.
	//
	// example:
	//
	// 500
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	// The used quota of API calls in your resource plan. Unit: calls.
	//
	// example:
	//
	// 9987
	QuotaUsed *int32 `json:"QuotaUsed,omitempty" xml:"QuotaUsed,omitempty"`
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi) GetQuotaLimit() *int32 {
	return s.QuotaLimit
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi) GetQuotaPackage() *int32 {
	return s.QuotaPackage
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi) GetQuotaUsed() *int32 {
	return s.QuotaUsed
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi) SetQuotaLimit(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi {
	s.QuotaLimit = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi) SetQuotaPackage(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi {
	s.QuotaPackage = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi) SetQuotaUsed(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi {
	s.QuotaUsed = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor struct {
	// The total quota of the time series for custom monitoring.
	//
	// example:
	//
	// 1200
	QuotaLimit *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	// The quota of the time series for custom monitoring in your resource plan.
	//
	// example:
	//
	// 1000
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	// The used quota of the time series for custom monitoring in your resource plan.
	//
	// example:
	//
	// 8
	QuotaUsed *int32 `json:"QuotaUsed,omitempty" xml:"QuotaUsed,omitempty"`
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor) GetQuotaLimit() *int32 {
	return s.QuotaLimit
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor) GetQuotaPackage() *int32 {
	return s.QuotaPackage
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor) GetQuotaUsed() *int32 {
	return s.QuotaUsed
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor) SetQuotaLimit(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor {
	s.QuotaLimit = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor) SetQuotaPackage(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor {
	s.QuotaPackage = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor) SetQuotaUsed(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor {
	s.QuotaUsed = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEnterpriseQuota struct {
	// The ID of the instance monitored by Hybrid Cloud Monitoring.
	//
	// example:
	//
	// cms_enterprise_public_cn-7mz27pd****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The description of Hybrid Cloud Monitoring.
	//
	// example:
	//
	// ENTERPRISE
	SuitInfo *string `json:"SuitInfo,omitempty" xml:"SuitInfo,omitempty"`
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEnterpriseQuota) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEnterpriseQuota) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEnterpriseQuota) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEnterpriseQuota) GetSuitInfo() *string {
	return s.SuitInfo
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEnterpriseQuota) SetInstanceId(v string) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEnterpriseQuota {
	s.InstanceId = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEnterpriseQuota) SetSuitInfo(v string) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEnterpriseQuota {
	s.SuitInfo = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEnterpriseQuota) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor struct {
	// The total quota of events that can be reported in event monitoring. The total quota is the value that is multiplied by 10,000.
	//
	// example:
	//
	// 55
	QuotaLimit *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	// The quota of events that can be reported in event monitoring in your resource plan. The total quota is the value that is multiplied by 10,000.
	//
	// example:
	//
	// 50
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	// The used quota of events that can be reported in event monitoring in your resource plan. The total quota is the value that is multiplied by 10,000.
	//
	// example:
	//
	// 2
	QuotaUsed *int32 `json:"QuotaUsed,omitempty" xml:"QuotaUsed,omitempty"`
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor) GetQuotaLimit() *int32 {
	return s.QuotaLimit
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor) GetQuotaPackage() *int32 {
	return s.QuotaPackage
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor) GetQuotaUsed() *int32 {
	return s.QuotaUsed
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor) SetQuotaLimit(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor {
	s.QuotaLimit = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor) SetQuotaPackage(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor {
	s.QuotaPackage = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor) SetQuotaUsed(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor {
	s.QuotaUsed = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor struct {
	// The total quota of processed log data in log monitoring. Unit: MB/min.
	//
	// example:
	//
	// 150
	QuotaLimit *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	// The quota of processed log data in log monitoring in your resource plan. Unit: MB/min.
	//
	// example:
	//
	// 150
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	// The used quota of processed log data in log monitoring in your resource plan. Unit: MB/min.
	//
	// example:
	//
	// 80
	QuotaUsed *int32 `json:"QuotaUsed,omitempty" xml:"QuotaUsed,omitempty"`
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor) GetQuotaLimit() *int32 {
	return s.QuotaLimit
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor) GetQuotaPackage() *int32 {
	return s.QuotaPackage
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor) GetQuotaUsed() *int32 {
	return s.QuotaUsed
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor) SetQuotaLimit(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor {
	s.QuotaLimit = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor) SetQuotaPackage(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor {
	s.QuotaPackage = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor) SetQuotaUsed(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor {
	s.QuotaUsed = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone struct {
	// The total quota of alert phone calls. Unit: calls.
	//
	// example:
	//
	// 550
	QuotaLimit *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	// The quota of alert phone calls in your resource plan. Unit: calls.
	//
	// example:
	//
	// 500
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	// The used quota of alert phone calls in your resource plan. Unit: calls.
	//
	// example:
	//
	// 100
	QuotaUsed *int32 `json:"QuotaUsed,omitempty" xml:"QuotaUsed,omitempty"`
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone) GetQuotaLimit() *int32 {
	return s.QuotaLimit
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone) GetQuotaPackage() *int32 {
	return s.QuotaPackage
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone) GetQuotaUsed() *int32 {
	return s.QuotaUsed
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone) SetQuotaLimit(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone {
	s.QuotaLimit = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone) SetQuotaPackage(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone {
	s.QuotaPackage = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone) SetQuotaUsed(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone {
	s.QuotaUsed = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS struct {
	// The total quota of alert text messages. Unit: messages.
	//
	// example:
	//
	// 550
	QuotaLimit *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	// The quota of alert text messages in your resource plan. Unit: messages.
	//
	// example:
	//
	// 500
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	// The used quota of alert text messages in your resource plan. Unit: messages.
	//
	// example:
	//
	// 38
	QuotaUsed *int32 `json:"QuotaUsed,omitempty" xml:"QuotaUsed,omitempty"`
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS) GetQuotaLimit() *int32 {
	return s.QuotaLimit
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS) GetQuotaPackage() *int32 {
	return s.QuotaPackage
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS) GetQuotaUsed() *int32 {
	return s.QuotaUsed
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS) SetQuotaLimit(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS {
	s.QuotaLimit = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS) SetQuotaPackage(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS {
	s.QuotaPackage = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS) SetQuotaUsed(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS {
	s.QuotaUsed = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorBrowser struct {
	// The total quota of browser detection tasks.
	//
	// example:
	//
	// 50
	QuotaLimit *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	// The quota of browser detection tasks in your resource plan.
	//
	// example:
	//
	// 50
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	// The used quota of browser detection tasks in your resource plan.
	//
	// example:
	//
	// 15
	QuotaUsed *int32 `json:"QuotaUsed,omitempty" xml:"QuotaUsed,omitempty"`
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorBrowser) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorBrowser) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorBrowser) GetQuotaLimit() *int32 {
	return s.QuotaLimit
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorBrowser) GetQuotaPackage() *int32 {
	return s.QuotaPackage
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorBrowser) GetQuotaUsed() *int32 {
	return s.QuotaUsed
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorBrowser) SetQuotaLimit(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorBrowser {
	s.QuotaLimit = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorBrowser) SetQuotaPackage(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorBrowser {
	s.QuotaPackage = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorBrowser) SetQuotaUsed(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorBrowser {
	s.QuotaUsed = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorBrowser) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe struct {
	// The total quota of ECS detection points for site monitoring.
	//
	// > The value indicates the maximum number of ECS detection points that you can select for a site monitoring task.
	//
	// example:
	//
	// 5
	QuotaLimit *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	// The quota of ECS detection points for site monitoring in your resource plan.
	//
	// example:
	//
	// 5
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	// The used quota of ECS detection points for site monitoring in your resource plan.
	//
	// > The value indicates the total number of ECS detection points that are used by existing site monitoring tasks.
	//
	// example:
	//
	// 20
	QuotaUsed *int32 `json:"QuotaUsed,omitempty" xml:"QuotaUsed,omitempty"`
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe) GetQuotaLimit() *int32 {
	return s.QuotaLimit
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe) GetQuotaPackage() *int32 {
	return s.QuotaPackage
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe) GetQuotaUsed() *int32 {
	return s.QuotaUsed
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe) SetQuotaLimit(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe {
	s.QuotaLimit = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe) SetQuotaPackage(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe {
	s.QuotaPackage = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe) SetQuotaUsed(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe {
	s.QuotaUsed = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorMobile struct {
	// The total quota of mobile detection tasks.
	//
	// example:
	//
	// 50
	QuotaLimit *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	// The quota of mobile detection tasks in your resource plan.
	//
	// example:
	//
	// 50
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	// The used quota of mobile detection tasks in your resource plan.
	//
	// example:
	//
	// 15
	QuotaUsed *int32 `json:"QuotaUsed,omitempty" xml:"QuotaUsed,omitempty"`
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorMobile) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorMobile) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorMobile) GetQuotaLimit() *int32 {
	return s.QuotaLimit
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorMobile) GetQuotaPackage() *int32 {
	return s.QuotaPackage
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorMobile) GetQuotaUsed() *int32 {
	return s.QuotaUsed
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorMobile) SetQuotaLimit(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorMobile {
	s.QuotaLimit = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorMobile) SetQuotaPackage(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorMobile {
	s.QuotaPackage = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorMobile) SetQuotaUsed(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorMobile {
	s.QuotaUsed = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorMobile) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe struct {
	// The total quota of carrier detection points for site monitoring.
	//
	// example:
	//
	// 5
	QuotaLimit *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	// The quota of carrier detection points for site monitoring in your resource plan.
	//
	// example:
	//
	// 5
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	// The used quota of carrier detection points for site monitoring in your resource plan.
	//
	// example:
	//
	// 0
	QuotaUsed *int32 `json:"QuotaUsed,omitempty" xml:"QuotaUsed,omitempty"`
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe) GetQuotaLimit() *int32 {
	return s.QuotaLimit
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe) GetQuotaPackage() *int32 {
	return s.QuotaPackage
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe) GetQuotaUsed() *int32 {
	return s.QuotaUsed
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe) SetQuotaLimit(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe {
	s.QuotaLimit = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe) SetQuotaPackage(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe {
	s.QuotaPackage = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe) SetQuotaUsed(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe {
	s.QuotaUsed = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask struct {
	// The total quota of site monitoring tasks.
	//
	// example:
	//
	// 25
	QuotaLimit *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	// The quota of site monitoring tasks in your resource plan.
	//
	// example:
	//
	// 20
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	// The used quota of site monitoring tasks in your resource plan.
	//
	// example:
	//
	// 15
	QuotaUsed *int32 `json:"QuotaUsed,omitempty" xml:"QuotaUsed,omitempty"`
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask) GetQuotaLimit() *int32 {
	return s.QuotaLimit
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask) GetQuotaPackage() *int32 {
	return s.QuotaPackage
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask) GetQuotaUsed() *int32 {
	return s.QuotaUsed
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask) SetQuotaLimit(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask {
	s.QuotaLimit = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask) SetQuotaPackage(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask {
	s.QuotaPackage = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask) SetQuotaUsed(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask {
	s.QuotaUsed = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask) Validate() error {
	return dara.Validate(s)
}
