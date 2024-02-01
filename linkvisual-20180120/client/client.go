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

type AddEventRecordPlanDeviceRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	PlanId        *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	StreamType    *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s AddEventRecordPlanDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddEventRecordPlanDeviceRequest) GoString() string {
	return s.String()
}

func (s *AddEventRecordPlanDeviceRequest) SetDeviceName(v string) *AddEventRecordPlanDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *AddEventRecordPlanDeviceRequest) SetIotId(v string) *AddEventRecordPlanDeviceRequest {
	s.IotId = &v
	return s
}

func (s *AddEventRecordPlanDeviceRequest) SetIotInstanceId(v string) *AddEventRecordPlanDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *AddEventRecordPlanDeviceRequest) SetPlanId(v string) *AddEventRecordPlanDeviceRequest {
	s.PlanId = &v
	return s
}

func (s *AddEventRecordPlanDeviceRequest) SetProductKey(v string) *AddEventRecordPlanDeviceRequest {
	s.ProductKey = &v
	return s
}

func (s *AddEventRecordPlanDeviceRequest) SetStreamType(v int32) *AddEventRecordPlanDeviceRequest {
	s.StreamType = &v
	return s
}

type AddEventRecordPlanDeviceResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddEventRecordPlanDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddEventRecordPlanDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *AddEventRecordPlanDeviceResponseBody) SetCode(v string) *AddEventRecordPlanDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *AddEventRecordPlanDeviceResponseBody) SetErrorMessage(v string) *AddEventRecordPlanDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddEventRecordPlanDeviceResponseBody) SetRequestId(v string) *AddEventRecordPlanDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddEventRecordPlanDeviceResponseBody) SetSuccess(v bool) *AddEventRecordPlanDeviceResponseBody {
	s.Success = &v
	return s
}

type AddEventRecordPlanDeviceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddEventRecordPlanDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddEventRecordPlanDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddEventRecordPlanDeviceResponse) GoString() string {
	return s.String()
}

func (s *AddEventRecordPlanDeviceResponse) SetHeaders(v map[string]*string) *AddEventRecordPlanDeviceResponse {
	s.Headers = v
	return s
}

func (s *AddEventRecordPlanDeviceResponse) SetStatusCode(v int32) *AddEventRecordPlanDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddEventRecordPlanDeviceResponse) SetBody(v *AddEventRecordPlanDeviceResponseBody) *AddEventRecordPlanDeviceResponse {
	s.Body = v
	return s
}

type AddFaceDeviceGroupRequest struct {
	DeviceGroupName *string `json:"DeviceGroupName,omitempty" xml:"DeviceGroupName,omitempty"`
	IsolationId     *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
}

func (s AddFaceDeviceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFaceDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *AddFaceDeviceGroupRequest) SetDeviceGroupName(v string) *AddFaceDeviceGroupRequest {
	s.DeviceGroupName = &v
	return s
}

func (s *AddFaceDeviceGroupRequest) SetIsolationId(v string) *AddFaceDeviceGroupRequest {
	s.IsolationId = &v
	return s
}

type AddFaceDeviceGroupResponseBody struct {
	Code         *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *AddFaceDeviceGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddFaceDeviceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFaceDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceDeviceGroupResponseBody) SetCode(v string) *AddFaceDeviceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *AddFaceDeviceGroupResponseBody) SetData(v *AddFaceDeviceGroupResponseBodyData) *AddFaceDeviceGroupResponseBody {
	s.Data = v
	return s
}

func (s *AddFaceDeviceGroupResponseBody) SetErrorMessage(v string) *AddFaceDeviceGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddFaceDeviceGroupResponseBody) SetRequestId(v string) *AddFaceDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFaceDeviceGroupResponseBody) SetSuccess(v bool) *AddFaceDeviceGroupResponseBody {
	s.Success = &v
	return s
}

type AddFaceDeviceGroupResponseBodyData struct {
	DeviceGroupId   *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	DeviceGroupName *string `json:"DeviceGroupName,omitempty" xml:"DeviceGroupName,omitempty"`
	ModifiedTime    *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
}

func (s AddFaceDeviceGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddFaceDeviceGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddFaceDeviceGroupResponseBodyData) SetDeviceGroupId(v string) *AddFaceDeviceGroupResponseBodyData {
	s.DeviceGroupId = &v
	return s
}

func (s *AddFaceDeviceGroupResponseBodyData) SetDeviceGroupName(v string) *AddFaceDeviceGroupResponseBodyData {
	s.DeviceGroupName = &v
	return s
}

func (s *AddFaceDeviceGroupResponseBodyData) SetModifiedTime(v string) *AddFaceDeviceGroupResponseBodyData {
	s.ModifiedTime = &v
	return s
}

type AddFaceDeviceGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFaceDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFaceDeviceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddFaceDeviceGroupResponse) GoString() string {
	return s.String()
}

func (s *AddFaceDeviceGroupResponse) SetHeaders(v map[string]*string) *AddFaceDeviceGroupResponse {
	s.Headers = v
	return s
}

func (s *AddFaceDeviceGroupResponse) SetStatusCode(v int32) *AddFaceDeviceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFaceDeviceGroupResponse) SetBody(v *AddFaceDeviceGroupResponseBody) *AddFaceDeviceGroupResponse {
	s.Body = v
	return s
}

type AddFaceDeviceToDeviceGroupRequest struct {
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	IsolationId   *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s AddFaceDeviceToDeviceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFaceDeviceToDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *AddFaceDeviceToDeviceGroupRequest) SetDeviceGroupId(v string) *AddFaceDeviceToDeviceGroupRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *AddFaceDeviceToDeviceGroupRequest) SetDeviceName(v string) *AddFaceDeviceToDeviceGroupRequest {
	s.DeviceName = &v
	return s
}

func (s *AddFaceDeviceToDeviceGroupRequest) SetIotInstanceId(v string) *AddFaceDeviceToDeviceGroupRequest {
	s.IotInstanceId = &v
	return s
}

func (s *AddFaceDeviceToDeviceGroupRequest) SetIsolationId(v string) *AddFaceDeviceToDeviceGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *AddFaceDeviceToDeviceGroupRequest) SetProductKey(v string) *AddFaceDeviceToDeviceGroupRequest {
	s.ProductKey = &v
	return s
}

type AddFaceDeviceToDeviceGroupResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddFaceDeviceToDeviceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFaceDeviceToDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceDeviceToDeviceGroupResponseBody) SetCode(v string) *AddFaceDeviceToDeviceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *AddFaceDeviceToDeviceGroupResponseBody) SetErrorMessage(v string) *AddFaceDeviceToDeviceGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddFaceDeviceToDeviceGroupResponseBody) SetRequestId(v string) *AddFaceDeviceToDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFaceDeviceToDeviceGroupResponseBody) SetSuccess(v bool) *AddFaceDeviceToDeviceGroupResponseBody {
	s.Success = &v
	return s
}

type AddFaceDeviceToDeviceGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFaceDeviceToDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFaceDeviceToDeviceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddFaceDeviceToDeviceGroupResponse) GoString() string {
	return s.String()
}

func (s *AddFaceDeviceToDeviceGroupResponse) SetHeaders(v map[string]*string) *AddFaceDeviceToDeviceGroupResponse {
	s.Headers = v
	return s
}

func (s *AddFaceDeviceToDeviceGroupResponse) SetStatusCode(v int32) *AddFaceDeviceToDeviceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFaceDeviceToDeviceGroupResponse) SetBody(v *AddFaceDeviceToDeviceGroupResponseBody) *AddFaceDeviceToDeviceGroupResponse {
	s.Body = v
	return s
}

type AddFaceUserRequest struct {
	CustomUserId *string `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	FacePicUrl   *string `json:"FacePicUrl,omitempty" xml:"FacePicUrl,omitempty"`
	IsolationId  *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Params       *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s AddFaceUserRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserRequest) GoString() string {
	return s.String()
}

func (s *AddFaceUserRequest) SetCustomUserId(v string) *AddFaceUserRequest {
	s.CustomUserId = &v
	return s
}

func (s *AddFaceUserRequest) SetFacePicUrl(v string) *AddFaceUserRequest {
	s.FacePicUrl = &v
	return s
}

func (s *AddFaceUserRequest) SetIsolationId(v string) *AddFaceUserRequest {
	s.IsolationId = &v
	return s
}

func (s *AddFaceUserRequest) SetName(v string) *AddFaceUserRequest {
	s.Name = &v
	return s
}

func (s *AddFaceUserRequest) SetParams(v string) *AddFaceUserRequest {
	s.Params = &v
	return s
}

type AddFaceUserResponseBody struct {
	Code         *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *AddFaceUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddFaceUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceUserResponseBody) SetCode(v string) *AddFaceUserResponseBody {
	s.Code = &v
	return s
}

func (s *AddFaceUserResponseBody) SetData(v *AddFaceUserResponseBodyData) *AddFaceUserResponseBody {
	s.Data = v
	return s
}

func (s *AddFaceUserResponseBody) SetErrorMessage(v string) *AddFaceUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddFaceUserResponseBody) SetRequestId(v string) *AddFaceUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFaceUserResponseBody) SetSuccess(v bool) *AddFaceUserResponseBody {
	s.Success = &v
	return s
}

type AddFaceUserResponseBodyData struct {
	CustomUserId *string `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Params       *string `json:"Params,omitempty" xml:"Params,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddFaceUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddFaceUserResponseBodyData) SetCustomUserId(v string) *AddFaceUserResponseBodyData {
	s.CustomUserId = &v
	return s
}

func (s *AddFaceUserResponseBodyData) SetName(v string) *AddFaceUserResponseBodyData {
	s.Name = &v
	return s
}

func (s *AddFaceUserResponseBodyData) SetParams(v string) *AddFaceUserResponseBodyData {
	s.Params = &v
	return s
}

func (s *AddFaceUserResponseBodyData) SetUserId(v string) *AddFaceUserResponseBodyData {
	s.UserId = &v
	return s
}

type AddFaceUserResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFaceUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFaceUserResponse) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserResponse) GoString() string {
	return s.String()
}

func (s *AddFaceUserResponse) SetHeaders(v map[string]*string) *AddFaceUserResponse {
	s.Headers = v
	return s
}

func (s *AddFaceUserResponse) SetStatusCode(v int32) *AddFaceUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFaceUserResponse) SetBody(v *AddFaceUserResponseBody) *AddFaceUserResponse {
	s.Body = v
	return s
}

type AddFaceUserGroupRequest struct {
	IsolationId   *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s AddFaceUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserGroupRequest) GoString() string {
	return s.String()
}

func (s *AddFaceUserGroupRequest) SetIsolationId(v string) *AddFaceUserGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *AddFaceUserGroupRequest) SetUserGroupName(v string) *AddFaceUserGroupRequest {
	s.UserGroupName = &v
	return s
}

type AddFaceUserGroupResponseBody struct {
	Code         *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *AddFaceUserGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddFaceUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceUserGroupResponseBody) SetCode(v string) *AddFaceUserGroupResponseBody {
	s.Code = &v
	return s
}

func (s *AddFaceUserGroupResponseBody) SetData(v *AddFaceUserGroupResponseBodyData) *AddFaceUserGroupResponseBody {
	s.Data = v
	return s
}

func (s *AddFaceUserGroupResponseBody) SetErrorMessage(v string) *AddFaceUserGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddFaceUserGroupResponseBody) SetRequestId(v string) *AddFaceUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFaceUserGroupResponseBody) SetSuccess(v bool) *AddFaceUserGroupResponseBody {
	s.Success = &v
	return s
}

type AddFaceUserGroupResponseBodyData struct {
	ModifiedTime  *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	UserGroupId   *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s AddFaceUserGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddFaceUserGroupResponseBodyData) SetModifiedTime(v string) *AddFaceUserGroupResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *AddFaceUserGroupResponseBodyData) SetUserGroupId(v string) *AddFaceUserGroupResponseBodyData {
	s.UserGroupId = &v
	return s
}

func (s *AddFaceUserGroupResponseBodyData) SetUserGroupName(v string) *AddFaceUserGroupResponseBodyData {
	s.UserGroupName = &v
	return s
}

type AddFaceUserGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFaceUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFaceUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserGroupResponse) GoString() string {
	return s.String()
}

func (s *AddFaceUserGroupResponse) SetHeaders(v map[string]*string) *AddFaceUserGroupResponse {
	s.Headers = v
	return s
}

func (s *AddFaceUserGroupResponse) SetStatusCode(v int32) *AddFaceUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFaceUserGroupResponse) SetBody(v *AddFaceUserGroupResponseBody) *AddFaceUserGroupResponse {
	s.Body = v
	return s
}

type AddFaceUserGroupAndDeviceGroupRelationRequest struct {
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	IsolationId   *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	Relation      *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	UserGroupId   *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s AddFaceUserGroupAndDeviceGroupRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserGroupAndDeviceGroupRelationRequest) GoString() string {
	return s.String()
}

func (s *AddFaceUserGroupAndDeviceGroupRelationRequest) SetDeviceGroupId(v string) *AddFaceUserGroupAndDeviceGroupRelationRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationRequest) SetIotInstanceId(v string) *AddFaceUserGroupAndDeviceGroupRelationRequest {
	s.IotInstanceId = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationRequest) SetIsolationId(v string) *AddFaceUserGroupAndDeviceGroupRelationRequest {
	s.IsolationId = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationRequest) SetRelation(v string) *AddFaceUserGroupAndDeviceGroupRelationRequest {
	s.Relation = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationRequest) SetUserGroupId(v string) *AddFaceUserGroupAndDeviceGroupRelationRequest {
	s.UserGroupId = &v
	return s
}

type AddFaceUserGroupAndDeviceGroupRelationResponseBody struct {
	Code         *string                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *AddFaceUserGroupAndDeviceGroupRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddFaceUserGroupAndDeviceGroupRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserGroupAndDeviceGroupRelationResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBody) SetCode(v string) *AddFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.Code = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBody) SetData(v *AddFaceUserGroupAndDeviceGroupRelationResponseBodyData) *AddFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.Data = v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBody) SetErrorMessage(v string) *AddFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBody) SetRequestId(v string) *AddFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBody) SetSuccess(v bool) *AddFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.Success = &v
	return s
}

type AddFaceUserGroupAndDeviceGroupRelationResponseBodyData struct {
	ControlId    *string `json:"ControlId,omitempty" xml:"ControlId,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
}

func (s AddFaceUserGroupAndDeviceGroupRelationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserGroupAndDeviceGroupRelationResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBodyData) SetControlId(v string) *AddFaceUserGroupAndDeviceGroupRelationResponseBodyData {
	s.ControlId = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBodyData) SetModifiedTime(v string) *AddFaceUserGroupAndDeviceGroupRelationResponseBodyData {
	s.ModifiedTime = &v
	return s
}

type AddFaceUserGroupAndDeviceGroupRelationResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFaceUserGroupAndDeviceGroupRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFaceUserGroupAndDeviceGroupRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserGroupAndDeviceGroupRelationResponse) GoString() string {
	return s.String()
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponse) SetHeaders(v map[string]*string) *AddFaceUserGroupAndDeviceGroupRelationResponse {
	s.Headers = v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponse) SetStatusCode(v int32) *AddFaceUserGroupAndDeviceGroupRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponse) SetBody(v *AddFaceUserGroupAndDeviceGroupRelationResponseBody) *AddFaceUserGroupAndDeviceGroupRelationResponse {
	s.Body = v
	return s
}

type AddFaceUserPictureRequest struct {
	FacePicUrl  *string `json:"FacePicUrl,omitempty" xml:"FacePicUrl,omitempty"`
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddFaceUserPictureRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserPictureRequest) GoString() string {
	return s.String()
}

func (s *AddFaceUserPictureRequest) SetFacePicUrl(v string) *AddFaceUserPictureRequest {
	s.FacePicUrl = &v
	return s
}

func (s *AddFaceUserPictureRequest) SetIsolationId(v string) *AddFaceUserPictureRequest {
	s.IsolationId = &v
	return s
}

func (s *AddFaceUserPictureRequest) SetUserId(v string) *AddFaceUserPictureRequest {
	s.UserId = &v
	return s
}

type AddFaceUserPictureResponseBody struct {
	Code         *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddFaceUserPictureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserPictureResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceUserPictureResponseBody) SetCode(v string) *AddFaceUserPictureResponseBody {
	s.Code = &v
	return s
}

func (s *AddFaceUserPictureResponseBody) SetData(v map[string]interface{}) *AddFaceUserPictureResponseBody {
	s.Data = v
	return s
}

func (s *AddFaceUserPictureResponseBody) SetErrorMessage(v string) *AddFaceUserPictureResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddFaceUserPictureResponseBody) SetRequestId(v string) *AddFaceUserPictureResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFaceUserPictureResponseBody) SetSuccess(v bool) *AddFaceUserPictureResponseBody {
	s.Success = &v
	return s
}

type AddFaceUserPictureResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFaceUserPictureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFaceUserPictureResponse) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserPictureResponse) GoString() string {
	return s.String()
}

func (s *AddFaceUserPictureResponse) SetHeaders(v map[string]*string) *AddFaceUserPictureResponse {
	s.Headers = v
	return s
}

func (s *AddFaceUserPictureResponse) SetStatusCode(v int32) *AddFaceUserPictureResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFaceUserPictureResponse) SetBody(v *AddFaceUserPictureResponseBody) *AddFaceUserPictureResponse {
	s.Body = v
	return s
}

type AddFaceUserToUserGroupRequest struct {
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddFaceUserToUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserToUserGroupRequest) GoString() string {
	return s.String()
}

func (s *AddFaceUserToUserGroupRequest) SetIsolationId(v string) *AddFaceUserToUserGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *AddFaceUserToUserGroupRequest) SetUserGroupId(v string) *AddFaceUserToUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *AddFaceUserToUserGroupRequest) SetUserId(v string) *AddFaceUserToUserGroupRequest {
	s.UserId = &v
	return s
}

type AddFaceUserToUserGroupResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddFaceUserToUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserToUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceUserToUserGroupResponseBody) SetCode(v string) *AddFaceUserToUserGroupResponseBody {
	s.Code = &v
	return s
}

func (s *AddFaceUserToUserGroupResponseBody) SetErrorMessage(v string) *AddFaceUserToUserGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddFaceUserToUserGroupResponseBody) SetRequestId(v string) *AddFaceUserToUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFaceUserToUserGroupResponseBody) SetSuccess(v bool) *AddFaceUserToUserGroupResponseBody {
	s.Success = &v
	return s
}

type AddFaceUserToUserGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFaceUserToUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFaceUserToUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserToUserGroupResponse) GoString() string {
	return s.String()
}

func (s *AddFaceUserToUserGroupResponse) SetHeaders(v map[string]*string) *AddFaceUserToUserGroupResponse {
	s.Headers = v
	return s
}

func (s *AddFaceUserToUserGroupResponse) SetStatusCode(v int32) *AddFaceUserToUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFaceUserToUserGroupResponse) SetBody(v *AddFaceUserToUserGroupResponseBody) *AddFaceUserToUserGroupResponse {
	s.Body = v
	return s
}

type AddRecordPlanDeviceRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	PlanId        *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	StreamType    *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s AddRecordPlanDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRecordPlanDeviceRequest) GoString() string {
	return s.String()
}

func (s *AddRecordPlanDeviceRequest) SetDeviceName(v string) *AddRecordPlanDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *AddRecordPlanDeviceRequest) SetIotId(v string) *AddRecordPlanDeviceRequest {
	s.IotId = &v
	return s
}

func (s *AddRecordPlanDeviceRequest) SetIotInstanceId(v string) *AddRecordPlanDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *AddRecordPlanDeviceRequest) SetPlanId(v string) *AddRecordPlanDeviceRequest {
	s.PlanId = &v
	return s
}

func (s *AddRecordPlanDeviceRequest) SetProductKey(v string) *AddRecordPlanDeviceRequest {
	s.ProductKey = &v
	return s
}

func (s *AddRecordPlanDeviceRequest) SetStreamType(v int32) *AddRecordPlanDeviceRequest {
	s.StreamType = &v
	return s
}

type AddRecordPlanDeviceResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddRecordPlanDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddRecordPlanDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *AddRecordPlanDeviceResponseBody) SetCode(v string) *AddRecordPlanDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *AddRecordPlanDeviceResponseBody) SetErrorMessage(v string) *AddRecordPlanDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddRecordPlanDeviceResponseBody) SetRequestId(v string) *AddRecordPlanDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRecordPlanDeviceResponseBody) SetSuccess(v bool) *AddRecordPlanDeviceResponseBody {
	s.Success = &v
	return s
}

type AddRecordPlanDeviceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRecordPlanDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRecordPlanDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddRecordPlanDeviceResponse) GoString() string {
	return s.String()
}

func (s *AddRecordPlanDeviceResponse) SetHeaders(v map[string]*string) *AddRecordPlanDeviceResponse {
	s.Headers = v
	return s
}

func (s *AddRecordPlanDeviceResponse) SetStatusCode(v int32) *AddRecordPlanDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRecordPlanDeviceResponse) SetBody(v *AddRecordPlanDeviceResponseBody) *AddRecordPlanDeviceResponse {
	s.Body = v
	return s
}

type BatchQueryVisionDeviceInfoRequest struct {
	DeviceNameList []*string `json:"DeviceNameList,omitempty" xml:"DeviceNameList,omitempty" type:"Repeated"`
	IotIdList      []*string `json:"IotIdList,omitempty" xml:"IotIdList,omitempty" type:"Repeated"`
	IotInstanceId  *string   `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey     *string   `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s BatchQueryVisionDeviceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryVisionDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryVisionDeviceInfoRequest) SetDeviceNameList(v []*string) *BatchQueryVisionDeviceInfoRequest {
	s.DeviceNameList = v
	return s
}

func (s *BatchQueryVisionDeviceInfoRequest) SetIotIdList(v []*string) *BatchQueryVisionDeviceInfoRequest {
	s.IotIdList = v
	return s
}

func (s *BatchQueryVisionDeviceInfoRequest) SetIotInstanceId(v string) *BatchQueryVisionDeviceInfoRequest {
	s.IotInstanceId = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoRequest) SetProductKey(v string) *BatchQueryVisionDeviceInfoRequest {
	s.ProductKey = &v
	return s
}

type BatchQueryVisionDeviceInfoResponseBody struct {
	Code         *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *BatchQueryVisionDeviceInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                     `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchQueryVisionDeviceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryVisionDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQueryVisionDeviceInfoResponseBody) SetCode(v string) *BatchQueryVisionDeviceInfoResponseBody {
	s.Code = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBody) SetData(v *BatchQueryVisionDeviceInfoResponseBodyData) *BatchQueryVisionDeviceInfoResponseBody {
	s.Data = v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBody) SetErrorMessage(v string) *BatchQueryVisionDeviceInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBody) SetRequestId(v string) *BatchQueryVisionDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBody) SetSuccess(v bool) *BatchQueryVisionDeviceInfoResponseBody {
	s.Success = &v
	return s
}

type BatchQueryVisionDeviceInfoResponseBodyData struct {
	DeviceInfoList []*BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList `json:"DeviceInfoList,omitempty" xml:"DeviceInfoList,omitempty" type:"Repeated"`
}

func (s BatchQueryVisionDeviceInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryVisionDeviceInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchQueryVisionDeviceInfoResponseBodyData) SetDeviceInfoList(v []*BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList) *BatchQueryVisionDeviceInfoResponseBodyData {
	s.DeviceInfoList = v
	return s
}

type BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList struct {
	Description    *string                                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceType     *int32                                                                  `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	GbDeviceInfo   *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo   `json:"GbDeviceInfo,omitempty" xml:"GbDeviceInfo,omitempty" type:"Struct"`
	IotId          *string                                                                 `json:"IotId,omitempty" xml:"IotId,omitempty"`
	RtmpDeviceInfo *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo `json:"RtmpDeviceInfo,omitempty" xml:"RtmpDeviceInfo,omitempty" type:"Struct"`
}

func (s BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList) GoString() string {
	return s.String()
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList) SetDescription(v string) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList {
	s.Description = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList) SetDeviceType(v int32) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList {
	s.DeviceType = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList) SetGbDeviceInfo(v *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList {
	s.GbDeviceInfo = v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList) SetIotId(v string) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList {
	s.IotId = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList) SetRtmpDeviceInfo(v *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList {
	s.RtmpDeviceInfo = v
	return s
}

type BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo struct {
	DeviceProtocol *int32  `json:"DeviceProtocol,omitempty" xml:"DeviceProtocol,omitempty"`
	GbId           *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	NetProtocol    *int32  `json:"NetProtocol,omitempty" xml:"NetProtocol,omitempty"`
	NickName       *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	Password       *string `json:"Password,omitempty" xml:"Password,omitempty"`
	SubProductKey  *string `json:"SubProductKey,omitempty" xml:"SubProductKey,omitempty"`
}

func (s BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo) GoString() string {
	return s.String()
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo) SetDeviceProtocol(v int32) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo {
	s.DeviceProtocol = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo) SetGbId(v string) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo {
	s.GbId = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo) SetNetProtocol(v int32) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo {
	s.NetProtocol = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo) SetNickName(v string) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo {
	s.NickName = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo) SetPassword(v string) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo {
	s.Password = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo) SetSubProductKey(v string) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo {
	s.SubProductKey = &v
	return s
}

type BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo struct {
	PullAuthKey       *string `json:"PullAuthKey,omitempty" xml:"PullAuthKey,omitempty"`
	PullKeyExpireTime *int32  `json:"PullKeyExpireTime,omitempty" xml:"PullKeyExpireTime,omitempty"`
	PushAuthKey       *string `json:"PushAuthKey,omitempty" xml:"PushAuthKey,omitempty"`
	PushKeyExpireTime *int32  `json:"PushKeyExpireTime,omitempty" xml:"PushKeyExpireTime,omitempty"`
	PushUrlSample     *string `json:"PushUrlSample,omitempty" xml:"PushUrlSample,omitempty"`
	StreamName        *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	StreamStatus      *int32  `json:"StreamStatus,omitempty" xml:"StreamStatus,omitempty"`
}

func (s BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) GoString() string {
	return s.String()
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) SetPullAuthKey(v string) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo {
	s.PullAuthKey = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) SetPullKeyExpireTime(v int32) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo {
	s.PullKeyExpireTime = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) SetPushAuthKey(v string) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo {
	s.PushAuthKey = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) SetPushKeyExpireTime(v int32) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo {
	s.PushKeyExpireTime = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) SetPushUrlSample(v string) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo {
	s.PushUrlSample = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) SetStreamName(v string) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo {
	s.StreamName = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) SetStreamStatus(v int32) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo {
	s.StreamStatus = &v
	return s
}

type BatchQueryVisionDeviceInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchQueryVisionDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchQueryVisionDeviceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryVisionDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *BatchQueryVisionDeviceInfoResponse) SetHeaders(v map[string]*string) *BatchQueryVisionDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponse) SetStatusCode(v int32) *BatchQueryVisionDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponse) SetBody(v *BatchQueryVisionDeviceInfoResponseBody) *BatchQueryVisionDeviceInfoResponse {
	s.Body = v
	return s
}

type BindPictureSearchAppWithDevicesRequest struct {
	AppInstanceId *string   `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	DeviceIotIds  []*string `json:"DeviceIotIds,omitempty" xml:"DeviceIotIds,omitempty" type:"Repeated"`
	IotInstanceId *string   `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
}

func (s BindPictureSearchAppWithDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s BindPictureSearchAppWithDevicesRequest) GoString() string {
	return s.String()
}

func (s *BindPictureSearchAppWithDevicesRequest) SetAppInstanceId(v string) *BindPictureSearchAppWithDevicesRequest {
	s.AppInstanceId = &v
	return s
}

func (s *BindPictureSearchAppWithDevicesRequest) SetDeviceIotIds(v []*string) *BindPictureSearchAppWithDevicesRequest {
	s.DeviceIotIds = v
	return s
}

func (s *BindPictureSearchAppWithDevicesRequest) SetIotInstanceId(v string) *BindPictureSearchAppWithDevicesRequest {
	s.IotInstanceId = &v
	return s
}

type BindPictureSearchAppWithDevicesResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindPictureSearchAppWithDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindPictureSearchAppWithDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *BindPictureSearchAppWithDevicesResponseBody) SetCode(v string) *BindPictureSearchAppWithDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *BindPictureSearchAppWithDevicesResponseBody) SetErrorMessage(v string) *BindPictureSearchAppWithDevicesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BindPictureSearchAppWithDevicesResponseBody) SetRequestId(v string) *BindPictureSearchAppWithDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindPictureSearchAppWithDevicesResponseBody) SetSuccess(v bool) *BindPictureSearchAppWithDevicesResponseBody {
	s.Success = &v
	return s
}

type BindPictureSearchAppWithDevicesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindPictureSearchAppWithDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindPictureSearchAppWithDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s BindPictureSearchAppWithDevicesResponse) GoString() string {
	return s.String()
}

func (s *BindPictureSearchAppWithDevicesResponse) SetHeaders(v map[string]*string) *BindPictureSearchAppWithDevicesResponse {
	s.Headers = v
	return s
}

func (s *BindPictureSearchAppWithDevicesResponse) SetStatusCode(v int32) *BindPictureSearchAppWithDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *BindPictureSearchAppWithDevicesResponse) SetBody(v *BindPictureSearchAppWithDevicesResponseBody) *BindPictureSearchAppWithDevicesResponse {
	s.Body = v
	return s
}

type CheckFaceUserDoExistOnDeviceRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	IsolationId   *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CheckFaceUserDoExistOnDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckFaceUserDoExistOnDeviceRequest) GoString() string {
	return s.String()
}

func (s *CheckFaceUserDoExistOnDeviceRequest) SetDeviceName(v string) *CheckFaceUserDoExistOnDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceRequest) SetIotInstanceId(v string) *CheckFaceUserDoExistOnDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceRequest) SetIsolationId(v string) *CheckFaceUserDoExistOnDeviceRequest {
	s.IsolationId = &v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceRequest) SetProductKey(v string) *CheckFaceUserDoExistOnDeviceRequest {
	s.ProductKey = &v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceRequest) SetUserId(v string) *CheckFaceUserDoExistOnDeviceRequest {
	s.UserId = &v
	return s
}

type CheckFaceUserDoExistOnDeviceResponseBody struct {
	Code         *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *CheckFaceUserDoExistOnDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckFaceUserDoExistOnDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckFaceUserDoExistOnDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *CheckFaceUserDoExistOnDeviceResponseBody) SetCode(v string) *CheckFaceUserDoExistOnDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceResponseBody) SetData(v *CheckFaceUserDoExistOnDeviceResponseBodyData) *CheckFaceUserDoExistOnDeviceResponseBody {
	s.Data = v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceResponseBody) SetErrorMessage(v string) *CheckFaceUserDoExistOnDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceResponseBody) SetRequestId(v string) *CheckFaceUserDoExistOnDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceResponseBody) SetSuccess(v bool) *CheckFaceUserDoExistOnDeviceResponseBody {
	s.Success = &v
	return s
}

type CheckFaceUserDoExistOnDeviceResponseBodyData struct {
	DoExist *bool `json:"DoExist,omitempty" xml:"DoExist,omitempty"`
}

func (s CheckFaceUserDoExistOnDeviceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CheckFaceUserDoExistOnDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckFaceUserDoExistOnDeviceResponseBodyData) SetDoExist(v bool) *CheckFaceUserDoExistOnDeviceResponseBodyData {
	s.DoExist = &v
	return s
}

type CheckFaceUserDoExistOnDeviceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckFaceUserDoExistOnDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckFaceUserDoExistOnDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckFaceUserDoExistOnDeviceResponse) GoString() string {
	return s.String()
}

func (s *CheckFaceUserDoExistOnDeviceResponse) SetHeaders(v map[string]*string) *CheckFaceUserDoExistOnDeviceResponse {
	s.Headers = v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceResponse) SetStatusCode(v int32) *CheckFaceUserDoExistOnDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceResponse) SetBody(v *CheckFaceUserDoExistOnDeviceResponseBody) *CheckFaceUserDoExistOnDeviceResponse {
	s.Body = v
	return s
}

type ClearFaceDeviceDBRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	IsolationId   *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s ClearFaceDeviceDBRequest) String() string {
	return tea.Prettify(s)
}

func (s ClearFaceDeviceDBRequest) GoString() string {
	return s.String()
}

func (s *ClearFaceDeviceDBRequest) SetDeviceName(v string) *ClearFaceDeviceDBRequest {
	s.DeviceName = &v
	return s
}

func (s *ClearFaceDeviceDBRequest) SetIotInstanceId(v string) *ClearFaceDeviceDBRequest {
	s.IotInstanceId = &v
	return s
}

func (s *ClearFaceDeviceDBRequest) SetIsolationId(v string) *ClearFaceDeviceDBRequest {
	s.IsolationId = &v
	return s
}

func (s *ClearFaceDeviceDBRequest) SetProductKey(v string) *ClearFaceDeviceDBRequest {
	s.ProductKey = &v
	return s
}

type ClearFaceDeviceDBResponseBody struct {
	Code         *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ClearFaceDeviceDBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClearFaceDeviceDBResponseBody) GoString() string {
	return s.String()
}

func (s *ClearFaceDeviceDBResponseBody) SetCode(v string) *ClearFaceDeviceDBResponseBody {
	s.Code = &v
	return s
}

func (s *ClearFaceDeviceDBResponseBody) SetData(v map[string]interface{}) *ClearFaceDeviceDBResponseBody {
	s.Data = v
	return s
}

func (s *ClearFaceDeviceDBResponseBody) SetErrorMessage(v string) *ClearFaceDeviceDBResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ClearFaceDeviceDBResponseBody) SetRequestId(v string) *ClearFaceDeviceDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClearFaceDeviceDBResponseBody) SetSuccess(v bool) *ClearFaceDeviceDBResponseBody {
	s.Success = &v
	return s
}

type ClearFaceDeviceDBResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearFaceDeviceDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearFaceDeviceDBResponse) String() string {
	return tea.Prettify(s)
}

func (s ClearFaceDeviceDBResponse) GoString() string {
	return s.String()
}

func (s *ClearFaceDeviceDBResponse) SetHeaders(v map[string]*string) *ClearFaceDeviceDBResponse {
	s.Headers = v
	return s
}

func (s *ClearFaceDeviceDBResponse) SetStatusCode(v int32) *ClearFaceDeviceDBResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearFaceDeviceDBResponse) SetBody(v *ClearFaceDeviceDBResponseBody) *ClearFaceDeviceDBResponse {
	s.Body = v
	return s
}

type CreateEventRecordPlanRequest struct {
	EventTypes        *string `json:"EventTypes,omitempty" xml:"EventTypes,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PreRecordDuration *int32  `json:"PreRecordDuration,omitempty" xml:"PreRecordDuration,omitempty"`
	RecordDuration    *int32  `json:"RecordDuration,omitempty" xml:"RecordDuration,omitempty"`
	TemplateId        *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateEventRecordPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEventRecordPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateEventRecordPlanRequest) SetEventTypes(v string) *CreateEventRecordPlanRequest {
	s.EventTypes = &v
	return s
}

func (s *CreateEventRecordPlanRequest) SetName(v string) *CreateEventRecordPlanRequest {
	s.Name = &v
	return s
}

func (s *CreateEventRecordPlanRequest) SetPreRecordDuration(v int32) *CreateEventRecordPlanRequest {
	s.PreRecordDuration = &v
	return s
}

func (s *CreateEventRecordPlanRequest) SetRecordDuration(v int32) *CreateEventRecordPlanRequest {
	s.RecordDuration = &v
	return s
}

func (s *CreateEventRecordPlanRequest) SetTemplateId(v string) *CreateEventRecordPlanRequest {
	s.TemplateId = &v
	return s
}

type CreateEventRecordPlanResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateEventRecordPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEventRecordPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEventRecordPlanResponseBody) SetCode(v string) *CreateEventRecordPlanResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEventRecordPlanResponseBody) SetData(v string) *CreateEventRecordPlanResponseBody {
	s.Data = &v
	return s
}

func (s *CreateEventRecordPlanResponseBody) SetErrorMessage(v string) *CreateEventRecordPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateEventRecordPlanResponseBody) SetRequestId(v string) *CreateEventRecordPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEventRecordPlanResponseBody) SetSuccess(v bool) *CreateEventRecordPlanResponseBody {
	s.Success = &v
	return s
}

type CreateEventRecordPlanResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEventRecordPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEventRecordPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEventRecordPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateEventRecordPlanResponse) SetHeaders(v map[string]*string) *CreateEventRecordPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateEventRecordPlanResponse) SetStatusCode(v int32) *CreateEventRecordPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEventRecordPlanResponse) SetBody(v *CreateEventRecordPlanResponseBody) *CreateEventRecordPlanResponse {
	s.Body = v
	return s
}

type CreateGbDeviceRequest struct {
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceName       *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceType       *int32  `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	GbId             *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	IotInstanceId    *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	MediaNetProtocol *string `json:"MediaNetProtocol,omitempty" xml:"MediaNetProtocol,omitempty"`
	Password         *string `json:"Password,omitempty" xml:"Password,omitempty"`
	ProductKey       *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	SubProductKey    *string `json:"SubProductKey,omitempty" xml:"SubProductKey,omitempty"`
}

func (s CreateGbDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGbDeviceRequest) GoString() string {
	return s.String()
}

func (s *CreateGbDeviceRequest) SetDescription(v string) *CreateGbDeviceRequest {
	s.Description = &v
	return s
}

func (s *CreateGbDeviceRequest) SetDeviceName(v string) *CreateGbDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *CreateGbDeviceRequest) SetDeviceType(v int32) *CreateGbDeviceRequest {
	s.DeviceType = &v
	return s
}

func (s *CreateGbDeviceRequest) SetGbId(v string) *CreateGbDeviceRequest {
	s.GbId = &v
	return s
}

func (s *CreateGbDeviceRequest) SetIotInstanceId(v string) *CreateGbDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *CreateGbDeviceRequest) SetMediaNetProtocol(v string) *CreateGbDeviceRequest {
	s.MediaNetProtocol = &v
	return s
}

func (s *CreateGbDeviceRequest) SetPassword(v string) *CreateGbDeviceRequest {
	s.Password = &v
	return s
}

func (s *CreateGbDeviceRequest) SetProductKey(v string) *CreateGbDeviceRequest {
	s.ProductKey = &v
	return s
}

func (s *CreateGbDeviceRequest) SetSubProductKey(v string) *CreateGbDeviceRequest {
	s.SubProductKey = &v
	return s
}

type CreateGbDeviceResponseBody struct {
	Code         *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *CreateGbDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateGbDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGbDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGbDeviceResponseBody) SetCode(v string) *CreateGbDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGbDeviceResponseBody) SetData(v *CreateGbDeviceResponseBodyData) *CreateGbDeviceResponseBody {
	s.Data = v
	return s
}

func (s *CreateGbDeviceResponseBody) SetErrorMessage(v string) *CreateGbDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateGbDeviceResponseBody) SetRequestId(v string) *CreateGbDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGbDeviceResponseBody) SetSuccess(v bool) *CreateGbDeviceResponseBody {
	s.Success = &v
	return s
}

type CreateGbDeviceResponseBodyData struct {
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
}

func (s CreateGbDeviceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateGbDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateGbDeviceResponseBodyData) SetDeviceName(v string) *CreateGbDeviceResponseBodyData {
	s.DeviceName = &v
	return s
}

func (s *CreateGbDeviceResponseBodyData) SetIotId(v string) *CreateGbDeviceResponseBodyData {
	s.IotId = &v
	return s
}

type CreateGbDeviceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGbDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGbDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGbDeviceResponse) GoString() string {
	return s.String()
}

func (s *CreateGbDeviceResponse) SetHeaders(v map[string]*string) *CreateGbDeviceResponse {
	s.Headers = v
	return s
}

func (s *CreateGbDeviceResponse) SetStatusCode(v int32) *CreateGbDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGbDeviceResponse) SetBody(v *CreateGbDeviceResponseBody) *CreateGbDeviceResponse {
	s.Body = v
	return s
}

type CreateLocalFileUploadJobRequest struct {
	IotInstanceId *string                                    `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	TimeSlot      []*CreateLocalFileUploadJobRequestTimeSlot `json:"TimeSlot,omitempty" xml:"TimeSlot,omitempty" type:"Repeated"`
}

func (s CreateLocalFileUploadJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLocalFileUploadJobRequest) GoString() string {
	return s.String()
}

func (s *CreateLocalFileUploadJobRequest) SetIotInstanceId(v string) *CreateLocalFileUploadJobRequest {
	s.IotInstanceId = &v
	return s
}

func (s *CreateLocalFileUploadJobRequest) SetTimeSlot(v []*CreateLocalFileUploadJobRequestTimeSlot) *CreateLocalFileUploadJobRequest {
	s.TimeSlot = v
	return s
}

type CreateLocalFileUploadJobRequestTimeSlot struct {
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	EndTime    *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	StartTime  *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateLocalFileUploadJobRequestTimeSlot) String() string {
	return tea.Prettify(s)
}

func (s CreateLocalFileUploadJobRequestTimeSlot) GoString() string {
	return s.String()
}

func (s *CreateLocalFileUploadJobRequestTimeSlot) SetDeviceName(v string) *CreateLocalFileUploadJobRequestTimeSlot {
	s.DeviceName = &v
	return s
}

func (s *CreateLocalFileUploadJobRequestTimeSlot) SetEndTime(v int32) *CreateLocalFileUploadJobRequestTimeSlot {
	s.EndTime = &v
	return s
}

func (s *CreateLocalFileUploadJobRequestTimeSlot) SetIotId(v string) *CreateLocalFileUploadJobRequestTimeSlot {
	s.IotId = &v
	return s
}

func (s *CreateLocalFileUploadJobRequestTimeSlot) SetProductKey(v string) *CreateLocalFileUploadJobRequestTimeSlot {
	s.ProductKey = &v
	return s
}

func (s *CreateLocalFileUploadJobRequestTimeSlot) SetStartTime(v int32) *CreateLocalFileUploadJobRequestTimeSlot {
	s.StartTime = &v
	return s
}

type CreateLocalFileUploadJobResponseBody struct {
	Code         *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *CreateLocalFileUploadJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateLocalFileUploadJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLocalFileUploadJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLocalFileUploadJobResponseBody) SetCode(v string) *CreateLocalFileUploadJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateLocalFileUploadJobResponseBody) SetData(v *CreateLocalFileUploadJobResponseBodyData) *CreateLocalFileUploadJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateLocalFileUploadJobResponseBody) SetErrorMessage(v string) *CreateLocalFileUploadJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateLocalFileUploadJobResponseBody) SetRequestId(v string) *CreateLocalFileUploadJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLocalFileUploadJobResponseBody) SetSuccess(v bool) *CreateLocalFileUploadJobResponseBody {
	s.Success = &v
	return s
}

type CreateLocalFileUploadJobResponseBodyData struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CreateLocalFileUploadJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateLocalFileUploadJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateLocalFileUploadJobResponseBodyData) SetJobId(v string) *CreateLocalFileUploadJobResponseBodyData {
	s.JobId = &v
	return s
}

type CreateLocalFileUploadJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLocalFileUploadJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLocalFileUploadJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLocalFileUploadJobResponse) GoString() string {
	return s.String()
}

func (s *CreateLocalFileUploadJobResponse) SetHeaders(v map[string]*string) *CreateLocalFileUploadJobResponse {
	s.Headers = v
	return s
}

func (s *CreateLocalFileUploadJobResponse) SetStatusCode(v int32) *CreateLocalFileUploadJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLocalFileUploadJobResponse) SetBody(v *CreateLocalFileUploadJobResponseBody) *CreateLocalFileUploadJobResponse {
	s.Body = v
	return s
}

type CreateLocalRecordDownloadByTimeJobRequest struct {
	BeginTime     *int32   `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	DeviceName    *string  `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	EndTime       *int32   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IotId         *string  `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string  `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string  `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	Speed         *float32 `json:"Speed,omitempty" xml:"Speed,omitempty"`
}

func (s CreateLocalRecordDownloadByTimeJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLocalRecordDownloadByTimeJobRequest) GoString() string {
	return s.String()
}

func (s *CreateLocalRecordDownloadByTimeJobRequest) SetBeginTime(v int32) *CreateLocalRecordDownloadByTimeJobRequest {
	s.BeginTime = &v
	return s
}

func (s *CreateLocalRecordDownloadByTimeJobRequest) SetDeviceName(v string) *CreateLocalRecordDownloadByTimeJobRequest {
	s.DeviceName = &v
	return s
}

func (s *CreateLocalRecordDownloadByTimeJobRequest) SetEndTime(v int32) *CreateLocalRecordDownloadByTimeJobRequest {
	s.EndTime = &v
	return s
}

func (s *CreateLocalRecordDownloadByTimeJobRequest) SetIotId(v string) *CreateLocalRecordDownloadByTimeJobRequest {
	s.IotId = &v
	return s
}

func (s *CreateLocalRecordDownloadByTimeJobRequest) SetIotInstanceId(v string) *CreateLocalRecordDownloadByTimeJobRequest {
	s.IotInstanceId = &v
	return s
}

func (s *CreateLocalRecordDownloadByTimeJobRequest) SetProductKey(v string) *CreateLocalRecordDownloadByTimeJobRequest {
	s.ProductKey = &v
	return s
}

func (s *CreateLocalRecordDownloadByTimeJobRequest) SetSpeed(v float32) *CreateLocalRecordDownloadByTimeJobRequest {
	s.Speed = &v
	return s
}

type CreateLocalRecordDownloadByTimeJobResponseBody struct {
	Code         *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *CreateLocalRecordDownloadByTimeJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateLocalRecordDownloadByTimeJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLocalRecordDownloadByTimeJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLocalRecordDownloadByTimeJobResponseBody) SetCode(v string) *CreateLocalRecordDownloadByTimeJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateLocalRecordDownloadByTimeJobResponseBody) SetData(v *CreateLocalRecordDownloadByTimeJobResponseBodyData) *CreateLocalRecordDownloadByTimeJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateLocalRecordDownloadByTimeJobResponseBody) SetErrorMessage(v string) *CreateLocalRecordDownloadByTimeJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateLocalRecordDownloadByTimeJobResponseBody) SetRequestId(v string) *CreateLocalRecordDownloadByTimeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLocalRecordDownloadByTimeJobResponseBody) SetSuccess(v bool) *CreateLocalRecordDownloadByTimeJobResponseBody {
	s.Success = &v
	return s
}

type CreateLocalRecordDownloadByTimeJobResponseBodyData struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CreateLocalRecordDownloadByTimeJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateLocalRecordDownloadByTimeJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateLocalRecordDownloadByTimeJobResponseBodyData) SetJobId(v string) *CreateLocalRecordDownloadByTimeJobResponseBodyData {
	s.JobId = &v
	return s
}

type CreateLocalRecordDownloadByTimeJobResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLocalRecordDownloadByTimeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLocalRecordDownloadByTimeJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLocalRecordDownloadByTimeJobResponse) GoString() string {
	return s.String()
}

func (s *CreateLocalRecordDownloadByTimeJobResponse) SetHeaders(v map[string]*string) *CreateLocalRecordDownloadByTimeJobResponse {
	s.Headers = v
	return s
}

func (s *CreateLocalRecordDownloadByTimeJobResponse) SetStatusCode(v int32) *CreateLocalRecordDownloadByTimeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLocalRecordDownloadByTimeJobResponse) SetBody(v *CreateLocalRecordDownloadByTimeJobResponseBody) *CreateLocalRecordDownloadByTimeJobResponse {
	s.Body = v
	return s
}

type CreatePictureSearchAppRequest struct {
	Desc          *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreatePictureSearchAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePictureSearchAppRequest) GoString() string {
	return s.String()
}

func (s *CreatePictureSearchAppRequest) SetDesc(v string) *CreatePictureSearchAppRequest {
	s.Desc = &v
	return s
}

func (s *CreatePictureSearchAppRequest) SetIotInstanceId(v string) *CreatePictureSearchAppRequest {
	s.IotInstanceId = &v
	return s
}

func (s *CreatePictureSearchAppRequest) SetName(v string) *CreatePictureSearchAppRequest {
	s.Name = &v
	return s
}

type CreatePictureSearchAppResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePictureSearchAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePictureSearchAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePictureSearchAppResponseBody) SetCode(v string) *CreatePictureSearchAppResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePictureSearchAppResponseBody) SetData(v string) *CreatePictureSearchAppResponseBody {
	s.Data = &v
	return s
}

func (s *CreatePictureSearchAppResponseBody) SetErrorMessage(v string) *CreatePictureSearchAppResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreatePictureSearchAppResponseBody) SetRequestId(v string) *CreatePictureSearchAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePictureSearchAppResponseBody) SetSuccess(v bool) *CreatePictureSearchAppResponseBody {
	s.Success = &v
	return s
}

type CreatePictureSearchAppResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePictureSearchAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePictureSearchAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePictureSearchAppResponse) GoString() string {
	return s.String()
}

func (s *CreatePictureSearchAppResponse) SetHeaders(v map[string]*string) *CreatePictureSearchAppResponse {
	s.Headers = v
	return s
}

func (s *CreatePictureSearchAppResponse) SetStatusCode(v int32) *CreatePictureSearchAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePictureSearchAppResponse) SetBody(v *CreatePictureSearchAppResponseBody) *CreatePictureSearchAppResponse {
	s.Body = v
	return s
}

type CreatePictureSearchJobRequest struct {
	AppInstanceId     *string  `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	BodyThreshold     *float32 `json:"BodyThreshold,omitempty" xml:"BodyThreshold,omitempty"`
	EndTime           *int64   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PictureSearchType *int32   `json:"PictureSearchType,omitempty" xml:"PictureSearchType,omitempty"`
	SearchPicUrl      *string  `json:"SearchPicUrl,omitempty" xml:"SearchPicUrl,omitempty"`
	StartTime         *int64   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Threshold         *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreatePictureSearchJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePictureSearchJobRequest) GoString() string {
	return s.String()
}

func (s *CreatePictureSearchJobRequest) SetAppInstanceId(v string) *CreatePictureSearchJobRequest {
	s.AppInstanceId = &v
	return s
}

func (s *CreatePictureSearchJobRequest) SetBodyThreshold(v float32) *CreatePictureSearchJobRequest {
	s.BodyThreshold = &v
	return s
}

func (s *CreatePictureSearchJobRequest) SetEndTime(v int64) *CreatePictureSearchJobRequest {
	s.EndTime = &v
	return s
}

func (s *CreatePictureSearchJobRequest) SetPictureSearchType(v int32) *CreatePictureSearchJobRequest {
	s.PictureSearchType = &v
	return s
}

func (s *CreatePictureSearchJobRequest) SetSearchPicUrl(v string) *CreatePictureSearchJobRequest {
	s.SearchPicUrl = &v
	return s
}

func (s *CreatePictureSearchJobRequest) SetStartTime(v int64) *CreatePictureSearchJobRequest {
	s.StartTime = &v
	return s
}

func (s *CreatePictureSearchJobRequest) SetThreshold(v float32) *CreatePictureSearchJobRequest {
	s.Threshold = &v
	return s
}

type CreatePictureSearchJobResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePictureSearchJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePictureSearchJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePictureSearchJobResponseBody) SetCode(v string) *CreatePictureSearchJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePictureSearchJobResponseBody) SetData(v string) *CreatePictureSearchJobResponseBody {
	s.Data = &v
	return s
}

func (s *CreatePictureSearchJobResponseBody) SetErrorMessage(v string) *CreatePictureSearchJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreatePictureSearchJobResponseBody) SetRequestId(v string) *CreatePictureSearchJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePictureSearchJobResponseBody) SetSuccess(v bool) *CreatePictureSearchJobResponseBody {
	s.Success = &v
	return s
}

type CreatePictureSearchJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePictureSearchJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePictureSearchJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePictureSearchJobResponse) GoString() string {
	return s.String()
}

func (s *CreatePictureSearchJobResponse) SetHeaders(v map[string]*string) *CreatePictureSearchJobResponse {
	s.Headers = v
	return s
}

func (s *CreatePictureSearchJobResponse) SetStatusCode(v int32) *CreatePictureSearchJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePictureSearchJobResponse) SetBody(v *CreatePictureSearchJobResponseBody) *CreatePictureSearchJobResponse {
	s.Body = v
	return s
}

type CreateRecordDownloadByTimeJobRequest struct {
	BeginTime     *int32  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	EndTime       *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	RecordType    *int32  `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	StreamType    *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s CreateRecordDownloadByTimeJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRecordDownloadByTimeJobRequest) GoString() string {
	return s.String()
}

func (s *CreateRecordDownloadByTimeJobRequest) SetBeginTime(v int32) *CreateRecordDownloadByTimeJobRequest {
	s.BeginTime = &v
	return s
}

func (s *CreateRecordDownloadByTimeJobRequest) SetDeviceName(v string) *CreateRecordDownloadByTimeJobRequest {
	s.DeviceName = &v
	return s
}

func (s *CreateRecordDownloadByTimeJobRequest) SetEndTime(v int32) *CreateRecordDownloadByTimeJobRequest {
	s.EndTime = &v
	return s
}

func (s *CreateRecordDownloadByTimeJobRequest) SetIotId(v string) *CreateRecordDownloadByTimeJobRequest {
	s.IotId = &v
	return s
}

func (s *CreateRecordDownloadByTimeJobRequest) SetIotInstanceId(v string) *CreateRecordDownloadByTimeJobRequest {
	s.IotInstanceId = &v
	return s
}

func (s *CreateRecordDownloadByTimeJobRequest) SetProductKey(v string) *CreateRecordDownloadByTimeJobRequest {
	s.ProductKey = &v
	return s
}

func (s *CreateRecordDownloadByTimeJobRequest) SetRecordType(v int32) *CreateRecordDownloadByTimeJobRequest {
	s.RecordType = &v
	return s
}

func (s *CreateRecordDownloadByTimeJobRequest) SetStreamType(v int32) *CreateRecordDownloadByTimeJobRequest {
	s.StreamType = &v
	return s
}

type CreateRecordDownloadByTimeJobResponseBody struct {
	Code         *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *CreateRecordDownloadByTimeJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRecordDownloadByTimeJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRecordDownloadByTimeJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRecordDownloadByTimeJobResponseBody) SetCode(v string) *CreateRecordDownloadByTimeJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRecordDownloadByTimeJobResponseBody) SetData(v *CreateRecordDownloadByTimeJobResponseBodyData) *CreateRecordDownloadByTimeJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateRecordDownloadByTimeJobResponseBody) SetErrorMessage(v string) *CreateRecordDownloadByTimeJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateRecordDownloadByTimeJobResponseBody) SetRequestId(v string) *CreateRecordDownloadByTimeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRecordDownloadByTimeJobResponseBody) SetSuccess(v bool) *CreateRecordDownloadByTimeJobResponseBody {
	s.Success = &v
	return s
}

type CreateRecordDownloadByTimeJobResponseBodyData struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CreateRecordDownloadByTimeJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateRecordDownloadByTimeJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateRecordDownloadByTimeJobResponseBodyData) SetJobId(v string) *CreateRecordDownloadByTimeJobResponseBodyData {
	s.JobId = &v
	return s
}

type CreateRecordDownloadByTimeJobResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRecordDownloadByTimeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRecordDownloadByTimeJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRecordDownloadByTimeJobResponse) GoString() string {
	return s.String()
}

func (s *CreateRecordDownloadByTimeJobResponse) SetHeaders(v map[string]*string) *CreateRecordDownloadByTimeJobResponse {
	s.Headers = v
	return s
}

func (s *CreateRecordDownloadByTimeJobResponse) SetStatusCode(v int32) *CreateRecordDownloadByTimeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRecordDownloadByTimeJobResponse) SetBody(v *CreateRecordDownloadByTimeJobResponseBody) *CreateRecordDownloadByTimeJobResponse {
	s.Body = v
	return s
}

type CreateRecordPlanRequest struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateRecordPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRecordPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateRecordPlanRequest) SetName(v string) *CreateRecordPlanRequest {
	s.Name = &v
	return s
}

func (s *CreateRecordPlanRequest) SetTemplateId(v string) *CreateRecordPlanRequest {
	s.TemplateId = &v
	return s
}

type CreateRecordPlanResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRecordPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRecordPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRecordPlanResponseBody) SetCode(v string) *CreateRecordPlanResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRecordPlanResponseBody) SetData(v string) *CreateRecordPlanResponseBody {
	s.Data = &v
	return s
}

func (s *CreateRecordPlanResponseBody) SetErrorMessage(v string) *CreateRecordPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateRecordPlanResponseBody) SetRequestId(v string) *CreateRecordPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRecordPlanResponseBody) SetSuccess(v bool) *CreateRecordPlanResponseBody {
	s.Success = &v
	return s
}

type CreateRecordPlanResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRecordPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRecordPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRecordPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateRecordPlanResponse) SetHeaders(v map[string]*string) *CreateRecordPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateRecordPlanResponse) SetStatusCode(v int32) *CreateRecordPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRecordPlanResponse) SetBody(v *CreateRecordPlanResponseBody) *CreateRecordPlanResponse {
	s.Body = v
	return s
}

type CreateRtmpDeviceRequest struct {
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceName        *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotInstanceId     *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey        *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	PullAuthKey       *string `json:"PullAuthKey,omitempty" xml:"PullAuthKey,omitempty"`
	PullKeyExpireTime *int32  `json:"PullKeyExpireTime,omitempty" xml:"PullKeyExpireTime,omitempty"`
	PushAuthKey       *string `json:"PushAuthKey,omitempty" xml:"PushAuthKey,omitempty"`
	PushKeyExpireTime *int32  `json:"PushKeyExpireTime,omitempty" xml:"PushKeyExpireTime,omitempty"`
	SubStreamName     *string `json:"SubStreamName,omitempty" xml:"SubStreamName,omitempty"`
}

func (s CreateRtmpDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRtmpDeviceRequest) GoString() string {
	return s.String()
}

func (s *CreateRtmpDeviceRequest) SetDescription(v string) *CreateRtmpDeviceRequest {
	s.Description = &v
	return s
}

func (s *CreateRtmpDeviceRequest) SetDeviceName(v string) *CreateRtmpDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *CreateRtmpDeviceRequest) SetIotInstanceId(v string) *CreateRtmpDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *CreateRtmpDeviceRequest) SetProductKey(v string) *CreateRtmpDeviceRequest {
	s.ProductKey = &v
	return s
}

func (s *CreateRtmpDeviceRequest) SetPullAuthKey(v string) *CreateRtmpDeviceRequest {
	s.PullAuthKey = &v
	return s
}

func (s *CreateRtmpDeviceRequest) SetPullKeyExpireTime(v int32) *CreateRtmpDeviceRequest {
	s.PullKeyExpireTime = &v
	return s
}

func (s *CreateRtmpDeviceRequest) SetPushAuthKey(v string) *CreateRtmpDeviceRequest {
	s.PushAuthKey = &v
	return s
}

func (s *CreateRtmpDeviceRequest) SetPushKeyExpireTime(v int32) *CreateRtmpDeviceRequest {
	s.PushKeyExpireTime = &v
	return s
}

func (s *CreateRtmpDeviceRequest) SetSubStreamName(v string) *CreateRtmpDeviceRequest {
	s.SubStreamName = &v
	return s
}

type CreateRtmpDeviceResponseBody struct {
	Code         *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *CreateRtmpDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRtmpDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRtmpDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRtmpDeviceResponseBody) SetCode(v string) *CreateRtmpDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRtmpDeviceResponseBody) SetData(v *CreateRtmpDeviceResponseBodyData) *CreateRtmpDeviceResponseBody {
	s.Data = v
	return s
}

func (s *CreateRtmpDeviceResponseBody) SetErrorMessage(v string) *CreateRtmpDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateRtmpDeviceResponseBody) SetRequestId(v string) *CreateRtmpDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRtmpDeviceResponseBody) SetSuccess(v bool) *CreateRtmpDeviceResponseBody {
	s.Success = &v
	return s
}

type CreateRtmpDeviceResponseBodyData struct {
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s CreateRtmpDeviceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateRtmpDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateRtmpDeviceResponseBodyData) SetDeviceName(v string) *CreateRtmpDeviceResponseBodyData {
	s.DeviceName = &v
	return s
}

func (s *CreateRtmpDeviceResponseBodyData) SetIotId(v string) *CreateRtmpDeviceResponseBodyData {
	s.IotId = &v
	return s
}

func (s *CreateRtmpDeviceResponseBodyData) SetStreamName(v string) *CreateRtmpDeviceResponseBodyData {
	s.StreamName = &v
	return s
}

type CreateRtmpDeviceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRtmpDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRtmpDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRtmpDeviceResponse) GoString() string {
	return s.String()
}

func (s *CreateRtmpDeviceResponse) SetHeaders(v map[string]*string) *CreateRtmpDeviceResponse {
	s.Headers = v
	return s
}

func (s *CreateRtmpDeviceResponse) SetStatusCode(v int32) *CreateRtmpDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRtmpDeviceResponse) SetBody(v *CreateRtmpDeviceResponseBody) *CreateRtmpDeviceResponse {
	s.Body = v
	return s
}

type CreateStreamPushJobRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	JobType       *int32  `json:"JobType,omitempty" xml:"JobType,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	PushUrl       *string `json:"PushUrl,omitempty" xml:"PushUrl,omitempty"`
	StreamType    *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s CreateStreamPushJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStreamPushJobRequest) GoString() string {
	return s.String()
}

func (s *CreateStreamPushJobRequest) SetDeviceName(v string) *CreateStreamPushJobRequest {
	s.DeviceName = &v
	return s
}

func (s *CreateStreamPushJobRequest) SetIotId(v string) *CreateStreamPushJobRequest {
	s.IotId = &v
	return s
}

func (s *CreateStreamPushJobRequest) SetIotInstanceId(v string) *CreateStreamPushJobRequest {
	s.IotInstanceId = &v
	return s
}

func (s *CreateStreamPushJobRequest) SetJobType(v int32) *CreateStreamPushJobRequest {
	s.JobType = &v
	return s
}

func (s *CreateStreamPushJobRequest) SetProductKey(v string) *CreateStreamPushJobRequest {
	s.ProductKey = &v
	return s
}

func (s *CreateStreamPushJobRequest) SetPushUrl(v string) *CreateStreamPushJobRequest {
	s.PushUrl = &v
	return s
}

func (s *CreateStreamPushJobRequest) SetStreamType(v int32) *CreateStreamPushJobRequest {
	s.StreamType = &v
	return s
}

type CreateStreamPushJobResponseBody struct {
	Code         *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *CreateStreamPushJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateStreamPushJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateStreamPushJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStreamPushJobResponseBody) SetCode(v string) *CreateStreamPushJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateStreamPushJobResponseBody) SetData(v *CreateStreamPushJobResponseBodyData) *CreateStreamPushJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateStreamPushJobResponseBody) SetErrorMessage(v string) *CreateStreamPushJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateStreamPushJobResponseBody) SetRequestId(v string) *CreateStreamPushJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStreamPushJobResponseBody) SetSuccess(v bool) *CreateStreamPushJobResponseBody {
	s.Success = &v
	return s
}

type CreateStreamPushJobResponseBodyData struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CreateStreamPushJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateStreamPushJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateStreamPushJobResponseBodyData) SetJobId(v string) *CreateStreamPushJobResponseBodyData {
	s.JobId = &v
	return s
}

type CreateStreamPushJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStreamPushJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStreamPushJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStreamPushJobResponse) GoString() string {
	return s.String()
}

func (s *CreateStreamPushJobResponse) SetHeaders(v map[string]*string) *CreateStreamPushJobResponse {
	s.Headers = v
	return s
}

func (s *CreateStreamPushJobResponse) SetStatusCode(v int32) *CreateStreamPushJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStreamPushJobResponse) SetBody(v *CreateStreamPushJobResponseBody) *CreateStreamPushJobResponse {
	s.Body = v
	return s
}

type CreateStreamSnapshotJobRequest struct {
	DeviceName       *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId            *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId    *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey       *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	SnapshotInterval *int32  `json:"SnapshotInterval,omitempty" xml:"SnapshotInterval,omitempty"`
	StreamType       *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s CreateStreamSnapshotJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStreamSnapshotJobRequest) GoString() string {
	return s.String()
}

func (s *CreateStreamSnapshotJobRequest) SetDeviceName(v string) *CreateStreamSnapshotJobRequest {
	s.DeviceName = &v
	return s
}

func (s *CreateStreamSnapshotJobRequest) SetIotId(v string) *CreateStreamSnapshotJobRequest {
	s.IotId = &v
	return s
}

func (s *CreateStreamSnapshotJobRequest) SetIotInstanceId(v string) *CreateStreamSnapshotJobRequest {
	s.IotInstanceId = &v
	return s
}

func (s *CreateStreamSnapshotJobRequest) SetProductKey(v string) *CreateStreamSnapshotJobRequest {
	s.ProductKey = &v
	return s
}

func (s *CreateStreamSnapshotJobRequest) SetSnapshotInterval(v int32) *CreateStreamSnapshotJobRequest {
	s.SnapshotInterval = &v
	return s
}

func (s *CreateStreamSnapshotJobRequest) SetStreamType(v int32) *CreateStreamSnapshotJobRequest {
	s.StreamType = &v
	return s
}

type CreateStreamSnapshotJobResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateStreamSnapshotJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateStreamSnapshotJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStreamSnapshotJobResponseBody) SetCode(v string) *CreateStreamSnapshotJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateStreamSnapshotJobResponseBody) SetErrorMessage(v string) *CreateStreamSnapshotJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateStreamSnapshotJobResponseBody) SetRequestId(v string) *CreateStreamSnapshotJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStreamSnapshotJobResponseBody) SetSuccess(v bool) *CreateStreamSnapshotJobResponseBody {
	s.Success = &v
	return s
}

type CreateStreamSnapshotJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStreamSnapshotJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStreamSnapshotJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStreamSnapshotJobResponse) GoString() string {
	return s.String()
}

func (s *CreateStreamSnapshotJobResponse) SetHeaders(v map[string]*string) *CreateStreamSnapshotJobResponse {
	s.Headers = v
	return s
}

func (s *CreateStreamSnapshotJobResponse) SetStatusCode(v int32) *CreateStreamSnapshotJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStreamSnapshotJobResponse) SetBody(v *CreateStreamSnapshotJobResponseBody) *CreateStreamSnapshotJobResponse {
	s.Body = v
	return s
}

type CreateTimeTemplateRequest struct {
	AllDay       *int32                                   `json:"AllDay,omitempty" xml:"AllDay,omitempty"`
	Name         *string                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	TimeSections []*CreateTimeTemplateRequestTimeSections `json:"TimeSections,omitempty" xml:"TimeSections,omitempty" type:"Repeated"`
}

func (s CreateTimeTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTimeTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTimeTemplateRequest) SetAllDay(v int32) *CreateTimeTemplateRequest {
	s.AllDay = &v
	return s
}

func (s *CreateTimeTemplateRequest) SetName(v string) *CreateTimeTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateTimeTemplateRequest) SetTimeSections(v []*CreateTimeTemplateRequestTimeSections) *CreateTimeTemplateRequest {
	s.TimeSections = v
	return s
}

type CreateTimeTemplateRequestTimeSections struct {
	Begin     *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	DayOfWeek *int32 `json:"DayOfWeek,omitempty" xml:"DayOfWeek,omitempty"`
	End       *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s CreateTimeTemplateRequestTimeSections) String() string {
	return tea.Prettify(s)
}

func (s CreateTimeTemplateRequestTimeSections) GoString() string {
	return s.String()
}

func (s *CreateTimeTemplateRequestTimeSections) SetBegin(v int32) *CreateTimeTemplateRequestTimeSections {
	s.Begin = &v
	return s
}

func (s *CreateTimeTemplateRequestTimeSections) SetDayOfWeek(v int32) *CreateTimeTemplateRequestTimeSections {
	s.DayOfWeek = &v
	return s
}

func (s *CreateTimeTemplateRequestTimeSections) SetEnd(v int32) *CreateTimeTemplateRequestTimeSections {
	s.End = &v
	return s
}

type CreateTimeTemplateResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTimeTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTimeTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTimeTemplateResponseBody) SetCode(v string) *CreateTimeTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTimeTemplateResponseBody) SetData(v string) *CreateTimeTemplateResponseBody {
	s.Data = &v
	return s
}

func (s *CreateTimeTemplateResponseBody) SetErrorMessage(v string) *CreateTimeTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateTimeTemplateResponseBody) SetRequestId(v string) *CreateTimeTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTimeTemplateResponseBody) SetSuccess(v bool) *CreateTimeTemplateResponseBody {
	s.Success = &v
	return s
}

type CreateTimeTemplateResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTimeTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTimeTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTimeTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateTimeTemplateResponse) SetHeaders(v map[string]*string) *CreateTimeTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateTimeTemplateResponse) SetStatusCode(v int32) *CreateTimeTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTimeTemplateResponse) SetBody(v *CreateTimeTemplateResponseBody) *CreateTimeTemplateResponse {
	s.Body = v
	return s
}

type DeleteEventRecordPlanRequest struct {
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s DeleteEventRecordPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventRecordPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventRecordPlanRequest) SetPlanId(v string) *DeleteEventRecordPlanRequest {
	s.PlanId = &v
	return s
}

type DeleteEventRecordPlanResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteEventRecordPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventRecordPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventRecordPlanResponseBody) SetCode(v string) *DeleteEventRecordPlanResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEventRecordPlanResponseBody) SetErrorMessage(v string) *DeleteEventRecordPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteEventRecordPlanResponseBody) SetRequestId(v string) *DeleteEventRecordPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEventRecordPlanResponseBody) SetSuccess(v bool) *DeleteEventRecordPlanResponseBody {
	s.Success = &v
	return s
}

type DeleteEventRecordPlanResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEventRecordPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEventRecordPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventRecordPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventRecordPlanResponse) SetHeaders(v map[string]*string) *DeleteEventRecordPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventRecordPlanResponse) SetStatusCode(v int32) *DeleteEventRecordPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEventRecordPlanResponse) SetBody(v *DeleteEventRecordPlanResponseBody) *DeleteEventRecordPlanResponse {
	s.Body = v
	return s
}

type DeleteEventRecordPlanDeviceRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	StreamType    *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s DeleteEventRecordPlanDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventRecordPlanDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventRecordPlanDeviceRequest) SetDeviceName(v string) *DeleteEventRecordPlanDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *DeleteEventRecordPlanDeviceRequest) SetIotId(v string) *DeleteEventRecordPlanDeviceRequest {
	s.IotId = &v
	return s
}

func (s *DeleteEventRecordPlanDeviceRequest) SetIotInstanceId(v string) *DeleteEventRecordPlanDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *DeleteEventRecordPlanDeviceRequest) SetProductKey(v string) *DeleteEventRecordPlanDeviceRequest {
	s.ProductKey = &v
	return s
}

func (s *DeleteEventRecordPlanDeviceRequest) SetStreamType(v int32) *DeleteEventRecordPlanDeviceRequest {
	s.StreamType = &v
	return s
}

type DeleteEventRecordPlanDeviceResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteEventRecordPlanDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventRecordPlanDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventRecordPlanDeviceResponseBody) SetCode(v string) *DeleteEventRecordPlanDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEventRecordPlanDeviceResponseBody) SetErrorMessage(v string) *DeleteEventRecordPlanDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteEventRecordPlanDeviceResponseBody) SetRequestId(v string) *DeleteEventRecordPlanDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEventRecordPlanDeviceResponseBody) SetSuccess(v bool) *DeleteEventRecordPlanDeviceResponseBody {
	s.Success = &v
	return s
}

type DeleteEventRecordPlanDeviceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEventRecordPlanDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEventRecordPlanDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventRecordPlanDeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventRecordPlanDeviceResponse) SetHeaders(v map[string]*string) *DeleteEventRecordPlanDeviceResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventRecordPlanDeviceResponse) SetStatusCode(v int32) *DeleteEventRecordPlanDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEventRecordPlanDeviceResponse) SetBody(v *DeleteEventRecordPlanDeviceResponseBody) *DeleteEventRecordPlanDeviceResponse {
	s.Body = v
	return s
}

type DeleteFaceDeviceGroupRequest struct {
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	IsolationId   *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
}

func (s DeleteFaceDeviceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceDeviceGroupRequest) SetDeviceGroupId(v string) *DeleteFaceDeviceGroupRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *DeleteFaceDeviceGroupRequest) SetIsolationId(v string) *DeleteFaceDeviceGroupRequest {
	s.IsolationId = &v
	return s
}

type DeleteFaceDeviceGroupResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFaceDeviceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceDeviceGroupResponseBody) SetCode(v string) *DeleteFaceDeviceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFaceDeviceGroupResponseBody) SetErrorMessage(v string) *DeleteFaceDeviceGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFaceDeviceGroupResponseBody) SetRequestId(v string) *DeleteFaceDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFaceDeviceGroupResponseBody) SetSuccess(v bool) *DeleteFaceDeviceGroupResponseBody {
	s.Success = &v
	return s
}

type DeleteFaceDeviceGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFaceDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFaceDeviceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceDeviceGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceDeviceGroupResponse) SetHeaders(v map[string]*string) *DeleteFaceDeviceGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceDeviceGroupResponse) SetStatusCode(v int32) *DeleteFaceDeviceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaceDeviceGroupResponse) SetBody(v *DeleteFaceDeviceGroupResponseBody) *DeleteFaceDeviceGroupResponse {
	s.Body = v
	return s
}

type DeleteFaceUserRequest struct {
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteFaceUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserRequest) SetIsolationId(v string) *DeleteFaceUserRequest {
	s.IsolationId = &v
	return s
}

func (s *DeleteFaceUserRequest) SetUserId(v string) *DeleteFaceUserRequest {
	s.UserId = &v
	return s
}

type DeleteFaceUserResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFaceUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserResponseBody) SetCode(v string) *DeleteFaceUserResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFaceUserResponseBody) SetErrorMessage(v string) *DeleteFaceUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFaceUserResponseBody) SetRequestId(v string) *DeleteFaceUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFaceUserResponseBody) SetSuccess(v bool) *DeleteFaceUserResponseBody {
	s.Success = &v
	return s
}

type DeleteFaceUserResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFaceUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFaceUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserResponse) SetHeaders(v map[string]*string) *DeleteFaceUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceUserResponse) SetStatusCode(v int32) *DeleteFaceUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaceUserResponse) SetBody(v *DeleteFaceUserResponseBody) *DeleteFaceUserResponse {
	s.Body = v
	return s
}

type DeleteFaceUserGroupRequest struct {
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s DeleteFaceUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceUserGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserGroupRequest) SetIsolationId(v string) *DeleteFaceUserGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *DeleteFaceUserGroupRequest) SetUserGroupId(v string) *DeleteFaceUserGroupRequest {
	s.UserGroupId = &v
	return s
}

type DeleteFaceUserGroupResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFaceUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserGroupResponseBody) SetCode(v string) *DeleteFaceUserGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFaceUserGroupResponseBody) SetErrorMessage(v string) *DeleteFaceUserGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFaceUserGroupResponseBody) SetRequestId(v string) *DeleteFaceUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFaceUserGroupResponseBody) SetSuccess(v bool) *DeleteFaceUserGroupResponseBody {
	s.Success = &v
	return s
}

type DeleteFaceUserGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFaceUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFaceUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceUserGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserGroupResponse) SetHeaders(v map[string]*string) *DeleteFaceUserGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceUserGroupResponse) SetStatusCode(v int32) *DeleteFaceUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaceUserGroupResponse) SetBody(v *DeleteFaceUserGroupResponseBody) *DeleteFaceUserGroupResponse {
	s.Body = v
	return s
}

type DeleteFaceUserGroupAndDeviceGroupRelationRequest struct {
	ControlId   *string `json:"ControlId,omitempty" xml:"ControlId,omitempty"`
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
}

func (s DeleteFaceUserGroupAndDeviceGroupRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceUserGroupAndDeviceGroupRelationRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationRequest) SetControlId(v string) *DeleteFaceUserGroupAndDeviceGroupRelationRequest {
	s.ControlId = &v
	return s
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationRequest) SetIsolationId(v string) *DeleteFaceUserGroupAndDeviceGroupRelationRequest {
	s.IsolationId = &v
	return s
}

type DeleteFaceUserGroupAndDeviceGroupRelationResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFaceUserGroupAndDeviceGroupRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceUserGroupAndDeviceGroupRelationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationResponseBody) SetCode(v string) *DeleteFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationResponseBody) SetErrorMessage(v string) *DeleteFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationResponseBody) SetRequestId(v string) *DeleteFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationResponseBody) SetSuccess(v bool) *DeleteFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.Success = &v
	return s
}

type DeleteFaceUserGroupAndDeviceGroupRelationResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFaceUserGroupAndDeviceGroupRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFaceUserGroupAndDeviceGroupRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceUserGroupAndDeviceGroupRelationResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationResponse) SetHeaders(v map[string]*string) *DeleteFaceUserGroupAndDeviceGroupRelationResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationResponse) SetStatusCode(v int32) *DeleteFaceUserGroupAndDeviceGroupRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationResponse) SetBody(v *DeleteFaceUserGroupAndDeviceGroupRelationResponseBody) *DeleteFaceUserGroupAndDeviceGroupRelationResponse {
	s.Body = v
	return s
}

type DeleteFaceUserPictureRequest struct {
	FacePicMd5  *string `json:"FacePicMd5,omitempty" xml:"FacePicMd5,omitempty"`
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteFaceUserPictureRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceUserPictureRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserPictureRequest) SetFacePicMd5(v string) *DeleteFaceUserPictureRequest {
	s.FacePicMd5 = &v
	return s
}

func (s *DeleteFaceUserPictureRequest) SetIsolationId(v string) *DeleteFaceUserPictureRequest {
	s.IsolationId = &v
	return s
}

func (s *DeleteFaceUserPictureRequest) SetUserId(v string) *DeleteFaceUserPictureRequest {
	s.UserId = &v
	return s
}

type DeleteFaceUserPictureResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFaceUserPictureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceUserPictureResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserPictureResponseBody) SetCode(v string) *DeleteFaceUserPictureResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFaceUserPictureResponseBody) SetErrorMessage(v string) *DeleteFaceUserPictureResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFaceUserPictureResponseBody) SetRequestId(v string) *DeleteFaceUserPictureResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFaceUserPictureResponseBody) SetSuccess(v bool) *DeleteFaceUserPictureResponseBody {
	s.Success = &v
	return s
}

type DeleteFaceUserPictureResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFaceUserPictureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFaceUserPictureResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceUserPictureResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserPictureResponse) SetHeaders(v map[string]*string) *DeleteFaceUserPictureResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceUserPictureResponse) SetStatusCode(v int32) *DeleteFaceUserPictureResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaceUserPictureResponse) SetBody(v *DeleteFaceUserPictureResponseBody) *DeleteFaceUserPictureResponse {
	s.Body = v
	return s
}

type DeleteGbDeviceRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s DeleteGbDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGbDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteGbDeviceRequest) SetDeviceName(v string) *DeleteGbDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *DeleteGbDeviceRequest) SetIotId(v string) *DeleteGbDeviceRequest {
	s.IotId = &v
	return s
}

func (s *DeleteGbDeviceRequest) SetIotInstanceId(v string) *DeleteGbDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *DeleteGbDeviceRequest) SetProductKey(v string) *DeleteGbDeviceRequest {
	s.ProductKey = &v
	return s
}

type DeleteGbDeviceResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteGbDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGbDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGbDeviceResponseBody) SetCode(v string) *DeleteGbDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGbDeviceResponseBody) SetErrorMessage(v string) *DeleteGbDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteGbDeviceResponseBody) SetRequestId(v string) *DeleteGbDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGbDeviceResponseBody) SetSuccess(v bool) *DeleteGbDeviceResponseBody {
	s.Success = &v
	return s
}

type DeleteGbDeviceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGbDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGbDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGbDeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteGbDeviceResponse) SetHeaders(v map[string]*string) *DeleteGbDeviceResponse {
	s.Headers = v
	return s
}

func (s *DeleteGbDeviceResponse) SetStatusCode(v int32) *DeleteGbDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGbDeviceResponse) SetBody(v *DeleteGbDeviceResponseBody) *DeleteGbDeviceResponse {
	s.Body = v
	return s
}

type DeleteLocalFileUploadJobRequest struct {
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	JobId         *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DeleteLocalFileUploadJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLocalFileUploadJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteLocalFileUploadJobRequest) SetIotInstanceId(v string) *DeleteLocalFileUploadJobRequest {
	s.IotInstanceId = &v
	return s
}

func (s *DeleteLocalFileUploadJobRequest) SetJobId(v string) *DeleteLocalFileUploadJobRequest {
	s.JobId = &v
	return s
}

type DeleteLocalFileUploadJobResponseBody struct {
	Code         *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteLocalFileUploadJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLocalFileUploadJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLocalFileUploadJobResponseBody) SetCode(v string) *DeleteLocalFileUploadJobResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteLocalFileUploadJobResponseBody) SetData(v map[string]interface{}) *DeleteLocalFileUploadJobResponseBody {
	s.Data = v
	return s
}

func (s *DeleteLocalFileUploadJobResponseBody) SetErrorMessage(v string) *DeleteLocalFileUploadJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteLocalFileUploadJobResponseBody) SetRequestId(v string) *DeleteLocalFileUploadJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLocalFileUploadJobResponseBody) SetSuccess(v bool) *DeleteLocalFileUploadJobResponseBody {
	s.Success = &v
	return s
}

type DeleteLocalFileUploadJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLocalFileUploadJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLocalFileUploadJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLocalFileUploadJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteLocalFileUploadJobResponse) SetHeaders(v map[string]*string) *DeleteLocalFileUploadJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteLocalFileUploadJobResponse) SetStatusCode(v int32) *DeleteLocalFileUploadJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLocalFileUploadJobResponse) SetBody(v *DeleteLocalFileUploadJobResponseBody) *DeleteLocalFileUploadJobResponse {
	s.Body = v
	return s
}

type DeletePictureRequest struct {
	DeviceName    *string   `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string   `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string   `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	PictureIdList []*string `json:"PictureIdList,omitempty" xml:"PictureIdList,omitempty" type:"Repeated"`
	ProductKey    *string   `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s DeletePictureRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePictureRequest) GoString() string {
	return s.String()
}

func (s *DeletePictureRequest) SetDeviceName(v string) *DeletePictureRequest {
	s.DeviceName = &v
	return s
}

func (s *DeletePictureRequest) SetIotId(v string) *DeletePictureRequest {
	s.IotId = &v
	return s
}

func (s *DeletePictureRequest) SetIotInstanceId(v string) *DeletePictureRequest {
	s.IotInstanceId = &v
	return s
}

func (s *DeletePictureRequest) SetPictureIdList(v []*string) *DeletePictureRequest {
	s.PictureIdList = v
	return s
}

func (s *DeletePictureRequest) SetProductKey(v string) *DeletePictureRequest {
	s.ProductKey = &v
	return s
}

type DeletePictureResponseBody struct {
	Code         *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *DeletePictureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeletePictureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePictureResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePictureResponseBody) SetCode(v string) *DeletePictureResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePictureResponseBody) SetData(v *DeletePictureResponseBodyData) *DeletePictureResponseBody {
	s.Data = v
	return s
}

func (s *DeletePictureResponseBody) SetErrorMessage(v string) *DeletePictureResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeletePictureResponseBody) SetRequestId(v string) *DeletePictureResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePictureResponseBody) SetSuccess(v bool) *DeletePictureResponseBody {
	s.Success = &v
	return s
}

type DeletePictureResponseBodyData struct {
	DeletedCount *int32 `json:"DeletedCount,omitempty" xml:"DeletedCount,omitempty"`
}

func (s DeletePictureResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeletePictureResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeletePictureResponseBodyData) SetDeletedCount(v int32) *DeletePictureResponseBodyData {
	s.DeletedCount = &v
	return s
}

type DeletePictureResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePictureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePictureResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePictureResponse) GoString() string {
	return s.String()
}

func (s *DeletePictureResponse) SetHeaders(v map[string]*string) *DeletePictureResponse {
	s.Headers = v
	return s
}

func (s *DeletePictureResponse) SetStatusCode(v int32) *DeletePictureResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePictureResponse) SetBody(v *DeletePictureResponseBody) *DeletePictureResponse {
	s.Body = v
	return s
}

type DeleteRecordRequest struct {
	DeviceName    *string   `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	FileNameList  []*string `json:"FileNameList,omitempty" xml:"FileNameList,omitempty" type:"Repeated"`
	IotId         *string   `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string   `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string   `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s DeleteRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecordRequest) SetDeviceName(v string) *DeleteRecordRequest {
	s.DeviceName = &v
	return s
}

func (s *DeleteRecordRequest) SetFileNameList(v []*string) *DeleteRecordRequest {
	s.FileNameList = v
	return s
}

func (s *DeleteRecordRequest) SetIotId(v string) *DeleteRecordRequest {
	s.IotId = &v
	return s
}

func (s *DeleteRecordRequest) SetIotInstanceId(v string) *DeleteRecordRequest {
	s.IotInstanceId = &v
	return s
}

func (s *DeleteRecordRequest) SetProductKey(v string) *DeleteRecordRequest {
	s.ProductKey = &v
	return s
}

type DeleteRecordResponseBody struct {
	Code         *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *DeleteRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRecordResponseBody) SetCode(v string) *DeleteRecordResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRecordResponseBody) SetData(v *DeleteRecordResponseBodyData) *DeleteRecordResponseBody {
	s.Data = v
	return s
}

func (s *DeleteRecordResponseBody) SetErrorMessage(v string) *DeleteRecordResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRecordResponseBody) SetRequestId(v string) *DeleteRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRecordResponseBody) SetSuccess(v bool) *DeleteRecordResponseBody {
	s.Success = &v
	return s
}

type DeleteRecordResponseBodyData struct {
	DeletedCount *int32 `json:"DeletedCount,omitempty" xml:"DeletedCount,omitempty"`
}

func (s DeleteRecordResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteRecordResponseBodyData) SetDeletedCount(v int32) *DeleteRecordResponseBodyData {
	s.DeletedCount = &v
	return s
}

type DeleteRecordResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordResponse) GoString() string {
	return s.String()
}

func (s *DeleteRecordResponse) SetHeaders(v map[string]*string) *DeleteRecordResponse {
	s.Headers = v
	return s
}

func (s *DeleteRecordResponse) SetStatusCode(v int32) *DeleteRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRecordResponse) SetBody(v *DeleteRecordResponseBody) *DeleteRecordResponse {
	s.Body = v
	return s
}

type DeleteRecordPlanRequest struct {
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s DeleteRecordPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecordPlanRequest) SetPlanId(v string) *DeleteRecordPlanRequest {
	s.PlanId = &v
	return s
}

type DeleteRecordPlanResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRecordPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRecordPlanResponseBody) SetCode(v string) *DeleteRecordPlanResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRecordPlanResponseBody) SetErrorMessage(v string) *DeleteRecordPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRecordPlanResponseBody) SetRequestId(v string) *DeleteRecordPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRecordPlanResponseBody) SetSuccess(v bool) *DeleteRecordPlanResponseBody {
	s.Success = &v
	return s
}

type DeleteRecordPlanResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRecordPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRecordPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteRecordPlanResponse) SetHeaders(v map[string]*string) *DeleteRecordPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteRecordPlanResponse) SetStatusCode(v int32) *DeleteRecordPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRecordPlanResponse) SetBody(v *DeleteRecordPlanResponseBody) *DeleteRecordPlanResponse {
	s.Body = v
	return s
}

type DeleteRecordPlanDeviceRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	StreamType    *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s DeleteRecordPlanDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordPlanDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecordPlanDeviceRequest) SetDeviceName(v string) *DeleteRecordPlanDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *DeleteRecordPlanDeviceRequest) SetIotId(v string) *DeleteRecordPlanDeviceRequest {
	s.IotId = &v
	return s
}

func (s *DeleteRecordPlanDeviceRequest) SetIotInstanceId(v string) *DeleteRecordPlanDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *DeleteRecordPlanDeviceRequest) SetProductKey(v string) *DeleteRecordPlanDeviceRequest {
	s.ProductKey = &v
	return s
}

func (s *DeleteRecordPlanDeviceRequest) SetStreamType(v int32) *DeleteRecordPlanDeviceRequest {
	s.StreamType = &v
	return s
}

type DeleteRecordPlanDeviceResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRecordPlanDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordPlanDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRecordPlanDeviceResponseBody) SetCode(v string) *DeleteRecordPlanDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRecordPlanDeviceResponseBody) SetErrorMessage(v string) *DeleteRecordPlanDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRecordPlanDeviceResponseBody) SetRequestId(v string) *DeleteRecordPlanDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRecordPlanDeviceResponseBody) SetSuccess(v bool) *DeleteRecordPlanDeviceResponseBody {
	s.Success = &v
	return s
}

type DeleteRecordPlanDeviceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRecordPlanDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRecordPlanDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordPlanDeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteRecordPlanDeviceResponse) SetHeaders(v map[string]*string) *DeleteRecordPlanDeviceResponse {
	s.Headers = v
	return s
}

func (s *DeleteRecordPlanDeviceResponse) SetStatusCode(v int32) *DeleteRecordPlanDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRecordPlanDeviceResponse) SetBody(v *DeleteRecordPlanDeviceResponseBody) *DeleteRecordPlanDeviceResponse {
	s.Body = v
	return s
}

type DeleteRtmpDeviceRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s DeleteRtmpDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRtmpDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteRtmpDeviceRequest) SetDeviceName(v string) *DeleteRtmpDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *DeleteRtmpDeviceRequest) SetIotId(v string) *DeleteRtmpDeviceRequest {
	s.IotId = &v
	return s
}

func (s *DeleteRtmpDeviceRequest) SetIotInstanceId(v string) *DeleteRtmpDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *DeleteRtmpDeviceRequest) SetProductKey(v string) *DeleteRtmpDeviceRequest {
	s.ProductKey = &v
	return s
}

type DeleteRtmpDeviceResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRtmpDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRtmpDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRtmpDeviceResponseBody) SetCode(v string) *DeleteRtmpDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRtmpDeviceResponseBody) SetErrorMessage(v string) *DeleteRtmpDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRtmpDeviceResponseBody) SetRequestId(v string) *DeleteRtmpDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRtmpDeviceResponseBody) SetSuccess(v bool) *DeleteRtmpDeviceResponseBody {
	s.Success = &v
	return s
}

type DeleteRtmpDeviceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRtmpDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRtmpDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRtmpDeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteRtmpDeviceResponse) SetHeaders(v map[string]*string) *DeleteRtmpDeviceResponse {
	s.Headers = v
	return s
}

func (s *DeleteRtmpDeviceResponse) SetStatusCode(v int32) *DeleteRtmpDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRtmpDeviceResponse) SetBody(v *DeleteRtmpDeviceResponseBody) *DeleteRtmpDeviceResponse {
	s.Body = v
	return s
}

type DeleteRtmpKeyRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	Type          *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeleteRtmpKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRtmpKeyRequest) GoString() string {
	return s.String()
}

func (s *DeleteRtmpKeyRequest) SetDeviceName(v string) *DeleteRtmpKeyRequest {
	s.DeviceName = &v
	return s
}

func (s *DeleteRtmpKeyRequest) SetIotId(v string) *DeleteRtmpKeyRequest {
	s.IotId = &v
	return s
}

func (s *DeleteRtmpKeyRequest) SetIotInstanceId(v string) *DeleteRtmpKeyRequest {
	s.IotInstanceId = &v
	return s
}

func (s *DeleteRtmpKeyRequest) SetProductKey(v string) *DeleteRtmpKeyRequest {
	s.ProductKey = &v
	return s
}

func (s *DeleteRtmpKeyRequest) SetType(v int32) *DeleteRtmpKeyRequest {
	s.Type = &v
	return s
}

type DeleteRtmpKeyResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRtmpKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRtmpKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRtmpKeyResponseBody) SetCode(v string) *DeleteRtmpKeyResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRtmpKeyResponseBody) SetErrorMessage(v string) *DeleteRtmpKeyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRtmpKeyResponseBody) SetRequestId(v string) *DeleteRtmpKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRtmpKeyResponseBody) SetSuccess(v bool) *DeleteRtmpKeyResponseBody {
	s.Success = &v
	return s
}

type DeleteRtmpKeyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRtmpKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRtmpKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRtmpKeyResponse) GoString() string {
	return s.String()
}

func (s *DeleteRtmpKeyResponse) SetHeaders(v map[string]*string) *DeleteRtmpKeyResponse {
	s.Headers = v
	return s
}

func (s *DeleteRtmpKeyResponse) SetStatusCode(v int32) *DeleteRtmpKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRtmpKeyResponse) SetBody(v *DeleteRtmpKeyResponseBody) *DeleteRtmpKeyResponse {
	s.Body = v
	return s
}

type DeleteStreamPushJobRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	JobId         *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s DeleteStreamPushJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteStreamPushJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteStreamPushJobRequest) SetDeviceName(v string) *DeleteStreamPushJobRequest {
	s.DeviceName = &v
	return s
}

func (s *DeleteStreamPushJobRequest) SetIotId(v string) *DeleteStreamPushJobRequest {
	s.IotId = &v
	return s
}

func (s *DeleteStreamPushJobRequest) SetIotInstanceId(v string) *DeleteStreamPushJobRequest {
	s.IotInstanceId = &v
	return s
}

func (s *DeleteStreamPushJobRequest) SetJobId(v string) *DeleteStreamPushJobRequest {
	s.JobId = &v
	return s
}

func (s *DeleteStreamPushJobRequest) SetProductKey(v string) *DeleteStreamPushJobRequest {
	s.ProductKey = &v
	return s
}

type DeleteStreamPushJobResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteStreamPushJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteStreamPushJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStreamPushJobResponseBody) SetCode(v string) *DeleteStreamPushJobResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteStreamPushJobResponseBody) SetErrorMessage(v string) *DeleteStreamPushJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteStreamPushJobResponseBody) SetRequestId(v string) *DeleteStreamPushJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStreamPushJobResponseBody) SetSuccess(v bool) *DeleteStreamPushJobResponseBody {
	s.Success = &v
	return s
}

type DeleteStreamPushJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStreamPushJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStreamPushJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteStreamPushJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteStreamPushJobResponse) SetHeaders(v map[string]*string) *DeleteStreamPushJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteStreamPushJobResponse) SetStatusCode(v int32) *DeleteStreamPushJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStreamPushJobResponse) SetBody(v *DeleteStreamPushJobResponseBody) *DeleteStreamPushJobResponse {
	s.Body = v
	return s
}

type DeleteStreamSnapshotJobRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	StreamType    *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s DeleteStreamSnapshotJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteStreamSnapshotJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteStreamSnapshotJobRequest) SetDeviceName(v string) *DeleteStreamSnapshotJobRequest {
	s.DeviceName = &v
	return s
}

func (s *DeleteStreamSnapshotJobRequest) SetIotId(v string) *DeleteStreamSnapshotJobRequest {
	s.IotId = &v
	return s
}

func (s *DeleteStreamSnapshotJobRequest) SetIotInstanceId(v string) *DeleteStreamSnapshotJobRequest {
	s.IotInstanceId = &v
	return s
}

func (s *DeleteStreamSnapshotJobRequest) SetProductKey(v string) *DeleteStreamSnapshotJobRequest {
	s.ProductKey = &v
	return s
}

func (s *DeleteStreamSnapshotJobRequest) SetStreamType(v int32) *DeleteStreamSnapshotJobRequest {
	s.StreamType = &v
	return s
}

type DeleteStreamSnapshotJobResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteStreamSnapshotJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteStreamSnapshotJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStreamSnapshotJobResponseBody) SetCode(v string) *DeleteStreamSnapshotJobResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteStreamSnapshotJobResponseBody) SetErrorMessage(v string) *DeleteStreamSnapshotJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteStreamSnapshotJobResponseBody) SetRequestId(v string) *DeleteStreamSnapshotJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStreamSnapshotJobResponseBody) SetSuccess(v bool) *DeleteStreamSnapshotJobResponseBody {
	s.Success = &v
	return s
}

type DeleteStreamSnapshotJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStreamSnapshotJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStreamSnapshotJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteStreamSnapshotJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteStreamSnapshotJobResponse) SetHeaders(v map[string]*string) *DeleteStreamSnapshotJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteStreamSnapshotJobResponse) SetStatusCode(v int32) *DeleteStreamSnapshotJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStreamSnapshotJobResponse) SetBody(v *DeleteStreamSnapshotJobResponseBody) *DeleteStreamSnapshotJobResponse {
	s.Body = v
	return s
}

type DeleteTimeTemplateRequest struct {
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteTimeTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTimeTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteTimeTemplateRequest) SetTemplateId(v string) *DeleteTimeTemplateRequest {
	s.TemplateId = &v
	return s
}

type DeleteTimeTemplateResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTimeTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTimeTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTimeTemplateResponseBody) SetCode(v string) *DeleteTimeTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTimeTemplateResponseBody) SetErrorMessage(v string) *DeleteTimeTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteTimeTemplateResponseBody) SetRequestId(v string) *DeleteTimeTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTimeTemplateResponseBody) SetSuccess(v bool) *DeleteTimeTemplateResponseBody {
	s.Success = &v
	return s
}

type DeleteTimeTemplateResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTimeTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTimeTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTimeTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteTimeTemplateResponse) SetHeaders(v map[string]*string) *DeleteTimeTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteTimeTemplateResponse) SetStatusCode(v int32) *DeleteTimeTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTimeTemplateResponse) SetBody(v *DeleteTimeTemplateResponseBody) *DeleteTimeTemplateResponse {
	s.Body = v
	return s
}

type DetectUserFaceByUrlRequest struct {
	FacePicUrl *string `json:"FacePicUrl,omitempty" xml:"FacePicUrl,omitempty"`
}

func (s DetectUserFaceByUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectUserFaceByUrlRequest) GoString() string {
	return s.String()
}

func (s *DetectUserFaceByUrlRequest) SetFacePicUrl(v string) *DetectUserFaceByUrlRequest {
	s.FacePicUrl = &v
	return s
}

type DetectUserFaceByUrlResponseBody struct {
	Code         *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *DetectUserFaceByUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetectUserFaceByUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectUserFaceByUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DetectUserFaceByUrlResponseBody) SetCode(v string) *DetectUserFaceByUrlResponseBody {
	s.Code = &v
	return s
}

func (s *DetectUserFaceByUrlResponseBody) SetData(v *DetectUserFaceByUrlResponseBodyData) *DetectUserFaceByUrlResponseBody {
	s.Data = v
	return s
}

func (s *DetectUserFaceByUrlResponseBody) SetErrorMessage(v string) *DetectUserFaceByUrlResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DetectUserFaceByUrlResponseBody) SetRequestId(v string) *DetectUserFaceByUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectUserFaceByUrlResponseBody) SetSuccess(v bool) *DetectUserFaceByUrlResponseBody {
	s.Success = &v
	return s
}

type DetectUserFaceByUrlResponseBodyData struct {
	Data []*DetectUserFaceByUrlResponseBodyDataData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DetectUserFaceByUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectUserFaceByUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectUserFaceByUrlResponseBodyData) SetData(v []*DetectUserFaceByUrlResponseBodyDataData) *DetectUserFaceByUrlResponseBodyData {
	s.Data = v
	return s
}

type DetectUserFaceByUrlResponseBodyDataData struct {
	Age                *int32                                            `json:"Age,omitempty" xml:"Age,omitempty"`
	BlurScore          *float32                                          `json:"BlurScore,omitempty" xml:"BlurScore,omitempty"`
	FaceProbability    *float32                                          `json:"FaceProbability,omitempty" xml:"FaceProbability,omitempty"`
	FaceRects          *DetectUserFaceByUrlResponseBodyDataDataFaceRects `json:"FaceRects,omitempty" xml:"FaceRects,omitempty" type:"Struct"`
	Gender             *int32                                            `json:"Gender,omitempty" xml:"Gender,omitempty"`
	GoodForLibrary     *bool                                             `json:"GoodForLibrary,omitempty" xml:"GoodForLibrary,omitempty"`
	GoodForRecognition *bool                                             `json:"GoodForRecognition,omitempty" xml:"GoodForRecognition,omitempty"`
	Landmarks          *DetectUserFaceByUrlResponseBodyDataDataLandmarks `json:"Landmarks,omitempty" xml:"Landmarks,omitempty" type:"Struct"`
	OcclusionScore     *float32                                          `json:"OcclusionScore,omitempty" xml:"OcclusionScore,omitempty"`
	PoseScore          *float32                                          `json:"PoseScore,omitempty" xml:"PoseScore,omitempty"`
}

func (s DetectUserFaceByUrlResponseBodyDataData) String() string {
	return tea.Prettify(s)
}

func (s DetectUserFaceByUrlResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *DetectUserFaceByUrlResponseBodyDataData) SetAge(v int32) *DetectUserFaceByUrlResponseBodyDataData {
	s.Age = &v
	return s
}

func (s *DetectUserFaceByUrlResponseBodyDataData) SetBlurScore(v float32) *DetectUserFaceByUrlResponseBodyDataData {
	s.BlurScore = &v
	return s
}

func (s *DetectUserFaceByUrlResponseBodyDataData) SetFaceProbability(v float32) *DetectUserFaceByUrlResponseBodyDataData {
	s.FaceProbability = &v
	return s
}

func (s *DetectUserFaceByUrlResponseBodyDataData) SetFaceRects(v *DetectUserFaceByUrlResponseBodyDataDataFaceRects) *DetectUserFaceByUrlResponseBodyDataData {
	s.FaceRects = v
	return s
}

func (s *DetectUserFaceByUrlResponseBodyDataData) SetGender(v int32) *DetectUserFaceByUrlResponseBodyDataData {
	s.Gender = &v
	return s
}

func (s *DetectUserFaceByUrlResponseBodyDataData) SetGoodForLibrary(v bool) *DetectUserFaceByUrlResponseBodyDataData {
	s.GoodForLibrary = &v
	return s
}

func (s *DetectUserFaceByUrlResponseBodyDataData) SetGoodForRecognition(v bool) *DetectUserFaceByUrlResponseBodyDataData {
	s.GoodForRecognition = &v
	return s
}

func (s *DetectUserFaceByUrlResponseBodyDataData) SetLandmarks(v *DetectUserFaceByUrlResponseBodyDataDataLandmarks) *DetectUserFaceByUrlResponseBodyDataData {
	s.Landmarks = v
	return s
}

func (s *DetectUserFaceByUrlResponseBodyDataData) SetOcclusionScore(v float32) *DetectUserFaceByUrlResponseBodyDataData {
	s.OcclusionScore = &v
	return s
}

func (s *DetectUserFaceByUrlResponseBodyDataData) SetPoseScore(v float32) *DetectUserFaceByUrlResponseBodyDataData {
	s.PoseScore = &v
	return s
}

type DetectUserFaceByUrlResponseBodyDataDataFaceRects struct {
	FaceRectsData []*string `json:"FaceRectsData,omitempty" xml:"FaceRectsData,omitempty" type:"Repeated"`
}

func (s DetectUserFaceByUrlResponseBodyDataDataFaceRects) String() string {
	return tea.Prettify(s)
}

func (s DetectUserFaceByUrlResponseBodyDataDataFaceRects) GoString() string {
	return s.String()
}

func (s *DetectUserFaceByUrlResponseBodyDataDataFaceRects) SetFaceRectsData(v []*string) *DetectUserFaceByUrlResponseBodyDataDataFaceRects {
	s.FaceRectsData = v
	return s
}

type DetectUserFaceByUrlResponseBodyDataDataLandmarks struct {
	LandmarksData []*string `json:"LandmarksData,omitempty" xml:"LandmarksData,omitempty" type:"Repeated"`
}

func (s DetectUserFaceByUrlResponseBodyDataDataLandmarks) String() string {
	return tea.Prettify(s)
}

func (s DetectUserFaceByUrlResponseBodyDataDataLandmarks) GoString() string {
	return s.String()
}

func (s *DetectUserFaceByUrlResponseBodyDataDataLandmarks) SetLandmarksData(v []*string) *DetectUserFaceByUrlResponseBodyDataDataLandmarks {
	s.LandmarksData = v
	return s
}

type DetectUserFaceByUrlResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectUserFaceByUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectUserFaceByUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectUserFaceByUrlResponse) GoString() string {
	return s.String()
}

func (s *DetectUserFaceByUrlResponse) SetHeaders(v map[string]*string) *DetectUserFaceByUrlResponse {
	s.Headers = v
	return s
}

func (s *DetectUserFaceByUrlResponse) SetStatusCode(v int32) *DetectUserFaceByUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectUserFaceByUrlResponse) SetBody(v *DetectUserFaceByUrlResponseBody) *DetectUserFaceByUrlResponse {
	s.Body = v
	return s
}

type EnableGbSubDeviceRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	SubDeviceId   *string `json:"SubDeviceId,omitempty" xml:"SubDeviceId,omitempty"`
}

func (s EnableGbSubDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableGbSubDeviceRequest) GoString() string {
	return s.String()
}

func (s *EnableGbSubDeviceRequest) SetDeviceName(v string) *EnableGbSubDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *EnableGbSubDeviceRequest) SetIotId(v string) *EnableGbSubDeviceRequest {
	s.IotId = &v
	return s
}

func (s *EnableGbSubDeviceRequest) SetIotInstanceId(v string) *EnableGbSubDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *EnableGbSubDeviceRequest) SetProductKey(v string) *EnableGbSubDeviceRequest {
	s.ProductKey = &v
	return s
}

func (s *EnableGbSubDeviceRequest) SetSubDeviceId(v string) *EnableGbSubDeviceRequest {
	s.SubDeviceId = &v
	return s
}

type EnableGbSubDeviceResponseBody struct {
	Code         *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *EnableGbSubDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableGbSubDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableGbSubDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *EnableGbSubDeviceResponseBody) SetCode(v string) *EnableGbSubDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *EnableGbSubDeviceResponseBody) SetData(v *EnableGbSubDeviceResponseBodyData) *EnableGbSubDeviceResponseBody {
	s.Data = v
	return s
}

func (s *EnableGbSubDeviceResponseBody) SetErrorMessage(v string) *EnableGbSubDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *EnableGbSubDeviceResponseBody) SetRequestId(v string) *EnableGbSubDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableGbSubDeviceResponseBody) SetSuccess(v bool) *EnableGbSubDeviceResponseBody {
	s.Success = &v
	return s
}

type EnableGbSubDeviceResponseBodyData struct {
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s EnableGbSubDeviceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EnableGbSubDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *EnableGbSubDeviceResponseBodyData) SetDeviceName(v string) *EnableGbSubDeviceResponseBodyData {
	s.DeviceName = &v
	return s
}

func (s *EnableGbSubDeviceResponseBodyData) SetProductKey(v string) *EnableGbSubDeviceResponseBodyData {
	s.ProductKey = &v
	return s
}

type EnableGbSubDeviceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableGbSubDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableGbSubDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableGbSubDeviceResponse) GoString() string {
	return s.String()
}

func (s *EnableGbSubDeviceResponse) SetHeaders(v map[string]*string) *EnableGbSubDeviceResponse {
	s.Headers = v
	return s
}

func (s *EnableGbSubDeviceResponse) SetStatusCode(v int32) *EnableGbSubDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableGbSubDeviceResponse) SetBody(v *EnableGbSubDeviceResponseBody) *EnableGbSubDeviceResponse {
	s.Body = v
	return s
}

type GetPictureSearchJobStatusRequest struct {
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	JobId         *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetPictureSearchJobStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPictureSearchJobStatusRequest) GoString() string {
	return s.String()
}

func (s *GetPictureSearchJobStatusRequest) SetAppInstanceId(v string) *GetPictureSearchJobStatusRequest {
	s.AppInstanceId = &v
	return s
}

func (s *GetPictureSearchJobStatusRequest) SetJobId(v string) *GetPictureSearchJobStatusRequest {
	s.JobId = &v
	return s
}

type GetPictureSearchJobStatusResponseBody struct {
	Code         *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *GetPictureSearchJobStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPictureSearchJobStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPictureSearchJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetPictureSearchJobStatusResponseBody) SetCode(v string) *GetPictureSearchJobStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetPictureSearchJobStatusResponseBody) SetData(v *GetPictureSearchJobStatusResponseBodyData) *GetPictureSearchJobStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetPictureSearchJobStatusResponseBody) SetErrorMessage(v string) *GetPictureSearchJobStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPictureSearchJobStatusResponseBody) SetRequestId(v string) *GetPictureSearchJobStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPictureSearchJobStatusResponseBody) SetSuccess(v bool) *GetPictureSearchJobStatusResponseBody {
	s.Success = &v
	return s
}

type GetPictureSearchJobStatusResponseBodyData struct {
	CreateTime   *int64   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EndTime      *int64   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	JobId        *string  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobStatus    *int32   `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	SearchPicUrl *string  `json:"SearchPicUrl,omitempty" xml:"SearchPicUrl,omitempty"`
	StartTime    *int64   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Threshold    *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s GetPictureSearchJobStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPictureSearchJobStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPictureSearchJobStatusResponseBodyData) SetCreateTime(v int64) *GetPictureSearchJobStatusResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetPictureSearchJobStatusResponseBodyData) SetEndTime(v int64) *GetPictureSearchJobStatusResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetPictureSearchJobStatusResponseBodyData) SetJobId(v string) *GetPictureSearchJobStatusResponseBodyData {
	s.JobId = &v
	return s
}

func (s *GetPictureSearchJobStatusResponseBodyData) SetJobStatus(v int32) *GetPictureSearchJobStatusResponseBodyData {
	s.JobStatus = &v
	return s
}

func (s *GetPictureSearchJobStatusResponseBodyData) SetSearchPicUrl(v string) *GetPictureSearchJobStatusResponseBodyData {
	s.SearchPicUrl = &v
	return s
}

func (s *GetPictureSearchJobStatusResponseBodyData) SetStartTime(v int64) *GetPictureSearchJobStatusResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetPictureSearchJobStatusResponseBodyData) SetThreshold(v float32) *GetPictureSearchJobStatusResponseBodyData {
	s.Threshold = &v
	return s
}

type GetPictureSearchJobStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPictureSearchJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPictureSearchJobStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPictureSearchJobStatusResponse) GoString() string {
	return s.String()
}

func (s *GetPictureSearchJobStatusResponse) SetHeaders(v map[string]*string) *GetPictureSearchJobStatusResponse {
	s.Headers = v
	return s
}

func (s *GetPictureSearchJobStatusResponse) SetStatusCode(v int32) *GetPictureSearchJobStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPictureSearchJobStatusResponse) SetBody(v *GetPictureSearchJobStatusResponseBody) *GetPictureSearchJobStatusResponse {
	s.Body = v
	return s
}

type PictureSearchPictureRequest struct {
	AppInstanceId     *string  `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	ContainPicUrl     *bool    `json:"ContainPicUrl,omitempty" xml:"ContainPicUrl,omitempty"`
	CurrentPage       *int32   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime           *int64   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize          *int32   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PictureSearchType *int32   `json:"PictureSearchType,omitempty" xml:"PictureSearchType,omitempty"`
	SearchPicUrl      *string  `json:"SearchPicUrl,omitempty" xml:"SearchPicUrl,omitempty"`
	StartTime         *int64   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Threshold         *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s PictureSearchPictureRequest) String() string {
	return tea.Prettify(s)
}

func (s PictureSearchPictureRequest) GoString() string {
	return s.String()
}

func (s *PictureSearchPictureRequest) SetAppInstanceId(v string) *PictureSearchPictureRequest {
	s.AppInstanceId = &v
	return s
}

func (s *PictureSearchPictureRequest) SetContainPicUrl(v bool) *PictureSearchPictureRequest {
	s.ContainPicUrl = &v
	return s
}

func (s *PictureSearchPictureRequest) SetCurrentPage(v int32) *PictureSearchPictureRequest {
	s.CurrentPage = &v
	return s
}

func (s *PictureSearchPictureRequest) SetEndTime(v int64) *PictureSearchPictureRequest {
	s.EndTime = &v
	return s
}

func (s *PictureSearchPictureRequest) SetPageSize(v int32) *PictureSearchPictureRequest {
	s.PageSize = &v
	return s
}

func (s *PictureSearchPictureRequest) SetPictureSearchType(v int32) *PictureSearchPictureRequest {
	s.PictureSearchType = &v
	return s
}

func (s *PictureSearchPictureRequest) SetSearchPicUrl(v string) *PictureSearchPictureRequest {
	s.SearchPicUrl = &v
	return s
}

func (s *PictureSearchPictureRequest) SetStartTime(v int64) *PictureSearchPictureRequest {
	s.StartTime = &v
	return s
}

func (s *PictureSearchPictureRequest) SetThreshold(v float32) *PictureSearchPictureRequest {
	s.Threshold = &v
	return s
}

type PictureSearchPictureResponseBody struct {
	Code         *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *PictureSearchPictureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                               `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PictureSearchPictureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PictureSearchPictureResponseBody) GoString() string {
	return s.String()
}

func (s *PictureSearchPictureResponseBody) SetCode(v string) *PictureSearchPictureResponseBody {
	s.Code = &v
	return s
}

func (s *PictureSearchPictureResponseBody) SetData(v *PictureSearchPictureResponseBodyData) *PictureSearchPictureResponseBody {
	s.Data = v
	return s
}

func (s *PictureSearchPictureResponseBody) SetErrorMessage(v string) *PictureSearchPictureResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *PictureSearchPictureResponseBody) SetRequestId(v string) *PictureSearchPictureResponseBody {
	s.RequestId = &v
	return s
}

func (s *PictureSearchPictureResponseBody) SetSuccess(v bool) *PictureSearchPictureResponseBody {
	s.Success = &v
	return s
}

type PictureSearchPictureResponseBodyData struct {
	CurrentPage *int32                                          `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageCount   *int32                                          `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageData    []*PictureSearchPictureResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	PageSize    *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total       *int32                                          `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s PictureSearchPictureResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PictureSearchPictureResponseBodyData) GoString() string {
	return s.String()
}

func (s *PictureSearchPictureResponseBodyData) SetCurrentPage(v int32) *PictureSearchPictureResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *PictureSearchPictureResponseBodyData) SetPageCount(v int32) *PictureSearchPictureResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *PictureSearchPictureResponseBodyData) SetPageData(v []*PictureSearchPictureResponseBodyDataPageData) *PictureSearchPictureResponseBodyData {
	s.PageData = v
	return s
}

func (s *PictureSearchPictureResponseBodyData) SetPageSize(v int32) *PictureSearchPictureResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *PictureSearchPictureResponseBodyData) SetTotal(v int32) *PictureSearchPictureResponseBodyData {
	s.Total = &v
	return s
}

type PictureSearchPictureResponseBodyDataPageData struct {
	EventTime    *int64   `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	GatewayIotId *string  `json:"GatewayIotId,omitempty" xml:"GatewayIotId,omitempty"`
	IotId        *string  `json:"IotId,omitempty" xml:"IotId,omitempty"`
	PicUrl       *string  `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	Threshold    *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	VectorId     *string  `json:"VectorId,omitempty" xml:"VectorId,omitempty"`
	VectorType   *int32   `json:"VectorType,omitempty" xml:"VectorType,omitempty"`
}

func (s PictureSearchPictureResponseBodyDataPageData) String() string {
	return tea.Prettify(s)
}

func (s PictureSearchPictureResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *PictureSearchPictureResponseBodyDataPageData) SetEventTime(v int64) *PictureSearchPictureResponseBodyDataPageData {
	s.EventTime = &v
	return s
}

func (s *PictureSearchPictureResponseBodyDataPageData) SetGatewayIotId(v string) *PictureSearchPictureResponseBodyDataPageData {
	s.GatewayIotId = &v
	return s
}

func (s *PictureSearchPictureResponseBodyDataPageData) SetIotId(v string) *PictureSearchPictureResponseBodyDataPageData {
	s.IotId = &v
	return s
}

func (s *PictureSearchPictureResponseBodyDataPageData) SetPicUrl(v string) *PictureSearchPictureResponseBodyDataPageData {
	s.PicUrl = &v
	return s
}

func (s *PictureSearchPictureResponseBodyDataPageData) SetThreshold(v float32) *PictureSearchPictureResponseBodyDataPageData {
	s.Threshold = &v
	return s
}

func (s *PictureSearchPictureResponseBodyDataPageData) SetVectorId(v string) *PictureSearchPictureResponseBodyDataPageData {
	s.VectorId = &v
	return s
}

func (s *PictureSearchPictureResponseBodyDataPageData) SetVectorType(v int32) *PictureSearchPictureResponseBodyDataPageData {
	s.VectorType = &v
	return s
}

type PictureSearchPictureResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PictureSearchPictureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PictureSearchPictureResponse) String() string {
	return tea.Prettify(s)
}

func (s PictureSearchPictureResponse) GoString() string {
	return s.String()
}

func (s *PictureSearchPictureResponse) SetHeaders(v map[string]*string) *PictureSearchPictureResponse {
	s.Headers = v
	return s
}

func (s *PictureSearchPictureResponse) SetStatusCode(v int32) *PictureSearchPictureResponse {
	s.StatusCode = &v
	return s
}

func (s *PictureSearchPictureResponse) SetBody(v *PictureSearchPictureResponseBody) *PictureSearchPictureResponse {
	s.Body = v
	return s
}

type QueryCarProcessEventsRequest struct {
	ActionType    *int32  `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	AreaIndex     *int32  `json:"AreaIndex,omitempty" xml:"AreaIndex,omitempty"`
	BeginTime     *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime       *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PlateNo       *string `json:"PlateNo,omitempty" xml:"PlateNo,omitempty"`
	SubDeviceName *string `json:"SubDeviceName,omitempty" xml:"SubDeviceName,omitempty"`
	SubIotId      *string `json:"SubIotId,omitempty" xml:"SubIotId,omitempty"`
	SubProductKey *string `json:"SubProductKey,omitempty" xml:"SubProductKey,omitempty"`
}

func (s QueryCarProcessEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCarProcessEventsRequest) GoString() string {
	return s.String()
}

func (s *QueryCarProcessEventsRequest) SetActionType(v int32) *QueryCarProcessEventsRequest {
	s.ActionType = &v
	return s
}

func (s *QueryCarProcessEventsRequest) SetAreaIndex(v int32) *QueryCarProcessEventsRequest {
	s.AreaIndex = &v
	return s
}

func (s *QueryCarProcessEventsRequest) SetBeginTime(v int64) *QueryCarProcessEventsRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryCarProcessEventsRequest) SetCurrentPage(v int32) *QueryCarProcessEventsRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryCarProcessEventsRequest) SetEndTime(v int64) *QueryCarProcessEventsRequest {
	s.EndTime = &v
	return s
}

func (s *QueryCarProcessEventsRequest) SetIotInstanceId(v string) *QueryCarProcessEventsRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryCarProcessEventsRequest) SetPageSize(v int32) *QueryCarProcessEventsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryCarProcessEventsRequest) SetPlateNo(v string) *QueryCarProcessEventsRequest {
	s.PlateNo = &v
	return s
}

func (s *QueryCarProcessEventsRequest) SetSubDeviceName(v string) *QueryCarProcessEventsRequest {
	s.SubDeviceName = &v
	return s
}

func (s *QueryCarProcessEventsRequest) SetSubIotId(v string) *QueryCarProcessEventsRequest {
	s.SubIotId = &v
	return s
}

func (s *QueryCarProcessEventsRequest) SetSubProductKey(v string) *QueryCarProcessEventsRequest {
	s.SubProductKey = &v
	return s
}

type QueryCarProcessEventsResponseBody struct {
	Code         *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryCarProcessEventsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCarProcessEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCarProcessEventsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCarProcessEventsResponseBody) SetCode(v string) *QueryCarProcessEventsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCarProcessEventsResponseBody) SetData(v *QueryCarProcessEventsResponseBodyData) *QueryCarProcessEventsResponseBody {
	s.Data = v
	return s
}

func (s *QueryCarProcessEventsResponseBody) SetErrorMessage(v string) *QueryCarProcessEventsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryCarProcessEventsResponseBody) SetRequestId(v string) *QueryCarProcessEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCarProcessEventsResponseBody) SetSuccess(v bool) *QueryCarProcessEventsResponseBody {
	s.Success = &v
	return s
}

type QueryCarProcessEventsResponseBodyData struct {
	CurrentPage *int32                                           `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageCount   *int32                                           `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageData    []*QueryCarProcessEventsResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	PageSize    *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total       *int32                                           `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryCarProcessEventsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryCarProcessEventsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCarProcessEventsResponseBodyData) SetCurrentPage(v int32) *QueryCarProcessEventsResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyData) SetPageCount(v int32) *QueryCarProcessEventsResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyData) SetPageData(v []*QueryCarProcessEventsResponseBodyDataPageData) *QueryCarProcessEventsResponseBodyData {
	s.PageData = v
	return s
}

func (s *QueryCarProcessEventsResponseBodyData) SetPageSize(v int32) *QueryCarProcessEventsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyData) SetTotal(v int32) *QueryCarProcessEventsResponseBodyData {
	s.Total = &v
	return s
}

type QueryCarProcessEventsResponseBodyDataPageData struct {
	ActionType        *int32  `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	AreaIndex         *int32  `json:"AreaIndex,omitempty" xml:"AreaIndex,omitempty"`
	Confidence        *int32  `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	EventId           *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	EventPicId        *string `json:"EventPicId,omitempty" xml:"EventPicId,omitempty"`
	EventPicUrl       *string `json:"EventPicUrl,omitempty" xml:"EventPicUrl,omitempty"`
	EventTime         *int64  `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	EventType         *int32  `json:"EventType,omitempty" xml:"EventType,omitempty"`
	IotId             *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	PlateNo           *string `json:"PlateNo,omitempty" xml:"PlateNo,omitempty"`
	SubDeviceName     *string `json:"SubDeviceName,omitempty" xml:"SubDeviceName,omitempty"`
	SubDeviceNickName *string `json:"SubDeviceNickName,omitempty" xml:"SubDeviceNickName,omitempty"`
	SubIotId          *string `json:"SubIotId,omitempty" xml:"SubIotId,omitempty"`
	SubProductKey     *string `json:"SubProductKey,omitempty" xml:"SubProductKey,omitempty"`
	TaskId            *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryCarProcessEventsResponseBodyDataPageData) String() string {
	return tea.Prettify(s)
}

func (s QueryCarProcessEventsResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetActionType(v int32) *QueryCarProcessEventsResponseBodyDataPageData {
	s.ActionType = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetAreaIndex(v int32) *QueryCarProcessEventsResponseBodyDataPageData {
	s.AreaIndex = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetConfidence(v int32) *QueryCarProcessEventsResponseBodyDataPageData {
	s.Confidence = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetEventId(v string) *QueryCarProcessEventsResponseBodyDataPageData {
	s.EventId = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetEventPicId(v string) *QueryCarProcessEventsResponseBodyDataPageData {
	s.EventPicId = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetEventPicUrl(v string) *QueryCarProcessEventsResponseBodyDataPageData {
	s.EventPicUrl = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetEventTime(v int64) *QueryCarProcessEventsResponseBodyDataPageData {
	s.EventTime = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetEventType(v int32) *QueryCarProcessEventsResponseBodyDataPageData {
	s.EventType = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetIotId(v string) *QueryCarProcessEventsResponseBodyDataPageData {
	s.IotId = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetPlateNo(v string) *QueryCarProcessEventsResponseBodyDataPageData {
	s.PlateNo = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetSubDeviceName(v string) *QueryCarProcessEventsResponseBodyDataPageData {
	s.SubDeviceName = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetSubDeviceNickName(v string) *QueryCarProcessEventsResponseBodyDataPageData {
	s.SubDeviceNickName = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetSubIotId(v string) *QueryCarProcessEventsResponseBodyDataPageData {
	s.SubIotId = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetSubProductKey(v string) *QueryCarProcessEventsResponseBodyDataPageData {
	s.SubProductKey = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetTaskId(v string) *QueryCarProcessEventsResponseBodyDataPageData {
	s.TaskId = &v
	return s
}

type QueryCarProcessEventsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCarProcessEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCarProcessEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCarProcessEventsResponse) GoString() string {
	return s.String()
}

func (s *QueryCarProcessEventsResponse) SetHeaders(v map[string]*string) *QueryCarProcessEventsResponse {
	s.Headers = v
	return s
}

func (s *QueryCarProcessEventsResponse) SetStatusCode(v int32) *QueryCarProcessEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCarProcessEventsResponse) SetBody(v *QueryCarProcessEventsResponseBody) *QueryCarProcessEventsResponse {
	s.Body = v
	return s
}

type QueryDeviceEventRequest struct {
	BeginTime     *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	EndTime       *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EventType     *int32  `json:"EventType,omitempty" xml:"EventType,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryDeviceEventRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceEventRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventRequest) SetBeginTime(v int64) *QueryDeviceEventRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryDeviceEventRequest) SetCurrentPage(v int32) *QueryDeviceEventRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryDeviceEventRequest) SetDeviceName(v string) *QueryDeviceEventRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryDeviceEventRequest) SetEndTime(v int64) *QueryDeviceEventRequest {
	s.EndTime = &v
	return s
}

func (s *QueryDeviceEventRequest) SetEventType(v int32) *QueryDeviceEventRequest {
	s.EventType = &v
	return s
}

func (s *QueryDeviceEventRequest) SetIotId(v string) *QueryDeviceEventRequest {
	s.IotId = &v
	return s
}

func (s *QueryDeviceEventRequest) SetIotInstanceId(v string) *QueryDeviceEventRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryDeviceEventRequest) SetPageSize(v int32) *QueryDeviceEventRequest {
	s.PageSize = &v
	return s
}

func (s *QueryDeviceEventRequest) SetProductKey(v string) *QueryDeviceEventRequest {
	s.ProductKey = &v
	return s
}

type QueryDeviceEventResponseBody struct {
	Code         *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryDeviceEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDeviceEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceEventResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventResponseBody) SetCode(v string) *QueryDeviceEventResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDeviceEventResponseBody) SetData(v *QueryDeviceEventResponseBodyData) *QueryDeviceEventResponseBody {
	s.Data = v
	return s
}

func (s *QueryDeviceEventResponseBody) SetErrorMessage(v string) *QueryDeviceEventResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryDeviceEventResponseBody) SetRequestId(v string) *QueryDeviceEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDeviceEventResponseBody) SetSuccess(v bool) *QueryDeviceEventResponseBody {
	s.Success = &v
	return s
}

type QueryDeviceEventResponseBodyData struct {
	List      []*QueryDeviceEventResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Page      *int32                                  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageCount *int32                                  `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageSize  *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int32                                  `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryDeviceEventResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventResponseBodyData) SetList(v []*QueryDeviceEventResponseBodyDataList) *QueryDeviceEventResponseBodyData {
	s.List = v
	return s
}

func (s *QueryDeviceEventResponseBodyData) SetPage(v int32) *QueryDeviceEventResponseBodyData {
	s.Page = &v
	return s
}

func (s *QueryDeviceEventResponseBodyData) SetPageCount(v int32) *QueryDeviceEventResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryDeviceEventResponseBodyData) SetPageSize(v int32) *QueryDeviceEventResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryDeviceEventResponseBodyData) SetTotal(v int32) *QueryDeviceEventResponseBodyData {
	s.Total = &v
	return s
}

type QueryDeviceEventResponseBodyDataList struct {
	EventData  *string `json:"EventData,omitempty" xml:"EventData,omitempty"`
	EventDesc  *string `json:"EventDesc,omitempty" xml:"EventDesc,omitempty"`
	EventId    *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	EventPicId *string `json:"EventPicId,omitempty" xml:"EventPicId,omitempty"`
	EventTime  *string `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	EventType  *int32  `json:"EventType,omitempty" xml:"EventType,omitempty"`
}

func (s QueryDeviceEventResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceEventResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventResponseBodyDataList) SetEventData(v string) *QueryDeviceEventResponseBodyDataList {
	s.EventData = &v
	return s
}

func (s *QueryDeviceEventResponseBodyDataList) SetEventDesc(v string) *QueryDeviceEventResponseBodyDataList {
	s.EventDesc = &v
	return s
}

func (s *QueryDeviceEventResponseBodyDataList) SetEventId(v string) *QueryDeviceEventResponseBodyDataList {
	s.EventId = &v
	return s
}

func (s *QueryDeviceEventResponseBodyDataList) SetEventPicId(v string) *QueryDeviceEventResponseBodyDataList {
	s.EventPicId = &v
	return s
}

func (s *QueryDeviceEventResponseBodyDataList) SetEventTime(v string) *QueryDeviceEventResponseBodyDataList {
	s.EventTime = &v
	return s
}

func (s *QueryDeviceEventResponseBodyDataList) SetEventType(v int32) *QueryDeviceEventResponseBodyDataList {
	s.EventType = &v
	return s
}

type QueryDeviceEventResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDeviceEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDeviceEventResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceEventResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventResponse) SetHeaders(v map[string]*string) *QueryDeviceEventResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceEventResponse) SetStatusCode(v int32) *QueryDeviceEventResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceEventResponse) SetBody(v *QueryDeviceEventResponseBody) *QueryDeviceEventResponse {
	s.Body = v
	return s
}

type QueryDeviceEventPictureRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	EventId       *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryDeviceEventPictureRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceEventPictureRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventPictureRequest) SetDeviceName(v string) *QueryDeviceEventPictureRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryDeviceEventPictureRequest) SetEventId(v string) *QueryDeviceEventPictureRequest {
	s.EventId = &v
	return s
}

func (s *QueryDeviceEventPictureRequest) SetIotId(v string) *QueryDeviceEventPictureRequest {
	s.IotId = &v
	return s
}

func (s *QueryDeviceEventPictureRequest) SetIotInstanceId(v string) *QueryDeviceEventPictureRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryDeviceEventPictureRequest) SetProductKey(v string) *QueryDeviceEventPictureRequest {
	s.ProductKey = &v
	return s
}

type QueryDeviceEventPictureResponseBody struct {
	Code         *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDeviceEventPictureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceEventPictureResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventPictureResponseBody) SetCode(v int32) *QueryDeviceEventPictureResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDeviceEventPictureResponseBody) SetData(v string) *QueryDeviceEventPictureResponseBody {
	s.Data = &v
	return s
}

func (s *QueryDeviceEventPictureResponseBody) SetErrorMessage(v string) *QueryDeviceEventPictureResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryDeviceEventPictureResponseBody) SetRequestId(v string) *QueryDeviceEventPictureResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDeviceEventPictureResponseBody) SetSuccess(v bool) *QueryDeviceEventPictureResponseBody {
	s.Success = &v
	return s
}

type QueryDeviceEventPictureResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDeviceEventPictureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDeviceEventPictureResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceEventPictureResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventPictureResponse) SetHeaders(v map[string]*string) *QueryDeviceEventPictureResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceEventPictureResponse) SetStatusCode(v int32) *QueryDeviceEventPictureResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceEventPictureResponse) SetBody(v *QueryDeviceEventPictureResponseBody) *QueryDeviceEventPictureResponse {
	s.Body = v
	return s
}

type QueryDeviceEventRecordRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	EventId       *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryDeviceEventRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceEventRecordRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventRecordRequest) SetDeviceName(v string) *QueryDeviceEventRecordRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryDeviceEventRecordRequest) SetEventId(v string) *QueryDeviceEventRecordRequest {
	s.EventId = &v
	return s
}

func (s *QueryDeviceEventRecordRequest) SetIotId(v string) *QueryDeviceEventRecordRequest {
	s.IotId = &v
	return s
}

func (s *QueryDeviceEventRecordRequest) SetIotInstanceId(v string) *QueryDeviceEventRecordRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryDeviceEventRecordRequest) SetProductKey(v string) *QueryDeviceEventRecordRequest {
	s.ProductKey = &v
	return s
}

type QueryDeviceEventRecordResponseBody struct {
	Code         *int32                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         []*QueryDeviceEventRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMessage *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDeviceEventRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceEventRecordResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventRecordResponseBody) SetCode(v int32) *QueryDeviceEventRecordResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDeviceEventRecordResponseBody) SetData(v []*QueryDeviceEventRecordResponseBodyData) *QueryDeviceEventRecordResponseBody {
	s.Data = v
	return s
}

func (s *QueryDeviceEventRecordResponseBody) SetErrorMessage(v string) *QueryDeviceEventRecordResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryDeviceEventRecordResponseBody) SetRequestId(v string) *QueryDeviceEventRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDeviceEventRecordResponseBody) SetSuccess(v bool) *QueryDeviceEventRecordResponseBody {
	s.Success = &v
	return s
}

type QueryDeviceEventRecordResponseBodyData struct {
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FileName  *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	VodUrl    *string `json:"VodUrl,omitempty" xml:"VodUrl,omitempty"`
}

func (s QueryDeviceEventRecordResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceEventRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventRecordResponseBodyData) SetBeginTime(v string) *QueryDeviceEventRecordResponseBodyData {
	s.BeginTime = &v
	return s
}

func (s *QueryDeviceEventRecordResponseBodyData) SetEndTime(v string) *QueryDeviceEventRecordResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *QueryDeviceEventRecordResponseBodyData) SetFileName(v string) *QueryDeviceEventRecordResponseBodyData {
	s.FileName = &v
	return s
}

func (s *QueryDeviceEventRecordResponseBodyData) SetVodUrl(v string) *QueryDeviceEventRecordResponseBodyData {
	s.VodUrl = &v
	return s
}

type QueryDeviceEventRecordResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDeviceEventRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDeviceEventRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceEventRecordResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventRecordResponse) SetHeaders(v map[string]*string) *QueryDeviceEventRecordResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceEventRecordResponse) SetStatusCode(v int32) *QueryDeviceEventRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceEventRecordResponse) SetBody(v *QueryDeviceEventRecordResponseBody) *QueryDeviceEventRecordResponse {
	s.Body = v
	return s
}

type QueryDevicePictureByListRequest struct {
	DeviceName    *string   `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	ExpireTime    *int32    `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	IotId         *string   `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string   `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	PictureIdList []*string `json:"PictureIdList,omitempty" xml:"PictureIdList,omitempty" type:"Repeated"`
	PictureType   *int32    `json:"PictureType,omitempty" xml:"PictureType,omitempty"`
	ProductKey    *string   `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	ThumbWidth    *int32    `json:"ThumbWidth,omitempty" xml:"ThumbWidth,omitempty"`
}

func (s QueryDevicePictureByListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePictureByListRequest) GoString() string {
	return s.String()
}

func (s *QueryDevicePictureByListRequest) SetDeviceName(v string) *QueryDevicePictureByListRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryDevicePictureByListRequest) SetExpireTime(v int32) *QueryDevicePictureByListRequest {
	s.ExpireTime = &v
	return s
}

func (s *QueryDevicePictureByListRequest) SetIotId(v string) *QueryDevicePictureByListRequest {
	s.IotId = &v
	return s
}

func (s *QueryDevicePictureByListRequest) SetIotInstanceId(v string) *QueryDevicePictureByListRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryDevicePictureByListRequest) SetPictureIdList(v []*string) *QueryDevicePictureByListRequest {
	s.PictureIdList = v
	return s
}

func (s *QueryDevicePictureByListRequest) SetPictureType(v int32) *QueryDevicePictureByListRequest {
	s.PictureType = &v
	return s
}

func (s *QueryDevicePictureByListRequest) SetProductKey(v string) *QueryDevicePictureByListRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryDevicePictureByListRequest) SetThumbWidth(v int32) *QueryDevicePictureByListRequest {
	s.ThumbWidth = &v
	return s
}

type QueryDevicePictureByListResponseBody struct {
	Code         *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryDevicePictureByListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDevicePictureByListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePictureByListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDevicePictureByListResponseBody) SetCode(v string) *QueryDevicePictureByListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDevicePictureByListResponseBody) SetData(v *QueryDevicePictureByListResponseBodyData) *QueryDevicePictureByListResponseBody {
	s.Data = v
	return s
}

func (s *QueryDevicePictureByListResponseBody) SetErrorMessage(v string) *QueryDevicePictureByListResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryDevicePictureByListResponseBody) SetRequestId(v string) *QueryDevicePictureByListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDevicePictureByListResponseBody) SetSuccess(v bool) *QueryDevicePictureByListResponseBody {
	s.Success = &v
	return s
}

type QueryDevicePictureByListResponseBodyData struct {
	PicData []*QueryDevicePictureByListResponseBodyDataPicData `json:"picData,omitempty" xml:"picData,omitempty" type:"Repeated"`
}

func (s QueryDevicePictureByListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePictureByListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDevicePictureByListResponseBodyData) SetPicData(v []*QueryDevicePictureByListResponseBodyDataPicData) *QueryDevicePictureByListResponseBodyData {
	s.PicData = v
	return s
}

type QueryDevicePictureByListResponseBodyDataPicData struct {
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	PicCreateTime *int64  `json:"PicCreateTime,omitempty" xml:"PicCreateTime,omitempty"`
	PicId         *string `json:"PicId,omitempty" xml:"PicId,omitempty"`
	PicUrl        *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	ThumbUrl      *string `json:"ThumbUrl,omitempty" xml:"ThumbUrl,omitempty"`
}

func (s QueryDevicePictureByListResponseBodyDataPicData) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePictureByListResponseBodyDataPicData) GoString() string {
	return s.String()
}

func (s *QueryDevicePictureByListResponseBodyDataPicData) SetIotId(v string) *QueryDevicePictureByListResponseBodyDataPicData {
	s.IotId = &v
	return s
}

func (s *QueryDevicePictureByListResponseBodyDataPicData) SetPicCreateTime(v int64) *QueryDevicePictureByListResponseBodyDataPicData {
	s.PicCreateTime = &v
	return s
}

func (s *QueryDevicePictureByListResponseBodyDataPicData) SetPicId(v string) *QueryDevicePictureByListResponseBodyDataPicData {
	s.PicId = &v
	return s
}

func (s *QueryDevicePictureByListResponseBodyDataPicData) SetPicUrl(v string) *QueryDevicePictureByListResponseBodyDataPicData {
	s.PicUrl = &v
	return s
}

func (s *QueryDevicePictureByListResponseBodyDataPicData) SetThumbUrl(v string) *QueryDevicePictureByListResponseBodyDataPicData {
	s.ThumbUrl = &v
	return s
}

type QueryDevicePictureByListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDevicePictureByListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDevicePictureByListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePictureByListResponse) GoString() string {
	return s.String()
}

func (s *QueryDevicePictureByListResponse) SetHeaders(v map[string]*string) *QueryDevicePictureByListResponse {
	s.Headers = v
	return s
}

func (s *QueryDevicePictureByListResponse) SetStatusCode(v int32) *QueryDevicePictureByListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDevicePictureByListResponse) SetBody(v *QueryDevicePictureByListResponseBody) *QueryDevicePictureByListResponse {
	s.Body = v
	return s
}

type QueryDevicePictureFileRequest struct {
	CaptureId     *string `json:"CaptureId,omitempty" xml:"CaptureId,omitempty"`
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	ExpireTime    *int32  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	PictureType   *int32  `json:"PictureType,omitempty" xml:"PictureType,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	ThumbWidth    *int32  `json:"ThumbWidth,omitempty" xml:"ThumbWidth,omitempty"`
}

func (s QueryDevicePictureFileRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePictureFileRequest) GoString() string {
	return s.String()
}

func (s *QueryDevicePictureFileRequest) SetCaptureId(v string) *QueryDevicePictureFileRequest {
	s.CaptureId = &v
	return s
}

func (s *QueryDevicePictureFileRequest) SetDeviceName(v string) *QueryDevicePictureFileRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryDevicePictureFileRequest) SetExpireTime(v int32) *QueryDevicePictureFileRequest {
	s.ExpireTime = &v
	return s
}

func (s *QueryDevicePictureFileRequest) SetIotId(v string) *QueryDevicePictureFileRequest {
	s.IotId = &v
	return s
}

func (s *QueryDevicePictureFileRequest) SetIotInstanceId(v string) *QueryDevicePictureFileRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryDevicePictureFileRequest) SetPictureType(v int32) *QueryDevicePictureFileRequest {
	s.PictureType = &v
	return s
}

func (s *QueryDevicePictureFileRequest) SetProductKey(v string) *QueryDevicePictureFileRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryDevicePictureFileRequest) SetThumbWidth(v int32) *QueryDevicePictureFileRequest {
	s.ThumbWidth = &v
	return s
}

type QueryDevicePictureFileResponseBody struct {
	Code         *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryDevicePictureFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDevicePictureFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePictureFileResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDevicePictureFileResponseBody) SetCode(v string) *QueryDevicePictureFileResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDevicePictureFileResponseBody) SetData(v *QueryDevicePictureFileResponseBodyData) *QueryDevicePictureFileResponseBody {
	s.Data = v
	return s
}

func (s *QueryDevicePictureFileResponseBody) SetErrorMessage(v string) *QueryDevicePictureFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryDevicePictureFileResponseBody) SetRequestId(v string) *QueryDevicePictureFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDevicePictureFileResponseBody) SetSuccess(v bool) *QueryDevicePictureFileResponseBody {
	s.Success = &v
	return s
}

type QueryDevicePictureFileResponseBodyData struct {
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	PicCreateTime *int64  `json:"PicCreateTime,omitempty" xml:"PicCreateTime,omitempty"`
	PicId         *string `json:"PicId,omitempty" xml:"PicId,omitempty"`
	PicUrl        *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	ThumbUrl      *string `json:"ThumbUrl,omitempty" xml:"ThumbUrl,omitempty"`
}

func (s QueryDevicePictureFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePictureFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDevicePictureFileResponseBodyData) SetIotId(v string) *QueryDevicePictureFileResponseBodyData {
	s.IotId = &v
	return s
}

func (s *QueryDevicePictureFileResponseBodyData) SetPicCreateTime(v int64) *QueryDevicePictureFileResponseBodyData {
	s.PicCreateTime = &v
	return s
}

func (s *QueryDevicePictureFileResponseBodyData) SetPicId(v string) *QueryDevicePictureFileResponseBodyData {
	s.PicId = &v
	return s
}

func (s *QueryDevicePictureFileResponseBodyData) SetPicUrl(v string) *QueryDevicePictureFileResponseBodyData {
	s.PicUrl = &v
	return s
}

func (s *QueryDevicePictureFileResponseBodyData) SetThumbUrl(v string) *QueryDevicePictureFileResponseBodyData {
	s.ThumbUrl = &v
	return s
}

type QueryDevicePictureFileResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDevicePictureFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDevicePictureFileResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePictureFileResponse) GoString() string {
	return s.String()
}

func (s *QueryDevicePictureFileResponse) SetHeaders(v map[string]*string) *QueryDevicePictureFileResponse {
	s.Headers = v
	return s
}

func (s *QueryDevicePictureFileResponse) SetStatusCode(v int32) *QueryDevicePictureFileResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDevicePictureFileResponse) SetBody(v *QueryDevicePictureFileResponseBody) *QueryDevicePictureFileResponse {
	s.Body = v
	return s
}

type QueryDevicePictureLifeCycleRequest struct {
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
}

func (s QueryDevicePictureLifeCycleRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePictureLifeCycleRequest) GoString() string {
	return s.String()
}

func (s *QueryDevicePictureLifeCycleRequest) SetIotId(v string) *QueryDevicePictureLifeCycleRequest {
	s.IotId = &v
	return s
}

type QueryDevicePictureLifeCycleResponseBody struct {
	Code         *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryDevicePictureLifeCycleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDevicePictureLifeCycleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePictureLifeCycleResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDevicePictureLifeCycleResponseBody) SetCode(v string) *QueryDevicePictureLifeCycleResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDevicePictureLifeCycleResponseBody) SetData(v *QueryDevicePictureLifeCycleResponseBodyData) *QueryDevicePictureLifeCycleResponseBody {
	s.Data = v
	return s
}

func (s *QueryDevicePictureLifeCycleResponseBody) SetErrorMessage(v string) *QueryDevicePictureLifeCycleResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryDevicePictureLifeCycleResponseBody) SetRequestId(v string) *QueryDevicePictureLifeCycleResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDevicePictureLifeCycleResponseBody) SetSuccess(v bool) *QueryDevicePictureLifeCycleResponseBody {
	s.Success = &v
	return s
}

type QueryDevicePictureLifeCycleResponseBodyData struct {
	Day   *int32  `json:"Day,omitempty" xml:"Day,omitempty"`
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
}

func (s QueryDevicePictureLifeCycleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePictureLifeCycleResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDevicePictureLifeCycleResponseBodyData) SetDay(v int32) *QueryDevicePictureLifeCycleResponseBodyData {
	s.Day = &v
	return s
}

func (s *QueryDevicePictureLifeCycleResponseBodyData) SetIotId(v string) *QueryDevicePictureLifeCycleResponseBodyData {
	s.IotId = &v
	return s
}

type QueryDevicePictureLifeCycleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDevicePictureLifeCycleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDevicePictureLifeCycleResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePictureLifeCycleResponse) GoString() string {
	return s.String()
}

func (s *QueryDevicePictureLifeCycleResponse) SetHeaders(v map[string]*string) *QueryDevicePictureLifeCycleResponse {
	s.Headers = v
	return s
}

func (s *QueryDevicePictureLifeCycleResponse) SetStatusCode(v int32) *QueryDevicePictureLifeCycleResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDevicePictureLifeCycleResponse) SetBody(v *QueryDevicePictureLifeCycleResponseBody) *QueryDevicePictureLifeCycleResponse {
	s.Body = v
	return s
}

type QueryDeviceRecordLifeCycleRequest struct {
	DeviceList []*string `json:"DeviceList,omitempty" xml:"DeviceList,omitempty" type:"Repeated"`
}

func (s QueryDeviceRecordLifeCycleRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceRecordLifeCycleRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceRecordLifeCycleRequest) SetDeviceList(v []*string) *QueryDeviceRecordLifeCycleRequest {
	s.DeviceList = v
	return s
}

type QueryDeviceRecordLifeCycleResponseBody struct {
	Code         *int32                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         []*QueryDeviceRecordLifeCycleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMessage *string                                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDeviceRecordLifeCycleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceRecordLifeCycleResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceRecordLifeCycleResponseBody) SetCode(v int32) *QueryDeviceRecordLifeCycleResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDeviceRecordLifeCycleResponseBody) SetData(v []*QueryDeviceRecordLifeCycleResponseBodyData) *QueryDeviceRecordLifeCycleResponseBody {
	s.Data = v
	return s
}

func (s *QueryDeviceRecordLifeCycleResponseBody) SetErrorMessage(v string) *QueryDeviceRecordLifeCycleResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryDeviceRecordLifeCycleResponseBody) SetRequestId(v string) *QueryDeviceRecordLifeCycleResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDeviceRecordLifeCycleResponseBody) SetSuccess(v bool) *QueryDeviceRecordLifeCycleResponseBody {
	s.Success = &v
	return s
}

type QueryDeviceRecordLifeCycleResponseBodyData struct {
	Day   *int32  `json:"Day,omitempty" xml:"Day,omitempty"`
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
}

func (s QueryDeviceRecordLifeCycleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceRecordLifeCycleResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDeviceRecordLifeCycleResponseBodyData) SetDay(v int32) *QueryDeviceRecordLifeCycleResponseBodyData {
	s.Day = &v
	return s
}

func (s *QueryDeviceRecordLifeCycleResponseBodyData) SetIotId(v string) *QueryDeviceRecordLifeCycleResponseBodyData {
	s.IotId = &v
	return s
}

type QueryDeviceRecordLifeCycleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDeviceRecordLifeCycleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDeviceRecordLifeCycleResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceRecordLifeCycleResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceRecordLifeCycleResponse) SetHeaders(v map[string]*string) *QueryDeviceRecordLifeCycleResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceRecordLifeCycleResponse) SetStatusCode(v int32) *QueryDeviceRecordLifeCycleResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceRecordLifeCycleResponse) SetBody(v *QueryDeviceRecordLifeCycleResponseBody) *QueryDeviceRecordLifeCycleResponse {
	s.Body = v
	return s
}

type QueryDeviceVodUrlRequest struct {
	DeviceName       *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	EnableStun       *bool   `json:"EnableStun,omitempty" xml:"EnableStun,omitempty"`
	EncryptType      *int32  `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	FileName         *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	IotId            *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId    *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	PlayUnLimited    *bool   `json:"PlayUnLimited,omitempty" xml:"PlayUnLimited,omitempty"`
	ProductKey       *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	Scheme           *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
	SeekTime         *int32  `json:"SeekTime,omitempty" xml:"SeekTime,omitempty"`
	ShouldEncrypt    *bool   `json:"ShouldEncrypt,omitempty" xml:"ShouldEncrypt,omitempty"`
	UrlValidDuration *int32  `json:"UrlValidDuration,omitempty" xml:"UrlValidDuration,omitempty"`
}

func (s QueryDeviceVodUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceVodUrlRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceVodUrlRequest) SetDeviceName(v string) *QueryDeviceVodUrlRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) SetEnableStun(v bool) *QueryDeviceVodUrlRequest {
	s.EnableStun = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) SetEncryptType(v int32) *QueryDeviceVodUrlRequest {
	s.EncryptType = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) SetFileName(v string) *QueryDeviceVodUrlRequest {
	s.FileName = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) SetIotId(v string) *QueryDeviceVodUrlRequest {
	s.IotId = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) SetIotInstanceId(v string) *QueryDeviceVodUrlRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) SetPlayUnLimited(v bool) *QueryDeviceVodUrlRequest {
	s.PlayUnLimited = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) SetProductKey(v string) *QueryDeviceVodUrlRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) SetScheme(v string) *QueryDeviceVodUrlRequest {
	s.Scheme = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) SetSeekTime(v int32) *QueryDeviceVodUrlRequest {
	s.SeekTime = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) SetShouldEncrypt(v bool) *QueryDeviceVodUrlRequest {
	s.ShouldEncrypt = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) SetUrlValidDuration(v int32) *QueryDeviceVodUrlRequest {
	s.UrlValidDuration = &v
	return s
}

type QueryDeviceVodUrlResponseBody struct {
	Code         *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryDeviceVodUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDeviceVodUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceVodUrlResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceVodUrlResponseBody) SetCode(v string) *QueryDeviceVodUrlResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDeviceVodUrlResponseBody) SetData(v *QueryDeviceVodUrlResponseBodyData) *QueryDeviceVodUrlResponseBody {
	s.Data = v
	return s
}

func (s *QueryDeviceVodUrlResponseBody) SetErrorMessage(v string) *QueryDeviceVodUrlResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryDeviceVodUrlResponseBody) SetRequestId(v string) *QueryDeviceVodUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDeviceVodUrlResponseBody) SetSuccess(v bool) *QueryDeviceVodUrlResponseBody {
	s.Success = &v
	return s
}

type QueryDeviceVodUrlResponseBodyData struct {
	DecryptKey *string `json:"DecryptKey,omitempty" xml:"DecryptKey,omitempty"`
	StunInfo   *string `json:"StunInfo,omitempty" xml:"StunInfo,omitempty"`
	VodUrl     *string `json:"VodUrl,omitempty" xml:"VodUrl,omitempty"`
}

func (s QueryDeviceVodUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceVodUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDeviceVodUrlResponseBodyData) SetDecryptKey(v string) *QueryDeviceVodUrlResponseBodyData {
	s.DecryptKey = &v
	return s
}

func (s *QueryDeviceVodUrlResponseBodyData) SetStunInfo(v string) *QueryDeviceVodUrlResponseBodyData {
	s.StunInfo = &v
	return s
}

func (s *QueryDeviceVodUrlResponseBodyData) SetVodUrl(v string) *QueryDeviceVodUrlResponseBodyData {
	s.VodUrl = &v
	return s
}

type QueryDeviceVodUrlResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDeviceVodUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDeviceVodUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceVodUrlResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceVodUrlResponse) SetHeaders(v map[string]*string) *QueryDeviceVodUrlResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceVodUrlResponse) SetStatusCode(v int32) *QueryDeviceVodUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceVodUrlResponse) SetBody(v *QueryDeviceVodUrlResponseBody) *QueryDeviceVodUrlResponse {
	s.Body = v
	return s
}

type QueryDeviceVodUrlByTimeRequest struct {
	BeginTime        *int32  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	DeviceName       *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	EnableStun       *bool   `json:"EnableStun,omitempty" xml:"EnableStun,omitempty"`
	EncryptType      *int32  `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	EndTime          *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IotId            *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId    *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	PlayUnLimited    *bool   `json:"PlayUnLimited,omitempty" xml:"PlayUnLimited,omitempty"`
	ProductKey       *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	Scheme           *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
	SeekTime         *int32  `json:"SeekTime,omitempty" xml:"SeekTime,omitempty"`
	ShouldEncrypt    *bool   `json:"ShouldEncrypt,omitempty" xml:"ShouldEncrypt,omitempty"`
	UrlValidDuration *int32  `json:"UrlValidDuration,omitempty" xml:"UrlValidDuration,omitempty"`
}

func (s QueryDeviceVodUrlByTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceVodUrlByTimeRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceVodUrlByTimeRequest) SetBeginTime(v int32) *QueryDeviceVodUrlByTimeRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeRequest) SetDeviceName(v string) *QueryDeviceVodUrlByTimeRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeRequest) SetEnableStun(v bool) *QueryDeviceVodUrlByTimeRequest {
	s.EnableStun = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeRequest) SetEncryptType(v int32) *QueryDeviceVodUrlByTimeRequest {
	s.EncryptType = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeRequest) SetEndTime(v int32) *QueryDeviceVodUrlByTimeRequest {
	s.EndTime = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeRequest) SetIotId(v string) *QueryDeviceVodUrlByTimeRequest {
	s.IotId = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeRequest) SetIotInstanceId(v string) *QueryDeviceVodUrlByTimeRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeRequest) SetPlayUnLimited(v bool) *QueryDeviceVodUrlByTimeRequest {
	s.PlayUnLimited = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeRequest) SetProductKey(v string) *QueryDeviceVodUrlByTimeRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeRequest) SetScheme(v string) *QueryDeviceVodUrlByTimeRequest {
	s.Scheme = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeRequest) SetSeekTime(v int32) *QueryDeviceVodUrlByTimeRequest {
	s.SeekTime = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeRequest) SetShouldEncrypt(v bool) *QueryDeviceVodUrlByTimeRequest {
	s.ShouldEncrypt = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeRequest) SetUrlValidDuration(v int32) *QueryDeviceVodUrlByTimeRequest {
	s.UrlValidDuration = &v
	return s
}

type QueryDeviceVodUrlByTimeResponseBody struct {
	Code         *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryDeviceVodUrlByTimeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDeviceVodUrlByTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceVodUrlByTimeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceVodUrlByTimeResponseBody) SetCode(v string) *QueryDeviceVodUrlByTimeResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeResponseBody) SetData(v *QueryDeviceVodUrlByTimeResponseBodyData) *QueryDeviceVodUrlByTimeResponseBody {
	s.Data = v
	return s
}

func (s *QueryDeviceVodUrlByTimeResponseBody) SetErrorMessage(v string) *QueryDeviceVodUrlByTimeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeResponseBody) SetRequestId(v string) *QueryDeviceVodUrlByTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeResponseBody) SetSuccess(v bool) *QueryDeviceVodUrlByTimeResponseBody {
	s.Success = &v
	return s
}

type QueryDeviceVodUrlByTimeResponseBodyData struct {
	DecryptKey *string `json:"DecryptKey,omitempty" xml:"DecryptKey,omitempty"`
	StunInfo   *string `json:"StunInfo,omitempty" xml:"StunInfo,omitempty"`
	VodUrl     *string `json:"VodUrl,omitempty" xml:"VodUrl,omitempty"`
}

func (s QueryDeviceVodUrlByTimeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceVodUrlByTimeResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDeviceVodUrlByTimeResponseBodyData) SetDecryptKey(v string) *QueryDeviceVodUrlByTimeResponseBodyData {
	s.DecryptKey = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeResponseBodyData) SetStunInfo(v string) *QueryDeviceVodUrlByTimeResponseBodyData {
	s.StunInfo = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeResponseBodyData) SetVodUrl(v string) *QueryDeviceVodUrlByTimeResponseBodyData {
	s.VodUrl = &v
	return s
}

type QueryDeviceVodUrlByTimeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDeviceVodUrlByTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDeviceVodUrlByTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceVodUrlByTimeResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceVodUrlByTimeResponse) SetHeaders(v map[string]*string) *QueryDeviceVodUrlByTimeResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceVodUrlByTimeResponse) SetStatusCode(v int32) *QueryDeviceVodUrlByTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeResponse) SetBody(v *QueryDeviceVodUrlByTimeResponseBody) *QueryDeviceVodUrlByTimeResponse {
	s.Body = v
	return s
}

type QueryEventRecordPlanDetailRequest struct {
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s QueryEventRecordPlanDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDetailRequest) SetPlanId(v string) *QueryEventRecordPlanDetailRequest {
	s.PlanId = &v
	return s
}

type QueryEventRecordPlanDetailResponseBody struct {
	Code         *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryEventRecordPlanDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                     `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryEventRecordPlanDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDetailResponseBody) SetCode(v string) *QueryEventRecordPlanDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryEventRecordPlanDetailResponseBody) SetData(v *QueryEventRecordPlanDetailResponseBodyData) *QueryEventRecordPlanDetailResponseBody {
	s.Data = v
	return s
}

func (s *QueryEventRecordPlanDetailResponseBody) SetErrorMessage(v string) *QueryEventRecordPlanDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryEventRecordPlanDetailResponseBody) SetRequestId(v string) *QueryEventRecordPlanDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEventRecordPlanDetailResponseBody) SetSuccess(v bool) *QueryEventRecordPlanDetailResponseBody {
	s.Success = &v
	return s
}

type QueryEventRecordPlanDetailResponseBodyData struct {
	Name              *string                                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	PlanId            *string                                                 `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	PreRecordDuration *int32                                                  `json:"PreRecordDuration,omitempty" xml:"PreRecordDuration,omitempty"`
	RecordDuration    *int32                                                  `json:"RecordDuration,omitempty" xml:"RecordDuration,omitempty"`
	TemplateId        *string                                                 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateInfo      *QueryEventRecordPlanDetailResponseBodyDataTemplateInfo `json:"TemplateInfo,omitempty" xml:"TemplateInfo,omitempty" type:"Struct"`
}

func (s QueryEventRecordPlanDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDetailResponseBodyData) SetName(v string) *QueryEventRecordPlanDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryEventRecordPlanDetailResponseBodyData) SetPlanId(v string) *QueryEventRecordPlanDetailResponseBodyData {
	s.PlanId = &v
	return s
}

func (s *QueryEventRecordPlanDetailResponseBodyData) SetPreRecordDuration(v int32) *QueryEventRecordPlanDetailResponseBodyData {
	s.PreRecordDuration = &v
	return s
}

func (s *QueryEventRecordPlanDetailResponseBodyData) SetRecordDuration(v int32) *QueryEventRecordPlanDetailResponseBodyData {
	s.RecordDuration = &v
	return s
}

func (s *QueryEventRecordPlanDetailResponseBodyData) SetTemplateId(v string) *QueryEventRecordPlanDetailResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *QueryEventRecordPlanDetailResponseBodyData) SetTemplateInfo(v *QueryEventRecordPlanDetailResponseBodyDataTemplateInfo) *QueryEventRecordPlanDetailResponseBodyData {
	s.TemplateInfo = v
	return s
}

type QueryEventRecordPlanDetailResponseBodyDataTemplateInfo struct {
	AllDay          *int32                                                                   `json:"AllDay,omitempty" xml:"AllDay,omitempty"`
	Default         *int32                                                                   `json:"Default,omitempty" xml:"Default,omitempty"`
	Name            *string                                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateId      *string                                                                  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TimeSectionList []*QueryEventRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList `json:"TimeSectionList,omitempty" xml:"TimeSectionList,omitempty" type:"Repeated"`
}

func (s QueryEventRecordPlanDetailResponseBodyDataTemplateInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDetailResponseBodyDataTemplateInfo) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDetailResponseBodyDataTemplateInfo) SetAllDay(v int32) *QueryEventRecordPlanDetailResponseBodyDataTemplateInfo {
	s.AllDay = &v
	return s
}

func (s *QueryEventRecordPlanDetailResponseBodyDataTemplateInfo) SetDefault(v int32) *QueryEventRecordPlanDetailResponseBodyDataTemplateInfo {
	s.Default = &v
	return s
}

func (s *QueryEventRecordPlanDetailResponseBodyDataTemplateInfo) SetName(v string) *QueryEventRecordPlanDetailResponseBodyDataTemplateInfo {
	s.Name = &v
	return s
}

func (s *QueryEventRecordPlanDetailResponseBodyDataTemplateInfo) SetTemplateId(v string) *QueryEventRecordPlanDetailResponseBodyDataTemplateInfo {
	s.TemplateId = &v
	return s
}

func (s *QueryEventRecordPlanDetailResponseBodyDataTemplateInfo) SetTimeSectionList(v []*QueryEventRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList) *QueryEventRecordPlanDetailResponseBodyDataTemplateInfo {
	s.TimeSectionList = v
	return s
}

type QueryEventRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList struct {
	Begin     *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	DayOfWeek *int32 `json:"DayOfWeek,omitempty" xml:"DayOfWeek,omitempty"`
	End       *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s QueryEventRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList) SetBegin(v int32) *QueryEventRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList {
	s.Begin = &v
	return s
}

func (s *QueryEventRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList) SetDayOfWeek(v int32) *QueryEventRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList {
	s.DayOfWeek = &v
	return s
}

func (s *QueryEventRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList) SetEnd(v int32) *QueryEventRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList {
	s.End = &v
	return s
}

type QueryEventRecordPlanDetailResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEventRecordPlanDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEventRecordPlanDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDetailResponse) SetHeaders(v map[string]*string) *QueryEventRecordPlanDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryEventRecordPlanDetailResponse) SetStatusCode(v int32) *QueryEventRecordPlanDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEventRecordPlanDetailResponse) SetBody(v *QueryEventRecordPlanDetailResponseBody) *QueryEventRecordPlanDetailResponse {
	s.Body = v
	return s
}

type QueryEventRecordPlanDeviceByDeviceRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	StreamType    *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s QueryEventRecordPlanDeviceByDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDeviceByDeviceRequest) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDeviceByDeviceRequest) SetDeviceName(v string) *QueryEventRecordPlanDeviceByDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceRequest) SetIotId(v string) *QueryEventRecordPlanDeviceByDeviceRequest {
	s.IotId = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceRequest) SetIotInstanceId(v string) *QueryEventRecordPlanDeviceByDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceRequest) SetProductKey(v string) *QueryEventRecordPlanDeviceByDeviceRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceRequest) SetStreamType(v int32) *QueryEventRecordPlanDeviceByDeviceRequest {
	s.StreamType = &v
	return s
}

type QueryEventRecordPlanDeviceByDeviceResponseBody struct {
	Code         *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryEventRecordPlanDeviceByDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryEventRecordPlanDeviceByDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDeviceByDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBody) SetCode(v string) *QueryEventRecordPlanDeviceByDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBody) SetData(v *QueryEventRecordPlanDeviceByDeviceResponseBodyData) *QueryEventRecordPlanDeviceByDeviceResponseBody {
	s.Data = v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBody) SetErrorMessage(v string) *QueryEventRecordPlanDeviceByDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBody) SetRequestId(v string) *QueryEventRecordPlanDeviceByDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBody) SetSuccess(v bool) *QueryEventRecordPlanDeviceByDeviceResponseBody {
	s.Success = &v
	return s
}

type QueryEventRecordPlanDeviceByDeviceResponseBodyData struct {
	Name              *string                                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	PlanId            *string                                                         `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	PreRecordDuration *int32                                                          `json:"PreRecordDuration,omitempty" xml:"PreRecordDuration,omitempty"`
	RecordDuration    *int32                                                          `json:"RecordDuration,omitempty" xml:"RecordDuration,omitempty"`
	TemplateId        *string                                                         `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateInfo      *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo `json:"TemplateInfo,omitempty" xml:"TemplateInfo,omitempty" type:"Struct"`
}

func (s QueryEventRecordPlanDeviceByDeviceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDeviceByDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBodyData) SetName(v string) *QueryEventRecordPlanDeviceByDeviceResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBodyData) SetPlanId(v string) *QueryEventRecordPlanDeviceByDeviceResponseBodyData {
	s.PlanId = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBodyData) SetPreRecordDuration(v int32) *QueryEventRecordPlanDeviceByDeviceResponseBodyData {
	s.PreRecordDuration = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBodyData) SetRecordDuration(v int32) *QueryEventRecordPlanDeviceByDeviceResponseBodyData {
	s.RecordDuration = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBodyData) SetTemplateId(v string) *QueryEventRecordPlanDeviceByDeviceResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBodyData) SetTemplateInfo(v *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo) *QueryEventRecordPlanDeviceByDeviceResponseBodyData {
	s.TemplateInfo = v
	return s
}

type QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo struct {
	AllDay          *int32                                                                           `json:"AllDay,omitempty" xml:"AllDay,omitempty"`
	Default         *int32                                                                           `json:"Default,omitempty" xml:"Default,omitempty"`
	Name            *string                                                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateId      *string                                                                          `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TimeSectionList []*QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList `json:"TimeSectionList,omitempty" xml:"TimeSectionList,omitempty" type:"Repeated"`
}

func (s QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo) SetAllDay(v int32) *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo {
	s.AllDay = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo) SetDefault(v int32) *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo {
	s.Default = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo) SetName(v string) *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo {
	s.Name = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo) SetTemplateId(v string) *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo {
	s.TemplateId = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo) SetTimeSectionList(v []*QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList) *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo {
	s.TimeSectionList = v
	return s
}

type QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList struct {
	Begin     *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	DayOfWeek *int32 `json:"DayOfWeek,omitempty" xml:"DayOfWeek,omitempty"`
	End       *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList) SetBegin(v int32) *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList {
	s.Begin = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList) SetDayOfWeek(v int32) *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList {
	s.DayOfWeek = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList) SetEnd(v int32) *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList {
	s.End = &v
	return s
}

type QueryEventRecordPlanDeviceByDeviceResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEventRecordPlanDeviceByDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEventRecordPlanDeviceByDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDeviceByDeviceResponse) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDeviceByDeviceResponse) SetHeaders(v map[string]*string) *QueryEventRecordPlanDeviceByDeviceResponse {
	s.Headers = v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponse) SetStatusCode(v int32) *QueryEventRecordPlanDeviceByDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponse) SetBody(v *QueryEventRecordPlanDeviceByDeviceResponseBody) *QueryEventRecordPlanDeviceByDeviceResponse {
	s.Body = v
	return s
}

type QueryEventRecordPlanDeviceByPlanRequest struct {
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PlanId      *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s QueryEventRecordPlanDeviceByPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDeviceByPlanRequest) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDeviceByPlanRequest) SetCurrentPage(v int32) *QueryEventRecordPlanDeviceByPlanRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanRequest) SetPageSize(v int32) *QueryEventRecordPlanDeviceByPlanRequest {
	s.PageSize = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanRequest) SetPlanId(v string) *QueryEventRecordPlanDeviceByPlanRequest {
	s.PlanId = &v
	return s
}

type QueryEventRecordPlanDeviceByPlanResponseBody struct {
	Code         *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryEventRecordPlanDeviceByPlanResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryEventRecordPlanDeviceByPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDeviceByPlanResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDeviceByPlanResponseBody) SetCode(v string) *QueryEventRecordPlanDeviceByPlanResponseBody {
	s.Code = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanResponseBody) SetData(v *QueryEventRecordPlanDeviceByPlanResponseBodyData) *QueryEventRecordPlanDeviceByPlanResponseBody {
	s.Data = v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanResponseBody) SetErrorMessage(v string) *QueryEventRecordPlanDeviceByPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanResponseBody) SetRequestId(v string) *QueryEventRecordPlanDeviceByPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanResponseBody) SetSuccess(v bool) *QueryEventRecordPlanDeviceByPlanResponseBody {
	s.Success = &v
	return s
}

type QueryEventRecordPlanDeviceByPlanResponseBodyData struct {
	List      []*QueryEventRecordPlanDeviceByPlanResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Page      *int32                                                  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageCount *int32                                                  `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageSize  *int32                                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int32                                                  `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryEventRecordPlanDeviceByPlanResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDeviceByPlanResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDeviceByPlanResponseBodyData) SetList(v []*QueryEventRecordPlanDeviceByPlanResponseBodyDataList) *QueryEventRecordPlanDeviceByPlanResponseBodyData {
	s.List = v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanResponseBodyData) SetPage(v int32) *QueryEventRecordPlanDeviceByPlanResponseBodyData {
	s.Page = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanResponseBodyData) SetPageCount(v int32) *QueryEventRecordPlanDeviceByPlanResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanResponseBodyData) SetPageSize(v int32) *QueryEventRecordPlanDeviceByPlanResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanResponseBodyData) SetTotal(v int32) *QueryEventRecordPlanDeviceByPlanResponseBodyData {
	s.Total = &v
	return s
}

type QueryEventRecordPlanDeviceByPlanResponseBodyDataList struct {
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	StreamType *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s QueryEventRecordPlanDeviceByPlanResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDeviceByPlanResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDeviceByPlanResponseBodyDataList) SetIotId(v string) *QueryEventRecordPlanDeviceByPlanResponseBodyDataList {
	s.IotId = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanResponseBodyDataList) SetStreamType(v int32) *QueryEventRecordPlanDeviceByPlanResponseBodyDataList {
	s.StreamType = &v
	return s
}

type QueryEventRecordPlanDeviceByPlanResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEventRecordPlanDeviceByPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEventRecordPlanDeviceByPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDeviceByPlanResponse) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDeviceByPlanResponse) SetHeaders(v map[string]*string) *QueryEventRecordPlanDeviceByPlanResponse {
	s.Headers = v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanResponse) SetStatusCode(v int32) *QueryEventRecordPlanDeviceByPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanResponse) SetBody(v *QueryEventRecordPlanDeviceByPlanResponseBody) *QueryEventRecordPlanDeviceByPlanResponse {
	s.Body = v
	return s
}

type QueryEventRecordPlansRequest struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryEventRecordPlansRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlansRequest) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlansRequest) SetCurrentPage(v int32) *QueryEventRecordPlansRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryEventRecordPlansRequest) SetPageSize(v int32) *QueryEventRecordPlansRequest {
	s.PageSize = &v
	return s
}

type QueryEventRecordPlansResponseBody struct {
	Code         *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryEventRecordPlansResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryEventRecordPlansResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlansResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlansResponseBody) SetCode(v string) *QueryEventRecordPlansResponseBody {
	s.Code = &v
	return s
}

func (s *QueryEventRecordPlansResponseBody) SetData(v *QueryEventRecordPlansResponseBodyData) *QueryEventRecordPlansResponseBody {
	s.Data = v
	return s
}

func (s *QueryEventRecordPlansResponseBody) SetErrorMessage(v string) *QueryEventRecordPlansResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryEventRecordPlansResponseBody) SetRequestId(v string) *QueryEventRecordPlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEventRecordPlansResponseBody) SetSuccess(v bool) *QueryEventRecordPlansResponseBody {
	s.Success = &v
	return s
}

type QueryEventRecordPlansResponseBodyData struct {
	List      []*QueryEventRecordPlansResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Page      *int32                                       `json:"Page,omitempty" xml:"Page,omitempty"`
	PageCount *int32                                       `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageSize  *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int32                                       `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryEventRecordPlansResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlansResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlansResponseBodyData) SetList(v []*QueryEventRecordPlansResponseBodyDataList) *QueryEventRecordPlansResponseBodyData {
	s.List = v
	return s
}

func (s *QueryEventRecordPlansResponseBodyData) SetPage(v int32) *QueryEventRecordPlansResponseBodyData {
	s.Page = &v
	return s
}

func (s *QueryEventRecordPlansResponseBodyData) SetPageCount(v int32) *QueryEventRecordPlansResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryEventRecordPlansResponseBodyData) SetPageSize(v int32) *QueryEventRecordPlansResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryEventRecordPlansResponseBodyData) SetTotal(v int32) *QueryEventRecordPlansResponseBodyData {
	s.Total = &v
	return s
}

type QueryEventRecordPlansResponseBodyDataList struct {
	EventType         *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PlanId            *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	PreRecordDuration *int32  `json:"PreRecordDuration,omitempty" xml:"PreRecordDuration,omitempty"`
	RecordDuration    *int32  `json:"RecordDuration,omitempty" xml:"RecordDuration,omitempty"`
	TemplateId        *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s QueryEventRecordPlansResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlansResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlansResponseBodyDataList) SetEventType(v string) *QueryEventRecordPlansResponseBodyDataList {
	s.EventType = &v
	return s
}

func (s *QueryEventRecordPlansResponseBodyDataList) SetName(v string) *QueryEventRecordPlansResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *QueryEventRecordPlansResponseBodyDataList) SetPlanId(v string) *QueryEventRecordPlansResponseBodyDataList {
	s.PlanId = &v
	return s
}

func (s *QueryEventRecordPlansResponseBodyDataList) SetPreRecordDuration(v int32) *QueryEventRecordPlansResponseBodyDataList {
	s.PreRecordDuration = &v
	return s
}

func (s *QueryEventRecordPlansResponseBodyDataList) SetRecordDuration(v int32) *QueryEventRecordPlansResponseBodyDataList {
	s.RecordDuration = &v
	return s
}

func (s *QueryEventRecordPlansResponseBodyDataList) SetTemplateId(v string) *QueryEventRecordPlansResponseBodyDataList {
	s.TemplateId = &v
	return s
}

type QueryEventRecordPlansResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEventRecordPlansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEventRecordPlansResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlansResponse) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlansResponse) SetHeaders(v map[string]*string) *QueryEventRecordPlansResponse {
	s.Headers = v
	return s
}

func (s *QueryEventRecordPlansResponse) SetStatusCode(v int32) *QueryEventRecordPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEventRecordPlansResponse) SetBody(v *QueryEventRecordPlansResponseBody) *QueryEventRecordPlansResponse {
	s.Body = v
	return s
}

type QueryFaceAllDeviceGroupRequest struct {
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	IsolationId   *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	PageNo        *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryFaceAllDeviceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceAllDeviceGroupRequest) SetIotInstanceId(v string) *QueryFaceAllDeviceGroupRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryFaceAllDeviceGroupRequest) SetIsolationId(v string) *QueryFaceAllDeviceGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceAllDeviceGroupRequest) SetPageNo(v int32) *QueryFaceAllDeviceGroupRequest {
	s.PageNo = &v
	return s
}

func (s *QueryFaceAllDeviceGroupRequest) SetPageSize(v int32) *QueryFaceAllDeviceGroupRequest {
	s.PageSize = &v
	return s
}

type QueryFaceAllDeviceGroupResponseBody struct {
	Code         *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryFaceAllDeviceGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceAllDeviceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceAllDeviceGroupResponseBody) SetCode(v string) *QueryFaceAllDeviceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceAllDeviceGroupResponseBody) SetData(v *QueryFaceAllDeviceGroupResponseBodyData) *QueryFaceAllDeviceGroupResponseBody {
	s.Data = v
	return s
}

func (s *QueryFaceAllDeviceGroupResponseBody) SetErrorMessage(v string) *QueryFaceAllDeviceGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFaceAllDeviceGroupResponseBody) SetRequestId(v string) *QueryFaceAllDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFaceAllDeviceGroupResponseBody) SetSuccess(v bool) *QueryFaceAllDeviceGroupResponseBody {
	s.Success = &v
	return s
}

type QueryFaceAllDeviceGroupResponseBodyData struct {
	DeviceGroupList []*QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList `json:"DeviceGroupList,omitempty" xml:"DeviceGroupList,omitempty" type:"Repeated"`
	PageNo          *int32                                                    `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize        *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total           *int32                                                    `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryFaceAllDeviceGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllDeviceGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceAllDeviceGroupResponseBodyData) SetDeviceGroupList(v []*QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList) *QueryFaceAllDeviceGroupResponseBodyData {
	s.DeviceGroupList = v
	return s
}

func (s *QueryFaceAllDeviceGroupResponseBodyData) SetPageNo(v int32) *QueryFaceAllDeviceGroupResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *QueryFaceAllDeviceGroupResponseBodyData) SetPageSize(v int32) *QueryFaceAllDeviceGroupResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryFaceAllDeviceGroupResponseBodyData) SetTotal(v int32) *QueryFaceAllDeviceGroupResponseBodyData {
	s.Total = &v
	return s
}

type QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList struct {
	DeviceGroupId   *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	DeviceGroupName *string `json:"DeviceGroupName,omitempty" xml:"DeviceGroupName,omitempty"`
	ModifiedTime    *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
}

func (s QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList) GoString() string {
	return s.String()
}

func (s *QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList) SetDeviceGroupId(v string) *QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList {
	s.DeviceGroupId = &v
	return s
}

func (s *QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList) SetDeviceGroupName(v string) *QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList {
	s.DeviceGroupName = &v
	return s
}

func (s *QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList) SetModifiedTime(v string) *QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList {
	s.ModifiedTime = &v
	return s
}

type QueryFaceAllDeviceGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceAllDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceAllDeviceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllDeviceGroupResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceAllDeviceGroupResponse) SetHeaders(v map[string]*string) *QueryFaceAllDeviceGroupResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceAllDeviceGroupResponse) SetStatusCode(v int32) *QueryFaceAllDeviceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceAllDeviceGroupResponse) SetBody(v *QueryFaceAllDeviceGroupResponseBody) *QueryFaceAllDeviceGroupResponse {
	s.Body = v
	return s
}

type QueryFaceAllUserGroupRequest struct {
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	PageNo      *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryFaceAllUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllUserGroupRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserGroupRequest) SetIsolationId(v string) *QueryFaceAllUserGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceAllUserGroupRequest) SetPageNo(v int32) *QueryFaceAllUserGroupRequest {
	s.PageNo = &v
	return s
}

func (s *QueryFaceAllUserGroupRequest) SetPageSize(v int32) *QueryFaceAllUserGroupRequest {
	s.PageSize = &v
	return s
}

type QueryFaceAllUserGroupResponseBody struct {
	Code         *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryFaceAllUserGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceAllUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserGroupResponseBody) SetCode(v string) *QueryFaceAllUserGroupResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceAllUserGroupResponseBody) SetData(v *QueryFaceAllUserGroupResponseBodyData) *QueryFaceAllUserGroupResponseBody {
	s.Data = v
	return s
}

func (s *QueryFaceAllUserGroupResponseBody) SetErrorMessage(v string) *QueryFaceAllUserGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFaceAllUserGroupResponseBody) SetRequestId(v string) *QueryFaceAllUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFaceAllUserGroupResponseBody) SetSuccess(v bool) *QueryFaceAllUserGroupResponseBody {
	s.Success = &v
	return s
}

type QueryFaceAllUserGroupResponseBodyData struct {
	PageNo        *int32                                                `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize      *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total         *int32                                                `json:"Total,omitempty" xml:"Total,omitempty"`
	UserGroupList []*QueryFaceAllUserGroupResponseBodyDataUserGroupList `json:"UserGroupList,omitempty" xml:"UserGroupList,omitempty" type:"Repeated"`
}

func (s QueryFaceAllUserGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllUserGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserGroupResponseBodyData) SetPageNo(v int32) *QueryFaceAllUserGroupResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *QueryFaceAllUserGroupResponseBodyData) SetPageSize(v int32) *QueryFaceAllUserGroupResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryFaceAllUserGroupResponseBodyData) SetTotal(v int32) *QueryFaceAllUserGroupResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryFaceAllUserGroupResponseBodyData) SetUserGroupList(v []*QueryFaceAllUserGroupResponseBodyDataUserGroupList) *QueryFaceAllUserGroupResponseBodyData {
	s.UserGroupList = v
	return s
}

type QueryFaceAllUserGroupResponseBodyDataUserGroupList struct {
	ModifiedTime  *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	UserGroupId   *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s QueryFaceAllUserGroupResponseBodyDataUserGroupList) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllUserGroupResponseBodyDataUserGroupList) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserGroupResponseBodyDataUserGroupList) SetModifiedTime(v string) *QueryFaceAllUserGroupResponseBodyDataUserGroupList {
	s.ModifiedTime = &v
	return s
}

func (s *QueryFaceAllUserGroupResponseBodyDataUserGroupList) SetUserGroupId(v string) *QueryFaceAllUserGroupResponseBodyDataUserGroupList {
	s.UserGroupId = &v
	return s
}

func (s *QueryFaceAllUserGroupResponseBodyDataUserGroupList) SetUserGroupName(v string) *QueryFaceAllUserGroupResponseBodyDataUserGroupList {
	s.UserGroupName = &v
	return s
}

type QueryFaceAllUserGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceAllUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceAllUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllUserGroupResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserGroupResponse) SetHeaders(v map[string]*string) *QueryFaceAllUserGroupResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceAllUserGroupResponse) SetStatusCode(v int32) *QueryFaceAllUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceAllUserGroupResponse) SetBody(v *QueryFaceAllUserGroupResponseBody) *QueryFaceAllUserGroupResponse {
	s.Body = v
	return s
}

type QueryFaceAllUserGroupAndDeviceGroupRelationRequest struct {
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	PageNo      *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryFaceAllUserGroupAndDeviceGroupRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllUserGroupAndDeviceGroupRelationRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationRequest) SetIsolationId(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationRequest) SetPageNo(v int32) *QueryFaceAllUserGroupAndDeviceGroupRelationRequest {
	s.PageNo = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationRequest) SetPageSize(v int32) *QueryFaceAllUserGroupAndDeviceGroupRelationRequest {
	s.PageSize = &v
	return s
}

type QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody struct {
	Code         *string                                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) SetCode(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) SetData(v *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody {
	s.Data = v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) SetErrorMessage(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) SetRequestId(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) SetSuccess(v bool) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody {
	s.Success = &v
	return s
}

type QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData struct {
	List     []*QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Page     *int32                                                             `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *int32                                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int32                                                             `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData) SetList(v []*QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData {
	s.List = v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData) SetPage(v int32) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData {
	s.Page = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData) SetPageSize(v int32) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData) SetTotal(v int32) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData {
	s.Total = &v
	return s
}

type QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList struct {
	ControlId     *string `json:"ControlId,omitempty" xml:"ControlId,omitempty"`
	ControlType   *string `json:"ControlType,omitempty" xml:"ControlType,omitempty"`
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	ModifiedTime  *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	UserGroupId   *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) SetControlId(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList {
	s.ControlId = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) SetControlType(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList {
	s.ControlType = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) SetDeviceGroupId(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList {
	s.DeviceGroupId = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) SetModifiedTime(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList {
	s.ModifiedTime = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) SetUserGroupId(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList {
	s.UserGroupId = &v
	return s
}

type QueryFaceAllUserGroupAndDeviceGroupRelationResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceAllUserGroupAndDeviceGroupRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllUserGroupAndDeviceGroupRelationResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponse) SetHeaders(v map[string]*string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponse) SetStatusCode(v int32) *QueryFaceAllUserGroupAndDeviceGroupRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponse) SetBody(v *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) *QueryFaceAllUserGroupAndDeviceGroupRelationResponse {
	s.Body = v
	return s
}

type QueryFaceAllUserIdsByGroupIdRequest struct {
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	PageNo      *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s QueryFaceAllUserIdsByGroupIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllUserIdsByGroupIdRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserIdsByGroupIdRequest) SetIsolationId(v string) *QueryFaceAllUserIdsByGroupIdRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdRequest) SetPageNo(v int32) *QueryFaceAllUserIdsByGroupIdRequest {
	s.PageNo = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdRequest) SetPageSize(v int32) *QueryFaceAllUserIdsByGroupIdRequest {
	s.PageSize = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdRequest) SetUserGroupId(v string) *QueryFaceAllUserIdsByGroupIdRequest {
	s.UserGroupId = &v
	return s
}

type QueryFaceAllUserIdsByGroupIdResponseBody struct {
	Code         *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryFaceAllUserIdsByGroupIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceAllUserIdsByGroupIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllUserIdsByGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBody) SetCode(v string) *QueryFaceAllUserIdsByGroupIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBody) SetData(v *QueryFaceAllUserIdsByGroupIdResponseBodyData) *QueryFaceAllUserIdsByGroupIdResponseBody {
	s.Data = v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBody) SetErrorMessage(v string) *QueryFaceAllUserIdsByGroupIdResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBody) SetRequestId(v string) *QueryFaceAllUserIdsByGroupIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBody) SetSuccess(v bool) *QueryFaceAllUserIdsByGroupIdResponseBody {
	s.Success = &v
	return s
}

type QueryFaceAllUserIdsByGroupIdResponseBodyData struct {
	List     []*QueryFaceAllUserIdsByGroupIdResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Page     *int32                                              `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int32                                              `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryFaceAllUserIdsByGroupIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllUserIdsByGroupIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyData) SetList(v []*QueryFaceAllUserIdsByGroupIdResponseBodyDataList) *QueryFaceAllUserIdsByGroupIdResponseBodyData {
	s.List = v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyData) SetPage(v int32) *QueryFaceAllUserIdsByGroupIdResponseBodyData {
	s.Page = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyData) SetPageSize(v int32) *QueryFaceAllUserIdsByGroupIdResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyData) SetTotal(v int32) *QueryFaceAllUserIdsByGroupIdResponseBodyData {
	s.Total = &v
	return s
}

type QueryFaceAllUserIdsByGroupIdResponseBodyDataList struct {
	CustomUserId *string `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Params       *string `json:"Params,omitempty" xml:"Params,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryFaceAllUserIdsByGroupIdResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllUserIdsByGroupIdResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyDataList) SetCustomUserId(v string) *QueryFaceAllUserIdsByGroupIdResponseBodyDataList {
	s.CustomUserId = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyDataList) SetName(v string) *QueryFaceAllUserIdsByGroupIdResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyDataList) SetParams(v string) *QueryFaceAllUserIdsByGroupIdResponseBodyDataList {
	s.Params = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyDataList) SetUserId(v string) *QueryFaceAllUserIdsByGroupIdResponseBodyDataList {
	s.UserId = &v
	return s
}

type QueryFaceAllUserIdsByGroupIdResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceAllUserIdsByGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceAllUserIdsByGroupIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllUserIdsByGroupIdResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserIdsByGroupIdResponse) SetHeaders(v map[string]*string) *QueryFaceAllUserIdsByGroupIdResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponse) SetStatusCode(v int32) *QueryFaceAllUserIdsByGroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponse) SetBody(v *QueryFaceAllUserIdsByGroupIdResponseBody) *QueryFaceAllUserIdsByGroupIdResponse {
	s.Body = v
	return s
}

type QueryFaceCustomUserIdByUserIdRequest struct {
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryFaceCustomUserIdByUserIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceCustomUserIdByUserIdRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceCustomUserIdByUserIdRequest) SetIsolationId(v string) *QueryFaceCustomUserIdByUserIdRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceCustomUserIdByUserIdRequest) SetUserId(v string) *QueryFaceCustomUserIdByUserIdRequest {
	s.UserId = &v
	return s
}

type QueryFaceCustomUserIdByUserIdResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceCustomUserIdByUserIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceCustomUserIdByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceCustomUserIdByUserIdResponseBody) SetCode(v string) *QueryFaceCustomUserIdByUserIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceCustomUserIdByUserIdResponseBody) SetData(v string) *QueryFaceCustomUserIdByUserIdResponseBody {
	s.Data = &v
	return s
}

func (s *QueryFaceCustomUserIdByUserIdResponseBody) SetErrorMessage(v string) *QueryFaceCustomUserIdByUserIdResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFaceCustomUserIdByUserIdResponseBody) SetRequestId(v string) *QueryFaceCustomUserIdByUserIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFaceCustomUserIdByUserIdResponseBody) SetSuccess(v bool) *QueryFaceCustomUserIdByUserIdResponseBody {
	s.Success = &v
	return s
}

type QueryFaceCustomUserIdByUserIdResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceCustomUserIdByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceCustomUserIdByUserIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceCustomUserIdByUserIdResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceCustomUserIdByUserIdResponse) SetHeaders(v map[string]*string) *QueryFaceCustomUserIdByUserIdResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceCustomUserIdByUserIdResponse) SetStatusCode(v int32) *QueryFaceCustomUserIdByUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceCustomUserIdByUserIdResponse) SetBody(v *QueryFaceCustomUserIdByUserIdResponseBody) *QueryFaceCustomUserIdByUserIdResponse {
	s.Body = v
	return s
}

type QueryFaceDeviceGroupsByDeviceRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	IsolationId   *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	PageNo        *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryFaceDeviceGroupsByDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceDeviceGroupsByDeviceRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceDeviceGroupsByDeviceRequest) SetDeviceName(v string) *QueryFaceDeviceGroupsByDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceRequest) SetIotInstanceId(v string) *QueryFaceDeviceGroupsByDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceRequest) SetIsolationId(v string) *QueryFaceDeviceGroupsByDeviceRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceRequest) SetPageNo(v int32) *QueryFaceDeviceGroupsByDeviceRequest {
	s.PageNo = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceRequest) SetPageSize(v int32) *QueryFaceDeviceGroupsByDeviceRequest {
	s.PageSize = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceRequest) SetProductKey(v string) *QueryFaceDeviceGroupsByDeviceRequest {
	s.ProductKey = &v
	return s
}

type QueryFaceDeviceGroupsByDeviceResponseBody struct {
	Code         *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryFaceDeviceGroupsByDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceDeviceGroupsByDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceDeviceGroupsByDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBody) SetCode(v string) *QueryFaceDeviceGroupsByDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBody) SetData(v *QueryFaceDeviceGroupsByDeviceResponseBodyData) *QueryFaceDeviceGroupsByDeviceResponseBody {
	s.Data = v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBody) SetErrorMessage(v string) *QueryFaceDeviceGroupsByDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBody) SetRequestId(v string) *QueryFaceDeviceGroupsByDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBody) SetSuccess(v bool) *QueryFaceDeviceGroupsByDeviceResponseBody {
	s.Success = &v
	return s
}

type QueryFaceDeviceGroupsByDeviceResponseBodyData struct {
	DeviceGroupList []*QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList `json:"DeviceGroupList,omitempty" xml:"DeviceGroupList,omitempty" type:"Repeated"`
	PageNo          *int32                                                          `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize        *int32                                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total           *int32                                                          `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryFaceDeviceGroupsByDeviceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceDeviceGroupsByDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBodyData) SetDeviceGroupList(v []*QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList) *QueryFaceDeviceGroupsByDeviceResponseBodyData {
	s.DeviceGroupList = v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBodyData) SetPageNo(v int32) *QueryFaceDeviceGroupsByDeviceResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBodyData) SetPageSize(v int32) *QueryFaceDeviceGroupsByDeviceResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBodyData) SetTotal(v int32) *QueryFaceDeviceGroupsByDeviceResponseBodyData {
	s.Total = &v
	return s
}

type QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList struct {
	DeviceGroupId   *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	DeviceGroupName *string `json:"DeviceGroupName,omitempty" xml:"DeviceGroupName,omitempty"`
	ModifiedTime    *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
}

func (s QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList) GoString() string {
	return s.String()
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList) SetDeviceGroupId(v string) *QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList {
	s.DeviceGroupId = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList) SetDeviceGroupName(v string) *QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList {
	s.DeviceGroupName = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList) SetModifiedTime(v string) *QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList {
	s.ModifiedTime = &v
	return s
}

type QueryFaceDeviceGroupsByDeviceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceDeviceGroupsByDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceDeviceGroupsByDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceDeviceGroupsByDeviceResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceDeviceGroupsByDeviceResponse) SetHeaders(v map[string]*string) *QueryFaceDeviceGroupsByDeviceResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceResponse) SetStatusCode(v int32) *QueryFaceDeviceGroupsByDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceResponse) SetBody(v *QueryFaceDeviceGroupsByDeviceResponseBody) *QueryFaceDeviceGroupsByDeviceResponse {
	s.Body = v
	return s
}

type QueryFaceUserRequest struct {
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryFaceUserRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceUserRequest) SetIsolationId(v string) *QueryFaceUserRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceUserRequest) SetUserId(v string) *QueryFaceUserRequest {
	s.UserId = &v
	return s
}

type QueryFaceUserResponseBody struct {
	Code         *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryFaceUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceUserResponseBody) SetCode(v string) *QueryFaceUserResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceUserResponseBody) SetData(v *QueryFaceUserResponseBodyData) *QueryFaceUserResponseBody {
	s.Data = v
	return s
}

func (s *QueryFaceUserResponseBody) SetErrorMessage(v string) *QueryFaceUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFaceUserResponseBody) SetRequestId(v string) *QueryFaceUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFaceUserResponseBody) SetSuccess(v bool) *QueryFaceUserResponseBody {
	s.Success = &v
	return s
}

type QueryFaceUserResponseBodyData struct {
	CustomUserId *string                                     `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	FacePicList  []*QueryFaceUserResponseBodyDataFacePicList `json:"FacePicList,omitempty" xml:"FacePicList,omitempty" type:"Repeated"`
	Name         *string                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	Params       *string                                     `json:"Params,omitempty" xml:"Params,omitempty"`
	UserId       *string                                     `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryFaceUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceUserResponseBodyData) SetCustomUserId(v string) *QueryFaceUserResponseBodyData {
	s.CustomUserId = &v
	return s
}

func (s *QueryFaceUserResponseBodyData) SetFacePicList(v []*QueryFaceUserResponseBodyDataFacePicList) *QueryFaceUserResponseBodyData {
	s.FacePicList = v
	return s
}

func (s *QueryFaceUserResponseBodyData) SetName(v string) *QueryFaceUserResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryFaceUserResponseBodyData) SetParams(v string) *QueryFaceUserResponseBodyData {
	s.Params = &v
	return s
}

func (s *QueryFaceUserResponseBodyData) SetUserId(v string) *QueryFaceUserResponseBodyData {
	s.UserId = &v
	return s
}

type QueryFaceUserResponseBodyDataFacePicList struct {
	FaceMd5        *string                                                   `json:"FaceMd5,omitempty" xml:"FaceMd5,omitempty"`
	FaceUrl        *string                                                   `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	FeatureDTOList []*QueryFaceUserResponseBodyDataFacePicListFeatureDTOList `json:"FeatureDTOList,omitempty" xml:"FeatureDTOList,omitempty" type:"Repeated"`
}

func (s QueryFaceUserResponseBodyDataFacePicList) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserResponseBodyDataFacePicList) GoString() string {
	return s.String()
}

func (s *QueryFaceUserResponseBodyDataFacePicList) SetFaceMd5(v string) *QueryFaceUserResponseBodyDataFacePicList {
	s.FaceMd5 = &v
	return s
}

func (s *QueryFaceUserResponseBodyDataFacePicList) SetFaceUrl(v string) *QueryFaceUserResponseBodyDataFacePicList {
	s.FaceUrl = &v
	return s
}

func (s *QueryFaceUserResponseBodyDataFacePicList) SetFeatureDTOList(v []*QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) *QueryFaceUserResponseBodyDataFacePicList {
	s.FeatureDTOList = v
	return s
}

type QueryFaceUserResponseBodyDataFacePicListFeatureDTOList struct {
	AlgorithmName     *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	AlgorithmProvider *string `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	AlgorithmVersion  *string `json:"AlgorithmVersion,omitempty" xml:"AlgorithmVersion,omitempty"`
	ErrorCode         *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage      *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	FaceMd5           *string `json:"FaceMd5,omitempty" xml:"FaceMd5,omitempty"`
}

func (s QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) GoString() string {
	return s.String()
}

func (s *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) SetAlgorithmName(v string) *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList {
	s.AlgorithmName = &v
	return s
}

func (s *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) SetAlgorithmProvider(v string) *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList {
	s.AlgorithmProvider = &v
	return s
}

func (s *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) SetAlgorithmVersion(v string) *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList {
	s.AlgorithmVersion = &v
	return s
}

func (s *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) SetErrorCode(v string) *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList {
	s.ErrorCode = &v
	return s
}

func (s *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) SetErrorMessage(v string) *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) SetFaceMd5(v string) *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList {
	s.FaceMd5 = &v
	return s
}

type QueryFaceUserResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceUserResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceUserResponse) SetHeaders(v map[string]*string) *QueryFaceUserResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceUserResponse) SetStatusCode(v int32) *QueryFaceUserResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceUserResponse) SetBody(v *QueryFaceUserResponseBody) *QueryFaceUserResponse {
	s.Body = v
	return s
}

type QueryFaceUserBatchRequest struct {
	IsolationId *string   `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	UserIdList  []*string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty" type:"Repeated"`
}

func (s QueryFaceUserBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserBatchRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceUserBatchRequest) SetIsolationId(v string) *QueryFaceUserBatchRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceUserBatchRequest) SetUserIdList(v []*string) *QueryFaceUserBatchRequest {
	s.UserIdList = v
	return s
}

type QueryFaceUserBatchResponseBody struct {
	Code         *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         []*QueryFaceUserBatchResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMessage *string                               `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceUserBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserBatchResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceUserBatchResponseBody) SetCode(v string) *QueryFaceUserBatchResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceUserBatchResponseBody) SetData(v []*QueryFaceUserBatchResponseBodyData) *QueryFaceUserBatchResponseBody {
	s.Data = v
	return s
}

func (s *QueryFaceUserBatchResponseBody) SetErrorMessage(v string) *QueryFaceUserBatchResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFaceUserBatchResponseBody) SetRequestId(v string) *QueryFaceUserBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFaceUserBatchResponseBody) SetSuccess(v bool) *QueryFaceUserBatchResponseBody {
	s.Success = &v
	return s
}

type QueryFaceUserBatchResponseBodyData struct {
	CreateTime   *int64                                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomUserId *string                                          `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	FacePicList  []*QueryFaceUserBatchResponseBodyDataFacePicList `json:"FacePicList,omitempty" xml:"FacePicList,omitempty" type:"Repeated"`
	ModifyTime   *int64                                           `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Name         *string                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	Params       *string                                          `json:"Params,omitempty" xml:"Params,omitempty"`
	UserId       *string                                          `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryFaceUserBatchResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserBatchResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceUserBatchResponseBodyData) SetCreateTime(v int64) *QueryFaceUserBatchResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *QueryFaceUserBatchResponseBodyData) SetCustomUserId(v string) *QueryFaceUserBatchResponseBodyData {
	s.CustomUserId = &v
	return s
}

func (s *QueryFaceUserBatchResponseBodyData) SetFacePicList(v []*QueryFaceUserBatchResponseBodyDataFacePicList) *QueryFaceUserBatchResponseBodyData {
	s.FacePicList = v
	return s
}

func (s *QueryFaceUserBatchResponseBodyData) SetModifyTime(v int64) *QueryFaceUserBatchResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *QueryFaceUserBatchResponseBodyData) SetName(v string) *QueryFaceUserBatchResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryFaceUserBatchResponseBodyData) SetParams(v string) *QueryFaceUserBatchResponseBodyData {
	s.Params = &v
	return s
}

func (s *QueryFaceUserBatchResponseBodyData) SetUserId(v string) *QueryFaceUserBatchResponseBodyData {
	s.UserId = &v
	return s
}

type QueryFaceUserBatchResponseBodyDataFacePicList struct {
	FaceMd5        *string                                                        `json:"FaceMd5,omitempty" xml:"FaceMd5,omitempty"`
	FaceUrl        *string                                                        `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	FeatureDTOList []*QueryFaceUserBatchResponseBodyDataFacePicListFeatureDTOList `json:"FeatureDTOList,omitempty" xml:"FeatureDTOList,omitempty" type:"Repeated"`
}

func (s QueryFaceUserBatchResponseBodyDataFacePicList) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserBatchResponseBodyDataFacePicList) GoString() string {
	return s.String()
}

func (s *QueryFaceUserBatchResponseBodyDataFacePicList) SetFaceMd5(v string) *QueryFaceUserBatchResponseBodyDataFacePicList {
	s.FaceMd5 = &v
	return s
}

func (s *QueryFaceUserBatchResponseBodyDataFacePicList) SetFaceUrl(v string) *QueryFaceUserBatchResponseBodyDataFacePicList {
	s.FaceUrl = &v
	return s
}

func (s *QueryFaceUserBatchResponseBodyDataFacePicList) SetFeatureDTOList(v []*QueryFaceUserBatchResponseBodyDataFacePicListFeatureDTOList) *QueryFaceUserBatchResponseBodyDataFacePicList {
	s.FeatureDTOList = v
	return s
}

type QueryFaceUserBatchResponseBodyDataFacePicListFeatureDTOList struct {
	AlgorithmName     *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	AlgorithmProvider *string `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	AlgorithmVersion  *string `json:"AlgorithmVersion,omitempty" xml:"AlgorithmVersion,omitempty"`
	ErrorCode         *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage      *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	FaceMd5           *string `json:"FaceMd5,omitempty" xml:"FaceMd5,omitempty"`
}

func (s QueryFaceUserBatchResponseBodyDataFacePicListFeatureDTOList) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserBatchResponseBodyDataFacePicListFeatureDTOList) GoString() string {
	return s.String()
}

func (s *QueryFaceUserBatchResponseBodyDataFacePicListFeatureDTOList) SetAlgorithmName(v string) *QueryFaceUserBatchResponseBodyDataFacePicListFeatureDTOList {
	s.AlgorithmName = &v
	return s
}

func (s *QueryFaceUserBatchResponseBodyDataFacePicListFeatureDTOList) SetAlgorithmProvider(v string) *QueryFaceUserBatchResponseBodyDataFacePicListFeatureDTOList {
	s.AlgorithmProvider = &v
	return s
}

func (s *QueryFaceUserBatchResponseBodyDataFacePicListFeatureDTOList) SetAlgorithmVersion(v string) *QueryFaceUserBatchResponseBodyDataFacePicListFeatureDTOList {
	s.AlgorithmVersion = &v
	return s
}

func (s *QueryFaceUserBatchResponseBodyDataFacePicListFeatureDTOList) SetErrorCode(v string) *QueryFaceUserBatchResponseBodyDataFacePicListFeatureDTOList {
	s.ErrorCode = &v
	return s
}

func (s *QueryFaceUserBatchResponseBodyDataFacePicListFeatureDTOList) SetErrorMessage(v string) *QueryFaceUserBatchResponseBodyDataFacePicListFeatureDTOList {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFaceUserBatchResponseBodyDataFacePicListFeatureDTOList) SetFaceMd5(v string) *QueryFaceUserBatchResponseBodyDataFacePicListFeatureDTOList {
	s.FaceMd5 = &v
	return s
}

type QueryFaceUserBatchResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceUserBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceUserBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserBatchResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceUserBatchResponse) SetHeaders(v map[string]*string) *QueryFaceUserBatchResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceUserBatchResponse) SetStatusCode(v int32) *QueryFaceUserBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceUserBatchResponse) SetBody(v *QueryFaceUserBatchResponseBody) *QueryFaceUserBatchResponse {
	s.Body = v
	return s
}

type QueryFaceUserByNameRequest struct {
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNo      *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Params      *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s QueryFaceUserByNameRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserByNameRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceUserByNameRequest) SetIsolationId(v string) *QueryFaceUserByNameRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceUserByNameRequest) SetName(v string) *QueryFaceUserByNameRequest {
	s.Name = &v
	return s
}

func (s *QueryFaceUserByNameRequest) SetPageNo(v int32) *QueryFaceUserByNameRequest {
	s.PageNo = &v
	return s
}

func (s *QueryFaceUserByNameRequest) SetPageSize(v int32) *QueryFaceUserByNameRequest {
	s.PageSize = &v
	return s
}

func (s *QueryFaceUserByNameRequest) SetParams(v string) *QueryFaceUserByNameRequest {
	s.Params = &v
	return s
}

type QueryFaceUserByNameResponseBody struct {
	Code         *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryFaceUserByNameResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceUserByNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserByNameResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceUserByNameResponseBody) SetCode(v string) *QueryFaceUserByNameResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceUserByNameResponseBody) SetData(v *QueryFaceUserByNameResponseBodyData) *QueryFaceUserByNameResponseBody {
	s.Data = v
	return s
}

func (s *QueryFaceUserByNameResponseBody) SetErrorMessage(v string) *QueryFaceUserByNameResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFaceUserByNameResponseBody) SetRequestId(v string) *QueryFaceUserByNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFaceUserByNameResponseBody) SetSuccess(v bool) *QueryFaceUserByNameResponseBody {
	s.Success = &v
	return s
}

type QueryFaceUserByNameResponseBodyData struct {
	List     []*QueryFaceUserByNameResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Page     *int32                                     `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int32                                     `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryFaceUserByNameResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserByNameResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceUserByNameResponseBodyData) SetList(v []*QueryFaceUserByNameResponseBodyDataList) *QueryFaceUserByNameResponseBodyData {
	s.List = v
	return s
}

func (s *QueryFaceUserByNameResponseBodyData) SetPage(v int32) *QueryFaceUserByNameResponseBodyData {
	s.Page = &v
	return s
}

func (s *QueryFaceUserByNameResponseBodyData) SetPageSize(v int32) *QueryFaceUserByNameResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryFaceUserByNameResponseBodyData) SetTotal(v int32) *QueryFaceUserByNameResponseBodyData {
	s.Total = &v
	return s
}

type QueryFaceUserByNameResponseBodyDataList struct {
	CreateTime   *int64                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomUserId *string                                               `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	FacePicList  []*QueryFaceUserByNameResponseBodyDataListFacePicList `json:"FacePicList,omitempty" xml:"FacePicList,omitempty" type:"Repeated"`
	ModifyTime   *int64                                                `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Name         *string                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	Params       *string                                               `json:"Params,omitempty" xml:"Params,omitempty"`
	UserId       *string                                               `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryFaceUserByNameResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserByNameResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryFaceUserByNameResponseBodyDataList) SetCreateTime(v int64) *QueryFaceUserByNameResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *QueryFaceUserByNameResponseBodyDataList) SetCustomUserId(v string) *QueryFaceUserByNameResponseBodyDataList {
	s.CustomUserId = &v
	return s
}

func (s *QueryFaceUserByNameResponseBodyDataList) SetFacePicList(v []*QueryFaceUserByNameResponseBodyDataListFacePicList) *QueryFaceUserByNameResponseBodyDataList {
	s.FacePicList = v
	return s
}

func (s *QueryFaceUserByNameResponseBodyDataList) SetModifyTime(v int64) *QueryFaceUserByNameResponseBodyDataList {
	s.ModifyTime = &v
	return s
}

func (s *QueryFaceUserByNameResponseBodyDataList) SetName(v string) *QueryFaceUserByNameResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *QueryFaceUserByNameResponseBodyDataList) SetParams(v string) *QueryFaceUserByNameResponseBodyDataList {
	s.Params = &v
	return s
}

func (s *QueryFaceUserByNameResponseBodyDataList) SetUserId(v string) *QueryFaceUserByNameResponseBodyDataList {
	s.UserId = &v
	return s
}

type QueryFaceUserByNameResponseBodyDataListFacePicList struct {
	FaceMd5        *string                                                             `json:"FaceMd5,omitempty" xml:"FaceMd5,omitempty"`
	FaceUrl        *string                                                             `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	FeatureDTOList []*QueryFaceUserByNameResponseBodyDataListFacePicListFeatureDTOList `json:"FeatureDTOList,omitempty" xml:"FeatureDTOList,omitempty" type:"Repeated"`
}

func (s QueryFaceUserByNameResponseBodyDataListFacePicList) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserByNameResponseBodyDataListFacePicList) GoString() string {
	return s.String()
}

func (s *QueryFaceUserByNameResponseBodyDataListFacePicList) SetFaceMd5(v string) *QueryFaceUserByNameResponseBodyDataListFacePicList {
	s.FaceMd5 = &v
	return s
}

func (s *QueryFaceUserByNameResponseBodyDataListFacePicList) SetFaceUrl(v string) *QueryFaceUserByNameResponseBodyDataListFacePicList {
	s.FaceUrl = &v
	return s
}

func (s *QueryFaceUserByNameResponseBodyDataListFacePicList) SetFeatureDTOList(v []*QueryFaceUserByNameResponseBodyDataListFacePicListFeatureDTOList) *QueryFaceUserByNameResponseBodyDataListFacePicList {
	s.FeatureDTOList = v
	return s
}

type QueryFaceUserByNameResponseBodyDataListFacePicListFeatureDTOList struct {
	AlgorithmName     *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	AlgorithmProvider *string `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	AlgorithmVersion  *string `json:"AlgorithmVersion,omitempty" xml:"AlgorithmVersion,omitempty"`
	ErrorCode         *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage      *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	FaceMd5           *string `json:"FaceMd5,omitempty" xml:"FaceMd5,omitempty"`
}

func (s QueryFaceUserByNameResponseBodyDataListFacePicListFeatureDTOList) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserByNameResponseBodyDataListFacePicListFeatureDTOList) GoString() string {
	return s.String()
}

func (s *QueryFaceUserByNameResponseBodyDataListFacePicListFeatureDTOList) SetAlgorithmName(v string) *QueryFaceUserByNameResponseBodyDataListFacePicListFeatureDTOList {
	s.AlgorithmName = &v
	return s
}

func (s *QueryFaceUserByNameResponseBodyDataListFacePicListFeatureDTOList) SetAlgorithmProvider(v string) *QueryFaceUserByNameResponseBodyDataListFacePicListFeatureDTOList {
	s.AlgorithmProvider = &v
	return s
}

func (s *QueryFaceUserByNameResponseBodyDataListFacePicListFeatureDTOList) SetAlgorithmVersion(v string) *QueryFaceUserByNameResponseBodyDataListFacePicListFeatureDTOList {
	s.AlgorithmVersion = &v
	return s
}

func (s *QueryFaceUserByNameResponseBodyDataListFacePicListFeatureDTOList) SetErrorCode(v string) *QueryFaceUserByNameResponseBodyDataListFacePicListFeatureDTOList {
	s.ErrorCode = &v
	return s
}

func (s *QueryFaceUserByNameResponseBodyDataListFacePicListFeatureDTOList) SetErrorMessage(v string) *QueryFaceUserByNameResponseBodyDataListFacePicListFeatureDTOList {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFaceUserByNameResponseBodyDataListFacePicListFeatureDTOList) SetFaceMd5(v string) *QueryFaceUserByNameResponseBodyDataListFacePicListFeatureDTOList {
	s.FaceMd5 = &v
	return s
}

type QueryFaceUserByNameResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceUserByNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceUserByNameResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserByNameResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceUserByNameResponse) SetHeaders(v map[string]*string) *QueryFaceUserByNameResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceUserByNameResponse) SetStatusCode(v int32) *QueryFaceUserByNameResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceUserByNameResponse) SetBody(v *QueryFaceUserByNameResponseBody) *QueryFaceUserByNameResponse {
	s.Body = v
	return s
}

type QueryFaceUserGroupRequest struct {
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	PageNo      *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryFaceUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserGroupRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceUserGroupRequest) SetIsolationId(v string) *QueryFaceUserGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceUserGroupRequest) SetPageNo(v int32) *QueryFaceUserGroupRequest {
	s.PageNo = &v
	return s
}

func (s *QueryFaceUserGroupRequest) SetPageSize(v int32) *QueryFaceUserGroupRequest {
	s.PageSize = &v
	return s
}

func (s *QueryFaceUserGroupRequest) SetUserId(v string) *QueryFaceUserGroupRequest {
	s.UserId = &v
	return s
}

type QueryFaceUserGroupResponseBody struct {
	Code         *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryFaceUserGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceUserGroupResponseBody) SetCode(v string) *QueryFaceUserGroupResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceUserGroupResponseBody) SetData(v *QueryFaceUserGroupResponseBodyData) *QueryFaceUserGroupResponseBody {
	s.Data = v
	return s
}

func (s *QueryFaceUserGroupResponseBody) SetErrorMessage(v string) *QueryFaceUserGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFaceUserGroupResponseBody) SetRequestId(v string) *QueryFaceUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFaceUserGroupResponseBody) SetSuccess(v bool) *QueryFaceUserGroupResponseBody {
	s.Success = &v
	return s
}

type QueryFaceUserGroupResponseBodyData struct {
	PageNo        *int32                                             `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize      *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total         *int32                                             `json:"Total,omitempty" xml:"Total,omitempty"`
	UserGroupList []*QueryFaceUserGroupResponseBodyDataUserGroupList `json:"UserGroupList,omitempty" xml:"UserGroupList,omitempty" type:"Repeated"`
}

func (s QueryFaceUserGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceUserGroupResponseBodyData) SetPageNo(v int32) *QueryFaceUserGroupResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *QueryFaceUserGroupResponseBodyData) SetPageSize(v int32) *QueryFaceUserGroupResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryFaceUserGroupResponseBodyData) SetTotal(v int32) *QueryFaceUserGroupResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryFaceUserGroupResponseBodyData) SetUserGroupList(v []*QueryFaceUserGroupResponseBodyDataUserGroupList) *QueryFaceUserGroupResponseBodyData {
	s.UserGroupList = v
	return s
}

type QueryFaceUserGroupResponseBodyDataUserGroupList struct {
	ModifiedTime  *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	UserGroupId   *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s QueryFaceUserGroupResponseBodyDataUserGroupList) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserGroupResponseBodyDataUserGroupList) GoString() string {
	return s.String()
}

func (s *QueryFaceUserGroupResponseBodyDataUserGroupList) SetModifiedTime(v string) *QueryFaceUserGroupResponseBodyDataUserGroupList {
	s.ModifiedTime = &v
	return s
}

func (s *QueryFaceUserGroupResponseBodyDataUserGroupList) SetUserGroupId(v string) *QueryFaceUserGroupResponseBodyDataUserGroupList {
	s.UserGroupId = &v
	return s
}

func (s *QueryFaceUserGroupResponseBodyDataUserGroupList) SetUserGroupName(v string) *QueryFaceUserGroupResponseBodyDataUserGroupList {
	s.UserGroupName = &v
	return s
}

type QueryFaceUserGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserGroupResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceUserGroupResponse) SetHeaders(v map[string]*string) *QueryFaceUserGroupResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceUserGroupResponse) SetStatusCode(v int32) *QueryFaceUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceUserGroupResponse) SetBody(v *QueryFaceUserGroupResponseBody) *QueryFaceUserGroupResponse {
	s.Body = v
	return s
}

type QueryFaceUserGroupAndDeviceGroupRelationRequest struct {
	ControlId   *string `json:"ControlId,omitempty" xml:"ControlId,omitempty"`
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
}

func (s QueryFaceUserGroupAndDeviceGroupRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserGroupAndDeviceGroupRelationRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationRequest) SetControlId(v string) *QueryFaceUserGroupAndDeviceGroupRelationRequest {
	s.ControlId = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationRequest) SetIsolationId(v string) *QueryFaceUserGroupAndDeviceGroupRelationRequest {
	s.IsolationId = &v
	return s
}

type QueryFaceUserGroupAndDeviceGroupRelationResponseBody struct {
	Code         *string                                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceUserGroupAndDeviceGroupRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserGroupAndDeviceGroupRelationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBody) SetCode(v string) *QueryFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBody) SetData(v *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) *QueryFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.Data = v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBody) SetErrorMessage(v string) *QueryFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBody) SetRequestId(v string) *QueryFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBody) SetSuccess(v bool) *QueryFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.Success = &v
	return s
}

type QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData struct {
	ControlId     *string `json:"ControlId,omitempty" xml:"ControlId,omitempty"`
	ControlType   *string `json:"ControlType,omitempty" xml:"ControlType,omitempty"`
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	ModifiedTime  *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	UserGroupId   *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) SetControlId(v string) *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData {
	s.ControlId = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) SetControlType(v string) *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData {
	s.ControlType = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) SetDeviceGroupId(v string) *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData {
	s.DeviceGroupId = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) SetModifiedTime(v string) *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) SetUserGroupId(v string) *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData {
	s.UserGroupId = &v
	return s
}

type QueryFaceUserGroupAndDeviceGroupRelationResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceUserGroupAndDeviceGroupRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceUserGroupAndDeviceGroupRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserGroupAndDeviceGroupRelationResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponse) SetHeaders(v map[string]*string) *QueryFaceUserGroupAndDeviceGroupRelationResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponse) SetStatusCode(v int32) *QueryFaceUserGroupAndDeviceGroupRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponse) SetBody(v *QueryFaceUserGroupAndDeviceGroupRelationResponseBody) *QueryFaceUserGroupAndDeviceGroupRelationResponse {
	s.Body = v
	return s
}

type QueryFaceUserIdByCustomUserIdRequest struct {
	CustomUserId *string `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	IsolationId  *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
}

func (s QueryFaceUserIdByCustomUserIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserIdByCustomUserIdRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceUserIdByCustomUserIdRequest) SetCustomUserId(v string) *QueryFaceUserIdByCustomUserIdRequest {
	s.CustomUserId = &v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdRequest) SetIsolationId(v string) *QueryFaceUserIdByCustomUserIdRequest {
	s.IsolationId = &v
	return s
}

type QueryFaceUserIdByCustomUserIdResponseBody struct {
	Code         *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryFaceUserIdByCustomUserIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceUserIdByCustomUserIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserIdByCustomUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceUserIdByCustomUserIdResponseBody) SetCode(v string) *QueryFaceUserIdByCustomUserIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdResponseBody) SetData(v *QueryFaceUserIdByCustomUserIdResponseBodyData) *QueryFaceUserIdByCustomUserIdResponseBody {
	s.Data = v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdResponseBody) SetErrorMessage(v string) *QueryFaceUserIdByCustomUserIdResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdResponseBody) SetRequestId(v string) *QueryFaceUserIdByCustomUserIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdResponseBody) SetSuccess(v bool) *QueryFaceUserIdByCustomUserIdResponseBody {
	s.Success = &v
	return s
}

type QueryFaceUserIdByCustomUserIdResponseBodyData struct {
	CustomUserId *string `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Params       *string `json:"Params,omitempty" xml:"Params,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryFaceUserIdByCustomUserIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserIdByCustomUserIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceUserIdByCustomUserIdResponseBodyData) SetCustomUserId(v string) *QueryFaceUserIdByCustomUserIdResponseBodyData {
	s.CustomUserId = &v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdResponseBodyData) SetName(v string) *QueryFaceUserIdByCustomUserIdResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdResponseBodyData) SetParams(v string) *QueryFaceUserIdByCustomUserIdResponseBodyData {
	s.Params = &v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdResponseBodyData) SetUserId(v string) *QueryFaceUserIdByCustomUserIdResponseBodyData {
	s.UserId = &v
	return s
}

type QueryFaceUserIdByCustomUserIdResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceUserIdByCustomUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceUserIdByCustomUserIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserIdByCustomUserIdResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceUserIdByCustomUserIdResponse) SetHeaders(v map[string]*string) *QueryFaceUserIdByCustomUserIdResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdResponse) SetStatusCode(v int32) *QueryFaceUserIdByCustomUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdResponse) SetBody(v *QueryFaceUserIdByCustomUserIdResponseBody) *QueryFaceUserIdByCustomUserIdResponse {
	s.Body = v
	return s
}

type QueryGbSubDeviceListRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageStart     *int32  `json:"PageStart,omitempty" xml:"PageStart,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryGbSubDeviceListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGbSubDeviceListRequest) GoString() string {
	return s.String()
}

func (s *QueryGbSubDeviceListRequest) SetDeviceName(v string) *QueryGbSubDeviceListRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryGbSubDeviceListRequest) SetIotId(v string) *QueryGbSubDeviceListRequest {
	s.IotId = &v
	return s
}

func (s *QueryGbSubDeviceListRequest) SetIotInstanceId(v string) *QueryGbSubDeviceListRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryGbSubDeviceListRequest) SetPageSize(v int32) *QueryGbSubDeviceListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryGbSubDeviceListRequest) SetPageStart(v int32) *QueryGbSubDeviceListRequest {
	s.PageStart = &v
	return s
}

func (s *QueryGbSubDeviceListRequest) SetProductKey(v string) *QueryGbSubDeviceListRequest {
	s.ProductKey = &v
	return s
}

type QueryGbSubDeviceListResponseBody struct {
	Code         *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryGbSubDeviceListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                               `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryGbSubDeviceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGbSubDeviceListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGbSubDeviceListResponseBody) SetCode(v string) *QueryGbSubDeviceListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryGbSubDeviceListResponseBody) SetData(v *QueryGbSubDeviceListResponseBodyData) *QueryGbSubDeviceListResponseBody {
	s.Data = v
	return s
}

func (s *QueryGbSubDeviceListResponseBody) SetErrorMessage(v string) *QueryGbSubDeviceListResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryGbSubDeviceListResponseBody) SetRequestId(v string) *QueryGbSubDeviceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryGbSubDeviceListResponseBody) SetSuccess(v bool) *QueryGbSubDeviceListResponseBody {
	s.Success = &v
	return s
}

type QueryGbSubDeviceListResponseBodyData struct {
	GbSubDeviceList []*QueryGbSubDeviceListResponseBodyDataGbSubDeviceList `json:"GbSubDeviceList,omitempty" xml:"GbSubDeviceList,omitempty" type:"Repeated"`
	Total           *int32                                                 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryGbSubDeviceListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryGbSubDeviceListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryGbSubDeviceListResponseBodyData) SetGbSubDeviceList(v []*QueryGbSubDeviceListResponseBodyDataGbSubDeviceList) *QueryGbSubDeviceListResponseBodyData {
	s.GbSubDeviceList = v
	return s
}

func (s *QueryGbSubDeviceListResponseBodyData) SetTotal(v int32) *QueryGbSubDeviceListResponseBodyData {
	s.Total = &v
	return s
}

type QueryGbSubDeviceListResponseBodyDataGbSubDeviceList struct {
	DeviceEnable *int32  `json:"DeviceEnable,omitempty" xml:"DeviceEnable,omitempty"`
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DeviceName   *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId        *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	ProductKey   *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryGbSubDeviceListResponseBodyDataGbSubDeviceList) String() string {
	return tea.Prettify(s)
}

func (s QueryGbSubDeviceListResponseBodyDataGbSubDeviceList) GoString() string {
	return s.String()
}

func (s *QueryGbSubDeviceListResponseBodyDataGbSubDeviceList) SetDeviceEnable(v int32) *QueryGbSubDeviceListResponseBodyDataGbSubDeviceList {
	s.DeviceEnable = &v
	return s
}

func (s *QueryGbSubDeviceListResponseBodyDataGbSubDeviceList) SetDeviceId(v string) *QueryGbSubDeviceListResponseBodyDataGbSubDeviceList {
	s.DeviceId = &v
	return s
}

func (s *QueryGbSubDeviceListResponseBodyDataGbSubDeviceList) SetDeviceName(v string) *QueryGbSubDeviceListResponseBodyDataGbSubDeviceList {
	s.DeviceName = &v
	return s
}

func (s *QueryGbSubDeviceListResponseBodyDataGbSubDeviceList) SetIotId(v string) *QueryGbSubDeviceListResponseBodyDataGbSubDeviceList {
	s.IotId = &v
	return s
}

func (s *QueryGbSubDeviceListResponseBodyDataGbSubDeviceList) SetProductKey(v string) *QueryGbSubDeviceListResponseBodyDataGbSubDeviceList {
	s.ProductKey = &v
	return s
}

type QueryGbSubDeviceListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGbSubDeviceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGbSubDeviceListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGbSubDeviceListResponse) GoString() string {
	return s.String()
}

func (s *QueryGbSubDeviceListResponse) SetHeaders(v map[string]*string) *QueryGbSubDeviceListResponse {
	s.Headers = v
	return s
}

func (s *QueryGbSubDeviceListResponse) SetStatusCode(v int32) *QueryGbSubDeviceListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGbSubDeviceListResponse) SetBody(v *QueryGbSubDeviceListResponseBody) *QueryGbSubDeviceListResponse {
	s.Body = v
	return s
}

type QueryLiveStreamingRequest struct {
	CacheDuration    *int32  `json:"CacheDuration,omitempty" xml:"CacheDuration,omitempty"`
	DeviceName       *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	EnableStun       *bool   `json:"EnableStun,omitempty" xml:"EnableStun,omitempty"`
	EncryptType      *int32  `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	ForceIFrame      *bool   `json:"ForceIFrame,omitempty" xml:"ForceIFrame,omitempty"`
	IotId            *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId    *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	PlayUnLimited    *bool   `json:"PlayUnLimited,omitempty" xml:"PlayUnLimited,omitempty"`
	ProductKey       *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	Scheme           *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
	ShouldEncrypt    *bool   `json:"ShouldEncrypt,omitempty" xml:"ShouldEncrypt,omitempty"`
	StreamType       *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	UrlValidDuration *int32  `json:"UrlValidDuration,omitempty" xml:"UrlValidDuration,omitempty"`
}

func (s QueryLiveStreamingRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveStreamingRequest) GoString() string {
	return s.String()
}

func (s *QueryLiveStreamingRequest) SetCacheDuration(v int32) *QueryLiveStreamingRequest {
	s.CacheDuration = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetDeviceName(v string) *QueryLiveStreamingRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetEnableStun(v bool) *QueryLiveStreamingRequest {
	s.EnableStun = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetEncryptType(v int32) *QueryLiveStreamingRequest {
	s.EncryptType = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetForceIFrame(v bool) *QueryLiveStreamingRequest {
	s.ForceIFrame = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetIotId(v string) *QueryLiveStreamingRequest {
	s.IotId = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetIotInstanceId(v string) *QueryLiveStreamingRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetPlayUnLimited(v bool) *QueryLiveStreamingRequest {
	s.PlayUnLimited = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetProductKey(v string) *QueryLiveStreamingRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetScheme(v string) *QueryLiveStreamingRequest {
	s.Scheme = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetShouldEncrypt(v bool) *QueryLiveStreamingRequest {
	s.ShouldEncrypt = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetStreamType(v int32) *QueryLiveStreamingRequest {
	s.StreamType = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetUrlValidDuration(v int32) *QueryLiveStreamingRequest {
	s.UrlValidDuration = &v
	return s
}

type QueryLiveStreamingResponseBody struct {
	Code         *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryLiveStreamingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryLiveStreamingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveStreamingResponseBody) GoString() string {
	return s.String()
}

func (s *QueryLiveStreamingResponseBody) SetCode(v string) *QueryLiveStreamingResponseBody {
	s.Code = &v
	return s
}

func (s *QueryLiveStreamingResponseBody) SetData(v *QueryLiveStreamingResponseBodyData) *QueryLiveStreamingResponseBody {
	s.Data = v
	return s
}

func (s *QueryLiveStreamingResponseBody) SetErrorMessage(v string) *QueryLiveStreamingResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryLiveStreamingResponseBody) SetRequestId(v string) *QueryLiveStreamingResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryLiveStreamingResponseBody) SetSuccess(v bool) *QueryLiveStreamingResponseBody {
	s.Success = &v
	return s
}

type QueryLiveStreamingResponseBodyData struct {
	Path            *string `json:"Path,omitempty" xml:"Path,omitempty"`
	RelayDecryptKey *string `json:"RelayDecryptKey,omitempty" xml:"RelayDecryptKey,omitempty"`
	StunInfo        *string `json:"StunInfo,omitempty" xml:"StunInfo,omitempty"`
}

func (s QueryLiveStreamingResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveStreamingResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryLiveStreamingResponseBodyData) SetPath(v string) *QueryLiveStreamingResponseBodyData {
	s.Path = &v
	return s
}

func (s *QueryLiveStreamingResponseBodyData) SetRelayDecryptKey(v string) *QueryLiveStreamingResponseBodyData {
	s.RelayDecryptKey = &v
	return s
}

func (s *QueryLiveStreamingResponseBodyData) SetStunInfo(v string) *QueryLiveStreamingResponseBodyData {
	s.StunInfo = &v
	return s
}

type QueryLiveStreamingResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryLiveStreamingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryLiveStreamingResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveStreamingResponse) GoString() string {
	return s.String()
}

func (s *QueryLiveStreamingResponse) SetHeaders(v map[string]*string) *QueryLiveStreamingResponse {
	s.Headers = v
	return s
}

func (s *QueryLiveStreamingResponse) SetStatusCode(v int32) *QueryLiveStreamingResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryLiveStreamingResponse) SetBody(v *QueryLiveStreamingResponseBody) *QueryLiveStreamingResponse {
	s.Body = v
	return s
}

type QueryLocalFileUploadJobRequest struct {
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	JobId         *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s QueryLocalFileUploadJobRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryLocalFileUploadJobRequest) GoString() string {
	return s.String()
}

func (s *QueryLocalFileUploadJobRequest) SetIotInstanceId(v string) *QueryLocalFileUploadJobRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryLocalFileUploadJobRequest) SetJobId(v string) *QueryLocalFileUploadJobRequest {
	s.JobId = &v
	return s
}

type QueryLocalFileUploadJobResponseBody struct {
	Code         *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryLocalFileUploadJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryLocalFileUploadJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryLocalFileUploadJobResponseBody) GoString() string {
	return s.String()
}

func (s *QueryLocalFileUploadJobResponseBody) SetCode(v string) *QueryLocalFileUploadJobResponseBody {
	s.Code = &v
	return s
}

func (s *QueryLocalFileUploadJobResponseBody) SetData(v *QueryLocalFileUploadJobResponseBodyData) *QueryLocalFileUploadJobResponseBody {
	s.Data = v
	return s
}

func (s *QueryLocalFileUploadJobResponseBody) SetErrorMessage(v string) *QueryLocalFileUploadJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryLocalFileUploadJobResponseBody) SetRequestId(v string) *QueryLocalFileUploadJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryLocalFileUploadJobResponseBody) SetSuccess(v bool) *QueryLocalFileUploadJobResponseBody {
	s.Success = &v
	return s
}

type QueryLocalFileUploadJobResponseBodyData struct {
	ResultList []*QueryLocalFileUploadJobResponseBodyDataResultList `json:"ResultList,omitempty" xml:"ResultList,omitempty" type:"Repeated"`
	Status     *int32                                               `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryLocalFileUploadJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryLocalFileUploadJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryLocalFileUploadJobResponseBodyData) SetResultList(v []*QueryLocalFileUploadJobResponseBodyDataResultList) *QueryLocalFileUploadJobResponseBodyData {
	s.ResultList = v
	return s
}

func (s *QueryLocalFileUploadJobResponseBodyData) SetStatus(v int32) *QueryLocalFileUploadJobResponseBodyData {
	s.Status = &v
	return s
}

type QueryLocalFileUploadJobResponseBodyDataResultList struct {
	Code          *int32                                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	DeviceName    *string                                                      `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	FileList      []*QueryLocalFileUploadJobResponseBodyDataResultListFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	IotId         *string                                                      `json:"IotId,omitempty" xml:"IotId,omitempty"`
	ProductKey    *string                                                      `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	SlotEndTime   *int32                                                       `json:"SlotEndTime,omitempty" xml:"SlotEndTime,omitempty"`
	SlotStartTime *int32                                                       `json:"SlotStartTime,omitempty" xml:"SlotStartTime,omitempty"`
	SlotStatus    *int32                                                       `json:"SlotStatus,omitempty" xml:"SlotStatus,omitempty"`
}

func (s QueryLocalFileUploadJobResponseBodyDataResultList) String() string {
	return tea.Prettify(s)
}

func (s QueryLocalFileUploadJobResponseBodyDataResultList) GoString() string {
	return s.String()
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultList) SetCode(v int32) *QueryLocalFileUploadJobResponseBodyDataResultList {
	s.Code = &v
	return s
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultList) SetDeviceName(v string) *QueryLocalFileUploadJobResponseBodyDataResultList {
	s.DeviceName = &v
	return s
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultList) SetFileList(v []*QueryLocalFileUploadJobResponseBodyDataResultListFileList) *QueryLocalFileUploadJobResponseBodyDataResultList {
	s.FileList = v
	return s
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultList) SetIotId(v string) *QueryLocalFileUploadJobResponseBodyDataResultList {
	s.IotId = &v
	return s
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultList) SetProductKey(v string) *QueryLocalFileUploadJobResponseBodyDataResultList {
	s.ProductKey = &v
	return s
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultList) SetSlotEndTime(v int32) *QueryLocalFileUploadJobResponseBodyDataResultList {
	s.SlotEndTime = &v
	return s
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultList) SetSlotStartTime(v int32) *QueryLocalFileUploadJobResponseBodyDataResultList {
	s.SlotStartTime = &v
	return s
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultList) SetSlotStatus(v int32) *QueryLocalFileUploadJobResponseBodyDataResultList {
	s.SlotStatus = &v
	return s
}

type QueryLocalFileUploadJobResponseBodyDataResultListFileList struct {
	FileEndTime   *int32  `json:"FileEndTime,omitempty" xml:"FileEndTime,omitempty"`
	FileName      *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileStartTime *int32  `json:"FileStartTime,omitempty" xml:"FileStartTime,omitempty"`
}

func (s QueryLocalFileUploadJobResponseBodyDataResultListFileList) String() string {
	return tea.Prettify(s)
}

func (s QueryLocalFileUploadJobResponseBodyDataResultListFileList) GoString() string {
	return s.String()
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultListFileList) SetFileEndTime(v int32) *QueryLocalFileUploadJobResponseBodyDataResultListFileList {
	s.FileEndTime = &v
	return s
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultListFileList) SetFileName(v string) *QueryLocalFileUploadJobResponseBodyDataResultListFileList {
	s.FileName = &v
	return s
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultListFileList) SetFileStartTime(v int32) *QueryLocalFileUploadJobResponseBodyDataResultListFileList {
	s.FileStartTime = &v
	return s
}

type QueryLocalFileUploadJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryLocalFileUploadJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryLocalFileUploadJobResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryLocalFileUploadJobResponse) GoString() string {
	return s.String()
}

func (s *QueryLocalFileUploadJobResponse) SetHeaders(v map[string]*string) *QueryLocalFileUploadJobResponse {
	s.Headers = v
	return s
}

func (s *QueryLocalFileUploadJobResponse) SetStatusCode(v int32) *QueryLocalFileUploadJobResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryLocalFileUploadJobResponse) SetBody(v *QueryLocalFileUploadJobResponseBody) *QueryLocalFileUploadJobResponse {
	s.Body = v
	return s
}

type QueryMonthRecordRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	Month         *string `json:"Month,omitempty" xml:"Month,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryMonthRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMonthRecordRequest) GoString() string {
	return s.String()
}

func (s *QueryMonthRecordRequest) SetDeviceName(v string) *QueryMonthRecordRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryMonthRecordRequest) SetIotId(v string) *QueryMonthRecordRequest {
	s.IotId = &v
	return s
}

func (s *QueryMonthRecordRequest) SetIotInstanceId(v string) *QueryMonthRecordRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryMonthRecordRequest) SetMonth(v string) *QueryMonthRecordRequest {
	s.Month = &v
	return s
}

func (s *QueryMonthRecordRequest) SetProductKey(v string) *QueryMonthRecordRequest {
	s.ProductKey = &v
	return s
}

type QueryMonthRecordResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMonthRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMonthRecordResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMonthRecordResponseBody) SetCode(v string) *QueryMonthRecordResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMonthRecordResponseBody) SetData(v string) *QueryMonthRecordResponseBody {
	s.Data = &v
	return s
}

func (s *QueryMonthRecordResponseBody) SetErrorMessage(v string) *QueryMonthRecordResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryMonthRecordResponseBody) SetRequestId(v string) *QueryMonthRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMonthRecordResponseBody) SetSuccess(v bool) *QueryMonthRecordResponseBody {
	s.Success = &v
	return s
}

type QueryMonthRecordResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMonthRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMonthRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMonthRecordResponse) GoString() string {
	return s.String()
}

func (s *QueryMonthRecordResponse) SetHeaders(v map[string]*string) *QueryMonthRecordResponse {
	s.Headers = v
	return s
}

func (s *QueryMonthRecordResponse) SetStatusCode(v int32) *QueryMonthRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMonthRecordResponse) SetBody(v *QueryMonthRecordResponseBody) *QueryMonthRecordResponse {
	s.Body = v
	return s
}

type QueryPictureFilesRequest struct {
	BeginTime     *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	EndTime       *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PictureSource *int32  `json:"PictureSource,omitempty" xml:"PictureSource,omitempty"`
	PictureType   *int32  `json:"PictureType,omitempty" xml:"PictureType,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryPictureFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureFilesRequest) GoString() string {
	return s.String()
}

func (s *QueryPictureFilesRequest) SetBeginTime(v int64) *QueryPictureFilesRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryPictureFilesRequest) SetCurrentPage(v int32) *QueryPictureFilesRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryPictureFilesRequest) SetDeviceName(v string) *QueryPictureFilesRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryPictureFilesRequest) SetEndTime(v int64) *QueryPictureFilesRequest {
	s.EndTime = &v
	return s
}

func (s *QueryPictureFilesRequest) SetIotId(v string) *QueryPictureFilesRequest {
	s.IotId = &v
	return s
}

func (s *QueryPictureFilesRequest) SetIotInstanceId(v string) *QueryPictureFilesRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryPictureFilesRequest) SetPageSize(v int32) *QueryPictureFilesRequest {
	s.PageSize = &v
	return s
}

func (s *QueryPictureFilesRequest) SetPictureSource(v int32) *QueryPictureFilesRequest {
	s.PictureSource = &v
	return s
}

func (s *QueryPictureFilesRequest) SetPictureType(v int32) *QueryPictureFilesRequest {
	s.PictureType = &v
	return s
}

func (s *QueryPictureFilesRequest) SetProductKey(v string) *QueryPictureFilesRequest {
	s.ProductKey = &v
	return s
}

type QueryPictureFilesResponseBody struct {
	Code         *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryPictureFilesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryPictureFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureFilesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPictureFilesResponseBody) SetCode(v string) *QueryPictureFilesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPictureFilesResponseBody) SetData(v *QueryPictureFilesResponseBodyData) *QueryPictureFilesResponseBody {
	s.Data = v
	return s
}

func (s *QueryPictureFilesResponseBody) SetErrorMessage(v string) *QueryPictureFilesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryPictureFilesResponseBody) SetRequestId(v string) *QueryPictureFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPictureFilesResponseBody) SetSuccess(v bool) *QueryPictureFilesResponseBody {
	s.Success = &v
	return s
}

type QueryPictureFilesResponseBodyData struct {
	List     []*QueryPictureFilesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Page     *int32                                   `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryPictureFilesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureFilesResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPictureFilesResponseBodyData) SetList(v []*QueryPictureFilesResponseBodyDataList) *QueryPictureFilesResponseBodyData {
	s.List = v
	return s
}

func (s *QueryPictureFilesResponseBodyData) SetPage(v int32) *QueryPictureFilesResponseBodyData {
	s.Page = &v
	return s
}

func (s *QueryPictureFilesResponseBodyData) SetPageSize(v int32) *QueryPictureFilesResponseBodyData {
	s.PageSize = &v
	return s
}

type QueryPictureFilesResponseBodyDataList struct {
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	PicCreateTime *int64  `json:"PicCreateTime,omitempty" xml:"PicCreateTime,omitempty"`
	PicId         *string `json:"PicId,omitempty" xml:"PicId,omitempty"`
	PicUrl        *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	ThumbUrl      *string `json:"ThumbUrl,omitempty" xml:"ThumbUrl,omitempty"`
}

func (s QueryPictureFilesResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureFilesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryPictureFilesResponseBodyDataList) SetIotId(v string) *QueryPictureFilesResponseBodyDataList {
	s.IotId = &v
	return s
}

func (s *QueryPictureFilesResponseBodyDataList) SetPicCreateTime(v int64) *QueryPictureFilesResponseBodyDataList {
	s.PicCreateTime = &v
	return s
}

func (s *QueryPictureFilesResponseBodyDataList) SetPicId(v string) *QueryPictureFilesResponseBodyDataList {
	s.PicId = &v
	return s
}

func (s *QueryPictureFilesResponseBodyDataList) SetPicUrl(v string) *QueryPictureFilesResponseBodyDataList {
	s.PicUrl = &v
	return s
}

func (s *QueryPictureFilesResponseBodyDataList) SetThumbUrl(v string) *QueryPictureFilesResponseBodyDataList {
	s.ThumbUrl = &v
	return s
}

type QueryPictureFilesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPictureFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPictureFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureFilesResponse) GoString() string {
	return s.String()
}

func (s *QueryPictureFilesResponse) SetHeaders(v map[string]*string) *QueryPictureFilesResponse {
	s.Headers = v
	return s
}

func (s *QueryPictureFilesResponse) SetStatusCode(v int32) *QueryPictureFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPictureFilesResponse) SetBody(v *QueryPictureFilesResponseBody) *QueryPictureFilesResponse {
	s.Body = v
	return s
}

type QueryPictureSearchAiboxesRequest struct {
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryPictureSearchAiboxesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchAiboxesRequest) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchAiboxesRequest) SetAppInstanceId(v string) *QueryPictureSearchAiboxesRequest {
	s.AppInstanceId = &v
	return s
}

func (s *QueryPictureSearchAiboxesRequest) SetCurrentPage(v int32) *QueryPictureSearchAiboxesRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryPictureSearchAiboxesRequest) SetIotInstanceId(v string) *QueryPictureSearchAiboxesRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryPictureSearchAiboxesRequest) SetPageSize(v int32) *QueryPictureSearchAiboxesRequest {
	s.PageSize = &v
	return s
}

type QueryPictureSearchAiboxesResponseBody struct {
	Code         *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryPictureSearchAiboxesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryPictureSearchAiboxesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchAiboxesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchAiboxesResponseBody) SetCode(v string) *QueryPictureSearchAiboxesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPictureSearchAiboxesResponseBody) SetData(v *QueryPictureSearchAiboxesResponseBodyData) *QueryPictureSearchAiboxesResponseBody {
	s.Data = v
	return s
}

func (s *QueryPictureSearchAiboxesResponseBody) SetErrorMessage(v string) *QueryPictureSearchAiboxesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryPictureSearchAiboxesResponseBody) SetRequestId(v string) *QueryPictureSearchAiboxesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPictureSearchAiboxesResponseBody) SetSuccess(v bool) *QueryPictureSearchAiboxesResponseBody {
	s.Success = &v
	return s
}

type QueryPictureSearchAiboxesResponseBodyData struct {
	CurrentPage *int32                                               `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageCount   *int32                                               `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageData    []*QueryPictureSearchAiboxesResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	PageSize    *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total       *int32                                               `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryPictureSearchAiboxesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchAiboxesResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchAiboxesResponseBodyData) SetCurrentPage(v int32) *QueryPictureSearchAiboxesResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *QueryPictureSearchAiboxesResponseBodyData) SetPageCount(v int32) *QueryPictureSearchAiboxesResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryPictureSearchAiboxesResponseBodyData) SetPageData(v []*QueryPictureSearchAiboxesResponseBodyDataPageData) *QueryPictureSearchAiboxesResponseBodyData {
	s.PageData = v
	return s
}

func (s *QueryPictureSearchAiboxesResponseBodyData) SetPageSize(v int32) *QueryPictureSearchAiboxesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryPictureSearchAiboxesResponseBodyData) SetTotal(v int32) *QueryPictureSearchAiboxesResponseBodyData {
	s.Total = &v
	return s
}

type QueryPictureSearchAiboxesResponseBodyDataPageData struct {
	IotId    *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
}

func (s QueryPictureSearchAiboxesResponseBodyDataPageData) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchAiboxesResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchAiboxesResponseBodyDataPageData) SetIotId(v string) *QueryPictureSearchAiboxesResponseBodyDataPageData {
	s.IotId = &v
	return s
}

func (s *QueryPictureSearchAiboxesResponseBodyDataPageData) SetNickName(v string) *QueryPictureSearchAiboxesResponseBodyDataPageData {
	s.NickName = &v
	return s
}

type QueryPictureSearchAiboxesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPictureSearchAiboxesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPictureSearchAiboxesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchAiboxesResponse) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchAiboxesResponse) SetHeaders(v map[string]*string) *QueryPictureSearchAiboxesResponse {
	s.Headers = v
	return s
}

func (s *QueryPictureSearchAiboxesResponse) SetStatusCode(v int32) *QueryPictureSearchAiboxesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPictureSearchAiboxesResponse) SetBody(v *QueryPictureSearchAiboxesResponseBody) *QueryPictureSearchAiboxesResponse {
	s.Body = v
	return s
}

type QueryPictureSearchAppsRequest struct {
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryPictureSearchAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchAppsRequest) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchAppsRequest) SetCurrentPage(v int32) *QueryPictureSearchAppsRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryPictureSearchAppsRequest) SetIotInstanceId(v string) *QueryPictureSearchAppsRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryPictureSearchAppsRequest) SetPageSize(v int32) *QueryPictureSearchAppsRequest {
	s.PageSize = &v
	return s
}

type QueryPictureSearchAppsResponseBody struct {
	Code         *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryPictureSearchAppsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryPictureSearchAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchAppsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchAppsResponseBody) SetCode(v string) *QueryPictureSearchAppsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPictureSearchAppsResponseBody) SetData(v *QueryPictureSearchAppsResponseBodyData) *QueryPictureSearchAppsResponseBody {
	s.Data = v
	return s
}

func (s *QueryPictureSearchAppsResponseBody) SetErrorMessage(v string) *QueryPictureSearchAppsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryPictureSearchAppsResponseBody) SetRequestId(v string) *QueryPictureSearchAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPictureSearchAppsResponseBody) SetSuccess(v bool) *QueryPictureSearchAppsResponseBody {
	s.Success = &v
	return s
}

type QueryPictureSearchAppsResponseBodyData struct {
	CurrentPage *int32                                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageCount   *int32                                            `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageData    []*QueryPictureSearchAppsResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	PageSize    *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total       *int32                                            `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryPictureSearchAppsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchAppsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchAppsResponseBodyData) SetCurrentPage(v int32) *QueryPictureSearchAppsResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *QueryPictureSearchAppsResponseBodyData) SetPageCount(v int32) *QueryPictureSearchAppsResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryPictureSearchAppsResponseBodyData) SetPageData(v []*QueryPictureSearchAppsResponseBodyDataPageData) *QueryPictureSearchAppsResponseBodyData {
	s.PageData = v
	return s
}

func (s *QueryPictureSearchAppsResponseBodyData) SetPageSize(v int32) *QueryPictureSearchAppsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryPictureSearchAppsResponseBodyData) SetTotal(v int32) *QueryPictureSearchAppsResponseBodyData {
	s.Total = &v
	return s
}

type QueryPictureSearchAppsResponseBodyDataPageData struct {
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	AppTemplateId *string `json:"AppTemplateId,omitempty" xml:"AppTemplateId,omitempty"`
	CreateTime    *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifiedTime  *int64  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s QueryPictureSearchAppsResponseBodyDataPageData) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchAppsResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchAppsResponseBodyDataPageData) SetAppInstanceId(v string) *QueryPictureSearchAppsResponseBodyDataPageData {
	s.AppInstanceId = &v
	return s
}

func (s *QueryPictureSearchAppsResponseBodyDataPageData) SetAppTemplateId(v string) *QueryPictureSearchAppsResponseBodyDataPageData {
	s.AppTemplateId = &v
	return s
}

func (s *QueryPictureSearchAppsResponseBodyDataPageData) SetCreateTime(v int64) *QueryPictureSearchAppsResponseBodyDataPageData {
	s.CreateTime = &v
	return s
}

func (s *QueryPictureSearchAppsResponseBodyDataPageData) SetDescription(v string) *QueryPictureSearchAppsResponseBodyDataPageData {
	s.Description = &v
	return s
}

func (s *QueryPictureSearchAppsResponseBodyDataPageData) SetModifiedTime(v int64) *QueryPictureSearchAppsResponseBodyDataPageData {
	s.ModifiedTime = &v
	return s
}

func (s *QueryPictureSearchAppsResponseBodyDataPageData) SetName(v string) *QueryPictureSearchAppsResponseBodyDataPageData {
	s.Name = &v
	return s
}

func (s *QueryPictureSearchAppsResponseBodyDataPageData) SetVersion(v string) *QueryPictureSearchAppsResponseBodyDataPageData {
	s.Version = &v
	return s
}

type QueryPictureSearchAppsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPictureSearchAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPictureSearchAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchAppsResponse) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchAppsResponse) SetHeaders(v map[string]*string) *QueryPictureSearchAppsResponse {
	s.Headers = v
	return s
}

func (s *QueryPictureSearchAppsResponse) SetStatusCode(v int32) *QueryPictureSearchAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPictureSearchAppsResponse) SetBody(v *QueryPictureSearchAppsResponseBody) *QueryPictureSearchAppsResponse {
	s.Body = v
	return s
}

type QueryPictureSearchDevicesRequest struct {
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryPictureSearchDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchDevicesRequest) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchDevicesRequest) SetAppInstanceId(v string) *QueryPictureSearchDevicesRequest {
	s.AppInstanceId = &v
	return s
}

func (s *QueryPictureSearchDevicesRequest) SetCurrentPage(v int32) *QueryPictureSearchDevicesRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryPictureSearchDevicesRequest) SetPageSize(v int32) *QueryPictureSearchDevicesRequest {
	s.PageSize = &v
	return s
}

type QueryPictureSearchDevicesResponseBody struct {
	Code         *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryPictureSearchDevicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryPictureSearchDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchDevicesResponseBody) SetCode(v string) *QueryPictureSearchDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPictureSearchDevicesResponseBody) SetData(v *QueryPictureSearchDevicesResponseBodyData) *QueryPictureSearchDevicesResponseBody {
	s.Data = v
	return s
}

func (s *QueryPictureSearchDevicesResponseBody) SetErrorMessage(v string) *QueryPictureSearchDevicesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryPictureSearchDevicesResponseBody) SetRequestId(v string) *QueryPictureSearchDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPictureSearchDevicesResponseBody) SetSuccess(v bool) *QueryPictureSearchDevicesResponseBody {
	s.Success = &v
	return s
}

type QueryPictureSearchDevicesResponseBodyData struct {
	CurrentPage *int32                                               `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageCount   *int32                                               `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageData    []*QueryPictureSearchDevicesResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	PageSize    *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total       *int32                                               `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryPictureSearchDevicesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchDevicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchDevicesResponseBodyData) SetCurrentPage(v int32) *QueryPictureSearchDevicesResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *QueryPictureSearchDevicesResponseBodyData) SetPageCount(v int32) *QueryPictureSearchDevicesResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryPictureSearchDevicesResponseBodyData) SetPageData(v []*QueryPictureSearchDevicesResponseBodyDataPageData) *QueryPictureSearchDevicesResponseBodyData {
	s.PageData = v
	return s
}

func (s *QueryPictureSearchDevicesResponseBodyData) SetPageSize(v int32) *QueryPictureSearchDevicesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryPictureSearchDevicesResponseBodyData) SetTotal(v int32) *QueryPictureSearchDevicesResponseBodyData {
	s.Total = &v
	return s
}

type QueryPictureSearchDevicesResponseBodyDataPageData struct {
	IotId    *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
}

func (s QueryPictureSearchDevicesResponseBodyDataPageData) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchDevicesResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchDevicesResponseBodyDataPageData) SetIotId(v string) *QueryPictureSearchDevicesResponseBodyDataPageData {
	s.IotId = &v
	return s
}

func (s *QueryPictureSearchDevicesResponseBodyDataPageData) SetNickName(v string) *QueryPictureSearchDevicesResponseBodyDataPageData {
	s.NickName = &v
	return s
}

type QueryPictureSearchDevicesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPictureSearchDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPictureSearchDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchDevicesResponse) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchDevicesResponse) SetHeaders(v map[string]*string) *QueryPictureSearchDevicesResponse {
	s.Headers = v
	return s
}

func (s *QueryPictureSearchDevicesResponse) SetStatusCode(v int32) *QueryPictureSearchDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPictureSearchDevicesResponse) SetBody(v *QueryPictureSearchDevicesResponseBody) *QueryPictureSearchDevicesResponse {
	s.Body = v
	return s
}

type QueryPictureSearchJobRequest struct {
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	JobStatus     *int32  `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryPictureSearchJobRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchJobRequest) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchJobRequest) SetAppInstanceId(v string) *QueryPictureSearchJobRequest {
	s.AppInstanceId = &v
	return s
}

func (s *QueryPictureSearchJobRequest) SetCurrentPage(v int32) *QueryPictureSearchJobRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryPictureSearchJobRequest) SetJobStatus(v int32) *QueryPictureSearchJobRequest {
	s.JobStatus = &v
	return s
}

func (s *QueryPictureSearchJobRequest) SetPageSize(v int32) *QueryPictureSearchJobRequest {
	s.PageSize = &v
	return s
}

type QueryPictureSearchJobResponseBody struct {
	Code         *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryPictureSearchJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryPictureSearchJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchJobResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchJobResponseBody) SetCode(v string) *QueryPictureSearchJobResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPictureSearchJobResponseBody) SetData(v *QueryPictureSearchJobResponseBodyData) *QueryPictureSearchJobResponseBody {
	s.Data = v
	return s
}

func (s *QueryPictureSearchJobResponseBody) SetErrorMessage(v string) *QueryPictureSearchJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryPictureSearchJobResponseBody) SetRequestId(v string) *QueryPictureSearchJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPictureSearchJobResponseBody) SetSuccess(v bool) *QueryPictureSearchJobResponseBody {
	s.Success = &v
	return s
}

type QueryPictureSearchJobResponseBodyData struct {
	CurrentPage *int32                                           `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageCount   *int32                                           `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageData    []*QueryPictureSearchJobResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	PageSize    *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total       *int32                                           `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryPictureSearchJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchJobResponseBodyData) SetCurrentPage(v int32) *QueryPictureSearchJobResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *QueryPictureSearchJobResponseBodyData) SetPageCount(v int32) *QueryPictureSearchJobResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryPictureSearchJobResponseBodyData) SetPageData(v []*QueryPictureSearchJobResponseBodyDataPageData) *QueryPictureSearchJobResponseBodyData {
	s.PageData = v
	return s
}

func (s *QueryPictureSearchJobResponseBodyData) SetPageSize(v int32) *QueryPictureSearchJobResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryPictureSearchJobResponseBodyData) SetTotal(v int32) *QueryPictureSearchJobResponseBodyData {
	s.Total = &v
	return s
}

type QueryPictureSearchJobResponseBodyDataPageData struct {
	CreateTime   *int64   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EndTime      *int64   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	JobId        *string  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobStatus    *int32   `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	SearchPicUrl *string  `json:"SearchPicUrl,omitempty" xml:"SearchPicUrl,omitempty"`
	StartTime    *int64   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Threshold    *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s QueryPictureSearchJobResponseBodyDataPageData) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchJobResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchJobResponseBodyDataPageData) SetCreateTime(v int64) *QueryPictureSearchJobResponseBodyDataPageData {
	s.CreateTime = &v
	return s
}

func (s *QueryPictureSearchJobResponseBodyDataPageData) SetEndTime(v int64) *QueryPictureSearchJobResponseBodyDataPageData {
	s.EndTime = &v
	return s
}

func (s *QueryPictureSearchJobResponseBodyDataPageData) SetJobId(v string) *QueryPictureSearchJobResponseBodyDataPageData {
	s.JobId = &v
	return s
}

func (s *QueryPictureSearchJobResponseBodyDataPageData) SetJobStatus(v int32) *QueryPictureSearchJobResponseBodyDataPageData {
	s.JobStatus = &v
	return s
}

func (s *QueryPictureSearchJobResponseBodyDataPageData) SetSearchPicUrl(v string) *QueryPictureSearchJobResponseBodyDataPageData {
	s.SearchPicUrl = &v
	return s
}

func (s *QueryPictureSearchJobResponseBodyDataPageData) SetStartTime(v int64) *QueryPictureSearchJobResponseBodyDataPageData {
	s.StartTime = &v
	return s
}

func (s *QueryPictureSearchJobResponseBodyDataPageData) SetThreshold(v float32) *QueryPictureSearchJobResponseBodyDataPageData {
	s.Threshold = &v
	return s
}

type QueryPictureSearchJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPictureSearchJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPictureSearchJobResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchJobResponse) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchJobResponse) SetHeaders(v map[string]*string) *QueryPictureSearchJobResponse {
	s.Headers = v
	return s
}

func (s *QueryPictureSearchJobResponse) SetStatusCode(v int32) *QueryPictureSearchJobResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPictureSearchJobResponse) SetBody(v *QueryPictureSearchJobResponseBody) *QueryPictureSearchJobResponse {
	s.Body = v
	return s
}

type QueryPictureSearchJobResultRequest struct {
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	JobId         *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryPictureSearchJobResultRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchJobResultRequest) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchJobResultRequest) SetAppInstanceId(v string) *QueryPictureSearchJobResultRequest {
	s.AppInstanceId = &v
	return s
}

func (s *QueryPictureSearchJobResultRequest) SetCurrentPage(v int32) *QueryPictureSearchJobResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryPictureSearchJobResultRequest) SetJobId(v string) *QueryPictureSearchJobResultRequest {
	s.JobId = &v
	return s
}

func (s *QueryPictureSearchJobResultRequest) SetPageSize(v int32) *QueryPictureSearchJobResultRequest {
	s.PageSize = &v
	return s
}

type QueryPictureSearchJobResultResponseBody struct {
	Code         *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryPictureSearchJobResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryPictureSearchJobResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchJobResultResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchJobResultResponseBody) SetCode(v string) *QueryPictureSearchJobResultResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPictureSearchJobResultResponseBody) SetData(v *QueryPictureSearchJobResultResponseBodyData) *QueryPictureSearchJobResultResponseBody {
	s.Data = v
	return s
}

func (s *QueryPictureSearchJobResultResponseBody) SetErrorMessage(v string) *QueryPictureSearchJobResultResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryPictureSearchJobResultResponseBody) SetRequestId(v string) *QueryPictureSearchJobResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPictureSearchJobResultResponseBody) SetSuccess(v bool) *QueryPictureSearchJobResultResponseBody {
	s.Success = &v
	return s
}

type QueryPictureSearchJobResultResponseBodyData struct {
	CurrentPage *int32                                                 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageCount   *int32                                                 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageData    []*QueryPictureSearchJobResultResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	PageSize    *int32                                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total       *int32                                                 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryPictureSearchJobResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchJobResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchJobResultResponseBodyData) SetCurrentPage(v int32) *QueryPictureSearchJobResultResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *QueryPictureSearchJobResultResponseBodyData) SetPageCount(v int32) *QueryPictureSearchJobResultResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryPictureSearchJobResultResponseBodyData) SetPageData(v []*QueryPictureSearchJobResultResponseBodyDataPageData) *QueryPictureSearchJobResultResponseBodyData {
	s.PageData = v
	return s
}

func (s *QueryPictureSearchJobResultResponseBodyData) SetPageSize(v int32) *QueryPictureSearchJobResultResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryPictureSearchJobResultResponseBodyData) SetTotal(v int32) *QueryPictureSearchJobResultResponseBodyData {
	s.Total = &v
	return s
}

type QueryPictureSearchJobResultResponseBodyDataPageData struct {
	DeviceNickName *string  `json:"DeviceNickName,omitempty" xml:"DeviceNickName,omitempty"`
	EventTime      *int64   `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	GatewayIotId   *string  `json:"GatewayIotId,omitempty" xml:"GatewayIotId,omitempty"`
	IotId          *string  `json:"IotId,omitempty" xml:"IotId,omitempty"`
	PicUrl         *string  `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	Threshold      *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	VectorId       *string  `json:"VectorId,omitempty" xml:"VectorId,omitempty"`
}

func (s QueryPictureSearchJobResultResponseBodyDataPageData) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchJobResultResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchJobResultResponseBodyDataPageData) SetDeviceNickName(v string) *QueryPictureSearchJobResultResponseBodyDataPageData {
	s.DeviceNickName = &v
	return s
}

func (s *QueryPictureSearchJobResultResponseBodyDataPageData) SetEventTime(v int64) *QueryPictureSearchJobResultResponseBodyDataPageData {
	s.EventTime = &v
	return s
}

func (s *QueryPictureSearchJobResultResponseBodyDataPageData) SetGatewayIotId(v string) *QueryPictureSearchJobResultResponseBodyDataPageData {
	s.GatewayIotId = &v
	return s
}

func (s *QueryPictureSearchJobResultResponseBodyDataPageData) SetIotId(v string) *QueryPictureSearchJobResultResponseBodyDataPageData {
	s.IotId = &v
	return s
}

func (s *QueryPictureSearchJobResultResponseBodyDataPageData) SetPicUrl(v string) *QueryPictureSearchJobResultResponseBodyDataPageData {
	s.PicUrl = &v
	return s
}

func (s *QueryPictureSearchJobResultResponseBodyDataPageData) SetThreshold(v float32) *QueryPictureSearchJobResultResponseBodyDataPageData {
	s.Threshold = &v
	return s
}

func (s *QueryPictureSearchJobResultResponseBodyDataPageData) SetVectorId(v string) *QueryPictureSearchJobResultResponseBodyDataPageData {
	s.VectorId = &v
	return s
}

type QueryPictureSearchJobResultResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPictureSearchJobResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPictureSearchJobResultResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchJobResultResponse) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchJobResultResponse) SetHeaders(v map[string]*string) *QueryPictureSearchJobResultResponse {
	s.Headers = v
	return s
}

func (s *QueryPictureSearchJobResultResponse) SetStatusCode(v int32) *QueryPictureSearchJobResultResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPictureSearchJobResultResponse) SetBody(v *QueryPictureSearchJobResultResponseBody) *QueryPictureSearchJobResultResponse {
	s.Body = v
	return s
}

type QueryRecordRequest struct {
	BeginTime     *int32  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	EndTime       *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	NeedSnapshot  *bool   `json:"NeedSnapshot,omitempty" xml:"NeedSnapshot,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	RecordType    *int32  `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	StreamType    *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s QueryRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordRequest) SetBeginTime(v int32) *QueryRecordRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryRecordRequest) SetCurrentPage(v int32) *QueryRecordRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryRecordRequest) SetDeviceName(v string) *QueryRecordRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryRecordRequest) SetEndTime(v int32) *QueryRecordRequest {
	s.EndTime = &v
	return s
}

func (s *QueryRecordRequest) SetIotId(v string) *QueryRecordRequest {
	s.IotId = &v
	return s
}

func (s *QueryRecordRequest) SetIotInstanceId(v string) *QueryRecordRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryRecordRequest) SetNeedSnapshot(v bool) *QueryRecordRequest {
	s.NeedSnapshot = &v
	return s
}

func (s *QueryRecordRequest) SetPageSize(v int32) *QueryRecordRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRecordRequest) SetProductKey(v string) *QueryRecordRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryRecordRequest) SetRecordType(v int32) *QueryRecordRequest {
	s.RecordType = &v
	return s
}

func (s *QueryRecordRequest) SetStreamType(v int32) *QueryRecordRequest {
	s.StreamType = &v
	return s
}

type QueryRecordResponseBody struct {
	Code         *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecordResponseBody) SetCode(v string) *QueryRecordResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRecordResponseBody) SetData(v *QueryRecordResponseBodyData) *QueryRecordResponseBody {
	s.Data = v
	return s
}

func (s *QueryRecordResponseBody) SetErrorMessage(v string) *QueryRecordResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryRecordResponseBody) SetRequestId(v string) *QueryRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRecordResponseBody) SetSuccess(v bool) *QueryRecordResponseBody {
	s.Success = &v
	return s
}

type QueryRecordResponseBodyData struct {
	List     []*QueryRecordResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Page     *int32                             `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryRecordResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRecordResponseBodyData) SetList(v []*QueryRecordResponseBodyDataList) *QueryRecordResponseBodyData {
	s.List = v
	return s
}

func (s *QueryRecordResponseBodyData) SetPage(v int32) *QueryRecordResponseBodyData {
	s.Page = &v
	return s
}

func (s *QueryRecordResponseBodyData) SetPageSize(v int32) *QueryRecordResponseBodyData {
	s.PageSize = &v
	return s
}

type QueryRecordResponseBodyDataList struct {
	BeginTime        *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime          *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EventType        *int32  `json:"EventType,omitempty" xml:"EventType,omitempty"`
	FileName         *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileSize         *int32  `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	RecordType       *int32  `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	SnapshotUrl      *string `json:"SnapshotUrl,omitempty" xml:"SnapshotUrl,omitempty"`
	StreamType       *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	VideoFrameNumber *int32  `json:"VideoFrameNumber,omitempty" xml:"VideoFrameNumber,omitempty"`
}

func (s QueryRecordResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryRecordResponseBodyDataList) SetBeginTime(v string) *QueryRecordResponseBodyDataList {
	s.BeginTime = &v
	return s
}

func (s *QueryRecordResponseBodyDataList) SetEndTime(v string) *QueryRecordResponseBodyDataList {
	s.EndTime = &v
	return s
}

func (s *QueryRecordResponseBodyDataList) SetEventType(v int32) *QueryRecordResponseBodyDataList {
	s.EventType = &v
	return s
}

func (s *QueryRecordResponseBodyDataList) SetFileName(v string) *QueryRecordResponseBodyDataList {
	s.FileName = &v
	return s
}

func (s *QueryRecordResponseBodyDataList) SetFileSize(v int32) *QueryRecordResponseBodyDataList {
	s.FileSize = &v
	return s
}

func (s *QueryRecordResponseBodyDataList) SetRecordType(v int32) *QueryRecordResponseBodyDataList {
	s.RecordType = &v
	return s
}

func (s *QueryRecordResponseBodyDataList) SetSnapshotUrl(v string) *QueryRecordResponseBodyDataList {
	s.SnapshotUrl = &v
	return s
}

func (s *QueryRecordResponseBodyDataList) SetStreamType(v int32) *QueryRecordResponseBodyDataList {
	s.StreamType = &v
	return s
}

func (s *QueryRecordResponseBodyDataList) SetVideoFrameNumber(v int32) *QueryRecordResponseBodyDataList {
	s.VideoFrameNumber = &v
	return s
}

type QueryRecordResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordResponse) SetHeaders(v map[string]*string) *QueryRecordResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordResponse) SetStatusCode(v int32) *QueryRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordResponse) SetBody(v *QueryRecordResponseBody) *QueryRecordResponse {
	s.Body = v
	return s
}

type QueryRecordByRecordIdRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	RecordId      *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s QueryRecordByRecordIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordByRecordIdRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordByRecordIdRequest) SetDeviceName(v string) *QueryRecordByRecordIdRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryRecordByRecordIdRequest) SetIotId(v string) *QueryRecordByRecordIdRequest {
	s.IotId = &v
	return s
}

func (s *QueryRecordByRecordIdRequest) SetIotInstanceId(v string) *QueryRecordByRecordIdRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryRecordByRecordIdRequest) SetProductKey(v string) *QueryRecordByRecordIdRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryRecordByRecordIdRequest) SetRecordId(v string) *QueryRecordByRecordIdRequest {
	s.RecordId = &v
	return s
}

type QueryRecordByRecordIdResponseBody struct {
	Code         *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         []*QueryRecordByRecordIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMessage *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRecordByRecordIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordByRecordIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecordByRecordIdResponseBody) SetCode(v string) *QueryRecordByRecordIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRecordByRecordIdResponseBody) SetData(v []*QueryRecordByRecordIdResponseBodyData) *QueryRecordByRecordIdResponseBody {
	s.Data = v
	return s
}

func (s *QueryRecordByRecordIdResponseBody) SetErrorMessage(v string) *QueryRecordByRecordIdResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryRecordByRecordIdResponseBody) SetRequestId(v string) *QueryRecordByRecordIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRecordByRecordIdResponseBody) SetSuccess(v bool) *QueryRecordByRecordIdResponseBody {
	s.Success = &v
	return s
}

type QueryRecordByRecordIdResponseBodyData struct {
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FileName  *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	VodUrl    *string `json:"VodUrl,omitempty" xml:"VodUrl,omitempty"`
}

func (s QueryRecordByRecordIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordByRecordIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRecordByRecordIdResponseBodyData) SetBeginTime(v string) *QueryRecordByRecordIdResponseBodyData {
	s.BeginTime = &v
	return s
}

func (s *QueryRecordByRecordIdResponseBodyData) SetEndTime(v string) *QueryRecordByRecordIdResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *QueryRecordByRecordIdResponseBodyData) SetFileName(v string) *QueryRecordByRecordIdResponseBodyData {
	s.FileName = &v
	return s
}

func (s *QueryRecordByRecordIdResponseBodyData) SetVodUrl(v string) *QueryRecordByRecordIdResponseBodyData {
	s.VodUrl = &v
	return s
}

type QueryRecordByRecordIdResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecordByRecordIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecordByRecordIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordByRecordIdResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordByRecordIdResponse) SetHeaders(v map[string]*string) *QueryRecordByRecordIdResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordByRecordIdResponse) SetStatusCode(v int32) *QueryRecordByRecordIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordByRecordIdResponse) SetBody(v *QueryRecordByRecordIdResponseBody) *QueryRecordByRecordIdResponse {
	s.Body = v
	return s
}

type QueryRecordDownloadJobByIdRequest struct {
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	JobId         *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s QueryRecordDownloadJobByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordDownloadJobByIdRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordDownloadJobByIdRequest) SetIotInstanceId(v string) *QueryRecordDownloadJobByIdRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryRecordDownloadJobByIdRequest) SetJobId(v string) *QueryRecordDownloadJobByIdRequest {
	s.JobId = &v
	return s
}

type QueryRecordDownloadJobByIdResponseBody struct {
	Code         *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryRecordDownloadJobByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                     `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRecordDownloadJobByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordDownloadJobByIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecordDownloadJobByIdResponseBody) SetCode(v string) *QueryRecordDownloadJobByIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRecordDownloadJobByIdResponseBody) SetData(v *QueryRecordDownloadJobByIdResponseBodyData) *QueryRecordDownloadJobByIdResponseBody {
	s.Data = v
	return s
}

func (s *QueryRecordDownloadJobByIdResponseBody) SetErrorMessage(v string) *QueryRecordDownloadJobByIdResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryRecordDownloadJobByIdResponseBody) SetRequestId(v string) *QueryRecordDownloadJobByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRecordDownloadJobByIdResponseBody) SetSuccess(v bool) *QueryRecordDownloadJobByIdResponseBody {
	s.Success = &v
	return s
}

type QueryRecordDownloadJobByIdResponseBodyData struct {
	BeginTime    *int32  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime      *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	IotId        *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	JobErrorCode *int32  `json:"JobErrorCode,omitempty" xml:"JobErrorCode,omitempty"`
	Progress     *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RecordType   *int32  `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	Status       *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	StreamType   *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	Type         *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s QueryRecordDownloadJobByIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordDownloadJobByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRecordDownloadJobByIdResponseBodyData) SetBeginTime(v int32) *QueryRecordDownloadJobByIdResponseBodyData {
	s.BeginTime = &v
	return s
}

func (s *QueryRecordDownloadJobByIdResponseBodyData) SetEndTime(v int32) *QueryRecordDownloadJobByIdResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *QueryRecordDownloadJobByIdResponseBodyData) SetFileName(v string) *QueryRecordDownloadJobByIdResponseBodyData {
	s.FileName = &v
	return s
}

func (s *QueryRecordDownloadJobByIdResponseBodyData) SetIotId(v string) *QueryRecordDownloadJobByIdResponseBodyData {
	s.IotId = &v
	return s
}

func (s *QueryRecordDownloadJobByIdResponseBodyData) SetJobErrorCode(v int32) *QueryRecordDownloadJobByIdResponseBodyData {
	s.JobErrorCode = &v
	return s
}

func (s *QueryRecordDownloadJobByIdResponseBodyData) SetProgress(v int32) *QueryRecordDownloadJobByIdResponseBodyData {
	s.Progress = &v
	return s
}

func (s *QueryRecordDownloadJobByIdResponseBodyData) SetRecordType(v int32) *QueryRecordDownloadJobByIdResponseBodyData {
	s.RecordType = &v
	return s
}

func (s *QueryRecordDownloadJobByIdResponseBodyData) SetStatus(v int32) *QueryRecordDownloadJobByIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryRecordDownloadJobByIdResponseBodyData) SetStreamType(v int32) *QueryRecordDownloadJobByIdResponseBodyData {
	s.StreamType = &v
	return s
}

func (s *QueryRecordDownloadJobByIdResponseBodyData) SetType(v int32) *QueryRecordDownloadJobByIdResponseBodyData {
	s.Type = &v
	return s
}

func (s *QueryRecordDownloadJobByIdResponseBodyData) SetUrl(v string) *QueryRecordDownloadJobByIdResponseBodyData {
	s.Url = &v
	return s
}

type QueryRecordDownloadJobByIdResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecordDownloadJobByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecordDownloadJobByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordDownloadJobByIdResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordDownloadJobByIdResponse) SetHeaders(v map[string]*string) *QueryRecordDownloadJobByIdResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordDownloadJobByIdResponse) SetStatusCode(v int32) *QueryRecordDownloadJobByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordDownloadJobByIdResponse) SetBody(v *QueryRecordDownloadJobByIdResponseBody) *QueryRecordDownloadJobByIdResponse {
	s.Body = v
	return s
}

type QueryRecordDownloadJobListRequest struct {
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryRecordDownloadJobListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordDownloadJobListRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordDownloadJobListRequest) SetCurrentPage(v int32) *QueryRecordDownloadJobListRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryRecordDownloadJobListRequest) SetDeviceName(v string) *QueryRecordDownloadJobListRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryRecordDownloadJobListRequest) SetIotId(v string) *QueryRecordDownloadJobListRequest {
	s.IotId = &v
	return s
}

func (s *QueryRecordDownloadJobListRequest) SetIotInstanceId(v string) *QueryRecordDownloadJobListRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryRecordDownloadJobListRequest) SetPageSize(v int32) *QueryRecordDownloadJobListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRecordDownloadJobListRequest) SetProductKey(v string) *QueryRecordDownloadJobListRequest {
	s.ProductKey = &v
	return s
}

type QueryRecordDownloadJobListResponseBody struct {
	Code         *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryRecordDownloadJobListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                     `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRecordDownloadJobListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordDownloadJobListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecordDownloadJobListResponseBody) SetCode(v string) *QueryRecordDownloadJobListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRecordDownloadJobListResponseBody) SetData(v *QueryRecordDownloadJobListResponseBodyData) *QueryRecordDownloadJobListResponseBody {
	s.Data = v
	return s
}

func (s *QueryRecordDownloadJobListResponseBody) SetErrorMessage(v string) *QueryRecordDownloadJobListResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryRecordDownloadJobListResponseBody) SetRequestId(v string) *QueryRecordDownloadJobListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRecordDownloadJobListResponseBody) SetSuccess(v bool) *QueryRecordDownloadJobListResponseBody {
	s.Success = &v
	return s
}

type QueryRecordDownloadJobListResponseBodyData struct {
	JobList []*QueryRecordDownloadJobListResponseBodyDataJobList `json:"JobList,omitempty" xml:"JobList,omitempty" type:"Repeated"`
	Total   *int32                                               `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryRecordDownloadJobListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordDownloadJobListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRecordDownloadJobListResponseBodyData) SetJobList(v []*QueryRecordDownloadJobListResponseBodyDataJobList) *QueryRecordDownloadJobListResponseBodyData {
	s.JobList = v
	return s
}

func (s *QueryRecordDownloadJobListResponseBodyData) SetTotal(v int32) *QueryRecordDownloadJobListResponseBodyData {
	s.Total = &v
	return s
}

type QueryRecordDownloadJobListResponseBodyDataJobList struct {
	BeginTime    *int32  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime      *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	IotId        *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	JobErrorCode *int32  `json:"JobErrorCode,omitempty" xml:"JobErrorCode,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Progress     *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RecordType   *int32  `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	Status       *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	StreamType   *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	Type         *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s QueryRecordDownloadJobListResponseBodyDataJobList) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordDownloadJobListResponseBodyDataJobList) GoString() string {
	return s.String()
}

func (s *QueryRecordDownloadJobListResponseBodyDataJobList) SetBeginTime(v int32) *QueryRecordDownloadJobListResponseBodyDataJobList {
	s.BeginTime = &v
	return s
}

func (s *QueryRecordDownloadJobListResponseBodyDataJobList) SetEndTime(v int32) *QueryRecordDownloadJobListResponseBodyDataJobList {
	s.EndTime = &v
	return s
}

func (s *QueryRecordDownloadJobListResponseBodyDataJobList) SetFileName(v string) *QueryRecordDownloadJobListResponseBodyDataJobList {
	s.FileName = &v
	return s
}

func (s *QueryRecordDownloadJobListResponseBodyDataJobList) SetIotId(v string) *QueryRecordDownloadJobListResponseBodyDataJobList {
	s.IotId = &v
	return s
}

func (s *QueryRecordDownloadJobListResponseBodyDataJobList) SetJobErrorCode(v int32) *QueryRecordDownloadJobListResponseBodyDataJobList {
	s.JobErrorCode = &v
	return s
}

func (s *QueryRecordDownloadJobListResponseBodyDataJobList) SetJobId(v string) *QueryRecordDownloadJobListResponseBodyDataJobList {
	s.JobId = &v
	return s
}

func (s *QueryRecordDownloadJobListResponseBodyDataJobList) SetProgress(v int32) *QueryRecordDownloadJobListResponseBodyDataJobList {
	s.Progress = &v
	return s
}

func (s *QueryRecordDownloadJobListResponseBodyDataJobList) SetRecordType(v int32) *QueryRecordDownloadJobListResponseBodyDataJobList {
	s.RecordType = &v
	return s
}

func (s *QueryRecordDownloadJobListResponseBodyDataJobList) SetStatus(v int32) *QueryRecordDownloadJobListResponseBodyDataJobList {
	s.Status = &v
	return s
}

func (s *QueryRecordDownloadJobListResponseBodyDataJobList) SetStreamType(v int32) *QueryRecordDownloadJobListResponseBodyDataJobList {
	s.StreamType = &v
	return s
}

func (s *QueryRecordDownloadJobListResponseBodyDataJobList) SetType(v int32) *QueryRecordDownloadJobListResponseBodyDataJobList {
	s.Type = &v
	return s
}

func (s *QueryRecordDownloadJobListResponseBodyDataJobList) SetUrl(v string) *QueryRecordDownloadJobListResponseBodyDataJobList {
	s.Url = &v
	return s
}

type QueryRecordDownloadJobListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecordDownloadJobListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecordDownloadJobListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordDownloadJobListResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordDownloadJobListResponse) SetHeaders(v map[string]*string) *QueryRecordDownloadJobListResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordDownloadJobListResponse) SetStatusCode(v int32) *QueryRecordDownloadJobListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordDownloadJobListResponse) SetBody(v *QueryRecordDownloadJobListResponseBody) *QueryRecordDownloadJobListResponse {
	s.Body = v
	return s
}

type QueryRecordDownloadUrlRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	FileName      *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryRecordDownloadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordDownloadUrlRequest) SetDeviceName(v string) *QueryRecordDownloadUrlRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryRecordDownloadUrlRequest) SetFileName(v string) *QueryRecordDownloadUrlRequest {
	s.FileName = &v
	return s
}

func (s *QueryRecordDownloadUrlRequest) SetIotId(v string) *QueryRecordDownloadUrlRequest {
	s.IotId = &v
	return s
}

func (s *QueryRecordDownloadUrlRequest) SetIotInstanceId(v string) *QueryRecordDownloadUrlRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryRecordDownloadUrlRequest) SetProductKey(v string) *QueryRecordDownloadUrlRequest {
	s.ProductKey = &v
	return s
}

type QueryRecordDownloadUrlResponseBody struct {
	Code         *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryRecordDownloadUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRecordDownloadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecordDownloadUrlResponseBody) SetCode(v string) *QueryRecordDownloadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRecordDownloadUrlResponseBody) SetData(v *QueryRecordDownloadUrlResponseBodyData) *QueryRecordDownloadUrlResponseBody {
	s.Data = v
	return s
}

func (s *QueryRecordDownloadUrlResponseBody) SetErrorMessage(v string) *QueryRecordDownloadUrlResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryRecordDownloadUrlResponseBody) SetRequestId(v string) *QueryRecordDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRecordDownloadUrlResponseBody) SetSuccess(v bool) *QueryRecordDownloadUrlResponseBody {
	s.Success = &v
	return s
}

type QueryRecordDownloadUrlResponseBodyData struct {
	Progress *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status   *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Url      *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s QueryRecordDownloadUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordDownloadUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRecordDownloadUrlResponseBodyData) SetProgress(v int32) *QueryRecordDownloadUrlResponseBodyData {
	s.Progress = &v
	return s
}

func (s *QueryRecordDownloadUrlResponseBodyData) SetStatus(v int32) *QueryRecordDownloadUrlResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryRecordDownloadUrlResponseBodyData) SetUrl(v string) *QueryRecordDownloadUrlResponseBodyData {
	s.Url = &v
	return s
}

type QueryRecordDownloadUrlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecordDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecordDownloadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordDownloadUrlResponse) SetHeaders(v map[string]*string) *QueryRecordDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordDownloadUrlResponse) SetStatusCode(v int32) *QueryRecordDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordDownloadUrlResponse) SetBody(v *QueryRecordDownloadUrlResponseBody) *QueryRecordDownloadUrlResponse {
	s.Body = v
	return s
}

type QueryRecordPlanDetailRequest struct {
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s QueryRecordPlanDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDetailRequest) SetPlanId(v string) *QueryRecordPlanDetailRequest {
	s.PlanId = &v
	return s
}

type QueryRecordPlanDetailResponseBody struct {
	Code         *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryRecordPlanDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRecordPlanDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDetailResponseBody) SetCode(v string) *QueryRecordPlanDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRecordPlanDetailResponseBody) SetData(v *QueryRecordPlanDetailResponseBodyData) *QueryRecordPlanDetailResponseBody {
	s.Data = v
	return s
}

func (s *QueryRecordPlanDetailResponseBody) SetErrorMessage(v string) *QueryRecordPlanDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryRecordPlanDetailResponseBody) SetRequestId(v string) *QueryRecordPlanDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRecordPlanDetailResponseBody) SetSuccess(v bool) *QueryRecordPlanDetailResponseBody {
	s.Success = &v
	return s
}

type QueryRecordPlanDetailResponseBodyData struct {
	Name         *string                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	PlanId       *string                                            `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	TemplateId   *string                                            `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateInfo *QueryRecordPlanDetailResponseBodyDataTemplateInfo `json:"TemplateInfo,omitempty" xml:"TemplateInfo,omitempty" type:"Struct"`
}

func (s QueryRecordPlanDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDetailResponseBodyData) SetName(v string) *QueryRecordPlanDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryRecordPlanDetailResponseBodyData) SetPlanId(v string) *QueryRecordPlanDetailResponseBodyData {
	s.PlanId = &v
	return s
}

func (s *QueryRecordPlanDetailResponseBodyData) SetTemplateId(v string) *QueryRecordPlanDetailResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *QueryRecordPlanDetailResponseBodyData) SetTemplateInfo(v *QueryRecordPlanDetailResponseBodyDataTemplateInfo) *QueryRecordPlanDetailResponseBodyData {
	s.TemplateInfo = v
	return s
}

type QueryRecordPlanDetailResponseBodyDataTemplateInfo struct {
	AllDay          *int32                                                              `json:"AllDay,omitempty" xml:"AllDay,omitempty"`
	Default         *int32                                                              `json:"Default,omitempty" xml:"Default,omitempty"`
	Name            *string                                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateId      *string                                                             `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TimeSectionList []*QueryRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList `json:"TimeSectionList,omitempty" xml:"TimeSectionList,omitempty" type:"Repeated"`
}

func (s QueryRecordPlanDetailResponseBodyDataTemplateInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDetailResponseBodyDataTemplateInfo) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDetailResponseBodyDataTemplateInfo) SetAllDay(v int32) *QueryRecordPlanDetailResponseBodyDataTemplateInfo {
	s.AllDay = &v
	return s
}

func (s *QueryRecordPlanDetailResponseBodyDataTemplateInfo) SetDefault(v int32) *QueryRecordPlanDetailResponseBodyDataTemplateInfo {
	s.Default = &v
	return s
}

func (s *QueryRecordPlanDetailResponseBodyDataTemplateInfo) SetName(v string) *QueryRecordPlanDetailResponseBodyDataTemplateInfo {
	s.Name = &v
	return s
}

func (s *QueryRecordPlanDetailResponseBodyDataTemplateInfo) SetTemplateId(v string) *QueryRecordPlanDetailResponseBodyDataTemplateInfo {
	s.TemplateId = &v
	return s
}

func (s *QueryRecordPlanDetailResponseBodyDataTemplateInfo) SetTimeSectionList(v []*QueryRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList) *QueryRecordPlanDetailResponseBodyDataTemplateInfo {
	s.TimeSectionList = v
	return s
}

type QueryRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList struct {
	Begin     *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	DayOfWeek *int32 `json:"DayOfWeek,omitempty" xml:"DayOfWeek,omitempty"`
	End       *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s QueryRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList) SetBegin(v int32) *QueryRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList {
	s.Begin = &v
	return s
}

func (s *QueryRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList) SetDayOfWeek(v int32) *QueryRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList {
	s.DayOfWeek = &v
	return s
}

func (s *QueryRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList) SetEnd(v int32) *QueryRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList {
	s.End = &v
	return s
}

type QueryRecordPlanDetailResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecordPlanDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecordPlanDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDetailResponse) SetHeaders(v map[string]*string) *QueryRecordPlanDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordPlanDetailResponse) SetStatusCode(v int32) *QueryRecordPlanDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordPlanDetailResponse) SetBody(v *QueryRecordPlanDetailResponseBody) *QueryRecordPlanDetailResponse {
	s.Body = v
	return s
}

type QueryRecordPlanDeviceByDeviceRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	StreamType    *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s QueryRecordPlanDeviceByDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDeviceByDeviceRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDeviceByDeviceRequest) SetDeviceName(v string) *QueryRecordPlanDeviceByDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceRequest) SetIotId(v string) *QueryRecordPlanDeviceByDeviceRequest {
	s.IotId = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceRequest) SetIotInstanceId(v string) *QueryRecordPlanDeviceByDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceRequest) SetProductKey(v string) *QueryRecordPlanDeviceByDeviceRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceRequest) SetStreamType(v int32) *QueryRecordPlanDeviceByDeviceRequest {
	s.StreamType = &v
	return s
}

type QueryRecordPlanDeviceByDeviceResponseBody struct {
	Code         *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryRecordPlanDeviceByDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRecordPlanDeviceByDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDeviceByDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDeviceByDeviceResponseBody) SetCode(v string) *QueryRecordPlanDeviceByDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceResponseBody) SetData(v *QueryRecordPlanDeviceByDeviceResponseBodyData) *QueryRecordPlanDeviceByDeviceResponseBody {
	s.Data = v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceResponseBody) SetErrorMessage(v string) *QueryRecordPlanDeviceByDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceResponseBody) SetRequestId(v string) *QueryRecordPlanDeviceByDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceResponseBody) SetSuccess(v bool) *QueryRecordPlanDeviceByDeviceResponseBody {
	s.Success = &v
	return s
}

type QueryRecordPlanDeviceByDeviceResponseBodyData struct {
	Name         *string                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	PlanId       *string                                                    `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	TemplateId   *string                                                    `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateInfo *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo `json:"TemplateInfo,omitempty" xml:"TemplateInfo,omitempty" type:"Struct"`
}

func (s QueryRecordPlanDeviceByDeviceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDeviceByDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDeviceByDeviceResponseBodyData) SetName(v string) *QueryRecordPlanDeviceByDeviceResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceResponseBodyData) SetPlanId(v string) *QueryRecordPlanDeviceByDeviceResponseBodyData {
	s.PlanId = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceResponseBodyData) SetTemplateId(v string) *QueryRecordPlanDeviceByDeviceResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceResponseBodyData) SetTemplateInfo(v *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo) *QueryRecordPlanDeviceByDeviceResponseBodyData {
	s.TemplateInfo = v
	return s
}

type QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo struct {
	AllDay          *int32                                                                      `json:"AllDay,omitempty" xml:"AllDay,omitempty"`
	Default         *int32                                                                      `json:"Default,omitempty" xml:"Default,omitempty"`
	Name            *string                                                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateId      *string                                                                     `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TimeSectionList []*QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList `json:"TimeSectionList,omitempty" xml:"TimeSectionList,omitempty" type:"Repeated"`
}

func (s QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo) SetAllDay(v int32) *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo {
	s.AllDay = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo) SetDefault(v int32) *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo {
	s.Default = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo) SetName(v string) *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo {
	s.Name = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo) SetTemplateId(v string) *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo {
	s.TemplateId = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo) SetTimeSectionList(v []*QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList) *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo {
	s.TimeSectionList = v
	return s
}

type QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList struct {
	Begin     *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	DayOfWeek *int32 `json:"DayOfWeek,omitempty" xml:"DayOfWeek,omitempty"`
	End       *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList) SetBegin(v int32) *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList {
	s.Begin = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList) SetDayOfWeek(v int32) *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList {
	s.DayOfWeek = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList) SetEnd(v int32) *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList {
	s.End = &v
	return s
}

type QueryRecordPlanDeviceByDeviceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecordPlanDeviceByDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecordPlanDeviceByDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDeviceByDeviceResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDeviceByDeviceResponse) SetHeaders(v map[string]*string) *QueryRecordPlanDeviceByDeviceResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceResponse) SetStatusCode(v int32) *QueryRecordPlanDeviceByDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceResponse) SetBody(v *QueryRecordPlanDeviceByDeviceResponseBody) *QueryRecordPlanDeviceByDeviceResponse {
	s.Body = v
	return s
}

type QueryRecordPlanDeviceByPlanRequest struct {
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PlanId      *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s QueryRecordPlanDeviceByPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDeviceByPlanRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDeviceByPlanRequest) SetCurrentPage(v int32) *QueryRecordPlanDeviceByPlanRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryRecordPlanDeviceByPlanRequest) SetPageSize(v int32) *QueryRecordPlanDeviceByPlanRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRecordPlanDeviceByPlanRequest) SetPlanId(v string) *QueryRecordPlanDeviceByPlanRequest {
	s.PlanId = &v
	return s
}

type QueryRecordPlanDeviceByPlanResponseBody struct {
	Code         *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryRecordPlanDeviceByPlanResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRecordPlanDeviceByPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDeviceByPlanResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDeviceByPlanResponseBody) SetCode(v string) *QueryRecordPlanDeviceByPlanResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRecordPlanDeviceByPlanResponseBody) SetData(v *QueryRecordPlanDeviceByPlanResponseBodyData) *QueryRecordPlanDeviceByPlanResponseBody {
	s.Data = v
	return s
}

func (s *QueryRecordPlanDeviceByPlanResponseBody) SetErrorMessage(v string) *QueryRecordPlanDeviceByPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryRecordPlanDeviceByPlanResponseBody) SetRequestId(v string) *QueryRecordPlanDeviceByPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRecordPlanDeviceByPlanResponseBody) SetSuccess(v bool) *QueryRecordPlanDeviceByPlanResponseBody {
	s.Success = &v
	return s
}

type QueryRecordPlanDeviceByPlanResponseBodyData struct {
	List      []*QueryRecordPlanDeviceByPlanResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Page      *int32                                             `json:"Page,omitempty" xml:"Page,omitempty"`
	PageCount *int32                                             `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageSize  *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int32                                             `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryRecordPlanDeviceByPlanResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDeviceByPlanResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDeviceByPlanResponseBodyData) SetList(v []*QueryRecordPlanDeviceByPlanResponseBodyDataList) *QueryRecordPlanDeviceByPlanResponseBodyData {
	s.List = v
	return s
}

func (s *QueryRecordPlanDeviceByPlanResponseBodyData) SetPage(v int32) *QueryRecordPlanDeviceByPlanResponseBodyData {
	s.Page = &v
	return s
}

func (s *QueryRecordPlanDeviceByPlanResponseBodyData) SetPageCount(v int32) *QueryRecordPlanDeviceByPlanResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryRecordPlanDeviceByPlanResponseBodyData) SetPageSize(v int32) *QueryRecordPlanDeviceByPlanResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryRecordPlanDeviceByPlanResponseBodyData) SetTotal(v int32) *QueryRecordPlanDeviceByPlanResponseBodyData {
	s.Total = &v
	return s
}

type QueryRecordPlanDeviceByPlanResponseBodyDataList struct {
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	StreamType *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s QueryRecordPlanDeviceByPlanResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDeviceByPlanResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDeviceByPlanResponseBodyDataList) SetIotId(v string) *QueryRecordPlanDeviceByPlanResponseBodyDataList {
	s.IotId = &v
	return s
}

func (s *QueryRecordPlanDeviceByPlanResponseBodyDataList) SetStreamType(v int32) *QueryRecordPlanDeviceByPlanResponseBodyDataList {
	s.StreamType = &v
	return s
}

type QueryRecordPlanDeviceByPlanResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecordPlanDeviceByPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecordPlanDeviceByPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDeviceByPlanResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDeviceByPlanResponse) SetHeaders(v map[string]*string) *QueryRecordPlanDeviceByPlanResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordPlanDeviceByPlanResponse) SetStatusCode(v int32) *QueryRecordPlanDeviceByPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordPlanDeviceByPlanResponse) SetBody(v *QueryRecordPlanDeviceByPlanResponseBody) *QueryRecordPlanDeviceByPlanResponse {
	s.Body = v
	return s
}

type QueryRecordPlansRequest struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryRecordPlansRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlansRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordPlansRequest) SetCurrentPage(v int32) *QueryRecordPlansRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryRecordPlansRequest) SetPageSize(v int32) *QueryRecordPlansRequest {
	s.PageSize = &v
	return s
}

type QueryRecordPlansResponseBody struct {
	Code         *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryRecordPlansResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRecordPlansResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlansResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecordPlansResponseBody) SetCode(v string) *QueryRecordPlansResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRecordPlansResponseBody) SetData(v *QueryRecordPlansResponseBodyData) *QueryRecordPlansResponseBody {
	s.Data = v
	return s
}

func (s *QueryRecordPlansResponseBody) SetErrorMessage(v string) *QueryRecordPlansResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryRecordPlansResponseBody) SetRequestId(v string) *QueryRecordPlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRecordPlansResponseBody) SetSuccess(v bool) *QueryRecordPlansResponseBody {
	s.Success = &v
	return s
}

type QueryRecordPlansResponseBodyData struct {
	List      []*QueryRecordPlansResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Page      *int32                                  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageCount *int32                                  `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageSize  *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int32                                  `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryRecordPlansResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlansResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRecordPlansResponseBodyData) SetList(v []*QueryRecordPlansResponseBodyDataList) *QueryRecordPlansResponseBodyData {
	s.List = v
	return s
}

func (s *QueryRecordPlansResponseBodyData) SetPage(v int32) *QueryRecordPlansResponseBodyData {
	s.Page = &v
	return s
}

func (s *QueryRecordPlansResponseBodyData) SetPageCount(v int32) *QueryRecordPlansResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryRecordPlansResponseBodyData) SetPageSize(v int32) *QueryRecordPlansResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryRecordPlansResponseBodyData) SetTotal(v int32) *QueryRecordPlansResponseBodyData {
	s.Total = &v
	return s
}

type QueryRecordPlansResponseBodyDataList struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PlanId     *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s QueryRecordPlansResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlansResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryRecordPlansResponseBodyDataList) SetName(v string) *QueryRecordPlansResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *QueryRecordPlansResponseBodyDataList) SetPlanId(v string) *QueryRecordPlansResponseBodyDataList {
	s.PlanId = &v
	return s
}

func (s *QueryRecordPlansResponseBodyDataList) SetTemplateId(v string) *QueryRecordPlansResponseBodyDataList {
	s.TemplateId = &v
	return s
}

type QueryRecordPlansResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecordPlansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecordPlansResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlansResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordPlansResponse) SetHeaders(v map[string]*string) *QueryRecordPlansResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordPlansResponse) SetStatusCode(v int32) *QueryRecordPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordPlansResponse) SetBody(v *QueryRecordPlansResponseBody) *QueryRecordPlansResponse {
	s.Body = v
	return s
}

type QueryRecordUrlRequest struct {
	DeviceName       *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	FileName         *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	IotId            *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId    *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey       *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	SeekTime         *int32  `json:"SeekTime,omitempty" xml:"SeekTime,omitempty"`
	UrlValidDuration *int32  `json:"UrlValidDuration,omitempty" xml:"UrlValidDuration,omitempty"`
}

func (s QueryRecordUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordUrlRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordUrlRequest) SetDeviceName(v string) *QueryRecordUrlRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryRecordUrlRequest) SetFileName(v string) *QueryRecordUrlRequest {
	s.FileName = &v
	return s
}

func (s *QueryRecordUrlRequest) SetIotId(v string) *QueryRecordUrlRequest {
	s.IotId = &v
	return s
}

func (s *QueryRecordUrlRequest) SetIotInstanceId(v string) *QueryRecordUrlRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryRecordUrlRequest) SetProductKey(v string) *QueryRecordUrlRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryRecordUrlRequest) SetSeekTime(v int32) *QueryRecordUrlRequest {
	s.SeekTime = &v
	return s
}

func (s *QueryRecordUrlRequest) SetUrlValidDuration(v int32) *QueryRecordUrlRequest {
	s.UrlValidDuration = &v
	return s
}

type QueryRecordUrlResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRecordUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordUrlResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecordUrlResponseBody) SetCode(v string) *QueryRecordUrlResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRecordUrlResponseBody) SetData(v string) *QueryRecordUrlResponseBody {
	s.Data = &v
	return s
}

func (s *QueryRecordUrlResponseBody) SetErrorMessage(v string) *QueryRecordUrlResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryRecordUrlResponseBody) SetRequestId(v string) *QueryRecordUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRecordUrlResponseBody) SetSuccess(v bool) *QueryRecordUrlResponseBody {
	s.Success = &v
	return s
}

type QueryRecordUrlResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecordUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecordUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordUrlResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordUrlResponse) SetHeaders(v map[string]*string) *QueryRecordUrlResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordUrlResponse) SetStatusCode(v int32) *QueryRecordUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordUrlResponse) SetBody(v *QueryRecordUrlResponseBody) *QueryRecordUrlResponse {
	s.Body = v
	return s
}

type QueryRecordUrlByTimeRequest struct {
	BeginTime        *int32  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	DeviceName       *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	EndTime          *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IotId            *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId    *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey       *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	StreamType       *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	UrlValidDuration *int32  `json:"UrlValidDuration,omitempty" xml:"UrlValidDuration,omitempty"`
}

func (s QueryRecordUrlByTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordUrlByTimeRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordUrlByTimeRequest) SetBeginTime(v int32) *QueryRecordUrlByTimeRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryRecordUrlByTimeRequest) SetDeviceName(v string) *QueryRecordUrlByTimeRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryRecordUrlByTimeRequest) SetEndTime(v int32) *QueryRecordUrlByTimeRequest {
	s.EndTime = &v
	return s
}

func (s *QueryRecordUrlByTimeRequest) SetIotId(v string) *QueryRecordUrlByTimeRequest {
	s.IotId = &v
	return s
}

func (s *QueryRecordUrlByTimeRequest) SetIotInstanceId(v string) *QueryRecordUrlByTimeRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryRecordUrlByTimeRequest) SetProductKey(v string) *QueryRecordUrlByTimeRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryRecordUrlByTimeRequest) SetStreamType(v int32) *QueryRecordUrlByTimeRequest {
	s.StreamType = &v
	return s
}

func (s *QueryRecordUrlByTimeRequest) SetUrlValidDuration(v int32) *QueryRecordUrlByTimeRequest {
	s.UrlValidDuration = &v
	return s
}

type QueryRecordUrlByTimeResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRecordUrlByTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordUrlByTimeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecordUrlByTimeResponseBody) SetCode(v string) *QueryRecordUrlByTimeResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRecordUrlByTimeResponseBody) SetData(v string) *QueryRecordUrlByTimeResponseBody {
	s.Data = &v
	return s
}

func (s *QueryRecordUrlByTimeResponseBody) SetErrorMessage(v string) *QueryRecordUrlByTimeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryRecordUrlByTimeResponseBody) SetRequestId(v string) *QueryRecordUrlByTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRecordUrlByTimeResponseBody) SetSuccess(v bool) *QueryRecordUrlByTimeResponseBody {
	s.Success = &v
	return s
}

type QueryRecordUrlByTimeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecordUrlByTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecordUrlByTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordUrlByTimeResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordUrlByTimeResponse) SetHeaders(v map[string]*string) *QueryRecordUrlByTimeResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordUrlByTimeResponse) SetStatusCode(v int32) *QueryRecordUrlByTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordUrlByTimeResponse) SetBody(v *QueryRecordUrlByTimeResponseBody) *QueryRecordUrlByTimeResponse {
	s.Body = v
	return s
}

type QueryRtmpKeyRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryRtmpKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRtmpKeyRequest) GoString() string {
	return s.String()
}

func (s *QueryRtmpKeyRequest) SetDeviceName(v string) *QueryRtmpKeyRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryRtmpKeyRequest) SetIotId(v string) *QueryRtmpKeyRequest {
	s.IotId = &v
	return s
}

func (s *QueryRtmpKeyRequest) SetIotInstanceId(v string) *QueryRtmpKeyRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryRtmpKeyRequest) SetProductKey(v string) *QueryRtmpKeyRequest {
	s.ProductKey = &v
	return s
}

type QueryRtmpKeyResponseBody struct {
	Code         *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryRtmpKeyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRtmpKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRtmpKeyResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRtmpKeyResponseBody) SetCode(v string) *QueryRtmpKeyResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRtmpKeyResponseBody) SetData(v *QueryRtmpKeyResponseBodyData) *QueryRtmpKeyResponseBody {
	s.Data = v
	return s
}

func (s *QueryRtmpKeyResponseBody) SetErrorMessage(v string) *QueryRtmpKeyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryRtmpKeyResponseBody) SetRequestId(v string) *QueryRtmpKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRtmpKeyResponseBody) SetSuccess(v bool) *QueryRtmpKeyResponseBody {
	s.Success = &v
	return s
}

type QueryRtmpKeyResponseBodyData struct {
	PullAuthKey       *string `json:"PullAuthKey,omitempty" xml:"PullAuthKey,omitempty"`
	PullKeyExpireTime *int32  `json:"PullKeyExpireTime,omitempty" xml:"PullKeyExpireTime,omitempty"`
	PushAuthKey       *string `json:"PushAuthKey,omitempty" xml:"PushAuthKey,omitempty"`
	PushKeyExpireTime *int32  `json:"PushKeyExpireTime,omitempty" xml:"PushKeyExpireTime,omitempty"`
	StreamName        *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s QueryRtmpKeyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryRtmpKeyResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRtmpKeyResponseBodyData) SetPullAuthKey(v string) *QueryRtmpKeyResponseBodyData {
	s.PullAuthKey = &v
	return s
}

func (s *QueryRtmpKeyResponseBodyData) SetPullKeyExpireTime(v int32) *QueryRtmpKeyResponseBodyData {
	s.PullKeyExpireTime = &v
	return s
}

func (s *QueryRtmpKeyResponseBodyData) SetPushAuthKey(v string) *QueryRtmpKeyResponseBodyData {
	s.PushAuthKey = &v
	return s
}

func (s *QueryRtmpKeyResponseBodyData) SetPushKeyExpireTime(v int32) *QueryRtmpKeyResponseBodyData {
	s.PushKeyExpireTime = &v
	return s
}

func (s *QueryRtmpKeyResponseBodyData) SetStreamName(v string) *QueryRtmpKeyResponseBodyData {
	s.StreamName = &v
	return s
}

type QueryRtmpKeyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRtmpKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRtmpKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRtmpKeyResponse) GoString() string {
	return s.String()
}

func (s *QueryRtmpKeyResponse) SetHeaders(v map[string]*string) *QueryRtmpKeyResponse {
	s.Headers = v
	return s
}

func (s *QueryRtmpKeyResponse) SetStatusCode(v int32) *QueryRtmpKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRtmpKeyResponse) SetBody(v *QueryRtmpKeyResponseBody) *QueryRtmpKeyResponse {
	s.Body = v
	return s
}

type QueryStreamPushJobRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	JobId         *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryStreamPushJobRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryStreamPushJobRequest) GoString() string {
	return s.String()
}

func (s *QueryStreamPushJobRequest) SetDeviceName(v string) *QueryStreamPushJobRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryStreamPushJobRequest) SetIotId(v string) *QueryStreamPushJobRequest {
	s.IotId = &v
	return s
}

func (s *QueryStreamPushJobRequest) SetIotInstanceId(v string) *QueryStreamPushJobRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryStreamPushJobRequest) SetJobId(v string) *QueryStreamPushJobRequest {
	s.JobId = &v
	return s
}

func (s *QueryStreamPushJobRequest) SetProductKey(v string) *QueryStreamPushJobRequest {
	s.ProductKey = &v
	return s
}

type QueryStreamPushJobResponseBody struct {
	Code         *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryStreamPushJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryStreamPushJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryStreamPushJobResponseBody) GoString() string {
	return s.String()
}

func (s *QueryStreamPushJobResponseBody) SetCode(v string) *QueryStreamPushJobResponseBody {
	s.Code = &v
	return s
}

func (s *QueryStreamPushJobResponseBody) SetData(v *QueryStreamPushJobResponseBodyData) *QueryStreamPushJobResponseBody {
	s.Data = v
	return s
}

func (s *QueryStreamPushJobResponseBody) SetErrorMessage(v string) *QueryStreamPushJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryStreamPushJobResponseBody) SetRequestId(v string) *QueryStreamPushJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryStreamPushJobResponseBody) SetSuccess(v bool) *QueryStreamPushJobResponseBody {
	s.Success = &v
	return s
}

type QueryStreamPushJobResponseBodyData struct {
	CreateTime *int32  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	JobType    *int32  `json:"JobType,omitempty" xml:"JobType,omitempty"`
	PushStatus *int32  `json:"PushStatus,omitempty" xml:"PushStatus,omitempty"`
	PushUrl    *string `json:"PushUrl,omitempty" xml:"PushUrl,omitempty"`
	StreamType *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s QueryStreamPushJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryStreamPushJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryStreamPushJobResponseBodyData) SetCreateTime(v int32) *QueryStreamPushJobResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *QueryStreamPushJobResponseBodyData) SetJobType(v int32) *QueryStreamPushJobResponseBodyData {
	s.JobType = &v
	return s
}

func (s *QueryStreamPushJobResponseBodyData) SetPushStatus(v int32) *QueryStreamPushJobResponseBodyData {
	s.PushStatus = &v
	return s
}

func (s *QueryStreamPushJobResponseBodyData) SetPushUrl(v string) *QueryStreamPushJobResponseBodyData {
	s.PushUrl = &v
	return s
}

func (s *QueryStreamPushJobResponseBodyData) SetStreamType(v int32) *QueryStreamPushJobResponseBodyData {
	s.StreamType = &v
	return s
}

type QueryStreamPushJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryStreamPushJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryStreamPushJobResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryStreamPushJobResponse) GoString() string {
	return s.String()
}

func (s *QueryStreamPushJobResponse) SetHeaders(v map[string]*string) *QueryStreamPushJobResponse {
	s.Headers = v
	return s
}

func (s *QueryStreamPushJobResponse) SetStatusCode(v int32) *QueryStreamPushJobResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryStreamPushJobResponse) SetBody(v *QueryStreamPushJobResponseBody) *QueryStreamPushJobResponse {
	s.Body = v
	return s
}

type QueryStreamPushJobListRequest struct {
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	JobType       *int32  `json:"JobType,omitempty" xml:"JobType,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryStreamPushJobListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryStreamPushJobListRequest) GoString() string {
	return s.String()
}

func (s *QueryStreamPushJobListRequest) SetCurrentPage(v int32) *QueryStreamPushJobListRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryStreamPushJobListRequest) SetDeviceName(v string) *QueryStreamPushJobListRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryStreamPushJobListRequest) SetIotId(v string) *QueryStreamPushJobListRequest {
	s.IotId = &v
	return s
}

func (s *QueryStreamPushJobListRequest) SetIotInstanceId(v string) *QueryStreamPushJobListRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryStreamPushJobListRequest) SetJobType(v int32) *QueryStreamPushJobListRequest {
	s.JobType = &v
	return s
}

func (s *QueryStreamPushJobListRequest) SetPageSize(v int32) *QueryStreamPushJobListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryStreamPushJobListRequest) SetProductKey(v string) *QueryStreamPushJobListRequest {
	s.ProductKey = &v
	return s
}

type QueryStreamPushJobListResponseBody struct {
	Code         *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryStreamPushJobListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryStreamPushJobListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryStreamPushJobListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryStreamPushJobListResponseBody) SetCode(v string) *QueryStreamPushJobListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryStreamPushJobListResponseBody) SetData(v *QueryStreamPushJobListResponseBodyData) *QueryStreamPushJobListResponseBody {
	s.Data = v
	return s
}

func (s *QueryStreamPushJobListResponseBody) SetErrorMessage(v string) *QueryStreamPushJobListResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryStreamPushJobListResponseBody) SetRequestId(v string) *QueryStreamPushJobListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryStreamPushJobListResponseBody) SetSuccess(v bool) *QueryStreamPushJobListResponseBody {
	s.Success = &v
	return s
}

type QueryStreamPushJobListResponseBodyData struct {
	JobList []*QueryStreamPushJobListResponseBodyDataJobList `json:"JobList,omitempty" xml:"JobList,omitempty" type:"Repeated"`
	Total   *int32                                           `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryStreamPushJobListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryStreamPushJobListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryStreamPushJobListResponseBodyData) SetJobList(v []*QueryStreamPushJobListResponseBodyDataJobList) *QueryStreamPushJobListResponseBodyData {
	s.JobList = v
	return s
}

func (s *QueryStreamPushJobListResponseBodyData) SetTotal(v int32) *QueryStreamPushJobListResponseBodyData {
	s.Total = &v
	return s
}

type QueryStreamPushJobListResponseBodyDataJobList struct {
	CreateTime *int32  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobType    *int32  `json:"JobType,omitempty" xml:"JobType,omitempty"`
	PushStatus *int32  `json:"PushStatus,omitempty" xml:"PushStatus,omitempty"`
	PushUrl    *string `json:"PushUrl,omitempty" xml:"PushUrl,omitempty"`
	StreamType *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s QueryStreamPushJobListResponseBodyDataJobList) String() string {
	return tea.Prettify(s)
}

func (s QueryStreamPushJobListResponseBodyDataJobList) GoString() string {
	return s.String()
}

func (s *QueryStreamPushJobListResponseBodyDataJobList) SetCreateTime(v int32) *QueryStreamPushJobListResponseBodyDataJobList {
	s.CreateTime = &v
	return s
}

func (s *QueryStreamPushJobListResponseBodyDataJobList) SetJobId(v string) *QueryStreamPushJobListResponseBodyDataJobList {
	s.JobId = &v
	return s
}

func (s *QueryStreamPushJobListResponseBodyDataJobList) SetJobType(v int32) *QueryStreamPushJobListResponseBodyDataJobList {
	s.JobType = &v
	return s
}

func (s *QueryStreamPushJobListResponseBodyDataJobList) SetPushStatus(v int32) *QueryStreamPushJobListResponseBodyDataJobList {
	s.PushStatus = &v
	return s
}

func (s *QueryStreamPushJobListResponseBodyDataJobList) SetPushUrl(v string) *QueryStreamPushJobListResponseBodyDataJobList {
	s.PushUrl = &v
	return s
}

func (s *QueryStreamPushJobListResponseBodyDataJobList) SetStreamType(v int32) *QueryStreamPushJobListResponseBodyDataJobList {
	s.StreamType = &v
	return s
}

type QueryStreamPushJobListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryStreamPushJobListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryStreamPushJobListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryStreamPushJobListResponse) GoString() string {
	return s.String()
}

func (s *QueryStreamPushJobListResponse) SetHeaders(v map[string]*string) *QueryStreamPushJobListResponse {
	s.Headers = v
	return s
}

func (s *QueryStreamPushJobListResponse) SetStatusCode(v int32) *QueryStreamPushJobListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryStreamPushJobListResponse) SetBody(v *QueryStreamPushJobListResponseBody) *QueryStreamPushJobListResponse {
	s.Body = v
	return s
}

type QueryStreamSnapshotJobRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryStreamSnapshotJobRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryStreamSnapshotJobRequest) GoString() string {
	return s.String()
}

func (s *QueryStreamSnapshotJobRequest) SetDeviceName(v string) *QueryStreamSnapshotJobRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryStreamSnapshotJobRequest) SetIotId(v string) *QueryStreamSnapshotJobRequest {
	s.IotId = &v
	return s
}

func (s *QueryStreamSnapshotJobRequest) SetIotInstanceId(v string) *QueryStreamSnapshotJobRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryStreamSnapshotJobRequest) SetProductKey(v string) *QueryStreamSnapshotJobRequest {
	s.ProductKey = &v
	return s
}

type QueryStreamSnapshotJobResponseBody struct {
	Code         *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryStreamSnapshotJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryStreamSnapshotJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryStreamSnapshotJobResponseBody) GoString() string {
	return s.String()
}

func (s *QueryStreamSnapshotJobResponseBody) SetCode(v string) *QueryStreamSnapshotJobResponseBody {
	s.Code = &v
	return s
}

func (s *QueryStreamSnapshotJobResponseBody) SetData(v *QueryStreamSnapshotJobResponseBodyData) *QueryStreamSnapshotJobResponseBody {
	s.Data = v
	return s
}

func (s *QueryStreamSnapshotJobResponseBody) SetErrorMessage(v string) *QueryStreamSnapshotJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryStreamSnapshotJobResponseBody) SetRequestId(v string) *QueryStreamSnapshotJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryStreamSnapshotJobResponseBody) SetSuccess(v bool) *QueryStreamSnapshotJobResponseBody {
	s.Success = &v
	return s
}

type QueryStreamSnapshotJobResponseBodyData struct {
	JobList []*QueryStreamSnapshotJobResponseBodyDataJobList `json:"JobList,omitempty" xml:"JobList,omitempty" type:"Repeated"`
}

func (s QueryStreamSnapshotJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryStreamSnapshotJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryStreamSnapshotJobResponseBodyData) SetJobList(v []*QueryStreamSnapshotJobResponseBodyDataJobList) *QueryStreamSnapshotJobResponseBodyData {
	s.JobList = v
	return s
}

type QueryStreamSnapshotJobResponseBodyDataJobList struct {
	SnapshotInterval *int32 `json:"SnapshotInterval,omitempty" xml:"SnapshotInterval,omitempty"`
	StreamType       *int32 `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s QueryStreamSnapshotJobResponseBodyDataJobList) String() string {
	return tea.Prettify(s)
}

func (s QueryStreamSnapshotJobResponseBodyDataJobList) GoString() string {
	return s.String()
}

func (s *QueryStreamSnapshotJobResponseBodyDataJobList) SetSnapshotInterval(v int32) *QueryStreamSnapshotJobResponseBodyDataJobList {
	s.SnapshotInterval = &v
	return s
}

func (s *QueryStreamSnapshotJobResponseBodyDataJobList) SetStreamType(v int32) *QueryStreamSnapshotJobResponseBodyDataJobList {
	s.StreamType = &v
	return s
}

type QueryStreamSnapshotJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryStreamSnapshotJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryStreamSnapshotJobResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryStreamSnapshotJobResponse) GoString() string {
	return s.String()
}

func (s *QueryStreamSnapshotJobResponse) SetHeaders(v map[string]*string) *QueryStreamSnapshotJobResponse {
	s.Headers = v
	return s
}

func (s *QueryStreamSnapshotJobResponse) SetStatusCode(v int32) *QueryStreamSnapshotJobResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryStreamSnapshotJobResponse) SetBody(v *QueryStreamSnapshotJobResponseBody) *QueryStreamSnapshotJobResponse {
	s.Body = v
	return s
}

type QueryTimeTemplateRequest struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryTimeTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTimeTemplateRequest) GoString() string {
	return s.String()
}

func (s *QueryTimeTemplateRequest) SetCurrentPage(v int32) *QueryTimeTemplateRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryTimeTemplateRequest) SetPageSize(v int32) *QueryTimeTemplateRequest {
	s.PageSize = &v
	return s
}

type QueryTimeTemplateResponseBody struct {
	Code         *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryTimeTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTimeTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTimeTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTimeTemplateResponseBody) SetCode(v string) *QueryTimeTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTimeTemplateResponseBody) SetData(v *QueryTimeTemplateResponseBodyData) *QueryTimeTemplateResponseBody {
	s.Data = v
	return s
}

func (s *QueryTimeTemplateResponseBody) SetErrorMessage(v string) *QueryTimeTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryTimeTemplateResponseBody) SetRequestId(v string) *QueryTimeTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTimeTemplateResponseBody) SetSuccess(v bool) *QueryTimeTemplateResponseBody {
	s.Success = &v
	return s
}

type QueryTimeTemplateResponseBodyData struct {
	List      []*QueryTimeTemplateResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Page      *int32                                   `json:"Page,omitempty" xml:"Page,omitempty"`
	PageCount *int32                                   `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageSize  *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int32                                   `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryTimeTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTimeTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTimeTemplateResponseBodyData) SetList(v []*QueryTimeTemplateResponseBodyDataList) *QueryTimeTemplateResponseBodyData {
	s.List = v
	return s
}

func (s *QueryTimeTemplateResponseBodyData) SetPage(v int32) *QueryTimeTemplateResponseBodyData {
	s.Page = &v
	return s
}

func (s *QueryTimeTemplateResponseBodyData) SetPageCount(v int32) *QueryTimeTemplateResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryTimeTemplateResponseBodyData) SetPageSize(v int32) *QueryTimeTemplateResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryTimeTemplateResponseBodyData) SetTotal(v int32) *QueryTimeTemplateResponseBodyData {
	s.Total = &v
	return s
}

type QueryTimeTemplateResponseBodyDataList struct {
	AllDay          *int32                                                  `json:"AllDay,omitempty" xml:"AllDay,omitempty"`
	Default         *int32                                                  `json:"Default,omitempty" xml:"Default,omitempty"`
	Name            *string                                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateId      *string                                                 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TimeSectionList []*QueryTimeTemplateResponseBodyDataListTimeSectionList `json:"TimeSectionList,omitempty" xml:"TimeSectionList,omitempty" type:"Repeated"`
}

func (s QueryTimeTemplateResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryTimeTemplateResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryTimeTemplateResponseBodyDataList) SetAllDay(v int32) *QueryTimeTemplateResponseBodyDataList {
	s.AllDay = &v
	return s
}

func (s *QueryTimeTemplateResponseBodyDataList) SetDefault(v int32) *QueryTimeTemplateResponseBodyDataList {
	s.Default = &v
	return s
}

func (s *QueryTimeTemplateResponseBodyDataList) SetName(v string) *QueryTimeTemplateResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *QueryTimeTemplateResponseBodyDataList) SetTemplateId(v string) *QueryTimeTemplateResponseBodyDataList {
	s.TemplateId = &v
	return s
}

func (s *QueryTimeTemplateResponseBodyDataList) SetTimeSectionList(v []*QueryTimeTemplateResponseBodyDataListTimeSectionList) *QueryTimeTemplateResponseBodyDataList {
	s.TimeSectionList = v
	return s
}

type QueryTimeTemplateResponseBodyDataListTimeSectionList struct {
	Begin     *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	DayOfWeek *int32 `json:"DayOfWeek,omitempty" xml:"DayOfWeek,omitempty"`
	End       *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s QueryTimeTemplateResponseBodyDataListTimeSectionList) String() string {
	return tea.Prettify(s)
}

func (s QueryTimeTemplateResponseBodyDataListTimeSectionList) GoString() string {
	return s.String()
}

func (s *QueryTimeTemplateResponseBodyDataListTimeSectionList) SetBegin(v int32) *QueryTimeTemplateResponseBodyDataListTimeSectionList {
	s.Begin = &v
	return s
}

func (s *QueryTimeTemplateResponseBodyDataListTimeSectionList) SetDayOfWeek(v int32) *QueryTimeTemplateResponseBodyDataListTimeSectionList {
	s.DayOfWeek = &v
	return s
}

func (s *QueryTimeTemplateResponseBodyDataListTimeSectionList) SetEnd(v int32) *QueryTimeTemplateResponseBodyDataListTimeSectionList {
	s.End = &v
	return s
}

type QueryTimeTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTimeTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTimeTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTimeTemplateResponse) GoString() string {
	return s.String()
}

func (s *QueryTimeTemplateResponse) SetHeaders(v map[string]*string) *QueryTimeTemplateResponse {
	s.Headers = v
	return s
}

func (s *QueryTimeTemplateResponse) SetStatusCode(v int32) *QueryTimeTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTimeTemplateResponse) SetBody(v *QueryTimeTemplateResponseBody) *QueryTimeTemplateResponse {
	s.Body = v
	return s
}

type QueryTimeTemplateDetailRequest struct {
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s QueryTimeTemplateDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTimeTemplateDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryTimeTemplateDetailRequest) SetTemplateId(v string) *QueryTimeTemplateDetailRequest {
	s.TemplateId = &v
	return s
}

type QueryTimeTemplateDetailResponseBody struct {
	Code         *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryTimeTemplateDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTimeTemplateDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTimeTemplateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTimeTemplateDetailResponseBody) SetCode(v string) *QueryTimeTemplateDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTimeTemplateDetailResponseBody) SetData(v *QueryTimeTemplateDetailResponseBodyData) *QueryTimeTemplateDetailResponseBody {
	s.Data = v
	return s
}

func (s *QueryTimeTemplateDetailResponseBody) SetErrorMessage(v string) *QueryTimeTemplateDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryTimeTemplateDetailResponseBody) SetRequestId(v string) *QueryTimeTemplateDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTimeTemplateDetailResponseBody) SetSuccess(v bool) *QueryTimeTemplateDetailResponseBody {
	s.Success = &v
	return s
}

type QueryTimeTemplateDetailResponseBodyData struct {
	AllDay          *int32                                                    `json:"AllDay,omitempty" xml:"AllDay,omitempty"`
	Default         *int32                                                    `json:"Default,omitempty" xml:"Default,omitempty"`
	Name            *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateId      *string                                                   `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TimeSectionList []*QueryTimeTemplateDetailResponseBodyDataTimeSectionList `json:"TimeSectionList,omitempty" xml:"TimeSectionList,omitempty" type:"Repeated"`
}

func (s QueryTimeTemplateDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTimeTemplateDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTimeTemplateDetailResponseBodyData) SetAllDay(v int32) *QueryTimeTemplateDetailResponseBodyData {
	s.AllDay = &v
	return s
}

func (s *QueryTimeTemplateDetailResponseBodyData) SetDefault(v int32) *QueryTimeTemplateDetailResponseBodyData {
	s.Default = &v
	return s
}

func (s *QueryTimeTemplateDetailResponseBodyData) SetName(v string) *QueryTimeTemplateDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryTimeTemplateDetailResponseBodyData) SetTemplateId(v string) *QueryTimeTemplateDetailResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *QueryTimeTemplateDetailResponseBodyData) SetTimeSectionList(v []*QueryTimeTemplateDetailResponseBodyDataTimeSectionList) *QueryTimeTemplateDetailResponseBodyData {
	s.TimeSectionList = v
	return s
}

type QueryTimeTemplateDetailResponseBodyDataTimeSectionList struct {
	Begin     *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	DayOfWeek *int32 `json:"DayOfWeek,omitempty" xml:"DayOfWeek,omitempty"`
	End       *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s QueryTimeTemplateDetailResponseBodyDataTimeSectionList) String() string {
	return tea.Prettify(s)
}

func (s QueryTimeTemplateDetailResponseBodyDataTimeSectionList) GoString() string {
	return s.String()
}

func (s *QueryTimeTemplateDetailResponseBodyDataTimeSectionList) SetBegin(v int32) *QueryTimeTemplateDetailResponseBodyDataTimeSectionList {
	s.Begin = &v
	return s
}

func (s *QueryTimeTemplateDetailResponseBodyDataTimeSectionList) SetDayOfWeek(v int32) *QueryTimeTemplateDetailResponseBodyDataTimeSectionList {
	s.DayOfWeek = &v
	return s
}

func (s *QueryTimeTemplateDetailResponseBodyDataTimeSectionList) SetEnd(v int32) *QueryTimeTemplateDetailResponseBodyDataTimeSectionList {
	s.End = &v
	return s
}

type QueryTimeTemplateDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTimeTemplateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTimeTemplateDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTimeTemplateDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryTimeTemplateDetailResponse) SetHeaders(v map[string]*string) *QueryTimeTemplateDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryTimeTemplateDetailResponse) SetStatusCode(v int32) *QueryTimeTemplateDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTimeTemplateDetailResponse) SetBody(v *QueryTimeTemplateDetailResponseBody) *QueryTimeTemplateDetailResponse {
	s.Body = v
	return s
}

type QueryVisionDeviceInfoRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryVisionDeviceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryVisionDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryVisionDeviceInfoRequest) SetDeviceName(v string) *QueryVisionDeviceInfoRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryVisionDeviceInfoRequest) SetIotId(v string) *QueryVisionDeviceInfoRequest {
	s.IotId = &v
	return s
}

func (s *QueryVisionDeviceInfoRequest) SetIotInstanceId(v string) *QueryVisionDeviceInfoRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryVisionDeviceInfoRequest) SetProductKey(v string) *QueryVisionDeviceInfoRequest {
	s.ProductKey = &v
	return s
}

type QueryVisionDeviceInfoResponseBody struct {
	Code         *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryVisionDeviceInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryVisionDeviceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryVisionDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVisionDeviceInfoResponseBody) SetCode(v string) *QueryVisionDeviceInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBody) SetData(v *QueryVisionDeviceInfoResponseBodyData) *QueryVisionDeviceInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryVisionDeviceInfoResponseBody) SetErrorMessage(v string) *QueryVisionDeviceInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBody) SetRequestId(v string) *QueryVisionDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBody) SetSuccess(v bool) *QueryVisionDeviceInfoResponseBody {
	s.Success = &v
	return s
}

type QueryVisionDeviceInfoResponseBodyData struct {
	Description    *string                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceType     *int32                                               `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	GbDeviceInfo   *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo   `json:"GbDeviceInfo,omitempty" xml:"GbDeviceInfo,omitempty" type:"Struct"`
	RtmpDeviceInfo *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo `json:"RtmpDeviceInfo,omitempty" xml:"RtmpDeviceInfo,omitempty" type:"Struct"`
}

func (s QueryVisionDeviceInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryVisionDeviceInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryVisionDeviceInfoResponseBodyData) SetDescription(v string) *QueryVisionDeviceInfoResponseBodyData {
	s.Description = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyData) SetDeviceType(v int32) *QueryVisionDeviceInfoResponseBodyData {
	s.DeviceType = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyData) SetGbDeviceInfo(v *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo) *QueryVisionDeviceInfoResponseBodyData {
	s.GbDeviceInfo = v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyData) SetRtmpDeviceInfo(v *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) *QueryVisionDeviceInfoResponseBodyData {
	s.RtmpDeviceInfo = v
	return s
}

type QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo struct {
	DeviceProtocol *int32  `json:"DeviceProtocol,omitempty" xml:"DeviceProtocol,omitempty"`
	GbId           *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	NetProtocol    *int32  `json:"NetProtocol,omitempty" xml:"NetProtocol,omitempty"`
	NickName       *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	Password       *string `json:"Password,omitempty" xml:"Password,omitempty"`
	SubProductKey  *string `json:"SubProductKey,omitempty" xml:"SubProductKey,omitempty"`
}

func (s QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo) GoString() string {
	return s.String()
}

func (s *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo) SetDeviceProtocol(v int32) *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo {
	s.DeviceProtocol = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo) SetGbId(v string) *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo {
	s.GbId = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo) SetNetProtocol(v int32) *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo {
	s.NetProtocol = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo) SetNickName(v string) *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo {
	s.NickName = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo) SetPassword(v string) *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo {
	s.Password = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo) SetSubProductKey(v string) *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo {
	s.SubProductKey = &v
	return s
}

type QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo struct {
	PullAuthKey       *string `json:"PullAuthKey,omitempty" xml:"PullAuthKey,omitempty"`
	PullKeyExpireTime *int32  `json:"PullKeyExpireTime,omitempty" xml:"PullKeyExpireTime,omitempty"`
	PushAuthKey       *string `json:"PushAuthKey,omitempty" xml:"PushAuthKey,omitempty"`
	PushKeyExpireTime *int32  `json:"PushKeyExpireTime,omitempty" xml:"PushKeyExpireTime,omitempty"`
	PushUrlSample     *string `json:"PushUrlSample,omitempty" xml:"PushUrlSample,omitempty"`
	StreamName        *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	StreamStatus      *int32  `json:"StreamStatus,omitempty" xml:"StreamStatus,omitempty"`
}

func (s QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) GoString() string {
	return s.String()
}

func (s *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) SetPullAuthKey(v string) *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo {
	s.PullAuthKey = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) SetPullKeyExpireTime(v int32) *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo {
	s.PullKeyExpireTime = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) SetPushAuthKey(v string) *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo {
	s.PushAuthKey = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) SetPushKeyExpireTime(v int32) *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo {
	s.PushKeyExpireTime = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) SetPushUrlSample(v string) *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo {
	s.PushUrlSample = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) SetStreamName(v string) *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo {
	s.StreamName = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) SetStreamStatus(v int32) *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo {
	s.StreamStatus = &v
	return s
}

type QueryVisionDeviceInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryVisionDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryVisionDeviceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryVisionDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryVisionDeviceInfoResponse) SetHeaders(v map[string]*string) *QueryVisionDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryVisionDeviceInfoResponse) SetStatusCode(v int32) *QueryVisionDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryVisionDeviceInfoResponse) SetBody(v *QueryVisionDeviceInfoResponseBody) *QueryVisionDeviceInfoResponse {
	s.Body = v
	return s
}

type QueryVoiceIntercomRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	Scheme        *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
}

func (s QueryVoiceIntercomRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryVoiceIntercomRequest) GoString() string {
	return s.String()
}

func (s *QueryVoiceIntercomRequest) SetDeviceName(v string) *QueryVoiceIntercomRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryVoiceIntercomRequest) SetIotId(v string) *QueryVoiceIntercomRequest {
	s.IotId = &v
	return s
}

func (s *QueryVoiceIntercomRequest) SetIotInstanceId(v string) *QueryVoiceIntercomRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryVoiceIntercomRequest) SetProductKey(v string) *QueryVoiceIntercomRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryVoiceIntercomRequest) SetScheme(v string) *QueryVoiceIntercomRequest {
	s.Scheme = &v
	return s
}

type QueryVoiceIntercomResponseBody struct {
	Code         *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryVoiceIntercomResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryVoiceIntercomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryVoiceIntercomResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVoiceIntercomResponseBody) SetCode(v string) *QueryVoiceIntercomResponseBody {
	s.Code = &v
	return s
}

func (s *QueryVoiceIntercomResponseBody) SetData(v *QueryVoiceIntercomResponseBodyData) *QueryVoiceIntercomResponseBody {
	s.Data = v
	return s
}

func (s *QueryVoiceIntercomResponseBody) SetErrorMessage(v string) *QueryVoiceIntercomResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryVoiceIntercomResponseBody) SetRequestId(v string) *QueryVoiceIntercomResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryVoiceIntercomResponseBody) SetSuccess(v bool) *QueryVoiceIntercomResponseBody {
	s.Success = &v
	return s
}

type QueryVoiceIntercomResponseBodyData struct {
	CryptoKey *QueryVoiceIntercomResponseBodyDataCryptoKey `json:"CryptoKey,omitempty" xml:"CryptoKey,omitempty" type:"Struct"`
	Url       *string                                      `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s QueryVoiceIntercomResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryVoiceIntercomResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryVoiceIntercomResponseBodyData) SetCryptoKey(v *QueryVoiceIntercomResponseBodyDataCryptoKey) *QueryVoiceIntercomResponseBodyData {
	s.CryptoKey = v
	return s
}

func (s *QueryVoiceIntercomResponseBodyData) SetUrl(v string) *QueryVoiceIntercomResponseBodyData {
	s.Url = &v
	return s
}

type QueryVoiceIntercomResponseBodyDataCryptoKey struct {
	Iv  *string `json:"Iv,omitempty" xml:"Iv,omitempty"`
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s QueryVoiceIntercomResponseBodyDataCryptoKey) String() string {
	return tea.Prettify(s)
}

func (s QueryVoiceIntercomResponseBodyDataCryptoKey) GoString() string {
	return s.String()
}

func (s *QueryVoiceIntercomResponseBodyDataCryptoKey) SetIv(v string) *QueryVoiceIntercomResponseBodyDataCryptoKey {
	s.Iv = &v
	return s
}

func (s *QueryVoiceIntercomResponseBodyDataCryptoKey) SetKey(v string) *QueryVoiceIntercomResponseBodyDataCryptoKey {
	s.Key = &v
	return s
}

type QueryVoiceIntercomResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryVoiceIntercomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryVoiceIntercomResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryVoiceIntercomResponse) GoString() string {
	return s.String()
}

func (s *QueryVoiceIntercomResponse) SetHeaders(v map[string]*string) *QueryVoiceIntercomResponse {
	s.Headers = v
	return s
}

func (s *QueryVoiceIntercomResponse) SetStatusCode(v int32) *QueryVoiceIntercomResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryVoiceIntercomResponse) SetBody(v *QueryVoiceIntercomResponseBody) *QueryVoiceIntercomResponse {
	s.Body = v
	return s
}

type RefreshGbSubDeviceListRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s RefreshGbSubDeviceListRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshGbSubDeviceListRequest) GoString() string {
	return s.String()
}

func (s *RefreshGbSubDeviceListRequest) SetDeviceName(v string) *RefreshGbSubDeviceListRequest {
	s.DeviceName = &v
	return s
}

func (s *RefreshGbSubDeviceListRequest) SetIotId(v string) *RefreshGbSubDeviceListRequest {
	s.IotId = &v
	return s
}

func (s *RefreshGbSubDeviceListRequest) SetIotInstanceId(v string) *RefreshGbSubDeviceListRequest {
	s.IotInstanceId = &v
	return s
}

func (s *RefreshGbSubDeviceListRequest) SetProductKey(v string) *RefreshGbSubDeviceListRequest {
	s.ProductKey = &v
	return s
}

type RefreshGbSubDeviceListResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RefreshGbSubDeviceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshGbSubDeviceListResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshGbSubDeviceListResponseBody) SetCode(v string) *RefreshGbSubDeviceListResponseBody {
	s.Code = &v
	return s
}

func (s *RefreshGbSubDeviceListResponseBody) SetErrorMessage(v string) *RefreshGbSubDeviceListResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RefreshGbSubDeviceListResponseBody) SetRequestId(v string) *RefreshGbSubDeviceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshGbSubDeviceListResponseBody) SetSuccess(v bool) *RefreshGbSubDeviceListResponseBody {
	s.Success = &v
	return s
}

type RefreshGbSubDeviceListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshGbSubDeviceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshGbSubDeviceListResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshGbSubDeviceListResponse) GoString() string {
	return s.String()
}

func (s *RefreshGbSubDeviceListResponse) SetHeaders(v map[string]*string) *RefreshGbSubDeviceListResponse {
	s.Headers = v
	return s
}

func (s *RefreshGbSubDeviceListResponse) SetStatusCode(v int32) *RefreshGbSubDeviceListResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshGbSubDeviceListResponse) SetBody(v *RefreshGbSubDeviceListResponseBody) *RefreshGbSubDeviceListResponse {
	s.Body = v
	return s
}

type RemoveFaceDeviceFromDeviceGroupRequest struct {
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	IsolationId   *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s RemoveFaceDeviceFromDeviceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveFaceDeviceFromDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveFaceDeviceFromDeviceGroupRequest) SetDeviceGroupId(v string) *RemoveFaceDeviceFromDeviceGroupRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupRequest) SetDeviceName(v string) *RemoveFaceDeviceFromDeviceGroupRequest {
	s.DeviceName = &v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupRequest) SetIotInstanceId(v string) *RemoveFaceDeviceFromDeviceGroupRequest {
	s.IotInstanceId = &v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupRequest) SetIsolationId(v string) *RemoveFaceDeviceFromDeviceGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupRequest) SetProductKey(v string) *RemoveFaceDeviceFromDeviceGroupRequest {
	s.ProductKey = &v
	return s
}

type RemoveFaceDeviceFromDeviceGroupResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveFaceDeviceFromDeviceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveFaceDeviceFromDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveFaceDeviceFromDeviceGroupResponseBody) SetCode(v string) *RemoveFaceDeviceFromDeviceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupResponseBody) SetErrorMessage(v string) *RemoveFaceDeviceFromDeviceGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupResponseBody) SetRequestId(v string) *RemoveFaceDeviceFromDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupResponseBody) SetSuccess(v bool) *RemoveFaceDeviceFromDeviceGroupResponseBody {
	s.Success = &v
	return s
}

type RemoveFaceDeviceFromDeviceGroupResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveFaceDeviceFromDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveFaceDeviceFromDeviceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveFaceDeviceFromDeviceGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveFaceDeviceFromDeviceGroupResponse) SetHeaders(v map[string]*string) *RemoveFaceDeviceFromDeviceGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupResponse) SetStatusCode(v int32) *RemoveFaceDeviceFromDeviceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupResponse) SetBody(v *RemoveFaceDeviceFromDeviceGroupResponseBody) *RemoveFaceDeviceFromDeviceGroupResponse {
	s.Body = v
	return s
}

type RemoveFaceUserFromUserGroupRequest struct {
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RemoveFaceUserFromUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveFaceUserFromUserGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveFaceUserFromUserGroupRequest) SetIsolationId(v string) *RemoveFaceUserFromUserGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *RemoveFaceUserFromUserGroupRequest) SetUserGroupId(v string) *RemoveFaceUserFromUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *RemoveFaceUserFromUserGroupRequest) SetUserId(v string) *RemoveFaceUserFromUserGroupRequest {
	s.UserId = &v
	return s
}

type RemoveFaceUserFromUserGroupResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveFaceUserFromUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveFaceUserFromUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveFaceUserFromUserGroupResponseBody) SetCode(v string) *RemoveFaceUserFromUserGroupResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveFaceUserFromUserGroupResponseBody) SetErrorMessage(v string) *RemoveFaceUserFromUserGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RemoveFaceUserFromUserGroupResponseBody) SetRequestId(v string) *RemoveFaceUserFromUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveFaceUserFromUserGroupResponseBody) SetSuccess(v bool) *RemoveFaceUserFromUserGroupResponseBody {
	s.Success = &v
	return s
}

type RemoveFaceUserFromUserGroupResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveFaceUserFromUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveFaceUserFromUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveFaceUserFromUserGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveFaceUserFromUserGroupResponse) SetHeaders(v map[string]*string) *RemoveFaceUserFromUserGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveFaceUserFromUserGroupResponse) SetStatusCode(v int32) *RemoveFaceUserFromUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveFaceUserFromUserGroupResponse) SetBody(v *RemoveFaceUserFromUserGroupResponseBody) *RemoveFaceUserFromUserGroupResponse {
	s.Body = v
	return s
}

type SetDevicePictureLifeCycleRequest struct {
	Day           *int32  `json:"Day,omitempty" xml:"Day,omitempty"`
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s SetDevicePictureLifeCycleRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDevicePictureLifeCycleRequest) GoString() string {
	return s.String()
}

func (s *SetDevicePictureLifeCycleRequest) SetDay(v int32) *SetDevicePictureLifeCycleRequest {
	s.Day = &v
	return s
}

func (s *SetDevicePictureLifeCycleRequest) SetDeviceName(v string) *SetDevicePictureLifeCycleRequest {
	s.DeviceName = &v
	return s
}

func (s *SetDevicePictureLifeCycleRequest) SetIotId(v string) *SetDevicePictureLifeCycleRequest {
	s.IotId = &v
	return s
}

func (s *SetDevicePictureLifeCycleRequest) SetIotInstanceId(v string) *SetDevicePictureLifeCycleRequest {
	s.IotInstanceId = &v
	return s
}

func (s *SetDevicePictureLifeCycleRequest) SetProductKey(v string) *SetDevicePictureLifeCycleRequest {
	s.ProductKey = &v
	return s
}

type SetDevicePictureLifeCycleResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetDevicePictureLifeCycleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDevicePictureLifeCycleResponseBody) GoString() string {
	return s.String()
}

func (s *SetDevicePictureLifeCycleResponseBody) SetCode(v string) *SetDevicePictureLifeCycleResponseBody {
	s.Code = &v
	return s
}

func (s *SetDevicePictureLifeCycleResponseBody) SetErrorMessage(v string) *SetDevicePictureLifeCycleResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SetDevicePictureLifeCycleResponseBody) SetRequestId(v string) *SetDevicePictureLifeCycleResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDevicePictureLifeCycleResponseBody) SetSuccess(v bool) *SetDevicePictureLifeCycleResponseBody {
	s.Success = &v
	return s
}

type SetDevicePictureLifeCycleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDevicePictureLifeCycleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDevicePictureLifeCycleResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDevicePictureLifeCycleResponse) GoString() string {
	return s.String()
}

func (s *SetDevicePictureLifeCycleResponse) SetHeaders(v map[string]*string) *SetDevicePictureLifeCycleResponse {
	s.Headers = v
	return s
}

func (s *SetDevicePictureLifeCycleResponse) SetStatusCode(v int32) *SetDevicePictureLifeCycleResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDevicePictureLifeCycleResponse) SetBody(v *SetDevicePictureLifeCycleResponseBody) *SetDevicePictureLifeCycleResponse {
	s.Body = v
	return s
}

type SetDeviceRecordLifeCycleRequest struct {
	Day           *int32  `json:"Day,omitempty" xml:"Day,omitempty"`
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s SetDeviceRecordLifeCycleRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDeviceRecordLifeCycleRequest) GoString() string {
	return s.String()
}

func (s *SetDeviceRecordLifeCycleRequest) SetDay(v int32) *SetDeviceRecordLifeCycleRequest {
	s.Day = &v
	return s
}

func (s *SetDeviceRecordLifeCycleRequest) SetDeviceName(v string) *SetDeviceRecordLifeCycleRequest {
	s.DeviceName = &v
	return s
}

func (s *SetDeviceRecordLifeCycleRequest) SetIotId(v string) *SetDeviceRecordLifeCycleRequest {
	s.IotId = &v
	return s
}

func (s *SetDeviceRecordLifeCycleRequest) SetIotInstanceId(v string) *SetDeviceRecordLifeCycleRequest {
	s.IotInstanceId = &v
	return s
}

func (s *SetDeviceRecordLifeCycleRequest) SetProductKey(v string) *SetDeviceRecordLifeCycleRequest {
	s.ProductKey = &v
	return s
}

type SetDeviceRecordLifeCycleResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetDeviceRecordLifeCycleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDeviceRecordLifeCycleResponseBody) GoString() string {
	return s.String()
}

func (s *SetDeviceRecordLifeCycleResponseBody) SetCode(v string) *SetDeviceRecordLifeCycleResponseBody {
	s.Code = &v
	return s
}

func (s *SetDeviceRecordLifeCycleResponseBody) SetErrorMessage(v string) *SetDeviceRecordLifeCycleResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SetDeviceRecordLifeCycleResponseBody) SetRequestId(v string) *SetDeviceRecordLifeCycleResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDeviceRecordLifeCycleResponseBody) SetSuccess(v bool) *SetDeviceRecordLifeCycleResponseBody {
	s.Success = &v
	return s
}

type SetDeviceRecordLifeCycleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDeviceRecordLifeCycleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDeviceRecordLifeCycleResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDeviceRecordLifeCycleResponse) GoString() string {
	return s.String()
}

func (s *SetDeviceRecordLifeCycleResponse) SetHeaders(v map[string]*string) *SetDeviceRecordLifeCycleResponse {
	s.Headers = v
	return s
}

func (s *SetDeviceRecordLifeCycleResponse) SetStatusCode(v int32) *SetDeviceRecordLifeCycleResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDeviceRecordLifeCycleResponse) SetBody(v *SetDeviceRecordLifeCycleResponseBody) *SetDeviceRecordLifeCycleResponse {
	s.Body = v
	return s
}

type StopLiveStreamingRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	StreamType    *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s StopLiveStreamingRequest) String() string {
	return tea.Prettify(s)
}

func (s StopLiveStreamingRequest) GoString() string {
	return s.String()
}

func (s *StopLiveStreamingRequest) SetDeviceName(v string) *StopLiveStreamingRequest {
	s.DeviceName = &v
	return s
}

func (s *StopLiveStreamingRequest) SetIotId(v string) *StopLiveStreamingRequest {
	s.IotId = &v
	return s
}

func (s *StopLiveStreamingRequest) SetIotInstanceId(v string) *StopLiveStreamingRequest {
	s.IotInstanceId = &v
	return s
}

func (s *StopLiveStreamingRequest) SetProductKey(v string) *StopLiveStreamingRequest {
	s.ProductKey = &v
	return s
}

func (s *StopLiveStreamingRequest) SetStreamType(v int32) *StopLiveStreamingRequest {
	s.StreamType = &v
	return s
}

type StopLiveStreamingResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopLiveStreamingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopLiveStreamingResponseBody) GoString() string {
	return s.String()
}

func (s *StopLiveStreamingResponseBody) SetCode(v string) *StopLiveStreamingResponseBody {
	s.Code = &v
	return s
}

func (s *StopLiveStreamingResponseBody) SetErrorMessage(v string) *StopLiveStreamingResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StopLiveStreamingResponseBody) SetRequestId(v string) *StopLiveStreamingResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopLiveStreamingResponseBody) SetSuccess(v bool) *StopLiveStreamingResponseBody {
	s.Success = &v
	return s
}

type StopLiveStreamingResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopLiveStreamingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopLiveStreamingResponse) String() string {
	return tea.Prettify(s)
}

func (s StopLiveStreamingResponse) GoString() string {
	return s.String()
}

func (s *StopLiveStreamingResponse) SetHeaders(v map[string]*string) *StopLiveStreamingResponse {
	s.Headers = v
	return s
}

func (s *StopLiveStreamingResponse) SetStatusCode(v int32) *StopLiveStreamingResponse {
	s.StatusCode = &v
	return s
}

func (s *StopLiveStreamingResponse) SetBody(v *StopLiveStreamingResponseBody) *StopLiveStreamingResponse {
	s.Body = v
	return s
}

type StopTriggeredRecordRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	RecordId      *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s StopTriggeredRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s StopTriggeredRecordRequest) GoString() string {
	return s.String()
}

func (s *StopTriggeredRecordRequest) SetDeviceName(v string) *StopTriggeredRecordRequest {
	s.DeviceName = &v
	return s
}

func (s *StopTriggeredRecordRequest) SetIotId(v string) *StopTriggeredRecordRequest {
	s.IotId = &v
	return s
}

func (s *StopTriggeredRecordRequest) SetIotInstanceId(v string) *StopTriggeredRecordRequest {
	s.IotInstanceId = &v
	return s
}

func (s *StopTriggeredRecordRequest) SetProductKey(v string) *StopTriggeredRecordRequest {
	s.ProductKey = &v
	return s
}

func (s *StopTriggeredRecordRequest) SetRecordId(v string) *StopTriggeredRecordRequest {
	s.RecordId = &v
	return s
}

type StopTriggeredRecordResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopTriggeredRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopTriggeredRecordResponseBody) GoString() string {
	return s.String()
}

func (s *StopTriggeredRecordResponseBody) SetCode(v string) *StopTriggeredRecordResponseBody {
	s.Code = &v
	return s
}

func (s *StopTriggeredRecordResponseBody) SetErrorMessage(v string) *StopTriggeredRecordResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StopTriggeredRecordResponseBody) SetRequestId(v string) *StopTriggeredRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopTriggeredRecordResponseBody) SetSuccess(v bool) *StopTriggeredRecordResponseBody {
	s.Success = &v
	return s
}

type StopTriggeredRecordResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopTriggeredRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopTriggeredRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s StopTriggeredRecordResponse) GoString() string {
	return s.String()
}

func (s *StopTriggeredRecordResponse) SetHeaders(v map[string]*string) *StopTriggeredRecordResponse {
	s.Headers = v
	return s
}

func (s *StopTriggeredRecordResponse) SetStatusCode(v int32) *StopTriggeredRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *StopTriggeredRecordResponse) SetBody(v *StopTriggeredRecordResponseBody) *StopTriggeredRecordResponse {
	s.Body = v
	return s
}

type TransferDeviceInstanceRequest struct {
	DeviceNameList   []*string `json:"DeviceNameList,omitempty" xml:"DeviceNameList,omitempty" type:"Repeated"`
	ProductKey       *string   `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	SourceInstanceId *string   `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	TargetInstanceId *string   `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
}

func (s TransferDeviceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferDeviceInstanceRequest) GoString() string {
	return s.String()
}

func (s *TransferDeviceInstanceRequest) SetDeviceNameList(v []*string) *TransferDeviceInstanceRequest {
	s.DeviceNameList = v
	return s
}

func (s *TransferDeviceInstanceRequest) SetProductKey(v string) *TransferDeviceInstanceRequest {
	s.ProductKey = &v
	return s
}

func (s *TransferDeviceInstanceRequest) SetSourceInstanceId(v string) *TransferDeviceInstanceRequest {
	s.SourceInstanceId = &v
	return s
}

func (s *TransferDeviceInstanceRequest) SetTargetInstanceId(v string) *TransferDeviceInstanceRequest {
	s.TargetInstanceId = &v
	return s
}

type TransferDeviceInstanceResponseBody struct {
	Code         *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *TransferDeviceInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TransferDeviceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferDeviceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *TransferDeviceInstanceResponseBody) SetCode(v string) *TransferDeviceInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *TransferDeviceInstanceResponseBody) SetData(v *TransferDeviceInstanceResponseBodyData) *TransferDeviceInstanceResponseBody {
	s.Data = v
	return s
}

func (s *TransferDeviceInstanceResponseBody) SetErrorMessage(v string) *TransferDeviceInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *TransferDeviceInstanceResponseBody) SetRequestId(v string) *TransferDeviceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferDeviceInstanceResponseBody) SetSuccess(v bool) *TransferDeviceInstanceResponseBody {
	s.Success = &v
	return s
}

type TransferDeviceInstanceResponseBodyData struct {
	FailedList  []*TransferDeviceInstanceResponseBodyDataFailedList  `json:"FailedList,omitempty" xml:"FailedList,omitempty" type:"Repeated"`
	SuccessList []*TransferDeviceInstanceResponseBodyDataSuccessList `json:"SuccessList,omitempty" xml:"SuccessList,omitempty" type:"Repeated"`
}

func (s TransferDeviceInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TransferDeviceInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *TransferDeviceInstanceResponseBodyData) SetFailedList(v []*TransferDeviceInstanceResponseBodyDataFailedList) *TransferDeviceInstanceResponseBodyData {
	s.FailedList = v
	return s
}

func (s *TransferDeviceInstanceResponseBodyData) SetSuccessList(v []*TransferDeviceInstanceResponseBodyDataSuccessList) *TransferDeviceInstanceResponseBodyData {
	s.SuccessList = v
	return s
}

type TransferDeviceInstanceResponseBodyDataFailedList struct {
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s TransferDeviceInstanceResponseBodyDataFailedList) String() string {
	return tea.Prettify(s)
}

func (s TransferDeviceInstanceResponseBodyDataFailedList) GoString() string {
	return s.String()
}

func (s *TransferDeviceInstanceResponseBodyDataFailedList) SetDeviceName(v string) *TransferDeviceInstanceResponseBodyDataFailedList {
	s.DeviceName = &v
	return s
}

func (s *TransferDeviceInstanceResponseBodyDataFailedList) SetMessage(v string) *TransferDeviceInstanceResponseBodyDataFailedList {
	s.Message = &v
	return s
}

type TransferDeviceInstanceResponseBodyDataSuccessList struct {
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s TransferDeviceInstanceResponseBodyDataSuccessList) String() string {
	return tea.Prettify(s)
}

func (s TransferDeviceInstanceResponseBodyDataSuccessList) GoString() string {
	return s.String()
}

func (s *TransferDeviceInstanceResponseBodyDataSuccessList) SetDeviceName(v string) *TransferDeviceInstanceResponseBodyDataSuccessList {
	s.DeviceName = &v
	return s
}

func (s *TransferDeviceInstanceResponseBodyDataSuccessList) SetMessage(v string) *TransferDeviceInstanceResponseBodyDataSuccessList {
	s.Message = &v
	return s
}

type TransferDeviceInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransferDeviceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferDeviceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferDeviceInstanceResponse) GoString() string {
	return s.String()
}

func (s *TransferDeviceInstanceResponse) SetHeaders(v map[string]*string) *TransferDeviceInstanceResponse {
	s.Headers = v
	return s
}

func (s *TransferDeviceInstanceResponse) SetStatusCode(v int32) *TransferDeviceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferDeviceInstanceResponse) SetBody(v *TransferDeviceInstanceResponseBody) *TransferDeviceInstanceResponse {
	s.Body = v
	return s
}

type TriggerCapturePictureRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s TriggerCapturePictureRequest) String() string {
	return tea.Prettify(s)
}

func (s TriggerCapturePictureRequest) GoString() string {
	return s.String()
}

func (s *TriggerCapturePictureRequest) SetDeviceName(v string) *TriggerCapturePictureRequest {
	s.DeviceName = &v
	return s
}

func (s *TriggerCapturePictureRequest) SetIotId(v string) *TriggerCapturePictureRequest {
	s.IotId = &v
	return s
}

func (s *TriggerCapturePictureRequest) SetIotInstanceId(v string) *TriggerCapturePictureRequest {
	s.IotInstanceId = &v
	return s
}

func (s *TriggerCapturePictureRequest) SetProductKey(v string) *TriggerCapturePictureRequest {
	s.ProductKey = &v
	return s
}

type TriggerCapturePictureResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TriggerCapturePictureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TriggerCapturePictureResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerCapturePictureResponseBody) SetCode(v string) *TriggerCapturePictureResponseBody {
	s.Code = &v
	return s
}

func (s *TriggerCapturePictureResponseBody) SetData(v string) *TriggerCapturePictureResponseBody {
	s.Data = &v
	return s
}

func (s *TriggerCapturePictureResponseBody) SetErrorMessage(v string) *TriggerCapturePictureResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *TriggerCapturePictureResponseBody) SetRequestId(v string) *TriggerCapturePictureResponseBody {
	s.RequestId = &v
	return s
}

func (s *TriggerCapturePictureResponseBody) SetSuccess(v bool) *TriggerCapturePictureResponseBody {
	s.Success = &v
	return s
}

type TriggerCapturePictureResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TriggerCapturePictureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TriggerCapturePictureResponse) String() string {
	return tea.Prettify(s)
}

func (s TriggerCapturePictureResponse) GoString() string {
	return s.String()
}

func (s *TriggerCapturePictureResponse) SetHeaders(v map[string]*string) *TriggerCapturePictureResponse {
	s.Headers = v
	return s
}

func (s *TriggerCapturePictureResponse) SetStatusCode(v int32) *TriggerCapturePictureResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerCapturePictureResponse) SetBody(v *TriggerCapturePictureResponseBody) *TriggerCapturePictureResponse {
	s.Body = v
	return s
}

type TriggerRecordRequest struct {
	DeviceName        *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId             *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId     *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	PreRecordDuration *int32  `json:"PreRecordDuration,omitempty" xml:"PreRecordDuration,omitempty"`
	ProductKey        *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	RecordDuration    *int32  `json:"RecordDuration,omitempty" xml:"RecordDuration,omitempty"`
	StreamType        *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s TriggerRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s TriggerRecordRequest) GoString() string {
	return s.String()
}

func (s *TriggerRecordRequest) SetDeviceName(v string) *TriggerRecordRequest {
	s.DeviceName = &v
	return s
}

func (s *TriggerRecordRequest) SetIotId(v string) *TriggerRecordRequest {
	s.IotId = &v
	return s
}

func (s *TriggerRecordRequest) SetIotInstanceId(v string) *TriggerRecordRequest {
	s.IotInstanceId = &v
	return s
}

func (s *TriggerRecordRequest) SetPreRecordDuration(v int32) *TriggerRecordRequest {
	s.PreRecordDuration = &v
	return s
}

func (s *TriggerRecordRequest) SetProductKey(v string) *TriggerRecordRequest {
	s.ProductKey = &v
	return s
}

func (s *TriggerRecordRequest) SetRecordDuration(v int32) *TriggerRecordRequest {
	s.RecordDuration = &v
	return s
}

func (s *TriggerRecordRequest) SetStreamType(v int32) *TriggerRecordRequest {
	s.StreamType = &v
	return s
}

type TriggerRecordResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TriggerRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TriggerRecordResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerRecordResponseBody) SetCode(v string) *TriggerRecordResponseBody {
	s.Code = &v
	return s
}

func (s *TriggerRecordResponseBody) SetData(v string) *TriggerRecordResponseBody {
	s.Data = &v
	return s
}

func (s *TriggerRecordResponseBody) SetErrorMessage(v string) *TriggerRecordResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *TriggerRecordResponseBody) SetRequestId(v string) *TriggerRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *TriggerRecordResponseBody) SetSuccess(v bool) *TriggerRecordResponseBody {
	s.Success = &v
	return s
}

type TriggerRecordResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TriggerRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TriggerRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s TriggerRecordResponse) GoString() string {
	return s.String()
}

func (s *TriggerRecordResponse) SetHeaders(v map[string]*string) *TriggerRecordResponse {
	s.Headers = v
	return s
}

func (s *TriggerRecordResponse) SetStatusCode(v int32) *TriggerRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerRecordResponse) SetBody(v *TriggerRecordResponseBody) *TriggerRecordResponse {
	s.Body = v
	return s
}

type UnbindPictureSearchAppWithDevicesRequest struct {
	AppInstanceId *string   `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	DeviceIotIds  []*string `json:"DeviceIotIds,omitempty" xml:"DeviceIotIds,omitempty" type:"Repeated"`
	IotInstanceId *string   `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
}

func (s UnbindPictureSearchAppWithDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindPictureSearchAppWithDevicesRequest) GoString() string {
	return s.String()
}

func (s *UnbindPictureSearchAppWithDevicesRequest) SetAppInstanceId(v string) *UnbindPictureSearchAppWithDevicesRequest {
	s.AppInstanceId = &v
	return s
}

func (s *UnbindPictureSearchAppWithDevicesRequest) SetDeviceIotIds(v []*string) *UnbindPictureSearchAppWithDevicesRequest {
	s.DeviceIotIds = v
	return s
}

func (s *UnbindPictureSearchAppWithDevicesRequest) SetIotInstanceId(v string) *UnbindPictureSearchAppWithDevicesRequest {
	s.IotInstanceId = &v
	return s
}

type UnbindPictureSearchAppWithDevicesResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnbindPictureSearchAppWithDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindPictureSearchAppWithDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindPictureSearchAppWithDevicesResponseBody) SetCode(v string) *UnbindPictureSearchAppWithDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindPictureSearchAppWithDevicesResponseBody) SetErrorMessage(v string) *UnbindPictureSearchAppWithDevicesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UnbindPictureSearchAppWithDevicesResponseBody) SetRequestId(v string) *UnbindPictureSearchAppWithDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindPictureSearchAppWithDevicesResponseBody) SetSuccess(v bool) *UnbindPictureSearchAppWithDevicesResponseBody {
	s.Success = &v
	return s
}

type UnbindPictureSearchAppWithDevicesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindPictureSearchAppWithDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindPictureSearchAppWithDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindPictureSearchAppWithDevicesResponse) GoString() string {
	return s.String()
}

func (s *UnbindPictureSearchAppWithDevicesResponse) SetHeaders(v map[string]*string) *UnbindPictureSearchAppWithDevicesResponse {
	s.Headers = v
	return s
}

func (s *UnbindPictureSearchAppWithDevicesResponse) SetStatusCode(v int32) *UnbindPictureSearchAppWithDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindPictureSearchAppWithDevicesResponse) SetBody(v *UnbindPictureSearchAppWithDevicesResponseBody) *UnbindPictureSearchAppWithDevicesResponse {
	s.Body = v
	return s
}

type UpdateEventRecordPlanRequest struct {
	EventTypes        *string `json:"EventTypes,omitempty" xml:"EventTypes,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PlanId            *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	PreRecordDuration *int32  `json:"PreRecordDuration,omitempty" xml:"PreRecordDuration,omitempty"`
	RecordDuration    *int32  `json:"RecordDuration,omitempty" xml:"RecordDuration,omitempty"`
	TemplateId        *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateEventRecordPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventRecordPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateEventRecordPlanRequest) SetEventTypes(v string) *UpdateEventRecordPlanRequest {
	s.EventTypes = &v
	return s
}

func (s *UpdateEventRecordPlanRequest) SetName(v string) *UpdateEventRecordPlanRequest {
	s.Name = &v
	return s
}

func (s *UpdateEventRecordPlanRequest) SetPlanId(v string) *UpdateEventRecordPlanRequest {
	s.PlanId = &v
	return s
}

func (s *UpdateEventRecordPlanRequest) SetPreRecordDuration(v int32) *UpdateEventRecordPlanRequest {
	s.PreRecordDuration = &v
	return s
}

func (s *UpdateEventRecordPlanRequest) SetRecordDuration(v int32) *UpdateEventRecordPlanRequest {
	s.RecordDuration = &v
	return s
}

func (s *UpdateEventRecordPlanRequest) SetTemplateId(v string) *UpdateEventRecordPlanRequest {
	s.TemplateId = &v
	return s
}

type UpdateEventRecordPlanResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateEventRecordPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventRecordPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEventRecordPlanResponseBody) SetCode(v string) *UpdateEventRecordPlanResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEventRecordPlanResponseBody) SetErrorMessage(v string) *UpdateEventRecordPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateEventRecordPlanResponseBody) SetRequestId(v string) *UpdateEventRecordPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEventRecordPlanResponseBody) SetSuccess(v bool) *UpdateEventRecordPlanResponseBody {
	s.Success = &v
	return s
}

type UpdateEventRecordPlanResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEventRecordPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEventRecordPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventRecordPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateEventRecordPlanResponse) SetHeaders(v map[string]*string) *UpdateEventRecordPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateEventRecordPlanResponse) SetStatusCode(v int32) *UpdateEventRecordPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEventRecordPlanResponse) SetBody(v *UpdateEventRecordPlanResponseBody) *UpdateEventRecordPlanResponse {
	s.Body = v
	return s
}

type UpdateFaceUserRequest struct {
	CustomUserId *string `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	FacePicUrl   *string `json:"FacePicUrl,omitempty" xml:"FacePicUrl,omitempty"`
	IsolationId  *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Params       *string `json:"Params,omitempty" xml:"Params,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateFaceUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateFaceUserRequest) SetCustomUserId(v string) *UpdateFaceUserRequest {
	s.CustomUserId = &v
	return s
}

func (s *UpdateFaceUserRequest) SetFacePicUrl(v string) *UpdateFaceUserRequest {
	s.FacePicUrl = &v
	return s
}

func (s *UpdateFaceUserRequest) SetIsolationId(v string) *UpdateFaceUserRequest {
	s.IsolationId = &v
	return s
}

func (s *UpdateFaceUserRequest) SetName(v string) *UpdateFaceUserRequest {
	s.Name = &v
	return s
}

func (s *UpdateFaceUserRequest) SetParams(v string) *UpdateFaceUserRequest {
	s.Params = &v
	return s
}

func (s *UpdateFaceUserRequest) SetUserId(v string) *UpdateFaceUserRequest {
	s.UserId = &v
	return s
}

type UpdateFaceUserResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateFaceUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFaceUserResponseBody) SetCode(v string) *UpdateFaceUserResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFaceUserResponseBody) SetErrorMessage(v string) *UpdateFaceUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateFaceUserResponseBody) SetRequestId(v string) *UpdateFaceUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFaceUserResponseBody) SetSuccess(v bool) *UpdateFaceUserResponseBody {
	s.Success = &v
	return s
}

type UpdateFaceUserResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFaceUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFaceUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateFaceUserResponse) SetHeaders(v map[string]*string) *UpdateFaceUserResponse {
	s.Headers = v
	return s
}

func (s *UpdateFaceUserResponse) SetStatusCode(v int32) *UpdateFaceUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFaceUserResponse) SetBody(v *UpdateFaceUserResponseBody) *UpdateFaceUserResponse {
	s.Body = v
	return s
}

type UpdateFaceUserGroupAndDeviceGroupRelationRequest struct {
	ControlId   *string `json:"ControlId,omitempty" xml:"ControlId,omitempty"`
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	Relation    *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
}

func (s UpdateFaceUserGroupAndDeviceGroupRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceUserGroupAndDeviceGroupRelationRequest) GoString() string {
	return s.String()
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationRequest) SetControlId(v string) *UpdateFaceUserGroupAndDeviceGroupRelationRequest {
	s.ControlId = &v
	return s
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationRequest) SetIsolationId(v string) *UpdateFaceUserGroupAndDeviceGroupRelationRequest {
	s.IsolationId = &v
	return s
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationRequest) SetRelation(v string) *UpdateFaceUserGroupAndDeviceGroupRelationRequest {
	s.Relation = &v
	return s
}

type UpdateFaceUserGroupAndDeviceGroupRelationResponseBody struct {
	Code         *string                                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *UpdateFaceUserGroupAndDeviceGroupRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateFaceUserGroupAndDeviceGroupRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceUserGroupAndDeviceGroupRelationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationResponseBody) SetCode(v string) *UpdateFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationResponseBody) SetData(v *UpdateFaceUserGroupAndDeviceGroupRelationResponseBodyData) *UpdateFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.Data = v
	return s
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationResponseBody) SetErrorMessage(v string) *UpdateFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationResponseBody) SetRequestId(v string) *UpdateFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationResponseBody) SetSuccess(v bool) *UpdateFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.Success = &v
	return s
}

type UpdateFaceUserGroupAndDeviceGroupRelationResponseBodyData struct {
	ControlId    *string `json:"ControlId,omitempty" xml:"ControlId,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
}

func (s UpdateFaceUserGroupAndDeviceGroupRelationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceUserGroupAndDeviceGroupRelationResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationResponseBodyData) SetControlId(v string) *UpdateFaceUserGroupAndDeviceGroupRelationResponseBodyData {
	s.ControlId = &v
	return s
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationResponseBodyData) SetModifiedTime(v string) *UpdateFaceUserGroupAndDeviceGroupRelationResponseBodyData {
	s.ModifiedTime = &v
	return s
}

type UpdateFaceUserGroupAndDeviceGroupRelationResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFaceUserGroupAndDeviceGroupRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFaceUserGroupAndDeviceGroupRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceUserGroupAndDeviceGroupRelationResponse) GoString() string {
	return s.String()
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationResponse) SetHeaders(v map[string]*string) *UpdateFaceUserGroupAndDeviceGroupRelationResponse {
	s.Headers = v
	return s
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationResponse) SetStatusCode(v int32) *UpdateFaceUserGroupAndDeviceGroupRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationResponse) SetBody(v *UpdateFaceUserGroupAndDeviceGroupRelationResponseBody) *UpdateFaceUserGroupAndDeviceGroupRelationResponse {
	s.Body = v
	return s
}

type UpdateGbDeviceRequest struct {
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	GbId          *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	Password      *string `json:"Password,omitempty" xml:"Password,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s UpdateGbDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGbDeviceRequest) GoString() string {
	return s.String()
}

func (s *UpdateGbDeviceRequest) SetDescription(v string) *UpdateGbDeviceRequest {
	s.Description = &v
	return s
}

func (s *UpdateGbDeviceRequest) SetDeviceName(v string) *UpdateGbDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *UpdateGbDeviceRequest) SetGbId(v string) *UpdateGbDeviceRequest {
	s.GbId = &v
	return s
}

func (s *UpdateGbDeviceRequest) SetIotId(v string) *UpdateGbDeviceRequest {
	s.IotId = &v
	return s
}

func (s *UpdateGbDeviceRequest) SetIotInstanceId(v string) *UpdateGbDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *UpdateGbDeviceRequest) SetPassword(v string) *UpdateGbDeviceRequest {
	s.Password = &v
	return s
}

func (s *UpdateGbDeviceRequest) SetProductKey(v string) *UpdateGbDeviceRequest {
	s.ProductKey = &v
	return s
}

type UpdateGbDeviceResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateGbDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGbDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGbDeviceResponseBody) SetCode(v string) *UpdateGbDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGbDeviceResponseBody) SetErrorMessage(v string) *UpdateGbDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateGbDeviceResponseBody) SetRequestId(v string) *UpdateGbDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGbDeviceResponseBody) SetSuccess(v bool) *UpdateGbDeviceResponseBody {
	s.Success = &v
	return s
}

type UpdateGbDeviceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGbDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGbDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGbDeviceResponse) GoString() string {
	return s.String()
}

func (s *UpdateGbDeviceResponse) SetHeaders(v map[string]*string) *UpdateGbDeviceResponse {
	s.Headers = v
	return s
}

func (s *UpdateGbDeviceResponse) SetStatusCode(v int32) *UpdateGbDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGbDeviceResponse) SetBody(v *UpdateGbDeviceResponseBody) *UpdateGbDeviceResponse {
	s.Body = v
	return s
}

type UpdateInstanceInternetProtocolRequest struct {
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	IpVersion     *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
}

func (s UpdateInstanceInternetProtocolRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceInternetProtocolRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInternetProtocolRequest) SetIotInstanceId(v string) *UpdateInstanceInternetProtocolRequest {
	s.IotInstanceId = &v
	return s
}

func (s *UpdateInstanceInternetProtocolRequest) SetIpVersion(v string) *UpdateInstanceInternetProtocolRequest {
	s.IpVersion = &v
	return s
}

type UpdateInstanceInternetProtocolResponseBody struct {
	Code         *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInstanceInternetProtocolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceInternetProtocolResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInternetProtocolResponseBody) SetCode(v string) *UpdateInstanceInternetProtocolResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceInternetProtocolResponseBody) SetData(v map[string]interface{}) *UpdateInstanceInternetProtocolResponseBody {
	s.Data = v
	return s
}

func (s *UpdateInstanceInternetProtocolResponseBody) SetErrorMessage(v string) *UpdateInstanceInternetProtocolResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateInstanceInternetProtocolResponseBody) SetRequestId(v string) *UpdateInstanceInternetProtocolResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceInternetProtocolResponseBody) SetSuccess(v bool) *UpdateInstanceInternetProtocolResponseBody {
	s.Success = &v
	return s
}

type UpdateInstanceInternetProtocolResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceInternetProtocolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceInternetProtocolResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceInternetProtocolResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInternetProtocolResponse) SetHeaders(v map[string]*string) *UpdateInstanceInternetProtocolResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceInternetProtocolResponse) SetStatusCode(v int32) *UpdateInstanceInternetProtocolResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceInternetProtocolResponse) SetBody(v *UpdateInstanceInternetProtocolResponseBody) *UpdateInstanceInternetProtocolResponse {
	s.Body = v
	return s
}

type UpdatePictureSearchAppRequest struct {
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdatePictureSearchAppRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePictureSearchAppRequest) GoString() string {
	return s.String()
}

func (s *UpdatePictureSearchAppRequest) SetAppInstanceId(v string) *UpdatePictureSearchAppRequest {
	s.AppInstanceId = &v
	return s
}

func (s *UpdatePictureSearchAppRequest) SetDescription(v string) *UpdatePictureSearchAppRequest {
	s.Description = &v
	return s
}

func (s *UpdatePictureSearchAppRequest) SetName(v string) *UpdatePictureSearchAppRequest {
	s.Name = &v
	return s
}

type UpdatePictureSearchAppResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdatePictureSearchAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePictureSearchAppResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePictureSearchAppResponseBody) SetCode(v string) *UpdatePictureSearchAppResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePictureSearchAppResponseBody) SetErrorMessage(v string) *UpdatePictureSearchAppResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdatePictureSearchAppResponseBody) SetRequestId(v string) *UpdatePictureSearchAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePictureSearchAppResponseBody) SetSuccess(v bool) *UpdatePictureSearchAppResponseBody {
	s.Success = &v
	return s
}

type UpdatePictureSearchAppResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePictureSearchAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePictureSearchAppResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePictureSearchAppResponse) GoString() string {
	return s.String()
}

func (s *UpdatePictureSearchAppResponse) SetHeaders(v map[string]*string) *UpdatePictureSearchAppResponse {
	s.Headers = v
	return s
}

func (s *UpdatePictureSearchAppResponse) SetStatusCode(v int32) *UpdatePictureSearchAppResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePictureSearchAppResponse) SetBody(v *UpdatePictureSearchAppResponseBody) *UpdatePictureSearchAppResponse {
	s.Body = v
	return s
}

type UpdateRecordPlanRequest struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PlanId     *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateRecordPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecordPlanRequest) SetName(v string) *UpdateRecordPlanRequest {
	s.Name = &v
	return s
}

func (s *UpdateRecordPlanRequest) SetPlanId(v string) *UpdateRecordPlanRequest {
	s.PlanId = &v
	return s
}

func (s *UpdateRecordPlanRequest) SetTemplateId(v string) *UpdateRecordPlanRequest {
	s.TemplateId = &v
	return s
}

type UpdateRecordPlanResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRecordPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRecordPlanResponseBody) SetCode(v string) *UpdateRecordPlanResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRecordPlanResponseBody) SetErrorMessage(v string) *UpdateRecordPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateRecordPlanResponseBody) SetRequestId(v string) *UpdateRecordPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRecordPlanResponseBody) SetSuccess(v bool) *UpdateRecordPlanResponseBody {
	s.Success = &v
	return s
}

type UpdateRecordPlanResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRecordPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRecordPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateRecordPlanResponse) SetHeaders(v map[string]*string) *UpdateRecordPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateRecordPlanResponse) SetStatusCode(v int32) *UpdateRecordPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRecordPlanResponse) SetBody(v *UpdateRecordPlanResponseBody) *UpdateRecordPlanResponse {
	s.Body = v
	return s
}

type UpdateRtmpKeyRequest struct {
	DeviceName        *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId             *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId     *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey        *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	PullAuthKey       *string `json:"PullAuthKey,omitempty" xml:"PullAuthKey,omitempty"`
	PullKeyExpireTime *int32  `json:"PullKeyExpireTime,omitempty" xml:"PullKeyExpireTime,omitempty"`
	PushAuthKey       *string `json:"PushAuthKey,omitempty" xml:"PushAuthKey,omitempty"`
	PushKeyExpireTime *int32  `json:"PushKeyExpireTime,omitempty" xml:"PushKeyExpireTime,omitempty"`
}

func (s UpdateRtmpKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRtmpKeyRequest) GoString() string {
	return s.String()
}

func (s *UpdateRtmpKeyRequest) SetDeviceName(v string) *UpdateRtmpKeyRequest {
	s.DeviceName = &v
	return s
}

func (s *UpdateRtmpKeyRequest) SetIotId(v string) *UpdateRtmpKeyRequest {
	s.IotId = &v
	return s
}

func (s *UpdateRtmpKeyRequest) SetIotInstanceId(v string) *UpdateRtmpKeyRequest {
	s.IotInstanceId = &v
	return s
}

func (s *UpdateRtmpKeyRequest) SetProductKey(v string) *UpdateRtmpKeyRequest {
	s.ProductKey = &v
	return s
}

func (s *UpdateRtmpKeyRequest) SetPullAuthKey(v string) *UpdateRtmpKeyRequest {
	s.PullAuthKey = &v
	return s
}

func (s *UpdateRtmpKeyRequest) SetPullKeyExpireTime(v int32) *UpdateRtmpKeyRequest {
	s.PullKeyExpireTime = &v
	return s
}

func (s *UpdateRtmpKeyRequest) SetPushAuthKey(v string) *UpdateRtmpKeyRequest {
	s.PushAuthKey = &v
	return s
}

func (s *UpdateRtmpKeyRequest) SetPushKeyExpireTime(v int32) *UpdateRtmpKeyRequest {
	s.PushKeyExpireTime = &v
	return s
}

type UpdateRtmpKeyResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRtmpKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRtmpKeyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRtmpKeyResponseBody) SetCode(v string) *UpdateRtmpKeyResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRtmpKeyResponseBody) SetErrorMessage(v string) *UpdateRtmpKeyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateRtmpKeyResponseBody) SetRequestId(v string) *UpdateRtmpKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRtmpKeyResponseBody) SetSuccess(v bool) *UpdateRtmpKeyResponseBody {
	s.Success = &v
	return s
}

type UpdateRtmpKeyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRtmpKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRtmpKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRtmpKeyResponse) GoString() string {
	return s.String()
}

func (s *UpdateRtmpKeyResponse) SetHeaders(v map[string]*string) *UpdateRtmpKeyResponse {
	s.Headers = v
	return s
}

func (s *UpdateRtmpKeyResponse) SetStatusCode(v int32) *UpdateRtmpKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRtmpKeyResponse) SetBody(v *UpdateRtmpKeyResponseBody) *UpdateRtmpKeyResponse {
	s.Body = v
	return s
}

type UpdateTimeTemplateRequest struct {
	AllDay       *int32                                   `json:"AllDay,omitempty" xml:"AllDay,omitempty"`
	Name         *string                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateId   *string                                  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TimeSections []*UpdateTimeTemplateRequestTimeSections `json:"TimeSections,omitempty" xml:"TimeSections,omitempty" type:"Repeated"`
}

func (s UpdateTimeTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTimeTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateTimeTemplateRequest) SetAllDay(v int32) *UpdateTimeTemplateRequest {
	s.AllDay = &v
	return s
}

func (s *UpdateTimeTemplateRequest) SetName(v string) *UpdateTimeTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdateTimeTemplateRequest) SetTemplateId(v string) *UpdateTimeTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateTimeTemplateRequest) SetTimeSections(v []*UpdateTimeTemplateRequestTimeSections) *UpdateTimeTemplateRequest {
	s.TimeSections = v
	return s
}

type UpdateTimeTemplateRequestTimeSections struct {
	Begin     *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	DayOfWeek *int32 `json:"DayOfWeek,omitempty" xml:"DayOfWeek,omitempty"`
	End       *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s UpdateTimeTemplateRequestTimeSections) String() string {
	return tea.Prettify(s)
}

func (s UpdateTimeTemplateRequestTimeSections) GoString() string {
	return s.String()
}

func (s *UpdateTimeTemplateRequestTimeSections) SetBegin(v int32) *UpdateTimeTemplateRequestTimeSections {
	s.Begin = &v
	return s
}

func (s *UpdateTimeTemplateRequestTimeSections) SetDayOfWeek(v int32) *UpdateTimeTemplateRequestTimeSections {
	s.DayOfWeek = &v
	return s
}

func (s *UpdateTimeTemplateRequestTimeSections) SetEnd(v int32) *UpdateTimeTemplateRequestTimeSections {
	s.End = &v
	return s
}

type UpdateTimeTemplateResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTimeTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTimeTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTimeTemplateResponseBody) SetCode(v string) *UpdateTimeTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTimeTemplateResponseBody) SetErrorMessage(v string) *UpdateTimeTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateTimeTemplateResponseBody) SetRequestId(v string) *UpdateTimeTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTimeTemplateResponseBody) SetSuccess(v bool) *UpdateTimeTemplateResponseBody {
	s.Success = &v
	return s
}

type UpdateTimeTemplateResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTimeTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTimeTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTimeTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateTimeTemplateResponse) SetHeaders(v map[string]*string) *UpdateTimeTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateTimeTemplateResponse) SetStatusCode(v int32) *UpdateTimeTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTimeTemplateResponse) SetBody(v *UpdateTimeTemplateResponseBody) *UpdateTimeTemplateResponse {
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
		"ap-northeast-1":              tea.String("linkvisual.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("linkvisual.aliyuncs.com"),
		"ap-south-1":                  tea.String("linkvisual.aliyuncs.com"),
		"ap-southeast-1":              tea.String("linkvisual.aliyuncs.com"),
		"ap-southeast-2":              tea.String("linkvisual.aliyuncs.com"),
		"ap-southeast-3":              tea.String("linkvisual.aliyuncs.com"),
		"ap-southeast-5":              tea.String("linkvisual.aliyuncs.com"),
		"cn-beijing":                  tea.String("linkvisual.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("linkvisual.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("linkvisual.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("linkvisual.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("linkvisual.aliyuncs.com"),
		"cn-chengdu":                  tea.String("linkvisual.aliyuncs.com"),
		"cn-edge-1":                   tea.String("linkvisual.aliyuncs.com"),
		"cn-fujian":                   tea.String("linkvisual.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("linkvisual.aliyuncs.com"),
		"cn-hongkong":                 tea.String("linkvisual.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("linkvisual.aliyuncs.com"),
		"cn-huhehaote":                tea.String("linkvisual.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("linkvisual.aliyuncs.com"),
		"cn-qingdao":                  tea.String("linkvisual.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("linkvisual.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("linkvisual.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("linkvisual.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("linkvisual.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("linkvisual.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("linkvisual.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("linkvisual.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("linkvisual.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("linkvisual.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("linkvisual.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("linkvisual.aliyuncs.com"),
		"cn-wuhan":                    tea.String("linkvisual.aliyuncs.com"),
		"cn-yushanfang":               tea.String("linkvisual.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("linkvisual.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("linkvisual.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("linkvisual.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("linkvisual.aliyuncs.com"),
		"eu-central-1":                tea.String("linkvisual.aliyuncs.com"),
		"eu-west-1":                   tea.String("linkvisual.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("linkvisual.aliyuncs.com"),
		"me-east-1":                   tea.String("linkvisual.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("linkvisual.aliyuncs.com"),
		"us-east-1":                   tea.String("linkvisual.aliyuncs.com"),
		"us-west-1":                   tea.String("linkvisual.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("linkvisual"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddEventRecordPlanDeviceWithOptions(request *AddEventRecordPlanDeviceRequest, runtime *util.RuntimeOptions) (_result *AddEventRecordPlanDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.StreamType)) {
		query["StreamType"] = request.StreamType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddEventRecordPlanDevice"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddEventRecordPlanDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddEventRecordPlanDevice(request *AddEventRecordPlanDeviceRequest) (_result *AddEventRecordPlanDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddEventRecordPlanDeviceResponse{}
	_body, _err := client.AddEventRecordPlanDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddFaceDeviceGroupWithOptions(request *AddFaceDeviceGroupRequest, runtime *util.RuntimeOptions) (_result *AddFaceDeviceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceGroupName)) {
		query["DeviceGroupName"] = request.DeviceGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddFaceDeviceGroup"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddFaceDeviceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddFaceDeviceGroup(request *AddFaceDeviceGroupRequest) (_result *AddFaceDeviceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddFaceDeviceGroupResponse{}
	_body, _err := client.AddFaceDeviceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddFaceDeviceToDeviceGroupWithOptions(request *AddFaceDeviceToDeviceGroupRequest, runtime *util.RuntimeOptions) (_result *AddFaceDeviceToDeviceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceGroupId)) {
		query["DeviceGroupId"] = request.DeviceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddFaceDeviceToDeviceGroup"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddFaceDeviceToDeviceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddFaceDeviceToDeviceGroup(request *AddFaceDeviceToDeviceGroupRequest) (_result *AddFaceDeviceToDeviceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddFaceDeviceToDeviceGroupResponse{}
	_body, _err := client.AddFaceDeviceToDeviceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddFaceUserWithOptions(request *AddFaceUserRequest, runtime *util.RuntimeOptions) (_result *AddFaceUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomUserId)) {
		query["CustomUserId"] = request.CustomUserId
	}

	if !tea.BoolValue(util.IsUnset(request.FacePicUrl)) {
		query["FacePicUrl"] = request.FacePicUrl
	}

	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		query["Params"] = request.Params
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddFaceUser"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddFaceUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddFaceUser(request *AddFaceUserRequest) (_result *AddFaceUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddFaceUserResponse{}
	_body, _err := client.AddFaceUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddFaceUserGroupWithOptions(request *AddFaceUserGroupRequest, runtime *util.RuntimeOptions) (_result *AddFaceUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupName)) {
		query["UserGroupName"] = request.UserGroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddFaceUserGroup"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddFaceUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddFaceUserGroup(request *AddFaceUserGroupRequest) (_result *AddFaceUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddFaceUserGroupResponse{}
	_body, _err := client.AddFaceUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddFaceUserGroupAndDeviceGroupRelationWithOptions(request *AddFaceUserGroupAndDeviceGroupRelationRequest, runtime *util.RuntimeOptions) (_result *AddFaceUserGroupAndDeviceGroupRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceGroupId)) {
		query["DeviceGroupId"] = request.DeviceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	if !tea.BoolValue(util.IsUnset(request.Relation)) {
		query["Relation"] = request.Relation
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddFaceUserGroupAndDeviceGroupRelation"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddFaceUserGroupAndDeviceGroupRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddFaceUserGroupAndDeviceGroupRelation(request *AddFaceUserGroupAndDeviceGroupRelationRequest) (_result *AddFaceUserGroupAndDeviceGroupRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddFaceUserGroupAndDeviceGroupRelationResponse{}
	_body, _err := client.AddFaceUserGroupAndDeviceGroupRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddFaceUserPictureWithOptions(request *AddFaceUserPictureRequest, runtime *util.RuntimeOptions) (_result *AddFaceUserPictureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FacePicUrl)) {
		query["FacePicUrl"] = request.FacePicUrl
	}

	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddFaceUserPicture"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddFaceUserPictureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddFaceUserPicture(request *AddFaceUserPictureRequest) (_result *AddFaceUserPictureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddFaceUserPictureResponse{}
	_body, _err := client.AddFaceUserPictureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddFaceUserToUserGroupWithOptions(request *AddFaceUserToUserGroupRequest, runtime *util.RuntimeOptions) (_result *AddFaceUserToUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddFaceUserToUserGroup"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddFaceUserToUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddFaceUserToUserGroup(request *AddFaceUserToUserGroupRequest) (_result *AddFaceUserToUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddFaceUserToUserGroupResponse{}
	_body, _err := client.AddFaceUserToUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddRecordPlanDeviceWithOptions(request *AddRecordPlanDeviceRequest, runtime *util.RuntimeOptions) (_result *AddRecordPlanDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.StreamType)) {
		query["StreamType"] = request.StreamType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddRecordPlanDevice"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddRecordPlanDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddRecordPlanDevice(request *AddRecordPlanDeviceRequest) (_result *AddRecordPlanDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddRecordPlanDeviceResponse{}
	_body, _err := client.AddRecordPlanDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchQueryVisionDeviceInfoWithOptions(request *BatchQueryVisionDeviceInfoRequest, runtime *util.RuntimeOptions) (_result *BatchQueryVisionDeviceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceNameList)) {
		query["DeviceNameList"] = request.DeviceNameList
	}

	if !tea.BoolValue(util.IsUnset(request.IotIdList)) {
		query["IotIdList"] = request.IotIdList
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchQueryVisionDeviceInfo"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchQueryVisionDeviceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchQueryVisionDeviceInfo(request *BatchQueryVisionDeviceInfoRequest) (_result *BatchQueryVisionDeviceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchQueryVisionDeviceInfoResponse{}
	_body, _err := client.BatchQueryVisionDeviceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindPictureSearchAppWithDevicesWithOptions(request *BindPictureSearchAppWithDevicesRequest, runtime *util.RuntimeOptions) (_result *BindPictureSearchAppWithDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceId)) {
		query["AppInstanceId"] = request.AppInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceIotIds)) {
		query["DeviceIotIds"] = request.DeviceIotIds
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindPictureSearchAppWithDevices"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindPictureSearchAppWithDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindPictureSearchAppWithDevices(request *BindPictureSearchAppWithDevicesRequest) (_result *BindPictureSearchAppWithDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindPictureSearchAppWithDevicesResponse{}
	_body, _err := client.BindPictureSearchAppWithDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckFaceUserDoExistOnDeviceWithOptions(request *CheckFaceUserDoExistOnDeviceRequest, runtime *util.RuntimeOptions) (_result *CheckFaceUserDoExistOnDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckFaceUserDoExistOnDevice"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckFaceUserDoExistOnDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckFaceUserDoExistOnDevice(request *CheckFaceUserDoExistOnDeviceRequest) (_result *CheckFaceUserDoExistOnDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckFaceUserDoExistOnDeviceResponse{}
	_body, _err := client.CheckFaceUserDoExistOnDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClearFaceDeviceDBWithOptions(request *ClearFaceDeviceDBRequest, runtime *util.RuntimeOptions) (_result *ClearFaceDeviceDBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ClearFaceDeviceDB"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ClearFaceDeviceDBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClearFaceDeviceDB(request *ClearFaceDeviceDBRequest) (_result *ClearFaceDeviceDBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClearFaceDeviceDBResponse{}
	_body, _err := client.ClearFaceDeviceDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEventRecordPlanWithOptions(request *CreateEventRecordPlanRequest, runtime *util.RuntimeOptions) (_result *CreateEventRecordPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventTypes)) {
		query["EventTypes"] = request.EventTypes
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PreRecordDuration)) {
		query["PreRecordDuration"] = request.PreRecordDuration
	}

	if !tea.BoolValue(util.IsUnset(request.RecordDuration)) {
		query["RecordDuration"] = request.RecordDuration
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEventRecordPlan"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEventRecordPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEventRecordPlan(request *CreateEventRecordPlanRequest) (_result *CreateEventRecordPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEventRecordPlanResponse{}
	_body, _err := client.CreateEventRecordPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGbDeviceWithOptions(request *CreateGbDeviceRequest, runtime *util.RuntimeOptions) (_result *CreateGbDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceType)) {
		query["DeviceType"] = request.DeviceType
	}

	if !tea.BoolValue(util.IsUnset(request.GbId)) {
		query["GbId"] = request.GbId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MediaNetProtocol)) {
		query["MediaNetProtocol"] = request.MediaNetProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.SubProductKey)) {
		query["SubProductKey"] = request.SubProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGbDevice"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGbDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGbDevice(request *CreateGbDeviceRequest) (_result *CreateGbDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGbDeviceResponse{}
	_body, _err := client.CreateGbDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLocalFileUploadJobWithOptions(request *CreateLocalFileUploadJobRequest, runtime *util.RuntimeOptions) (_result *CreateLocalFileUploadJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TimeSlot)) {
		query["TimeSlot"] = request.TimeSlot
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLocalFileUploadJob"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLocalFileUploadJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLocalFileUploadJob(request *CreateLocalFileUploadJobRequest) (_result *CreateLocalFileUploadJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLocalFileUploadJobResponse{}
	_body, _err := client.CreateLocalFileUploadJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLocalRecordDownloadByTimeJobWithOptions(request *CreateLocalRecordDownloadByTimeJobRequest, runtime *util.RuntimeOptions) (_result *CreateLocalRecordDownloadByTimeJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.Speed)) {
		query["Speed"] = request.Speed
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLocalRecordDownloadByTimeJob"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLocalRecordDownloadByTimeJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLocalRecordDownloadByTimeJob(request *CreateLocalRecordDownloadByTimeJobRequest) (_result *CreateLocalRecordDownloadByTimeJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLocalRecordDownloadByTimeJobResponse{}
	_body, _err := client.CreateLocalRecordDownloadByTimeJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePictureSearchAppWithOptions(request *CreatePictureSearchAppRequest, runtime *util.RuntimeOptions) (_result *CreatePictureSearchAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Desc)) {
		query["Desc"] = request.Desc
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePictureSearchApp"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePictureSearchAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePictureSearchApp(request *CreatePictureSearchAppRequest) (_result *CreatePictureSearchAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePictureSearchAppResponse{}
	_body, _err := client.CreatePictureSearchAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePictureSearchJobWithOptions(request *CreatePictureSearchJobRequest, runtime *util.RuntimeOptions) (_result *CreatePictureSearchJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceId)) {
		query["AppInstanceId"] = request.AppInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.BodyThreshold)) {
		query["BodyThreshold"] = request.BodyThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PictureSearchType)) {
		query["PictureSearchType"] = request.PictureSearchType
	}

	if !tea.BoolValue(util.IsUnset(request.SearchPicUrl)) {
		query["SearchPicUrl"] = request.SearchPicUrl
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Threshold)) {
		query["Threshold"] = request.Threshold
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePictureSearchJob"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePictureSearchJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePictureSearchJob(request *CreatePictureSearchJobRequest) (_result *CreatePictureSearchJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePictureSearchJobResponse{}
	_body, _err := client.CreatePictureSearchJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRecordDownloadByTimeJobWithOptions(request *CreateRecordDownloadByTimeJobRequest, runtime *util.RuntimeOptions) (_result *CreateRecordDownloadByTimeJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.RecordType)) {
		query["RecordType"] = request.RecordType
	}

	if !tea.BoolValue(util.IsUnset(request.StreamType)) {
		query["StreamType"] = request.StreamType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRecordDownloadByTimeJob"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRecordDownloadByTimeJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRecordDownloadByTimeJob(request *CreateRecordDownloadByTimeJobRequest) (_result *CreateRecordDownloadByTimeJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRecordDownloadByTimeJobResponse{}
	_body, _err := client.CreateRecordDownloadByTimeJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRecordPlanWithOptions(request *CreateRecordPlanRequest, runtime *util.RuntimeOptions) (_result *CreateRecordPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRecordPlan"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRecordPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRecordPlan(request *CreateRecordPlanRequest) (_result *CreateRecordPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRecordPlanResponse{}
	_body, _err := client.CreateRecordPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRtmpDeviceWithOptions(request *CreateRtmpDeviceRequest, runtime *util.RuntimeOptions) (_result *CreateRtmpDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.PullAuthKey)) {
		query["PullAuthKey"] = request.PullAuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.PullKeyExpireTime)) {
		query["PullKeyExpireTime"] = request.PullKeyExpireTime
	}

	if !tea.BoolValue(util.IsUnset(request.PushAuthKey)) {
		query["PushAuthKey"] = request.PushAuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.PushKeyExpireTime)) {
		query["PushKeyExpireTime"] = request.PushKeyExpireTime
	}

	if !tea.BoolValue(util.IsUnset(request.SubStreamName)) {
		query["SubStreamName"] = request.SubStreamName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRtmpDevice"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRtmpDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRtmpDevice(request *CreateRtmpDeviceRequest) (_result *CreateRtmpDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRtmpDeviceResponse{}
	_body, _err := client.CreateRtmpDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateStreamPushJobWithOptions(request *CreateStreamPushJobRequest, runtime *util.RuntimeOptions) (_result *CreateStreamPushJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JobType)) {
		query["JobType"] = request.JobType
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.PushUrl)) {
		query["PushUrl"] = request.PushUrl
	}

	if !tea.BoolValue(util.IsUnset(request.StreamType)) {
		query["StreamType"] = request.StreamType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateStreamPushJob"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateStreamPushJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateStreamPushJob(request *CreateStreamPushJobRequest) (_result *CreateStreamPushJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateStreamPushJobResponse{}
	_body, _err := client.CreateStreamPushJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateStreamSnapshotJobWithOptions(request *CreateStreamSnapshotJobRequest, runtime *util.RuntimeOptions) (_result *CreateStreamSnapshotJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotInterval)) {
		query["SnapshotInterval"] = request.SnapshotInterval
	}

	if !tea.BoolValue(util.IsUnset(request.StreamType)) {
		query["StreamType"] = request.StreamType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateStreamSnapshotJob"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateStreamSnapshotJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateStreamSnapshotJob(request *CreateStreamSnapshotJobRequest) (_result *CreateStreamSnapshotJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateStreamSnapshotJobResponse{}
	_body, _err := client.CreateStreamSnapshotJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTimeTemplateWithOptions(request *CreateTimeTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateTimeTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllDay)) {
		query["AllDay"] = request.AllDay
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.TimeSections)) {
		query["TimeSections"] = request.TimeSections
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTimeTemplate"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTimeTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTimeTemplate(request *CreateTimeTemplateRequest) (_result *CreateTimeTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTimeTemplateResponse{}
	_body, _err := client.CreateTimeTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEventRecordPlanWithOptions(request *DeleteEventRecordPlanRequest, runtime *util.RuntimeOptions) (_result *DeleteEventRecordPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEventRecordPlan"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEventRecordPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEventRecordPlan(request *DeleteEventRecordPlanRequest) (_result *DeleteEventRecordPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEventRecordPlanResponse{}
	_body, _err := client.DeleteEventRecordPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEventRecordPlanDeviceWithOptions(request *DeleteEventRecordPlanDeviceRequest, runtime *util.RuntimeOptions) (_result *DeleteEventRecordPlanDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.StreamType)) {
		query["StreamType"] = request.StreamType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEventRecordPlanDevice"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEventRecordPlanDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEventRecordPlanDevice(request *DeleteEventRecordPlanDeviceRequest) (_result *DeleteEventRecordPlanDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEventRecordPlanDeviceResponse{}
	_body, _err := client.DeleteEventRecordPlanDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFaceDeviceGroupWithOptions(request *DeleteFaceDeviceGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteFaceDeviceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceGroupId)) {
		query["DeviceGroupId"] = request.DeviceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFaceDeviceGroup"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFaceDeviceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFaceDeviceGroup(request *DeleteFaceDeviceGroupRequest) (_result *DeleteFaceDeviceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFaceDeviceGroupResponse{}
	_body, _err := client.DeleteFaceDeviceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFaceUserWithOptions(request *DeleteFaceUserRequest, runtime *util.RuntimeOptions) (_result *DeleteFaceUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFaceUser"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFaceUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFaceUser(request *DeleteFaceUserRequest) (_result *DeleteFaceUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFaceUserResponse{}
	_body, _err := client.DeleteFaceUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFaceUserGroupWithOptions(request *DeleteFaceUserGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteFaceUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFaceUserGroup"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFaceUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFaceUserGroup(request *DeleteFaceUserGroupRequest) (_result *DeleteFaceUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFaceUserGroupResponse{}
	_body, _err := client.DeleteFaceUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFaceUserGroupAndDeviceGroupRelationWithOptions(request *DeleteFaceUserGroupAndDeviceGroupRelationRequest, runtime *util.RuntimeOptions) (_result *DeleteFaceUserGroupAndDeviceGroupRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ControlId)) {
		query["ControlId"] = request.ControlId
	}

	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFaceUserGroupAndDeviceGroupRelation"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFaceUserGroupAndDeviceGroupRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFaceUserGroupAndDeviceGroupRelation(request *DeleteFaceUserGroupAndDeviceGroupRelationRequest) (_result *DeleteFaceUserGroupAndDeviceGroupRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFaceUserGroupAndDeviceGroupRelationResponse{}
	_body, _err := client.DeleteFaceUserGroupAndDeviceGroupRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFaceUserPictureWithOptions(request *DeleteFaceUserPictureRequest, runtime *util.RuntimeOptions) (_result *DeleteFaceUserPictureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FacePicMd5)) {
		query["FacePicMd5"] = request.FacePicMd5
	}

	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFaceUserPicture"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFaceUserPictureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFaceUserPicture(request *DeleteFaceUserPictureRequest) (_result *DeleteFaceUserPictureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFaceUserPictureResponse{}
	_body, _err := client.DeleteFaceUserPictureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGbDeviceWithOptions(request *DeleteGbDeviceRequest, runtime *util.RuntimeOptions) (_result *DeleteGbDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGbDevice"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGbDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGbDevice(request *DeleteGbDeviceRequest) (_result *DeleteGbDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGbDeviceResponse{}
	_body, _err := client.DeleteGbDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLocalFileUploadJobWithOptions(request *DeleteLocalFileUploadJobRequest, runtime *util.RuntimeOptions) (_result *DeleteLocalFileUploadJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLocalFileUploadJob"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLocalFileUploadJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLocalFileUploadJob(request *DeleteLocalFileUploadJobRequest) (_result *DeleteLocalFileUploadJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLocalFileUploadJobResponse{}
	_body, _err := client.DeleteLocalFileUploadJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePictureWithOptions(request *DeletePictureRequest, runtime *util.RuntimeOptions) (_result *DeletePictureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PictureIdList)) {
		query["PictureIdList"] = request.PictureIdList
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePicture"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePictureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePicture(request *DeletePictureRequest) (_result *DeletePictureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePictureResponse{}
	_body, _err := client.DeletePictureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRecordWithOptions(request *DeleteRecordRequest, runtime *util.RuntimeOptions) (_result *DeleteRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.FileNameList)) {
		query["FileNameList"] = request.FileNameList
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRecord"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRecord(request *DeleteRecordRequest) (_result *DeleteRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRecordResponse{}
	_body, _err := client.DeleteRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRecordPlanWithOptions(request *DeleteRecordPlanRequest, runtime *util.RuntimeOptions) (_result *DeleteRecordPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRecordPlan"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRecordPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRecordPlan(request *DeleteRecordPlanRequest) (_result *DeleteRecordPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRecordPlanResponse{}
	_body, _err := client.DeleteRecordPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRecordPlanDeviceWithOptions(request *DeleteRecordPlanDeviceRequest, runtime *util.RuntimeOptions) (_result *DeleteRecordPlanDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.StreamType)) {
		query["StreamType"] = request.StreamType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRecordPlanDevice"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRecordPlanDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRecordPlanDevice(request *DeleteRecordPlanDeviceRequest) (_result *DeleteRecordPlanDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRecordPlanDeviceResponse{}
	_body, _err := client.DeleteRecordPlanDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRtmpDeviceWithOptions(request *DeleteRtmpDeviceRequest, runtime *util.RuntimeOptions) (_result *DeleteRtmpDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRtmpDevice"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRtmpDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRtmpDevice(request *DeleteRtmpDeviceRequest) (_result *DeleteRtmpDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRtmpDeviceResponse{}
	_body, _err := client.DeleteRtmpDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRtmpKeyWithOptions(request *DeleteRtmpKeyRequest, runtime *util.RuntimeOptions) (_result *DeleteRtmpKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRtmpKey"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRtmpKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRtmpKey(request *DeleteRtmpKeyRequest) (_result *DeleteRtmpKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRtmpKeyResponse{}
	_body, _err := client.DeleteRtmpKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteStreamPushJobWithOptions(request *DeleteStreamPushJobRequest, runtime *util.RuntimeOptions) (_result *DeleteStreamPushJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteStreamPushJob"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteStreamPushJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteStreamPushJob(request *DeleteStreamPushJobRequest) (_result *DeleteStreamPushJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteStreamPushJobResponse{}
	_body, _err := client.DeleteStreamPushJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteStreamSnapshotJobWithOptions(request *DeleteStreamSnapshotJobRequest, runtime *util.RuntimeOptions) (_result *DeleteStreamSnapshotJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.StreamType)) {
		query["StreamType"] = request.StreamType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteStreamSnapshotJob"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteStreamSnapshotJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteStreamSnapshotJob(request *DeleteStreamSnapshotJobRequest) (_result *DeleteStreamSnapshotJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteStreamSnapshotJobResponse{}
	_body, _err := client.DeleteStreamSnapshotJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTimeTemplateWithOptions(request *DeleteTimeTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteTimeTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTimeTemplate"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTimeTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTimeTemplate(request *DeleteTimeTemplateRequest) (_result *DeleteTimeTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTimeTemplateResponse{}
	_body, _err := client.DeleteTimeTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectUserFaceByUrlWithOptions(request *DetectUserFaceByUrlRequest, runtime *util.RuntimeOptions) (_result *DetectUserFaceByUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FacePicUrl)) {
		query["FacePicUrl"] = request.FacePicUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectUserFaceByUrl"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectUserFaceByUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectUserFaceByUrl(request *DetectUserFaceByUrlRequest) (_result *DetectUserFaceByUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectUserFaceByUrlResponse{}
	_body, _err := client.DetectUserFaceByUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableGbSubDeviceWithOptions(request *EnableGbSubDeviceRequest, runtime *util.RuntimeOptions) (_result *EnableGbSubDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.SubDeviceId)) {
		query["SubDeviceId"] = request.SubDeviceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableGbSubDevice"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableGbSubDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableGbSubDevice(request *EnableGbSubDeviceRequest) (_result *EnableGbSubDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableGbSubDeviceResponse{}
	_body, _err := client.EnableGbSubDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPictureSearchJobStatusWithOptions(request *GetPictureSearchJobStatusRequest, runtime *util.RuntimeOptions) (_result *GetPictureSearchJobStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceId)) {
		query["AppInstanceId"] = request.AppInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPictureSearchJobStatus"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPictureSearchJobStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPictureSearchJobStatus(request *GetPictureSearchJobStatusRequest) (_result *GetPictureSearchJobStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPictureSearchJobStatusResponse{}
	_body, _err := client.GetPictureSearchJobStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PictureSearchPictureWithOptions(request *PictureSearchPictureRequest, runtime *util.RuntimeOptions) (_result *PictureSearchPictureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceId)) {
		query["AppInstanceId"] = request.AppInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ContainPicUrl)) {
		query["ContainPicUrl"] = request.ContainPicUrl
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PictureSearchType)) {
		query["PictureSearchType"] = request.PictureSearchType
	}

	if !tea.BoolValue(util.IsUnset(request.SearchPicUrl)) {
		query["SearchPicUrl"] = request.SearchPicUrl
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Threshold)) {
		query["Threshold"] = request.Threshold
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PictureSearchPicture"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PictureSearchPictureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PictureSearchPicture(request *PictureSearchPictureRequest) (_result *PictureSearchPictureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PictureSearchPictureResponse{}
	_body, _err := client.PictureSearchPictureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCarProcessEventsWithOptions(request *QueryCarProcessEventsRequest, runtime *util.RuntimeOptions) (_result *QueryCarProcessEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionType)) {
		query["ActionType"] = request.ActionType
	}

	if !tea.BoolValue(util.IsUnset(request.AreaIndex)) {
		query["AreaIndex"] = request.AreaIndex
	}

	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PlateNo)) {
		query["PlateNo"] = request.PlateNo
	}

	if !tea.BoolValue(util.IsUnset(request.SubDeviceName)) {
		query["SubDeviceName"] = request.SubDeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.SubIotId)) {
		query["SubIotId"] = request.SubIotId
	}

	if !tea.BoolValue(util.IsUnset(request.SubProductKey)) {
		query["SubProductKey"] = request.SubProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCarProcessEvents"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCarProcessEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCarProcessEvents(request *QueryCarProcessEventsRequest) (_result *QueryCarProcessEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCarProcessEventsResponse{}
	_body, _err := client.QueryCarProcessEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDeviceEventWithOptions(request *QueryDeviceEventRequest, runtime *util.RuntimeOptions) (_result *QueryDeviceEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EventType)) {
		query["EventType"] = request.EventType
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDeviceEvent"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDeviceEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDeviceEvent(request *QueryDeviceEventRequest) (_result *QueryDeviceEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDeviceEventResponse{}
	_body, _err := client.QueryDeviceEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDeviceEventPictureWithOptions(request *QueryDeviceEventPictureRequest, runtime *util.RuntimeOptions) (_result *QueryDeviceEventPictureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.EventId)) {
		query["EventId"] = request.EventId
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDeviceEventPicture"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDeviceEventPictureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDeviceEventPicture(request *QueryDeviceEventPictureRequest) (_result *QueryDeviceEventPictureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDeviceEventPictureResponse{}
	_body, _err := client.QueryDeviceEventPictureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDeviceEventRecordWithOptions(request *QueryDeviceEventRecordRequest, runtime *util.RuntimeOptions) (_result *QueryDeviceEventRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.EventId)) {
		query["EventId"] = request.EventId
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDeviceEventRecord"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDeviceEventRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDeviceEventRecord(request *QueryDeviceEventRecordRequest) (_result *QueryDeviceEventRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDeviceEventRecordResponse{}
	_body, _err := client.QueryDeviceEventRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDevicePictureByListWithOptions(request *QueryDevicePictureByListRequest, runtime *util.RuntimeOptions) (_result *QueryDevicePictureByListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireTime)) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PictureIdList)) {
		query["PictureIdList"] = request.PictureIdList
	}

	if !tea.BoolValue(util.IsUnset(request.PictureType)) {
		query["PictureType"] = request.PictureType
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.ThumbWidth)) {
		query["ThumbWidth"] = request.ThumbWidth
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDevicePictureByList"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDevicePictureByListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDevicePictureByList(request *QueryDevicePictureByListRequest) (_result *QueryDevicePictureByListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDevicePictureByListResponse{}
	_body, _err := client.QueryDevicePictureByListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDevicePictureFileWithOptions(request *QueryDevicePictureFileRequest, runtime *util.RuntimeOptions) (_result *QueryDevicePictureFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CaptureId)) {
		query["CaptureId"] = request.CaptureId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireTime)) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PictureType)) {
		query["PictureType"] = request.PictureType
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.ThumbWidth)) {
		query["ThumbWidth"] = request.ThumbWidth
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDevicePictureFile"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDevicePictureFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDevicePictureFile(request *QueryDevicePictureFileRequest) (_result *QueryDevicePictureFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDevicePictureFileResponse{}
	_body, _err := client.QueryDevicePictureFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDevicePictureLifeCycleWithOptions(request *QueryDevicePictureLifeCycleRequest, runtime *util.RuntimeOptions) (_result *QueryDevicePictureLifeCycleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDevicePictureLifeCycle"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDevicePictureLifeCycleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDevicePictureLifeCycle(request *QueryDevicePictureLifeCycleRequest) (_result *QueryDevicePictureLifeCycleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDevicePictureLifeCycleResponse{}
	_body, _err := client.QueryDevicePictureLifeCycleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDeviceRecordLifeCycleWithOptions(request *QueryDeviceRecordLifeCycleRequest, runtime *util.RuntimeOptions) (_result *QueryDeviceRecordLifeCycleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceList)) {
		query["DeviceList"] = request.DeviceList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDeviceRecordLifeCycle"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDeviceRecordLifeCycleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDeviceRecordLifeCycle(request *QueryDeviceRecordLifeCycleRequest) (_result *QueryDeviceRecordLifeCycleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDeviceRecordLifeCycleResponse{}
	_body, _err := client.QueryDeviceRecordLifeCycleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDeviceVodUrlWithOptions(request *QueryDeviceVodUrlRequest, runtime *util.RuntimeOptions) (_result *QueryDeviceVodUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.EnableStun)) {
		query["EnableStun"] = request.EnableStun
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptType)) {
		query["EncryptType"] = request.EncryptType
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PlayUnLimited)) {
		query["PlayUnLimited"] = request.PlayUnLimited
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.Scheme)) {
		query["Scheme"] = request.Scheme
	}

	if !tea.BoolValue(util.IsUnset(request.SeekTime)) {
		query["SeekTime"] = request.SeekTime
	}

	if !tea.BoolValue(util.IsUnset(request.ShouldEncrypt)) {
		query["ShouldEncrypt"] = request.ShouldEncrypt
	}

	if !tea.BoolValue(util.IsUnset(request.UrlValidDuration)) {
		query["UrlValidDuration"] = request.UrlValidDuration
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDeviceVodUrl"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDeviceVodUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDeviceVodUrl(request *QueryDeviceVodUrlRequest) (_result *QueryDeviceVodUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDeviceVodUrlResponse{}
	_body, _err := client.QueryDeviceVodUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDeviceVodUrlByTimeWithOptions(request *QueryDeviceVodUrlByTimeRequest, runtime *util.RuntimeOptions) (_result *QueryDeviceVodUrlByTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.EnableStun)) {
		query["EnableStun"] = request.EnableStun
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptType)) {
		query["EncryptType"] = request.EncryptType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PlayUnLimited)) {
		query["PlayUnLimited"] = request.PlayUnLimited
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.Scheme)) {
		query["Scheme"] = request.Scheme
	}

	if !tea.BoolValue(util.IsUnset(request.SeekTime)) {
		query["SeekTime"] = request.SeekTime
	}

	if !tea.BoolValue(util.IsUnset(request.ShouldEncrypt)) {
		query["ShouldEncrypt"] = request.ShouldEncrypt
	}

	if !tea.BoolValue(util.IsUnset(request.UrlValidDuration)) {
		query["UrlValidDuration"] = request.UrlValidDuration
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDeviceVodUrlByTime"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDeviceVodUrlByTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDeviceVodUrlByTime(request *QueryDeviceVodUrlByTimeRequest) (_result *QueryDeviceVodUrlByTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDeviceVodUrlByTimeResponse{}
	_body, _err := client.QueryDeviceVodUrlByTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryEventRecordPlanDetailWithOptions(request *QueryEventRecordPlanDetailRequest, runtime *util.RuntimeOptions) (_result *QueryEventRecordPlanDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryEventRecordPlanDetail"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryEventRecordPlanDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryEventRecordPlanDetail(request *QueryEventRecordPlanDetailRequest) (_result *QueryEventRecordPlanDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryEventRecordPlanDetailResponse{}
	_body, _err := client.QueryEventRecordPlanDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryEventRecordPlanDeviceByDeviceWithOptions(request *QueryEventRecordPlanDeviceByDeviceRequest, runtime *util.RuntimeOptions) (_result *QueryEventRecordPlanDeviceByDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.StreamType)) {
		query["StreamType"] = request.StreamType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryEventRecordPlanDeviceByDevice"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryEventRecordPlanDeviceByDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryEventRecordPlanDeviceByDevice(request *QueryEventRecordPlanDeviceByDeviceRequest) (_result *QueryEventRecordPlanDeviceByDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryEventRecordPlanDeviceByDeviceResponse{}
	_body, _err := client.QueryEventRecordPlanDeviceByDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryEventRecordPlanDeviceByPlanWithOptions(request *QueryEventRecordPlanDeviceByPlanRequest, runtime *util.RuntimeOptions) (_result *QueryEventRecordPlanDeviceByPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryEventRecordPlanDeviceByPlan"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryEventRecordPlanDeviceByPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryEventRecordPlanDeviceByPlan(request *QueryEventRecordPlanDeviceByPlanRequest) (_result *QueryEventRecordPlanDeviceByPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryEventRecordPlanDeviceByPlanResponse{}
	_body, _err := client.QueryEventRecordPlanDeviceByPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryEventRecordPlansWithOptions(request *QueryEventRecordPlansRequest, runtime *util.RuntimeOptions) (_result *QueryEventRecordPlansResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryEventRecordPlans"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryEventRecordPlansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryEventRecordPlans(request *QueryEventRecordPlansRequest) (_result *QueryEventRecordPlansResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryEventRecordPlansResponse{}
	_body, _err := client.QueryEventRecordPlansWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFaceAllDeviceGroupWithOptions(request *QueryFaceAllDeviceGroupRequest, runtime *util.RuntimeOptions) (_result *QueryFaceAllDeviceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryFaceAllDeviceGroup"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryFaceAllDeviceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryFaceAllDeviceGroup(request *QueryFaceAllDeviceGroupRequest) (_result *QueryFaceAllDeviceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryFaceAllDeviceGroupResponse{}
	_body, _err := client.QueryFaceAllDeviceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFaceAllUserGroupWithOptions(request *QueryFaceAllUserGroupRequest, runtime *util.RuntimeOptions) (_result *QueryFaceAllUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryFaceAllUserGroup"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryFaceAllUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryFaceAllUserGroup(request *QueryFaceAllUserGroupRequest) (_result *QueryFaceAllUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryFaceAllUserGroupResponse{}
	_body, _err := client.QueryFaceAllUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFaceAllUserGroupAndDeviceGroupRelationWithOptions(request *QueryFaceAllUserGroupAndDeviceGroupRelationRequest, runtime *util.RuntimeOptions) (_result *QueryFaceAllUserGroupAndDeviceGroupRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryFaceAllUserGroupAndDeviceGroupRelation"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryFaceAllUserGroupAndDeviceGroupRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryFaceAllUserGroupAndDeviceGroupRelation(request *QueryFaceAllUserGroupAndDeviceGroupRelationRequest) (_result *QueryFaceAllUserGroupAndDeviceGroupRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryFaceAllUserGroupAndDeviceGroupRelationResponse{}
	_body, _err := client.QueryFaceAllUserGroupAndDeviceGroupRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFaceAllUserIdsByGroupIdWithOptions(request *QueryFaceAllUserIdsByGroupIdRequest, runtime *util.RuntimeOptions) (_result *QueryFaceAllUserIdsByGroupIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryFaceAllUserIdsByGroupId"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryFaceAllUserIdsByGroupIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryFaceAllUserIdsByGroupId(request *QueryFaceAllUserIdsByGroupIdRequest) (_result *QueryFaceAllUserIdsByGroupIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryFaceAllUserIdsByGroupIdResponse{}
	_body, _err := client.QueryFaceAllUserIdsByGroupIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFaceCustomUserIdByUserIdWithOptions(request *QueryFaceCustomUserIdByUserIdRequest, runtime *util.RuntimeOptions) (_result *QueryFaceCustomUserIdByUserIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryFaceCustomUserIdByUserId"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryFaceCustomUserIdByUserIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryFaceCustomUserIdByUserId(request *QueryFaceCustomUserIdByUserIdRequest) (_result *QueryFaceCustomUserIdByUserIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryFaceCustomUserIdByUserIdResponse{}
	_body, _err := client.QueryFaceCustomUserIdByUserIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFaceDeviceGroupsByDeviceWithOptions(request *QueryFaceDeviceGroupsByDeviceRequest, runtime *util.RuntimeOptions) (_result *QueryFaceDeviceGroupsByDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryFaceDeviceGroupsByDevice"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryFaceDeviceGroupsByDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryFaceDeviceGroupsByDevice(request *QueryFaceDeviceGroupsByDeviceRequest) (_result *QueryFaceDeviceGroupsByDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryFaceDeviceGroupsByDeviceResponse{}
	_body, _err := client.QueryFaceDeviceGroupsByDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFaceUserWithOptions(request *QueryFaceUserRequest, runtime *util.RuntimeOptions) (_result *QueryFaceUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryFaceUser"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryFaceUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryFaceUser(request *QueryFaceUserRequest) (_result *QueryFaceUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryFaceUserResponse{}
	_body, _err := client.QueryFaceUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFaceUserBatchWithOptions(request *QueryFaceUserBatchRequest, runtime *util.RuntimeOptions) (_result *QueryFaceUserBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdList)) {
		query["UserIdList"] = request.UserIdList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryFaceUserBatch"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryFaceUserBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryFaceUserBatch(request *QueryFaceUserBatchRequest) (_result *QueryFaceUserBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryFaceUserBatchResponse{}
	_body, _err := client.QueryFaceUserBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFaceUserByNameWithOptions(request *QueryFaceUserByNameRequest, runtime *util.RuntimeOptions) (_result *QueryFaceUserByNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		query["Params"] = request.Params
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryFaceUserByName"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryFaceUserByNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryFaceUserByName(request *QueryFaceUserByNameRequest) (_result *QueryFaceUserByNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryFaceUserByNameResponse{}
	_body, _err := client.QueryFaceUserByNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFaceUserGroupWithOptions(request *QueryFaceUserGroupRequest, runtime *util.RuntimeOptions) (_result *QueryFaceUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryFaceUserGroup"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryFaceUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryFaceUserGroup(request *QueryFaceUserGroupRequest) (_result *QueryFaceUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryFaceUserGroupResponse{}
	_body, _err := client.QueryFaceUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFaceUserGroupAndDeviceGroupRelationWithOptions(request *QueryFaceUserGroupAndDeviceGroupRelationRequest, runtime *util.RuntimeOptions) (_result *QueryFaceUserGroupAndDeviceGroupRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ControlId)) {
		query["ControlId"] = request.ControlId
	}

	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryFaceUserGroupAndDeviceGroupRelation"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryFaceUserGroupAndDeviceGroupRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryFaceUserGroupAndDeviceGroupRelation(request *QueryFaceUserGroupAndDeviceGroupRelationRequest) (_result *QueryFaceUserGroupAndDeviceGroupRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryFaceUserGroupAndDeviceGroupRelationResponse{}
	_body, _err := client.QueryFaceUserGroupAndDeviceGroupRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFaceUserIdByCustomUserIdWithOptions(request *QueryFaceUserIdByCustomUserIdRequest, runtime *util.RuntimeOptions) (_result *QueryFaceUserIdByCustomUserIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomUserId)) {
		query["CustomUserId"] = request.CustomUserId
	}

	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryFaceUserIdByCustomUserId"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryFaceUserIdByCustomUserIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryFaceUserIdByCustomUserId(request *QueryFaceUserIdByCustomUserIdRequest) (_result *QueryFaceUserIdByCustomUserIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryFaceUserIdByCustomUserIdResponse{}
	_body, _err := client.QueryFaceUserIdByCustomUserIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryGbSubDeviceListWithOptions(request *QueryGbSubDeviceListRequest, runtime *util.RuntimeOptions) (_result *QueryGbSubDeviceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PageStart)) {
		query["PageStart"] = request.PageStart
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryGbSubDeviceList"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGbSubDeviceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryGbSubDeviceList(request *QueryGbSubDeviceListRequest) (_result *QueryGbSubDeviceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryGbSubDeviceListResponse{}
	_body, _err := client.QueryGbSubDeviceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryLiveStreamingWithOptions(request *QueryLiveStreamingRequest, runtime *util.RuntimeOptions) (_result *QueryLiveStreamingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CacheDuration)) {
		query["CacheDuration"] = request.CacheDuration
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.EnableStun)) {
		query["EnableStun"] = request.EnableStun
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptType)) {
		query["EncryptType"] = request.EncryptType
	}

	if !tea.BoolValue(util.IsUnset(request.ForceIFrame)) {
		query["ForceIFrame"] = request.ForceIFrame
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PlayUnLimited)) {
		query["PlayUnLimited"] = request.PlayUnLimited
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.Scheme)) {
		query["Scheme"] = request.Scheme
	}

	if !tea.BoolValue(util.IsUnset(request.ShouldEncrypt)) {
		query["ShouldEncrypt"] = request.ShouldEncrypt
	}

	if !tea.BoolValue(util.IsUnset(request.StreamType)) {
		query["StreamType"] = request.StreamType
	}

	if !tea.BoolValue(util.IsUnset(request.UrlValidDuration)) {
		query["UrlValidDuration"] = request.UrlValidDuration
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryLiveStreaming"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryLiveStreamingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryLiveStreaming(request *QueryLiveStreamingRequest) (_result *QueryLiveStreamingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryLiveStreamingResponse{}
	_body, _err := client.QueryLiveStreamingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryLocalFileUploadJobWithOptions(request *QueryLocalFileUploadJobRequest, runtime *util.RuntimeOptions) (_result *QueryLocalFileUploadJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryLocalFileUploadJob"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryLocalFileUploadJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryLocalFileUploadJob(request *QueryLocalFileUploadJobRequest) (_result *QueryLocalFileUploadJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryLocalFileUploadJobResponse{}
	_body, _err := client.QueryLocalFileUploadJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMonthRecordWithOptions(request *QueryMonthRecordRequest, runtime *util.RuntimeOptions) (_result *QueryMonthRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Month)) {
		query["Month"] = request.Month
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMonthRecord"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMonthRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMonthRecord(request *QueryMonthRecordRequest) (_result *QueryMonthRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMonthRecordResponse{}
	_body, _err := client.QueryMonthRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPictureFilesWithOptions(request *QueryPictureFilesRequest, runtime *util.RuntimeOptions) (_result *QueryPictureFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PictureSource)) {
		query["PictureSource"] = request.PictureSource
	}

	if !tea.BoolValue(util.IsUnset(request.PictureType)) {
		query["PictureType"] = request.PictureType
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPictureFiles"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPictureFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPictureFiles(request *QueryPictureFilesRequest) (_result *QueryPictureFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPictureFilesResponse{}
	_body, _err := client.QueryPictureFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPictureSearchAiboxesWithOptions(request *QueryPictureSearchAiboxesRequest, runtime *util.RuntimeOptions) (_result *QueryPictureSearchAiboxesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceId)) {
		query["AppInstanceId"] = request.AppInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPictureSearchAiboxes"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPictureSearchAiboxesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPictureSearchAiboxes(request *QueryPictureSearchAiboxesRequest) (_result *QueryPictureSearchAiboxesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPictureSearchAiboxesResponse{}
	_body, _err := client.QueryPictureSearchAiboxesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPictureSearchAppsWithOptions(request *QueryPictureSearchAppsRequest, runtime *util.RuntimeOptions) (_result *QueryPictureSearchAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPictureSearchApps"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPictureSearchAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPictureSearchApps(request *QueryPictureSearchAppsRequest) (_result *QueryPictureSearchAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPictureSearchAppsResponse{}
	_body, _err := client.QueryPictureSearchAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPictureSearchDevicesWithOptions(request *QueryPictureSearchDevicesRequest, runtime *util.RuntimeOptions) (_result *QueryPictureSearchDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceId)) {
		query["AppInstanceId"] = request.AppInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPictureSearchDevices"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPictureSearchDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPictureSearchDevices(request *QueryPictureSearchDevicesRequest) (_result *QueryPictureSearchDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPictureSearchDevicesResponse{}
	_body, _err := client.QueryPictureSearchDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPictureSearchJobWithOptions(request *QueryPictureSearchJobRequest, runtime *util.RuntimeOptions) (_result *QueryPictureSearchJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceId)) {
		query["AppInstanceId"] = request.AppInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.JobStatus)) {
		query["JobStatus"] = request.JobStatus
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPictureSearchJob"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPictureSearchJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPictureSearchJob(request *QueryPictureSearchJobRequest) (_result *QueryPictureSearchJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPictureSearchJobResponse{}
	_body, _err := client.QueryPictureSearchJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPictureSearchJobResultWithOptions(request *QueryPictureSearchJobResultRequest, runtime *util.RuntimeOptions) (_result *QueryPictureSearchJobResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceId)) {
		query["AppInstanceId"] = request.AppInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPictureSearchJobResult"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPictureSearchJobResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPictureSearchJobResult(request *QueryPictureSearchJobResultRequest) (_result *QueryPictureSearchJobResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPictureSearchJobResultResponse{}
	_body, _err := client.QueryPictureSearchJobResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRecordWithOptions(request *QueryRecordRequest, runtime *util.RuntimeOptions) (_result *QueryRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NeedSnapshot)) {
		query["NeedSnapshot"] = request.NeedSnapshot
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.RecordType)) {
		query["RecordType"] = request.RecordType
	}

	if !tea.BoolValue(util.IsUnset(request.StreamType)) {
		query["StreamType"] = request.StreamType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRecord"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRecord(request *QueryRecordRequest) (_result *QueryRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRecordResponse{}
	_body, _err := client.QueryRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRecordByRecordIdWithOptions(request *QueryRecordByRecordIdRequest, runtime *util.RuntimeOptions) (_result *QueryRecordByRecordIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.RecordId)) {
		query["RecordId"] = request.RecordId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRecordByRecordId"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRecordByRecordIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRecordByRecordId(request *QueryRecordByRecordIdRequest) (_result *QueryRecordByRecordIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRecordByRecordIdResponse{}
	_body, _err := client.QueryRecordByRecordIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRecordDownloadJobByIdWithOptions(request *QueryRecordDownloadJobByIdRequest, runtime *util.RuntimeOptions) (_result *QueryRecordDownloadJobByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRecordDownloadJobById"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRecordDownloadJobByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRecordDownloadJobById(request *QueryRecordDownloadJobByIdRequest) (_result *QueryRecordDownloadJobByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRecordDownloadJobByIdResponse{}
	_body, _err := client.QueryRecordDownloadJobByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRecordDownloadJobListWithOptions(request *QueryRecordDownloadJobListRequest, runtime *util.RuntimeOptions) (_result *QueryRecordDownloadJobListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRecordDownloadJobList"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRecordDownloadJobListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRecordDownloadJobList(request *QueryRecordDownloadJobListRequest) (_result *QueryRecordDownloadJobListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRecordDownloadJobListResponse{}
	_body, _err := client.QueryRecordDownloadJobListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRecordDownloadUrlWithOptions(request *QueryRecordDownloadUrlRequest, runtime *util.RuntimeOptions) (_result *QueryRecordDownloadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRecordDownloadUrl"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRecordDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRecordDownloadUrl(request *QueryRecordDownloadUrlRequest) (_result *QueryRecordDownloadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRecordDownloadUrlResponse{}
	_body, _err := client.QueryRecordDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRecordPlanDetailWithOptions(request *QueryRecordPlanDetailRequest, runtime *util.RuntimeOptions) (_result *QueryRecordPlanDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRecordPlanDetail"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRecordPlanDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRecordPlanDetail(request *QueryRecordPlanDetailRequest) (_result *QueryRecordPlanDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRecordPlanDetailResponse{}
	_body, _err := client.QueryRecordPlanDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRecordPlanDeviceByDeviceWithOptions(request *QueryRecordPlanDeviceByDeviceRequest, runtime *util.RuntimeOptions) (_result *QueryRecordPlanDeviceByDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.StreamType)) {
		query["StreamType"] = request.StreamType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRecordPlanDeviceByDevice"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRecordPlanDeviceByDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRecordPlanDeviceByDevice(request *QueryRecordPlanDeviceByDeviceRequest) (_result *QueryRecordPlanDeviceByDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRecordPlanDeviceByDeviceResponse{}
	_body, _err := client.QueryRecordPlanDeviceByDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRecordPlanDeviceByPlanWithOptions(request *QueryRecordPlanDeviceByPlanRequest, runtime *util.RuntimeOptions) (_result *QueryRecordPlanDeviceByPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRecordPlanDeviceByPlan"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRecordPlanDeviceByPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRecordPlanDeviceByPlan(request *QueryRecordPlanDeviceByPlanRequest) (_result *QueryRecordPlanDeviceByPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRecordPlanDeviceByPlanResponse{}
	_body, _err := client.QueryRecordPlanDeviceByPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRecordPlansWithOptions(request *QueryRecordPlansRequest, runtime *util.RuntimeOptions) (_result *QueryRecordPlansResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRecordPlans"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRecordPlansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRecordPlans(request *QueryRecordPlansRequest) (_result *QueryRecordPlansResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRecordPlansResponse{}
	_body, _err := client.QueryRecordPlansWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRecordUrlWithOptions(request *QueryRecordUrlRequest, runtime *util.RuntimeOptions) (_result *QueryRecordUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.SeekTime)) {
		query["SeekTime"] = request.SeekTime
	}

	if !tea.BoolValue(util.IsUnset(request.UrlValidDuration)) {
		query["UrlValidDuration"] = request.UrlValidDuration
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRecordUrl"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRecordUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRecordUrl(request *QueryRecordUrlRequest) (_result *QueryRecordUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRecordUrlResponse{}
	_body, _err := client.QueryRecordUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRecordUrlByTimeWithOptions(request *QueryRecordUrlByTimeRequest, runtime *util.RuntimeOptions) (_result *QueryRecordUrlByTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.StreamType)) {
		query["StreamType"] = request.StreamType
	}

	if !tea.BoolValue(util.IsUnset(request.UrlValidDuration)) {
		query["UrlValidDuration"] = request.UrlValidDuration
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRecordUrlByTime"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRecordUrlByTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRecordUrlByTime(request *QueryRecordUrlByTimeRequest) (_result *QueryRecordUrlByTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRecordUrlByTimeResponse{}
	_body, _err := client.QueryRecordUrlByTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRtmpKeyWithOptions(request *QueryRtmpKeyRequest, runtime *util.RuntimeOptions) (_result *QueryRtmpKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRtmpKey"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRtmpKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRtmpKey(request *QueryRtmpKeyRequest) (_result *QueryRtmpKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRtmpKeyResponse{}
	_body, _err := client.QueryRtmpKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryStreamPushJobWithOptions(request *QueryStreamPushJobRequest, runtime *util.RuntimeOptions) (_result *QueryStreamPushJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryStreamPushJob"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryStreamPushJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryStreamPushJob(request *QueryStreamPushJobRequest) (_result *QueryStreamPushJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryStreamPushJobResponse{}
	_body, _err := client.QueryStreamPushJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryStreamPushJobListWithOptions(request *QueryStreamPushJobListRequest, runtime *util.RuntimeOptions) (_result *QueryStreamPushJobListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JobType)) {
		query["JobType"] = request.JobType
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryStreamPushJobList"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryStreamPushJobListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryStreamPushJobList(request *QueryStreamPushJobListRequest) (_result *QueryStreamPushJobListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryStreamPushJobListResponse{}
	_body, _err := client.QueryStreamPushJobListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryStreamSnapshotJobWithOptions(request *QueryStreamSnapshotJobRequest, runtime *util.RuntimeOptions) (_result *QueryStreamSnapshotJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryStreamSnapshotJob"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryStreamSnapshotJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryStreamSnapshotJob(request *QueryStreamSnapshotJobRequest) (_result *QueryStreamSnapshotJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryStreamSnapshotJobResponse{}
	_body, _err := client.QueryStreamSnapshotJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTimeTemplateWithOptions(request *QueryTimeTemplateRequest, runtime *util.RuntimeOptions) (_result *QueryTimeTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTimeTemplate"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTimeTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTimeTemplate(request *QueryTimeTemplateRequest) (_result *QueryTimeTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTimeTemplateResponse{}
	_body, _err := client.QueryTimeTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTimeTemplateDetailWithOptions(request *QueryTimeTemplateDetailRequest, runtime *util.RuntimeOptions) (_result *QueryTimeTemplateDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTimeTemplateDetail"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTimeTemplateDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTimeTemplateDetail(request *QueryTimeTemplateDetailRequest) (_result *QueryTimeTemplateDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTimeTemplateDetailResponse{}
	_body, _err := client.QueryTimeTemplateDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryVisionDeviceInfoWithOptions(request *QueryVisionDeviceInfoRequest, runtime *util.RuntimeOptions) (_result *QueryVisionDeviceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryVisionDeviceInfo"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryVisionDeviceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryVisionDeviceInfo(request *QueryVisionDeviceInfoRequest) (_result *QueryVisionDeviceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryVisionDeviceInfoResponse{}
	_body, _err := client.QueryVisionDeviceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryVoiceIntercomWithOptions(request *QueryVoiceIntercomRequest, runtime *util.RuntimeOptions) (_result *QueryVoiceIntercomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.Scheme)) {
		query["Scheme"] = request.Scheme
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryVoiceIntercom"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryVoiceIntercomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryVoiceIntercom(request *QueryVoiceIntercomRequest) (_result *QueryVoiceIntercomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryVoiceIntercomResponse{}
	_body, _err := client.QueryVoiceIntercomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshGbSubDeviceListWithOptions(request *RefreshGbSubDeviceListRequest, runtime *util.RuntimeOptions) (_result *RefreshGbSubDeviceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshGbSubDeviceList"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshGbSubDeviceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshGbSubDeviceList(request *RefreshGbSubDeviceListRequest) (_result *RefreshGbSubDeviceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshGbSubDeviceListResponse{}
	_body, _err := client.RefreshGbSubDeviceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveFaceDeviceFromDeviceGroupWithOptions(request *RemoveFaceDeviceFromDeviceGroupRequest, runtime *util.RuntimeOptions) (_result *RemoveFaceDeviceFromDeviceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceGroupId)) {
		query["DeviceGroupId"] = request.DeviceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveFaceDeviceFromDeviceGroup"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveFaceDeviceFromDeviceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveFaceDeviceFromDeviceGroup(request *RemoveFaceDeviceFromDeviceGroupRequest) (_result *RemoveFaceDeviceFromDeviceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveFaceDeviceFromDeviceGroupResponse{}
	_body, _err := client.RemoveFaceDeviceFromDeviceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveFaceUserFromUserGroupWithOptions(request *RemoveFaceUserFromUserGroupRequest, runtime *util.RuntimeOptions) (_result *RemoveFaceUserFromUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveFaceUserFromUserGroup"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveFaceUserFromUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveFaceUserFromUserGroup(request *RemoveFaceUserFromUserGroupRequest) (_result *RemoveFaceUserFromUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveFaceUserFromUserGroupResponse{}
	_body, _err := client.RemoveFaceUserFromUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDevicePictureLifeCycleWithOptions(request *SetDevicePictureLifeCycleRequest, runtime *util.RuntimeOptions) (_result *SetDevicePictureLifeCycleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Day)) {
		query["Day"] = request.Day
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDevicePictureLifeCycle"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDevicePictureLifeCycleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDevicePictureLifeCycle(request *SetDevicePictureLifeCycleRequest) (_result *SetDevicePictureLifeCycleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDevicePictureLifeCycleResponse{}
	_body, _err := client.SetDevicePictureLifeCycleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDeviceRecordLifeCycleWithOptions(request *SetDeviceRecordLifeCycleRequest, runtime *util.RuntimeOptions) (_result *SetDeviceRecordLifeCycleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Day)) {
		query["Day"] = request.Day
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDeviceRecordLifeCycle"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDeviceRecordLifeCycleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDeviceRecordLifeCycle(request *SetDeviceRecordLifeCycleRequest) (_result *SetDeviceRecordLifeCycleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDeviceRecordLifeCycleResponse{}
	_body, _err := client.SetDeviceRecordLifeCycleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopLiveStreamingWithOptions(request *StopLiveStreamingRequest, runtime *util.RuntimeOptions) (_result *StopLiveStreamingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.StreamType)) {
		query["StreamType"] = request.StreamType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopLiveStreaming"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopLiveStreamingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopLiveStreaming(request *StopLiveStreamingRequest) (_result *StopLiveStreamingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopLiveStreamingResponse{}
	_body, _err := client.StopLiveStreamingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopTriggeredRecordWithOptions(request *StopTriggeredRecordRequest, runtime *util.RuntimeOptions) (_result *StopTriggeredRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.RecordId)) {
		query["RecordId"] = request.RecordId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopTriggeredRecord"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopTriggeredRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopTriggeredRecord(request *StopTriggeredRecordRequest) (_result *StopTriggeredRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopTriggeredRecordResponse{}
	_body, _err := client.StopTriggeredRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransferDeviceInstanceWithOptions(request *TransferDeviceInstanceRequest, runtime *util.RuntimeOptions) (_result *TransferDeviceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceNameList)) {
		query["DeviceNameList"] = request.DeviceNameList
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.SourceInstanceId)) {
		query["SourceInstanceId"] = request.SourceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetInstanceId)) {
		query["TargetInstanceId"] = request.TargetInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TransferDeviceInstance"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TransferDeviceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferDeviceInstance(request *TransferDeviceInstanceRequest) (_result *TransferDeviceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransferDeviceInstanceResponse{}
	_body, _err := client.TransferDeviceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TriggerCapturePictureWithOptions(request *TriggerCapturePictureRequest, runtime *util.RuntimeOptions) (_result *TriggerCapturePictureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TriggerCapturePicture"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TriggerCapturePictureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TriggerCapturePicture(request *TriggerCapturePictureRequest) (_result *TriggerCapturePictureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TriggerCapturePictureResponse{}
	_body, _err := client.TriggerCapturePictureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TriggerRecordWithOptions(request *TriggerRecordRequest, runtime *util.RuntimeOptions) (_result *TriggerRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PreRecordDuration)) {
		query["PreRecordDuration"] = request.PreRecordDuration
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.RecordDuration)) {
		query["RecordDuration"] = request.RecordDuration
	}

	if !tea.BoolValue(util.IsUnset(request.StreamType)) {
		query["StreamType"] = request.StreamType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TriggerRecord"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TriggerRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TriggerRecord(request *TriggerRecordRequest) (_result *TriggerRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TriggerRecordResponse{}
	_body, _err := client.TriggerRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindPictureSearchAppWithDevicesWithOptions(request *UnbindPictureSearchAppWithDevicesRequest, runtime *util.RuntimeOptions) (_result *UnbindPictureSearchAppWithDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceId)) {
		query["AppInstanceId"] = request.AppInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceIotIds)) {
		query["DeviceIotIds"] = request.DeviceIotIds
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindPictureSearchAppWithDevices"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindPictureSearchAppWithDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindPictureSearchAppWithDevices(request *UnbindPictureSearchAppWithDevicesRequest) (_result *UnbindPictureSearchAppWithDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindPictureSearchAppWithDevicesResponse{}
	_body, _err := client.UnbindPictureSearchAppWithDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEventRecordPlanWithOptions(request *UpdateEventRecordPlanRequest, runtime *util.RuntimeOptions) (_result *UpdateEventRecordPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventTypes)) {
		query["EventTypes"] = request.EventTypes
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.PreRecordDuration)) {
		query["PreRecordDuration"] = request.PreRecordDuration
	}

	if !tea.BoolValue(util.IsUnset(request.RecordDuration)) {
		query["RecordDuration"] = request.RecordDuration
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEventRecordPlan"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEventRecordPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEventRecordPlan(request *UpdateEventRecordPlanRequest) (_result *UpdateEventRecordPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateEventRecordPlanResponse{}
	_body, _err := client.UpdateEventRecordPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFaceUserWithOptions(request *UpdateFaceUserRequest, runtime *util.RuntimeOptions) (_result *UpdateFaceUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomUserId)) {
		query["CustomUserId"] = request.CustomUserId
	}

	if !tea.BoolValue(util.IsUnset(request.FacePicUrl)) {
		query["FacePicUrl"] = request.FacePicUrl
	}

	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		query["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFaceUser"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFaceUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFaceUser(request *UpdateFaceUserRequest) (_result *UpdateFaceUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFaceUserResponse{}
	_body, _err := client.UpdateFaceUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFaceUserGroupAndDeviceGroupRelationWithOptions(request *UpdateFaceUserGroupAndDeviceGroupRelationRequest, runtime *util.RuntimeOptions) (_result *UpdateFaceUserGroupAndDeviceGroupRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ControlId)) {
		query["ControlId"] = request.ControlId
	}

	if !tea.BoolValue(util.IsUnset(request.IsolationId)) {
		query["IsolationId"] = request.IsolationId
	}

	if !tea.BoolValue(util.IsUnset(request.Relation)) {
		query["Relation"] = request.Relation
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFaceUserGroupAndDeviceGroupRelation"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFaceUserGroupAndDeviceGroupRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFaceUserGroupAndDeviceGroupRelation(request *UpdateFaceUserGroupAndDeviceGroupRelationRequest) (_result *UpdateFaceUserGroupAndDeviceGroupRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFaceUserGroupAndDeviceGroupRelationResponse{}
	_body, _err := client.UpdateFaceUserGroupAndDeviceGroupRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGbDeviceWithOptions(request *UpdateGbDeviceRequest, runtime *util.RuntimeOptions) (_result *UpdateGbDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.GbId)) {
		query["GbId"] = request.GbId
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGbDevice"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGbDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGbDevice(request *UpdateGbDeviceRequest) (_result *UpdateGbDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGbDeviceResponse{}
	_body, _err := client.UpdateGbDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInstanceInternetProtocolWithOptions(request *UpdateInstanceInternetProtocolRequest, runtime *util.RuntimeOptions) (_result *UpdateInstanceInternetProtocolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IpVersion)) {
		query["IpVersion"] = request.IpVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceInternetProtocol"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceInternetProtocolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInstanceInternetProtocol(request *UpdateInstanceInternetProtocolRequest) (_result *UpdateInstanceInternetProtocolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateInstanceInternetProtocolResponse{}
	_body, _err := client.UpdateInstanceInternetProtocolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePictureSearchAppWithOptions(request *UpdatePictureSearchAppRequest, runtime *util.RuntimeOptions) (_result *UpdatePictureSearchAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceId)) {
		query["AppInstanceId"] = request.AppInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePictureSearchApp"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePictureSearchAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePictureSearchApp(request *UpdatePictureSearchAppRequest) (_result *UpdatePictureSearchAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePictureSearchAppResponse{}
	_body, _err := client.UpdatePictureSearchAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRecordPlanWithOptions(request *UpdateRecordPlanRequest, runtime *util.RuntimeOptions) (_result *UpdateRecordPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRecordPlan"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRecordPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRecordPlan(request *UpdateRecordPlanRequest) (_result *UpdateRecordPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRecordPlanResponse{}
	_body, _err := client.UpdateRecordPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRtmpKeyWithOptions(request *UpdateRtmpKeyRequest, runtime *util.RuntimeOptions) (_result *UpdateRtmpKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.PullAuthKey)) {
		query["PullAuthKey"] = request.PullAuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.PullKeyExpireTime)) {
		query["PullKeyExpireTime"] = request.PullKeyExpireTime
	}

	if !tea.BoolValue(util.IsUnset(request.PushAuthKey)) {
		query["PushAuthKey"] = request.PushAuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.PushKeyExpireTime)) {
		query["PushKeyExpireTime"] = request.PushKeyExpireTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRtmpKey"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRtmpKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRtmpKey(request *UpdateRtmpKeyRequest) (_result *UpdateRtmpKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRtmpKeyResponse{}
	_body, _err := client.UpdateRtmpKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTimeTemplateWithOptions(request *UpdateTimeTemplateRequest, runtime *util.RuntimeOptions) (_result *UpdateTimeTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllDay)) {
		query["AllDay"] = request.AllDay
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TimeSections)) {
		query["TimeSections"] = request.TimeSections
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTimeTemplate"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTimeTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTimeTemplate(request *UpdateTimeTemplateRequest) (_result *UpdateTimeTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTimeTemplateResponse{}
	_body, _err := client.UpdateTimeTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
