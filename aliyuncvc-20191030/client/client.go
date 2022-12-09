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

type ActiveDeviceRequest struct {
	ActiveCode    *string `json:"ActiveCode,omitempty" xml:"ActiveCode,omitempty"`
	DeviceVersion *string `json:"DeviceVersion,omitempty" xml:"DeviceVersion,omitempty"`
	IP            *string `json:"IP,omitempty" xml:"IP,omitempty"`
	Mac           *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	SN            *string `json:"SN,omitempty" xml:"SN,omitempty"`
}

func (s ActiveDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s ActiveDeviceRequest) GoString() string {
	return s.String()
}

func (s *ActiveDeviceRequest) SetActiveCode(v string) *ActiveDeviceRequest {
	s.ActiveCode = &v
	return s
}

func (s *ActiveDeviceRequest) SetDeviceVersion(v string) *ActiveDeviceRequest {
	s.DeviceVersion = &v
	return s
}

func (s *ActiveDeviceRequest) SetIP(v string) *ActiveDeviceRequest {
	s.IP = &v
	return s
}

func (s *ActiveDeviceRequest) SetMac(v string) *ActiveDeviceRequest {
	s.Mac = &v
	return s
}

func (s *ActiveDeviceRequest) SetSN(v string) *ActiveDeviceRequest {
	s.SN = &v
	return s
}

type ActiveDeviceResponseBody struct {
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Token     *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s ActiveDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActiveDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *ActiveDeviceResponseBody) SetErrorCode(v int32) *ActiveDeviceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ActiveDeviceResponseBody) SetMessage(v string) *ActiveDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *ActiveDeviceResponseBody) SetRequestId(v string) *ActiveDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActiveDeviceResponseBody) SetSuccess(v bool) *ActiveDeviceResponseBody {
	s.Success = &v
	return s
}

func (s *ActiveDeviceResponseBody) SetToken(v string) *ActiveDeviceResponseBody {
	s.Token = &v
	return s
}

type ActiveDeviceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ActiveDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ActiveDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s ActiveDeviceResponse) GoString() string {
	return s.String()
}

func (s *ActiveDeviceResponse) SetHeaders(v map[string]*string) *ActiveDeviceResponse {
	s.Headers = v
	return s
}

func (s *ActiveDeviceResponse) SetStatusCode(v int32) *ActiveDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *ActiveDeviceResponse) SetBody(v *ActiveDeviceResponseBody) *ActiveDeviceResponse {
	s.Body = v
	return s
}

type ActiveMeetingRequest struct {
	MeetingCode *string `json:"MeetingCode,omitempty" xml:"MeetingCode,omitempty"`
	MeetingUUID *string `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
}

func (s ActiveMeetingRequest) String() string {
	return tea.Prettify(s)
}

func (s ActiveMeetingRequest) GoString() string {
	return s.String()
}

func (s *ActiveMeetingRequest) SetMeetingCode(v string) *ActiveMeetingRequest {
	s.MeetingCode = &v
	return s
}

func (s *ActiveMeetingRequest) SetMeetingUUID(v string) *ActiveMeetingRequest {
	s.MeetingUUID = &v
	return s
}

type ActiveMeetingResponseBody struct {
	ErrorCode   *int32                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	MeetingInfo *ActiveMeetingResponseBodyMeetingInfo `json:"MeetingInfo,omitempty" xml:"MeetingInfo,omitempty" type:"Struct"`
	Message     *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ActiveMeetingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActiveMeetingResponseBody) GoString() string {
	return s.String()
}

func (s *ActiveMeetingResponseBody) SetErrorCode(v int32) *ActiveMeetingResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ActiveMeetingResponseBody) SetMeetingInfo(v *ActiveMeetingResponseBodyMeetingInfo) *ActiveMeetingResponseBody {
	s.MeetingInfo = v
	return s
}

func (s *ActiveMeetingResponseBody) SetMessage(v string) *ActiveMeetingResponseBody {
	s.Message = &v
	return s
}

func (s *ActiveMeetingResponseBody) SetRequestId(v string) *ActiveMeetingResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActiveMeetingResponseBody) SetSuccess(v bool) *ActiveMeetingResponseBody {
	s.Success = &v
	return s
}

type ActiveMeetingResponseBodyMeetingInfo struct {
	MeetingCode *string `json:"MeetingCode,omitempty" xml:"MeetingCode,omitempty"`
	ValidTime   *int64  `json:"ValidTime,omitempty" xml:"ValidTime,omitempty"`
}

func (s ActiveMeetingResponseBodyMeetingInfo) String() string {
	return tea.Prettify(s)
}

func (s ActiveMeetingResponseBodyMeetingInfo) GoString() string {
	return s.String()
}

func (s *ActiveMeetingResponseBodyMeetingInfo) SetMeetingCode(v string) *ActiveMeetingResponseBodyMeetingInfo {
	s.MeetingCode = &v
	return s
}

func (s *ActiveMeetingResponseBodyMeetingInfo) SetValidTime(v int64) *ActiveMeetingResponseBodyMeetingInfo {
	s.ValidTime = &v
	return s
}

type ActiveMeetingResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ActiveMeetingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ActiveMeetingResponse) String() string {
	return tea.Prettify(s)
}

func (s ActiveMeetingResponse) GoString() string {
	return s.String()
}

func (s *ActiveMeetingResponse) SetHeaders(v map[string]*string) *ActiveMeetingResponse {
	s.Headers = v
	return s
}

func (s *ActiveMeetingResponse) SetStatusCode(v int32) *ActiveMeetingResponse {
	s.StatusCode = &v
	return s
}

func (s *ActiveMeetingResponse) SetBody(v *ActiveMeetingResponseBody) *ActiveMeetingResponse {
	s.Body = v
	return s
}

type BatchCreateDeviceRequest struct {
	SN *string `json:"SN,omitempty" xml:"SN,omitempty"`
}

func (s BatchCreateDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateDeviceRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateDeviceRequest) SetSN(v string) *BatchCreateDeviceRequest {
	s.SN = &v
	return s
}

type BatchCreateDeviceResponseBody struct {
	Devices   []*BatchCreateDeviceResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	ErrorCode *int32                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchCreateDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateDeviceResponseBody) SetDevices(v []*BatchCreateDeviceResponseBodyDevices) *BatchCreateDeviceResponseBody {
	s.Devices = v
	return s
}

func (s *BatchCreateDeviceResponseBody) SetErrorCode(v int32) *BatchCreateDeviceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BatchCreateDeviceResponseBody) SetMessage(v string) *BatchCreateDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *BatchCreateDeviceResponseBody) SetRequestId(v string) *BatchCreateDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCreateDeviceResponseBody) SetSuccess(v bool) *BatchCreateDeviceResponseBody {
	s.Success = &v
	return s
}

type BatchCreateDeviceResponseBodyDevices struct {
	ActiveCode      *string `json:"ActiveCode,omitempty" xml:"ActiveCode,omitempty"`
	DeviceErrorCode *int32  `json:"DeviceErrorCode,omitempty" xml:"DeviceErrorCode,omitempty"`
	DeviceMessage   *string `json:"DeviceMessage,omitempty" xml:"DeviceMessage,omitempty"`
	SN              *string `json:"SN,omitempty" xml:"SN,omitempty"`
}

func (s BatchCreateDeviceResponseBodyDevices) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateDeviceResponseBodyDevices) GoString() string {
	return s.String()
}

func (s *BatchCreateDeviceResponseBodyDevices) SetActiveCode(v string) *BatchCreateDeviceResponseBodyDevices {
	s.ActiveCode = &v
	return s
}

func (s *BatchCreateDeviceResponseBodyDevices) SetDeviceErrorCode(v int32) *BatchCreateDeviceResponseBodyDevices {
	s.DeviceErrorCode = &v
	return s
}

func (s *BatchCreateDeviceResponseBodyDevices) SetDeviceMessage(v string) *BatchCreateDeviceResponseBodyDevices {
	s.DeviceMessage = &v
	return s
}

func (s *BatchCreateDeviceResponseBodyDevices) SetSN(v string) *BatchCreateDeviceResponseBodyDevices {
	s.SN = &v
	return s
}

type BatchCreateDeviceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchCreateDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchCreateDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateDeviceResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateDeviceResponse) SetHeaders(v map[string]*string) *BatchCreateDeviceResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateDeviceResponse) SetStatusCode(v int32) *BatchCreateDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCreateDeviceResponse) SetBody(v *BatchCreateDeviceResponseBody) *BatchCreateDeviceResponse {
	s.Body = v
	return s
}

type BatchJoinMeetingRequest struct {
	MeetingCode *string `json:"MeetingCode,omitempty" xml:"MeetingCode,omitempty"`
	Password    *string `json:"Password,omitempty" xml:"Password,omitempty"`
	RtcEngine   *string `json:"RtcEngine,omitempty" xml:"RtcEngine,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s BatchJoinMeetingRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchJoinMeetingRequest) GoString() string {
	return s.String()
}

func (s *BatchJoinMeetingRequest) SetMeetingCode(v string) *BatchJoinMeetingRequest {
	s.MeetingCode = &v
	return s
}

func (s *BatchJoinMeetingRequest) SetPassword(v string) *BatchJoinMeetingRequest {
	s.Password = &v
	return s
}

func (s *BatchJoinMeetingRequest) SetRtcEngine(v string) *BatchJoinMeetingRequest {
	s.RtcEngine = &v
	return s
}

func (s *BatchJoinMeetingRequest) SetUserId(v string) *BatchJoinMeetingRequest {
	s.UserId = &v
	return s
}

type BatchJoinMeetingResponseBody struct {
	ErrorCode   *int32                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	MeetingInfo *BatchJoinMeetingResponseBodyMeetingInfo `json:"MeetingInfo,omitempty" xml:"MeetingInfo,omitempty" type:"Struct"`
	Message     *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchJoinMeetingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchJoinMeetingResponseBody) GoString() string {
	return s.String()
}

func (s *BatchJoinMeetingResponseBody) SetErrorCode(v int32) *BatchJoinMeetingResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BatchJoinMeetingResponseBody) SetMeetingInfo(v *BatchJoinMeetingResponseBodyMeetingInfo) *BatchJoinMeetingResponseBody {
	s.MeetingInfo = v
	return s
}

func (s *BatchJoinMeetingResponseBody) SetMessage(v string) *BatchJoinMeetingResponseBody {
	s.Message = &v
	return s
}

func (s *BatchJoinMeetingResponseBody) SetRequestId(v string) *BatchJoinMeetingResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchJoinMeetingResponseBody) SetSuccess(v bool) *BatchJoinMeetingResponseBody {
	s.Success = &v
	return s
}

type BatchJoinMeetingResponseBodyMeetingInfo struct {
	ClientAppId   *string                                              `json:"ClientAppId,omitempty" xml:"ClientAppId,omitempty"`
	MeetingAppId  *string                                              `json:"MeetingAppId,omitempty" xml:"MeetingAppId,omitempty"`
	MeetingCode   *string                                              `json:"MeetingCode,omitempty" xml:"MeetingCode,omitempty"`
	MeetingDomain *string                                              `json:"MeetingDomain,omitempty" xml:"MeetingDomain,omitempty"`
	MeetingUUID   *string                                              `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
	MemberList    []*BatchJoinMeetingResponseBodyMeetingInfoMemberList `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
	SlsInfo       *BatchJoinMeetingResponseBodyMeetingInfoSlsInfo      `json:"SlsInfo,omitempty" xml:"SlsInfo,omitempty" type:"Struct"`
}

func (s BatchJoinMeetingResponseBodyMeetingInfo) String() string {
	return tea.Prettify(s)
}

func (s BatchJoinMeetingResponseBodyMeetingInfo) GoString() string {
	return s.String()
}

func (s *BatchJoinMeetingResponseBodyMeetingInfo) SetClientAppId(v string) *BatchJoinMeetingResponseBodyMeetingInfo {
	s.ClientAppId = &v
	return s
}

func (s *BatchJoinMeetingResponseBodyMeetingInfo) SetMeetingAppId(v string) *BatchJoinMeetingResponseBodyMeetingInfo {
	s.MeetingAppId = &v
	return s
}

func (s *BatchJoinMeetingResponseBodyMeetingInfo) SetMeetingCode(v string) *BatchJoinMeetingResponseBodyMeetingInfo {
	s.MeetingCode = &v
	return s
}

func (s *BatchJoinMeetingResponseBodyMeetingInfo) SetMeetingDomain(v string) *BatchJoinMeetingResponseBodyMeetingInfo {
	s.MeetingDomain = &v
	return s
}

func (s *BatchJoinMeetingResponseBodyMeetingInfo) SetMeetingUUID(v string) *BatchJoinMeetingResponseBodyMeetingInfo {
	s.MeetingUUID = &v
	return s
}

func (s *BatchJoinMeetingResponseBodyMeetingInfo) SetMemberList(v []*BatchJoinMeetingResponseBodyMeetingInfoMemberList) *BatchJoinMeetingResponseBodyMeetingInfo {
	s.MemberList = v
	return s
}

func (s *BatchJoinMeetingResponseBodyMeetingInfo) SetSlsInfo(v *BatchJoinMeetingResponseBodyMeetingInfoSlsInfo) *BatchJoinMeetingResponseBodyMeetingInfo {
	s.SlsInfo = v
	return s
}

type BatchJoinMeetingResponseBodyMeetingInfoMemberList struct {
	MeetingToken *string `json:"MeetingToken,omitempty" xml:"MeetingToken,omitempty"`
	MemberUUID   *string `json:"MemberUUID,omitempty" xml:"MemberUUID,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s BatchJoinMeetingResponseBodyMeetingInfoMemberList) String() string {
	return tea.Prettify(s)
}

func (s BatchJoinMeetingResponseBodyMeetingInfoMemberList) GoString() string {
	return s.String()
}

func (s *BatchJoinMeetingResponseBodyMeetingInfoMemberList) SetMeetingToken(v string) *BatchJoinMeetingResponseBodyMeetingInfoMemberList {
	s.MeetingToken = &v
	return s
}

func (s *BatchJoinMeetingResponseBodyMeetingInfoMemberList) SetMemberUUID(v string) *BatchJoinMeetingResponseBodyMeetingInfoMemberList {
	s.MemberUUID = &v
	return s
}

func (s *BatchJoinMeetingResponseBodyMeetingInfoMemberList) SetUserId(v string) *BatchJoinMeetingResponseBodyMeetingInfoMemberList {
	s.UserId = &v
	return s
}

type BatchJoinMeetingResponseBodyMeetingInfoSlsInfo struct {
	LogServiceEndpoint *string `json:"LogServiceEndpoint,omitempty" xml:"LogServiceEndpoint,omitempty"`
	Logstore           *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	Project            *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s BatchJoinMeetingResponseBodyMeetingInfoSlsInfo) String() string {
	return tea.Prettify(s)
}

func (s BatchJoinMeetingResponseBodyMeetingInfoSlsInfo) GoString() string {
	return s.String()
}

func (s *BatchJoinMeetingResponseBodyMeetingInfoSlsInfo) SetLogServiceEndpoint(v string) *BatchJoinMeetingResponseBodyMeetingInfoSlsInfo {
	s.LogServiceEndpoint = &v
	return s
}

func (s *BatchJoinMeetingResponseBodyMeetingInfoSlsInfo) SetLogstore(v string) *BatchJoinMeetingResponseBodyMeetingInfoSlsInfo {
	s.Logstore = &v
	return s
}

func (s *BatchJoinMeetingResponseBodyMeetingInfoSlsInfo) SetProject(v string) *BatchJoinMeetingResponseBodyMeetingInfoSlsInfo {
	s.Project = &v
	return s
}

type BatchJoinMeetingResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchJoinMeetingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchJoinMeetingResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchJoinMeetingResponse) GoString() string {
	return s.String()
}

func (s *BatchJoinMeetingResponse) SetHeaders(v map[string]*string) *BatchJoinMeetingResponse {
	s.Headers = v
	return s
}

func (s *BatchJoinMeetingResponse) SetStatusCode(v int32) *BatchJoinMeetingResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchJoinMeetingResponse) SetBody(v *BatchJoinMeetingResponseBody) *BatchJoinMeetingResponse {
	s.Body = v
	return s
}

type BatchJoinMeetingInternationalRequest struct {
	MeetingCode *string `json:"MeetingCode,omitempty" xml:"MeetingCode,omitempty"`
	Password    *string `json:"Password,omitempty" xml:"Password,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s BatchJoinMeetingInternationalRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchJoinMeetingInternationalRequest) GoString() string {
	return s.String()
}

func (s *BatchJoinMeetingInternationalRequest) SetMeetingCode(v string) *BatchJoinMeetingInternationalRequest {
	s.MeetingCode = &v
	return s
}

func (s *BatchJoinMeetingInternationalRequest) SetPassword(v string) *BatchJoinMeetingInternationalRequest {
	s.Password = &v
	return s
}

func (s *BatchJoinMeetingInternationalRequest) SetUserId(v string) *BatchJoinMeetingInternationalRequest {
	s.UserId = &v
	return s
}

type BatchJoinMeetingInternationalResponseBody struct {
	ErrorCode   *int32                                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	MeetingInfo *BatchJoinMeetingInternationalResponseBodyMeetingInfo `json:"MeetingInfo,omitempty" xml:"MeetingInfo,omitempty" type:"Struct"`
	Message     *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchJoinMeetingInternationalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchJoinMeetingInternationalResponseBody) GoString() string {
	return s.String()
}

func (s *BatchJoinMeetingInternationalResponseBody) SetErrorCode(v int32) *BatchJoinMeetingInternationalResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BatchJoinMeetingInternationalResponseBody) SetMeetingInfo(v *BatchJoinMeetingInternationalResponseBodyMeetingInfo) *BatchJoinMeetingInternationalResponseBody {
	s.MeetingInfo = v
	return s
}

func (s *BatchJoinMeetingInternationalResponseBody) SetMessage(v string) *BatchJoinMeetingInternationalResponseBody {
	s.Message = &v
	return s
}

func (s *BatchJoinMeetingInternationalResponseBody) SetRequestId(v string) *BatchJoinMeetingInternationalResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchJoinMeetingInternationalResponseBody) SetSuccess(v bool) *BatchJoinMeetingInternationalResponseBody {
	s.Success = &v
	return s
}

type BatchJoinMeetingInternationalResponseBodyMeetingInfo struct {
	ClientAppId   *string                                                           `json:"ClientAppId,omitempty" xml:"ClientAppId,omitempty"`
	MeetingAppId  *string                                                           `json:"MeetingAppId,omitempty" xml:"MeetingAppId,omitempty"`
	MeetingCode   *string                                                           `json:"MeetingCode,omitempty" xml:"MeetingCode,omitempty"`
	MeetingDomain *string                                                           `json:"MeetingDomain,omitempty" xml:"MeetingDomain,omitempty"`
	MeetingToken  *string                                                           `json:"MeetingToken,omitempty" xml:"MeetingToken,omitempty"`
	MeetingUUID   *string                                                           `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
	MemberList    []*BatchJoinMeetingInternationalResponseBodyMeetingInfoMemberList `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
	SlsInfo       *BatchJoinMeetingInternationalResponseBodyMeetingInfoSlsInfo      `json:"SlsInfo,omitempty" xml:"SlsInfo,omitempty" type:"Struct"`
}

func (s BatchJoinMeetingInternationalResponseBodyMeetingInfo) String() string {
	return tea.Prettify(s)
}

func (s BatchJoinMeetingInternationalResponseBodyMeetingInfo) GoString() string {
	return s.String()
}

func (s *BatchJoinMeetingInternationalResponseBodyMeetingInfo) SetClientAppId(v string) *BatchJoinMeetingInternationalResponseBodyMeetingInfo {
	s.ClientAppId = &v
	return s
}

func (s *BatchJoinMeetingInternationalResponseBodyMeetingInfo) SetMeetingAppId(v string) *BatchJoinMeetingInternationalResponseBodyMeetingInfo {
	s.MeetingAppId = &v
	return s
}

func (s *BatchJoinMeetingInternationalResponseBodyMeetingInfo) SetMeetingCode(v string) *BatchJoinMeetingInternationalResponseBodyMeetingInfo {
	s.MeetingCode = &v
	return s
}

func (s *BatchJoinMeetingInternationalResponseBodyMeetingInfo) SetMeetingDomain(v string) *BatchJoinMeetingInternationalResponseBodyMeetingInfo {
	s.MeetingDomain = &v
	return s
}

func (s *BatchJoinMeetingInternationalResponseBodyMeetingInfo) SetMeetingToken(v string) *BatchJoinMeetingInternationalResponseBodyMeetingInfo {
	s.MeetingToken = &v
	return s
}

func (s *BatchJoinMeetingInternationalResponseBodyMeetingInfo) SetMeetingUUID(v string) *BatchJoinMeetingInternationalResponseBodyMeetingInfo {
	s.MeetingUUID = &v
	return s
}

func (s *BatchJoinMeetingInternationalResponseBodyMeetingInfo) SetMemberList(v []*BatchJoinMeetingInternationalResponseBodyMeetingInfoMemberList) *BatchJoinMeetingInternationalResponseBodyMeetingInfo {
	s.MemberList = v
	return s
}

func (s *BatchJoinMeetingInternationalResponseBodyMeetingInfo) SetSlsInfo(v *BatchJoinMeetingInternationalResponseBodyMeetingInfoSlsInfo) *BatchJoinMeetingInternationalResponseBodyMeetingInfo {
	s.SlsInfo = v
	return s
}

type BatchJoinMeetingInternationalResponseBodyMeetingInfoMemberList struct {
	MeetingToken *string `json:"MeetingToken,omitempty" xml:"MeetingToken,omitempty"`
	MemberUUID   *string `json:"MemberUUID,omitempty" xml:"MemberUUID,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s BatchJoinMeetingInternationalResponseBodyMeetingInfoMemberList) String() string {
	return tea.Prettify(s)
}

func (s BatchJoinMeetingInternationalResponseBodyMeetingInfoMemberList) GoString() string {
	return s.String()
}

func (s *BatchJoinMeetingInternationalResponseBodyMeetingInfoMemberList) SetMeetingToken(v string) *BatchJoinMeetingInternationalResponseBodyMeetingInfoMemberList {
	s.MeetingToken = &v
	return s
}

func (s *BatchJoinMeetingInternationalResponseBodyMeetingInfoMemberList) SetMemberUUID(v string) *BatchJoinMeetingInternationalResponseBodyMeetingInfoMemberList {
	s.MemberUUID = &v
	return s
}

func (s *BatchJoinMeetingInternationalResponseBodyMeetingInfoMemberList) SetUserId(v string) *BatchJoinMeetingInternationalResponseBodyMeetingInfoMemberList {
	s.UserId = &v
	return s
}

type BatchJoinMeetingInternationalResponseBodyMeetingInfoSlsInfo struct {
	LogServiceEndpoint *string `json:"LogServiceEndpoint,omitempty" xml:"LogServiceEndpoint,omitempty"`
	Logstore           *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	Project            *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s BatchJoinMeetingInternationalResponseBodyMeetingInfoSlsInfo) String() string {
	return tea.Prettify(s)
}

func (s BatchJoinMeetingInternationalResponseBodyMeetingInfoSlsInfo) GoString() string {
	return s.String()
}

func (s *BatchJoinMeetingInternationalResponseBodyMeetingInfoSlsInfo) SetLogServiceEndpoint(v string) *BatchJoinMeetingInternationalResponseBodyMeetingInfoSlsInfo {
	s.LogServiceEndpoint = &v
	return s
}

func (s *BatchJoinMeetingInternationalResponseBodyMeetingInfoSlsInfo) SetLogstore(v string) *BatchJoinMeetingInternationalResponseBodyMeetingInfoSlsInfo {
	s.Logstore = &v
	return s
}

func (s *BatchJoinMeetingInternationalResponseBodyMeetingInfoSlsInfo) SetProject(v string) *BatchJoinMeetingInternationalResponseBodyMeetingInfoSlsInfo {
	s.Project = &v
	return s
}

type BatchJoinMeetingInternationalResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchJoinMeetingInternationalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchJoinMeetingInternationalResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchJoinMeetingInternationalResponse) GoString() string {
	return s.String()
}

func (s *BatchJoinMeetingInternationalResponse) SetHeaders(v map[string]*string) *BatchJoinMeetingInternationalResponse {
	s.Headers = v
	return s
}

func (s *BatchJoinMeetingInternationalResponse) SetStatusCode(v int32) *BatchJoinMeetingInternationalResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchJoinMeetingInternationalResponse) SetBody(v *BatchJoinMeetingInternationalResponseBody) *BatchJoinMeetingInternationalResponse {
	s.Body = v
	return s
}

type CallDeviceRequest struct {
	InviteName      *string `json:"InviteName,omitempty" xml:"InviteName,omitempty"`
	JoinMeetingFlag *bool   `json:"JoinMeetingFlag,omitempty" xml:"JoinMeetingFlag,omitempty"`
	MeetingCode     *string `json:"MeetingCode,omitempty" xml:"MeetingCode,omitempty"`
	OperateUserId   *string `json:"OperateUserId,omitempty" xml:"OperateUserId,omitempty"`
	SN              *string `json:"SN,omitempty" xml:"SN,omitempty"`
}

func (s CallDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s CallDeviceRequest) GoString() string {
	return s.String()
}

func (s *CallDeviceRequest) SetInviteName(v string) *CallDeviceRequest {
	s.InviteName = &v
	return s
}

func (s *CallDeviceRequest) SetJoinMeetingFlag(v bool) *CallDeviceRequest {
	s.JoinMeetingFlag = &v
	return s
}

func (s *CallDeviceRequest) SetMeetingCode(v string) *CallDeviceRequest {
	s.MeetingCode = &v
	return s
}

func (s *CallDeviceRequest) SetOperateUserId(v string) *CallDeviceRequest {
	s.OperateUserId = &v
	return s
}

func (s *CallDeviceRequest) SetSN(v string) *CallDeviceRequest {
	s.SN = &v
	return s
}

type CallDeviceResponseBody struct {
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CallDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CallDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *CallDeviceResponseBody) SetErrorCode(v int32) *CallDeviceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CallDeviceResponseBody) SetMessage(v string) *CallDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *CallDeviceResponseBody) SetMessageId(v string) *CallDeviceResponseBody {
	s.MessageId = &v
	return s
}

func (s *CallDeviceResponseBody) SetRequestId(v string) *CallDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CallDeviceResponseBody) SetSuccess(v bool) *CallDeviceResponseBody {
	s.Success = &v
	return s
}

type CallDeviceResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CallDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CallDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s CallDeviceResponse) GoString() string {
	return s.String()
}

func (s *CallDeviceResponse) SetHeaders(v map[string]*string) *CallDeviceResponse {
	s.Headers = v
	return s
}

func (s *CallDeviceResponse) SetStatusCode(v int32) *CallDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *CallDeviceResponse) SetBody(v *CallDeviceResponseBody) *CallDeviceResponse {
	s.Body = v
	return s
}

type ConferenceToLiveRequest struct {
	LiveName         *string `json:"LiveName,omitempty" xml:"LiveName,omitempty"`
	MeetingUUID      *string `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
	OpenPasswordFlag *bool   `json:"OpenPasswordFlag,omitempty" xml:"OpenPasswordFlag,omitempty"`
	Password         *string `json:"Password,omitempty" xml:"Password,omitempty"`
	UserId           *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ConferenceToLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s ConferenceToLiveRequest) GoString() string {
	return s.String()
}

func (s *ConferenceToLiveRequest) SetLiveName(v string) *ConferenceToLiveRequest {
	s.LiveName = &v
	return s
}

func (s *ConferenceToLiveRequest) SetMeetingUUID(v string) *ConferenceToLiveRequest {
	s.MeetingUUID = &v
	return s
}

func (s *ConferenceToLiveRequest) SetOpenPasswordFlag(v bool) *ConferenceToLiveRequest {
	s.OpenPasswordFlag = &v
	return s
}

func (s *ConferenceToLiveRequest) SetPassword(v string) *ConferenceToLiveRequest {
	s.Password = &v
	return s
}

func (s *ConferenceToLiveRequest) SetUserId(v string) *ConferenceToLiveRequest {
	s.UserId = &v
	return s
}

type ConferenceToLiveResponseBody struct {
	ErrorCode *int32                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	LiveInfo  *ConferenceToLiveResponseBodyLiveInfo `json:"LiveInfo,omitempty" xml:"LiveInfo,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConferenceToLiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConferenceToLiveResponseBody) GoString() string {
	return s.String()
}

func (s *ConferenceToLiveResponseBody) SetErrorCode(v int32) *ConferenceToLiveResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ConferenceToLiveResponseBody) SetLiveInfo(v *ConferenceToLiveResponseBodyLiveInfo) *ConferenceToLiveResponseBody {
	s.LiveInfo = v
	return s
}

func (s *ConferenceToLiveResponseBody) SetMessage(v string) *ConferenceToLiveResponseBody {
	s.Message = &v
	return s
}

func (s *ConferenceToLiveResponseBody) SetRequestId(v string) *ConferenceToLiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConferenceToLiveResponseBody) SetSuccess(v bool) *ConferenceToLiveResponseBody {
	s.Success = &v
	return s
}

type ConferenceToLiveResponseBodyLiveInfo struct {
	LiveUUID *string `json:"LiveUUID,omitempty" xml:"LiveUUID,omitempty"`
}

func (s ConferenceToLiveResponseBodyLiveInfo) String() string {
	return tea.Prettify(s)
}

func (s ConferenceToLiveResponseBodyLiveInfo) GoString() string {
	return s.String()
}

func (s *ConferenceToLiveResponseBodyLiveInfo) SetLiveUUID(v string) *ConferenceToLiveResponseBodyLiveInfo {
	s.LiveUUID = &v
	return s
}

type ConferenceToLiveResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ConferenceToLiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConferenceToLiveResponse) String() string {
	return tea.Prettify(s)
}

func (s ConferenceToLiveResponse) GoString() string {
	return s.String()
}

func (s *ConferenceToLiveResponse) SetHeaders(v map[string]*string) *ConferenceToLiveResponse {
	s.Headers = v
	return s
}

func (s *ConferenceToLiveResponse) SetStatusCode(v int32) *ConferenceToLiveResponse {
	s.StatusCode = &v
	return s
}

func (s *ConferenceToLiveResponse) SetBody(v *ConferenceToLiveResponseBody) *ConferenceToLiveResponse {
	s.Body = v
	return s
}

type CreateDeviceMeetingRequest struct {
	MeetingName     *string `json:"MeetingName,omitempty" xml:"MeetingName,omitempty"`
	OpenPasswordtag *bool   `json:"OpenPasswordtag,omitempty" xml:"OpenPasswordtag,omitempty"`
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	SN              *string `json:"SN,omitempty" xml:"SN,omitempty"`
	Token           *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s CreateDeviceMeetingRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceMeetingRequest) GoString() string {
	return s.String()
}

func (s *CreateDeviceMeetingRequest) SetMeetingName(v string) *CreateDeviceMeetingRequest {
	s.MeetingName = &v
	return s
}

func (s *CreateDeviceMeetingRequest) SetOpenPasswordtag(v bool) *CreateDeviceMeetingRequest {
	s.OpenPasswordtag = &v
	return s
}

func (s *CreateDeviceMeetingRequest) SetPassword(v string) *CreateDeviceMeetingRequest {
	s.Password = &v
	return s
}

func (s *CreateDeviceMeetingRequest) SetSN(v string) *CreateDeviceMeetingRequest {
	s.SN = &v
	return s
}

func (s *CreateDeviceMeetingRequest) SetToken(v string) *CreateDeviceMeetingRequest {
	s.Token = &v
	return s
}

type CreateDeviceMeetingResponseBody struct {
	Devices   *CreateDeviceMeetingResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Struct"`
	ErrorCode *int32                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDeviceMeetingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceMeetingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeviceMeetingResponseBody) SetDevices(v *CreateDeviceMeetingResponseBodyDevices) *CreateDeviceMeetingResponseBody {
	s.Devices = v
	return s
}

func (s *CreateDeviceMeetingResponseBody) SetErrorCode(v int32) *CreateDeviceMeetingResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDeviceMeetingResponseBody) SetMessage(v string) *CreateDeviceMeetingResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDeviceMeetingResponseBody) SetRequestId(v string) *CreateDeviceMeetingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDeviceMeetingResponseBody) SetSuccess(v bool) *CreateDeviceMeetingResponseBody {
	s.Success = &v
	return s
}

