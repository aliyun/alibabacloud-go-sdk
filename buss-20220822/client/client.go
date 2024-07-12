// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BusinessResultServiceRequest struct {
	ActionCode    *string                `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	BussinessCode *string                `json:"BussinessCode,omitempty" xml:"BussinessCode,omitempty"`
	ErrCode       *string                `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage    *string                `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId     *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result        map[string]interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
	Success       *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BusinessResultServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s BusinessResultServiceRequest) GoString() string {
	return s.String()
}

func (s *BusinessResultServiceRequest) SetActionCode(v string) *BusinessResultServiceRequest {
	s.ActionCode = &v
	return s
}

func (s *BusinessResultServiceRequest) SetBussinessCode(v string) *BusinessResultServiceRequest {
	s.BussinessCode = &v
	return s
}

func (s *BusinessResultServiceRequest) SetErrCode(v string) *BusinessResultServiceRequest {
	s.ErrCode = &v
	return s
}

func (s *BusinessResultServiceRequest) SetErrMessage(v string) *BusinessResultServiceRequest {
	s.ErrMessage = &v
	return s
}

func (s *BusinessResultServiceRequest) SetRequestId(v string) *BusinessResultServiceRequest {
	s.RequestId = &v
	return s
}

func (s *BusinessResultServiceRequest) SetResult(v map[string]interface{}) *BusinessResultServiceRequest {
	s.Result = v
	return s
}

func (s *BusinessResultServiceRequest) SetSuccess(v bool) *BusinessResultServiceRequest {
	s.Success = &v
	return s
}

type BusinessResultServiceShrinkRequest struct {
	ActionCode    *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	BussinessCode *string `json:"BussinessCode,omitempty" xml:"BussinessCode,omitempty"`
	ErrCode       *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage    *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultShrink  *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Success       *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BusinessResultServiceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BusinessResultServiceShrinkRequest) GoString() string {
	return s.String()
}

func (s *BusinessResultServiceShrinkRequest) SetActionCode(v string) *BusinessResultServiceShrinkRequest {
	s.ActionCode = &v
	return s
}

func (s *BusinessResultServiceShrinkRequest) SetBussinessCode(v string) *BusinessResultServiceShrinkRequest {
	s.BussinessCode = &v
	return s
}

func (s *BusinessResultServiceShrinkRequest) SetErrCode(v string) *BusinessResultServiceShrinkRequest {
	s.ErrCode = &v
	return s
}

func (s *BusinessResultServiceShrinkRequest) SetErrMessage(v string) *BusinessResultServiceShrinkRequest {
	s.ErrMessage = &v
	return s
}

func (s *BusinessResultServiceShrinkRequest) SetRequestId(v string) *BusinessResultServiceShrinkRequest {
	s.RequestId = &v
	return s
}

func (s *BusinessResultServiceShrinkRequest) SetResultShrink(v string) *BusinessResultServiceShrinkRequest {
	s.ResultShrink = &v
	return s
}

func (s *BusinessResultServiceShrinkRequest) SetSuccess(v bool) *BusinessResultServiceShrinkRequest {
	s.Success = &v
	return s
}

type BusinessResultServiceResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Count   *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BusinessResultServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BusinessResultServiceResponseBody) GoString() string {
	return s.String()
}

func (s *BusinessResultServiceResponseBody) SetCode(v string) *BusinessResultServiceResponseBody {
	s.Code = &v
	return s
}

func (s *BusinessResultServiceResponseBody) SetCount(v int32) *BusinessResultServiceResponseBody {
	s.Count = &v
	return s
}

func (s *BusinessResultServiceResponseBody) SetData(v string) *BusinessResultServiceResponseBody {
	s.Data = &v
	return s
}

func (s *BusinessResultServiceResponseBody) SetMessage(v string) *BusinessResultServiceResponseBody {
	s.Message = &v
	return s
}

func (s *BusinessResultServiceResponseBody) SetSuccess(v bool) *BusinessResultServiceResponseBody {
	s.Success = &v
	return s
}

type BusinessResultServiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BusinessResultServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BusinessResultServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s BusinessResultServiceResponse) GoString() string {
	return s.String()
}

func (s *BusinessResultServiceResponse) SetHeaders(v map[string]*string) *BusinessResultServiceResponse {
	s.Headers = v
	return s
}

func (s *BusinessResultServiceResponse) SetStatusCode(v int32) *BusinessResultServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *BusinessResultServiceResponse) SetBody(v *BusinessResultServiceResponseBody) *BusinessResultServiceResponse {
	s.Body = v
	return s
}

type CreateUserInvestigationInfoQueryTaskRequest struct {
	// This parameter is required.
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// This parameter is required.
	EmployeeId *string `json:"employeeId,omitempty" xml:"employeeId,omitempty"`
	// This parameter is required.
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	OssFileName *string `json:"ossFileName,omitempty" xml:"ossFileName,omitempty"`
	// This parameter is required.
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateUserInvestigationInfoQueryTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserInvestigationInfoQueryTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateUserInvestigationInfoQueryTaskRequest) SetDataType(v string) *CreateUserInvestigationInfoQueryTaskRequest {
	s.DataType = &v
	return s
}

func (s *CreateUserInvestigationInfoQueryTaskRequest) SetEmployeeId(v string) *CreateUserInvestigationInfoQueryTaskRequest {
	s.EmployeeId = &v
	return s
}

func (s *CreateUserInvestigationInfoQueryTaskRequest) SetEndTime(v int64) *CreateUserInvestigationInfoQueryTaskRequest {
	s.EndTime = &v
	return s
}

func (s *CreateUserInvestigationInfoQueryTaskRequest) SetOssFileName(v string) *CreateUserInvestigationInfoQueryTaskRequest {
	s.OssFileName = &v
	return s
}

func (s *CreateUserInvestigationInfoQueryTaskRequest) SetStartTime(v int64) *CreateUserInvestigationInfoQueryTaskRequest {
	s.StartTime = &v
	return s
}

func (s *CreateUserInvestigationInfoQueryTaskRequest) SetUserId(v string) *CreateUserInvestigationInfoQueryTaskRequest {
	s.UserId = &v
	return s
}

type CreateUserInvestigationInfoQueryTaskResponseBody struct {
	Code    *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *CreateUserInvestigationInfoQueryTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateUserInvestigationInfoQueryTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserInvestigationInfoQueryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserInvestigationInfoQueryTaskResponseBody) SetCode(v string) *CreateUserInvestigationInfoQueryTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateUserInvestigationInfoQueryTaskResponseBody) SetData(v *CreateUserInvestigationInfoQueryTaskResponseBodyData) *CreateUserInvestigationInfoQueryTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateUserInvestigationInfoQueryTaskResponseBody) SetMessage(v string) *CreateUserInvestigationInfoQueryTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateUserInvestigationInfoQueryTaskResponseBody) SetSuccess(v bool) *CreateUserInvestigationInfoQueryTaskResponseBody {
	s.Success = &v
	return s
}

type CreateUserInvestigationInfoQueryTaskResponseBodyData struct {
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateUserInvestigationInfoQueryTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateUserInvestigationInfoQueryTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateUserInvestigationInfoQueryTaskResponseBodyData) SetTaskId(v string) *CreateUserInvestigationInfoQueryTaskResponseBodyData {
	s.TaskId = &v
	return s
}

type CreateUserInvestigationInfoQueryTaskResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserInvestigationInfoQueryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserInvestigationInfoQueryTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserInvestigationInfoQueryTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateUserInvestigationInfoQueryTaskResponse) SetHeaders(v map[string]*string) *CreateUserInvestigationInfoQueryTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateUserInvestigationInfoQueryTaskResponse) SetStatusCode(v int32) *CreateUserInvestigationInfoQueryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserInvestigationInfoQueryTaskResponse) SetBody(v *CreateUserInvestigationInfoQueryTaskResponseBody) *CreateUserInvestigationInfoQueryTaskResponse {
	s.Body = v
	return s
}

type FindInstanceInfoRequest struct {
	BussinessCode *string                `json:"bussinessCode,omitempty" xml:"bussinessCode,omitempty"`
	Domain        *string                `json:"domain,omitempty" xml:"domain,omitempty"`
	EndTime       *int64                 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Extras        map[string]interface{} `json:"extras,omitempty" xml:"extras,omitempty"`
	Ip            *string                `json:"ip,omitempty" xml:"ip,omitempty"`
	NeedDNS       *bool                  `json:"needDNS,omitempty" xml:"needDNS,omitempty"`
	StartTime     *int64                 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Url           *string                `json:"url,omitempty" xml:"url,omitempty"`
	UserId        *string                `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s FindInstanceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s FindInstanceInfoRequest) GoString() string {
	return s.String()
}

func (s *FindInstanceInfoRequest) SetBussinessCode(v string) *FindInstanceInfoRequest {
	s.BussinessCode = &v
	return s
}

func (s *FindInstanceInfoRequest) SetDomain(v string) *FindInstanceInfoRequest {
	s.Domain = &v
	return s
}

func (s *FindInstanceInfoRequest) SetEndTime(v int64) *FindInstanceInfoRequest {
	s.EndTime = &v
	return s
}

func (s *FindInstanceInfoRequest) SetExtras(v map[string]interface{}) *FindInstanceInfoRequest {
	s.Extras = v
	return s
}

func (s *FindInstanceInfoRequest) SetIp(v string) *FindInstanceInfoRequest {
	s.Ip = &v
	return s
}

func (s *FindInstanceInfoRequest) SetNeedDNS(v bool) *FindInstanceInfoRequest {
	s.NeedDNS = &v
	return s
}

func (s *FindInstanceInfoRequest) SetStartTime(v int64) *FindInstanceInfoRequest {
	s.StartTime = &v
	return s
}

func (s *FindInstanceInfoRequest) SetUrl(v string) *FindInstanceInfoRequest {
	s.Url = &v
	return s
}

func (s *FindInstanceInfoRequest) SetUserId(v string) *FindInstanceInfoRequest {
	s.UserId = &v
	return s
}

type FindInstanceInfoShrinkRequest struct {
	BussinessCode *string `json:"bussinessCode,omitempty" xml:"bussinessCode,omitempty"`
	Domain        *string `json:"domain,omitempty" xml:"domain,omitempty"`
	EndTime       *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ExtrasShrink  *string `json:"extras,omitempty" xml:"extras,omitempty"`
	Ip            *string `json:"ip,omitempty" xml:"ip,omitempty"`
	NeedDNS       *bool   `json:"needDNS,omitempty" xml:"needDNS,omitempty"`
	StartTime     *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Url           *string `json:"url,omitempty" xml:"url,omitempty"`
	UserId        *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s FindInstanceInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s FindInstanceInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *FindInstanceInfoShrinkRequest) SetBussinessCode(v string) *FindInstanceInfoShrinkRequest {
	s.BussinessCode = &v
	return s
}

func (s *FindInstanceInfoShrinkRequest) SetDomain(v string) *FindInstanceInfoShrinkRequest {
	s.Domain = &v
	return s
}

func (s *FindInstanceInfoShrinkRequest) SetEndTime(v int64) *FindInstanceInfoShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *FindInstanceInfoShrinkRequest) SetExtrasShrink(v string) *FindInstanceInfoShrinkRequest {
	s.ExtrasShrink = &v
	return s
}

