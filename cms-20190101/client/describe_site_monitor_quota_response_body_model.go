// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteMonitorQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSiteMonitorQuotaResponseBody
	GetCode() *string
	SetData(v *DescribeSiteMonitorQuotaResponseBodyData) *DescribeSiteMonitorQuotaResponseBody
	GetData() *DescribeSiteMonitorQuotaResponseBodyData
	SetMessage(v string) *DescribeSiteMonitorQuotaResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSiteMonitorQuotaResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeSiteMonitorQuotaResponseBody
	GetSuccess() *string
}

type DescribeSiteMonitorQuotaResponseBody struct {
	// The responses code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The quota.
	Data *DescribeSiteMonitorQuotaResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 26860260-76C6-404E-AB7A-EB98D36A6885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSiteMonitorQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorQuotaResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSiteMonitorQuotaResponseBody) GetData() *DescribeSiteMonitorQuotaResponseBodyData {
	return s.Data
}

func (s *DescribeSiteMonitorQuotaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSiteMonitorQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSiteMonitorQuotaResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSiteMonitorQuotaResponseBody) SetCode(v string) *DescribeSiteMonitorQuotaResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSiteMonitorQuotaResponseBody) SetData(v *DescribeSiteMonitorQuotaResponseBodyData) *DescribeSiteMonitorQuotaResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSiteMonitorQuotaResponseBody) SetMessage(v string) *DescribeSiteMonitorQuotaResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSiteMonitorQuotaResponseBody) SetRequestId(v string) *DescribeSiteMonitorQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSiteMonitorQuotaResponseBody) SetSuccess(v string) *DescribeSiteMonitorQuotaResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSiteMonitorQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSiteMonitorQuotaResponseBodyData struct {
	// Indicates whether second-level monitoring is enabled. Valid values:
	//
	// 	- true: Second-level monitoring is enabled.
	//
	// 	- false: Second-level monitoring is disabled.
	//
	// example:
	//
	// false
	SecondMonitor *bool `json:"SecondMonitor,omitempty" xml:"SecondMonitor,omitempty"`
	// The quota of detection points that are provided by Alibaba Cloud. Five detection points are provided for free.
	//
	// example:
	//
	// 5
	SiteMonitorIdcQuota *int32 `json:"SiteMonitorIdcQuota,omitempty" xml:"SiteMonitorIdcQuota,omitempty"`
	// The quota of detection points that are not provided by Alibaba Cloud. Default value: 0.
	//
	// example:
	//
	// 0
	SiteMonitorOperatorQuotaQuota *int32 `json:"SiteMonitorOperatorQuotaQuota,omitempty" xml:"SiteMonitorOperatorQuotaQuota,omitempty"`
	// The used quota of site monitoring tasks.
	//
	// example:
	//
	// 6
	SiteMonitorQuotaTaskUsed *int32 `json:"SiteMonitorQuotaTaskUsed,omitempty" xml:"SiteMonitorQuotaTaskUsed,omitempty"`
	// The quota of site monitoring tasks.
	//
	// example:
	//
	// 10
	SiteMonitorTaskQuota *int32 `json:"SiteMonitorTaskQuota,omitempty" xml:"SiteMonitorTaskQuota,omitempty"`
	// The version of site monitoring. Valid values:
	//
	// 	- V1
	//
	// 	- V2
	//
	// example:
	//
	// V1
	SiteMonitorVersion *string `json:"SiteMonitorVersion,omitempty" xml:"SiteMonitorVersion,omitempty"`
}

func (s DescribeSiteMonitorQuotaResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorQuotaResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorQuotaResponseBodyData) GetSecondMonitor() *bool {
	return s.SecondMonitor
}

func (s *DescribeSiteMonitorQuotaResponseBodyData) GetSiteMonitorIdcQuota() *int32 {
	return s.SiteMonitorIdcQuota
}

func (s *DescribeSiteMonitorQuotaResponseBodyData) GetSiteMonitorOperatorQuotaQuota() *int32 {
	return s.SiteMonitorOperatorQuotaQuota
}

func (s *DescribeSiteMonitorQuotaResponseBodyData) GetSiteMonitorQuotaTaskUsed() *int32 {
	return s.SiteMonitorQuotaTaskUsed
}

func (s *DescribeSiteMonitorQuotaResponseBodyData) GetSiteMonitorTaskQuota() *int32 {
	return s.SiteMonitorTaskQuota
}

func (s *DescribeSiteMonitorQuotaResponseBodyData) GetSiteMonitorVersion() *string {
	return s.SiteMonitorVersion
}

func (s *DescribeSiteMonitorQuotaResponseBodyData) SetSecondMonitor(v bool) *DescribeSiteMonitorQuotaResponseBodyData {
	s.SecondMonitor = &v
	return s
}

func (s *DescribeSiteMonitorQuotaResponseBodyData) SetSiteMonitorIdcQuota(v int32) *DescribeSiteMonitorQuotaResponseBodyData {
	s.SiteMonitorIdcQuota = &v
	return s
}

func (s *DescribeSiteMonitorQuotaResponseBodyData) SetSiteMonitorOperatorQuotaQuota(v int32) *DescribeSiteMonitorQuotaResponseBodyData {
	s.SiteMonitorOperatorQuotaQuota = &v
	return s
}

func (s *DescribeSiteMonitorQuotaResponseBodyData) SetSiteMonitorQuotaTaskUsed(v int32) *DescribeSiteMonitorQuotaResponseBodyData {
	s.SiteMonitorQuotaTaskUsed = &v
	return s
}

func (s *DescribeSiteMonitorQuotaResponseBodyData) SetSiteMonitorTaskQuota(v int32) *DescribeSiteMonitorQuotaResponseBodyData {
	s.SiteMonitorTaskQuota = &v
	return s
}

func (s *DescribeSiteMonitorQuotaResponseBodyData) SetSiteMonitorVersion(v string) *DescribeSiteMonitorQuotaResponseBodyData {
	s.SiteMonitorVersion = &v
	return s
}

func (s *DescribeSiteMonitorQuotaResponseBodyData) Validate() error {
	return dara.Validate(s)
}
