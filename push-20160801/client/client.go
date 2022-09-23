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

type BindAliasRequest struct {
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	AppKey    *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	DeviceId  *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s BindAliasRequest) String() string {
	return tea.Prettify(s)
}

func (s BindAliasRequest) GoString() string {
	return s.String()
}

func (s *BindAliasRequest) SetAliasName(v string) *BindAliasRequest {
	s.AliasName = &v
	return s
}

func (s *BindAliasRequest) SetAppKey(v int64) *BindAliasRequest {
	s.AppKey = &v
	return s
}

func (s *BindAliasRequest) SetDeviceId(v string) *BindAliasRequest {
	s.DeviceId = &v
	return s
}

type BindAliasResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindAliasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindAliasResponseBody) GoString() string {
	return s.String()
}

func (s *BindAliasResponseBody) SetRequestId(v string) *BindAliasResponseBody {
	s.RequestId = &v
	return s
}

type BindAliasResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BindAliasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s BindAliasResponse) GoString() string {
	return s.String()
}

func (s *BindAliasResponse) SetHeaders(v map[string]*string) *BindAliasResponse {
	s.Headers = v
	return s
}

func (s *BindAliasResponse) SetStatusCode(v int32) *BindAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *BindAliasResponse) SetBody(v *BindAliasResponseBody) *BindAliasResponse {
	s.Body = v
	return s
}

type BindPhoneRequest struct {
	AppKey      *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	DeviceId    *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s BindPhoneRequest) String() string {
	return tea.Prettify(s)
}

func (s BindPhoneRequest) GoString() string {
	return s.String()
}

func (s *BindPhoneRequest) SetAppKey(v int64) *BindPhoneRequest {
	s.AppKey = &v
	return s
}

func (s *BindPhoneRequest) SetDeviceId(v string) *BindPhoneRequest {
	s.DeviceId = &v
	return s
}

func (s *BindPhoneRequest) SetPhoneNumber(v string) *BindPhoneRequest {
	s.PhoneNumber = &v
	return s
}

type BindPhoneResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindPhoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindPhoneResponseBody) GoString() string {
	return s.String()
}

func (s *BindPhoneResponseBody) SetRequestId(v string) *BindPhoneResponseBody {
	s.RequestId = &v
	return s
}

type BindPhoneResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BindPhoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindPhoneResponse) String() string {
	return tea.Prettify(s)
}

func (s BindPhoneResponse) GoString() string {
	return s.String()
}

func (s *BindPhoneResponse) SetHeaders(v map[string]*string) *BindPhoneResponse {
	s.Headers = v
	return s
}

func (s *BindPhoneResponse) SetStatusCode(v int32) *BindPhoneResponse {
	s.StatusCode = &v
	return s
}

func (s *BindPhoneResponse) SetBody(v *BindPhoneResponseBody) *BindPhoneResponse {
	s.Body = v
	return s
}

type BindTagRequest struct {
	AppKey    *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	ClientKey *string `json:"ClientKey,omitempty" xml:"ClientKey,omitempty"`
	KeyType   *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	TagName   *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s BindTagRequest) String() string {
	return tea.Prettify(s)
}

func (s BindTagRequest) GoString() string {
	return s.String()
}

func (s *BindTagRequest) SetAppKey(v int64) *BindTagRequest {
	s.AppKey = &v
	return s
}

func (s *BindTagRequest) SetClientKey(v string) *BindTagRequest {
	s.ClientKey = &v
	return s
}

func (s *BindTagRequest) SetKeyType(v string) *BindTagRequest {
	s.KeyType = &v
	return s
}

func (s *BindTagRequest) SetTagName(v string) *BindTagRequest {
	s.TagName = &v
	return s
}

type BindTagResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindTagResponseBody) GoString() string {
	return s.String()
}

func (s *BindTagResponseBody) SetRequestId(v string) *BindTagResponseBody {
	s.RequestId = &v
	return s
}

type BindTagResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BindTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindTagResponse) String() string {
	return tea.Prettify(s)
}

func (s BindTagResponse) GoString() string {
	return s.String()
}

func (s *BindTagResponse) SetHeaders(v map[string]*string) *BindTagResponse {
	s.Headers = v
	return s
}

func (s *BindTagResponse) SetStatusCode(v int32) *BindTagResponse {
	s.StatusCode = &v
	return s
}

func (s *BindTagResponse) SetBody(v *BindTagResponseBody) *BindTagResponse {
	s.Body = v
	return s
}

type CancelPushRequest struct {
	AppKey    *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	MessageId *int64 `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s CancelPushRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelPushRequest) GoString() string {
	return s.String()
}

func (s *CancelPushRequest) SetAppKey(v int64) *CancelPushRequest {
	s.AppKey = &v
	return s
}

func (s *CancelPushRequest) SetMessageId(v int64) *CancelPushRequest {
	s.MessageId = &v
	return s
}

type CancelPushResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelPushResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelPushResponseBody) GoString() string {
	return s.String()
}

func (s *CancelPushResponseBody) SetRequestId(v string) *CancelPushResponseBody {
	s.RequestId = &v
	return s
}

type CancelPushResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelPushResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelPushResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelPushResponse) GoString() string {
	return s.String()
}

func (s *CancelPushResponse) SetHeaders(v map[string]*string) *CancelPushResponse {
	s.Headers = v
	return s
}

func (s *CancelPushResponse) SetStatusCode(v int32) *CancelPushResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelPushResponse) SetBody(v *CancelPushResponseBody) *CancelPushResponse {
	s.Body = v
	return s
}

type CheckCertificateRequest struct {
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s CheckCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckCertificateRequest) GoString() string {
	return s.String()
}

func (s *CheckCertificateRequest) SetAppKey(v int64) *CheckCertificateRequest {
	s.AppKey = &v
	return s
}

type CheckCertificateResponseBody struct {
	Android             *bool                                            `json:"Android,omitempty" xml:"Android,omitempty"`
	DevelopmentCertInfo *CheckCertificateResponseBodyDevelopmentCertInfo `json:"DevelopmentCertInfo,omitempty" xml:"DevelopmentCertInfo,omitempty" type:"Struct"`
	IOS                 *bool                                            `json:"IOS,omitempty" xml:"IOS,omitempty"`
	ProductionCertInfo  *CheckCertificateResponseBodyProductionCertInfo  `json:"ProductionCertInfo,omitempty" xml:"ProductionCertInfo,omitempty" type:"Struct"`
	RequestId           *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *CheckCertificateResponseBody) SetAndroid(v bool) *CheckCertificateResponseBody {
	s.Android = &v
	return s
}

func (s *CheckCertificateResponseBody) SetDevelopmentCertInfo(v *CheckCertificateResponseBodyDevelopmentCertInfo) *CheckCertificateResponseBody {
	s.DevelopmentCertInfo = v
	return s
}

func (s *CheckCertificateResponseBody) SetIOS(v bool) *CheckCertificateResponseBody {
	s.IOS = &v
	return s
}

func (s *CheckCertificateResponseBody) SetProductionCertInfo(v *CheckCertificateResponseBodyProductionCertInfo) *CheckCertificateResponseBody {
	s.ProductionCertInfo = v
	return s
}

func (s *CheckCertificateResponseBody) SetRequestId(v string) *CheckCertificateResponseBody {
	s.RequestId = &v
	return s
}

type CheckCertificateResponseBodyDevelopmentCertInfo struct {
	ExipreTime *int64  `json:"ExipreTime,omitempty" xml:"ExipreTime,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CheckCertificateResponseBodyDevelopmentCertInfo) String() string {
	return tea.Prettify(s)
}

func (s CheckCertificateResponseBodyDevelopmentCertInfo) GoString() string {
	return s.String()
}

func (s *CheckCertificateResponseBodyDevelopmentCertInfo) SetExipreTime(v int64) *CheckCertificateResponseBodyDevelopmentCertInfo {
	s.ExipreTime = &v
	return s
}

func (s *CheckCertificateResponseBodyDevelopmentCertInfo) SetStatus(v string) *CheckCertificateResponseBodyDevelopmentCertInfo {
	s.Status = &v
	return s
}

type CheckCertificateResponseBodyProductionCertInfo struct {
	ExipreTime *int64  `json:"ExipreTime,omitempty" xml:"ExipreTime,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CheckCertificateResponseBodyProductionCertInfo) String() string {
	return tea.Prettify(s)
}

func (s CheckCertificateResponseBodyProductionCertInfo) GoString() string {
	return s.String()
}

func (s *CheckCertificateResponseBodyProductionCertInfo) SetExipreTime(v int64) *CheckCertificateResponseBodyProductionCertInfo {
	s.ExipreTime = &v
	return s
}

func (s *CheckCertificateResponseBodyProductionCertInfo) SetStatus(v string) *CheckCertificateResponseBodyProductionCertInfo {
	s.Status = &v
	return s
}

type CheckCertificateResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckCertificateResponse) GoString() string {
	return s.String()
}

func (s *CheckCertificateResponse) SetHeaders(v map[string]*string) *CheckCertificateResponse {
	s.Headers = v
	return s
}

func (s *CheckCertificateResponse) SetStatusCode(v int32) *CheckCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckCertificateResponse) SetBody(v *CheckCertificateResponseBody) *CheckCertificateResponse {
	s.Body = v
	return s
}

type CheckDeviceRequest struct {
	AppKey   *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s CheckDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckDeviceRequest) GoString() string {
	return s.String()
}

func (s *CheckDeviceRequest) SetAppKey(v int64) *CheckDeviceRequest {
	s.AppKey = &v
	return s
}

func (s *CheckDeviceRequest) SetDeviceId(v string) *CheckDeviceRequest {
	s.DeviceId = &v
	return s
}

type CheckDeviceResponseBody struct {
	Available *bool   `json:"Available,omitempty" xml:"Available,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDeviceResponseBody) SetAvailable(v bool) *CheckDeviceResponseBody {
	s.Available = &v
	return s
}

func (s *CheckDeviceResponseBody) SetRequestId(v string) *CheckDeviceResponseBody {
	s.RequestId = &v
	return s
}

type CheckDeviceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckDeviceResponse) GoString() string {
	return s.String()
}

func (s *CheckDeviceResponse) SetHeaders(v map[string]*string) *CheckDeviceResponse {
	s.Headers = v
	return s
}

func (s *CheckDeviceResponse) SetStatusCode(v int32) *CheckDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDeviceResponse) SetBody(v *CheckDeviceResponseBody) *CheckDeviceResponse {
	s.Body = v
	return s
}

type CheckDevicesRequest struct {
	AppKey    *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	DeviceIds *string `json:"DeviceIds,omitempty" xml:"DeviceIds,omitempty"`
}

func (s CheckDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckDevicesRequest) GoString() string {
	return s.String()
}

func (s *CheckDevicesRequest) SetAppKey(v int64) *CheckDevicesRequest {
	s.AppKey = &v
	return s
}

func (s *CheckDevicesRequest) SetDeviceIds(v string) *CheckDevicesRequest {
	s.DeviceIds = &v
	return s
}

type CheckDevicesResponseBody struct {
	DeviceCheckInfos *CheckDevicesResponseBodyDeviceCheckInfos `json:"DeviceCheckInfos,omitempty" xml:"DeviceCheckInfos,omitempty" type:"Struct"`
	RequestId        *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDevicesResponseBody) SetDeviceCheckInfos(v *CheckDevicesResponseBodyDeviceCheckInfos) *CheckDevicesResponseBody {
	s.DeviceCheckInfos = v
	return s
}

func (s *CheckDevicesResponseBody) SetRequestId(v string) *CheckDevicesResponseBody {
	s.RequestId = &v
	return s
}

type CheckDevicesResponseBodyDeviceCheckInfos struct {
	DeviceCheckInfo []*CheckDevicesResponseBodyDeviceCheckInfosDeviceCheckInfo `json:"DeviceCheckInfo,omitempty" xml:"DeviceCheckInfo,omitempty" type:"Repeated"`
}

func (s CheckDevicesResponseBodyDeviceCheckInfos) String() string {
	return tea.Prettify(s)
}

func (s CheckDevicesResponseBodyDeviceCheckInfos) GoString() string {
	return s.String()
}

func (s *CheckDevicesResponseBodyDeviceCheckInfos) SetDeviceCheckInfo(v []*CheckDevicesResponseBodyDeviceCheckInfosDeviceCheckInfo) *CheckDevicesResponseBodyDeviceCheckInfos {
	s.DeviceCheckInfo = v
	return s
}

type CheckDevicesResponseBodyDeviceCheckInfosDeviceCheckInfo struct {
	Available *bool   `json:"Available,omitempty" xml:"Available,omitempty"`
	DeviceId  *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s CheckDevicesResponseBodyDeviceCheckInfosDeviceCheckInfo) String() string {
	return tea.Prettify(s)
}

func (s CheckDevicesResponseBodyDeviceCheckInfosDeviceCheckInfo) GoString() string {
	return s.String()
}

func (s *CheckDevicesResponseBodyDeviceCheckInfosDeviceCheckInfo) SetAvailable(v bool) *CheckDevicesResponseBodyDeviceCheckInfosDeviceCheckInfo {
	s.Available = &v
	return s
}

func (s *CheckDevicesResponseBodyDeviceCheckInfosDeviceCheckInfo) SetDeviceId(v string) *CheckDevicesResponseBodyDeviceCheckInfosDeviceCheckInfo {
	s.DeviceId = &v
	return s
}

type CheckDevicesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckDevicesResponse) GoString() string {
	return s.String()
}

func (s *CheckDevicesResponse) SetHeaders(v map[string]*string) *CheckDevicesResponse {
	s.Headers = v
	return s
}

func (s *CheckDevicesResponse) SetStatusCode(v int32) *CheckDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDevicesResponse) SetBody(v *CheckDevicesResponseBody) *CheckDevicesResponse {
	s.Body = v
	return s
}

type CompleteContinuouslyPushRequest struct {
	AppKey    *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s CompleteContinuouslyPushRequest) String() string {
	return tea.Prettify(s)
}

func (s CompleteContinuouslyPushRequest) GoString() string {
	return s.String()
}

func (s *CompleteContinuouslyPushRequest) SetAppKey(v int64) *CompleteContinuouslyPushRequest {
	s.AppKey = &v
	return s
}

func (s *CompleteContinuouslyPushRequest) SetMessageId(v string) *CompleteContinuouslyPushRequest {
	s.MessageId = &v
	return s
}

type CompleteContinuouslyPushResponseBody struct {
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CompleteContinuouslyPushResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CompleteContinuouslyPushResponseBody) GoString() string {
	return s.String()
}

func (s *CompleteContinuouslyPushResponseBody) SetMessageId(v string) *CompleteContinuouslyPushResponseBody {
	s.MessageId = &v
	return s
}

func (s *CompleteContinuouslyPushResponseBody) SetRequestId(v string) *CompleteContinuouslyPushResponseBody {
	s.RequestId = &v
	return s
}

type CompleteContinuouslyPushResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CompleteContinuouslyPushResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CompleteContinuouslyPushResponse) String() string {
	return tea.Prettify(s)
}

func (s CompleteContinuouslyPushResponse) GoString() string {
	return s.String()
}

func (s *CompleteContinuouslyPushResponse) SetHeaders(v map[string]*string) *CompleteContinuouslyPushResponse {
	s.Headers = v
	return s
}

func (s *CompleteContinuouslyPushResponse) SetStatusCode(v int32) *CompleteContinuouslyPushResponse {
	s.StatusCode = &v
	return s
}

func (s *CompleteContinuouslyPushResponse) SetBody(v *CompleteContinuouslyPushResponseBody) *CompleteContinuouslyPushResponse {
	s.Body = v
	return s
}

type ContinuouslyPushRequest struct {
	AppKey      *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	MessageId   *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	Target      *string `json:"Target,omitempty" xml:"Target,omitempty"`
	TargetValue *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
}

func (s ContinuouslyPushRequest) String() string {
	return tea.Prettify(s)
}

func (s ContinuouslyPushRequest) GoString() string {
	return s.String()
}

func (s *ContinuouslyPushRequest) SetAppKey(v int64) *ContinuouslyPushRequest {
	s.AppKey = &v
	return s
}

func (s *ContinuouslyPushRequest) SetMessageId(v string) *ContinuouslyPushRequest {
	s.MessageId = &v
	return s
}

func (s *ContinuouslyPushRequest) SetTarget(v string) *ContinuouslyPushRequest {
	s.Target = &v
	return s
}

func (s *ContinuouslyPushRequest) SetTargetValue(v string) *ContinuouslyPushRequest {
	s.TargetValue = &v
	return s
}