func (s *FindInstanceInfoShrinkRequest) SetIp(v string) *FindInstanceInfoShrinkRequest {
	s.Ip = &v
	return s
}

func (s *FindInstanceInfoShrinkRequest) SetNeedDNS(v bool) *FindInstanceInfoShrinkRequest {
	s.NeedDNS = &v
	return s
}

func (s *FindInstanceInfoShrinkRequest) SetStartTime(v int64) *FindInstanceInfoShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *FindInstanceInfoShrinkRequest) SetUrl(v string) *FindInstanceInfoShrinkRequest {
	s.Url = &v
	return s
}

func (s *FindInstanceInfoShrinkRequest) SetUserId(v string) *FindInstanceInfoShrinkRequest {
	s.UserId = &v
	return s
}

type FindInstanceInfoResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Count     *int32                            `json:"Count,omitempty" xml:"Count,omitempty"`
	Data      *FindInstanceInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FindInstanceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindInstanceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *FindInstanceInfoResponseBody) SetCode(v string) *FindInstanceInfoResponseBody {
	s.Code = &v
	return s
}

func (s *FindInstanceInfoResponseBody) SetCount(v int32) *FindInstanceInfoResponseBody {
	s.Count = &v
	return s
}

func (s *FindInstanceInfoResponseBody) SetData(v *FindInstanceInfoResponseBodyData) *FindInstanceInfoResponseBody {
	s.Data = v
	return s
}

func (s *FindInstanceInfoResponseBody) SetMessage(v string) *FindInstanceInfoResponseBody {
	s.Message = &v
	return s
}

func (s *FindInstanceInfoResponseBody) SetRequestId(v string) *FindInstanceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindInstanceInfoResponseBody) SetSuccess(v bool) *FindInstanceInfoResponseBody {
	s.Success = &v
	return s
}

type FindInstanceInfoResponseBodyData struct {
	PegInstanceInfoList []*FindInstanceInfoResponseBodyDataPegInstanceInfoList `json:"PegInstanceInfoList,omitempty" xml:"PegInstanceInfoList,omitempty" type:"Repeated"`
}

func (s FindInstanceInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FindInstanceInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *FindInstanceInfoResponseBodyData) SetPegInstanceInfoList(v []*FindInstanceInfoResponseBodyDataPegInstanceInfoList) *FindInstanceInfoResponseBodyData {
	s.PegInstanceInfoList = v
	return s
}

type FindInstanceInfoResponseBodyDataPegInstanceInfoList struct {
	BussinessCode      *string                                                      `json:"BussinessCode,omitempty" xml:"BussinessCode,omitempty"`
	Coordinate         map[string]interface{}                                       `json:"Coordinate,omitempty" xml:"Coordinate,omitempty"`
	IdType             *string                                                      `json:"IdType,omitempty" xml:"IdType,omitempty"`
	InstanceId         *string                                                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ServiceCreatedTime *string                                                      `json:"ServiceCreatedTime,omitempty" xml:"ServiceCreatedTime,omitempty"`
	UserId             *string                                                      `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserInfo           *FindInstanceInfoResponseBodyDataPegInstanceInfoListUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s FindInstanceInfoResponseBodyDataPegInstanceInfoList) String() string {
	return tea.Prettify(s)
}

func (s FindInstanceInfoResponseBodyDataPegInstanceInfoList) GoString() string {
	return s.String()
}

func (s *FindInstanceInfoResponseBodyDataPegInstanceInfoList) SetBussinessCode(v string) *FindInstanceInfoResponseBodyDataPegInstanceInfoList {
	s.BussinessCode = &v
	return s
}

func (s *FindInstanceInfoResponseBodyDataPegInstanceInfoList) SetCoordinate(v map[string]interface{}) *FindInstanceInfoResponseBodyDataPegInstanceInfoList {
	s.Coordinate = v
	return s
}

func (s *FindInstanceInfoResponseBodyDataPegInstanceInfoList) SetIdType(v string) *FindInstanceInfoResponseBodyDataPegInstanceInfoList {
	s.IdType = &v
	return s
}

func (s *FindInstanceInfoResponseBodyDataPegInstanceInfoList) SetInstanceId(v string) *FindInstanceInfoResponseBodyDataPegInstanceInfoList {
	s.InstanceId = &v
	return s
}

func (s *FindInstanceInfoResponseBodyDataPegInstanceInfoList) SetServiceCreatedTime(v string) *FindInstanceInfoResponseBodyDataPegInstanceInfoList {
	s.ServiceCreatedTime = &v
	return s
}

func (s *FindInstanceInfoResponseBodyDataPegInstanceInfoList) SetUserId(v string) *FindInstanceInfoResponseBodyDataPegInstanceInfoList {
	s.UserId = &v
	return s
}

func (s *FindInstanceInfoResponseBodyDataPegInstanceInfoList) SetUserInfo(v *FindInstanceInfoResponseBodyDataPegInstanceInfoListUserInfo) *FindInstanceInfoResponseBodyDataPegInstanceInfoList {
	s.UserInfo = v
	return s
}

type FindInstanceInfoResponseBodyDataPegInstanceInfoListUserInfo struct {
	GcLevel        *string `json:"GcLevel,omitempty" xml:"GcLevel,omitempty"`
	HitWhiteReason *string `json:"HitWhiteReason,omitempty" xml:"HitWhiteReason,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserSite       *string `json:"UserSite,omitempty" xml:"UserSite,omitempty"`
	WhiteUser      *bool   `json:"WhiteUser,omitempty" xml:"WhiteUser,omitempty"`
}

func (s FindInstanceInfoResponseBodyDataPegInstanceInfoListUserInfo) String() string {
	return tea.Prettify(s)
}

func (s FindInstanceInfoResponseBodyDataPegInstanceInfoListUserInfo) GoString() string {
	return s.String()
}

func (s *FindInstanceInfoResponseBodyDataPegInstanceInfoListUserInfo) SetGcLevel(v string) *FindInstanceInfoResponseBodyDataPegInstanceInfoListUserInfo {
	s.GcLevel = &v
	return s
}

func (s *FindInstanceInfoResponseBodyDataPegInstanceInfoListUserInfo) SetHitWhiteReason(v string) *FindInstanceInfoResponseBodyDataPegInstanceInfoListUserInfo {
	s.HitWhiteReason = &v
	return s
}

func (s *FindInstanceInfoResponseBodyDataPegInstanceInfoListUserInfo) SetUserId(v string) *FindInstanceInfoResponseBodyDataPegInstanceInfoListUserInfo {
	s.UserId = &v
	return s
}

func (s *FindInstanceInfoResponseBodyDataPegInstanceInfoListUserInfo) SetUserSite(v string) *FindInstanceInfoResponseBodyDataPegInstanceInfoListUserInfo {
	s.UserSite = &v
	return s
}

func (s *FindInstanceInfoResponseBodyDataPegInstanceInfoListUserInfo) SetWhiteUser(v bool) *FindInstanceInfoResponseBodyDataPegInstanceInfoListUserInfo {
	s.WhiteUser = &v
	return s
}

type FindInstanceInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FindInstanceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FindInstanceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s FindInstanceInfoResponse) GoString() string {
	return s.String()
}

func (s *FindInstanceInfoResponse) SetHeaders(v map[string]*string) *FindInstanceInfoResponse {
	s.Headers = v
	return s
}

func (s *FindInstanceInfoResponse) SetStatusCode(v int32) *FindInstanceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *FindInstanceInfoResponse) SetBody(v *FindInstanceInfoResponseBody) *FindInstanceInfoResponse {
	s.Body = v
	return s
}

type FindUserAvailbleResourcesRequest struct {
	BussinessCode *string `json:"bussinessCode,omitempty" xml:"bussinessCode,omitempty"`
	CurrPage      *int32  `json:"currPage,omitempty" xml:"currPage,omitempty"`
	PageSize      *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ResourceType  *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	UserId        *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s FindUserAvailbleResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s FindUserAvailbleResourcesRequest) GoString() string {
	return s.String()
}

func (s *FindUserAvailbleResourcesRequest) SetBussinessCode(v string) *FindUserAvailbleResourcesRequest {
	s.BussinessCode = &v
	return s
}

func (s *FindUserAvailbleResourcesRequest) SetCurrPage(v int32) *FindUserAvailbleResourcesRequest {
	s.CurrPage = &v
	return s
}

func (s *FindUserAvailbleResourcesRequest) SetPageSize(v int32) *FindUserAvailbleResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *FindUserAvailbleResourcesRequest) SetResourceType(v string) *FindUserAvailbleResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *FindUserAvailbleResourcesRequest) SetUserId(v string) *FindUserAvailbleResourcesRequest {
	s.UserId = &v
	return s
}

type FindUserAvailbleResourcesResponseBody struct {
	Code       *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Count      *int64                                     `json:"Count,omitempty" xml:"Count,omitempty"`
	Data       *FindUserAvailbleResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	MaxResults *int32                                     `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message    *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	// This parameter is required.
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s FindUserAvailbleResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindUserAvailbleResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *FindUserAvailbleResourcesResponseBody) SetCode(v string) *FindUserAvailbleResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *FindUserAvailbleResourcesResponseBody) SetCount(v int64) *FindUserAvailbleResourcesResponseBody {
	s.Count = &v
	return s
}

func (s *FindUserAvailbleResourcesResponseBody) SetData(v *FindUserAvailbleResourcesResponseBodyData) *FindUserAvailbleResourcesResponseBody {
	s.Data = v
	return s
}

func (s *FindUserAvailbleResourcesResponseBody) SetMaxResults(v int32) *FindUserAvailbleResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *FindUserAvailbleResourcesResponseBody) SetMessage(v string) *FindUserAvailbleResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *FindUserAvailbleResourcesResponseBody) SetNextToken(v string) *FindUserAvailbleResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *FindUserAvailbleResourcesResponseBody) SetRequestId(v string) *FindUserAvailbleResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindUserAvailbleResourcesResponseBody) SetSuccess(v bool) *FindUserAvailbleResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *FindUserAvailbleResourcesResponseBody) SetTotalCount(v int32) *FindUserAvailbleResourcesResponseBody {
	s.TotalCount = &v
	return s
}

type FindUserAvailbleResourcesResponseBodyData struct {
	PegCoordinates []*FindUserAvailbleResourcesResponseBodyDataPegCoordinates `json:"PegCoordinates,omitempty" xml:"PegCoordinates,omitempty" type:"Repeated"`
	UserInfo       *FindUserAvailbleResourcesResponseBodyDataUserInfo         `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s FindUserAvailbleResourcesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FindUserAvailbleResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *FindUserAvailbleResourcesResponseBodyData) SetPegCoordinates(v []*FindUserAvailbleResourcesResponseBodyDataPegCoordinates) *FindUserAvailbleResourcesResponseBodyData {
	s.PegCoordinates = v
	return s
}

func (s *FindUserAvailbleResourcesResponseBodyData) SetUserInfo(v *FindUserAvailbleResourcesResponseBodyDataUserInfo) *FindUserAvailbleResourcesResponseBodyData {
	s.UserInfo = v
	return s
}