type CreateDeviceMeetingResponseBodyDevices struct {
	MeetingCode *string `json:"MeetingCode,omitempty" xml:"MeetingCode,omitempty"`
	MeetingUUID *string `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
}

func (s CreateDeviceMeetingResponseBodyDevices) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceMeetingResponseBodyDevices) GoString() string {
	return s.String()
}

func (s *CreateDeviceMeetingResponseBodyDevices) SetMeetingCode(v string) *CreateDeviceMeetingResponseBodyDevices {
	s.MeetingCode = &v
	return s
}

func (s *CreateDeviceMeetingResponseBodyDevices) SetMeetingUUID(v string) *CreateDeviceMeetingResponseBodyDevices {
	s.MeetingUUID = &v
	return s
}

type CreateDeviceMeetingResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDeviceMeetingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDeviceMeetingResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceMeetingResponse) GoString() string {
	return s.String()
}

func (s *CreateDeviceMeetingResponse) SetHeaders(v map[string]*string) *CreateDeviceMeetingResponse {
	s.Headers = v
	return s
}

func (s *CreateDeviceMeetingResponse) SetStatusCode(v int32) *CreateDeviceMeetingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeviceMeetingResponse) SetBody(v *CreateDeviceMeetingResponseBody) *CreateDeviceMeetingResponse {
	s.Body = v
	return s
}

type CreateEvaluationRequest struct {
	CreateTime  *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Evaluation  *string `json:"Evaluation,omitempty" xml:"Evaluation,omitempty"`
	MeetingUUID *string `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
	MemberUUID  *string `json:"MemberUUID,omitempty" xml:"MemberUUID,omitempty"`
	Memo        *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	Score       *string `json:"Score,omitempty" xml:"Score,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateEvaluationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEvaluationRequest) GoString() string {
	return s.String()
}

func (s *CreateEvaluationRequest) SetCreateTime(v int64) *CreateEvaluationRequest {
	s.CreateTime = &v
	return s
}

func (s *CreateEvaluationRequest) SetDescription(v string) *CreateEvaluationRequest {
	s.Description = &v
	return s
}

func (s *CreateEvaluationRequest) SetEvaluation(v string) *CreateEvaluationRequest {
	s.Evaluation = &v
	return s
}

func (s *CreateEvaluationRequest) SetMeetingUUID(v string) *CreateEvaluationRequest {
	s.MeetingUUID = &v
	return s
}

func (s *CreateEvaluationRequest) SetMemberUUID(v string) *CreateEvaluationRequest {
	s.MemberUUID = &v
	return s
}

func (s *CreateEvaluationRequest) SetMemo(v string) *CreateEvaluationRequest {
	s.Memo = &v
	return s
}

func (s *CreateEvaluationRequest) SetScore(v string) *CreateEvaluationRequest {
	s.Score = &v
	return s
}

func (s *CreateEvaluationRequest) SetUserId(v string) *CreateEvaluationRequest {
	s.UserId = &v
	return s
}

type CreateEvaluationResponseBody struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateEvaluationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEvaluationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEvaluationResponseBody) SetErrorCode(v string) *CreateEvaluationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateEvaluationResponseBody) SetMessage(v string) *CreateEvaluationResponseBody {
	s.Message = &v
	return s
}

func (s *CreateEvaluationResponseBody) SetRequestId(v string) *CreateEvaluationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEvaluationResponseBody) SetSuccess(v bool) *CreateEvaluationResponseBody {
	s.Success = &v
	return s
}

type CreateEvaluationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateEvaluationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEvaluationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEvaluationResponse) GoString() string {
	return s.String()
}

func (s *CreateEvaluationResponse) SetHeaders(v map[string]*string) *CreateEvaluationResponse {
	s.Headers = v
	return s
}

func (s *CreateEvaluationResponse) SetStatusCode(v int32) *CreateEvaluationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEvaluationResponse) SetBody(v *CreateEvaluationResponseBody) *CreateEvaluationResponse {
	s.Body = v
	return s
}

type CreateLiveRequest struct {
	LiveName         *string `json:"LiveName,omitempty" xml:"LiveName,omitempty"`
	OpenPasswordFlag *bool   `json:"OpenPasswordFlag,omitempty" xml:"OpenPasswordFlag,omitempty"`
	Password         *string `json:"Password,omitempty" xml:"Password,omitempty"`
	UserId           *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveRequest) SetLiveName(v string) *CreateLiveRequest {
	s.LiveName = &v
	return s
}

func (s *CreateLiveRequest) SetOpenPasswordFlag(v bool) *CreateLiveRequest {
	s.OpenPasswordFlag = &v
	return s
}

func (s *CreateLiveRequest) SetPassword(v string) *CreateLiveRequest {
	s.Password = &v
	return s
}

func (s *CreateLiveRequest) SetUserId(v string) *CreateLiveRequest {
	s.UserId = &v
	return s
}

type CreateLiveResponseBody struct {
	ErrorCode *int32                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	LiveInfo  *CreateLiveResponseBodyLiveInfo `json:"LiveInfo,omitempty" xml:"LiveInfo,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateLiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLiveResponseBody) SetErrorCode(v int32) *CreateLiveResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateLiveResponseBody) SetLiveInfo(v *CreateLiveResponseBodyLiveInfo) *CreateLiveResponseBody {
	s.LiveInfo = v
	return s
}

func (s *CreateLiveResponseBody) SetMessage(v string) *CreateLiveResponseBody {
	s.Message = &v
	return s
}

func (s *CreateLiveResponseBody) SetRequestId(v string) *CreateLiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLiveResponseBody) SetSuccess(v bool) *CreateLiveResponseBody {
	s.Success = &v
	return s
}

type CreateLiveResponseBodyLiveInfo struct {
	LiveUUID   *string `json:"LiveUUID,omitempty" xml:"LiveUUID,omitempty"`
	StreamUUID *string `json:"StreamUUID,omitempty" xml:"StreamUUID,omitempty"`
}

func (s CreateLiveResponseBodyLiveInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveResponseBodyLiveInfo) GoString() string {
	return s.String()
}

func (s *CreateLiveResponseBodyLiveInfo) SetLiveUUID(v string) *CreateLiveResponseBodyLiveInfo {
	s.LiveUUID = &v
	return s
}

func (s *CreateLiveResponseBodyLiveInfo) SetStreamUUID(v string) *CreateLiveResponseBodyLiveInfo {
	s.StreamUUID = &v
	return s
}

type CreateLiveResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateLiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLiveResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveResponse) GoString() string {
	return s.String()
}

func (s *CreateLiveResponse) SetHeaders(v map[string]*string) *CreateLiveResponse {
	s.Headers = v
	return s
}

func (s *CreateLiveResponse) SetStatusCode(v int32) *CreateLiveResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLiveResponse) SetBody(v *CreateLiveResponseBody) *CreateLiveResponse {
	s.Body = v
	return s
}

type CreateMeetingRequest struct {
	MasterEnableFlag *bool   `json:"MasterEnableFlag,omitempty" xml:"MasterEnableFlag,omitempty"`
	MeetingMode      *string `json:"MeetingMode,omitempty" xml:"MeetingMode,omitempty"`
	MeetingName      *string `json:"MeetingName,omitempty" xml:"MeetingName,omitempty"`
	OpenPasswordFlag *bool   `json:"OpenPasswordFlag,omitempty" xml:"OpenPasswordFlag,omitempty"`
	Password         *string `json:"Password,omitempty" xml:"Password,omitempty"`
	TenantCode       *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	UserId           *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateMeetingRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingRequest) GoString() string {
	return s.String()
}

func (s *CreateMeetingRequest) SetMasterEnableFlag(v bool) *CreateMeetingRequest {
	s.MasterEnableFlag = &v
	return s
}

func (s *CreateMeetingRequest) SetMeetingMode(v string) *CreateMeetingRequest {
	s.MeetingMode = &v
	return s
}

func (s *CreateMeetingRequest) SetMeetingName(v string) *CreateMeetingRequest {
	s.MeetingName = &v
	return s
}

func (s *CreateMeetingRequest) SetOpenPasswordFlag(v bool) *CreateMeetingRequest {
	s.OpenPasswordFlag = &v
	return s
}

func (s *CreateMeetingRequest) SetPassword(v string) *CreateMeetingRequest {
	s.Password = &v
	return s
}

func (s *CreateMeetingRequest) SetTenantCode(v string) *CreateMeetingRequest {
	s.TenantCode = &v
	return s
}

func (s *CreateMeetingRequest) SetUserId(v string) *CreateMeetingRequest {
	s.UserId = &v
	return s
}

type CreateMeetingResponseBody struct {
	ErrorCode   *int32                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	MeetingInfo *CreateMeetingResponseBodyMeetingInfo `json:"MeetingInfo,omitempty" xml:"MeetingInfo,omitempty" type:"Struct"`
	Message     *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMeetingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMeetingResponseBody) SetErrorCode(v int32) *CreateMeetingResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateMeetingResponseBody) SetMeetingInfo(v *CreateMeetingResponseBodyMeetingInfo) *CreateMeetingResponseBody {
	s.MeetingInfo = v
	return s
}

func (s *CreateMeetingResponseBody) SetMessage(v string) *CreateMeetingResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMeetingResponseBody) SetRequestId(v string) *CreateMeetingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMeetingResponseBody) SetSuccess(v bool) *CreateMeetingResponseBody {
	s.Success = &v
	return s
}

type CreateMeetingResponseBodyMeetingInfo struct {
	MeetingCode *string `json:"MeetingCode,omitempty" xml:"MeetingCode,omitempty"`
	MeetingUUID *string `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
}

func (s CreateMeetingResponseBodyMeetingInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingResponseBodyMeetingInfo) GoString() string {
	return s.String()
}

func (s *CreateMeetingResponseBodyMeetingInfo) SetMeetingCode(v string) *CreateMeetingResponseBodyMeetingInfo {
	s.MeetingCode = &v
	return s
}

func (s *CreateMeetingResponseBodyMeetingInfo) SetMeetingUUID(v string) *CreateMeetingResponseBodyMeetingInfo {
	s.MeetingUUID = &v
	return s
}

type CreateMeetingResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMeetingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMeetingResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingResponse) GoString() string {
	return s.String()
}

func (s *CreateMeetingResponse) SetHeaders(v map[string]*string) *CreateMeetingResponse {
	s.Headers = v
	return s
}

func (s *CreateMeetingResponse) SetStatusCode(v int32) *CreateMeetingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMeetingResponse) SetBody(v *CreateMeetingResponseBody) *CreateMeetingResponse {
	s.Body = v
	return s
}

type CreateMeetingInternationalRequest struct {
	MeetingName      *string `json:"MeetingName,omitempty" xml:"MeetingName,omitempty"`
	OpenPasswordFlag *string `json:"OpenPasswordFlag,omitempty" xml:"OpenPasswordFlag,omitempty"`
	Password         *string `json:"Password,omitempty" xml:"Password,omitempty"`
	UserId           *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateMeetingInternationalRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingInternationalRequest) GoString() string {
	return s.String()
}

func (s *CreateMeetingInternationalRequest) SetMeetingName(v string) *CreateMeetingInternationalRequest {
	s.MeetingName = &v
	return s
}

func (s *CreateMeetingInternationalRequest) SetOpenPasswordFlag(v string) *CreateMeetingInternationalRequest {
	s.OpenPasswordFlag = &v
	return s
}

func (s *CreateMeetingInternationalRequest) SetPassword(v string) *CreateMeetingInternationalRequest {
	s.Password = &v
	return s
}

func (s *CreateMeetingInternationalRequest) SetUserId(v string) *CreateMeetingInternationalRequest {
	s.UserId = &v
	return s
}

type CreateMeetingInternationalResponseBody struct {
	ErrorCode   *int32                                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	MeetingInfo *CreateMeetingInternationalResponseBodyMeetingInfo `json:"MeetingInfo,omitempty" xml:"MeetingInfo,omitempty" type:"Struct"`
	Message     *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMeetingInternationalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingInternationalResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMeetingInternationalResponseBody) SetErrorCode(v int32) *CreateMeetingInternationalResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateMeetingInternationalResponseBody) SetMeetingInfo(v *CreateMeetingInternationalResponseBodyMeetingInfo) *CreateMeetingInternationalResponseBody {
	s.MeetingInfo = v
	return s
}

func (s *CreateMeetingInternationalResponseBody) SetMessage(v string) *CreateMeetingInternationalResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMeetingInternationalResponseBody) SetRequestId(v string) *CreateMeetingInternationalResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMeetingInternationalResponseBody) SetSuccess(v bool) *CreateMeetingInternationalResponseBody {
	s.Success = &v
	return s
}

type CreateMeetingInternationalResponseBodyMeetingInfo struct {
	MeetingCode *string `json:"MeetingCode,omitempty" xml:"MeetingCode,omitempty"`
	MeetingUUID *string `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
}

func (s CreateMeetingInternationalResponseBodyMeetingInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingInternationalResponseBodyMeetingInfo) GoString() string {
	return s.String()
}

func (s *CreateMeetingInternationalResponseBodyMeetingInfo) SetMeetingCode(v string) *CreateMeetingInternationalResponseBodyMeetingInfo {
	s.MeetingCode = &v
	return s
}

func (s *CreateMeetingInternationalResponseBodyMeetingInfo) SetMeetingUUID(v string) *CreateMeetingInternationalResponseBodyMeetingInfo {
	s.MeetingUUID = &v
	return s
}

type CreateMeetingInternationalResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMeetingInternationalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMeetingInternationalResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingInternationalResponse) GoString() string {
	return s.String()
}

func (s *CreateMeetingInternationalResponse) SetHeaders(v map[string]*string) *CreateMeetingInternationalResponse {
	s.Headers = v
	return s
}

func (s *CreateMeetingInternationalResponse) SetStatusCode(v int32) *CreateMeetingInternationalResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMeetingInternationalResponse) SetBody(v *CreateMeetingInternationalResponseBody) *CreateMeetingInternationalResponse {
	s.Body = v
	return s
}

type CreateUserRequest struct {
	Count      *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	TenantCode *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	UserInfo   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s CreateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) SetCount(v int32) *CreateUserRequest {
	s.Count = &v
	return s
}

func (s *CreateUserRequest) SetTenantCode(v string) *CreateUserRequest {
	s.TenantCode = &v
	return s
}

func (s *CreateUserRequest) SetUserInfo(v string) *CreateUserRequest {
	s.UserInfo = &v
	return s
}

type CreateUserResponseBody struct {
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBody) SetErrorCode(v int32) *CreateUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateUserResponseBody) SetMessage(v string) *CreateUserResponseBody {
	s.Message = &v
	return s
}

func (s *CreateUserResponseBody) SetRequestId(v string) *CreateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserResponseBody) SetSuccess(v bool) *CreateUserResponseBody {
	s.Success = &v
	return s
}

func (s *CreateUserResponseBody) SetUserId(v string) *CreateUserResponseBody {
	s.UserId = &v
	return s
}

type CreateUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponse) GoString() string {
	return s.String()
}

func (s *CreateUserResponse) SetHeaders(v map[string]*string) *CreateUserResponse {
	s.Headers = v
	return s
}

func (s *CreateUserResponse) SetStatusCode(v int32) *CreateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserResponse) SetBody(v *CreateUserResponseBody) *CreateUserResponse {
	s.Body = v
	return s
}

type CreateUserInternationalRequest struct {
	Count    *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	UserInfo *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s CreateUserInternationalRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserInternationalRequest) GoString() string {
	return s.String()
}

func (s *CreateUserInternationalRequest) SetCount(v int32) *CreateUserInternationalRequest {
	s.Count = &v
	return s
}

func (s *CreateUserInternationalRequest) SetUserInfo(v string) *CreateUserInternationalRequest {
	s.UserInfo = &v
	return s
}

type CreateUserInternationalResponseBody struct {
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateUserInternationalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserInternationalResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserInternationalResponseBody) SetErrorCode(v int32) *CreateUserInternationalResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateUserInternationalResponseBody) SetMessage(v string) *CreateUserInternationalResponseBody {
	s.Message = &v
	return s
}

func (s *CreateUserInternationalResponseBody) SetRequestId(v string) *CreateUserInternationalResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserInternationalResponseBody) SetSuccess(v bool) *CreateUserInternationalResponseBody {
	s.Success = &v
	return s
}

func (s *CreateUserInternationalResponseBody) SetUserId(v string) *CreateUserInternationalResponseBody {
	s.UserId = &v
	return s
}

type CreateUserInternationalResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateUserInternationalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUserInternationalResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserInternationalResponse) GoString() string {
	return s.String()
}

func (s *CreateUserInternationalResponse) SetHeaders(v map[string]*string) *CreateUserInternationalResponse {
	s.Headers = v
	return s
}

func (s *CreateUserInternationalResponse) SetStatusCode(v int32) *CreateUserInternationalResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserInternationalResponse) SetBody(v *CreateUserInternationalResponseBody) *CreateUserInternationalResponse {
	s.Body = v
	return s
}

type CustomGonggeLayoutRequest struct {
	LayoutSolution *string `json:"LayoutSolution,omitempty" xml:"LayoutSolution,omitempty"`
	MeetingUUID    *string `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
}

func (s CustomGonggeLayoutRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomGonggeLayoutRequest) GoString() string {
	return s.String()
}

func (s *CustomGonggeLayoutRequest) SetLayoutSolution(v string) *CustomGonggeLayoutRequest {
	s.LayoutSolution = &v
	return s
}

func (s *CustomGonggeLayoutRequest) SetMeetingUUID(v string) *CustomGonggeLayoutRequest {
	s.MeetingUUID = &v
	return s
}

type CustomGonggeLayoutResponseBody struct {
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CustomGonggeLayoutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomGonggeLayoutResponseBody) GoString() string {
	return s.String()
}

func (s *CustomGonggeLayoutResponseBody) SetErrorCode(v int32) *CustomGonggeLayoutResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CustomGonggeLayoutResponseBody) SetMessage(v string) *CustomGonggeLayoutResponseBody {
	s.Message = &v
	return s
}

func (s *CustomGonggeLayoutResponseBody) SetRequestId(v string) *CustomGonggeLayoutResponseBody {
	s.RequestId = &v
	return s
}

func (s *CustomGonggeLayoutResponseBody) SetSuccess(v bool) *CustomGonggeLayoutResponseBody {
	s.Success = &v
	return s
}

type CustomGonggeLayoutResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CustomGonggeLayoutResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomGonggeLayoutResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomGonggeLayoutResponse) GoString() string {
	return s.String()
}

func (s *CustomGonggeLayoutResponse) SetHeaders(v map[string]*string) *CustomGonggeLayoutResponse {
	s.Headers = v
	return s
}

func (s *CustomGonggeLayoutResponse) SetStatusCode(v int32) *CustomGonggeLayoutResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomGonggeLayoutResponse) SetBody(v *CustomGonggeLayoutResponseBody) *CustomGonggeLayoutResponse {
	s.Body = v
	return s
}

type CustomLayoutRequest struct {
	LayoutInfo *string `json:"LayoutInfo,omitempty" xml:"LayoutInfo,omitempty"`
	LiveUUID   *string `json:"LiveUUID,omitempty" xml:"LiveUUID,omitempty"`
}

func (s CustomLayoutRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomLayoutRequest) GoString() string {
	return s.String()
}

func (s *CustomLayoutRequest) SetLayoutInfo(v string) *CustomLayoutRequest {
	s.LayoutInfo = &v
	return s
}

func (s *CustomLayoutRequest) SetLiveUUID(v string) *CustomLayoutRequest {
	s.LiveUUID = &v
	return s
}

type CustomLayoutResponseBody struct {
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CustomLayoutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomLayoutResponseBody) GoString() string {
	return s.String()
}

func (s *CustomLayoutResponseBody) SetErrorCode(v int32) *CustomLayoutResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CustomLayoutResponseBody) SetMessage(v string) *CustomLayoutResponseBody {
	s.Message = &v
	return s
}

func (s *CustomLayoutResponseBody) SetRequestId(v string) *CustomLayoutResponseBody {
	s.RequestId = &v
	return s
}

func (s *CustomLayoutResponseBody) SetSuccess(v bool) *CustomLayoutResponseBody {
	s.Success = &v
	return s
}

type CustomLayoutResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CustomLayoutResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomLayoutResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomLayoutResponse) GoString() string {
	return s.String()
}

func (s *CustomLayoutResponse) SetHeaders(v map[string]*string) *CustomLayoutResponse {
	s.Headers = v
	return s
}

func (s *CustomLayoutResponse) SetStatusCode(v int32) *CustomLayoutResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomLayoutResponse) SetBody(v *CustomLayoutResponseBody) *CustomLayoutResponse {
	s.Body = v
	return s
}

type DeleteDeviceRequest struct {
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SN      *string `json:"SN,omitempty" xml:"SN,omitempty"`
}

func (s DeleteDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeviceRequest) SetGroupId(v string) *DeleteDeviceRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteDeviceRequest) SetSN(v string) *DeleteDeviceRequest {
	s.SN = &v
	return s
}

type DeleteDeviceResponseBody struct {
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeviceResponseBody) SetErrorCode(v int32) *DeleteDeviceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDeviceResponseBody) SetMessage(v string) *DeleteDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDeviceResponseBody) SetRequestId(v string) *DeleteDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDeviceResponseBody) SetSuccess(v bool) *DeleteDeviceResponseBody {
	s.Success = &v
	return s
}

type DeleteDeviceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeviceResponse) SetHeaders(v map[string]*string) *DeleteDeviceResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeviceResponse) SetStatusCode(v int32) *DeleteDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeviceResponse) SetBody(v *DeleteDeviceResponseBody) *DeleteDeviceResponse {
	s.Body = v
	return s
}

type DeleteLiveRequest struct {
	LiveUUID *string `json:"LiveUUID,omitempty" xml:"LiveUUID,omitempty"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveRequest) SetLiveUUID(v string) *DeleteLiveRequest {
	s.LiveUUID = &v
	return s
}

func (s *DeleteLiveRequest) SetUserId(v string) *DeleteLiveRequest {
	s.UserId = &v
	return s
}

type DeleteLiveResponseBody struct {
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteLiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveResponseBody) SetErrorCode(v int32) *DeleteLiveResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteLiveResponseBody) SetMessage(v string) *DeleteLiveResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteLiveResponseBody) SetRequestId(v string) *DeleteLiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveResponseBody) SetSuccess(v bool) *DeleteLiveResponseBody {
	s.Success = &v
	return s
}

type DeleteLiveResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteLiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLiveResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveResponse) SetHeaders(v map[string]*string) *DeleteLiveResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveResponse) SetStatusCode(v int32) *DeleteLiveResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveResponse) SetBody(v *DeleteLiveResponseBody) *DeleteLiveResponse {
	s.Body = v
	return s
}

type DeleteMeetingRequest struct {
	MeetingUUID *string `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
	TenantCode  *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteMeetingRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMeetingRequest) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRequest) SetMeetingUUID(v string) *DeleteMeetingRequest {
	s.MeetingUUID = &v
	return s
}

func (s *DeleteMeetingRequest) SetTenantCode(v string) *DeleteMeetingRequest {
	s.TenantCode = &v
	return s
}

func (s *DeleteMeetingRequest) SetUserId(v string) *DeleteMeetingRequest {
	s.UserId = &v
	return s
}

type DeleteMeetingResponseBody struct {
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMeetingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMeetingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMeetingResponseBody) SetErrorCode(v int32) *DeleteMeetingResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteMeetingResponseBody) SetMessage(v string) *DeleteMeetingResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMeetingResponseBody) SetRequestId(v string) *DeleteMeetingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMeetingResponseBody) SetSuccess(v bool) *DeleteMeetingResponseBody {
	s.Success = &v
	return s
}

type DeleteMeetingResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteMeetingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMeetingResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMeetingResponse) GoString() string {
	return s.String()
}

func (s *DeleteMeetingResponse) SetHeaders(v map[string]*string) *DeleteMeetingResponse {
	s.Headers = v
	return s
}

func (s *DeleteMeetingResponse) SetStatusCode(v int32) *DeleteMeetingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMeetingResponse) SetBody(v *DeleteMeetingResponseBody) *DeleteMeetingResponse {
	s.Body = v
	return s
}

type DeleteUserRequest struct {
	Count    *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	UserInfo *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s DeleteUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserRequest) SetCount(v int32) *DeleteUserRequest {
	s.Count = &v
	return s
}

func (s *DeleteUserRequest) SetUserInfo(v string) *DeleteUserRequest {
	s.UserInfo = &v
	return s
}

type DeleteUserResponseBody struct {
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserResponseBody) SetErrorCode(v int32) *DeleteUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteUserResponseBody) SetMessage(v string) *DeleteUserResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteUserResponseBody) SetRequestId(v string) *DeleteUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserResponseBody) SetSuccess(v bool) *DeleteUserResponseBody {
	s.Success = &v
	return s
}

type DeleteUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserResponse) SetHeaders(v map[string]*string) *DeleteUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserResponse) SetStatusCode(v int32) *DeleteUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserResponse) SetBody(v *DeleteUserResponseBody) *DeleteUserResponse {
	s.Body = v
	return s
}

type EnableLiveSpeakerRequest struct {
	EnableSpeakerFlag *bool   `json:"EnableSpeakerFlag,omitempty" xml:"EnableSpeakerFlag,omitempty"`
	LiveUUID          *string `json:"LiveUUID,omitempty" xml:"LiveUUID,omitempty"`
}

func (s EnableLiveSpeakerRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableLiveSpeakerRequest) GoString() string {
	return s.String()
}

func (s *EnableLiveSpeakerRequest) SetEnableSpeakerFlag(v bool) *EnableLiveSpeakerRequest {
	s.EnableSpeakerFlag = &v
	return s
}

func (s *EnableLiveSpeakerRequest) SetLiveUUID(v string) *EnableLiveSpeakerRequest {
	s.LiveUUID = &v
	return s
}

type EnableLiveSpeakerResponseBody struct {
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableLiveSpeakerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableLiveSpeakerResponseBody) GoString() string {
	return s.String()
}

func (s *EnableLiveSpeakerResponseBody) SetErrorCode(v int32) *EnableLiveSpeakerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *EnableLiveSpeakerResponseBody) SetMessage(v string) *EnableLiveSpeakerResponseBody {
	s.Message = &v
	return s
}

func (s *EnableLiveSpeakerResponseBody) SetRequestId(v string) *EnableLiveSpeakerResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableLiveSpeakerResponseBody) SetSuccess(v bool) *EnableLiveSpeakerResponseBody {
	s.Success = &v
	return s
}

type EnableLiveSpeakerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableLiveSpeakerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableLiveSpeakerResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableLiveSpeakerResponse) GoString() string {
	return s.String()
}

func (s *EnableLiveSpeakerResponse) SetHeaders(v map[string]*string) *EnableLiveSpeakerResponse {
	s.Headers = v
	return s
}

func (s *EnableLiveSpeakerResponse) SetStatusCode(v int32) *EnableLiveSpeakerResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableLiveSpeakerResponse) SetBody(v *EnableLiveSpeakerResponseBody) *EnableLiveSpeakerResponse {
	s.Body = v
	return s
}

type EndDeviceMeetingRequest struct {
	MeetingUUID *string `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
	SN          *string `json:"SN,omitempty" xml:"SN,omitempty"`
}

func (s EndDeviceMeetingRequest) String() string {
	return tea.Prettify(s)
}

func (s EndDeviceMeetingRequest) GoString() string {
	return s.String()
}

func (s *EndDeviceMeetingRequest) SetMeetingUUID(v string) *EndDeviceMeetingRequest {
	s.MeetingUUID = &v
	return s
}

func (s *EndDeviceMeetingRequest) SetSN(v string) *EndDeviceMeetingRequest {
	s.SN = &v
	return s
}

type EndDeviceMeetingResponseBody struct {
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EndDeviceMeetingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EndDeviceMeetingResponseBody) GoString() string {
	return s.String()
}

func (s *EndDeviceMeetingResponseBody) SetErrorCode(v int32) *EndDeviceMeetingResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *EndDeviceMeetingResponseBody) SetMessage(v string) *EndDeviceMeetingResponseBody {
	s.Message = &v
	return s
}

func (s *EndDeviceMeetingResponseBody) SetRequestId(v string) *EndDeviceMeetingResponseBody {
	s.RequestId = &v
	return s
}

func (s *EndDeviceMeetingResponseBody) SetSuccess(v bool) *EndDeviceMeetingResponseBody {
	s.Success = &v
	return s
}

type EndDeviceMeetingResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EndDeviceMeetingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EndDeviceMeetingResponse) String() string {
	return tea.Prettify(s)
}

func (s EndDeviceMeetingResponse) GoString() string {
	return s.String()
}

func (s *EndDeviceMeetingResponse) SetHeaders(v map[string]*string) *EndDeviceMeetingResponse {
	s.Headers = v
	return s
}

func (s *EndDeviceMeetingResponse) SetStatusCode(v int32) *EndDeviceMeetingResponse {
	s.StatusCode = &v
	return s
}

func (s *EndDeviceMeetingResponse) SetBody(v *EndDeviceMeetingResponseBody) *EndDeviceMeetingResponse {
	s.Body = v
	return s
}

type EndLiveRequest struct {
	LiveUUID *string `json:"LiveUUID,omitempty" xml:"LiveUUID,omitempty"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s EndLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s EndLiveRequest) GoString() string {
	return s.String()
}

func (s *EndLiveRequest) SetLiveUUID(v string) *EndLiveRequest {
	s.LiveUUID = &v
	return s
}

func (s *EndLiveRequest) SetUserId(v string) *EndLiveRequest {
	s.UserId = &v
	return s
}

type EndLiveResponseBody struct {
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EndLiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EndLiveResponseBody) GoString() string {
	return s.String()
}

func (s *EndLiveResponseBody) SetErrorCode(v int32) *EndLiveResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *EndLiveResponseBody) SetMessage(v string) *EndLiveResponseBody {
	s.Message = &v
	return s
}

func (s *EndLiveResponseBody) SetRequestId(v string) *EndLiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *EndLiveResponseBody) SetSuccess(v bool) *EndLiveResponseBody {
	s.Success = &v
	return s
}

type EndLiveResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EndLiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EndLiveResponse) String() string {
	return tea.Prettify(s)
}

func (s EndLiveResponse) GoString() string {
	return s.String()
}

func (s *EndLiveResponse) SetHeaders(v map[string]*string) *EndLiveResponse {
	s.Headers = v
	return s
}

func (s *EndLiveResponse) SetStatusCode(v int32) *EndLiveResponse {
	s.StatusCode = &v
	return s
}

func (s *EndLiveResponse) SetBody(v *EndLiveResponseBody) *EndLiveResponse {
	s.Body = v
	return s
}

type GetAccountInfoResponseBody struct {
	AccountInfo *GetAccountInfoResponseBodyAccountInfo `json:"AccountInfo,omitempty" xml:"AccountInfo,omitempty" type:"Struct"`
	ErrorCode   *int32                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message     *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAccountInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccountInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountInfoResponseBody) SetAccountInfo(v *GetAccountInfoResponseBodyAccountInfo) *GetAccountInfoResponseBody {
	s.AccountInfo = v
	return s
}

func (s *GetAccountInfoResponseBody) SetErrorCode(v int32) *GetAccountInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetAccountInfoResponseBody) SetMessage(v string) *GetAccountInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetAccountInfoResponseBody) SetRequestId(v string) *GetAccountInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountInfoResponseBody) SetSuccess(v bool) *GetAccountInfoResponseBody {
	s.Success = &v
	return s
}

type GetAccountInfoResponseBodyAccountInfo struct {
	AccountApplicationMax    *int32 `json:"AccountApplicationMax,omitempty" xml:"AccountApplicationMax,omitempty"`
	AccountApplicationNumber *int32 `json:"AccountApplicationNumber,omitempty" xml:"AccountApplicationNumber,omitempty"`
	AccountConcurrentMax     *int32 `json:"AccountConcurrentMax,omitempty" xml:"AccountConcurrentMax,omitempty"`
}

func (s GetAccountInfoResponseBodyAccountInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAccountInfoResponseBodyAccountInfo) GoString() string {
	return s.String()
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetAccountApplicationMax(v int32) *GetAccountInfoResponseBodyAccountInfo {
	s.AccountApplicationMax = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetAccountApplicationNumber(v int32) *GetAccountInfoResponseBodyAccountInfo {
	s.AccountApplicationNumber = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetAccountConcurrentMax(v int32) *GetAccountInfoResponseBodyAccountInfo {
	s.AccountConcurrentMax = &v
	return s
}

type GetAccountInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAccountInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAccountInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAccountInfoResponse) SetHeaders(v map[string]*string) *GetAccountInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAccountInfoResponse) SetStatusCode(v int32) *GetAccountInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountInfoResponse) SetBody(v *GetAccountInfoResponseBody) *GetAccountInfoResponse {
	s.Body = v
	return s
}

type GetConferenceConcurrencyStatisticsResponseBody struct {
	Data      *GetConferenceConcurrencyStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *int32                                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetConferenceConcurrencyStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConferenceConcurrencyStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetConferenceConcurrencyStatisticsResponseBody) SetData(v *GetConferenceConcurrencyStatisticsResponseBodyData) *GetConferenceConcurrencyStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *GetConferenceConcurrencyStatisticsResponseBody) SetErrorCode(v int32) *GetConferenceConcurrencyStatisticsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetConferenceConcurrencyStatisticsResponseBody) SetMessage(v string) *GetConferenceConcurrencyStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *GetConferenceConcurrencyStatisticsResponseBody) SetRequestId(v string) *GetConferenceConcurrencyStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConferenceConcurrencyStatisticsResponseBody) SetSuccess(v bool) *GetConferenceConcurrencyStatisticsResponseBody {
	s.Success = &v
	return s
}

type GetConferenceConcurrencyStatisticsResponseBodyData struct {
	BuyConcurrency *int32 `json:"BuyConcurrency,omitempty" xml:"BuyConcurrency,omitempty"`
	CurrentMeeting *int32 `json:"CurrentMeeting,omitempty" xml:"CurrentMeeting,omitempty"`
	UseConcurrency *int32 `json:"UseConcurrency,omitempty" xml:"UseConcurrency,omitempty"`
}

func (s GetConferenceConcurrencyStatisticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetConferenceConcurrencyStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetConferenceConcurrencyStatisticsResponseBodyData) SetBuyConcurrency(v int32) *GetConferenceConcurrencyStatisticsResponseBodyData {
	s.BuyConcurrency = &v
	return s
}

func (s *GetConferenceConcurrencyStatisticsResponseBodyData) SetCurrentMeeting(v int32) *GetConferenceConcurrencyStatisticsResponseBodyData {
	s.CurrentMeeting = &v
	return s
}

func (s *GetConferenceConcurrencyStatisticsResponseBodyData) SetUseConcurrency(v int32) *GetConferenceConcurrencyStatisticsResponseBodyData {
	s.UseConcurrency = &v
	return s
}

type GetConferenceConcurrencyStatisticsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetConferenceConcurrencyStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConferenceConcurrencyStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConferenceConcurrencyStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetConferenceConcurrencyStatisticsResponse) SetHeaders(v map[string]*string) *GetConferenceConcurrencyStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetConferenceConcurrencyStatisticsResponse) SetStatusCode(v int32) *GetConferenceConcurrencyStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConferenceConcurrencyStatisticsResponse) SetBody(v *GetConferenceConcurrencyStatisticsResponseBody) *GetConferenceConcurrencyStatisticsResponse {
	s.Body = v
	return s
}

type GetDeviceActiveCodeRequest struct {
	SN *string `json:"SN,omitempty" xml:"SN,omitempty"`
}

func (s GetDeviceActiveCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceActiveCodeRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceActiveCodeRequest) SetSN(v string) *GetDeviceActiveCodeRequest {
	s.SN = &v
	return s
}

type GetDeviceActiveCodeResponseBody struct {
	Devices   []*GetDeviceActiveCodeResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	ErrorCode *int32                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDeviceActiveCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceActiveCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceActiveCodeResponseBody) SetDevices(v []*GetDeviceActiveCodeResponseBodyDevices) *GetDeviceActiveCodeResponseBody {
	s.Devices = v
	return s
}

func (s *GetDeviceActiveCodeResponseBody) SetErrorCode(v int32) *GetDeviceActiveCodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDeviceActiveCodeResponseBody) SetMessage(v string) *GetDeviceActiveCodeResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceActiveCodeResponseBody) SetSuccess(v bool) *GetDeviceActiveCodeResponseBody {
	s.Success = &v
	return s
}

type GetDeviceActiveCodeResponseBodyDevices struct {
	ActiveCode      *string `json:"ActiveCode,omitempty" xml:"ActiveCode,omitempty"`
	DeviceErrorCode *int32  `json:"DeviceErrorCode,omitempty" xml:"DeviceErrorCode,omitempty"`
	DeviceMessage   *string `json:"DeviceMessage,omitempty" xml:"DeviceMessage,omitempty"`
	SN              *string `json:"SN,omitempty" xml:"SN,omitempty"`
}

func (s GetDeviceActiveCodeResponseBodyDevices) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceActiveCodeResponseBodyDevices) GoString() string {
	return s.String()
}

func (s *GetDeviceActiveCodeResponseBodyDevices) SetActiveCode(v string) *GetDeviceActiveCodeResponseBodyDevices {
	s.ActiveCode = &v
	return s
}

func (s *GetDeviceActiveCodeResponseBodyDevices) SetDeviceErrorCode(v int32) *GetDeviceActiveCodeResponseBodyDevices {
	s.DeviceErrorCode = &v
	return s
}

func (s *GetDeviceActiveCodeResponseBodyDevices) SetDeviceMessage(v string) *GetDeviceActiveCodeResponseBodyDevices {
	s.DeviceMessage = &v
	return s
}

func (s *GetDeviceActiveCodeResponseBodyDevices) SetSN(v string) *GetDeviceActiveCodeResponseBodyDevices {
	s.SN = &v
	return s
}

type GetDeviceActiveCodeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDeviceActiveCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeviceActiveCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceActiveCodeResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceActiveCodeResponse) SetHeaders(v map[string]*string) *GetDeviceActiveCodeResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceActiveCodeResponse) SetStatusCode(v int32) *GetDeviceActiveCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceActiveCodeResponse) SetBody(v *GetDeviceActiveCodeResponseBody) *GetDeviceActiveCodeResponse {
	s.Body = v
	return s
}

type GetDeviceInfoRequest struct {
	CastScreenCode *string `json:"CastScreenCode,omitempty" xml:"CastScreenCode,omitempty"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Sn             *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	TenantCode     *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
}

