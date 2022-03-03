// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetCardDetailRequest struct {
	DestroyCard *bool   `json:"DestroyCard,omitempty" xml:"DestroyCard,omitempty"`
	Iccid       *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ShowPsim    *bool   `json:"ShowPsim,omitempty" xml:"ShowPsim,omitempty"`
}

func (s GetCardDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCardDetailRequest) GoString() string {
	return s.String()
}

func (s *GetCardDetailRequest) SetDestroyCard(v bool) *GetCardDetailRequest {
	s.DestroyCard = &v
	return s
}

func (s *GetCardDetailRequest) SetIccid(v string) *GetCardDetailRequest {
	s.Iccid = &v
	return s
}

func (s *GetCardDetailRequest) SetInstanceId(v string) *GetCardDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCardDetailRequest) SetShowPsim(v bool) *GetCardDetailRequest {
	s.ShowPsim = &v
	return s
}

type GetCardDetailResponseBody struct {
	Code             *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *GetCardDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage     *string                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string                        `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCardDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCardDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetCardDetailResponseBody) SetCode(v string) *GetCardDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetCardDetailResponseBody) SetData(v *GetCardDetailResponseBodyData) *GetCardDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetCardDetailResponseBody) SetErrorMessage(v string) *GetCardDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetCardDetailResponseBody) SetLocalizedMessage(v string) *GetCardDetailResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *GetCardDetailResponseBody) SetRequestId(v string) *GetCardDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCardDetailResponseBody) SetSuccess(v bool) *GetCardDetailResponseBody {
	s.Success = &v
	return s
}

type GetCardDetailResponseBodyData struct {
	ListPsimCards []*GetCardDetailResponseBodyDataListPsimCards `json:"ListPsimCards,omitempty" xml:"ListPsimCards,omitempty" type:"Repeated"`
	VsimCardInfo  *GetCardDetailResponseBodyDataVsimCardInfo    `json:"VsimCardInfo,omitempty" xml:"VsimCardInfo,omitempty" type:"Struct"`
}

func (s GetCardDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCardDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCardDetailResponseBodyData) SetListPsimCards(v []*GetCardDetailResponseBodyDataListPsimCards) *GetCardDetailResponseBodyData {
	s.ListPsimCards = v
	return s
}

func (s *GetCardDetailResponseBodyData) SetVsimCardInfo(v *GetCardDetailResponseBodyDataVsimCardInfo) *GetCardDetailResponseBodyData {
	s.VsimCardInfo = v
	return s
}

type GetCardDetailResponseBodyDataListPsimCards struct {
	ApnName               *string   `json:"ApnName,omitempty" xml:"ApnName,omitempty"`
	CertifyStatus         *string   `json:"CertifyStatus,omitempty" xml:"CertifyStatus,omitempty"`
	Iccid                 *string   `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	Imsi                  []*string `json:"Imsi,omitempty" xml:"Imsi,omitempty" type:"Repeated"`
	Msisdn                []*string `json:"Msisdn,omitempty" xml:"Msisdn,omitempty" type:"Repeated"`
	OsStatus              *string   `json:"OsStatus,omitempty" xml:"OsStatus,omitempty"`
	PeriodAddFlow         *string   `json:"PeriodAddFlow,omitempty" xml:"PeriodAddFlow,omitempty"`
	PeriodSmsUse          *string   `json:"PeriodSmsUse,omitempty" xml:"PeriodSmsUse,omitempty"`
	PrivateNetworkSegment *string   `json:"PrivateNetworkSegment,omitempty" xml:"PrivateNetworkSegment,omitempty"`
	Status                *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	Vendor                *string   `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s GetCardDetailResponseBodyDataListPsimCards) String() string {
	return tea.Prettify(s)
}

func (s GetCardDetailResponseBodyDataListPsimCards) GoString() string {
	return s.String()
}

func (s *GetCardDetailResponseBodyDataListPsimCards) SetApnName(v string) *GetCardDetailResponseBodyDataListPsimCards {
	s.ApnName = &v
	return s
}

func (s *GetCardDetailResponseBodyDataListPsimCards) SetCertifyStatus(v string) *GetCardDetailResponseBodyDataListPsimCards {
	s.CertifyStatus = &v
	return s
}

func (s *GetCardDetailResponseBodyDataListPsimCards) SetIccid(v string) *GetCardDetailResponseBodyDataListPsimCards {
	s.Iccid = &v
	return s
}

func (s *GetCardDetailResponseBodyDataListPsimCards) SetImsi(v []*string) *GetCardDetailResponseBodyDataListPsimCards {
	s.Imsi = v
	return s
}

func (s *GetCardDetailResponseBodyDataListPsimCards) SetMsisdn(v []*string) *GetCardDetailResponseBodyDataListPsimCards {
	s.Msisdn = v
	return s
}

func (s *GetCardDetailResponseBodyDataListPsimCards) SetOsStatus(v string) *GetCardDetailResponseBodyDataListPsimCards {
	s.OsStatus = &v
	return s
}

func (s *GetCardDetailResponseBodyDataListPsimCards) SetPeriodAddFlow(v string) *GetCardDetailResponseBodyDataListPsimCards {
	s.PeriodAddFlow = &v
	return s
}

func (s *GetCardDetailResponseBodyDataListPsimCards) SetPeriodSmsUse(v string) *GetCardDetailResponseBodyDataListPsimCards {
	s.PeriodSmsUse = &v
	return s
}

func (s *GetCardDetailResponseBodyDataListPsimCards) SetPrivateNetworkSegment(v string) *GetCardDetailResponseBodyDataListPsimCards {
	s.PrivateNetworkSegment = &v
	return s
}

func (s *GetCardDetailResponseBodyDataListPsimCards) SetStatus(v string) *GetCardDetailResponseBodyDataListPsimCards {
	s.Status = &v
	return s
}

func (s *GetCardDetailResponseBodyDataListPsimCards) SetVendor(v string) *GetCardDetailResponseBodyDataListPsimCards {
	s.Vendor = &v
	return s
}

type GetCardDetailResponseBodyDataVsimCardInfo struct {
	ActiveTime                    *string                                             `json:"ActiveTime,omitempty" xml:"ActiveTime,omitempty"`
	ActiveType                    *string                                             `json:"ActiveType,omitempty" xml:"ActiveType,omitempty"`
	AliFee                        *string                                             `json:"AliFee,omitempty" xml:"AliFee,omitempty"`
	AliyunOrderId                 *string                                             `json:"AliyunOrderId,omitempty" xml:"AliyunOrderId,omitempty"`
	ApnName                       *string                                             `json:"ApnName,omitempty" xml:"ApnName,omitempty"`
	AutoLimitResume               *bool                                               `json:"AutoLimitResume,omitempty" xml:"AutoLimitResume,omitempty"`
	AutoRebindReuse               *bool                                               `json:"AutoRebindReuse,omitempty" xml:"AutoRebindReuse,omitempty"`
	CardLimitSpeedThreshold       *int32                                              `json:"CardLimitSpeedThreshold,omitempty" xml:"CardLimitSpeedThreshold,omitempty"`
	CardLimitStopThreshold        *int32                                              `json:"CardLimitStopThreshold,omitempty" xml:"CardLimitStopThreshold,omitempty"`
	CertifyStatus                 *string                                             `json:"CertifyStatus,omitempty" xml:"CertifyStatus,omitempty"`
	CertifyType                   *string                                             `json:"CertifyType,omitempty" xml:"CertifyType,omitempty"`
	CredentialInstanceId          *string                                             `json:"CredentialInstanceId,omitempty" xml:"CredentialInstanceId,omitempty"`
	CredentialLimitSpeedThreshold *int32                                              `json:"CredentialLimitSpeedThreshold,omitempty" xml:"CredentialLimitSpeedThreshold,omitempty"`
	CredentialLimitStopThreshold  *int32                                              `json:"CredentialLimitStopThreshold,omitempty" xml:"CredentialLimitStopThreshold,omitempty"`
	CredentialNo                  *string                                             `json:"CredentialNo,omitempty" xml:"CredentialNo,omitempty"`
	CredentialType                *string                                             `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	DataLevel                     *string                                             `json:"DataLevel,omitempty" xml:"DataLevel,omitempty"`
	DataType                      *string                                             `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DeviceImei                    *string                                             `json:"DeviceImei,omitempty" xml:"DeviceImei,omitempty"`
	DirectionalGroupId            *string                                             `json:"DirectionalGroupId,omitempty" xml:"DirectionalGroupId,omitempty"`
	DirectionalGroupName          *string                                             `json:"DirectionalGroupName,omitempty" xml:"DirectionalGroupName,omitempty"`
	ExpireTime                    *string                                             `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	FlowThresholdUnit             *string                                             `json:"FlowThresholdUnit,omitempty" xml:"FlowThresholdUnit,omitempty"`
	Iccid                         *string                                             `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	Imsi                          []*string                                           `json:"Imsi,omitempty" xml:"Imsi,omitempty" type:"Repeated"`
	IsAutoRecharge                *bool                                               `json:"IsAutoRecharge,omitempty" xml:"IsAutoRecharge,omitempty"`
	Msisdn                        []*string                                           `json:"Msisdn,omitempty" xml:"Msisdn,omitempty" type:"Repeated"`
	NotifyId                      *string                                             `json:"NotifyId,omitempty" xml:"NotifyId,omitempty"`
	OpenAccountTime               *string                                             `json:"OpenAccountTime,omitempty" xml:"OpenAccountTime,omitempty"`
	OsStatus                      *string                                             `json:"OsStatus,omitempty" xml:"OsStatus,omitempty"`
	Period                        *string                                             `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodAddFlow                 *string                                             `json:"PeriodAddFlow,omitempty" xml:"PeriodAddFlow,omitempty"`
	PeriodRestFlow                *string                                             `json:"PeriodRestFlow,omitempty" xml:"PeriodRestFlow,omitempty"`
	PeriodSmsUse                  *string                                             `json:"PeriodSmsUse,omitempty" xml:"PeriodSmsUse,omitempty"`
	PrivateNetworkSegment         *string                                             `json:"PrivateNetworkSegment,omitempty" xml:"PrivateNetworkSegment,omitempty"`
	SimType                       *string                                             `json:"SimType,omitempty" xml:"SimType,omitempty"`
	Status                        *string                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	TagList                       []*GetCardDetailResponseBodyDataVsimCardInfoTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	Vendor                        *string                                             `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	VsimInstanceId                *int32                                              `json:"VsimInstanceId,omitempty" xml:"VsimInstanceId,omitempty"`
}

func (s GetCardDetailResponseBodyDataVsimCardInfo) String() string {
	return tea.Prettify(s)
}

func (s GetCardDetailResponseBodyDataVsimCardInfo) GoString() string {
	return s.String()
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetActiveTime(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.ActiveTime = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetActiveType(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.ActiveType = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetAliFee(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.AliFee = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetAliyunOrderId(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.AliyunOrderId = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetApnName(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.ApnName = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetAutoLimitResume(v bool) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.AutoLimitResume = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetAutoRebindReuse(v bool) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.AutoRebindReuse = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetCardLimitSpeedThreshold(v int32) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.CardLimitSpeedThreshold = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetCardLimitStopThreshold(v int32) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.CardLimitStopThreshold = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetCertifyStatus(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.CertifyStatus = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetCertifyType(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.CertifyType = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetCredentialInstanceId(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.CredentialInstanceId = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetCredentialLimitSpeedThreshold(v int32) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.CredentialLimitSpeedThreshold = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetCredentialLimitStopThreshold(v int32) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.CredentialLimitStopThreshold = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetCredentialNo(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.CredentialNo = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetCredentialType(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.CredentialType = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetDataLevel(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.DataLevel = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetDataType(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.DataType = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetDeviceImei(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.DeviceImei = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetDirectionalGroupId(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.DirectionalGroupId = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetDirectionalGroupName(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.DirectionalGroupName = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetExpireTime(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.ExpireTime = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetFlowThresholdUnit(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.FlowThresholdUnit = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetIccid(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.Iccid = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetImsi(v []*string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.Imsi = v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetIsAutoRecharge(v bool) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.IsAutoRecharge = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetMsisdn(v []*string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.Msisdn = v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetNotifyId(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.NotifyId = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetOpenAccountTime(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.OpenAccountTime = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetOsStatus(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.OsStatus = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetPeriod(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.Period = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetPeriodAddFlow(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.PeriodAddFlow = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetPeriodRestFlow(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.PeriodRestFlow = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetPeriodSmsUse(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.PeriodSmsUse = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetPrivateNetworkSegment(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.PrivateNetworkSegment = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetSimType(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.SimType = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetStatus(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.Status = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetTagList(v []*GetCardDetailResponseBodyDataVsimCardInfoTagList) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.TagList = v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetVendor(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.Vendor = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetVsimInstanceId(v int32) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.VsimInstanceId = &v
	return s
}

type GetCardDetailResponseBodyDataVsimCardInfoTagList struct {
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s GetCardDetailResponseBodyDataVsimCardInfoTagList) String() string {
	return tea.Prettify(s)
}

func (s GetCardDetailResponseBodyDataVsimCardInfoTagList) GoString() string {
	return s.String()
}

func (s *GetCardDetailResponseBodyDataVsimCardInfoTagList) SetId(v int64) *GetCardDetailResponseBodyDataVsimCardInfoTagList {
	s.Id = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfoTagList) SetTagName(v string) *GetCardDetailResponseBodyDataVsimCardInfoTagList {
	s.TagName = &v
	return s
}

type GetCardDetailResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCardDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCardDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCardDetailResponse) GoString() string {
	return s.String()
}

func (s *GetCardDetailResponse) SetHeaders(v map[string]*string) *GetCardDetailResponse {
	s.Headers = v
	return s
}

func (s *GetCardDetailResponse) SetBody(v *GetCardDetailResponseBody) *GetCardDetailResponse {
	s.Body = v
	return s
}

type GetCardFlowInfoRequest struct {
	DateList []*string `json:"DateList,omitempty" xml:"DateList,omitempty" type:"Repeated"`
	Iccid    *string   `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
}

