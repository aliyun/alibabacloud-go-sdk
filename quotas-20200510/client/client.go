// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateQuotaAlarmRequest struct {
	// The name of the quota alert.
	AlarmName *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	// The abbreviation of the cloud service name.
	//
	// >  For more information, see [Alibaba Cloud services that support Quota Center](~~182368~~).
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	QuotaActionCode *string                                   `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	QuotaDimensions []*CreateQuotaAlarmRequestQuotaDimensions `json:"QuotaDimensions,omitempty" xml:"QuotaDimensions,omitempty" type:"Repeated"`
	// The numeric value of the alert threshold. The value must meet the following requirements:
	//
	// *   If the `ThresholdType` parameter is set to `used` and the used quota is greater than or equal to a specified value, you receive an alert. The alert threshold must be greater than the current used quota.
	// *   If the `ThresholdType` parameter is set to `usable` and the available quota is less than or equal to a specified value, you received an alert. The alert threshold must be less than the current available quota.
	//
	// >  You must set one of the Threshold and ThresholdPercent parameters.
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The percentage of the alert threshold. Valid values:
	//
	// *   If the `ThresholdType` parameter is set to `used` and the percentage of the used quota in the total quota is greater than or equal to a specified value, you receive an alert. Value range: (50%, 100%].
	// *   If the `ThresholdType` parameter is set to `usable` and the percentage of the available quota in the total quota is less than or equal to a specified value, you receive an alert. Value range: (0%, 50%].
	//
	// >  You must set one of the Threshold and ThresholdPercent parameters.
	ThresholdPercent *float32 `json:"ThresholdPercent,omitempty" xml:"ThresholdPercent,omitempty"`
	// The type of the quota alert. Valid values:
	//
	// *   used: The alert is created for the used quota.
	// *   usable: The alert is created for the available quota.
	ThresholdType *string `json:"ThresholdType,omitempty" xml:"ThresholdType,omitempty"`
	// The webhook URL. Quota Center sends the alert notification to a specified URL by using an HTTP POST request.
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
	// The dimension keys.
	//
	// The value range of N changes based on the number of dimensions that are supported by the related cloud service.
	//
	// >  If you set the ProductCode parameter to ecs, ecs-spec, actiontrail, or ess, this parameter is required.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The dimension values.
	//
	// The value range of N changes based on the number of dimensions that are supported by the related cloud service.
	//
	// >  If you set the ProductCode parameter to ecs, ecs-spec, actiontrail, or ess, this parameter is required.
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
	// The ID of the alert.
	AlarmId *string `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// The ID of the request.
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateQuotaAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// *   Sync: The application is reviewed in a synchronous manner. Quota Center automatically reviews the application. The result is returned immediately after you submit the application. However, the chance of an approval for an application that is reviewed in Sync mode is lower than the chance of an approval for an application that is reviewed in Async mode. The validity period of the new quota value is 1 hour.
	// *   Async: The application is reviewed in an asynchronous manner. An Alibaba Cloud support engineer reviews the application. The chance of an approval for an application that is reviewed in Async mode is higher than the chance of an approval for an application that is reviewed in Sync mode. The validity period of the new quota value is one month.
	//
	// > This parameter is available only for ECS Quotas by Instance Type.
	AuditMode *string `json:"AuditMode,omitempty" xml:"AuditMode,omitempty"`
	// The requested value of the quota.
	//
	// > Applications are reviewed by the technical support team of each Alibaba Cloud service. To increase the success rate of your application, you must specify a reasonable quota value and detailed reasons when you submit an application to increase the value of the quota.
	DesireValue *float32 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The quota dimensions.
	Dimensions []*CreateQuotaApplicationRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The end time of the validity period of the quota. Specify the value in UTC. This parameter is valid only if you set the QuotaCategory parameter to WhiteListLabel.
	//
	// > If you do not specify an end time, the default end time is 99 years after the quota application is submitted.
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The language of the quota alert notification. Valid values:
	//
	// *   zh (default value): Chinese
	// *   en: English
	EnvLanguage *string `json:"EnvLanguage,omitempty" xml:"EnvLanguage,omitempty"`
	// The start time of the validity period of the quota. Specify the value in UTC. This parameter is valid only if you set the QuotaCategory parameter to WhiteListLabel.
	//
	// > If you do not specify a start time, the default start time is the time when the quota application is submitted.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Specifies whether to send a notification about the application result. Valid values:
	//
	// *   0 (default value): no
	// *   3: yes
	NoticeType *int32 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// > For more information, see [Alibaba Cloud services that support Quota Center](~~182368~~).
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The type of the quota.
	//
	// *   CommonQuota: general quota
	// *   FlowControl: API rate limit
	// *   WhiteListLabel: whitelist quota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The reason for the application.
	//
	// > Applications are reviewed by the technical support team of each Alibaba Cloud service. To increase the success rate of your application, you must specify a reasonable quota value and detailed reasons when you submit an application to increase the value of the quota.
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
	// The dimension keys.
	//
	// The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// > This parameter is required if you set the ProductCode parameter to ecs, ecs-spec, actiontrail, or ess.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The dimension values.
	//
	// The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// > This parameter is required if you set the ProductCode parameter to ecs, ecs-spec, actiontrail, or ess.
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
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The time when the application was submitted.
	ApplyTime *string `json:"ApplyTime,omitempty" xml:"ApplyTime,omitempty"`
	// The quota value that is approved.
	ApproveValue *float32 `json:"ApproveValue,omitempty" xml:"ApproveValue,omitempty"`
	// The result of the application.
	AuditReason *string `json:"AuditReason,omitempty" xml:"AuditReason,omitempty"`
	// The requested value of the quota.
	DesireValue *int32 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The quota dimensions.
	Dimension map[string]interface{} `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The time when the new quota value takes effect.
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The time when the new quota expires.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether Quota Center sends a notification about the application result. Valid values:
	//
	// *   0: Quota Center sends a notification.
	// *   3: Quota Center does not send a notification.
	NoticeType *int64 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the quota.
	QuotaArn *string `json:"QuotaArn,omitempty" xml:"QuotaArn,omitempty"`
	// The description of the quota.
	QuotaDescription *string `json:"QuotaDescription,omitempty" xml:"QuotaDescription,omitempty"`
	// The name of the quota.
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The unit of the new quota value.
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// The reason for the application.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the application. Valid values:
	//
	// *   Disagree: The application is rejected.
	// *   Agree: The application is approved.
	// *   Process: The application is being reviewed.
	// *   Cancel: The application is canceled.
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateQuotaApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreateTemplateQuotaItemRequest struct {
	// The requested value of the quota.
	DesireValue *float32 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The quota dimensions.
	Dimensions []*CreateTemplateQuotaItemRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The start time of the validity period of the quota. Specify the value in UTC. This parameter is valid only if you set the QuotaCategory parameter to WhiteListLabel.
	//
	// > If you do not specify this parameter, the quota takes effect immediately.
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The language of the quota alert notification. Valid values:
	//
	// *   zh (default value): Chinese
	// *   en: English
	EnvLanguage *string `json:"EnvLanguage,omitempty" xml:"EnvLanguage,omitempty"`
	// The end time of the validity period of the quota. Specify the value in UTC. This parameter is valid only if you set the QuotaCategory parameter to WhiteListLabel.
	//
	// > If the value of this parameter is empty, no end time is specified.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Specifies whether to send a notification about the application result. Valid values:
	//
	// *   0 (default value): no
	// *   3: yes
	NoticeType *int64 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// > For more information, see [Alibaba Cloud services that support Quota Center](~~182368~~).
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The type of the quota. Valid values:
	//
	// *   CommonQuota: general quota
	// *   WhiteListLabel: whitelist quota
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
	// The dimension keys.
	//
	// The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// > This parameter is required if you set the ProductCode parameter to ecs, ecs-spec, actiontrail, or ess.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The dimension values.
	//
	// The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// > This parameter is required if you set the ProductCode parameter to ecs, ecs-spec, actiontrail, or ess.
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
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the request.
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTemplateQuotaItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteQuotaAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the request.
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTemplateQuotaItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The details of dimensions.
	Dimensions []*GetProductQuotaRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The abbreviation of the cloud service name.
	//
	// >  For more information, see [Alibaba Cloud services that support Quota Center](~~182368~~).
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
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
	// The dimension keys.
	//
	// >
	// *   The value range of N varies based on the number of dimensions that are supported by the related cloud service.
	// *   If you call the operation to query the details about a quota that belongs to a cloud service that supports dimensions, you must configure this parameter. You must configure the `Dimensions.N.Key` and `Dimensions.N.Value` parameters at the same time. The following cloud services support dimensions: Elastic Compute Service (ECS) whose service code is ecs, Enterprise Distributed Application Service (EDAS) whose service code is edas, ECS Quotas by Instance Type whose service code is ecs-spec, and Auto Scaling (ESS) whose service code is ess.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The dimension values.
	//
	// >
	// *   The value range of N varies based on the number of dimensions that are supported by the related cloud service.
	// *   If you call the operation to query the details about a quota that belongs to a cloud service that supports dimensions, you must configure this parameter. You must configure the `Dimensions.N.Key` and `Dimensions.N.Value` parameters at the same time. The following cloud services support dimensions: Elastic Compute Service (ECS) whose service code is ecs, Enterprise Distributed Application Service (EDAS) whose service code is edas, ECS Quotas by Instance Type whose service code is ecs-spec, and Auto Scaling (ESS) whose service code is ess.
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
	// The details about the quota.
	Quota *GetProductQuotaResponseBodyQuota `json:"Quota,omitempty" xml:"Quota,omitempty" type:"Struct"`
	// The ID of the request.
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
	// *   true
	// *   false
	Adjustable *bool `json:"Adjustable,omitempty" xml:"Adjustable,omitempty"`
	// The range of the quota value, for example, `[802,10000]`.
	ApplicableRange []*float32 `json:"ApplicableRange,omitempty" xml:"ApplicableRange,omitempty" type:"Repeated"`
	// The type of the adjustable value. Valid values:
	//
	// *   continuous
	// *   discontinuous
	ApplicableType *string `json:"ApplicableType,omitempty" xml:"ApplicableType,omitempty"`
	// Indicates whether the system shows the used value of the quota. Valid values:
	//
	// *   true
	// *   false
	Consumable *bool `json:"Consumable,omitempty" xml:"Consumable,omitempty"`
	// The quota dimensions. Format: `{"regionId":"Region"}`.
	Dimensions    map[string]interface{} `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	EffectiveTime *string                `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	ExpireTime    *string                `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The calculation cycle of the quota.
	Period *GetProductQuotaResponseBodyQuotaPeriod `json:"Period,omitempty" xml:"Period,omitempty" type:"Struct"`
	// The abbreviation of the Alibaba Cloud service name.
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the quota.
	QuotaArn      *string `json:"QuotaArn,omitempty" xml:"QuotaArn,omitempty"`
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The description of the quota.
	QuotaDescription *string `json:"QuotaDescription,omitempty" xml:"QuotaDescription,omitempty"`
	// The details about the quota.
	QuotaItems []*GetProductQuotaResponseBodyQuotaQuotaItems `json:"QuotaItems,omitempty" xml:"QuotaItems,omitempty" type:"Repeated"`
	// The name of the quota.
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The type of the quota. Valid values:
	//
	// - privilege
	// - normal (default value)
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// The unit of the quota.
	//
	// >  The unit of each quota is unique. For example, the quota whose ID is `q_cbdch3` represents the maximum number of Container Service for Kubernetes (ACK) clusters. The unit of this quota is clusters. The quota whose ID is `q_security-groups` represents the maximum number of security groups. The unit of this quota is security groups.
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// The range of the quota value that can be requested for the current quota item. When you configure a quota template, you can use the range as a reference.
	//
	// - If the value of the ApplicableType parameter is continuous and the value of the ApplicableRange parameter is [802,1000], the quota value ranges from 802 to 1,000.
	// - If the value of the ApplicableType parameter is discontinuous and the value of the ApplicableRange parameter is [10,20,50,100], the quota value is 10, 20, 50, or 100.
	SupportedRange []*float32 `json:"SupportedRange,omitempty" xml:"SupportedRange,omitempty" type:"Repeated"`
	// The value of the quota.
	TotalQuota *float32 `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	// The used quota.
	TotalUsage *float32 `json:"TotalUsage,omitempty" xml:"TotalUsage,omitempty"`
	// The reason why the quota is not adjustable. Valid values:
	//
	// *   nonactivated: The service is not activated.
	// *   applicationProcess: The application is being processed.
	// *   limitReached: The quota limit is reached.
	// *   supportTicketRequired: The quota can be increased only by submitting a ticket.
	UnadjustableDetail *string `json:"UnadjustableDetail,omitempty" xml:"UnadjustableDetail,omitempty"`
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

type GetProductQuotaResponseBodyQuotaPeriod struct {
	// The unit of the calculation cycle of the quota. Valid values:
	//
	// *   second
	// *   minute
	// *   hour
	// *   day
	// *   week
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The value of the calculation cycle of the quota.
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
	Quota *string `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// The unit of the quota.
	//
	// >  The unit of each quota is unique. For example, the quota whose ID is `q_cbdch3` represents the maximum number of Container Service for Kubernetes (ACK) clusters. The unit of this quota is clusters. The quota whose ID is `q_security-groups` represents the maximum number of security groups. The unit of this quota is Number of security groups.
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// The category of the quota. Valid values:
	//
	// *   BaseQuota: base quota
	// *   ReservedQuota: reserved quota
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The used quota.
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

type GetProductQuotaResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetProductQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The dimension details that are supported by the cloud service.
	DependentDimensions []*GetProductQuotaDimensionRequestDependentDimensions `json:"DependentDimensions,omitempty" xml:"DependentDimensions,omitempty" type:"Repeated"`
	// The dimension key.
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The abbreviation of the cloud service name.
	//
	// >  For more information, see [Alibaba Cloud services that support Quota Center](~~182368~~).
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
	// The dimension keys that are supported by the cloud service.
	//
	// >  The value range of N varies based on the number of dimensions that are supported by the related cloud service.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The dimension values that are supported by the cloud service.
	//
	// >  The value range of N varies based on the number of dimensions that are supported by the related cloud service.
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
	// The details about the quota dimension.
	QuotaDimension *GetProductQuotaDimensionResponseBodyQuotaDimension `json:"QuotaDimension,omitempty" xml:"QuotaDimension,omitempty" type:"Struct"`
	// The ID of the request.
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
	// The quota dimensions that are supported by the cloud service.
	DependentDimensions []*string `json:"DependentDimensions,omitempty" xml:"DependentDimensions,omitempty" type:"Repeated"`
	// The dimension key. Valid values:
	//
	// *   regionId: region ID
	// *   zoneId: zone ID
	// *   chargeType: billing method
	// *   networkType: network type
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The details about the dimension value.
	DimensionValueDetail []*GetProductQuotaDimensionResponseBodyQuotaDimensionDimensionValueDetail `json:"DimensionValueDetail,omitempty" xml:"DimensionValueDetail,omitempty" type:"Repeated"`
	// The dimension values.
	DimensionValues []*string `json:"DimensionValues,omitempty" xml:"DimensionValues,omitempty" type:"Repeated"`
	// The name of the dimension.
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
	// The name of the dimension value.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The dimension value.
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetProductQuotaDimensionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// For more information about how to query the ID of a quota alert, see [ListQuotaAlarms](~~184348~~).
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
	AlarmId *string `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// The name of the quota alert.
	AlarmName *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	// The time when the quota alert was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The notification methods of the quota alert. Valid values:
	//
	// *   sms: SMS messages
	// *   email: emails
	NotifyChannels []*string `json:"NotifyChannels,omitempty" xml:"NotifyChannels,omitempty" type:"Repeated"`
	// The alert contact.
	NotifyTarget *string `json:"NotifyTarget,omitempty" xml:"NotifyTarget,omitempty"`
	// The abbreviation of the cloud service name.
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The quota dimension.
	QuotaDimension map[string]interface{} `json:"QuotaDimension,omitempty" xml:"QuotaDimension,omitempty"`
	// The used quota.
	QuotaUsage *float32 `json:"QuotaUsage,omitempty" xml:"QuotaUsage,omitempty"`
	// The quota value.
	QuotaValue *float32 `json:"QuotaValue,omitempty" xml:"QuotaValue,omitempty"`
	// The numeric value of the alert threshold.
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The percentage of the alert threshold.
	ThresholdPercent *float32 `json:"ThresholdPercent,omitempty" xml:"ThresholdPercent,omitempty"`
	// The type of the quota alert. Valid values:
	//
	// *   used: The alert is created for the used quota.
	// *   usable: The alert is created for the available quota.
	ThresholdType *string `json:"ThresholdType,omitempty" xml:"ThresholdType,omitempty"`
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

type GetQuotaAlarmResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetQuotaAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The time when the application was submitted.
	ApplyTime *string `json:"ApplyTime,omitempty" xml:"ApplyTime,omitempty"`
	// The approved quota value.
	ApproveValue *float32 `json:"ApproveValue,omitempty" xml:"ApproveValue,omitempty"`
	// The result of the application.
	AuditReason *string `json:"AuditReason,omitempty" xml:"AuditReason,omitempty"`
	// The quota for which you apply.
	DesireValue *int32 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The quota dimensions.
	//
	// Format: `{"regionId":"Region"}`.
	Dimension map[string]interface{} `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The time when the quota took effect.
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The time when the quota expired.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The notification method. Valid values:
	//
	// *   0: Quota Center does not send a notification.
	// *   1: Quota Center sends an email notification.
	// *   2: Quota Center sends an SMS notification.
	NoticeType *int64 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The abbreviation of the cloud service name.
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the quota.
	QuotaArn *string `json:"QuotaArn,omitempty" xml:"QuotaArn,omitempty"`
	// The description of the quota.
	QuotaDescription *string `json:"QuotaDescription,omitempty" xml:"QuotaDescription,omitempty"`
	// The name of the quota.
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The unit of the quota.
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// The reason for the application.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The status of the application. Valid values:
	//
	// *   Disagree: The application is rejected.
	// *   Agree: The application is approved.
	// *   Process: The application is pending approval.
	// *   Cancel: The application is closed.
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetQuotaApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type GetQuotaTemplateServiceStatusRequest struct {
	// The ID of the resource directory.
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
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the quota template. Valid values:
	//
	// *   \-1: disabled
	// *   1: enabled
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetQuotaTemplateServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The end of the time range to query.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The keyword that is used to execute the query.
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The maximum number of records to be returned for the query.
	//
	// Valid values: 1 to 200. Default value: 30.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to mark the location where the query is started. An empty value indicates that the query is executed from the start.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The abbreviation of the cloud service name.
	//
	// >  For more information about the Alibaba Cloud services that support Quota Center, see [Alibaba Cloud services that support Quota Center](~~182368~~).
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The start of the time range to query.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListAlarmHistoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmHistoriesRequest) GoString() string {
	return s.String()
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
	// The details of the alert records.
	AlarmHistories []*ListAlarmHistoriesResponseBodyAlarmHistories `json:"AlarmHistories,omitempty" xml:"AlarmHistories,omitempty" type:"Repeated"`
	// The maximum number of records returned for the query.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to mark the location where the query is ended. An empty value indicates that all the data is queried.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records returned for the query.
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
	AlarmName *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	// The time when the quota alert was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The notification methods of the quota alert. Valid values:
	//
	// *   sms: short messages
	// *   email: emails
	NotifyChannels []*string `json:"NotifyChannels,omitempty" xml:"NotifyChannels,omitempty" type:"Repeated"`
	// The alert contact.
	NotifyTarget *string `json:"NotifyTarget,omitempty" xml:"NotifyTarget,omitempty"`
	// The abbreviation of the cloud service name.
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The used quota.
	QuotaUsage *float32 `json:"QuotaUsage,omitempty" xml:"QuotaUsage,omitempty"`
	// The numeric value of the alert threshold.
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The percentage of the alert threshold.
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAlarmHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The abbreviation of the cloud service name.
	//
	// >  For more information about the Alibaba Cloud services that support Quota Center, see [Alibaba Cloud services that support Quota Center](~~182368~~).
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
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
	// The list of quotas on which the specified quota depends.
	Quotas []*ListDependentQuotasResponseBodyQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	// The ID of the request.
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
	// The dimensions of a quota on which the specified quota depends.
	Dimensions []*ListDependentQuotasResponseBodyQuotasDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The abbreviation of the cloud service name.
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The relationship percentage between the specified quota and the quotas on which the specified quota depends.
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
	// The dimension of a quota on which the specified quota depends.
	DependentDimension []*string `json:"DependentDimension,omitempty" xml:"DependentDimension,omitempty" type:"Repeated"`
	// The dimension key.
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDependentQuotasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The maximum number of records that can be returned for the query.
	//
	// Valid values: 1 to 200. Default value: 30.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position from which you want to start the query. If you leave this parameter empty, the query starts from the beginning.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The service code.
	//
	// >  For more information, see [Alibaba Cloud services that support Quota Center](~~182368~~).
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
	// The number of entries returned per page.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position at which the query ends. An empty value indicates that all data is returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records that are returned for the query.
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
	// The keys of the dimension group.
	DimensionKeys []*string `json:"DimensionKeys,omitempty" xml:"DimensionKeys,omitempty" type:"Repeated"`
	// The code of the dimension group.
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
	// The name of the dimension group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The service code.
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProductDimensionGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The maximum number of records that you want to return for the query.
	//
	// Valid values: 1 to 200. Default value: 30.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position where you want to start the query. An empty value indicates that the query starts from the beginning.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The abbreviation of the cloud service name.
	//
	// >  For more information, see [Alibaba Cloud services that support Quota Center](~~182368~~).
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The type of the quota. Valid values:
	//
	// *   FlowControl: API rate limit
	// *   CommonQuota: general quota
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
	// The number of records returned for the query.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position where the query ends. An empty value indicates that all the data is returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The quota dimensions.
	QuotaDimensions []*ListProductQuotaDimensionsResponseBodyQuotaDimensions `json:"QuotaDimensions,omitempty" xml:"QuotaDimensions,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records returned for the query.
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
	// The quota dimensions that are supported by the cloud service.
	DependentDimensions []*string `json:"DependentDimensions,omitempty" xml:"DependentDimensions,omitempty" type:"Repeated"`
	// The dimension key. Valid values:
	//
	// *   RegionId: the region ID
	// *   zoneId: the zone ID
	// *   chargeType: the billing method
	// *   networkType: the network type
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The details about the dimension value.
	DimensionValueDetail []*ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail `json:"DimensionValueDetail,omitempty" xml:"DimensionValueDetail,omitempty" type:"Repeated"`
	// The list of the dimension values.
	DimensionValues []*string `json:"DimensionValues,omitempty" xml:"DimensionValues,omitempty" type:"Repeated"`
	// The name of the quota dimension.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the quota dimension is required when you query quota dimensions. Valid values:
	//
	// *   true
	// *   false
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
	// The name of the dimension value.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The dimension value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail) String() string {
	return tea.Prettify(s)
}

func (s ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail) GoString() string {
	return s.String()
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail) SetName(v string) *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail {
	s.Name = &v
	return s
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail) SetValue(v string) *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail {
	s.Value = &v
	return s
}

type ListProductQuotaDimensionsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProductQuotaDimensionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
	// The keyword that you want to use to search for the quotas.
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The maximum number of records that can be returned for the query.
	//
	// Valid values: 1 to 100. Default value: 30.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position from which you want to start the query. If you leave this parameter empty, the query starts from the beginning.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// > For more information, see [Alibaba Cloud services that support Quota Center](~~182368~~).
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The type of the quota. Valid values:
	//
	// *   CommonQuota: general quota
	// *   FlowControl: API rate limit
	// *   WhiteListLabel: whitelist quota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The field based on which you want to sort the returned records. This parameter is available only for quotas that belong to ECS Quotas by Instance Type. Valid values:
	//
	// *   TIME: The returned records are sorted by the last update time.
	// *   TOTAL: The returned records are sorted by the usage of the total quota.
	// *   RESERVED: The returned records are sorted by the usage of the reserved quota.
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	// The order in which you want to sort the returned records. This parameter is available only for quotas that belong to ECS Quotas by Instance Type. Valid values:
	//
	// *   Ascending: ascending order
	// *   Descending: descending order
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
	// > The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the dimension.
	//
	// > The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
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
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position at which the query ends. An empty value indicates that all data is returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The details of the quotas.
	Quotas []*ListProductQuotasResponseBodyQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records that are returned for the query.
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
	// *   true: The quota is adjustable.
	// *   false: The quota is not adjustable.
	Adjustable *bool `json:"Adjustable,omitempty" xml:"Adjustable,omitempty"`
	// None.
	ApplicableRange []*float32 `json:"ApplicableRange,omitempty" xml:"ApplicableRange,omitempty" type:"Repeated"`
	// The type of the adjustable value. Valid values:
	//
	// *   continuous
	// *   discontinuous
	ApplicableType *string `json:"ApplicableType,omitempty" xml:"ApplicableType,omitempty"`
	// Indicates whether the system shows the used value of the quota. Valid values:
	//
	// *   true: The system shows the used value of the quota.
	// *   false: The system does not show the used value of the quota.
	Consumable *bool `json:"Consumable,omitempty" xml:"Consumable,omitempty"`
	// The quota dimensions. Format: `{"regionId":"Region"}`.
	Dimensions map[string]interface{} `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The start time of the validity period of the quota. The value is displayed in UTC.
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The end time of the validity period of the quota. The value is displayed in UTC.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The calculation cycle of the quota.
	Period *ListProductQuotasResponseBodyQuotasPeriod `json:"Period,omitempty" xml:"Period,omitempty" type:"Struct"`
	// The abbreviation of the Alibaba Cloud service name.
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the quota.
	QuotaArn *string `json:"QuotaArn,omitempty" xml:"QuotaArn,omitempty"`
	// The type of the quota.
	//
	// *   CommonQuota: general quota
	// *   FlowControl: API rate limit
	// *   WhiteListLabel: whitelist quota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The description of the quota.
	QuotaDescription *string `json:"QuotaDescription,omitempty" xml:"QuotaDescription,omitempty"`
	// The details of the quotas.
	QuotaItems []*ListProductQuotasResponseBodyQuotasQuotaItems `json:"QuotaItems,omitempty" xml:"QuotaItems,omitempty" type:"Repeated"`
	// The name of the quota.
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The type of the quota. Valid values:
	//
	// *   privilege
	// *   normal
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// The unit of the new quota value.
	//
	// **
	//
	// **The unit of each quota is unique.** For example, the quota whose ID is `q_cbdch3` represents the maximum number of Container Service for Kubernetes (ACK) clusters. The unit of this quota is clusters. The quota whose ID is `q_security-groups` represents the maximum number of security groups. The unit of this quota is security groups.
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// None.
	SupportedRange []*float32 `json:"SupportedRange,omitempty" xml:"SupportedRange,omitempty" type:"Repeated"`
	// The value of the quota.
	TotalQuota *float32 `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	// The used quota.
	TotalUsage *float32 `json:"TotalUsage,omitempty" xml:"TotalUsage,omitempty"`
	// The reason why the quota is not adjustable. Valid values:
	//
	// *   nonactivated: The service is not activated.
	// *   applicationProcess: The application is being processed.
	// *   limitReached: The quota limit is reached.
	UnadjustableDetail *string `json:"UnadjustableDetail,omitempty" xml:"UnadjustableDetail,omitempty"`
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

type ListProductQuotasResponseBodyQuotasPeriod struct {
	// The unit of the calculation cycle. Valid values:
	//
	// *   second
	// *   minute
	// *   hour
	// *   day
	// *   week
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The value of the calculation cycle.
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
	// The value of the quota.
	Quota *string `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// The unit of the new quota value.
	//
	// **
	//
	// **The unit of each quota is unique.** For example, the quota whose ID is `q_cbdch3` represents the maximum number of Container Service for Kubernetes (ACK) clusters. The unit of this quota is clusters. The quota whose ID is `q_security-groups` represents the maximum number of security groups. The unit of this quota is security groups.
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// The category of the quota. Valid values:
	//
	// *   BaseQuota: base quota
	// *   ReservedQuota: reserved quota
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The used quota.
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

type ListProductQuotasResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProductQuotasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The maximum number of records to be returned for the query.
	//
	// Valid values: 1 to 200. Default value: 30.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to mark the location where the query is started. An empty value indicates that the query is executed from the start.
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
	// The maximum number of records returned for the query.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to mark the location where the query is ended. An empty value indicates that all the data is queried.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information of the cloud service.
	ProductInfo []*ListProductsResponseBodyProductInfo `json:"ProductInfo,omitempty" xml:"ProductInfo,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records returned for the query.
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
	// Indicates whether the cloud service supports general quotas. Valid values:
	//
	// *   support: The cloud service supports general quotas.
	// *   unsupport: The cloud service does not support general quotas.
	CommonQuotaSupport *string `json:"CommonQuotaSupport,omitempty" xml:"CommonQuotaSupport,omitempty"`
	// Indicates whether the cloud service supports dynamic quota adjustment. Valid values:
	//
	// *   true
	// *   false
	Dynamic *bool `json:"Dynamic,omitempty" xml:"Dynamic,omitempty"`
	// Indicates whether the cloud service supports API rate limits. Valid values:
	//
	// *   support: The cloud service supports API rate limits.
	// *   unsupport: The cloud service does not support API rate limits.
	FlowControlSupport *string `json:"FlowControlSupport,omitempty" xml:"FlowControlSupport,omitempty"`
	// The abbreviation of the cloud service name.
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The name of the cloud service.
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The name of the cloud service.
	ProductNameEn *string `json:"ProductNameEn,omitempty" xml:"ProductNameEn,omitempty"`
	// The ID of the service category.
	SecondCategoryId *int64 `json:"SecondCategoryId,omitempty" xml:"SecondCategoryId,omitempty"`
	// The name of the service category.
	SecondCategoryName *string `json:"SecondCategoryName,omitempty" xml:"SecondCategoryName,omitempty"`
	// The name of the service category.
	SecondCategoryNameEn       *string `json:"SecondCategoryNameEn,omitempty" xml:"SecondCategoryNameEn,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProductsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The name of the quota alert.
	AlarmName *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	// The maximum number of records that you want to return for the query.
	//
	// Valid values: 1 to 200. Default value: 30.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position from which you want to start the query.
	//
	// >  An empty value indicates that the query starts from the beginning.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The abbreviation of the cloud service name.
	//
	// >  For more information, see [Alibaba Cloud services that support Quota Center](~~182368~~).
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	QuotaActionCode *string                                  `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
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
	// The dimension keys.
	//
	// >  The value range of N varies based on the number of dimensions that are supported by the related cloud service.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The dimension values.
	//
	// >  The value range of N varies based on the number of dimensions that are supported by the related cloud service.
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
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position at which the query ends.
	//
	// >  If an empty value is returned, all data is queried.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The details about the quota alert.
	QuotaAlarms []*ListQuotaAlarmsResponseBodyQuotaAlarms `json:"QuotaAlarms,omitempty" xml:"QuotaAlarms,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of quota alerts.
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
	// The ID of the quota alert.
	AlarmId *string `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// The name of the quota alert.
	AlarmName *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	// The time when the quota alert was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the alert threshold was reached. Valid values:
	//
	// *   false
	// *   true
	ExceedThreshold *bool `json:"ExceedThreshold,omitempty" xml:"ExceedThreshold,omitempty"`
	// The notification method. Valid values:
	//
	// *   sms: SMS messages
	// *   email: emails
	NotifyChannels []*string `json:"NotifyChannels,omitempty" xml:"NotifyChannels,omitempty" type:"Repeated"`
	// The alert contact. Valid value: accountContact.
	NotifyTarget *string `json:"NotifyTarget,omitempty" xml:"NotifyTarget,omitempty"`
	// The abbreviation of the cloud service name.
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The quota dimensions.
	QuotaDimensions map[string]interface{} `json:"QuotaDimensions,omitempty" xml:"QuotaDimensions,omitempty"`
	// The used quota.
	QuotaUsage *float32 `json:"QuotaUsage,omitempty" xml:"QuotaUsage,omitempty"`
	// The value of the quota.
	QuotaValue *float32 `json:"QuotaValue,omitempty" xml:"QuotaValue,omitempty"`
	// The numeric value of the alert threshold.
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The percentage of the alert threshold.
	ThresholdPercent *float32 `json:"ThresholdPercent,omitempty" xml:"ThresholdPercent,omitempty"`
	// The type of the quota alert. Valid values:
	//
	// *   used: The alert is created for the used quota.
	// *   usable: The alert is created for the available quota.
	ThresholdType *string `json:"ThresholdType,omitempty" xml:"ThresholdType,omitempty"`
	// The webhook URL.
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListQuotaAlarmsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The maximum number of records that can be returned for the query. Valid values: 1 to 100. Default value: 30.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position from which you want to start the query.
	//
	// > An empty value indicates that the query starts from the beginning.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// > For more information, see [Alibaba Cloud services that support Quota Center](~~182368~~).
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// 配额种类。取值：
	// - CommonQuota：通用配额。
	// - WhiteListLabel：权益配额。
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
	// > The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the dimension.
	//
	// > The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
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
	// The maximum number of records that are returned for the query.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position at which the query ends.
	//
	// > An empty value indicates that all data is returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The quota templates that are returned.
	QuotaApplicationTemplates []*ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates `json:"QuotaApplicationTemplates,omitempty" xml:"QuotaApplicationTemplates,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records that are returned for the query.
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
	// *   continuous
	// *   discontinuous
	ApplicableType *string `json:"ApplicableType,omitempty" xml:"ApplicableType,omitempty"`
	// The requested value of the quota.
	DesireValue *float32 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The quota dimensions.
	//
	// Format: {"regionId":"Region"}.
	Dimensions map[string]interface{} `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// 配额生效的UTC时间。
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The language of the quota alert notification. Valid values:
	//
	// *   zh: Chinese
	// *   en: English
	EnvLanguage *string `json:"EnvLanguage,omitempty" xml:"EnvLanguage,omitempty"`
	// 配额失效的UTC时间。
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the quota template.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether Quota Center sends a notification about the application result. Valid values:
	//
	// *   0: no
	// *   3: yes
	NoticeType *int32 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// 配额类型。
	// - CommonQuota：通用配额。
	// - WhiteListLabel：权益配额。
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The description of the quota.
	QuotaDescription *string `json:"QuotaDescription,omitempty" xml:"QuotaDescription,omitempty"`
	// The name of the quota.
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

type ListQuotaApplicationTemplatesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListQuotaApplicationTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Dimensions []*ListQuotaApplicationsRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The keyword that you want to use to search for the application.
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The maximum number of records that you want to return for the query.
	//
	// Valid values: 1 to 200. Default value: 30.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position from which you want to start the query. An empty value indicates that the query starts from the beginning.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The abbreviation of the cloud service name.
	//
	// >  For more information, see [Alibaba Cloud services that support Quota Center](~~182368~~).
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The type of the quota. Valid values:
	//
	// *   FlowControl: API rate limit
	// *   CommonQuota: general quota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The status of the application. Valid values:
	//
	// *   Disagree: The application is rejected.
	// *   Agree: The application is approved.
	// *   Process: The application is pending approval.
	// *   Cancel: The application is closed.
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
	// The dimension keys.
	//
	// >  The value range of N varies based on the number of dimensions that are supported by the related cloud service.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The dimension values.
	//
	// >  The value range of N varies based on the number of dimensions that are supported by the related cloud service.
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
	// The number of records that are returned for the query.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position at which the query ends. An empty value indicates that all data is returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The details about the applications.
	QuotaApplications []*ListQuotaApplicationsResponseBodyQuotaApplications `json:"QuotaApplications,omitempty" xml:"QuotaApplications,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of applications.
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
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The time when the application was submitted.
	ApplyTime *string `json:"ApplyTime,omitempty" xml:"ApplyTime,omitempty"`
	// The quota value that is approved.
	ApproveValue *float32 `json:"ApproveValue,omitempty" xml:"ApproveValue,omitempty"`
	// The result of the application.
	AuditReason *string `json:"AuditReason,omitempty" xml:"AuditReason,omitempty"`
	// The remarks of the application.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The quota value that is approved.
	DesireValue *float32 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The quota dimension of the application.
	Dimension map[string]interface{} `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The time when the application took effect.
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The time when the application expired.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether Quota Center sends a notification about the application result. Valid values:
	//
	// *   0: Quota Center sends a notification.
	// *   3: Quota Center does not send a notification.
	NoticeType *int32 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The calculation cycle of the quota.
	Period *ListQuotaApplicationsResponseBodyQuotaApplicationsPeriod `json:"Period,omitempty" xml:"Period,omitempty" type:"Struct"`
	// The abbreviation of the cloud service name.
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the quota.
	QuotaArn *string `json:"QuotaArn,omitempty" xml:"QuotaArn,omitempty"`
	// The description of the quota.
	QuotaDescription *string `json:"QuotaDescription,omitempty" xml:"QuotaDescription,omitempty"`
	// The name of the quota.
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The unit of the quota.
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// The reason for the application.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The status of the application. Valid values:
	//
	// *   Disagree: The application is rejected.
	// *   Agree: The application is approved.
	// *   Process: The application is pending approval.
	// *   Cancel: The application is closed.
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
	// The unit of the calculation cycle of the quota. Valid values:
	//
	// *   second
	// *   minute
	// *   hour
	// *   day
	// *   week
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The value of the calculation cycle of the quota.
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListQuotaApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ModifyQuotaTemplateServiceStatusRequest struct {
	// The status of the quota template. Valid values:
	//
	// *   \-1: disabled
	// *   1: enabled
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
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the quota template. Valid values:
	//
	// *   \-1: disabled
	// *   1: enabled
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
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyQuotaTemplateServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The requested value of the quota.
	DesireValue *float32 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The quota dimensions.
	Dimensions []*ModifyTemplateQuotaItemRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The start time of the validity period of the quota. Specify the value in UTC. This parameter is valid only if you set the QuotaCategory parameter to WhiteListLabel.
	//
	// > If you do not specify this parameter, the quota takes effect immediately.
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The language of the quota alert notification. Valid values:
	//
	// *   zh (default value): Chinese
	// *   en: English
	EnvLanguage *string `json:"EnvLanguage,omitempty" xml:"EnvLanguage,omitempty"`
	// The end time of the validity period of the quota. Specify the value in UTC. This parameter is valid only if you set the QuotaCategory parameter to WhiteListLabel.
	//
	// > If the value of this parameter is empty, no end time is specified.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the quota template. For more information about how to obtain the ID of a quota template, see [ListQuotaApplicationTemplates](~~450403~~).
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Specifies whether to send a notification about the application result. Valid values:
	//
	// *   0 (default value): no
	// *   3: yes
	NoticeType *int64 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// > For more information, see [Alibaba Cloud services that support Quota Center](~~182368~~).
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The type of the quota. Valid values:
	//
	// *   CommonQuota: general quota
	// *   WhiteListLabel: whitelist quota
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
	// The dimension keys.
	//
	// The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// > This parameter is required if you set the ProductCode parameter to ecs, ecs-spec, actiontrail, or ess.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The dimension values.
	//
	// The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// > This parameter is required if you set the ProductCode parameter to ecs, ecs-spec, actiontrail, or ess.
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
	// The ID of the quota template.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the request.
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyTemplateQuotaItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type UpdateQuotaAlarmRequest struct {
	// The ID of the quota alert.
	AlarmId *string `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// The name of the quota alert.
	AlarmName *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	// The numeric value of the alert threshold. Valid values:
	//
	// *   If the `ThresholdType` parameter is set to `used` and the used quota is greater than or equal to a specified value, you receive alert notifications. The alert threshold must be greater than the current used quota.
	// *   If the `ThresholdType` parameter is set to `usable` and the available quota is less than or equal to a specified value, you receive alert notifications. The alert threshold must be less than the current available quota.
	//
	// >  You must set one of the Threshold and ThresholdPercent parameters.
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The percentage of the alert threshold. Valid values:
	//
	// *   If the `ThresholdType` parameter is set to `used` and the percentage of the used quota in the total quota is greater than or equal to a specified value, you receive alert notifications. Value range: (50%, 100%].
	// *   If the `ThresholdType` parameter is set to `usable` and the percentage of the available quota in the total quota is less than or equal to a specified value, you receive alert notifications. Value range: (0%, 50%].
	//
	// >  You must set one of the Threshold and ThresholdPercent parameters.
	ThresholdPercent *float32 `json:"ThresholdPercent,omitempty" xml:"ThresholdPercent,omitempty"`
	// The type of the quota alert. Valid values:
	//
	// *   used: The alert is created for the used quota.
	// *   usable: The alert is created for the available quota.
	ThresholdType *string `json:"ThresholdType,omitempty" xml:"ThresholdType,omitempty"`
	// The webhook URL. Quota Center sends the alert notification to a specified URL by using an HTTP POST request.
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateQuotaAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

/**
 * In this example, the operation is called to create a quota alert for a quota whose ID is `q_hvnoqv`. This quota represents the maximum number of rules that can be created by a user. The quota belongs to Cloud Config whose service code is `config`.
 *
 * @param request CreateQuotaAlarmRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateQuotaAlarmResponse
 */
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

/**
 * In this example, the operation is called to create a quota alert for a quota whose ID is `q_hvnoqv`. This quota represents the maximum number of rules that can be created by a user. The quota belongs to Cloud Config whose service code is `config`.
 *
 * @param request CreateQuotaAlarmRequest
 * @return CreateQuotaAlarmResponse
 */
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

/**
 * In this example, the operation is called to submit an application to increase the value of a quota whose ID is `q_security-groups` and whose name is Maximum Number of Security Groups. This quota belongs to Elastic Compute Service (ECS). The requested value of the quota is `804`, the application reason is `Scale Out`, and the region of the quota is `cn-hangzhou`.
 *
 * @param request CreateQuotaApplicationRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateQuotaApplicationResponse
 */
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

/**
 * In this example, the operation is called to submit an application to increase the value of a quota whose ID is `q_security-groups` and whose name is Maximum Number of Security Groups. This quota belongs to Elastic Compute Service (ECS). The requested value of the quota is `804`, the application reason is `Scale Out`, and the region of the quota is `cn-hangzhou`.
 *
 * @param request CreateQuotaApplicationRequest
 * @return CreateQuotaApplicationResponse
 */
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

/**
 * In this example, the operation is called to delete a quota alert whose ID is `6b512ab7-da3a-4142-b529-2b2a9294****`.
 *
 * @param request DeleteQuotaAlarmRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteQuotaAlarmResponse
 */
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

/**
 * In this example, the operation is called to delete a quota alert whose ID is `6b512ab7-da3a-4142-b529-2b2a9294****`.
 *
 * @param request DeleteQuotaAlarmRequest
 * @return DeleteQuotaAlarmResponse
 */
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

/**
 * In this example, the operation is called to query the details about a quota whose ID is `q_security-groups` and whose name is Maximum Number of Security Groups. This quota belongs to Elastic Compute Service (ECS). The query result shows the details about the quota. The details include the name, ID, description, quota value, used quota, unit, and dimension of the quota. In this example, the quota name is `Maximum Number of Security Groups`. The quota ID is `q_security-groups`. The description is `The maximum number of security groups that can be created for the current account`. The quota value is `801`. The used quota is `26`. The quota unit is `security groups`. The quota dimension is `{"regionId":"cn-hangzhou"}`.
 *
 * @param request GetProductQuotaRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetProductQuotaResponse
 */
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

/**
 * In this example, the operation is called to query the details about a quota whose ID is `q_security-groups` and whose name is Maximum Number of Security Groups. This quota belongs to Elastic Compute Service (ECS). The query result shows the details about the quota. The details include the name, ID, description, quota value, used quota, unit, and dimension of the quota. In this example, the quota name is `Maximum Number of Security Groups`. The quota ID is `q_security-groups`. The description is `The maximum number of security groups that can be created for the current account`. The quota value is `801`. The used quota is `26`. The quota unit is `security groups`. The quota dimension is `{"regionId":"cn-hangzhou"}`.
 *
 * @param request GetProductQuotaRequest
 * @return GetProductQuotaResponse
 */
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

/**
 * In this example, the operation is called to query the details about a quota dimension whose key is `regionId`. The quota dimension belongs to ECS Quotas by Instance Type whose service code is ecs-spec. The following query result is returned:
 * *   The values of the quota dimension include `cn-shenzhen`, `cn-beijing`, and `cn-hangzhou`.
 * *   The name of the quota dimension is `region`.
 *
 * @param request GetProductQuotaDimensionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetProductQuotaDimensionResponse
 */
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

/**
 * In this example, the operation is called to query the details about a quota dimension whose key is `regionId`. The quota dimension belongs to ECS Quotas by Instance Type whose service code is ecs-spec. The following query result is returned:
 * *   The values of the quota dimension include `cn-shenzhen`, `cn-beijing`, and `cn-hangzhou`.
 * *   The name of the quota dimension is `region`.
 *
 * @param request GetProductQuotaDimensionRequest
 * @return GetProductQuotaDimensionResponse
 */
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

/**
 * In this example, the operation is called to query the details of a quota alert whose ID is `78d7e436-4b25-4897-84b5-d7b656bb****`. The details of the alert are returned. The query result includes the alert ID, alert name, alert contact, and the time when the quota alert was created.
 *
 * @param request GetQuotaAlarmRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetQuotaAlarmResponse
 */
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

/**
 * In this example, the operation is called to query the details of a quota alert whose ID is `78d7e436-4b25-4897-84b5-d7b656bb****`. The details of the alert are returned. The query result includes the alert ID, alert name, alert contact, and the time when the quota alert was created.
 *
 * @param request GetQuotaAlarmRequest
 * @return GetQuotaAlarmResponse
 */
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

/**
 * In this example, the operation is called to query the details about an application whose ID is `d314d6ae-867d-484c-9009-3d421a80****`. The query result shows the details about the application. The details include the application ID, application time, expected quota value, and application result.
 *
 * @param request GetQuotaApplicationRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetQuotaApplicationResponse
 */
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

/**
 * In this example, the operation is called to query the details about an application whose ID is `d314d6ae-867d-484c-9009-3d421a80****`. The query result shows the details about the application. The details include the application ID, application time, expected quota value, and application result.
 *
 * @param request GetQuotaApplicationRequest
 * @return GetQuotaApplicationResponse
 */
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

func (client *Client) ListAlarmHistoriesWithOptions(request *ListAlarmHistoriesRequest, runtime *util.RuntimeOptions) (_result *ListAlarmHistoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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

/**
 * In this example, the operation is called to query the list of quotas. A quota whose ID is `q_i5uzm3` depends on these quotas. The name of the quota is Maximum Number of Nodes. This quota belongs to Container Service for Kubernetes (ACK). The query result indicates that the specified quota depends on the following three quotas:
 * *   An Elastic Compute Service (ECS) quota whose ID is `q_elastic-network-interfaces`. This quota is the maximum number of ENIs (Secondary ENIs) that can be owned by the current account. The regions of the quota dimension include `cn-shenzhen`, `cn-beijing`, `cn-hangzhou`.
 * *   A Server Load Balancer (SLB) quota whose ID is `q_fh20b0`. This quota is the number of servers that can be attached to the backend of an SLB instance.
 * *   An SLB quota whose ID is `q_3mmbsp`. This quota is the number of SLB instances that can be owned by a user.
 *
 * @param request ListDependentQuotasRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListDependentQuotasResponse
 */
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

/**
 * In this example, the operation is called to query the list of quotas. A quota whose ID is `q_i5uzm3` depends on these quotas. The name of the quota is Maximum Number of Nodes. This quota belongs to Container Service for Kubernetes (ACK). The query result indicates that the specified quota depends on the following three quotas:
 * *   An Elastic Compute Service (ECS) quota whose ID is `q_elastic-network-interfaces`. This quota is the maximum number of ENIs (Secondary ENIs) that can be owned by the current account. The regions of the quota dimension include `cn-shenzhen`, `cn-beijing`, `cn-hangzhou`.
 * *   A Server Load Balancer (SLB) quota whose ID is `q_fh20b0`. This quota is the number of servers that can be attached to the backend of an SLB instance.
 * *   An SLB quota whose ID is `q_3mmbsp`. This quota is the number of SLB instances that can be owned by a user.
 *
 * @param request ListDependentQuotasRequest
 * @return ListDependentQuotasResponse
 */
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

/**
 * This topic provides an example on how to call the ListProductDimensionGroups operation to query the dimension groups of Object Storage Service (OSS). In this example, a dimension group is returned. The group name is `OSS_Group`, the group code is `oss_wf1ngqmd7q`, and the group key is `chargeType`.
 *
 * @param request ListProductDimensionGroupsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListProductDimensionGroupsResponse
 */
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

/**
 * This topic provides an example on how to call the ListProductDimensionGroups operation to query the dimension groups of Object Storage Service (OSS). In this example, a dimension group is returned. The group name is `OSS_Group`, the group code is `oss_wf1ngqmd7q`, and the group key is `chargeType`.
 *
 * @param request ListProductDimensionGroupsRequest
 * @return ListProductDimensionGroupsResponse
 */
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

/**
 * In this example, the operation is called to query the quota dimensions that are supported by Elastic Compute Service (ECS). The query result shows all the quota dimensions that are supported by ECS.
 *
 * @param request ListProductQuotaDimensionsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListProductQuotaDimensionsResponse
 */
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

/**
 * In this example, the operation is called to query the quota dimensions that are supported by Elastic Compute Service (ECS). The query result shows all the quota dimensions that are supported by ECS.
 *
 * @param request ListProductQuotaDimensionsRequest
 * @return ListProductQuotaDimensionsResponse
 */
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

/**
 * In this example, the operation is called to query the quotas whose instance type is `ecs.g5.2xlarge`. The quotas belong to ECS Quotas by Instance Type. The query result includes the name, ID, unit, dimensions, and cycle of each quota.
 *
 * @param request ListProductQuotasRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListProductQuotasResponse
 */
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

/**
 * In this example, the operation is called to query the quotas whose instance type is `ecs.g5.2xlarge`. The quotas belong to ECS Quotas by Instance Type. The query result includes the name, ID, unit, dimensions, and cycle of each quota.
 *
 * @param request ListProductQuotasRequest
 * @return ListProductQuotasResponse
 */
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

/**
 * In this example, the operation is called to query the details about an application that is submitted to increase a quota whose ID is `q_i5uzm3` and whose name is Maximum Number of Nodes. This quota belongs to Container Service for Kubernetes (ACK). The query result shows the details about the application. The details include the application ID, application time, requested quota, and application result. In this example, the application ID is `b926571d-cc09-4711-b547-58a615f0****`. The application time is `2021-01-15T09:13:53Z`. The expected quota value is `101`. The application result is `Agree`.
 *
 * @param request ListQuotaApplicationsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListQuotaApplicationsResponse
 */
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

/**
 * In this example, the operation is called to query the details about an application that is submitted to increase a quota whose ID is `q_i5uzm3` and whose name is Maximum Number of Nodes. This quota belongs to Container Service for Kubernetes (ACK). The query result shows the details about the application. The details include the application ID, application time, requested quota, and application result. In this example, the application ID is `b926571d-cc09-4711-b547-58a615f0****`. The application time is `2021-01-15T09:13:53Z`. The expected quota value is `101`. The application result is `Agree`.
 *
 * @param request ListQuotaApplicationsRequest
 * @return ListQuotaApplicationsResponse
 */
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

/**
 * By default, the status of a quota template is enabled.
 *
 * @param request ModifyQuotaTemplateServiceStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyQuotaTemplateServiceStatusResponse
 */
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

/**
 * By default, the status of a quota template is enabled.
 *
 * @param request ModifyQuotaTemplateServiceStatusRequest
 * @return ModifyQuotaTemplateServiceStatusResponse
 */
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

/**
 * In this example, the operation is called to modify the information about a quota alert whose ID is `a2efa7fc-832f-47bb-8054-39e28012****` and name is `rules`. The alert threshold is changed from `150` to `160`.
 *
 * @param request UpdateQuotaAlarmRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateQuotaAlarmResponse
 */
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

/**
 * In this example, the operation is called to modify the information about a quota alert whose ID is `a2efa7fc-832f-47bb-8054-39e28012****` and name is `rules`. The alert threshold is changed from `150` to `160`.
 *
 * @param request UpdateQuotaAlarmRequest
 * @return UpdateQuotaAlarmResponse
 */
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