func (s GetDeviceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceInfoRequest) SetCastScreenCode(v string) *GetDeviceInfoRequest {
	s.CastScreenCode = &v
	return s
}

func (s *GetDeviceInfoRequest) SetGroupId(v string) *GetDeviceInfoRequest {
	s.GroupId = &v
	return s
}

func (s *GetDeviceInfoRequest) SetSn(v string) *GetDeviceInfoRequest {
	s.Sn = &v
	return s
}

func (s *GetDeviceInfoRequest) SetTenantCode(v string) *GetDeviceInfoRequest {
	s.TenantCode = &v
	return s
}

type GetDeviceInfoResponseBody struct {
	Device    *GetDeviceInfoResponseBodyDevice `json:"Device,omitempty" xml:"Device,omitempty" type:"Struct"`
	ErrorCode *int32                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDeviceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceInfoResponseBody) SetDevice(v *GetDeviceInfoResponseBodyDevice) *GetDeviceInfoResponseBody {
	s.Device = v
	return s
}

func (s *GetDeviceInfoResponseBody) SetErrorCode(v int32) *GetDeviceInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDeviceInfoResponseBody) SetMessage(v string) *GetDeviceInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceInfoResponseBody) SetRequestId(v string) *GetDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceInfoResponseBody) SetSuccess(v bool) *GetDeviceInfoResponseBody {
	s.Success = &v
	return s
}

type GetDeviceInfoResponseBodyDevice struct {
	ActivationCode *string `json:"ActivationCode,omitempty" xml:"ActivationCode,omitempty"`
	CastScreenCode *string `json:"CastScreenCode,omitempty" xml:"CastScreenCode,omitempty"`
	IP             *string `json:"IP,omitempty" xml:"IP,omitempty"`
	Mac            *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	Port           *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Sn             *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	Ssid           *string `json:"Ssid,omitempty" xml:"Ssid,omitempty"`
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDeviceInfoResponseBodyDevice) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceInfoResponseBodyDevice) GoString() string {
	return s.String()
}

func (s *GetDeviceInfoResponseBodyDevice) SetActivationCode(v string) *GetDeviceInfoResponseBodyDevice {
	s.ActivationCode = &v
	return s
}

func (s *GetDeviceInfoResponseBodyDevice) SetCastScreenCode(v string) *GetDeviceInfoResponseBodyDevice {
	s.CastScreenCode = &v
	return s
}

func (s *GetDeviceInfoResponseBodyDevice) SetIP(v string) *GetDeviceInfoResponseBodyDevice {
	s.IP = &v
	return s
}

func (s *GetDeviceInfoResponseBodyDevice) SetMac(v string) *GetDeviceInfoResponseBodyDevice {
	s.Mac = &v
	return s
}

func (s *GetDeviceInfoResponseBodyDevice) SetPort(v string) *GetDeviceInfoResponseBodyDevice {
	s.Port = &v
	return s
}

func (s *GetDeviceInfoResponseBodyDevice) SetSn(v string) *GetDeviceInfoResponseBodyDevice {
	s.Sn = &v
	return s
}

func (s *GetDeviceInfoResponseBodyDevice) SetSsid(v string) *GetDeviceInfoResponseBodyDevice {
	s.Ssid = &v
	return s
}

func (s *GetDeviceInfoResponseBodyDevice) SetStatus(v int32) *GetDeviceInfoResponseBodyDevice {
	s.Status = &v
	return s
}

type GetDeviceInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeviceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceInfoResponse) SetHeaders(v map[string]*string) *GetDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceInfoResponse) SetStatusCode(v int32) *GetDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceInfoResponse) SetBody(v *GetDeviceInfoResponseBody) *GetDeviceInfoResponse {
	s.Body = v
	return s
}

type GetDeviceListRequest struct {
	CastScreenCode *string `json:"CastScreenCode,omitempty" xml:"CastScreenCode,omitempty"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s GetDeviceListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceListRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceListRequest) SetCastScreenCode(v string) *GetDeviceListRequest {
	s.CastScreenCode = &v
	return s
}

func (s *GetDeviceListRequest) SetGroupId(v string) *GetDeviceListRequest {
	s.GroupId = &v
	return s
}

type GetDeviceListResponseBody struct {
	Data      []*GetDeviceListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode *int32                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDeviceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceListResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceListResponseBody) SetData(v []*GetDeviceListResponseBodyData) *GetDeviceListResponseBody {
	s.Data = v
	return s
}

func (s *GetDeviceListResponseBody) SetErrorCode(v int32) *GetDeviceListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDeviceListResponseBody) SetMessage(v string) *GetDeviceListResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceListResponseBody) SetRequestId(v string) *GetDeviceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceListResponseBody) SetSuccess(v bool) *GetDeviceListResponseBody {
	s.Success = &v
	return s
}

type GetDeviceListResponseBodyData struct {
	ActivationCode *string `json:"ActivationCode,omitempty" xml:"ActivationCode,omitempty"`
	CastScreenCode *string `json:"CastScreenCode,omitempty" xml:"CastScreenCode,omitempty"`
	IP             *string `json:"IP,omitempty" xml:"IP,omitempty"`
	Mac            *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	Port           *string `json:"Port,omitempty" xml:"Port,omitempty"`
	SN             *string `json:"SN,omitempty" xml:"SN,omitempty"`
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDeviceListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeviceListResponseBodyData) SetActivationCode(v string) *GetDeviceListResponseBodyData {
	s.ActivationCode = &v
	return s
}

func (s *GetDeviceListResponseBodyData) SetCastScreenCode(v string) *GetDeviceListResponseBodyData {
	s.CastScreenCode = &v
	return s
}

func (s *GetDeviceListResponseBodyData) SetIP(v string) *GetDeviceListResponseBodyData {
	s.IP = &v
	return s
}

func (s *GetDeviceListResponseBodyData) SetMac(v string) *GetDeviceListResponseBodyData {
	s.Mac = &v
	return s
}

func (s *GetDeviceListResponseBodyData) SetPort(v string) *GetDeviceListResponseBodyData {
	s.Port = &v
	return s
}

func (s *GetDeviceListResponseBodyData) SetSN(v string) *GetDeviceListResponseBodyData {
	s.SN = &v
	return s
}

func (s *GetDeviceListResponseBodyData) SetStatus(v int32) *GetDeviceListResponseBodyData {
	s.Status = &v
	return s
}

type GetDeviceListResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDeviceListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeviceListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceListResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceListResponse) SetHeaders(v map[string]*string) *GetDeviceListResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceListResponse) SetStatusCode(v int32) *GetDeviceListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceListResponse) SetBody(v *GetDeviceListResponseBody) *GetDeviceListResponse {
	s.Body = v
	return s
}

type GetDeviceTokenRequest struct {
	SN    *string `json:"SN,omitempty" xml:"SN,omitempty"`
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetDeviceTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceTokenRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceTokenRequest) SetSN(v string) *GetDeviceTokenRequest {
	s.SN = &v
	return s
}

func (s *GetDeviceTokenRequest) SetToken(v string) *GetDeviceTokenRequest {
	s.Token = &v
	return s
}

type GetDeviceTokenResponseBody struct {
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Token     *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetDeviceTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceTokenResponseBody) SetErrorCode(v int32) *GetDeviceTokenResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDeviceTokenResponseBody) SetMessage(v string) *GetDeviceTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceTokenResponseBody) SetRequestId(v string) *GetDeviceTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceTokenResponseBody) SetSuccess(v bool) *GetDeviceTokenResponseBody {
	s.Success = &v
	return s
}

func (s *GetDeviceTokenResponseBody) SetToken(v string) *GetDeviceTokenResponseBody {
	s.Token = &v
	return s
}

type GetDeviceTokenResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDeviceTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeviceTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceTokenResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceTokenResponse) SetHeaders(v map[string]*string) *GetDeviceTokenResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceTokenResponse) SetStatusCode(v int32) *GetDeviceTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceTokenResponse) SetBody(v *GetDeviceTokenResponseBody) *GetDeviceTokenResponse {
	s.Body = v
	return s
}

type GetGrtnTokenResponseBody struct {
	ErrorCode *int64                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	GrtnInfo  *GetGrtnTokenResponseBodyGrtnInfo `json:"GrtnInfo,omitempty" xml:"GrtnInfo,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGrtnTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGrtnTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetGrtnTokenResponseBody) SetErrorCode(v int64) *GetGrtnTokenResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetGrtnTokenResponseBody) SetGrtnInfo(v *GetGrtnTokenResponseBodyGrtnInfo) *GetGrtnTokenResponseBody {
	s.GrtnInfo = v
	return s
}

func (s *GetGrtnTokenResponseBody) SetMessage(v string) *GetGrtnTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetGrtnTokenResponseBody) SetRequestId(v string) *GetGrtnTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGrtnTokenResponseBody) SetSuccess(v bool) *GetGrtnTokenResponseBody {
	s.Success = &v
	return s
}

type GetGrtnTokenResponseBodyGrtnInfo struct {
	Agent     *string `json:"Agent,omitempty" xml:"Agent,omitempty"`
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Gslb      *string `json:"Gslb,omitempty" xml:"Gslb,omitempty"`
	Nonce     *string `json:"Nonce,omitempty" xml:"Nonce,omitempty"`
	Timestamp *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Token     *string `json:"Token,omitempty" xml:"Token,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetGrtnTokenResponseBodyGrtnInfo) String() string {
	return tea.Prettify(s)
}

func (s GetGrtnTokenResponseBodyGrtnInfo) GoString() string {
	return s.String()
}

func (s *GetGrtnTokenResponseBodyGrtnInfo) SetAgent(v string) *GetGrtnTokenResponseBodyGrtnInfo {
	s.Agent = &v
	return s
}

func (s *GetGrtnTokenResponseBodyGrtnInfo) SetAppId(v string) *GetGrtnTokenResponseBodyGrtnInfo {
	s.AppId = &v
	return s
}

func (s *GetGrtnTokenResponseBodyGrtnInfo) SetChannelId(v string) *GetGrtnTokenResponseBodyGrtnInfo {
	s.ChannelId = &v
	return s
}

func (s *GetGrtnTokenResponseBodyGrtnInfo) SetGslb(v string) *GetGrtnTokenResponseBodyGrtnInfo {
	s.Gslb = &v
	return s
}

func (s *GetGrtnTokenResponseBodyGrtnInfo) SetNonce(v string) *GetGrtnTokenResponseBodyGrtnInfo {
	s.Nonce = &v
	return s
}

func (s *GetGrtnTokenResponseBodyGrtnInfo) SetTimestamp(v int64) *GetGrtnTokenResponseBodyGrtnInfo {
	s.Timestamp = &v
	return s
}

func (s *GetGrtnTokenResponseBodyGrtnInfo) SetToken(v string) *GetGrtnTokenResponseBodyGrtnInfo {
	s.Token = &v
	return s
}

func (s *GetGrtnTokenResponseBodyGrtnInfo) SetUserId(v string) *GetGrtnTokenResponseBodyGrtnInfo {
	s.UserId = &v
	return s
}

type GetGrtnTokenResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetGrtnTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGrtnTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGrtnTokenResponse) GoString() string {
	return s.String()
}

func (s *GetGrtnTokenResponse) SetHeaders(v map[string]*string) *GetGrtnTokenResponse {
	s.Headers = v
	return s
}

func (s *GetGrtnTokenResponse) SetStatusCode(v int32) *GetGrtnTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGrtnTokenResponse) SetBody(v *GetGrtnTokenResponseBody) *GetGrtnTokenResponse {
	s.Body = v
	return s
}

type GetMeetingRequest struct {
	MeetingUUID *string `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
	TenantCode  *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
}

func (s GetMeetingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMeetingRequest) GoString() string {
	return s.String()
}

func (s *GetMeetingRequest) SetMeetingUUID(v string) *GetMeetingRequest {
	s.MeetingUUID = &v
	return s
}

func (s *GetMeetingRequest) SetTenantCode(v string) *GetMeetingRequest {
	s.TenantCode = &v
	return s
}

type GetMeetingResponseBody struct {
	ErrorCode   *int32                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	MeetingInfo *GetMeetingResponseBodyMeetingInfo `json:"MeetingInfo,omitempty" xml:"MeetingInfo,omitempty" type:"Struct"`
	Message     *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMeetingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMeetingResponseBody) GoString() string {
	return s.String()
}

func (s *GetMeetingResponseBody) SetErrorCode(v int32) *GetMeetingResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMeetingResponseBody) SetMeetingInfo(v *GetMeetingResponseBodyMeetingInfo) *GetMeetingResponseBody {
	s.MeetingInfo = v
	return s
}

func (s *GetMeetingResponseBody) SetMessage(v string) *GetMeetingResponseBody {
	s.Message = &v
	return s
}

func (s *GetMeetingResponseBody) SetRequestId(v string) *GetMeetingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMeetingResponseBody) SetSuccess(v bool) *GetMeetingResponseBody {
	s.Success = &v
	return s
}

type GetMeetingResponseBodyMeetingInfo struct {
	AppointmentType   *int32                                         `json:"AppointmentType,omitempty" xml:"AppointmentType,omitempty"`
	CreateTime        *int64                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EndTime           *int64                                         `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MasterEnableFlag  *bool                                          `json:"MasterEnableFlag,omitempty" xml:"MasterEnableFlag,omitempty"`
	MeetingCode       *string                                        `json:"MeetingCode,omitempty" xml:"MeetingCode,omitempty"`
	MeetingCodeExpire *bool                                          `json:"MeetingCodeExpire,omitempty" xml:"MeetingCodeExpire,omitempty"`
	MeetingName       *string                                        `json:"MeetingName,omitempty" xml:"MeetingName,omitempty"`
	MeetingUUID       *string                                        `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
	MemberList        []*GetMeetingResponseBodyMeetingInfoMemberList `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
	OpenPasswordFlag  *bool                                          `json:"OpenPasswordFlag,omitempty" xml:"OpenPasswordFlag,omitempty"`
	Password          *string                                        `json:"Password,omitempty" xml:"Password,omitempty"`
	StartTime         *int64                                         `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status            *int32                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId            *string                                        `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ValidTime         *int64                                         `json:"ValidTime,omitempty" xml:"ValidTime,omitempty"`
}

func (s GetMeetingResponseBodyMeetingInfo) String() string {
	return tea.Prettify(s)
}

func (s GetMeetingResponseBodyMeetingInfo) GoString() string {
	return s.String()
}

func (s *GetMeetingResponseBodyMeetingInfo) SetAppointmentType(v int32) *GetMeetingResponseBodyMeetingInfo {
	s.AppointmentType = &v
	return s
}

func (s *GetMeetingResponseBodyMeetingInfo) SetCreateTime(v int64) *GetMeetingResponseBodyMeetingInfo {
	s.CreateTime = &v
	return s
}

func (s *GetMeetingResponseBodyMeetingInfo) SetEndTime(v int64) *GetMeetingResponseBodyMeetingInfo {
	s.EndTime = &v
	return s
}

func (s *GetMeetingResponseBodyMeetingInfo) SetMasterEnableFlag(v bool) *GetMeetingResponseBodyMeetingInfo {
	s.MasterEnableFlag = &v
	return s
}

func (s *GetMeetingResponseBodyMeetingInfo) SetMeetingCode(v string) *GetMeetingResponseBodyMeetingInfo {
	s.MeetingCode = &v
	return s
}

func (s *GetMeetingResponseBodyMeetingInfo) SetMeetingCodeExpire(v bool) *GetMeetingResponseBodyMeetingInfo {
	s.MeetingCodeExpire = &v
	return s
}

func (s *GetMeetingResponseBodyMeetingInfo) SetMeetingName(v string) *GetMeetingResponseBodyMeetingInfo {
	s.MeetingName = &v
	return s
}

func (s *GetMeetingResponseBodyMeetingInfo) SetMeetingUUID(v string) *GetMeetingResponseBodyMeetingInfo {
	s.MeetingUUID = &v
	return s
}

func (s *GetMeetingResponseBodyMeetingInfo) SetMemberList(v []*GetMeetingResponseBodyMeetingInfoMemberList) *GetMeetingResponseBodyMeetingInfo {
	s.MemberList = v
	return s
}

func (s *GetMeetingResponseBodyMeetingInfo) SetOpenPasswordFlag(v bool) *GetMeetingResponseBodyMeetingInfo {
	s.OpenPasswordFlag = &v
	return s
}

func (s *GetMeetingResponseBodyMeetingInfo) SetPassword(v string) *GetMeetingResponseBodyMeetingInfo {
	s.Password = &v
	return s
}

func (s *GetMeetingResponseBodyMeetingInfo) SetStartTime(v int64) *GetMeetingResponseBodyMeetingInfo {
	s.StartTime = &v
	return s
}

func (s *GetMeetingResponseBodyMeetingInfo) SetStatus(v int32) *GetMeetingResponseBodyMeetingInfo {
	s.Status = &v
	return s
}

func (s *GetMeetingResponseBodyMeetingInfo) SetUserId(v string) *GetMeetingResponseBodyMeetingInfo {
	s.UserId = &v
	return s
}

func (s *GetMeetingResponseBodyMeetingInfo) SetValidTime(v int64) *GetMeetingResponseBodyMeetingInfo {
	s.ValidTime = &v
	return s
}

type GetMeetingResponseBodyMeetingInfoMemberList struct {
	CreatorFlag   *bool   `json:"CreatorFlag,omitempty" xml:"CreatorFlag,omitempty"`
	MasterFlag    *bool   `json:"MasterFlag,omitempty" xml:"MasterFlag,omitempty"`
	MemberUUID    *string `json:"MemberUUID,omitempty" xml:"MemberUUID,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserAvatarUrl *string `json:"UserAvatarUrl,omitempty" xml:"UserAvatarUrl,omitempty"`
	UserDeptName  *string `json:"UserDeptName,omitempty" xml:"UserDeptName,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName      *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetMeetingResponseBodyMeetingInfoMemberList) String() string {
	return tea.Prettify(s)
}

func (s GetMeetingResponseBodyMeetingInfoMemberList) GoString() string {
	return s.String()
}

func (s *GetMeetingResponseBodyMeetingInfoMemberList) SetCreatorFlag(v bool) *GetMeetingResponseBodyMeetingInfoMemberList {
	s.CreatorFlag = &v
	return s
}

func (s *GetMeetingResponseBodyMeetingInfoMemberList) SetMasterFlag(v bool) *GetMeetingResponseBodyMeetingInfoMemberList {
	s.MasterFlag = &v
	return s
}

func (s *GetMeetingResponseBodyMeetingInfoMemberList) SetMemberUUID(v string) *GetMeetingResponseBodyMeetingInfoMemberList {
	s.MemberUUID = &v
	return s
}

func (s *GetMeetingResponseBodyMeetingInfoMemberList) SetStatus(v string) *GetMeetingResponseBodyMeetingInfoMemberList {
	s.Status = &v
	return s
}

func (s *GetMeetingResponseBodyMeetingInfoMemberList) SetUserAvatarUrl(v string) *GetMeetingResponseBodyMeetingInfoMemberList {
	s.UserAvatarUrl = &v
	return s
}

func (s *GetMeetingResponseBodyMeetingInfoMemberList) SetUserDeptName(v string) *GetMeetingResponseBodyMeetingInfoMemberList {
	s.UserDeptName = &v
	return s
}

func (s *GetMeetingResponseBodyMeetingInfoMemberList) SetUserId(v string) *GetMeetingResponseBodyMeetingInfoMemberList {
	s.UserId = &v
	return s
}

func (s *GetMeetingResponseBodyMeetingInfoMemberList) SetUserName(v string) *GetMeetingResponseBodyMeetingInfoMemberList {
	s.UserName = &v
	return s
}

type GetMeetingResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMeetingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMeetingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMeetingResponse) GoString() string {
	return s.String()
}

func (s *GetMeetingResponse) SetHeaders(v map[string]*string) *GetMeetingResponse {
	s.Headers = v
	return s
}

func (s *GetMeetingResponse) SetStatusCode(v int32) *GetMeetingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMeetingResponse) SetBody(v *GetMeetingResponseBody) *GetMeetingResponse {
	s.Body = v
	return s
}

type GetMeetingConcurrencyRequest struct {
	MeetingUUID *string `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
	TenantCode  *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
}

func (s GetMeetingConcurrencyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMeetingConcurrencyRequest) GoString() string {
	return s.String()
}

func (s *GetMeetingConcurrencyRequest) SetMeetingUUID(v string) *GetMeetingConcurrencyRequest {
	s.MeetingUUID = &v
	return s
}

func (s *GetMeetingConcurrencyRequest) SetTenantCode(v string) *GetMeetingConcurrencyRequest {
	s.TenantCode = &v
	return s
}

type GetMeetingConcurrencyResponseBody struct {
	Data      *GetMeetingConcurrencyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *int32                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMeetingConcurrencyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMeetingConcurrencyResponseBody) GoString() string {
	return s.String()
}

func (s *GetMeetingConcurrencyResponseBody) SetData(v *GetMeetingConcurrencyResponseBodyData) *GetMeetingConcurrencyResponseBody {
	s.Data = v
	return s
}

func (s *GetMeetingConcurrencyResponseBody) SetErrorCode(v int32) *GetMeetingConcurrencyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMeetingConcurrencyResponseBody) SetMessage(v string) *GetMeetingConcurrencyResponseBody {
	s.Message = &v
	return s
}

func (s *GetMeetingConcurrencyResponseBody) SetRequestId(v string) *GetMeetingConcurrencyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMeetingConcurrencyResponseBody) SetSuccess(v bool) *GetMeetingConcurrencyResponseBody {
	s.Success = &v
	return s
}

type GetMeetingConcurrencyResponseBodyData struct {
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MeetingCode  *string `json:"MeetingCode,omitempty" xml:"MeetingCode,omitempty"`
	MeetingName  *string `json:"MeetingName,omitempty" xml:"MeetingName,omitempty"`
	OfflineCount *int32  `json:"OfflineCount,omitempty" xml:"OfflineCount,omitempty"`
	OnlineCount  *int32  `json:"OnlineCount,omitempty" xml:"OnlineCount,omitempty"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetMeetingConcurrencyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMeetingConcurrencyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMeetingConcurrencyResponseBodyData) SetEndTime(v int64) *GetMeetingConcurrencyResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetMeetingConcurrencyResponseBodyData) SetMeetingCode(v string) *GetMeetingConcurrencyResponseBodyData {
	s.MeetingCode = &v
	return s
}

func (s *GetMeetingConcurrencyResponseBodyData) SetMeetingName(v string) *GetMeetingConcurrencyResponseBodyData {
	s.MeetingName = &v
	return s
}

func (s *GetMeetingConcurrencyResponseBodyData) SetOfflineCount(v int32) *GetMeetingConcurrencyResponseBodyData {
	s.OfflineCount = &v
	return s
}

func (s *GetMeetingConcurrencyResponseBodyData) SetOnlineCount(v int32) *GetMeetingConcurrencyResponseBodyData {
	s.OnlineCount = &v
	return s
}

func (s *GetMeetingConcurrencyResponseBodyData) SetStartTime(v int64) *GetMeetingConcurrencyResponseBodyData {
	s.StartTime = &v
	return s
}

type GetMeetingConcurrencyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMeetingConcurrencyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMeetingConcurrencyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMeetingConcurrencyResponse) GoString() string {
	return s.String()
}

func (s *GetMeetingConcurrencyResponse) SetHeaders(v map[string]*string) *GetMeetingConcurrencyResponse {
	s.Headers = v
	return s
}

func (s *GetMeetingConcurrencyResponse) SetStatusCode(v int32) *GetMeetingConcurrencyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMeetingConcurrencyResponse) SetBody(v *GetMeetingConcurrencyResponseBody) *GetMeetingConcurrencyResponse {
	s.Body = v
	return s
}

type GetMeetingInternationalRequest struct {
	MeetingUUID *string `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
}

func (s GetMeetingInternationalRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMeetingInternationalRequest) GoString() string {
	return s.String()
}

func (s *GetMeetingInternationalRequest) SetMeetingUUID(v string) *GetMeetingInternationalRequest {
	s.MeetingUUID = &v
	return s
}

type GetMeetingInternationalResponseBody struct {
	ErrorCode   *int32                                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	MeetingInfo *GetMeetingInternationalResponseBodyMeetingInfo `json:"MeetingInfo,omitempty" xml:"MeetingInfo,omitempty" type:"Struct"`
	Message     *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMeetingInternationalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMeetingInternationalResponseBody) GoString() string {
	return s.String()
}

func (s *GetMeetingInternationalResponseBody) SetErrorCode(v int32) *GetMeetingInternationalResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMeetingInternationalResponseBody) SetMeetingInfo(v *GetMeetingInternationalResponseBodyMeetingInfo) *GetMeetingInternationalResponseBody {
	s.MeetingInfo = v
	return s
}

func (s *GetMeetingInternationalResponseBody) SetMessage(v string) *GetMeetingInternationalResponseBody {
	s.Message = &v
	return s
}

func (s *GetMeetingInternationalResponseBody) SetRequestId(v string) *GetMeetingInternationalResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMeetingInternationalResponseBody) SetSuccess(v bool) *GetMeetingInternationalResponseBody {
	s.Success = &v
	return s
}

type GetMeetingInternationalResponseBodyMeetingInfo struct {
	CreateTime  *int64                                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	MeetingCode *string                                                     `json:"MeetingCode,omitempty" xml:"MeetingCode,omitempty"`
	MeetingName *string                                                     `json:"MeetingName,omitempty" xml:"MeetingName,omitempty"`
	MeetingUUID *string                                                     `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
	MemberList  []*GetMeetingInternationalResponseBodyMeetingInfoMemberList `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
	Password    *string                                                     `json:"Password,omitempty" xml:"Password,omitempty"`
	UserId      *string                                                     `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ValidTime   *int64                                                      `json:"ValidTime,omitempty" xml:"ValidTime,omitempty"`
}

func (s GetMeetingInternationalResponseBodyMeetingInfo) String() string {
	return tea.Prettify(s)
}

func (s GetMeetingInternationalResponseBodyMeetingInfo) GoString() string {
	return s.String()
}

func (s *GetMeetingInternationalResponseBodyMeetingInfo) SetCreateTime(v int64) *GetMeetingInternationalResponseBodyMeetingInfo {
	s.CreateTime = &v
	return s
}

func (s *GetMeetingInternationalResponseBodyMeetingInfo) SetMeetingCode(v string) *GetMeetingInternationalResponseBodyMeetingInfo {
	s.MeetingCode = &v
	return s
}

func (s *GetMeetingInternationalResponseBodyMeetingInfo) SetMeetingName(v string) *GetMeetingInternationalResponseBodyMeetingInfo {
	s.MeetingName = &v
	return s
}

func (s *GetMeetingInternationalResponseBodyMeetingInfo) SetMeetingUUID(v string) *GetMeetingInternationalResponseBodyMeetingInfo {
	s.MeetingUUID = &v
	return s
}

func (s *GetMeetingInternationalResponseBodyMeetingInfo) SetMemberList(v []*GetMeetingInternationalResponseBodyMeetingInfoMemberList) *GetMeetingInternationalResponseBodyMeetingInfo {
	s.MemberList = v
	return s
}

func (s *GetMeetingInternationalResponseBodyMeetingInfo) SetPassword(v string) *GetMeetingInternationalResponseBodyMeetingInfo {
	s.Password = &v
	return s
}

func (s *GetMeetingInternationalResponseBodyMeetingInfo) SetUserId(v string) *GetMeetingInternationalResponseBodyMeetingInfo {
	s.UserId = &v
	return s
}

func (s *GetMeetingInternationalResponseBodyMeetingInfo) SetValidTime(v int64) *GetMeetingInternationalResponseBodyMeetingInfo {
	s.ValidTime = &v
	return s
}

type GetMeetingInternationalResponseBodyMeetingInfoMemberList struct {
	MemberUUID    *string `json:"MemberUUID,omitempty" xml:"MemberUUID,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserAvatarUrl *string `json:"UserAvatarUrl,omitempty" xml:"UserAvatarUrl,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName      *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetMeetingInternationalResponseBodyMeetingInfoMemberList) String() string {
	return tea.Prettify(s)
}

func (s GetMeetingInternationalResponseBodyMeetingInfoMemberList) GoString() string {
	return s.String()
}

func (s *GetMeetingInternationalResponseBodyMeetingInfoMemberList) SetMemberUUID(v string) *GetMeetingInternationalResponseBodyMeetingInfoMemberList {
	s.MemberUUID = &v
	return s
}

func (s *GetMeetingInternationalResponseBodyMeetingInfoMemberList) SetStatus(v string) *GetMeetingInternationalResponseBodyMeetingInfoMemberList {
	s.Status = &v
	return s
}

func (s *GetMeetingInternationalResponseBodyMeetingInfoMemberList) SetUserAvatarUrl(v string) *GetMeetingInternationalResponseBodyMeetingInfoMemberList {
	s.UserAvatarUrl = &v
	return s
}

func (s *GetMeetingInternationalResponseBodyMeetingInfoMemberList) SetUserId(v string) *GetMeetingInternationalResponseBodyMeetingInfoMemberList {
	s.UserId = &v
	return s
}

func (s *GetMeetingInternationalResponseBodyMeetingInfoMemberList) SetUserName(v string) *GetMeetingInternationalResponseBodyMeetingInfoMemberList {
	s.UserName = &v
	return s
}

type GetMeetingInternationalResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMeetingInternationalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMeetingInternationalResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMeetingInternationalResponse) GoString() string {
	return s.String()
}

func (s *GetMeetingInternationalResponse) SetHeaders(v map[string]*string) *GetMeetingInternationalResponse {
	s.Headers = v
	return s
}

func (s *GetMeetingInternationalResponse) SetStatusCode(v int32) *GetMeetingInternationalResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMeetingInternationalResponse) SetBody(v *GetMeetingInternationalResponseBody) *GetMeetingInternationalResponse {
	s.Body = v
	return s
}