func (s GetCardFlowInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCardFlowInfoRequest) GoString() string {
	return s.String()
}

func (s *GetCardFlowInfoRequest) SetDateList(v []*string) *GetCardFlowInfoRequest {
	s.DateList = v
	return s
}

func (s *GetCardFlowInfoRequest) SetIccid(v string) *GetCardFlowInfoRequest {
	s.Iccid = &v
	return s
}

type GetCardFlowInfoResponseBody struct {
	Code             *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *GetCardFlowInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage     *string                          `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string                          `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCardFlowInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCardFlowInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetCardFlowInfoResponseBody) SetCode(v string) *GetCardFlowInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetCardFlowInfoResponseBody) SetData(v *GetCardFlowInfoResponseBodyData) *GetCardFlowInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetCardFlowInfoResponseBody) SetErrorMessage(v string) *GetCardFlowInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetCardFlowInfoResponseBody) SetLocalizedMessage(v string) *GetCardFlowInfoResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *GetCardFlowInfoResponseBody) SetRequestId(v string) *GetCardFlowInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCardFlowInfoResponseBody) SetSuccess(v bool) *GetCardFlowInfoResponseBody {
	s.Success = &v
	return s
}

type GetCardFlowInfoResponseBodyData struct {
	ListCardMonthFlow []*GetCardFlowInfoResponseBodyDataListCardMonthFlow `json:"ListCardMonthFlow,omitempty" xml:"ListCardMonthFlow,omitempty" type:"Repeated"`
	ListPackageDTO    []*GetCardFlowInfoResponseBodyDataListPackageDTO    `json:"ListPackageDTO,omitempty" xml:"ListPackageDTO,omitempty" type:"Repeated"`
	ListVendorDetail  []*GetCardFlowInfoResponseBodyDataListVendorDetail  `json:"ListVendorDetail,omitempty" xml:"ListVendorDetail,omitempty" type:"Repeated"`
}

