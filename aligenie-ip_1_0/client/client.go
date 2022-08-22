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

type DeviceControlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeviceControlHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeviceControlHeaders) GoString() string {
	return s.String()
}

func (s *DeviceControlHeaders) SetCommonHeaders(v map[string]*string) *DeviceControlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeviceControlHeaders) SetXAcsAligenieAccessToken(v string) *DeviceControlHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *DeviceControlHeaders) SetAuthorization(v string) *DeviceControlHeaders {
	s.Authorization = &v
	return s
}

type DeviceControlRequest struct {
	Payload  *DeviceControlRequestPayload  `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo *DeviceControlRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s DeviceControlRequest) String() string {
	return tea.Prettify(s)
}

func (s DeviceControlRequest) GoString() string {
	return s.String()
}

func (s *DeviceControlRequest) SetPayload(v *DeviceControlRequestPayload) *DeviceControlRequest {
	s.Payload = v
	return s
}

func (s *DeviceControlRequest) SetUserInfo(v *DeviceControlRequestUserInfo) *DeviceControlRequest {
	s.UserInfo = v
	return s
}

type DeviceControlRequestPayload struct {
	Category     *string            `json:"Category,omitempty" xml:"Category,omitempty"`
	Cmd          *string            `json:"Cmd,omitempty" xml:"Cmd,omitempty"`
	Current      *string            `json:"Current,omitempty" xml:"Current,omitempty"`
	Device       *string            `json:"Device,omitempty" xml:"Device,omitempty"`
	DeviceNumber *string            `json:"DeviceNumber,omitempty" xml:"DeviceNumber,omitempty"`
	ExtendInfo   *string            `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	Location     *string            `json:"Location,omitempty" xml:"Location,omitempty"`
	Properties   map[string]*string `json:"Properties,omitempty" xml:"Properties,omitempty"`
}

func (s DeviceControlRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s DeviceControlRequestPayload) GoString() string {
	return s.String()
}

func (s *DeviceControlRequestPayload) SetCategory(v string) *DeviceControlRequestPayload {
	s.Category = &v
	return s
}

func (s *DeviceControlRequestPayload) SetCmd(v string) *DeviceControlRequestPayload {
	s.Cmd = &v
	return s
}

func (s *DeviceControlRequestPayload) SetCurrent(v string) *DeviceControlRequestPayload {
	s.Current = &v
	return s
}

func (s *DeviceControlRequestPayload) SetDevice(v string) *DeviceControlRequestPayload {
	s.Device = &v
	return s
}

func (s *DeviceControlRequestPayload) SetDeviceNumber(v string) *DeviceControlRequestPayload {
	s.DeviceNumber = &v
	return s
}

func (s *DeviceControlRequestPayload) SetExtendInfo(v string) *DeviceControlRequestPayload {
	s.ExtendInfo = &v
	return s
}

func (s *DeviceControlRequestPayload) SetLocation(v string) *DeviceControlRequestPayload {
	s.Location = &v
	return s
}

func (s *DeviceControlRequestPayload) SetProperties(v map[string]*string) *DeviceControlRequestPayload {
	s.Properties = v
	return s
}

type DeviceControlRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DeviceControlRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s DeviceControlRequestUserInfo) GoString() string {
	return s.String()
}

func (s *DeviceControlRequestUserInfo) SetEncodeKey(v string) *DeviceControlRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *DeviceControlRequestUserInfo) SetEncodeType(v string) *DeviceControlRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *DeviceControlRequestUserInfo) SetId(v string) *DeviceControlRequestUserInfo {
	s.Id = &v
	return s
}

func (s *DeviceControlRequestUserInfo) SetIdType(v string) *DeviceControlRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *DeviceControlRequestUserInfo) SetOrganizationId(v string) *DeviceControlRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type DeviceControlShrinkRequest struct {
	PayloadShrink  *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s DeviceControlShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeviceControlShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeviceControlShrinkRequest) SetPayloadShrink(v string) *DeviceControlShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *DeviceControlShrinkRequest) SetUserInfoShrink(v string) *DeviceControlShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type DeviceControlResponseBody struct {
	Code      *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DeviceControlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeviceControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeviceControlResponseBody) GoString() string {
	return s.String()
}

func (s *DeviceControlResponseBody) SetCode(v int32) *DeviceControlResponseBody {
	s.Code = &v
	return s
}

func (s *DeviceControlResponseBody) SetMessage(v string) *DeviceControlResponseBody {
	s.Message = &v
	return s
}

func (s *DeviceControlResponseBody) SetRequestId(v string) *DeviceControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeviceControlResponseBody) SetResult(v *DeviceControlResponseBodyResult) *DeviceControlResponseBody {
	s.Result = v
	return s
}

type DeviceControlResponseBodyResult struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeviceControlResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeviceControlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeviceControlResponseBodyResult) SetStatus(v string) *DeviceControlResponseBodyResult {
	s.Status = &v
	return s
}

type DeviceControlResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeviceControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeviceControlResponse) String() string {
	return tea.Prettify(s)
}

func (s DeviceControlResponse) GoString() string {
	return s.String()
}

func (s *DeviceControlResponse) SetHeaders(v map[string]*string) *DeviceControlResponse {
	s.Headers = v
	return s
}

func (s *DeviceControlResponse) SetStatusCode(v int32) *DeviceControlResponse {
	s.StatusCode = &v
	return s
}

func (s *DeviceControlResponse) SetBody(v *DeviceControlResponseBody) *DeviceControlResponse {
	s.Body = v
	return s
}

type GetHotelHomeBackImageAndModesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelHomeBackImageAndModesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetHotelHomeBackImageAndModesHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelHomeBackImageAndModesHeaders) SetCommonHeaders(v map[string]*string) *GetHotelHomeBackImageAndModesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelHomeBackImageAndModesHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelHomeBackImageAndModesHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesHeaders) SetAuthorization(v string) *GetHotelHomeBackImageAndModesHeaders {
	s.Authorization = &v
	return s
}

type GetHotelHomeBackImageAndModesRequest struct {
	UserInfo *GetHotelHomeBackImageAndModesRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetHotelHomeBackImageAndModesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelHomeBackImageAndModesRequest) GoString() string {
	return s.String()
}

func (s *GetHotelHomeBackImageAndModesRequest) SetUserInfo(v *GetHotelHomeBackImageAndModesRequestUserInfo) *GetHotelHomeBackImageAndModesRequest {
	s.UserInfo = v
	return s
}

type GetHotelHomeBackImageAndModesRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetHotelHomeBackImageAndModesRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s GetHotelHomeBackImageAndModesRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetHotelHomeBackImageAndModesRequestUserInfo) SetEncodeKey(v string) *GetHotelHomeBackImageAndModesRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesRequestUserInfo) SetEncodeType(v string) *GetHotelHomeBackImageAndModesRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesRequestUserInfo) SetId(v string) *GetHotelHomeBackImageAndModesRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesRequestUserInfo) SetIdType(v string) *GetHotelHomeBackImageAndModesRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesRequestUserInfo) SetOrganizationId(v string) *GetHotelHomeBackImageAndModesRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type GetHotelHomeBackImageAndModesShrinkRequest struct {
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetHotelHomeBackImageAndModesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelHomeBackImageAndModesShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetHotelHomeBackImageAndModesShrinkRequest) SetUserInfoShrink(v string) *GetHotelHomeBackImageAndModesShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type GetHotelHomeBackImageAndModesResponseBody struct {
	Code      *int32                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetHotelHomeBackImageAndModesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetHotelHomeBackImageAndModesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotelHomeBackImageAndModesResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelHomeBackImageAndModesResponseBody) SetCode(v int32) *GetHotelHomeBackImageAndModesResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponseBody) SetMessage(v string) *GetHotelHomeBackImageAndModesResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponseBody) SetRequestId(v string) *GetHotelHomeBackImageAndModesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponseBody) SetResult(v *GetHotelHomeBackImageAndModesResponseBodyResult) *GetHotelHomeBackImageAndModesResponseBody {
	s.Result = v
	return s
}