type FindUserAvailbleResourcesResponseBodyDataPegCoordinates struct {
	BussinessCode      *string                `json:"BussinessCode,omitempty" xml:"BussinessCode,omitempty"`
	ChargeType         *string                `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Coordinate         map[string]interface{} `json:"Coordinate,omitempty" xml:"Coordinate,omitempty"`
	IdType             *string                `json:"IdType,omitempty" xml:"IdType,omitempty"`
	InstanceId         *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region             *string                `json:"Region,omitempty" xml:"Region,omitempty"`
	ReleaseTime        *string                `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	ResCreateTime      *string                `json:"ResCreateTime,omitempty" xml:"ResCreateTime,omitempty"`
	ResourceStatus     *string                `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	ResourceType       *string                `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ServiceCreatedTime *string                `json:"ServiceCreatedTime,omitempty" xml:"ServiceCreatedTime,omitempty"`
	UserId             *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s FindUserAvailbleResourcesResponseBodyDataPegCoordinates) String() string {
	return tea.Prettify(s)
}

func (s FindUserAvailbleResourcesResponseBodyDataPegCoordinates) GoString() string {
	return s.String()
}

func (s *FindUserAvailbleResourcesResponseBodyDataPegCoordinates) SetBussinessCode(v string) *FindUserAvailbleResourcesResponseBodyDataPegCoordinates {
	s.BussinessCode = &v
	return s
}

func (s *FindUserAvailbleResourcesResponseBodyDataPegCoordinates) SetChargeType(v string) *FindUserAvailbleResourcesResponseBodyDataPegCoordinates {
	s.ChargeType = &v
	return s
}

func (s *FindUserAvailbleResourcesResponseBodyDataPegCoordinates) SetCoordinate(v map[string]interface{}) *FindUserAvailbleResourcesResponseBodyDataPegCoordinates {
	s.Coordinate = v
	return s
}

func (s *FindUserAvailbleResourcesResponseBodyDataPegCoordinates) SetIdType(v string) *FindUserAvailbleResourcesResponseBodyDataPegCoordinates {
	s.IdType = &v
	return s
}

func (s *FindUserAvailbleResourcesResponseBodyDataPegCoordinates) SetInstanceId(v string) *FindUserAvailbleResourcesResponseBodyDataPegCoordinates {
	s.InstanceId = &v
	return s
}

func (s *FindUserAvailbleResourcesResponseBodyDataPegCoordinates) SetRegion(v string) *FindUserAvailbleResourcesResponseBodyDataPegCoordinates {
	s.Region = &v
	return s
}

func (s *FindUserAvailbleResourcesResponseBodyDataPegCoordinates) SetReleaseTime(v string) *FindUserAvailbleResourcesResponseBodyDataPegCoordinates {
	s.ReleaseTime = &v
	return s
}

func (s *FindUserAvailbleResourcesResponseBodyDataPegCoordinates) SetResCreateTime(v string) *FindUserAvailbleResourcesResponseBodyDataPegCoordinates {
	s.ResCreateTime = &v
	return s
}

func (s *FindUserAvailbleResourcesResponseBodyDataPegCoordinates) SetResourceStatus(v string) *FindUserAvailbleResourcesResponseBodyDataPegCoordinates {
	s.ResourceStatus = &v
	return s
}

func (s *FindUserAvailbleResourcesResponseBodyDataPegCoordinates) SetResourceType(v string) *FindUserAvailbleResourcesResponseBodyDataPegCoordinates {
	s.ResourceType = &v
	return s
}

func (s *FindUserAvailbleResourcesResponseBodyDataPegCoordinates) SetServiceCreatedTime(v string) *FindUserAvailbleResourcesResponseBodyDataPegCoordinates {
	s.ServiceCreatedTime = &v
	return s
}

func (s *FindUserAvailbleResourcesResponseBodyDataPegCoordinates) SetUserId(v string) *FindUserAvailbleResourcesResponseBodyDataPegCoordinates {
	s.UserId = &v
	return s
}

type FindUserAvailbleResourcesResponseBodyDataUserInfo struct {
	GcLevel        *string `json:"GcLevel,omitempty" xml:"GcLevel,omitempty"`
	HitWhiteReason *string `json:"HitWhiteReason,omitempty" xml:"HitWhiteReason,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserSite       *string `json:"UserSite,omitempty" xml:"UserSite,omitempty"`
	WhiteUser      *bool   `json:"WhiteUser,omitempty" xml:"WhiteUser,omitempty"`
}

func (s FindUserAvailbleResourcesResponseBodyDataUserInfo) String() string {
	return tea.Prettify(s)
}

func (s FindUserAvailbleResourcesResponseBodyDataUserInfo) GoString() string {
	return s.String()
}

func (s *FindUserAvailbleResourcesResponseBodyDataUserInfo) SetGcLevel(v string) *FindUserAvailbleResourcesResponseBodyDataUserInfo {
	s.GcLevel = &v
	return s
}

func (s *FindUserAvailbleResourcesResponseBodyDataUserInfo) SetHitWhiteReason(v string) *FindUserAvailbleResourcesResponseBodyDataUserInfo {
	s.HitWhiteReason = &v
	return s
}

func (s *FindUserAvailbleResourcesResponseBodyDataUserInfo) SetUserId(v string) *FindUserAvailbleResourcesResponseBodyDataUserInfo {
	s.UserId = &v
	return s
}

func (s *FindUserAvailbleResourcesResponseBodyDataUserInfo) SetUserSite(v string) *FindUserAvailbleResourcesResponseBodyDataUserInfo {
	s.UserSite = &v
	return s
}

func (s *FindUserAvailbleResourcesResponseBodyDataUserInfo) SetWhiteUser(v bool) *FindUserAvailbleResourcesResponseBodyDataUserInfo {
	s.WhiteUser = &v
	return s
}

type FindUserAvailbleResourcesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FindUserAvailbleResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FindUserAvailbleResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s FindUserAvailbleResourcesResponse) GoString() string {
	return s.String()
}

func (s *FindUserAvailbleResourcesResponse) SetHeaders(v map[string]*string) *FindUserAvailbleResourcesResponse {
	s.Headers = v
	return s
}

func (s *FindUserAvailbleResourcesResponse) SetStatusCode(v int32) *FindUserAvailbleResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *FindUserAvailbleResourcesResponse) SetBody(v *FindUserAvailbleResourcesResponseBody) *FindUserAvailbleResourcesResponse {
	s.Body = v
	return s
}

type PunishResourceSearchRequest struct {
	ActionCodes    []*string `json:"ActionCodes,omitempty" xml:"ActionCodes,omitempty" type:"Repeated"`
	BussinessCodes []*string `json:"BussinessCodes,omitempty" xml:"BussinessCodes,omitempty" type:"Repeated"`
	Class          *string   `json:"Class,omitempty" xml:"Class,omitempty"`
	Domain         *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	EndDate        *int64    `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	InstanceId     *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ip             *string   `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Page           *int64    `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize       *int64    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SourceCodes    []*string `json:"SourceCodes,omitempty" xml:"SourceCodes,omitempty" type:"Repeated"`
	StartDate      *int64    `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Status         *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	Url            *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	UserIds        []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s PunishResourceSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s PunishResourceSearchRequest) GoString() string {
	return s.String()
}

func (s *PunishResourceSearchRequest) SetActionCodes(v []*string) *PunishResourceSearchRequest {
	s.ActionCodes = v
	return s
}

func (s *PunishResourceSearchRequest) SetBussinessCodes(v []*string) *PunishResourceSearchRequest {
	s.BussinessCodes = v
	return s
}

func (s *PunishResourceSearchRequest) SetClass(v string) *PunishResourceSearchRequest {
	s.Class = &v
	return s
}

func (s *PunishResourceSearchRequest) SetDomain(v string) *PunishResourceSearchRequest {
	s.Domain = &v
	return s
}

func (s *PunishResourceSearchRequest) SetEndDate(v int64) *PunishResourceSearchRequest {
	s.EndDate = &v
	return s
}

func (s *PunishResourceSearchRequest) SetInstanceId(v string) *PunishResourceSearchRequest {
	s.InstanceId = &v
	return s
}

func (s *PunishResourceSearchRequest) SetIp(v string) *PunishResourceSearchRequest {
	s.Ip = &v
	return s
}

func (s *PunishResourceSearchRequest) SetPage(v int64) *PunishResourceSearchRequest {
	s.Page = &v
	return s
}

func (s *PunishResourceSearchRequest) SetPageSize(v int64) *PunishResourceSearchRequest {
	s.PageSize = &v
	return s
}

func (s *PunishResourceSearchRequest) SetSourceCodes(v []*string) *PunishResourceSearchRequest {
	s.SourceCodes = v
	return s
}

func (s *PunishResourceSearchRequest) SetStartDate(v int64) *PunishResourceSearchRequest {
	s.StartDate = &v
	return s
}

func (s *PunishResourceSearchRequest) SetStatus(v string) *PunishResourceSearchRequest {
	s.Status = &v
	return s
}

func (s *PunishResourceSearchRequest) SetUrl(v string) *PunishResourceSearchRequest {
	s.Url = &v
	return s
}

func (s *PunishResourceSearchRequest) SetUserIds(v []*string) *PunishResourceSearchRequest {
	s.UserIds = v
	return s
}

type PunishResourceSearchShrinkRequest struct {
	ActionCodesShrink    *string `json:"ActionCodes,omitempty" xml:"ActionCodes,omitempty"`
	BussinessCodesShrink *string `json:"BussinessCodes,omitempty" xml:"BussinessCodes,omitempty"`
	Class                *string `json:"Class,omitempty" xml:"Class,omitempty"`
	Domain               *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	EndDate              *int64  `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ip                   *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Page                 *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SourceCodesShrink    *string `json:"SourceCodes,omitempty" xml:"SourceCodes,omitempty"`
	StartDate            *int64  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Url                  *string `json:"Url,omitempty" xml:"Url,omitempty"`
	UserIdsShrink        *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s PunishResourceSearchShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s PunishResourceSearchShrinkRequest) GoString() string {
	return s.String()
}

func (s *PunishResourceSearchShrinkRequest) SetActionCodesShrink(v string) *PunishResourceSearchShrinkRequest {
	s.ActionCodesShrink = &v
	return s
}

func (s *PunishResourceSearchShrinkRequest) SetBussinessCodesShrink(v string) *PunishResourceSearchShrinkRequest {
	s.BussinessCodesShrink = &v
	return s
}

func (s *PunishResourceSearchShrinkRequest) SetClass(v string) *PunishResourceSearchShrinkRequest {
	s.Class = &v
	return s
}

func (s *PunishResourceSearchShrinkRequest) SetDomain(v string) *PunishResourceSearchShrinkRequest {
	s.Domain = &v
	return s
}

func (s *PunishResourceSearchShrinkRequest) SetEndDate(v int64) *PunishResourceSearchShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *PunishResourceSearchShrinkRequest) SetInstanceId(v string) *PunishResourceSearchShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *PunishResourceSearchShrinkRequest) SetIp(v string) *PunishResourceSearchShrinkRequest {
	s.Ip = &v
	return s
}

func (s *PunishResourceSearchShrinkRequest) SetPage(v int64) *PunishResourceSearchShrinkRequest {
	s.Page = &v
	return s
}

func (s *PunishResourceSearchShrinkRequest) SetPageSize(v int64) *PunishResourceSearchShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *PunishResourceSearchShrinkRequest) SetSourceCodesShrink(v string) *PunishResourceSearchShrinkRequest {
	s.SourceCodesShrink = &v
	return s
}

func (s *PunishResourceSearchShrinkRequest) SetStartDate(v int64) *PunishResourceSearchShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *PunishResourceSearchShrinkRequest) SetStatus(v string) *PunishResourceSearchShrinkRequest {
	s.Status = &v
	return s
}

