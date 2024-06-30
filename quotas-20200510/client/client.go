// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateQuotaAlarmRequest struct {
	// This parameter is required.
	AlarmName *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// config
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// q_hvnoqv
	QuotaActionCode *string                                   `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	QuotaDimensions []*CreateQuotaAlarmRequestQuotaDimensions `json:"QuotaDimensions,omitempty" xml:"QuotaDimensions,omitempty" type:"Repeated"`
	// example:
	//
	// 150
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// example:
	//
	// 50
	ThresholdPercent *float32 `json:"ThresholdPercent,omitempty" xml:"ThresholdPercent,omitempty"`
	// example:
	//
	// used
	ThresholdType *string `json:"ThresholdType,omitempty" xml:"ThresholdType,omitempty"`
	// example:
	//
	// https://alert.aliyun.com/callback
	WebHook *string `json:"WebHook,omitempty" xml:"WebHook,omitempty"`
}

func (s CreateQuotaAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaAlarmRequest) GoString() string {
	return s.String()
}

func (s *CreateQuotaAlarmRequest) SetAlarmName(v string) *CreateQuotaAlarmRequest {
	s.AlarmName = &v
	return s
}

func (s *CreateQuotaAlarmRequest) SetProductCode(v string) *CreateQuotaAlarmRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateQuotaAlarmRequest) SetQuotaActionCode(v string) *CreateQuotaAlarmRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *CreateQuotaAlarmRequest) SetQuotaDimensions(v []*CreateQuotaAlarmRequestQuotaDimensions) *CreateQuotaAlarmRequest {
	s.QuotaDimensions = v
	return s
}

func (s *CreateQuotaAlarmRequest) SetThreshold(v float32) *CreateQuotaAlarmRequest {
	s.Threshold = &v
	return s
}

func (s *CreateQuotaAlarmRequest) SetThresholdPercent(v float32) *CreateQuotaAlarmRequest {
	s.ThresholdPercent = &v
	return s
}

func (s *CreateQuotaAlarmRequest) SetThresholdType(v string) *CreateQuotaAlarmRequest {
	s.ThresholdType = &v
	return s
}

func (s *CreateQuotaAlarmRequest) SetWebHook(v string) *CreateQuotaAlarmRequest {
	s.WebHook = &v
	return s
}

type CreateQuotaAlarmRequestQuotaDimensions struct {
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateQuotaAlarmRequestQuotaDimensions) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaAlarmRequestQuotaDimensions) GoString() string {
	return s.String()
}

func (s *CreateQuotaAlarmRequestQuotaDimensions) SetKey(v string) *CreateQuotaAlarmRequestQuotaDimensions {
	s.Key = &v
	return s
}

func (s *CreateQuotaAlarmRequestQuotaDimensions) SetValue(v string) *CreateQuotaAlarmRequestQuotaDimensions {
	s.Value = &v
	return s
}

type CreateQuotaAlarmResponseBody struct {
	// example:
	//
	// 18b3be23-b7b0-4d45-91bc-d0c331aa****
	AlarmId *string `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// example:
	//
	// BD219E2B-E687-45EE-B5F3-61FB730551B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateQuotaAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQuotaAlarmResponseBody) SetAlarmId(v string) *CreateQuotaAlarmResponseBody {
	s.AlarmId = &v
	return s
}

func (s *CreateQuotaAlarmResponseBody) SetRequestId(v string) *CreateQuotaAlarmResponseBody {
	s.RequestId = &v
	return s
}

type CreateQuotaAlarmResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQuotaAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQuotaAlarmResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaAlarmResponse) GoString() string {
	return s.String()
}

func (s *CreateQuotaAlarmResponse) SetHeaders(v map[string]*string) *CreateQuotaAlarmResponse {
	s.Headers = v
	return s
}

func (s *CreateQuotaAlarmResponse) SetStatusCode(v int32) *CreateQuotaAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQuotaAlarmResponse) SetBody(v *CreateQuotaAlarmResponseBody) *CreateQuotaAlarmResponse {
	s.Body = v
	return s
}

type CreateQuotaApplicationRequest struct {
	// The mode in which you want the application to be reviewed. Valid values:
	//
	// 	- Sync: The application is reviewed in a synchronous manner. Quota Center automatically reviews the application. The result is returned immediately after you submit the application. However, the chance of an approval for an application that is reviewed in Sync mode is lower than the chance of an approval for an application that is reviewed in Async mode. The validity period of the new quota value is 1 hour.
	//
	// 	- Async: The application is reviewed in an asynchronous manner. An Alibaba Cloud support engineer reviews the application. The chance of an approval for an application that is reviewed in Async mode is higher than the chance of an approval for an application that is reviewed in Sync mode. The validity period of the new quota value is one month.
	//
	// > This parameter is available only for ECS Quotas by Instance Type.
	//
	// example:
	//
	// Sync
	AuditMode *string `json:"AuditMode,omitempty" xml:"AuditMode,omitempty"`
	// The requested value of the quota.
	//
	// >
	//
	// 	- You can specify the DesireValue parameter based on the values of the `TotalUsage` and `ApplicableRange` parameters that are returned by the [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html) operation.
	//
	// 	- Applications are reviewed by the technical support team of each Alibaba Cloud service. To increase the success rate of your application, you must specify a reasonable quota value and detailed reasons when you submit an application to increase the value of the quota.
	//
	// This parameter is required.
	//
	// example:
	//
	// 804
	DesireValue *float32 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The quota dimensions.
	Dimensions []*CreateQuotaApplicationRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The end time of the validity period of the quota. Specify the value in UTC. This parameter is valid only if you set the QuotaCategory parameter to WhiteListLabel.
	//
	// >  If you do not specify an end time, the default end time is 99 years after the quota application is submitted.
	//
	// example:
	//
	// 2021-01-19T09:25:56Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The language of the quota alert notification. Valid values:
	//
	// 	- zh (default value): Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	EnvLanguage *string `json:"EnvLanguage,omitempty" xml:"EnvLanguage,omitempty"`
	// The start time of the validity period of the quota. Specify the value in UTC. This parameter is valid only if you set the QuotaCategory parameter to WhiteListLabel.
	//
	// >  If you do not specify a start time, the default start time is the time when the quota application is submitted.
	//
	// example:
	//
	// 2021-01-20T09:25:56Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Specifies whether to send a notification about the application result. Valid values:
	//
	// 	- 0 (default value): sends a notification about the application result.
	//
	// 	- 3: A notification about the application result is sent.
	//
	// example:
	//
	// 0
	NoticeType *int32 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// >  To query the abbreviation of an Alibaba Cloud service name, check the `ProductCode` parameter that is described in [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// >  To query the quota ID of an Alibaba Cloud service, check the `QuotaActionCode` parameter that is described in [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// q_security-groups
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The type of the quota.
	//
	// 	- CommonQuota (default value): general quota
	//
	// 	- FlowControl: API rate limit
	//
	// 	- WhiteListLabel: whitelist quota
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The reason for the application.
	//
	// > Applications are reviewed by the technical support team of each Alibaba Cloud service. To increase the success rate of your application, you must specify a reasonable quota value and detailed reasons when you submit an application to increase the value of the quota.
	//
	// This parameter is required.
	//
	// example:
	//
	// Scale Out
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s CreateQuotaApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateQuotaApplicationRequest) SetAuditMode(v string) *CreateQuotaApplicationRequest {
	s.AuditMode = &v
	return s
}

func (s *CreateQuotaApplicationRequest) SetDesireValue(v float32) *CreateQuotaApplicationRequest {
	s.DesireValue = &v
	return s
}

func (s *CreateQuotaApplicationRequest) SetDimensions(v []*CreateQuotaApplicationRequestDimensions) *CreateQuotaApplicationRequest {
	s.Dimensions = v
	return s
}

func (s *CreateQuotaApplicationRequest) SetEffectiveTime(v string) *CreateQuotaApplicationRequest {
	s.EffectiveTime = &v
	return s
}

func (s *CreateQuotaApplicationRequest) SetEnvLanguage(v string) *CreateQuotaApplicationRequest {
	s.EnvLanguage = &v
	return s
}

func (s *CreateQuotaApplicationRequest) SetExpireTime(v string) *CreateQuotaApplicationRequest {
	s.ExpireTime = &v
	return s
}

func (s *CreateQuotaApplicationRequest) SetNoticeType(v int32) *CreateQuotaApplicationRequest {
	s.NoticeType = &v
	return s
}

func (s *CreateQuotaApplicationRequest) SetProductCode(v string) *CreateQuotaApplicationRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateQuotaApplicationRequest) SetQuotaActionCode(v string) *CreateQuotaApplicationRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *CreateQuotaApplicationRequest) SetQuotaCategory(v string) *CreateQuotaApplicationRequest {
	s.QuotaCategory = &v
	return s
}

func (s *CreateQuotaApplicationRequest) SetReason(v string) *CreateQuotaApplicationRequest {
	s.Reason = &v
	return s
}

type CreateQuotaApplicationRequestDimensions struct {
	// The key of the dimension.
	//
	// >
	//
	// 	- The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// 	- This parameter is required if you set the `ProductCode` parameter to `ecs`, `ecs-spec`, `actiontrail`, or `ess`.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the dimension.
	//
	// >
	//
	// 	- The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// 	- This parameter is required if you set the `ProductCode` parameter to `ecs`, `ecs-spec`, `actiontrail`, or `ess`.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateQuotaApplicationRequestDimensions) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaApplicationRequestDimensions) GoString() string {
	return s.String()
}

func (s *CreateQuotaApplicationRequestDimensions) SetKey(v string) *CreateQuotaApplicationRequestDimensions {
	s.Key = &v
	return s
}

func (s *CreateQuotaApplicationRequestDimensions) SetValue(v string) *CreateQuotaApplicationRequestDimensions {
	s.Value = &v
	return s
}

type CreateQuotaApplicationResponseBody struct {
	// The ID of the application.
	//
	// example:
	//
	// d314d6ae-867d-484c-9009-3d421a80****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The time when the application was submitted.
	//
	// example:
	//
	// 2021-01-19T09:25:56Z
	ApplyTime *string `json:"ApplyTime,omitempty" xml:"ApplyTime,omitempty"`
	// The quota value that is approved.
	//
	// example:
	//
	// 804
	ApproveValue *float32 `json:"ApproveValue,omitempty" xml:"ApproveValue,omitempty"`
	// The result of the application.
	//
	// example:
	//
	// Agree
	AuditReason *string `json:"AuditReason,omitempty" xml:"AuditReason,omitempty"`
	// The requested value of the quota.
	//
	// example:
	//
	// 12
	DesireValue *int32 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The quota dimension.
	//
	// example:
	//
	// {"regionId":"cn-hangzhou"}
	Dimension map[string]interface{} `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The time when the new quota value takes effect.
	//
	// example:
	//
	// 2021-01-19T09:25:56Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The time when the new quota expires.
	//
	// example:
	//
	// 2021-01-20T09:25:56Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether Quota Center sends a notification about the application result. Valid values:
	//
	// 	- 0: Quota Center does not send a notification.
	//
	// 	- 3: Quota Center sends a notification.
	//
	// example:
	//
	// 3
	NoticeType *int64 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// example:
	//
	// ecs-spec
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// ecs.c5.large
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the quota.
	//
	// example:
	//
	// acs:quotas:cn-hangzhou:*:quota/ecs/ecs.m2.medium/prepaid/classic/instancetype/cn-hangzhou-b
	QuotaArn *string `json:"QuotaArn,omitempty" xml:"QuotaArn,omitempty"`
	// The description of the quota.
	//
	// example:
	//
	// ecs.c5.large
	QuotaDescription *string `json:"QuotaDescription,omitempty" xml:"QuotaDescription,omitempty"`
	// The name of the quota.
	//
	// example:
	//
	// ecs.c5.large
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The unit of the quota.
	//
	// example:
	//
	// AMOUNT
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// The reason for the application.
	//
	// example:
	//
	// Scale Out
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D47B3A10-CDAC-5412-B2EE-EC9A3DBE9053
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the application. Valid values:
	//
	// 	- Disagree: The application is rejected.
	//
	// 	- Agree: The application is approved.
	//
	// 	- Process: The application is being reviewed.
	//
	// 	- Cancel: The application is canceled.
	//
	// example:
	//
	// Process
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateQuotaApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQuotaApplicationResponseBody) SetApplicationId(v string) *CreateQuotaApplicationResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetApplyTime(v string) *CreateQuotaApplicationResponseBody {
	s.ApplyTime = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetApproveValue(v float32) *CreateQuotaApplicationResponseBody {
	s.ApproveValue = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetAuditReason(v string) *CreateQuotaApplicationResponseBody {
	s.AuditReason = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetDesireValue(v int32) *CreateQuotaApplicationResponseBody {
	s.DesireValue = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetDimension(v map[string]interface{}) *CreateQuotaApplicationResponseBody {
	s.Dimension = v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetEffectiveTime(v string) *CreateQuotaApplicationResponseBody {
	s.EffectiveTime = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetExpireTime(v string) *CreateQuotaApplicationResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetNoticeType(v int64) *CreateQuotaApplicationResponseBody {
	s.NoticeType = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetProductCode(v string) *CreateQuotaApplicationResponseBody {
	s.ProductCode = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetQuotaActionCode(v string) *CreateQuotaApplicationResponseBody {
	s.QuotaActionCode = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetQuotaArn(v string) *CreateQuotaApplicationResponseBody {
	s.QuotaArn = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetQuotaDescription(v string) *CreateQuotaApplicationResponseBody {
	s.QuotaDescription = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetQuotaName(v string) *CreateQuotaApplicationResponseBody {
	s.QuotaName = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetQuotaUnit(v string) *CreateQuotaApplicationResponseBody {
	s.QuotaUnit = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetReason(v string) *CreateQuotaApplicationResponseBody {
	s.Reason = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetRequestId(v string) *CreateQuotaApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetStatus(v string) *CreateQuotaApplicationResponseBody {
	s.Status = &v
	return s
}

type CreateQuotaApplicationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQuotaApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQuotaApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaApplicationResponse) GoString() string {
	return s.String()
}

func (s *CreateQuotaApplicationResponse) SetHeaders(v map[string]*string) *CreateQuotaApplicationResponse {
	s.Headers = v
	return s
}

func (s *CreateQuotaApplicationResponse) SetStatusCode(v int32) *CreateQuotaApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQuotaApplicationResponse) SetBody(v *CreateQuotaApplicationResponseBody) *CreateQuotaApplicationResponse {
	s.Body = v
	return s
}

type CreateQuotaApplicationsForTemplateRequest struct {
	// The Alibaba Cloud accounts that correspond to the resource directory members for which the quotas are applied.
	//
	// >  You can submit a quota increase application for a maximum of 50 members at a time. For more information about the members of a resource directory, see [ListAccounts](https://help.aliyun.com/document_detail/604207.html).
	//
	// This parameter is required.
	AliyunUids []*string `json:"AliyunUids,omitempty" xml:"AliyunUids,omitempty" type:"Repeated"`
	// The requested value of the quota.
	//
	// >
	//
	// 	- You can specify DesireValue based on the values of `TotalUsage` and `ApplicableRange` in the response to the [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html) operation.
	//
	// 	- Applications are reviewed by the technical support team of each Alibaba Cloud service. To increase the success rate of your application, specify a reasonable quota value and a detailed reason.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	DesireValue *float64 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The quota dimensions.
	Dimensions []*CreateQuotaApplicationsForTemplateRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The start time of the validity period of the quota. Specify the value in UTC. This parameter is valid only if you set the QuotaCategory parameter to WhiteListLabel.
	//
	// >  If you do not specify a start time, the value is the time when the quota application is submitted.
	//
	// example:
	//
	// 2021-01-19T09:25:56Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The language of the notification about the application result. Valid values:
	//
	// 	- zh (default): Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	EnvLanguage *string `json:"EnvLanguage,omitempty" xml:"EnvLanguage,omitempty"`
	// The end time of the validity period of the quota. Specify the value in UTC. This parameter is valid only if you set the QuotaCategory parameter to WhiteListLabel.
	//
	// >  If you do not specify an end time, the value is 99 years after the start time of the validity period.
	//
	// example:
	//
	// 2021-01-20T09:25:56Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Specifies whether to send a notification about the application result. Valid values:
	//
	// 	- 0 (default): no
	//
	// 	- 3: yes
	//
	// example:
	//
	// 0
	NoticeType *int32 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// >  To query the abbreviation of an Alibaba Cloud service name, call the [ListProducts](https://help.aliyun.com/document_detail/440555.html) operation and check the value of `ProductCode` in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs-spec
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// >  To query the quota ID of an Alibaba Cloud service, call the [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html) and check the value of `QuotaActionCode` in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs.g5.2xlarge
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The quota type. Valid values:
	//
	// 	- CommonQuota: general quota
	//
	// 	- FlowControl: API rate limit
	//
	// 	- WhiteListLabel: privilege
	//
	// This parameter is required.
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The reason for the quota application.
	//
	// >  Applications are reviewed by the technical support team of each Alibaba Cloud service. To increase the success rate of your application, you must specify a reasonable quota value and detailed reasons when you submit the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// Scale Out
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s CreateQuotaApplicationsForTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaApplicationsForTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateQuotaApplicationsForTemplateRequest) SetAliyunUids(v []*string) *CreateQuotaApplicationsForTemplateRequest {
	s.AliyunUids = v
	return s
}

func (s *CreateQuotaApplicationsForTemplateRequest) SetDesireValue(v float64) *CreateQuotaApplicationsForTemplateRequest {
	s.DesireValue = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateRequest) SetDimensions(v []*CreateQuotaApplicationsForTemplateRequestDimensions) *CreateQuotaApplicationsForTemplateRequest {
	s.Dimensions = v
	return s
}

func (s *CreateQuotaApplicationsForTemplateRequest) SetEffectiveTime(v string) *CreateQuotaApplicationsForTemplateRequest {
	s.EffectiveTime = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateRequest) SetEnvLanguage(v string) *CreateQuotaApplicationsForTemplateRequest {
	s.EnvLanguage = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateRequest) SetExpireTime(v string) *CreateQuotaApplicationsForTemplateRequest {
	s.ExpireTime = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateRequest) SetNoticeType(v int32) *CreateQuotaApplicationsForTemplateRequest {
	s.NoticeType = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateRequest) SetProductCode(v string) *CreateQuotaApplicationsForTemplateRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateRequest) SetQuotaActionCode(v string) *CreateQuotaApplicationsForTemplateRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateRequest) SetQuotaCategory(v string) *CreateQuotaApplicationsForTemplateRequest {
	s.QuotaCategory = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateRequest) SetReason(v string) *CreateQuotaApplicationsForTemplateRequest {
	s.Reason = &v
	return s
}

type CreateQuotaApplicationsForTemplateRequestDimensions struct {
	// The key of the dimension.
	//
	// >
	//
	// 	- The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// 	- This parameter is required if you set the `ProductCode` parameter to `ecs`, `ecs-spec`, `actiontrail`, or `ess`.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the dimension.
	//
	// >
	//
	// 	- The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// 	- This parameter is required if you set the `ProductCode` parameter to `ecs`, `ecs-spec`, `actiontrail`, or `ess`.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateQuotaApplicationsForTemplateRequestDimensions) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaApplicationsForTemplateRequestDimensions) GoString() string {
	return s.String()
}

func (s *CreateQuotaApplicationsForTemplateRequestDimensions) SetKey(v string) *CreateQuotaApplicationsForTemplateRequestDimensions {
	s.Key = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateRequestDimensions) SetValue(v string) *CreateQuotaApplicationsForTemplateRequestDimensions {
	s.Value = &v
	return s
}

type CreateQuotaApplicationsForTemplateResponseBody struct {
	// The Alibaba Cloud accounts for which the quotas are applied.
	AliyunUids []*string `json:"AliyunUids,omitempty" xml:"AliyunUids,omitempty" type:"Repeated"`
	// The ID of the quota application batch.
	//
	// example:
	//
	// d314d6ae-867d-484c-9009-3d421a80****
	BatchQuotaApplicationId *string `json:"BatchQuotaApplicationId,omitempty" xml:"BatchQuotaApplicationId,omitempty"`
	// The Alibaba Cloud accounts of the members in a resource directory whose quota increase request is rejected, and the reason for the rejection.
	FailResults []*CreateQuotaApplicationsForTemplateResponseBodyFailResults `json:"FailResults,omitempty" xml:"FailResults,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 8FF8CAF0-29D9-4F11-B6A4-FD2CBCA016D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateQuotaApplicationsForTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaApplicationsForTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQuotaApplicationsForTemplateResponseBody) SetAliyunUids(v []*string) *CreateQuotaApplicationsForTemplateResponseBody {
	s.AliyunUids = v
	return s
}

func (s *CreateQuotaApplicationsForTemplateResponseBody) SetBatchQuotaApplicationId(v string) *CreateQuotaApplicationsForTemplateResponseBody {
	s.BatchQuotaApplicationId = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateResponseBody) SetFailResults(v []*CreateQuotaApplicationsForTemplateResponseBodyFailResults) *CreateQuotaApplicationsForTemplateResponseBody {
	s.FailResults = v
	return s
}

func (s *CreateQuotaApplicationsForTemplateResponseBody) SetRequestId(v string) *CreateQuotaApplicationsForTemplateResponseBody {
	s.RequestId = &v
	return s
}

type CreateQuotaApplicationsForTemplateResponseBodyFailResults struct {
	// The Alibaba Cloud account of the members in a resource directory whose quota increase request is rejected.
	//
	// example:
	//
	// 135048337611****
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// The reason for the rejection.
	//
	// example:
	//
	// The quota adjustment application is being processed. Please try again later.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s CreateQuotaApplicationsForTemplateResponseBodyFailResults) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaApplicationsForTemplateResponseBodyFailResults) GoString() string {
	return s.String()
}

func (s *CreateQuotaApplicationsForTemplateResponseBodyFailResults) SetAliyunUid(v string) *CreateQuotaApplicationsForTemplateResponseBodyFailResults {
	s.AliyunUid = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateResponseBodyFailResults) SetReason(v string) *CreateQuotaApplicationsForTemplateResponseBodyFailResults {
	s.Reason = &v
	return s
}

type CreateQuotaApplicationsForTemplateResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQuotaApplicationsForTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQuotaApplicationsForTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaApplicationsForTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateQuotaApplicationsForTemplateResponse) SetHeaders(v map[string]*string) *CreateQuotaApplicationsForTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateQuotaApplicationsForTemplateResponse) SetStatusCode(v int32) *CreateQuotaApplicationsForTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateResponse) SetBody(v *CreateQuotaApplicationsForTemplateResponseBody) *CreateQuotaApplicationsForTemplateResponse {
	s.Body = v
	return s
}

type CreateTemplateQuotaItemRequest struct {
	// The requested value of the quota.
	//
	// >
	//
	// 	- You can specify DesireValue based on the values of `TotalUsage` and `ApplicableRange` in the response to the [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html) operation.
	//
	// 	- Applications are reviewed by the technical support team for each cloud service. To increase the success rate of your application, specify a reasonable quota value and a detailed reason.
	//
	// This parameter is required.
	//
	// example:
	//
	// 804
	DesireValue *float32 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The quota dimensions.
	Dimensions []*CreateTemplateQuotaItemRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The start time of the validity period of the quota. Specify the value in UTC. This parameter is valid only if you set the QuotaCategory parameter to WhiteListLabel.
	//
	// > If you leave this parameter empty, the quota takes effect immediately.
	//
	// example:
	//
	// 2021-01-19T09:25:56Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The language of the quota alert notification. Valid values:
	//
	// 	- zh (default value): Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	EnvLanguage *string `json:"EnvLanguage,omitempty" xml:"EnvLanguage,omitempty"`
	// The end time of the validity period of the quota. Specify the value in UTC. This parameter is valid only if you set the QuotaCategory parameter to WhiteListLabel.
	//
	// > If you leave this parameter empty, no end time is specified.
	//
	// example:
	//
	// 2021-01-20T09:25:56Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Specifies whether to send a notification about the application result. Valid values:
	//
	// 	- 0 (default value): no
	//
	// 	- 3: yes
	//
	// example:
	//
	// 0
	NoticeType *int64 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// >  To query the abbreviation of an Alibaba Cloud service name, call the [ListProducts](https://help.aliyun.com/document_detail/440555.html) operation and check the value of `ProductCode` in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// >  To obtain the quota ID of an Alibaba Cloud service, call the [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html) operation and check the value of `QuotaActionCode` in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// q_security-groups
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The quota type. Valid values:
	//
	// 	- CommonQuota: general quota
	//
	// 	- WhiteListLabel: privilege
	//
	// 	- FlowControl: API rate limit
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
}

func (s CreateTemplateQuotaItemRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateQuotaItemRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateQuotaItemRequest) SetDesireValue(v float32) *CreateTemplateQuotaItemRequest {
	s.DesireValue = &v
	return s
}

func (s *CreateTemplateQuotaItemRequest) SetDimensions(v []*CreateTemplateQuotaItemRequestDimensions) *CreateTemplateQuotaItemRequest {
	s.Dimensions = v
	return s
}

func (s *CreateTemplateQuotaItemRequest) SetEffectiveTime(v string) *CreateTemplateQuotaItemRequest {
	s.EffectiveTime = &v
	return s
}

func (s *CreateTemplateQuotaItemRequest) SetEnvLanguage(v string) *CreateTemplateQuotaItemRequest {
	s.EnvLanguage = &v
	return s
}

func (s *CreateTemplateQuotaItemRequest) SetExpireTime(v string) *CreateTemplateQuotaItemRequest {
	s.ExpireTime = &v
	return s
}

func (s *CreateTemplateQuotaItemRequest) SetNoticeType(v int64) *CreateTemplateQuotaItemRequest {
	s.NoticeType = &v
	return s
}

func (s *CreateTemplateQuotaItemRequest) SetProductCode(v string) *CreateTemplateQuotaItemRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateTemplateQuotaItemRequest) SetQuotaActionCode(v string) *CreateTemplateQuotaItemRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *CreateTemplateQuotaItemRequest) SetQuotaCategory(v string) *CreateTemplateQuotaItemRequest {
	s.QuotaCategory = &v
	return s
}