type GetHotelHomeBackImageAndModesResponseBodyResult struct {
	BackgroundImage *string                                                    `json:"BackgroundImage,omitempty" xml:"BackgroundImage,omitempty"`
	HotelName       *string                                                    `json:"HotelName,omitempty" xml:"HotelName,omitempty"`
	ModeList        []*GetHotelHomeBackImageAndModesResponseBodyResultModeList `json:"ModeList,omitempty" xml:"ModeList,omitempty" type:"Repeated"`
}

func (s GetHotelHomeBackImageAndModesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetHotelHomeBackImageAndModesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetHotelHomeBackImageAndModesResponseBodyResult) SetBackgroundImage(v string) *GetHotelHomeBackImageAndModesResponseBodyResult {
	s.BackgroundImage = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponseBodyResult) SetHotelName(v string) *GetHotelHomeBackImageAndModesResponseBodyResult {
	s.HotelName = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponseBodyResult) SetModeList(v []*GetHotelHomeBackImageAndModesResponseBodyResultModeList) *GetHotelHomeBackImageAndModesResponseBodyResult {
	s.ModeList = v
	return s
}

type GetHotelHomeBackImageAndModesResponseBodyResultModeList struct {
	CnName *string `json:"CnName,omitempty" xml:"CnName,omitempty"`
	Code   *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Icon   *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
}

func (s GetHotelHomeBackImageAndModesResponseBodyResultModeList) String() string {
	return tea.Prettify(s)
}

func (s GetHotelHomeBackImageAndModesResponseBodyResultModeList) GoString() string {
	return s.String()
}

func (s *GetHotelHomeBackImageAndModesResponseBodyResultModeList) SetCnName(v string) *GetHotelHomeBackImageAndModesResponseBodyResultModeList {
	s.CnName = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponseBodyResultModeList) SetCode(v string) *GetHotelHomeBackImageAndModesResponseBodyResultModeList {
	s.Code = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponseBodyResultModeList) SetIcon(v string) *GetHotelHomeBackImageAndModesResponseBodyResultModeList {
	s.Icon = &v
	return s
}

type GetHotelHomeBackImageAndModesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHotelHomeBackImageAndModesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotelHomeBackImageAndModesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotelHomeBackImageAndModesResponse) GoString() string {
	return s.String()
}

func (s *GetHotelHomeBackImageAndModesResponse) SetHeaders(v map[string]*string) *GetHotelHomeBackImageAndModesResponse {
	s.Headers = v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponse) SetStatusCode(v int32) *GetHotelHomeBackImageAndModesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponse) SetBody(v *GetHotelHomeBackImageAndModesResponseBody) *GetHotelHomeBackImageAndModesResponse {
	s.Body = v
	return s
}

type GetHotelOrderDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelOrderDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetHotelOrderDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelOrderDetailHeaders) SetCommonHeaders(v map[string]*string) *GetHotelOrderDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelOrderDetailHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelOrderDetailHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelOrderDetailHeaders) SetAuthorization(v string) *GetHotelOrderDetailHeaders {
	s.Authorization = &v
	return s
}

type GetHotelOrderDetailRequest struct {
	Payload *GetHotelOrderDetailRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
}

func (s GetHotelOrderDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelOrderDetailRequest) GoString() string {
	return s.String()
}

func (s *GetHotelOrderDetailRequest) SetPayload(v *GetHotelOrderDetailRequestPayload) *GetHotelOrderDetailRequest {
	s.Payload = v
	return s
}

type GetHotelOrderDetailRequestPayload struct {
	OrderNo *string `json:"OrderNo,omitempty" xml:"OrderNo,omitempty"`
}

func (s GetHotelOrderDetailRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s GetHotelOrderDetailRequestPayload) GoString() string {
	return s.String()
}

func (s *GetHotelOrderDetailRequestPayload) SetOrderNo(v string) *GetHotelOrderDetailRequestPayload {
	s.OrderNo = &v
	return s
}

type GetHotelOrderDetailShrinkRequest struct {
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
}

func (s GetHotelOrderDetailShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelOrderDetailShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetHotelOrderDetailShrinkRequest) SetPayloadShrink(v string) *GetHotelOrderDetailShrinkRequest {
	s.PayloadShrink = &v
	return s
}

type GetHotelOrderDetailResponseBody struct {
	Code      *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*GetHotelOrderDetailResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s GetHotelOrderDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotelOrderDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelOrderDetailResponseBody) SetCode(v int32) *GetHotelOrderDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotelOrderDetailResponseBody) SetMessage(v string) *GetHotelOrderDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelOrderDetailResponseBody) SetRequestId(v string) *GetHotelOrderDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelOrderDetailResponseBody) SetResult(v []*GetHotelOrderDetailResponseBodyResult) *GetHotelOrderDetailResponseBody {
	s.Result = v
	return s
}