func (s *PunishResourceSearchShrinkRequest) SetUrl(v string) *PunishResourceSearchShrinkRequest {
	s.Url = &v
	return s
}

func (s *PunishResourceSearchShrinkRequest) SetUserIdsShrink(v string) *PunishResourceSearchShrinkRequest {
	s.UserIdsShrink = &v
	return s
}

type PunishResourceSearchResponseBody struct {
	Code       *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	DataList   []*PunishResourceSearchResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	Message    *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	Success    *string                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	ViewCount  *int32                                      `json:"ViewCount,omitempty" xml:"ViewCount,omitempty"`
}

func (s PunishResourceSearchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PunishResourceSearchResponseBody) GoString() string {
	return s.String()
}

func (s *PunishResourceSearchResponseBody) SetCode(v string) *PunishResourceSearchResponseBody {
	s.Code = &v
	return s
}

func (s *PunishResourceSearchResponseBody) SetDataList(v []*PunishResourceSearchResponseBodyDataList) *PunishResourceSearchResponseBody {
	s.DataList = v
	return s
}

func (s *PunishResourceSearchResponseBody) SetMessage(v string) *PunishResourceSearchResponseBody {
	s.Message = &v
	return s
}

func (s *PunishResourceSearchResponseBody) SetSuccess(v string) *PunishResourceSearchResponseBody {
	s.Success = &v
	return s
}

func (s *PunishResourceSearchResponseBody) SetTotalCount(v int32) *PunishResourceSearchResponseBody {
	s.TotalCount = &v
	return s
}

func (s *PunishResourceSearchResponseBody) SetViewCount(v int32) *PunishResourceSearchResponseBody {
	s.ViewCount = &v
	return s
}

type PunishResourceSearchResponseBodyDataList struct {
	ActionCode    *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	BussinessCode *string `json:"BussinessCode,omitempty" xml:"BussinessCode,omitempty"`
	Class         *string `json:"Class,omitempty" xml:"Class,omitempty"`
	Coordinate    *string `json:"Coordinate,omitempty" xml:"Coordinate,omitempty"`
	Creator       *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Deleted       *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Extras        *string `json:"Extras,omitempty" xml:"Extras,omitempty"`
	GmtCreated    *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified   *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ip            *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Modifier      *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	ObjectId      *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	ObjectType    *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	ObjectValue   *string `json:"ObjectValue,omitempty" xml:"ObjectValue,omitempty"`
	PunishFrom    *string `json:"PunishFrom,omitempty" xml:"PunishFrom,omitempty"`
	Reason        *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	RequestId     *int64  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SourceCode    *string `json:"SourceCode,omitempty" xml:"SourceCode,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Url           *string `json:"Url,omitempty" xml:"Url,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s PunishResourceSearchResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s PunishResourceSearchResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *PunishResourceSearchResponseBodyDataList) SetActionCode(v string) *PunishResourceSearchResponseBodyDataList {
	s.ActionCode = &v
	return s
}

func (s *PunishResourceSearchResponseBodyDataList) SetBussinessCode(v string) *PunishResourceSearchResponseBodyDataList {
	s.BussinessCode = &v
	return s
}

func (s *PunishResourceSearchResponseBodyDataList) SetClass(v string) *PunishResourceSearchResponseBodyDataList {
	s.Class = &v
	return s
}

func (s *PunishResourceSearchResponseBodyDataList) SetCoordinate(v string) *PunishResourceSearchResponseBodyDataList {
	s.Coordinate = &v
	return s
}

func (s *PunishResourceSearchResponseBodyDataList) SetCreator(v string) *PunishResourceSearchResponseBodyDataList {
	s.Creator = &v
	return s
}

func (s *PunishResourceSearchResponseBodyDataList) SetDeleted(v bool) *PunishResourceSearchResponseBodyDataList {
	s.Deleted = &v
	return s
}

func (s *PunishResourceSearchResponseBodyDataList) SetDomain(v string) *PunishResourceSearchResponseBodyDataList {
	s.Domain = &v
	return s
}

func (s *PunishResourceSearchResponseBodyDataList) SetExtras(v string) *PunishResourceSearchResponseBodyDataList {
	s.Extras = &v
	return s
}

func (s *PunishResourceSearchResponseBodyDataList) SetGmtCreated(v string) *PunishResourceSearchResponseBodyDataList {
	s.GmtCreated = &v
	return s
}

func (s *PunishResourceSearchResponseBodyDataList) SetGmtModified(v string) *PunishResourceSearchResponseBodyDataList {
	s.GmtModified = &v
	return s
}

func (s *PunishResourceSearchResponseBodyDataList) SetId(v int64) *PunishResourceSearchResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *PunishResourceSearchResponseBodyDataList) SetInstanceId(v string) *PunishResourceSearchResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *PunishResourceSearchResponseBodyDataList) SetIp(v string) *PunishResourceSearchResponseBodyDataList {
	s.Ip = &v
	return s
}

func (s *PunishResourceSearchResponseBodyDataList) SetModifier(v string) *PunishResourceSearchResponseBodyDataList {
	s.Modifier = &v
	return s
}

func (s *PunishResourceSearchResponseBodyDataList) SetObjectId(v string) *PunishResourceSearchResponseBodyDataList {
	s.ObjectId = &v
	return s
}

func (s *PunishResourceSearchResponseBodyDataList) SetObjectType(v string) *PunishResourceSearchResponseBodyDataList {
	s.ObjectType = &v
	return s
}

func (s *PunishResourceSearchResponseBodyDataList) SetObjectValue(v string) *PunishResourceSearchResponseBodyDataList {
	s.ObjectValue = &v
	return s
}

func (s *PunishResourceSearchResponseBodyDataList) SetPunishFrom(v string) *PunishResourceSearchResponseBodyDataList {
	s.PunishFrom = &v
	return s
}

func (s *PunishResourceSearchResponseBodyDataList) SetReason(v string) *PunishResourceSearchResponseBodyDataList {
	s.Reason = &v
	return s
}

func (s *PunishResourceSearchResponseBodyDataList) SetRequestId(v int64) *PunishResourceSearchResponseBodyDataList {
	s.RequestId = &v
	return s
}

func (s *PunishResourceSearchResponseBodyDataList) SetSourceCode(v string) *PunishResourceSearchResponseBodyDataList {
	s.SourceCode = &v
	return s
}

func (s *PunishResourceSearchResponseBodyDataList) SetStatus(v string) *PunishResourceSearchResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *PunishResourceSearchResponseBodyDataList) SetUrl(v string) *PunishResourceSearchResponseBodyDataList {
	s.Url = &v
	return s
}

func (s *PunishResourceSearchResponseBodyDataList) SetUserId(v string) *PunishResourceSearchResponseBodyDataList {
	s.UserId = &v
	return s
}

type PunishResourceSearchResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PunishResourceSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PunishResourceSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s PunishResourceSearchResponse) GoString() string {
	return s.String()
}

func (s *PunishResourceSearchResponse) SetHeaders(v map[string]*string) *PunishResourceSearchResponse {
	s.Headers = v
	return s
}

func (s *PunishResourceSearchResponse) SetStatusCode(v int32) *PunishResourceSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *PunishResourceSearchResponse) SetBody(v *PunishResourceSearchResponseBody) *PunishResourceSearchResponse {
	s.Body = v
	return s
}