type ContinuouslyPushResponseBody struct {
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ContinuouslyPushResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ContinuouslyPushResponseBody) GoString() string {
	return s.String()
}

func (s *ContinuouslyPushResponseBody) SetMessageId(v string) *ContinuouslyPushResponseBody {
	s.MessageId = &v
	return s
}

func (s *ContinuouslyPushResponseBody) SetRequestId(v string) *ContinuouslyPushResponseBody {
	s.RequestId = &v
	return s
}

type ContinuouslyPushResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ContinuouslyPushResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ContinuouslyPushResponse) String() string {
	return tea.Prettify(s)
}

func (s ContinuouslyPushResponse) GoString() string {
	return s.String()
}

func (s *ContinuouslyPushResponse) SetHeaders(v map[string]*string) *ContinuouslyPushResponse {
	s.Headers = v
	return s
}

func (s *ContinuouslyPushResponse) SetStatusCode(v int32) *ContinuouslyPushResponse {
	s.StatusCode = &v
	return s
}

func (s *ContinuouslyPushResponse) SetBody(v *ContinuouslyPushResponseBody) *ContinuouslyPushResponse {
	s.Body = v
	return s
}

type ListSummaryAppsResponseBody struct {
	RequestId       *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SummaryAppInfos *ListSummaryAppsResponseBodySummaryAppInfos `json:"SummaryAppInfos,omitempty" xml:"SummaryAppInfos,omitempty" type:"Struct"`
}

func (s ListSummaryAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSummaryAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSummaryAppsResponseBody) SetRequestId(v string) *ListSummaryAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSummaryAppsResponseBody) SetSummaryAppInfos(v *ListSummaryAppsResponseBodySummaryAppInfos) *ListSummaryAppsResponseBody {
	s.SummaryAppInfos = v
	return s
}

type ListSummaryAppsResponseBodySummaryAppInfos struct {
	SummaryAppInfo []*ListSummaryAppsResponseBodySummaryAppInfosSummaryAppInfo `json:"SummaryAppInfo,omitempty" xml:"SummaryAppInfo,omitempty" type:"Repeated"`
}

func (s ListSummaryAppsResponseBodySummaryAppInfos) String() string {
	return tea.Prettify(s)
}

func (s ListSummaryAppsResponseBodySummaryAppInfos) GoString() string {
	return s.String()
}

func (s *ListSummaryAppsResponseBodySummaryAppInfos) SetSummaryAppInfo(v []*ListSummaryAppsResponseBodySummaryAppInfosSummaryAppInfo) *ListSummaryAppsResponseBodySummaryAppInfos {
	s.SummaryAppInfo = v
	return s
}

type ListSummaryAppsResponseBodySummaryAppInfosSummaryAppInfo struct {
	AppKey  *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s ListSummaryAppsResponseBodySummaryAppInfosSummaryAppInfo) String() string {
	return tea.Prettify(s)
}

func (s ListSummaryAppsResponseBodySummaryAppInfosSummaryAppInfo) GoString() string {
	return s.String()
}

func (s *ListSummaryAppsResponseBodySummaryAppInfosSummaryAppInfo) SetAppKey(v int64) *ListSummaryAppsResponseBodySummaryAppInfosSummaryAppInfo {
	s.AppKey = &v
	return s
}

func (s *ListSummaryAppsResponseBodySummaryAppInfosSummaryAppInfo) SetAppName(v string) *ListSummaryAppsResponseBodySummaryAppInfosSummaryAppInfo {
	s.AppName = &v
	return s
}

type ListSummaryAppsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSummaryAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSummaryAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSummaryAppsResponse) GoString() string {
	return s.String()
}

func (s *ListSummaryAppsResponse) SetHeaders(v map[string]*string) *ListSummaryAppsResponse {
	s.Headers = v
	return s
}

func (s *ListSummaryAppsResponse) SetStatusCode(v int32) *ListSummaryAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSummaryAppsResponse) SetBody(v *ListSummaryAppsResponseBody) *ListSummaryAppsResponse {
	s.Body = v
	return s
}

type ListTagsRequest struct {
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s ListTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagsRequest) GoString() string {
	return s.String()
}

func (s *ListTagsRequest) SetAppKey(v int64) *ListTagsRequest {
	s.AppKey = &v
	return s
}

type ListTagsResponseBody struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagInfos  *ListTagsResponseBodyTagInfos `json:"TagInfos,omitempty" xml:"TagInfos,omitempty" type:"Struct"`
}

func (s ListTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBody) SetRequestId(v string) *ListTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagsResponseBody) SetTagInfos(v *ListTagsResponseBodyTagInfos) *ListTagsResponseBody {
	s.TagInfos = v
	return s
}

type ListTagsResponseBodyTagInfos struct {
	TagInfo []*ListTagsResponseBodyTagInfosTagInfo `json:"TagInfo,omitempty" xml:"TagInfo,omitempty" type:"Repeated"`
}

func (s ListTagsResponseBodyTagInfos) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponseBodyTagInfos) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBodyTagInfos) SetTagInfo(v []*ListTagsResponseBodyTagInfosTagInfo) *ListTagsResponseBodyTagInfos {
	s.TagInfo = v
	return s
}

type ListTagsResponseBodyTagInfosTagInfo struct {
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s ListTagsResponseBodyTagInfosTagInfo) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponseBodyTagInfosTagInfo) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBodyTagInfosTagInfo) SetTagName(v string) *ListTagsResponseBodyTagInfosTagInfo {
	s.TagName = &v
	return s
}

type ListTagsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponse) GoString() string {
	return s.String()
}

func (s *ListTagsResponse) SetHeaders(v map[string]*string) *ListTagsResponse {
	s.Headers = v
	return s
}

func (s *ListTagsResponse) SetStatusCode(v int32) *ListTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagsResponse) SetBody(v *ListTagsResponseBody) *ListTagsResponse {
	s.Body = v
	return s
}

type MassPushRequest struct {
	AppKey   *int64                     `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	PushTask []*MassPushRequestPushTask `json:"PushTask,omitempty" xml:"PushTask,omitempty" type:"Repeated"`
}

func (s MassPushRequest) String() string {
	return tea.Prettify(s)
}

func (s MassPushRequest) GoString() string {
	return s.String()
}

func (s *MassPushRequest) SetAppKey(v int64) *MassPushRequest {
	s.AppKey = &v
	return s
}

func (s *MassPushRequest) SetPushTask(v []*MassPushRequestPushTask) *MassPushRequest {
	s.PushTask = v
	return s
}

type MassPushRequestPushTask struct {
	AndroidActivity                  *string  `json:"AndroidActivity,omitempty" xml:"AndroidActivity,omitempty"`
	AndroidBigBody                   *string  `json:"AndroidBigBody,omitempty" xml:"AndroidBigBody,omitempty"`
	AndroidBigPictureUrl             *string  `json:"AndroidBigPictureUrl,omitempty" xml:"AndroidBigPictureUrl,omitempty"`
	AndroidBigTitle                  *string  `json:"AndroidBigTitle,omitempty" xml:"AndroidBigTitle,omitempty"`
	AndroidExtParameters             *string  `json:"AndroidExtParameters,omitempty" xml:"AndroidExtParameters,omitempty"`
	AndroidImageUrl                  *string  `json:"AndroidImageUrl,omitempty" xml:"AndroidImageUrl,omitempty"`
	AndroidInboxBody                 *string  `json:"AndroidInboxBody,omitempty" xml:"AndroidInboxBody,omitempty"`
	AndroidMessageHuaweiCategory     *string  `json:"AndroidMessageHuaweiCategory,omitempty" xml:"AndroidMessageHuaweiCategory,omitempty"`
	AndroidMessageHuaweiUrgency      *string  `json:"AndroidMessageHuaweiUrgency,omitempty" xml:"AndroidMessageHuaweiUrgency,omitempty"`
	AndroidMusic                     *string  `json:"AndroidMusic,omitempty" xml:"AndroidMusic,omitempty"`
	AndroidNotificationBarPriority   *int32   `json:"AndroidNotificationBarPriority,omitempty" xml:"AndroidNotificationBarPriority,omitempty"`
	AndroidNotificationBarType       *int32   `json:"AndroidNotificationBarType,omitempty" xml:"AndroidNotificationBarType,omitempty"`
	AndroidNotificationChannel       *string  `json:"AndroidNotificationChannel,omitempty" xml:"AndroidNotificationChannel,omitempty"`
	AndroidNotificationHuaweiChannel *string  `json:"AndroidNotificationHuaweiChannel,omitempty" xml:"AndroidNotificationHuaweiChannel,omitempty"`
	AndroidNotificationNotifyId      *int32   `json:"AndroidNotificationNotifyId,omitempty" xml:"AndroidNotificationNotifyId,omitempty"`
	AndroidNotificationVivoChannel   *string  `json:"AndroidNotificationVivoChannel,omitempty" xml:"AndroidNotificationVivoChannel,omitempty"`
	AndroidNotificationXiaomiChannel *string  `json:"AndroidNotificationXiaomiChannel,omitempty" xml:"AndroidNotificationXiaomiChannel,omitempty"`
	AndroidNotifyType                *string  `json:"AndroidNotifyType,omitempty" xml:"AndroidNotifyType,omitempty"`
	AndroidOpenType                  *string  `json:"AndroidOpenType,omitempty" xml:"AndroidOpenType,omitempty"`
	AndroidOpenUrl                   *string  `json:"AndroidOpenUrl,omitempty" xml:"AndroidOpenUrl,omitempty"`
	AndroidPopupActivity             *string  `json:"AndroidPopupActivity,omitempty" xml:"AndroidPopupActivity,omitempty"`
	AndroidPopupBody                 *string  `json:"AndroidPopupBody,omitempty" xml:"AndroidPopupBody,omitempty"`
	AndroidPopupTitle                *string  `json:"AndroidPopupTitle,omitempty" xml:"AndroidPopupTitle,omitempty"`
	AndroidRemind                    *bool    `json:"AndroidRemind,omitempty" xml:"AndroidRemind,omitempty"`
	AndroidRenderStyle               *string  `json:"AndroidRenderStyle,omitempty" xml:"AndroidRenderStyle,omitempty"`
	AndroidVivoPushMode              *int32   `json:"AndroidVivoPushMode,omitempty" xml:"AndroidVivoPushMode,omitempty"`
	AndroidXiaoMiActivity            *string  `json:"AndroidXiaoMiActivity,omitempty" xml:"AndroidXiaoMiActivity,omitempty"`
	AndroidXiaoMiNotifyBody          *string  `json:"AndroidXiaoMiNotifyBody,omitempty" xml:"AndroidXiaoMiNotifyBody,omitempty"`
	AndroidXiaoMiNotifyTitle         *string  `json:"AndroidXiaoMiNotifyTitle,omitempty" xml:"AndroidXiaoMiNotifyTitle,omitempty"`
	AndroidXiaomiBigPictureUrl       *string  `json:"AndroidXiaomiBigPictureUrl,omitempty" xml:"AndroidXiaomiBigPictureUrl,omitempty"`
	AndroidXiaomiImageUrl            *string  `json:"AndroidXiaomiImageUrl,omitempty" xml:"AndroidXiaomiImageUrl,omitempty"`
	Body                             *string  `json:"Body,omitempty" xml:"Body,omitempty"`
	DeviceType                       *string  `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	ExpireTime                       *string  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	JobKey                           *string  `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
	PushTime                         *string  `json:"PushTime,omitempty" xml:"PushTime,omitempty"`
	PushType                         *string  `json:"PushType,omitempty" xml:"PushType,omitempty"`
	SendChannels                     *string  `json:"SendChannels,omitempty" xml:"SendChannels,omitempty"`
	SendSpeed                        *int32   `json:"SendSpeed,omitempty" xml:"SendSpeed,omitempty"`
	StoreOffline                     *bool    `json:"StoreOffline,omitempty" xml:"StoreOffline,omitempty"`
	Target                           *string  `json:"Target,omitempty" xml:"Target,omitempty"`
	TargetValue                      *string  `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
	Title                            *string  `json:"Title,omitempty" xml:"Title,omitempty"`
	Trim                             *bool    `json:"Trim,omitempty" xml:"Trim,omitempty"`
	IOSApnsEnv                       *string  `json:"iOSApnsEnv,omitempty" xml:"iOSApnsEnv,omitempty"`
	IOSBadge                         *int32   `json:"iOSBadge,omitempty" xml:"iOSBadge,omitempty"`
	IOSBadgeAutoIncrement            *bool    `json:"iOSBadgeAutoIncrement,omitempty" xml:"iOSBadgeAutoIncrement,omitempty"`
	IOSExtParameters                 *string  `json:"iOSExtParameters,omitempty" xml:"iOSExtParameters,omitempty"`
	IOSInterruptionLevel             *string  `json:"iOSInterruptionLevel,omitempty" xml:"iOSInterruptionLevel,omitempty"`
	IOSMusic                         *string  `json:"iOSMusic,omitempty" xml:"iOSMusic,omitempty"`
	IOSMutableContent                *bool    `json:"iOSMutableContent,omitempty" xml:"iOSMutableContent,omitempty"`
	IOSNotificationCategory          *string  `json:"iOSNotificationCategory,omitempty" xml:"iOSNotificationCategory,omitempty"`
	IOSNotificationCollapseId        *string  `json:"iOSNotificationCollapseId,omitempty" xml:"iOSNotificationCollapseId,omitempty"`
	IOSNotificationThreadId          *string  `json:"iOSNotificationThreadId,omitempty" xml:"iOSNotificationThreadId,omitempty"`
	IOSRelevanceScore                *float64 `json:"iOSRelevanceScore,omitempty" xml:"iOSRelevanceScore,omitempty"`
	IOSRemind                        *bool    `json:"iOSRemind,omitempty" xml:"iOSRemind,omitempty"`
	IOSRemindBody                    *string  `json:"iOSRemindBody,omitempty" xml:"iOSRemindBody,omitempty"`
	IOSSilentNotification            *bool    `json:"iOSSilentNotification,omitempty" xml:"iOSSilentNotification,omitempty"`
	IOSSubtitle                      *string  `json:"iOSSubtitle,omitempty" xml:"iOSSubtitle,omitempty"`
}

func (s MassPushRequestPushTask) String() string {
	return tea.Prettify(s)
}

func (s MassPushRequestPushTask) GoString() string {
	return s.String()
}