type GetHotelOrderDetailResponseBodyResult struct {
	ApplyAmt  *int64  `json:"ApplyAmt,omitempty" xml:"ApplyAmt,omitempty"`
	GmtCreate *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	ItemUrl   *string `json:"ItemUrl,omitempty" xml:"ItemUrl,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Quantity  *int64  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
}

func (s GetHotelOrderDetailResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetHotelOrderDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetHotelOrderDetailResponseBodyResult) SetApplyAmt(v int64) *GetHotelOrderDetailResponseBodyResult {
	s.ApplyAmt = &v
	return s
}

func (s *GetHotelOrderDetailResponseBodyResult) SetGmtCreate(v int64) *GetHotelOrderDetailResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *GetHotelOrderDetailResponseBodyResult) SetItemUrl(v string) *GetHotelOrderDetailResponseBodyResult {
	s.ItemUrl = &v
	return s
}

func (s *GetHotelOrderDetailResponseBodyResult) SetName(v string) *GetHotelOrderDetailResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetHotelOrderDetailResponseBodyResult) SetQuantity(v int64) *GetHotelOrderDetailResponseBodyResult {
	s.Quantity = &v
	return s
}

type GetHotelOrderDetailResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHotelOrderDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotelOrderDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotelOrderDetailResponse) GoString() string {
	return s.String()
}

func (s *GetHotelOrderDetailResponse) SetHeaders(v map[string]*string) *GetHotelOrderDetailResponse {
	s.Headers = v
	return s
}

func (s *GetHotelOrderDetailResponse) SetStatusCode(v int32) *GetHotelOrderDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelOrderDetailResponse) SetBody(v *GetHotelOrderDetailResponseBody) *GetHotelOrderDetailResponse {
	s.Body = v
	return s
}

type GetHotelSampleUtterancesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelSampleUtterancesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetHotelSampleUtterancesHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelSampleUtterancesHeaders) SetCommonHeaders(v map[string]*string) *GetHotelSampleUtterancesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelSampleUtterancesHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelSampleUtterancesHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelSampleUtterancesHeaders) SetAuthorization(v string) *GetHotelSampleUtterancesHeaders {
	s.Authorization = &v
	return s
}

type GetHotelSampleUtterancesRequest struct {
	UserInfo *GetHotelSampleUtterancesRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetHotelSampleUtterancesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelSampleUtterancesRequest) GoString() string {
	return s.String()
}

func (s *GetHotelSampleUtterancesRequest) SetUserInfo(v *GetHotelSampleUtterancesRequestUserInfo) *GetHotelSampleUtterancesRequest {
	s.UserInfo = v
	return s
}

type GetHotelSampleUtterancesRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetHotelSampleUtterancesRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s GetHotelSampleUtterancesRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetHotelSampleUtterancesRequestUserInfo) SetEncodeKey(v string) *GetHotelSampleUtterancesRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetHotelSampleUtterancesRequestUserInfo) SetEncodeType(v string) *GetHotelSampleUtterancesRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetHotelSampleUtterancesRequestUserInfo) SetId(v string) *GetHotelSampleUtterancesRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetHotelSampleUtterancesRequestUserInfo) SetIdType(v string) *GetHotelSampleUtterancesRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetHotelSampleUtterancesRequestUserInfo) SetOrganizationId(v string) *GetHotelSampleUtterancesRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type GetHotelSampleUtterancesShrinkRequest struct {
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetHotelSampleUtterancesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelSampleUtterancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetHotelSampleUtterancesShrinkRequest) SetUserInfoShrink(v string) *GetHotelSampleUtterancesShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type GetHotelSampleUtterancesResponseBody struct {
	Code      *int32    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s GetHotelSampleUtterancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotelSampleUtterancesResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelSampleUtterancesResponseBody) SetCode(v int32) *GetHotelSampleUtterancesResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotelSampleUtterancesResponseBody) SetMessage(v string) *GetHotelSampleUtterancesResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelSampleUtterancesResponseBody) SetRequestId(v string) *GetHotelSampleUtterancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelSampleUtterancesResponseBody) SetResult(v []*string) *GetHotelSampleUtterancesResponseBody {
	s.Result = v
	return s
}

type GetHotelSampleUtterancesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHotelSampleUtterancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotelSampleUtterancesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotelSampleUtterancesResponse) GoString() string {
	return s.String()
}

func (s *GetHotelSampleUtterancesResponse) SetHeaders(v map[string]*string) *GetHotelSampleUtterancesResponse {
	s.Headers = v
	return s
}

func (s *GetHotelSampleUtterancesResponse) SetStatusCode(v int32) *GetHotelSampleUtterancesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelSampleUtterancesResponse) SetBody(v *GetHotelSampleUtterancesResponseBody) *GetHotelSampleUtterancesResponse {
	s.Body = v
	return s
}

type GetHotelScreenSaverHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelScreenSaverHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetHotelScreenSaverHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverHeaders) SetCommonHeaders(v map[string]*string) *GetHotelScreenSaverHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelScreenSaverHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelScreenSaverHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelScreenSaverHeaders) SetAuthorization(v string) *GetHotelScreenSaverHeaders {
	s.Authorization = &v
	return s
}

type GetHotelScreenSaverRequest struct {
	UserInfo *GetHotelScreenSaverRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetHotelScreenSaverRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelScreenSaverRequest) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverRequest) SetUserInfo(v *GetHotelScreenSaverRequestUserInfo) *GetHotelScreenSaverRequest {
	s.UserInfo = v
	return s
}

type GetHotelScreenSaverRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetHotelScreenSaverRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s GetHotelScreenSaverRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverRequestUserInfo) SetEncodeKey(v string) *GetHotelScreenSaverRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetHotelScreenSaverRequestUserInfo) SetEncodeType(v string) *GetHotelScreenSaverRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetHotelScreenSaverRequestUserInfo) SetId(v string) *GetHotelScreenSaverRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetHotelScreenSaverRequestUserInfo) SetIdType(v string) *GetHotelScreenSaverRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetHotelScreenSaverRequestUserInfo) SetOrganizationId(v string) *GetHotelScreenSaverRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type GetHotelScreenSaverShrinkRequest struct {
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetHotelScreenSaverShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelScreenSaverShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverShrinkRequest) SetUserInfoShrink(v string) *GetHotelScreenSaverShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type GetHotelScreenSaverResponseBody struct {
	Code      *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetHotelScreenSaverResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetHotelScreenSaverResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotelScreenSaverResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverResponseBody) SetCode(v int32) *GetHotelScreenSaverResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotelScreenSaverResponseBody) SetMessage(v string) *GetHotelScreenSaverResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelScreenSaverResponseBody) SetRequestId(v string) *GetHotelScreenSaverResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelScreenSaverResponseBody) SetResult(v *GetHotelScreenSaverResponseBodyResult) *GetHotelScreenSaverResponseBody {
	s.Result = v
	return s
}

type GetHotelScreenSaverResponseBodyResult struct {
	PicUrl    *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	StyleCode *string `json:"StyleCode,omitempty" xml:"StyleCode,omitempty"`
}

func (s GetHotelScreenSaverResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetHotelScreenSaverResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverResponseBodyResult) SetPicUrl(v string) *GetHotelScreenSaverResponseBodyResult {
	s.PicUrl = &v
	return s
}

func (s *GetHotelScreenSaverResponseBodyResult) SetStyleCode(v string) *GetHotelScreenSaverResponseBodyResult {
	s.StyleCode = &v
	return s
}

type GetHotelScreenSaverResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHotelScreenSaverResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotelScreenSaverResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotelScreenSaverResponse) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverResponse) SetHeaders(v map[string]*string) *GetHotelScreenSaverResponse {
	s.Headers = v
	return s
}

func (s *GetHotelScreenSaverResponse) SetStatusCode(v int32) *GetHotelScreenSaverResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelScreenSaverResponse) SetBody(v *GetHotelScreenSaverResponseBody) *GetHotelScreenSaverResponse {
	s.Body = v
	return s
}

type ListHotelControlDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListHotelControlDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListHotelControlDeviceHeaders) GoString() string {
	return s.String()
}

func (s *ListHotelControlDeviceHeaders) SetCommonHeaders(v map[string]*string) *ListHotelControlDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListHotelControlDeviceHeaders) SetXAcsAligenieAccessToken(v string) *ListHotelControlDeviceHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListHotelControlDeviceHeaders) SetAuthorization(v string) *ListHotelControlDeviceHeaders {
	s.Authorization = &v
	return s
}

type ListHotelControlDeviceRequest struct {
	UserInfo *ListHotelControlDeviceRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListHotelControlDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotelControlDeviceRequest) GoString() string {
	return s.String()
}

func (s *ListHotelControlDeviceRequest) SetUserInfo(v *ListHotelControlDeviceRequestUserInfo) *ListHotelControlDeviceRequest {
	s.UserInfo = v
	return s
}

type ListHotelControlDeviceRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListHotelControlDeviceRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s ListHotelControlDeviceRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListHotelControlDeviceRequestUserInfo) SetEncodeKey(v string) *ListHotelControlDeviceRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListHotelControlDeviceRequestUserInfo) SetEncodeType(v string) *ListHotelControlDeviceRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListHotelControlDeviceRequestUserInfo) SetId(v string) *ListHotelControlDeviceRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListHotelControlDeviceRequestUserInfo) SetIdType(v string) *ListHotelControlDeviceRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListHotelControlDeviceRequestUserInfo) SetOrganizationId(v string) *ListHotelControlDeviceRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type ListHotelControlDeviceShrinkRequest struct {
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListHotelControlDeviceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotelControlDeviceShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHotelControlDeviceShrinkRequest) SetUserInfoShrink(v string) *ListHotelControlDeviceShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type ListHotelControlDeviceResponseBody struct {
	Code      *int32               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []map[string]*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListHotelControlDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotelControlDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotelControlDeviceResponseBody) SetCode(v int32) *ListHotelControlDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotelControlDeviceResponseBody) SetMessage(v string) *ListHotelControlDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotelControlDeviceResponseBody) SetRequestId(v string) *ListHotelControlDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotelControlDeviceResponseBody) SetResult(v []map[string]*string) *ListHotelControlDeviceResponseBody {
	s.Result = v
	return s
}

type ListHotelControlDeviceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListHotelControlDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHotelControlDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotelControlDeviceResponse) GoString() string {
	return s.String()
}

func (s *ListHotelControlDeviceResponse) SetHeaders(v map[string]*string) *ListHotelControlDeviceResponse {
	s.Headers = v
	return s
}

func (s *ListHotelControlDeviceResponse) SetStatusCode(v int32) *ListHotelControlDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotelControlDeviceResponse) SetBody(v *ListHotelControlDeviceResponseBody) *ListHotelControlDeviceResponse {
	s.Body = v
	return s
}

type ListHotelOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListHotelOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListHotelOrderHeaders) GoString() string {
	return s.String()
}

func (s *ListHotelOrderHeaders) SetCommonHeaders(v map[string]*string) *ListHotelOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListHotelOrderHeaders) SetXAcsAligenieAccessToken(v string) *ListHotelOrderHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListHotelOrderHeaders) SetAuthorization(v string) *ListHotelOrderHeaders {
	s.Authorization = &v
	return s
}

type ListHotelOrderRequest struct {
	Payload  *ListHotelOrderRequestPayload  `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo *ListHotelOrderRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListHotelOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotelOrderRequest) GoString() string {
	return s.String()
}