type RiskEventSyncRequest struct {
	Deleted *bool `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	// This parameter is required.
	DiscoveryTime *int64 `json:"DiscoveryTime,omitempty" xml:"DiscoveryTime,omitempty"`
	// This parameter is required.
	EventName   *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	EventNumber *string `json:"EventNumber,omitempty" xml:"EventNumber,omitempty"`
	RelevanceBu *string `json:"RelevanceBu,omitempty" xml:"RelevanceBu,omitempty"`
	// This parameter is required.
	RiskDetail *string `json:"RiskDetail,omitempty" xml:"RiskDetail,omitempty"`
	// This parameter is required.
	RiskEffectType *string `json:"RiskEffectType,omitempty" xml:"RiskEffectType,omitempty"`
	RiskLevel      *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// This parameter is required.
	RiskType *string `json:"RiskType,omitempty" xml:"RiskType,omitempty"`
	// This parameter is required.
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// This parameter is required.
	Submitter *string `json:"Submitter,omitempty" xml:"Submitter,omitempty"`
}

func (s RiskEventSyncRequest) String() string {
	return tea.Prettify(s)
}

func (s RiskEventSyncRequest) GoString() string {
	return s.String()
}

func (s *RiskEventSyncRequest) SetDeleted(v bool) *RiskEventSyncRequest {
	s.Deleted = &v
	return s
}

func (s *RiskEventSyncRequest) SetDiscoveryTime(v int64) *RiskEventSyncRequest {
	s.DiscoveryTime = &v
	return s
}

func (s *RiskEventSyncRequest) SetEventName(v string) *RiskEventSyncRequest {
	s.EventName = &v
	return s
}

func (s *RiskEventSyncRequest) SetEventNumber(v string) *RiskEventSyncRequest {
	s.EventNumber = &v
	return s
}

func (s *RiskEventSyncRequest) SetRelevanceBu(v string) *RiskEventSyncRequest {
	s.RelevanceBu = &v
	return s
}

func (s *RiskEventSyncRequest) SetRiskDetail(v string) *RiskEventSyncRequest {
	s.RiskDetail = &v
	return s
}

func (s *RiskEventSyncRequest) SetRiskEffectType(v string) *RiskEventSyncRequest {
	s.RiskEffectType = &v
	return s
}

func (s *RiskEventSyncRequest) SetRiskLevel(v string) *RiskEventSyncRequest {
	s.RiskLevel = &v
	return s
}

func (s *RiskEventSyncRequest) SetRiskType(v string) *RiskEventSyncRequest {
	s.RiskType = &v
	return s
}

func (s *RiskEventSyncRequest) SetSource(v string) *RiskEventSyncRequest {
	s.Source = &v
	return s
}

func (s *RiskEventSyncRequest) SetSubmitter(v string) *RiskEventSyncRequest {
	s.Submitter = &v
	return s
}

type RiskEventSyncResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Count   *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RiskEventSyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RiskEventSyncResponseBody) GoString() string {
	return s.String()
}

func (s *RiskEventSyncResponseBody) SetCode(v string) *RiskEventSyncResponseBody {
	s.Code = &v
	return s
}

func (s *RiskEventSyncResponseBody) SetCount(v int32) *RiskEventSyncResponseBody {
	s.Count = &v
	return s
}

func (s *RiskEventSyncResponseBody) SetData(v string) *RiskEventSyncResponseBody {
	s.Data = &v
	return s
}

func (s *RiskEventSyncResponseBody) SetMessage(v string) *RiskEventSyncResponseBody {
	s.Message = &v
	return s
}

func (s *RiskEventSyncResponseBody) SetSuccess(v bool) *RiskEventSyncResponseBody {
	s.Success = &v
	return s
}

type RiskEventSyncResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RiskEventSyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RiskEventSyncResponse) String() string {
	return tea.Prettify(s)
}

func (s RiskEventSyncResponse) GoString() string {
	return s.String()
}

func (s *RiskEventSyncResponse) SetHeaders(v map[string]*string) *RiskEventSyncResponse {
	s.Headers = v
	return s
}

func (s *RiskEventSyncResponse) SetStatusCode(v int32) *RiskEventSyncResponse {
	s.StatusCode = &v
	return s
}

func (s *RiskEventSyncResponse) SetBody(v *RiskEventSyncResponseBody) *RiskEventSyncResponse {
	s.Body = v
	return s
}

type SearchPunishEventsRequest struct {
	// This parameter is required.
	AliUid         *string   `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	BussinessCodes []*string `json:"BussinessCodes,omitempty" xml:"BussinessCodes,omitempty" type:"Repeated"`
	CaseCodes      []*string `json:"CaseCodes,omitempty" xml:"CaseCodes,omitempty" type:"Repeated"`
	EventCodes     []*string `json:"EventCodes,omitempty" xml:"EventCodes,omitempty" type:"Repeated"`
	ResourceId     *string   `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s SearchPunishEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchPunishEventsRequest) GoString() string {
	return s.String()
}

func (s *SearchPunishEventsRequest) SetAliUid(v string) *SearchPunishEventsRequest {
	s.AliUid = &v
	return s
}

func (s *SearchPunishEventsRequest) SetBussinessCodes(v []*string) *SearchPunishEventsRequest {
	s.BussinessCodes = v
	return s
}

func (s *SearchPunishEventsRequest) SetCaseCodes(v []*string) *SearchPunishEventsRequest {
	s.CaseCodes = v
	return s
}

func (s *SearchPunishEventsRequest) SetEventCodes(v []*string) *SearchPunishEventsRequest {
	s.EventCodes = v
	return s
}

func (s *SearchPunishEventsRequest) SetResourceId(v string) *SearchPunishEventsRequest {
	s.ResourceId = &v
	return s
}

type SearchPunishEventsShrinkRequest struct {
	// This parameter is required.
	AliUid               *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	BussinessCodesShrink *string `json:"BussinessCodes,omitempty" xml:"BussinessCodes,omitempty"`
	CaseCodesShrink      *string `json:"CaseCodes,omitempty" xml:"CaseCodes,omitempty"`
	EventCodesShrink     *string `json:"EventCodes,omitempty" xml:"EventCodes,omitempty"`
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s SearchPunishEventsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchPunishEventsShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchPunishEventsShrinkRequest) SetAliUid(v string) *SearchPunishEventsShrinkRequest {
	s.AliUid = &v
	return s
}

func (s *SearchPunishEventsShrinkRequest) SetBussinessCodesShrink(v string) *SearchPunishEventsShrinkRequest {
	s.BussinessCodesShrink = &v
	return s
}

func (s *SearchPunishEventsShrinkRequest) SetCaseCodesShrink(v string) *SearchPunishEventsShrinkRequest {
	s.CaseCodesShrink = &v
	return s
}

func (s *SearchPunishEventsShrinkRequest) SetEventCodesShrink(v string) *SearchPunishEventsShrinkRequest {
	s.EventCodesShrink = &v
	return s
}

func (s *SearchPunishEventsShrinkRequest) SetResourceId(v string) *SearchPunishEventsShrinkRequest {
	s.ResourceId = &v
	return s
}

type SearchPunishEventsResponseBody struct {
	Code       *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	DataList   []*SearchPunishEventsResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	Message    *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchPunishEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchPunishEventsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchPunishEventsResponseBody) SetCode(v string) *SearchPunishEventsResponseBody {
	s.Code = &v
	return s
}

func (s *SearchPunishEventsResponseBody) SetDataList(v []*SearchPunishEventsResponseBodyDataList) *SearchPunishEventsResponseBody {
	s.DataList = v
	return s
}

func (s *SearchPunishEventsResponseBody) SetMessage(v string) *SearchPunishEventsResponseBody {
	s.Message = &v
	return s
}

func (s *SearchPunishEventsResponseBody) SetRequestId(v string) *SearchPunishEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchPunishEventsResponseBody) SetTotalCount(v int32) *SearchPunishEventsResponseBody {
	s.TotalCount = &v
	return s
}

type SearchPunishEventsResponseBodyDataList struct {
	BussinessCode *string `json:"BussinessCode,omitempty" xml:"BussinessCode,omitempty"`
	CaseCode      *string `json:"CaseCode,omitempty" xml:"CaseCode,omitempty"`
	RecordsCount  *int32  `json:"RecordsCount,omitempty" xml:"RecordsCount,omitempty"`
	ResourceId    *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TipsCode      *string `json:"TipsCode,omitempty" xml:"TipsCode,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchPunishEventsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s SearchPunishEventsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *SearchPunishEventsResponseBodyDataList) SetBussinessCode(v string) *SearchPunishEventsResponseBodyDataList {
	s.BussinessCode = &v
	return s
}

func (s *SearchPunishEventsResponseBodyDataList) SetCaseCode(v string) *SearchPunishEventsResponseBodyDataList {
	s.CaseCode = &v
	return s
}

func (s *SearchPunishEventsResponseBodyDataList) SetRecordsCount(v int32) *SearchPunishEventsResponseBodyDataList {
	s.RecordsCount = &v
	return s
}

func (s *SearchPunishEventsResponseBodyDataList) SetResourceId(v string) *SearchPunishEventsResponseBodyDataList {
	s.ResourceId = &v
	return s
}

func (s *SearchPunishEventsResponseBodyDataList) SetTipsCode(v string) *SearchPunishEventsResponseBodyDataList {
	s.TipsCode = &v
	return s
}

func (s *SearchPunishEventsResponseBodyDataList) SetUserId(v string) *SearchPunishEventsResponseBodyDataList {
	s.UserId = &v
	return s
}

type SearchPunishEventsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchPunishEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchPunishEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchPunishEventsResponse) GoString() string {
	return s.String()
}

func (s *SearchPunishEventsResponse) SetHeaders(v map[string]*string) *SearchPunishEventsResponse {
	s.Headers = v
	return s
}

func (s *SearchPunishEventsResponse) SetStatusCode(v int32) *SearchPunishEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchPunishEventsResponse) SetBody(v *SearchPunishEventsResponseBody) *SearchPunishEventsResponse {
	s.Body = v
	return s
}

type SearchPunishRecordsRequest struct {
	ActionCodes []*string `json:"ActionCodes,omitempty" xml:"ActionCodes,omitempty" type:"Repeated"`
	// This parameter is required.
	AliUid         *string   `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	BussinessCodes *string   `json:"BussinessCodes,omitempty" xml:"BussinessCodes,omitempty"`
	CaseCodes      []*string `json:"CaseCodes,omitempty" xml:"CaseCodes,omitempty" type:"Repeated"`
	Domain         *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	EndTime        *int64    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EventCodes     []*string `json:"EventCodes,omitempty" xml:"EventCodes,omitempty" type:"Repeated"`
	Ip             *string   `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Page           *string   `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize       *string   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PunishStatus   []*string `json:"PunishStatus,omitempty" xml:"PunishStatus,omitempty" type:"Repeated"`
	ResourceId     *string   `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	SourceCodes    []*string `json:"SourceCodes,omitempty" xml:"SourceCodes,omitempty" type:"Repeated"`
	StartTime      *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Url            *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	UrlFuzzy       *string   `json:"UrlFuzzy,omitempty" xml:"UrlFuzzy,omitempty"`
}

func (s SearchPunishRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchPunishRecordsRequest) GoString() string {
	return s.String()
}

func (s *SearchPunishRecordsRequest) SetActionCodes(v []*string) *SearchPunishRecordsRequest {
	s.ActionCodes = v
	return s
}

func (s *SearchPunishRecordsRequest) SetAliUid(v string) *SearchPunishRecordsRequest {
	s.AliUid = &v
	return s
}

func (s *SearchPunishRecordsRequest) SetBussinessCodes(v string) *SearchPunishRecordsRequest {
	s.BussinessCodes = &v
	return s
}

func (s *SearchPunishRecordsRequest) SetCaseCodes(v []*string) *SearchPunishRecordsRequest {
	s.CaseCodes = v
	return s
}

func (s *SearchPunishRecordsRequest) SetDomain(v string) *SearchPunishRecordsRequest {
	s.Domain = &v
	return s
}

func (s *SearchPunishRecordsRequest) SetEndTime(v int64) *SearchPunishRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *SearchPunishRecordsRequest) SetEventCodes(v []*string) *SearchPunishRecordsRequest {
	s.EventCodes = v
	return s
}

func (s *SearchPunishRecordsRequest) SetIp(v string) *SearchPunishRecordsRequest {
	s.Ip = &v
	return s
}

func (s *SearchPunishRecordsRequest) SetPage(v string) *SearchPunishRecordsRequest {
	s.Page = &v
	return s
}

func (s *SearchPunishRecordsRequest) SetPageSize(v string) *SearchPunishRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *SearchPunishRecordsRequest) SetPunishStatus(v []*string) *SearchPunishRecordsRequest {
	s.PunishStatus = v
	return s
}

func (s *SearchPunishRecordsRequest) SetResourceId(v string) *SearchPunishRecordsRequest {
	s.ResourceId = &v
	return s
}

func (s *SearchPunishRecordsRequest) SetSourceCodes(v []*string) *SearchPunishRecordsRequest {
	s.SourceCodes = v
	return s
}

func (s *SearchPunishRecordsRequest) SetStartTime(v int64) *SearchPunishRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *SearchPunishRecordsRequest) SetUrl(v string) *SearchPunishRecordsRequest {
	s.Url = &v
	return s
}

func (s *SearchPunishRecordsRequest) SetUrlFuzzy(v string) *SearchPunishRecordsRequest {
	s.UrlFuzzy = &v
	return s
}