type CreateTemplateQuotaItemRequestDimensions struct {
	// The key of the dimension.
	//
	// >
	//
	// 	- The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// 	- This parameter is required if you set the `ProductCode` parameter to `ecs`, `ecs-spec`, `actiontrail`, or `ess`.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the dimension.
	//
	// >
	//
	// 	- The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// 	- This parameter is required if you set the `ProductCode` parameter to `ecs`, `ecs-spec`, `actiontrail`, or `ess`.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTemplateQuotaItemRequestDimensions) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateQuotaItemRequestDimensions) GoString() string {
	return s.String()
}

func (s *CreateTemplateQuotaItemRequestDimensions) SetKey(v string) *CreateTemplateQuotaItemRequestDimensions {
	s.Key = &v
	return s
}

func (s *CreateTemplateQuotaItemRequestDimensions) SetValue(v string) *CreateTemplateQuotaItemRequestDimensions {
	s.Value = &v
	return s
}

type CreateTemplateQuotaItemResponseBody struct {
	// The ID of the quota template.
	//
	// example:
	//
	// 1****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D47B3A10-CDAC-5412-B2EE-EC9A3DBE9053
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTemplateQuotaItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateQuotaItemResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateQuotaItemResponseBody) SetId(v string) *CreateTemplateQuotaItemResponseBody {
	s.Id = &v
	return s
}

func (s *CreateTemplateQuotaItemResponseBody) SetRequestId(v string) *CreateTemplateQuotaItemResponseBody {
	s.RequestId = &v
	return s
}

type CreateTemplateQuotaItemResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTemplateQuotaItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTemplateQuotaItemResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateQuotaItemResponse) GoString() string {
	return s.String()
}

func (s *CreateTemplateQuotaItemResponse) SetHeaders(v map[string]*string) *CreateTemplateQuotaItemResponse {
	s.Headers = v
	return s
}

func (s *CreateTemplateQuotaItemResponse) SetStatusCode(v int32) *CreateTemplateQuotaItemResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTemplateQuotaItemResponse) SetBody(v *CreateTemplateQuotaItemResponseBody) *CreateTemplateQuotaItemResponse {
	s.Body = v
	return s
}

type DeleteQuotaAlarmRequest struct {
	// The ID of the quota alert.
	//
	// >  You can call the [ListQuotaAlarms](https://help.aliyun.com/document_detail/440561.html) operation to obtain the ID of a quota alert.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6b512ab7-da3a-4142-b529-2b2a9294****
	AlarmId *string `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
}

func (s DeleteQuotaAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteQuotaAlarmRequest) GoString() string {
	return s.String()
}

func (s *DeleteQuotaAlarmRequest) SetAlarmId(v string) *DeleteQuotaAlarmRequest {
	s.AlarmId = &v
	return s
}

type DeleteQuotaAlarmResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A95C65B3-7CF4-469E-B1D5-1CA0628A6411
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteQuotaAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteQuotaAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQuotaAlarmResponseBody) SetRequestId(v string) *DeleteQuotaAlarmResponseBody {
	s.RequestId = &v
	return s
}

type DeleteQuotaAlarmResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQuotaAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQuotaAlarmResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteQuotaAlarmResponse) GoString() string {
	return s.String()
}

func (s *DeleteQuotaAlarmResponse) SetHeaders(v map[string]*string) *DeleteQuotaAlarmResponse {
	s.Headers = v
	return s
}

func (s *DeleteQuotaAlarmResponse) SetStatusCode(v int32) *DeleteQuotaAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQuotaAlarmResponse) SetBody(v *DeleteQuotaAlarmResponseBody) *DeleteQuotaAlarmResponse {
	s.Body = v
	return s
}

type DeleteTemplateQuotaItemRequest struct {
	// The ID of the quota template.
	//
	// You can call the [ListQuotaApplicationTemplates](https://help.aliyun.com/document_detail/450403.html) operation to obtain the ID of a quota template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteTemplateQuotaItemRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateQuotaItemRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplateQuotaItemRequest) SetId(v string) *DeleteTemplateQuotaItemRequest {
	s.Id = &v
	return s
}

type DeleteTemplateQuotaItemResponseBody struct {
	// The ID of the quota template.
	//
	// example:
	//
	// 1****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D47B3A10-CDAC-5412-B2EE-EC9A3DBE9053
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTemplateQuotaItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateQuotaItemResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplateQuotaItemResponseBody) SetId(v string) *DeleteTemplateQuotaItemResponseBody {
	s.Id = &v
	return s
}

func (s *DeleteTemplateQuotaItemResponseBody) SetRequestId(v string) *DeleteTemplateQuotaItemResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTemplateQuotaItemResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTemplateQuotaItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTemplateQuotaItemResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateQuotaItemResponse) GoString() string {
	return s.String()
}

func (s *DeleteTemplateQuotaItemResponse) SetHeaders(v map[string]*string) *DeleteTemplateQuotaItemResponse {
	s.Headers = v
	return s
}

func (s *DeleteTemplateQuotaItemResponse) SetStatusCode(v int32) *DeleteTemplateQuotaItemResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTemplateQuotaItemResponse) SetBody(v *DeleteTemplateQuotaItemResponseBody) *DeleteTemplateQuotaItemResponse {
	s.Body = v
	return s
}

type GetProductQuotaRequest struct {
	// The quota dimensions.
	//
	// example:
	//
	// {\\"regionId\\":\\"cn-beijing\\"}
	Dimensions []*GetProductQuotaRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// > For more information, see [Alibaba Cloud services that support Quota Center](https://help.aliyun.com/document_detail/182368.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	//
	// This parameter is required.
	//
	// example:
	//
	// q_security-groups
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
}

func (s GetProductQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProductQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetProductQuotaRequest) SetDimensions(v []*GetProductQuotaRequestDimensions) *GetProductQuotaRequest {
	s.Dimensions = v
	return s
}

func (s *GetProductQuotaRequest) SetProductCode(v string) *GetProductQuotaRequest {
	s.ProductCode = &v
	return s
}

func (s *GetProductQuotaRequest) SetQuotaActionCode(v string) *GetProductQuotaRequest {
	s.QuotaActionCode = &v
	return s
}

type GetProductQuotaRequestDimensions struct {
	// The key of the dimension.
	//
	// >- The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// >- If you call the operation to query the details of a quota that belongs to a cloud service that supports dimensions, you must configure this parameter. You must configure the `Dimensions.N.Key` and `Dimensions.N.Value` parameters at the same time. The following cloud services support dimensions: ECS whose service code is ecs, Enterprise Distributed Application Service (EDAS) whose service code is edas, ECS Quotas by Instance Type whose service code is ecs-spec, and Auto Scaling (ESS) whose service code is ess.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the dimension.
	//
	// > - The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// > - If you call the operation to query the details of a quota that belongs to a cloud service that supports dimensions, you must configure this parameter. You must configure the `Dimensions.N.Key` and `Dimensions.N.Value` parameters at the same time. The following cloud services support dimensions: ECS whose service code is ecs, EDAS whose service code is edas, ECS Quotas by Instance Type whose service code is ecs-spec, and ESS whose service code is ess.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetProductQuotaRequestDimensions) String() string {
	return tea.Prettify(s)
}

func (s GetProductQuotaRequestDimensions) GoString() string {
	return s.String()
}

func (s *GetProductQuotaRequestDimensions) SetKey(v string) *GetProductQuotaRequestDimensions {
	s.Key = &v
	return s
}

func (s *GetProductQuotaRequestDimensions) SetValue(v string) *GetProductQuotaRequestDimensions {
	s.Value = &v
	return s
}

type GetProductQuotaResponseBody struct {
	// The details of the quota.
	Quota *GetProductQuotaResponseBodyQuota `json:"Quota,omitempty" xml:"Quota,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 8FF8CAF0-29D9-4F11-B6A4-FD2CBCA016D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProductQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProductQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductQuotaResponseBody) SetQuota(v *GetProductQuotaResponseBodyQuota) *GetProductQuotaResponseBody {
	s.Quota = v
	return s
}

func (s *GetProductQuotaResponseBody) SetRequestId(v string) *GetProductQuotaResponseBody {
	s.RequestId = &v
	return s
}

type GetProductQuotaResponseBodyQuota struct {
	// Indicates whether the quota is adjustable. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Adjustable *bool `json:"Adjustable,omitempty" xml:"Adjustable,omitempty"`
	// The range of the quota value.
	ApplicableRange []*float32 `json:"ApplicableRange,omitempty" xml:"ApplicableRange,omitempty" type:"Repeated"`
	// The type of the adjustable value. Valid values:
	//
	// 	- continuous
	//
	// 	- discontinuous
	//
	// example:
	//
	// continuous
	ApplicableType *string `json:"ApplicableType,omitempty" xml:"ApplicableType,omitempty"`
	// The reason for submitting a quota increase request.
	//
	// example:
	//
	// The business xxx is expected to grow by 50%.
	ApplyReasonTips *string `json:"ApplyReasonTips,omitempty" xml:"ApplyReasonTips,omitempty"`
	// Indicates whether the system shows the used value of the quota. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Consumable *bool `json:"Consumable,omitempty" xml:"Consumable,omitempty"`
	// The quota dimension. Format: `{"regionId":"Region"}`.
	//
	// example:
	//
	// {"regionId":"cn-hangzhou"}
	Dimensions map[string]interface{} `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The start time of the validity period of the quota. Specify the value in UTC.
	//
	// example:
	//
	// 2022-09-28T06:06:00Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The end time of the validity period of the quota. Specify the value in UTC.
	//
	// example:
	//
	// 2022-09-29T06:06:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the quota is a global quota. Valid values:
	//
	// 	- true: The quota is shared in all regions.
	//
	// 	- false: The quota is independently used in a region.
	//
	// example:
	//
	// true
	GlobalQuota *bool `json:"GlobalQuota,omitempty" xml:"GlobalQuota,omitempty"`
	// The calculation cycle of the quota.
	Period *GetProductQuotaResponseBodyQuotaPeriod `json:"Period,omitempty" xml:"Period,omitempty" type:"Struct"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	//
	// example:
	//
	// q_security-groups
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the quota.
	//
	// example:
	//
	// acs:quotas:cn-hangzhou:120886317861****:quota/ecs/q_security-groups/
	QuotaArn *string `json:"QuotaArn,omitempty" xml:"QuotaArn,omitempty"`
	// The type of the quota. Valid values:
	//
	// 	- CommonQuota: general quota
	//
	// 	- FlowControl: API rate limit
	//
	// 	- WhiteListLabel: whitelist quota
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The description of the quota.
	//
	// example:
	//
	// The maximum number of security groups that can be owned by the current account
	QuotaDescription *string `json:"QuotaDescription,omitempty" xml:"QuotaDescription,omitempty"`
	// The details of the quotas.
	QuotaItems []*GetProductQuotaResponseBodyQuotaQuotaItems `json:"QuotaItems,omitempty" xml:"QuotaItems,omitempty" type:"Repeated"`
	// The name of the quota.
	//
	// example:
	//
	// Maximum Number of Security Groups
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The type of the quota. Valid values:
	//
	// 	- privilege
	//
	// 	- normal (default value)
	//
	// example:
	//
	// normal
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// The unit of the new quota value.
	//
	// > The unit of each quota is unique.*	- For example, the quota whose ID is `q_cbdch3` represents the maximum number of Container Service for Kubernetes (ACK) clusters. The unit of this quota is clusters. The quota whose ID is `q_security-groups` represents the maximum number of security groups. The unit of this quota is security groups.
	//
	// example:
	//
	// Count
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// The range of the quota value.
	SupportedRange []*float32 `json:"SupportedRange,omitempty" xml:"SupportedRange,omitempty" type:"Repeated"`
	// The value of the quota.
	//
	// example:
	//
	// 801
	TotalQuota *float32 `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	// The used quota.
	//
	// example:
	//
	// 26
	TotalUsage *float32 `json:"TotalUsage,omitempty" xml:"TotalUsage,omitempty"`
	// The reason why the quota is not adjustable. Valid values:
	//
	// 	- nonactivated: The service is not activated.
	//
	// 	- applicationProcess: The application is being processed.
	//
	// 	- limitReached: The quota limit is reached.
	//
	// 	- supportTicketRequired: The quota can be increased only by submitting a ticket.
	//
	// example:
	//
	// limitReached
	UnadjustableDetail *string                                      `json:"UnadjustableDetail,omitempty" xml:"UnadjustableDetail,omitempty"`
	UsageMetric        *GetProductQuotaResponseBodyQuotaUsageMetric `json:"UsageMetric,omitempty" xml:"UsageMetric,omitempty" type:"Struct"`
}

func (s GetProductQuotaResponseBodyQuota) String() string {
	return tea.Prettify(s)
}

func (s GetProductQuotaResponseBodyQuota) GoString() string {
	return s.String()
}

func (s *GetProductQuotaResponseBodyQuota) SetAdjustable(v bool) *GetProductQuotaResponseBodyQuota {
	s.Adjustable = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetApplicableRange(v []*float32) *GetProductQuotaResponseBodyQuota {
	s.ApplicableRange = v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetApplicableType(v string) *GetProductQuotaResponseBodyQuota {
	s.ApplicableType = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetApplyReasonTips(v string) *GetProductQuotaResponseBodyQuota {
	s.ApplyReasonTips = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetConsumable(v bool) *GetProductQuotaResponseBodyQuota {
	s.Consumable = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetDimensions(v map[string]interface{}) *GetProductQuotaResponseBodyQuota {
	s.Dimensions = v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetEffectiveTime(v string) *GetProductQuotaResponseBodyQuota {
	s.EffectiveTime = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetExpireTime(v string) *GetProductQuotaResponseBodyQuota {
	s.ExpireTime = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetGlobalQuota(v bool) *GetProductQuotaResponseBodyQuota {
	s.GlobalQuota = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetPeriod(v *GetProductQuotaResponseBodyQuotaPeriod) *GetProductQuotaResponseBodyQuota {
	s.Period = v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetProductCode(v string) *GetProductQuotaResponseBodyQuota {
	s.ProductCode = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetQuotaActionCode(v string) *GetProductQuotaResponseBodyQuota {
	s.QuotaActionCode = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetQuotaArn(v string) *GetProductQuotaResponseBodyQuota {
	s.QuotaArn = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetQuotaCategory(v string) *GetProductQuotaResponseBodyQuota {
	s.QuotaCategory = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetQuotaDescription(v string) *GetProductQuotaResponseBodyQuota {
	s.QuotaDescription = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetQuotaItems(v []*GetProductQuotaResponseBodyQuotaQuotaItems) *GetProductQuotaResponseBodyQuota {
	s.QuotaItems = v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetQuotaName(v string) *GetProductQuotaResponseBodyQuota {
	s.QuotaName = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetQuotaType(v string) *GetProductQuotaResponseBodyQuota {
	s.QuotaType = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetQuotaUnit(v string) *GetProductQuotaResponseBodyQuota {
	s.QuotaUnit = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetSupportedRange(v []*float32) *GetProductQuotaResponseBodyQuota {
	s.SupportedRange = v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetTotalQuota(v float32) *GetProductQuotaResponseBodyQuota {
	s.TotalQuota = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetTotalUsage(v float32) *GetProductQuotaResponseBodyQuota {
	s.TotalUsage = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetUnadjustableDetail(v string) *GetProductQuotaResponseBodyQuota {
	s.UnadjustableDetail = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetUsageMetric(v *GetProductQuotaResponseBodyQuotaUsageMetric) *GetProductQuotaResponseBodyQuota {
	s.UsageMetric = v
	return s
}

type GetProductQuotaResponseBodyQuotaPeriod struct {
	// The unit of the calculation cycle of the quota. Valid values:
	//
	// 	- second
	//
	// 	- minute
	//
	// 	- hour
	//
	// 	- day
	//
	// 	- week
	//
	// example:
	//
	// Day
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The value of the calculation cycle of the quota.
	//
	// example:
	//
	// 1
	PeriodValue *int32 `json:"PeriodValue,omitempty" xml:"PeriodValue,omitempty"`
}

func (s GetProductQuotaResponseBodyQuotaPeriod) String() string {
	return tea.Prettify(s)
}

func (s GetProductQuotaResponseBodyQuotaPeriod) GoString() string {
	return s.String()
}

func (s *GetProductQuotaResponseBodyQuotaPeriod) SetPeriodUnit(v string) *GetProductQuotaResponseBodyQuotaPeriod {
	s.PeriodUnit = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuotaPeriod) SetPeriodValue(v int32) *GetProductQuotaResponseBodyQuotaPeriod {
	s.PeriodValue = &v
	return s
}

type GetProductQuotaResponseBodyQuotaQuotaItems struct {
	// The value of the quota.
	//
	// example:
	//
	// 801
	Quota *string `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// The unit of the quota.
	//
	// >  The unit of each quota is unique. For example, the quota whose ID is `q_cbdch3` represents the maximum number of ACK clusters. The unit of this quota is clusters. The quota whose ID is `q_security-groups` represents the maximum number of security groups. The unit of this quota is security groups.
	//
	// example:
	//
	// Count
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// The category of the quota. Valid values:
	//
	// 	- BaseQuota: base quota.
	//
	// 	- ReservedQuota: reserved quota.
	//
	// example:
	//
	// BaseQuota
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The used quota.
	//
	// example:
	//
	// 26
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s GetProductQuotaResponseBodyQuotaQuotaItems) String() string {
	return tea.Prettify(s)
}

func (s GetProductQuotaResponseBodyQuotaQuotaItems) GoString() string {
	return s.String()
}

func (s *GetProductQuotaResponseBodyQuotaQuotaItems) SetQuota(v string) *GetProductQuotaResponseBodyQuotaQuotaItems {
	s.Quota = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuotaQuotaItems) SetQuotaUnit(v string) *GetProductQuotaResponseBodyQuotaQuotaItems {
	s.QuotaUnit = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuotaQuotaItems) SetType(v string) *GetProductQuotaResponseBodyQuotaQuotaItems {
	s.Type = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuotaQuotaItems) SetUsage(v string) *GetProductQuotaResponseBodyQuotaQuotaItems {
	s.Usage = &v
	return s
}

type GetProductQuotaResponseBodyQuotaUsageMetric struct {
	MetricDimensions map[string]*string `json:"MetricDimensions,omitempty" xml:"MetricDimensions,omitempty"`
	MetricName       *string            `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	MetricNamespace  *string            `json:"MetricNamespace,omitempty" xml:"MetricNamespace,omitempty"`
}

func (s GetProductQuotaResponseBodyQuotaUsageMetric) String() string {
	return tea.Prettify(s)
}

func (s GetProductQuotaResponseBodyQuotaUsageMetric) GoString() string {
	return s.String()
}

func (s *GetProductQuotaResponseBodyQuotaUsageMetric) SetMetricDimensions(v map[string]*string) *GetProductQuotaResponseBodyQuotaUsageMetric {
	s.MetricDimensions = v
	return s
}

func (s *GetProductQuotaResponseBodyQuotaUsageMetric) SetMetricName(v string) *GetProductQuotaResponseBodyQuotaUsageMetric {
	s.MetricName = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuotaUsageMetric) SetMetricNamespace(v string) *GetProductQuotaResponseBodyQuotaUsageMetric {
	s.MetricNamespace = &v
	return s
}

type GetProductQuotaResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProductQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProductQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProductQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetProductQuotaResponse) SetHeaders(v map[string]*string) *GetProductQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetProductQuotaResponse) SetStatusCode(v int32) *GetProductQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProductQuotaResponse) SetBody(v *GetProductQuotaResponseBody) *GetProductQuotaResponse {
	s.Body = v
	return s
}

type GetProductQuotaDimensionRequest struct {
	// The information about quota dimensions.
	DependentDimensions []*GetProductQuotaDimensionRequestDependentDimensions `json:"DependentDimensions,omitempty" xml:"DependentDimensions,omitempty" type:"Repeated"`
	// The key of the quota dimension.
	//
	// example:
	//
	// regionId
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// > For more information, see [Alibaba Cloud services that support Quota Center](https://help.aliyun.com/document_detail/182368.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs-spec
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s GetProductQuotaDimensionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProductQuotaDimensionRequest) GoString() string {
	return s.String()
}

func (s *GetProductQuotaDimensionRequest) SetDependentDimensions(v []*GetProductQuotaDimensionRequestDependentDimensions) *GetProductQuotaDimensionRequest {
	s.DependentDimensions = v
	return s
}

func (s *GetProductQuotaDimensionRequest) SetDimensionKey(v string) *GetProductQuotaDimensionRequest {
	s.DimensionKey = &v
	return s
}

func (s *GetProductQuotaDimensionRequest) SetProductCode(v string) *GetProductQuotaDimensionRequest {
	s.ProductCode = &v
	return s
}

type GetProductQuotaDimensionRequestDependentDimensions struct {
	// The key of the quota dimension on which the quota dimension that you want to query is dependent.
	//
	// > The value range of N varies based on the number of quota dimensions that are supported by the specified Alibaba Cloud service.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the quota dimension on which the quota dimension that you want to query is dependent.
	//
	// > The value range of N varies based on the number of quota dimensions that are supported by the specified Alibaba Cloud service.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetProductQuotaDimensionRequestDependentDimensions) String() string {
	return tea.Prettify(s)
}

func (s GetProductQuotaDimensionRequestDependentDimensions) GoString() string {
	return s.String()
}

func (s *GetProductQuotaDimensionRequestDependentDimensions) SetKey(v string) *GetProductQuotaDimensionRequestDependentDimensions {
	s.Key = &v
	return s
}

func (s *GetProductQuotaDimensionRequestDependentDimensions) SetValue(v string) *GetProductQuotaDimensionRequestDependentDimensions {
	s.Value = &v
	return s
}

type GetProductQuotaDimensionResponseBody struct {
	// The details of the quota dimension.
	QuotaDimension *GetProductQuotaDimensionResponseBodyQuotaDimension `json:"QuotaDimension,omitempty" xml:"QuotaDimension,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1FA5F0E2-368E-4BA4-A8D0-6060FC6BB8F3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProductQuotaDimensionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProductQuotaDimensionResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductQuotaDimensionResponseBody) SetQuotaDimension(v *GetProductQuotaDimensionResponseBodyQuotaDimension) *GetProductQuotaDimensionResponseBody {
	s.QuotaDimension = v
	return s
}

func (s *GetProductQuotaDimensionResponseBody) SetRequestId(v string) *GetProductQuotaDimensionResponseBody {
	s.RequestId = &v
	return s
}

type GetProductQuotaDimensionResponseBodyQuotaDimension struct {
	// The quota dimensions on which the quota dimension that you want to query is dependent.
	DependentDimensions []*string `json:"DependentDimensions,omitempty" xml:"DependentDimensions,omitempty" type:"Repeated"`
	// The key of the quota dimension. Valid values:
	//
	// 	- regionId: the region ID.
	//
	// 	- zoneId: the zone ID.
	//
	// 	- chargeType: the billing method.
	//
	// 	- networkType: the network type.
	//
	// example:
	//
	// regionId
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The details of the quota dimension value.
	DimensionValueDetail []*GetProductQuotaDimensionResponseBodyQuotaDimensionDimensionValueDetail `json:"DimensionValueDetail,omitempty" xml:"DimensionValueDetail,omitempty" type:"Repeated"`
	// The values of the quota dimension.
	DimensionValues []*string `json:"DimensionValues,omitempty" xml:"DimensionValues,omitempty" type:"Repeated"`
	// The name of the quota dimension.
	//
	// example:
	//
	// region
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetProductQuotaDimensionResponseBodyQuotaDimension) String() string {
	return tea.Prettify(s)
}

func (s GetProductQuotaDimensionResponseBodyQuotaDimension) GoString() string {
	return s.String()
}

func (s *GetProductQuotaDimensionResponseBodyQuotaDimension) SetDependentDimensions(v []*string) *GetProductQuotaDimensionResponseBodyQuotaDimension {
	s.DependentDimensions = v
	return s
}

func (s *GetProductQuotaDimensionResponseBodyQuotaDimension) SetDimensionKey(v string) *GetProductQuotaDimensionResponseBodyQuotaDimension {
	s.DimensionKey = &v
	return s
}

func (s *GetProductQuotaDimensionResponseBodyQuotaDimension) SetDimensionValueDetail(v []*GetProductQuotaDimensionResponseBodyQuotaDimensionDimensionValueDetail) *GetProductQuotaDimensionResponseBodyQuotaDimension {
	s.DimensionValueDetail = v
	return s
}

func (s *GetProductQuotaDimensionResponseBodyQuotaDimension) SetDimensionValues(v []*string) *GetProductQuotaDimensionResponseBodyQuotaDimension {
	s.DimensionValues = v
	return s
}

func (s *GetProductQuotaDimensionResponseBodyQuotaDimension) SetName(v string) *GetProductQuotaDimensionResponseBodyQuotaDimension {
	s.Name = &v
	return s
}

type GetProductQuotaDimensionResponseBodyQuotaDimensionDimensionValueDetail struct {
	// The name of the quota dimension.
	//
	// example:
	//
	// cn-hangzhou
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the quota dimension.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetProductQuotaDimensionResponseBodyQuotaDimensionDimensionValueDetail) String() string {
	return tea.Prettify(s)
}

func (s GetProductQuotaDimensionResponseBodyQuotaDimensionDimensionValueDetail) GoString() string {
	return s.String()
}

func (s *GetProductQuotaDimensionResponseBodyQuotaDimensionDimensionValueDetail) SetName(v string) *GetProductQuotaDimensionResponseBodyQuotaDimensionDimensionValueDetail {
	s.Name = &v
	return s
}

func (s *GetProductQuotaDimensionResponseBodyQuotaDimensionDimensionValueDetail) SetValue(v string) *GetProductQuotaDimensionResponseBodyQuotaDimensionDimensionValueDetail {
	s.Value = &v
	return s
}

type GetProductQuotaDimensionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProductQuotaDimensionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProductQuotaDimensionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProductQuotaDimensionResponse) GoString() string {
	return s.String()
}

func (s *GetProductQuotaDimensionResponse) SetHeaders(v map[string]*string) *GetProductQuotaDimensionResponse {
	s.Headers = v
	return s
}