func (s *ListHotelOrderRequest) SetPayload(v *ListHotelOrderRequestPayload) *ListHotelOrderRequest {
	s.Payload = v
	return s
}

func (s *ListHotelOrderRequest) SetUserInfo(v *ListHotelOrderRequestUserInfo) *ListHotelOrderRequest {
	s.UserInfo = v
	return s
}

type ListHotelOrderRequestPayload struct {
	Page *ListHotelOrderRequestPayloadPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
}

func (s ListHotelOrderRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s ListHotelOrderRequestPayload) GoString() string {
	return s.String()
}

func (s *ListHotelOrderRequestPayload) SetPage(v *ListHotelOrderRequestPayloadPage) *ListHotelOrderRequestPayload {
	s.Page = v
	return s
}

type ListHotelOrderRequestPayloadPage struct {
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListHotelOrderRequestPayloadPage) String() string {
	return tea.Prettify(s)
}

func (s ListHotelOrderRequestPayloadPage) GoString() string {
	return s.String()
}

func (s *ListHotelOrderRequestPayloadPage) SetPageNumber(v int64) *ListHotelOrderRequestPayloadPage {
	s.PageNumber = &v
	return s
}

func (s *ListHotelOrderRequestPayloadPage) SetPageSize(v int64) *ListHotelOrderRequestPayloadPage {
	s.PageSize = &v
	return s
}

type ListHotelOrderRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListHotelOrderRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s ListHotelOrderRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListHotelOrderRequestUserInfo) SetEncodeKey(v string) *ListHotelOrderRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListHotelOrderRequestUserInfo) SetEncodeType(v string) *ListHotelOrderRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListHotelOrderRequestUserInfo) SetId(v string) *ListHotelOrderRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListHotelOrderRequestUserInfo) SetIdType(v string) *ListHotelOrderRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListHotelOrderRequestUserInfo) SetOrganizationId(v string) *ListHotelOrderRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type ListHotelOrderShrinkRequest struct {
	PayloadShrink  *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListHotelOrderShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotelOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHotelOrderShrinkRequest) SetPayloadShrink(v string) *ListHotelOrderShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *ListHotelOrderShrinkRequest) SetUserInfoShrink(v string) *ListHotelOrderShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type ListHotelOrderResponseBody struct {
	Code      *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	Page      *ListHotelOrderResponseBodyPage     `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListHotelOrderResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListHotelOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotelOrderResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotelOrderResponseBody) SetCode(v int32) *ListHotelOrderResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotelOrderResponseBody) SetMessage(v string) *ListHotelOrderResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotelOrderResponseBody) SetPage(v *ListHotelOrderResponseBodyPage) *ListHotelOrderResponseBody {
	s.Page = v
	return s
}

func (s *ListHotelOrderResponseBody) SetRequestId(v string) *ListHotelOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotelOrderResponseBody) SetResult(v []*ListHotelOrderResponseBodyResult) *ListHotelOrderResponseBody {
	s.Result = v
	return s
}

type ListHotelOrderResponseBodyPage struct {
	HasNext    *bool  `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total      *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	TotalPage  *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListHotelOrderResponseBodyPage) String() string {
	return tea.Prettify(s)
}

func (s ListHotelOrderResponseBodyPage) GoString() string {
	return s.String()
}

func (s *ListHotelOrderResponseBodyPage) SetHasNext(v bool) *ListHotelOrderResponseBodyPage {
	s.HasNext = &v
	return s
}

func (s *ListHotelOrderResponseBodyPage) SetPageNumber(v int32) *ListHotelOrderResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *ListHotelOrderResponseBodyPage) SetPageSize(v int32) *ListHotelOrderResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *ListHotelOrderResponseBodyPage) SetTotal(v int32) *ListHotelOrderResponseBodyPage {
	s.Total = &v
	return s
}

func (s *ListHotelOrderResponseBodyPage) SetTotalPage(v int32) *ListHotelOrderResponseBodyPage {
	s.TotalPage = &v
	return s
}

type ListHotelOrderResponseBodyResult struct {
	ApplyAmt    *int64  `json:"ApplyAmt,omitempty" xml:"ApplyAmt,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	OrderNo     *string `json:"OrderNo,omitempty" xml:"OrderNo,omitempty"`
	Quantity    *int64  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	RoomNo      *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	TypeIconUrl *string `json:"TypeIconUrl,omitempty" xml:"TypeIconUrl,omitempty"`
	TypeName    *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s ListHotelOrderResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListHotelOrderResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListHotelOrderResponseBodyResult) SetApplyAmt(v int64) *ListHotelOrderResponseBodyResult {
	s.ApplyAmt = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetGmtCreate(v int64) *ListHotelOrderResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetOrderNo(v string) *ListHotelOrderResponseBodyResult {
	s.OrderNo = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetQuantity(v int64) *ListHotelOrderResponseBodyResult {
	s.Quantity = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetRoomNo(v string) *ListHotelOrderResponseBodyResult {
	s.RoomNo = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetStatus(v string) *ListHotelOrderResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetType(v string) *ListHotelOrderResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetTypeIconUrl(v string) *ListHotelOrderResponseBodyResult {
	s.TypeIconUrl = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetTypeName(v string) *ListHotelOrderResponseBodyResult {
	s.TypeName = &v
	return s
}

type ListHotelOrderResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListHotelOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHotelOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotelOrderResponse) GoString() string {
	return s.String()
}

func (s *ListHotelOrderResponse) SetHeaders(v map[string]*string) *ListHotelOrderResponse {
	s.Headers = v
	return s
}

func (s *ListHotelOrderResponse) SetStatusCode(v int32) *ListHotelOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotelOrderResponse) SetBody(v *ListHotelOrderResponseBody) *ListHotelOrderResponse {
	s.Body = v
	return s
}

type ListHotelSceneItemHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListHotelSceneItemHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemHeaders) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemHeaders) SetCommonHeaders(v map[string]*string) *ListHotelSceneItemHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListHotelSceneItemHeaders) SetXAcsAligenieAccessToken(v string) *ListHotelSceneItemHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListHotelSceneItemHeaders) SetAuthorization(v string) *ListHotelSceneItemHeaders {
	s.Authorization = &v
	return s
}

type ListHotelSceneItemRequest struct {
	Payload  *ListHotelSceneItemRequestPayload  `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo *ListHotelSceneItemRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListHotelSceneItemRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemRequest) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemRequest) SetPayload(v *ListHotelSceneItemRequestPayload) *ListHotelSceneItemRequest {
	s.Payload = v
	return s
}

func (s *ListHotelSceneItemRequest) SetUserInfo(v *ListHotelSceneItemRequestUserInfo) *ListHotelSceneItemRequest {
	s.UserInfo = v
	return s
}

type ListHotelSceneItemRequestPayload struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListHotelSceneItemRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemRequestPayload) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemRequestPayload) SetType(v string) *ListHotelSceneItemRequestPayload {
	s.Type = &v
	return s
}

type ListHotelSceneItemRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListHotelSceneItemRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemRequestUserInfo) SetEncodeKey(v string) *ListHotelSceneItemRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListHotelSceneItemRequestUserInfo) SetEncodeType(v string) *ListHotelSceneItemRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListHotelSceneItemRequestUserInfo) SetId(v string) *ListHotelSceneItemRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListHotelSceneItemRequestUserInfo) SetIdType(v string) *ListHotelSceneItemRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListHotelSceneItemRequestUserInfo) SetOrganizationId(v string) *ListHotelSceneItemRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type ListHotelSceneItemShrinkRequest struct {
	PayloadShrink  *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListHotelSceneItemShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemShrinkRequest) SetPayloadShrink(v string) *ListHotelSceneItemShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *ListHotelSceneItemShrinkRequest) SetUserInfoShrink(v string) *ListHotelSceneItemShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type ListHotelSceneItemResponseBody struct {
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	Page      *ListHotelSceneItemResponseBodyPage   `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListHotelSceneItemResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListHotelSceneItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemResponseBody) SetCode(v int32) *ListHotelSceneItemResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotelSceneItemResponseBody) SetMessage(v string) *ListHotelSceneItemResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotelSceneItemResponseBody) SetPage(v *ListHotelSceneItemResponseBodyPage) *ListHotelSceneItemResponseBody {
	s.Page = v
	return s
}

func (s *ListHotelSceneItemResponseBody) SetRequestId(v string) *ListHotelSceneItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotelSceneItemResponseBody) SetResult(v *ListHotelSceneItemResponseBodyResult) *ListHotelSceneItemResponseBody {
	s.Result = v
	return s
}

type ListHotelSceneItemResponseBodyPage struct {
	HasNext    *bool  `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total      *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	TotalPage  *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListHotelSceneItemResponseBodyPage) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemResponseBodyPage) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemResponseBodyPage) SetHasNext(v bool) *ListHotelSceneItemResponseBodyPage {
	s.HasNext = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyPage) SetPageNumber(v int32) *ListHotelSceneItemResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyPage) SetPageSize(v int32) *ListHotelSceneItemResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyPage) SetTotal(v int32) *ListHotelSceneItemResponseBodyPage {
	s.Total = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyPage) SetTotalPage(v int32) *ListHotelSceneItemResponseBodyPage {
	s.TotalPage = &v
	return s
}

type ListHotelSceneItemResponseBodyResult struct {
	SecondCategoryList []*ListHotelSceneItemResponseBodyResultSecondCategoryList `json:"SecondCategoryList,omitempty" xml:"SecondCategoryList,omitempty" type:"Repeated"`
}

func (s ListHotelSceneItemResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemResponseBodyResult) SetSecondCategoryList(v []*ListHotelSceneItemResponseBodyResultSecondCategoryList) *ListHotelSceneItemResponseBodyResult {
	s.SecondCategoryList = v
	return s
}

type ListHotelSceneItemResponseBodyResultSecondCategoryList struct {
	ItemList           []*ListHotelSceneItemResponseBodyResultSecondCategoryListItemList `json:"ItemList,omitempty" xml:"ItemList,omitempty" type:"Repeated"`
	SecondCategoryName *string                                                           `json:"SecondCategoryName,omitempty" xml:"SecondCategoryName,omitempty"`
}

func (s ListHotelSceneItemResponseBodyResultSecondCategoryList) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemResponseBodyResultSecondCategoryList) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryList) SetItemList(v []*ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) *ListHotelSceneItemResponseBodyResultSecondCategoryList {
	s.ItemList = v
	return s
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryList) SetSecondCategoryName(v string) *ListHotelSceneItemResponseBodyResultSecondCategoryList {
	s.SecondCategoryName = &v
	return s
}

type ListHotelSceneItemResponseBodyResultSecondCategoryListItemList struct {
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Icon     *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Price    *int64  `json:"Price,omitempty" xml:"Price,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) SetCategory(v string) *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList {
	s.Category = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) SetIcon(v string) *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList {
	s.Icon = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) SetId(v string) *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList {
	s.Id = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) SetName(v string) *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList {
	s.Name = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) SetPrice(v int64) *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList {
	s.Price = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) SetStatus(v string) *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList {
	s.Status = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) SetType(v string) *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList {
	s.Type = &v
	return s
}

type ListHotelSceneItemResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListHotelSceneItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHotelSceneItemResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotelSceneItemResponse) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemResponse) SetHeaders(v map[string]*string) *ListHotelSceneItemResponse {
	s.Headers = v
	return s
}

func (s *ListHotelSceneItemResponse) SetStatusCode(v int32) *ListHotelSceneItemResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotelSceneItemResponse) SetBody(v *ListHotelSceneItemResponseBody) *ListHotelSceneItemResponse {
	s.Body = v
	return s
}

type ListHotelServiceCategoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListHotelServiceCategoryHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListHotelServiceCategoryHeaders) GoString() string {
	return s.String()
}

func (s *ListHotelServiceCategoryHeaders) SetCommonHeaders(v map[string]*string) *ListHotelServiceCategoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListHotelServiceCategoryHeaders) SetXAcsAligenieAccessToken(v string) *ListHotelServiceCategoryHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListHotelServiceCategoryHeaders) SetAuthorization(v string) *ListHotelServiceCategoryHeaders {
	s.Authorization = &v
	return s
}

type ListHotelServiceCategoryRequest struct {
	Payload *ListHotelServiceCategoryRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
}

func (s ListHotelServiceCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotelServiceCategoryRequest) GoString() string {
	return s.String()
}

func (s *ListHotelServiceCategoryRequest) SetPayload(v *ListHotelServiceCategoryRequestPayload) *ListHotelServiceCategoryRequest {
	s.Payload = v
	return s
}

type ListHotelServiceCategoryRequestPayload struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListHotelServiceCategoryRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s ListHotelServiceCategoryRequestPayload) GoString() string {
	return s.String()
}

func (s *ListHotelServiceCategoryRequestPayload) SetType(v string) *ListHotelServiceCategoryRequestPayload {
	s.Type = &v
	return s
}

type ListHotelServiceCategoryShrinkRequest struct {
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
}

func (s ListHotelServiceCategoryShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotelServiceCategoryShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHotelServiceCategoryShrinkRequest) SetPayloadShrink(v string) *ListHotelServiceCategoryShrinkRequest {
	s.PayloadShrink = &v
	return s
}

type ListHotelServiceCategoryResponseBody struct {
	Code      *int32                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListHotelServiceCategoryResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListHotelServiceCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotelServiceCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotelServiceCategoryResponseBody) SetCode(v int32) *ListHotelServiceCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotelServiceCategoryResponseBody) SetMessage(v string) *ListHotelServiceCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotelServiceCategoryResponseBody) SetRequestId(v string) *ListHotelServiceCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotelServiceCategoryResponseBody) SetResult(v []*ListHotelServiceCategoryResponseBodyResult) *ListHotelServiceCategoryResponseBody {
	s.Result = v
	return s
}

type ListHotelServiceCategoryResponseBodyResult struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListHotelServiceCategoryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListHotelServiceCategoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListHotelServiceCategoryResponseBodyResult) SetCode(v string) *ListHotelServiceCategoryResponseBodyResult {
	s.Code = &v
	return s
}

func (s *ListHotelServiceCategoryResponseBodyResult) SetDesc(v string) *ListHotelServiceCategoryResponseBodyResult {
	s.Desc = &v
	return s
}

func (s *ListHotelServiceCategoryResponseBodyResult) SetIcon(v string) *ListHotelServiceCategoryResponseBodyResult {
	s.Icon = &v
	return s
}

func (s *ListHotelServiceCategoryResponseBodyResult) SetName(v string) *ListHotelServiceCategoryResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListHotelServiceCategoryResponseBodyResult) SetType(v string) *ListHotelServiceCategoryResponseBodyResult {
	s.Type = &v
	return s
}

type ListHotelServiceCategoryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListHotelServiceCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHotelServiceCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotelServiceCategoryResponse) GoString() string {
	return s.String()
}

func (s *ListHotelServiceCategoryResponse) SetHeaders(v map[string]*string) *ListHotelServiceCategoryResponse {
	s.Headers = v
	return s
}

func (s *ListHotelServiceCategoryResponse) SetStatusCode(v int32) *ListHotelServiceCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotelServiceCategoryResponse) SetBody(v *ListHotelServiceCategoryResponseBody) *ListHotelServiceCategoryResponse {
	s.Body = v
	return s
}

type QueryDeviceStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s QueryDeviceStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatusHeaders) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusHeaders) SetCommonHeaders(v map[string]*string) *QueryDeviceStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDeviceStatusHeaders) SetXAcsAligenieAccessToken(v string) *QueryDeviceStatusHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *QueryDeviceStatusHeaders) SetAuthorization(v string) *QueryDeviceStatusHeaders {
	s.Authorization = &v
	return s
}

type QueryDeviceStatusRequest struct {
	Payload  *QueryDeviceStatusRequestPayload  `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo *QueryDeviceStatusRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s QueryDeviceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusRequest) SetPayload(v *QueryDeviceStatusRequestPayload) *QueryDeviceStatusRequest {
	s.Payload = v
	return s
}