func (s GetCardFlowInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCardFlowInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCardFlowInfoResponseBodyData) SetListCardMonthFlow(v []*GetCardFlowInfoResponseBodyDataListCardMonthFlow) *GetCardFlowInfoResponseBodyData {
	s.ListCardMonthFlow = v
	return s
}

func (s *GetCardFlowInfoResponseBodyData) SetListPackageDTO(v []*GetCardFlowInfoResponseBodyDataListPackageDTO) *GetCardFlowInfoResponseBodyData {
	s.ListPackageDTO = v
	return s
}

func (s *GetCardFlowInfoResponseBodyData) SetListVendorDetail(v []*GetCardFlowInfoResponseBodyDataListVendorDetail) *GetCardFlowInfoResponseBodyData {
	s.ListVendorDetail = v
	return s
}

type GetCardFlowInfoResponseBodyDataListCardMonthFlow struct {
	FlowCount   *string                                                        `json:"FlowCount,omitempty" xml:"FlowCount,omitempty"`
	ListDayFlow []*GetCardFlowInfoResponseBodyDataListCardMonthFlowListDayFlow `json:"ListDayFlow,omitempty" xml:"ListDayFlow,omitempty" type:"Repeated"`
	Month       *string                                                        `json:"Month,omitempty" xml:"Month,omitempty"`
}