type GetMeetingMemberRequest struct {
	MeetingUUID *string `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
}

func (s GetMeetingMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMeetingMemberRequest) GoString() string {
	return s.String()
}

func (s *GetMeetingMemberRequest) SetMeetingUUID(v string) *GetMeetingMemberRequest {
	s.MeetingUUID = &v
	return s
}

type GetMeetingMemberResponseBody struct {
	ErrorCode *int32                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Members   map[string]interface{} `json:"Members,omitempty" xml:"Members,omitempty"`
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMeetingMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMeetingMemberResponseBody) GoString() string {
	return s.String()
}

func (s *GetMeetingMemberResponseBody) SetErrorCode(v int32) *GetMeetingMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMeetingMemberResponseBody) SetMembers(v map[string]interface{}) *GetMeetingMemberResponseBody {
	s.Members = v
	return s
}

func (s *GetMeetingMemberResponseBody) SetMessage(v string) *GetMeetingMemberResponseBody {
	s.Message = &v
	return s
}

func (s *GetMeetingMemberResponseBody) SetRequestId(v string) *GetMeetingMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMeetingMemberResponseBody) SetSuccess(v bool) *GetMeetingMemberResponseBody {
	s.Success = &v
	return s
}

type GetMeetingMemberResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMeetingMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMeetingMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMeetingMemberResponse) GoString() string {
	return s.String()
}

func (s *GetMeetingMemberResponse) SetHeaders(v map[string]*string) *GetMeetingMemberResponse {
	s.Headers = v
	return s
}

func (s *GetMeetingMemberResponse) SetStatusCode(v int32) *GetMeetingMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMeetingMemberResponse) SetBody(v *GetMeetingMemberResponseBody) *GetMeetingMemberResponse {
	s.Body = v
	return s
}

type GetMemberStatusRequest struct {
	MeetingUUID *string `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetMemberStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMemberStatusRequest) GoString() string {
	return s.String()
}

func (s *GetMemberStatusRequest) SetMeetingUUID(v string) *GetMemberStatusRequest {
	s.MeetingUUID = &v
	return s
}

func (s *GetMemberStatusRequest) SetUserId(v string) *GetMemberStatusRequest {
	s.UserId = &v
	return s
}

type GetMemberStatusResponseBody struct {
	ErrorCode *int64  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMemberStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMemberStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemberStatusResponseBody) SetErrorCode(v int64) *GetMemberStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMemberStatusResponseBody) SetMessage(v string) *GetMemberStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetMemberStatusResponseBody) SetRequestId(v string) *GetMemberStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemberStatusResponseBody) SetStatus(v string) *GetMemberStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetMemberStatusResponseBody) SetSuccess(v bool) *GetMemberStatusResponseBody {
	s.Success = &v
	return s
}

type GetMemberStatusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMemberStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMemberStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMemberStatusResponse) GoString() string {
	return s.String()
}

func (s *GetMemberStatusResponse) SetHeaders(v map[string]*string) *GetMemberStatusResponse {
	s.Headers = v
	return s
}

func (s *GetMemberStatusResponse) SetStatusCode(v int32) *GetMemberStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMemberStatusResponse) SetBody(v *GetMemberStatusResponseBody) *GetMemberStatusResponse {
	s.Body = v
	return s
}

type GetScreenVerificationCodeRequest struct {
	CastScreenCode *string `json:"CastScreenCode,omitempty" xml:"CastScreenCode,omitempty"`
}

func (s GetScreenVerificationCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetScreenVerificationCodeRequest) GoString() string {
	return s.String()
}

func (s *GetScreenVerificationCodeRequest) SetCastScreenCode(v string) *GetScreenVerificationCodeRequest {
	s.CastScreenCode = &v
	return s
}

type GetScreenVerificationCodeResponseBody struct {
	ErrorCode              *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message                *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScreenVerificationCode *string `json:"ScreenVerificationCode,omitempty" xml:"ScreenVerificationCode,omitempty"`
	Success                *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetScreenVerificationCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetScreenVerificationCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetScreenVerificationCodeResponseBody) SetErrorCode(v int32) *GetScreenVerificationCodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetScreenVerificationCodeResponseBody) SetMessage(v string) *GetScreenVerificationCodeResponseBody {
	s.Message = &v
	return s
}

func (s *GetScreenVerificationCodeResponseBody) SetRequestId(v string) *GetScreenVerificationCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScreenVerificationCodeResponseBody) SetScreenVerificationCode(v string) *GetScreenVerificationCodeResponseBody {
	s.ScreenVerificationCode = &v
	return s
}

func (s *GetScreenVerificationCodeResponseBody) SetSuccess(v bool) *GetScreenVerificationCodeResponseBody {
	s.Success = &v
	return s
}

type GetScreenVerificationCodeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetScreenVerificationCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetScreenVerificationCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetScreenVerificationCodeResponse) GoString() string {
	return s.String()
}

func (s *GetScreenVerificationCodeResponse) SetHeaders(v map[string]*string) *GetScreenVerificationCodeResponse {
	s.Headers = v
	return s
}

func (s *GetScreenVerificationCodeResponse) SetStatusCode(v int32) *GetScreenVerificationCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScreenVerificationCodeResponse) SetBody(v *GetScreenVerificationCodeResponseBody) *GetScreenVerificationCodeResponse {
	s.Body = v
	return s
}

type GetStatisticRequest struct {
	EndTime   *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetStatisticRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStatisticRequest) GoString() string {
	return s.String()
}

func (s *GetStatisticRequest) SetEndTime(v int64) *GetStatisticRequest {
	s.EndTime = &v
	return s
}

func (s *GetStatisticRequest) SetStartTime(v int64) *GetStatisticRequest {
	s.StartTime = &v
	return s
}

type GetStatisticResponseBody struct {
	ErrorCode     *int32                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message       *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId     *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StatisticInfo *GetStatisticResponseBodyStatisticInfo `json:"StatisticInfo,omitempty" xml:"StatisticInfo,omitempty" type:"Struct"`
	Success       *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetStatisticResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *GetStatisticResponseBody) SetErrorCode(v int32) *GetStatisticResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetStatisticResponseBody) SetMessage(v string) *GetStatisticResponseBody {
	s.Message = &v
	return s
}

func (s *GetStatisticResponseBody) SetRequestId(v string) *GetStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStatisticResponseBody) SetStatisticInfo(v *GetStatisticResponseBodyStatisticInfo) *GetStatisticResponseBody {
	s.StatisticInfo = v
	return s
}

func (s *GetStatisticResponseBody) SetSuccess(v bool) *GetStatisticResponseBody {
	s.Success = &v
	return s
}

type GetStatisticResponseBodyStatisticInfo struct {
	MaxConcurrency  *int64 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	MeetingDuration *int64 `json:"MeetingDuration,omitempty" xml:"MeetingDuration,omitempty"`
	MeetingNumber   *int64 `json:"MeetingNumber,omitempty" xml:"MeetingNumber,omitempty"`
	MemberNumber    *int64 `json:"MemberNumber,omitempty" xml:"MemberNumber,omitempty"`
}

func (s GetStatisticResponseBodyStatisticInfo) String() string {
	return tea.Prettify(s)
}

func (s GetStatisticResponseBodyStatisticInfo) GoString() string {
	return s.String()
}

func (s *GetStatisticResponseBodyStatisticInfo) SetMaxConcurrency(v int64) *GetStatisticResponseBodyStatisticInfo {
	s.MaxConcurrency = &v
	return s
}

func (s *GetStatisticResponseBodyStatisticInfo) SetMeetingDuration(v int64) *GetStatisticResponseBodyStatisticInfo {
	s.MeetingDuration = &v
	return s
}

func (s *GetStatisticResponseBodyStatisticInfo) SetMeetingNumber(v int64) *GetStatisticResponseBodyStatisticInfo {
	s.MeetingNumber = &v
	return s
}

func (s *GetStatisticResponseBodyStatisticInfo) SetMemberNumber(v int64) *GetStatisticResponseBodyStatisticInfo {
	s.MemberNumber = &v
	return s
}

type GetStatisticResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStatisticResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStatisticResponse) GoString() string {
	return s.String()
}

func (s *GetStatisticResponse) SetHeaders(v map[string]*string) *GetStatisticResponse {
	s.Headers = v
	return s
}

func (s *GetStatisticResponse) SetStatusCode(v int32) *GetStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStatisticResponse) SetBody(v *GetStatisticResponseBody) *GetStatisticResponse {
	s.Body = v
	return s
}

type GetUserRequest struct {
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) SetUserId(v string) *GetUserRequest {
	s.UserId = &v
	return s
}

type GetUserResponseBody struct {
	ErrorCode *int32                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
	UserInfo  *GetUserResponseBodyUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) SetErrorCode(v int32) *GetUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetUserResponseBody) SetMessage(v string) *GetUserResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserResponseBody) SetSuccess(v bool) *GetUserResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserResponseBody) SetUserInfo(v *GetUserResponseBodyUserInfo) *GetUserResponseBody {
	s.UserInfo = v
	return s
}

type GetUserResponseBodyUserInfo struct {
	CreateTime    *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DepartId      *string `json:"DepartId,omitempty" xml:"DepartId,omitempty"`
	DepartName    *string `json:"DepartName,omitempty" xml:"DepartName,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	JobName       *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	UserAvatarUrl *string `json:"UserAvatarUrl,omitempty" xml:"UserAvatarUrl,omitempty"`
	UserEmail     *string `json:"UserEmail,omitempty" xml:"UserEmail,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserMobile    *string `json:"UserMobile,omitempty" xml:"UserMobile,omitempty"`
	UserName      *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserTel       *string `json:"UserTel,omitempty" xml:"UserTel,omitempty"`
}

func (s GetUserResponseBodyUserInfo) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyUserInfo) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUserInfo) SetCreateTime(v int64) *GetUserResponseBodyUserInfo {
	s.CreateTime = &v
	return s
}

func (s *GetUserResponseBodyUserInfo) SetDepartId(v string) *GetUserResponseBodyUserInfo {
	s.DepartId = &v
	return s
}

func (s *GetUserResponseBodyUserInfo) SetDepartName(v string) *GetUserResponseBodyUserInfo {
	s.DepartName = &v
	return s
}

func (s *GetUserResponseBodyUserInfo) SetGroupId(v string) *GetUserResponseBodyUserInfo {
	s.GroupId = &v
	return s
}

func (s *GetUserResponseBodyUserInfo) SetGroupName(v string) *GetUserResponseBodyUserInfo {
	s.GroupName = &v
	return s
}

func (s *GetUserResponseBodyUserInfo) SetJobName(v string) *GetUserResponseBodyUserInfo {
	s.JobName = &v
	return s
}

func (s *GetUserResponseBodyUserInfo) SetUserAvatarUrl(v string) *GetUserResponseBodyUserInfo {
	s.UserAvatarUrl = &v
	return s
}

func (s *GetUserResponseBodyUserInfo) SetUserEmail(v string) *GetUserResponseBodyUserInfo {
	s.UserEmail = &v
	return s
}

func (s *GetUserResponseBodyUserInfo) SetUserId(v string) *GetUserResponseBodyUserInfo {
	s.UserId = &v
	return s
}

func (s *GetUserResponseBodyUserInfo) SetUserMobile(v string) *GetUserResponseBodyUserInfo {
	s.UserMobile = &v
	return s
}

func (s *GetUserResponseBodyUserInfo) SetUserName(v string) *GetUserResponseBodyUserInfo {
	s.UserName = &v
	return s
}

func (s *GetUserResponseBodyUserInfo) SetUserTel(v string) *GetUserResponseBodyUserInfo {
	s.UserTel = &v
	return s
}

type GetUserResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponse) GoString() string {
	return s.String()
}

func (s *GetUserResponse) SetHeaders(v map[string]*string) *GetUserResponse {
	s.Headers = v
	return s
}

func (s *GetUserResponse) SetStatusCode(v int32) *GetUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserResponse) SetBody(v *GetUserResponseBody) *GetUserResponse {
	s.Body = v
	return s
}

type GetWebSocketTokenRequest struct {
	OldToken  *string `json:"OldToken,omitempty" xml:"OldToken,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GetWebSocketTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWebSocketTokenRequest) GoString() string {
	return s.String()
}

func (s *GetWebSocketTokenRequest) SetOldToken(v string) *GetWebSocketTokenRequest {
	s.OldToken = &v
	return s
}

func (s *GetWebSocketTokenRequest) SetSessionId(v string) *GetWebSocketTokenRequest {
	s.SessionId = &v
	return s
}

type GetWebSocketTokenResponseBody struct {
	Data      *GetWebSocketTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *int32                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWebSocketTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWebSocketTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetWebSocketTokenResponseBody) SetData(v *GetWebSocketTokenResponseBodyData) *GetWebSocketTokenResponseBody {
	s.Data = v
	return s
}

func (s *GetWebSocketTokenResponseBody) SetErrorCode(v int32) *GetWebSocketTokenResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetWebSocketTokenResponseBody) SetMessage(v string) *GetWebSocketTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetWebSocketTokenResponseBody) SetRequestId(v string) *GetWebSocketTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWebSocketTokenResponseBody) SetSuccess(v bool) *GetWebSocketTokenResponseBody {
	s.Success = &v
	return s
}

type GetWebSocketTokenResponseBodyData struct {
	AuthWsUrl         *string `json:"AuthWsUrl,omitempty" xml:"AuthWsUrl,omitempty"`
	Token             *string `json:"Token,omitempty" xml:"Token,omitempty"`
	WsOuterReConnTime *string `json:"WsOuterReConnTime,omitempty" xml:"WsOuterReConnTime,omitempty"`
}

func (s GetWebSocketTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetWebSocketTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWebSocketTokenResponseBodyData) SetAuthWsUrl(v string) *GetWebSocketTokenResponseBodyData {
	s.AuthWsUrl = &v
	return s
}

func (s *GetWebSocketTokenResponseBodyData) SetToken(v string) *GetWebSocketTokenResponseBodyData {
	s.Token = &v
	return s
}

func (s *GetWebSocketTokenResponseBodyData) SetWsOuterReConnTime(v string) *GetWebSocketTokenResponseBodyData {
	s.WsOuterReConnTime = &v
	return s
}

type GetWebSocketTokenResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetWebSocketTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWebSocketTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWebSocketTokenResponse) GoString() string {
	return s.String()
}

func (s *GetWebSocketTokenResponse) SetHeaders(v map[string]*string) *GetWebSocketTokenResponse {
	s.Headers = v
	return s
}

func (s *GetWebSocketTokenResponse) SetStatusCode(v int32) *GetWebSocketTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWebSocketTokenResponse) SetBody(v *GetWebSocketTokenResponseBody) *GetWebSocketTokenResponse {
	s.Body = v
	return s
}

type InviteUserRequest struct {
	MeetingUUID *string `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
	UserIds     *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s InviteUserRequest) String() string {
	return tea.Prettify(s)
}

func (s InviteUserRequest) GoString() string {
	return s.String()
}

func (s *InviteUserRequest) SetMeetingUUID(v string) *InviteUserRequest {
	s.MeetingUUID = &v
	return s
}

func (s *InviteUserRequest) SetUserIds(v string) *InviteUserRequest {
	s.UserIds = &v
	return s
}

type InviteUserResponseBody struct {
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InviteUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InviteUserResponseBody) GoString() string {
	return s.String()
}

func (s *InviteUserResponseBody) SetErrorCode(v int32) *InviteUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *InviteUserResponseBody) SetMessage(v string) *InviteUserResponseBody {
	s.Message = &v
	return s
}

func (s *InviteUserResponseBody) SetRequestId(v string) *InviteUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *InviteUserResponseBody) SetSuccess(v bool) *InviteUserResponseBody {
	s.Success = &v
	return s
}

type InviteUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InviteUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InviteUserResponse) String() string {
	return tea.Prettify(s)
}

func (s InviteUserResponse) GoString() string {
	return s.String()
}

func (s *InviteUserResponse) SetHeaders(v map[string]*string) *InviteUserResponse {
	s.Headers = v
	return s
}

func (s *InviteUserResponse) SetStatusCode(v int32) *InviteUserResponse {
	s.StatusCode = &v
	return s
}

func (s *InviteUserResponse) SetBody(v *InviteUserResponseBody) *InviteUserResponse {
	s.Body = v
	return s
}

type JoinDeviceMeetingRequest struct {
	MeetingCode *string `json:"MeetingCode,omitempty" xml:"MeetingCode,omitempty"`
	Password    *string `json:"Password,omitempty" xml:"Password,omitempty"`
	SN          *string `json:"SN,omitempty" xml:"SN,omitempty"`
	Token       *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s JoinDeviceMeetingRequest) String() string {
	return tea.Prettify(s)
}

func (s JoinDeviceMeetingRequest) GoString() string {
	return s.String()
}

func (s *JoinDeviceMeetingRequest) SetMeetingCode(v string) *JoinDeviceMeetingRequest {
	s.MeetingCode = &v
	return s
}

func (s *JoinDeviceMeetingRequest) SetPassword(v string) *JoinDeviceMeetingRequest {
	s.Password = &v
	return s
}

func (s *JoinDeviceMeetingRequest) SetSN(v string) *JoinDeviceMeetingRequest {
	s.SN = &v
	return s
}

func (s *JoinDeviceMeetingRequest) SetToken(v string) *JoinDeviceMeetingRequest {
	s.Token = &v
	return s
}

type JoinDeviceMeetingResponseBody struct {
	Device    *JoinDeviceMeetingResponseBodyDevice `json:"Device,omitempty" xml:"Device,omitempty" type:"Struct"`
	ErrorCode *int32                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s JoinDeviceMeetingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s JoinDeviceMeetingResponseBody) GoString() string {
	return s.String()
}

func (s *JoinDeviceMeetingResponseBody) SetDevice(v *JoinDeviceMeetingResponseBodyDevice) *JoinDeviceMeetingResponseBody {
	s.Device = v
	return s
}

func (s *JoinDeviceMeetingResponseBody) SetErrorCode(v int32) *JoinDeviceMeetingResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *JoinDeviceMeetingResponseBody) SetMessage(v string) *JoinDeviceMeetingResponseBody {
	s.Message = &v
	return s
}

func (s *JoinDeviceMeetingResponseBody) SetRequestId(v string) *JoinDeviceMeetingResponseBody {
	s.RequestId = &v
	return s
}

func (s *JoinDeviceMeetingResponseBody) SetSuccess(v bool) *JoinDeviceMeetingResponseBody {
	s.Success = &v
	return s
}

type JoinDeviceMeetingResponseBodyDevice struct {
	ClientAppId   *string                                     `json:"ClientAppId,omitempty" xml:"ClientAppId,omitempty"`
	MeetingAppId  *string                                     `json:"MeetingAppId,omitempty" xml:"MeetingAppId,omitempty"`
	MeetingCode   *string                                     `json:"MeetingCode,omitempty" xml:"MeetingCode,omitempty"`
	MeetingDomain *string                                     `json:"MeetingDomain,omitempty" xml:"MeetingDomain,omitempty"`
	MeetingToken  *string                                     `json:"MeetingToken,omitempty" xml:"MeetingToken,omitempty"`
	MeetingUUID   *string                                     `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
	MemberUUID    *string                                     `json:"MemberUUID,omitempty" xml:"MemberUUID,omitempty"`
	SlsInfo       *JoinDeviceMeetingResponseBodyDeviceSlsInfo `json:"SlsInfo,omitempty" xml:"SlsInfo,omitempty" type:"Struct"`
}

func (s JoinDeviceMeetingResponseBodyDevice) String() string {
	return tea.Prettify(s)
}

func (s JoinDeviceMeetingResponseBodyDevice) GoString() string {
	return s.String()
}

func (s *JoinDeviceMeetingResponseBodyDevice) SetClientAppId(v string) *JoinDeviceMeetingResponseBodyDevice {
	s.ClientAppId = &v
	return s
}

func (s *JoinDeviceMeetingResponseBodyDevice) SetMeetingAppId(v string) *JoinDeviceMeetingResponseBodyDevice {
	s.MeetingAppId = &v
	return s
}

func (s *JoinDeviceMeetingResponseBodyDevice) SetMeetingCode(v string) *JoinDeviceMeetingResponseBodyDevice {
	s.MeetingCode = &v
	return s
}

func (s *JoinDeviceMeetingResponseBodyDevice) SetMeetingDomain(v string) *JoinDeviceMeetingResponseBodyDevice {
	s.MeetingDomain = &v
	return s
}

func (s *JoinDeviceMeetingResponseBodyDevice) SetMeetingToken(v string) *JoinDeviceMeetingResponseBodyDevice {
	s.MeetingToken = &v
	return s
}

func (s *JoinDeviceMeetingResponseBodyDevice) SetMeetingUUID(v string) *JoinDeviceMeetingResponseBodyDevice {
	s.MeetingUUID = &v
	return s
}

func (s *JoinDeviceMeetingResponseBodyDevice) SetMemberUUID(v string) *JoinDeviceMeetingResponseBodyDevice {
	s.MemberUUID = &v
	return s
}

func (s *JoinDeviceMeetingResponseBodyDevice) SetSlsInfo(v *JoinDeviceMeetingResponseBodyDeviceSlsInfo) *JoinDeviceMeetingResponseBodyDevice {
	s.SlsInfo = v
	return s
}

type JoinDeviceMeetingResponseBodyDeviceSlsInfo struct {
	LogServiceEndpoint *string `json:"LogServiceEndpoint,omitempty" xml:"LogServiceEndpoint,omitempty"`
	Logstore           *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	Project            *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s JoinDeviceMeetingResponseBodyDeviceSlsInfo) String() string {
	return tea.Prettify(s)
}

func (s JoinDeviceMeetingResponseBodyDeviceSlsInfo) GoString() string {
	return s.String()
}

func (s *JoinDeviceMeetingResponseBodyDeviceSlsInfo) SetLogServiceEndpoint(v string) *JoinDeviceMeetingResponseBodyDeviceSlsInfo {
	s.LogServiceEndpoint = &v
	return s
}

func (s *JoinDeviceMeetingResponseBodyDeviceSlsInfo) SetLogstore(v string) *JoinDeviceMeetingResponseBodyDeviceSlsInfo {
	s.Logstore = &v
	return s
}

func (s *JoinDeviceMeetingResponseBodyDeviceSlsInfo) SetProject(v string) *JoinDeviceMeetingResponseBodyDeviceSlsInfo {
	s.Project = &v
	return s
}

type JoinDeviceMeetingResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *JoinDeviceMeetingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s JoinDeviceMeetingResponse) String() string {
	return tea.Prettify(s)
}

func (s JoinDeviceMeetingResponse) GoString() string {
	return s.String()
}

func (s *JoinDeviceMeetingResponse) SetHeaders(v map[string]*string) *JoinDeviceMeetingResponse {
	s.Headers = v
	return s
}

func (s *JoinDeviceMeetingResponse) SetStatusCode(v int32) *JoinDeviceMeetingResponse {
	s.StatusCode = &v
	return s
}

func (s *JoinDeviceMeetingResponse) SetBody(v *JoinDeviceMeetingResponseBody) *JoinDeviceMeetingResponse {
	s.Body = v
	return s
}

type JoinLiveRequest struct {
	LiveUUID *string `json:"LiveUUID,omitempty" xml:"LiveUUID,omitempty"`
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s JoinLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s JoinLiveRequest) GoString() string {
	return s.String()
}

func (s *JoinLiveRequest) SetLiveUUID(v string) *JoinLiveRequest {
	s.LiveUUID = &v
	return s
}

func (s *JoinLiveRequest) SetPassword(v string) *JoinLiveRequest {
	s.Password = &v
	return s
}

func (s *JoinLiveRequest) SetUserId(v string) *JoinLiveRequest {
	s.UserId = &v
	return s
}

type JoinLiveResponseBody struct {
	ErrorCode   *int32                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	MeetingInfo *JoinLiveResponseBodyMeetingInfo `json:"MeetingInfo,omitempty" xml:"MeetingInfo,omitempty" type:"Struct"`
	Message     *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s JoinLiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s JoinLiveResponseBody) GoString() string {
	return s.String()
}

func (s *JoinLiveResponseBody) SetErrorCode(v int32) *JoinLiveResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *JoinLiveResponseBody) SetMeetingInfo(v *JoinLiveResponseBodyMeetingInfo) *JoinLiveResponseBody {
	s.MeetingInfo = v
	return s
}

func (s *JoinLiveResponseBody) SetMessage(v string) *JoinLiveResponseBody {
	s.Message = &v
	return s
}

func (s *JoinLiveResponseBody) SetRequestId(v string) *JoinLiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *JoinLiveResponseBody) SetSuccess(v bool) *JoinLiveResponseBody {
	s.Success = &v
	return s
}

type JoinLiveResponseBodyMeetingInfo struct {
	ClientAppId   *string                                 `json:"ClientAppId,omitempty" xml:"ClientAppId,omitempty"`
	MeetingAppId  *string                                 `json:"MeetingAppId,omitempty" xml:"MeetingAppId,omitempty"`
	MeetingCode   *string                                 `json:"MeetingCode,omitempty" xml:"MeetingCode,omitempty"`
	MeetingDomain *string                                 `json:"MeetingDomain,omitempty" xml:"MeetingDomain,omitempty"`
	MeetingToken  *string                                 `json:"MeetingToken,omitempty" xml:"MeetingToken,omitempty"`
	MeetingUUID   *string                                 `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
	MemberUUID    *string                                 `json:"MemberUUID,omitempty" xml:"MemberUUID,omitempty"`
	SlsInfo       *JoinLiveResponseBodyMeetingInfoSlsInfo `json:"SlsInfo,omitempty" xml:"SlsInfo,omitempty" type:"Struct"`
}

func (s JoinLiveResponseBodyMeetingInfo) String() string {
	return tea.Prettify(s)
}

func (s JoinLiveResponseBodyMeetingInfo) GoString() string {
	return s.String()
}

func (s *JoinLiveResponseBodyMeetingInfo) SetClientAppId(v string) *JoinLiveResponseBodyMeetingInfo {
	s.ClientAppId = &v
	return s
}

func (s *JoinLiveResponseBodyMeetingInfo) SetMeetingAppId(v string) *JoinLiveResponseBodyMeetingInfo {
	s.MeetingAppId = &v
	return s
}

func (s *JoinLiveResponseBodyMeetingInfo) SetMeetingCode(v string) *JoinLiveResponseBodyMeetingInfo {
	s.MeetingCode = &v
	return s
}

func (s *JoinLiveResponseBodyMeetingInfo) SetMeetingDomain(v string) *JoinLiveResponseBodyMeetingInfo {
	s.MeetingDomain = &v
	return s
}

func (s *JoinLiveResponseBodyMeetingInfo) SetMeetingToken(v string) *JoinLiveResponseBodyMeetingInfo {
	s.MeetingToken = &v
	return s
}

func (s *JoinLiveResponseBodyMeetingInfo) SetMeetingUUID(v string) *JoinLiveResponseBodyMeetingInfo {
	s.MeetingUUID = &v
	return s
}

func (s *JoinLiveResponseBodyMeetingInfo) SetMemberUUID(v string) *JoinLiveResponseBodyMeetingInfo {
	s.MemberUUID = &v
	return s
}

func (s *JoinLiveResponseBodyMeetingInfo) SetSlsInfo(v *JoinLiveResponseBodyMeetingInfoSlsInfo) *JoinLiveResponseBodyMeetingInfo {
	s.SlsInfo = v
	return s
}

type JoinLiveResponseBodyMeetingInfoSlsInfo struct {
	LogServiceEndpoint *string `json:"LogServiceEndpoint,omitempty" xml:"LogServiceEndpoint,omitempty"`
	Logstore           *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	Project            *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s JoinLiveResponseBodyMeetingInfoSlsInfo) String() string {
	return tea.Prettify(s)
}

func (s JoinLiveResponseBodyMeetingInfoSlsInfo) GoString() string {
	return s.String()
}

func (s *JoinLiveResponseBodyMeetingInfoSlsInfo) SetLogServiceEndpoint(v string) *JoinLiveResponseBodyMeetingInfoSlsInfo {
	s.LogServiceEndpoint = &v
	return s
}

func (s *JoinLiveResponseBodyMeetingInfoSlsInfo) SetLogstore(v string) *JoinLiveResponseBodyMeetingInfoSlsInfo {
	s.Logstore = &v
	return s
}

func (s *JoinLiveResponseBodyMeetingInfoSlsInfo) SetProject(v string) *JoinLiveResponseBodyMeetingInfoSlsInfo {
	s.Project = &v
	return s
}

type JoinLiveResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *JoinLiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s JoinLiveResponse) String() string {
	return tea.Prettify(s)
}

func (s JoinLiveResponse) GoString() string {
	return s.String()
}

func (s *JoinLiveResponse) SetHeaders(v map[string]*string) *JoinLiveResponse {
	s.Headers = v
	return s
}

func (s *JoinLiveResponse) SetStatusCode(v int32) *JoinLiveResponse {
	s.StatusCode = &v
	return s
}

func (s *JoinLiveResponse) SetBody(v *JoinLiveResponseBody) *JoinLiveResponse {
	s.Body = v
	return s
}

type JoinMeetingRequest struct {
	MeetingCode *string `json:"MeetingCode,omitempty" xml:"MeetingCode,omitempty"`
	Password    *string `json:"Password,omitempty" xml:"Password,omitempty"`
	RtcEngine   *string `json:"RtcEngine,omitempty" xml:"RtcEngine,omitempty"`
	TenantCode  *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s JoinMeetingRequest) String() string {
	return tea.Prettify(s)
}

func (s JoinMeetingRequest) GoString() string {
	return s.String()
}

func (s *JoinMeetingRequest) SetMeetingCode(v string) *JoinMeetingRequest {
	s.MeetingCode = &v
	return s
}

func (s *JoinMeetingRequest) SetPassword(v string) *JoinMeetingRequest {
	s.Password = &v
	return s
}

func (s *JoinMeetingRequest) SetRtcEngine(v string) *JoinMeetingRequest {
	s.RtcEngine = &v
	return s
}

func (s *JoinMeetingRequest) SetTenantCode(v string) *JoinMeetingRequest {
	s.TenantCode = &v
	return s
}

func (s *JoinMeetingRequest) SetUserId(v string) *JoinMeetingRequest {
	s.UserId = &v
	return s
}

type JoinMeetingResponseBody struct {
	ErrorCode   *int64                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	MeetingInfo *JoinMeetingResponseBodyMeetingInfo `json:"MeetingInfo,omitempty" xml:"MeetingInfo,omitempty" type:"Struct"`
	Message     *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s JoinMeetingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s JoinMeetingResponseBody) GoString() string {
	return s.String()
}

func (s *JoinMeetingResponseBody) SetErrorCode(v int64) *JoinMeetingResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *JoinMeetingResponseBody) SetMeetingInfo(v *JoinMeetingResponseBodyMeetingInfo) *JoinMeetingResponseBody {
	s.MeetingInfo = v
	return s
}

func (s *JoinMeetingResponseBody) SetMessage(v string) *JoinMeetingResponseBody {
	s.Message = &v
	return s
}

func (s *JoinMeetingResponseBody) SetRequestId(v string) *JoinMeetingResponseBody {
	s.RequestId = &v
	return s
}

func (s *JoinMeetingResponseBody) SetSuccess(v bool) *JoinMeetingResponseBody {
	s.Success = &v
	return s
}

