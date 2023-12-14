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

type AddMemberBasicInfoRequest struct {
	Body *AddMemberBasicInfoRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s AddMemberBasicInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s AddMemberBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *AddMemberBasicInfoRequest) SetBody(v *AddMemberBasicInfoRequestBody) *AddMemberBasicInfoRequest {
	s.Body = v
	return s
}

type AddMemberBasicInfoRequestBody struct {
	Area           *string                                  `json:"Area,omitempty" xml:"Area,omitempty"`
	Avatar         *string                                  `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	Birthday       *string                                  `json:"Birthday,omitempty" xml:"Birthday,omitempty"`
	Channels       []*AddMemberBasicInfoRequestBodyChannels `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
	City           *string                                  `json:"City,omitempty" xml:"City,omitempty"`
	Country        *string                                  `json:"Country,omitempty" xml:"Country,omitempty"`
	Email          *string                                  `json:"Email,omitempty" xml:"Email,omitempty"`
	Extra          *string                                  `json:"Extra,omitempty" xml:"Extra,omitempty"`
	GmtCreate      *string                                  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	MemberName     *string                                  `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	MemberNickName *string                                  `json:"MemberNickName,omitempty" xml:"MemberNickName,omitempty"`
	Mobile         *string                                  `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	OpenMerchantId *string                                  `json:"OpenMerchantId,omitempty" xml:"OpenMerchantId,omitempty"`
	PlatFormType   *string                                  `json:"PlatFormType,omitempty" xml:"PlatFormType,omitempty"`
	Province       *string                                  `json:"Province,omitempty" xml:"Province,omitempty"`
	Sex            *string                                  `json:"Sex,omitempty" xml:"Sex,omitempty"`
}

func (s AddMemberBasicInfoRequestBody) String() string {
	return tea.Prettify(s)
}

func (s AddMemberBasicInfoRequestBody) GoString() string {
	return s.String()
}

