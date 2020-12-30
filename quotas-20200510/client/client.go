// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateQuotaAlarmRequest struct {
	ProductCode      *string                                   `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	QuotaActionCode  *string                                   `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	AlarmName        *string                                   `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	Threshold        *float32                                  `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	ThresholdPercent *float32                                  `json:"ThresholdPercent,omitempty" xml:"ThresholdPercent,omitempty"`
	WebHook          *string                                   `json:"WebHook,omitempty" xml:"WebHook,omitempty"`
	QuotaDimensions  []*CreateQuotaAlarmRequestQuotaDimensions `json:"QuotaDimensions,omitempty" xml:"QuotaDimensions,omitempty" type:"Repeated"`
}

func (s CreateQuotaAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaAlarmRequest) GoString() string {
	return s.String()
}

func (s *CreateQuotaAlarmRequest) SetProductCode(v string) *CreateQuotaAlarmRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateQuotaAlarmRequest) SetQuotaActionCode(v string) *CreateQuotaAlarmRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *CreateQuotaAlarmRequest) SetAlarmName(v string) *CreateQuotaAlarmRequest {
	s.AlarmName = &v
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

func (s *CreateQuotaAlarmRequest) SetWebHook(v string) *CreateQuotaAlarmRequest {
	s.WebHook = &v
	return s
}

func (s *CreateQuotaAlarmRequest) SetQuotaDimensions(v []*CreateQuotaAlarmRequestQuotaDimensions) *CreateQuotaAlarmRequest {
	s.QuotaDimensions = v
	return s
}

type CreateQuotaAlarmRequestQuotaDimensions struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	AlarmId   *string `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateQuotaAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateQuotaAlarmResponse) SetBody(v *CreateQuotaAlarmResponseBody) *CreateQuotaAlarmResponse {
	s.Body = v
	return s
}

type CreateQuotaApplicationRequest struct {
	ProductCode     *string                                    `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	QuotaActionCode *string                                    `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	DesireValue     *float32                                   `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	Reason          *string                                    `json:"Reason,omitempty" xml:"Reason,omitempty"`
	NoticeType      *int32                                     `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	Dimensions      []*CreateQuotaApplicationRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
}

func (s CreateQuotaApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateQuotaApplicationRequest) SetProductCode(v string) *CreateQuotaApplicationRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateQuotaApplicationRequest) SetQuotaActionCode(v string) *CreateQuotaApplicationRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *CreateQuotaApplicationRequest) SetDesireValue(v float32) *CreateQuotaApplicationRequest {
	s.DesireValue = &v
	return s
}

func (s *CreateQuotaApplicationRequest) SetReason(v string) *CreateQuotaApplicationRequest {
	s.Reason = &v
	return s
}

func (s *CreateQuotaApplicationRequest) SetNoticeType(v int32) *CreateQuotaApplicationRequest {
	s.NoticeType = &v
	return s
}

func (s *CreateQuotaApplicationRequest) SetDimensions(v []*CreateQuotaApplicationRequestDimensions) *CreateQuotaApplicationRequest {
	s.Dimensions = v
	return s
}

type CreateQuotaApplicationRequestDimensions struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s CreateQuotaApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQuotaApplicationResponseBody) SetRequestId(v string) *CreateQuotaApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQuotaApplicationResponseBody) SetApplicationId(v string) *CreateQuotaApplicationResponseBody {
	s.ApplicationId = &v
	return s
}

type CreateQuotaApplicationResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateQuotaApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateQuotaApplicationResponse) SetBody(v *CreateQuotaApplicationResponseBody) *CreateQuotaApplicationResponse {
	s.Body = v
	return s
}

type DeleteQuotaAlarmRequest struct {
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteQuotaAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteQuotaAlarmResponse) SetBody(v *DeleteQuotaAlarmResponseBody) *DeleteQuotaAlarmResponse {
	s.Body = v
	return s
}

type GetProductQuotaRequest struct {
	ProductCode     *string                             `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	QuotaActionCode *string                             `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	Dimensions      []*GetProductQuotaRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
}

func (s GetProductQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProductQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetProductQuotaRequest) SetProductCode(v string) *GetProductQuotaRequest {
	s.ProductCode = &v
	return s
}

func (s *GetProductQuotaRequest) SetQuotaActionCode(v string) *GetProductQuotaRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *GetProductQuotaRequest) SetDimensions(v []*GetProductQuotaRequestDimensions) *GetProductQuotaRequest {
	s.Dimensions = v
	return s
}

type GetProductQuotaRequestDimensions struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Quota     *GetProductQuotaResponseBodyQuota `json:"Quota,omitempty" xml:"Quota,omitempty" type:"Struct"`
}

func (s GetProductQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProductQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductQuotaResponseBody) SetRequestId(v string) *GetProductQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProductQuotaResponseBody) SetQuota(v *GetProductQuotaResponseBodyQuota) *GetProductQuotaResponseBody {
	s.Quota = v
	return s
}

type GetProductQuotaResponseBodyQuota struct {
	QuotaUnit          *string                                       `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	QuotaActionCode    *string                                       `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	TotalUsage         *float32                                      `json:"TotalUsage,omitempty" xml:"TotalUsage,omitempty"`
	ApplicableRange    []*float32                                    `json:"ApplicableRange,omitempty" xml:"ApplicableRange,omitempty" type:"Repeated"`
	QuotaType          *string                                       `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	QuotaDescription   *string                                       `json:"QuotaDescription,omitempty" xml:"QuotaDescription,omitempty"`
	Period             *GetProductQuotaResponseBodyQuotaPeriod       `json:"Period,omitempty" xml:"Period,omitempty" type:"Struct"`
	QuotaArn           *string                                       `json:"QuotaArn,omitempty" xml:"QuotaArn,omitempty"`
	ApplicableType     *string                                       `json:"ApplicableType,omitempty" xml:"ApplicableType,omitempty"`
	QuotaItems         []*GetProductQuotaResponseBodyQuotaQuotaItems `json:"QuotaItems,omitempty" xml:"QuotaItems,omitempty" type:"Repeated"`
	Dimensions         map[string]interface{}                        `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	Adjustable         *bool                                         `json:"Adjustable,omitempty" xml:"Adjustable,omitempty"`
	QuotaName          *string                                       `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	UnadjustableDetail *string                                       `json:"UnadjustableDetail,omitempty" xml:"UnadjustableDetail,omitempty"`
	Consumable         *bool                                         `json:"Consumable,omitempty" xml:"Consumable,omitempty"`
	TotalQuota         *float32                                      `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	ProductCode        *string                                       `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s GetProductQuotaResponseBodyQuota) String() string {
	return tea.Prettify(s)
}