type JoinMeetingResponseBodyMeetingInfo struct {
	ClientAppId   *string                                     `json:"ClientAppId,omitempty" xml:"ClientAppId,omitempty"`
	EndTime       *int64                                      `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GrtnInfo      *JoinMeetingResponseBodyMeetingInfoGrtnInfo `json:"GrtnInfo,omitempty" xml:"GrtnInfo,omitempty" type:"Struct"`
	MeetingAppId  *string                                     `json:"MeetingAppId,omitempty" xml:"MeetingAppId,omitempty"`
	MeetingCode   *string                                     `json:"MeetingCode,omitempty" xml:"MeetingCode,omitempty"`
	MeetingDomain *string                                     `json:"MeetingDomain,omitempty" xml:"MeetingDomain,omitempty"`
	MeetingName   *string                                     `json:"MeetingName,omitempty" xml:"MeetingName,omitempty"`
	MeetingToken  *string                                     `json:"MeetingToken,omitempty" xml:"MeetingToken,omitempty"`
	MeetingUUID   *string                                     `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
	MemberUUID    *string                                     `json:"MemberUUID,omitempty" xml:"MemberUUID,omitempty"`
	SlsInfo       *JoinMeetingResponseBodyMeetingInfoSlsInfo  `json:"SlsInfo,omitempty" xml:"SlsInfo,omitempty" type:"Struct"`
	StartTime     *int64                                      `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s JoinMeetingResponseBodyMeetingInfo) String() string {
	return tea.Prettify(s)
}

func (s JoinMeetingResponseBodyMeetingInfo) GoString() string {
	return s.String()
}

func (s *JoinMeetingResponseBodyMeetingInfo) SetClientAppId(v string) *JoinMeetingResponseBodyMeetingInfo {
	s.ClientAppId = &v
	return s
}

func (s *JoinMeetingResponseBodyMeetingInfo) SetEndTime(v int64) *JoinMeetingResponseBodyMeetingInfo {
	s.EndTime = &v
	return s
}

func (s *JoinMeetingResponseBodyMeetingInfo) SetGrtnInfo(v *JoinMeetingResponseBodyMeetingInfoGrtnInfo) *JoinMeetingResponseBodyMeetingInfo {
	s.GrtnInfo = v
	return s
}

func (s *JoinMeetingResponseBodyMeetingInfo) SetMeetingAppId(v string) *JoinMeetingResponseBodyMeetingInfo {
	s.MeetingAppId = &v
	return s
}

func (s *JoinMeetingResponseBodyMeetingInfo) SetMeetingCode(v string) *JoinMeetingResponseBodyMeetingInfo {
	s.MeetingCode = &v
	return s
}

func (s *JoinMeetingResponseBodyMeetingInfo) SetMeetingDomain(v string) *JoinMeetingResponseBodyMeetingInfo {
	s.MeetingDomain = &v
	return s
}

func (s *JoinMeetingResponseBodyMeetingInfo) SetMeetingName(v string) *JoinMeetingResponseBodyMeetingInfo {
	s.MeetingName = &v
	return s
}

func (s *JoinMeetingResponseBodyMeetingInfo) SetMeetingToken(v string) *JoinMeetingResponseBodyMeetingInfo {
	s.MeetingToken = &v
	return s
}

func (s *JoinMeetingResponseBodyMeetingInfo) SetMeetingUUID(v string) *JoinMeetingResponseBodyMeetingInfo {
	s.MeetingUUID = &v
	return s
}

func (s *JoinMeetingResponseBodyMeetingInfo) SetMemberUUID(v string) *JoinMeetingResponseBodyMeetingInfo {
	s.MemberUUID = &v
	return s
}

func (s *JoinMeetingResponseBodyMeetingInfo) SetSlsInfo(v *JoinMeetingResponseBodyMeetingInfoSlsInfo) *JoinMeetingResponseBodyMeetingInfo {
	s.SlsInfo = v
	return s
}

func (s *JoinMeetingResponseBodyMeetingInfo) SetStartTime(v int64) *JoinMeetingResponseBodyMeetingInfo {
	s.StartTime = &v
	return s
}

type JoinMeetingResponseBodyMeetingInfoGrtnInfo struct {
	Agent     *string   `json:"Agent,omitempty" xml:"Agent,omitempty"`
	AppId     *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId *int64    `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Gslb      []*string `json:"Gslb,omitempty" xml:"Gslb,omitempty" type:"Repeated"`
	Nonce     *string   `json:"Nonce,omitempty" xml:"Nonce,omitempty"`
	Timestamp *int64    `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Token     *string   `json:"Token,omitempty" xml:"Token,omitempty"`
	UserId    *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s JoinMeetingResponseBodyMeetingInfoGrtnInfo) String() string {
	return tea.Prettify(s)
}

func (s JoinMeetingResponseBodyMeetingInfoGrtnInfo) GoString() string {
	return s.String()
}

func (s *JoinMeetingResponseBodyMeetingInfoGrtnInfo) SetAgent(v string) *JoinMeetingResponseBodyMeetingInfoGrtnInfo {
	s.Agent = &v
	return s
}

func (s *JoinMeetingResponseBodyMeetingInfoGrtnInfo) SetAppId(v string) *JoinMeetingResponseBodyMeetingInfoGrtnInfo {
	s.AppId = &v
	return s
}

func (s *JoinMeetingResponseBodyMeetingInfoGrtnInfo) SetChannelId(v int64) *JoinMeetingResponseBodyMeetingInfoGrtnInfo {
	s.ChannelId = &v
	return s
}

func (s *JoinMeetingResponseBodyMeetingInfoGrtnInfo) SetGslb(v []*string) *JoinMeetingResponseBodyMeetingInfoGrtnInfo {
	s.Gslb = v
	return s
}

func (s *JoinMeetingResponseBodyMeetingInfoGrtnInfo) SetNonce(v string) *JoinMeetingResponseBodyMeetingInfoGrtnInfo {
	s.Nonce = &v
	return s
}

func (s *JoinMeetingResponseBodyMeetingInfoGrtnInfo) SetTimestamp(v int64) *JoinMeetingResponseBodyMeetingInfoGrtnInfo {
	s.Timestamp = &v
	return s
}

func (s *JoinMeetingResponseBodyMeetingInfoGrtnInfo) SetToken(v string) *JoinMeetingResponseBodyMeetingInfoGrtnInfo {
	s.Token = &v
	return s
}

func (s *JoinMeetingResponseBodyMeetingInfoGrtnInfo) SetUserId(v string) *JoinMeetingResponseBodyMeetingInfoGrtnInfo {
	s.UserId = &v
	return s
}

type JoinMeetingResponseBodyMeetingInfoSlsInfo struct {
	LogServiceEndpoint *string `json:"LogServiceEndpoint,omitempty" xml:"LogServiceEndpoint,omitempty"`
	Logstore           *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	Project            *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s JoinMeetingResponseBodyMeetingInfoSlsInfo) String() string {
	return tea.Prettify(s)
}

func (s JoinMeetingResponseBodyMeetingInfoSlsInfo) GoString() string {
	return s.String()
}

func (s *JoinMeetingResponseBodyMeetingInfoSlsInfo) SetLogServiceEndpoint(v string) *JoinMeetingResponseBodyMeetingInfoSlsInfo {
	s.LogServiceEndpoint = &v
	return s
}

func (s *JoinMeetingResponseBodyMeetingInfoSlsInfo) SetLogstore(v string) *JoinMeetingResponseBodyMeetingInfoSlsInfo {
	s.Logstore = &v
	return s
}

func (s *JoinMeetingResponseBodyMeetingInfoSlsInfo) SetProject(v string) *JoinMeetingResponseBodyMeetingInfoSlsInfo {
	s.Project = &v
	return s
}

type JoinMeetingResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *JoinMeetingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s JoinMeetingResponse) String() string {
	return tea.Prettify(s)
}

func (s JoinMeetingResponse) GoString() string {
	return s.String()
}

func (s *JoinMeetingResponse) SetHeaders(v map[string]*string) *JoinMeetingResponse {
	s.Headers = v
	return s
}

func (s *JoinMeetingResponse) SetStatusCode(v int32) *JoinMeetingResponse {
	s.StatusCode = &v
	return s
}

func (s *JoinMeetingResponse) SetBody(v *JoinMeetingResponseBody) *JoinMeetingResponse {
	s.Body = v
	return s
}

type JoinMeetingInternationalRequest struct {
	MeetingCode *string `json:"MeetingCode,omitempty" xml:"MeetingCode,omitempty"`
	Password    *string `json:"Password,omitempty" xml:"Password,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s JoinMeetingInternationalRequest) String() string {
	return tea.Prettify(s)
}

func (s JoinMeetingInternationalRequest) GoString() string {
	return s.String()
}

func (s *JoinMeetingInternationalRequest) SetMeetingCode(v string) *JoinMeetingInternationalRequest {
	s.MeetingCode = &v
	return s
}

func (s *JoinMeetingInternationalRequest) SetPassword(v string) *JoinMeetingInternationalRequest {
	s.Password = &v
	return s
}

func (s *JoinMeetingInternationalRequest) SetUserId(v string) *JoinMeetingInternationalRequest {
	s.UserId = &v
	return s
}

type JoinMeetingInternationalResponseBody struct {
	ErrorCode   *int32                                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	MeetingInfo *JoinMeetingInternationalResponseBodyMeetingInfo `json:"MeetingInfo,omitempty" xml:"MeetingInfo,omitempty" type:"Struct"`
	Message     *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s JoinMeetingInternationalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s JoinMeetingInternationalResponseBody) GoString() string {
	return s.String()
}

func (s *JoinMeetingInternationalResponseBody) SetErrorCode(v int32) *JoinMeetingInternationalResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *JoinMeetingInternationalResponseBody) SetMeetingInfo(v *JoinMeetingInternationalResponseBodyMeetingInfo) *JoinMeetingInternationalResponseBody {
	s.MeetingInfo = v
	return s
}

func (s *JoinMeetingInternationalResponseBody) SetMessage(v string) *JoinMeetingInternationalResponseBody {
	s.Message = &v
	return s
}

func (s *JoinMeetingInternationalResponseBody) SetRequestId(v string) *JoinMeetingInternationalResponseBody {
	s.RequestId = &v
	return s
}

func (s *JoinMeetingInternationalResponseBody) SetSuccess(v bool) *JoinMeetingInternationalResponseBody {
	s.Success = &v
	return s
}

type JoinMeetingInternationalResponseBodyMeetingInfo struct {
	ClientAppId   *string                                                 `json:"ClientAppId,omitempty" xml:"ClientAppId,omitempty"`
	MeetingAppId  *string                                                 `json:"MeetingAppId,omitempty" xml:"MeetingAppId,omitempty"`
	MeetingCode   *string                                                 `json:"MeetingCode,omitempty" xml:"MeetingCode,omitempty"`
	MeetingDomain *string                                                 `json:"MeetingDomain,omitempty" xml:"MeetingDomain,omitempty"`
	MeetingToken  *string                                                 `json:"MeetingToken,omitempty" xml:"MeetingToken,omitempty"`
	MeetingUUID   *string                                                 `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
	MemberUUID    *string                                                 `json:"MemberUUID,omitempty" xml:"MemberUUID,omitempty"`
	SlsInfo       *JoinMeetingInternationalResponseBodyMeetingInfoSlsInfo `json:"SlsInfo,omitempty" xml:"SlsInfo,omitempty" type:"Struct"`
}

func (s JoinMeetingInternationalResponseBodyMeetingInfo) String() string {
	return tea.Prettify(s)
}

func (s JoinMeetingInternationalResponseBodyMeetingInfo) GoString() string {
	return s.String()
}

func (s *JoinMeetingInternationalResponseBodyMeetingInfo) SetClientAppId(v string) *JoinMeetingInternationalResponseBodyMeetingInfo {
	s.ClientAppId = &v
	return s
}

func (s *JoinMeetingInternationalResponseBodyMeetingInfo) SetMeetingAppId(v string) *JoinMeetingInternationalResponseBodyMeetingInfo {
	s.MeetingAppId = &v
	return s
}

func (s *JoinMeetingInternationalResponseBodyMeetingInfo) SetMeetingCode(v string) *JoinMeetingInternationalResponseBodyMeetingInfo {
	s.MeetingCode = &v
	return s
}

func (s *JoinMeetingInternationalResponseBodyMeetingInfo) SetMeetingDomain(v string) *JoinMeetingInternationalResponseBodyMeetingInfo {
	s.MeetingDomain = &v
	return s
}

func (s *JoinMeetingInternationalResponseBodyMeetingInfo) SetMeetingToken(v string) *JoinMeetingInternationalResponseBodyMeetingInfo {
	s.MeetingToken = &v
	return s
}

func (s *JoinMeetingInternationalResponseBodyMeetingInfo) SetMeetingUUID(v string) *JoinMeetingInternationalResponseBodyMeetingInfo {
	s.MeetingUUID = &v
	return s
}

func (s *JoinMeetingInternationalResponseBodyMeetingInfo) SetMemberUUID(v string) *JoinMeetingInternationalResponseBodyMeetingInfo {
	s.MemberUUID = &v
	return s
}

func (s *JoinMeetingInternationalResponseBodyMeetingInfo) SetSlsInfo(v *JoinMeetingInternationalResponseBodyMeetingInfoSlsInfo) *JoinMeetingInternationalResponseBodyMeetingInfo {
	s.SlsInfo = v
	return s
}

type JoinMeetingInternationalResponseBodyMeetingInfoSlsInfo struct {
	LogServiceEndpoint *string `json:"LogServiceEndpoint,omitempty" xml:"LogServiceEndpoint,omitempty"`
	Logstore           *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	Project            *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s JoinMeetingInternationalResponseBodyMeetingInfoSlsInfo) String() string {
	return tea.Prettify(s)
}

func (s JoinMeetingInternationalResponseBodyMeetingInfoSlsInfo) GoString() string {
	return s.String()
}

func (s *JoinMeetingInternationalResponseBodyMeetingInfoSlsInfo) SetLogServiceEndpoint(v string) *JoinMeetingInternationalResponseBodyMeetingInfoSlsInfo {
	s.LogServiceEndpoint = &v
	return s
}

func (s *JoinMeetingInternationalResponseBodyMeetingInfoSlsInfo) SetLogstore(v string) *JoinMeetingInternationalResponseBodyMeetingInfoSlsInfo {
	s.Logstore = &v
	return s
}

func (s *JoinMeetingInternationalResponseBodyMeetingInfoSlsInfo) SetProject(v string) *JoinMeetingInternationalResponseBodyMeetingInfoSlsInfo {
	s.Project = &v
	return s
}

type JoinMeetingInternationalResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *JoinMeetingInternationalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s JoinMeetingInternationalResponse) String() string {
	return tea.Prettify(s)
}

func (s JoinMeetingInternationalResponse) GoString() string {
	return s.String()
}

func (s *JoinMeetingInternationalResponse) SetHeaders(v map[string]*string) *JoinMeetingInternationalResponse {
	s.Headers = v
	return s
}

func (s *JoinMeetingInternationalResponse) SetStatusCode(v int32) *JoinMeetingInternationalResponse {
	s.StatusCode = &v
	return s
}

func (s *JoinMeetingInternationalResponse) SetBody(v *JoinMeetingInternationalResponseBody) *JoinMeetingInternationalResponse {
	s.Body = v
	return s
}

type ListConferenceDevicesRequest struct {
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s ListConferenceDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConferenceDevicesRequest) GoString() string {
	return s.String()
}

func (s *ListConferenceDevicesRequest) SetPageNumber(v int32) *ListConferenceDevicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListConferenceDevicesRequest) SetPageSize(v int32) *ListConferenceDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListConferenceDevicesRequest) SetSerialNumber(v string) *ListConferenceDevicesRequest {
	s.SerialNumber = &v
	return s
}

type ListConferenceDevicesResponseBody struct {
	ConferencesDatas *ListConferenceDevicesResponseBodyConferencesDatas `json:"ConferencesDatas,omitempty" xml:"ConferencesDatas,omitempty" type:"Struct"`
	ErrorCode        *int32                                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message          *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId        *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListConferenceDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConferenceDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListConferenceDevicesResponseBody) SetConferencesDatas(v *ListConferenceDevicesResponseBodyConferencesDatas) *ListConferenceDevicesResponseBody {
	s.ConferencesDatas = v
	return s
}

func (s *ListConferenceDevicesResponseBody) SetErrorCode(v int32) *ListConferenceDevicesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListConferenceDevicesResponseBody) SetMessage(v string) *ListConferenceDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListConferenceDevicesResponseBody) SetRequestId(v string) *ListConferenceDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConferenceDevicesResponseBody) SetSuccess(v bool) *ListConferenceDevicesResponseBody {
	s.Success = &v
	return s
}

type ListConferenceDevicesResponseBodyConferencesDatas struct {
	Conferences []*ListConferenceDevicesResponseBodyConferencesDatasConferences `json:"Conferences,omitempty" xml:"Conferences,omitempty" type:"Repeated"`
	PageNumber  *int32                                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total       *int32                                                          `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListConferenceDevicesResponseBodyConferencesDatas) String() string {
	return tea.Prettify(s)
}

func (s ListConferenceDevicesResponseBodyConferencesDatas) GoString() string {
	return s.String()
}

func (s *ListConferenceDevicesResponseBodyConferencesDatas) SetConferences(v []*ListConferenceDevicesResponseBodyConferencesDatasConferences) *ListConferenceDevicesResponseBodyConferencesDatas {
	s.Conferences = v
	return s
}

func (s *ListConferenceDevicesResponseBodyConferencesDatas) SetPageNumber(v int32) *ListConferenceDevicesResponseBodyConferencesDatas {
	s.PageNumber = &v
	return s
}

func (s *ListConferenceDevicesResponseBodyConferencesDatas) SetPageSize(v int32) *ListConferenceDevicesResponseBodyConferencesDatas {
	s.PageSize = &v
	return s
}

func (s *ListConferenceDevicesResponseBodyConferencesDatas) SetTotal(v int32) *ListConferenceDevicesResponseBodyConferencesDatas {
	s.Total = &v
	return s
}

type ListConferenceDevicesResponseBodyConferencesDatasConferences struct {
	ActivationCode    *string `json:"ActivationCode,omitempty" xml:"ActivationCode,omitempty"`
	CastScreenCode    *string `json:"CastScreenCode,omitempty" xml:"CastScreenCode,omitempty"`
	ConferenceCode    *string `json:"ConferenceCode,omitempty" xml:"ConferenceCode,omitempty"`
	ConferenceName    *string `json:"ConferenceName,omitempty" xml:"ConferenceName,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeviceModel       *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	Manufacturer      *string `json:"Manufacturer,omitempty" xml:"Manufacturer,omitempty"`
	PictureUrl        *string `json:"PictureUrl,omitempty" xml:"PictureUrl,omitempty"`
	SN                *string `json:"SN,omitempty" xml:"SN,omitempty"`
	StartUpPictureUrl *string `json:"StartUpPictureUrl,omitempty" xml:"StartUpPictureUrl,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListConferenceDevicesResponseBodyConferencesDatasConferences) String() string {
	return tea.Prettify(s)
}

func (s ListConferenceDevicesResponseBodyConferencesDatasConferences) GoString() string {
	return s.String()
}

func (s *ListConferenceDevicesResponseBodyConferencesDatasConferences) SetActivationCode(v string) *ListConferenceDevicesResponseBodyConferencesDatasConferences {
	s.ActivationCode = &v
	return s
}

func (s *ListConferenceDevicesResponseBodyConferencesDatasConferences) SetCastScreenCode(v string) *ListConferenceDevicesResponseBodyConferencesDatasConferences {
	s.CastScreenCode = &v
	return s
}

func (s *ListConferenceDevicesResponseBodyConferencesDatasConferences) SetConferenceCode(v string) *ListConferenceDevicesResponseBodyConferencesDatasConferences {
	s.ConferenceCode = &v
	return s
}

func (s *ListConferenceDevicesResponseBodyConferencesDatasConferences) SetConferenceName(v string) *ListConferenceDevicesResponseBodyConferencesDatasConferences {
	s.ConferenceName = &v
	return s
}

func (s *ListConferenceDevicesResponseBodyConferencesDatasConferences) SetCreateTime(v string) *ListConferenceDevicesResponseBodyConferencesDatasConferences {
	s.CreateTime = &v
	return s
}

func (s *ListConferenceDevicesResponseBodyConferencesDatasConferences) SetDeviceModel(v string) *ListConferenceDevicesResponseBodyConferencesDatasConferences {
	s.DeviceModel = &v
	return s
}

func (s *ListConferenceDevicesResponseBodyConferencesDatasConferences) SetManufacturer(v string) *ListConferenceDevicesResponseBodyConferencesDatasConferences {
	s.Manufacturer = &v
	return s
}

func (s *ListConferenceDevicesResponseBodyConferencesDatasConferences) SetPictureUrl(v string) *ListConferenceDevicesResponseBodyConferencesDatasConferences {
	s.PictureUrl = &v
	return s
}

func (s *ListConferenceDevicesResponseBodyConferencesDatasConferences) SetSN(v string) *ListConferenceDevicesResponseBodyConferencesDatasConferences {
	s.SN = &v
	return s
}

func (s *ListConferenceDevicesResponseBodyConferencesDatasConferences) SetStartUpPictureUrl(v string) *ListConferenceDevicesResponseBodyConferencesDatasConferences {
	s.StartUpPictureUrl = &v
	return s
}

func (s *ListConferenceDevicesResponseBodyConferencesDatasConferences) SetStatus(v string) *ListConferenceDevicesResponseBodyConferencesDatasConferences {
	s.Status = &v
	return s
}

type ListConferenceDevicesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListConferenceDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConferenceDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConferenceDevicesResponse) GoString() string {
	return s.String()
}

func (s *ListConferenceDevicesResponse) SetHeaders(v map[string]*string) *ListConferenceDevicesResponse {
	s.Headers = v
	return s
}

func (s *ListConferenceDevicesResponse) SetStatusCode(v int32) *ListConferenceDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConferenceDevicesResponse) SetBody(v *ListConferenceDevicesResponseBody) *ListConferenceDevicesResponse {
	s.Body = v
	return s
}

type ListDeviceIpRequest struct {
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SN      *string `json:"SN,omitempty" xml:"SN,omitempty"`
}

func (s ListDeviceIpRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceIpRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceIpRequest) SetGroupId(v string) *ListDeviceIpRequest {
	s.GroupId = &v
	return s
}

func (s *ListDeviceIpRequest) SetSN(v string) *ListDeviceIpRequest {
	s.SN = &v
	return s
}

type ListDeviceIpResponseBody struct {
	Devices   []*ListDeviceIpResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	ErrorCode *int32                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDeviceIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceIpResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceIpResponseBody) SetDevices(v []*ListDeviceIpResponseBodyDevices) *ListDeviceIpResponseBody {
	s.Devices = v
	return s
}

func (s *ListDeviceIpResponseBody) SetErrorCode(v int32) *ListDeviceIpResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDeviceIpResponseBody) SetMessage(v string) *ListDeviceIpResponseBody {
	s.Message = &v
	return s
}

func (s *ListDeviceIpResponseBody) SetRequestId(v string) *ListDeviceIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceIpResponseBody) SetSuccess(v bool) *ListDeviceIpResponseBody {
	s.Success = &v
	return s
}

type ListDeviceIpResponseBodyDevices struct {
	Ip         *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Mac        *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	Port       *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ScreenCode *string `json:"ScreenCode,omitempty" xml:"ScreenCode,omitempty"`
	SsId       *string `json:"SsId,omitempty" xml:"SsId,omitempty"`
}

func (s ListDeviceIpResponseBodyDevices) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceIpResponseBodyDevices) GoString() string {
	return s.String()
}

func (s *ListDeviceIpResponseBodyDevices) SetIp(v string) *ListDeviceIpResponseBodyDevices {
	s.Ip = &v
	return s
}

func (s *ListDeviceIpResponseBodyDevices) SetMac(v string) *ListDeviceIpResponseBodyDevices {
	s.Mac = &v
	return s
}

func (s *ListDeviceIpResponseBodyDevices) SetPort(v string) *ListDeviceIpResponseBodyDevices {
	s.Port = &v
	return s
}

func (s *ListDeviceIpResponseBodyDevices) SetScreenCode(v string) *ListDeviceIpResponseBodyDevices {
	s.ScreenCode = &v
	return s
}

func (s *ListDeviceIpResponseBodyDevices) SetSsId(v string) *ListDeviceIpResponseBodyDevices {
	s.SsId = &v
	return s
}

type ListDeviceIpResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDeviceIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeviceIpResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceIpResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceIpResponse) SetHeaders(v map[string]*string) *ListDeviceIpResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceIpResponse) SetStatusCode(v int32) *ListDeviceIpResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceIpResponse) SetBody(v *ListDeviceIpResponseBody) *ListDeviceIpResponse {
	s.Body = v
	return s
}

type ListDevicesRequest struct {
	CastScreenCode *string `json:"CastScreenCode,omitempty" xml:"CastScreenCode,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SN             *string `json:"SN,omitempty" xml:"SN,omitempty"`
}

func (s ListDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesRequest) GoString() string {
	return s.String()
}

func (s *ListDevicesRequest) SetCastScreenCode(v string) *ListDevicesRequest {
	s.CastScreenCode = &v
	return s
}

func (s *ListDevicesRequest) SetPageNumber(v int32) *ListDevicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDevicesRequest) SetPageSize(v int32) *ListDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDevicesRequest) SetSN(v string) *ListDevicesRequest {
	s.SN = &v
	return s
}

type ListDevicesResponseBody struct {
	Data      *ListDevicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *int32                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBody) SetData(v *ListDevicesResponseBodyData) *ListDevicesResponseBody {
	s.Data = v
	return s
}

func (s *ListDevicesResponseBody) SetErrorCode(v int32) *ListDevicesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDevicesResponseBody) SetMessage(v string) *ListDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListDevicesResponseBody) SetRequestId(v string) *ListDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDevicesResponseBody) SetSuccess(v bool) *ListDevicesResponseBody {
	s.Success = &v
	return s
}

type ListDevicesResponseBodyData struct {
	Devices    []*ListDevicesResponseBodyDataDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total      *int32                                `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDevicesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBodyData) SetDevices(v []*ListDevicesResponseBodyDataDevices) *ListDevicesResponseBodyData {
	s.Devices = v
	return s
}

func (s *ListDevicesResponseBodyData) SetPageNumber(v int32) *ListDevicesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetPageSize(v int32) *ListDevicesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetTotal(v int32) *ListDevicesResponseBodyData {
	s.Total = &v
	return s
}

type ListDevicesResponseBodyDataDevices struct {
	ActivationCode    *string `json:"ActivationCode,omitempty" xml:"ActivationCode,omitempty"`
	CastScreenCode    *string `json:"CastScreenCode,omitempty" xml:"CastScreenCode,omitempty"`
	ConferenceCode    *string `json:"ConferenceCode,omitempty" xml:"ConferenceCode,omitempty"`
	ConferenceName    *string `json:"ConferenceName,omitempty" xml:"ConferenceName,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PictureUrl        *string `json:"PictureUrl,omitempty" xml:"PictureUrl,omitempty"`
	SN                *string `json:"SN,omitempty" xml:"SN,omitempty"`
	StartUpPictureUrl *string `json:"StartUpPictureUrl,omitempty" xml:"StartUpPictureUrl,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDevicesResponseBodyDataDevices) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponseBodyDataDevices) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBodyDataDevices) SetActivationCode(v string) *ListDevicesResponseBodyDataDevices {
	s.ActivationCode = &v
	return s
}

func (s *ListDevicesResponseBodyDataDevices) SetCastScreenCode(v string) *ListDevicesResponseBodyDataDevices {
	s.CastScreenCode = &v
	return s
}

func (s *ListDevicesResponseBodyDataDevices) SetConferenceCode(v string) *ListDevicesResponseBodyDataDevices {
	s.ConferenceCode = &v
	return s
}

func (s *ListDevicesResponseBodyDataDevices) SetConferenceName(v string) *ListDevicesResponseBodyDataDevices {
	s.ConferenceName = &v
	return s
}

func (s *ListDevicesResponseBodyDataDevices) SetCreateTime(v string) *ListDevicesResponseBodyDataDevices {
	s.CreateTime = &v
	return s
}

func (s *ListDevicesResponseBodyDataDevices) SetPictureUrl(v string) *ListDevicesResponseBodyDataDevices {
	s.PictureUrl = &v
	return s
}

func (s *ListDevicesResponseBodyDataDevices) SetSN(v string) *ListDevicesResponseBodyDataDevices {
	s.SN = &v
	return s
}

func (s *ListDevicesResponseBodyDataDevices) SetStartUpPictureUrl(v string) *ListDevicesResponseBodyDataDevices {
	s.StartUpPictureUrl = &v
	return s
}

func (s *ListDevicesResponseBodyDataDevices) SetStatus(v string) *ListDevicesResponseBodyDataDevices {
	s.Status = &v
	return s
}

type ListDevicesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponse) GoString() string {
	return s.String()
}

func (s *ListDevicesResponse) SetHeaders(v map[string]*string) *ListDevicesResponse {
	s.Headers = v
	return s
}

func (s *ListDevicesResponse) SetStatusCode(v int32) *ListDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDevicesResponse) SetBody(v *ListDevicesResponseBody) *ListDevicesResponse {
	s.Body = v
	return s
}

type ListEvaluationsResponseBody struct {
	ErrorCode      *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	UserEvaluation *string `json:"UserEvaluation,omitempty" xml:"UserEvaluation,omitempty"`
}

func (s ListEvaluationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEvaluationsResponseBody) SetErrorCode(v int32) *ListEvaluationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListEvaluationsResponseBody) SetMessage(v string) *ListEvaluationsResponseBody {
	s.Message = &v
	return s
}

func (s *ListEvaluationsResponseBody) SetRequestId(v string) *ListEvaluationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEvaluationsResponseBody) SetSuccess(v bool) *ListEvaluationsResponseBody {
	s.Success = &v
	return s
}

func (s *ListEvaluationsResponseBody) SetUserEvaluation(v string) *ListEvaluationsResponseBody {
	s.UserEvaluation = &v
	return s
}

type ListEvaluationsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListEvaluationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEvaluationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationsResponse) GoString() string {
	return s.String()
}

func (s *ListEvaluationsResponse) SetHeaders(v map[string]*string) *ListEvaluationsResponse {
	s.Headers = v
	return s
}

func (s *ListEvaluationsResponse) SetStatusCode(v int32) *ListEvaluationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEvaluationsResponse) SetBody(v *ListEvaluationsResponseBody) *ListEvaluationsResponse {
	s.Body = v
	return s
}

type ListIsvStatisticsRequest struct {
	EndTime   *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListIsvStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIsvStatisticsRequest) GoString() string {
	return s.String()
}

func (s *ListIsvStatisticsRequest) SetEndTime(v int64) *ListIsvStatisticsRequest {
	s.EndTime = &v
	return s
}

func (s *ListIsvStatisticsRequest) SetStartTime(v int64) *ListIsvStatisticsRequest {
	s.StartTime = &v
	return s
}

type ListIsvStatisticsResponseBody struct {
	Data      *ListIsvStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *int32                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListIsvStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIsvStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIsvStatisticsResponseBody) SetData(v *ListIsvStatisticsResponseBodyData) *ListIsvStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *ListIsvStatisticsResponseBody) SetErrorCode(v int32) *ListIsvStatisticsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListIsvStatisticsResponseBody) SetMessage(v string) *ListIsvStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *ListIsvStatisticsResponseBody) SetRequestId(v string) *ListIsvStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIsvStatisticsResponseBody) SetSuccess(v bool) *ListIsvStatisticsResponseBody {
	s.Success = &v
	return s
}

type ListIsvStatisticsResponseBodyData struct {
	Statistics []*ListIsvStatisticsResponseBodyDataStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Repeated"`
	Total      *ListIsvStatisticsResponseBodyDataTotal        `json:"Total,omitempty" xml:"Total,omitempty" type:"Struct"`
}

func (s ListIsvStatisticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListIsvStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIsvStatisticsResponseBodyData) SetStatistics(v []*ListIsvStatisticsResponseBodyDataStatistics) *ListIsvStatisticsResponseBodyData {
	s.Statistics = v
	return s
}

func (s *ListIsvStatisticsResponseBodyData) SetTotal(v *ListIsvStatisticsResponseBodyDataTotal) *ListIsvStatisticsResponseBodyData {
	s.Total = v
	return s
}

type ListIsvStatisticsResponseBodyDataStatistics struct {
	Day           *string `json:"Day,omitempty" xml:"Day,omitempty"`
	MeetingLength *string `json:"MeetingLength,omitempty" xml:"MeetingLength,omitempty"`
	MeetingNumber *string `json:"MeetingNumber,omitempty" xml:"MeetingNumber,omitempty"`
	MemberNumber  *string `json:"MemberNumber,omitempty" xml:"MemberNumber,omitempty"`
}

func (s ListIsvStatisticsResponseBodyDataStatistics) String() string {
	return tea.Prettify(s)
}

func (s ListIsvStatisticsResponseBodyDataStatistics) GoString() string {
	return s.String()
}

func (s *ListIsvStatisticsResponseBodyDataStatistics) SetDay(v string) *ListIsvStatisticsResponseBodyDataStatistics {
	s.Day = &v
	return s
}

func (s *ListIsvStatisticsResponseBodyDataStatistics) SetMeetingLength(v string) *ListIsvStatisticsResponseBodyDataStatistics {
	s.MeetingLength = &v
	return s
}

func (s *ListIsvStatisticsResponseBodyDataStatistics) SetMeetingNumber(v string) *ListIsvStatisticsResponseBodyDataStatistics {
	s.MeetingNumber = &v
	return s
}

func (s *ListIsvStatisticsResponseBodyDataStatistics) SetMemberNumber(v string) *ListIsvStatisticsResponseBodyDataStatistics {
	s.MemberNumber = &v
	return s
}

type ListIsvStatisticsResponseBodyDataTotal struct {
	MeetingLength *int32 `json:"MeetingLength,omitempty" xml:"MeetingLength,omitempty"`
	MeetingNumber *int32 `json:"MeetingNumber,omitempty" xml:"MeetingNumber,omitempty"`
	MemberNumber  *int32 `json:"MemberNumber,omitempty" xml:"MemberNumber,omitempty"`
}

func (s ListIsvStatisticsResponseBodyDataTotal) String() string {
	return tea.Prettify(s)
}

func (s ListIsvStatisticsResponseBodyDataTotal) GoString() string {
	return s.String()
}

func (s *ListIsvStatisticsResponseBodyDataTotal) SetMeetingLength(v int32) *ListIsvStatisticsResponseBodyDataTotal {
	s.MeetingLength = &v
	return s
}

func (s *ListIsvStatisticsResponseBodyDataTotal) SetMeetingNumber(v int32) *ListIsvStatisticsResponseBodyDataTotal {
	s.MeetingNumber = &v
	return s
}

func (s *ListIsvStatisticsResponseBodyDataTotal) SetMemberNumber(v int32) *ListIsvStatisticsResponseBodyDataTotal {
	s.MemberNumber = &v
	return s
}

type ListIsvStatisticsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListIsvStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIsvStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIsvStatisticsResponse) GoString() string {
	return s.String()
}

func (s *ListIsvStatisticsResponse) SetHeaders(v map[string]*string) *ListIsvStatisticsResponse {
	s.Headers = v
	return s
}

func (s *ListIsvStatisticsResponse) SetStatusCode(v int32) *ListIsvStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIsvStatisticsResponse) SetBody(v *ListIsvStatisticsResponseBody) *ListIsvStatisticsResponse {
	s.Body = v
	return s
}

