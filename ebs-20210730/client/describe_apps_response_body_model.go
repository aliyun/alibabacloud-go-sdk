// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeAppsResponseBody
	GetAccessDeniedDetail() *string
	SetAppInfos(v []*DescribeAppsResponseBodyAppInfos) *DescribeAppsResponseBody
	GetAppInfos() []*DescribeAppsResponseBodyAppInfos
	SetCode(v string) *DescribeAppsResponseBody
	GetCode() *string
	SetDynamicCode(v string) *DescribeAppsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeAppsResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *DescribeAppsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeAppsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAppsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAppsResponseBody
	GetSuccess() *bool
	SetUserCode(v string) *DescribeAppsResponseBody
	GetUserCode() *string
}

type DescribeAppsResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string                             `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	AppInfos           []*DescribeAppsResponseBodyAppInfos `json:"AppInfos,omitempty" xml:"AppInfos,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// OK
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// -
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A1C00637-AC84-5EFD-89B5-D5CE39F0F2B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// OK
	UserCode *string `json:"UserCode,omitempty" xml:"UserCode,omitempty"`
}

func (s DescribeAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeAppsResponseBody) GetAppInfos() []*DescribeAppsResponseBodyAppInfos {
	return s.AppInfos
}

func (s *DescribeAppsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeAppsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeAppsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeAppsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeAppsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAppsResponseBody) GetUserCode() *string {
	return s.UserCode
}

func (s *DescribeAppsResponseBody) SetAccessDeniedDetail(v string) *DescribeAppsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeAppsResponseBody) SetAppInfos(v []*DescribeAppsResponseBodyAppInfos) *DescribeAppsResponseBody {
	s.AppInfos = v
	return s
}

func (s *DescribeAppsResponseBody) SetCode(v string) *DescribeAppsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAppsResponseBody) SetDynamicCode(v string) *DescribeAppsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeAppsResponseBody) SetDynamicMessage(v string) *DescribeAppsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeAppsResponseBody) SetHttpStatusCode(v int32) *DescribeAppsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAppsResponseBody) SetMessage(v string) *DescribeAppsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAppsResponseBody) SetRequestId(v string) *DescribeAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppsResponseBody) SetSuccess(v bool) *DescribeAppsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAppsResponseBody) SetUserCode(v string) *DescribeAppsResponseBody {
	s.UserCode = &v
	return s
}