type SearchPunishRecordsShrinkRequest struct {
	ActionCodesShrink *string `json:"ActionCodes,omitempty" xml:"ActionCodes,omitempty"`
	// This parameter is required.
	AliUid             *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	BussinessCodes     *string `json:"BussinessCodes,omitempty" xml:"BussinessCodes,omitempty"`
	CaseCodesShrink    *string `json:"CaseCodes,omitempty" xml:"CaseCodes,omitempty"`
	Domain             *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	EndTime            *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EventCodesShrink   *string `json:"EventCodes,omitempty" xml:"EventCodes,omitempty"`
	Ip                 *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Page               *string `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize           *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PunishStatusShrink *string `json:"PunishStatus,omitempty" xml:"PunishStatus,omitempty"`
	ResourceId         *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	SourceCodesShrink  *string `json:"SourceCodes,omitempty" xml:"SourceCodes,omitempty"`
	StartTime          *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Url                *string `json:"Url,omitempty" xml:"Url,omitempty"`
	UrlFuzzy           *string `json:"UrlFuzzy,omitempty" xml:"UrlFuzzy,omitempty"`
}

func (s SearchPunishRecordsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchPunishRecordsShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchPunishRecordsShrinkRequest) SetActionCodesShrink(v string) *SearchPunishRecordsShrinkRequest {
	s.ActionCodesShrink = &v
	return s
}

func (s *SearchPunishRecordsShrinkRequest) SetAliUid(v string) *SearchPunishRecordsShrinkRequest {
	s.AliUid = &v
	return s
}

func (s *SearchPunishRecordsShrinkRequest) SetBussinessCodes(v string) *SearchPunishRecordsShrinkRequest {
	s.BussinessCodes = &v
	return s
}

func (s *SearchPunishRecordsShrinkRequest) SetCaseCodesShrink(v string) *SearchPunishRecordsShrinkRequest {
	s.CaseCodesShrink = &v
	return s
}

func (s *SearchPunishRecordsShrinkRequest) SetDomain(v string) *SearchPunishRecordsShrinkRequest {
	s.Domain = &v
	return s
}

func (s *SearchPunishRecordsShrinkRequest) SetEndTime(v int64) *SearchPunishRecordsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *SearchPunishRecordsShrinkRequest) SetEventCodesShrink(v string) *SearchPunishRecordsShrinkRequest {
	s.EventCodesShrink = &v
	return s
}

func (s *SearchPunishRecordsShrinkRequest) SetIp(v string) *SearchPunishRecordsShrinkRequest {
	s.Ip = &v
	return s
}

func (s *SearchPunishRecordsShrinkRequest) SetPage(v string) *SearchPunishRecordsShrinkRequest {
	s.Page = &v
	return s
}

func (s *SearchPunishRecordsShrinkRequest) SetPageSize(v string) *SearchPunishRecordsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *SearchPunishRecordsShrinkRequest) SetPunishStatusShrink(v string) *SearchPunishRecordsShrinkRequest {
	s.PunishStatusShrink = &v
	return s
}

func (s *SearchPunishRecordsShrinkRequest) SetResourceId(v string) *SearchPunishRecordsShrinkRequest {
	s.ResourceId = &v
	return s
}

func (s *SearchPunishRecordsShrinkRequest) SetSourceCodesShrink(v string) *SearchPunishRecordsShrinkRequest {
	s.SourceCodesShrink = &v
	return s
}

func (s *SearchPunishRecordsShrinkRequest) SetStartTime(v int64) *SearchPunishRecordsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *SearchPunishRecordsShrinkRequest) SetUrl(v string) *SearchPunishRecordsShrinkRequest {
	s.Url = &v
	return s
}

func (s *SearchPunishRecordsShrinkRequest) SetUrlFuzzy(v string) *SearchPunishRecordsShrinkRequest {
	s.UrlFuzzy = &v
	return s
}

type SearchPunishRecordsResponseBody struct {
	Code       *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	DataList   []*SearchPunishRecordsResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	Message    *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchPunishRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchPunishRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchPunishRecordsResponseBody) SetCode(v string) *SearchPunishRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *SearchPunishRecordsResponseBody) SetDataList(v []*SearchPunishRecordsResponseBodyDataList) *SearchPunishRecordsResponseBody {
	s.DataList = v
	return s
}

func (s *SearchPunishRecordsResponseBody) SetMessage(v string) *SearchPunishRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *SearchPunishRecordsResponseBody) SetRequestId(v string) *SearchPunishRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchPunishRecordsResponseBody) SetSuccess(v bool) *SearchPunishRecordsResponseBody {
	s.Success = &v
	return s
}

func (s *SearchPunishRecordsResponseBody) SetTotalCount(v int32) *SearchPunishRecordsResponseBody {
	s.TotalCount = &v
	return s
}

type SearchPunishRecordsResponseBodyDataList struct {
	ActionCode    *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	AntiStatus    *string `json:"AntiStatus,omitempty" xml:"AntiStatus,omitempty"`
	BussinessCode *string `json:"BussinessCode,omitempty" xml:"BussinessCode,omitempty"`
	CaseCode      *string `json:"CaseCode,omitempty" xml:"CaseCode,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	EventCode     *string `json:"EventCode,omitempty" xml:"EventCode,omitempty"`
	Ip            *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	PunishStatus  *string `json:"PunishStatus,omitempty" xml:"PunishStatus,omitempty"`
	Reason        *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	ResourceId    *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TipsCode      *string `json:"TipsCode,omitempty" xml:"TipsCode,omitempty"`
	Url           *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SearchPunishRecordsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s SearchPunishRecordsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *SearchPunishRecordsResponseBodyDataList) SetActionCode(v string) *SearchPunishRecordsResponseBodyDataList {
	s.ActionCode = &v
	return s
}

func (s *SearchPunishRecordsResponseBodyDataList) SetAntiStatus(v string) *SearchPunishRecordsResponseBodyDataList {
	s.AntiStatus = &v
	return s
}

func (s *SearchPunishRecordsResponseBodyDataList) SetBussinessCode(v string) *SearchPunishRecordsResponseBodyDataList {
	s.BussinessCode = &v
	return s
}

func (s *SearchPunishRecordsResponseBodyDataList) SetCaseCode(v string) *SearchPunishRecordsResponseBodyDataList {
	s.CaseCode = &v
	return s
}

func (s *SearchPunishRecordsResponseBodyDataList) SetCreateTime(v string) *SearchPunishRecordsResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *SearchPunishRecordsResponseBodyDataList) SetDomain(v string) *SearchPunishRecordsResponseBodyDataList {
	s.Domain = &v
	return s
}

func (s *SearchPunishRecordsResponseBodyDataList) SetEventCode(v string) *SearchPunishRecordsResponseBodyDataList {
	s.EventCode = &v
	return s
}

func (s *SearchPunishRecordsResponseBodyDataList) SetIp(v string) *SearchPunishRecordsResponseBodyDataList {
	s.Ip = &v
	return s
}

func (s *SearchPunishRecordsResponseBodyDataList) SetPunishStatus(v string) *SearchPunishRecordsResponseBodyDataList {
	s.PunishStatus = &v
	return s
}

func (s *SearchPunishRecordsResponseBodyDataList) SetReason(v string) *SearchPunishRecordsResponseBodyDataList {
	s.Reason = &v
	return s
}

func (s *SearchPunishRecordsResponseBodyDataList) SetResourceId(v string) *SearchPunishRecordsResponseBodyDataList {
	s.ResourceId = &v
	return s
}

func (s *SearchPunishRecordsResponseBodyDataList) SetTipsCode(v string) *SearchPunishRecordsResponseBodyDataList {
	s.TipsCode = &v
	return s
}

func (s *SearchPunishRecordsResponseBodyDataList) SetUrl(v string) *SearchPunishRecordsResponseBodyDataList {
	s.Url = &v
	return s
}

type SearchPunishRecordsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchPunishRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchPunishRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchPunishRecordsResponse) GoString() string {
	return s.String()
}

func (s *SearchPunishRecordsResponse) SetHeaders(v map[string]*string) *SearchPunishRecordsResponse {
	s.Headers = v
	return s
}

func (s *SearchPunishRecordsResponse) SetStatusCode(v int32) *SearchPunishRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchPunishRecordsResponse) SetBody(v *SearchPunishRecordsResponseBody) *SearchPunishRecordsResponse {
	s.Body = v
	return s
}

type SearchPunishRequestRequest struct {
	AntiStatuses   []*string `json:"AntiStatuses,omitempty" xml:"AntiStatuses,omitempty" type:"Repeated"`
	BussinessCodes []*string `json:"BussinessCodes,omitempty" xml:"BussinessCodes,omitempty" type:"Repeated"`
	Class          *string   `json:"Class,omitempty" xml:"Class,omitempty"`
	EndDate        *int64    `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	EventCodes     []*string `json:"EventCodes,omitempty" xml:"EventCodes,omitempty" type:"Repeated"`
	ExtRequestId   *string   `json:"ExtRequestId,omitempty" xml:"ExtRequestId,omitempty"`
	Id             *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string   `json:"IdType,omitempty" xml:"IdType,omitempty"`
	InstanceId     *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Page           *int64    `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize       *int64    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PunishDomain   *string   `json:"PunishDomain,omitempty" xml:"PunishDomain,omitempty"`
	PunishIp       *string   `json:"PunishIp,omitempty" xml:"PunishIp,omitempty"`
	PunishStatuses []*string `json:"PunishStatuses,omitempty" xml:"PunishStatuses,omitempty" type:"Repeated"`
	PunishUrl      *string   `json:"PunishUrl,omitempty" xml:"PunishUrl,omitempty"`
	PunishUrlFull  *string   `json:"PunishUrlFull,omitempty" xml:"PunishUrlFull,omitempty"`
	SourceCodes    []*string `json:"SourceCodes,omitempty" xml:"SourceCodes,omitempty" type:"Repeated"`
	StartDate      *int64    `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	UserIds        []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s SearchPunishRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchPunishRequestRequest) GoString() string {
	return s.String()
}

func (s *SearchPunishRequestRequest) SetAntiStatuses(v []*string) *SearchPunishRequestRequest {
	s.AntiStatuses = v
	return s
}

func (s *SearchPunishRequestRequest) SetBussinessCodes(v []*string) *SearchPunishRequestRequest {
	s.BussinessCodes = v
	return s
}

func (s *SearchPunishRequestRequest) SetClass(v string) *SearchPunishRequestRequest {
	s.Class = &v
	return s
}

func (s *SearchPunishRequestRequest) SetEndDate(v int64) *SearchPunishRequestRequest {
	s.EndDate = &v
	return s
}

func (s *SearchPunishRequestRequest) SetEventCodes(v []*string) *SearchPunishRequestRequest {
	s.EventCodes = v
	return s
}

func (s *SearchPunishRequestRequest) SetExtRequestId(v string) *SearchPunishRequestRequest {
	s.ExtRequestId = &v
	return s
}

func (s *SearchPunishRequestRequest) SetId(v int64) *SearchPunishRequestRequest {
	s.Id = &v
	return s
}

func (s *SearchPunishRequestRequest) SetIdType(v string) *SearchPunishRequestRequest {
	s.IdType = &v
	return s
}

func (s *SearchPunishRequestRequest) SetInstanceId(v string) *SearchPunishRequestRequest {
	s.InstanceId = &v
	return s
}

func (s *SearchPunishRequestRequest) SetPage(v int64) *SearchPunishRequestRequest {
	s.Page = &v
	return s
}

func (s *SearchPunishRequestRequest) SetPageSize(v int64) *SearchPunishRequestRequest {
	s.PageSize = &v
	return s
}

func (s *SearchPunishRequestRequest) SetPunishDomain(v string) *SearchPunishRequestRequest {
	s.PunishDomain = &v
	return s
}

func (s *SearchPunishRequestRequest) SetPunishIp(v string) *SearchPunishRequestRequest {
	s.PunishIp = &v
	return s
}

func (s *SearchPunishRequestRequest) SetPunishStatuses(v []*string) *SearchPunishRequestRequest {
	s.PunishStatuses = v
	return s
}

func (s *SearchPunishRequestRequest) SetPunishUrl(v string) *SearchPunishRequestRequest {
	s.PunishUrl = &v
	return s
}

func (s *SearchPunishRequestRequest) SetPunishUrlFull(v string) *SearchPunishRequestRequest {
	s.PunishUrlFull = &v
	return s
}

func (s *SearchPunishRequestRequest) SetSourceCodes(v []*string) *SearchPunishRequestRequest {
	s.SourceCodes = v
	return s
}

func (s *SearchPunishRequestRequest) SetStartDate(v int64) *SearchPunishRequestRequest {
	s.StartDate = &v
	return s
}

func (s *SearchPunishRequestRequest) SetUserIds(v []*string) *SearchPunishRequestRequest {
	s.UserIds = v
	return s
}

type SearchPunishRequestShrinkRequest struct {
	AntiStatusesShrink   *string `json:"AntiStatuses,omitempty" xml:"AntiStatuses,omitempty"`
	BussinessCodesShrink *string `json:"BussinessCodes,omitempty" xml:"BussinessCodes,omitempty"`
	Class                *string `json:"Class,omitempty" xml:"Class,omitempty"`
	EndDate              *int64  `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	EventCodesShrink     *string `json:"EventCodes,omitempty" xml:"EventCodes,omitempty"`
	ExtRequestId         *string `json:"ExtRequestId,omitempty" xml:"ExtRequestId,omitempty"`
	Id                   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType               *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Page                 *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PunishDomain         *string `json:"PunishDomain,omitempty" xml:"PunishDomain,omitempty"`
	PunishIp             *string `json:"PunishIp,omitempty" xml:"PunishIp,omitempty"`
	PunishStatusesShrink *string `json:"PunishStatuses,omitempty" xml:"PunishStatuses,omitempty"`
	PunishUrl            *string `json:"PunishUrl,omitempty" xml:"PunishUrl,omitempty"`
	PunishUrlFull        *string `json:"PunishUrlFull,omitempty" xml:"PunishUrlFull,omitempty"`
	SourceCodesShrink    *string `json:"SourceCodes,omitempty" xml:"SourceCodes,omitempty"`
	StartDate            *int64  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	UserIdsShrink        *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s SearchPunishRequestShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchPunishRequestShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchPunishRequestShrinkRequest) SetAntiStatusesShrink(v string) *SearchPunishRequestShrinkRequest {
	s.AntiStatusesShrink = &v
	return s
}