type ListMembersRequest struct {
	MeetingUUID *string `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
}

func (s ListMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMembersRequest) GoString() string {
	return s.String()
}

func (s *ListMembersRequest) SetMeetingUUID(v string) *ListMembersRequest {
	s.MeetingUUID = &v
	return s
}

type ListMembersResponseBody struct {
	ErrorCode   *int32                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	MeetingInfo *ListMembersResponseBodyMeetingInfo `json:"MeetingInfo,omitempty" xml:"MeetingInfo,omitempty" type:"Struct"`
	Message     *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListMembersResponseBody) SetErrorCode(v int32) *ListMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListMembersResponseBody) SetMeetingInfo(v *ListMembersResponseBodyMeetingInfo) *ListMembersResponseBody {
	s.MeetingInfo = v
	return s
}

func (s *ListMembersResponseBody) SetMessage(v string) *ListMembersResponseBody {
	s.Message = &v
	return s
}

func (s *ListMembersResponseBody) SetRequestId(v string) *ListMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMembersResponseBody) SetSuccess(v bool) *ListMembersResponseBody {
	s.Success = &v
	return s
}

type ListMembersResponseBodyMeetingInfo struct {
	CreateTime  *int64                                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	MeetingCode *string                                          `json:"MeetingCode,omitempty" xml:"MeetingCode,omitempty"`
	MeetingName *string                                          `json:"MeetingName,omitempty" xml:"MeetingName,omitempty"`
	MeetingUUID *string                                          `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
	MemberInfos []*ListMembersResponseBodyMeetingInfoMemberInfos `json:"MemberInfos,omitempty" xml:"MemberInfos,omitempty" type:"Repeated"`
	Memo        *string                                          `json:"Memo,omitempty" xml:"Memo,omitempty"`
	UserId      *string                                          `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName    *string                                          `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListMembersResponseBodyMeetingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListMembersResponseBodyMeetingInfo) GoString() string {
	return s.String()
}

func (s *ListMembersResponseBodyMeetingInfo) SetCreateTime(v int64) *ListMembersResponseBodyMeetingInfo {
	s.CreateTime = &v
	return s
}

func (s *ListMembersResponseBodyMeetingInfo) SetMeetingCode(v string) *ListMembersResponseBodyMeetingInfo {
	s.MeetingCode = &v
	return s
}

func (s *ListMembersResponseBodyMeetingInfo) SetMeetingName(v string) *ListMembersResponseBodyMeetingInfo {
	s.MeetingName = &v
	return s
}

func (s *ListMembersResponseBodyMeetingInfo) SetMeetingUUID(v string) *ListMembersResponseBodyMeetingInfo {
	s.MeetingUUID = &v
	return s
}

func (s *ListMembersResponseBodyMeetingInfo) SetMemberInfos(v []*ListMembersResponseBodyMeetingInfoMemberInfos) *ListMembersResponseBodyMeetingInfo {
	s.MemberInfos = v
	return s
}

func (s *ListMembersResponseBodyMeetingInfo) SetMemo(v string) *ListMembersResponseBodyMeetingInfo {
	s.Memo = &v
	return s
}

func (s *ListMembersResponseBodyMeetingInfo) SetUserId(v string) *ListMembersResponseBodyMeetingInfo {
	s.UserId = &v
	return s
}

func (s *ListMembersResponseBodyMeetingInfo) SetUserName(v string) *ListMembersResponseBodyMeetingInfo {
	s.UserName = &v
	return s
}

type ListMembersResponseBodyMeetingInfoMemberInfos struct {
	BeginTime  *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MemberUUID *string `json:"MemberUUID,omitempty" xml:"MemberUUID,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName   *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListMembersResponseBodyMeetingInfoMemberInfos) String() string {
	return tea.Prettify(s)
}

func (s ListMembersResponseBodyMeetingInfoMemberInfos) GoString() string {
	return s.String()
}

func (s *ListMembersResponseBodyMeetingInfoMemberInfos) SetBeginTime(v int64) *ListMembersResponseBodyMeetingInfoMemberInfos {
	s.BeginTime = &v
	return s
}

func (s *ListMembersResponseBodyMeetingInfoMemberInfos) SetEndTime(v int64) *ListMembersResponseBodyMeetingInfoMemberInfos {
	s.EndTime = &v
	return s
}

func (s *ListMembersResponseBodyMeetingInfoMemberInfos) SetMemberUUID(v string) *ListMembersResponseBodyMeetingInfoMemberInfos {
	s.MemberUUID = &v
	return s
}

func (s *ListMembersResponseBodyMeetingInfoMemberInfos) SetStatus(v string) *ListMembersResponseBodyMeetingInfoMemberInfos {
	s.Status = &v
	return s
}

func (s *ListMembersResponseBodyMeetingInfoMemberInfos) SetUserId(v string) *ListMembersResponseBodyMeetingInfoMemberInfos {
	s.UserId = &v
	return s
}

func (s *ListMembersResponseBodyMeetingInfoMemberInfos) SetUserName(v string) *ListMembersResponseBodyMeetingInfoMemberInfos {
	s.UserName = &v
	return s
}

type ListMembersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMembersResponse) GoString() string {
	return s.String()
}

func (s *ListMembersResponse) SetHeaders(v map[string]*string) *ListMembersResponse {
	s.Headers = v
	return s
}

func (s *ListMembersResponse) SetStatusCode(v int32) *ListMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMembersResponse) SetBody(v *ListMembersResponseBody) *ListMembersResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetPageNumber(v int32) *ListUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v int32) *ListUsersRequest {
	s.PageSize = &v
	return s
}

type ListUsersResponseBody struct {
	Data      *ListUsersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *int32                     `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetData(v *ListUsersResponseBodyData) *ListUsersResponseBody {
	s.Data = v
	return s
}

func (s *ListUsersResponseBody) SetErrorCode(v int32) *ListUsersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListUsersResponseBody) SetMessage(v string) *ListUsersResponseBody {
	s.Message = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetSuccess(v bool) *ListUsersResponseBody {
	s.Success = &v
	return s
}

type ListUsersResponseBodyData struct {
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UserInfos  []*ListUsersResponseBodyDataUserInfos `json:"UserInfos,omitempty" xml:"UserInfos,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyData) SetPageNumber(v int32) *ListUsersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListUsersResponseBodyData) SetPageSize(v int32) *ListUsersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListUsersResponseBodyData) SetTotalCount(v int32) *ListUsersResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListUsersResponseBodyData) SetUserInfos(v []*ListUsersResponseBodyDataUserInfos) *ListUsersResponseBodyData {
	s.UserInfos = v
	return s
}

type ListUsersResponseBodyDataUserInfos struct {
	CreateTime    *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DepartId      *string `json:"DepartId,omitempty" xml:"DepartId,omitempty"`
	DepartName    *string `json:"DepartName,omitempty" xml:"DepartName,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	JobName       *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	UserAvatarUrl *string `json:"UserAvatarUrl,omitempty" xml:"UserAvatarUrl,omitempty"`
	UserEmail     *string `json:"UserEmail,omitempty" xml:"UserEmail,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserMobile    *string `json:"UserMobile,omitempty" xml:"UserMobile,omitempty"`
	UserName      *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserTel       *string `json:"UserTel,omitempty" xml:"UserTel,omitempty"`
}

func (s ListUsersResponseBodyDataUserInfos) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyDataUserInfos) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyDataUserInfos) SetCreateTime(v int64) *ListUsersResponseBodyDataUserInfos {
	s.CreateTime = &v
	return s
}

func (s *ListUsersResponseBodyDataUserInfos) SetDepartId(v string) *ListUsersResponseBodyDataUserInfos {
	s.DepartId = &v
	return s
}

func (s *ListUsersResponseBodyDataUserInfos) SetDepartName(v string) *ListUsersResponseBodyDataUserInfos {
	s.DepartName = &v
	return s
}

func (s *ListUsersResponseBodyDataUserInfos) SetGroupId(v string) *ListUsersResponseBodyDataUserInfos {
	s.GroupId = &v
	return s
}

func (s *ListUsersResponseBodyDataUserInfos) SetGroupName(v string) *ListUsersResponseBodyDataUserInfos {
	s.GroupName = &v
	return s
}

func (s *ListUsersResponseBodyDataUserInfos) SetJobName(v string) *ListUsersResponseBodyDataUserInfos {
	s.JobName = &v
	return s
}

func (s *ListUsersResponseBodyDataUserInfos) SetUserAvatarUrl(v string) *ListUsersResponseBodyDataUserInfos {
	s.UserAvatarUrl = &v
	return s
}

func (s *ListUsersResponseBodyDataUserInfos) SetUserEmail(v string) *ListUsersResponseBodyDataUserInfos {
	s.UserEmail = &v
	return s
}

func (s *ListUsersResponseBodyDataUserInfos) SetUserId(v string) *ListUsersResponseBodyDataUserInfos {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseBodyDataUserInfos) SetUserMobile(v string) *ListUsersResponseBodyDataUserInfos {
	s.UserMobile = &v
	return s
}

func (s *ListUsersResponseBodyDataUserInfos) SetUserName(v string) *ListUsersResponseBodyDataUserInfos {
	s.UserName = &v
	return s
}

func (s *ListUsersResponseBodyDataUserInfos) SetUserTel(v string) *ListUsersResponseBodyDataUserInfos {
	s.UserTel = &v
	return s
}

type ListUsersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponse) GoString() string {
	return s.String()
}

func (s *ListUsersResponse) SetHeaders(v map[string]*string) *ListUsersResponse {
	s.Headers = v
	return s
}

func (s *ListUsersResponse) SetStatusCode(v int32) *ListUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type ModifyDeviceBackgroundRequest struct {
	Picture      *string `json:"Picture,omitempty" xml:"Picture,omitempty"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s ModifyDeviceBackgroundRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeviceBackgroundRequest) GoString() string {
	return s.String()
}

func (s *ModifyDeviceBackgroundRequest) SetPicture(v string) *ModifyDeviceBackgroundRequest {
	s.Picture = &v
	return s
}

func (s *ModifyDeviceBackgroundRequest) SetSerialNumber(v string) *ModifyDeviceBackgroundRequest {
	s.SerialNumber = &v
	return s
}

type ModifyDeviceBackgroundResponseBody struct {
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDeviceBackgroundResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeviceBackgroundResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDeviceBackgroundResponseBody) SetErrorCode(v int32) *ModifyDeviceBackgroundResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ModifyDeviceBackgroundResponseBody) SetMessage(v string) *ModifyDeviceBackgroundResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyDeviceBackgroundResponseBody) SetRequestId(v string) *ModifyDeviceBackgroundResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDeviceBackgroundResponseBody) SetSuccess(v bool) *ModifyDeviceBackgroundResponseBody {
	s.Success = &v
	return s
}

type ModifyDeviceBackgroundResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDeviceBackgroundResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDeviceBackgroundResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeviceBackgroundResponse) GoString() string {
	return s.String()
}

func (s *ModifyDeviceBackgroundResponse) SetHeaders(v map[string]*string) *ModifyDeviceBackgroundResponse {
	s.Headers = v
	return s
}

func (s *ModifyDeviceBackgroundResponse) SetStatusCode(v int32) *ModifyDeviceBackgroundResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDeviceBackgroundResponse) SetBody(v *ModifyDeviceBackgroundResponseBody) *ModifyDeviceBackgroundResponse {
	s.Body = v
	return s
}

type ModifyMeetingPasswordRequest struct {
	MeetingUUID      *string `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
	OpenPasswordFlag *bool   `json:"OpenPasswordFlag,omitempty" xml:"OpenPasswordFlag,omitempty"`
	Password         *string `json:"Password,omitempty" xml:"Password,omitempty"`
	UserId           *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ModifyMeetingPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyMeetingPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyMeetingPasswordRequest) SetMeetingUUID(v string) *ModifyMeetingPasswordRequest {
	s.MeetingUUID = &v
	return s
}

func (s *ModifyMeetingPasswordRequest) SetOpenPasswordFlag(v bool) *ModifyMeetingPasswordRequest {
	s.OpenPasswordFlag = &v
	return s
}

func (s *ModifyMeetingPasswordRequest) SetPassword(v string) *ModifyMeetingPasswordRequest {
	s.Password = &v
	return s
}

func (s *ModifyMeetingPasswordRequest) SetUserId(v string) *ModifyMeetingPasswordRequest {
	s.UserId = &v
	return s
}

type ModifyMeetingPasswordResponseBody struct {
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyMeetingPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyMeetingPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMeetingPasswordResponseBody) SetErrorCode(v int32) *ModifyMeetingPasswordResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ModifyMeetingPasswordResponseBody) SetMessage(v string) *ModifyMeetingPasswordResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyMeetingPasswordResponseBody) SetRequestId(v string) *ModifyMeetingPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMeetingPasswordResponseBody) SetSuccess(v bool) *ModifyMeetingPasswordResponseBody {
	s.Success = &v
	return s
}

type ModifyMeetingPasswordResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyMeetingPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyMeetingPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyMeetingPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifyMeetingPasswordResponse) SetHeaders(v map[string]*string) *ModifyMeetingPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifyMeetingPasswordResponse) SetStatusCode(v int32) *ModifyMeetingPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMeetingPasswordResponse) SetBody(v *ModifyMeetingPasswordResponseBody) *ModifyMeetingPasswordResponse {
	s.Body = v
	return s
}

type ModifyMeetingPasswordInternationalRequest struct {
	MeetingUUID      *string `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
	OpenPasswordFlag *bool   `json:"OpenPasswordFlag,omitempty" xml:"OpenPasswordFlag,omitempty"`
	Password         *string `json:"Password,omitempty" xml:"Password,omitempty"`
	UserId           *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ModifyMeetingPasswordInternationalRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyMeetingPasswordInternationalRequest) GoString() string {
	return s.String()
}

func (s *ModifyMeetingPasswordInternationalRequest) SetMeetingUUID(v string) *ModifyMeetingPasswordInternationalRequest {
	s.MeetingUUID = &v
	return s
}

func (s *ModifyMeetingPasswordInternationalRequest) SetOpenPasswordFlag(v bool) *ModifyMeetingPasswordInternationalRequest {
	s.OpenPasswordFlag = &v
	return s
}

func (s *ModifyMeetingPasswordInternationalRequest) SetPassword(v string) *ModifyMeetingPasswordInternationalRequest {
	s.Password = &v
	return s
}

func (s *ModifyMeetingPasswordInternationalRequest) SetUserId(v string) *ModifyMeetingPasswordInternationalRequest {
	s.UserId = &v
	return s
}

type ModifyMeetingPasswordInternationalResponseBody struct {
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyMeetingPasswordInternationalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyMeetingPasswordInternationalResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMeetingPasswordInternationalResponseBody) SetErrorCode(v int32) *ModifyMeetingPasswordInternationalResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ModifyMeetingPasswordInternationalResponseBody) SetMessage(v string) *ModifyMeetingPasswordInternationalResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyMeetingPasswordInternationalResponseBody) SetRequestId(v string) *ModifyMeetingPasswordInternationalResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMeetingPasswordInternationalResponseBody) SetSuccess(v bool) *ModifyMeetingPasswordInternationalResponseBody {
	s.Success = &v
	return s
}

type ModifyMeetingPasswordInternationalResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyMeetingPasswordInternationalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyMeetingPasswordInternationalResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyMeetingPasswordInternationalResponse) GoString() string {
	return s.String()
}

func (s *ModifyMeetingPasswordInternationalResponse) SetHeaders(v map[string]*string) *ModifyMeetingPasswordInternationalResponse {
	s.Headers = v
	return s
}

func (s *ModifyMeetingPasswordInternationalResponse) SetStatusCode(v int32) *ModifyMeetingPasswordInternationalResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMeetingPasswordInternationalResponse) SetBody(v *ModifyMeetingPasswordInternationalResponseBody) *ModifyMeetingPasswordInternationalResponse {
	s.Body = v
	return s
}

type QueryMeetingMemberActionRequest struct {
	EndTime        *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MeetingUUID    *string `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
	MeetingUnitKey *string `json:"MeetingUnitKey,omitempty" xml:"MeetingUnitKey,omitempty"`
	MemberUUID     *string `json:"MemberUUID,omitempty" xml:"MemberUUID,omitempty"`
	StartTime      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryMeetingMemberActionRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingMemberActionRequest) GoString() string {
	return s.String()
}

func (s *QueryMeetingMemberActionRequest) SetEndTime(v int64) *QueryMeetingMemberActionRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMeetingMemberActionRequest) SetMeetingUUID(v string) *QueryMeetingMemberActionRequest {
	s.MeetingUUID = &v
	return s
}

func (s *QueryMeetingMemberActionRequest) SetMeetingUnitKey(v string) *QueryMeetingMemberActionRequest {
	s.MeetingUnitKey = &v
	return s
}

func (s *QueryMeetingMemberActionRequest) SetMemberUUID(v string) *QueryMeetingMemberActionRequest {
	s.MemberUUID = &v
	return s
}

func (s *QueryMeetingMemberActionRequest) SetStartTime(v int64) *QueryMeetingMemberActionRequest {
	s.StartTime = &v
	return s
}

type QueryMeetingMemberActionResponseBody struct {
	AudioStatus    *int32                                               `json:"AudioStatus,omitempty" xml:"AudioStatus,omitempty"`
	ErrorCode      *int32                                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorCodeCount *int64                                               `json:"ErrorCodeCount,omitempty" xml:"ErrorCodeCount,omitempty"`
	ErrorCodeList  []*QueryMeetingMemberActionResponseBodyErrorCodeList `json:"ErrorCodeList,omitempty" xml:"ErrorCodeList,omitempty" type:"Repeated"`
	MeetingStatus  *int32                                               `json:"MeetingStatus,omitempty" xml:"MeetingStatus,omitempty"`
	Message        *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                                `json:"Success,omitempty" xml:"Success,omitempty"`
	VideoStatus    *int32                                               `json:"VideoStatus,omitempty" xml:"VideoStatus,omitempty"`
}

func (s QueryMeetingMemberActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingMemberActionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMeetingMemberActionResponseBody) SetAudioStatus(v int32) *QueryMeetingMemberActionResponseBody {
	s.AudioStatus = &v
	return s
}

func (s *QueryMeetingMemberActionResponseBody) SetErrorCode(v int32) *QueryMeetingMemberActionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryMeetingMemberActionResponseBody) SetErrorCodeCount(v int64) *QueryMeetingMemberActionResponseBody {
	s.ErrorCodeCount = &v
	return s
}

func (s *QueryMeetingMemberActionResponseBody) SetErrorCodeList(v []*QueryMeetingMemberActionResponseBodyErrorCodeList) *QueryMeetingMemberActionResponseBody {
	s.ErrorCodeList = v
	return s
}

func (s *QueryMeetingMemberActionResponseBody) SetMeetingStatus(v int32) *QueryMeetingMemberActionResponseBody {
	s.MeetingStatus = &v
	return s
}

func (s *QueryMeetingMemberActionResponseBody) SetMessage(v string) *QueryMeetingMemberActionResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMeetingMemberActionResponseBody) SetRequestId(v string) *QueryMeetingMemberActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMeetingMemberActionResponseBody) SetSuccess(v bool) *QueryMeetingMemberActionResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMeetingMemberActionResponseBody) SetVideoStatus(v int32) *QueryMeetingMemberActionResponseBody {
	s.VideoStatus = &v
	return s
}

type QueryMeetingMemberActionResponseBodyErrorCodeList struct {
	ErrorCodeCount *int64 `json:"ErrorCodeCount,omitempty" xml:"ErrorCodeCount,omitempty"`
	Time           *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s QueryMeetingMemberActionResponseBodyErrorCodeList) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingMemberActionResponseBodyErrorCodeList) GoString() string {
	return s.String()
}

func (s *QueryMeetingMemberActionResponseBodyErrorCodeList) SetErrorCodeCount(v int64) *QueryMeetingMemberActionResponseBodyErrorCodeList {
	s.ErrorCodeCount = &v
	return s
}

func (s *QueryMeetingMemberActionResponseBodyErrorCodeList) SetTime(v int64) *QueryMeetingMemberActionResponseBodyErrorCodeList {
	s.Time = &v
	return s
}

type QueryMeetingMemberActionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMeetingMemberActionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMeetingMemberActionResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingMemberActionResponse) GoString() string {
	return s.String()
}

func (s *QueryMeetingMemberActionResponse) SetHeaders(v map[string]*string) *QueryMeetingMemberActionResponse {
	s.Headers = v
	return s
}

func (s *QueryMeetingMemberActionResponse) SetStatusCode(v int32) *QueryMeetingMemberActionResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMeetingMemberActionResponse) SetBody(v *QueryMeetingMemberActionResponseBody) *QueryMeetingMemberActionResponse {
	s.Body = v
	return s
}

type RefreshDeviceScreenCodeRequest struct {
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s RefreshDeviceScreenCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshDeviceScreenCodeRequest) GoString() string {
	return s.String()
}

func (s *RefreshDeviceScreenCodeRequest) SetSerialNumber(v string) *RefreshDeviceScreenCodeRequest {
	s.SerialNumber = &v
	return s
}

type RefreshDeviceScreenCodeResponseBody struct {
	Data      *RefreshDeviceScreenCodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *int32                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RefreshDeviceScreenCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshDeviceScreenCodeResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshDeviceScreenCodeResponseBody) SetData(v *RefreshDeviceScreenCodeResponseBodyData) *RefreshDeviceScreenCodeResponseBody {
	s.Data = v
	return s
}

func (s *RefreshDeviceScreenCodeResponseBody) SetErrorCode(v int32) *RefreshDeviceScreenCodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RefreshDeviceScreenCodeResponseBody) SetMessage(v string) *RefreshDeviceScreenCodeResponseBody {
	s.Message = &v
	return s
}

func (s *RefreshDeviceScreenCodeResponseBody) SetRequestId(v string) *RefreshDeviceScreenCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshDeviceScreenCodeResponseBody) SetSuccess(v bool) *RefreshDeviceScreenCodeResponseBody {
	s.Success = &v
	return s
}

type RefreshDeviceScreenCodeResponseBodyData struct {
	ScreenCode   *string `json:"ScreenCode,omitempty" xml:"ScreenCode,omitempty"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s RefreshDeviceScreenCodeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RefreshDeviceScreenCodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *RefreshDeviceScreenCodeResponseBodyData) SetScreenCode(v string) *RefreshDeviceScreenCodeResponseBodyData {
	s.ScreenCode = &v
	return s
}

func (s *RefreshDeviceScreenCodeResponseBodyData) SetSerialNumber(v string) *RefreshDeviceScreenCodeResponseBodyData {
	s.SerialNumber = &v
	return s
}

type RefreshDeviceScreenCodeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RefreshDeviceScreenCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefreshDeviceScreenCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshDeviceScreenCodeResponse) GoString() string {
	return s.String()
}

func (s *RefreshDeviceScreenCodeResponse) SetHeaders(v map[string]*string) *RefreshDeviceScreenCodeResponse {
	s.Headers = v
	return s
}

func (s *RefreshDeviceScreenCodeResponse) SetStatusCode(v int32) *RefreshDeviceScreenCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshDeviceScreenCodeResponse) SetBody(v *RefreshDeviceScreenCodeResponseBody) *RefreshDeviceScreenCodeResponse {
	s.Body = v
	return s
}

type RegisterDeviceRequest struct {
	DeviceVersion *string `json:"DeviceVersion,omitempty" xml:"DeviceVersion,omitempty"`
	IP            *string `json:"IP,omitempty" xml:"IP,omitempty"`
	Mac           *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	SN            *string `json:"SN,omitempty" xml:"SN,omitempty"`
	Token         *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s RegisterDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceRequest) GoString() string {
	return s.String()
}

func (s *RegisterDeviceRequest) SetDeviceVersion(v string) *RegisterDeviceRequest {
	s.DeviceVersion = &v
	return s
}

func (s *RegisterDeviceRequest) SetIP(v string) *RegisterDeviceRequest {
	s.IP = &v
	return s
}

func (s *RegisterDeviceRequest) SetMac(v string) *RegisterDeviceRequest {
	s.Mac = &v
	return s
}

func (s *RegisterDeviceRequest) SetSN(v string) *RegisterDeviceRequest {
	s.SN = &v
	return s
}

func (s *RegisterDeviceRequest) SetToken(v string) *RegisterDeviceRequest {
	s.Token = &v
	return s
}

type RegisterDeviceResponseBody struct {
	DeviceInfo *RegisterDeviceResponseBodyDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	ErrorCode  *int32                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message    *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RegisterDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterDeviceResponseBody) SetDeviceInfo(v *RegisterDeviceResponseBodyDeviceInfo) *RegisterDeviceResponseBody {
	s.DeviceInfo = v
	return s
}

func (s *RegisterDeviceResponseBody) SetErrorCode(v int32) *RegisterDeviceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RegisterDeviceResponseBody) SetMessage(v string) *RegisterDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *RegisterDeviceResponseBody) SetRequestId(v string) *RegisterDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterDeviceResponseBody) SetSuccess(v bool) *RegisterDeviceResponseBody {
	s.Success = &v
	return s
}

type RegisterDeviceResponseBodyDeviceInfo struct {
	AuthWsChannelConfig *RegisterDeviceResponseBodyDeviceInfoAuthWsChannelConfig `json:"AuthWsChannelConfig,omitempty" xml:"AuthWsChannelConfig,omitempty" type:"Struct"`
	ChannelType         *string                                                  `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	DeviceName          *string                                                  `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceSessionId     *string                                                  `json:"DeviceSessionId,omitempty" xml:"DeviceSessionId,omitempty"`
	MessageKey          *string                                                  `json:"MessageKey,omitempty" xml:"MessageKey,omitempty"`
	MqttParam           *RegisterDeviceResponseBodyDeviceInfoMqttParam           `json:"MqttParam,omitempty" xml:"MqttParam,omitempty" type:"Struct"`
	RegisterTime        *int64                                                   `json:"RegisterTime,omitempty" xml:"RegisterTime,omitempty"`
	ScreenCode          *string                                                  `json:"ScreenCode,omitempty" xml:"ScreenCode,omitempty"`
	SlsConfig           *RegisterDeviceResponseBodyDeviceInfoSlsConfig           `json:"SlsConfig,omitempty" xml:"SlsConfig,omitempty" type:"Struct"`
}

func (s RegisterDeviceResponseBodyDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceResponseBodyDeviceInfo) GoString() string {
	return s.String()
}

func (s *RegisterDeviceResponseBodyDeviceInfo) SetAuthWsChannelConfig(v *RegisterDeviceResponseBodyDeviceInfoAuthWsChannelConfig) *RegisterDeviceResponseBodyDeviceInfo {
	s.AuthWsChannelConfig = v
	return s
}

func (s *RegisterDeviceResponseBodyDeviceInfo) SetChannelType(v string) *RegisterDeviceResponseBodyDeviceInfo {
	s.ChannelType = &v
	return s
}

func (s *RegisterDeviceResponseBodyDeviceInfo) SetDeviceName(v string) *RegisterDeviceResponseBodyDeviceInfo {
	s.DeviceName = &v
	return s
}

func (s *RegisterDeviceResponseBodyDeviceInfo) SetDeviceSessionId(v string) *RegisterDeviceResponseBodyDeviceInfo {
	s.DeviceSessionId = &v
	return s
}

func (s *RegisterDeviceResponseBodyDeviceInfo) SetMessageKey(v string) *RegisterDeviceResponseBodyDeviceInfo {
	s.MessageKey = &v
	return s
}

func (s *RegisterDeviceResponseBodyDeviceInfo) SetMqttParam(v *RegisterDeviceResponseBodyDeviceInfoMqttParam) *RegisterDeviceResponseBodyDeviceInfo {
	s.MqttParam = v
	return s
}

func (s *RegisterDeviceResponseBodyDeviceInfo) SetRegisterTime(v int64) *RegisterDeviceResponseBodyDeviceInfo {
	s.RegisterTime = &v
	return s
}

func (s *RegisterDeviceResponseBodyDeviceInfo) SetScreenCode(v string) *RegisterDeviceResponseBodyDeviceInfo {
	s.ScreenCode = &v
	return s
}

func (s *RegisterDeviceResponseBodyDeviceInfo) SetSlsConfig(v *RegisterDeviceResponseBodyDeviceInfoSlsConfig) *RegisterDeviceResponseBodyDeviceInfo {
	s.SlsConfig = v
	return s
}

type RegisterDeviceResponseBodyDeviceInfoAuthWsChannelConfig struct {
	AuthWsUrl         *string `json:"AuthWsUrl,omitempty" xml:"AuthWsUrl,omitempty"`
	Token             *string `json:"Token,omitempty" xml:"Token,omitempty"`
	WsOuterReconnTime *int32  `json:"WsOuterReconnTime,omitempty" xml:"WsOuterReconnTime,omitempty"`
}

func (s RegisterDeviceResponseBodyDeviceInfoAuthWsChannelConfig) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceResponseBodyDeviceInfoAuthWsChannelConfig) GoString() string {
	return s.String()
}

func (s *RegisterDeviceResponseBodyDeviceInfoAuthWsChannelConfig) SetAuthWsUrl(v string) *RegisterDeviceResponseBodyDeviceInfoAuthWsChannelConfig {
	s.AuthWsUrl = &v
	return s
}

func (s *RegisterDeviceResponseBodyDeviceInfoAuthWsChannelConfig) SetToken(v string) *RegisterDeviceResponseBodyDeviceInfoAuthWsChannelConfig {
	s.Token = &v
	return s
}

func (s *RegisterDeviceResponseBodyDeviceInfoAuthWsChannelConfig) SetWsOuterReconnTime(v int32) *RegisterDeviceResponseBodyDeviceInfoAuthWsChannelConfig {
	s.WsOuterReconnTime = &v
	return s
}

type RegisterDeviceResponseBodyDeviceInfoMqttParam struct {
	CleanSession     *string `json:"CleanSession,omitempty" xml:"CleanSession,omitempty"`
	ClientId         *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	GroupId          *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Host             *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Password         *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Port             *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ReconnectTimeout *string `json:"ReconnectTimeout,omitempty" xml:"ReconnectTimeout,omitempty"`
	SDKClientPort    *string `json:"SDKClientPort,omitempty" xml:"SDKClientPort,omitempty"`
	TLSPort          *string `json:"TLSPort,omitempty" xml:"TLSPort,omitempty"`
	Topic            *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	UseTLS           *string `json:"UseTLS,omitempty" xml:"UseTLS,omitempty"`
	UserName         *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s RegisterDeviceResponseBodyDeviceInfoMqttParam) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceResponseBodyDeviceInfoMqttParam) GoString() string {
	return s.String()
}

func (s *RegisterDeviceResponseBodyDeviceInfoMqttParam) SetCleanSession(v string) *RegisterDeviceResponseBodyDeviceInfoMqttParam {
	s.CleanSession = &v
	return s
}

func (s *RegisterDeviceResponseBodyDeviceInfoMqttParam) SetClientId(v string) *RegisterDeviceResponseBodyDeviceInfoMqttParam {
	s.ClientId = &v
	return s
}

func (s *RegisterDeviceResponseBodyDeviceInfoMqttParam) SetGroupId(v string) *RegisterDeviceResponseBodyDeviceInfoMqttParam {
	s.GroupId = &v
	return s
}

func (s *RegisterDeviceResponseBodyDeviceInfoMqttParam) SetHost(v string) *RegisterDeviceResponseBodyDeviceInfoMqttParam {
	s.Host = &v
	return s
}

func (s *RegisterDeviceResponseBodyDeviceInfoMqttParam) SetPassword(v string) *RegisterDeviceResponseBodyDeviceInfoMqttParam {
	s.Password = &v
	return s
}

func (s *RegisterDeviceResponseBodyDeviceInfoMqttParam) SetPort(v string) *RegisterDeviceResponseBodyDeviceInfoMqttParam {
	s.Port = &v
	return s
}

func (s *RegisterDeviceResponseBodyDeviceInfoMqttParam) SetReconnectTimeout(v string) *RegisterDeviceResponseBodyDeviceInfoMqttParam {
	s.ReconnectTimeout = &v
	return s
}

func (s *RegisterDeviceResponseBodyDeviceInfoMqttParam) SetSDKClientPort(v string) *RegisterDeviceResponseBodyDeviceInfoMqttParam {
	s.SDKClientPort = &v
	return s
}

func (s *RegisterDeviceResponseBodyDeviceInfoMqttParam) SetTLSPort(v string) *RegisterDeviceResponseBodyDeviceInfoMqttParam {
	s.TLSPort = &v
	return s
}

func (s *RegisterDeviceResponseBodyDeviceInfoMqttParam) SetTopic(v string) *RegisterDeviceResponseBodyDeviceInfoMqttParam {
	s.Topic = &v
	return s
}

func (s *RegisterDeviceResponseBodyDeviceInfoMqttParam) SetUseTLS(v string) *RegisterDeviceResponseBodyDeviceInfoMqttParam {
	s.UseTLS = &v
	return s
}

func (s *RegisterDeviceResponseBodyDeviceInfoMqttParam) SetUserName(v string) *RegisterDeviceResponseBodyDeviceInfoMqttParam {
	s.UserName = &v
	return s
}

type RegisterDeviceResponseBodyDeviceInfoSlsConfig struct {
	LogServiceEndpoint *string `json:"LogServiceEndpoint,omitempty" xml:"LogServiceEndpoint,omitempty"`
	LogStore           *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	Project            *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s RegisterDeviceResponseBodyDeviceInfoSlsConfig) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceResponseBodyDeviceInfoSlsConfig) GoString() string {
	return s.String()
}

func (s *RegisterDeviceResponseBodyDeviceInfoSlsConfig) SetLogServiceEndpoint(v string) *RegisterDeviceResponseBodyDeviceInfoSlsConfig {
	s.LogServiceEndpoint = &v
	return s
}

func (s *RegisterDeviceResponseBodyDeviceInfoSlsConfig) SetLogStore(v string) *RegisterDeviceResponseBodyDeviceInfoSlsConfig {
	s.LogStore = &v
	return s
}

func (s *RegisterDeviceResponseBodyDeviceInfoSlsConfig) SetProject(v string) *RegisterDeviceResponseBodyDeviceInfoSlsConfig {
	s.Project = &v
	return s
}

type RegisterDeviceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RegisterDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceResponse) GoString() string {
	return s.String()
}