func (s GetCardFlowInfoResponseBodyDataListCardMonthFlow) String() string {
	return tea.Prettify(s)
}

func (s GetCardFlowInfoResponseBodyDataListCardMonthFlow) GoString() string {
	return s.String()
}

func (s *GetCardFlowInfoResponseBodyDataListCardMonthFlow) SetFlowCount(v string) *GetCardFlowInfoResponseBodyDataListCardMonthFlow {
	s.FlowCount = &v
	return s
}

func (s *GetCardFlowInfoResponseBodyDataListCardMonthFlow) SetListDayFlow(v []*GetCardFlowInfoResponseBodyDataListCardMonthFlowListDayFlow) *GetCardFlowInfoResponseBodyDataListCardMonthFlow {
	s.ListDayFlow = v
	return s
}

func (s *GetCardFlowInfoResponseBodyDataListCardMonthFlow) SetMonth(v string) *GetCardFlowInfoResponseBodyDataListCardMonthFlow {
	s.Month = &v
	return s
}

type GetCardFlowInfoResponseBodyDataListCardMonthFlowListDayFlow struct {
	Day  *string `json:"Day,omitempty" xml:"Day,omitempty"`
	Flow *string `json:"Flow,omitempty" xml:"Flow,omitempty"`
}

func (s GetCardFlowInfoResponseBodyDataListCardMonthFlowListDayFlow) String() string {
	return tea.Prettify(s)
}

func (s GetCardFlowInfoResponseBodyDataListCardMonthFlowListDayFlow) GoString() string {
	return s.String()
}

func (s *GetCardFlowInfoResponseBodyDataListCardMonthFlowListDayFlow) SetDay(v string) *GetCardFlowInfoResponseBodyDataListCardMonthFlowListDayFlow {
	s.Day = &v
	return s
}

func (s *GetCardFlowInfoResponseBodyDataListCardMonthFlowListDayFlow) SetFlow(v string) *GetCardFlowInfoResponseBodyDataListCardMonthFlowListDayFlow {
	s.Flow = &v
	return s
}

type GetCardFlowInfoResponseBodyDataListPackageDTO struct {
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	ExpireTime    *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	PackageName   *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	Remark        *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s GetCardFlowInfoResponseBodyDataListPackageDTO) String() string {
	return tea.Prettify(s)
}