func (s *GetProductQuotaDimensionResponse) SetStatusCode(v int32) *GetProductQuotaDimensionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProductQuotaDimensionResponse) SetBody(v *GetProductQuotaDimensionResponseBody) *GetProductQuotaDimensionResponse {
	s.Body = v
	return s
}

type GetQuotaAlarmRequest struct {
	// The ID of the quota alert.
	//
	// For more information about how to query the ID of a quota alert, see [ListQuotaAlarms](https://help.aliyun.com/document_detail/184348.html).
	//
	// example:
	//
	// 78d7e436-4b25-4897-84b5-d7b656bb****
	AlarmId *string `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
}

func (s GetQuotaAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaAlarmRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaAlarmRequest) SetAlarmId(v string) *GetQuotaAlarmRequest {
	s.AlarmId = &v
	return s
}

type GetQuotaAlarmResponseBody struct {
	// The details of the quota alert.
	QuotaAlarm *GetQuotaAlarmResponseBodyQuotaAlarm `json:"QuotaAlarm,omitempty" xml:"QuotaAlarm,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 81B9D511-F3DD-43B1-9A81-1795DDB52ADF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetQuotaAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaAlarmResponseBody) SetQuotaAlarm(v *GetQuotaAlarmResponseBodyQuotaAlarm) *GetQuotaAlarmResponseBody {
	s.QuotaAlarm = v
	return s
}

func (s *GetQuotaAlarmResponseBody) SetRequestId(v string) *GetQuotaAlarmResponseBody {
	s.RequestId = &v
	return s
}

type GetQuotaAlarmResponseBodyQuotaAlarm struct {
	// The ID of the quota alert.
	//
	// example:
	//
	// 78d7e436-4b25-4897-84b5-d7b656bb****
	AlarmId *string `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// The name of the quota alert.
	//
	// example:
	//
	// tf-testacccn-hangzhouquotasquotaalarm81611
	AlarmName *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	// The time when the quota alert was created.
	//
	// example:
	//
	// 2021-01-21T03:47:28Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The alert notification methods.
	NotifyChannels []*string `json:"NotifyChannels,omitempty" xml:"NotifyChannels,omitempty" type:"Repeated"`
	// The alert contact.
	//
	// example:
	//
	// accountContact
	NotifyTarget *string `json:"NotifyTarget,omitempty" xml:"NotifyTarget,omitempty"`
	// The abbreviation of the cloud service name.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	//
	// example:
	//
	// q_security-groups
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The quota dimension.
	//
	// example:
	//
	// {"regionId":"cn-hangzhou"}
	QuotaDimension map[string]interface{} `json:"QuotaDimension,omitempty" xml:"QuotaDimension,omitempty"`
	// The used quota.
	//
	// example:
	//
	// 28
	QuotaUsage *float32 `json:"QuotaUsage,omitempty" xml:"QuotaUsage,omitempty"`
	// The quota value.
	//
	// example:
	//
	// 804
	QuotaValue *float32 `json:"QuotaValue,omitempty" xml:"QuotaValue,omitempty"`
	// The numeric value of the alert threshold.
	//
	// example:
	//
	// 29
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The percentage of the alert threshold.
	//
	// example:
	//
	// 50
	ThresholdPercent *float32 `json:"ThresholdPercent,omitempty" xml:"ThresholdPercent,omitempty"`
	// The type of the quota alert. Valid values:
	//
	// 	- used: The alert is created for the used quota.
	//
	// 	- usable: The alert is created for the available quota.
	//
	// example:
	//
	// used
	ThresholdType *string `json:"ThresholdType,omitempty" xml:"ThresholdType,omitempty"`
	// The webhook URL. Quota Center sends alert notifications to the specified URL by using HTTP POST requests.
	//
	// example:
	//
	// https://alert.aliyun.com/callback
	Webhook *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s GetQuotaAlarmResponseBodyQuotaAlarm) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaAlarmResponseBodyQuotaAlarm) GoString() string {
	return s.String()
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetAlarmId(v string) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.AlarmId = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetAlarmName(v string) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.AlarmName = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetCreateTime(v string) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.CreateTime = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetNotifyChannels(v []*string) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.NotifyChannels = v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetNotifyTarget(v string) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.NotifyTarget = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetProductCode(v string) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.ProductCode = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetQuotaActionCode(v string) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.QuotaActionCode = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetQuotaDimension(v map[string]interface{}) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.QuotaDimension = v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetQuotaUsage(v float32) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.QuotaUsage = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetQuotaValue(v float32) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.QuotaValue = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetThreshold(v float32) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.Threshold = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetThresholdPercent(v float32) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.ThresholdPercent = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetThresholdType(v string) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.ThresholdType = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetWebhook(v string) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.Webhook = &v
	return s
}

type GetQuotaAlarmResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQuotaAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQuotaAlarmResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaAlarmResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaAlarmResponse) SetHeaders(v map[string]*string) *GetQuotaAlarmResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaAlarmResponse) SetStatusCode(v int32) *GetQuotaAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQuotaAlarmResponse) SetBody(v *GetQuotaAlarmResponseBody) *GetQuotaAlarmResponse {
	s.Body = v
	return s
}

type GetQuotaApplicationRequest struct {
	// The ID of the application.
	//
	// example:
	//
	// d314d6ae-867d-484c-9009-3d421a80****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s GetQuotaApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaApplicationRequest) SetApplicationId(v string) *GetQuotaApplicationRequest {
	s.ApplicationId = &v
	return s
}

type GetQuotaApplicationResponseBody struct {
	// The details about the application.
	QuotaApplication *GetQuotaApplicationResponseBodyQuotaApplication `json:"QuotaApplication,omitempty" xml:"QuotaApplication,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 7BBD1D37-094C-4485-8B7D-64682F82BC18
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetQuotaApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaApplicationResponseBody) SetQuotaApplication(v *GetQuotaApplicationResponseBodyQuotaApplication) *GetQuotaApplicationResponseBody {
	s.QuotaApplication = v
	return s
}

func (s *GetQuotaApplicationResponseBody) SetRequestId(v string) *GetQuotaApplicationResponseBody {
	s.RequestId = &v
	return s
}

type GetQuotaApplicationResponseBodyQuotaApplication struct {
	// The ID of the application.
	//
	// example:
	//
	// d314d6ae-867d-484c-9009-3d421a80****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The time when the application was submitted.
	//
	// example:
	//
	// 2021-01-19T09:25:56Z
	ApplyTime *string `json:"ApplyTime,omitempty" xml:"ApplyTime,omitempty"`
	// The approved quota value.
	//
	// example:
	//
	// 10
	ApproveValue *float32 `json:"ApproveValue,omitempty" xml:"ApproveValue,omitempty"`
	// The result of the application.
	//
	// example:
	//
	// Agree
	AuditReason *string `json:"AuditReason,omitempty" xml:"AuditReason,omitempty"`
	// The expected value of the quota.
	//
	// example:
	//
	// 804
	DesireValue *int32 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The dimension.
	//
	// Format: `{"regionId":"Region"}`.
	//
	// example:
	//
	// ["cn-shanghai","cn-hangzhou"]
	Dimension map[string]interface{} `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The time when the new quota value takes effect.
	//
	// example:
	//
	// 2021-01-19 15:30:00
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The time when the new quota expires.
	//
	// example:
	//
	// 2023-06-29 15:30:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The method of that is used to send alert notifications. Valid values:
	//
	// 	- 0: Quota Center does not send a notification.
	//
	// 	- 1: Quota Center sends an email notification.
	//
	// 	- 2: Quota Center sends an SMS notification.
	//
	// example:
	//
	// 0
	NoticeType *int64 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	//
	// example:
	//
	// q_security-groups
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the quota.
	//
	// example:
	//
	// acs:quotas:cn-hangzhou:120886317861****:quota/ecs/q_security-groups/
	QuotaArn *string `json:"QuotaArn,omitempty" xml:"QuotaArn,omitempty"`
	// The description of the quota.
	//
	// example:
	//
	// The maximum number of security groups that can be owned by the current account
	QuotaDescription *string `json:"QuotaDescription,omitempty" xml:"QuotaDescription,omitempty"`
	// The name of the quota.
	//
	// example:
	//
	// Maximum Number of Security Groups
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The unit of the new quota value.
	//
	// example:
	//
	// Count
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// The reason for the application.
	//
	// example:
	//
	// Scale Out
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The status of the application. Valid values:
	//
	// 	- Disagree: The application is rejected.
	//
	// 	- Agree: The application is approved.
	//
	// 	- Process: The application is being reviewed.
	//
	// 	- Cancel: The application is closed.
	//
	// example:
	//
	// Agree
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetQuotaApplicationResponseBodyQuotaApplication) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaApplicationResponseBodyQuotaApplication) GoString() string {
	return s.String()
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetApplicationId(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.ApplicationId = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetApplyTime(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.ApplyTime = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetApproveValue(v float32) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.ApproveValue = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetAuditReason(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.AuditReason = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetDesireValue(v int32) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.DesireValue = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetDimension(v map[string]interface{}) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.Dimension = v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetEffectiveTime(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.EffectiveTime = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetExpireTime(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.ExpireTime = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetNoticeType(v int64) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.NoticeType = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetProductCode(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.ProductCode = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetQuotaActionCode(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.QuotaActionCode = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetQuotaArn(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.QuotaArn = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetQuotaDescription(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.QuotaDescription = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetQuotaName(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.QuotaName = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetQuotaUnit(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.QuotaUnit = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetReason(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.Reason = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetStatus(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.Status = &v
	return s
}

type GetQuotaApplicationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQuotaApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQuotaApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaApplicationResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaApplicationResponse) SetHeaders(v map[string]*string) *GetQuotaApplicationResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaApplicationResponse) SetStatusCode(v int32) *GetQuotaApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQuotaApplicationResponse) SetBody(v *GetQuotaApplicationResponseBody) *GetQuotaApplicationResponse {
	s.Body = v
	return s
}

type GetQuotaApplicationApprovalRequest struct {
	// The quota application ID.
	//
	// For more information about how to obtain the ID of a quota application, see [ListQuotaApplications](https://help.aliyun.com/document_detail/440568.html).
	//
	// example:
	//
	// d314d6ae-867d-484c-9009-3d42****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s GetQuotaApplicationApprovalRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaApplicationApprovalRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaApplicationApprovalRequest) SetApplicationId(v string) *GetQuotaApplicationApprovalRequest {
	s.ApplicationId = &v
	return s
}

type GetQuotaApplicationApprovalResponseBody struct {
	// Indicates whether retries are allowed. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// false
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// PARAMETER.ILLEGALL
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// Parameter[%s] is required.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The parameters whose values are invalid.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// RAM.PERMISSION.DENIED
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// You are not authorized to do this action or the API input parameter is invalid.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The information about quota application approval.
	Module *GetQuotaApplicationApprovalResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7ED584FB-ECBF-4A2A-969D-F54D25EFABF9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQuotaApplicationApprovalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaApplicationApprovalResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaApplicationApprovalResponseBody) SetAllowRetry(v bool) *GetQuotaApplicationApprovalResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetQuotaApplicationApprovalResponseBody) SetDynamicCode(v string) *GetQuotaApplicationApprovalResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetQuotaApplicationApprovalResponseBody) SetDynamicMessage(v string) *GetQuotaApplicationApprovalResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetQuotaApplicationApprovalResponseBody) SetErrorArgs(v []interface{}) *GetQuotaApplicationApprovalResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetQuotaApplicationApprovalResponseBody) SetErrorCode(v string) *GetQuotaApplicationApprovalResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetQuotaApplicationApprovalResponseBody) SetErrorMsg(v string) *GetQuotaApplicationApprovalResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetQuotaApplicationApprovalResponseBody) SetHttpStatusCode(v int32) *GetQuotaApplicationApprovalResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetQuotaApplicationApprovalResponseBody) SetModule(v *GetQuotaApplicationApprovalResponseBodyModule) *GetQuotaApplicationApprovalResponseBody {
	s.Module = v
	return s
}

func (s *GetQuotaApplicationApprovalResponseBody) SetRequestId(v string) *GetQuotaApplicationApprovalResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQuotaApplicationApprovalResponseBody) SetSuccess(v bool) *GetQuotaApplicationApprovalResponseBody {
	s.Success = &v
	return s
}

type GetQuotaApplicationApprovalResponseBodyModule struct {
	// The average amount of time required to approve quota applications. Unit: hour.
	//
	// example:
	//
	// 24
	ApprovalHours *int32 `json:"ApprovalHours,omitempty" xml:"ApprovalHours,omitempty"`
	// The interval between two consecutive approval reminders. Unit: hour.
	//
	// example:
	//
	// 24
	ReminderIntervalHours *int32 `json:"ReminderIntervalHours,omitempty" xml:"ReminderIntervalHours,omitempty"`
	// Indicates whether approval reminders are supported for quota applications. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// false
	SupportReminder *bool `json:"SupportReminder,omitempty" xml:"SupportReminder,omitempty"`
	// The reason why approval reminders are not supported for quota applications.
	//
	// example:
	//
	// can only be remind once within the interval
	UnsupportReminderReason *string `json:"UnsupportReminderReason,omitempty" xml:"UnsupportReminderReason,omitempty"`
}

func (s GetQuotaApplicationApprovalResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaApplicationApprovalResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetQuotaApplicationApprovalResponseBodyModule) SetApprovalHours(v int32) *GetQuotaApplicationApprovalResponseBodyModule {
	s.ApprovalHours = &v
	return s
}

func (s *GetQuotaApplicationApprovalResponseBodyModule) SetReminderIntervalHours(v int32) *GetQuotaApplicationApprovalResponseBodyModule {
	s.ReminderIntervalHours = &v
	return s
}

func (s *GetQuotaApplicationApprovalResponseBodyModule) SetSupportReminder(v bool) *GetQuotaApplicationApprovalResponseBodyModule {
	s.SupportReminder = &v
	return s
}

func (s *GetQuotaApplicationApprovalResponseBodyModule) SetUnsupportReminderReason(v string) *GetQuotaApplicationApprovalResponseBodyModule {
	s.UnsupportReminderReason = &v
	return s
}

type GetQuotaApplicationApprovalResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQuotaApplicationApprovalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQuotaApplicationApprovalResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaApplicationApprovalResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaApplicationApprovalResponse) SetHeaders(v map[string]*string) *GetQuotaApplicationApprovalResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaApplicationApprovalResponse) SetStatusCode(v int32) *GetQuotaApplicationApprovalResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQuotaApplicationApprovalResponse) SetBody(v *GetQuotaApplicationApprovalResponseBody) *GetQuotaApplicationApprovalResponse {
	s.Body = v
	return s
}

type GetQuotaTemplateServiceStatusRequest struct {
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-pG****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
}

func (s GetQuotaTemplateServiceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaTemplateServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaTemplateServiceStatusRequest) SetResourceDirectoryId(v string) *GetQuotaTemplateServiceStatusRequest {
	s.ResourceDirectoryId = &v
	return s
}

type GetQuotaTemplateServiceStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D47B3A10-CDAC-5412-B2EE-EC9A3DBE9053
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the quota template.
	TemplateServiceStatus *GetQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus `json:"TemplateServiceStatus,omitempty" xml:"TemplateServiceStatus,omitempty" type:"Struct"`
}

func (s GetQuotaTemplateServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaTemplateServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaTemplateServiceStatusResponseBody) SetRequestId(v string) *GetQuotaTemplateServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQuotaTemplateServiceStatusResponseBody) SetTemplateServiceStatus(v *GetQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) *GetQuotaTemplateServiceStatusResponseBody {
	s.TemplateServiceStatus = v
	return s
}

type GetQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus struct {
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-pG****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The state of the quota template. Valid values:
	//
	// 	- \\-1: The quota template is disabled.
	//
	// 	- 1: The quota template is enabled.
	//
	// example:
	//
	// 1
	ServiceStatus *int32 `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
}

func (s GetQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) GoString() string {
	return s.String()
}

func (s *GetQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) SetResourceDirectoryId(v string) *GetQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus {
	s.ResourceDirectoryId = &v
	return s
}

func (s *GetQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) SetServiceStatus(v int32) *GetQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus {
	s.ServiceStatus = &v
	return s
}

type GetQuotaTemplateServiceStatusResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQuotaTemplateServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQuotaTemplateServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaTemplateServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaTemplateServiceStatusResponse) SetHeaders(v map[string]*string) *GetQuotaTemplateServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaTemplateServiceStatusResponse) SetStatusCode(v int32) *GetQuotaTemplateServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQuotaTemplateServiceStatusResponse) SetBody(v *GetQuotaTemplateServiceStatusResponseBody) *GetQuotaTemplateServiceStatusResponse {
	s.Body = v
	return s
}

type ListAlarmHistoriesRequest struct {
	// The ID of the alert.
	//
	// example:
	//
	// 18b3be23-b7b0-4d45-91bc-d0c331aa****
	AlarmId *string `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// The end of the time range to query.
	//
	// example:
	//
	// 20201024
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The keyword that is used for the query.
	//
	// example:
	//
	// Quantity
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The maximum number of records that can be returned for the query.
	//
	// Valid values: 1 to 200. Default value: 30.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position from which you want to start the query. If you leave this parameter empty, the query starts from the beginning.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// > For more information, see [Alibaba Cloud services that support Quota Center](https://help.aliyun.com/document_detail/182368.html).
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The beginning of the time range to query.
	//
	// example:
	//
	// 20201020
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListAlarmHistoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmHistoriesRequest) GoString() string {
	return s.String()
}

func (s *ListAlarmHistoriesRequest) SetAlarmId(v string) *ListAlarmHistoriesRequest {
	s.AlarmId = &v
	return s
}

func (s *ListAlarmHistoriesRequest) SetEndTime(v int64) *ListAlarmHistoriesRequest {
	s.EndTime = &v
	return s
}

func (s *ListAlarmHistoriesRequest) SetKeyword(v string) *ListAlarmHistoriesRequest {
	s.Keyword = &v
	return s
}

func (s *ListAlarmHistoriesRequest) SetMaxResults(v int32) *ListAlarmHistoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAlarmHistoriesRequest) SetNextToken(v string) *ListAlarmHistoriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAlarmHistoriesRequest) SetProductCode(v string) *ListAlarmHistoriesRequest {
	s.ProductCode = &v
	return s
}

func (s *ListAlarmHistoriesRequest) SetStartTime(v int64) *ListAlarmHistoriesRequest {
	s.StartTime = &v
	return s
}

type ListAlarmHistoriesResponseBody struct {
	// The details of the quota alert records.
	AlarmHistories []*ListAlarmHistoriesResponseBodyAlarmHistories `json:"AlarmHistories,omitempty" xml:"AlarmHistories,omitempty" type:"Repeated"`
	// The maximum number of records that are returned for the query.
	//
	// example:
	//
	// 4
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position at which the query ends. An empty value indicates that all data is returned.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CB38DDF9-B1E0-48C1-9966-19C443C2841E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records that are returned for the query.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAlarmHistoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmHistoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlarmHistoriesResponseBody) SetAlarmHistories(v []*ListAlarmHistoriesResponseBodyAlarmHistories) *ListAlarmHistoriesResponseBody {
	s.AlarmHistories = v
	return s
}

func (s *ListAlarmHistoriesResponseBody) SetMaxResults(v int32) *ListAlarmHistoriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAlarmHistoriesResponseBody) SetNextToken(v string) *ListAlarmHistoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAlarmHistoriesResponseBody) SetRequestId(v string) *ListAlarmHistoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlarmHistoriesResponseBody) SetTotalCount(v int32) *ListAlarmHistoriesResponseBody {
	s.TotalCount = &v
	return s
}

type ListAlarmHistoriesResponseBodyAlarmHistories struct {
	// The name of the quota alert.
	//
	// example:
	//
	// security_groups
	AlarmName *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	// The time when the quota alert rule was created.
	//
	// example:
	//
	// 2021-01-24T09:20:09Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The notification methods of the quota alert.
	NotifyChannels []*string `json:"NotifyChannels,omitempty" xml:"NotifyChannels,omitempty" type:"Repeated"`
	// The quota alert contact.
	//
	// example:
	//
	// accountContact
	NotifyTarget *string `json:"NotifyTarget,omitempty" xml:"NotifyTarget,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// q_security-groups
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The used quota.
	//
	// example:
	//
	// 31
	QuotaUsage *float32 `json:"QuotaUsage,omitempty" xml:"QuotaUsage,omitempty"`
	// The threshold to trigger quota alerts.
	//
	// example:
	//
	// 29
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The percentage of the quota alert threshold.
	//
	// example:
	//
	// 80
	ThresholdPercent *float32 `json:"ThresholdPercent,omitempty" xml:"ThresholdPercent,omitempty"`
}

func (s ListAlarmHistoriesResponseBodyAlarmHistories) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmHistoriesResponseBodyAlarmHistories) GoString() string {
	return s.String()
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) SetAlarmName(v string) *ListAlarmHistoriesResponseBodyAlarmHistories {
	s.AlarmName = &v
	return s
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) SetCreateTime(v string) *ListAlarmHistoriesResponseBodyAlarmHistories {
	s.CreateTime = &v
	return s
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) SetNotifyChannels(v []*string) *ListAlarmHistoriesResponseBodyAlarmHistories {
	s.NotifyChannels = v
	return s
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) SetNotifyTarget(v string) *ListAlarmHistoriesResponseBodyAlarmHistories {
	s.NotifyTarget = &v
	return s
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) SetProductCode(v string) *ListAlarmHistoriesResponseBodyAlarmHistories {
	s.ProductCode = &v
	return s
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) SetQuotaActionCode(v string) *ListAlarmHistoriesResponseBodyAlarmHistories {
	s.QuotaActionCode = &v
	return s
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) SetQuotaUsage(v float32) *ListAlarmHistoriesResponseBodyAlarmHistories {
	s.QuotaUsage = &v
	return s
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) SetThreshold(v float32) *ListAlarmHistoriesResponseBodyAlarmHistories {
	s.Threshold = &v
	return s
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) SetThresholdPercent(v float32) *ListAlarmHistoriesResponseBodyAlarmHistories {
	s.ThresholdPercent = &v
	return s
}

type ListAlarmHistoriesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlarmHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlarmHistoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmHistoriesResponse) GoString() string {
	return s.String()
}

func (s *ListAlarmHistoriesResponse) SetHeaders(v map[string]*string) *ListAlarmHistoriesResponse {
	s.Headers = v
	return s
}

func (s *ListAlarmHistoriesResponse) SetStatusCode(v int32) *ListAlarmHistoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlarmHistoriesResponse) SetBody(v *ListAlarmHistoriesResponseBody) *ListAlarmHistoriesResponse {
	s.Body = v
	return s
}

type ListDependentQuotasRequest struct {
	// The abbreviation of the Alibaba Cloud service name.
	//
	// > For more information, see [Alibaba Cloud services that support Quota Center](https://help.aliyun.com/document_detail/182368.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// csk
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// q_i5uzm3
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
}

func (s ListDependentQuotasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDependentQuotasRequest) GoString() string {
	return s.String()
}

func (s *ListDependentQuotasRequest) SetProductCode(v string) *ListDependentQuotasRequest {
	s.ProductCode = &v
	return s
}

func (s *ListDependentQuotasRequest) SetQuotaActionCode(v string) *ListDependentQuotasRequest {
	s.QuotaActionCode = &v
	return s
}

type ListDependentQuotasResponseBody struct {
	// The quotas on which the specified quota depends.
	Quotas []*ListDependentQuotasResponseBodyQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 920D8A47-26BB-49FA-A09F-F98D7DAA55F3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDependentQuotasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDependentQuotasResponseBody) GoString() string {
	return s.String()
}

func (s *ListDependentQuotasResponseBody) SetQuotas(v []*ListDependentQuotasResponseBodyQuotas) *ListDependentQuotasResponseBody {
	s.Quotas = v
	return s
}

func (s *ListDependentQuotasResponseBody) SetRequestId(v string) *ListDependentQuotasResponseBody {
	s.RequestId = &v
	return s
}

type ListDependentQuotasResponseBodyQuotas struct {
	// The dimensions of the quotas on which the specified quota depends.
	Dimensions []*ListDependentQuotasResponseBodyQuotasDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// q_elastic-network-interfaces
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The relationship percentage between the specified quota and the quotas on which the specified quota depends.
	//
	// example:
	//
	// 50
	Scale *float32 `json:"Scale,omitempty" xml:"Scale,omitempty"`
}

func (s ListDependentQuotasResponseBodyQuotas) String() string {
	return tea.Prettify(s)
}

func (s ListDependentQuotasResponseBodyQuotas) GoString() string {
	return s.String()
}

func (s *ListDependentQuotasResponseBodyQuotas) SetDimensions(v []*ListDependentQuotasResponseBodyQuotasDimensions) *ListDependentQuotasResponseBodyQuotas {
	s.Dimensions = v
	return s
}

func (s *ListDependentQuotasResponseBodyQuotas) SetProductCode(v string) *ListDependentQuotasResponseBodyQuotas {
	s.ProductCode = &v
	return s
}

func (s *ListDependentQuotasResponseBodyQuotas) SetQuotaActionCode(v string) *ListDependentQuotasResponseBodyQuotas {
	s.QuotaActionCode = &v
	return s
}

func (s *ListDependentQuotasResponseBodyQuotas) SetScale(v float32) *ListDependentQuotasResponseBodyQuotas {
	s.Scale = &v
	return s
}

type ListDependentQuotasResponseBodyQuotasDimensions struct {
	// The dimensions of the quotas on which the specified quota depends.
	DependentDimension []*string `json:"DependentDimension,omitempty" xml:"DependentDimension,omitempty" type:"Repeated"`
	// The key of the quota dimension.
	//
	// example:
	//
	// regionId
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The dimension values.
	DimensionValues []*string `json:"DimensionValues,omitempty" xml:"DimensionValues,omitempty" type:"Repeated"`
}

func (s ListDependentQuotasResponseBodyQuotasDimensions) String() string {
	return tea.Prettify(s)
}

func (s ListDependentQuotasResponseBodyQuotasDimensions) GoString() string {
	return s.String()
}

func (s *ListDependentQuotasResponseBodyQuotasDimensions) SetDependentDimension(v []*string) *ListDependentQuotasResponseBodyQuotasDimensions {
	s.DependentDimension = v
	return s
}