func (s *SearchPunishRequestShrinkRequest) SetBussinessCodesShrink(v string) *SearchPunishRequestShrinkRequest {
	s.BussinessCodesShrink = &v
	return s
}

func (s *SearchPunishRequestShrinkRequest) SetClass(v string) *SearchPunishRequestShrinkRequest {
	s.Class = &v
	return s
}

func (s *SearchPunishRequestShrinkRequest) SetEndDate(v int64) *SearchPunishRequestShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *SearchPunishRequestShrinkRequest) SetEventCodesShrink(v string) *SearchPunishRequestShrinkRequest {
	s.EventCodesShrink = &v
	return s
}

func (s *SearchPunishRequestShrinkRequest) SetExtRequestId(v string) *SearchPunishRequestShrinkRequest {
	s.ExtRequestId = &v
	return s
}

func (s *SearchPunishRequestShrinkRequest) SetId(v int64) *SearchPunishRequestShrinkRequest {
	s.Id = &v
	return s
}

func (s *SearchPunishRequestShrinkRequest) SetIdType(v string) *SearchPunishRequestShrinkRequest {
	s.IdType = &v
	return s
}

func (s *SearchPunishRequestShrinkRequest) SetInstanceId(v string) *SearchPunishRequestShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *SearchPunishRequestShrinkRequest) SetPage(v int64) *SearchPunishRequestShrinkRequest {
	s.Page = &v
	return s
}

func (s *SearchPunishRequestShrinkRequest) SetPageSize(v int64) *SearchPunishRequestShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *SearchPunishRequestShrinkRequest) SetPunishDomain(v string) *SearchPunishRequestShrinkRequest {
	s.PunishDomain = &v
	return s
}

func (s *SearchPunishRequestShrinkRequest) SetPunishIp(v string) *SearchPunishRequestShrinkRequest {
	s.PunishIp = &v
	return s
}

func (s *SearchPunishRequestShrinkRequest) SetPunishStatusesShrink(v string) *SearchPunishRequestShrinkRequest {
	s.PunishStatusesShrink = &v
	return s
}

func (s *SearchPunishRequestShrinkRequest) SetPunishUrl(v string) *SearchPunishRequestShrinkRequest {
	s.PunishUrl = &v
	return s
}

func (s *SearchPunishRequestShrinkRequest) SetPunishUrlFull(v string) *SearchPunishRequestShrinkRequest {
	s.PunishUrlFull = &v
	return s
}

func (s *SearchPunishRequestShrinkRequest) SetSourceCodesShrink(v string) *SearchPunishRequestShrinkRequest {
	s.SourceCodesShrink = &v
	return s
}

func (s *SearchPunishRequestShrinkRequest) SetStartDate(v int64) *SearchPunishRequestShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *SearchPunishRequestShrinkRequest) SetUserIdsShrink(v string) *SearchPunishRequestShrinkRequest {
	s.UserIdsShrink = &v
	return s
}

type SearchPunishRequestResponseBody struct {
	Class      *string                                    `json:"Class,omitempty" xml:"Class,omitempty"`
	Code       *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Count      *int64                                     `json:"Count,omitempty" xml:"Count,omitempty"`
	DataList   []*SearchPunishRequestResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	Message    *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	Success    *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	ViewCount  *int64                                     `json:"ViewCount,omitempty" xml:"ViewCount,omitempty"`
}

func (s SearchPunishRequestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchPunishRequestResponseBody) GoString() string {
	return s.String()
}

func (s *SearchPunishRequestResponseBody) SetClass(v string) *SearchPunishRequestResponseBody {
	s.Class = &v
	return s
}

func (s *SearchPunishRequestResponseBody) SetCode(v string) *SearchPunishRequestResponseBody {
	s.Code = &v
	return s
}

func (s *SearchPunishRequestResponseBody) SetCount(v int64) *SearchPunishRequestResponseBody {
	s.Count = &v
	return s
}

func (s *SearchPunishRequestResponseBody) SetDataList(v []*SearchPunishRequestResponseBodyDataList) *SearchPunishRequestResponseBody {
	s.DataList = v
	return s
}

func (s *SearchPunishRequestResponseBody) SetMessage(v string) *SearchPunishRequestResponseBody {
	s.Message = &v
	return s
}

func (s *SearchPunishRequestResponseBody) SetSuccess(v bool) *SearchPunishRequestResponseBody {
	s.Success = &v
	return s
}

func (s *SearchPunishRequestResponseBody) SetTotalCount(v int64) *SearchPunishRequestResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchPunishRequestResponseBody) SetViewCount(v int64) *SearchPunishRequestResponseBody {
	s.ViewCount = &v
	return s
}