func (s GetCardFlowInfoResponseBodyDataListPackageDTO) GoString() string {
	return s.String()
}

func (s *GetCardFlowInfoResponseBodyDataListPackageDTO) SetEffectiveTime(v string) *GetCardFlowInfoResponseBodyDataListPackageDTO {
	s.EffectiveTime = &v
	return s
}

func (s *GetCardFlowInfoResponseBodyDataListPackageDTO) SetExpireTime(v string) *GetCardFlowInfoResponseBodyDataListPackageDTO {
	s.ExpireTime = &v
	return s
}

func (s *GetCardFlowInfoResponseBodyDataListPackageDTO) SetPackageName(v string) *GetCardFlowInfoResponseBodyDataListPackageDTO {
	s.PackageName = &v
	return s
}

func (s *GetCardFlowInfoResponseBodyDataListPackageDTO) SetRemark(v string) *GetCardFlowInfoResponseBodyDataListPackageDTO {
	s.Remark = &v
	return s
}

type GetCardFlowInfoResponseBodyDataListVendorDetail struct {
	NetWorkDelay   *string `json:"NetWorkDelay,omitempty" xml:"NetWorkDelay,omitempty"`
	SignalStrength *string `json:"SignalStrength,omitempty" xml:"SignalStrength,omitempty"`
	Vendor         *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s GetCardFlowInfoResponseBodyDataListVendorDetail) String() string {
	return tea.Prettify(s)
}

func (s GetCardFlowInfoResponseBodyDataListVendorDetail) GoString() string {
	return s.String()
}

func (s *GetCardFlowInfoResponseBodyDataListVendorDetail) SetNetWorkDelay(v string) *GetCardFlowInfoResponseBodyDataListVendorDetail {
	s.NetWorkDelay = &v
	return s
}

func (s *GetCardFlowInfoResponseBodyDataListVendorDetail) SetSignalStrength(v string) *GetCardFlowInfoResponseBodyDataListVendorDetail {
	s.SignalStrength = &v
	return s
}

func (s *GetCardFlowInfoResponseBodyDataListVendorDetail) SetVendor(v string) *GetCardFlowInfoResponseBodyDataListVendorDetail {
	s.Vendor = &v
	return s
}

type GetCardFlowInfoResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCardFlowInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCardFlowInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCardFlowInfoResponse) GoString() string {
	return s.String()
}

func (s *GetCardFlowInfoResponse) SetHeaders(v map[string]*string) *GetCardFlowInfoResponse {
	s.Headers = v
	return s
}

func (s *GetCardFlowInfoResponse) SetBody(v *GetCardFlowInfoResponseBody) *GetCardFlowInfoResponse {
	s.Body = v
	return s
}