func (s *QueryDeviceStatusRequest) SetUserInfo(v *QueryDeviceStatusRequestUserInfo) *QueryDeviceStatusRequest {
	s.UserInfo = v
	return s
}

type QueryDeviceStatusRequestPayload struct {
	LocationDevices []*QueryDeviceStatusRequestPayloadLocationDevices `json:"LocationDevices,omitempty" xml:"LocationDevices,omitempty" type:"Repeated"`
	Properties      map[string]*string                                `json:"Properties,omitempty" xml:"Properties,omitempty"`
}

func (s QueryDeviceStatusRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatusRequestPayload) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusRequestPayload) SetLocationDevices(v []*QueryDeviceStatusRequestPayloadLocationDevices) *QueryDeviceStatusRequestPayload {
	s.LocationDevices = v
	return s
}

func (s *QueryDeviceStatusRequestPayload) SetProperties(v map[string]*string) *QueryDeviceStatusRequestPayload {
	s.Properties = v
	return s
}

type QueryDeviceStatusRequestPayloadLocationDevices struct {
	DeviceNumber *string `json:"DeviceNumber,omitempty" xml:"DeviceNumber,omitempty"`
	DeviceType   *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	Location     *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s QueryDeviceStatusRequestPayloadLocationDevices) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatusRequestPayloadLocationDevices) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusRequestPayloadLocationDevices) SetDeviceNumber(v string) *QueryDeviceStatusRequestPayloadLocationDevices {
	s.DeviceNumber = &v
	return s
}

func (s *QueryDeviceStatusRequestPayloadLocationDevices) SetDeviceType(v string) *QueryDeviceStatusRequestPayloadLocationDevices {
	s.DeviceType = &v
	return s
}

func (s *QueryDeviceStatusRequestPayloadLocationDevices) SetLocation(v string) *QueryDeviceStatusRequestPayloadLocationDevices {
	s.Location = &v
	return s
}

type QueryDeviceStatusRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s QueryDeviceStatusRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatusRequestUserInfo) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusRequestUserInfo) SetEncodeKey(v string) *QueryDeviceStatusRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *QueryDeviceStatusRequestUserInfo) SetEncodeType(v string) *QueryDeviceStatusRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *QueryDeviceStatusRequestUserInfo) SetId(v string) *QueryDeviceStatusRequestUserInfo {
	s.Id = &v
	return s
}

func (s *QueryDeviceStatusRequestUserInfo) SetIdType(v string) *QueryDeviceStatusRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *QueryDeviceStatusRequestUserInfo) SetOrganizationId(v string) *QueryDeviceStatusRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type QueryDeviceStatusShrinkRequest struct {
	PayloadShrink  *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s QueryDeviceStatusShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusShrinkRequest) SetPayloadShrink(v string) *QueryDeviceStatusShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *QueryDeviceStatusShrinkRequest) SetUserInfoShrink(v string) *QueryDeviceStatusShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type QueryDeviceStatusResponseBody struct {
	Code      *int32               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []map[string]*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s QueryDeviceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusResponseBody) SetCode(v int32) *QueryDeviceStatusResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDeviceStatusResponseBody) SetMessage(v string) *QueryDeviceStatusResponseBody {
	s.Message = &v
	return s
}

func (s *QueryDeviceStatusResponseBody) SetRequestId(v string) *QueryDeviceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDeviceStatusResponseBody) SetResult(v []map[string]*string) *QueryDeviceStatusResponseBody {
	s.Result = v
	return s
}

type QueryDeviceStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDeviceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDeviceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusResponse) SetHeaders(v map[string]*string) *QueryDeviceStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceStatusResponse) SetStatusCode(v int32) *QueryDeviceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceStatusResponse) SetBody(v *QueryDeviceStatusResponseBody) *QueryDeviceStatusResponse {
	s.Body = v
	return s
}

type QueryHotelProductHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s QueryHotelProductHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryHotelProductHeaders) GoString() string {
	return s.String()
}

func (s *QueryHotelProductHeaders) SetCommonHeaders(v map[string]*string) *QueryHotelProductHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryHotelProductHeaders) SetXAcsAligenieAccessToken(v string) *QueryHotelProductHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *QueryHotelProductHeaders) SetAuthorization(v string) *QueryHotelProductHeaders {
	s.Authorization = &v
	return s
}

type QueryHotelProductRequest struct {
	UserInfo *QueryHotelProductRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s QueryHotelProductRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHotelProductRequest) GoString() string {
	return s.String()
}

func (s *QueryHotelProductRequest) SetUserInfo(v *QueryHotelProductRequestUserInfo) *QueryHotelProductRequest {
	s.UserInfo = v
	return s
}

type QueryHotelProductRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s QueryHotelProductRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryHotelProductRequestUserInfo) GoString() string {
	return s.String()
}

func (s *QueryHotelProductRequestUserInfo) SetEncodeKey(v string) *QueryHotelProductRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *QueryHotelProductRequestUserInfo) SetEncodeType(v string) *QueryHotelProductRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *QueryHotelProductRequestUserInfo) SetId(v string) *QueryHotelProductRequestUserInfo {
	s.Id = &v
	return s
}

func (s *QueryHotelProductRequestUserInfo) SetIdType(v string) *QueryHotelProductRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *QueryHotelProductRequestUserInfo) SetOrganizationId(v string) *QueryHotelProductRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type QueryHotelProductShrinkRequest struct {
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s QueryHotelProductShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHotelProductShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryHotelProductShrinkRequest) SetUserInfoShrink(v string) *QueryHotelProductShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type QueryHotelProductResponseBody struct {
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryHotelProductResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s QueryHotelProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHotelProductResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHotelProductResponseBody) SetCode(v int32) *QueryHotelProductResponseBody {
	s.Code = &v
	return s
}

func (s *QueryHotelProductResponseBody) SetMessage(v string) *QueryHotelProductResponseBody {
	s.Message = &v
	return s
}

func (s *QueryHotelProductResponseBody) SetRequestId(v string) *QueryHotelProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryHotelProductResponseBody) SetResult(v *QueryHotelProductResponseBodyResult) *QueryHotelProductResponseBody {
	s.Result = v
	return s
}

type QueryHotelProductResponseBodyResult struct {
	HotelId     *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	HotelName   *string `json:"HotelName,omitempty" xml:"HotelName,omitempty"`
	ProductKey  *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
}

func (s QueryHotelProductResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryHotelProductResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryHotelProductResponseBodyResult) SetHotelId(v string) *QueryHotelProductResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *QueryHotelProductResponseBodyResult) SetHotelName(v string) *QueryHotelProductResponseBodyResult {
	s.HotelName = &v
	return s
}

func (s *QueryHotelProductResponseBodyResult) SetProductKey(v string) *QueryHotelProductResponseBodyResult {
	s.ProductKey = &v
	return s
}

func (s *QueryHotelProductResponseBodyResult) SetProductName(v string) *QueryHotelProductResponseBodyResult {
	s.ProductName = &v
	return s
}

type QueryHotelProductResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryHotelProductResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryHotelProductResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHotelProductResponse) GoString() string {
	return s.String()
}

func (s *QueryHotelProductResponse) SetHeaders(v map[string]*string) *QueryHotelProductResponse {
	s.Headers = v
	return s
}

func (s *QueryHotelProductResponse) SetStatusCode(v int32) *QueryHotelProductResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHotelProductResponse) SetBody(v *QueryHotelProductResponseBody) *QueryHotelProductResponse {
	s.Body = v
	return s
}

type RoomCheckOutHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s RoomCheckOutHeaders) String() string {
	return tea.Prettify(s)
}

func (s RoomCheckOutHeaders) GoString() string {
	return s.String()
}