type SearchPunishRequestResponseBodyDataList struct {
	AntiPunishRespTime *string `json:"AntiPunishRespTime,omitempty" xml:"AntiPunishRespTime,omitempty"`
	AntiPunishTime     *string `json:"AntiPunishTime,omitempty" xml:"AntiPunishTime,omitempty"`
	AntiResult         *string `json:"AntiResult,omitempty" xml:"AntiResult,omitempty"`
	AntiStatus         *string `json:"AntiStatus,omitempty" xml:"AntiStatus,omitempty"`
	BussinessCode      *string `json:"BussinessCode,omitempty" xml:"BussinessCode,omitempty"`
	CaseCode           *string `json:"CaseCode,omitempty" xml:"CaseCode,omitempty"`
	CaseExtendCode     *string `json:"CaseExtendCode,omitempty" xml:"CaseExtendCode,omitempty"`
	CaseSubCode        *string `json:"CaseSubCode,omitempty" xml:"CaseSubCode,omitempty"`
	Class              *string `json:"Class,omitempty" xml:"Class,omitempty"`
	Creator            *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Deleted            *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	EventCode          *string `json:"EventCode,omitempty" xml:"EventCode,omitempty"`
	ExpectedRemoveTime *string `json:"ExpectedRemoveTime,omitempty" xml:"ExpectedRemoveTime,omitempty"`
	ExtRequestId       *string `json:"ExtRequestId,omitempty" xml:"ExtRequestId,omitempty"`
	GmtCreated         *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified        *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType             *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IpString           *string `json:"IpString,omitempty" xml:"IpString,omitempty"`
	Modifier           *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	PunishDomain       *string `json:"PunishDomain,omitempty" xml:"PunishDomain,omitempty"`
	PunishIp           *string `json:"PunishIp,omitempty" xml:"PunishIp,omitempty"`
	PunishRequest      *string `json:"PunishRequest,omitempty" xml:"PunishRequest,omitempty"`
	PunishRespTime     *string `json:"PunishRespTime,omitempty" xml:"PunishRespTime,omitempty"`
	PunishResult       *string `json:"PunishResult,omitempty" xml:"PunishResult,omitempty"`
	PunishStatus       *string `json:"PunishStatus,omitempty" xml:"PunishStatus,omitempty"`
	PunishTime         *string `json:"PunishTime,omitempty" xml:"PunishTime,omitempty"`
	PunishUrl          *string `json:"PunishUrl,omitempty" xml:"PunishUrl,omitempty"`
	Reason             *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	SourceCode         *string `json:"SourceCode,omitempty" xml:"SourceCode,omitempty"`
	UserId             *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchPunishRequestResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s SearchPunishRequestResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *SearchPunishRequestResponseBodyDataList) SetAntiPunishRespTime(v string) *SearchPunishRequestResponseBodyDataList {
	s.AntiPunishRespTime = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetAntiPunishTime(v string) *SearchPunishRequestResponseBodyDataList {
	s.AntiPunishTime = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetAntiResult(v string) *SearchPunishRequestResponseBodyDataList {
	s.AntiResult = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetAntiStatus(v string) *SearchPunishRequestResponseBodyDataList {
	s.AntiStatus = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetBussinessCode(v string) *SearchPunishRequestResponseBodyDataList {
	s.BussinessCode = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetCaseCode(v string) *SearchPunishRequestResponseBodyDataList {
	s.CaseCode = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetCaseExtendCode(v string) *SearchPunishRequestResponseBodyDataList {
	s.CaseExtendCode = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetCaseSubCode(v string) *SearchPunishRequestResponseBodyDataList {
	s.CaseSubCode = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetClass(v string) *SearchPunishRequestResponseBodyDataList {
	s.Class = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetCreator(v string) *SearchPunishRequestResponseBodyDataList {
	s.Creator = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetDeleted(v bool) *SearchPunishRequestResponseBodyDataList {
	s.Deleted = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetEventCode(v string) *SearchPunishRequestResponseBodyDataList {
	s.EventCode = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetExpectedRemoveTime(v string) *SearchPunishRequestResponseBodyDataList {
	s.ExpectedRemoveTime = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetExtRequestId(v string) *SearchPunishRequestResponseBodyDataList {
	s.ExtRequestId = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetGmtCreated(v string) *SearchPunishRequestResponseBodyDataList {
	s.GmtCreated = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetGmtModified(v string) *SearchPunishRequestResponseBodyDataList {
	s.GmtModified = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetId(v int64) *SearchPunishRequestResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetIdType(v string) *SearchPunishRequestResponseBodyDataList {
	s.IdType = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetInstanceId(v string) *SearchPunishRequestResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetIpString(v string) *SearchPunishRequestResponseBodyDataList {
	s.IpString = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetModifier(v string) *SearchPunishRequestResponseBodyDataList {
	s.Modifier = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetPunishDomain(v string) *SearchPunishRequestResponseBodyDataList {
	s.PunishDomain = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetPunishIp(v string) *SearchPunishRequestResponseBodyDataList {
	s.PunishIp = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetPunishRequest(v string) *SearchPunishRequestResponseBodyDataList {
	s.PunishRequest = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetPunishRespTime(v string) *SearchPunishRequestResponseBodyDataList {
	s.PunishRespTime = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetPunishResult(v string) *SearchPunishRequestResponseBodyDataList {
	s.PunishResult = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetPunishStatus(v string) *SearchPunishRequestResponseBodyDataList {
	s.PunishStatus = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetPunishTime(v string) *SearchPunishRequestResponseBodyDataList {
	s.PunishTime = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetPunishUrl(v string) *SearchPunishRequestResponseBodyDataList {
	s.PunishUrl = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetReason(v string) *SearchPunishRequestResponseBodyDataList {
	s.Reason = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetSourceCode(v string) *SearchPunishRequestResponseBodyDataList {
	s.SourceCode = &v
	return s
}

func (s *SearchPunishRequestResponseBodyDataList) SetUserId(v string) *SearchPunishRequestResponseBodyDataList {
	s.UserId = &v
	return s
}

type SearchPunishRequestResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchPunishRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchPunishRequestResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchPunishRequestResponse) GoString() string {
	return s.String()
}

func (s *SearchPunishRequestResponse) SetHeaders(v map[string]*string) *SearchPunishRequestResponse {
	s.Headers = v
	return s
}

func (s *SearchPunishRequestResponse) SetStatusCode(v int32) *SearchPunishRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchPunishRequestResponse) SetBody(v *SearchPunishRequestResponseBody) *SearchPunishRequestResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("buss"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 处罚请求异步接口回调
//
// @param tmpReq - BusinessResultServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BusinessResultServiceResponse
func (client *Client) BusinessResultServiceWithOptions(tmpReq *BusinessResultServiceRequest, runtime *util.RuntimeOptions) (_result *BusinessResultServiceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BusinessResultServiceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Result)) {
		request.ResultShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Result, tea.String("Result"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BusinessResultService"),
		Version:     tea.String("2022-08-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BusinessResultServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 处罚请求异步接口回调
//
// @param request - BusinessResultServiceRequest
//
// @return BusinessResultServiceResponse
func (client *Client) BusinessResultService(request *BusinessResultServiceRequest) (_result *BusinessResultServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BusinessResultServiceResponse{}
	_body, _err := client.BusinessResultServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 协查中心查询任务接口
//
// @param request - CreateUserInvestigationInfoQueryTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserInvestigationInfoQueryTaskResponse
func (client *Client) CreateUserInvestigationInfoQueryTaskWithOptions(request *CreateUserInvestigationInfoQueryTaskRequest, runtime *util.RuntimeOptions) (_result *CreateUserInvestigationInfoQueryTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUserInvestigationInfoQueryTask"),
		Version:     tea.String("2022-08-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserInvestigationInfoQueryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 协查中心查询任务接口
//
// @param request - CreateUserInvestigationInfoQueryTaskRequest
//
// @return CreateUserInvestigationInfoQueryTaskResponse
func (client *Client) CreateUserInvestigationInfoQueryTask(request *CreateUserInvestigationInfoQueryTaskRequest) (_result *CreateUserInvestigationInfoQueryTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserInvestigationInfoQueryTaskResponse{}
	_body, _err := client.CreateUserInvestigationInfoQueryTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 反查资源
//
// @param tmpReq - FindInstanceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FindInstanceInfoResponse
func (client *Client) FindInstanceInfoWithOptions(tmpReq *FindInstanceInfoRequest, runtime *util.RuntimeOptions) (_result *FindInstanceInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &FindInstanceInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Extras)) {
		request.ExtrasShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extras, tea.String("extras"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FindInstanceInfo"),
		Version:     tea.String("2022-08-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FindInstanceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 反查资源
//
// @param request - FindInstanceInfoRequest
//
// @return FindInstanceInfoResponse
func (client *Client) FindInstanceInfo(request *FindInstanceInfoRequest) (_result *FindInstanceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FindInstanceInfoResponse{}
	_body, _err := client.FindInstanceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据用户id查询对应产品下全部可用资产信息接口
//
// @param request - FindUserAvailbleResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FindUserAvailbleResourcesResponse
func (client *Client) FindUserAvailbleResourcesWithOptions(request *FindUserAvailbleResourcesRequest, runtime *util.RuntimeOptions) (_result *FindUserAvailbleResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FindUserAvailbleResources"),
		Version:     tea.String("2022-08-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FindUserAvailbleResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据用户id查询对应产品下全部可用资产信息接口
//
// @param request - FindUserAvailbleResourcesRequest
//
// @return FindUserAvailbleResourcesResponse
func (client *Client) FindUserAvailbleResources(request *FindUserAvailbleResourcesRequest) (_result *FindUserAvailbleResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FindUserAvailbleResourcesResponse{}
	_body, _err := client.FindUserAvailbleResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 处罚资源搜索
//
// @param tmpReq - PunishResourceSearchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PunishResourceSearchResponse
func (client *Client) PunishResourceSearchWithOptions(tmpReq *PunishResourceSearchRequest, runtime *util.RuntimeOptions) (_result *PunishResourceSearchResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &PunishResourceSearchShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ActionCodes)) {
		request.ActionCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ActionCodes, tea.String("ActionCodes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.BussinessCodes)) {
		request.BussinessCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BussinessCodes, tea.String("BussinessCodes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SourceCodes)) {
		request.SourceCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceCodes, tea.String("SourceCodes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserIds)) {
		request.UserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIds, tea.String("UserIds"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PunishResourceSearch"),
		Version:     tea.String("2022-08-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PunishResourceSearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 处罚资源搜索
//
// @param request - PunishResourceSearchRequest
//
// @return PunishResourceSearchResponse
func (client *Client) PunishResourceSearch(request *PunishResourceSearchRequest) (_result *PunishResourceSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PunishResourceSearchResponse{}
	_body, _err := client.PunishResourceSearchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 风险事件同步
//
// @param request - RiskEventSyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RiskEventSyncResponse
func (client *Client) RiskEventSyncWithOptions(request *RiskEventSyncRequest, runtime *util.RuntimeOptions) (_result *RiskEventSyncResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Deleted)) {
		body["Deleted"] = request.Deleted
	}

	if !tea.BoolValue(util.IsUnset(request.DiscoveryTime)) {
		body["DiscoveryTime"] = request.DiscoveryTime
	}

	if !tea.BoolValue(util.IsUnset(request.EventName)) {
		body["EventName"] = request.EventName
	}

	if !tea.BoolValue(util.IsUnset(request.EventNumber)) {
		body["EventNumber"] = request.EventNumber
	}

	if !tea.BoolValue(util.IsUnset(request.RelevanceBu)) {
		body["RelevanceBu"] = request.RelevanceBu
	}

	if !tea.BoolValue(util.IsUnset(request.RiskDetail)) {
		body["RiskDetail"] = request.RiskDetail
	}

	if !tea.BoolValue(util.IsUnset(request.RiskEffectType)) {
		body["RiskEffectType"] = request.RiskEffectType
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevel)) {
		body["RiskLevel"] = request.RiskLevel
	}

	if !tea.BoolValue(util.IsUnset(request.RiskType)) {
		body["RiskType"] = request.RiskType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.Submitter)) {
		body["Submitter"] = request.Submitter
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RiskEventSync"),
		Version:     tea.String("2022-08-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RiskEventSyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 风险事件同步
//
// @param request - RiskEventSyncRequest
//
// @return RiskEventSyncResponse
func (client *Client) RiskEventSync(request *RiskEventSyncRequest) (_result *RiskEventSyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RiskEventSyncResponse{}
	_body, _err := client.RiskEventSyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 管控事件查询
//
// @param tmpReq - SearchPunishEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchPunishEventsResponse
func (client *Client) SearchPunishEventsWithOptions(tmpReq *SearchPunishEventsRequest, runtime *util.RuntimeOptions) (_result *SearchPunishEventsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SearchPunishEventsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.BussinessCodes)) {
		request.BussinessCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BussinessCodes, tea.String("BussinessCodes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.CaseCodes)) {
		request.CaseCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CaseCodes, tea.String("CaseCodes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.EventCodes)) {
		request.EventCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventCodes, tea.String("EventCodes"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.BussinessCodesShrink)) {
		query["BussinessCodes"] = request.BussinessCodesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CaseCodesShrink)) {
		query["CaseCodes"] = request.CaseCodesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.EventCodesShrink)) {
		query["EventCodes"] = request.EventCodesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchPunishEvents"),
		Version:     tea.String("2022-08-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchPunishEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 管控事件查询
//
// @param request - SearchPunishEventsRequest
//
// @return SearchPunishEventsResponse
func (client *Client) SearchPunishEvents(request *SearchPunishEventsRequest) (_result *SearchPunishEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchPunishEventsResponse{}
	_body, _err := client.SearchPunishEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 管控事件查询
//
// @param tmpReq - SearchPunishRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchPunishRecordsResponse
func (client *Client) SearchPunishRecordsWithOptions(tmpReq *SearchPunishRecordsRequest, runtime *util.RuntimeOptions) (_result *SearchPunishRecordsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SearchPunishRecordsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ActionCodes)) {
		request.ActionCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ActionCodes, tea.String("ActionCodes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.CaseCodes)) {
		request.CaseCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CaseCodes, tea.String("CaseCodes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.EventCodes)) {
		request.EventCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventCodes, tea.String("EventCodes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.PunishStatus)) {
		request.PunishStatusShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PunishStatus, tea.String("PunishStatus"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SourceCodes)) {
		request.SourceCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceCodes, tea.String("SourceCodes"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionCodesShrink)) {
		query["ActionCodes"] = request.ActionCodesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.BussinessCodes)) {
		query["BussinessCodes"] = request.BussinessCodes
	}

	if !tea.BoolValue(util.IsUnset(request.CaseCodesShrink)) {
		query["CaseCodes"] = request.CaseCodesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EventCodesShrink)) {
		query["EventCodes"] = request.EventCodesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PunishStatusShrink)) {
		query["PunishStatus"] = request.PunishStatusShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCodesShrink)) {
		query["SourceCodes"] = request.SourceCodesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	if !tea.BoolValue(util.IsUnset(request.UrlFuzzy)) {
		query["UrlFuzzy"] = request.UrlFuzzy
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchPunishRecords"),
		Version:     tea.String("2022-08-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchPunishRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 管控事件查询
//
// @param request - SearchPunishRecordsRequest
//
// @return SearchPunishRecordsResponse
func (client *Client) SearchPunishRecords(request *SearchPunishRecordsRequest) (_result *SearchPunishRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchPunishRecordsResponse{}
	_body, _err := client.SearchPunishRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 处罚记录查询
//
// @param tmpReq - SearchPunishRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchPunishRequestResponse
func (client *Client) SearchPunishRequestWithOptions(tmpReq *SearchPunishRequestRequest, runtime *util.RuntimeOptions) (_result *SearchPunishRequestResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SearchPunishRequestShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AntiStatuses)) {
		request.AntiStatusesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AntiStatuses, tea.String("AntiStatuses"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.BussinessCodes)) {
		request.BussinessCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BussinessCodes, tea.String("BussinessCodes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.EventCodes)) {
		request.EventCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventCodes, tea.String("EventCodes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.PunishStatuses)) {
		request.PunishStatusesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PunishStatuses, tea.String("PunishStatuses"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SourceCodes)) {
		request.SourceCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceCodes, tea.String("SourceCodes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserIds)) {
		request.UserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIds, tea.String("UserIds"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchPunishRequest"),
		Version:     tea.String("2022-08-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchPunishRequestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 处罚记录查询
//
// @param request - SearchPunishRequestRequest
//
// @return SearchPunishRequestResponse
func (client *Client) SearchPunishRequest(request *SearchPunishRequestRequest) (_result *SearchPunishRequestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchPunishRequestResponse{}
	_body, _err := client.SearchPunishRequestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