func (s *ListDependentQuotasResponseBodyQuotasDimensions) SetDimensionKey(v string) *ListDependentQuotasResponseBodyQuotasDimensions {
	s.DimensionKey = &v
	return s
}

func (s *ListDependentQuotasResponseBodyQuotasDimensions) SetDimensionValues(v []*string) *ListDependentQuotasResponseBodyQuotasDimensions {
	s.DimensionValues = v
	return s
}

type ListDependentQuotasResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDependentQuotasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDependentQuotasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDependentQuotasResponse) GoString() string {
	return s.String()
}

func (s *ListDependentQuotasResponse) SetHeaders(v map[string]*string) *ListDependentQuotasResponse {
	s.Headers = v
	return s
}

func (s *ListDependentQuotasResponse) SetStatusCode(v int32) *ListDependentQuotasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDependentQuotasResponse) SetBody(v *ListDependentQuotasResponseBody) *ListDependentQuotasResponse {
	s.Body = v
	return s
}

type ListProductDimensionGroupsRequest struct {
	// The maximum number of records that can be returned for the query. Valid values: 1 to 200. Default value: 30.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position from which you want to start the query. If you leave this parameter empty, the query starts from the beginning.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The service code.
	//
	// > For more information, see [Alibaba Cloud services that support Quota Center](https://help.aliyun.com/document_detail/182368.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// oss
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s ListProductDimensionGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductDimensionGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListProductDimensionGroupsRequest) SetMaxResults(v int32) *ListProductDimensionGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProductDimensionGroupsRequest) SetNextToken(v string) *ListProductDimensionGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListProductDimensionGroupsRequest) SetProductCode(v string) *ListProductDimensionGroupsRequest {
	s.ProductCode = &v
	return s
}

type ListProductDimensionGroupsResponseBody struct {
	// The dimension groups.
	DimensionGroups []*ListProductDimensionGroupsResponseBodyDimensionGroups `json:"DimensionGroups,omitempty" xml:"DimensionGroups,omitempty" type:"Repeated"`
	// The maximum number of records that are returned for the query.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position at which the query ends. An empty value indicates that all data is returned.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 057D210F-F2FC-5329-A536-26C16628BB09
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records that are returned for the query.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProductDimensionGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductDimensionGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductDimensionGroupsResponseBody) SetDimensionGroups(v []*ListProductDimensionGroupsResponseBodyDimensionGroups) *ListProductDimensionGroupsResponseBody {
	s.DimensionGroups = v
	return s
}

func (s *ListProductDimensionGroupsResponseBody) SetMaxResults(v int32) *ListProductDimensionGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListProductDimensionGroupsResponseBody) SetNextToken(v string) *ListProductDimensionGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProductDimensionGroupsResponseBody) SetRequestId(v string) *ListProductDimensionGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductDimensionGroupsResponseBody) SetTotalCount(v int32) *ListProductDimensionGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListProductDimensionGroupsResponseBodyDimensionGroups struct {
	// The key of the dimension group.
	DimensionKeys []*string `json:"DimensionKeys,omitempty" xml:"DimensionKeys,omitempty" type:"Repeated"`
	// The code of the dimension group.
	//
	// example:
	//
	// oss_wf1ngqmd7q
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
	// The name of the dimension group.
	//
	// example:
	//
	// OSS_Group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The service code.
	//
	// example:
	//
	// oss
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s ListProductDimensionGroupsResponseBodyDimensionGroups) String() string {
	return tea.Prettify(s)
}

func (s ListProductDimensionGroupsResponseBodyDimensionGroups) GoString() string {
	return s.String()
}

func (s *ListProductDimensionGroupsResponseBodyDimensionGroups) SetDimensionKeys(v []*string) *ListProductDimensionGroupsResponseBodyDimensionGroups {
	s.DimensionKeys = v
	return s
}

func (s *ListProductDimensionGroupsResponseBodyDimensionGroups) SetGroupCode(v string) *ListProductDimensionGroupsResponseBodyDimensionGroups {
	s.GroupCode = &v
	return s
}

func (s *ListProductDimensionGroupsResponseBodyDimensionGroups) SetGroupName(v string) *ListProductDimensionGroupsResponseBodyDimensionGroups {
	s.GroupName = &v
	return s
}

func (s *ListProductDimensionGroupsResponseBodyDimensionGroups) SetProductCode(v string) *ListProductDimensionGroupsResponseBodyDimensionGroups {
	s.ProductCode = &v
	return s
}

type ListProductDimensionGroupsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProductDimensionGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProductDimensionGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProductDimensionGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListProductDimensionGroupsResponse) SetHeaders(v map[string]*string) *ListProductDimensionGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListProductDimensionGroupsResponse) SetStatusCode(v int32) *ListProductDimensionGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductDimensionGroupsResponse) SetBody(v *ListProductDimensionGroupsResponseBody) *ListProductDimensionGroupsResponse {
	s.Body = v
	return s
}

type ListProductQuotaDimensionsRequest struct {
	// The maximum number of records that can be returned for the query.
	//
	// Valid values: 1 to 200. Default value: 30.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position from which you want to start the query. If you leave this parameter empty, the query starts from the beginning.
	//
	// example:
	//
	// 0
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// > For more information, see [Alibaba Cloud services that support Quota Center](https://help.aliyun.com/document_detail/182368.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The type of the quota. Valid values:
	//
	// 	- FlowControl: API rate limit.
	//
	// 	- CommonQuota: general quota. This is the default value.
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
}

func (s ListProductQuotaDimensionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductQuotaDimensionsRequest) GoString() string {
	return s.String()
}

func (s *ListProductQuotaDimensionsRequest) SetMaxResults(v int32) *ListProductQuotaDimensionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProductQuotaDimensionsRequest) SetNextToken(v string) *ListProductQuotaDimensionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListProductQuotaDimensionsRequest) SetProductCode(v string) *ListProductQuotaDimensionsRequest {
	s.ProductCode = &v
	return s
}

func (s *ListProductQuotaDimensionsRequest) SetQuotaCategory(v string) *ListProductQuotaDimensionsRequest {
	s.QuotaCategory = &v
	return s
}

type ListProductQuotaDimensionsResponseBody struct {
	// The maximum number of records that are returned for the query.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position at which the query ends. An empty value indicates that all data is returned.
	//
	// example:
	//
	// 10
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The quota dimensions.
	QuotaDimensions []*ListProductQuotaDimensionsResponseBodyQuotaDimensions `json:"QuotaDimensions,omitempty" xml:"QuotaDimensions,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 7ED584FB-ECBF-4A2A-969D-F54D25EFABF9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records that are returned for the query.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProductQuotaDimensionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductQuotaDimensionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductQuotaDimensionsResponseBody) SetMaxResults(v int32) *ListProductQuotaDimensionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListProductQuotaDimensionsResponseBody) SetNextToken(v string) *ListProductQuotaDimensionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProductQuotaDimensionsResponseBody) SetQuotaDimensions(v []*ListProductQuotaDimensionsResponseBodyQuotaDimensions) *ListProductQuotaDimensionsResponseBody {
	s.QuotaDimensions = v
	return s
}

func (s *ListProductQuotaDimensionsResponseBody) SetRequestId(v string) *ListProductQuotaDimensionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductQuotaDimensionsResponseBody) SetTotalCount(v int32) *ListProductQuotaDimensionsResponseBody {
	s.TotalCount = &v
	return s
}

type ListProductQuotaDimensionsResponseBodyQuotaDimensions struct {
	// The quota dimensions on which the quota dimension that you want to query is dependent.
	DependentDimensions []*string `json:"DependentDimensions,omitempty" xml:"DependentDimensions,omitempty" type:"Repeated"`
	// The key of the quota dimension. Valid values:
	//
	// 	- regionId: the region ID.
	//
	// 	- zoneId: the zone ID.
	//
	// 	- chargeType: the billing method.
	//
	// 	- networkType: the network type.
	//
	// example:
	//
	// zoneId
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The details about the dimension value.
	DimensionValueDetail []*ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail `json:"DimensionValueDetail,omitempty" xml:"DimensionValueDetail,omitempty" type:"Repeated"`
	// The dimension values.
	DimensionValues []*string `json:"DimensionValues,omitempty" xml:"DimensionValues,omitempty" type:"Repeated"`
	// The name of the quota dimension.
	//
	// example:
	//
	// Zone
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the quota dimension is required when you query quota dimensions. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Requisite *bool `json:"Requisite,omitempty" xml:"Requisite,omitempty"`
}

func (s ListProductQuotaDimensionsResponseBodyQuotaDimensions) String() string {
	return tea.Prettify(s)
}

func (s ListProductQuotaDimensionsResponseBodyQuotaDimensions) GoString() string {
	return s.String()
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensions) SetDependentDimensions(v []*string) *ListProductQuotaDimensionsResponseBodyQuotaDimensions {
	s.DependentDimensions = v
	return s
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensions) SetDimensionKey(v string) *ListProductQuotaDimensionsResponseBodyQuotaDimensions {
	s.DimensionKey = &v
	return s
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensions) SetDimensionValueDetail(v []*ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail) *ListProductQuotaDimensionsResponseBodyQuotaDimensions {
	s.DimensionValueDetail = v
	return s
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensions) SetDimensionValues(v []*string) *ListProductQuotaDimensionsResponseBodyQuotaDimensions {
	s.DimensionValues = v
	return s
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensions) SetName(v string) *ListProductQuotaDimensionsResponseBodyQuotaDimensions {
	s.Name = &v
	return s
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensions) SetRequisite(v bool) *ListProductQuotaDimensionsResponseBodyQuotaDimensions {
	s.Requisite = &v
	return s
}

type ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail struct {
	// The quota dimensions on which the quota dimension that you want to query is dependent.
	DependentDimensions []*ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetailDependentDimensions `json:"DependentDimensions,omitempty" xml:"DependentDimensions,omitempty" type:"Repeated"`
	// The name of the quota dimension.
	//
	// example:
	//
	// cn-hangzhou
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the quota dimension.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail) String() string {
	return tea.Prettify(s)
}

func (s ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail) GoString() string {
	return s.String()
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail) SetDependentDimensions(v []*ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetailDependentDimensions) *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail {
	s.DependentDimensions = v
	return s
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail) SetName(v string) *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail {
	s.Name = &v
	return s
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail) SetValue(v string) *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail {
	s.Value = &v
	return s
}

type ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetailDependentDimensions struct {
	// The key of the quota dimension on which the quota dimension that you want to query is dependent.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the quota dimension on which the quota dimension that you want to query is dependent.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetailDependentDimensions) String() string {
	return tea.Prettify(s)
}

func (s ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetailDependentDimensions) GoString() string {
	return s.String()
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetailDependentDimensions) SetKey(v string) *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetailDependentDimensions {
	s.Key = &v
	return s
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetailDependentDimensions) SetValue(v string) *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetailDependentDimensions {
	s.Value = &v
	return s
}

type ListProductQuotaDimensionsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProductQuotaDimensionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProductQuotaDimensionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProductQuotaDimensionsResponse) GoString() string {
	return s.String()
}

func (s *ListProductQuotaDimensionsResponse) SetHeaders(v map[string]*string) *ListProductQuotaDimensionsResponse {
	s.Headers = v
	return s
}

func (s *ListProductQuotaDimensionsResponse) SetStatusCode(v int32) *ListProductQuotaDimensionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductQuotaDimensionsResponse) SetBody(v *ListProductQuotaDimensionsResponseBody) *ListProductQuotaDimensionsResponse {
	s.Body = v
	return s
}

type ListProductQuotasRequest struct {
	// The quota dimensions.
	Dimensions []*ListProductQuotasRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The code of the dimension group.
	//
	// example:
	//
	// entconsole_w1j3msbo2g
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
	// The keyword that you want to use to search for the quotas.
	//
	// example:
	//
	// ecs-spec
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The maximum number of records that can be returned for the query.
	//
	// Valid values: 1 to 100. Default value: 30.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position from which you want to start the query. If you leave this parameter empty, the query starts from the beginning.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// >  To query the abbreviation of an Alibaba Cloud service name, call the [ListProducts](https://help.aliyun.com/document_detail/440555.html) operation and check the value of the `ProductCode` parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs-spec
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// ecs.g5.2xlarge
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The type of the quota. Valid values:
	//
	// 	- CommonQuota (default value): general quota
	//
	// 	- FlowControl: API rate limit
	//
	// 	- WhiteListLabel: whitelist quota
	//
	// example:
	//
	// FlowControl
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The field based on which you want to sort the returned records. Valid values:
	//
	// 	- TIME: The returned records are sorted by the last update time.
	//
	// 	- TOTAL: The returned records are sorted by the usage of the total quota.
	//
	// 	- RESERVED: The returned records are sorted by the usage of the reserved quota.
	//
	// >  This parameter is available only for quotas that belong to ECS Quotas by Instance Type. You can leave this parameter empty.
	//
	// example:
	//
	// TIME
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	// The order in which you want to sort the returned records. Valid values:
	//
	// 	- Ascending: ascending order
	//
	// 	- Descending: descending order
	//
	// >  This parameter is available only for quotas that belong to ECS Quotas by Instance Type. You can leave this parameter empty.
	//
	// example:
	//
	// Ascending
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListProductQuotasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductQuotasRequest) GoString() string {
	return s.String()
}

func (s *ListProductQuotasRequest) SetDimensions(v []*ListProductQuotasRequestDimensions) *ListProductQuotasRequest {
	s.Dimensions = v
	return s
}

func (s *ListProductQuotasRequest) SetGroupCode(v string) *ListProductQuotasRequest {
	s.GroupCode = &v
	return s
}

func (s *ListProductQuotasRequest) SetKeyWord(v string) *ListProductQuotasRequest {
	s.KeyWord = &v
	return s
}

func (s *ListProductQuotasRequest) SetMaxResults(v int32) *ListProductQuotasRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProductQuotasRequest) SetNextToken(v string) *ListProductQuotasRequest {
	s.NextToken = &v
	return s
}

func (s *ListProductQuotasRequest) SetProductCode(v string) *ListProductQuotasRequest {
	s.ProductCode = &v
	return s
}

func (s *ListProductQuotasRequest) SetQuotaActionCode(v string) *ListProductQuotasRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *ListProductQuotasRequest) SetQuotaCategory(v string) *ListProductQuotasRequest {
	s.QuotaCategory = &v
	return s
}

func (s *ListProductQuotasRequest) SetSortField(v string) *ListProductQuotasRequest {
	s.SortField = &v
	return s
}

func (s *ListProductQuotasRequest) SetSortOrder(v string) *ListProductQuotasRequest {
	s.SortOrder = &v
	return s
}

type ListProductQuotasRequestDimensions struct {
	// The key of the dimension.
	//
	// >  The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the dimension.
	//
	// >  The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListProductQuotasRequestDimensions) String() string {
	return tea.Prettify(s)
}

func (s ListProductQuotasRequestDimensions) GoString() string {
	return s.String()
}

func (s *ListProductQuotasRequestDimensions) SetKey(v string) *ListProductQuotasRequestDimensions {
	s.Key = &v
	return s
}

func (s *ListProductQuotasRequestDimensions) SetValue(v string) *ListProductQuotasRequestDimensions {
	s.Value = &v
	return s
}

type ListProductQuotasResponseBody struct {
	// The maximum number of records that are returned for the query.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position at which the query ends. An empty value indicates that all data is returned.
	//
	// example:
	//
	// AAAAAd98/tlL5GF2aM7UMKQGM8LZesIPr0CbfxASQvHV/pwcmVKNfdBbW8OPld3NvG9Cy8+dNcyFzyUttQA3IONfBhRGpXFyiVoTgK+dupBsP2mX
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The queried quotas.
	Quotas []*ListProductQuotasResponseBodyQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// D0131FD5-5397-44FE-BF5A-4B7165B813CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records that are returned for the query.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProductQuotasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductQuotasResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductQuotasResponseBody) SetMaxResults(v int32) *ListProductQuotasResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListProductQuotasResponseBody) SetNextToken(v string) *ListProductQuotasResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProductQuotasResponseBody) SetQuotas(v []*ListProductQuotasResponseBodyQuotas) *ListProductQuotasResponseBody {
	s.Quotas = v
	return s
}

func (s *ListProductQuotasResponseBody) SetRequestId(v string) *ListProductQuotasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductQuotasResponseBody) SetTotalCount(v int32) *ListProductQuotasResponseBody {
	s.TotalCount = &v
	return s
}

type ListProductQuotasResponseBodyQuotas struct {
	// Indicates whether the quota is adjustable. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Adjustable *bool `json:"Adjustable,omitempty" xml:"Adjustable,omitempty"`
	// N/A
	ApplicableRange []*float32 `json:"ApplicableRange,omitempty" xml:"ApplicableRange,omitempty" type:"Repeated"`
	// The type of the adjustable value. Valid values:
	//
	// 	- continuous
	//
	// 	- discontinuous
	//
	// example:
	//
	// discontinuous
	ApplicableType *string `json:"ApplicableType,omitempty" xml:"ApplicableType,omitempty"`
	// The reason for submitting a quota increase request.
	//
	// example:
	//
	// The business xxx is expected to grow by 50%.
	ApplyReasonTips *string `json:"ApplyReasonTips,omitempty" xml:"ApplyReasonTips,omitempty"`
	// Indicates whether the system shows the used value of the quota. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Consumable *bool `json:"Consumable,omitempty" xml:"Consumable,omitempty"`
	// The quota dimensions. Format: `{"regionId":"Region"}`.
	//
	// example:
	//
	// {"regionId":"cn-hangzhou"}
	Dimensions map[string]interface{} `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The start time of the validity period of the quota. The value is displayed in UTC.
	//
	// example:
	//
	// 2022-09-28T06:07:00Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The end time of the validity period of the quota. The value is displayed in UTC.
	//
	// example:
	//
	// 2022-09-29T06:07:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the quota is a global quota. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	GlobalQuota *bool `json:"GlobalQuota,omitempty" xml:"GlobalQuota,omitempty"`
	// The calculation cycle of the quota.
	Period *ListProductQuotasResponseBodyQuotasPeriod `json:"Period,omitempty" xml:"Period,omitempty" type:"Struct"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// example:
	//
	// ecs-spec
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// ecs.g5.2xlarge
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the quota.
	//
	// example:
	//
	// acs:quotas:cn-hangzhou:107992689699****:quota/ecs/ecs.g5.2xlarge/postpaid/vpc/cn-hangzhou/instancetype/cn-hangzhou-i
	QuotaArn *string `json:"QuotaArn,omitempty" xml:"QuotaArn,omitempty"`
	// The type of the quota. Valid values:
	//
	// 	- CommonQuota: general quota
	//
	// 	- FlowControl: API rate limit
	//
	// 	- WhiteListLabel: privilege
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The description of the quota.
	//
	// example:
	//
	// ecs.g5.2xlarge
	QuotaDescription *string `json:"QuotaDescription,omitempty" xml:"QuotaDescription,omitempty"`
	// The details of the quota.
	QuotaItems []*ListProductQuotasResponseBodyQuotasQuotaItems `json:"QuotaItems,omitempty" xml:"QuotaItems,omitempty" type:"Repeated"`
	// The quota name.
	//
	// example:
	//
	// ecs.g5.2xlarge
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The type of the quota. Valid values:
	//
	// 	- privilege
	//
	// 	- normal
	//
	// example:
	//
	// privilege
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// The unit of the quota.
	//
	// example:
	//
	// AMOUNT
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// N/A
	SupportedRange []*float32 `json:"SupportedRange,omitempty" xml:"SupportedRange,omitempty" type:"Repeated"`
	// The quota value.
	//
	// example:
	//
	// 200
	TotalQuota *float32 `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	// The quota usage.
	//
	// example:
	//
	// 1
	TotalUsage *float32 `json:"TotalUsage,omitempty" xml:"TotalUsage,omitempty"`
	// The reason why the quota is not adjustable. Valid values:
	//
	// 	- nonactivated: The service is not activated.
	//
	// 	- applicationProcess: The application is being processed.
	//
	// 	- limitReached: The quota limit is reached.
	//
	// example:
	//
	// applicationProcess
	UnadjustableDetail *string                                         `json:"UnadjustableDetail,omitempty" xml:"UnadjustableDetail,omitempty"`
	UsageMetric        *ListProductQuotasResponseBodyQuotasUsageMetric `json:"UsageMetric,omitempty" xml:"UsageMetric,omitempty" type:"Struct"`
}

func (s ListProductQuotasResponseBodyQuotas) String() string {
	return tea.Prettify(s)
}

func (s ListProductQuotasResponseBodyQuotas) GoString() string {
	return s.String()
}