type GetCredentialPoolStatisticsRequest struct {
	CredentialNO *string `json:"CredentialNO,omitempty" xml:"CredentialNO,omitempty"`
	Date         *string `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s GetCredentialPoolStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCredentialPoolStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetCredentialPoolStatisticsRequest) SetCredentialNO(v string) *GetCredentialPoolStatisticsRequest {
	s.CredentialNO = &v
	return s
}

func (s *GetCredentialPoolStatisticsRequest) SetDate(v string) *GetCredentialPoolStatisticsRequest {
	s.Date = &v
	return s
}

type GetCredentialPoolStatisticsResponseBody struct {
	Code         *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *GetCredentialPoolStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCredentialPoolStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCredentialPoolStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetCredentialPoolStatisticsResponseBody) SetCode(v string) *GetCredentialPoolStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBody) SetData(v *GetCredentialPoolStatisticsResponseBodyData) *GetCredentialPoolStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBody) SetErrorMessage(v string) *GetCredentialPoolStatisticsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBody) SetRequestId(v string) *GetCredentialPoolStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBody) SetSuccess(v bool) *GetCredentialPoolStatisticsResponseBody {
	s.Success = &v
	return s
}

type GetCredentialPoolStatisticsResponseBodyData struct {
	CardActiveNum          *int64  `json:"CardActiveNum,omitempty" xml:"CardActiveNum,omitempty"`
	CardTotalNum           *int64  `json:"CardTotalNum,omitempty" xml:"CardTotalNum,omitempty"`
	CredentialInstanceId   *string `json:"CredentialInstanceId,omitempty" xml:"CredentialInstanceId,omitempty"`
	CredentialNO           *string `json:"CredentialNO,omitempty" xml:"CredentialNO,omitempty"`
	CredentialType         *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	EffectiveAvailableFlow *string `json:"EffectiveAvailableFlow,omitempty" xml:"EffectiveAvailableFlow,omitempty"`
	EffectiveTotalFlow     *string `json:"EffectiveTotalFlow,omitempty" xml:"EffectiveTotalFlow,omitempty"`
	MonthFeatureFee        *int64  `json:"MonthFeatureFee,omitempty" xml:"MonthFeatureFee,omitempty"`
	PoolAvaiable           *string `json:"PoolAvaiable,omitempty" xml:"PoolAvaiable,omitempty"`
	PoolGrandTotal         *string `json:"PoolGrandTotal,omitempty" xml:"PoolGrandTotal,omitempty"`
	PoolGrandTotalUsed     *string `json:"PoolGrandTotalUsed,omitempty" xml:"PoolGrandTotalUsed,omitempty"`
	PoolOutUsed            *string `json:"PoolOutUsed,omitempty" xml:"PoolOutUsed,omitempty"`
	PoolUsed               *string `json:"PoolUsed,omitempty" xml:"PoolUsed,omitempty"`
	SmsUsed                *int64  `json:"SmsUsed,omitempty" xml:"SmsUsed,omitempty"`
}

func (s GetCredentialPoolStatisticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCredentialPoolStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetCardActiveNum(v int64) *GetCredentialPoolStatisticsResponseBodyData {
	s.CardActiveNum = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetCardTotalNum(v int64) *GetCredentialPoolStatisticsResponseBodyData {
	s.CardTotalNum = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetCredentialInstanceId(v string) *GetCredentialPoolStatisticsResponseBodyData {
	s.CredentialInstanceId = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetCredentialNO(v string) *GetCredentialPoolStatisticsResponseBodyData {
	s.CredentialNO = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetCredentialType(v string) *GetCredentialPoolStatisticsResponseBodyData {
	s.CredentialType = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetEffectiveAvailableFlow(v string) *GetCredentialPoolStatisticsResponseBodyData {
	s.EffectiveAvailableFlow = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetEffectiveTotalFlow(v string) *GetCredentialPoolStatisticsResponseBodyData {
	s.EffectiveTotalFlow = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetMonthFeatureFee(v int64) *GetCredentialPoolStatisticsResponseBodyData {
	s.MonthFeatureFee = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetPoolAvaiable(v string) *GetCredentialPoolStatisticsResponseBodyData {
	s.PoolAvaiable = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetPoolGrandTotal(v string) *GetCredentialPoolStatisticsResponseBodyData {
	s.PoolGrandTotal = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetPoolGrandTotalUsed(v string) *GetCredentialPoolStatisticsResponseBodyData {
	s.PoolGrandTotalUsed = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetPoolOutUsed(v string) *GetCredentialPoolStatisticsResponseBodyData {
	s.PoolOutUsed = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetPoolUsed(v string) *GetCredentialPoolStatisticsResponseBodyData {
	s.PoolUsed = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetSmsUsed(v int64) *GetCredentialPoolStatisticsResponseBodyData {
	s.SmsUsed = &v
	return s
}

type GetCredentialPoolStatisticsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCredentialPoolStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCredentialPoolStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCredentialPoolStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetCredentialPoolStatisticsResponse) SetHeaders(v map[string]*string) *GetCredentialPoolStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetCredentialPoolStatisticsResponse) SetBody(v *GetCredentialPoolStatisticsResponseBody) *GetCredentialPoolStatisticsResponse {
	s.Body = v
	return s
}

type RebindResumeSingleCardRequest struct {
	Iccid      *string                `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	OptMsisdns map[string]interface{} `json:"OptMsisdns,omitempty" xml:"OptMsisdns,omitempty"`
}

func (s RebindResumeSingleCardRequest) String() string {
	return tea.Prettify(s)
}

func (s RebindResumeSingleCardRequest) GoString() string {
	return s.String()
}

func (s *RebindResumeSingleCardRequest) SetIccid(v string) *RebindResumeSingleCardRequest {
	s.Iccid = &v
	return s
}

func (s *RebindResumeSingleCardRequest) SetOptMsisdns(v map[string]interface{}) *RebindResumeSingleCardRequest {
	s.OptMsisdns = v
	return s
}

type RebindResumeSingleCardShrinkRequest struct {
	Iccid            *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	OptMsisdnsShrink *string `json:"OptMsisdns,omitempty" xml:"OptMsisdns,omitempty"`
}

func (s RebindResumeSingleCardShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RebindResumeSingleCardShrinkRequest) GoString() string {
	return s.String()
}

func (s *RebindResumeSingleCardShrinkRequest) SetIccid(v string) *RebindResumeSingleCardShrinkRequest {
	s.Iccid = &v
	return s
}

func (s *RebindResumeSingleCardShrinkRequest) SetOptMsisdnsShrink(v string) *RebindResumeSingleCardShrinkRequest {
	s.OptMsisdnsShrink = &v
	return s
}

type RebindResumeSingleCardResponseBody struct {
	Code             *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage     *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RebindResumeSingleCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebindResumeSingleCardResponseBody) GoString() string {
	return s.String()
}

func (s *RebindResumeSingleCardResponseBody) SetCode(v string) *RebindResumeSingleCardResponseBody {
	s.Code = &v
	return s
}

func (s *RebindResumeSingleCardResponseBody) SetData(v bool) *RebindResumeSingleCardResponseBody {
	s.Data = &v
	return s
}