func (s *AddMemberBasicInfoRequestBody) SetArea(v string) *AddMemberBasicInfoRequestBody {
	s.Area = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetAvatar(v string) *AddMemberBasicInfoRequestBody {
	s.Avatar = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetBirthday(v string) *AddMemberBasicInfoRequestBody {
	s.Birthday = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetChannels(v []*AddMemberBasicInfoRequestBodyChannels) *AddMemberBasicInfoRequestBody {
	s.Channels = v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetCity(v string) *AddMemberBasicInfoRequestBody {
	s.City = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetCountry(v string) *AddMemberBasicInfoRequestBody {
	s.Country = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetEmail(v string) *AddMemberBasicInfoRequestBody {
	s.Email = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetExtra(v string) *AddMemberBasicInfoRequestBody {
	s.Extra = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetGmtCreate(v string) *AddMemberBasicInfoRequestBody {
	s.GmtCreate = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetMemberName(v string) *AddMemberBasicInfoRequestBody {
	s.MemberName = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetMemberNickName(v string) *AddMemberBasicInfoRequestBody {
	s.MemberNickName = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetMobile(v string) *AddMemberBasicInfoRequestBody {
	s.Mobile = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetOpenMerchantId(v string) *AddMemberBasicInfoRequestBody {
	s.OpenMerchantId = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetPlatFormType(v string) *AddMemberBasicInfoRequestBody {
	s.PlatFormType = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetProvince(v string) *AddMemberBasicInfoRequestBody {
	s.Province = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetSex(v string) *AddMemberBasicInfoRequestBody {
	s.Sex = &v
	return s
}

type AddMemberBasicInfoRequestBodyChannels struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelCode    *string `json:"ChannelCode,omitempty" xml:"ChannelCode,omitempty"`
	ChannelOpenId  *string `json:"ChannelOpenId,omitempty" xml:"ChannelOpenId,omitempty"`
	ChannelUnionId *string `json:"ChannelUnionId,omitempty" xml:"ChannelUnionId,omitempty"`
	Scene          *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s AddMemberBasicInfoRequestBodyChannels) String() string {
	return tea.Prettify(s)
}

func (s AddMemberBasicInfoRequestBodyChannels) GoString() string {
	return s.String()
}

func (s *AddMemberBasicInfoRequestBodyChannels) SetAppId(v string) *AddMemberBasicInfoRequestBodyChannels {
	s.AppId = &v
	return s
}

func (s *AddMemberBasicInfoRequestBodyChannels) SetChannelCode(v string) *AddMemberBasicInfoRequestBodyChannels {
	s.ChannelCode = &v
	return s
}

func (s *AddMemberBasicInfoRequestBodyChannels) SetChannelOpenId(v string) *AddMemberBasicInfoRequestBodyChannels {
	s.ChannelOpenId = &v
	return s
}

func (s *AddMemberBasicInfoRequestBodyChannels) SetChannelUnionId(v string) *AddMemberBasicInfoRequestBodyChannels {
	s.ChannelUnionId = &v
	return s
}

func (s *AddMemberBasicInfoRequestBodyChannels) SetScene(v string) *AddMemberBasicInfoRequestBodyChannels {
	s.Scene = &v
	return s
}

type AddMemberBasicInfoShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMemberBasicInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddMemberBasicInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddMemberBasicInfoShrinkRequest) SetBodyShrink(v string) *AddMemberBasicInfoShrinkRequest {
	s.BodyShrink = &v
	return s
}

type AddMemberBasicInfoResponseBody struct {
	ErrorCode     *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage  *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	OuterMemberId *string `json:"OuterMemberId,omitempty" xml:"OuterMemberId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success       *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddMemberBasicInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddMemberBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *AddMemberBasicInfoResponseBody) SetErrorCode(v string) *AddMemberBasicInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddMemberBasicInfoResponseBody) SetErrorMessage(v string) *AddMemberBasicInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddMemberBasicInfoResponseBody) SetOuterMemberId(v string) *AddMemberBasicInfoResponseBody {
	s.OuterMemberId = &v
	return s
}

func (s *AddMemberBasicInfoResponseBody) SetRequestId(v string) *AddMemberBasicInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMemberBasicInfoResponseBody) SetSuccess(v bool) *AddMemberBasicInfoResponseBody {
	s.Success = &v
	return s
}

type AddMemberBasicInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddMemberBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddMemberBasicInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s AddMemberBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *AddMemberBasicInfoResponse) SetHeaders(v map[string]*string) *AddMemberBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *AddMemberBasicInfoResponse) SetStatusCode(v int32) *AddMemberBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMemberBasicInfoResponse) SetBody(v *AddMemberBasicInfoResponseBody) *AddMemberBasicInfoResponse {
	s.Body = v
	return s
}

type MemberAccountDetailPageQueryRequest struct {
	Body *MemberAccountDetailPageQueryRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s MemberAccountDetailPageQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s MemberAccountDetailPageQueryRequest) GoString() string {
	return s.String()
}

func (s *MemberAccountDetailPageQueryRequest) SetBody(v *MemberAccountDetailPageQueryRequestBody) *MemberAccountDetailPageQueryRequest {
	s.Body = v
	return s
}

type MemberAccountDetailPageQueryRequestBody struct {
	AccountType    *int32  `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OpenMerchantId *string `json:"OpenMerchantId,omitempty" xml:"OpenMerchantId,omitempty"`
	OuterMemberId  *string `json:"OuterMemberId,omitempty" xml:"OuterMemberId,omitempty"`
	Page           *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PlatFormType   *string `json:"PlatFormType,omitempty" xml:"PlatFormType,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s MemberAccountDetailPageQueryRequestBody) String() string {
	return tea.Prettify(s)
}

func (s MemberAccountDetailPageQueryRequestBody) GoString() string {
	return s.String()
}

func (s *MemberAccountDetailPageQueryRequestBody) SetAccountType(v int32) *MemberAccountDetailPageQueryRequestBody {
	s.AccountType = &v
	return s
}

func (s *MemberAccountDetailPageQueryRequestBody) SetEndTime(v string) *MemberAccountDetailPageQueryRequestBody {
	s.EndTime = &v
	return s
}

func (s *MemberAccountDetailPageQueryRequestBody) SetOpenMerchantId(v string) *MemberAccountDetailPageQueryRequestBody {
	s.OpenMerchantId = &v
	return s
}

func (s *MemberAccountDetailPageQueryRequestBody) SetOuterMemberId(v string) *MemberAccountDetailPageQueryRequestBody {
	s.OuterMemberId = &v
	return s
}

func (s *MemberAccountDetailPageQueryRequestBody) SetPage(v int32) *MemberAccountDetailPageQueryRequestBody {
	s.Page = &v
	return s
}

func (s *MemberAccountDetailPageQueryRequestBody) SetPageSize(v int32) *MemberAccountDetailPageQueryRequestBody {
	s.PageSize = &v
	return s
}

func (s *MemberAccountDetailPageQueryRequestBody) SetPlatFormType(v string) *MemberAccountDetailPageQueryRequestBody {
	s.PlatFormType = &v
	return s
}

func (s *MemberAccountDetailPageQueryRequestBody) SetStartTime(v string) *MemberAccountDetailPageQueryRequestBody {
	s.StartTime = &v
	return s
}

type MemberAccountDetailPageQueryShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MemberAccountDetailPageQueryShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s MemberAccountDetailPageQueryShrinkRequest) GoString() string {
	return s.String()
}

func (s *MemberAccountDetailPageQueryShrinkRequest) SetBodyShrink(v string) *MemberAccountDetailPageQueryShrinkRequest {
	s.BodyShrink = &v
	return s
}

type MemberAccountDetailPageQueryResponseBody struct {
	Data         []*MemberAccountDetailPageQueryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode    *string                                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s MemberAccountDetailPageQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MemberAccountDetailPageQueryResponseBody) GoString() string {
	return s.String()
}

func (s *MemberAccountDetailPageQueryResponseBody) SetData(v []*MemberAccountDetailPageQueryResponseBodyData) *MemberAccountDetailPageQueryResponseBody {
	s.Data = v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBody) SetErrorCode(v string) *MemberAccountDetailPageQueryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBody) SetErrorMessage(v string) *MemberAccountDetailPageQueryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBody) SetRequestId(v string) *MemberAccountDetailPageQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBody) SetSuccess(v string) *MemberAccountDetailPageQueryResponseBody {
	s.Success = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBody) SetTotalCount(v int32) *MemberAccountDetailPageQueryResponseBody {
	s.TotalCount = &v
	return s
}

type MemberAccountDetailPageQueryResponseBodyData struct {
	AccountBalance *string `json:"AccountBalance,omitempty" xml:"AccountBalance,omitempty"`
	AccountType    *int32  `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	ActivityType   *string `json:"ActivityType,omitempty" xml:"ActivityType,omitempty"`
	ChannelCode    *string `json:"ChannelCode,omitempty" xml:"ChannelCode,omitempty"`
	DetailValue    *string `json:"DetailValue,omitempty" xml:"DetailValue,omitempty"`
	Extra          *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	GmtCreate      *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified    *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	OpenMerchantId *string `json:"OpenMerchantId,omitempty" xml:"OpenMerchantId,omitempty"`
	OperateType    *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	OuterMemberId  *string `json:"OuterMemberId,omitempty" xml:"OuterMemberId,omitempty"`
	Remark         *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s MemberAccountDetailPageQueryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s MemberAccountDetailPageQueryResponseBodyData) GoString() string {
	return s.String()
}

func (s *MemberAccountDetailPageQueryResponseBodyData) SetAccountBalance(v string) *MemberAccountDetailPageQueryResponseBodyData {
	s.AccountBalance = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBodyData) SetAccountType(v int32) *MemberAccountDetailPageQueryResponseBodyData {
	s.AccountType = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBodyData) SetActivityType(v string) *MemberAccountDetailPageQueryResponseBodyData {
	s.ActivityType = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBodyData) SetChannelCode(v string) *MemberAccountDetailPageQueryResponseBodyData {
	s.ChannelCode = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBodyData) SetDetailValue(v string) *MemberAccountDetailPageQueryResponseBodyData {
	s.DetailValue = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBodyData) SetExtra(v string) *MemberAccountDetailPageQueryResponseBodyData {
	s.Extra = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBodyData) SetGmtCreate(v string) *MemberAccountDetailPageQueryResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBodyData) SetGmtModified(v string) *MemberAccountDetailPageQueryResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBodyData) SetOpenMerchantId(v string) *MemberAccountDetailPageQueryResponseBodyData {
	s.OpenMerchantId = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBodyData) SetOperateType(v string) *MemberAccountDetailPageQueryResponseBodyData {
	s.OperateType = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBodyData) SetOuterMemberId(v string) *MemberAccountDetailPageQueryResponseBodyData {
	s.OuterMemberId = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBodyData) SetRemark(v string) *MemberAccountDetailPageQueryResponseBodyData {
	s.Remark = &v
	return s
}

type MemberAccountDetailPageQueryResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MemberAccountDetailPageQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MemberAccountDetailPageQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s MemberAccountDetailPageQueryResponse) GoString() string {
	return s.String()
}

func (s *MemberAccountDetailPageQueryResponse) SetHeaders(v map[string]*string) *MemberAccountDetailPageQueryResponse {
	s.Headers = v
	return s
}

func (s *MemberAccountDetailPageQueryResponse) SetStatusCode(v int32) *MemberAccountDetailPageQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponse) SetBody(v *MemberAccountDetailPageQueryResponseBody) *MemberAccountDetailPageQueryResponse {
	s.Body = v
	return s
}

type MemberPointChangeRequest struct {
	Body *MemberPointChangeRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s MemberPointChangeRequest) String() string {
	return tea.Prettify(s)
}

func (s MemberPointChangeRequest) GoString() string {
	return s.String()
}

func (s *MemberPointChangeRequest) SetBody(v *MemberPointChangeRequestBody) *MemberPointChangeRequest {
	s.Body = v
	return s
}

type MemberPointChangeRequestBody struct {
	AccountType    *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	ChannelCode    *string `json:"ChannelCode,omitempty" xml:"ChannelCode,omitempty"`
	Extra          *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	OpenMerchantId *string `json:"OpenMerchantId,omitempty" xml:"OpenMerchantId,omitempty"`
	OperateType    *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	OuterMemberId  *string `json:"OuterMemberId,omitempty" xml:"OuterMemberId,omitempty"`
	PlatFormType   *string `json:"PlatFormType,omitempty" xml:"PlatFormType,omitempty"`
	Quantity       *string `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	SerialNo       *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
}

func (s MemberPointChangeRequestBody) String() string {
	return tea.Prettify(s)
}

func (s MemberPointChangeRequestBody) GoString() string {
	return s.String()
}

func (s *MemberPointChangeRequestBody) SetAccountType(v string) *MemberPointChangeRequestBody {
	s.AccountType = &v
	return s
}

func (s *MemberPointChangeRequestBody) SetChannelCode(v string) *MemberPointChangeRequestBody {
	s.ChannelCode = &v
	return s
}

func (s *MemberPointChangeRequestBody) SetExtra(v string) *MemberPointChangeRequestBody {
	s.Extra = &v
	return s
}

func (s *MemberPointChangeRequestBody) SetOpenMerchantId(v string) *MemberPointChangeRequestBody {
	s.OpenMerchantId = &v
	return s
}

func (s *MemberPointChangeRequestBody) SetOperateType(v string) *MemberPointChangeRequestBody {
	s.OperateType = &v
	return s
}

func (s *MemberPointChangeRequestBody) SetOuterMemberId(v string) *MemberPointChangeRequestBody {
	s.OuterMemberId = &v
	return s
}

func (s *MemberPointChangeRequestBody) SetPlatFormType(v string) *MemberPointChangeRequestBody {
	s.PlatFormType = &v
	return s
}

func (s *MemberPointChangeRequestBody) SetQuantity(v string) *MemberPointChangeRequestBody {
	s.Quantity = &v
	return s
}

func (s *MemberPointChangeRequestBody) SetSerialNo(v string) *MemberPointChangeRequestBody {
	s.SerialNo = &v
	return s
}

type MemberPointChangeShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MemberPointChangeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s MemberPointChangeShrinkRequest) GoString() string {
	return s.String()
}

func (s *MemberPointChangeShrinkRequest) SetBodyShrink(v string) *MemberPointChangeShrinkRequest {
	s.BodyShrink = &v
	return s
}

type MemberPointChangeResponseBody struct {
	AccountBalance *string `json:"AccountBalance,omitempty" xml:"AccountBalance,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LevelName      *string `json:"LevelName,omitempty" xml:"LevelName,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MemberPointChangeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MemberPointChangeResponseBody) GoString() string {
	return s.String()
}

func (s *MemberPointChangeResponseBody) SetAccountBalance(v string) *MemberPointChangeResponseBody {
	s.AccountBalance = &v
	return s
}

func (s *MemberPointChangeResponseBody) SetErrorCode(v string) *MemberPointChangeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *MemberPointChangeResponseBody) SetErrorMessage(v string) *MemberPointChangeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *MemberPointChangeResponseBody) SetLevelName(v string) *MemberPointChangeResponseBody {
	s.LevelName = &v
	return s
}

func (s *MemberPointChangeResponseBody) SetRequestId(v string) *MemberPointChangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *MemberPointChangeResponseBody) SetSuccess(v string) *MemberPointChangeResponseBody {
	s.Success = &v
	return s
}

type MemberPointChangeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MemberPointChangeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MemberPointChangeResponse) String() string {
	return tea.Prettify(s)
}

func (s MemberPointChangeResponse) GoString() string {
	return s.String()
}

func (s *MemberPointChangeResponse) SetHeaders(v map[string]*string) *MemberPointChangeResponse {
	s.Headers = v
	return s
}

func (s *MemberPointChangeResponse) SetStatusCode(v int32) *MemberPointChangeResponse {
	s.StatusCode = &v
	return s
}

func (s *MemberPointChangeResponse) SetBody(v *MemberPointChangeResponseBody) *MemberPointChangeResponse {
	s.Body = v
	return s
}

type QueryMemberBasicInfoRequest struct {
	Body *QueryMemberBasicInfoRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s QueryMemberBasicInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMemberBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryMemberBasicInfoRequest) SetBody(v *QueryMemberBasicInfoRequestBody) *QueryMemberBasicInfoRequest {
	s.Body = v
	return s
}

type QueryMemberBasicInfoRequestBody struct {
	ChannelCode    *string `json:"ChannelCode,omitempty" xml:"ChannelCode,omitempty"`
	ChannelOpenId  *string `json:"ChannelOpenId,omitempty" xml:"ChannelOpenId,omitempty"`
	Mobile         *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	OpenMerchantId *string `json:"OpenMerchantId,omitempty" xml:"OpenMerchantId,omitempty"`
	OuterMemberId  *string `json:"OuterMemberId,omitempty" xml:"OuterMemberId,omitempty"`
	PlatFormType   *string `json:"PlatFormType,omitempty" xml:"PlatFormType,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryMemberBasicInfoRequestBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMemberBasicInfoRequestBody) GoString() string {
	return s.String()
}

func (s *QueryMemberBasicInfoRequestBody) SetChannelCode(v string) *QueryMemberBasicInfoRequestBody {
	s.ChannelCode = &v
	return s
}

func (s *QueryMemberBasicInfoRequestBody) SetChannelOpenId(v string) *QueryMemberBasicInfoRequestBody {
	s.ChannelOpenId = &v
	return s
}

func (s *QueryMemberBasicInfoRequestBody) SetMobile(v string) *QueryMemberBasicInfoRequestBody {
	s.Mobile = &v
	return s
}

func (s *QueryMemberBasicInfoRequestBody) SetOpenMerchantId(v string) *QueryMemberBasicInfoRequestBody {
	s.OpenMerchantId = &v
	return s
}

func (s *QueryMemberBasicInfoRequestBody) SetOuterMemberId(v string) *QueryMemberBasicInfoRequestBody {
	s.OuterMemberId = &v
	return s
}

func (s *QueryMemberBasicInfoRequestBody) SetPlatFormType(v string) *QueryMemberBasicInfoRequestBody {
	s.PlatFormType = &v
	return s
}

func (s *QueryMemberBasicInfoRequestBody) SetType(v string) *QueryMemberBasicInfoRequestBody {
	s.Type = &v
	return s
}

type QueryMemberBasicInfoShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMemberBasicInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMemberBasicInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryMemberBasicInfoShrinkRequest) SetBodyShrink(v string) *QueryMemberBasicInfoShrinkRequest {
	s.BodyShrink = &v
	return s
}

type QueryMemberBasicInfoResponseBody struct {
	Data         *QueryMemberBasicInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode    *string                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                               `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMemberBasicInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMemberBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMemberBasicInfoResponseBody) SetData(v *QueryMemberBasicInfoResponseBodyData) *QueryMemberBasicInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryMemberBasicInfoResponseBody) SetErrorCode(v string) *QueryMemberBasicInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBody) SetErrorMessage(v string) *QueryMemberBasicInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBody) SetRequestId(v string) *QueryMemberBasicInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBody) SetSuccess(v bool) *QueryMemberBasicInfoResponseBody {
	s.Success = &v
	return s
}

type QueryMemberBasicInfoResponseBodyData struct {
	Area           *string `json:"Area,omitempty" xml:"Area,omitempty"`
	Avatar         *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	Birthday       *string `json:"Birthday,omitempty" xml:"Birthday,omitempty"`
	City           *string `json:"City,omitempty" xml:"City,omitempty"`
	Country        *string `json:"Country,omitempty" xml:"Country,omitempty"`
	Email          *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Extra          *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LevelName      *string `json:"LevelName,omitempty" xml:"LevelName,omitempty"`
	LevelNum       *string `json:"LevelNum,omitempty" xml:"LevelNum,omitempty"`
	MemberName     *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	MemberNickName *string `json:"MemberNickName,omitempty" xml:"MemberNickName,omitempty"`
	Mobile         *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	OpenMerchantId *string `json:"OpenMerchantId,omitempty" xml:"OpenMerchantId,omitempty"`
	OuterMemberId  *string `json:"OuterMemberId,omitempty" xml:"OuterMemberId,omitempty"`
	Points         *string `json:"Points,omitempty" xml:"Points,omitempty"`
	Province       *string `json:"Province,omitempty" xml:"Province,omitempty"`
	Score          *string `json:"Score,omitempty" xml:"Score,omitempty"`
	Sex            *string `json:"Sex,omitempty" xml:"Sex,omitempty"`
}

func (s QueryMemberBasicInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryMemberBasicInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMemberBasicInfoResponseBodyData) SetArea(v string) *QueryMemberBasicInfoResponseBodyData {
	s.Area = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetAvatar(v string) *QueryMemberBasicInfoResponseBodyData {
	s.Avatar = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetBirthday(v string) *QueryMemberBasicInfoResponseBodyData {
	s.Birthday = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetCity(v string) *QueryMemberBasicInfoResponseBodyData {
	s.City = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetCountry(v string) *QueryMemberBasicInfoResponseBodyData {
	s.Country = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetEmail(v string) *QueryMemberBasicInfoResponseBodyData {
	s.Email = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetExtra(v string) *QueryMemberBasicInfoResponseBodyData {
	s.Extra = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetLevelName(v string) *QueryMemberBasicInfoResponseBodyData {
	s.LevelName = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetLevelNum(v string) *QueryMemberBasicInfoResponseBodyData {
	s.LevelNum = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetMemberName(v string) *QueryMemberBasicInfoResponseBodyData {
	s.MemberName = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetMemberNickName(v string) *QueryMemberBasicInfoResponseBodyData {
	s.MemberNickName = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetMobile(v string) *QueryMemberBasicInfoResponseBodyData {
	s.Mobile = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetOpenMerchantId(v string) *QueryMemberBasicInfoResponseBodyData {
	s.OpenMerchantId = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetOuterMemberId(v string) *QueryMemberBasicInfoResponseBodyData {
	s.OuterMemberId = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetPoints(v string) *QueryMemberBasicInfoResponseBodyData {
	s.Points = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetProvince(v string) *QueryMemberBasicInfoResponseBodyData {
	s.Province = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetScore(v string) *QueryMemberBasicInfoResponseBodyData {
	s.Score = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetSex(v string) *QueryMemberBasicInfoResponseBodyData {
	s.Sex = &v
	return s
}

type QueryMemberBasicInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMemberBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMemberBasicInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMemberBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryMemberBasicInfoResponse) SetHeaders(v map[string]*string) *QueryMemberBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryMemberBasicInfoResponse) SetStatusCode(v int32) *QueryMemberBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMemberBasicInfoResponse) SetBody(v *QueryMemberBasicInfoResponseBody) *QueryMemberBasicInfoResponse {
	s.Body = v
	return s
}

type SyncCardInfoRequest struct {
	Body *SyncCardInfoRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s SyncCardInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncCardInfoRequest) GoString() string {
	return s.String()
}

func (s *SyncCardInfoRequest) SetBody(v *SyncCardInfoRequestBody) *SyncCardInfoRequest {
	s.Body = v
	return s
}

type SyncCardInfoRequestBody struct {
	BuyerId        *string                                 `json:"BuyerId,omitempty" xml:"BuyerId,omitempty"`
	Extra          *string                                 `json:"Extra,omitempty" xml:"Extra,omitempty"`
	Gifters        *SyncCardInfoRequestBodyGifters         `json:"Gifters,omitempty" xml:"Gifters,omitempty" type:"Struct"`
	OccurredAt     *string                                 `json:"OccurredAt,omitempty" xml:"OccurredAt,omitempty"`
	OpenMerchantId *string                                 `json:"OpenMerchantId,omitempty" xml:"OpenMerchantId,omitempty"`
	OrderId        *string                                 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PresentDetail  []*SyncCardInfoRequestBodyPresentDetail `json:"PresentDetail,omitempty" xml:"PresentDetail,omitempty" type:"Repeated"`
	ReceivedAt     *string                                 `json:"ReceivedAt,omitempty" xml:"ReceivedAt,omitempty"`
	Recipient      *SyncCardInfoRequestBodyRecipient       `json:"Recipient,omitempty" xml:"Recipient,omitempty" type:"Struct"`
	Status         *string                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	Theme          *string                                 `json:"Theme,omitempty" xml:"Theme,omitempty"`
	TransferId     *string                                 `json:"TransferId,omitempty" xml:"TransferId,omitempty"`
	TransferredAt  *string                                 `json:"TransferredAt,omitempty" xml:"TransferredAt,omitempty"`
}

func (s SyncCardInfoRequestBody) String() string {
	return tea.Prettify(s)
}

func (s SyncCardInfoRequestBody) GoString() string {
	return s.String()
}

func (s *SyncCardInfoRequestBody) SetBuyerId(v string) *SyncCardInfoRequestBody {
	s.BuyerId = &v
	return s
}

func (s *SyncCardInfoRequestBody) SetExtra(v string) *SyncCardInfoRequestBody {
	s.Extra = &v
	return s
}

func (s *SyncCardInfoRequestBody) SetGifters(v *SyncCardInfoRequestBodyGifters) *SyncCardInfoRequestBody {
	s.Gifters = v
	return s
}

func (s *SyncCardInfoRequestBody) SetOccurredAt(v string) *SyncCardInfoRequestBody {
	s.OccurredAt = &v
	return s
}

func (s *SyncCardInfoRequestBody) SetOpenMerchantId(v string) *SyncCardInfoRequestBody {
	s.OpenMerchantId = &v
	return s
}

func (s *SyncCardInfoRequestBody) SetOrderId(v string) *SyncCardInfoRequestBody {
	s.OrderId = &v
	return s
}

func (s *SyncCardInfoRequestBody) SetPresentDetail(v []*SyncCardInfoRequestBodyPresentDetail) *SyncCardInfoRequestBody {
	s.PresentDetail = v
	return s
}

func (s *SyncCardInfoRequestBody) SetReceivedAt(v string) *SyncCardInfoRequestBody {
	s.ReceivedAt = &v
	return s
}

func (s *SyncCardInfoRequestBody) SetRecipient(v *SyncCardInfoRequestBodyRecipient) *SyncCardInfoRequestBody {
	s.Recipient = v
	return s
}

func (s *SyncCardInfoRequestBody) SetStatus(v string) *SyncCardInfoRequestBody {
	s.Status = &v
	return s
}

func (s *SyncCardInfoRequestBody) SetTheme(v string) *SyncCardInfoRequestBody {
	s.Theme = &v
	return s
}

func (s *SyncCardInfoRequestBody) SetTransferId(v string) *SyncCardInfoRequestBody {
	s.TransferId = &v
	return s
}

func (s *SyncCardInfoRequestBody) SetTransferredAt(v string) *SyncCardInfoRequestBody {
	s.TransferredAt = &v
	return s
}

type SyncCardInfoRequestBodyGifters struct {
	HeaderImg *string `json:"HeaderImg,omitempty" xml:"HeaderImg,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Nickname  *string `json:"Nickname,omitempty" xml:"Nickname,omitempty"`
	// openId
	OpenId *string `json:"OpenId,omitempty" xml:"OpenId,omitempty"`
	Phone  *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s SyncCardInfoRequestBodyGifters) String() string {
	return tea.Prettify(s)
}

func (s SyncCardInfoRequestBodyGifters) GoString() string {
	return s.String()
}

func (s *SyncCardInfoRequestBodyGifters) SetHeaderImg(v string) *SyncCardInfoRequestBodyGifters {
	s.HeaderImg = &v
	return s
}

func (s *SyncCardInfoRequestBodyGifters) SetId(v string) *SyncCardInfoRequestBodyGifters {
	s.Id = &v
	return s
}

func (s *SyncCardInfoRequestBodyGifters) SetNickname(v string) *SyncCardInfoRequestBodyGifters {
	s.Nickname = &v
	return s
}

func (s *SyncCardInfoRequestBodyGifters) SetOpenId(v string) *SyncCardInfoRequestBodyGifters {
	s.OpenId = &v
	return s
}

func (s *SyncCardInfoRequestBodyGifters) SetPhone(v string) *SyncCardInfoRequestBodyGifters {
	s.Phone = &v
	return s
}

type SyncCardInfoRequestBodyPresentDetail struct {
	Count  *int64   `json:"Count,omitempty" xml:"Count,omitempty"`
	ItemId *string  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Price  *float64 `json:"Price,omitempty" xml:"Price,omitempty"`
	SkuId  *string  `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
}

func (s SyncCardInfoRequestBodyPresentDetail) String() string {
	return tea.Prettify(s)
}

func (s SyncCardInfoRequestBodyPresentDetail) GoString() string {
	return s.String()
}

func (s *SyncCardInfoRequestBodyPresentDetail) SetCount(v int64) *SyncCardInfoRequestBodyPresentDetail {
	s.Count = &v
	return s
}

func (s *SyncCardInfoRequestBodyPresentDetail) SetItemId(v string) *SyncCardInfoRequestBodyPresentDetail {
	s.ItemId = &v
	return s
}

func (s *SyncCardInfoRequestBodyPresentDetail) SetName(v string) *SyncCardInfoRequestBodyPresentDetail {
	s.Name = &v
	return s
}

func (s *SyncCardInfoRequestBodyPresentDetail) SetPrice(v float64) *SyncCardInfoRequestBodyPresentDetail {
	s.Price = &v
	return s
}

func (s *SyncCardInfoRequestBodyPresentDetail) SetSkuId(v string) *SyncCardInfoRequestBodyPresentDetail {
	s.SkuId = &v
	return s
}

type SyncCardInfoRequestBodyRecipient struct {
	HeaderImg *string `json:"HeaderImg,omitempty" xml:"HeaderImg,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Nickname  *string `json:"Nickname,omitempty" xml:"Nickname,omitempty"`
	// openId
	OpenId *string `json:"OpenId,omitempty" xml:"OpenId,omitempty"`
	Phone  *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s SyncCardInfoRequestBodyRecipient) String() string {
	return tea.Prettify(s)
}

func (s SyncCardInfoRequestBodyRecipient) GoString() string {
	return s.String()
}

func (s *SyncCardInfoRequestBodyRecipient) SetHeaderImg(v string) *SyncCardInfoRequestBodyRecipient {
	s.HeaderImg = &v
	return s
}

func (s *SyncCardInfoRequestBodyRecipient) SetId(v string) *SyncCardInfoRequestBodyRecipient {
	s.Id = &v
	return s
}

func (s *SyncCardInfoRequestBodyRecipient) SetNickname(v string) *SyncCardInfoRequestBodyRecipient {
	s.Nickname = &v
	return s
}

func (s *SyncCardInfoRequestBodyRecipient) SetOpenId(v string) *SyncCardInfoRequestBodyRecipient {
	s.OpenId = &v
	return s
}

func (s *SyncCardInfoRequestBodyRecipient) SetPhone(v string) *SyncCardInfoRequestBodyRecipient {
	s.Phone = &v
	return s
}

type SyncCardInfoShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncCardInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncCardInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *SyncCardInfoShrinkRequest) SetBodyShrink(v string) *SyncCardInfoShrinkRequest {
	s.BodyShrink = &v
	return s
}

type SyncCardInfoResponseBody struct {
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncCardInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncCardInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SyncCardInfoResponseBody) SetErrorCode(v string) *SyncCardInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SyncCardInfoResponseBody) SetErrorMessage(v string) *SyncCardInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SyncCardInfoResponseBody) SetHttpStatusCode(v string) *SyncCardInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SyncCardInfoResponseBody) SetMessage(v string) *SyncCardInfoResponseBody {
	s.Message = &v
	return s
}

func (s *SyncCardInfoResponseBody) SetRequestId(v string) *SyncCardInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncCardInfoResponseBody) SetSuccess(v string) *SyncCardInfoResponseBody {
	s.Success = &v
	return s
}

type SyncCardInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SyncCardInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncCardInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncCardInfoResponse) GoString() string {
	return s.String()
}

func (s *SyncCardInfoResponse) SetHeaders(v map[string]*string) *SyncCardInfoResponse {
	s.Headers = v
	return s
}

func (s *SyncCardInfoResponse) SetStatusCode(v int32) *SyncCardInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncCardInfoResponse) SetBody(v *SyncCardInfoResponseBody) *SyncCardInfoResponse {
	s.Body = v
	return s
}

type SyncMemberBehaviorInfoRequest struct {
	Body *SyncMemberBehaviorInfoRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s SyncMemberBehaviorInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncMemberBehaviorInfoRequest) GoString() string {
	return s.String()
}

func (s *SyncMemberBehaviorInfoRequest) SetBody(v *SyncMemberBehaviorInfoRequestBody) *SyncMemberBehaviorInfoRequest {
	s.Body = v
	return s
}

type SyncMemberBehaviorInfoRequestBody struct {
	ActionDuration  *string `json:"ActionDuration,omitempty" xml:"ActionDuration,omitempty"`
	ActionEndDate   *string `json:"ActionEndDate,omitempty" xml:"ActionEndDate,omitempty"`
	ActionResult    *bool   `json:"ActionResult,omitempty" xml:"ActionResult,omitempty"`
	ActionStartDate *string `json:"ActionStartDate,omitempty" xml:"ActionStartDate,omitempty"`
	ActionSubType   *string `json:"ActionSubType,omitempty" xml:"ActionSubType,omitempty"`
	ActionType      *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	Extra           *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	OpenMerchantId  *string `json:"OpenMerchantId,omitempty" xml:"OpenMerchantId,omitempty"`
	OuterMemberId   *string `json:"OuterMemberId,omitempty" xml:"OuterMemberId,omitempty"`
	PlatFormType    *string `json:"PlatFormType,omitempty" xml:"PlatFormType,omitempty"`
}

func (s SyncMemberBehaviorInfoRequestBody) String() string {
	return tea.Prettify(s)
}

func (s SyncMemberBehaviorInfoRequestBody) GoString() string {
	return s.String()
}

func (s *SyncMemberBehaviorInfoRequestBody) SetActionDuration(v string) *SyncMemberBehaviorInfoRequestBody {
	s.ActionDuration = &v
	return s
}

func (s *SyncMemberBehaviorInfoRequestBody) SetActionEndDate(v string) *SyncMemberBehaviorInfoRequestBody {
	s.ActionEndDate = &v
	return s
}

func (s *SyncMemberBehaviorInfoRequestBody) SetActionResult(v bool) *SyncMemberBehaviorInfoRequestBody {
	s.ActionResult = &v
	return s
}

func (s *SyncMemberBehaviorInfoRequestBody) SetActionStartDate(v string) *SyncMemberBehaviorInfoRequestBody {
	s.ActionStartDate = &v
	return s
}

func (s *SyncMemberBehaviorInfoRequestBody) SetActionSubType(v string) *SyncMemberBehaviorInfoRequestBody {
	s.ActionSubType = &v
	return s
}

func (s *SyncMemberBehaviorInfoRequestBody) SetActionType(v string) *SyncMemberBehaviorInfoRequestBody {
	s.ActionType = &v
	return s
}

func (s *SyncMemberBehaviorInfoRequestBody) SetExtra(v string) *SyncMemberBehaviorInfoRequestBody {
	s.Extra = &v
	return s
}

func (s *SyncMemberBehaviorInfoRequestBody) SetOpenMerchantId(v string) *SyncMemberBehaviorInfoRequestBody {
	s.OpenMerchantId = &v
	return s
}

func (s *SyncMemberBehaviorInfoRequestBody) SetOuterMemberId(v string) *SyncMemberBehaviorInfoRequestBody {
	s.OuterMemberId = &v
	return s
}

func (s *SyncMemberBehaviorInfoRequestBody) SetPlatFormType(v string) *SyncMemberBehaviorInfoRequestBody {
	s.PlatFormType = &v
	return s
}

type SyncMemberBehaviorInfoShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncMemberBehaviorInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncMemberBehaviorInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *SyncMemberBehaviorInfoShrinkRequest) SetBodyShrink(v string) *SyncMemberBehaviorInfoShrinkRequest {
	s.BodyShrink = &v
	return s
}

type SyncMemberBehaviorInfoResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncMemberBehaviorInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncMemberBehaviorInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SyncMemberBehaviorInfoResponseBody) SetErrorCode(v string) *SyncMemberBehaviorInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SyncMemberBehaviorInfoResponseBody) SetErrorMessage(v string) *SyncMemberBehaviorInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SyncMemberBehaviorInfoResponseBody) SetRequestId(v string) *SyncMemberBehaviorInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncMemberBehaviorInfoResponseBody) SetSuccess(v bool) *SyncMemberBehaviorInfoResponseBody {
	s.Success = &v
	return s
}

type SyncMemberBehaviorInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SyncMemberBehaviorInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncMemberBehaviorInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncMemberBehaviorInfoResponse) GoString() string {
	return s.String()
}

func (s *SyncMemberBehaviorInfoResponse) SetHeaders(v map[string]*string) *SyncMemberBehaviorInfoResponse {
	s.Headers = v
	return s
}

func (s *SyncMemberBehaviorInfoResponse) SetStatusCode(v int32) *SyncMemberBehaviorInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncMemberBehaviorInfoResponse) SetBody(v *SyncMemberBehaviorInfoResponseBody) *SyncMemberBehaviorInfoResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("retailadvqa"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddMemberBasicInfoWithOptions(tmpReq *AddMemberBasicInfoRequest, runtime *util.RuntimeOptions) (_result *AddMemberBasicInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddMemberBasicInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["body"] = request.BodyShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddMemberBasicInfo"),
		Version:     tea.String("2023-04-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddMemberBasicInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddMemberBasicInfo(request *AddMemberBasicInfoRequest) (_result *AddMemberBasicInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddMemberBasicInfoResponse{}
	_body, _err := client.AddMemberBasicInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MemberAccountDetailPageQueryWithOptions(tmpReq *MemberAccountDetailPageQueryRequest, runtime *util.RuntimeOptions) (_result *MemberAccountDetailPageQueryResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &MemberAccountDetailPageQueryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["body"] = request.BodyShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MemberAccountDetailPageQuery"),
		Version:     tea.String("2023-04-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MemberAccountDetailPageQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MemberAccountDetailPageQuery(request *MemberAccountDetailPageQueryRequest) (_result *MemberAccountDetailPageQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MemberAccountDetailPageQueryResponse{}
	_body, _err := client.MemberAccountDetailPageQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MemberPointChangeWithOptions(tmpReq *MemberPointChangeRequest, runtime *util.RuntimeOptions) (_result *MemberPointChangeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &MemberPointChangeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["body"] = request.BodyShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MemberPointChange"),
		Version:     tea.String("2023-04-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MemberPointChangeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MemberPointChange(request *MemberPointChangeRequest) (_result *MemberPointChangeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MemberPointChangeResponse{}
	_body, _err := client.MemberPointChangeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMemberBasicInfoWithOptions(tmpReq *QueryMemberBasicInfoRequest, runtime *util.RuntimeOptions) (_result *QueryMemberBasicInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryMemberBasicInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["body"] = request.BodyShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMemberBasicInfo"),
		Version:     tea.String("2023-04-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMemberBasicInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMemberBasicInfo(request *QueryMemberBasicInfoRequest) (_result *QueryMemberBasicInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMemberBasicInfoResponse{}
	_body, _err := client.QueryMemberBasicInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncCardInfoWithOptions(tmpReq *SyncCardInfoRequest, runtime *util.RuntimeOptions) (_result *SyncCardInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SyncCardInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["body"] = request.BodyShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SyncCardInfo"),
		Version:     tea.String("2023-04-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncCardInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncCardInfo(request *SyncCardInfoRequest) (_result *SyncCardInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SyncCardInfoResponse{}
	_body, _err := client.SyncCardInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncMemberBehaviorInfoWithOptions(tmpReq *SyncMemberBehaviorInfoRequest, runtime *util.RuntimeOptions) (_result *SyncMemberBehaviorInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SyncMemberBehaviorInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["body"] = request.BodyShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SyncMemberBehaviorInfo"),
		Version:     tea.String("2023-04-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncMemberBehaviorInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncMemberBehaviorInfo(request *SyncMemberBehaviorInfoRequest) (_result *SyncMemberBehaviorInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SyncMemberBehaviorInfoResponse{}
	_body, _err := client.SyncMemberBehaviorInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