func (s *ListProductQuotasResponseBodyQuotas) SetAdjustable(v bool) *ListProductQuotasResponseBodyQuotas {
	s.Adjustable = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetApplicableRange(v []*float32) *ListProductQuotasResponseBodyQuotas {
	s.ApplicableRange = v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetApplicableType(v string) *ListProductQuotasResponseBodyQuotas {
	s.ApplicableType = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetApplyReasonTips(v string) *ListProductQuotasResponseBodyQuotas {
	s.ApplyReasonTips = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetConsumable(v bool) *ListProductQuotasResponseBodyQuotas {
	s.Consumable = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetDimensions(v map[string]interface{}) *ListProductQuotasResponseBodyQuotas {
	s.Dimensions = v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetEffectiveTime(v string) *ListProductQuotasResponseBodyQuotas {
	s.EffectiveTime = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetExpireTime(v string) *ListProductQuotasResponseBodyQuotas {
	s.ExpireTime = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetGlobalQuota(v bool) *ListProductQuotasResponseBodyQuotas {
	s.GlobalQuota = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetPeriod(v *ListProductQuotasResponseBodyQuotasPeriod) *ListProductQuotasResponseBodyQuotas {
	s.Period = v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetProductCode(v string) *ListProductQuotasResponseBodyQuotas {
	s.ProductCode = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetQuotaActionCode(v string) *ListProductQuotasResponseBodyQuotas {
	s.QuotaActionCode = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetQuotaArn(v string) *ListProductQuotasResponseBodyQuotas {
	s.QuotaArn = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetQuotaCategory(v string) *ListProductQuotasResponseBodyQuotas {
	s.QuotaCategory = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetQuotaDescription(v string) *ListProductQuotasResponseBodyQuotas {
	s.QuotaDescription = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetQuotaItems(v []*ListProductQuotasResponseBodyQuotasQuotaItems) *ListProductQuotasResponseBodyQuotas {
	s.QuotaItems = v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetQuotaName(v string) *ListProductQuotasResponseBodyQuotas {
	s.QuotaName = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetQuotaType(v string) *ListProductQuotasResponseBodyQuotas {
	s.QuotaType = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetQuotaUnit(v string) *ListProductQuotasResponseBodyQuotas {
	s.QuotaUnit = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetSupportedRange(v []*float32) *ListProductQuotasResponseBodyQuotas {
	s.SupportedRange = v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetTotalQuota(v float32) *ListProductQuotasResponseBodyQuotas {
	s.TotalQuota = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetTotalUsage(v float32) *ListProductQuotasResponseBodyQuotas {
	s.TotalUsage = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetUnadjustableDetail(v string) *ListProductQuotasResponseBodyQuotas {
	s.UnadjustableDetail = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetUsageMetric(v *ListProductQuotasResponseBodyQuotasUsageMetric) *ListProductQuotasResponseBodyQuotas {
	s.UsageMetric = v
	return s
}

type ListProductQuotasResponseBodyQuotasPeriod struct {
	// The unit of the calculation cycle. Valid values:
	//
	// 	- second
	//
	// 	- minute
	//
	// 	- hour
	//
	// 	- day
	//
	// 	- week
	//
	// example:
	//
	// day
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The value of the calculation cycle.
	//
	// example:
	//
	// 1
	PeriodValue *int32 `json:"PeriodValue,omitempty" xml:"PeriodValue,omitempty"`
}

func (s ListProductQuotasResponseBodyQuotasPeriod) String() string {
	return tea.Prettify(s)
}

func (s ListProductQuotasResponseBodyQuotasPeriod) GoString() string {
	return s.String()
}

func (s *ListProductQuotasResponseBodyQuotasPeriod) SetPeriodUnit(v string) *ListProductQuotasResponseBodyQuotasPeriod {
	s.PeriodUnit = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotasPeriod) SetPeriodValue(v int32) *ListProductQuotasResponseBodyQuotasPeriod {
	s.PeriodValue = &v
	return s
}

type ListProductQuotasResponseBodyQuotasQuotaItems struct {
	// The quota value.
	//
	// example:
	//
	// 10
	Quota *string `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// The unit of the quota.
	//
	// example:
	//
	// AMOUNT
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// The category of the quota. Valid values:
	//
	// 	- BaseQuota: base quota
	//
	// 	- ReservedQuota: reserved quota
	//
	// example:
	//
	// BaseQuota
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The quota usage.
	//
	// example:
	//
	// 1
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s ListProductQuotasResponseBodyQuotasQuotaItems) String() string {
	return tea.Prettify(s)
}

func (s ListProductQuotasResponseBodyQuotasQuotaItems) GoString() string {
	return s.String()
}

func (s *ListProductQuotasResponseBodyQuotasQuotaItems) SetQuota(v string) *ListProductQuotasResponseBodyQuotasQuotaItems {
	s.Quota = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotasQuotaItems) SetQuotaUnit(v string) *ListProductQuotasResponseBodyQuotasQuotaItems {
	s.QuotaUnit = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotasQuotaItems) SetType(v string) *ListProductQuotasResponseBodyQuotasQuotaItems {
	s.Type = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotasQuotaItems) SetUsage(v string) *ListProductQuotasResponseBodyQuotasQuotaItems {
	s.Usage = &v
	return s
}

type ListProductQuotasResponseBodyQuotasUsageMetric struct {
	MetricDimensions map[string]*string `json:"MetricDimensions,omitempty" xml:"MetricDimensions,omitempty"`
	MetricName       *string            `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	MetricNamespace  *string            `json:"MetricNamespace,omitempty" xml:"MetricNamespace,omitempty"`
}

func (s ListProductQuotasResponseBodyQuotasUsageMetric) String() string {
	return tea.Prettify(s)
}

func (s ListProductQuotasResponseBodyQuotasUsageMetric) GoString() string {
	return s.String()
}

func (s *ListProductQuotasResponseBodyQuotasUsageMetric) SetMetricDimensions(v map[string]*string) *ListProductQuotasResponseBodyQuotasUsageMetric {
	s.MetricDimensions = v
	return s
}

func (s *ListProductQuotasResponseBodyQuotasUsageMetric) SetMetricName(v string) *ListProductQuotasResponseBodyQuotasUsageMetric {
	s.MetricName = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotasUsageMetric) SetMetricNamespace(v string) *ListProductQuotasResponseBodyQuotasUsageMetric {
	s.MetricNamespace = &v
	return s
}

type ListProductQuotasResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProductQuotasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProductQuotasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProductQuotasResponse) GoString() string {
	return s.String()
}

func (s *ListProductQuotasResponse) SetHeaders(v map[string]*string) *ListProductQuotasResponse {
	s.Headers = v
	return s
}

func (s *ListProductQuotasResponse) SetStatusCode(v int32) *ListProductQuotasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductQuotasResponse) SetBody(v *ListProductQuotasResponseBody) *ListProductQuotasResponse {
	s.Body = v
	return s
}

type ListProductsRequest struct {
	// The maximum number of records that can be returned for the query.
	//
	// Valid values: 1 to 200. Default value: 30.
	//
	// example:
	//
	// 4
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position from which you want to start the query. If you leave this parameter empty, the query starts from the beginning.
	//
	// example:
	//
	// 4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListProductsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductsRequest) GoString() string {
	return s.String()
}

func (s *ListProductsRequest) SetMaxResults(v int32) *ListProductsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProductsRequest) SetNextToken(v string) *ListProductsRequest {
	s.NextToken = &v
	return s
}

type ListProductsResponseBody struct {
	// The maximum number of records that are returned for the query.
	//
	// example:
	//
	// 4
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position at which the query ends. An empty value indicates that all data is returned.
	//
	// example:
	//
	// 4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information of the Alibaba Cloud service.
	ProductInfo []*ListProductsResponseBodyProductInfo `json:"ProductInfo,omitempty" xml:"ProductInfo,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 1DA9C136-11BC-4C39-ADC6-B86276128072
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records that are returned for the query.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProductsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBody) SetMaxResults(v int32) *ListProductsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListProductsResponseBody) SetNextToken(v string) *ListProductsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProductsResponseBody) SetProductInfo(v []*ListProductsResponseBodyProductInfo) *ListProductsResponseBody {
	s.ProductInfo = v
	return s
}

func (s *ListProductsResponseBody) SetRequestId(v string) *ListProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductsResponseBody) SetTotalCount(v int32) *ListProductsResponseBody {
	s.TotalCount = &v
	return s
}

type ListProductsResponseBodyProductInfo struct {
	// Indicates whether the Alibaba Cloud service supports general quotas. Valid values:
	//
	// 	- support: The Alibaba Cloud service supports general quotas.
	//
	// 	- unsupport: The Alibaba Cloud service does not support general quotas.
	//
	// example:
	//
	// support
	CommonQuotaSupport *string `json:"CommonQuotaSupport,omitempty" xml:"CommonQuotaSupport,omitempty"`
	// Indicates whether the Alibaba Cloud service supports dynamic quota adjustment. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Dynamic *bool `json:"Dynamic,omitempty" xml:"Dynamic,omitempty"`
	// Indicates whether the Alibaba Cloud service supports API rate limits. Valid values:
	//
	// 	- support: The Alibaba Cloud service supports API rate limits.
	//
	// 	- unsupport: The Alibaba Cloud service does not support API rate limits.
	//
	// example:
	//
	// unsupport
	FlowControlSupport *string `json:"FlowControlSupport,omitempty" xml:"FlowControlSupport,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The name of the Alibaba Cloud service.
	//
	// example:
	//
	// Elastic Compute Service (ECS)
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The English name of the Alibaba Cloud service.
	//
	// example:
	//
	// Elastic Compute Service
	ProductNameEn *string `json:"ProductNameEn,omitempty" xml:"ProductNameEn,omitempty"`
	// The ID of the service category.
	//
	// example:
	//
	// 5
	SecondCategoryId *int64 `json:"SecondCategoryId,omitempty" xml:"SecondCategoryId,omitempty"`
	// The name of the service category.
	//
	// example:
	//
	// Elastic Compute
	SecondCategoryName *string `json:"SecondCategoryName,omitempty" xml:"SecondCategoryName,omitempty"`
	// The English name of the service category.
	//
	// example:
	//
	// Elastic Compute
	SecondCategoryNameEn *string `json:"SecondCategoryNameEn,omitempty" xml:"SecondCategoryNameEn,omitempty"`
	// Indicates whether the Alibaba Cloud service supports whitelist quotas. Valid values:
	//
	// 	- support: The Alibaba Cloud service supports whitelist quotas.
	//
	// 	- unsupport: The Alibaba Cloud service does not support whitelist quotas.
	//
	// example:
	//
	// support
	WhiteListLabelQuotaSupport *string `json:"WhiteListLabelQuotaSupport,omitempty" xml:"WhiteListLabelQuotaSupport,omitempty"`
}

func (s ListProductsResponseBodyProductInfo) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBodyProductInfo) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyProductInfo) SetCommonQuotaSupport(v string) *ListProductsResponseBodyProductInfo {
	s.CommonQuotaSupport = &v
	return s
}

func (s *ListProductsResponseBodyProductInfo) SetDynamic(v bool) *ListProductsResponseBodyProductInfo {
	s.Dynamic = &v
	return s
}

func (s *ListProductsResponseBodyProductInfo) SetFlowControlSupport(v string) *ListProductsResponseBodyProductInfo {
	s.FlowControlSupport = &v
	return s
}

func (s *ListProductsResponseBodyProductInfo) SetProductCode(v string) *ListProductsResponseBodyProductInfo {
	s.ProductCode = &v
	return s
}

func (s *ListProductsResponseBodyProductInfo) SetProductName(v string) *ListProductsResponseBodyProductInfo {
	s.ProductName = &v
	return s
}

func (s *ListProductsResponseBodyProductInfo) SetProductNameEn(v string) *ListProductsResponseBodyProductInfo {
	s.ProductNameEn = &v
	return s
}

func (s *ListProductsResponseBodyProductInfo) SetSecondCategoryId(v int64) *ListProductsResponseBodyProductInfo {
	s.SecondCategoryId = &v
	return s
}

func (s *ListProductsResponseBodyProductInfo) SetSecondCategoryName(v string) *ListProductsResponseBodyProductInfo {
	s.SecondCategoryName = &v
	return s
}

func (s *ListProductsResponseBodyProductInfo) SetSecondCategoryNameEn(v string) *ListProductsResponseBodyProductInfo {
	s.SecondCategoryNameEn = &v
	return s
}

func (s *ListProductsResponseBodyProductInfo) SetWhiteListLabelQuotaSupport(v string) *ListProductsResponseBodyProductInfo {
	s.WhiteListLabelQuotaSupport = &v
	return s
}

type ListProductsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProductsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProductsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponse) GoString() string {
	return s.String()
}

func (s *ListProductsResponse) SetHeaders(v map[string]*string) *ListProductsResponse {
	s.Headers = v
	return s
}

func (s *ListProductsResponse) SetStatusCode(v int32) *ListProductsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductsResponse) SetBody(v *ListProductsResponseBody) *ListProductsResponse {
	s.Body = v
	return s
}

type ListQuotaAlarmsRequest struct {
	// The name of the alert.
	//
	// example:
	//
	// rules
	AlarmName *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	// The maximum number of records that you want to return for the query.
	//
	// Valid values: 1 to 200. Default value: 30.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position from which you want to start the query.
	//
	// > An empty value indicates that the query starts from the beginning.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// >  To query the abbreviation of an Alibaba Cloud service name, call the [ListProducts](https://help.aliyun.com/document_detail/440554.html) operation and check the value of `ProductCode` in the response.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// >
	//
	// 	- To obtain the quota ID of a cloud service, call the [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html) operation and check the value of `QuotaActionCode` in the response.
	//
	// 	- If you specify this parameter, you must specify `ProductCode`.
	//
	// example:
	//
	// q_hvnoqv
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The quota dimensions.
	QuotaDimensions []*ListQuotaAlarmsRequestQuotaDimensions `json:"QuotaDimensions,omitempty" xml:"QuotaDimensions,omitempty" type:"Repeated"`
}

func (s ListQuotaAlarmsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaAlarmsRequest) GoString() string {
	return s.String()
}

func (s *ListQuotaAlarmsRequest) SetAlarmName(v string) *ListQuotaAlarmsRequest {
	s.AlarmName = &v
	return s
}

func (s *ListQuotaAlarmsRequest) SetMaxResults(v int32) *ListQuotaAlarmsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListQuotaAlarmsRequest) SetNextToken(v string) *ListQuotaAlarmsRequest {
	s.NextToken = &v
	return s
}

func (s *ListQuotaAlarmsRequest) SetProductCode(v string) *ListQuotaAlarmsRequest {
	s.ProductCode = &v
	return s
}

func (s *ListQuotaAlarmsRequest) SetQuotaActionCode(v string) *ListQuotaAlarmsRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *ListQuotaAlarmsRequest) SetQuotaDimensions(v []*ListQuotaAlarmsRequestQuotaDimensions) *ListQuotaAlarmsRequest {
	s.QuotaDimensions = v
	return s
}

type ListQuotaAlarmsRequestQuotaDimensions struct {
	// The key of the dimension.
	//
	// >
	//
	// 	- The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// 	- This parameter is required if you set the `ProductCode` parameter to `ecs`, `ecs-spec`, `actiontrail`, or `ess`.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the dimension.
	//
	// >
	//
	// 	- The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// 	- This parameter is required if you set the `ProductCode` parameter to `ecs`, `ecs-spec`, `actiontrail`, or `ess`.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListQuotaAlarmsRequestQuotaDimensions) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaAlarmsRequestQuotaDimensions) GoString() string {
	return s.String()
}

func (s *ListQuotaAlarmsRequestQuotaDimensions) SetKey(v string) *ListQuotaAlarmsRequestQuotaDimensions {
	s.Key = &v
	return s
}

func (s *ListQuotaAlarmsRequestQuotaDimensions) SetValue(v string) *ListQuotaAlarmsRequestQuotaDimensions {
	s.Value = &v
	return s
}

type ListQuotaAlarmsResponseBody struct {
	// The maximum number of records that are returned for the query.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position at which the query ends.
	//
	// > An empty value indicates that all data is returned.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The details about the quota alert.
	QuotaAlarms []*ListQuotaAlarmsResponseBodyQuotaAlarms `json:"QuotaAlarms,omitempty" xml:"QuotaAlarms,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 87F3B755-3BD2-4C76-B36A-93247002918C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of quota alerts.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListQuotaAlarmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaAlarmsResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotaAlarmsResponseBody) SetMaxResults(v int32) *ListQuotaAlarmsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListQuotaAlarmsResponseBody) SetNextToken(v string) *ListQuotaAlarmsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListQuotaAlarmsResponseBody) SetQuotaAlarms(v []*ListQuotaAlarmsResponseBodyQuotaAlarms) *ListQuotaAlarmsResponseBody {
	s.QuotaAlarms = v
	return s
}

func (s *ListQuotaAlarmsResponseBody) SetRequestId(v string) *ListQuotaAlarmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuotaAlarmsResponseBody) SetTotalCount(v int32) *ListQuotaAlarmsResponseBody {
	s.TotalCount = &v
	return s
}

type ListQuotaAlarmsResponseBodyQuotaAlarms struct {
	// The ID of the alert.
	//
	// example:
	//
	// a2efa7fc-832f-47bb-8054-39e28012****
	AlarmId *string `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// The name of the alert event.
	//
	// example:
	//
	// rules
	AlarmName *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	// The time when the quota alert was created.
	//
	// example:
	//
	// 2020-11-27T07:23:34Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the alert threshold was reached. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// false
	ExceedThreshold *bool `json:"ExceedThreshold,omitempty" xml:"ExceedThreshold,omitempty"`
	// The alert notification methods.
	NotifyChannels []*string `json:"NotifyChannels,omitempty" xml:"NotifyChannels,omitempty" type:"Repeated"`
	// The alert contact. The value is accountContact.
	//
	// example:
	//
	// accountContact
	NotifyTarget *string `json:"NotifyTarget,omitempty" xml:"NotifyTarget,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// example:
	//
	// config
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	//
	// example:
	//
	// q_hvnoqv
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The quota dimensions.
	//
	// example:
	//
	// {"regionId":"cn-hangzhou"}
	QuotaDimensions map[string]interface{} `json:"QuotaDimensions,omitempty" xml:"QuotaDimensions,omitempty"`
	// The used quota.
	//
	// example:
	//
	// 73
	QuotaUsage *float32 `json:"QuotaUsage,omitempty" xml:"QuotaUsage,omitempty"`
	// The value of the quota.
	//
	// example:
	//
	// 200
	QuotaValue *float32 `json:"QuotaValue,omitempty" xml:"QuotaValue,omitempty"`
	// The numeric value of the alert threshold.
	//
	// example:
	//
	// 160
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The percentage of the alert threshold.
	//
	// example:
	//
	// 80
	ThresholdPercent *float32 `json:"ThresholdPercent,omitempty" xml:"ThresholdPercent,omitempty"`
	// The type of the quota alert. Valid values:
	//
	// 	- used: The alert is created for the used quota.
	//
	// 	- usable: The alert is created for the available quota.
	//
	// example:
	//
	// used
	ThresholdType *string `json:"ThresholdType,omitempty" xml:"ThresholdType,omitempty"`
	// The webhook URL.
	//
	// example:
	//
	// https://www.aliyun.com/webhook
	WebHook *string `json:"WebHook,omitempty" xml:"WebHook,omitempty"`
}

func (s ListQuotaAlarmsResponseBodyQuotaAlarms) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaAlarmsResponseBodyQuotaAlarms) GoString() string {
	return s.String()
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetAlarmId(v string) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.AlarmId = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetAlarmName(v string) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.AlarmName = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetCreateTime(v string) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.CreateTime = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetExceedThreshold(v bool) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.ExceedThreshold = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetNotifyChannels(v []*string) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.NotifyChannels = v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetNotifyTarget(v string) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.NotifyTarget = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetProductCode(v string) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.ProductCode = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetQuotaActionCode(v string) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.QuotaActionCode = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetQuotaDimensions(v map[string]interface{}) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.QuotaDimensions = v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetQuotaUsage(v float32) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.QuotaUsage = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetQuotaValue(v float32) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.QuotaValue = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetThreshold(v float32) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.Threshold = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetThresholdPercent(v float32) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.ThresholdPercent = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetThresholdType(v string) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.ThresholdType = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetWebHook(v string) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.WebHook = &v
	return s
}

type ListQuotaAlarmsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQuotaAlarmsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQuotaAlarmsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaAlarmsResponse) GoString() string {
	return s.String()
}

func (s *ListQuotaAlarmsResponse) SetHeaders(v map[string]*string) *ListQuotaAlarmsResponse {
	s.Headers = v
	return s
}

func (s *ListQuotaAlarmsResponse) SetStatusCode(v int32) *ListQuotaAlarmsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQuotaAlarmsResponse) SetBody(v *ListQuotaAlarmsResponseBody) *ListQuotaAlarmsResponse {
	s.Body = v
	return s
}

type ListQuotaApplicationTemplatesRequest struct {
	// The quota dimensions.
	Dimensions []*ListQuotaApplicationTemplatesRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The ID of the quota item.
	//
	// example:
	//
	// 1****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The maximum number of records that can be returned for the query. Valid values: 1 to 100. Default value: 30.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position from which you want to start the query.
	//
	// > If you leave this parameter empty, the query starts from the beginning.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// >  To query the abbreviation of an Alibaba Cloud service name, call the [ListProducts](https://help.aliyun.com/document_detail/440555.html) operation and check the value of `ProductCode` in the response.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	//
	// example:
	//
	// q_security-groups
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The type of the quota. Valid values:
	//
	// 	- CommonQuota: general quota
	//
	// 	- WhiteListLabel: privilege
	//
	// 	- FlowControl: API rate limit
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
}

func (s ListQuotaApplicationTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaApplicationTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationTemplatesRequest) SetDimensions(v []*ListQuotaApplicationTemplatesRequestDimensions) *ListQuotaApplicationTemplatesRequest {
	s.Dimensions = v
	return s
}

func (s *ListQuotaApplicationTemplatesRequest) SetId(v string) *ListQuotaApplicationTemplatesRequest {
	s.Id = &v
	return s
}

func (s *ListQuotaApplicationTemplatesRequest) SetMaxResults(v int32) *ListQuotaApplicationTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListQuotaApplicationTemplatesRequest) SetNextToken(v string) *ListQuotaApplicationTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListQuotaApplicationTemplatesRequest) SetProductCode(v string) *ListQuotaApplicationTemplatesRequest {
	s.ProductCode = &v
	return s
}

func (s *ListQuotaApplicationTemplatesRequest) SetQuotaActionCode(v string) *ListQuotaApplicationTemplatesRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *ListQuotaApplicationTemplatesRequest) SetQuotaCategory(v string) *ListQuotaApplicationTemplatesRequest {
	s.QuotaCategory = &v
	return s
}

type ListQuotaApplicationTemplatesRequestDimensions struct {
	// The key of the dimension.
	//
	// >
	//
	// 	- The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// 	- You must specify both Key and Value for each quota dimension.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the dimension.
	//
	// >
	//
	// 	- The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// 	- You must specify both Key and Value for each quota dimension.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListQuotaApplicationTemplatesRequestDimensions) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaApplicationTemplatesRequestDimensions) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationTemplatesRequestDimensions) SetKey(v string) *ListQuotaApplicationTemplatesRequestDimensions {
	s.Key = &v
	return s
}

func (s *ListQuotaApplicationTemplatesRequestDimensions) SetValue(v string) *ListQuotaApplicationTemplatesRequestDimensions {
	s.Value = &v
	return s
}

type ListQuotaApplicationTemplatesResponseBody struct {
	// The maximum number of records returned for the query.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position at which the query ends.
	//
	// > An empty value indicates that all data is returned.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The queried quota templates.
	QuotaApplicationTemplates []*ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates `json:"QuotaApplicationTemplates,omitempty" xml:"QuotaApplicationTemplates,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D47B3A10-CDAC-5412-B2EE-EC9A3DBE9053
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records returned for the query.
	//
	// example:
	//
	// 40
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListQuotaApplicationTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaApplicationTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationTemplatesResponseBody) SetMaxResults(v int32) *ListQuotaApplicationTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBody) SetNextToken(v string) *ListQuotaApplicationTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBody) SetQuotaApplicationTemplates(v []*ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) *ListQuotaApplicationTemplatesResponseBody {
	s.QuotaApplicationTemplates = v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBody) SetRequestId(v string) *ListQuotaApplicationTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBody) SetTotalCount(v int32) *ListQuotaApplicationTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

type ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates struct {
	// None
	ApplicableRange []*float32 `json:"ApplicableRange,omitempty" xml:"ApplicableRange,omitempty" type:"Repeated"`
	// The type of the adjustable value. Valid values:
	//
	// 	- continuous
	//
	// 	- discontinuous
	//
	// example:
	//
	// continuous
	ApplicableType *string `json:"ApplicableType,omitempty" xml:"ApplicableType,omitempty"`
	// The requested value of the quota.
	//
	// example:
	//
	// 802
	DesireValue *float32 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The quota dimensions.
	//
	// Format: {"regionId":"Region"}.
	//
	// example:
	//
	// {"regionId":"cn-hangzhou"}
	Dimensions map[string]interface{} `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The start time of the validity period of the quota. The value is displayed in UTC.
	//
	// example:
	//
	// 2022-09-28T06:07:00Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The language of the quota alert notification. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	EnvLanguage *string `json:"EnvLanguage,omitempty" xml:"EnvLanguage,omitempty"`
	// The end time of the validity period of the quota. The value is displayed in UTC.
	//
	// example:
	//
	// 2022-09-29T06:07:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the quota template.
	//
	// example:
	//
	// 1****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether Quota Center sends a notification about the application result. Valid values:
	//
	// 	- 0: no
	//
	// 	- 3: yes
	//
	// example:
	//
	// 0
	NoticeType *int32 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The calculation cycle of the quota.
	Period *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplatesPeriod `json:"Period,omitempty" xml:"Period,omitempty" type:"Struct"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	//
	// example:
	//
	// q_security-groups
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The type of the quota. Valid values:
	//
	// 	- CommonQuota: general quota
	//
	// 	- FlowControl: API rate limit
	//
	// 	- WhiteListLabel: privilege
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The description of the quota.
	//
	// example:
	//
	// The maximum number of security groups that can be created by the current account.
	QuotaDescription *string `json:"QuotaDescription,omitempty" xml:"QuotaDescription,omitempty"`
	// The name of the quota.
	//
	// example:
	//
	// Maximum Number of Security Groups
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
}

func (s ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetApplicableRange(v []*float32) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.ApplicableRange = v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetApplicableType(v string) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.ApplicableType = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetDesireValue(v float32) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.DesireValue = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetDimensions(v map[string]interface{}) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.Dimensions = v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetEffectiveTime(v string) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.EffectiveTime = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetEnvLanguage(v string) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.EnvLanguage = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetExpireTime(v string) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.ExpireTime = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetId(v string) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.Id = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetNoticeType(v int32) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.NoticeType = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetPeriod(v *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplatesPeriod) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.Period = v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetProductCode(v string) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.ProductCode = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetQuotaActionCode(v string) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.QuotaActionCode = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetQuotaCategory(v string) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.QuotaCategory = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetQuotaDescription(v string) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.QuotaDescription = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetQuotaName(v string) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.QuotaName = &v
	return s
}

type ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplatesPeriod struct {
	// The unit of the calculation cycle. Valid values:
	//
	// 	- second
	//
	// 	- minute
	//
	// 	- hour
	//
	// 	- day
	//
	// 	- week
	//
	// example:
	//
	// day
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The value of the calculation cycle.
	//
	// example:
	//
	// 1
	PeriodValue *int32 `json:"PeriodValue,omitempty" xml:"PeriodValue,omitempty"`
}

func (s ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplatesPeriod) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplatesPeriod) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplatesPeriod) SetPeriodUnit(v string) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplatesPeriod {
	s.PeriodUnit = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplatesPeriod) SetPeriodValue(v int32) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplatesPeriod {
	s.PeriodValue = &v
	return s
}

type ListQuotaApplicationTemplatesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQuotaApplicationTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQuotaApplicationTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaApplicationTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationTemplatesResponse) SetHeaders(v map[string]*string) *ListQuotaApplicationTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListQuotaApplicationTemplatesResponse) SetStatusCode(v int32) *ListQuotaApplicationTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponse) SetBody(v *ListQuotaApplicationTemplatesResponseBody) *ListQuotaApplicationTemplatesResponse {
	s.Body = v
	return s
}

type ListQuotaApplicationsRequest struct {
	// The quota dimensions.
	Dimensions []*ListQuotaApplicationsRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The keyword that you want to use to search for the application.
	//
	// example:
	//
	// Cluster
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The maximum number of records that can be returned for the query.
	//
	// Valid values: 1 to 200. Default value: 30.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position from which you want to start the query. If you leave this parameter empty, the query starts from the beginning.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// example:
	//
	// csk
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	//
	// example:
	//
	// q_i5uzm3
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The type of the quota. Valid values:
	//
	// 	- CommonQuota: general quota
	//
	// 	- FlowControl: API rate limit
	//
	// 	- WhiteListLabel: whitelist quota
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The status of the application. Valid values:
	//
	// 	- Disagree: rejects the application.
	//
	// 	- Agree: approves the application.
	//
	// 	- Process: reviews the application.
	//
	// 	- Cancel: cancels the application.
	//
	// example:
	//
	// Agree
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListQuotaApplicationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsRequest) SetDimensions(v []*ListQuotaApplicationsRequestDimensions) *ListQuotaApplicationsRequest {
	s.Dimensions = v
	return s
}

func (s *ListQuotaApplicationsRequest) SetKeyWord(v string) *ListQuotaApplicationsRequest {
	s.KeyWord = &v
	return s
}

func (s *ListQuotaApplicationsRequest) SetMaxResults(v int32) *ListQuotaApplicationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListQuotaApplicationsRequest) SetNextToken(v string) *ListQuotaApplicationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListQuotaApplicationsRequest) SetProductCode(v string) *ListQuotaApplicationsRequest {
	s.ProductCode = &v
	return s
}

func (s *ListQuotaApplicationsRequest) SetQuotaActionCode(v string) *ListQuotaApplicationsRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *ListQuotaApplicationsRequest) SetQuotaCategory(v string) *ListQuotaApplicationsRequest {
	s.QuotaCategory = &v
	return s
}

func (s *ListQuotaApplicationsRequest) SetStatus(v string) *ListQuotaApplicationsRequest {
	s.Status = &v
	return s
}

type ListQuotaApplicationsRequestDimensions struct {
	// The key of the dimension.
	//
	// >  The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the dimension.
	//
	// >  The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListQuotaApplicationsRequestDimensions) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaApplicationsRequestDimensions) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsRequestDimensions) SetKey(v string) *ListQuotaApplicationsRequestDimensions {
	s.Key = &v
	return s
}

func (s *ListQuotaApplicationsRequestDimensions) SetValue(v string) *ListQuotaApplicationsRequestDimensions {
	s.Value = &v
	return s
}

type ListQuotaApplicationsResponseBody struct {
	// The maximum number of records that are returned for the query.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position at which the query ends. An empty value indicates that all data is returned.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The details of the applications.
	QuotaApplications []*ListQuotaApplicationsResponseBodyQuotaApplications `json:"QuotaApplications,omitempty" xml:"QuotaApplications,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 730925FB-0BD0-40AC-AF3A-A6C6E9716879
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of applications.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListQuotaApplicationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsResponseBody) SetMaxResults(v int32) *ListQuotaApplicationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListQuotaApplicationsResponseBody) SetNextToken(v string) *ListQuotaApplicationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListQuotaApplicationsResponseBody) SetQuotaApplications(v []*ListQuotaApplicationsResponseBodyQuotaApplications) *ListQuotaApplicationsResponseBody {
	s.QuotaApplications = v
	return s
}

func (s *ListQuotaApplicationsResponseBody) SetRequestId(v string) *ListQuotaApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuotaApplicationsResponseBody) SetTotalCount(v int32) *ListQuotaApplicationsResponseBody {
	s.TotalCount = &v
	return s
}

type ListQuotaApplicationsResponseBodyQuotaApplications struct {
	// The ID of the application.
	//
	// example:
	//
	// b926571d-cc09-4711-b547-58a615f0****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The time when the application was submitted.
	//
	// example:
	//
	// 2021-01-15T09:13:53Z
	ApplyTime *string `json:"ApplyTime,omitempty" xml:"ApplyTime,omitempty"`
	// The quota value that is approved.
	//
	// example:
	//
	// 101
	ApproveValue *float32 `json:"ApproveValue,omitempty" xml:"ApproveValue,omitempty"`
	// The result of the application.
	//
	// example:
	//
	// Agree
	AuditReason *string `json:"AuditReason,omitempty" xml:"AuditReason,omitempty"`
	// The remarks of the application.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The quota value that is approved.
	//
	// example:
	//
	// 101
	DesireValue *float32 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The dimension of the application.
	//
	// example:
	//
	// {"regionId":"cn-hangzhou"}
	Dimension map[string]interface{} `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The time when the application took effect.
	//
	// example:
	//
	// 2021-01-15T11:46:25Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The time when the application expired.
	//
	// example:
	//
	// 2021-01-17T11:46:25Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether Quota Center sends a notification about the application result. Valid values:
	//
	// 	- 0: A notification about the application result is not sent.
	//
	// 	- 3: A notification about the application result is sent.
	//
	// example:
	//
	// 3
	NoticeType *int32 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The calculation cycle of the quota.
	Period *ListQuotaApplicationsResponseBodyQuotaApplicationsPeriod `json:"Period,omitempty" xml:"Period,omitempty" type:"Struct"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// example:
	//
	// csk
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	//
	// example:
	//
	// q_i5uzm3
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the quota.
	//
	// example:
	//
	// acs:quotas:*:120886317861****:quota/csk/q_i5uzm3
	QuotaArn *string `json:"QuotaArn,omitempty" xml:"QuotaArn,omitempty"`
	// The description of the quota.
	//
	// example:
	//
	// The maximum number of nodes in a cluster
	QuotaDescription *string `json:"QuotaDescription,omitempty" xml:"QuotaDescription,omitempty"`
	// The name of the quota.
	//
	// example:
	//
	// Maximum Number of Nodes
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The unit of the new quota value.
	//
	// example:
	//
	// Node
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// The reason for the application.
	//
	// example:
	//
	// Business expansion
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The status of the application. Valid values:
	//
	// 	- Disagree: The application is rejected.
	//
	// 	- Agree: The application is approved.
	//
	// 	- Process: The application is being reviewed.
	//
	// 	- Cancel: The application is canceled.
	//
	// example:
	//
	// Agree
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListQuotaApplicationsResponseBodyQuotaApplications) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaApplicationsResponseBodyQuotaApplications) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetApplicationId(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetApplyTime(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.ApplyTime = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetApproveValue(v float32) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.ApproveValue = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetAuditReason(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.AuditReason = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetComment(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.Comment = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetDesireValue(v float32) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.DesireValue = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetDimension(v map[string]interface{}) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.Dimension = v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetEffectiveTime(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.EffectiveTime = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetExpireTime(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.ExpireTime = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetNoticeType(v int32) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.NoticeType = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetPeriod(v *ListQuotaApplicationsResponseBodyQuotaApplicationsPeriod) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.Period = v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetProductCode(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.ProductCode = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetQuotaActionCode(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.QuotaActionCode = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetQuotaArn(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.QuotaArn = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetQuotaDescription(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.QuotaDescription = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetQuotaName(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.QuotaName = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetQuotaUnit(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.QuotaUnit = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetReason(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.Reason = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetStatus(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.Status = &v
	return s
}

type ListQuotaApplicationsResponseBodyQuotaApplicationsPeriod struct {
	// The unit of the calculation cycle. Valid values:
	//
	// 	- second
	//
	// 	- minute
	//
	// 	- hour
	//
	// 	- day
	//
	// 	- week
	//
	// example:
	//
	// second
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The value of the calculation cycle.
	//
	// example:
	//
	// 1
	PeriodValue *int64 `json:"PeriodValue,omitempty" xml:"PeriodValue,omitempty"`
}

func (s ListQuotaApplicationsResponseBodyQuotaApplicationsPeriod) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaApplicationsResponseBodyQuotaApplicationsPeriod) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplicationsPeriod) SetPeriodUnit(v string) *ListQuotaApplicationsResponseBodyQuotaApplicationsPeriod {
	s.PeriodUnit = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplicationsPeriod) SetPeriodValue(v int64) *ListQuotaApplicationsResponseBodyQuotaApplicationsPeriod {
	s.PeriodValue = &v
	return s
}

type ListQuotaApplicationsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQuotaApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQuotaApplicationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsResponse) SetHeaders(v map[string]*string) *ListQuotaApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListQuotaApplicationsResponse) SetStatusCode(v int32) *ListQuotaApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQuotaApplicationsResponse) SetBody(v *ListQuotaApplicationsResponseBody) *ListQuotaApplicationsResponse {
	s.Body = v
	return s
}

type ListQuotaApplicationsDetailForTemplateRequest struct {
	// The Alibaba Cloud account that is used to submit the quota increase application.
	//
	// example:
	//
	// 135048337611****
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// The ID of the quota application batch.
	//
	// example:
	//
	// d314d6ae-867d-484c-9009-3d421a80****
	BatchQuotaApplicationId *string `json:"BatchQuotaApplicationId,omitempty" xml:"BatchQuotaApplicationId,omitempty"`
	// The maximum number of records that can be returned for the query.
	//
	// Valid values: 1 to 100. Default value: 30.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position from which you want to start the query.
	//
	// >  An empty value indicates that the query starts from the beginning.
	//
	// example:
	//
	// 4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// >  To query the abbreviation of an Alibaba Cloud service name, call the [ListProducts](https://help.aliyun.com/document_detail/440555.html) operation and check the value of `ProductCode` in the response.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// ecs.c5.large
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The quota type. Valid values:
	//
	// 	- CommonQuota: general quota
	//
	// 	- FlowControl: API rate limit
	//
	// 	- WhiteListLabel: privilege
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The approval state of the quota increase application. Valid values:
	//
	// 	- Disagree: The application is rejected.
	//
	// 	- Agree: The application is approved.
	//
	// 	- Process: The application is in review.
	//
	// 	- Cancel: The application is canceled.
	//
	// example:
	//
	// Agree
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListQuotaApplicationsDetailForTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaApplicationsDetailForTemplateRequest) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsDetailForTemplateRequest) SetAliyunUid(v string) *ListQuotaApplicationsDetailForTemplateRequest {
	s.AliyunUid = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateRequest) SetBatchQuotaApplicationId(v string) *ListQuotaApplicationsDetailForTemplateRequest {
	s.BatchQuotaApplicationId = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateRequest) SetMaxResults(v int32) *ListQuotaApplicationsDetailForTemplateRequest {
	s.MaxResults = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateRequest) SetNextToken(v string) *ListQuotaApplicationsDetailForTemplateRequest {
	s.NextToken = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateRequest) SetProductCode(v string) *ListQuotaApplicationsDetailForTemplateRequest {
	s.ProductCode = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateRequest) SetQuotaActionCode(v string) *ListQuotaApplicationsDetailForTemplateRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateRequest) SetQuotaCategory(v string) *ListQuotaApplicationsDetailForTemplateRequest {
	s.QuotaCategory = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateRequest) SetStatus(v string) *ListQuotaApplicationsDetailForTemplateRequest {
	s.Status = &v
	return s
}

type ListQuotaApplicationsDetailForTemplateResponseBody struct {
	// The maximum number of records that can be returned for the query.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position from which you want to start the query.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The details of the quota increase application.
	QuotaApplications []*ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications `json:"QuotaApplications,omitempty" xml:"QuotaApplications,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D47B3A10-CDAC-5412-B2EE-EC9A3DBE9053
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records that are returned for the query.
	//
	// example:
	//
	// 9
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListQuotaApplicationsDetailForTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaApplicationsDetailForTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBody) SetMaxResults(v int32) *ListQuotaApplicationsDetailForTemplateResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBody) SetNextToken(v string) *ListQuotaApplicationsDetailForTemplateResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBody) SetQuotaApplications(v []*ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) *ListQuotaApplicationsDetailForTemplateResponseBody {
	s.QuotaApplications = v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBody) SetRequestId(v string) *ListQuotaApplicationsDetailForTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBody) SetTotalCount(v int32) *ListQuotaApplicationsDetailForTemplateResponseBody {
	s.TotalCount = &v
	return s
}

type ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 175376458581****
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// The ID of the quota increase application.
	//
	// example:
	//
	// b926571d-cc09-4711-b547-58a615f0****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The time when the quota increase application was submitted. The value is displayed in UTC.
	//
	// example:
	//
	// 2021-01-15T09:13:53Z
	ApplyTime *string `json:"ApplyTime,omitempty" xml:"ApplyTime,omitempty"`
	// The quota value that is approved.
	//
	// example:
	//
	// 101
	ApproveValue *float64 `json:"ApproveValue,omitempty" xml:"ApproveValue,omitempty"`
	// The approval result of the quota increase application.
	//
	// example:
	//
	// Agree
	AuditReason *string `json:"AuditReason,omitempty" xml:"AuditReason,omitempty"`
	// The ID of the quota application batch.
	//
	// example:
	//
	// d314d6ae-867d-484c-9009-3d421a80****
	BatchQuotaApplicationId *string `json:"BatchQuotaApplicationId,omitempty" xml:"BatchQuotaApplicationId,omitempty"`
	// The requested value of the quota.
	//
	// example:
	//
	// 60
	DesireValue *float64 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The start time of the validity period of the quota. The value is displayed in UTC.
	//
	// example:
	//
	// 2021-01-15T11:46:25Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The language of the quota alert notification. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	EnvLanguage *string `json:"EnvLanguage,omitempty" xml:"EnvLanguage,omitempty"`
	// The end time of the validity period of the quota. The value is displayed in UTC.
	//
	// example:
	//
	// 2021-01-17T11:46:25Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether Quota Center sends a notification about the application result. Valid values:
	//
	// 	- 0: no
	//
	// 	- 3: yes
	//
	// example:
	//
	// 3
	NoticeType *int32 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The calculation cycle of the quota.
	Period *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplicationsPeriod `json:"Period,omitempty" xml:"Period,omitempty" type:"Struct"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// ecs.n1.large
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the quota.
	//
	// example:
	//
	// acs:quotas:*:120886317861****:quota/csk/q_i5uzm3
	QuotaArn *string `json:"QuotaArn,omitempty" xml:"QuotaArn,omitempty"`
	// The quota type. Valid values:
	//
	// 	- CommonQuota (default): general quota
	//
	// 	- WhiteListLabel: privilege
	//
	// 	- FlowControl: API rate limit
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The description of the quota.
	//
	// example:
	//
	// The maximum number of nodes in a cluster
	QuotaDescription *string `json:"QuotaDescription,omitempty" xml:"QuotaDescription,omitempty"`
	// The quota dimensions.
	QuotaDimension map[string]*string `json:"QuotaDimension,omitempty" xml:"QuotaDimension,omitempty"`
	// The quota name.
	//
	// example:
	//
	// Maximum Number of Nodes
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The unit of the quota.
	//
	// example:
	//
	// GiB
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// The reason for the quota increase application.
	//
	// example:
	//
	// Business expansion
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The approval status of the quota increase application. Valid values:
	//
	// 	- Disagree: The application is rejected.
	//
	// 	- Agree: The application is approved.
	//
	// 	- Process: The application is pending approval.
	//
	// 	- Cancel: The application is canceled.
	//
	// example:
	//
	// Agree
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetAliyunUid(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.AliyunUid = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetApplicationId(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetApplyTime(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.ApplyTime = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetApproveValue(v float64) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.ApproveValue = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetAuditReason(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.AuditReason = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetBatchQuotaApplicationId(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.BatchQuotaApplicationId = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetDesireValue(v float64) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.DesireValue = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetEffectiveTime(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.EffectiveTime = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetEnvLanguage(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.EnvLanguage = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetExpireTime(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.ExpireTime = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetNoticeType(v int32) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.NoticeType = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetPeriod(v *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplicationsPeriod) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.Period = v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetProductCode(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.ProductCode = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetQuotaActionCode(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.QuotaActionCode = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetQuotaArn(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.QuotaArn = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetQuotaCategory(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.QuotaCategory = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetQuotaDescription(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.QuotaDescription = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetQuotaDimension(v map[string]*string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.QuotaDimension = v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetQuotaName(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.QuotaName = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetQuotaUnit(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.QuotaUnit = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetReason(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.Reason = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications) SetStatus(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplications {
	s.Status = &v
	return s
}

type ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplicationsPeriod struct {
	// The unit of the calculation cycle of the quota.
	//
	// example:
	//
	// second
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The value of the calculation cycle of the quota.
	//
	// example:
	//
	// 1
	PeriodValue *int32 `json:"PeriodValue,omitempty" xml:"PeriodValue,omitempty"`
}

func (s ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplicationsPeriod) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplicationsPeriod) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplicationsPeriod) SetPeriodUnit(v string) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplicationsPeriod {
	s.PeriodUnit = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplicationsPeriod) SetPeriodValue(v int32) *ListQuotaApplicationsDetailForTemplateResponseBodyQuotaApplicationsPeriod {
	s.PeriodValue = &v
	return s
}

type ListQuotaApplicationsDetailForTemplateResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQuotaApplicationsDetailForTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQuotaApplicationsDetailForTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaApplicationsDetailForTemplateResponse) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsDetailForTemplateResponse) SetHeaders(v map[string]*string) *ListQuotaApplicationsDetailForTemplateResponse {
	s.Headers = v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponse) SetStatusCode(v int32) *ListQuotaApplicationsDetailForTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponse) SetBody(v *ListQuotaApplicationsDetailForTemplateResponseBody) *ListQuotaApplicationsDetailForTemplateResponse {
	s.Body = v
	return s
}

type ListQuotaApplicationsForTemplateRequest struct {
	// The UTC time when the quota application ends.
	//
	// example:
	//
	// 2023-05-22T16:00:00Z
	ApplyEndTime *string `json:"ApplyEndTime,omitempty" xml:"ApplyEndTime,omitempty"`
	// The UTC time when the quota application starts.
	//
	// example:
	//
	// 2023-05-22T16:00:00Z
	ApplyStartTime *string `json:"ApplyStartTime,omitempty" xml:"ApplyStartTime,omitempty"`
	// The ID of the quota application batch.
	//
	// example:
	//
	// d314d6ae-867d-484c-9009-3d421a80****
	BatchQuotaApplicationId *string `json:"BatchQuotaApplicationId,omitempty" xml:"BatchQuotaApplicationId,omitempty"`
	// The maximum number of entries to return for a single request.
	//
	// Valid values: 1 to 100. Default value: 30.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position from which you want to start the query.
	//
	// >  An empty value indicates that the query starts from the beginning.
	//
	// example:
	//
	// 4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// >  To query the abbreviation of an Alibaba Cloud service name, call the [ListProducts](https://help.aliyun.com/document_detail/440555.html) operation and check the value of `ProductCode` in the response.
	//
	// example:
	//
	// ecs-spec
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// ecs.g5.2xlarge
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The quota type. Valid values:
	//
	// 	- CommonQuota: general quota
	//
	// 	- FlowControl: API rate limit
	//
	// 	- WhiteListLabel: privilege
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
}

func (s ListQuotaApplicationsForTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaApplicationsForTemplateRequest) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsForTemplateRequest) SetApplyEndTime(v string) *ListQuotaApplicationsForTemplateRequest {
	s.ApplyEndTime = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateRequest) SetApplyStartTime(v string) *ListQuotaApplicationsForTemplateRequest {
	s.ApplyStartTime = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateRequest) SetBatchQuotaApplicationId(v string) *ListQuotaApplicationsForTemplateRequest {
	s.BatchQuotaApplicationId = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateRequest) SetMaxResults(v int32) *ListQuotaApplicationsForTemplateRequest {
	s.MaxResults = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateRequest) SetNextToken(v string) *ListQuotaApplicationsForTemplateRequest {
	s.NextToken = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateRequest) SetProductCode(v string) *ListQuotaApplicationsForTemplateRequest {
	s.ProductCode = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateRequest) SetQuotaActionCode(v string) *ListQuotaApplicationsForTemplateRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateRequest) SetQuotaCategory(v string) *ListQuotaApplicationsForTemplateRequest {
	s.QuotaCategory = &v
	return s
}

type ListQuotaApplicationsForTemplateResponseBody struct {
	// The maximum number of records that can be returned for the query.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position from which you want to start the query.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The queried quota application records.
	QuotaBatchApplications []*ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications `json:"QuotaBatchApplications,omitempty" xml:"QuotaBatchApplications,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D47B3A10-CDAC-5412-B2EE-EC9A3DBE9053
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records that are returned for the query.
	//
	// example:
	//
	// 67
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListQuotaApplicationsForTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaApplicationsForTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsForTemplateResponseBody) SetMaxResults(v int32) *ListQuotaApplicationsForTemplateResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBody) SetNextToken(v string) *ListQuotaApplicationsForTemplateResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBody) SetQuotaBatchApplications(v []*ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) *ListQuotaApplicationsForTemplateResponseBody {
	s.QuotaBatchApplications = v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBody) SetRequestId(v string) *ListQuotaApplicationsForTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBody) SetTotalCount(v int32) *ListQuotaApplicationsForTemplateResponseBody {
	s.TotalCount = &v
	return s
}

type ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications struct {
	// The Alibaba Cloud accounts that correspond to the resource directory members for which the quotas are applied.
	AliyunUids []*string `json:"AliyunUids,omitempty" xml:"AliyunUids,omitempty" type:"Repeated"`
	// The time when the quota increase application was submitted. The value is displayed in UTC.
	//
	// example:
	//
	// 2022-09-23T02:38:18Z
	ApplyTime *string `json:"ApplyTime,omitempty" xml:"ApplyTime,omitempty"`
	// The number of applications in different approval states.
	AuditStatusVos []*ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplicationsAuditStatusVos `json:"AuditStatusVos,omitempty" xml:"AuditStatusVos,omitempty" type:"Repeated"`
	// The ID of the quota application batch.
	//
	// example:
	//
	// d314d6ae-867d-484c-9009-3d421a80****
	BatchQuotaApplicationId *string `json:"BatchQuotaApplicationId,omitempty" xml:"BatchQuotaApplicationId,omitempty"`
	// The requested value of the quota.
	//
	// example:
	//
	// 105
	DesireValue *float64 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The quota dimensions.
	//
	// Format example: {"regionId":"cn-hangzhou"}.
	//
	// example:
	//
	// {"regionId":"cn-hangzhou"}
	Dimensions map[string]interface{} `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The start time of the validity period of the quota. The value is displayed in UTC.
	//
	// example:
	//
	// 2023-05-22T16:00:00Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The end time of the validity period of the quota. The value is displayed in UTC.
	//
	// example:
	//
	// 2024-05-14T02:08:56Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// example:
	//
	// vpc
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// q_fhoz4k
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The quota type. Valid values:
	//
	// 	- CommonQuota: general quota
	//
	// 	- FlowControl: API rate limit
	//
	// 	- WhiteListLabel: privilege
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The reason for the quota increase application.
	//
	// example:
	//
	// Business expansion
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) SetAliyunUids(v []*string) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications {
	s.AliyunUids = v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) SetApplyTime(v string) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications {
	s.ApplyTime = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) SetAuditStatusVos(v []*ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplicationsAuditStatusVos) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications {
	s.AuditStatusVos = v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) SetBatchQuotaApplicationId(v string) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications {
	s.BatchQuotaApplicationId = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) SetDesireValue(v float64) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications {
	s.DesireValue = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) SetDimensions(v map[string]interface{}) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications {
	s.Dimensions = v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) SetEffectiveTime(v string) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications {
	s.EffectiveTime = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) SetExpireTime(v string) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications {
	s.ExpireTime = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) SetProductCode(v string) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications {
	s.ProductCode = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) SetQuotaActionCode(v string) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications {
	s.QuotaActionCode = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) SetQuotaCategory(v string) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications {
	s.QuotaCategory = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications) SetReason(v string) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplications {
	s.Reason = &v
	return s
}

type ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplicationsAuditStatusVos struct {
	// The number of approval tickets.
	//
	// example:
	//
	// 4
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The approval state of the quota increase application. Valid values:
	//
	// 	- Disagree: The application is rejected.
	//
	// 	- Agree: The application is approved.
	//
	// 	- Process: The application is in review.
	//
	// 	- Cancel: The application is canceled.
	//
	// example:
	//
	// Agree
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplicationsAuditStatusVos) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplicationsAuditStatusVos) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplicationsAuditStatusVos) SetCount(v int32) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplicationsAuditStatusVos {
	s.Count = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplicationsAuditStatusVos) SetStatus(v string) *ListQuotaApplicationsForTemplateResponseBodyQuotaBatchApplicationsAuditStatusVos {
	s.Status = &v
	return s
}

type ListQuotaApplicationsForTemplateResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQuotaApplicationsForTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQuotaApplicationsForTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaApplicationsForTemplateResponse) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsForTemplateResponse) SetHeaders(v map[string]*string) *ListQuotaApplicationsForTemplateResponse {
	s.Headers = v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponse) SetStatusCode(v int32) *ListQuotaApplicationsForTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponse) SetBody(v *ListQuotaApplicationsForTemplateResponseBody) *ListQuotaApplicationsForTemplateResponse {
	s.Body = v
	return s
}

type ModifyQuotaTemplateServiceStatusRequest struct {
	// The status of the quota template. Valid values:
	//
	// 	- \\-1: The quota template is disabled.
	//
	// 	- 1: The quota template is enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ServiceStatus *int32 `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
}

func (s ModifyQuotaTemplateServiceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyQuotaTemplateServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyQuotaTemplateServiceStatusRequest) SetServiceStatus(v int32) *ModifyQuotaTemplateServiceStatusRequest {
	s.ServiceStatus = &v
	return s
}

type ModifyQuotaTemplateServiceStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D47B3A10-CDAC-5412-B2EE-EC9A3DBE9053
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the quota template.
	TemplateServiceStatus *ModifyQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus `json:"TemplateServiceStatus,omitempty" xml:"TemplateServiceStatus,omitempty" type:"Struct"`
}

func (s ModifyQuotaTemplateServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyQuotaTemplateServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyQuotaTemplateServiceStatusResponseBody) SetRequestId(v string) *ModifyQuotaTemplateServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyQuotaTemplateServiceStatusResponseBody) SetTemplateServiceStatus(v *ModifyQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) *ModifyQuotaTemplateServiceStatusResponseBody {
	s.TemplateServiceStatus = v
	return s
}

type ModifyQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus struct {
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-pG****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the quota template. Valid values:
	//
	// 	- \\-1: The quota template is disabled.
	//
	// 	- 1: The quota template is enabled.
	//
	// example:
	//
	// 1
	ServiceStatus *int32 `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
}

func (s ModifyQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) String() string {
	return tea.Prettify(s)
}

func (s ModifyQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) GoString() string {
	return s.String()
}

func (s *ModifyQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) SetResourceDirectoryId(v string) *ModifyQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus {
	s.ResourceDirectoryId = &v
	return s
}

func (s *ModifyQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus) SetServiceStatus(v int32) *ModifyQuotaTemplateServiceStatusResponseBodyTemplateServiceStatus {
	s.ServiceStatus = &v
	return s
}

type ModifyQuotaTemplateServiceStatusResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyQuotaTemplateServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyQuotaTemplateServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyQuotaTemplateServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyQuotaTemplateServiceStatusResponse) SetHeaders(v map[string]*string) *ModifyQuotaTemplateServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyQuotaTemplateServiceStatusResponse) SetStatusCode(v int32) *ModifyQuotaTemplateServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyQuotaTemplateServiceStatusResponse) SetBody(v *ModifyQuotaTemplateServiceStatusResponseBody) *ModifyQuotaTemplateServiceStatusResponse {
	s.Body = v
	return s
}

type ModifyTemplateQuotaItemRequest struct {
	// example:
	//
	// 804
	DesireValue *float32                                    `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	Dimensions  []*ModifyTemplateQuotaItemRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// example:
	//
	// 2021-01-19T09:25:56Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// example:
	//
	// zh
	EnvLanguage *string `json:"EnvLanguage,omitempty" xml:"EnvLanguage,omitempty"`
	// example:
	//
	// 2021-01-20T09:25:56Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 0
	NoticeType *int64 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// q_security-groups
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
}

func (s ModifyTemplateQuotaItemRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTemplateQuotaItemRequest) GoString() string {
	return s.String()
}

func (s *ModifyTemplateQuotaItemRequest) SetDesireValue(v float32) *ModifyTemplateQuotaItemRequest {
	s.DesireValue = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) SetDimensions(v []*ModifyTemplateQuotaItemRequestDimensions) *ModifyTemplateQuotaItemRequest {
	s.Dimensions = v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) SetEffectiveTime(v string) *ModifyTemplateQuotaItemRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) SetEnvLanguage(v string) *ModifyTemplateQuotaItemRequest {
	s.EnvLanguage = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) SetExpireTime(v string) *ModifyTemplateQuotaItemRequest {
	s.ExpireTime = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) SetId(v string) *ModifyTemplateQuotaItemRequest {
	s.Id = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) SetNoticeType(v int64) *ModifyTemplateQuotaItemRequest {
	s.NoticeType = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) SetProductCode(v string) *ModifyTemplateQuotaItemRequest {
	s.ProductCode = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) SetQuotaActionCode(v string) *ModifyTemplateQuotaItemRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) SetQuotaCategory(v string) *ModifyTemplateQuotaItemRequest {
	s.QuotaCategory = &v
	return s
}