func (s GetProductQuotaResponseBodyQuota) GoString() string {
	return s.String()
}

func (s *GetProductQuotaResponseBodyQuota) SetQuotaUnit(v string) *GetProductQuotaResponseBodyQuota {
	s.QuotaUnit = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetQuotaActionCode(v string) *GetProductQuotaResponseBodyQuota {
	s.QuotaActionCode = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetTotalUsage(v float32) *GetProductQuotaResponseBodyQuota {
	s.TotalUsage = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetApplicableRange(v []*float32) *GetProductQuotaResponseBodyQuota {
	s.ApplicableRange = v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetQuotaType(v string) *GetProductQuotaResponseBodyQuota {
	s.QuotaType = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetQuotaDescription(v string) *GetProductQuotaResponseBodyQuota {
	s.QuotaDescription = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetPeriod(v *GetProductQuotaResponseBodyQuotaPeriod) *GetProductQuotaResponseBodyQuota {
	s.Period = v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetQuotaArn(v string) *GetProductQuotaResponseBodyQuota {
	s.QuotaArn = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetApplicableType(v string) *GetProductQuotaResponseBodyQuota {
	s.ApplicableType = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetQuotaItems(v []*GetProductQuotaResponseBodyQuotaQuotaItems) *GetProductQuotaResponseBodyQuota {
	s.QuotaItems = v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetDimensions(v map[string]interface{}) *GetProductQuotaResponseBodyQuota {
	s.Dimensions = v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetAdjustable(v bool) *GetProductQuotaResponseBodyQuota {
	s.Adjustable = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetQuotaName(v string) *GetProductQuotaResponseBodyQuota {
	s.QuotaName = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetUnadjustableDetail(v string) *GetProductQuotaResponseBodyQuota {
	s.UnadjustableDetail = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetConsumable(v bool) *GetProductQuotaResponseBodyQuota {
	s.Consumable = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetTotalQuota(v float32) *GetProductQuotaResponseBodyQuota {
	s.TotalQuota = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetProductCode(v string) *GetProductQuotaResponseBodyQuota {
	s.ProductCode = &v
	return s
}

type GetProductQuotaResponseBodyQuotaPeriod struct {
	PeriodValue *int32  `json:"PeriodValue,omitempty" xml:"PeriodValue,omitempty"`
	PeriodUnit  *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
}

func (s GetProductQuotaResponseBodyQuotaPeriod) String() string {
	return tea.Prettify(s)
}

func (s GetProductQuotaResponseBodyQuotaPeriod) GoString() string {
	return s.String()
}

func (s *GetProductQuotaResponseBodyQuotaPeriod) SetPeriodValue(v int32) *GetProductQuotaResponseBodyQuotaPeriod {
	s.PeriodValue = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuotaPeriod) SetPeriodUnit(v string) *GetProductQuotaResponseBodyQuotaPeriod {
	s.PeriodUnit = &v
	return s
}

type GetProductQuotaResponseBodyQuotaQuotaItems struct {
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Quota     *string `json:"Quota,omitempty" xml:"Quota,omitempty"`
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	Usage     *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s GetProductQuotaResponseBodyQuotaQuotaItems) String() string {
	return tea.Prettify(s)
}

func (s GetProductQuotaResponseBodyQuotaQuotaItems) GoString() string {
	return s.String()
}

func (s *GetProductQuotaResponseBodyQuotaQuotaItems) SetType(v string) *GetProductQuotaResponseBodyQuotaQuotaItems {
	s.Type = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuotaQuotaItems) SetQuota(v string) *GetProductQuotaResponseBodyQuotaQuotaItems {
	s.Quota = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuotaQuotaItems) SetQuotaUnit(v string) *GetProductQuotaResponseBodyQuotaQuotaItems {
	s.QuotaUnit = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuotaQuotaItems) SetUsage(v string) *GetProductQuotaResponseBodyQuotaQuotaItems {
	s.Usage = &v
	return s
}

type GetProductQuotaResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProductQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetProductQuotaResponse) SetBody(v *GetProductQuotaResponseBody) *GetProductQuotaResponse {
	s.Body = v
	return s
}

type GetProductQuotaDimensionRequest struct {
	ProductCode         *string                                               `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	DimensionKey        *string                                               `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	DependentDimensions []*GetProductQuotaDimensionRequestDependentDimensions `json:"DependentDimensions,omitempty" xml:"DependentDimensions,omitempty" type:"Repeated"`
}

func (s GetProductQuotaDimensionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProductQuotaDimensionRequest) GoString() string {
	return s.String()
}

func (s *GetProductQuotaDimensionRequest) SetProductCode(v string) *GetProductQuotaDimensionRequest {
	s.ProductCode = &v
	return s
}

func (s *GetProductQuotaDimensionRequest) SetDimensionKey(v string) *GetProductQuotaDimensionRequest {
	s.DimensionKey = &v
	return s
}

func (s *GetProductQuotaDimensionRequest) SetDependentDimensions(v []*GetProductQuotaDimensionRequestDependentDimensions) *GetProductQuotaDimensionRequest {
	s.DependentDimensions = v
	return s
}

type GetProductQuotaDimensionRequestDependentDimensions struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	QuotaDimension *GetProductQuotaDimensionResponseBodyQuotaDimension `json:"QuotaDimension,omitempty" xml:"QuotaDimension,omitempty" type:"Struct"`
}

func (s GetProductQuotaDimensionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProductQuotaDimensionResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductQuotaDimensionResponseBody) SetRequestId(v string) *GetProductQuotaDimensionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProductQuotaDimensionResponseBody) SetQuotaDimension(v *GetProductQuotaDimensionResponseBodyQuotaDimension) *GetProductQuotaDimensionResponseBody {
	s.QuotaDimension = v
	return s
}

type GetProductQuotaDimensionResponseBodyQuotaDimension struct {
	DimensionKey        *string   `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	DependentDimensions []*string `json:"DependentDimensions,omitempty" xml:"DependentDimensions,omitempty" type:"Repeated"`
	DimensionValues     []*string `json:"DimensionValues,omitempty" xml:"DimensionValues,omitempty" type:"Repeated"`
	Name                *string   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetProductQuotaDimensionResponseBodyQuotaDimension) String() string {
	return tea.Prettify(s)
}

func (s GetProductQuotaDimensionResponseBodyQuotaDimension) GoString() string {
	return s.String()
}

func (s *GetProductQuotaDimensionResponseBodyQuotaDimension) SetDimensionKey(v string) *GetProductQuotaDimensionResponseBodyQuotaDimension {
	s.DimensionKey = &v
	return s
}

func (s *GetProductQuotaDimensionResponseBodyQuotaDimension) SetDependentDimensions(v []*string) *GetProductQuotaDimensionResponseBodyQuotaDimension {
	s.DependentDimensions = v
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

type GetProductQuotaDimensionResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProductQuotaDimensionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetProductQuotaDimensionResponse) SetBody(v *GetProductQuotaDimensionResponseBody) *GetProductQuotaDimensionResponse {
	s.Body = v
	return s
}

type GetQuotaAlarmRequest struct {
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
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	QuotaAlarm *GetQuotaAlarmResponseBodyQuotaAlarm `json:"QuotaAlarm,omitempty" xml:"QuotaAlarm,omitempty" type:"Struct"`
}

func (s GetQuotaAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaAlarmResponseBody) SetRequestId(v string) *GetQuotaAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQuotaAlarmResponseBody) SetQuotaAlarm(v *GetQuotaAlarmResponseBodyQuotaAlarm) *GetQuotaAlarmResponseBody {
	s.QuotaAlarm = v
	return s
}

type GetQuotaAlarmResponseBodyQuotaAlarm struct {
	ThresholdPercent *float32               `json:"ThresholdPercent,omitempty" xml:"ThresholdPercent,omitempty"`
	QuotaDimension   map[string]interface{} `json:"QuotaDimension,omitempty" xml:"QuotaDimension,omitempty"`
	CreateTime       *string                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	QuotaActionCode  *string                `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	AlarmName        *string                `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	NotifyTarget     *string                `json:"NotifyTarget,omitempty" xml:"NotifyTarget,omitempty"`
	NotifyChannels   []*string              `json:"NotifyChannels,omitempty" xml:"NotifyChannels,omitempty" type:"Repeated"`
	QuotaUsage       *float32               `json:"QuotaUsage,omitempty" xml:"QuotaUsage,omitempty"`
	QuotaValue       *float32               `json:"QuotaValue,omitempty" xml:"QuotaValue,omitempty"`
	AlarmId          *string                `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	Threshold        *float32               `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	ProductCode      *string                `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s GetQuotaAlarmResponseBodyQuotaAlarm) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaAlarmResponseBodyQuotaAlarm) GoString() string {
	return s.String()
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetThresholdPercent(v float32) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.ThresholdPercent = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetQuotaDimension(v map[string]interface{}) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.QuotaDimension = v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetCreateTime(v string) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.CreateTime = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetQuotaActionCode(v string) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.QuotaActionCode = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetAlarmName(v string) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.AlarmName = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetNotifyTarget(v string) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.NotifyTarget = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetNotifyChannels(v []*string) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.NotifyChannels = v
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

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetAlarmId(v string) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.AlarmId = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetThreshold(v float32) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.Threshold = &v
	return s
}

func (s *GetQuotaAlarmResponseBodyQuotaAlarm) SetProductCode(v string) *GetQuotaAlarmResponseBodyQuotaAlarm {
	s.ProductCode = &v
	return s
}

type GetQuotaAlarmResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetQuotaAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetQuotaAlarmResponse) SetBody(v *GetQuotaAlarmResponseBody) *GetQuotaAlarmResponse {
	s.Body = v
	return s
}

type GetQuotaApplicationRequest struct {
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
	RequestId        *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	QuotaApplication *GetQuotaApplicationResponseBodyQuotaApplication `json:"QuotaApplication,omitempty" xml:"QuotaApplication,omitempty" type:"Struct"`
}

func (s GetQuotaApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaApplicationResponseBody) SetRequestId(v string) *GetQuotaApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQuotaApplicationResponseBody) SetQuotaApplication(v *GetQuotaApplicationResponseBodyQuotaApplication) *GetQuotaApplicationResponseBody {
	s.QuotaApplication = v
	return s
}

type GetQuotaApplicationResponseBodyQuotaApplication struct {
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	DesireValue      *int32  `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	QuotaActionCode  *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	QuotaName        *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	ApplicationId    *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	Reason           *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	AuditReason      *string `json:"AuditReason,omitempty" xml:"AuditReason,omitempty"`
	QuotaDescription *string `json:"QuotaDescription,omitempty" xml:"QuotaDescription,omitempty"`
	ProductCode      *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	QuotaArn         *string `json:"QuotaArn,omitempty" xml:"QuotaArn,omitempty"`
	ApplyTime        *string `json:"ApplyTime,omitempty" xml:"ApplyTime,omitempty"`
}

func (s GetQuotaApplicationResponseBodyQuotaApplication) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaApplicationResponseBodyQuotaApplication) GoString() string {
	return s.String()
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetStatus(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.Status = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetDesireValue(v int32) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.DesireValue = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetQuotaActionCode(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.QuotaActionCode = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetQuotaName(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.QuotaName = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetApplicationId(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.ApplicationId = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetReason(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.Reason = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetAuditReason(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.AuditReason = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetQuotaDescription(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.QuotaDescription = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetProductCode(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.ProductCode = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetQuotaArn(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.QuotaArn = &v
	return s
}

func (s *GetQuotaApplicationResponseBodyQuotaApplication) SetApplyTime(v string) *GetQuotaApplicationResponseBodyQuotaApplication {
	s.ApplyTime = &v
	return s
}

type GetQuotaApplicationResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetQuotaApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetQuotaApplicationResponse) SetBody(v *GetQuotaApplicationResponseBody) *GetQuotaApplicationResponse {
	s.Body = v
	return s
}

type ListAlarmHistoriesRequest struct {
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Keyword     *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s ListAlarmHistoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmHistoriesRequest) GoString() string {
	return s.String()
}

func (s *ListAlarmHistoriesRequest) SetNextToken(v string) *ListAlarmHistoriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAlarmHistoriesRequest) SetMaxResults(v int32) *ListAlarmHistoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAlarmHistoriesRequest) SetKeyword(v string) *ListAlarmHistoriesRequest {
	s.Keyword = &v
	return s
}

func (s *ListAlarmHistoriesRequest) SetStartTime(v int64) *ListAlarmHistoriesRequest {
	s.StartTime = &v
	return s
}

func (s *ListAlarmHistoriesRequest) SetEndTime(v int64) *ListAlarmHistoriesRequest {
	s.EndTime = &v
	return s
}

func (s *ListAlarmHistoriesRequest) SetProductCode(v string) *ListAlarmHistoriesRequest {
	s.ProductCode = &v
	return s
}

type ListAlarmHistoriesResponseBody struct {
	TotalCount     *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	NextToken      *string                                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxResults     *int32                                          `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	AlarmHistories []*ListAlarmHistoriesResponseBodyAlarmHistories `json:"AlarmHistories,omitempty" xml:"AlarmHistories,omitempty" type:"Repeated"`
}

func (s ListAlarmHistoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmHistoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlarmHistoriesResponseBody) SetTotalCount(v int32) *ListAlarmHistoriesResponseBody {
	s.TotalCount = &v
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

func (s *ListAlarmHistoriesResponseBody) SetMaxResults(v int32) *ListAlarmHistoriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAlarmHistoriesResponseBody) SetAlarmHistories(v []*ListAlarmHistoriesResponseBodyAlarmHistories) *ListAlarmHistoriesResponseBody {
	s.AlarmHistories = v
	return s
}

type ListAlarmHistoriesResponseBodyAlarmHistories struct {
	QuotaUsage       *float32  `json:"QuotaUsage,omitempty" xml:"QuotaUsage,omitempty"`
	ThresholdPercent *float32  `json:"ThresholdPercent,omitempty" xml:"ThresholdPercent,omitempty"`
	CreateTime       *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	QuotaActionCode  *string   `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	AlarmName        *string   `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	NotifyTarget     *string   `json:"NotifyTarget,omitempty" xml:"NotifyTarget,omitempty"`
	NotifyChannels   []*string `json:"NotifyChannels,omitempty" xml:"NotifyChannels,omitempty" type:"Repeated"`
	Threshold        *float32  `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	ProductCode      *string   `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s ListAlarmHistoriesResponseBodyAlarmHistories) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmHistoriesResponseBodyAlarmHistories) GoString() string {
	return s.String()
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) SetQuotaUsage(v float32) *ListAlarmHistoriesResponseBodyAlarmHistories {
	s.QuotaUsage = &v
	return s
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) SetThresholdPercent(v float32) *ListAlarmHistoriesResponseBodyAlarmHistories {
	s.ThresholdPercent = &v
	return s
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) SetCreateTime(v string) *ListAlarmHistoriesResponseBodyAlarmHistories {
	s.CreateTime = &v
	return s
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) SetQuotaActionCode(v string) *ListAlarmHistoriesResponseBodyAlarmHistories {
	s.QuotaActionCode = &v
	return s
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) SetAlarmName(v string) *ListAlarmHistoriesResponseBodyAlarmHistories {
	s.AlarmName = &v
	return s
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) SetNotifyTarget(v string) *ListAlarmHistoriesResponseBodyAlarmHistories {
	s.NotifyTarget = &v
	return s
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) SetNotifyChannels(v []*string) *ListAlarmHistoriesResponseBodyAlarmHistories {
	s.NotifyChannels = v
	return s
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) SetThreshold(v float32) *ListAlarmHistoriesResponseBodyAlarmHistories {
	s.Threshold = &v
	return s
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) SetProductCode(v string) *ListAlarmHistoriesResponseBodyAlarmHistories {
	s.ProductCode = &v
	return s
}

type ListAlarmHistoriesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAlarmHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListAlarmHistoriesResponse) SetBody(v *ListAlarmHistoriesResponseBody) *ListAlarmHistoriesResponse {
	s.Body = v
	return s
}

type ListDependentQuotasRequest struct {
	ProductCode     *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
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
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Quotas    []*ListDependentQuotasResponseBodyQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
}

func (s ListDependentQuotasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDependentQuotasResponseBody) GoString() string {
	return s.String()
}

func (s *ListDependentQuotasResponseBody) SetRequestId(v string) *ListDependentQuotasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDependentQuotasResponseBody) SetQuotas(v []*ListDependentQuotasResponseBodyQuotas) *ListDependentQuotasResponseBody {
	s.Quotas = v
	return s
}

type ListDependentQuotasResponseBodyQuotas struct {
	QuotaActionCode *string                                            `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	Dimensions      []*ListDependentQuotasResponseBodyQuotasDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	ProductCode     *string                                            `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	Scale           *float32                                           `json:"Scale,omitempty" xml:"Scale,omitempty"`
}

func (s ListDependentQuotasResponseBodyQuotas) String() string {
	return tea.Prettify(s)
}

func (s ListDependentQuotasResponseBodyQuotas) GoString() string {
	return s.String()
}

func (s *ListDependentQuotasResponseBodyQuotas) SetQuotaActionCode(v string) *ListDependentQuotasResponseBodyQuotas {
	s.QuotaActionCode = &v
	return s
}

func (s *ListDependentQuotasResponseBodyQuotas) SetDimensions(v []*ListDependentQuotasResponseBodyQuotasDimensions) *ListDependentQuotasResponseBodyQuotas {
	s.Dimensions = v
	return s
}

func (s *ListDependentQuotasResponseBodyQuotas) SetProductCode(v string) *ListDependentQuotasResponseBodyQuotas {
	s.ProductCode = &v
	return s
}

func (s *ListDependentQuotasResponseBodyQuotas) SetScale(v float32) *ListDependentQuotasResponseBodyQuotas {
	s.Scale = &v
	return s
}

type ListDependentQuotasResponseBodyQuotasDimensions struct {
	DimensionKey       *string   `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	DependentDimension []*string `json:"DependentDimension,omitempty" xml:"DependentDimension,omitempty" type:"Repeated"`
	DimensionValues    []*string `json:"DimensionValues,omitempty" xml:"DimensionValues,omitempty" type:"Repeated"`
}

func (s ListDependentQuotasResponseBodyQuotasDimensions) String() string {
	return tea.Prettify(s)
}

func (s ListDependentQuotasResponseBodyQuotasDimensions) GoString() string {
	return s.String()
}

func (s *ListDependentQuotasResponseBodyQuotasDimensions) SetDimensionKey(v string) *ListDependentQuotasResponseBodyQuotasDimensions {
	s.DimensionKey = &v
	return s
}

func (s *ListDependentQuotasResponseBodyQuotasDimensions) SetDependentDimension(v []*string) *ListDependentQuotasResponseBodyQuotasDimensions {
	s.DependentDimension = v
	return s
}

func (s *ListDependentQuotasResponseBodyQuotasDimensions) SetDimensionValues(v []*string) *ListDependentQuotasResponseBodyQuotasDimensions {
	s.DimensionValues = v
	return s
}

type ListDependentQuotasResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDependentQuotasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListDependentQuotasResponse) SetBody(v *ListDependentQuotasResponseBody) *ListDependentQuotasResponse {
	s.Body = v
	return s
}

type ListProductQuotaDimensionsRequest struct {
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s ListProductQuotaDimensionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductQuotaDimensionsRequest) GoString() string {
	return s.String()
}

func (s *ListProductQuotaDimensionsRequest) SetNextToken(v string) *ListProductQuotaDimensionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListProductQuotaDimensionsRequest) SetMaxResults(v int32) *ListProductQuotaDimensionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProductQuotaDimensionsRequest) SetProductCode(v string) *ListProductQuotaDimensionsRequest {
	s.ProductCode = &v
	return s
}

type ListProductQuotaDimensionsResponseBody struct {
	QuotaDimensions []*ListProductQuotaDimensionsResponseBodyQuotaDimensions `json:"QuotaDimensions,omitempty" xml:"QuotaDimensions,omitempty" type:"Repeated"`
	TotalCount      *int32                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	NextToken       *string                                                  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId       *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxResults      *int32                                                   `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListProductQuotaDimensionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductQuotaDimensionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductQuotaDimensionsResponseBody) SetQuotaDimensions(v []*ListProductQuotaDimensionsResponseBodyQuotaDimensions) *ListProductQuotaDimensionsResponseBody {
	s.QuotaDimensions = v
	return s
}

func (s *ListProductQuotaDimensionsResponseBody) SetTotalCount(v int32) *ListProductQuotaDimensionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListProductQuotaDimensionsResponseBody) SetNextToken(v string) *ListProductQuotaDimensionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProductQuotaDimensionsResponseBody) SetRequestId(v string) *ListProductQuotaDimensionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductQuotaDimensionsResponseBody) SetMaxResults(v int32) *ListProductQuotaDimensionsResponseBody {
	s.MaxResults = &v
	return s
}

type ListProductQuotaDimensionsResponseBodyQuotaDimensions struct {
	Requisite           *bool     `json:"Requisite,omitempty" xml:"Requisite,omitempty"`
	DimensionKey        *string   `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	DependentDimensions []*string `json:"DependentDimensions,omitempty" xml:"DependentDimensions,omitempty" type:"Repeated"`
	DimensionValues     []*string `json:"DimensionValues,omitempty" xml:"DimensionValues,omitempty" type:"Repeated"`
	Name                *string   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListProductQuotaDimensionsResponseBodyQuotaDimensions) String() string {
	return tea.Prettify(s)
}

func (s ListProductQuotaDimensionsResponseBodyQuotaDimensions) GoString() string {
	return s.String()
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensions) SetRequisite(v bool) *ListProductQuotaDimensionsResponseBodyQuotaDimensions {
	s.Requisite = &v
	return s
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensions) SetDimensionKey(v string) *ListProductQuotaDimensionsResponseBodyQuotaDimensions {
	s.DimensionKey = &v
	return s
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensions) SetDependentDimensions(v []*string) *ListProductQuotaDimensionsResponseBodyQuotaDimensions {
	s.DependentDimensions = v
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

type ListProductQuotaDimensionsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProductQuotaDimensionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListProductQuotaDimensionsResponse) SetBody(v *ListProductQuotaDimensionsResponseBody) *ListProductQuotaDimensionsResponse {
	s.Body = v
	return s
}

type ListProductQuotasRequest struct {
	NextToken       *string                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults      *int32                                `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	ProductCode     *string                               `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	QuotaActionCode *string                               `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	KeyWord         *string                               `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	SortField       *string                               `json:"SortField,omitempty" xml:"SortField,omitempty"`
	SortOrder       *string                               `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	Dimensions      []*ListProductQuotasRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
}

func (s ListProductQuotasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductQuotasRequest) GoString() string {
	return s.String()
}

func (s *ListProductQuotasRequest) SetNextToken(v string) *ListProductQuotasRequest {
	s.NextToken = &v
	return s
}

func (s *ListProductQuotasRequest) SetMaxResults(v int32) *ListProductQuotasRequest {
	s.MaxResults = &v
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

func (s *ListProductQuotasRequest) SetKeyWord(v string) *ListProductQuotasRequest {
	s.KeyWord = &v
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

func (s *ListProductQuotasRequest) SetDimensions(v []*ListProductQuotasRequestDimensions) *ListProductQuotasRequest {
	s.Dimensions = v
	return s
}

type ListProductQuotasRequestDimensions struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	TotalCount *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	NextToken  *string                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Quotas     []*ListProductQuotasResponseBodyQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	MaxResults *int32                                 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListProductQuotasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductQuotasResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductQuotasResponseBody) SetTotalCount(v int32) *ListProductQuotasResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListProductQuotasResponseBody) SetNextToken(v string) *ListProductQuotasResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProductQuotasResponseBody) SetRequestId(v string) *ListProductQuotasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductQuotasResponseBody) SetQuotas(v []*ListProductQuotasResponseBodyQuotas) *ListProductQuotasResponseBody {
	s.Quotas = v
	return s
}

func (s *ListProductQuotasResponseBody) SetMaxResults(v int32) *ListProductQuotasResponseBody {
	s.MaxResults = &v
	return s
}

type ListProductQuotasResponseBodyQuotas struct {
	QuotaUnit          *string                                          `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	QuotaActionCode    *string                                          `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	TotalUsage         *float32                                         `json:"TotalUsage,omitempty" xml:"TotalUsage,omitempty"`
	ApplicableRange    []*float32                                       `json:"ApplicableRange,omitempty" xml:"ApplicableRange,omitempty" type:"Repeated"`
	QuotaType          *string                                          `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	QuotaDescription   *string                                          `json:"QuotaDescription,omitempty" xml:"QuotaDescription,omitempty"`
	Period             *ListProductQuotasResponseBodyQuotasPeriod       `json:"Period,omitempty" xml:"Period,omitempty" type:"Struct"`
	QuotaArn           *string                                          `json:"QuotaArn,omitempty" xml:"QuotaArn,omitempty"`
	ApplicableType     *string                                          `json:"ApplicableType,omitempty" xml:"ApplicableType,omitempty"`
	QuotaItems         []*ListProductQuotasResponseBodyQuotasQuotaItems `json:"QuotaItems,omitempty" xml:"QuotaItems,omitempty" type:"Repeated"`
	Dimensions         map[string]interface{}                           `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	Adjustable         *bool                                            `json:"Adjustable,omitempty" xml:"Adjustable,omitempty"`
	QuotaName          *string                                          `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	UnadjustableDetail *string                                          `json:"UnadjustableDetail,omitempty" xml:"UnadjustableDetail,omitempty"`
	Consumable         *bool                                            `json:"Consumable,omitempty" xml:"Consumable,omitempty"`
	TotalQuota         *float32                                         `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	ProductCode        *string                                          `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s ListProductQuotasResponseBodyQuotas) String() string {
	return tea.Prettify(s)
}

func (s ListProductQuotasResponseBodyQuotas) GoString() string {
	return s.String()
}

func (s *ListProductQuotasResponseBodyQuotas) SetQuotaUnit(v string) *ListProductQuotasResponseBodyQuotas {
	s.QuotaUnit = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetQuotaActionCode(v string) *ListProductQuotasResponseBodyQuotas {
	s.QuotaActionCode = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetTotalUsage(v float32) *ListProductQuotasResponseBodyQuotas {
	s.TotalUsage = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetApplicableRange(v []*float32) *ListProductQuotasResponseBodyQuotas {
	s.ApplicableRange = v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetQuotaType(v string) *ListProductQuotasResponseBodyQuotas {
	s.QuotaType = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetQuotaDescription(v string) *ListProductQuotasResponseBodyQuotas {
	s.QuotaDescription = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetPeriod(v *ListProductQuotasResponseBodyQuotasPeriod) *ListProductQuotasResponseBodyQuotas {
	s.Period = v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetQuotaArn(v string) *ListProductQuotasResponseBodyQuotas {
	s.QuotaArn = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetApplicableType(v string) *ListProductQuotasResponseBodyQuotas {
	s.ApplicableType = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetQuotaItems(v []*ListProductQuotasResponseBodyQuotasQuotaItems) *ListProductQuotasResponseBodyQuotas {
	s.QuotaItems = v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetDimensions(v map[string]interface{}) *ListProductQuotasResponseBodyQuotas {
	s.Dimensions = v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetAdjustable(v bool) *ListProductQuotasResponseBodyQuotas {
	s.Adjustable = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetQuotaName(v string) *ListProductQuotasResponseBodyQuotas {
	s.QuotaName = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetUnadjustableDetail(v string) *ListProductQuotasResponseBodyQuotas {
	s.UnadjustableDetail = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetConsumable(v bool) *ListProductQuotasResponseBodyQuotas {
	s.Consumable = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetTotalQuota(v float32) *ListProductQuotasResponseBodyQuotas {
	s.TotalQuota = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetProductCode(v string) *ListProductQuotasResponseBodyQuotas {
	s.ProductCode = &v
	return s
}

type ListProductQuotasResponseBodyQuotasPeriod struct {
	PeriodValue *int32  `json:"PeriodValue,omitempty" xml:"PeriodValue,omitempty"`
	PeriodUnit  *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
}

func (s ListProductQuotasResponseBodyQuotasPeriod) String() string {
	return tea.Prettify(s)
}

func (s ListProductQuotasResponseBodyQuotasPeriod) GoString() string {
	return s.String()
}

func (s *ListProductQuotasResponseBodyQuotasPeriod) SetPeriodValue(v int32) *ListProductQuotasResponseBodyQuotasPeriod {
	s.PeriodValue = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotasPeriod) SetPeriodUnit(v string) *ListProductQuotasResponseBodyQuotasPeriod {
	s.PeriodUnit = &v
	return s
}

type ListProductQuotasResponseBodyQuotasQuotaItems struct {
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Quota     *string `json:"Quota,omitempty" xml:"Quota,omitempty"`
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	Usage     *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s ListProductQuotasResponseBodyQuotasQuotaItems) String() string {
	return tea.Prettify(s)
}

func (s ListProductQuotasResponseBodyQuotasQuotaItems) GoString() string {
	return s.String()
}

func (s *ListProductQuotasResponseBodyQuotasQuotaItems) SetType(v string) *ListProductQuotasResponseBodyQuotasQuotaItems {
	s.Type = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotasQuotaItems) SetQuota(v string) *ListProductQuotasResponseBodyQuotasQuotaItems {
	s.Quota = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotasQuotaItems) SetQuotaUnit(v string) *ListProductQuotasResponseBodyQuotasQuotaItems {
	s.QuotaUnit = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotasQuotaItems) SetUsage(v string) *ListProductQuotasResponseBodyQuotasQuotaItems {
	s.Usage = &v
	return s
}

type ListProductQuotasResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProductQuotasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListProductQuotasResponse) SetBody(v *ListProductQuotasResponseBody) *ListProductQuotasResponse {
	s.Body = v
	return s
}

type ListProductsRequest struct {
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListProductsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductsRequest) GoString() string {
	return s.String()
}

func (s *ListProductsRequest) SetNextToken(v string) *ListProductsRequest {
	s.NextToken = &v
	return s
}

func (s *ListProductsRequest) SetMaxResults(v int32) *ListProductsRequest {
	s.MaxResults = &v
	return s
}

type ListProductsResponseBody struct {
	ProductInfo []*ListProductsResponseBodyProductInfo `json:"ProductInfo,omitempty" xml:"ProductInfo,omitempty" type:"Repeated"`
	TotalCount  *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	NextToken   *string                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxResults  *int32                                 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListProductsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBody) SetProductInfo(v []*ListProductsResponseBodyProductInfo) *ListProductsResponseBody {
	s.ProductInfo = v
	return s
}

func (s *ListProductsResponseBody) SetTotalCount(v int32) *ListProductsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListProductsResponseBody) SetNextToken(v string) *ListProductsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProductsResponseBody) SetRequestId(v string) *ListProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductsResponseBody) SetMaxResults(v int32) *ListProductsResponseBody {
	s.MaxResults = &v
	return s
}

type ListProductsResponseBodyProductInfo struct {
	ProductName          *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	SecondCategoryId     *int64  `json:"SecondCategoryId,omitempty" xml:"SecondCategoryId,omitempty"`
	ProductNameEn        *string `json:"ProductNameEn,omitempty" xml:"ProductNameEn,omitempty"`
	Dynamic              *bool   `json:"Dynamic,omitempty" xml:"Dynamic,omitempty"`
	SecondCategoryNameEn *string `json:"SecondCategoryNameEn,omitempty" xml:"SecondCategoryNameEn,omitempty"`
	SecondCategoryName   *string `json:"SecondCategoryName,omitempty" xml:"SecondCategoryName,omitempty"`
	ProductCode          *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s ListProductsResponseBodyProductInfo) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBodyProductInfo) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyProductInfo) SetProductName(v string) *ListProductsResponseBodyProductInfo {
	s.ProductName = &v
	return s
}

func (s *ListProductsResponseBodyProductInfo) SetSecondCategoryId(v int64) *ListProductsResponseBodyProductInfo {
	s.SecondCategoryId = &v
	return s
}

func (s *ListProductsResponseBodyProductInfo) SetProductNameEn(v string) *ListProductsResponseBodyProductInfo {
	s.ProductNameEn = &v
	return s
}

func (s *ListProductsResponseBodyProductInfo) SetDynamic(v bool) *ListProductsResponseBodyProductInfo {
	s.Dynamic = &v
	return s
}

func (s *ListProductsResponseBodyProductInfo) SetSecondCategoryNameEn(v string) *ListProductsResponseBodyProductInfo {
	s.SecondCategoryNameEn = &v
	return s
}

func (s *ListProductsResponseBodyProductInfo) SetSecondCategoryName(v string) *ListProductsResponseBodyProductInfo {
	s.SecondCategoryName = &v
	return s
}

func (s *ListProductsResponseBodyProductInfo) SetProductCode(v string) *ListProductsResponseBodyProductInfo {
	s.ProductCode = &v
	return s
}

type ListProductsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProductsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListProductsResponse) SetBody(v *ListProductsResponseBody) *ListProductsResponse {
	s.Body = v
	return s
}

type ListQuotaAlarmsRequest struct {
	NextToken       *string                                  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults      *int32                                   `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	ProductCode     *string                                  `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	AlarmName       *string                                  `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	QuotaActionCode *string                                  `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	QuotaDimensions []*ListQuotaAlarmsRequestQuotaDimensions `json:"QuotaDimensions,omitempty" xml:"QuotaDimensions,omitempty" type:"Repeated"`
}

func (s ListQuotaAlarmsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaAlarmsRequest) GoString() string {
	return s.String()
}

func (s *ListQuotaAlarmsRequest) SetNextToken(v string) *ListQuotaAlarmsRequest {
	s.NextToken = &v
	return s
}

func (s *ListQuotaAlarmsRequest) SetMaxResults(v int32) *ListQuotaAlarmsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListQuotaAlarmsRequest) SetProductCode(v string) *ListQuotaAlarmsRequest {
	s.ProductCode = &v
	return s
}

func (s *ListQuotaAlarmsRequest) SetAlarmName(v string) *ListQuotaAlarmsRequest {
	s.AlarmName = &v
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	TotalCount  *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	QuotaAlarms []*ListQuotaAlarmsResponseBodyQuotaAlarms `json:"QuotaAlarms,omitempty" xml:"QuotaAlarms,omitempty" type:"Repeated"`
	NextToken   *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxResults  *int32                                    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListQuotaAlarmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaAlarmsResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotaAlarmsResponseBody) SetTotalCount(v int32) *ListQuotaAlarmsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListQuotaAlarmsResponseBody) SetQuotaAlarms(v []*ListQuotaAlarmsResponseBodyQuotaAlarms) *ListQuotaAlarmsResponseBody {
	s.QuotaAlarms = v
	return s
}

func (s *ListQuotaAlarmsResponseBody) SetNextToken(v string) *ListQuotaAlarmsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListQuotaAlarmsResponseBody) SetRequestId(v string) *ListQuotaAlarmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuotaAlarmsResponseBody) SetMaxResults(v int32) *ListQuotaAlarmsResponseBody {
	s.MaxResults = &v
	return s
}

type ListQuotaAlarmsResponseBodyQuotaAlarms struct {
	ThresholdPercent *float32               `json:"ThresholdPercent,omitempty" xml:"ThresholdPercent,omitempty"`
	QuotaDimensions  map[string]interface{} `json:"QuotaDimensions,omitempty" xml:"QuotaDimensions,omitempty"`
	CreateTime       *string                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	QuotaActionCode  *string                `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	AlarmName        *string                `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	NotifyTarget     *string                `json:"NotifyTarget,omitempty" xml:"NotifyTarget,omitempty"`
	NotifyChannels   []*string              `json:"NotifyChannels,omitempty" xml:"NotifyChannels,omitempty" type:"Repeated"`
	QuotaUsage       *float32               `json:"QuotaUsage,omitempty" xml:"QuotaUsage,omitempty"`
	QuotaValue       *float32               `json:"QuotaValue,omitempty" xml:"QuotaValue,omitempty"`
	AlarmId          *string                `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	Threshold        *float32               `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	ProductCode      *string                `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	WebHook          *string                `json:"WebHook,omitempty" xml:"WebHook,omitempty"`
	ExceedThreshold  *bool                  `json:"ExceedThreshold,omitempty" xml:"ExceedThreshold,omitempty"`
}

func (s ListQuotaAlarmsResponseBodyQuotaAlarms) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaAlarmsResponseBodyQuotaAlarms) GoString() string {
	return s.String()
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetThresholdPercent(v float32) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.ThresholdPercent = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetQuotaDimensions(v map[string]interface{}) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.QuotaDimensions = v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetCreateTime(v string) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.CreateTime = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetQuotaActionCode(v string) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.QuotaActionCode = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetAlarmName(v string) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.AlarmName = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetNotifyTarget(v string) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.NotifyTarget = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetNotifyChannels(v []*string) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.NotifyChannels = v
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

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetAlarmId(v string) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.AlarmId = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetThreshold(v float32) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.Threshold = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetProductCode(v string) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.ProductCode = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetWebHook(v string) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.WebHook = &v
	return s
}

func (s *ListQuotaAlarmsResponseBodyQuotaAlarms) SetExceedThreshold(v bool) *ListQuotaAlarmsResponseBodyQuotaAlarms {
	s.ExceedThreshold = &v
	return s
}

type ListQuotaAlarmsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListQuotaAlarmsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListQuotaAlarmsResponse) SetBody(v *ListQuotaAlarmsResponseBody) *ListQuotaAlarmsResponse {
	s.Body = v
	return s
}