func (s *RebindResumeSingleCardResponseBody) SetErrorMessage(v string) *RebindResumeSingleCardResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RebindResumeSingleCardResponseBody) SetLocalizedMessage(v string) *RebindResumeSingleCardResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *RebindResumeSingleCardResponseBody) SetRequestId(v string) *RebindResumeSingleCardResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebindResumeSingleCardResponseBody) SetSuccess(v bool) *RebindResumeSingleCardResponseBody {
	s.Success = &v
	return s
}

type RebindResumeSingleCardResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RebindResumeSingleCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RebindResumeSingleCardResponse) String() string {
	return tea.Prettify(s)
}

func (s RebindResumeSingleCardResponse) GoString() string {
	return s.String()
}

func (s *RebindResumeSingleCardResponse) SetHeaders(v map[string]*string) *RebindResumeSingleCardResponse {
	s.Headers = v
	return s
}

func (s *RebindResumeSingleCardResponse) SetBody(v *RebindResumeSingleCardResponseBody) *RebindResumeSingleCardResponse {
	s.Body = v
	return s
}

type ResumeSingleCardRequest struct {
	Iccid      *string                `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	OptMsisdns map[string]interface{} `json:"OptMsisdns,omitempty" xml:"OptMsisdns,omitempty"`
}

func (s ResumeSingleCardRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeSingleCardRequest) GoString() string {
	return s.String()
}

func (s *ResumeSingleCardRequest) SetIccid(v string) *ResumeSingleCardRequest {
	s.Iccid = &v
	return s
}

func (s *ResumeSingleCardRequest) SetOptMsisdns(v map[string]interface{}) *ResumeSingleCardRequest {
	s.OptMsisdns = v
	return s
}

type ResumeSingleCardShrinkRequest struct {
	Iccid            *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	OptMsisdnsShrink *string `json:"OptMsisdns,omitempty" xml:"OptMsisdns,omitempty"`
}

func (s ResumeSingleCardShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeSingleCardShrinkRequest) GoString() string {
	return s.String()
}

func (s *ResumeSingleCardShrinkRequest) SetIccid(v string) *ResumeSingleCardShrinkRequest {
	s.Iccid = &v
	return s
}

func (s *ResumeSingleCardShrinkRequest) SetOptMsisdnsShrink(v string) *ResumeSingleCardShrinkRequest {
	s.OptMsisdnsShrink = &v
	return s
}

type ResumeSingleCardResponseBody struct {
	Code             *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage     *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResumeSingleCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeSingleCardResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeSingleCardResponseBody) SetCode(v string) *ResumeSingleCardResponseBody {
	s.Code = &v
	return s
}

func (s *ResumeSingleCardResponseBody) SetData(v bool) *ResumeSingleCardResponseBody {
	s.Data = &v
	return s
}

func (s *ResumeSingleCardResponseBody) SetErrorMessage(v string) *ResumeSingleCardResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ResumeSingleCardResponseBody) SetLocalizedMessage(v string) *ResumeSingleCardResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *ResumeSingleCardResponseBody) SetRequestId(v string) *ResumeSingleCardResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeSingleCardResponseBody) SetSuccess(v bool) *ResumeSingleCardResponseBody {
	s.Success = &v
	return s
}

type ResumeSingleCardResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResumeSingleCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResumeSingleCardResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeSingleCardResponse) GoString() string {
	return s.String()
}

func (s *ResumeSingleCardResponse) SetHeaders(v map[string]*string) *ResumeSingleCardResponse {
	s.Headers = v
	return s
}

func (s *ResumeSingleCardResponse) SetBody(v *ResumeSingleCardResponseBody) *ResumeSingleCardResponse {
	s.Body = v
	return s
}

type StopSingleCardRequest struct {
	Iccid      *string                `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	OptMsisdns map[string]interface{} `json:"OptMsisdns,omitempty" xml:"OptMsisdns,omitempty"`
}

func (s StopSingleCardRequest) String() string {
	return tea.Prettify(s)
}

func (s StopSingleCardRequest) GoString() string {
	return s.String()
}

func (s *StopSingleCardRequest) SetIccid(v string) *StopSingleCardRequest {
	s.Iccid = &v
	return s
}

func (s *StopSingleCardRequest) SetOptMsisdns(v map[string]interface{}) *StopSingleCardRequest {
	s.OptMsisdns = v
	return s
}

type StopSingleCardShrinkRequest struct {
	Iccid            *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	OptMsisdnsShrink *string `json:"OptMsisdns,omitempty" xml:"OptMsisdns,omitempty"`
}

func (s StopSingleCardShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s StopSingleCardShrinkRequest) GoString() string {
	return s.String()
}

func (s *StopSingleCardShrinkRequest) SetIccid(v string) *StopSingleCardShrinkRequest {
	s.Iccid = &v
	return s
}

func (s *StopSingleCardShrinkRequest) SetOptMsisdnsShrink(v string) *StopSingleCardShrinkRequest {
	s.OptMsisdnsShrink = &v
	return s
}