type ModifyTemplateQuotaItemRequestDimensions struct {
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyTemplateQuotaItemRequestDimensions) String() string {
	return tea.Prettify(s)
}

func (s ModifyTemplateQuotaItemRequestDimensions) GoString() string {
	return s.String()
}

func (s *ModifyTemplateQuotaItemRequestDimensions) SetKey(v string) *ModifyTemplateQuotaItemRequestDimensions {
	s.Key = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequestDimensions) SetValue(v string) *ModifyTemplateQuotaItemRequestDimensions {
	s.Value = &v
	return s
}

type ModifyTemplateQuotaItemResponseBody struct {
	// example:
	//
	// 1****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// D47B3A10-CDAC-5412-B2EE-EC9A3DBE9053
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTemplateQuotaItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTemplateQuotaItemResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTemplateQuotaItemResponseBody) SetId(v string) *ModifyTemplateQuotaItemResponseBody {
	s.Id = &v
	return s
}

func (s *ModifyTemplateQuotaItemResponseBody) SetRequestId(v string) *ModifyTemplateQuotaItemResponseBody {
	s.RequestId = &v
	return s
}

type ModifyTemplateQuotaItemResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTemplateQuotaItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTemplateQuotaItemResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTemplateQuotaItemResponse) GoString() string {
	return s.String()
}

func (s *ModifyTemplateQuotaItemResponse) SetHeaders(v map[string]*string) *ModifyTemplateQuotaItemResponse {
	s.Headers = v
	return s
}

func (s *ModifyTemplateQuotaItemResponse) SetStatusCode(v int32) *ModifyTemplateQuotaItemResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTemplateQuotaItemResponse) SetBody(v *ModifyTemplateQuotaItemResponseBody) *ModifyTemplateQuotaItemResponse {
	s.Body = v
	return s
}

type RemindQuotaApplicationApprovalRequest struct {
	// The quota application ID.
	//
	// For more information about how to obtain the ID of a quota application, see [ListQuotaApplications](https://help.aliyun.com/document_detail/440568.html).
	//
	// example:
	//
	// 219f1ff6-6205-495f-bda7-90d2fe945e****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s RemindQuotaApplicationApprovalRequest) String() string {
	return tea.Prettify(s)
}

func (s RemindQuotaApplicationApprovalRequest) GoString() string {
	return s.String()
}

func (s *RemindQuotaApplicationApprovalRequest) SetApplicationId(v string) *RemindQuotaApplicationApprovalRequest {
	s.ApplicationId = &v
	return s
}

type RemindQuotaApplicationApprovalResponseBody struct {
	// Indicates whether retries are allowed. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// false
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// PARAMETER.ILLEGALL
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// Parameter[%s] is required.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The parameters whose values are invalid.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// RAM.PERMISSION.DENIED
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// You are not authorized to do this action or the API input parameter is invalid.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The quota application ID.
	//
	// example:
	//
	// 219f1ff6-6205-495f-bda7-90d2fe945e****
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D47B3A10-CDAC-5412-B2EE-EC9A3DBE9053
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
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemindQuotaApplicationApprovalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemindQuotaApplicationApprovalResponseBody) GoString() string {
	return s.String()
}

func (s *RemindQuotaApplicationApprovalResponseBody) SetAllowRetry(v bool) *RemindQuotaApplicationApprovalResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *RemindQuotaApplicationApprovalResponseBody) SetDynamicCode(v string) *RemindQuotaApplicationApprovalResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *RemindQuotaApplicationApprovalResponseBody) SetDynamicMessage(v string) *RemindQuotaApplicationApprovalResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *RemindQuotaApplicationApprovalResponseBody) SetErrorArgs(v []interface{}) *RemindQuotaApplicationApprovalResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *RemindQuotaApplicationApprovalResponseBody) SetErrorCode(v string) *RemindQuotaApplicationApprovalResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RemindQuotaApplicationApprovalResponseBody) SetErrorMsg(v string) *RemindQuotaApplicationApprovalResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *RemindQuotaApplicationApprovalResponseBody) SetHttpStatusCode(v int32) *RemindQuotaApplicationApprovalResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemindQuotaApplicationApprovalResponseBody) SetModule(v string) *RemindQuotaApplicationApprovalResponseBody {
	s.Module = &v
	return s
}

func (s *RemindQuotaApplicationApprovalResponseBody) SetRequestId(v string) *RemindQuotaApplicationApprovalResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemindQuotaApplicationApprovalResponseBody) SetSuccess(v bool) *RemindQuotaApplicationApprovalResponseBody {
	s.Success = &v
	return s
}

type RemindQuotaApplicationApprovalResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemindQuotaApplicationApprovalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemindQuotaApplicationApprovalResponse) String() string {
	return tea.Prettify(s)
}

func (s RemindQuotaApplicationApprovalResponse) GoString() string {
	return s.String()
}

func (s *RemindQuotaApplicationApprovalResponse) SetHeaders(v map[string]*string) *RemindQuotaApplicationApprovalResponse {
	s.Headers = v
	return s
}

func (s *RemindQuotaApplicationApprovalResponse) SetStatusCode(v int32) *RemindQuotaApplicationApprovalResponse {
	s.StatusCode = &v
	return s
}

func (s *RemindQuotaApplicationApprovalResponse) SetBody(v *RemindQuotaApplicationApprovalResponseBody) *RemindQuotaApplicationApprovalResponse {
	s.Body = v
	return s
}