type ListQuotaApplicationsRequest struct {
	NextToken       *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults      *int32                                    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	ProductCode     *string                                   `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	QuotaActionCode *string                                   `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	Status          *string                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	KeyWord         *string                                   `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	Dimensions      []*ListQuotaApplicationsRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
}

func (s ListQuotaApplicationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsRequest) SetNextToken(v string) *ListQuotaApplicationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListQuotaApplicationsRequest) SetMaxResults(v int32) *ListQuotaApplicationsRequest {
	s.MaxResults = &v
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

func (s *ListQuotaApplicationsRequest) SetStatus(v string) *ListQuotaApplicationsRequest {
	s.Status = &v
	return s
}

func (s *ListQuotaApplicationsRequest) SetKeyWord(v string) *ListQuotaApplicationsRequest {
	s.KeyWord = &v
	return s
}

func (s *ListQuotaApplicationsRequest) SetDimensions(v []*ListQuotaApplicationsRequestDimensions) *ListQuotaApplicationsRequest {
	s.Dimensions = v
	return s
}

type ListQuotaApplicationsRequestDimensions struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	TotalCount        *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	QuotaApplications []*ListQuotaApplicationsResponseBodyQuotaApplications `json:"QuotaApplications,omitempty" xml:"QuotaApplications,omitempty" type:"Repeated"`
	NextToken         *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId         *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxResults        *int32                                                `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListQuotaApplicationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsResponseBody) SetTotalCount(v int32) *ListQuotaApplicationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListQuotaApplicationsResponseBody) SetQuotaApplications(v []*ListQuotaApplicationsResponseBodyQuotaApplications) *ListQuotaApplicationsResponseBody {
	s.QuotaApplications = v
	return s
}