type StopSingleCardResponseBody struct {
	Code             *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage     *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopSingleCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopSingleCardResponseBody) GoString() string {
	return s.String()
}

func (s *StopSingleCardResponseBody) SetCode(v string) *StopSingleCardResponseBody {
	s.Code = &v
	return s
}

func (s *StopSingleCardResponseBody) SetData(v bool) *StopSingleCardResponseBody {
	s.Data = &v
	return s
}

func (s *StopSingleCardResponseBody) SetErrorMessage(v string) *StopSingleCardResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StopSingleCardResponseBody) SetLocalizedMessage(v string) *StopSingleCardResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *StopSingleCardResponseBody) SetRequestId(v string) *StopSingleCardResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopSingleCardResponseBody) SetSuccess(v bool) *StopSingleCardResponseBody {
	s.Success = &v
	return s
}

type StopSingleCardResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopSingleCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopSingleCardResponse) String() string {
	return tea.Prettify(s)
}

func (s StopSingleCardResponse) GoString() string {
	return s.String()
}

func (s *StopSingleCardResponse) SetHeaders(v map[string]*string) *StopSingleCardResponse {
	s.Headers = v
	return s
}

func (s *StopSingleCardResponse) SetBody(v *StopSingleCardResponseBody) *StopSingleCardResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("linkcard"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GetCardDetailWithOptions(request *GetCardDetailRequest, runtime *util.RuntimeOptions) (_result *GetCardDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestroyCard)) {
		query["DestroyCard"] = request.DestroyCard
	}

	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ShowPsim)) {
		query["ShowPsim"] = request.ShowPsim
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCardDetail"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCardDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCardDetail(request *GetCardDetailRequest) (_result *GetCardDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCardDetailResponse{}
	_body, _err := client.GetCardDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCardFlowInfoWithOptions(request *GetCardFlowInfoRequest, runtime *util.RuntimeOptions) (_result *GetCardFlowInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DateList)) {
		query["DateList"] = request.DateList
	}

	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCardFlowInfo"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCardFlowInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCardFlowInfo(request *GetCardFlowInfoRequest) (_result *GetCardFlowInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCardFlowInfoResponse{}
	_body, _err := client.GetCardFlowInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCredentialPoolStatisticsWithOptions(request *GetCredentialPoolStatisticsRequest, runtime *util.RuntimeOptions) (_result *GetCredentialPoolStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CredentialNO)) {
		query["CredentialNO"] = request.CredentialNO
	}

	if !tea.BoolValue(util.IsUnset(request.Date)) {
		query["Date"] = request.Date
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCredentialPoolStatistics"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCredentialPoolStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCredentialPoolStatistics(request *GetCredentialPoolStatisticsRequest) (_result *GetCredentialPoolStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCredentialPoolStatisticsResponse{}
	_body, _err := client.GetCredentialPoolStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RebindResumeSingleCardWithOptions(tmpReq *RebindResumeSingleCardRequest, runtime *util.RuntimeOptions) (_result *RebindResumeSingleCardResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RebindResumeSingleCardShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.OptMsisdns)) {
		request.OptMsisdnsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OptMsisdns, tea.String("OptMsisdns"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
	}

	if !tea.BoolValue(util.IsUnset(request.OptMsisdnsShrink)) {
		query["OptMsisdns"] = request.OptMsisdnsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RebindResumeSingleCard"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RebindResumeSingleCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RebindResumeSingleCard(request *RebindResumeSingleCardRequest) (_result *RebindResumeSingleCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RebindResumeSingleCardResponse{}
	_body, _err := client.RebindResumeSingleCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResumeSingleCardWithOptions(tmpReq *ResumeSingleCardRequest, runtime *util.RuntimeOptions) (_result *ResumeSingleCardResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ResumeSingleCardShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.OptMsisdns)) {
		request.OptMsisdnsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OptMsisdns, tea.String("OptMsisdns"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
	}

	if !tea.BoolValue(util.IsUnset(request.OptMsisdnsShrink)) {
		query["OptMsisdns"] = request.OptMsisdnsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResumeSingleCard"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResumeSingleCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResumeSingleCard(request *ResumeSingleCardRequest) (_result *ResumeSingleCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResumeSingleCardResponse{}
	_body, _err := client.ResumeSingleCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopSingleCardWithOptions(tmpReq *StopSingleCardRequest, runtime *util.RuntimeOptions) (_result *StopSingleCardResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &StopSingleCardShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.OptMsisdns)) {
		request.OptMsisdnsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OptMsisdns, tea.String("OptMsisdns"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
	}

	if !tea.BoolValue(util.IsUnset(request.OptMsisdnsShrink)) {
		query["OptMsisdns"] = request.OptMsisdnsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopSingleCard"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopSingleCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopSingleCard(request *StopSingleCardRequest) (_result *StopSingleCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopSingleCardResponse{}
	_body, _err := client.StopSingleCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