func (s *RegisterDeviceResponse) SetHeaders(v map[string]*string) *RegisterDeviceResponse {
	s.Headers = v
	return s
}

func (s *RegisterDeviceResponse) SetStatusCode(v int32) *RegisterDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterDeviceResponse) SetBody(v *RegisterDeviceResponseBody) *RegisterDeviceResponse {
	s.Body = v
	return s
}

type RegisterUemDeviceRequest struct {
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DeviceVersion *string `json:"DeviceVersion,omitempty" xml:"DeviceVersion,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	IP            *string `json:"IP,omitempty" xml:"IP,omitempty"`
	Mac           *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	OwnerId       *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s RegisterUemDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterUemDeviceRequest) GoString() string {
	return s.String()
}

func (s *RegisterUemDeviceRequest) SetDeviceId(v string) *RegisterUemDeviceRequest {
	s.DeviceId = &v
	return s
}

func (s *RegisterUemDeviceRequest) SetDeviceVersion(v string) *RegisterUemDeviceRequest {
	s.DeviceVersion = &v
	return s
}

func (s *RegisterUemDeviceRequest) SetGroupId(v string) *RegisterUemDeviceRequest {
	s.GroupId = &v
	return s
}

func (s *RegisterUemDeviceRequest) SetGroupName(v string) *RegisterUemDeviceRequest {
	s.GroupName = &v
	return s
}

func (s *RegisterUemDeviceRequest) SetIP(v string) *RegisterUemDeviceRequest {
	s.IP = &v
	return s
}

func (s *RegisterUemDeviceRequest) SetMac(v string) *RegisterUemDeviceRequest {
	s.Mac = &v
	return s
}

func (s *RegisterUemDeviceRequest) SetOwnerId(v string) *RegisterUemDeviceRequest {
	s.OwnerId = &v
	return s
}

type RegisterUemDeviceResponseBody struct {
	DeviceInfo *RegisterUemDeviceResponseBodyDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	ErrorCode  *int32                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message    *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RegisterUemDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterUemDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterUemDeviceResponseBody) SetDeviceInfo(v *RegisterUemDeviceResponseBodyDeviceInfo) *RegisterUemDeviceResponseBody {
	s.DeviceInfo = v
	return s
}

func (s *RegisterUemDeviceResponseBody) SetErrorCode(v int32) *RegisterUemDeviceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RegisterUemDeviceResponseBody) SetMessage(v string) *RegisterUemDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *RegisterUemDeviceResponseBody) SetRequestId(v string) *RegisterUemDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterUemDeviceResponseBody) SetSuccess(v bool) *RegisterUemDeviceResponseBody {
	s.Success = &v
	return s
}

type RegisterUemDeviceResponseBodyDeviceInfo struct {
	AuthWsChannelConfig *RegisterUemDeviceResponseBodyDeviceInfoAuthWsChannelConfig `json:"AuthWsChannelConfig,omitempty" xml:"AuthWsChannelConfig,omitempty" type:"Struct"`
	ChannelType         *string                                                     `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	DeviceName          *string                                                     `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceSessionId     *string                                                     `json:"DeviceSessionId,omitempty" xml:"DeviceSessionId,omitempty"`
	MessageKey          *string                                                     `json:"MessageKey,omitempty" xml:"MessageKey,omitempty"`
	MqttParam           *RegisterUemDeviceResponseBodyDeviceInfoMqttParam           `json:"MqttParam,omitempty" xml:"MqttParam,omitempty" type:"Struct"`
	RegisterTime        *int64                                                      `json:"RegisterTime,omitempty" xml:"RegisterTime,omitempty"`
	ScreenCode          *string                                                     `json:"ScreenCode,omitempty" xml:"ScreenCode,omitempty"`
	SlsConfig           *RegisterUemDeviceResponseBodyDeviceInfoSlsConfig           `json:"SlsConfig,omitempty" xml:"SlsConfig,omitempty" type:"Struct"`
	Token               *string                                                     `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s RegisterUemDeviceResponseBodyDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s RegisterUemDeviceResponseBodyDeviceInfo) GoString() string {
	return s.String()
}

func (s *RegisterUemDeviceResponseBodyDeviceInfo) SetAuthWsChannelConfig(v *RegisterUemDeviceResponseBodyDeviceInfoAuthWsChannelConfig) *RegisterUemDeviceResponseBodyDeviceInfo {
	s.AuthWsChannelConfig = v
	return s
}

func (s *RegisterUemDeviceResponseBodyDeviceInfo) SetChannelType(v string) *RegisterUemDeviceResponseBodyDeviceInfo {
	s.ChannelType = &v
	return s
}

func (s *RegisterUemDeviceResponseBodyDeviceInfo) SetDeviceName(v string) *RegisterUemDeviceResponseBodyDeviceInfo {
	s.DeviceName = &v
	return s
}

func (s *RegisterUemDeviceResponseBodyDeviceInfo) SetDeviceSessionId(v string) *RegisterUemDeviceResponseBodyDeviceInfo {
	s.DeviceSessionId = &v
	return s
}

func (s *RegisterUemDeviceResponseBodyDeviceInfo) SetMessageKey(v string) *RegisterUemDeviceResponseBodyDeviceInfo {
	s.MessageKey = &v
	return s
}

func (s *RegisterUemDeviceResponseBodyDeviceInfo) SetMqttParam(v *RegisterUemDeviceResponseBodyDeviceInfoMqttParam) *RegisterUemDeviceResponseBodyDeviceInfo {
	s.MqttParam = v
	return s
}

func (s *RegisterUemDeviceResponseBodyDeviceInfo) SetRegisterTime(v int64) *RegisterUemDeviceResponseBodyDeviceInfo {
	s.RegisterTime = &v
	return s
}

func (s *RegisterUemDeviceResponseBodyDeviceInfo) SetScreenCode(v string) *RegisterUemDeviceResponseBodyDeviceInfo {
	s.ScreenCode = &v
	return s
}

func (s *RegisterUemDeviceResponseBodyDeviceInfo) SetSlsConfig(v *RegisterUemDeviceResponseBodyDeviceInfoSlsConfig) *RegisterUemDeviceResponseBodyDeviceInfo {
	s.SlsConfig = v
	return s
}

func (s *RegisterUemDeviceResponseBodyDeviceInfo) SetToken(v string) *RegisterUemDeviceResponseBodyDeviceInfo {
	s.Token = &v
	return s
}

type RegisterUemDeviceResponseBodyDeviceInfoAuthWsChannelConfig struct {
	AuthWsUrl         *string `json:"AuthWsUrl,omitempty" xml:"AuthWsUrl,omitempty"`
	Token             *string `json:"Token,omitempty" xml:"Token,omitempty"`
	WsOuterReconnTime *int32  `json:"WsOuterReconnTime,omitempty" xml:"WsOuterReconnTime,omitempty"`
}

func (s RegisterUemDeviceResponseBodyDeviceInfoAuthWsChannelConfig) String() string {
	return tea.Prettify(s)
}

func (s RegisterUemDeviceResponseBodyDeviceInfoAuthWsChannelConfig) GoString() string {
	return s.String()
}

func (s *RegisterUemDeviceResponseBodyDeviceInfoAuthWsChannelConfig) SetAuthWsUrl(v string) *RegisterUemDeviceResponseBodyDeviceInfoAuthWsChannelConfig {
	s.AuthWsUrl = &v
	return s
}

func (s *RegisterUemDeviceResponseBodyDeviceInfoAuthWsChannelConfig) SetToken(v string) *RegisterUemDeviceResponseBodyDeviceInfoAuthWsChannelConfig {
	s.Token = &v
	return s
}

func (s *RegisterUemDeviceResponseBodyDeviceInfoAuthWsChannelConfig) SetWsOuterReconnTime(v int32) *RegisterUemDeviceResponseBodyDeviceInfoAuthWsChannelConfig {
	s.WsOuterReconnTime = &v
	return s
}

type RegisterUemDeviceResponseBodyDeviceInfoMqttParam struct {
	CleanSession     *string `json:"CleanSession,omitempty" xml:"CleanSession,omitempty"`
	ClientId         *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	GroupId          *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Host             *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Password         *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Port             *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ReconnectTimeout *string `json:"ReconnectTimeout,omitempty" xml:"ReconnectTimeout,omitempty"`
	SDKClientPort    *string `json:"SDKClientPort,omitempty" xml:"SDKClientPort,omitempty"`
	TLSPort          *string `json:"TLSPort,omitempty" xml:"TLSPort,omitempty"`
	Topic            *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	UseTLS           *string `json:"UseTLS,omitempty" xml:"UseTLS,omitempty"`
	UserName         *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s RegisterUemDeviceResponseBodyDeviceInfoMqttParam) String() string {
	return tea.Prettify(s)
}

func (s RegisterUemDeviceResponseBodyDeviceInfoMqttParam) GoString() string {
	return s.String()
}

func (s *RegisterUemDeviceResponseBodyDeviceInfoMqttParam) SetCleanSession(v string) *RegisterUemDeviceResponseBodyDeviceInfoMqttParam {
	s.CleanSession = &v
	return s
}

func (s *RegisterUemDeviceResponseBodyDeviceInfoMqttParam) SetClientId(v string) *RegisterUemDeviceResponseBodyDeviceInfoMqttParam {
	s.ClientId = &v
	return s
}

func (s *RegisterUemDeviceResponseBodyDeviceInfoMqttParam) SetGroupId(v string) *RegisterUemDeviceResponseBodyDeviceInfoMqttParam {
	s.GroupId = &v
	return s
}

func (s *RegisterUemDeviceResponseBodyDeviceInfoMqttParam) SetHost(v string) *RegisterUemDeviceResponseBodyDeviceInfoMqttParam {
	s.Host = &v
	return s
}

func (s *RegisterUemDeviceResponseBodyDeviceInfoMqttParam) SetPassword(v string) *RegisterUemDeviceResponseBodyDeviceInfoMqttParam {
	s.Password = &v
	return s
}

func (s *RegisterUemDeviceResponseBodyDeviceInfoMqttParam) SetPort(v string) *RegisterUemDeviceResponseBodyDeviceInfoMqttParam {
	s.Port = &v
	return s
}

func (s *RegisterUemDeviceResponseBodyDeviceInfoMqttParam) SetReconnectTimeout(v string) *RegisterUemDeviceResponseBodyDeviceInfoMqttParam {
	s.ReconnectTimeout = &v
	return s
}

func (s *RegisterUemDeviceResponseBodyDeviceInfoMqttParam) SetSDKClientPort(v string) *RegisterUemDeviceResponseBodyDeviceInfoMqttParam {
	s.SDKClientPort = &v
	return s
}

func (s *RegisterUemDeviceResponseBodyDeviceInfoMqttParam) SetTLSPort(v string) *RegisterUemDeviceResponseBodyDeviceInfoMqttParam {
	s.TLSPort = &v
	return s
}

func (s *RegisterUemDeviceResponseBodyDeviceInfoMqttParam) SetTopic(v string) *RegisterUemDeviceResponseBodyDeviceInfoMqttParam {
	s.Topic = &v
	return s
}

func (s *RegisterUemDeviceResponseBodyDeviceInfoMqttParam) SetUseTLS(v string) *RegisterUemDeviceResponseBodyDeviceInfoMqttParam {
	s.UseTLS = &v
	return s
}

func (s *RegisterUemDeviceResponseBodyDeviceInfoMqttParam) SetUserName(v string) *RegisterUemDeviceResponseBodyDeviceInfoMqttParam {
	s.UserName = &v
	return s
}

type RegisterUemDeviceResponseBodyDeviceInfoSlsConfig struct {
	LogServiceEndpoint *string `json:"LogServiceEndpoint,omitempty" xml:"LogServiceEndpoint,omitempty"`
	LogStore           *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	Project            *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s RegisterUemDeviceResponseBodyDeviceInfoSlsConfig) String() string {
	return tea.Prettify(s)
}

func (s RegisterUemDeviceResponseBodyDeviceInfoSlsConfig) GoString() string {
	return s.String()
}

func (s *RegisterUemDeviceResponseBodyDeviceInfoSlsConfig) SetLogServiceEndpoint(v string) *RegisterUemDeviceResponseBodyDeviceInfoSlsConfig {
	s.LogServiceEndpoint = &v
	return s
}

func (s *RegisterUemDeviceResponseBodyDeviceInfoSlsConfig) SetLogStore(v string) *RegisterUemDeviceResponseBodyDeviceInfoSlsConfig {
	s.LogStore = &v
	return s
}

func (s *RegisterUemDeviceResponseBodyDeviceInfoSlsConfig) SetProject(v string) *RegisterUemDeviceResponseBodyDeviceInfoSlsConfig {
	s.Project = &v
	return s
}

type RegisterUemDeviceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RegisterUemDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterUemDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterUemDeviceResponse) GoString() string {
	return s.String()
}

func (s *RegisterUemDeviceResponse) SetHeaders(v map[string]*string) *RegisterUemDeviceResponse {
	s.Headers = v
	return s
}

func (s *RegisterUemDeviceResponse) SetStatusCode(v int32) *RegisterUemDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterUemDeviceResponse) SetBody(v *RegisterUemDeviceResponseBody) *RegisterUemDeviceResponse {
	s.Body = v
	return s
}

type SendMeetingCommandRequest struct {
	Command            *string `json:"Command,omitempty" xml:"Command,omitempty"`
	MeetingUUID        *string `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
	MemberUUID         *string `json:"MemberUUID,omitempty" xml:"MemberUUID,omitempty"`
	OperatorMemberUUID *string `json:"OperatorMemberUUID,omitempty" xml:"OperatorMemberUUID,omitempty"`
	SendType           *int32  `json:"SendType,omitempty" xml:"SendType,omitempty"`
}

func (s SendMeetingCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMeetingCommandRequest) GoString() string {
	return s.String()
}

func (s *SendMeetingCommandRequest) SetCommand(v string) *SendMeetingCommandRequest {
	s.Command = &v
	return s
}

func (s *SendMeetingCommandRequest) SetMeetingUUID(v string) *SendMeetingCommandRequest {
	s.MeetingUUID = &v
	return s
}

func (s *SendMeetingCommandRequest) SetMemberUUID(v string) *SendMeetingCommandRequest {
	s.MemberUUID = &v
	return s
}

func (s *SendMeetingCommandRequest) SetOperatorMemberUUID(v string) *SendMeetingCommandRequest {
	s.OperatorMemberUUID = &v
	return s
}

func (s *SendMeetingCommandRequest) SetSendType(v int32) *SendMeetingCommandRequest {
	s.SendType = &v
	return s
}

type SendMeetingCommandResponseBody struct {
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendMeetingCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMeetingCommandResponseBody) GoString() string {
	return s.String()
}

func (s *SendMeetingCommandResponseBody) SetErrorCode(v int32) *SendMeetingCommandResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SendMeetingCommandResponseBody) SetMessage(v string) *SendMeetingCommandResponseBody {
	s.Message = &v
	return s
}

func (s *SendMeetingCommandResponseBody) SetRequestId(v string) *SendMeetingCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendMeetingCommandResponseBody) SetSuccess(v bool) *SendMeetingCommandResponseBody {
	s.Success = &v
	return s
}

type SendMeetingCommandResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendMeetingCommandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendMeetingCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s SendMeetingCommandResponse) GoString() string {
	return s.String()
}

func (s *SendMeetingCommandResponse) SetHeaders(v map[string]*string) *SendMeetingCommandResponse {
	s.Headers = v
	return s
}

func (s *SendMeetingCommandResponse) SetStatusCode(v int32) *SendMeetingCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *SendMeetingCommandResponse) SetBody(v *SendMeetingCommandResponseBody) *SendMeetingCommandResponse {
	s.Body = v
	return s
}

type SendScreenStartRequest struct {
	CastScreenCode *string `json:"CastScreenCode,omitempty" xml:"CastScreenCode,omitempty"`
	SessionId      *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s SendScreenStartRequest) String() string {
	return tea.Prettify(s)
}

func (s SendScreenStartRequest) GoString() string {
	return s.String()
}

func (s *SendScreenStartRequest) SetCastScreenCode(v string) *SendScreenStartRequest {
	s.CastScreenCode = &v
	return s
}

func (s *SendScreenStartRequest) SetSessionId(v string) *SendScreenStartRequest {
	s.SessionId = &v
	return s
}

type SendScreenStartResponseBody struct {
	Data      *SendScreenStartResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *int32                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendScreenStartResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendScreenStartResponseBody) GoString() string {
	return s.String()
}

func (s *SendScreenStartResponseBody) SetData(v *SendScreenStartResponseBodyData) *SendScreenStartResponseBody {
	s.Data = v
	return s
}

func (s *SendScreenStartResponseBody) SetErrorCode(v int32) *SendScreenStartResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SendScreenStartResponseBody) SetMessage(v string) *SendScreenStartResponseBody {
	s.Message = &v
	return s
}

func (s *SendScreenStartResponseBody) SetRequestId(v string) *SendScreenStartResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendScreenStartResponseBody) SetSuccess(v bool) *SendScreenStartResponseBody {
	s.Success = &v
	return s
}

type SendScreenStartResponseBodyData struct {
	AuthWsChannelConfig    *SendScreenStartResponseBodyDataAuthWsChannelConfig    `json:"AuthWsChannelConfig,omitempty" xml:"AuthWsChannelConfig,omitempty" type:"Struct"`
	IceServers             []*SendScreenStartResponseBodyDataIceServers           `json:"IceServers,omitempty" xml:"IceServers,omitempty" type:"Repeated"`
	ReceiverInfo           map[string]interface{}                                 `json:"ReceiverInfo,omitempty" xml:"ReceiverInfo,omitempty"`
	SessionId              *string                                                `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	ShareConfig            *SendScreenStartResponseBodyDataShareConfig            `json:"ShareConfig,omitempty" xml:"ShareConfig,omitempty" type:"Struct"`
	SignallingServerConfig *SendScreenStartResponseBodyDataSignallingServerConfig `json:"SignallingServerConfig,omitempty" xml:"SignallingServerConfig,omitempty" type:"Struct"`
}

func (s SendScreenStartResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SendScreenStartResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendScreenStartResponseBodyData) SetAuthWsChannelConfig(v *SendScreenStartResponseBodyDataAuthWsChannelConfig) *SendScreenStartResponseBodyData {
	s.AuthWsChannelConfig = v
	return s
}

func (s *SendScreenStartResponseBodyData) SetIceServers(v []*SendScreenStartResponseBodyDataIceServers) *SendScreenStartResponseBodyData {
	s.IceServers = v
	return s
}

func (s *SendScreenStartResponseBodyData) SetReceiverInfo(v map[string]interface{}) *SendScreenStartResponseBodyData {
	s.ReceiverInfo = v
	return s
}

func (s *SendScreenStartResponseBodyData) SetSessionId(v string) *SendScreenStartResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *SendScreenStartResponseBodyData) SetShareConfig(v *SendScreenStartResponseBodyDataShareConfig) *SendScreenStartResponseBodyData {
	s.ShareConfig = v
	return s
}

func (s *SendScreenStartResponseBodyData) SetSignallingServerConfig(v *SendScreenStartResponseBodyDataSignallingServerConfig) *SendScreenStartResponseBodyData {
	s.SignallingServerConfig = v
	return s
}

type SendScreenStartResponseBodyDataAuthWsChannelConfig struct {
	AuthWsUrl         *string `json:"AuthWsUrl,omitempty" xml:"AuthWsUrl,omitempty"`
	Token             *string `json:"Token,omitempty" xml:"Token,omitempty"`
	WsOuterReConnTime *int32  `json:"WsOuterReConnTime,omitempty" xml:"WsOuterReConnTime,omitempty"`
}

func (s SendScreenStartResponseBodyDataAuthWsChannelConfig) String() string {
	return tea.Prettify(s)
}

func (s SendScreenStartResponseBodyDataAuthWsChannelConfig) GoString() string {
	return s.String()
}

func (s *SendScreenStartResponseBodyDataAuthWsChannelConfig) SetAuthWsUrl(v string) *SendScreenStartResponseBodyDataAuthWsChannelConfig {
	s.AuthWsUrl = &v
	return s
}

func (s *SendScreenStartResponseBodyDataAuthWsChannelConfig) SetToken(v string) *SendScreenStartResponseBodyDataAuthWsChannelConfig {
	s.Token = &v
	return s
}

func (s *SendScreenStartResponseBodyDataAuthWsChannelConfig) SetWsOuterReConnTime(v int32) *SendScreenStartResponseBodyDataAuthWsChannelConfig {
	s.WsOuterReConnTime = &v
	return s
}

type SendScreenStartResponseBodyDataIceServers struct {
	Credential *string `json:"Credential,omitempty" xml:"Credential,omitempty"`
	Url        *string `json:"Url,omitempty" xml:"Url,omitempty"`
	UserName   *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s SendScreenStartResponseBodyDataIceServers) String() string {
	return tea.Prettify(s)
}

func (s SendScreenStartResponseBodyDataIceServers) GoString() string {
	return s.String()
}

func (s *SendScreenStartResponseBodyDataIceServers) SetCredential(v string) *SendScreenStartResponseBodyDataIceServers {
	s.Credential = &v
	return s
}

func (s *SendScreenStartResponseBodyDataIceServers) SetUrl(v string) *SendScreenStartResponseBodyDataIceServers {
	s.Url = &v
	return s
}

func (s *SendScreenStartResponseBodyDataIceServers) SetUserName(v string) *SendScreenStartResponseBodyDataIceServers {
	s.UserName = &v
	return s
}

type SendScreenStartResponseBodyDataShareConfig struct {
	MaxMultiScreenShareBitrate *int32 `json:"MaxMultiScreenShareBitrate,omitempty" xml:"MaxMultiScreenShareBitrate,omitempty"`
	MaxScreenShareBitrate      *int32 `json:"MaxScreenShareBitrate,omitempty" xml:"MaxScreenShareBitrate,omitempty"`
	ShareConfigMaxFrameRate    *int32 `json:"ShareConfigMaxFrameRate,omitempty" xml:"ShareConfigMaxFrameRate,omitempty"`
	ShareConfigMinFrameRate    *int32 `json:"ShareConfigMinFrameRate,omitempty" xml:"ShareConfigMinFrameRate,omitempty"`
}

func (s SendScreenStartResponseBodyDataShareConfig) String() string {
	return tea.Prettify(s)
}

func (s SendScreenStartResponseBodyDataShareConfig) GoString() string {
	return s.String()
}

func (s *SendScreenStartResponseBodyDataShareConfig) SetMaxMultiScreenShareBitrate(v int32) *SendScreenStartResponseBodyDataShareConfig {
	s.MaxMultiScreenShareBitrate = &v
	return s
}

func (s *SendScreenStartResponseBodyDataShareConfig) SetMaxScreenShareBitrate(v int32) *SendScreenStartResponseBodyDataShareConfig {
	s.MaxScreenShareBitrate = &v
	return s
}

func (s *SendScreenStartResponseBodyDataShareConfig) SetShareConfigMaxFrameRate(v int32) *SendScreenStartResponseBodyDataShareConfig {
	s.ShareConfigMaxFrameRate = &v
	return s
}

func (s *SendScreenStartResponseBodyDataShareConfig) SetShareConfigMinFrameRate(v int32) *SendScreenStartResponseBodyDataShareConfig {
	s.ShareConfigMinFrameRate = &v
	return s
}

type SendScreenStartResponseBodyDataSignallingServerConfig struct {
	ClientHeartBeatUrl       *string `json:"ClientHeartBeatUrl,omitempty" xml:"ClientHeartBeatUrl,omitempty"`
	GetMessageUrl            *string `json:"GetMessageUrl,omitempty" xml:"GetMessageUrl,omitempty"`
	GetReceiverInfoUrl       *string `json:"GetReceiverInfoUrl,omitempty" xml:"GetReceiverInfoUrl,omitempty"`
	SendMessageUrl           *string `json:"SendMessageUrl,omitempty" xml:"SendMessageUrl,omitempty"`
	ShareScreenMeetingRobUrl *string `json:"ShareScreenMeetingRobUrl,omitempty" xml:"ShareScreenMeetingRobUrl,omitempty"`
}

func (s SendScreenStartResponseBodyDataSignallingServerConfig) String() string {
	return tea.Prettify(s)
}

func (s SendScreenStartResponseBodyDataSignallingServerConfig) GoString() string {
	return s.String()
}

func (s *SendScreenStartResponseBodyDataSignallingServerConfig) SetClientHeartBeatUrl(v string) *SendScreenStartResponseBodyDataSignallingServerConfig {
	s.ClientHeartBeatUrl = &v
	return s
}

func (s *SendScreenStartResponseBodyDataSignallingServerConfig) SetGetMessageUrl(v string) *SendScreenStartResponseBodyDataSignallingServerConfig {
	s.GetMessageUrl = &v
	return s
}

func (s *SendScreenStartResponseBodyDataSignallingServerConfig) SetGetReceiverInfoUrl(v string) *SendScreenStartResponseBodyDataSignallingServerConfig {
	s.GetReceiverInfoUrl = &v
	return s
}

func (s *SendScreenStartResponseBodyDataSignallingServerConfig) SetSendMessageUrl(v string) *SendScreenStartResponseBodyDataSignallingServerConfig {
	s.SendMessageUrl = &v
	return s
}

func (s *SendScreenStartResponseBodyDataSignallingServerConfig) SetShareScreenMeetingRobUrl(v string) *SendScreenStartResponseBodyDataSignallingServerConfig {
	s.ShareScreenMeetingRobUrl = &v
	return s
}

type SendScreenStartResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendScreenStartResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendScreenStartResponse) String() string {
	return tea.Prettify(s)
}

func (s SendScreenStartResponse) GoString() string {
	return s.String()
}

func (s *SendScreenStartResponse) SetHeaders(v map[string]*string) *SendScreenStartResponse {
	s.Headers = v
	return s
}

func (s *SendScreenStartResponse) SetStatusCode(v int32) *SendScreenStartResponse {
	s.StatusCode = &v
	return s
}

func (s *SendScreenStartResponse) SetBody(v *SendScreenStartResponseBody) *SendScreenStartResponse {
	s.Body = v
	return s
}

type StartLiveRequest struct {
	LayoutInfo *string `json:"LayoutInfo,omitempty" xml:"LayoutInfo,omitempty"`
	LiveUUID   *string `json:"LiveUUID,omitempty" xml:"LiveUUID,omitempty"`
	PushInfo   *string `json:"PushInfo,omitempty" xml:"PushInfo,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s StartLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s StartLiveRequest) GoString() string {
	return s.String()
}

func (s *StartLiveRequest) SetLayoutInfo(v string) *StartLiveRequest {
	s.LayoutInfo = &v
	return s
}

func (s *StartLiveRequest) SetLiveUUID(v string) *StartLiveRequest {
	s.LiveUUID = &v
	return s
}

func (s *StartLiveRequest) SetPushInfo(v string) *StartLiveRequest {
	s.PushInfo = &v
	return s
}

func (s *StartLiveRequest) SetUserId(v string) *StartLiveRequest {
	s.UserId = &v
	return s
}

type StartLiveResponseBody struct {
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartLiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartLiveResponseBody) GoString() string {
	return s.String()
}

func (s *StartLiveResponseBody) SetErrorCode(v int32) *StartLiveResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StartLiveResponseBody) SetMessage(v string) *StartLiveResponseBody {
	s.Message = &v
	return s
}

func (s *StartLiveResponseBody) SetRequestId(v string) *StartLiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartLiveResponseBody) SetSuccess(v bool) *StartLiveResponseBody {
	s.Success = &v
	return s
}

type StartLiveResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartLiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartLiveResponse) String() string {
	return tea.Prettify(s)
}

func (s StartLiveResponse) GoString() string {
	return s.String()
}

func (s *StartLiveResponse) SetHeaders(v map[string]*string) *StartLiveResponse {
	s.Headers = v
	return s
}

func (s *StartLiveResponse) SetStatusCode(v int32) *StartLiveResponse {
	s.StatusCode = &v
	return s
}

func (s *StartLiveResponse) SetBody(v *StartLiveResponseBody) *StartLiveResponse {
	s.Body = v
	return s
}

type UpdateDeviceHeartBeatRequest struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s UpdateDeviceHeartBeatRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceHeartBeatRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeviceHeartBeatRequest) SetMessage(v string) *UpdateDeviceHeartBeatRequest {
	s.Message = &v
	return s
}

type UpdateDeviceHeartBeatResponseBody struct {
	DeviceInfo *UpdateDeviceHeartBeatResponseBodyDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	ErrorCode  *int32                                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message    *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDeviceHeartBeatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceHeartBeatResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeviceHeartBeatResponseBody) SetDeviceInfo(v *UpdateDeviceHeartBeatResponseBodyDeviceInfo) *UpdateDeviceHeartBeatResponseBody {
	s.DeviceInfo = v
	return s
}

func (s *UpdateDeviceHeartBeatResponseBody) SetErrorCode(v int32) *UpdateDeviceHeartBeatResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDeviceHeartBeatResponseBody) SetMessage(v string) *UpdateDeviceHeartBeatResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDeviceHeartBeatResponseBody) SetRequestId(v string) *UpdateDeviceHeartBeatResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDeviceHeartBeatResponseBody) SetSuccess(v bool) *UpdateDeviceHeartBeatResponseBody {
	s.Success = &v
	return s
}

type UpdateDeviceHeartBeatResponseBodyDeviceInfo struct {
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
}

func (s UpdateDeviceHeartBeatResponseBodyDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceHeartBeatResponseBodyDeviceInfo) GoString() string {
	return s.String()
}

func (s *UpdateDeviceHeartBeatResponseBodyDeviceInfo) SetChannelType(v string) *UpdateDeviceHeartBeatResponseBodyDeviceInfo {
	s.ChannelType = &v
	return s
}

type UpdateDeviceHeartBeatResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDeviceHeartBeatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDeviceHeartBeatResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceHeartBeatResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeviceHeartBeatResponse) SetHeaders(v map[string]*string) *UpdateDeviceHeartBeatResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeviceHeartBeatResponse) SetStatusCode(v int32) *UpdateDeviceHeartBeatResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDeviceHeartBeatResponse) SetBody(v *UpdateDeviceHeartBeatResponseBody) *UpdateDeviceHeartBeatResponse {
	s.Body = v
	return s
}

type UpdateDeviceStartupPictureRequest struct {
	Picture      *string `json:"Picture,omitempty" xml:"Picture,omitempty"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s UpdateDeviceStartupPictureRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceStartupPictureRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeviceStartupPictureRequest) SetPicture(v string) *UpdateDeviceStartupPictureRequest {
	s.Picture = &v
	return s
}

func (s *UpdateDeviceStartupPictureRequest) SetSerialNumber(v string) *UpdateDeviceStartupPictureRequest {
	s.SerialNumber = &v
	return s
}

type UpdateDeviceStartupPictureResponseBody struct {
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDeviceStartupPictureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceStartupPictureResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeviceStartupPictureResponseBody) SetErrorCode(v int32) *UpdateDeviceStartupPictureResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDeviceStartupPictureResponseBody) SetMessage(v string) *UpdateDeviceStartupPictureResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDeviceStartupPictureResponseBody) SetRequestId(v string) *UpdateDeviceStartupPictureResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDeviceStartupPictureResponseBody) SetSuccess(v bool) *UpdateDeviceStartupPictureResponseBody {
	s.Success = &v
	return s
}

type UpdateDeviceStartupPictureResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDeviceStartupPictureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDeviceStartupPictureResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceStartupPictureResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeviceStartupPictureResponse) SetHeaders(v map[string]*string) *UpdateDeviceStartupPictureResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeviceStartupPictureResponse) SetStatusCode(v int32) *UpdateDeviceStartupPictureResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDeviceStartupPictureResponse) SetBody(v *UpdateDeviceStartupPictureResponseBody) *UpdateDeviceStartupPictureResponse {
	s.Body = v
	return s
}