func (s *MassPushRequestPushTask) SetAndroidActivity(v string) *MassPushRequestPushTask {
	s.AndroidActivity = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidBigBody(v string) *MassPushRequestPushTask {
	s.AndroidBigBody = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidBigPictureUrl(v string) *MassPushRequestPushTask {
	s.AndroidBigPictureUrl = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidBigTitle(v string) *MassPushRequestPushTask {
	s.AndroidBigTitle = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidExtParameters(v string) *MassPushRequestPushTask {
	s.AndroidExtParameters = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidImageUrl(v string) *MassPushRequestPushTask {
	s.AndroidImageUrl = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidInboxBody(v string) *MassPushRequestPushTask {
	s.AndroidInboxBody = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidMessageHuaweiCategory(v string) *MassPushRequestPushTask {
	s.AndroidMessageHuaweiCategory = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidMessageHuaweiUrgency(v string) *MassPushRequestPushTask {
	s.AndroidMessageHuaweiUrgency = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidMusic(v string) *MassPushRequestPushTask {
	s.AndroidMusic = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidNotificationBarPriority(v int32) *MassPushRequestPushTask {
	s.AndroidNotificationBarPriority = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidNotificationBarType(v int32) *MassPushRequestPushTask {
	s.AndroidNotificationBarType = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidNotificationChannel(v string) *MassPushRequestPushTask {
	s.AndroidNotificationChannel = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidNotificationHuaweiChannel(v string) *MassPushRequestPushTask {
	s.AndroidNotificationHuaweiChannel = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidNotificationNotifyId(v int32) *MassPushRequestPushTask {
	s.AndroidNotificationNotifyId = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidNotificationVivoChannel(v string) *MassPushRequestPushTask {
	s.AndroidNotificationVivoChannel = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidNotificationXiaomiChannel(v string) *MassPushRequestPushTask {
	s.AndroidNotificationXiaomiChannel = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidNotifyType(v string) *MassPushRequestPushTask {
	s.AndroidNotifyType = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidOpenType(v string) *MassPushRequestPushTask {
	s.AndroidOpenType = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidOpenUrl(v string) *MassPushRequestPushTask {
	s.AndroidOpenUrl = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidPopupActivity(v string) *MassPushRequestPushTask {
	s.AndroidPopupActivity = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidPopupBody(v string) *MassPushRequestPushTask {
	s.AndroidPopupBody = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidPopupTitle(v string) *MassPushRequestPushTask {
	s.AndroidPopupTitle = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidRemind(v bool) *MassPushRequestPushTask {
	s.AndroidRemind = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidRenderStyle(v string) *MassPushRequestPushTask {
	s.AndroidRenderStyle = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidVivoPushMode(v int32) *MassPushRequestPushTask {
	s.AndroidVivoPushMode = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidXiaoMiActivity(v string) *MassPushRequestPushTask {
	s.AndroidXiaoMiActivity = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidXiaoMiNotifyBody(v string) *MassPushRequestPushTask {
	s.AndroidXiaoMiNotifyBody = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidXiaoMiNotifyTitle(v string) *MassPushRequestPushTask {
	s.AndroidXiaoMiNotifyTitle = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidXiaomiBigPictureUrl(v string) *MassPushRequestPushTask {
	s.AndroidXiaomiBigPictureUrl = &v
	return s
}

func (s *MassPushRequestPushTask) SetAndroidXiaomiImageUrl(v string) *MassPushRequestPushTask {
	s.AndroidXiaomiImageUrl = &v
	return s
}

func (s *MassPushRequestPushTask) SetBody(v string) *MassPushRequestPushTask {
	s.Body = &v
	return s
}

func (s *MassPushRequestPushTask) SetDeviceType(v string) *MassPushRequestPushTask {
	s.DeviceType = &v
	return s
}

func (s *MassPushRequestPushTask) SetExpireTime(v string) *MassPushRequestPushTask {
	s.ExpireTime = &v
	return s
}

func (s *MassPushRequestPushTask) SetJobKey(v string) *MassPushRequestPushTask {
	s.JobKey = &v
	return s
}

func (s *MassPushRequestPushTask) SetPushTime(v string) *MassPushRequestPushTask {
	s.PushTime = &v
	return s
}

func (s *MassPushRequestPushTask) SetPushType(v string) *MassPushRequestPushTask {
	s.PushType = &v
	return s
}

func (s *MassPushRequestPushTask) SetSendChannels(v string) *MassPushRequestPushTask {
	s.SendChannels = &v
	return s
}

func (s *MassPushRequestPushTask) SetSendSpeed(v int32) *MassPushRequestPushTask {
	s.SendSpeed = &v
	return s
}

func (s *MassPushRequestPushTask) SetStoreOffline(v bool) *MassPushRequestPushTask {
	s.StoreOffline = &v
	return s
}

func (s *MassPushRequestPushTask) SetTarget(v string) *MassPushRequestPushTask {
	s.Target = &v
	return s
}

func (s *MassPushRequestPushTask) SetTargetValue(v string) *MassPushRequestPushTask {
	s.TargetValue = &v
	return s
}

func (s *MassPushRequestPushTask) SetTitle(v string) *MassPushRequestPushTask {
	s.Title = &v
	return s
}

func (s *MassPushRequestPushTask) SetTrim(v bool) *MassPushRequestPushTask {
	s.Trim = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSApnsEnv(v string) *MassPushRequestPushTask {
	s.IOSApnsEnv = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSBadge(v int32) *MassPushRequestPushTask {
	s.IOSBadge = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSBadgeAutoIncrement(v bool) *MassPushRequestPushTask {
	s.IOSBadgeAutoIncrement = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSExtParameters(v string) *MassPushRequestPushTask {
	s.IOSExtParameters = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSInterruptionLevel(v string) *MassPushRequestPushTask {
	s.IOSInterruptionLevel = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSMusic(v string) *MassPushRequestPushTask {
	s.IOSMusic = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSMutableContent(v bool) *MassPushRequestPushTask {
	s.IOSMutableContent = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSNotificationCategory(v string) *MassPushRequestPushTask {
	s.IOSNotificationCategory = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSNotificationCollapseId(v string) *MassPushRequestPushTask {
	s.IOSNotificationCollapseId = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSNotificationThreadId(v string) *MassPushRequestPushTask {
	s.IOSNotificationThreadId = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSRelevanceScore(v float64) *MassPushRequestPushTask {
	s.IOSRelevanceScore = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSRemind(v bool) *MassPushRequestPushTask {
	s.IOSRemind = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSRemindBody(v string) *MassPushRequestPushTask {
	s.IOSRemindBody = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSSilentNotification(v bool) *MassPushRequestPushTask {
	s.IOSSilentNotification = &v
	return s
}

func (s *MassPushRequestPushTask) SetIOSSubtitle(v string) *MassPushRequestPushTask {
	s.IOSSubtitle = &v
	return s
}

type MassPushResponseBody struct {
	MessageIds *MassPushResponseBodyMessageIds `json:"MessageIds,omitempty" xml:"MessageIds,omitempty" type:"Struct"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MassPushResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MassPushResponseBody) GoString() string {
	return s.String()
}

func (s *MassPushResponseBody) SetMessageIds(v *MassPushResponseBodyMessageIds) *MassPushResponseBody {
	s.MessageIds = v
	return s
}

func (s *MassPushResponseBody) SetRequestId(v string) *MassPushResponseBody {
	s.RequestId = &v
	return s
}

type MassPushResponseBodyMessageIds struct {
	MessageId []*string `json:"MessageId,omitempty" xml:"MessageId,omitempty" type:"Repeated"`
}

func (s MassPushResponseBodyMessageIds) String() string {
	return tea.Prettify(s)
}

func (s MassPushResponseBodyMessageIds) GoString() string {
	return s.String()
}

func (s *MassPushResponseBodyMessageIds) SetMessageId(v []*string) *MassPushResponseBodyMessageIds {
	s.MessageId = v
	return s
}

type MassPushResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MassPushResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MassPushResponse) String() string {
	return tea.Prettify(s)
}

func (s MassPushResponse) GoString() string {
	return s.String()
}

func (s *MassPushResponse) SetHeaders(v map[string]*string) *MassPushResponse {
	s.Headers = v
	return s
}

func (s *MassPushResponse) SetStatusCode(v int32) *MassPushResponse {
	s.StatusCode = &v
	return s
}

func (s *MassPushResponse) SetBody(v *MassPushResponseBody) *MassPushResponse {
	s.Body = v
	return s
}

type PushRequest struct {
	AndroidActivity                  *string  `json:"AndroidActivity,omitempty" xml:"AndroidActivity,omitempty"`
	AndroidBigBody                   *string  `json:"AndroidBigBody,omitempty" xml:"AndroidBigBody,omitempty"`
	AndroidBigPictureUrl             *string  `json:"AndroidBigPictureUrl,omitempty" xml:"AndroidBigPictureUrl,omitempty"`
	AndroidBigTitle                  *string  `json:"AndroidBigTitle,omitempty" xml:"AndroidBigTitle,omitempty"`
	AndroidExtParameters             *string  `json:"AndroidExtParameters,omitempty" xml:"AndroidExtParameters,omitempty"`
	AndroidImageUrl                  *string  `json:"AndroidImageUrl,omitempty" xml:"AndroidImageUrl,omitempty"`
	AndroidInboxBody                 *string  `json:"AndroidInboxBody,omitempty" xml:"AndroidInboxBody,omitempty"`
	AndroidMessageHuaweiCategory     *string  `json:"AndroidMessageHuaweiCategory,omitempty" xml:"AndroidMessageHuaweiCategory,omitempty"`
	AndroidMessageHuaweiUrgency      *string  `json:"AndroidMessageHuaweiUrgency,omitempty" xml:"AndroidMessageHuaweiUrgency,omitempty"`
	AndroidMusic                     *string  `json:"AndroidMusic,omitempty" xml:"AndroidMusic,omitempty"`
	AndroidNotificationBarPriority   *int32   `json:"AndroidNotificationBarPriority,omitempty" xml:"AndroidNotificationBarPriority,omitempty"`
	AndroidNotificationBarType       *int32   `json:"AndroidNotificationBarType,omitempty" xml:"AndroidNotificationBarType,omitempty"`
	AndroidNotificationChannel       *string  `json:"AndroidNotificationChannel,omitempty" xml:"AndroidNotificationChannel,omitempty"`
	AndroidNotificationHuaweiChannel *string  `json:"AndroidNotificationHuaweiChannel,omitempty" xml:"AndroidNotificationHuaweiChannel,omitempty"`
	AndroidNotificationNotifyId      *int32   `json:"AndroidNotificationNotifyId,omitempty" xml:"AndroidNotificationNotifyId,omitempty"`
	AndroidNotificationVivoChannel   *string  `json:"AndroidNotificationVivoChannel,omitempty" xml:"AndroidNotificationVivoChannel,omitempty"`
	AndroidNotificationXiaomiChannel *string  `json:"AndroidNotificationXiaomiChannel,omitempty" xml:"AndroidNotificationXiaomiChannel,omitempty"`
	AndroidNotifyType                *string  `json:"AndroidNotifyType,omitempty" xml:"AndroidNotifyType,omitempty"`
	AndroidOpenType                  *string  `json:"AndroidOpenType,omitempty" xml:"AndroidOpenType,omitempty"`
	AndroidOpenUrl                   *string  `json:"AndroidOpenUrl,omitempty" xml:"AndroidOpenUrl,omitempty"`
	AndroidPopupActivity             *string  `json:"AndroidPopupActivity,omitempty" xml:"AndroidPopupActivity,omitempty"`
	AndroidPopupBody                 *string  `json:"AndroidPopupBody,omitempty" xml:"AndroidPopupBody,omitempty"`
	AndroidPopupTitle                *string  `json:"AndroidPopupTitle,omitempty" xml:"AndroidPopupTitle,omitempty"`
	AndroidRemind                    *bool    `json:"AndroidRemind,omitempty" xml:"AndroidRemind,omitempty"`
	AndroidRenderStyle               *int32   `json:"AndroidRenderStyle,omitempty" xml:"AndroidRenderStyle,omitempty"`
	AndroidVivoPushMode              *int32   `json:"AndroidVivoPushMode,omitempty" xml:"AndroidVivoPushMode,omitempty"`
	AndroidXiaoMiActivity            *string  `json:"AndroidXiaoMiActivity,omitempty" xml:"AndroidXiaoMiActivity,omitempty"`
	AndroidXiaoMiNotifyBody          *string  `json:"AndroidXiaoMiNotifyBody,omitempty" xml:"AndroidXiaoMiNotifyBody,omitempty"`
	AndroidXiaoMiNotifyTitle         *string  `json:"AndroidXiaoMiNotifyTitle,omitempty" xml:"AndroidXiaoMiNotifyTitle,omitempty"`
	AndroidXiaomiBigPictureUrl       *string  `json:"AndroidXiaomiBigPictureUrl,omitempty" xml:"AndroidXiaomiBigPictureUrl,omitempty"`
	AndroidXiaomiImageUrl            *string  `json:"AndroidXiaomiImageUrl,omitempty" xml:"AndroidXiaomiImageUrl,omitempty"`
	AppKey                           *int64   `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	Body                             *string  `json:"Body,omitempty" xml:"Body,omitempty"`
	DeviceType                       *string  `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	ExpireTime                       *string  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	JobKey                           *string  `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
	PushTime                         *string  `json:"PushTime,omitempty" xml:"PushTime,omitempty"`
	PushType                         *string  `json:"PushType,omitempty" xml:"PushType,omitempty"`
	SendChannels                     *string  `json:"SendChannels,omitempty" xml:"SendChannels,omitempty"`
	SendSpeed                        *int32   `json:"SendSpeed,omitempty" xml:"SendSpeed,omitempty"`
	SmsDelaySecs                     *int32   `json:"SmsDelaySecs,omitempty" xml:"SmsDelaySecs,omitempty"`
	SmsParams                        *string  `json:"SmsParams,omitempty" xml:"SmsParams,omitempty"`
	SmsSendPolicy                    *int32   `json:"SmsSendPolicy,omitempty" xml:"SmsSendPolicy,omitempty"`
	SmsSignName                      *string  `json:"SmsSignName,omitempty" xml:"SmsSignName,omitempty"`
	SmsTemplateName                  *string  `json:"SmsTemplateName,omitempty" xml:"SmsTemplateName,omitempty"`
	StoreOffline                     *bool    `json:"StoreOffline,omitempty" xml:"StoreOffline,omitempty"`
	Target                           *string  `json:"Target,omitempty" xml:"Target,omitempty"`
	TargetValue                      *string  `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
	Title                            *string  `json:"Title,omitempty" xml:"Title,omitempty"`
	Trim                             *bool    `json:"Trim,omitempty" xml:"Trim,omitempty"`
	IOSApnsEnv                       *string  `json:"iOSApnsEnv,omitempty" xml:"iOSApnsEnv,omitempty"`
	IOSBadge                         *int32   `json:"iOSBadge,omitempty" xml:"iOSBadge,omitempty"`
	IOSBadgeAutoIncrement            *bool    `json:"iOSBadgeAutoIncrement,omitempty" xml:"iOSBadgeAutoIncrement,omitempty"`
	IOSExtParameters                 *string  `json:"iOSExtParameters,omitempty" xml:"iOSExtParameters,omitempty"`
	IOSInterruptionLevel             *string  `json:"iOSInterruptionLevel,omitempty" xml:"iOSInterruptionLevel,omitempty"`
	IOSMusic                         *string  `json:"iOSMusic,omitempty" xml:"iOSMusic,omitempty"`
	IOSMutableContent                *bool    `json:"iOSMutableContent,omitempty" xml:"iOSMutableContent,omitempty"`
	IOSNotificationCategory          *string  `json:"iOSNotificationCategory,omitempty" xml:"iOSNotificationCategory,omitempty"`
	IOSNotificationCollapseId        *string  `json:"iOSNotificationCollapseId,omitempty" xml:"iOSNotificationCollapseId,omitempty"`
	IOSNotificationThreadId          *string  `json:"iOSNotificationThreadId,omitempty" xml:"iOSNotificationThreadId,omitempty"`
	IOSRelevanceScore                *float64 `json:"iOSRelevanceScore,omitempty" xml:"iOSRelevanceScore,omitempty"`
	IOSRemind                        *bool    `json:"iOSRemind,omitempty" xml:"iOSRemind,omitempty"`
	IOSRemindBody                    *string  `json:"iOSRemindBody,omitempty" xml:"iOSRemindBody,omitempty"`
	IOSSilentNotification            *bool    `json:"iOSSilentNotification,omitempty" xml:"iOSSilentNotification,omitempty"`
	IOSSubtitle                      *string  `json:"iOSSubtitle,omitempty" xml:"iOSSubtitle,omitempty"`
}

func (s PushRequest) String() string {
	return tea.Prettify(s)
}

func (s PushRequest) GoString() string {
	return s.String()
}

func (s *PushRequest) SetAndroidActivity(v string) *PushRequest {
	s.AndroidActivity = &v
	return s
}

func (s *PushRequest) SetAndroidBigBody(v string) *PushRequest {
	s.AndroidBigBody = &v
	return s
}

func (s *PushRequest) SetAndroidBigPictureUrl(v string) *PushRequest {
	s.AndroidBigPictureUrl = &v
	return s
}

func (s *PushRequest) SetAndroidBigTitle(v string) *PushRequest {
	s.AndroidBigTitle = &v
	return s
}

func (s *PushRequest) SetAndroidExtParameters(v string) *PushRequest {
	s.AndroidExtParameters = &v
	return s
}

func (s *PushRequest) SetAndroidImageUrl(v string) *PushRequest {
	s.AndroidImageUrl = &v
	return s
}

func (s *PushRequest) SetAndroidInboxBody(v string) *PushRequest {
	s.AndroidInboxBody = &v
	return s
}

func (s *PushRequest) SetAndroidMessageHuaweiCategory(v string) *PushRequest {
	s.AndroidMessageHuaweiCategory = &v
	return s
}

func (s *PushRequest) SetAndroidMessageHuaweiUrgency(v string) *PushRequest {
	s.AndroidMessageHuaweiUrgency = &v
	return s
}

func (s *PushRequest) SetAndroidMusic(v string) *PushRequest {
	s.AndroidMusic = &v
	return s
}

func (s *PushRequest) SetAndroidNotificationBarPriority(v int32) *PushRequest {
	s.AndroidNotificationBarPriority = &v
	return s
}

func (s *PushRequest) SetAndroidNotificationBarType(v int32) *PushRequest {
	s.AndroidNotificationBarType = &v
	return s
}

func (s *PushRequest) SetAndroidNotificationChannel(v string) *PushRequest {
	s.AndroidNotificationChannel = &v
	return s
}

func (s *PushRequest) SetAndroidNotificationHuaweiChannel(v string) *PushRequest {
	s.AndroidNotificationHuaweiChannel = &v
	return s
}

func (s *PushRequest) SetAndroidNotificationNotifyId(v int32) *PushRequest {
	s.AndroidNotificationNotifyId = &v
	return s
}

func (s *PushRequest) SetAndroidNotificationVivoChannel(v string) *PushRequest {
	s.AndroidNotificationVivoChannel = &v
	return s
}

func (s *PushRequest) SetAndroidNotificationXiaomiChannel(v string) *PushRequest {
	s.AndroidNotificationXiaomiChannel = &v
	return s
}

func (s *PushRequest) SetAndroidNotifyType(v string) *PushRequest {
	s.AndroidNotifyType = &v
	return s
}

func (s *PushRequest) SetAndroidOpenType(v string) *PushRequest {
	s.AndroidOpenType = &v
	return s
}

func (s *PushRequest) SetAndroidOpenUrl(v string) *PushRequest {
	s.AndroidOpenUrl = &v
	return s
}

func (s *PushRequest) SetAndroidPopupActivity(v string) *PushRequest {
	s.AndroidPopupActivity = &v
	return s
}

func (s *PushRequest) SetAndroidPopupBody(v string) *PushRequest {
	s.AndroidPopupBody = &v
	return s
}

func (s *PushRequest) SetAndroidPopupTitle(v string) *PushRequest {
	s.AndroidPopupTitle = &v
	return s
}

func (s *PushRequest) SetAndroidRemind(v bool) *PushRequest {
	s.AndroidRemind = &v
	return s
}

func (s *PushRequest) SetAndroidRenderStyle(v int32) *PushRequest {
	s.AndroidRenderStyle = &v
	return s
}

func (s *PushRequest) SetAndroidVivoPushMode(v int32) *PushRequest {
	s.AndroidVivoPushMode = &v
	return s
}

func (s *PushRequest) SetAndroidXiaoMiActivity(v string) *PushRequest {
	s.AndroidXiaoMiActivity = &v
	return s
}

func (s *PushRequest) SetAndroidXiaoMiNotifyBody(v string) *PushRequest {
	s.AndroidXiaoMiNotifyBody = &v
	return s
}

func (s *PushRequest) SetAndroidXiaoMiNotifyTitle(v string) *PushRequest {
	s.AndroidXiaoMiNotifyTitle = &v
	return s
}

func (s *PushRequest) SetAndroidXiaomiBigPictureUrl(v string) *PushRequest {
	s.AndroidXiaomiBigPictureUrl = &v
	return s
}

func (s *PushRequest) SetAndroidXiaomiImageUrl(v string) *PushRequest {
	s.AndroidXiaomiImageUrl = &v
	return s
}

func (s *PushRequest) SetAppKey(v int64) *PushRequest {
	s.AppKey = &v
	return s
}

func (s *PushRequest) SetBody(v string) *PushRequest {
	s.Body = &v
	return s
}

func (s *PushRequest) SetDeviceType(v string) *PushRequest {
	s.DeviceType = &v
	return s
}

func (s *PushRequest) SetExpireTime(v string) *PushRequest {
	s.ExpireTime = &v
	return s
}

func (s *PushRequest) SetJobKey(v string) *PushRequest {
	s.JobKey = &v
	return s
}

func (s *PushRequest) SetPushTime(v string) *PushRequest {
	s.PushTime = &v
	return s
}

func (s *PushRequest) SetPushType(v string) *PushRequest {
	s.PushType = &v
	return s
}

func (s *PushRequest) SetSendChannels(v string) *PushRequest {
	s.SendChannels = &v
	return s
}

func (s *PushRequest) SetSendSpeed(v int32) *PushRequest {
	s.SendSpeed = &v
	return s
}

func (s *PushRequest) SetSmsDelaySecs(v int32) *PushRequest {
	s.SmsDelaySecs = &v
	return s
}

func (s *PushRequest) SetSmsParams(v string) *PushRequest {
	s.SmsParams = &v
	return s
}

func (s *PushRequest) SetSmsSendPolicy(v int32) *PushRequest {
	s.SmsSendPolicy = &v
	return s
}

func (s *PushRequest) SetSmsSignName(v string) *PushRequest {
	s.SmsSignName = &v
	return s
}

func (s *PushRequest) SetSmsTemplateName(v string) *PushRequest {
	s.SmsTemplateName = &v
	return s
}

func (s *PushRequest) SetStoreOffline(v bool) *PushRequest {
	s.StoreOffline = &v
	return s
}

func (s *PushRequest) SetTarget(v string) *PushRequest {
	s.Target = &v
	return s
}

func (s *PushRequest) SetTargetValue(v string) *PushRequest {
	s.TargetValue = &v
	return s
}

func (s *PushRequest) SetTitle(v string) *PushRequest {
	s.Title = &v
	return s
}

func (s *PushRequest) SetTrim(v bool) *PushRequest {
	s.Trim = &v
	return s
}

func (s *PushRequest) SetIOSApnsEnv(v string) *PushRequest {
	s.IOSApnsEnv = &v
	return s
}

func (s *PushRequest) SetIOSBadge(v int32) *PushRequest {
	s.IOSBadge = &v
	return s
}

func (s *PushRequest) SetIOSBadgeAutoIncrement(v bool) *PushRequest {
	s.IOSBadgeAutoIncrement = &v
	return s
}

func (s *PushRequest) SetIOSExtParameters(v string) *PushRequest {
	s.IOSExtParameters = &v
	return s
}

func (s *PushRequest) SetIOSInterruptionLevel(v string) *PushRequest {
	s.IOSInterruptionLevel = &v
	return s
}

func (s *PushRequest) SetIOSMusic(v string) *PushRequest {
	s.IOSMusic = &v
	return s
}

func (s *PushRequest) SetIOSMutableContent(v bool) *PushRequest {
	s.IOSMutableContent = &v
	return s
}

func (s *PushRequest) SetIOSNotificationCategory(v string) *PushRequest {
	s.IOSNotificationCategory = &v
	return s
}

func (s *PushRequest) SetIOSNotificationCollapseId(v string) *PushRequest {
	s.IOSNotificationCollapseId = &v
	return s
}

func (s *PushRequest) SetIOSNotificationThreadId(v string) *PushRequest {
	s.IOSNotificationThreadId = &v
	return s
}

func (s *PushRequest) SetIOSRelevanceScore(v float64) *PushRequest {
	s.IOSRelevanceScore = &v
	return s
}

func (s *PushRequest) SetIOSRemind(v bool) *PushRequest {
	s.IOSRemind = &v
	return s
}

func (s *PushRequest) SetIOSRemindBody(v string) *PushRequest {
	s.IOSRemindBody = &v
	return s
}

func (s *PushRequest) SetIOSSilentNotification(v bool) *PushRequest {
	s.IOSSilentNotification = &v
	return s
}

func (s *PushRequest) SetIOSSubtitle(v string) *PushRequest {
	s.IOSSubtitle = &v
	return s
}

type PushResponseBody struct {
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PushResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushResponseBody) GoString() string {
	return s.String()
}

func (s *PushResponseBody) SetMessageId(v string) *PushResponseBody {
	s.MessageId = &v
	return s
}

func (s *PushResponseBody) SetRequestId(v string) *PushResponseBody {
	s.RequestId = &v
	return s
}

type PushResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PushResponseBody  `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PushResponse) String() string {
	return tea.Prettify(s)
}

func (s PushResponse) GoString() string {
	return s.String()
}

func (s *PushResponse) SetHeaders(v map[string]*string) *PushResponse {
	s.Headers = v
	return s
}

func (s *PushResponse) SetStatusCode(v int32) *PushResponse {
	s.StatusCode = &v
	return s
}

func (s *PushResponse) SetBody(v *PushResponseBody) *PushResponse {
	s.Body = v
	return s
}

type PushMessageToAndroidRequest struct {
	AppKey      *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	Body        *string `json:"Body,omitempty" xml:"Body,omitempty"`
	JobKey      *string `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
	Target      *string `json:"Target,omitempty" xml:"Target,omitempty"`
	TargetValue *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s PushMessageToAndroidRequest) String() string {
	return tea.Prettify(s)
}

func (s PushMessageToAndroidRequest) GoString() string {
	return s.String()
}

func (s *PushMessageToAndroidRequest) SetAppKey(v int64) *PushMessageToAndroidRequest {
	s.AppKey = &v
	return s
}

func (s *PushMessageToAndroidRequest) SetBody(v string) *PushMessageToAndroidRequest {
	s.Body = &v
	return s
}

func (s *PushMessageToAndroidRequest) SetJobKey(v string) *PushMessageToAndroidRequest {
	s.JobKey = &v
	return s
}

func (s *PushMessageToAndroidRequest) SetTarget(v string) *PushMessageToAndroidRequest {
	s.Target = &v
	return s
}

func (s *PushMessageToAndroidRequest) SetTargetValue(v string) *PushMessageToAndroidRequest {
	s.TargetValue = &v
	return s
}

func (s *PushMessageToAndroidRequest) SetTitle(v string) *PushMessageToAndroidRequest {
	s.Title = &v
	return s
}

type PushMessageToAndroidResponseBody struct {
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PushMessageToAndroidResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushMessageToAndroidResponseBody) GoString() string {
	return s.String()
}

func (s *PushMessageToAndroidResponseBody) SetMessageId(v string) *PushMessageToAndroidResponseBody {
	s.MessageId = &v
	return s
}

func (s *PushMessageToAndroidResponseBody) SetRequestId(v string) *PushMessageToAndroidResponseBody {
	s.RequestId = &v
	return s
}

type PushMessageToAndroidResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PushMessageToAndroidResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PushMessageToAndroidResponse) String() string {
	return tea.Prettify(s)
}

func (s PushMessageToAndroidResponse) GoString() string {
	return s.String()
}

func (s *PushMessageToAndroidResponse) SetHeaders(v map[string]*string) *PushMessageToAndroidResponse {
	s.Headers = v
	return s
}

func (s *PushMessageToAndroidResponse) SetStatusCode(v int32) *PushMessageToAndroidResponse {
	s.StatusCode = &v
	return s
}

func (s *PushMessageToAndroidResponse) SetBody(v *PushMessageToAndroidResponseBody) *PushMessageToAndroidResponse {
	s.Body = v
	return s
}

type PushMessageToiOSRequest struct {
	AppKey      *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	Body        *string `json:"Body,omitempty" xml:"Body,omitempty"`
	JobKey      *string `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
	Target      *string `json:"Target,omitempty" xml:"Target,omitempty"`
	TargetValue *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s PushMessageToiOSRequest) String() string {
	return tea.Prettify(s)
}

func (s PushMessageToiOSRequest) GoString() string {
	return s.String()
}

func (s *PushMessageToiOSRequest) SetAppKey(v int64) *PushMessageToiOSRequest {
	s.AppKey = &v
	return s
}

func (s *PushMessageToiOSRequest) SetBody(v string) *PushMessageToiOSRequest {
	s.Body = &v
	return s
}

func (s *PushMessageToiOSRequest) SetJobKey(v string) *PushMessageToiOSRequest {
	s.JobKey = &v
	return s
}

func (s *PushMessageToiOSRequest) SetTarget(v string) *PushMessageToiOSRequest {
	s.Target = &v
	return s
}

func (s *PushMessageToiOSRequest) SetTargetValue(v string) *PushMessageToiOSRequest {
	s.TargetValue = &v
	return s
}

func (s *PushMessageToiOSRequest) SetTitle(v string) *PushMessageToiOSRequest {
	s.Title = &v
	return s
}

type PushMessageToiOSResponseBody struct {
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PushMessageToiOSResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushMessageToiOSResponseBody) GoString() string {
	return s.String()
}

func (s *PushMessageToiOSResponseBody) SetMessageId(v string) *PushMessageToiOSResponseBody {
	s.MessageId = &v
	return s
}

func (s *PushMessageToiOSResponseBody) SetRequestId(v string) *PushMessageToiOSResponseBody {
	s.RequestId = &v
	return s
}

type PushMessageToiOSResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PushMessageToiOSResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PushMessageToiOSResponse) String() string {
	return tea.Prettify(s)
}

func (s PushMessageToiOSResponse) GoString() string {
	return s.String()
}

func (s *PushMessageToiOSResponse) SetHeaders(v map[string]*string) *PushMessageToiOSResponse {
	s.Headers = v
	return s
}

func (s *PushMessageToiOSResponse) SetStatusCode(v int32) *PushMessageToiOSResponse {
	s.StatusCode = &v
	return s
}

func (s *PushMessageToiOSResponse) SetBody(v *PushMessageToiOSResponseBody) *PushMessageToiOSResponse {
	s.Body = v
	return s
}

type PushNoticeToAndroidRequest struct {
	AppKey        *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	Body          *string `json:"Body,omitempty" xml:"Body,omitempty"`
	ExtParameters *string `json:"ExtParameters,omitempty" xml:"ExtParameters,omitempty"`
	JobKey        *string `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
	Target        *string `json:"Target,omitempty" xml:"Target,omitempty"`
	TargetValue   *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
	Title         *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s PushNoticeToAndroidRequest) String() string {
	return tea.Prettify(s)
}

func (s PushNoticeToAndroidRequest) GoString() string {
	return s.String()
}

func (s *PushNoticeToAndroidRequest) SetAppKey(v int64) *PushNoticeToAndroidRequest {
	s.AppKey = &v
	return s
}

func (s *PushNoticeToAndroidRequest) SetBody(v string) *PushNoticeToAndroidRequest {
	s.Body = &v
	return s
}

func (s *PushNoticeToAndroidRequest) SetExtParameters(v string) *PushNoticeToAndroidRequest {
	s.ExtParameters = &v
	return s
}

func (s *PushNoticeToAndroidRequest) SetJobKey(v string) *PushNoticeToAndroidRequest {
	s.JobKey = &v
	return s
}

func (s *PushNoticeToAndroidRequest) SetTarget(v string) *PushNoticeToAndroidRequest {
	s.Target = &v
	return s
}

func (s *PushNoticeToAndroidRequest) SetTargetValue(v string) *PushNoticeToAndroidRequest {
	s.TargetValue = &v
	return s
}

func (s *PushNoticeToAndroidRequest) SetTitle(v string) *PushNoticeToAndroidRequest {
	s.Title = &v
	return s
}

type PushNoticeToAndroidResponseBody struct {
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PushNoticeToAndroidResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushNoticeToAndroidResponseBody) GoString() string {
	return s.String()
}

func (s *PushNoticeToAndroidResponseBody) SetMessageId(v string) *PushNoticeToAndroidResponseBody {
	s.MessageId = &v
	return s
}

func (s *PushNoticeToAndroidResponseBody) SetRequestId(v string) *PushNoticeToAndroidResponseBody {
	s.RequestId = &v
	return s
}

type PushNoticeToAndroidResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PushNoticeToAndroidResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PushNoticeToAndroidResponse) String() string {
	return tea.Prettify(s)
}

func (s PushNoticeToAndroidResponse) GoString() string {
	return s.String()
}

func (s *PushNoticeToAndroidResponse) SetHeaders(v map[string]*string) *PushNoticeToAndroidResponse {
	s.Headers = v
	return s
}

func (s *PushNoticeToAndroidResponse) SetStatusCode(v int32) *PushNoticeToAndroidResponse {
	s.StatusCode = &v
	return s
}

func (s *PushNoticeToAndroidResponse) SetBody(v *PushNoticeToAndroidResponseBody) *PushNoticeToAndroidResponse {
	s.Body = v
	return s
}

type PushNoticeToiOSRequest struct {
	ApnsEnv       *string `json:"ApnsEnv,omitempty" xml:"ApnsEnv,omitempty"`
	AppKey        *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	Body          *string `json:"Body,omitempty" xml:"Body,omitempty"`
	ExtParameters *string `json:"ExtParameters,omitempty" xml:"ExtParameters,omitempty"`
	JobKey        *string `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
	Target        *string `json:"Target,omitempty" xml:"Target,omitempty"`
	TargetValue   *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
	Title         *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s PushNoticeToiOSRequest) String() string {
	return tea.Prettify(s)
}

func (s PushNoticeToiOSRequest) GoString() string {
	return s.String()
}

func (s *PushNoticeToiOSRequest) SetApnsEnv(v string) *PushNoticeToiOSRequest {
	s.ApnsEnv = &v
	return s
}

func (s *PushNoticeToiOSRequest) SetAppKey(v int64) *PushNoticeToiOSRequest {
	s.AppKey = &v
	return s
}

func (s *PushNoticeToiOSRequest) SetBody(v string) *PushNoticeToiOSRequest {
	s.Body = &v
	return s
}

func (s *PushNoticeToiOSRequest) SetExtParameters(v string) *PushNoticeToiOSRequest {
	s.ExtParameters = &v
	return s
}

func (s *PushNoticeToiOSRequest) SetJobKey(v string) *PushNoticeToiOSRequest {
	s.JobKey = &v
	return s
}

func (s *PushNoticeToiOSRequest) SetTarget(v string) *PushNoticeToiOSRequest {
	s.Target = &v
	return s
}

func (s *PushNoticeToiOSRequest) SetTargetValue(v string) *PushNoticeToiOSRequest {
	s.TargetValue = &v
	return s
}

func (s *PushNoticeToiOSRequest) SetTitle(v string) *PushNoticeToiOSRequest {
	s.Title = &v
	return s
}

type PushNoticeToiOSResponseBody struct {
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PushNoticeToiOSResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushNoticeToiOSResponseBody) GoString() string {
	return s.String()
}

func (s *PushNoticeToiOSResponseBody) SetMessageId(v string) *PushNoticeToiOSResponseBody {
	s.MessageId = &v
	return s
}

func (s *PushNoticeToiOSResponseBody) SetRequestId(v string) *PushNoticeToiOSResponseBody {
	s.RequestId = &v
	return s
}

type PushNoticeToiOSResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PushNoticeToiOSResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PushNoticeToiOSResponse) String() string {
	return tea.Prettify(s)
}

func (s PushNoticeToiOSResponse) GoString() string {
	return s.String()
}

func (s *PushNoticeToiOSResponse) SetHeaders(v map[string]*string) *PushNoticeToiOSResponse {
	s.Headers = v
	return s
}

func (s *PushNoticeToiOSResponse) SetStatusCode(v int32) *PushNoticeToiOSResponse {
	s.StatusCode = &v
	return s
}

func (s *PushNoticeToiOSResponse) SetBody(v *PushNoticeToiOSResponseBody) *PushNoticeToiOSResponse {
	s.Body = v
	return s
}

type QueryAliasesRequest struct {
	AppKey   *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s QueryAliasesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAliasesRequest) GoString() string {
	return s.String()
}

func (s *QueryAliasesRequest) SetAppKey(v int64) *QueryAliasesRequest {
	s.AppKey = &v
	return s
}

func (s *QueryAliasesRequest) SetDeviceId(v string) *QueryAliasesRequest {
	s.DeviceId = &v
	return s
}

type QueryAliasesResponseBody struct {
	AliasInfos *QueryAliasesResponseBodyAliasInfos `json:"AliasInfos,omitempty" xml:"AliasInfos,omitempty" type:"Struct"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryAliasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAliasesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAliasesResponseBody) SetAliasInfos(v *QueryAliasesResponseBodyAliasInfos) *QueryAliasesResponseBody {
	s.AliasInfos = v
	return s
}

func (s *QueryAliasesResponseBody) SetRequestId(v string) *QueryAliasesResponseBody {
	s.RequestId = &v
	return s
}

type QueryAliasesResponseBodyAliasInfos struct {
	AliasInfo []*QueryAliasesResponseBodyAliasInfosAliasInfo `json:"AliasInfo,omitempty" xml:"AliasInfo,omitempty" type:"Repeated"`
}

func (s QueryAliasesResponseBodyAliasInfos) String() string {
	return tea.Prettify(s)
}

func (s QueryAliasesResponseBodyAliasInfos) GoString() string {
	return s.String()
}

func (s *QueryAliasesResponseBodyAliasInfos) SetAliasInfo(v []*QueryAliasesResponseBodyAliasInfosAliasInfo) *QueryAliasesResponseBodyAliasInfos {
	s.AliasInfo = v
	return s
}

type QueryAliasesResponseBodyAliasInfosAliasInfo struct {
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
}

func (s QueryAliasesResponseBodyAliasInfosAliasInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryAliasesResponseBodyAliasInfosAliasInfo) GoString() string {
	return s.String()
}

func (s *QueryAliasesResponseBodyAliasInfosAliasInfo) SetAliasName(v string) *QueryAliasesResponseBodyAliasInfosAliasInfo {
	s.AliasName = &v
	return s
}

type QueryAliasesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryAliasesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAliasesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAliasesResponse) GoString() string {
	return s.String()
}

func (s *QueryAliasesResponse) SetHeaders(v map[string]*string) *QueryAliasesResponse {
	s.Headers = v
	return s
}

func (s *QueryAliasesResponse) SetStatusCode(v int32) *QueryAliasesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAliasesResponse) SetBody(v *QueryAliasesResponseBody) *QueryAliasesResponse {
	s.Body = v
	return s
}

type QueryDeviceInfoRequest struct {
	AppKey   *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s QueryDeviceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceInfoRequest) SetAppKey(v int64) *QueryDeviceInfoRequest {
	s.AppKey = &v
	return s
}

func (s *QueryDeviceInfoRequest) SetDeviceId(v string) *QueryDeviceInfoRequest {
	s.DeviceId = &v
	return s
}

type QueryDeviceInfoResponseBody struct {
	DeviceInfo *QueryDeviceInfoResponseBodyDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryDeviceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceInfoResponseBody) SetDeviceInfo(v *QueryDeviceInfoResponseBodyDeviceInfo) *QueryDeviceInfoResponseBody {
	s.DeviceInfo = v
	return s
}

func (s *QueryDeviceInfoResponseBody) SetRequestId(v string) *QueryDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

type QueryDeviceInfoResponseBodyDeviceInfo struct {
	Account        *string `json:"Account,omitempty" xml:"Account,omitempty"`
	Alias          *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	DeviceId       *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DeviceToken    *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	DeviceType     *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	LastOnlineTime *string `json:"LastOnlineTime,omitempty" xml:"LastOnlineTime,omitempty"`
	Online         *bool   `json:"Online,omitempty" xml:"Online,omitempty"`
	PhoneNumber    *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	PushEnabled    *bool   `json:"PushEnabled,omitempty" xml:"PushEnabled,omitempty"`
	Tags           *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s QueryDeviceInfoResponseBodyDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceInfoResponseBodyDeviceInfo) GoString() string {
	return s.String()
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) SetAccount(v string) *QueryDeviceInfoResponseBodyDeviceInfo {
	s.Account = &v
	return s
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) SetAlias(v string) *QueryDeviceInfoResponseBodyDeviceInfo {
	s.Alias = &v
	return s
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) SetDeviceId(v string) *QueryDeviceInfoResponseBodyDeviceInfo {
	s.DeviceId = &v
	return s
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) SetDeviceToken(v string) *QueryDeviceInfoResponseBodyDeviceInfo {
	s.DeviceToken = &v
	return s
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) SetDeviceType(v string) *QueryDeviceInfoResponseBodyDeviceInfo {
	s.DeviceType = &v
	return s
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) SetLastOnlineTime(v string) *QueryDeviceInfoResponseBodyDeviceInfo {
	s.LastOnlineTime = &v
	return s
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) SetOnline(v bool) *QueryDeviceInfoResponseBodyDeviceInfo {
	s.Online = &v
	return s
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) SetPhoneNumber(v string) *QueryDeviceInfoResponseBodyDeviceInfo {
	s.PhoneNumber = &v
	return s
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) SetPushEnabled(v bool) *QueryDeviceInfoResponseBodyDeviceInfo {
	s.PushEnabled = &v
	return s
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) SetTags(v string) *QueryDeviceInfoResponseBodyDeviceInfo {
	s.Tags = &v
	return s
}

type QueryDeviceInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDeviceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceInfoResponse) SetHeaders(v map[string]*string) *QueryDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceInfoResponse) SetStatusCode(v int32) *QueryDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceInfoResponse) SetBody(v *QueryDeviceInfoResponseBody) *QueryDeviceInfoResponse {
	s.Body = v
	return s
}

type QueryDeviceStatRequest struct {
	AppKey     *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	QueryType  *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryDeviceStatRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatRequest) SetAppKey(v int64) *QueryDeviceStatRequest {
	s.AppKey = &v
	return s
}

func (s *QueryDeviceStatRequest) SetDeviceType(v string) *QueryDeviceStatRequest {
	s.DeviceType = &v
	return s
}

func (s *QueryDeviceStatRequest) SetEndTime(v string) *QueryDeviceStatRequest {
	s.EndTime = &v
	return s
}

func (s *QueryDeviceStatRequest) SetQueryType(v string) *QueryDeviceStatRequest {
	s.QueryType = &v
	return s
}

func (s *QueryDeviceStatRequest) SetStartTime(v string) *QueryDeviceStatRequest {
	s.StartTime = &v
	return s
}

type QueryDeviceStatResponseBody struct {
	AppDeviceStats *QueryDeviceStatResponseBodyAppDeviceStats `json:"AppDeviceStats,omitempty" xml:"AppDeviceStats,omitempty" type:"Struct"`
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryDeviceStatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatResponseBody) SetAppDeviceStats(v *QueryDeviceStatResponseBodyAppDeviceStats) *QueryDeviceStatResponseBody {
	s.AppDeviceStats = v
	return s
}

func (s *QueryDeviceStatResponseBody) SetRequestId(v string) *QueryDeviceStatResponseBody {
	s.RequestId = &v
	return s
}

type QueryDeviceStatResponseBodyAppDeviceStats struct {
	AppDeviceStat []*QueryDeviceStatResponseBodyAppDeviceStatsAppDeviceStat `json:"AppDeviceStat,omitempty" xml:"AppDeviceStat,omitempty" type:"Repeated"`
}

func (s QueryDeviceStatResponseBodyAppDeviceStats) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatResponseBodyAppDeviceStats) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatResponseBodyAppDeviceStats) SetAppDeviceStat(v []*QueryDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) *QueryDeviceStatResponseBodyAppDeviceStats {
	s.AppDeviceStat = v
	return s
}

type QueryDeviceStatResponseBodyAppDeviceStatsAppDeviceStat struct {
	Count      *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	Time       *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s QueryDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) SetCount(v int64) *QueryDeviceStatResponseBodyAppDeviceStatsAppDeviceStat {
	s.Count = &v
	return s
}

func (s *QueryDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) SetDeviceType(v string) *QueryDeviceStatResponseBodyAppDeviceStatsAppDeviceStat {
	s.DeviceType = &v
	return s
}

func (s *QueryDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) SetTime(v string) *QueryDeviceStatResponseBodyAppDeviceStatsAppDeviceStat {
	s.Time = &v
	return s
}

type QueryDeviceStatResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDeviceStatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDeviceStatResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatResponse) SetHeaders(v map[string]*string) *QueryDeviceStatResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceStatResponse) SetStatusCode(v int32) *QueryDeviceStatResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceStatResponse) SetBody(v *QueryDeviceStatResponseBody) *QueryDeviceStatResponse {
	s.Body = v
	return s
}

type QueryDevicesByAccountRequest struct {
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	AppKey  *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s QueryDevicesByAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicesByAccountRequest) GoString() string {
	return s.String()
}

func (s *QueryDevicesByAccountRequest) SetAccount(v string) *QueryDevicesByAccountRequest {
	s.Account = &v
	return s
}

func (s *QueryDevicesByAccountRequest) SetAppKey(v int64) *QueryDevicesByAccountRequest {
	s.AppKey = &v
	return s
}

type QueryDevicesByAccountResponseBody struct {
	DeviceIds *QueryDevicesByAccountResponseBodyDeviceIds `json:"DeviceIds,omitempty" xml:"DeviceIds,omitempty" type:"Struct"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryDevicesByAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicesByAccountResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDevicesByAccountResponseBody) SetDeviceIds(v *QueryDevicesByAccountResponseBodyDeviceIds) *QueryDevicesByAccountResponseBody {
	s.DeviceIds = v
	return s
}

func (s *QueryDevicesByAccountResponseBody) SetRequestId(v string) *QueryDevicesByAccountResponseBody {
	s.RequestId = &v
	return s
}

type QueryDevicesByAccountResponseBodyDeviceIds struct {
	DeviceId []*string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty" type:"Repeated"`
}

func (s QueryDevicesByAccountResponseBodyDeviceIds) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicesByAccountResponseBodyDeviceIds) GoString() string {
	return s.String()
}

func (s *QueryDevicesByAccountResponseBodyDeviceIds) SetDeviceId(v []*string) *QueryDevicesByAccountResponseBodyDeviceIds {
	s.DeviceId = v
	return s
}

type QueryDevicesByAccountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDevicesByAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDevicesByAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicesByAccountResponse) GoString() string {
	return s.String()
}

func (s *QueryDevicesByAccountResponse) SetHeaders(v map[string]*string) *QueryDevicesByAccountResponse {
	s.Headers = v
	return s
}

func (s *QueryDevicesByAccountResponse) SetStatusCode(v int32) *QueryDevicesByAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDevicesByAccountResponse) SetBody(v *QueryDevicesByAccountResponseBody) *QueryDevicesByAccountResponse {
	s.Body = v
	return s
}

type QueryDevicesByAliasRequest struct {
	Alias  *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	AppKey *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s QueryDevicesByAliasRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicesByAliasRequest) GoString() string {
	return s.String()
}

func (s *QueryDevicesByAliasRequest) SetAlias(v string) *QueryDevicesByAliasRequest {
	s.Alias = &v
	return s
}

func (s *QueryDevicesByAliasRequest) SetAppKey(v int64) *QueryDevicesByAliasRequest {
	s.AppKey = &v
	return s
}

type QueryDevicesByAliasResponseBody struct {
	DeviceIds *QueryDevicesByAliasResponseBodyDeviceIds `json:"DeviceIds,omitempty" xml:"DeviceIds,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryDevicesByAliasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicesByAliasResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDevicesByAliasResponseBody) SetDeviceIds(v *QueryDevicesByAliasResponseBodyDeviceIds) *QueryDevicesByAliasResponseBody {
	s.DeviceIds = v
	return s
}

func (s *QueryDevicesByAliasResponseBody) SetRequestId(v string) *QueryDevicesByAliasResponseBody {
	s.RequestId = &v
	return s
}

type QueryDevicesByAliasResponseBodyDeviceIds struct {
	DeviceId []*string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty" type:"Repeated"`
}

func (s QueryDevicesByAliasResponseBodyDeviceIds) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicesByAliasResponseBodyDeviceIds) GoString() string {
	return s.String()
}

func (s *QueryDevicesByAliasResponseBodyDeviceIds) SetDeviceId(v []*string) *QueryDevicesByAliasResponseBodyDeviceIds {
	s.DeviceId = v
	return s
}

type QueryDevicesByAliasResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDevicesByAliasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDevicesByAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicesByAliasResponse) GoString() string {
	return s.String()
}

func (s *QueryDevicesByAliasResponse) SetHeaders(v map[string]*string) *QueryDevicesByAliasResponse {
	s.Headers = v
	return s
}

func (s *QueryDevicesByAliasResponse) SetStatusCode(v int32) *QueryDevicesByAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDevicesByAliasResponse) SetBody(v *QueryDevicesByAliasResponseBody) *QueryDevicesByAliasResponse {
	s.Body = v
	return s
}

type QueryPushRecordsRequest struct {
	AppKey    *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Keyword   *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Page      *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PushType  *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
	Source    *string `json:"Source,omitempty" xml:"Source,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Target    *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s QueryPushRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPushRecordsRequest) GoString() string {
	return s.String()
}

func (s *QueryPushRecordsRequest) SetAppKey(v int64) *QueryPushRecordsRequest {
	s.AppKey = &v
	return s
}

func (s *QueryPushRecordsRequest) SetEndTime(v string) *QueryPushRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *QueryPushRecordsRequest) SetKeyword(v string) *QueryPushRecordsRequest {
	s.Keyword = &v
	return s
}

func (s *QueryPushRecordsRequest) SetNextToken(v string) *QueryPushRecordsRequest {
	s.NextToken = &v
	return s
}

func (s *QueryPushRecordsRequest) SetPage(v int32) *QueryPushRecordsRequest {
	s.Page = &v
	return s
}

func (s *QueryPushRecordsRequest) SetPageSize(v int32) *QueryPushRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryPushRecordsRequest) SetPushType(v string) *QueryPushRecordsRequest {
	s.PushType = &v
	return s
}

func (s *QueryPushRecordsRequest) SetSource(v string) *QueryPushRecordsRequest {
	s.Source = &v
	return s
}

func (s *QueryPushRecordsRequest) SetStartTime(v string) *QueryPushRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *QueryPushRecordsRequest) SetTarget(v string) *QueryPushRecordsRequest {
	s.Target = &v
	return s
}

type QueryPushRecordsResponseBody struct {
	NextToken *string                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Page      *int32                                 `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize  *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PushInfos *QueryPushRecordsResponseBodyPushInfos `json:"PushInfos,omitempty" xml:"PushInfos,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int32                                 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryPushRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPushRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPushRecordsResponseBody) SetNextToken(v string) *QueryPushRecordsResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryPushRecordsResponseBody) SetPage(v int32) *QueryPushRecordsResponseBody {
	s.Page = &v
	return s
}

func (s *QueryPushRecordsResponseBody) SetPageSize(v int32) *QueryPushRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryPushRecordsResponseBody) SetPushInfos(v *QueryPushRecordsResponseBodyPushInfos) *QueryPushRecordsResponseBody {
	s.PushInfos = v
	return s
}

func (s *QueryPushRecordsResponseBody) SetRequestId(v string) *QueryPushRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPushRecordsResponseBody) SetTotal(v int32) *QueryPushRecordsResponseBody {
	s.Total = &v
	return s
}

type QueryPushRecordsResponseBodyPushInfos struct {
	PushInfo []*QueryPushRecordsResponseBodyPushInfosPushInfo `json:"PushInfo,omitempty" xml:"PushInfo,omitempty" type:"Repeated"`
}

func (s QueryPushRecordsResponseBodyPushInfos) String() string {
	return tea.Prettify(s)
}

func (s QueryPushRecordsResponseBodyPushInfos) GoString() string {
	return s.String()
}

func (s *QueryPushRecordsResponseBodyPushInfos) SetPushInfo(v []*QueryPushRecordsResponseBodyPushInfosPushInfo) *QueryPushRecordsResponseBodyPushInfos {
	s.PushInfo = v
	return s
}

type QueryPushRecordsResponseBodyPushInfosPushInfo struct {
	AppKey     *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	Body       *string `json:"Body,omitempty" xml:"Body,omitempty"`
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	MessageId  *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	PushTime   *string `json:"PushTime,omitempty" xml:"PushTime,omitempty"`
	PushType   *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
	Source     *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Target     *string `json:"Target,omitempty" xml:"Target,omitempty"`
	Title      *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s QueryPushRecordsResponseBodyPushInfosPushInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryPushRecordsResponseBodyPushInfosPushInfo) GoString() string {
	return s.String()
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) SetAppKey(v int64) *QueryPushRecordsResponseBodyPushInfosPushInfo {
	s.AppKey = &v
	return s
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) SetBody(v string) *QueryPushRecordsResponseBodyPushInfosPushInfo {
	s.Body = &v
	return s
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) SetDeviceType(v string) *QueryPushRecordsResponseBodyPushInfosPushInfo {
	s.DeviceType = &v
	return s
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) SetMessageId(v string) *QueryPushRecordsResponseBodyPushInfosPushInfo {
	s.MessageId = &v
	return s
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) SetPushTime(v string) *QueryPushRecordsResponseBodyPushInfosPushInfo {
	s.PushTime = &v
	return s
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) SetPushType(v string) *QueryPushRecordsResponseBodyPushInfosPushInfo {
	s.PushType = &v
	return s
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) SetSource(v string) *QueryPushRecordsResponseBodyPushInfosPushInfo {
	s.Source = &v
	return s
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) SetStatus(v string) *QueryPushRecordsResponseBodyPushInfosPushInfo {
	s.Status = &v
	return s
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) SetTarget(v string) *QueryPushRecordsResponseBodyPushInfosPushInfo {
	s.Target = &v
	return s
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) SetTitle(v string) *QueryPushRecordsResponseBodyPushInfosPushInfo {
	s.Title = &v
	return s
}

type QueryPushRecordsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryPushRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPushRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPushRecordsResponse) GoString() string {
	return s.String()
}

func (s *QueryPushRecordsResponse) SetHeaders(v map[string]*string) *QueryPushRecordsResponse {
	s.Headers = v
	return s
}

func (s *QueryPushRecordsResponse) SetStatusCode(v int32) *QueryPushRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPushRecordsResponse) SetBody(v *QueryPushRecordsResponseBody) *QueryPushRecordsResponse {
	s.Body = v
	return s
}

type QueryPushStatByAppRequest struct {
	AppKey      *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryPushStatByAppRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPushStatByAppRequest) GoString() string {
	return s.String()
}

func (s *QueryPushStatByAppRequest) SetAppKey(v int64) *QueryPushStatByAppRequest {
	s.AppKey = &v
	return s
}

func (s *QueryPushStatByAppRequest) SetEndTime(v string) *QueryPushStatByAppRequest {
	s.EndTime = &v
	return s
}

func (s *QueryPushStatByAppRequest) SetGranularity(v string) *QueryPushStatByAppRequest {
	s.Granularity = &v
	return s
}

func (s *QueryPushStatByAppRequest) SetStartTime(v string) *QueryPushStatByAppRequest {
	s.StartTime = &v
	return s
}

type QueryPushStatByAppResponseBody struct {
	AppPushStats *QueryPushStatByAppResponseBodyAppPushStats `json:"AppPushStats,omitempty" xml:"AppPushStats,omitempty" type:"Struct"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryPushStatByAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPushStatByAppResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPushStatByAppResponseBody) SetAppPushStats(v *QueryPushStatByAppResponseBodyAppPushStats) *QueryPushStatByAppResponseBody {
	s.AppPushStats = v
	return s
}

func (s *QueryPushStatByAppResponseBody) SetRequestId(v string) *QueryPushStatByAppResponseBody {
	s.RequestId = &v
	return s
}

type QueryPushStatByAppResponseBodyAppPushStats struct {
	AppPushStat []*QueryPushStatByAppResponseBodyAppPushStatsAppPushStat `json:"AppPushStat,omitempty" xml:"AppPushStat,omitempty" type:"Repeated"`
}

func (s QueryPushStatByAppResponseBodyAppPushStats) String() string {
	return tea.Prettify(s)
}

func (s QueryPushStatByAppResponseBodyAppPushStats) GoString() string {
	return s.String()
}

func (s *QueryPushStatByAppResponseBodyAppPushStats) SetAppPushStat(v []*QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) *QueryPushStatByAppResponseBodyAppPushStats {
	s.AppPushStat = v
	return s
}

type QueryPushStatByAppResponseBodyAppPushStatsAppPushStat struct {
	AcceptCount            *int64  `json:"AcceptCount,omitempty" xml:"AcceptCount,omitempty"`
	DeletedCount           *int64  `json:"DeletedCount,omitempty" xml:"DeletedCount,omitempty"`
	OpenedCount            *int64  `json:"OpenedCount,omitempty" xml:"OpenedCount,omitempty"`
	ReceivedCount          *int64  `json:"ReceivedCount,omitempty" xml:"ReceivedCount,omitempty"`
	SentCount              *int64  `json:"SentCount,omitempty" xml:"SentCount,omitempty"`
	SmsFailedCount         *int64  `json:"SmsFailedCount,omitempty" xml:"SmsFailedCount,omitempty"`
	SmsReceiveFailedCount  *int64  `json:"SmsReceiveFailedCount,omitempty" xml:"SmsReceiveFailedCount,omitempty"`
	SmsReceiveSuccessCount *int64  `json:"SmsReceiveSuccessCount,omitempty" xml:"SmsReceiveSuccessCount,omitempty"`
	SmsSentCount           *int64  `json:"SmsSentCount,omitempty" xml:"SmsSentCount,omitempty"`
	SmsSkipCount           *int64  `json:"SmsSkipCount,omitempty" xml:"SmsSkipCount,omitempty"`
	Time                   *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) String() string {
	return tea.Prettify(s)
}

func (s QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) GoString() string {
	return s.String()
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) SetAcceptCount(v int64) *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat {
	s.AcceptCount = &v
	return s
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) SetDeletedCount(v int64) *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat {
	s.DeletedCount = &v
	return s
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) SetOpenedCount(v int64) *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat {
	s.OpenedCount = &v
	return s
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) SetReceivedCount(v int64) *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat {
	s.ReceivedCount = &v
	return s
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) SetSentCount(v int64) *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat {
	s.SentCount = &v
	return s
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) SetSmsFailedCount(v int64) *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat {
	s.SmsFailedCount = &v
	return s
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) SetSmsReceiveFailedCount(v int64) *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat {
	s.SmsReceiveFailedCount = &v
	return s
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) SetSmsReceiveSuccessCount(v int64) *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat {
	s.SmsReceiveSuccessCount = &v
	return s
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) SetSmsSentCount(v int64) *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat {
	s.SmsSentCount = &v
	return s
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) SetSmsSkipCount(v int64) *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat {
	s.SmsSkipCount = &v
	return s
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) SetTime(v string) *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat {
	s.Time = &v
	return s
}

type QueryPushStatByAppResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryPushStatByAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPushStatByAppResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPushStatByAppResponse) GoString() string {
	return s.String()
}

func (s *QueryPushStatByAppResponse) SetHeaders(v map[string]*string) *QueryPushStatByAppResponse {
	s.Headers = v
	return s
}

func (s *QueryPushStatByAppResponse) SetStatusCode(v int32) *QueryPushStatByAppResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPushStatByAppResponse) SetBody(v *QueryPushStatByAppResponseBody) *QueryPushStatByAppResponse {
	s.Body = v
	return s
}

type QueryPushStatByMsgRequest struct {
	AppKey    *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	MessageId *int64 `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s QueryPushStatByMsgRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPushStatByMsgRequest) GoString() string {
	return s.String()
}

func (s *QueryPushStatByMsgRequest) SetAppKey(v int64) *QueryPushStatByMsgRequest {
	s.AppKey = &v
	return s
}

func (s *QueryPushStatByMsgRequest) SetMessageId(v int64) *QueryPushStatByMsgRequest {
	s.MessageId = &v
	return s
}

type QueryPushStatByMsgResponseBody struct {
	PushStats *QueryPushStatByMsgResponseBodyPushStats `json:"PushStats,omitempty" xml:"PushStats,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryPushStatByMsgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPushStatByMsgResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPushStatByMsgResponseBody) SetPushStats(v *QueryPushStatByMsgResponseBodyPushStats) *QueryPushStatByMsgResponseBody {
	s.PushStats = v
	return s
}

func (s *QueryPushStatByMsgResponseBody) SetRequestId(v string) *QueryPushStatByMsgResponseBody {
	s.RequestId = &v
	return s
}

type QueryPushStatByMsgResponseBodyPushStats struct {
	PushStat []*QueryPushStatByMsgResponseBodyPushStatsPushStat `json:"PushStat,omitempty" xml:"PushStat,omitempty" type:"Repeated"`
}

func (s QueryPushStatByMsgResponseBodyPushStats) String() string {
	return tea.Prettify(s)
}

func (s QueryPushStatByMsgResponseBodyPushStats) GoString() string {
	return s.String()
}

func (s *QueryPushStatByMsgResponseBodyPushStats) SetPushStat(v []*QueryPushStatByMsgResponseBodyPushStatsPushStat) *QueryPushStatByMsgResponseBodyPushStats {
	s.PushStat = v
	return s
}

type QueryPushStatByMsgResponseBodyPushStatsPushStat struct {
	AcceptCount            *int64  `json:"AcceptCount,omitempty" xml:"AcceptCount,omitempty"`
	DeletedCount           *int64  `json:"DeletedCount,omitempty" xml:"DeletedCount,omitempty"`
	MessageId              *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	OpenedCount            *int64  `json:"OpenedCount,omitempty" xml:"OpenedCount,omitempty"`
	ReceivedCount          *int64  `json:"ReceivedCount,omitempty" xml:"ReceivedCount,omitempty"`
	SentCount              *int64  `json:"SentCount,omitempty" xml:"SentCount,omitempty"`
	SmsFailedCount         *int64  `json:"SmsFailedCount,omitempty" xml:"SmsFailedCount,omitempty"`
	SmsReceiveFailedCount  *int64  `json:"SmsReceiveFailedCount,omitempty" xml:"SmsReceiveFailedCount,omitempty"`
	SmsReceiveSuccessCount *int64  `json:"SmsReceiveSuccessCount,omitempty" xml:"SmsReceiveSuccessCount,omitempty"`
	SmsSentCount           *int64  `json:"SmsSentCount,omitempty" xml:"SmsSentCount,omitempty"`
	SmsSkipCount           *int64  `json:"SmsSkipCount,omitempty" xml:"SmsSkipCount,omitempty"`
}

func (s QueryPushStatByMsgResponseBodyPushStatsPushStat) String() string {
	return tea.Prettify(s)
}

func (s QueryPushStatByMsgResponseBodyPushStatsPushStat) GoString() string {
	return s.String()
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) SetAcceptCount(v int64) *QueryPushStatByMsgResponseBodyPushStatsPushStat {
	s.AcceptCount = &v
	return s
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) SetDeletedCount(v int64) *QueryPushStatByMsgResponseBodyPushStatsPushStat {
	s.DeletedCount = &v
	return s
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) SetMessageId(v string) *QueryPushStatByMsgResponseBodyPushStatsPushStat {
	s.MessageId = &v
	return s
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) SetOpenedCount(v int64) *QueryPushStatByMsgResponseBodyPushStatsPushStat {
	s.OpenedCount = &v
	return s
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) SetReceivedCount(v int64) *QueryPushStatByMsgResponseBodyPushStatsPushStat {
	s.ReceivedCount = &v
	return s
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) SetSentCount(v int64) *QueryPushStatByMsgResponseBodyPushStatsPushStat {
	s.SentCount = &v
	return s
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) SetSmsFailedCount(v int64) *QueryPushStatByMsgResponseBodyPushStatsPushStat {
	s.SmsFailedCount = &v
	return s
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) SetSmsReceiveFailedCount(v int64) *QueryPushStatByMsgResponseBodyPushStatsPushStat {
	s.SmsReceiveFailedCount = &v
	return s
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) SetSmsReceiveSuccessCount(v int64) *QueryPushStatByMsgResponseBodyPushStatsPushStat {
	s.SmsReceiveSuccessCount = &v
	return s
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) SetSmsSentCount(v int64) *QueryPushStatByMsgResponseBodyPushStatsPushStat {
	s.SmsSentCount = &v
	return s
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) SetSmsSkipCount(v int64) *QueryPushStatByMsgResponseBodyPushStatsPushStat {
	s.SmsSkipCount = &v
	return s
}

type QueryPushStatByMsgResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryPushStatByMsgResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPushStatByMsgResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPushStatByMsgResponse) GoString() string {
	return s.String()
}

func (s *QueryPushStatByMsgResponse) SetHeaders(v map[string]*string) *QueryPushStatByMsgResponse {
	s.Headers = v
	return s
}

func (s *QueryPushStatByMsgResponse) SetStatusCode(v int32) *QueryPushStatByMsgResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPushStatByMsgResponse) SetBody(v *QueryPushStatByMsgResponseBody) *QueryPushStatByMsgResponse {
	s.Body = v
	return s
}

type QueryTagsRequest struct {
	AppKey    *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	ClientKey *string `json:"ClientKey,omitempty" xml:"ClientKey,omitempty"`
	KeyType   *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
}

func (s QueryTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTagsRequest) GoString() string {
	return s.String()
}

func (s *QueryTagsRequest) SetAppKey(v int64) *QueryTagsRequest {
	s.AppKey = &v
	return s
}

func (s *QueryTagsRequest) SetClientKey(v string) *QueryTagsRequest {
	s.ClientKey = &v
	return s
}

func (s *QueryTagsRequest) SetKeyType(v string) *QueryTagsRequest {
	s.KeyType = &v
	return s
}

type QueryTagsResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagInfos  *QueryTagsResponseBodyTagInfos `json:"TagInfos,omitempty" xml:"TagInfos,omitempty" type:"Struct"`
}

func (s QueryTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTagsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTagsResponseBody) SetRequestId(v string) *QueryTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTagsResponseBody) SetTagInfos(v *QueryTagsResponseBodyTagInfos) *QueryTagsResponseBody {
	s.TagInfos = v
	return s
}

type QueryTagsResponseBodyTagInfos struct {
	TagInfo []*QueryTagsResponseBodyTagInfosTagInfo `json:"TagInfo,omitempty" xml:"TagInfo,omitempty" type:"Repeated"`
}

func (s QueryTagsResponseBodyTagInfos) String() string {
	return tea.Prettify(s)
}

func (s QueryTagsResponseBodyTagInfos) GoString() string {
	return s.String()
}

func (s *QueryTagsResponseBodyTagInfos) SetTagInfo(v []*QueryTagsResponseBodyTagInfosTagInfo) *QueryTagsResponseBodyTagInfos {
	s.TagInfo = v
	return s
}

type QueryTagsResponseBodyTagInfosTagInfo struct {
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s QueryTagsResponseBodyTagInfosTagInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryTagsResponseBodyTagInfosTagInfo) GoString() string {
	return s.String()
}

func (s *QueryTagsResponseBodyTagInfosTagInfo) SetTagName(v string) *QueryTagsResponseBodyTagInfosTagInfo {
	s.TagName = &v
	return s
}

type QueryTagsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTagsResponse) GoString() string {
	return s.String()
}

func (s *QueryTagsResponse) SetHeaders(v map[string]*string) *QueryTagsResponse {
	s.Headers = v
	return s
}

func (s *QueryTagsResponse) SetStatusCode(v int32) *QueryTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTagsResponse) SetBody(v *QueryTagsResponseBody) *QueryTagsResponse {
	s.Body = v
	return s
}

type QueryUniqueDeviceStatRequest struct {
	AppKey      *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryUniqueDeviceStatRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUniqueDeviceStatRequest) GoString() string {
	return s.String()
}

func (s *QueryUniqueDeviceStatRequest) SetAppKey(v int64) *QueryUniqueDeviceStatRequest {
	s.AppKey = &v
	return s
}

func (s *QueryUniqueDeviceStatRequest) SetEndTime(v string) *QueryUniqueDeviceStatRequest {
	s.EndTime = &v
	return s
}

func (s *QueryUniqueDeviceStatRequest) SetGranularity(v string) *QueryUniqueDeviceStatRequest {
	s.Granularity = &v
	return s
}

func (s *QueryUniqueDeviceStatRequest) SetStartTime(v string) *QueryUniqueDeviceStatRequest {
	s.StartTime = &v
	return s
}

type QueryUniqueDeviceStatResponseBody struct {
	AppDeviceStats *QueryUniqueDeviceStatResponseBodyAppDeviceStats `json:"AppDeviceStats,omitempty" xml:"AppDeviceStats,omitempty" type:"Struct"`
	RequestId      *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryUniqueDeviceStatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUniqueDeviceStatResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUniqueDeviceStatResponseBody) SetAppDeviceStats(v *QueryUniqueDeviceStatResponseBodyAppDeviceStats) *QueryUniqueDeviceStatResponseBody {
	s.AppDeviceStats = v
	return s
}

func (s *QueryUniqueDeviceStatResponseBody) SetRequestId(v string) *QueryUniqueDeviceStatResponseBody {
	s.RequestId = &v
	return s
}

type QueryUniqueDeviceStatResponseBodyAppDeviceStats struct {
	AppDeviceStat []*QueryUniqueDeviceStatResponseBodyAppDeviceStatsAppDeviceStat `json:"AppDeviceStat,omitempty" xml:"AppDeviceStat,omitempty" type:"Repeated"`
}

func (s QueryUniqueDeviceStatResponseBodyAppDeviceStats) String() string {
	return tea.Prettify(s)
}

func (s QueryUniqueDeviceStatResponseBodyAppDeviceStats) GoString() string {
	return s.String()
}

func (s *QueryUniqueDeviceStatResponseBodyAppDeviceStats) SetAppDeviceStat(v []*QueryUniqueDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) *QueryUniqueDeviceStatResponseBodyAppDeviceStats {
	s.AppDeviceStat = v
	return s
}

type QueryUniqueDeviceStatResponseBodyAppDeviceStatsAppDeviceStat struct {
	Count *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	Time  *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s QueryUniqueDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) String() string {
	return tea.Prettify(s)
}

func (s QueryUniqueDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) GoString() string {
	return s.String()
}

func (s *QueryUniqueDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) SetCount(v int64) *QueryUniqueDeviceStatResponseBodyAppDeviceStatsAppDeviceStat {
	s.Count = &v
	return s
}

func (s *QueryUniqueDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) SetTime(v string) *QueryUniqueDeviceStatResponseBodyAppDeviceStatsAppDeviceStat {
	s.Time = &v
	return s
}

type QueryUniqueDeviceStatResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryUniqueDeviceStatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUniqueDeviceStatResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUniqueDeviceStatResponse) GoString() string {
	return s.String()
}

func (s *QueryUniqueDeviceStatResponse) SetHeaders(v map[string]*string) *QueryUniqueDeviceStatResponse {
	s.Headers = v
	return s
}

func (s *QueryUniqueDeviceStatResponse) SetStatusCode(v int32) *QueryUniqueDeviceStatResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUniqueDeviceStatResponse) SetBody(v *QueryUniqueDeviceStatResponseBody) *QueryUniqueDeviceStatResponse {
	s.Body = v
	return s
}

type RemoveTagRequest struct {
	AppKey  *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s RemoveTagRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveTagRequest) GoString() string {
	return s.String()
}

func (s *RemoveTagRequest) SetAppKey(v int64) *RemoveTagRequest {
	s.AppKey = &v
	return s
}

func (s *RemoveTagRequest) SetTagName(v string) *RemoveTagRequest {
	s.TagName = &v
	return s
}

type RemoveTagResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveTagResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveTagResponseBody) SetRequestId(v string) *RemoveTagResponseBody {
	s.RequestId = &v
	return s
}

type RemoveTagResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveTagResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveTagResponse) GoString() string {
	return s.String()
}

func (s *RemoveTagResponse) SetHeaders(v map[string]*string) *RemoveTagResponse {
	s.Headers = v
	return s
}

func (s *RemoveTagResponse) SetStatusCode(v int32) *RemoveTagResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveTagResponse) SetBody(v *RemoveTagResponseBody) *RemoveTagResponse {
	s.Body = v
	return s
}

type UnbindAliasRequest struct {
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	AppKey    *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	DeviceId  *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	UnbindAll *bool   `json:"UnbindAll,omitempty" xml:"UnbindAll,omitempty"`
}

func (s UnbindAliasRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindAliasRequest) GoString() string {
	return s.String()
}

func (s *UnbindAliasRequest) SetAliasName(v string) *UnbindAliasRequest {
	s.AliasName = &v
	return s
}

func (s *UnbindAliasRequest) SetAppKey(v int64) *UnbindAliasRequest {
	s.AppKey = &v
	return s
}

func (s *UnbindAliasRequest) SetDeviceId(v string) *UnbindAliasRequest {
	s.DeviceId = &v
	return s
}

func (s *UnbindAliasRequest) SetUnbindAll(v bool) *UnbindAliasRequest {
	s.UnbindAll = &v
	return s
}

type UnbindAliasResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindAliasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindAliasResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindAliasResponseBody) SetRequestId(v string) *UnbindAliasResponseBody {
	s.RequestId = &v
	return s
}

type UnbindAliasResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnbindAliasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindAliasResponse) GoString() string {
	return s.String()
}

func (s *UnbindAliasResponse) SetHeaders(v map[string]*string) *UnbindAliasResponse {
	s.Headers = v
	return s
}

func (s *UnbindAliasResponse) SetStatusCode(v int32) *UnbindAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindAliasResponse) SetBody(v *UnbindAliasResponseBody) *UnbindAliasResponse {
	s.Body = v
	return s
}

type UnbindPhoneRequest struct {
	AppKey   *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s UnbindPhoneRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindPhoneRequest) GoString() string {
	return s.String()
}

func (s *UnbindPhoneRequest) SetAppKey(v int64) *UnbindPhoneRequest {
	s.AppKey = &v
	return s
}

func (s *UnbindPhoneRequest) SetDeviceId(v string) *UnbindPhoneRequest {
	s.DeviceId = &v
	return s
}

type UnbindPhoneResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindPhoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindPhoneResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindPhoneResponseBody) SetRequestId(v string) *UnbindPhoneResponseBody {
	s.RequestId = &v
	return s
}

type UnbindPhoneResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnbindPhoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindPhoneResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindPhoneResponse) GoString() string {
	return s.String()
}

func (s *UnbindPhoneResponse) SetHeaders(v map[string]*string) *UnbindPhoneResponse {
	s.Headers = v
	return s
}

func (s *UnbindPhoneResponse) SetStatusCode(v int32) *UnbindPhoneResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindPhoneResponse) SetBody(v *UnbindPhoneResponseBody) *UnbindPhoneResponse {
	s.Body = v
	return s
}

type UnbindTagRequest struct {
	AppKey    *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	ClientKey *string `json:"ClientKey,omitempty" xml:"ClientKey,omitempty"`
	KeyType   *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	TagName   *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s UnbindTagRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindTagRequest) GoString() string {
	return s.String()
}

func (s *UnbindTagRequest) SetAppKey(v int64) *UnbindTagRequest {
	s.AppKey = &v
	return s
}

func (s *UnbindTagRequest) SetClientKey(v string) *UnbindTagRequest {
	s.ClientKey = &v
	return s
}

func (s *UnbindTagRequest) SetKeyType(v string) *UnbindTagRequest {
	s.KeyType = &v
	return s
}

func (s *UnbindTagRequest) SetTagName(v string) *UnbindTagRequest {
	s.TagName = &v
	return s
}

type UnbindTagResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindTagResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindTagResponseBody) SetRequestId(v string) *UnbindTagResponseBody {
	s.RequestId = &v
	return s
}

type UnbindTagResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnbindTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindTagResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindTagResponse) GoString() string {
	return s.String()
}

func (s *UnbindTagResponse) SetHeaders(v map[string]*string) *UnbindTagResponse {
	s.Headers = v
	return s
}

func (s *UnbindTagResponse) SetStatusCode(v int32) *UnbindTagResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindTagResponse) SetBody(v *UnbindTagResponseBody) *UnbindTagResponse {
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
		"ap-northeast-1":              tea.String("cloudpush.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("cloudpush.aliyuncs.com"),
		"ap-south-1":                  tea.String("cloudpush.aliyuncs.com"),
		"ap-southeast-1":              tea.String("cloudpush.aliyuncs.com"),
		"ap-southeast-2":              tea.String("cloudpush.aliyuncs.com"),
		"ap-southeast-3":              tea.String("cloudpush.aliyuncs.com"),
		"ap-southeast-5":              tea.String("cloudpush.aliyuncs.com"),
		"cn-beijing":                  tea.String("cloudpush.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("cloudpush.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("cloudpush.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("cloudpush.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("cloudpush.aliyuncs.com"),
		"cn-chengdu":                  tea.String("cloudpush.aliyuncs.com"),
		"cn-edge-1":                   tea.String("cloudpush.aliyuncs.com"),
		"cn-fujian":                   tea.String("cloudpush.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("cloudpush.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("cloudpush.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("cloudpush.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("cloudpush.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("cloudpush.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("cloudpush.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("cloudpush.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("cloudpush.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("cloudpush.aliyuncs.com"),
		"cn-hongkong":                 tea.String("cloudpush.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("cloudpush.aliyuncs.com"),
		"cn-huhehaote":                tea.String("cloudpush.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("cloudpush.aliyuncs.com"),
		"cn-qingdao":                  tea.String("cloudpush.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("cloudpush.aliyuncs.com"),
		"cn-shanghai":                 tea.String("cloudpush.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("cloudpush.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("cloudpush.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("cloudpush.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("cloudpush.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("cloudpush.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("cloudpush.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("cloudpush.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("cloudpush.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("cloudpush.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("cloudpush.aliyuncs.com"),
		"cn-wuhan":                    tea.String("cloudpush.aliyuncs.com"),
		"cn-yushanfang":               tea.String("cloudpush.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("cloudpush.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("cloudpush.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("cloudpush.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("cloudpush.aliyuncs.com"),
		"eu-central-1":                tea.String("cloudpush.aliyuncs.com"),
		"eu-west-1":                   tea.String("cloudpush.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("cloudpush.aliyuncs.com"),
		"me-east-1":                   tea.String("cloudpush.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("cloudpush.aliyuncs.com"),
		"us-east-1":                   tea.String("cloudpush.aliyuncs.com"),
		"us-west-1":                   tea.String("cloudpush.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("push"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) BindAliasWithOptions(request *BindAliasRequest, runtime *util.RuntimeOptions) (_result *BindAliasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliasName)) {
		query["AliasName"] = request.AliasName
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindAlias"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindAlias(request *BindAliasRequest) (_result *BindAliasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindAliasResponse{}
	_body, _err := client.BindAliasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindPhoneWithOptions(request *BindPhoneRequest, runtime *util.RuntimeOptions) (_result *BindPhoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindPhone"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindPhoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindPhone(request *BindPhoneRequest) (_result *BindPhoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindPhoneResponse{}
	_body, _err := client.BindPhoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindTagWithOptions(request *BindTagRequest, runtime *util.RuntimeOptions) (_result *BindTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.ClientKey)) {
		query["ClientKey"] = request.ClientKey
	}

	if !tea.BoolValue(util.IsUnset(request.KeyType)) {
		query["KeyType"] = request.KeyType
	}

	if !tea.BoolValue(util.IsUnset(request.TagName)) {
		query["TagName"] = request.TagName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindTag"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindTag(request *BindTagRequest) (_result *BindTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindTagResponse{}
	_body, _err := client.BindTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelPushWithOptions(request *CancelPushRequest, runtime *util.RuntimeOptions) (_result *CancelPushResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.MessageId)) {
		query["MessageId"] = request.MessageId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelPush"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelPushResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelPush(request *CancelPushRequest) (_result *CancelPushResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelPushResponse{}
	_body, _err := client.CancelPushWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckCertificateWithOptions(request *CheckCertificateRequest, runtime *util.RuntimeOptions) (_result *CheckCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckCertificate"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckCertificate(request *CheckCertificateRequest) (_result *CheckCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckCertificateResponse{}
	_body, _err := client.CheckCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckDeviceWithOptions(request *CheckDeviceRequest, runtime *util.RuntimeOptions) (_result *CheckDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckDevice"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckDevice(request *CheckDeviceRequest) (_result *CheckDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckDeviceResponse{}
	_body, _err := client.CheckDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckDevicesWithOptions(request *CheckDevicesRequest, runtime *util.RuntimeOptions) (_result *CheckDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceIds)) {
		query["DeviceIds"] = request.DeviceIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckDevices"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckDevices(request *CheckDevicesRequest) (_result *CheckDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckDevicesResponse{}
	_body, _err := client.CheckDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CompleteContinuouslyPushWithOptions(request *CompleteContinuouslyPushRequest, runtime *util.RuntimeOptions) (_result *CompleteContinuouslyPushResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.MessageId)) {
		query["MessageId"] = request.MessageId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CompleteContinuouslyPush"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CompleteContinuouslyPushResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CompleteContinuouslyPush(request *CompleteContinuouslyPushRequest) (_result *CompleteContinuouslyPushResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CompleteContinuouslyPushResponse{}
	_body, _err := client.CompleteContinuouslyPushWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ContinuouslyPushWithOptions(request *ContinuouslyPushRequest, runtime *util.RuntimeOptions) (_result *ContinuouslyPushResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.MessageId)) {
		query["MessageId"] = request.MessageId
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		query["Target"] = request.Target
	}

	if !tea.BoolValue(util.IsUnset(request.TargetValue)) {
		query["TargetValue"] = request.TargetValue
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ContinuouslyPush"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ContinuouslyPushResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ContinuouslyPush(request *ContinuouslyPushRequest) (_result *ContinuouslyPushResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ContinuouslyPushResponse{}
	_body, _err := client.ContinuouslyPushWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSummaryAppsWithOptions(runtime *util.RuntimeOptions) (_result *ListSummaryAppsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListSummaryApps"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSummaryAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSummaryApps() (_result *ListSummaryAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSummaryAppsResponse{}
	_body, _err := client.ListSummaryAppsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagsWithOptions(request *ListTagsRequest, runtime *util.RuntimeOptions) (_result *ListTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTags"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTags(request *ListTagsRequest) (_result *ListTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagsResponse{}
	_body, _err := client.ListTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MassPushWithOptions(request *MassPushRequest, runtime *util.RuntimeOptions) (_result *MassPushResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PushTask)) {
		body["PushTask"] = request.PushTask
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MassPush"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MassPushResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MassPush(request *MassPushRequest) (_result *MassPushResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MassPushResponse{}
	_body, _err := client.MassPushWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushWithOptions(request *PushRequest, runtime *util.RuntimeOptions) (_result *PushResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AndroidActivity)) {
		query["AndroidActivity"] = request.AndroidActivity
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidBigBody)) {
		query["AndroidBigBody"] = request.AndroidBigBody
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidBigPictureUrl)) {
		query["AndroidBigPictureUrl"] = request.AndroidBigPictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidBigTitle)) {
		query["AndroidBigTitle"] = request.AndroidBigTitle
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidExtParameters)) {
		query["AndroidExtParameters"] = request.AndroidExtParameters
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidImageUrl)) {
		query["AndroidImageUrl"] = request.AndroidImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidInboxBody)) {
		query["AndroidInboxBody"] = request.AndroidInboxBody
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidMessageHuaweiCategory)) {
		query["AndroidMessageHuaweiCategory"] = request.AndroidMessageHuaweiCategory
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidMessageHuaweiUrgency)) {
		query["AndroidMessageHuaweiUrgency"] = request.AndroidMessageHuaweiUrgency
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidMusic)) {
		query["AndroidMusic"] = request.AndroidMusic
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidNotificationBarPriority)) {
		query["AndroidNotificationBarPriority"] = request.AndroidNotificationBarPriority
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidNotificationBarType)) {
		query["AndroidNotificationBarType"] = request.AndroidNotificationBarType
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidNotificationChannel)) {
		query["AndroidNotificationChannel"] = request.AndroidNotificationChannel
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidNotificationHuaweiChannel)) {
		query["AndroidNotificationHuaweiChannel"] = request.AndroidNotificationHuaweiChannel
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidNotificationNotifyId)) {
		query["AndroidNotificationNotifyId"] = request.AndroidNotificationNotifyId
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidNotificationVivoChannel)) {
		query["AndroidNotificationVivoChannel"] = request.AndroidNotificationVivoChannel
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidNotificationXiaomiChannel)) {
		query["AndroidNotificationXiaomiChannel"] = request.AndroidNotificationXiaomiChannel
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidNotifyType)) {
		query["AndroidNotifyType"] = request.AndroidNotifyType
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidOpenType)) {
		query["AndroidOpenType"] = request.AndroidOpenType
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidOpenUrl)) {
		query["AndroidOpenUrl"] = request.AndroidOpenUrl
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidPopupActivity)) {
		query["AndroidPopupActivity"] = request.AndroidPopupActivity
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidPopupBody)) {
		query["AndroidPopupBody"] = request.AndroidPopupBody
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidPopupTitle)) {
		query["AndroidPopupTitle"] = request.AndroidPopupTitle
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidRemind)) {
		query["AndroidRemind"] = request.AndroidRemind
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidRenderStyle)) {
		query["AndroidRenderStyle"] = request.AndroidRenderStyle
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidVivoPushMode)) {
		query["AndroidVivoPushMode"] = request.AndroidVivoPushMode
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidXiaoMiActivity)) {
		query["AndroidXiaoMiActivity"] = request.AndroidXiaoMiActivity
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidXiaoMiNotifyBody)) {
		query["AndroidXiaoMiNotifyBody"] = request.AndroidXiaoMiNotifyBody
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidXiaoMiNotifyTitle)) {
		query["AndroidXiaoMiNotifyTitle"] = request.AndroidXiaoMiNotifyTitle
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidXiaomiBigPictureUrl)) {
		query["AndroidXiaomiBigPictureUrl"] = request.AndroidXiaomiBigPictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidXiaomiImageUrl)) {
		query["AndroidXiaomiImageUrl"] = request.AndroidXiaomiImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.Body)) {
		query["Body"] = request.Body
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceType)) {
		query["DeviceType"] = request.DeviceType
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireTime)) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !tea.BoolValue(util.IsUnset(request.JobKey)) {
		query["JobKey"] = request.JobKey
	}

	if !tea.BoolValue(util.IsUnset(request.PushTime)) {
		query["PushTime"] = request.PushTime
	}

	if !tea.BoolValue(util.IsUnset(request.PushType)) {
		query["PushType"] = request.PushType
	}

	if !tea.BoolValue(util.IsUnset(request.SendChannels)) {
		query["SendChannels"] = request.SendChannels
	}

	if !tea.BoolValue(util.IsUnset(request.SendSpeed)) {
		query["SendSpeed"] = request.SendSpeed
	}

	if !tea.BoolValue(util.IsUnset(request.SmsDelaySecs)) {
		query["SmsDelaySecs"] = request.SmsDelaySecs
	}

	if !tea.BoolValue(util.IsUnset(request.SmsParams)) {
		query["SmsParams"] = request.SmsParams
	}

	if !tea.BoolValue(util.IsUnset(request.SmsSendPolicy)) {
		query["SmsSendPolicy"] = request.SmsSendPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.SmsSignName)) {
		query["SmsSignName"] = request.SmsSignName
	}

	if !tea.BoolValue(util.IsUnset(request.SmsTemplateName)) {
		query["SmsTemplateName"] = request.SmsTemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.StoreOffline)) {
		query["StoreOffline"] = request.StoreOffline
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		query["Target"] = request.Target
	}

	if !tea.BoolValue(util.IsUnset(request.TargetValue)) {
		query["TargetValue"] = request.TargetValue
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Trim)) {
		query["Trim"] = request.Trim
	}

	if !tea.BoolValue(util.IsUnset(request.IOSApnsEnv)) {
		query["iOSApnsEnv"] = request.IOSApnsEnv
	}

	if !tea.BoolValue(util.IsUnset(request.IOSBadge)) {
		query["iOSBadge"] = request.IOSBadge
	}

	if !tea.BoolValue(util.IsUnset(request.IOSBadgeAutoIncrement)) {
		query["iOSBadgeAutoIncrement"] = request.IOSBadgeAutoIncrement
	}

	if !tea.BoolValue(util.IsUnset(request.IOSExtParameters)) {
		query["iOSExtParameters"] = request.IOSExtParameters
	}

	if !tea.BoolValue(util.IsUnset(request.IOSInterruptionLevel)) {
		query["iOSInterruptionLevel"] = request.IOSInterruptionLevel
	}

	if !tea.BoolValue(util.IsUnset(request.IOSMusic)) {
		query["iOSMusic"] = request.IOSMusic
	}

	if !tea.BoolValue(util.IsUnset(request.IOSMutableContent)) {
		query["iOSMutableContent"] = request.IOSMutableContent
	}

	if !tea.BoolValue(util.IsUnset(request.IOSNotificationCategory)) {
		query["iOSNotificationCategory"] = request.IOSNotificationCategory
	}

	if !tea.BoolValue(util.IsUnset(request.IOSNotificationCollapseId)) {
		query["iOSNotificationCollapseId"] = request.IOSNotificationCollapseId
	}

	if !tea.BoolValue(util.IsUnset(request.IOSNotificationThreadId)) {
		query["iOSNotificationThreadId"] = request.IOSNotificationThreadId
	}

	if !tea.BoolValue(util.IsUnset(request.IOSRelevanceScore)) {
		query["iOSRelevanceScore"] = request.IOSRelevanceScore
	}

	if !tea.BoolValue(util.IsUnset(request.IOSRemind)) {
		query["iOSRemind"] = request.IOSRemind
	}

	if !tea.BoolValue(util.IsUnset(request.IOSRemindBody)) {
		query["iOSRemindBody"] = request.IOSRemindBody
	}

	if !tea.BoolValue(util.IsUnset(request.IOSSilentNotification)) {
		query["iOSSilentNotification"] = request.IOSSilentNotification
	}

	if !tea.BoolValue(util.IsUnset(request.IOSSubtitle)) {
		query["iOSSubtitle"] = request.IOSSubtitle
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Push"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PushResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Push(request *PushRequest) (_result *PushResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PushResponse{}
	_body, _err := client.PushWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushMessageToAndroidWithOptions(request *PushMessageToAndroidRequest, runtime *util.RuntimeOptions) (_result *PushMessageToAndroidResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.Body)) {
		query["Body"] = request.Body
	}

	if !tea.BoolValue(util.IsUnset(request.JobKey)) {
		query["JobKey"] = request.JobKey
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		query["Target"] = request.Target
	}

	if !tea.BoolValue(util.IsUnset(request.TargetValue)) {
		query["TargetValue"] = request.TargetValue
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PushMessageToAndroid"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PushMessageToAndroidResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushMessageToAndroid(request *PushMessageToAndroidRequest) (_result *PushMessageToAndroidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PushMessageToAndroidResponse{}
	_body, _err := client.PushMessageToAndroidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushMessageToiOSWithOptions(request *PushMessageToiOSRequest, runtime *util.RuntimeOptions) (_result *PushMessageToiOSResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.Body)) {
		query["Body"] = request.Body
	}

	if !tea.BoolValue(util.IsUnset(request.JobKey)) {
		query["JobKey"] = request.JobKey
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		query["Target"] = request.Target
	}

	if !tea.BoolValue(util.IsUnset(request.TargetValue)) {
		query["TargetValue"] = request.TargetValue
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PushMessageToiOS"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PushMessageToiOSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushMessageToiOS(request *PushMessageToiOSRequest) (_result *PushMessageToiOSResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PushMessageToiOSResponse{}
	_body, _err := client.PushMessageToiOSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushNoticeToAndroidWithOptions(request *PushNoticeToAndroidRequest, runtime *util.RuntimeOptions) (_result *PushNoticeToAndroidResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.Body)) {
		query["Body"] = request.Body
	}

	if !tea.BoolValue(util.IsUnset(request.ExtParameters)) {
		query["ExtParameters"] = request.ExtParameters
	}

	if !tea.BoolValue(util.IsUnset(request.JobKey)) {
		query["JobKey"] = request.JobKey
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		query["Target"] = request.Target
	}

	if !tea.BoolValue(util.IsUnset(request.TargetValue)) {
		query["TargetValue"] = request.TargetValue
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PushNoticeToAndroid"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PushNoticeToAndroidResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushNoticeToAndroid(request *PushNoticeToAndroidRequest) (_result *PushNoticeToAndroidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PushNoticeToAndroidResponse{}
	_body, _err := client.PushNoticeToAndroidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushNoticeToiOSWithOptions(request *PushNoticeToiOSRequest, runtime *util.RuntimeOptions) (_result *PushNoticeToiOSResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApnsEnv)) {
		query["ApnsEnv"] = request.ApnsEnv
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.Body)) {
		query["Body"] = request.Body
	}

	if !tea.BoolValue(util.IsUnset(request.ExtParameters)) {
		query["ExtParameters"] = request.ExtParameters
	}

	if !tea.BoolValue(util.IsUnset(request.JobKey)) {
		query["JobKey"] = request.JobKey
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		query["Target"] = request.Target
	}

	if !tea.BoolValue(util.IsUnset(request.TargetValue)) {
		query["TargetValue"] = request.TargetValue
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PushNoticeToiOS"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PushNoticeToiOSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushNoticeToiOS(request *PushNoticeToiOSRequest) (_result *PushNoticeToiOSResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PushNoticeToiOSResponse{}
	_body, _err := client.PushNoticeToiOSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAliasesWithOptions(request *QueryAliasesRequest, runtime *util.RuntimeOptions) (_result *QueryAliasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAliases"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAliasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAliases(request *QueryAliasesRequest) (_result *QueryAliasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAliasesResponse{}
	_body, _err := client.QueryAliasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDeviceInfoWithOptions(request *QueryDeviceInfoRequest, runtime *util.RuntimeOptions) (_result *QueryDeviceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDeviceInfo"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDeviceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDeviceInfo(request *QueryDeviceInfoRequest) (_result *QueryDeviceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDeviceInfoResponse{}
	_body, _err := client.QueryDeviceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDeviceStatWithOptions(request *QueryDeviceStatRequest, runtime *util.RuntimeOptions) (_result *QueryDeviceStatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceType)) {
		query["DeviceType"] = request.DeviceType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.QueryType)) {
		query["QueryType"] = request.QueryType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDeviceStat"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDeviceStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDeviceStat(request *QueryDeviceStatRequest) (_result *QueryDeviceStatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDeviceStatResponse{}
	_body, _err := client.QueryDeviceStatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDevicesByAccountWithOptions(request *QueryDevicesByAccountRequest, runtime *util.RuntimeOptions) (_result *QueryDevicesByAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Account)) {
		query["Account"] = request.Account
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDevicesByAccount"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDevicesByAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDevicesByAccount(request *QueryDevicesByAccountRequest) (_result *QueryDevicesByAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDevicesByAccountResponse{}
	_body, _err := client.QueryDevicesByAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDevicesByAliasWithOptions(request *QueryDevicesByAliasRequest, runtime *util.RuntimeOptions) (_result *QueryDevicesByAliasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Alias)) {
		query["Alias"] = request.Alias
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDevicesByAlias"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDevicesByAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDevicesByAlias(request *QueryDevicesByAliasRequest) (_result *QueryDevicesByAliasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDevicesByAliasResponse{}
	_body, _err := client.QueryDevicesByAliasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPushRecordsWithOptions(request *QueryPushRecordsRequest, runtime *util.RuntimeOptions) (_result *QueryPushRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PushType)) {
		query["PushType"] = request.PushType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		query["Target"] = request.Target
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPushRecords"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPushRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPushRecords(request *QueryPushRecordsRequest) (_result *QueryPushRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPushRecordsResponse{}
	_body, _err := client.QueryPushRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPushStatByAppWithOptions(request *QueryPushStatByAppRequest, runtime *util.RuntimeOptions) (_result *QueryPushStatByAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Granularity)) {
		query["Granularity"] = request.Granularity
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPushStatByApp"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPushStatByAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPushStatByApp(request *QueryPushStatByAppRequest) (_result *QueryPushStatByAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPushStatByAppResponse{}
	_body, _err := client.QueryPushStatByAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPushStatByMsgWithOptions(request *QueryPushStatByMsgRequest, runtime *util.RuntimeOptions) (_result *QueryPushStatByMsgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.MessageId)) {
		query["MessageId"] = request.MessageId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPushStatByMsg"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPushStatByMsgResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPushStatByMsg(request *QueryPushStatByMsgRequest) (_result *QueryPushStatByMsgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPushStatByMsgResponse{}
	_body, _err := client.QueryPushStatByMsgWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTagsWithOptions(request *QueryTagsRequest, runtime *util.RuntimeOptions) (_result *QueryTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.ClientKey)) {
		query["ClientKey"] = request.ClientKey
	}

	if !tea.BoolValue(util.IsUnset(request.KeyType)) {
		query["KeyType"] = request.KeyType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTags"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTags(request *QueryTagsRequest) (_result *QueryTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTagsResponse{}
	_body, _err := client.QueryTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUniqueDeviceStatWithOptions(request *QueryUniqueDeviceStatRequest, runtime *util.RuntimeOptions) (_result *QueryUniqueDeviceStatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Granularity)) {
		query["Granularity"] = request.Granularity
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryUniqueDeviceStat"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUniqueDeviceStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUniqueDeviceStat(request *QueryUniqueDeviceStatRequest) (_result *QueryUniqueDeviceStatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryUniqueDeviceStatResponse{}
	_body, _err := client.QueryUniqueDeviceStatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveTagWithOptions(request *RemoveTagRequest, runtime *util.RuntimeOptions) (_result *RemoveTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.TagName)) {
		query["TagName"] = request.TagName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveTag"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveTag(request *RemoveTagRequest) (_result *RemoveTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveTagResponse{}
	_body, _err := client.RemoveTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindAliasWithOptions(request *UnbindAliasRequest, runtime *util.RuntimeOptions) (_result *UnbindAliasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliasName)) {
		query["AliasName"] = request.AliasName
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.UnbindAll)) {
		query["UnbindAll"] = request.UnbindAll
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindAlias"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindAlias(request *UnbindAliasRequest) (_result *UnbindAliasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindAliasResponse{}
	_body, _err := client.UnbindAliasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindPhoneWithOptions(request *UnbindPhoneRequest, runtime *util.RuntimeOptions) (_result *UnbindPhoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindPhone"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindPhoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindPhone(request *UnbindPhoneRequest) (_result *UnbindPhoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindPhoneResponse{}
	_body, _err := client.UnbindPhoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindTagWithOptions(request *UnbindTagRequest, runtime *util.RuntimeOptions) (_result *UnbindTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.ClientKey)) {
		query["ClientKey"] = request.ClientKey
	}

	if !tea.BoolValue(util.IsUnset(request.KeyType)) {
		query["KeyType"] = request.KeyType
	}

	if !tea.BoolValue(util.IsUnset(request.TagName)) {
		query["TagName"] = request.TagName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindTag"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindTag(request *UnbindTagRequest) (_result *UnbindTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindTagResponse{}
	_body, _err := client.UnbindTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