func (s *DescribeAppsResponseBody) Validate() error {
	if s.AppInfos != nil {
		for _, item := range s.AppInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAppsResponseBodyAppInfos struct {
	// example:
	//
	// app-bd5e3533
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// app1
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// default
	AppTags []*DescribeAppsResponseBodyAppInfosAppTags `json:"AppTags,omitempty" xml:"AppTags,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Default                *bool   `json:"Default,omitempty" xml:"Default,omitempty"`
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EventBridgeSendEnabled *bool   `json:"EventBridgeSendEnabled,omitempty" xml:"EventBridgeSendEnabled,omitempty"`
	// example:
	//
	// 2025-11-14T02:11:32Z
	ModifyTime         *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	MonitorSendEnabled *bool  `json:"MonitorSendEnabled,omitempty" xml:"MonitorSendEnabled,omitempty"`
	// example:
	//
	// true
	ReportSendEnabled *bool `json:"ReportSendEnabled,omitempty" xml:"ReportSendEnabled,omitempty"`
	SlsSendEnabled    *bool `json:"SlsSendEnabled,omitempty" xml:"SlsSendEnabled,omitempty"`
	// example:
	//
	// Weekly
	SubscribePeriod *string `json:"SubscribePeriod,omitempty" xml:"SubscribePeriod,omitempty"`
	// example:
	//
	// Subscribe
	SubscribeStatus *string `json:"SubscribeStatus,omitempty" xml:"SubscribeStatus,omitempty"`
}

func (s DescribeAppsResponseBodyAppInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBodyAppInfos) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyAppInfos) GetAppId() *string {
	return s.AppId
}

func (s *DescribeAppsResponseBodyAppInfos) GetAppName() *string {
	return s.AppName
}

func (s *DescribeAppsResponseBodyAppInfos) GetAppTags() []*DescribeAppsResponseBodyAppInfosAppTags {
	return s.AppTags
}

func (s *DescribeAppsResponseBodyAppInfos) GetDefault() *bool {
	return s.Default
}

func (s *DescribeAppsResponseBodyAppInfos) GetDescription() *string {
	return s.Description
}

func (s *DescribeAppsResponseBodyAppInfos) GetEventBridgeSendEnabled() *bool {
	return s.EventBridgeSendEnabled
}

func (s *DescribeAppsResponseBodyAppInfos) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *DescribeAppsResponseBodyAppInfos) GetMonitorSendEnabled() *bool {
	return s.MonitorSendEnabled
}

func (s *DescribeAppsResponseBodyAppInfos) GetReportSendEnabled() *bool {
	return s.ReportSendEnabled
}

func (s *DescribeAppsResponseBodyAppInfos) GetSlsSendEnabled() *bool {
	return s.SlsSendEnabled
}

func (s *DescribeAppsResponseBodyAppInfos) GetSubscribePeriod() *string {
	return s.SubscribePeriod
}

func (s *DescribeAppsResponseBodyAppInfos) GetSubscribeStatus() *string {
	return s.SubscribeStatus
}

func (s *DescribeAppsResponseBodyAppInfos) SetAppId(v string) *DescribeAppsResponseBodyAppInfos {
	s.AppId = &v
	return s
}

func (s *DescribeAppsResponseBodyAppInfos) SetAppName(v string) *DescribeAppsResponseBodyAppInfos {
	s.AppName = &v
	return s
}

func (s *DescribeAppsResponseBodyAppInfos) SetAppTags(v []*DescribeAppsResponseBodyAppInfosAppTags) *DescribeAppsResponseBodyAppInfos {
	s.AppTags = v
	return s
}

func (s *DescribeAppsResponseBodyAppInfos) SetDefault(v bool) *DescribeAppsResponseBodyAppInfos {
	s.Default = &v
	return s
}

func (s *DescribeAppsResponseBodyAppInfos) SetDescription(v string) *DescribeAppsResponseBodyAppInfos {
	s.Description = &v
	return s
}

func (s *DescribeAppsResponseBodyAppInfos) SetEventBridgeSendEnabled(v bool) *DescribeAppsResponseBodyAppInfos {
	s.EventBridgeSendEnabled = &v
	return s
}

func (s *DescribeAppsResponseBodyAppInfos) SetModifyTime(v int64) *DescribeAppsResponseBodyAppInfos {
	s.ModifyTime = &v
	return s
}

func (s *DescribeAppsResponseBodyAppInfos) SetMonitorSendEnabled(v bool) *DescribeAppsResponseBodyAppInfos {
	s.MonitorSendEnabled = &v
	return s
}

func (s *DescribeAppsResponseBodyAppInfos) SetReportSendEnabled(v bool) *DescribeAppsResponseBodyAppInfos {
	s.ReportSendEnabled = &v
	return s
}

func (s *DescribeAppsResponseBodyAppInfos) SetSlsSendEnabled(v bool) *DescribeAppsResponseBodyAppInfos {
	s.SlsSendEnabled = &v
	return s
}

func (s *DescribeAppsResponseBodyAppInfos) SetSubscribePeriod(v string) *DescribeAppsResponseBodyAppInfos {
	s.SubscribePeriod = &v
	return s
}

func (s *DescribeAppsResponseBodyAppInfos) SetSubscribeStatus(v string) *DescribeAppsResponseBodyAppInfos {
	s.SubscribeStatus = &v
	return s
}

func (s *DescribeAppsResponseBodyAppInfos) Validate() error {
	if s.AppTags != nil {
		for _, item := range s.AppTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAppsResponseBodyAppInfosAppTags struct {
	// example:
	//
	// ebs
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// test
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeAppsResponseBodyAppInfosAppTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBodyAppInfosAppTags) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyAppInfosAppTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeAppsResponseBodyAppInfosAppTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeAppsResponseBodyAppInfosAppTags) SetTagKey(v string) *DescribeAppsResponseBodyAppInfosAppTags {
	s.TagKey = &v
	return s
}

func (s *DescribeAppsResponseBodyAppInfosAppTags) SetTagValue(v string) *DescribeAppsResponseBodyAppInfosAppTags {
	s.TagValue = &v
	return s
}

func (s *DescribeAppsResponseBodyAppInfosAppTags) Validate() error {
	return dara.Validate(s)
}