type UpdateGonggeLayoutRequest struct {
	MeetingUUID *string `json:"MeetingUUID,omitempty" xml:"MeetingUUID,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
	VideoCount  *string `json:"VideoCount,omitempty" xml:"VideoCount,omitempty"`
}

func (s UpdateGonggeLayoutRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGonggeLayoutRequest) GoString() string {
	return s.String()
}

func (s *UpdateGonggeLayoutRequest) SetMeetingUUID(v string) *UpdateGonggeLayoutRequest {
	s.MeetingUUID = &v
	return s
}

func (s *UpdateGonggeLayoutRequest) SetValue(v string) *UpdateGonggeLayoutRequest {
	s.Value = &v
	return s
}

func (s *UpdateGonggeLayoutRequest) SetVideoCount(v string) *UpdateGonggeLayoutRequest {
	s.VideoCount = &v
	return s
}

type UpdateGonggeLayoutResponseBody struct {
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateGonggeLayoutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGonggeLayoutResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGonggeLayoutResponseBody) SetErrorCode(v int32) *UpdateGonggeLayoutResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateGonggeLayoutResponseBody) SetMessage(v string) *UpdateGonggeLayoutResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGonggeLayoutResponseBody) SetRequestId(v string) *UpdateGonggeLayoutResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGonggeLayoutResponseBody) SetSuccess(v bool) *UpdateGonggeLayoutResponseBody {
	s.Success = &v
	return s
}

type UpdateGonggeLayoutResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateGonggeLayoutResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGonggeLayoutResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGonggeLayoutResponse) GoString() string {
	return s.String()
}

func (s *UpdateGonggeLayoutResponse) SetHeaders(v map[string]*string) *UpdateGonggeLayoutResponse {
	s.Headers = v
	return s
}

func (s *UpdateGonggeLayoutResponse) SetStatusCode(v int32) *UpdateGonggeLayoutResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGonggeLayoutResponse) SetBody(v *UpdateGonggeLayoutResponseBody) *UpdateGonggeLayoutResponse {
	s.Body = v
	return s
}

type UpdateLivePasswordRequest struct {
	LiveUUID         *string `json:"LiveUUID,omitempty" xml:"LiveUUID,omitempty"`
	OpenPasswordFlag *bool   `json:"OpenPasswordFlag,omitempty" xml:"OpenPasswordFlag,omitempty"`
	Password         *string `json:"Password,omitempty" xml:"Password,omitempty"`
	UserId           *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateLivePasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLivePasswordRequest) GoString() string {
	return s.String()
}

func (s *UpdateLivePasswordRequest) SetLiveUUID(v string) *UpdateLivePasswordRequest {
	s.LiveUUID = &v
	return s
}

func (s *UpdateLivePasswordRequest) SetOpenPasswordFlag(v bool) *UpdateLivePasswordRequest {
	s.OpenPasswordFlag = &v
	return s
}

func (s *UpdateLivePasswordRequest) SetPassword(v string) *UpdateLivePasswordRequest {
	s.Password = &v
	return s
}

func (s *UpdateLivePasswordRequest) SetUserId(v string) *UpdateLivePasswordRequest {
	s.UserId = &v
	return s
}

type UpdateLivePasswordResponseBody struct {
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateLivePasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLivePasswordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLivePasswordResponseBody) SetErrorCode(v int32) *UpdateLivePasswordResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateLivePasswordResponseBody) SetMessage(v string) *UpdateLivePasswordResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateLivePasswordResponseBody) SetRequestId(v string) *UpdateLivePasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLivePasswordResponseBody) SetSuccess(v bool) *UpdateLivePasswordResponseBody {
	s.Success = &v
	return s
}

type UpdateLivePasswordResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateLivePasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLivePasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLivePasswordResponse) GoString() string {
	return s.String()
}

func (s *UpdateLivePasswordResponse) SetHeaders(v map[string]*string) *UpdateLivePasswordResponse {
	s.Headers = v
	return s
}

func (s *UpdateLivePasswordResponse) SetStatusCode(v int32) *UpdateLivePasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLivePasswordResponse) SetBody(v *UpdateLivePasswordResponseBody) *UpdateLivePasswordResponse {
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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("aliyuncvc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ActiveDeviceWithOptions(request *ActiveDeviceRequest, runtime *util.RuntimeOptions) (_result *ActiveDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActiveCode)) {
		body["ActiveCode"] = request.ActiveCode
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceVersion)) {
		body["DeviceVersion"] = request.DeviceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.IP)) {
		body["IP"] = request.IP
	}

	if !tea.BoolValue(util.IsUnset(request.Mac)) {
		body["Mac"] = request.Mac
	}

	if !tea.BoolValue(util.IsUnset(request.SN)) {
		body["SN"] = request.SN
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ActiveDevice"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ActiveDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ActiveDevice(request *ActiveDeviceRequest) (_result *ActiveDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ActiveDeviceResponse{}
	_body, _err := client.ActiveDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ActiveMeetingWithOptions(request *ActiveMeetingRequest, runtime *util.RuntimeOptions) (_result *ActiveMeetingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MeetingCode)) {
		query["MeetingCode"] = request.MeetingCode
	}

	if !tea.BoolValue(util.IsUnset(request.MeetingUUID)) {
		query["MeetingUUID"] = request.MeetingUUID
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ActiveMeeting"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ActiveMeetingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ActiveMeeting(request *ActiveMeetingRequest) (_result *ActiveMeetingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ActiveMeetingResponse{}
	_body, _err := client.ActiveMeetingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchCreateDeviceWithOptions(request *BatchCreateDeviceRequest, runtime *util.RuntimeOptions) (_result *BatchCreateDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SN)) {
		body["SN"] = request.SN
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchCreateDevice"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchCreateDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchCreateDevice(request *BatchCreateDeviceRequest) (_result *BatchCreateDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchCreateDeviceResponse{}
	_body, _err := client.BatchCreateDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchJoinMeetingWithOptions(request *BatchJoinMeetingRequest, runtime *util.RuntimeOptions) (_result *BatchJoinMeetingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MeetingCode)) {
		body["MeetingCode"] = request.MeetingCode
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.RtcEngine)) {
		body["RtcEngine"] = request.RtcEngine
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchJoinMeeting"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchJoinMeetingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchJoinMeeting(request *BatchJoinMeetingRequest) (_result *BatchJoinMeetingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchJoinMeetingResponse{}
	_body, _err := client.BatchJoinMeetingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchJoinMeetingInternationalWithOptions(request *BatchJoinMeetingInternationalRequest, runtime *util.RuntimeOptions) (_result *BatchJoinMeetingInternationalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MeetingCode)) {
		body["MeetingCode"] = request.MeetingCode
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchJoinMeetingInternational"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchJoinMeetingInternationalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchJoinMeetingInternational(request *BatchJoinMeetingInternationalRequest) (_result *BatchJoinMeetingInternationalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchJoinMeetingInternationalResponse{}
	_body, _err := client.BatchJoinMeetingInternationalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CallDeviceWithOptions(request *CallDeviceRequest, runtime *util.RuntimeOptions) (_result *CallDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JoinMeetingFlag)) {
		query["JoinMeetingFlag"] = request.JoinMeetingFlag
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InviteName)) {
		body["InviteName"] = request.InviteName
	}

	if !tea.BoolValue(util.IsUnset(request.MeetingCode)) {
		body["MeetingCode"] = request.MeetingCode
	}

	if !tea.BoolValue(util.IsUnset(request.OperateUserId)) {
		body["OperateUserId"] = request.OperateUserId
	}

	if !tea.BoolValue(util.IsUnset(request.SN)) {
		body["SN"] = request.SN
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CallDevice"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CallDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CallDevice(request *CallDeviceRequest) (_result *CallDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CallDeviceResponse{}
	_body, _err := client.CallDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConferenceToLiveWithOptions(request *ConferenceToLiveRequest, runtime *util.RuntimeOptions) (_result *ConferenceToLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveName)) {
		body["LiveName"] = request.LiveName
	}

	if !tea.BoolValue(util.IsUnset(request.MeetingUUID)) {
		body["MeetingUUID"] = request.MeetingUUID
	}

	if !tea.BoolValue(util.IsUnset(request.OpenPasswordFlag)) {
		body["OpenPasswordFlag"] = request.OpenPasswordFlag
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ConferenceToLive"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConferenceToLiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConferenceToLive(request *ConferenceToLiveRequest) (_result *ConferenceToLiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConferenceToLiveResponse{}
	_body, _err := client.ConferenceToLiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDeviceMeetingWithOptions(request *CreateDeviceMeetingRequest, runtime *util.RuntimeOptions) (_result *CreateDeviceMeetingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MeetingName)) {
		body["MeetingName"] = request.MeetingName
	}

	if !tea.BoolValue(util.IsUnset(request.OpenPasswordtag)) {
		body["OpenPasswordtag"] = request.OpenPasswordtag
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.SN)) {
		body["SN"] = request.SN
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		body["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDeviceMeeting"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDeviceMeetingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDeviceMeeting(request *CreateDeviceMeetingRequest) (_result *CreateDeviceMeetingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDeviceMeetingResponse{}
	_body, _err := client.CreateDeviceMeetingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEvaluationWithOptions(request *CreateEvaluationRequest, runtime *util.RuntimeOptions) (_result *CreateEvaluationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateTime)) {
		query["CreateTime"] = request.CreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Evaluation)) {
		query["Evaluation"] = request.Evaluation
	}

	if !tea.BoolValue(util.IsUnset(request.MeetingUUID)) {
		query["MeetingUUID"] = request.MeetingUUID
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUUID)) {
		query["MemberUUID"] = request.MemberUUID
	}

	if !tea.BoolValue(util.IsUnset(request.Memo)) {
		query["Memo"] = request.Memo
	}

	if !tea.BoolValue(util.IsUnset(request.Score)) {
		query["Score"] = request.Score
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEvaluation"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEvaluationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEvaluation(request *CreateEvaluationRequest) (_result *CreateEvaluationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEvaluationResponse{}
	_body, _err := client.CreateEvaluationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLiveWithOptions(request *CreateLiveRequest, runtime *util.RuntimeOptions) (_result *CreateLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveName)) {
		body["LiveName"] = request.LiveName
	}

	if !tea.BoolValue(util.IsUnset(request.OpenPasswordFlag)) {
		body["OpenPasswordFlag"] = request.OpenPasswordFlag
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLive"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLive(request *CreateLiveRequest) (_result *CreateLiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLiveResponse{}
	_body, _err := client.CreateLiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMeetingWithOptions(request *CreateMeetingRequest, runtime *util.RuntimeOptions) (_result *CreateMeetingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TenantCode)) {
		query["TenantCode"] = request.TenantCode
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MasterEnableFlag)) {
		body["MasterEnableFlag"] = request.MasterEnableFlag
	}

	if !tea.BoolValue(util.IsUnset(request.MeetingMode)) {
		body["MeetingMode"] = request.MeetingMode
	}

	if !tea.BoolValue(util.IsUnset(request.MeetingName)) {
		body["MeetingName"] = request.MeetingName
	}

	if !tea.BoolValue(util.IsUnset(request.OpenPasswordFlag)) {
		body["OpenPasswordFlag"] = request.OpenPasswordFlag
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMeeting"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMeetingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMeeting(request *CreateMeetingRequest) (_result *CreateMeetingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMeetingResponse{}
	_body, _err := client.CreateMeetingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMeetingInternationalWithOptions(request *CreateMeetingInternationalRequest, runtime *util.RuntimeOptions) (_result *CreateMeetingInternationalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MeetingName)) {
		body["MeetingName"] = request.MeetingName
	}

	if !tea.BoolValue(util.IsUnset(request.OpenPasswordFlag)) {
		body["OpenPasswordFlag"] = request.OpenPasswordFlag
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMeetingInternational"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMeetingInternationalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMeetingInternational(request *CreateMeetingInternationalRequest) (_result *CreateMeetingInternationalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMeetingInternationalResponse{}
	_body, _err := client.CreateMeetingInternationalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUserWithOptions(request *CreateUserRequest, runtime *util.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TenantCode)) {
		query["TenantCode"] = request.TenantCode
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Count)) {
		body["Count"] = request.Count
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfo)) {
		body["UserInfo"] = request.UserInfo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUser"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUser(request *CreateUserRequest) (_result *CreateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserResponse{}
	_body, _err := client.CreateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUserInternationalWithOptions(request *CreateUserInternationalRequest, runtime *util.RuntimeOptions) (_result *CreateUserInternationalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Count)) {
		body["Count"] = request.Count
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfo)) {
		body["UserInfo"] = request.UserInfo
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUserInternational"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserInternationalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUserInternational(request *CreateUserInternationalRequest) (_result *CreateUserInternationalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserInternationalResponse{}
	_body, _err := client.CreateUserInternationalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomGonggeLayoutWithOptions(request *CustomGonggeLayoutRequest, runtime *util.RuntimeOptions) (_result *CustomGonggeLayoutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LayoutSolution)) {
		body["LayoutSolution"] = request.LayoutSolution
	}

	if !tea.BoolValue(util.IsUnset(request.MeetingUUID)) {
		body["MeetingUUID"] = request.MeetingUUID
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CustomGonggeLayout"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CustomGonggeLayoutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomGonggeLayout(request *CustomGonggeLayoutRequest) (_result *CustomGonggeLayoutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CustomGonggeLayoutResponse{}
	_body, _err := client.CustomGonggeLayoutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomLayoutWithOptions(request *CustomLayoutRequest, runtime *util.RuntimeOptions) (_result *CustomLayoutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LayoutInfo)) {
		body["LayoutInfo"] = request.LayoutInfo
	}

	if !tea.BoolValue(util.IsUnset(request.LiveUUID)) {
		body["LiveUUID"] = request.LiveUUID
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CustomLayout"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CustomLayoutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomLayout(request *CustomLayoutRequest) (_result *CustomLayoutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CustomLayoutResponse{}
	_body, _err := client.CustomLayoutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDeviceWithOptions(request *DeleteDeviceRequest, runtime *util.RuntimeOptions) (_result *DeleteDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SN)) {
		body["SN"] = request.SN
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDevice"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDevice(request *DeleteDeviceRequest) (_result *DeleteDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDeviceResponse{}
	_body, _err := client.DeleteDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLiveWithOptions(request *DeleteLiveRequest, runtime *util.RuntimeOptions) (_result *DeleteLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveUUID)) {
		body["LiveUUID"] = request.LiveUUID
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLive"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLive(request *DeleteLiveRequest) (_result *DeleteLiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLiveResponse{}
	_body, _err := client.DeleteLiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMeetingWithOptions(request *DeleteMeetingRequest, runtime *util.RuntimeOptions) (_result *DeleteMeetingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MeetingUUID)) {
		query["MeetingUUID"] = request.MeetingUUID
	}

	if !tea.BoolValue(util.IsUnset(request.TenantCode)) {
		query["TenantCode"] = request.TenantCode
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMeeting"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMeetingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMeeting(request *DeleteMeetingRequest) (_result *DeleteMeetingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMeetingResponse{}
	_body, _err := client.DeleteMeetingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserWithOptions(request *DeleteUserRequest, runtime *util.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Count)) {
		body["Count"] = request.Count
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfo)) {
		body["UserInfo"] = request.UserInfo
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUser"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUser(request *DeleteUserRequest) (_result *DeleteUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserResponse{}
	_body, _err := client.DeleteUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableLiveSpeakerWithOptions(request *EnableLiveSpeakerRequest, runtime *util.RuntimeOptions) (_result *EnableLiveSpeakerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnableSpeakerFlag)) {
		body["EnableSpeakerFlag"] = request.EnableSpeakerFlag
	}

	if !tea.BoolValue(util.IsUnset(request.LiveUUID)) {
		body["LiveUUID"] = request.LiveUUID
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableLiveSpeaker"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableLiveSpeakerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableLiveSpeaker(request *EnableLiveSpeakerRequest) (_result *EnableLiveSpeakerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableLiveSpeakerResponse{}
	_body, _err := client.EnableLiveSpeakerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EndDeviceMeetingWithOptions(request *EndDeviceMeetingRequest, runtime *util.RuntimeOptions) (_result *EndDeviceMeetingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MeetingUUID)) {
		body["MeetingUUID"] = request.MeetingUUID
	}

	if !tea.BoolValue(util.IsUnset(request.SN)) {
		body["SN"] = request.SN
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EndDeviceMeeting"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EndDeviceMeetingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EndDeviceMeeting(request *EndDeviceMeetingRequest) (_result *EndDeviceMeetingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EndDeviceMeetingResponse{}
	_body, _err := client.EndDeviceMeetingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EndLiveWithOptions(request *EndLiveRequest, runtime *util.RuntimeOptions) (_result *EndLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveUUID)) {
		body["LiveUUID"] = request.LiveUUID
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EndLive"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EndLiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EndLive(request *EndLiveRequest) (_result *EndLiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EndLiveResponse{}
	_body, _err := client.EndLiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAccountInfoWithOptions(runtime *util.RuntimeOptions) (_result *GetAccountInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetAccountInfo"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccountInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAccountInfo() (_result *GetAccountInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccountInfoResponse{}
	_body, _err := client.GetAccountInfoWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConferenceConcurrencyStatisticsWithOptions(runtime *util.RuntimeOptions) (_result *GetConferenceConcurrencyStatisticsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetConferenceConcurrencyStatistics"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConferenceConcurrencyStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConferenceConcurrencyStatistics() (_result *GetConferenceConcurrencyStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConferenceConcurrencyStatisticsResponse{}
	_body, _err := client.GetConferenceConcurrencyStatisticsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeviceActiveCodeWithOptions(request *GetDeviceActiveCodeRequest, runtime *util.RuntimeOptions) (_result *GetDeviceActiveCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SN)) {
		body["SN"] = request.SN
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeviceActiveCode"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeviceActiveCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeviceActiveCode(request *GetDeviceActiveCodeRequest) (_result *GetDeviceActiveCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeviceActiveCodeResponse{}
	_body, _err := client.GetDeviceActiveCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeviceInfoWithOptions(request *GetDeviceInfoRequest, runtime *util.RuntimeOptions) (_result *GetDeviceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CastScreenCode)) {
		query["CastScreenCode"] = request.CastScreenCode
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		query["Sn"] = request.Sn
	}

	if !tea.BoolValue(util.IsUnset(request.TenantCode)) {
		query["TenantCode"] = request.TenantCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeviceInfo"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeviceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeviceInfo(request *GetDeviceInfoRequest) (_result *GetDeviceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeviceInfoResponse{}
	_body, _err := client.GetDeviceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeviceListWithOptions(request *GetDeviceListRequest, runtime *util.RuntimeOptions) (_result *GetDeviceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CastScreenCode)) {
		query["CastScreenCode"] = request.CastScreenCode
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeviceList"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeviceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeviceList(request *GetDeviceListRequest) (_result *GetDeviceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeviceListResponse{}
	_body, _err := client.GetDeviceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeviceTokenWithOptions(request *GetDeviceTokenRequest, runtime *util.RuntimeOptions) (_result *GetDeviceTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SN)) {
		query["SN"] = request.SN
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeviceToken"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeviceTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeviceToken(request *GetDeviceTokenRequest) (_result *GetDeviceTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeviceTokenResponse{}
	_body, _err := client.GetDeviceTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGrtnTokenWithOptions(runtime *util.RuntimeOptions) (_result *GetGrtnTokenResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetGrtnToken"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGrtnTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGrtnToken() (_result *GetGrtnTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGrtnTokenResponse{}
	_body, _err := client.GetGrtnTokenWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMeetingWithOptions(request *GetMeetingRequest, runtime *util.RuntimeOptions) (_result *GetMeetingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TenantCode)) {
		query["TenantCode"] = request.TenantCode
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MeetingUUID)) {
		body["MeetingUUID"] = request.MeetingUUID
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMeeting"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMeetingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMeeting(request *GetMeetingRequest) (_result *GetMeetingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMeetingResponse{}
	_body, _err := client.GetMeetingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMeetingConcurrencyWithOptions(request *GetMeetingConcurrencyRequest, runtime *util.RuntimeOptions) (_result *GetMeetingConcurrencyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MeetingUUID)) {
		query["MeetingUUID"] = request.MeetingUUID
	}

	if !tea.BoolValue(util.IsUnset(request.TenantCode)) {
		query["TenantCode"] = request.TenantCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMeetingConcurrency"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMeetingConcurrencyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMeetingConcurrency(request *GetMeetingConcurrencyRequest) (_result *GetMeetingConcurrencyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMeetingConcurrencyResponse{}
	_body, _err := client.GetMeetingConcurrencyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMeetingInternationalWithOptions(request *GetMeetingInternationalRequest, runtime *util.RuntimeOptions) (_result *GetMeetingInternationalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MeetingUUID)) {
		body["MeetingUUID"] = request.MeetingUUID
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMeetingInternational"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMeetingInternationalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMeetingInternational(request *GetMeetingInternationalRequest) (_result *GetMeetingInternationalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMeetingInternationalResponse{}
	_body, _err := client.GetMeetingInternationalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMeetingMemberWithOptions(request *GetMeetingMemberRequest, runtime *util.RuntimeOptions) (_result *GetMeetingMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MeetingUUID)) {
		query["MeetingUUID"] = request.MeetingUUID
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMeetingMember"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMeetingMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMeetingMember(request *GetMeetingMemberRequest) (_result *GetMeetingMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMeetingMemberResponse{}
	_body, _err := client.GetMeetingMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMemberStatusWithOptions(request *GetMemberStatusRequest, runtime *util.RuntimeOptions) (_result *GetMemberStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MeetingUUID)) {
		query["MeetingUUID"] = request.MeetingUUID
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMemberStatus"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMemberStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMemberStatus(request *GetMemberStatusRequest) (_result *GetMemberStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMemberStatusResponse{}
	_body, _err := client.GetMemberStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetScreenVerificationCodeWithOptions(request *GetScreenVerificationCodeRequest, runtime *util.RuntimeOptions) (_result *GetScreenVerificationCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CastScreenCode)) {
		body["CastScreenCode"] = request.CastScreenCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetScreenVerificationCode"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetScreenVerificationCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetScreenVerificationCode(request *GetScreenVerificationCodeRequest) (_result *GetScreenVerificationCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetScreenVerificationCodeResponse{}
	_body, _err := client.GetScreenVerificationCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStatisticWithOptions(request *GetStatisticRequest, runtime *util.RuntimeOptions) (_result *GetStatisticResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStatistic"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStatistic(request *GetStatisticRequest) (_result *GetStatisticResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStatisticResponse{}
	_body, _err := client.GetStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserWithOptions(request *GetUserRequest, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUser"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUser(request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWebSocketTokenWithOptions(request *GetWebSocketTokenRequest, runtime *util.RuntimeOptions) (_result *GetWebSocketTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OldToken)) {
		query["OldToken"] = request.OldToken
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWebSocketToken"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWebSocketTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWebSocketToken(request *GetWebSocketTokenRequest) (_result *GetWebSocketTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWebSocketTokenResponse{}
	_body, _err := client.GetWebSocketTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InviteUserWithOptions(request *InviteUserRequest, runtime *util.RuntimeOptions) (_result *InviteUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MeetingUUID)) {
		body["MeetingUUID"] = request.MeetingUUID
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["UserIds"] = request.UserIds
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InviteUser"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InviteUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InviteUser(request *InviteUserRequest) (_result *InviteUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InviteUserResponse{}
	_body, _err := client.InviteUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) JoinDeviceMeetingWithOptions(request *JoinDeviceMeetingRequest, runtime *util.RuntimeOptions) (_result *JoinDeviceMeetingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MeetingCode)) {
		body["MeetingCode"] = request.MeetingCode
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.SN)) {
		body["SN"] = request.SN
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		body["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("JoinDeviceMeeting"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &JoinDeviceMeetingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) JoinDeviceMeeting(request *JoinDeviceMeetingRequest) (_result *JoinDeviceMeetingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &JoinDeviceMeetingResponse{}
	_body, _err := client.JoinDeviceMeetingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) JoinLiveWithOptions(request *JoinLiveRequest, runtime *util.RuntimeOptions) (_result *JoinLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveUUID)) {
		body["LiveUUID"] = request.LiveUUID
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("JoinLive"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &JoinLiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) JoinLive(request *JoinLiveRequest) (_result *JoinLiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &JoinLiveResponse{}
	_body, _err := client.JoinLiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) JoinMeetingWithOptions(request *JoinMeetingRequest, runtime *util.RuntimeOptions) (_result *JoinMeetingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TenantCode)) {
		query["TenantCode"] = request.TenantCode
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MeetingCode)) {
		body["MeetingCode"] = request.MeetingCode
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.RtcEngine)) {
		body["RtcEngine"] = request.RtcEngine
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("JoinMeeting"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &JoinMeetingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) JoinMeeting(request *JoinMeetingRequest) (_result *JoinMeetingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &JoinMeetingResponse{}
	_body, _err := client.JoinMeetingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) JoinMeetingInternationalWithOptions(request *JoinMeetingInternationalRequest, runtime *util.RuntimeOptions) (_result *JoinMeetingInternationalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MeetingCode)) {
		body["MeetingCode"] = request.MeetingCode
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("JoinMeetingInternational"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &JoinMeetingInternationalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) JoinMeetingInternational(request *JoinMeetingInternationalRequest) (_result *JoinMeetingInternationalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &JoinMeetingInternationalResponse{}
	_body, _err := client.JoinMeetingInternationalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConferenceDevicesWithOptions(request *ListConferenceDevicesRequest, runtime *util.RuntimeOptions) (_result *ListConferenceDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		query["SerialNumber"] = request.SerialNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListConferenceDevices"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConferenceDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConferenceDevices(request *ListConferenceDevicesRequest) (_result *ListConferenceDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConferenceDevicesResponse{}
	_body, _err := client.ListConferenceDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeviceIpWithOptions(request *ListDeviceIpRequest, runtime *util.RuntimeOptions) (_result *ListDeviceIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SN)) {
		body["SN"] = request.SN
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeviceIp"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeviceIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeviceIp(request *ListDeviceIpRequest) (_result *ListDeviceIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeviceIpResponse{}
	_body, _err := client.ListDeviceIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDevicesWithOptions(request *ListDevicesRequest, runtime *util.RuntimeOptions) (_result *ListDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CastScreenCode)) {
		query["CastScreenCode"] = request.CastScreenCode
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SN)) {
		query["SN"] = request.SN
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDevices"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDevices(request *ListDevicesRequest) (_result *ListDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDevicesResponse{}
	_body, _err := client.ListDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEvaluationsWithOptions(runtime *util.RuntimeOptions) (_result *ListEvaluationsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListEvaluations"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEvaluationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEvaluations() (_result *ListEvaluationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEvaluationsResponse{}
	_body, _err := client.ListEvaluationsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIsvStatisticsWithOptions(request *ListIsvStatisticsRequest, runtime *util.RuntimeOptions) (_result *ListIsvStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIsvStatistics"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIsvStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIsvStatistics(request *ListIsvStatisticsRequest) (_result *ListIsvStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIsvStatisticsResponse{}
	_body, _err := client.ListIsvStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMembersWithOptions(request *ListMembersRequest, runtime *util.RuntimeOptions) (_result *ListMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MeetingUUID)) {
		query["MeetingUUID"] = request.MeetingUUID
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMembers"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMembers(request *ListMembersRequest) (_result *ListMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMembersResponse{}
	_body, _err := client.ListMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsers"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDeviceBackgroundWithOptions(request *ModifyDeviceBackgroundRequest, runtime *util.RuntimeOptions) (_result *ModifyDeviceBackgroundResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Picture)) {
		body["Picture"] = request.Picture
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		body["SerialNumber"] = request.SerialNumber
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDeviceBackground"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDeviceBackgroundResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDeviceBackground(request *ModifyDeviceBackgroundRequest) (_result *ModifyDeviceBackgroundResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDeviceBackgroundResponse{}
	_body, _err := client.ModifyDeviceBackgroundWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyMeetingPasswordWithOptions(request *ModifyMeetingPasswordRequest, runtime *util.RuntimeOptions) (_result *ModifyMeetingPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MeetingUUID)) {
		body["MeetingUUID"] = request.MeetingUUID
	}

	if !tea.BoolValue(util.IsUnset(request.OpenPasswordFlag)) {
		body["OpenPasswordFlag"] = request.OpenPasswordFlag
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyMeetingPassword"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyMeetingPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyMeetingPassword(request *ModifyMeetingPasswordRequest) (_result *ModifyMeetingPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyMeetingPasswordResponse{}
	_body, _err := client.ModifyMeetingPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyMeetingPasswordInternationalWithOptions(request *ModifyMeetingPasswordInternationalRequest, runtime *util.RuntimeOptions) (_result *ModifyMeetingPasswordInternationalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MeetingUUID)) {
		body["MeetingUUID"] = request.MeetingUUID
	}

	if !tea.BoolValue(util.IsUnset(request.OpenPasswordFlag)) {
		body["OpenPasswordFlag"] = request.OpenPasswordFlag
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyMeetingPasswordInternational"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyMeetingPasswordInternationalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyMeetingPasswordInternational(request *ModifyMeetingPasswordInternationalRequest) (_result *ModifyMeetingPasswordInternationalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyMeetingPasswordInternationalResponse{}
	_body, _err := client.ModifyMeetingPasswordInternationalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMeetingMemberActionWithOptions(request *QueryMeetingMemberActionRequest, runtime *util.RuntimeOptions) (_result *QueryMeetingMemberActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MeetingUUID)) {
		body["MeetingUUID"] = request.MeetingUUID
	}

	if !tea.BoolValue(util.IsUnset(request.MeetingUnitKey)) {
		body["MeetingUnitKey"] = request.MeetingUnitKey
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUUID)) {
		body["MemberUUID"] = request.MemberUUID
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMeetingMemberAction"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMeetingMemberActionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMeetingMemberAction(request *QueryMeetingMemberActionRequest) (_result *QueryMeetingMemberActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMeetingMemberActionResponse{}
	_body, _err := client.QueryMeetingMemberActionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshDeviceScreenCodeWithOptions(request *RefreshDeviceScreenCodeRequest, runtime *util.RuntimeOptions) (_result *RefreshDeviceScreenCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		body["SerialNumber"] = request.SerialNumber
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshDeviceScreenCode"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshDeviceScreenCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshDeviceScreenCode(request *RefreshDeviceScreenCodeRequest) (_result *RefreshDeviceScreenCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshDeviceScreenCodeResponse{}
	_body, _err := client.RefreshDeviceScreenCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterDeviceWithOptions(request *RegisterDeviceRequest, runtime *util.RuntimeOptions) (_result *RegisterDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceVersion)) {
		query["DeviceVersion"] = request.DeviceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.IP)) {
		query["IP"] = request.IP
	}

	if !tea.BoolValue(util.IsUnset(request.Mac)) {
		query["Mac"] = request.Mac
	}

	if !tea.BoolValue(util.IsUnset(request.SN)) {
		query["SN"] = request.SN
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RegisterDevice"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RegisterDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterDevice(request *RegisterDeviceRequest) (_result *RegisterDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterDeviceResponse{}
	_body, _err := client.RegisterDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterUemDeviceWithOptions(request *RegisterUemDeviceRequest, runtime *util.RuntimeOptions) (_result *RegisterUemDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceVersion)) {
		query["DeviceVersion"] = request.DeviceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.IP)) {
		query["IP"] = request.IP
	}

	if !tea.BoolValue(util.IsUnset(request.Mac)) {
		query["Mac"] = request.Mac
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RegisterUemDevice"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RegisterUemDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterUemDevice(request *RegisterUemDeviceRequest) (_result *RegisterUemDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterUemDeviceResponse{}
	_body, _err := client.RegisterUemDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendMeetingCommandWithOptions(request *SendMeetingCommandRequest, runtime *util.RuntimeOptions) (_result *SendMeetingCommandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MeetingUUID)) {
		query["MeetingUUID"] = request.MeetingUUID
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUUID)) {
		query["MemberUUID"] = request.MemberUUID
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Command)) {
		body["Command"] = request.Command
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorMemberUUID)) {
		body["OperatorMemberUUID"] = request.OperatorMemberUUID
	}

	if !tea.BoolValue(util.IsUnset(request.SendType)) {
		body["SendType"] = request.SendType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendMeetingCommand"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendMeetingCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendMeetingCommand(request *SendMeetingCommandRequest) (_result *SendMeetingCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendMeetingCommandResponse{}
	_body, _err := client.SendMeetingCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendScreenStartWithOptions(request *SendScreenStartRequest, runtime *util.RuntimeOptions) (_result *SendScreenStartResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CastScreenCode)) {
		body["CastScreenCode"] = request.CastScreenCode
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendScreenStart"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendScreenStartResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendScreenStart(request *SendScreenStartRequest) (_result *SendScreenStartResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendScreenStartResponse{}
	_body, _err := client.SendScreenStartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartLiveWithOptions(request *StartLiveRequest, runtime *util.RuntimeOptions) (_result *StartLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LayoutInfo)) {
		body["LayoutInfo"] = request.LayoutInfo
	}

	if !tea.BoolValue(util.IsUnset(request.LiveUUID)) {
		body["LiveUUID"] = request.LiveUUID
	}

	if !tea.BoolValue(util.IsUnset(request.PushInfo)) {
		body["PushInfo"] = request.PushInfo
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartLive"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartLiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartLive(request *StartLiveRequest) (_result *StartLiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartLiveResponse{}
	_body, _err := client.StartLiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDeviceHeartBeatWithOptions(request *UpdateDeviceHeartBeatRequest, runtime *util.RuntimeOptions) (_result *UpdateDeviceHeartBeatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Message)) {
		query["Message"] = request.Message
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDeviceHeartBeat"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDeviceHeartBeatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDeviceHeartBeat(request *UpdateDeviceHeartBeatRequest) (_result *UpdateDeviceHeartBeatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDeviceHeartBeatResponse{}
	_body, _err := client.UpdateDeviceHeartBeatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDeviceStartupPictureWithOptions(request *UpdateDeviceStartupPictureRequest, runtime *util.RuntimeOptions) (_result *UpdateDeviceStartupPictureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Picture)) {
		body["Picture"] = request.Picture
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		body["SerialNumber"] = request.SerialNumber
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDeviceStartupPicture"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDeviceStartupPictureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDeviceStartupPicture(request *UpdateDeviceStartupPictureRequest) (_result *UpdateDeviceStartupPictureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDeviceStartupPictureResponse{}
	_body, _err := client.UpdateDeviceStartupPictureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGonggeLayoutWithOptions(request *UpdateGonggeLayoutRequest, runtime *util.RuntimeOptions) (_result *UpdateGonggeLayoutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MeetingUUID)) {
		body["MeetingUUID"] = request.MeetingUUID
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		body["Value"] = request.Value
	}

	if !tea.BoolValue(util.IsUnset(request.VideoCount)) {
		body["VideoCount"] = request.VideoCount
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGonggeLayout"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGonggeLayoutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGonggeLayout(request *UpdateGonggeLayoutRequest) (_result *UpdateGonggeLayoutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGonggeLayoutResponse{}
	_body, _err := client.UpdateGonggeLayoutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLivePasswordWithOptions(request *UpdateLivePasswordRequest, runtime *util.RuntimeOptions) (_result *UpdateLivePasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveUUID)) {
		body["LiveUUID"] = request.LiveUUID
	}

	if !tea.BoolValue(util.IsUnset(request.OpenPasswordFlag)) {
		body["OpenPasswordFlag"] = request.OpenPasswordFlag
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLivePassword"),
		Version:     tea.String("2019-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLivePasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLivePassword(request *UpdateLivePasswordRequest) (_result *UpdateLivePasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLivePasswordResponse{}
	_body, _err := client.UpdateLivePasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