type UpdateQuotaAlarmRequest struct {
	// The ID of the quota alert.
	//
	// >  You can call the [ListQuotaAlarms](https://help.aliyun.com/document_detail/440561.html) operation to obtain the ID of a quota alert.
	//
	// This parameter is required.
	//
	// example:
	//
	// a2efa7fc-832f-47bb-8054-39e28012****
	AlarmId *string `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// The name of the quota alert.
	//
	// >  You can call the [ListQuotaAlarms](https://help.aliyun.com/document_detail/440561.html) operation to obtain the name of a quota alert.
	//
	// This parameter is required.
	//
	// example:
	//
	// rules
	AlarmName *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	// The numeric value of the alert threshold. Valid values:
	//
	// 	- If you set the `ThresholdType` parameter to `used`, you will receive an alert notification when the used quota is greater than or equal to the preset alert threshold. The alert threshold must be greater than the current used quota.
	//
	// 	- If you set the `ThresholdType` parameter to `usable`, you will receive an alert notification when the available quota is less than or equal to the preset alert threshold. The alert threshold must be less than the current available quota.
	//
	// > You must set one of the Threshold and ThresholdPercent parameters.
	//
	// example:
	//
	// 160
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The percentage of the alert threshold. Valid values:
	//
	// 	- If you set `ThresholdType` to `used`, you receive an alert notification when the used quota is greater than or equal to the preset percentage of the alert threshold. Value range: (50%, 100%].
	//
	// 	- If you set `ThresholdType` to `usable`, you receive an alert notification when the available quota is less than or equal to the preset percentage of the alert threshold. Value range: (0%, 50%].
	//
	// >  You must set one of Threshold and ThresholdPercent.
	//
	// example:
	//
	// 51
	ThresholdPercent *float32 `json:"ThresholdPercent,omitempty" xml:"ThresholdPercent,omitempty"`
	// The type of the quota alert. Valid values:
	//
	// 	- used (default): The alert is created for the used quota.
	//
	// 	- usable: The alert is created for the available quota.
	//
	// example:
	//
	// usable
	ThresholdType *string `json:"ThresholdType,omitempty" xml:"ThresholdType,omitempty"`
	// The webhook URL. Quota Center sends alert notifications to the specified URL by using HTTP POST requests.
	//
	// example:
	//
	// https://alert.aliyun.com/callback
	WebHook *string `json:"WebHook,omitempty" xml:"WebHook,omitempty"`
}

func (s UpdateQuotaAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaAlarmRequest) GoString() string {
	return s.String()
}

func (s *UpdateQuotaAlarmRequest) SetAlarmId(v string) *UpdateQuotaAlarmRequest {
	s.AlarmId = &v
	return s
}

func (s *UpdateQuotaAlarmRequest) SetAlarmName(v string) *UpdateQuotaAlarmRequest {
	s.AlarmName = &v
	return s
}

func (s *UpdateQuotaAlarmRequest) SetThreshold(v float32) *UpdateQuotaAlarmRequest {
	s.Threshold = &v
	return s
}

func (s *UpdateQuotaAlarmRequest) SetThresholdPercent(v float32) *UpdateQuotaAlarmRequest {
	s.ThresholdPercent = &v
	return s
}

func (s *UpdateQuotaAlarmRequest) SetThresholdType(v string) *UpdateQuotaAlarmRequest {
	s.ThresholdType = &v
	return s
}

func (s *UpdateQuotaAlarmRequest) SetWebHook(v string) *UpdateQuotaAlarmRequest {
	s.WebHook = &v
	return s
}

type UpdateQuotaAlarmResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A95C65B3-7CF4-469E-B1D5-1CA0628A6411
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateQuotaAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQuotaAlarmResponseBody) SetRequestId(v string) *UpdateQuotaAlarmResponseBody {
	s.RequestId = &v
	return s
}

type UpdateQuotaAlarmResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateQuotaAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateQuotaAlarmResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaAlarmResponse) GoString() string {
	return s.String()
}

func (s *UpdateQuotaAlarmResponse) SetHeaders(v map[string]*string) *UpdateQuotaAlarmResponse {
	s.Headers = v
	return s
}

func (s *UpdateQuotaAlarmResponse) SetStatusCode(v int32) *UpdateQuotaAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateQuotaAlarmResponse) SetBody(v *UpdateQuotaAlarmResponseBody) *UpdateQuotaAlarmResponse {
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
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("quotas"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// Summary:
//
// The value of the quota dimension.
//
// The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
//
// > This parameter is required if you set the ProductCode parameter to ecs, ecs-spec, actiontrail, or ess.
//
// Description:
//
// The ID of the alert.
//
// @param request - CreateQuotaAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQuotaAlarmResponse
func (client *Client) CreateQuotaAlarmWithOptions(request *CreateQuotaAlarmRequest, runtime *util.RuntimeOptions) (_result *CreateQuotaAlarmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmName)) {
		body["AlarmName"] = request.AlarmName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaActionCode)) {
		body["QuotaActionCode"] = request.QuotaActionCode
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaDimensions)) {
		body["QuotaDimensions"] = request.QuotaDimensions
	}

	if !tea.BoolValue(util.IsUnset(request.Threshold)) {
		body["Threshold"] = request.Threshold
	}

	if !tea.BoolValue(util.IsUnset(request.ThresholdPercent)) {
		body["ThresholdPercent"] = request.ThresholdPercent
	}

	if !tea.BoolValue(util.IsUnset(request.ThresholdType)) {
		body["ThresholdType"] = request.ThresholdType
	}

	if !tea.BoolValue(util.IsUnset(request.WebHook)) {
		body["WebHook"] = request.WebHook
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateQuotaAlarm"),
		Version:     tea.String("2020-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateQuotaAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The value of the quota dimension.
//
// The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
//
// > This parameter is required if you set the ProductCode parameter to ecs, ecs-spec, actiontrail, or ess.
//
// Description:
//
// The ID of the alert.
//
// @param request - CreateQuotaAlarmRequest
//
// @return CreateQuotaAlarmResponse
func (client *Client) CreateQuotaAlarm(request *CreateQuotaAlarmRequest) (_result *CreateQuotaAlarmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateQuotaAlarmResponse{}
	_body, _err := client.CreateQuotaAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits an application to increase a quota.
//
// Description:
//
// In this example, the operation is called to submit an application to increase the value of a quota whose ID is `q_security-groups` and whose name is Maximum Number of Security Groups. The quota belongs to Elastic Compute Service (ECS). The expected value of the quota is `804`, the application reason is `Scale Out`, and the ID of the region to which the quota belongs is `cn-hangzhou`.
//
// @param request - CreateQuotaApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQuotaApplicationResponse
func (client *Client) CreateQuotaApplicationWithOptions(request *CreateQuotaApplicationRequest, runtime *util.RuntimeOptions) (_result *CreateQuotaApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditMode)) {
		body["AuditMode"] = request.AuditMode
	}

	if !tea.BoolValue(util.IsUnset(request.DesireValue)) {
		body["DesireValue"] = request.DesireValue
	}

	if !tea.BoolValue(util.IsUnset(request.Dimensions)) {
		body["Dimensions"] = request.Dimensions
	}

	if !tea.BoolValue(util.IsUnset(request.EffectiveTime)) {
		body["EffectiveTime"] = request.EffectiveTime
	}

	if !tea.BoolValue(util.IsUnset(request.EnvLanguage)) {
		body["EnvLanguage"] = request.EnvLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireTime)) {
		body["ExpireTime"] = request.ExpireTime
	}

	if !tea.BoolValue(util.IsUnset(request.NoticeType)) {
		body["NoticeType"] = request.NoticeType
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaActionCode)) {
		body["QuotaActionCode"] = request.QuotaActionCode
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaCategory)) {
		body["QuotaCategory"] = request.QuotaCategory
	}

	if !tea.BoolValue(util.IsUnset(request.Reason)) {
		body["Reason"] = request.Reason
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateQuotaApplication"),
		Version:     tea.String("2020-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateQuotaApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits an application to increase a quota.
//
// Description:
//
// In this example, the operation is called to submit an application to increase the value of a quota whose ID is `q_security-groups` and whose name is Maximum Number of Security Groups. The quota belongs to Elastic Compute Service (ECS). The expected value of the quota is `804`, the application reason is `Scale Out`, and the ID of the region to which the quota belongs is `cn-hangzhou`.
//
// @param request - CreateQuotaApplicationRequest
//
// @return CreateQuotaApplicationResponse
func (client *Client) CreateQuotaApplication(request *CreateQuotaApplicationRequest) (_result *CreateQuotaApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateQuotaApplicationResponse{}
	_body, _err := client.CreateQuotaApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a quota increase application. After you add a quota item to a quota template, the system automatically submits quota applications only for new members in the resource directory. The quota values for existing members remain unchanged. If you want to increase the quota values of existing members, you can submit a quota application for the members by applying quota templates to the members.
//
// Description:
//
// ### [](#)QPS limit
//
// You can add a maximum of 10 quota items to a quota template at a time.
//
// @param request - CreateQuotaApplicationsForTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQuotaApplicationsForTemplateResponse
func (client *Client) CreateQuotaApplicationsForTemplateWithOptions(request *CreateQuotaApplicationsForTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateQuotaApplicationsForTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunUids)) {
		body["AliyunUids"] = request.AliyunUids
	}

	if !tea.BoolValue(util.IsUnset(request.DesireValue)) {
		body["DesireValue"] = request.DesireValue
	}

	if !tea.BoolValue(util.IsUnset(request.Dimensions)) {
		body["Dimensions"] = request.Dimensions
	}

	if !tea.BoolValue(util.IsUnset(request.EffectiveTime)) {
		body["EffectiveTime"] = request.EffectiveTime
	}

	if !tea.BoolValue(util.IsUnset(request.EnvLanguage)) {
		body["EnvLanguage"] = request.EnvLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireTime)) {
		body["ExpireTime"] = request.ExpireTime
	}

	if !tea.BoolValue(util.IsUnset(request.NoticeType)) {
		body["NoticeType"] = request.NoticeType
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaActionCode)) {
		body["QuotaActionCode"] = request.QuotaActionCode
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaCategory)) {
		body["QuotaCategory"] = request.QuotaCategory
	}

	if !tea.BoolValue(util.IsUnset(request.Reason)) {
		body["Reason"] = request.Reason
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateQuotaApplicationsForTemplate"),
		Version:     tea.String("2020-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateQuotaApplicationsForTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a quota increase application. After you add a quota item to a quota template, the system automatically submits quota applications only for new members in the resource directory. The quota values for existing members remain unchanged. If you want to increase the quota values of existing members, you can submit a quota application for the members by applying quota templates to the members.
//
// Description:
//
// ### [](#)QPS limit
//
// You can add a maximum of 10 quota items to a quota template at a time.
//
// @param request - CreateQuotaApplicationsForTemplateRequest
//
// @return CreateQuotaApplicationsForTemplateResponse
func (client *Client) CreateQuotaApplicationsForTemplate(request *CreateQuotaApplicationsForTemplateRequest) (_result *CreateQuotaApplicationsForTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateQuotaApplicationsForTemplateResponse{}
	_body, _err := client.CreateQuotaApplicationsForTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a quota template by using the management account of a resource directory. After you create a quota template, if a member is added to the resource directory, the quota template automatically submits a quota increase request for the member. The quota values for existing members remain unchanged. You can use a quota template to apply for increases on multiple quotas at the same time. This automated approach improves the efficiency of quota management across your organization.
//
// Description:
//
// ### [](#)Prerequisites
//
// You must set the `ServiceStatus` parameter to `1`. This ensures that the quota template is enabled.
//
// You can call the [GetQuotaTemplateServiceStatus](https://help.aliyun.com/document_detail/450407.html) operation to query the status of a quota template. If the `ServiceStatus` parameter is set to `0` or `-1`, you must call the [ModifyQuotaTemplateServiceStatus](https://help.aliyun.com/document_detail/450406.html) operation to set the ServiceStatus parameter to `1`.
//
// @param request - CreateTemplateQuotaItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTemplateQuotaItemResponse
func (client *Client) CreateTemplateQuotaItemWithOptions(request *CreateTemplateQuotaItemRequest, runtime *util.RuntimeOptions) (_result *CreateTemplateQuotaItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesireValue)) {
		body["DesireValue"] = request.DesireValue
	}

	if !tea.BoolValue(util.IsUnset(request.Dimensions)) {
		body["Dimensions"] = request.Dimensions
	}

	if !tea.BoolValue(util.IsUnset(request.EffectiveTime)) {
		body["EffectiveTime"] = request.EffectiveTime
	}

	if !tea.BoolValue(util.IsUnset(request.EnvLanguage)) {
		body["EnvLanguage"] = request.EnvLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireTime)) {
		body["ExpireTime"] = request.ExpireTime
	}

	if !tea.BoolValue(util.IsUnset(request.NoticeType)) {
		body["NoticeType"] = request.NoticeType
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaActionCode)) {
		body["QuotaActionCode"] = request.QuotaActionCode
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaCategory)) {
		body["QuotaCategory"] = request.QuotaCategory
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTemplateQuotaItem"),
		Version:     tea.String("2020-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTemplateQuotaItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a quota template by using the management account of a resource directory. After you create a quota template, if a member is added to the resource directory, the quota template automatically submits a quota increase request for the member. The quota values for existing members remain unchanged. You can use a quota template to apply for increases on multiple quotas at the same time. This automated approach improves the efficiency of quota management across your organization.
//
// Description:
//
// ### [](#)Prerequisites
//
// You must set the `ServiceStatus` parameter to `1`. This ensures that the quota template is enabled.
//
// You can call the [GetQuotaTemplateServiceStatus](https://help.aliyun.com/document_detail/450407.html) operation to query the status of a quota template. If the `ServiceStatus` parameter is set to `0` or `-1`, you must call the [ModifyQuotaTemplateServiceStatus](https://help.aliyun.com/document_detail/450406.html) operation to set the ServiceStatus parameter to `1`.
//
// @param request - CreateTemplateQuotaItemRequest
//
// @return CreateTemplateQuotaItemResponse
func (client *Client) CreateTemplateQuotaItem(request *CreateTemplateQuotaItemRequest) (_result *CreateTemplateQuotaItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTemplateQuotaItemResponse{}
	_body, _err := client.CreateTemplateQuotaItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a quota alert.
//
// Description:
//
// In this example, the operation is called to delete a quota alert whose ID is `6b512ab7-da3a-4142-b529-2b2a9294****`.
//
// @param request - DeleteQuotaAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQuotaAlarmResponse
func (client *Client) DeleteQuotaAlarmWithOptions(request *DeleteQuotaAlarmRequest, runtime *util.RuntimeOptions) (_result *DeleteQuotaAlarmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmId)) {
		body["AlarmId"] = request.AlarmId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteQuotaAlarm"),
		Version:     tea.String("2020-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteQuotaAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a quota alert.
//
// Description:
//
// In this example, the operation is called to delete a quota alert whose ID is `6b512ab7-da3a-4142-b529-2b2a9294****`.
//
// @param request - DeleteQuotaAlarmRequest
//
// @return DeleteQuotaAlarmResponse
func (client *Client) DeleteQuotaAlarm(request *DeleteQuotaAlarmRequest) (_result *DeleteQuotaAlarmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteQuotaAlarmResponse{}
	_body, _err := client.DeleteQuotaAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a quota template by using the management account of a resource directory. After you delete a quota template, if a member is added to the resource directory, the quota template no longer automatically submits a quota increase request for the member.
//
// @param request - DeleteTemplateQuotaItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTemplateQuotaItemResponse
func (client *Client) DeleteTemplateQuotaItemWithOptions(request *DeleteTemplateQuotaItemRequest, runtime *util.RuntimeOptions) (_result *DeleteTemplateQuotaItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTemplateQuotaItem"),
		Version:     tea.String("2020-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTemplateQuotaItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a quota template by using the management account of a resource directory. After you delete a quota template, if a member is added to the resource directory, the quota template no longer automatically submits a quota increase request for the member.
//
// @param request - DeleteTemplateQuotaItemRequest
//
// @return DeleteTemplateQuotaItemResponse
func (client *Client) DeleteTemplateQuotaItem(request *DeleteTemplateQuotaItemRequest) (_result *DeleteTemplateQuotaItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTemplateQuotaItemResponse{}
	_body, _err := client.DeleteTemplateQuotaItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the specified quota.
//
// Description:
//
// In this example, the operation is called to query the details of a quota whose ID is `q_security-groups` and whose name is Maximum Number of Security Groups. This quota belongs to Elastic Compute Service (ECS). The query result shows the details of the quota. The details include the name, ID, description, quota value, used quota, unit, and dimension of the quota. In this example, the quota name is `Maximum Number of Security Groups`. The quota ID is `q_security-groups`. The description is `The maximum number of security groups that can be created for the current account`. The quota value is `801`. The used quota is `26`. The quota unit is `Number of security groups`. The quota dimension is `{"regionId":"cn-hangzhou"}`.
//
// @param request - GetProductQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProductQuotaResponse
func (client *Client) GetProductQuotaWithOptions(request *GetProductQuotaRequest, runtime *util.RuntimeOptions) (_result *GetProductQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Dimensions)) {
		body["Dimensions"] = request.Dimensions
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaActionCode)) {
		body["QuotaActionCode"] = request.QuotaActionCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProductQuota"),
		Version:     tea.String("2020-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProductQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the specified quota.
//
// Description:
//
// In this example, the operation is called to query the details of a quota whose ID is `q_security-groups` and whose name is Maximum Number of Security Groups. This quota belongs to Elastic Compute Service (ECS). The query result shows the details of the quota. The details include the name, ID, description, quota value, used quota, unit, and dimension of the quota. In this example, the quota name is `Maximum Number of Security Groups`. The quota ID is `q_security-groups`. The description is `The maximum number of security groups that can be created for the current account`. The quota value is `801`. The used quota is `26`. The quota unit is `Number of security groups`. The quota dimension is `{"regionId":"cn-hangzhou"}`.
//
// @param request - GetProductQuotaRequest
//
// @return GetProductQuotaResponse
func (client *Client) GetProductQuota(request *GetProductQuotaRequest) (_result *GetProductQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProductQuotaResponse{}
	_body, _err := client.GetProductQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a quota dimension that is supported by an Alibaba Cloud service.
//
// Description:
//
// In this example, the operation is called to query the details of a quota dimension whose key is `regionId`. The quota dimension belongs to Elastic Compute Service (ECS) Quotas by Instance Type whose service code is ecs-spec. The following query results are returned:
//
// 	- The values of the quota dimension include `cn-shenzhen`, `cn-beijing`, and `cn-hangzhou`.
//
// 	- The name of the quota dimension is `region`.
//
// @param request - GetProductQuotaDimensionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProductQuotaDimensionResponse
func (client *Client) GetProductQuotaDimensionWithOptions(request *GetProductQuotaDimensionRequest, runtime *util.RuntimeOptions) (_result *GetProductQuotaDimensionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DependentDimensions)) {
		body["DependentDimensions"] = request.DependentDimensions
	}

	if !tea.BoolValue(util.IsUnset(request.DimensionKey)) {
		body["DimensionKey"] = request.DimensionKey
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProductQuotaDimension"),
		Version:     tea.String("2020-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProductQuotaDimensionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a quota dimension that is supported by an Alibaba Cloud service.
//
// Description:
//
// In this example, the operation is called to query the details of a quota dimension whose key is `regionId`. The quota dimension belongs to Elastic Compute Service (ECS) Quotas by Instance Type whose service code is ecs-spec. The following query results are returned:
//
// 	- The values of the quota dimension include `cn-shenzhen`, `cn-beijing`, and `cn-hangzhou`.
//
// 	- The name of the quota dimension is `region`.
//
// @param request - GetProductQuotaDimensionRequest
//
// @return GetProductQuotaDimensionResponse
func (client *Client) GetProductQuotaDimension(request *GetProductQuotaDimensionRequest) (_result *GetProductQuotaDimensionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProductQuotaDimensionResponse{}
	_body, _err := client.GetProductQuotaDimensionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// In this example, the operation is called to query the details of a quota alert. The details of the alert are returned. The query results include the alert ID, alert name, alert contact, and time when the quota alert was created.
//
// Description:
//
// In this example, the operation is called to query the details of a quota alert whose ID is `78d7e436-4b25-4897-84b5-d7b656bb****`. The details of the alert are returned. The query result includes the alert ID, alert name, alert contact, and the time when the quota alert was created.
//
// @param request - GetQuotaAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQuotaAlarmResponse
func (client *Client) GetQuotaAlarmWithOptions(request *GetQuotaAlarmRequest, runtime *util.RuntimeOptions) (_result *GetQuotaAlarmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmId)) {
		body["AlarmId"] = request.AlarmId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQuotaAlarm"),
		Version:     tea.String("2020-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQuotaAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// In this example, the operation is called to query the details of a quota alert. The details of the alert are returned. The query results include the alert ID, alert name, alert contact, and time when the quota alert was created.
//
// Description:
//
// In this example, the operation is called to query the details of a quota alert whose ID is `78d7e436-4b25-4897-84b5-d7b656bb****`. The details of the alert are returned. The query result includes the alert ID, alert name, alert contact, and the time when the quota alert was created.
//
// @param request - GetQuotaAlarmRequest
//
// @return GetQuotaAlarmResponse
func (client *Client) GetQuotaAlarm(request *GetQuotaAlarmRequest) (_result *GetQuotaAlarmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQuotaAlarmResponse{}
	_body, _err := client.GetQuotaAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about a specified application that is submitted to increase a quota.
//
// Description:
//
// In this example, the operation is called to query the details about an application whose ID is `d314d6ae-867d-484c-9009-3d421a80****`. The query result shows the details about the application. The details include the application ID, application time, expected quota value, and application result.
//
// @param request - GetQuotaApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQuotaApplicationResponse
func (client *Client) GetQuotaApplicationWithOptions(request *GetQuotaApplicationRequest, runtime *util.RuntimeOptions) (_result *GetQuotaApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		body["ApplicationId"] = request.ApplicationId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQuotaApplication"),
		Version:     tea.String("2020-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQuotaApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about a specified application that is submitted to increase a quota.
//
// Description:
//
// In this example, the operation is called to query the details about an application whose ID is `d314d6ae-867d-484c-9009-3d421a80****`. The query result shows the details about the application. The details include the application ID, application time, expected quota value, and application result.
//
// @param request - GetQuotaApplicationRequest
//
// @return GetQuotaApplicationResponse
func (client *Client) GetQuotaApplication(request *GetQuotaApplicationRequest) (_result *GetQuotaApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQuotaApplicationResponse{}
	_body, _err := client.GetQuotaApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about quota application approval, such as the average amount of time required for approval, whether approval reminders are supported, and the interval between two consecutive approval reminders.
//
// Description:
//
// ### [](#)Prerequisites
//
// Make sure that you have created an application for quota increase. For more information, see [CreateQuotaApplication](https://help.aliyun.com/document_detail/440566.html).
//
// @param request - GetQuotaApplicationApprovalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQuotaApplicationApprovalResponse
func (client *Client) GetQuotaApplicationApprovalWithOptions(request *GetQuotaApplicationApprovalRequest, runtime *util.RuntimeOptions) (_result *GetQuotaApplicationApprovalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		body["ApplicationId"] = request.ApplicationId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQuotaApplicationApproval"),
		Version:     tea.String("2020-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQuotaApplicationApprovalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about quota application approval, such as the average amount of time required for approval, whether approval reminders are supported, and the interval between two consecutive approval reminders.
//
// Description:
//
// ### [](#)Prerequisites
//
// Make sure that you have created an application for quota increase. For more information, see [CreateQuotaApplication](https://help.aliyun.com/document_detail/440566.html).
//
// @param request - GetQuotaApplicationApprovalRequest
//
// @return GetQuotaApplicationApprovalResponse
func (client *Client) GetQuotaApplicationApproval(request *GetQuotaApplicationApprovalRequest) (_result *GetQuotaApplicationApprovalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQuotaApplicationApprovalResponse{}
	_body, _err := client.GetQuotaApplicationApprovalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of a quota template.
//
// @param request - GetQuotaTemplateServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQuotaTemplateServiceStatusResponse
func (client *Client) GetQuotaTemplateServiceStatusWithOptions(request *GetQuotaTemplateServiceStatusRequest, runtime *util.RuntimeOptions) (_result *GetQuotaTemplateServiceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceDirectoryId)) {
		body["ResourceDirectoryId"] = request.ResourceDirectoryId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQuotaTemplateServiceStatus"),
		Version:     tea.String("2020-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQuotaTemplateServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a quota template.
//
// @param request - GetQuotaTemplateServiceStatusRequest
//
// @return GetQuotaTemplateServiceStatusResponse
func (client *Client) GetQuotaTemplateServiceStatus(request *GetQuotaTemplateServiceStatusRequest) (_result *GetQuotaTemplateServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQuotaTemplateServiceStatusResponse{}
	_body, _err := client.GetQuotaTemplateServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the alert records.
//
// @param request - ListAlarmHistoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlarmHistoriesResponse
func (client *Client) ListAlarmHistoriesWithOptions(request *ListAlarmHistoriesRequest, runtime *util.RuntimeOptions) (_result *ListAlarmHistoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmId)) {
		body["AlarmId"] = request.AlarmId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		body["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAlarmHistories"),
		Version:     tea.String("2020-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAlarmHistoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the alert records.
//
// @param request - ListAlarmHistoriesRequest
//
// @return ListAlarmHistoriesResponse
func (client *Client) ListAlarmHistories(request *ListAlarmHistoriesRequest) (_result *ListAlarmHistoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAlarmHistoriesResponse{}
	_body, _err := client.ListAlarmHistoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the quotas on which a specified quota depends.
//
// Description:
//
// In this example, the operation is called to query the quotas on which a Container Service for Kubernetes (ACK) quota whose ID is `q_i5uzm3` depends. This quota is the maximum number of nodes that can be created in an ACK cluster. The query result indicates that the specified quota depends on the following three quotas:
//
// 	- An Elastic Compute Service (ECS) quota whose ID is `q_elastic-network-interfaces`. This quota is the maximum number of ENIs (Secondary ENIs) that can be owned by an Alibaba Cloud account. The quota is available in the following regions: `cn-shenzhen`, `cn-beijing`, and `cn-hangzhou`.
//
// 	- A Server Load Balancer (SLB) quota whose ID is `q_fh20b0`. This quota is the number of servers that can be attached to the backend of an SLB instance.
//
// 	- An SLB quota whose ID is `q_3mmbsp`. This quota is the number of SLB instances that can be owned by an Alibaba Cloud account.
//
// @param request - ListDependentQuotasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDependentQuotasResponse
func (client *Client) ListDependentQuotasWithOptions(request *ListDependentQuotasRequest, runtime *util.RuntimeOptions) (_result *ListDependentQuotasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaActionCode)) {
		body["QuotaActionCode"] = request.QuotaActionCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDependentQuotas"),
		Version:     tea.String("2020-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDependentQuotasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the quotas on which a specified quota depends.
//
// Description:
//
// In this example, the operation is called to query the quotas on which a Container Service for Kubernetes (ACK) quota whose ID is `q_i5uzm3` depends. This quota is the maximum number of nodes that can be created in an ACK cluster. The query result indicates that the specified quota depends on the following three quotas:
//
// 	- An Elastic Compute Service (ECS) quota whose ID is `q_elastic-network-interfaces`. This quota is the maximum number of ENIs (Secondary ENIs) that can be owned by an Alibaba Cloud account. The quota is available in the following regions: `cn-shenzhen`, `cn-beijing`, and `cn-hangzhou`.
//
// 	- A Server Load Balancer (SLB) quota whose ID is `q_fh20b0`. This quota is the number of servers that can be attached to the backend of an SLB instance.
//
// 	- An SLB quota whose ID is `q_3mmbsp`. This quota is the number of SLB instances that can be owned by an Alibaba Cloud account.
//
// @param request - ListDependentQuotasRequest
//
// @return ListDependentQuotasResponse
func (client *Client) ListDependentQuotas(request *ListDependentQuotasRequest) (_result *ListDependentQuotasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDependentQuotasResponse{}
	_body, _err := client.ListDependentQuotasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the dimension groups of an Alibaba Cloud service.
//
// Description:
//
// This topic provides an example on how to call the ListProductDimensionGroups operation to query the dimension groups of Object Storage Service (OSS). In this example, a dimension group is returned. The group name is `OSS_Group`, the group code is `oss_wf1ngqmd7q`, and the group key is `chargeType`.
//
// @param request - ListProductDimensionGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductDimensionGroupsResponse
func (client *Client) ListProductDimensionGroupsWithOptions(request *ListProductDimensionGroupsRequest, runtime *util.RuntimeOptions) (_result *ListProductDimensionGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProductDimensionGroups"),
		Version:     tea.String("2020-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProductDimensionGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the dimension groups of an Alibaba Cloud service.
//
// Description:
//
// This topic provides an example on how to call the ListProductDimensionGroups operation to query the dimension groups of Object Storage Service (OSS). In this example, a dimension group is returned. The group name is `OSS_Group`, the group code is `oss_wf1ngqmd7q`, and the group key is `chargeType`.
//
// @param request - ListProductDimensionGroupsRequest
//
// @return ListProductDimensionGroupsResponse
func (client *Client) ListProductDimensionGroups(request *ListProductDimensionGroupsRequest) (_result *ListProductDimensionGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProductDimensionGroupsResponse{}
	_body, _err := client.ListProductDimensionGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the quota dimensions that are supported by the specified Alibaba Cloud service.
//
// Description:
//
// In this example, the operation is called to query the quota dimensions that are supported by Elastic Compute Service (ECS). The query results show all the quota dimensions that are supported by ECS.
//
// @param request - ListProductQuotaDimensionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductQuotaDimensionsResponse
func (client *Client) ListProductQuotaDimensionsWithOptions(request *ListProductQuotaDimensionsRequest, runtime *util.RuntimeOptions) (_result *ListProductQuotaDimensionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaCategory)) {
		body["QuotaCategory"] = request.QuotaCategory
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProductQuotaDimensions"),
		Version:     tea.String("2020-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProductQuotaDimensionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the quota dimensions that are supported by the specified Alibaba Cloud service.
//
// Description:
//
// In this example, the operation is called to query the quota dimensions that are supported by Elastic Compute Service (ECS). The query results show all the quota dimensions that are supported by ECS.
//
// @param request - ListProductQuotaDimensionsRequest
//
// @return ListProductQuotaDimensionsResponse
func (client *Client) ListProductQuotaDimensions(request *ListProductQuotaDimensionsRequest) (_result *ListProductQuotaDimensionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProductQuotaDimensionsResponse{}
	_body, _err := client.ListProductQuotaDimensionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the quotas of a specific Alibaba Cloud service.
//
// Description:
//
// In this example, the operation is called to query the quotas whose instance type is `ecs.g5.2xlarge`. The quotas belong to Elastic Compute Service (ECS) Quotas by Instance Type. The query result includes the name, ID, unit, dimensions, and cycle of each quota.
//
// @param request - ListProductQuotasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductQuotasResponse
func (client *Client) ListProductQuotasWithOptions(request *ListProductQuotasRequest, runtime *util.RuntimeOptions) (_result *ListProductQuotasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Dimensions)) {
		body["Dimensions"] = request.Dimensions
	}

	if !tea.BoolValue(util.IsUnset(request.GroupCode)) {
		body["GroupCode"] = request.GroupCode
	}

	if !tea.BoolValue(util.IsUnset(request.KeyWord)) {
		body["KeyWord"] = request.KeyWord
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaActionCode)) {
		body["QuotaActionCode"] = request.QuotaActionCode
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaCategory)) {
		body["QuotaCategory"] = request.QuotaCategory
	}

	if !tea.BoolValue(util.IsUnset(request.SortField)) {
		body["SortField"] = request.SortField
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		body["SortOrder"] = request.SortOrder
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProductQuotas"),
		Version:     tea.String("2020-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProductQuotasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the quotas of a specific Alibaba Cloud service.
//
// Description:
//
// In this example, the operation is called to query the quotas whose instance type is `ecs.g5.2xlarge`. The quotas belong to Elastic Compute Service (ECS) Quotas by Instance Type. The query result includes the name, ID, unit, dimensions, and cycle of each quota.
//
// @param request - ListProductQuotasRequest
//
// @return ListProductQuotasResponse
func (client *Client) ListProductQuotas(request *ListProductQuotasRequest) (_result *ListProductQuotasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProductQuotasResponse{}
	_body, _err := client.ListProductQuotasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Alibaba Cloud services that support Quota Center.
//
// Description:
//
// The services in the query result are the same as the services listed in [Alibaba Cloud services that support Quota Center](https://help.aliyun.com/document_detail/182368.html).
//
// @param request - ListProductsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductsResponse
func (client *Client) ListProductsWithOptions(request *ListProductsRequest, runtime *util.RuntimeOptions) (_result *ListProductsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProducts"),
		Version:     tea.String("2020-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProductsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Alibaba Cloud services that support Quota Center.
//
// Description:
//
// The services in the query result are the same as the services listed in [Alibaba Cloud services that support Quota Center](https://help.aliyun.com/document_detail/182368.html).
//
// @param request - ListProductsRequest
//
// @return ListProductsResponse
func (client *Client) ListProducts(request *ListProductsRequest) (_result *ListProductsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProductsResponse{}
	_body, _err := client.ListProductsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries quota alerts.
//
// @param request - ListQuotaAlarmsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotaAlarmsResponse
func (client *Client) ListQuotaAlarmsWithOptions(request *ListQuotaAlarmsRequest, runtime *util.RuntimeOptions) (_result *ListQuotaAlarmsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmName)) {
		body["AlarmName"] = request.AlarmName
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaActionCode)) {
		body["QuotaActionCode"] = request.QuotaActionCode
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaDimensions)) {
		body["QuotaDimensions"] = request.QuotaDimensions
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQuotaAlarms"),
		Version:     tea.String("2020-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQuotaAlarmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries quota alerts.
//
// @param request - ListQuotaAlarmsRequest
//
// @return ListQuotaAlarmsResponse
func (client *Client) ListQuotaAlarms(request *ListQuotaAlarmsRequest) (_result *ListQuotaAlarmsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListQuotaAlarmsResponse{}
	_body, _err := client.ListQuotaAlarmsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries quota templates by using the management account of a resource directory.
//
// @param request - ListQuotaApplicationTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotaApplicationTemplatesResponse
func (client *Client) ListQuotaApplicationTemplatesWithOptions(request *ListQuotaApplicationTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListQuotaApplicationTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Dimensions)) {
		body["Dimensions"] = request.Dimensions
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaActionCode)) {
		body["QuotaActionCode"] = request.QuotaActionCode
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaCategory)) {
		body["QuotaCategory"] = request.QuotaCategory
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQuotaApplicationTemplates"),
		Version:     tea.String("2020-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQuotaApplicationTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries quota templates by using the management account of a resource directory.
//
// @param request - ListQuotaApplicationTemplatesRequest
//
// @return ListQuotaApplicationTemplatesResponse
func (client *Client) ListQuotaApplicationTemplates(request *ListQuotaApplicationTemplatesRequest) (_result *ListQuotaApplicationTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListQuotaApplicationTemplatesResponse{}
	_body, _err := client.ListQuotaApplicationTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an application that is submitted to increase a quota.
//
// Description:
//
// In this example, the operation is called to query the details of an application that is submitted to increase a quota whose ID is `q_i5uzm3` and whose name is Maximum Number of Nodes. This quota belongs to Container Service for Kubernetes (ACK). The query result shows the details of the application. The details include the application ID, application time, requested quota, and application result. In this example, the application ID is `b926571d-cc09-4711-b547-58a615f0****`. The application time is `2021-01-15T09:13:53Z`. The expected quota value is `101`. The application result is `Agree`.
//
// @param request - ListQuotaApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotaApplicationsResponse
func (client *Client) ListQuotaApplicationsWithOptions(request *ListQuotaApplicationsRequest, runtime *util.RuntimeOptions) (_result *ListQuotaApplicationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Dimensions)) {
		body["Dimensions"] = request.Dimensions
	}

	if !tea.BoolValue(util.IsUnset(request.KeyWord)) {
		body["KeyWord"] = request.KeyWord
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaActionCode)) {
		body["QuotaActionCode"] = request.QuotaActionCode
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaCategory)) {
		body["QuotaCategory"] = request.QuotaCategory
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQuotaApplications"),
		Version:     tea.String("2020-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQuotaApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an application that is submitted to increase a quota.
//
// Description:
//
// In this example, the operation is called to query the details of an application that is submitted to increase a quota whose ID is `q_i5uzm3` and whose name is Maximum Number of Nodes. This quota belongs to Container Service for Kubernetes (ACK). The query result shows the details of the application. The details include the application ID, application time, requested quota, and application result. In this example, the application ID is `b926571d-cc09-4711-b547-58a615f0****`. The application time is `2021-01-15T09:13:53Z`. The expected quota value is `101`. The application result is `Agree`.
//
// @param request - ListQuotaApplicationsRequest
//
// @return ListQuotaApplicationsResponse
func (client *Client) ListQuotaApplications(request *ListQuotaApplicationsRequest) (_result *ListQuotaApplicationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListQuotaApplicationsResponse{}
	_body, _err := client.ListQuotaApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a quota increase application for member accounts in a resource directory.
//
// @param request - ListQuotaApplicationsDetailForTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotaApplicationsDetailForTemplateResponse
func (client *Client) ListQuotaApplicationsDetailForTemplateWithOptions(request *ListQuotaApplicationsDetailForTemplateRequest, runtime *util.RuntimeOptions) (_result *ListQuotaApplicationsDetailForTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunUid)) {
		body["AliyunUid"] = request.AliyunUid
	}

	if !tea.BoolValue(util.IsUnset(request.BatchQuotaApplicationId)) {
		body["BatchQuotaApplicationId"] = request.BatchQuotaApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaActionCode)) {
		body["QuotaActionCode"] = request.QuotaActionCode
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaCategory)) {
		body["QuotaCategory"] = request.QuotaCategory
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQuotaApplicationsDetailForTemplate"),
		Version:     tea.String("2020-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQuotaApplicationsDetailForTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a quota increase application for member accounts in a resource directory.
//
// @param request - ListQuotaApplicationsDetailForTemplateRequest
//
// @return ListQuotaApplicationsDetailForTemplateResponse
func (client *Client) ListQuotaApplicationsDetailForTemplate(request *ListQuotaApplicationsDetailForTemplateRequest) (_result *ListQuotaApplicationsDetailForTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListQuotaApplicationsDetailForTemplateResponse{}
	_body, _err := client.ListQuotaApplicationsDetailForTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the application records of a quota template that is used to apply for quotas for member accounts.
//
// @param request - ListQuotaApplicationsForTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotaApplicationsForTemplateResponse
func (client *Client) ListQuotaApplicationsForTemplateWithOptions(request *ListQuotaApplicationsForTemplateRequest, runtime *util.RuntimeOptions) (_result *ListQuotaApplicationsForTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplyEndTime)) {
		body["ApplyEndTime"] = request.ApplyEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ApplyStartTime)) {
		body["ApplyStartTime"] = request.ApplyStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.BatchQuotaApplicationId)) {
		body["BatchQuotaApplicationId"] = request.BatchQuotaApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaActionCode)) {
		body["QuotaActionCode"] = request.QuotaActionCode
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaCategory)) {
		body["QuotaCategory"] = request.QuotaCategory
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQuotaApplicationsForTemplate"),
		Version:     tea.String("2020-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQuotaApplicationsForTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the application records of a quota template that is used to apply for quotas for member accounts.
//
// @param request - ListQuotaApplicationsForTemplateRequest
//
// @return ListQuotaApplicationsForTemplateResponse
func (client *Client) ListQuotaApplicationsForTemplate(request *ListQuotaApplicationsForTemplateRequest) (_result *ListQuotaApplicationsForTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListQuotaApplicationsForTemplateResponse{}
	_body, _err := client.ListQuotaApplicationsForTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the status of a quota template. By default, the quota template is not configured. If the management account of a resource directory uses a quota template for the first time, you must enable the quota template.
//
// Description:
//
// ### [](#)Prerequisites
//
// A resource directory is enabled. For more information, see [EnableResourceDirectory](https://help.aliyun.com/document_detail/604185.html).
//
// ### [](#)Usage notes
//
// If the `ServiceStatus` parameter is set to `0` or `-1`, you can call this operation to set the parameter to `1`. Then, you can call the [CreateTemplateQuotaItem](https://help.aliyun.com/document_detail/450615.html) operation to create a quota template.
//
// @param request - ModifyQuotaTemplateServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyQuotaTemplateServiceStatusResponse
func (client *Client) ModifyQuotaTemplateServiceStatusWithOptions(request *ModifyQuotaTemplateServiceStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyQuotaTemplateServiceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceStatus)) {
		body["ServiceStatus"] = request.ServiceStatus
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyQuotaTemplateServiceStatus"),
		Version:     tea.String("2020-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyQuotaTemplateServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the status of a quota template. By default, the quota template is not configured. If the management account of a resource directory uses a quota template for the first time, you must enable the quota template.
//
// Description:
//
// ### [](#)Prerequisites
//
// A resource directory is enabled. For more information, see [EnableResourceDirectory](https://help.aliyun.com/document_detail/604185.html).
//
// ### [](#)Usage notes
//
// If the `ServiceStatus` parameter is set to `0` or `-1`, you can call this operation to set the parameter to `1`. Then, you can call the [CreateTemplateQuotaItem](https://help.aliyun.com/document_detail/450615.html) operation to create a quota template.
//
// @param request - ModifyQuotaTemplateServiceStatusRequest
//
// @return ModifyQuotaTemplateServiceStatusResponse
func (client *Client) ModifyQuotaTemplateServiceStatus(request *ModifyQuotaTemplateServiceStatusRequest) (_result *ModifyQuotaTemplateServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyQuotaTemplateServiceStatusResponse{}
	_body, _err := client.ModifyQuotaTemplateServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The ID of the quota template.
//
// @param request - ModifyTemplateQuotaItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTemplateQuotaItemResponse
func (client *Client) ModifyTemplateQuotaItemWithOptions(request *ModifyTemplateQuotaItemRequest, runtime *util.RuntimeOptions) (_result *ModifyTemplateQuotaItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QuotaCategory)) {
		query["QuotaCategory"] = request.QuotaCategory
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesireValue)) {
		body["DesireValue"] = request.DesireValue
	}

	if !tea.BoolValue(util.IsUnset(request.Dimensions)) {
		body["Dimensions"] = request.Dimensions
	}

	if !tea.BoolValue(util.IsUnset(request.EffectiveTime)) {
		body["EffectiveTime"] = request.EffectiveTime
	}

	if !tea.BoolValue(util.IsUnset(request.EnvLanguage)) {
		body["EnvLanguage"] = request.EnvLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireTime)) {
		body["ExpireTime"] = request.ExpireTime
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.NoticeType)) {
		body["NoticeType"] = request.NoticeType
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaActionCode)) {
		body["QuotaActionCode"] = request.QuotaActionCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTemplateQuotaItem"),
		Version:     tea.String("2020-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyTemplateQuotaItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The ID of the quota template.
//
// @param request - ModifyTemplateQuotaItemRequest
//
// @return ModifyTemplateQuotaItemResponse
func (client *Client) ModifyTemplateQuotaItem(request *ModifyTemplateQuotaItemRequest) (_result *ModifyTemplateQuotaItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTemplateQuotaItemResponse{}
	_body, _err := client.ModifyTemplateQuotaItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Reminds the approver of a quota application to review the application. This operation is applicable to quota applications that support the approval reminding feature.
//
// Description:
//
// >  You can call this operation to enable the approval reminder feature for quota applications that support this feature. To check whether this feature is supported, you can view the value of `SupportReminder` in the GetQuotaApplicationApproval operation. If the value of SupportReminder is `true`, this feature is supported.
//
// @param request - RemindQuotaApplicationApprovalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemindQuotaApplicationApprovalResponse
func (client *Client) RemindQuotaApplicationApprovalWithOptions(request *RemindQuotaApplicationApprovalRequest, runtime *util.RuntimeOptions) (_result *RemindQuotaApplicationApprovalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		body["ApplicationId"] = request.ApplicationId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemindQuotaApplicationApproval"),
		Version:     tea.String("2020-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemindQuotaApplicationApprovalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reminds the approver of a quota application to review the application. This operation is applicable to quota applications that support the approval reminding feature.
//
// Description:
//
// >  You can call this operation to enable the approval reminder feature for quota applications that support this feature. To check whether this feature is supported, you can view the value of `SupportReminder` in the GetQuotaApplicationApproval operation. If the value of SupportReminder is `true`, this feature is supported.
//
// @param request - RemindQuotaApplicationApprovalRequest
//
// @return RemindQuotaApplicationApprovalResponse
func (client *Client) RemindQuotaApplicationApproval(request *RemindQuotaApplicationApprovalRequest) (_result *RemindQuotaApplicationApprovalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemindQuotaApplicationApprovalResponse{}
	_body, _err := client.RemindQuotaApplicationApprovalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a quota alert rule.
//
// Description:
//
// In this example, the operation is called to modify the information about a quota alert whose ID is `a2efa7fc-832f-47bb-8054-39e28012****` and whose name is `rules`. The alert threshold is changed from `150` to `160`.
//
// @param request - UpdateQuotaAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQuotaAlarmResponse
func (client *Client) UpdateQuotaAlarmWithOptions(request *UpdateQuotaAlarmRequest, runtime *util.RuntimeOptions) (_result *UpdateQuotaAlarmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmId)) {
		body["AlarmId"] = request.AlarmId
	}

	if !tea.BoolValue(util.IsUnset(request.AlarmName)) {
		body["AlarmName"] = request.AlarmName
	}

	if !tea.BoolValue(util.IsUnset(request.Threshold)) {
		body["Threshold"] = request.Threshold
	}

	if !tea.BoolValue(util.IsUnset(request.ThresholdPercent)) {
		body["ThresholdPercent"] = request.ThresholdPercent
	}

	if !tea.BoolValue(util.IsUnset(request.ThresholdType)) {
		body["ThresholdType"] = request.ThresholdType
	}

	if !tea.BoolValue(util.IsUnset(request.WebHook)) {
		body["WebHook"] = request.WebHook
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateQuotaAlarm"),
		Version:     tea.String("2020-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateQuotaAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a quota alert rule.
//
// Description:
//
// In this example, the operation is called to modify the information about a quota alert whose ID is `a2efa7fc-832f-47bb-8054-39e28012****` and whose name is `rules`. The alert threshold is changed from `150` to `160`.
//
// @param request - UpdateQuotaAlarmRequest
//
// @return UpdateQuotaAlarmResponse
func (client *Client) UpdateQuotaAlarm(request *UpdateQuotaAlarmRequest) (_result *UpdateQuotaAlarmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateQuotaAlarmResponse{}
	_body, _err := client.UpdateQuotaAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
