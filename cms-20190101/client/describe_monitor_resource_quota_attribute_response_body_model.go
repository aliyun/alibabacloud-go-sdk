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
	// > 200 indicates success.
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
	// The quota details of CloudMonitor.
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
	if s.ResourceQuota != nil {
		if err := s.ResourceQuota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota struct {
	// The API quota.
	Api *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi `json:"Api,omitempty" xml:"Api,omitempty" type:"Struct"`
	// The quota of custom monitoring.
	CustomMonitor *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor `json:"CustomMonitor,omitempty" xml:"CustomMonitor,omitempty" type:"Struct"`
	// The quota details of Hybrid Cloud Monitoring.
	EnterpriseQuota *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEnterpriseQuota `json:"EnterpriseQuota,omitempty" xml:"EnterpriseQuota,omitempty" type:"Struct"`
	// The quota of event monitoring.
	EventMonitor *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor `json:"EventMonitor,omitempty" xml:"EventMonitor,omitempty" type:"Struct"`
	// The expiration time of the plan. The time follows the ISO 8601 standard in the YYYY-MM-DD format. The default time zone is UTC+8.
	//
	// example:
	//
	// 2021-02-28
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The subscription plan ID.
	//
	// example:
	//
	// cms_edition-cn-n6w20rn****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The quota of Log Monitoring.
	LogMonitor *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor `json:"LogMonitor,omitempty" xml:"LogMonitor,omitempty" type:"Struct"`
	// The quota of alert phone call packages.
	Phone *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone `json:"Phone,omitempty" xml:"Phone,omitempty" type:"Struct"`
	// The quota of alert text message packages.
	SMS *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS `json:"SMS,omitempty" xml:"SMS,omitempty" type:"Struct"`
	// The quota of browser-based monitoring tasks.
	SiteMonitorBrowser *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorBrowser `json:"SiteMonitorBrowser,omitempty" xml:"SiteMonitorBrowser,omitempty" type:"Struct"`
	// The quota of ECS detection points for site monitoring.
	SiteMonitorEcsProbe *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe `json:"SiteMonitorEcsProbe,omitempty" xml:"SiteMonitorEcsProbe,omitempty" type:"Struct"`
	// The quota of mobile monitoring tasks.
	SiteMonitorMobile *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorMobile `json:"SiteMonitorMobile,omitempty" xml:"SiteMonitorMobile,omitempty" type:"Struct"`
	// The quota of carrier detection points for site monitoring.
	SiteMonitorOperatorProbe *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe `json:"SiteMonitorOperatorProbe,omitempty" xml:"SiteMonitorOperatorProbe,omitempty" type:"Struct"`
	// The quota of site monitoring tasks.
	SiteMonitorTask *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask `json:"SiteMonitorTask,omitempty" xml:"SiteMonitorTask,omitempty" type:"Struct"`
	// The current edition of CloudMonitor. Valid values:
	//
	// - free: Free Edition.
	//
	// - pro: Pro Edition.
	//
	// - cms_post: Pay-as-you-go.
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
	if s.Api != nil {
		if err := s.Api.Validate(); err != nil {
			return err
		}
	}
	if s.CustomMonitor != nil {
		if err := s.CustomMonitor.Validate(); err != nil {
			return err
		}
	}
	if s.EnterpriseQuota != nil {
		if err := s.EnterpriseQuota.Validate(); err != nil {
			return err
		}
	}
	if s.EventMonitor != nil {
		if err := s.EventMonitor.Validate(); err != nil {
			return err
		}
	}
	if s.LogMonitor != nil {
		if err := s.LogMonitor.Validate(); err != nil {
			return err
		}
	}
	if s.Phone != nil {
		if err := s.Phone.Validate(); err != nil {
			return err
		}
	}
	if s.SMS != nil {
		if err := s.SMS.Validate(); err != nil {
			return err
		}
	}
	if s.SiteMonitorBrowser != nil {
		if err := s.SiteMonitorBrowser.Validate(); err != nil {
			return err
		}
	}
	if s.SiteMonitorEcsProbe != nil {
		if err := s.SiteMonitorEcsProbe.Validate(); err != nil {
			return err
		}
	}
	if s.SiteMonitorMobile != nil {
		if err := s.SiteMonitorMobile.Validate(); err != nil {
			return err
		}
	}
	if s.SiteMonitorOperatorProbe != nil {
		if err := s.SiteMonitorOperatorProbe.Validate(); err != nil {
			return err
		}
	}
	if s.SiteMonitorTask != nil {
		if err := s.SiteMonitorTask.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi struct {
	// The total API quota. Unit: 10,000 calls.
	//
	// example:
	//
	// 500
	QuotaLimit *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	// The API quota included in the plan. Unit: 10,000 calls.
	//
	// example:
	//
	// 500
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	// The used API quota in the plan. Unit: calls.
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
	// The total quota of custom monitoring time series. Unit: time series.
	//
	// example:
	//
	// 1200
	QuotaLimit *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	// The quota of custom monitoring time series included in the plan. Unit: time series.
	//
	// example:
	//
	// 1000
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	// The used quota of custom monitoring time series in the plan. Unit: time series.
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
	// The instance ID of Hybrid Cloud Monitoring.
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
	// The total quota of event monitoring. Unit: 10,000 events.
	//
	// example:
	//
	// 55
	QuotaLimit *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	// The quota of event monitoring included in the plan. Unit: 10,000 events.
	//
	// example:
	//
	// 50
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	// The used quota of event monitoring in the plan. Unit: 10,000 events.
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
	// The total quota of Log Monitoring. Unit: MByte/Min.
	//
	// example:
	//
	// 150
	QuotaLimit *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	// The quota of Log Monitoring included in the plan. Unit: MByte/Min.
	//
	// example:
	//
	// 150
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	// The used quota of Log Monitoring in the plan. Unit: MByte/Min.
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
	// The quota of alert phone call packages. Unit: calls.
	//
	// example:
	//
	// 500
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	// The used quota of alert phone call packages. Unit: calls.
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
	// The quota of alert text message packages. Unit: messages.
	//
	// example:
	//
	// 500
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	// The used quota of alert text message packages. Unit: messages.
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
	// The total quota of browser-based monitoring tasks. Unit: tasks.
	//
	// example:
	//
	// 50
	QuotaLimit *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	// The quota of browser-based monitoring tasks included in the plan. Unit: tasks.
	//
	// example:
	//
	// 50
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	// The used quota of browser-based monitoring tasks in the plan. Unit: tasks.
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
	// > The maximum number of ECS detection points that can be selected for a single site monitoring task.
	//
	// example:
	//
	// 5
	QuotaLimit *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	// The quota of ECS detection points for site monitoring included in the plan.
	//
	// example:
	//
	// 5
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	// The used quota of ECS detection points for site monitoring in the plan.
	//
	// > The total number of ECS detection points across all site monitoring tasks that have been created.
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
	// The total quota of mobile monitoring tasks. Unit: tasks.
	//
	// example:
	//
	// 50
	QuotaLimit *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	// The quota of mobile monitoring tasks included in the plan. Unit: tasks.
	//
	// example:
	//
	// 50
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	// The used quota of mobile monitoring tasks in the plan. Unit: tasks.
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
	// The quota of carrier detection points for site monitoring included in the plan.
	//
	// example:
	//
	// 5
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	// The used quota of carrier detection points for site monitoring in the plan.
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
	// The total quota of site monitoring tasks. Unit: tasks.
	//
	// example:
	//
	// 25
	QuotaLimit *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	// The quota of site monitoring tasks included in the plan. Unit: tasks.
	//
	// example:
	//
	// 20
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	// The used quota of site monitoring tasks in the plan. Unit: tasks.
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