func (s *ListQuotaApplicationsResponseBody) SetNextToken(v string) *ListQuotaApplicationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListQuotaApplicationsResponseBody) SetRequestId(v string) *ListQuotaApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuotaApplicationsResponseBody) SetMaxResults(v int32) *ListQuotaApplicationsResponseBody {
	s.MaxResults = &v
	return s
}

type ListQuotaApplicationsResponseBodyQuotaApplications struct {
	Status           *string                `json:"Status,omitempty" xml:"Status,omitempty"`
	Comment          *string                `json:"Comment,omitempty" xml:"Comment,omitempty"`
	ExpireTime       *string                `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	QuotaUnit        *string                `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	DesireValue      *float32               `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	NoticeType       *int32                 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	QuotaActionCode  *string                `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	Dimension        map[string]interface{} `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	QuotaDescription *string                `json:"QuotaDescription,omitempty" xml:"QuotaDescription,omitempty"`
	QuotaArn         *string                `json:"QuotaArn,omitempty" xml:"QuotaArn,omitempty"`
	EffectiveTime    *string                `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	ApproveValue     *float32               `json:"ApproveValue,omitempty" xml:"ApproveValue,omitempty"`
	QuotaName        *string                `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	ApplicationId    *string                `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	AuditReason      *string                `json:"AuditReason,omitempty" xml:"AuditReason,omitempty"`
	Reason           *string                `json:"Reason,omitempty" xml:"Reason,omitempty"`
	ApplyTime        *string                `json:"ApplyTime,omitempty" xml:"ApplyTime,omitempty"`
	ProductCode      *string                `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s ListQuotaApplicationsResponseBodyQuotaApplications) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaApplicationsResponseBodyQuotaApplications) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetStatus(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.Status = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetComment(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.Comment = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetExpireTime(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.ExpireTime = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetQuotaUnit(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.QuotaUnit = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetDesireValue(v float32) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.DesireValue = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetNoticeType(v int32) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.NoticeType = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetQuotaActionCode(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.QuotaActionCode = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetDimension(v map[string]interface{}) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.Dimension = v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetQuotaDescription(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.QuotaDescription = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetQuotaArn(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.QuotaArn = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetEffectiveTime(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.EffectiveTime = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetApproveValue(v float32) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.ApproveValue = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetQuotaName(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.QuotaName = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetApplicationId(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetAuditReason(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.AuditReason = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetReason(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.Reason = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetApplyTime(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.ApplyTime = &v
	return s
}

func (s *ListQuotaApplicationsResponseBodyQuotaApplications) SetProductCode(v string) *ListQuotaApplicationsResponseBodyQuotaApplications {
	s.ProductCode = &v
	return s
}

type ListQuotaApplicationsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListQuotaApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListQuotaApplicationsResponse) SetBody(v *ListQuotaApplicationsResponseBody) *ListQuotaApplicationsResponse {
	s.Body = v
	return s
}

type UpdateQuotaAlarmRequest struct {
	AlarmId          *string  `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	AlarmName        *string  `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	Threshold        *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	ThresholdPercent *float32 `json:"ThresholdPercent,omitempty" xml:"ThresholdPercent,omitempty"`
	WebHook          *string  `json:"WebHook,omitempty" xml:"WebHook,omitempty"`
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

func (s *UpdateQuotaAlarmRequest) SetWebHook(v string) *UpdateQuotaAlarmRequest {
	s.WebHook = &v
	return s
}

type UpdateQuotaAlarmResponseBody struct {
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateQuotaAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	client.EndpointRule = tea.String("")
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

func (client *Client) CreateQuotaAlarmWithOptions(request *CreateQuotaAlarmRequest, runtime *util.RuntimeOptions) (_result *CreateQuotaAlarmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateQuotaAlarmResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateQuotaAlarm"), tea.String("2020-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CreateQuotaApplicationWithOptions(request *CreateQuotaApplicationRequest, runtime *util.RuntimeOptions) (_result *CreateQuotaApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateQuotaApplicationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateQuotaApplication"), tea.String("2020-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DeleteQuotaAlarmWithOptions(request *DeleteQuotaAlarmRequest, runtime *util.RuntimeOptions) (_result *DeleteQuotaAlarmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteQuotaAlarmResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteQuotaAlarm"), tea.String("2020-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) GetProductQuotaWithOptions(request *GetProductQuotaRequest, runtime *util.RuntimeOptions) (_result *GetProductQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetProductQuotaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetProductQuota"), tea.String("2020-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) GetProductQuotaDimensionWithOptions(request *GetProductQuotaDimensionRequest, runtime *util.RuntimeOptions) (_result *GetProductQuotaDimensionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetProductQuotaDimensionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetProductQuotaDimension"), tea.String("2020-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) GetQuotaAlarmWithOptions(request *GetQuotaAlarmRequest, runtime *util.RuntimeOptions) (_result *GetQuotaAlarmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetQuotaAlarmResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetQuotaAlarm"), tea.String("2020-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) GetQuotaApplicationWithOptions(request *GetQuotaApplicationRequest, runtime *util.RuntimeOptions) (_result *GetQuotaApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetQuotaApplicationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetQuotaApplication"), tea.String("2020-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ListAlarmHistoriesWithOptions(request *ListAlarmHistoriesRequest, runtime *util.RuntimeOptions) (_result *ListAlarmHistoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAlarmHistoriesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAlarmHistories"), tea.String("2020-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListDependentQuotasWithOptions(request *ListDependentQuotasRequest, runtime *util.RuntimeOptions) (_result *ListDependentQuotasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDependentQuotasResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDependentQuotas"), tea.String("2020-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ListProductQuotaDimensionsWithOptions(request *ListProductQuotaDimensionsRequest, runtime *util.RuntimeOptions) (_result *ListProductQuotaDimensionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListProductQuotaDimensionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListProductQuotaDimensions"), tea.String("2020-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ListProductQuotasWithOptions(request *ListProductQuotasRequest, runtime *util.RuntimeOptions) (_result *ListProductQuotasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListProductQuotasResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListProductQuotas"), tea.String("2020-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListProductsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListProducts"), tea.String("2020-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListQuotaAlarmsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListQuotaAlarms"), tea.String("2020-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListQuotaApplicationsWithOptions(request *ListQuotaApplicationsRequest, runtime *util.RuntimeOptions) (_result *ListQuotaApplicationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListQuotaApplicationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListQuotaApplications"), tea.String("2020-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) UpdateQuotaAlarmWithOptions(request *UpdateQuotaAlarmRequest, runtime *util.RuntimeOptions) (_result *UpdateQuotaAlarmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateQuotaAlarmResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateQuotaAlarm"), tea.String("2020-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