func (s *RoomCheckOutHeaders) SetCommonHeaders(v map[string]*string) *RoomCheckOutHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RoomCheckOutHeaders) SetXAcsAligenieAccessToken(v string) *RoomCheckOutHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *RoomCheckOutHeaders) SetAuthorization(v string) *RoomCheckOutHeaders {
	s.Authorization = &v
	return s
}

type RoomCheckOutRequest struct {
	DeviceInfo *RoomCheckOutRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	UserInfo   *RoomCheckOutRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s RoomCheckOutRequest) String() string {
	return tea.Prettify(s)
}

func (s RoomCheckOutRequest) GoString() string {
	return s.String()
}

func (s *RoomCheckOutRequest) SetDeviceInfo(v *RoomCheckOutRequestDeviceInfo) *RoomCheckOutRequest {
	s.DeviceInfo = v
	return s
}

func (s *RoomCheckOutRequest) SetUserInfo(v *RoomCheckOutRequestUserInfo) *RoomCheckOutRequest {
	s.UserInfo = v
	return s
}

type RoomCheckOutRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s RoomCheckOutRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s RoomCheckOutRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *RoomCheckOutRequestDeviceInfo) SetEncodeKey(v string) *RoomCheckOutRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *RoomCheckOutRequestDeviceInfo) SetEncodeType(v string) *RoomCheckOutRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *RoomCheckOutRequestDeviceInfo) SetId(v string) *RoomCheckOutRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *RoomCheckOutRequestDeviceInfo) SetIdType(v string) *RoomCheckOutRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *RoomCheckOutRequestDeviceInfo) SetOrganizationId(v string) *RoomCheckOutRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type RoomCheckOutRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s RoomCheckOutRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s RoomCheckOutRequestUserInfo) GoString() string {
	return s.String()
}

func (s *RoomCheckOutRequestUserInfo) SetEncodeKey(v string) *RoomCheckOutRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *RoomCheckOutRequestUserInfo) SetEncodeType(v string) *RoomCheckOutRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *RoomCheckOutRequestUserInfo) SetId(v string) *RoomCheckOutRequestUserInfo {
	s.Id = &v
	return s
}

func (s *RoomCheckOutRequestUserInfo) SetIdType(v string) *RoomCheckOutRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *RoomCheckOutRequestUserInfo) SetOrganizationId(v string) *RoomCheckOutRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type RoomCheckOutShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s RoomCheckOutShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RoomCheckOutShrinkRequest) GoString() string {
	return s.String()
}

func (s *RoomCheckOutShrinkRequest) SetDeviceInfoShrink(v string) *RoomCheckOutShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *RoomCheckOutShrinkRequest) SetUserInfoShrink(v string) *RoomCheckOutShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type RoomCheckOutResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s RoomCheckOutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RoomCheckOutResponseBody) GoString() string {
	return s.String()
}

func (s *RoomCheckOutResponseBody) SetCode(v int32) *RoomCheckOutResponseBody {
	s.Code = &v
	return s
}

func (s *RoomCheckOutResponseBody) SetMessage(v string) *RoomCheckOutResponseBody {
	s.Message = &v
	return s
}

func (s *RoomCheckOutResponseBody) SetRequestId(v string) *RoomCheckOutResponseBody {
	s.RequestId = &v
	return s
}

func (s *RoomCheckOutResponseBody) SetResult(v bool) *RoomCheckOutResponseBody {
	s.Result = &v
	return s
}

type RoomCheckOutResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RoomCheckOutResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RoomCheckOutResponse) String() string {
	return tea.Prettify(s)
}

func (s RoomCheckOutResponse) GoString() string {
	return s.String()
}

func (s *RoomCheckOutResponse) SetHeaders(v map[string]*string) *RoomCheckOutResponse {
	s.Headers = v
	return s
}

func (s *RoomCheckOutResponse) SetStatusCode(v int32) *RoomCheckOutResponse {
	s.StatusCode = &v
	return s
}

func (s *RoomCheckOutResponse) SetBody(v *RoomCheckOutResponseBody) *RoomCheckOutResponse {
	s.Body = v
	return s
}

type SubmitHotelOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s SubmitHotelOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotelOrderHeaders) GoString() string {
	return s.String()
}

func (s *SubmitHotelOrderHeaders) SetCommonHeaders(v map[string]*string) *SubmitHotelOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SubmitHotelOrderHeaders) SetXAcsAligenieAccessToken(v string) *SubmitHotelOrderHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *SubmitHotelOrderHeaders) SetAuthorization(v string) *SubmitHotelOrderHeaders {
	s.Authorization = &v
	return s
}

type SubmitHotelOrderRequest struct {
	Payload  *SubmitHotelOrderRequestPayload  `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo *SubmitHotelOrderRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s SubmitHotelOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotelOrderRequest) GoString() string {
	return s.String()
}

func (s *SubmitHotelOrderRequest) SetPayload(v *SubmitHotelOrderRequestPayload) *SubmitHotelOrderRequest {
	s.Payload = v
	return s
}

func (s *SubmitHotelOrderRequest) SetUserInfo(v *SubmitHotelOrderRequestUserInfo) *SubmitHotelOrderRequest {
	s.UserInfo = v
	return s
}

type SubmitHotelOrderRequestPayload struct {
	ItemList []*SubmitHotelOrderRequestPayloadItemList `json:"ItemList,omitempty" xml:"ItemList,omitempty" type:"Repeated"`
	Type     *string                                   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitHotelOrderRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotelOrderRequestPayload) GoString() string {
	return s.String()
}

func (s *SubmitHotelOrderRequestPayload) SetItemList(v []*SubmitHotelOrderRequestPayloadItemList) *SubmitHotelOrderRequestPayload {
	s.ItemList = v
	return s
}

func (s *SubmitHotelOrderRequestPayload) SetType(v string) *SubmitHotelOrderRequestPayload {
	s.Type = &v
	return s
}

type SubmitHotelOrderRequestPayloadItemList struct {
	ItemId   *int64 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	Quantity *int64 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
}

func (s SubmitHotelOrderRequestPayloadItemList) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotelOrderRequestPayloadItemList) GoString() string {
	return s.String()
}

func (s *SubmitHotelOrderRequestPayloadItemList) SetItemId(v int64) *SubmitHotelOrderRequestPayloadItemList {
	s.ItemId = &v
	return s
}

func (s *SubmitHotelOrderRequestPayloadItemList) SetQuantity(v int64) *SubmitHotelOrderRequestPayloadItemList {
	s.Quantity = &v
	return s
}

type SubmitHotelOrderRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s SubmitHotelOrderRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotelOrderRequestUserInfo) GoString() string {
	return s.String()
}

func (s *SubmitHotelOrderRequestUserInfo) SetEncodeKey(v string) *SubmitHotelOrderRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *SubmitHotelOrderRequestUserInfo) SetEncodeType(v string) *SubmitHotelOrderRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *SubmitHotelOrderRequestUserInfo) SetId(v string) *SubmitHotelOrderRequestUserInfo {
	s.Id = &v
	return s
}

func (s *SubmitHotelOrderRequestUserInfo) SetIdType(v string) *SubmitHotelOrderRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *SubmitHotelOrderRequestUserInfo) SetOrganizationId(v string) *SubmitHotelOrderRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type SubmitHotelOrderShrinkRequest struct {
	PayloadShrink  *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s SubmitHotelOrderShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotelOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitHotelOrderShrinkRequest) SetPayloadShrink(v string) *SubmitHotelOrderShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *SubmitHotelOrderShrinkRequest) SetUserInfoShrink(v string) *SubmitHotelOrderShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type SubmitHotelOrderResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s SubmitHotelOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotelOrderResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitHotelOrderResponseBody) SetCode(v int32) *SubmitHotelOrderResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitHotelOrderResponseBody) SetMessage(v string) *SubmitHotelOrderResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitHotelOrderResponseBody) SetRequestId(v string) *SubmitHotelOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitHotelOrderResponseBody) SetResult(v string) *SubmitHotelOrderResponseBody {
	s.Result = &v
	return s
}

type SubmitHotelOrderResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitHotelOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitHotelOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotelOrderResponse) GoString() string {
	return s.String()
}

func (s *SubmitHotelOrderResponse) SetHeaders(v map[string]*string) *SubmitHotelOrderResponse {
	s.Headers = v
	return s
}

func (s *SubmitHotelOrderResponse) SetStatusCode(v int32) *SubmitHotelOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitHotelOrderResponse) SetBody(v *SubmitHotelOrderResponseBody) *SubmitHotelOrderResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("aligenie"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) DeviceControl(request *DeviceControlRequest) (_result *DeviceControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeviceControlHeaders{}
	_result = &DeviceControlResponse{}
	_body, _err := client.DeviceControlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeviceControlWithOptions(tmpReq *DeviceControlRequest, headers *DeviceControlHeaders, runtime *util.RuntimeOptions) (_result *DeviceControlResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeviceControlShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Payload))) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Payload), tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		query["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeviceControl"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/deviceControl"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeviceControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotelHomeBackImageAndModes(request *GetHotelHomeBackImageAndModesRequest) (_result *GetHotelHomeBackImageAndModesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetHotelHomeBackImageAndModesHeaders{}
	_result = &GetHotelHomeBackImageAndModesResponse{}
	_body, _err := client.GetHotelHomeBackImageAndModesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotelHomeBackImageAndModesWithOptions(tmpReq *GetHotelHomeBackImageAndModesRequest, headers *GetHotelHomeBackImageAndModesHeaders, runtime *util.RuntimeOptions) (_result *GetHotelHomeBackImageAndModesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetHotelHomeBackImageAndModesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotelHomeBackImageAndModes"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/getHotelHomeBackImageAndModes"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotelHomeBackImageAndModesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotelOrderDetail(request *GetHotelOrderDetailRequest) (_result *GetHotelOrderDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetHotelOrderDetailHeaders{}
	_result = &GetHotelOrderDetailResponse{}
	_body, _err := client.GetHotelOrderDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotelOrderDetailWithOptions(tmpReq *GetHotelOrderDetailRequest, headers *GetHotelOrderDetailHeaders, runtime *util.RuntimeOptions) (_result *GetHotelOrderDetailResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetHotelOrderDetailShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Payload))) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Payload), tea.String("Payload"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		query["Payload"] = request.PayloadShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotelOrderDetail"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/getHotelOrderDetail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotelOrderDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotelSampleUtterances(request *GetHotelSampleUtterancesRequest) (_result *GetHotelSampleUtterancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetHotelSampleUtterancesHeaders{}
	_result = &GetHotelSampleUtterancesResponse{}
	_body, _err := client.GetHotelSampleUtterancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotelSampleUtterancesWithOptions(tmpReq *GetHotelSampleUtterancesRequest, headers *GetHotelSampleUtterancesHeaders, runtime *util.RuntimeOptions) (_result *GetHotelSampleUtterancesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetHotelSampleUtterancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotelSampleUtterances"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/getHotelSampleUtterances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotelSampleUtterancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotelScreenSaver(request *GetHotelScreenSaverRequest) (_result *GetHotelScreenSaverResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetHotelScreenSaverHeaders{}
	_result = &GetHotelScreenSaverResponse{}
	_body, _err := client.GetHotelScreenSaverWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotelScreenSaverWithOptions(tmpReq *GetHotelScreenSaverRequest, headers *GetHotelScreenSaverHeaders, runtime *util.RuntimeOptions) (_result *GetHotelScreenSaverResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetHotelScreenSaverShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotelScreenSaver"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/getHotelScreenSaver"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotelScreenSaverResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHotelControlDevice(request *ListHotelControlDeviceRequest) (_result *ListHotelControlDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListHotelControlDeviceHeaders{}
	_result = &ListHotelControlDeviceResponse{}
	_body, _err := client.ListHotelControlDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHotelControlDeviceWithOptions(tmpReq *ListHotelControlDeviceRequest, headers *ListHotelControlDeviceHeaders, runtime *util.RuntimeOptions) (_result *ListHotelControlDeviceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListHotelControlDeviceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotelControlDevice"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/listHotelControlDevice"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHotelControlDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHotelOrder(request *ListHotelOrderRequest) (_result *ListHotelOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListHotelOrderHeaders{}
	_result = &ListHotelOrderResponse{}
	_body, _err := client.ListHotelOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHotelOrderWithOptions(tmpReq *ListHotelOrderRequest, headers *ListHotelOrderHeaders, runtime *util.RuntimeOptions) (_result *ListHotelOrderResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListHotelOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Payload))) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Payload), tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		query["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotelOrder"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/listHotelOrder"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHotelOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHotelSceneItem(request *ListHotelSceneItemRequest) (_result *ListHotelSceneItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListHotelSceneItemHeaders{}
	_result = &ListHotelSceneItemResponse{}
	_body, _err := client.ListHotelSceneItemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHotelSceneItemWithOptions(tmpReq *ListHotelSceneItemRequest, headers *ListHotelSceneItemHeaders, runtime *util.RuntimeOptions) (_result *ListHotelSceneItemResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListHotelSceneItemShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Payload))) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Payload), tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		query["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotelSceneItem"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/listHotelSceneItem"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHotelSceneItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHotelServiceCategory(request *ListHotelServiceCategoryRequest) (_result *ListHotelServiceCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListHotelServiceCategoryHeaders{}
	_result = &ListHotelServiceCategoryResponse{}
	_body, _err := client.ListHotelServiceCategoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHotelServiceCategoryWithOptions(tmpReq *ListHotelServiceCategoryRequest, headers *ListHotelServiceCategoryHeaders, runtime *util.RuntimeOptions) (_result *ListHotelServiceCategoryResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListHotelServiceCategoryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Payload))) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Payload), tea.String("Payload"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		query["Payload"] = request.PayloadShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotelServiceCategory"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/listHotelServiceCategory"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHotelServiceCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDeviceStatus(request *QueryDeviceStatusRequest) (_result *QueryDeviceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDeviceStatusHeaders{}
	_result = &QueryDeviceStatusResponse{}
	_body, _err := client.QueryDeviceStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDeviceStatusWithOptions(tmpReq *QueryDeviceStatusRequest, headers *QueryDeviceStatusHeaders, runtime *util.RuntimeOptions) (_result *QueryDeviceStatusResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryDeviceStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Payload))) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Payload), tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		query["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDeviceStatus"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/queryDeviceStatus"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDeviceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryHotelProduct(request *QueryHotelProductRequest) (_result *QueryHotelProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryHotelProductHeaders{}
	_result = &QueryHotelProductResponse{}
	_body, _err := client.QueryHotelProductWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryHotelProductWithOptions(tmpReq *QueryHotelProductRequest, headers *QueryHotelProductHeaders, runtime *util.RuntimeOptions) (_result *QueryHotelProductResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryHotelProductShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryHotelProduct"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/queryHotelProduct"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryHotelProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RoomCheckOut(request *RoomCheckOutRequest) (_result *RoomCheckOutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RoomCheckOutHeaders{}
	_result = &RoomCheckOutResponse{}
	_body, _err := client.RoomCheckOutWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RoomCheckOutWithOptions(tmpReq *RoomCheckOutRequest, headers *RoomCheckOutHeaders, runtime *util.RuntimeOptions) (_result *RoomCheckOutResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RoomCheckOutShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DeviceInfo))) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DeviceInfo), tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RoomCheckOut"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/roomCheckOut"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RoomCheckOutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitHotelOrder(request *SubmitHotelOrderRequest) (_result *SubmitHotelOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SubmitHotelOrderHeaders{}
	_result = &SubmitHotelOrderResponse{}
	_body, _err := client.SubmitHotelOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitHotelOrderWithOptions(tmpReq *SubmitHotelOrderRequest, headers *SubmitHotelOrderHeaders, runtime *util.RuntimeOptions) (_result *SubmitHotelOrderResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SubmitHotelOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Payload))) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Payload), tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		query["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitHotelOrder"),
		Version:     tea.String("ip_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/ip/submitHotelOrder"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitHotelOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
