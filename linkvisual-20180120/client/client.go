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

type AddEventRecordPlanDeviceRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId       *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	PlanId      *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	StreamType  *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s AddEventRecordPlanDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddEventRecordPlanDeviceRequest) GoString() string {
	return s.String()
}

func (s *AddEventRecordPlanDeviceRequest) SetApiProduct(v string) *AddEventRecordPlanDeviceRequest {
	s.ApiProduct = &v
	return s
}

func (s *AddEventRecordPlanDeviceRequest) SetApiRevision(v string) *AddEventRecordPlanDeviceRequest {
	s.ApiRevision = &v
	return s
}

func (s *AddEventRecordPlanDeviceRequest) SetIotId(v string) *AddEventRecordPlanDeviceRequest {
	s.IotId = &v
	return s
}

func (s *AddEventRecordPlanDeviceRequest) SetPlanId(v string) *AddEventRecordPlanDeviceRequest {
	s.PlanId = &v
	return s
}

func (s *AddEventRecordPlanDeviceRequest) SetStreamType(v int32) *AddEventRecordPlanDeviceRequest {
	s.StreamType = &v
	return s
}

type AddEventRecordPlanDeviceResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddEventRecordPlanDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddEventRecordPlanDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *AddEventRecordPlanDeviceResponseBody) SetRequestId(v string) *AddEventRecordPlanDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddEventRecordPlanDeviceResponseBody) SetErrorMessage(v string) *AddEventRecordPlanDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddEventRecordPlanDeviceResponseBody) SetCode(v string) *AddEventRecordPlanDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *AddEventRecordPlanDeviceResponseBody) SetSuccess(v bool) *AddEventRecordPlanDeviceResponseBody {
	s.Success = &v
	return s
}

type AddEventRecordPlanDeviceResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddEventRecordPlanDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddEventRecordPlanDeviceResponse) SetBody(v *AddEventRecordPlanDeviceResponseBody) *AddEventRecordPlanDeviceResponse {
	s.Body = v
	return s
}

type AddFaceDeviceGroupRequest struct {
	ApiProduct      *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision     *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId     *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	DeviceGroupName *string `json:"DeviceGroupName,omitempty" xml:"DeviceGroupName,omitempty"`
}

func (s AddFaceDeviceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFaceDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *AddFaceDeviceGroupRequest) SetApiProduct(v string) *AddFaceDeviceGroupRequest {
	s.ApiProduct = &v
	return s
}

func (s *AddFaceDeviceGroupRequest) SetApiRevision(v string) *AddFaceDeviceGroupRequest {
	s.ApiRevision = &v
	return s
}

func (s *AddFaceDeviceGroupRequest) SetIsolationId(v string) *AddFaceDeviceGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *AddFaceDeviceGroupRequest) SetDeviceGroupName(v string) *AddFaceDeviceGroupRequest {
	s.DeviceGroupName = &v
	return s
}

type AddFaceDeviceGroupResponseBody struct {
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *AddFaceDeviceGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddFaceDeviceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFaceDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceDeviceGroupResponseBody) SetRequestId(v string) *AddFaceDeviceGroupResponseBody {
	s.RequestId = &v
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

func (s *AddFaceDeviceGroupResponseBody) SetCode(v string) *AddFaceDeviceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *AddFaceDeviceGroupResponseBody) SetSuccess(v bool) *AddFaceDeviceGroupResponseBody {
	s.Success = &v
	return s
}

type AddFaceDeviceGroupResponseBodyData struct {
	DeviceGroupId   *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	ModifiedTime    *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	DeviceGroupName *string `json:"DeviceGroupName,omitempty" xml:"DeviceGroupName,omitempty"`
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

func (s *AddFaceDeviceGroupResponseBodyData) SetModifiedTime(v string) *AddFaceDeviceGroupResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *AddFaceDeviceGroupResponseBodyData) SetDeviceGroupName(v string) *AddFaceDeviceGroupResponseBodyData {
	s.DeviceGroupName = &v
	return s
}

type AddFaceDeviceGroupResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddFaceDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddFaceDeviceGroupResponse) SetBody(v *AddFaceDeviceGroupResponseBody) *AddFaceDeviceGroupResponse {
	s.Body = v
	return s
}

type AddFaceDeviceToDeviceGroupRequest struct {
	ApiProduct    *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId   *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
}

func (s AddFaceDeviceToDeviceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFaceDeviceToDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *AddFaceDeviceToDeviceGroupRequest) SetApiProduct(v string) *AddFaceDeviceToDeviceGroupRequest {
	s.ApiProduct = &v
	return s
}

func (s *AddFaceDeviceToDeviceGroupRequest) SetApiRevision(v string) *AddFaceDeviceToDeviceGroupRequest {
	s.ApiRevision = &v
	return s
}

func (s *AddFaceDeviceToDeviceGroupRequest) SetIsolationId(v string) *AddFaceDeviceToDeviceGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *AddFaceDeviceToDeviceGroupRequest) SetIotInstanceId(v string) *AddFaceDeviceToDeviceGroupRequest {
	s.IotInstanceId = &v
	return s
}

func (s *AddFaceDeviceToDeviceGroupRequest) SetProductKey(v string) *AddFaceDeviceToDeviceGroupRequest {
	s.ProductKey = &v
	return s
}

func (s *AddFaceDeviceToDeviceGroupRequest) SetDeviceName(v string) *AddFaceDeviceToDeviceGroupRequest {
	s.DeviceName = &v
	return s
}

func (s *AddFaceDeviceToDeviceGroupRequest) SetDeviceGroupId(v string) *AddFaceDeviceToDeviceGroupRequest {
	s.DeviceGroupId = &v
	return s
}

type AddFaceDeviceToDeviceGroupResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddFaceDeviceToDeviceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFaceDeviceToDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceDeviceToDeviceGroupResponseBody) SetRequestId(v string) *AddFaceDeviceToDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFaceDeviceToDeviceGroupResponseBody) SetErrorMessage(v string) *AddFaceDeviceToDeviceGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddFaceDeviceToDeviceGroupResponseBody) SetCode(v string) *AddFaceDeviceToDeviceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *AddFaceDeviceToDeviceGroupResponseBody) SetSuccess(v bool) *AddFaceDeviceToDeviceGroupResponseBody {
	s.Success = &v
	return s
}

type AddFaceDeviceToDeviceGroupResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddFaceDeviceToDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddFaceDeviceToDeviceGroupResponse) SetBody(v *AddFaceDeviceToDeviceGroupResponseBody) *AddFaceDeviceToDeviceGroupResponse {
	s.Body = v
	return s
}

type AddFaceUserRequest struct {
	ApiProduct   *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision  *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId  *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	FacePicUrl   *string `json:"FacePicUrl,omitempty" xml:"FacePicUrl,omitempty"`
	CustomUserId *string `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Params       *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s AddFaceUserRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserRequest) GoString() string {
	return s.String()
}

func (s *AddFaceUserRequest) SetApiProduct(v string) *AddFaceUserRequest {
	s.ApiProduct = &v
	return s
}

func (s *AddFaceUserRequest) SetApiRevision(v string) *AddFaceUserRequest {
	s.ApiRevision = &v
	return s
}

func (s *AddFaceUserRequest) SetIsolationId(v string) *AddFaceUserRequest {
	s.IsolationId = &v
	return s
}

func (s *AddFaceUserRequest) SetFacePicUrl(v string) *AddFaceUserRequest {
	s.FacePicUrl = &v
	return s
}

func (s *AddFaceUserRequest) SetCustomUserId(v string) *AddFaceUserRequest {
	s.CustomUserId = &v
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
	RequestId    *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *AddFaceUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddFaceUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceUserResponseBody) SetRequestId(v string) *AddFaceUserResponseBody {
	s.RequestId = &v
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

func (s *AddFaceUserResponseBody) SetCode(v string) *AddFaceUserResponseBody {
	s.Code = &v
	return s
}

func (s *AddFaceUserResponseBody) SetSuccess(v bool) *AddFaceUserResponseBody {
	s.Success = &v
	return s
}

type AddFaceUserResponseBodyData struct {
	Params       *string `json:"Params,omitempty" xml:"Params,omitempty"`
	CustomUserId *string `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s AddFaceUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddFaceUserResponseBodyData) SetParams(v string) *AddFaceUserResponseBodyData {
	s.Params = &v
	return s
}

func (s *AddFaceUserResponseBodyData) SetCustomUserId(v string) *AddFaceUserResponseBodyData {
	s.CustomUserId = &v
	return s
}

func (s *AddFaceUserResponseBodyData) SetUserId(v string) *AddFaceUserResponseBodyData {
	s.UserId = &v
	return s
}

func (s *AddFaceUserResponseBodyData) SetName(v string) *AddFaceUserResponseBodyData {
	s.Name = &v
	return s
}

type AddFaceUserResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddFaceUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddFaceUserResponse) SetBody(v *AddFaceUserResponseBody) *AddFaceUserResponse {
	s.Body = v
	return s
}

type AddFaceUserGroupRequest struct {
	ApiProduct    *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId   *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s AddFaceUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserGroupRequest) GoString() string {
	return s.String()
}

func (s *AddFaceUserGroupRequest) SetApiProduct(v string) *AddFaceUserGroupRequest {
	s.ApiProduct = &v
	return s
}

func (s *AddFaceUserGroupRequest) SetApiRevision(v string) *AddFaceUserGroupRequest {
	s.ApiRevision = &v
	return s
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
	RequestId    *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *AddFaceUserGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddFaceUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceUserGroupResponseBody) SetRequestId(v string) *AddFaceUserGroupResponseBody {
	s.RequestId = &v
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

func (s *AddFaceUserGroupResponseBody) SetCode(v string) *AddFaceUserGroupResponseBody {
	s.Code = &v
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddFaceUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddFaceUserGroupResponse) SetBody(v *AddFaceUserGroupResponseBody) *AddFaceUserGroupResponse {
	s.Body = v
	return s
}

type AddFaceUserGroupAndDeviceGroupRelationRequest struct {
	ApiProduct    *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId   *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	UserGroupId   *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	Relation      *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
}

func (s AddFaceUserGroupAndDeviceGroupRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserGroupAndDeviceGroupRelationRequest) GoString() string {
	return s.String()
}

func (s *AddFaceUserGroupAndDeviceGroupRelationRequest) SetApiProduct(v string) *AddFaceUserGroupAndDeviceGroupRelationRequest {
	s.ApiProduct = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationRequest) SetApiRevision(v string) *AddFaceUserGroupAndDeviceGroupRelationRequest {
	s.ApiRevision = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationRequest) SetIsolationId(v string) *AddFaceUserGroupAndDeviceGroupRelationRequest {
	s.IsolationId = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationRequest) SetIotInstanceId(v string) *AddFaceUserGroupAndDeviceGroupRelationRequest {
	s.IotInstanceId = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationRequest) SetUserGroupId(v string) *AddFaceUserGroupAndDeviceGroupRelationRequest {
	s.UserGroupId = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationRequest) SetDeviceGroupId(v string) *AddFaceUserGroupAndDeviceGroupRelationRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationRequest) SetRelation(v string) *AddFaceUserGroupAndDeviceGroupRelationRequest {
	s.Relation = &v
	return s
}

type AddFaceUserGroupAndDeviceGroupRelationResponseBody struct {
	RequestId    *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *AddFaceUserGroupAndDeviceGroupRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddFaceUserGroupAndDeviceGroupRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserGroupAndDeviceGroupRelationResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBody) SetRequestId(v string) *AddFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.RequestId = &v
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

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBody) SetCode(v string) *AddFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.Code = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBody) SetSuccess(v bool) *AddFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.Success = &v
	return s
}

type AddFaceUserGroupAndDeviceGroupRelationResponseBodyData struct {
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ControlId    *string `json:"ControlId,omitempty" xml:"ControlId,omitempty"`
}

func (s AddFaceUserGroupAndDeviceGroupRelationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserGroupAndDeviceGroupRelationResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBodyData) SetModifiedTime(v string) *AddFaceUserGroupAndDeviceGroupRelationResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBodyData) SetControlId(v string) *AddFaceUserGroupAndDeviceGroupRelationResponseBodyData {
	s.ControlId = &v
	return s
}

type AddFaceUserGroupAndDeviceGroupRelationResponse struct {
	Headers map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddFaceUserGroupAndDeviceGroupRelationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddFaceUserGroupAndDeviceGroupRelationResponse) SetBody(v *AddFaceUserGroupAndDeviceGroupRelationResponseBody) *AddFaceUserGroupAndDeviceGroupRelationResponse {
	s.Body = v
	return s
}

type AddFaceUserPictureRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	FacePicUrl  *string `json:"FacePicUrl,omitempty" xml:"FacePicUrl,omitempty"`
}

func (s AddFaceUserPictureRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserPictureRequest) GoString() string {
	return s.String()
}

func (s *AddFaceUserPictureRequest) SetApiProduct(v string) *AddFaceUserPictureRequest {
	s.ApiProduct = &v
	return s
}

func (s *AddFaceUserPictureRequest) SetApiRevision(v string) *AddFaceUserPictureRequest {
	s.ApiRevision = &v
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

func (s *AddFaceUserPictureRequest) SetFacePicUrl(v string) *AddFaceUserPictureRequest {
	s.FacePicUrl = &v
	return s
}

type AddFaceUserPictureResponseBody struct {
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddFaceUserPictureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserPictureResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceUserPictureResponseBody) SetRequestId(v string) *AddFaceUserPictureResponseBody {
	s.RequestId = &v
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

func (s *AddFaceUserPictureResponseBody) SetCode(v string) *AddFaceUserPictureResponseBody {
	s.Code = &v
	return s
}

func (s *AddFaceUserPictureResponseBody) SetSuccess(v bool) *AddFaceUserPictureResponseBody {
	s.Success = &v
	return s
}

type AddFaceUserPictureResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddFaceUserPictureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddFaceUserPictureResponse) SetBody(v *AddFaceUserPictureResponseBody) *AddFaceUserPictureResponse {
	s.Body = v
	return s
}

type AddFaceUserToUserGroupRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s AddFaceUserToUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserToUserGroupRequest) GoString() string {
	return s.String()
}

func (s *AddFaceUserToUserGroupRequest) SetApiProduct(v string) *AddFaceUserToUserGroupRequest {
	s.ApiProduct = &v
	return s
}

func (s *AddFaceUserToUserGroupRequest) SetApiRevision(v string) *AddFaceUserToUserGroupRequest {
	s.ApiRevision = &v
	return s
}

func (s *AddFaceUserToUserGroupRequest) SetIsolationId(v string) *AddFaceUserToUserGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *AddFaceUserToUserGroupRequest) SetUserId(v string) *AddFaceUserToUserGroupRequest {
	s.UserId = &v
	return s
}

func (s *AddFaceUserToUserGroupRequest) SetUserGroupId(v string) *AddFaceUserToUserGroupRequest {
	s.UserGroupId = &v
	return s
}

type AddFaceUserToUserGroupResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddFaceUserToUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFaceUserToUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceUserToUserGroupResponseBody) SetRequestId(v string) *AddFaceUserToUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFaceUserToUserGroupResponseBody) SetErrorMessage(v string) *AddFaceUserToUserGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddFaceUserToUserGroupResponseBody) SetCode(v string) *AddFaceUserToUserGroupResponseBody {
	s.Code = &v
	return s
}

func (s *AddFaceUserToUserGroupResponseBody) SetSuccess(v bool) *AddFaceUserToUserGroupResponseBody {
	s.Success = &v
	return s
}

type AddFaceUserToUserGroupResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddFaceUserToUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddFaceUserToUserGroupResponse) SetBody(v *AddFaceUserToUserGroupResponseBody) *AddFaceUserToUserGroupResponse {
	s.Body = v
	return s
}

type AddRecordPlanDeviceRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId       *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	PlanId      *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	StreamType  *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s AddRecordPlanDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRecordPlanDeviceRequest) GoString() string {
	return s.String()
}

func (s *AddRecordPlanDeviceRequest) SetApiProduct(v string) *AddRecordPlanDeviceRequest {
	s.ApiProduct = &v
	return s
}

func (s *AddRecordPlanDeviceRequest) SetApiRevision(v string) *AddRecordPlanDeviceRequest {
	s.ApiRevision = &v
	return s
}

func (s *AddRecordPlanDeviceRequest) SetIotId(v string) *AddRecordPlanDeviceRequest {
	s.IotId = &v
	return s
}

func (s *AddRecordPlanDeviceRequest) SetPlanId(v string) *AddRecordPlanDeviceRequest {
	s.PlanId = &v
	return s
}

func (s *AddRecordPlanDeviceRequest) SetStreamType(v int32) *AddRecordPlanDeviceRequest {
	s.StreamType = &v
	return s
}

type AddRecordPlanDeviceResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddRecordPlanDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddRecordPlanDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *AddRecordPlanDeviceResponseBody) SetRequestId(v string) *AddRecordPlanDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRecordPlanDeviceResponseBody) SetErrorMessage(v string) *AddRecordPlanDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddRecordPlanDeviceResponseBody) SetCode(v string) *AddRecordPlanDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *AddRecordPlanDeviceResponseBody) SetSuccess(v bool) *AddRecordPlanDeviceResponseBody {
	s.Success = &v
	return s
}

type AddRecordPlanDeviceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddRecordPlanDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddRecordPlanDeviceResponse) SetBody(v *AddRecordPlanDeviceResponseBody) *AddRecordPlanDeviceResponse {
	s.Body = v
	return s
}

type BindAIPlanWithDevicesRequest struct {
	ApiProduct  *string   `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string   `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	PlanId      *string   `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	IotIdList   []*string `json:"IotIdList,omitempty" xml:"IotIdList,omitempty" type:"Repeated"`
}

func (s BindAIPlanWithDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s BindAIPlanWithDevicesRequest) GoString() string {
	return s.String()
}

func (s *BindAIPlanWithDevicesRequest) SetApiProduct(v string) *BindAIPlanWithDevicesRequest {
	s.ApiProduct = &v
	return s
}

func (s *BindAIPlanWithDevicesRequest) SetApiRevision(v string) *BindAIPlanWithDevicesRequest {
	s.ApiRevision = &v
	return s
}

func (s *BindAIPlanWithDevicesRequest) SetPlanId(v string) *BindAIPlanWithDevicesRequest {
	s.PlanId = &v
	return s
}

func (s *BindAIPlanWithDevicesRequest) SetIotIdList(v []*string) *BindAIPlanWithDevicesRequest {
	s.IotIdList = v
	return s
}

type BindAIPlanWithDevicesResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s BindAIPlanWithDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindAIPlanWithDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *BindAIPlanWithDevicesResponseBody) SetRequestId(v string) *BindAIPlanWithDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAIPlanWithDevicesResponseBody) SetSuccess(v bool) *BindAIPlanWithDevicesResponseBody {
	s.Success = &v
	return s
}

func (s *BindAIPlanWithDevicesResponseBody) SetErrorMessage(v string) *BindAIPlanWithDevicesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BindAIPlanWithDevicesResponseBody) SetCode(v string) *BindAIPlanWithDevicesResponseBody {
	s.Code = &v
	return s
}

type BindAIPlanWithDevicesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindAIPlanWithDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindAIPlanWithDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s BindAIPlanWithDevicesResponse) GoString() string {
	return s.String()
}

func (s *BindAIPlanWithDevicesResponse) SetHeaders(v map[string]*string) *BindAIPlanWithDevicesResponse {
	s.Headers = v
	return s
}

func (s *BindAIPlanWithDevicesResponse) SetBody(v *BindAIPlanWithDevicesResponseBody) *BindAIPlanWithDevicesResponse {
	s.Body = v
	return s
}

type BindPictureSearchAppWithDevicesRequest struct {
	ApiProduct    *string   `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string   `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	AppInstanceId *string   `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	DeviceIotIds  []*string `json:"DeviceIotIds,omitempty" xml:"DeviceIotIds,omitempty" type:"Repeated"`
}

func (s BindPictureSearchAppWithDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s BindPictureSearchAppWithDevicesRequest) GoString() string {
	return s.String()
}

func (s *BindPictureSearchAppWithDevicesRequest) SetApiProduct(v string) *BindPictureSearchAppWithDevicesRequest {
	s.ApiProduct = &v
	return s
}

func (s *BindPictureSearchAppWithDevicesRequest) SetApiRevision(v string) *BindPictureSearchAppWithDevicesRequest {
	s.ApiRevision = &v
	return s
}

func (s *BindPictureSearchAppWithDevicesRequest) SetAppInstanceId(v string) *BindPictureSearchAppWithDevicesRequest {
	s.AppInstanceId = &v
	return s
}

func (s *BindPictureSearchAppWithDevicesRequest) SetDeviceIotIds(v []*string) *BindPictureSearchAppWithDevicesRequest {
	s.DeviceIotIds = v
	return s
}

type BindPictureSearchAppWithDevicesResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindPictureSearchAppWithDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindPictureSearchAppWithDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *BindPictureSearchAppWithDevicesResponseBody) SetRequestId(v string) *BindPictureSearchAppWithDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindPictureSearchAppWithDevicesResponseBody) SetErrorMessage(v string) *BindPictureSearchAppWithDevicesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BindPictureSearchAppWithDevicesResponseBody) SetCode(v string) *BindPictureSearchAppWithDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *BindPictureSearchAppWithDevicesResponseBody) SetSuccess(v bool) *BindPictureSearchAppWithDevicesResponseBody {
	s.Success = &v
	return s
}

type BindPictureSearchAppWithDevicesResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindPictureSearchAppWithDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *BindPictureSearchAppWithDevicesResponse) SetBody(v *BindPictureSearchAppWithDevicesResponseBody) *BindPictureSearchAppWithDevicesResponse {
	s.Body = v
	return s
}

type CheckFaceUserDoExistOnDeviceRequest struct {
	ApiProduct    *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId   *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
}

func (s CheckFaceUserDoExistOnDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckFaceUserDoExistOnDeviceRequest) GoString() string {
	return s.String()
}

func (s *CheckFaceUserDoExistOnDeviceRequest) SetApiProduct(v string) *CheckFaceUserDoExistOnDeviceRequest {
	s.ApiProduct = &v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceRequest) SetApiRevision(v string) *CheckFaceUserDoExistOnDeviceRequest {
	s.ApiRevision = &v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceRequest) SetIsolationId(v string) *CheckFaceUserDoExistOnDeviceRequest {
	s.IsolationId = &v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceRequest) SetIotInstanceId(v string) *CheckFaceUserDoExistOnDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceRequest) SetUserId(v string) *CheckFaceUserDoExistOnDeviceRequest {
	s.UserId = &v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceRequest) SetProductKey(v string) *CheckFaceUserDoExistOnDeviceRequest {
	s.ProductKey = &v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceRequest) SetDeviceName(v string) *CheckFaceUserDoExistOnDeviceRequest {
	s.DeviceName = &v
	return s
}

type CheckFaceUserDoExistOnDeviceResponseBody struct {
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *CheckFaceUserDoExistOnDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckFaceUserDoExistOnDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckFaceUserDoExistOnDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *CheckFaceUserDoExistOnDeviceResponseBody) SetRequestId(v string) *CheckFaceUserDoExistOnDeviceResponseBody {
	s.RequestId = &v
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

func (s *CheckFaceUserDoExistOnDeviceResponseBody) SetCode(v string) *CheckFaceUserDoExistOnDeviceResponseBody {
	s.Code = &v
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
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckFaceUserDoExistOnDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CheckFaceUserDoExistOnDeviceResponse) SetBody(v *CheckFaceUserDoExistOnDeviceResponseBody) *CheckFaceUserDoExistOnDeviceResponse {
	s.Body = v
	return s
}

type ClearFaceDeviceDBRequest struct {
	ApiProduct    *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId   *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
}

func (s ClearFaceDeviceDBRequest) String() string {
	return tea.Prettify(s)
}

func (s ClearFaceDeviceDBRequest) GoString() string {
	return s.String()
}

func (s *ClearFaceDeviceDBRequest) SetApiProduct(v string) *ClearFaceDeviceDBRequest {
	s.ApiProduct = &v
	return s
}

func (s *ClearFaceDeviceDBRequest) SetApiRevision(v string) *ClearFaceDeviceDBRequest {
	s.ApiRevision = &v
	return s
}

func (s *ClearFaceDeviceDBRequest) SetIsolationId(v string) *ClearFaceDeviceDBRequest {
	s.IsolationId = &v
	return s
}

func (s *ClearFaceDeviceDBRequest) SetIotInstanceId(v string) *ClearFaceDeviceDBRequest {
	s.IotInstanceId = &v
	return s
}

func (s *ClearFaceDeviceDBRequest) SetProductKey(v string) *ClearFaceDeviceDBRequest {
	s.ProductKey = &v
	return s
}

func (s *ClearFaceDeviceDBRequest) SetDeviceName(v string) *ClearFaceDeviceDBRequest {
	s.DeviceName = &v
	return s
}

type ClearFaceDeviceDBResponseBody struct {
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ClearFaceDeviceDBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClearFaceDeviceDBResponseBody) GoString() string {
	return s.String()
}

func (s *ClearFaceDeviceDBResponseBody) SetRequestId(v string) *ClearFaceDeviceDBResponseBody {
	s.RequestId = &v
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

func (s *ClearFaceDeviceDBResponseBody) SetCode(v string) *ClearFaceDeviceDBResponseBody {
	s.Code = &v
	return s
}

func (s *ClearFaceDeviceDBResponseBody) SetSuccess(v bool) *ClearFaceDeviceDBResponseBody {
	s.Success = &v
	return s
}

type ClearFaceDeviceDBResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ClearFaceDeviceDBResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ClearFaceDeviceDBResponse) SetBody(v *ClearFaceDeviceDBResponseBody) *ClearFaceDeviceDBResponse {
	s.Body = v
	return s
}

type ConfigAIActionRequest struct {
	ApiProduct         *string                                    `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision        *string                                    `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	ActionId           *string                                    `json:"ActionId,omitempty" xml:"ActionId,omitempty"`
	DataTypeConfigList []*ConfigAIActionRequestDataTypeConfigList `json:"DataTypeConfigList,omitempty" xml:"DataTypeConfigList,omitempty" type:"Repeated"`
	AlgoConfig         *string                                    `json:"AlgoConfig,omitempty" xml:"AlgoConfig,omitempty"`
}

func (s ConfigAIActionRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigAIActionRequest) GoString() string {
	return s.String()
}

func (s *ConfigAIActionRequest) SetApiProduct(v string) *ConfigAIActionRequest {
	s.ApiProduct = &v
	return s
}

func (s *ConfigAIActionRequest) SetApiRevision(v string) *ConfigAIActionRequest {
	s.ApiRevision = &v
	return s
}

func (s *ConfigAIActionRequest) SetActionId(v string) *ConfigAIActionRequest {
	s.ActionId = &v
	return s
}

func (s *ConfigAIActionRequest) SetDataTypeConfigList(v []*ConfigAIActionRequestDataTypeConfigList) *ConfigAIActionRequest {
	s.DataTypeConfigList = v
	return s
}

func (s *ConfigAIActionRequest) SetAlgoConfig(v string) *ConfigAIActionRequest {
	s.AlgoConfig = &v
	return s
}

type ConfigAIActionRequestDataTypeConfigList struct {
	Configs  *string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
}

func (s ConfigAIActionRequestDataTypeConfigList) String() string {
	return tea.Prettify(s)
}

func (s ConfigAIActionRequestDataTypeConfigList) GoString() string {
	return s.String()
}

func (s *ConfigAIActionRequestDataTypeConfigList) SetConfigs(v string) *ConfigAIActionRequestDataTypeConfigList {
	s.Configs = &v
	return s
}

func (s *ConfigAIActionRequestDataTypeConfigList) SetDataType(v string) *ConfigAIActionRequestDataTypeConfigList {
	s.DataType = &v
	return s
}

type ConfigAIActionResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ConfigAIActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigAIActionResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigAIActionResponseBody) SetRequestId(v string) *ConfigAIActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigAIActionResponseBody) SetSuccess(v bool) *ConfigAIActionResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigAIActionResponseBody) SetErrorMessage(v string) *ConfigAIActionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ConfigAIActionResponseBody) SetCode(v string) *ConfigAIActionResponseBody {
	s.Code = &v
	return s
}

type ConfigAIActionResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigAIActionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigAIActionResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigAIActionResponse) GoString() string {
	return s.String()
}

func (s *ConfigAIActionResponse) SetHeaders(v map[string]*string) *ConfigAIActionResponse {
	s.Headers = v
	return s
}

func (s *ConfigAIActionResponse) SetBody(v *ConfigAIActionResponseBody) *ConfigAIActionResponse {
	s.Body = v
	return s
}

type CreateAIAppRequest struct {
	ApiProduct         *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision        *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	AppTemplateId      *string `json:"AppTemplateId,omitempty" xml:"AppTemplateId,omitempty"`
	AppTemplateVersion *string `json:"AppTemplateVersion,omitempty" xml:"AppTemplateVersion,omitempty"`
	Level              *int32  `json:"Level,omitempty" xml:"Level,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateAIAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAIAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAIAppRequest) SetApiProduct(v string) *CreateAIAppRequest {
	s.ApiProduct = &v
	return s
}

func (s *CreateAIAppRequest) SetApiRevision(v string) *CreateAIAppRequest {
	s.ApiRevision = &v
	return s
}

func (s *CreateAIAppRequest) SetAppTemplateId(v string) *CreateAIAppRequest {
	s.AppTemplateId = &v
	return s
}

func (s *CreateAIAppRequest) SetAppTemplateVersion(v string) *CreateAIAppRequest {
	s.AppTemplateVersion = &v
	return s
}

func (s *CreateAIAppRequest) SetLevel(v int32) *CreateAIAppRequest {
	s.Level = &v
	return s
}

func (s *CreateAIAppRequest) SetName(v string) *CreateAIAppRequest {
	s.Name = &v
	return s
}

func (s *CreateAIAppRequest) SetDescription(v string) *CreateAIAppRequest {
	s.Description = &v
	return s
}

type CreateAIAppResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAIAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAIAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAIAppResponseBody) SetRequestId(v string) *CreateAIAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAIAppResponseBody) SetData(v string) *CreateAIAppResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAIAppResponseBody) SetErrorMessage(v string) *CreateAIAppResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateAIAppResponseBody) SetCode(v string) *CreateAIAppResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAIAppResponseBody) SetSuccess(v bool) *CreateAIAppResponseBody {
	s.Success = &v
	return s
}

type CreateAIAppResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAIAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAIAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAIAppResponse) GoString() string {
	return s.String()
}

func (s *CreateAIAppResponse) SetHeaders(v map[string]*string) *CreateAIAppResponse {
	s.Headers = v
	return s
}

func (s *CreateAIAppResponse) SetBody(v *CreateAIAppResponseBody) *CreateAIAppResponse {
	s.Body = v
	return s
}

type CreateAIPlanRequest struct {
	ApiProduct     *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision    *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PlanTemplateId *string `json:"PlanTemplateId,omitempty" xml:"PlanTemplateId,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateAIPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAIPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateAIPlanRequest) SetApiProduct(v string) *CreateAIPlanRequest {
	s.ApiProduct = &v
	return s
}

func (s *CreateAIPlanRequest) SetApiRevision(v string) *CreateAIPlanRequest {
	s.ApiRevision = &v
	return s
}

func (s *CreateAIPlanRequest) SetAppId(v string) *CreateAIPlanRequest {
	s.AppId = &v
	return s
}

func (s *CreateAIPlanRequest) SetPlanTemplateId(v string) *CreateAIPlanRequest {
	s.PlanTemplateId = &v
	return s
}

func (s *CreateAIPlanRequest) SetDescription(v string) *CreateAIPlanRequest {
	s.Description = &v
	return s
}

type CreateAIPlanResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAIPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAIPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAIPlanResponseBody) SetRequestId(v string) *CreateAIPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAIPlanResponseBody) SetData(v string) *CreateAIPlanResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAIPlanResponseBody) SetErrorMessage(v string) *CreateAIPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateAIPlanResponseBody) SetCode(v string) *CreateAIPlanResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAIPlanResponseBody) SetSuccess(v bool) *CreateAIPlanResponseBody {
	s.Success = &v
	return s
}

type CreateAIPlanResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAIPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAIPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAIPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateAIPlanResponse) SetHeaders(v map[string]*string) *CreateAIPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateAIPlanResponse) SetBody(v *CreateAIPlanResponseBody) *CreateAIPlanResponse {
	s.Body = v
	return s
}

type CreateAlgorithmRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateAlgorithmRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAlgorithmRequest) GoString() string {
	return s.String()
}

func (s *CreateAlgorithmRequest) SetApiProduct(v string) *CreateAlgorithmRequest {
	s.ApiProduct = &v
	return s
}

func (s *CreateAlgorithmRequest) SetApiRevision(v string) *CreateAlgorithmRequest {
	s.ApiRevision = &v
	return s
}

func (s *CreateAlgorithmRequest) SetName(v string) *CreateAlgorithmRequest {
	s.Name = &v
	return s
}

func (s *CreateAlgorithmRequest) SetDescription(v string) *CreateAlgorithmRequest {
	s.Description = &v
	return s
}

type CreateAlgorithmResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAlgorithmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAlgorithmResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlgorithmResponseBody) SetRequestId(v string) *CreateAlgorithmResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAlgorithmResponseBody) SetData(v int64) *CreateAlgorithmResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAlgorithmResponseBody) SetErrorMessage(v string) *CreateAlgorithmResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateAlgorithmResponseBody) SetCode(v string) *CreateAlgorithmResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAlgorithmResponseBody) SetSuccess(v bool) *CreateAlgorithmResponseBody {
	s.Success = &v
	return s
}

type CreateAlgorithmResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAlgorithmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAlgorithmResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAlgorithmResponse) GoString() string {
	return s.String()
}

func (s *CreateAlgorithmResponse) SetHeaders(v map[string]*string) *CreateAlgorithmResponse {
	s.Headers = v
	return s
}

func (s *CreateAlgorithmResponse) SetBody(v *CreateAlgorithmResponseBody) *CreateAlgorithmResponse {
	s.Body = v
	return s
}

type CreateDevicePurifyJobRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId       *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s CreateDevicePurifyJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDevicePurifyJobRequest) GoString() string {
	return s.String()
}

func (s *CreateDevicePurifyJobRequest) SetApiProduct(v string) *CreateDevicePurifyJobRequest {
	s.ApiProduct = &v
	return s
}

func (s *CreateDevicePurifyJobRequest) SetApiRevision(v string) *CreateDevicePurifyJobRequest {
	s.ApiRevision = &v
	return s
}

func (s *CreateDevicePurifyJobRequest) SetIotId(v string) *CreateDevicePurifyJobRequest {
	s.IotId = &v
	return s
}

func (s *CreateDevicePurifyJobRequest) SetStartTime(v int64) *CreateDevicePurifyJobRequest {
	s.StartTime = &v
	return s
}

func (s *CreateDevicePurifyJobRequest) SetEndTime(v int64) *CreateDevicePurifyJobRequest {
	s.EndTime = &v
	return s
}

type CreateDevicePurifyJobResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDevicePurifyJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDevicePurifyJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDevicePurifyJobResponseBody) SetRequestId(v string) *CreateDevicePurifyJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDevicePurifyJobResponseBody) SetData(v string) *CreateDevicePurifyJobResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDevicePurifyJobResponseBody) SetErrorMessage(v string) *CreateDevicePurifyJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDevicePurifyJobResponseBody) SetCode(v string) *CreateDevicePurifyJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDevicePurifyJobResponseBody) SetSuccess(v bool) *CreateDevicePurifyJobResponseBody {
	s.Success = &v
	return s
}

type CreateDevicePurifyJobResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDevicePurifyJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDevicePurifyJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDevicePurifyJobResponse) GoString() string {
	return s.String()
}

func (s *CreateDevicePurifyJobResponse) SetHeaders(v map[string]*string) *CreateDevicePurifyJobResponse {
	s.Headers = v
	return s
}

func (s *CreateDevicePurifyJobResponse) SetBody(v *CreateDevicePurifyJobResponseBody) *CreateDevicePurifyJobResponse {
	s.Body = v
	return s
}

type CreateDevicePurifyJobByPlanIdRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	PlanId      *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	Utc         *int64  `json:"Utc,omitempty" xml:"Utc,omitempty"`
}

func (s CreateDevicePurifyJobByPlanIdRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDevicePurifyJobByPlanIdRequest) GoString() string {
	return s.String()
}

func (s *CreateDevicePurifyJobByPlanIdRequest) SetApiProduct(v string) *CreateDevicePurifyJobByPlanIdRequest {
	s.ApiProduct = &v
	return s
}

func (s *CreateDevicePurifyJobByPlanIdRequest) SetApiRevision(v string) *CreateDevicePurifyJobByPlanIdRequest {
	s.ApiRevision = &v
	return s
}

func (s *CreateDevicePurifyJobByPlanIdRequest) SetPlanId(v string) *CreateDevicePurifyJobByPlanIdRequest {
	s.PlanId = &v
	return s
}

func (s *CreateDevicePurifyJobByPlanIdRequest) SetUtc(v int64) *CreateDevicePurifyJobByPlanIdRequest {
	s.Utc = &v
	return s
}

type CreateDevicePurifyJobByPlanIdResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDevicePurifyJobByPlanIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDevicePurifyJobByPlanIdResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDevicePurifyJobByPlanIdResponseBody) SetRequestId(v string) *CreateDevicePurifyJobByPlanIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDevicePurifyJobByPlanIdResponseBody) SetData(v string) *CreateDevicePurifyJobByPlanIdResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDevicePurifyJobByPlanIdResponseBody) SetErrorMessage(v string) *CreateDevicePurifyJobByPlanIdResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDevicePurifyJobByPlanIdResponseBody) SetCode(v string) *CreateDevicePurifyJobByPlanIdResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDevicePurifyJobByPlanIdResponseBody) SetSuccess(v bool) *CreateDevicePurifyJobByPlanIdResponseBody {
	s.Success = &v
	return s
}

type CreateDevicePurifyJobByPlanIdResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDevicePurifyJobByPlanIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDevicePurifyJobByPlanIdResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDevicePurifyJobByPlanIdResponse) GoString() string {
	return s.String()
}

func (s *CreateDevicePurifyJobByPlanIdResponse) SetHeaders(v map[string]*string) *CreateDevicePurifyJobByPlanIdResponse {
	s.Headers = v
	return s
}

func (s *CreateDevicePurifyJobByPlanIdResponse) SetBody(v *CreateDevicePurifyJobByPlanIdResponseBody) *CreateDevicePurifyJobByPlanIdResponse {
	s.Body = v
	return s
}

type CreateDevicePurifyPlanRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId       *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	StartTime   *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s CreateDevicePurifyPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDevicePurifyPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateDevicePurifyPlanRequest) SetApiProduct(v string) *CreateDevicePurifyPlanRequest {
	s.ApiProduct = &v
	return s
}

func (s *CreateDevicePurifyPlanRequest) SetApiRevision(v string) *CreateDevicePurifyPlanRequest {
	s.ApiRevision = &v
	return s
}

func (s *CreateDevicePurifyPlanRequest) SetIotId(v string) *CreateDevicePurifyPlanRequest {
	s.IotId = &v
	return s
}

func (s *CreateDevicePurifyPlanRequest) SetStartTime(v int32) *CreateDevicePurifyPlanRequest {
	s.StartTime = &v
	return s
}

func (s *CreateDevicePurifyPlanRequest) SetEndTime(v int32) *CreateDevicePurifyPlanRequest {
	s.EndTime = &v
	return s
}

type CreateDevicePurifyPlanResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDevicePurifyPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDevicePurifyPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDevicePurifyPlanResponseBody) SetRequestId(v string) *CreateDevicePurifyPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDevicePurifyPlanResponseBody) SetData(v string) *CreateDevicePurifyPlanResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDevicePurifyPlanResponseBody) SetErrorMessage(v string) *CreateDevicePurifyPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDevicePurifyPlanResponseBody) SetCode(v string) *CreateDevicePurifyPlanResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDevicePurifyPlanResponseBody) SetSuccess(v bool) *CreateDevicePurifyPlanResponseBody {
	s.Success = &v
	return s
}

type CreateDevicePurifyPlanResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDevicePurifyPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDevicePurifyPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDevicePurifyPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateDevicePurifyPlanResponse) SetHeaders(v map[string]*string) *CreateDevicePurifyPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateDevicePurifyPlanResponse) SetBody(v *CreateDevicePurifyPlanResponseBody) *CreateDevicePurifyPlanResponse {
	s.Body = v
	return s
}

type CreateEventRecordPlanRequest struct {
	ApiProduct        *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision       *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	EventTypes        *string `json:"EventTypes,omitempty" xml:"EventTypes,omitempty"`
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

func (s *CreateEventRecordPlanRequest) SetApiProduct(v string) *CreateEventRecordPlanRequest {
	s.ApiProduct = &v
	return s
}

func (s *CreateEventRecordPlanRequest) SetApiRevision(v string) *CreateEventRecordPlanRequest {
	s.ApiRevision = &v
	return s
}

func (s *CreateEventRecordPlanRequest) SetName(v string) *CreateEventRecordPlanRequest {
	s.Name = &v
	return s
}

func (s *CreateEventRecordPlanRequest) SetEventTypes(v string) *CreateEventRecordPlanRequest {
	s.EventTypes = &v
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
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateEventRecordPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEventRecordPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEventRecordPlanResponseBody) SetRequestId(v string) *CreateEventRecordPlanResponseBody {
	s.RequestId = &v
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

func (s *CreateEventRecordPlanResponseBody) SetCode(v string) *CreateEventRecordPlanResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEventRecordPlanResponseBody) SetSuccess(v bool) *CreateEventRecordPlanResponseBody {
	s.Success = &v
	return s
}

type CreateEventRecordPlanResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateEventRecordPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateEventRecordPlanResponse) SetBody(v *CreateEventRecordPlanResponseBody) *CreateEventRecordPlanResponse {
	s.Body = v
	return s
}

type CreateModelRequest struct {
	ApiProduct           *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision          *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	AlgorithmId          *int64  `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ModelPackageStandard *string `json:"ModelPackageStandard,omitempty" xml:"ModelPackageStandard,omitempty"`
	Hardware             *string `json:"Hardware,omitempty" xml:"Hardware,omitempty"`
	UploadModelPath      *string `json:"UploadModelPath,omitempty" xml:"UploadModelPath,omitempty"`
	NeedEncrypt          *bool   `json:"NeedEncrypt,omitempty" xml:"NeedEncrypt,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateModelRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateModelRequest) GoString() string {
	return s.String()
}

func (s *CreateModelRequest) SetApiProduct(v string) *CreateModelRequest {
	s.ApiProduct = &v
	return s
}

func (s *CreateModelRequest) SetApiRevision(v string) *CreateModelRequest {
	s.ApiRevision = &v
	return s
}

func (s *CreateModelRequest) SetAlgorithmId(v int64) *CreateModelRequest {
	s.AlgorithmId = &v
	return s
}

func (s *CreateModelRequest) SetName(v string) *CreateModelRequest {
	s.Name = &v
	return s
}

func (s *CreateModelRequest) SetModelPackageStandard(v string) *CreateModelRequest {
	s.ModelPackageStandard = &v
	return s
}

func (s *CreateModelRequest) SetHardware(v string) *CreateModelRequest {
	s.Hardware = &v
	return s
}

func (s *CreateModelRequest) SetUploadModelPath(v string) *CreateModelRequest {
	s.UploadModelPath = &v
	return s
}

func (s *CreateModelRequest) SetNeedEncrypt(v bool) *CreateModelRequest {
	s.NeedEncrypt = &v
	return s
}

func (s *CreateModelRequest) SetDescription(v string) *CreateModelRequest {
	s.Description = &v
	return s
}

type CreateModelResponseBody struct {
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateModelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelResponseBody) SetRequestId(v string) *CreateModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModelResponseBody) SetData(v map[string]interface{}) *CreateModelResponseBody {
	s.Data = v
	return s
}

func (s *CreateModelResponseBody) SetErrorMessage(v string) *CreateModelResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateModelResponseBody) SetCode(v string) *CreateModelResponseBody {
	s.Code = &v
	return s
}

func (s *CreateModelResponseBody) SetSuccess(v bool) *CreateModelResponseBody {
	s.Success = &v
	return s
}

type CreateModelResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateModelResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateModelResponse) GoString() string {
	return s.String()
}

func (s *CreateModelResponse) SetHeaders(v map[string]*string) *CreateModelResponse {
	s.Headers = v
	return s
}

func (s *CreateModelResponse) SetBody(v *CreateModelResponseBody) *CreateModelResponse {
	s.Body = v
	return s
}

type CreatePictureSearchAppRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Desc        *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
}

func (s CreatePictureSearchAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePictureSearchAppRequest) GoString() string {
	return s.String()
}

func (s *CreatePictureSearchAppRequest) SetApiProduct(v string) *CreatePictureSearchAppRequest {
	s.ApiProduct = &v
	return s
}

func (s *CreatePictureSearchAppRequest) SetApiRevision(v string) *CreatePictureSearchAppRequest {
	s.ApiRevision = &v
	return s
}

func (s *CreatePictureSearchAppRequest) SetName(v string) *CreatePictureSearchAppRequest {
	s.Name = &v
	return s
}

func (s *CreatePictureSearchAppRequest) SetDesc(v string) *CreatePictureSearchAppRequest {
	s.Desc = &v
	return s
}

type CreatePictureSearchAppResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePictureSearchAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePictureSearchAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePictureSearchAppResponseBody) SetRequestId(v string) *CreatePictureSearchAppResponseBody {
	s.RequestId = &v
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

func (s *CreatePictureSearchAppResponseBody) SetCode(v string) *CreatePictureSearchAppResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePictureSearchAppResponseBody) SetSuccess(v bool) *CreatePictureSearchAppResponseBody {
	s.Success = &v
	return s
}

type CreatePictureSearchAppResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreatePictureSearchAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreatePictureSearchAppResponse) SetBody(v *CreatePictureSearchAppResponseBody) *CreatePictureSearchAppResponse {
	s.Body = v
	return s
}

type CreateRecordPlanRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateId  *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateRecordPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRecordPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateRecordPlanRequest) SetApiProduct(v string) *CreateRecordPlanRequest {
	s.ApiProduct = &v
	return s
}

func (s *CreateRecordPlanRequest) SetApiRevision(v string) *CreateRecordPlanRequest {
	s.ApiRevision = &v
	return s
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
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRecordPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRecordPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRecordPlanResponseBody) SetRequestId(v string) *CreateRecordPlanResponseBody {
	s.RequestId = &v
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

func (s *CreateRecordPlanResponseBody) SetCode(v string) *CreateRecordPlanResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRecordPlanResponseBody) SetSuccess(v bool) *CreateRecordPlanResponseBody {
	s.Success = &v
	return s
}

type CreateRecordPlanResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRecordPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateRecordPlanResponse) SetBody(v *CreateRecordPlanResponseBody) *CreateRecordPlanResponse {
	s.Body = v
	return s
}

type CreateTimeTemplateRequest struct {
	ApiProduct   *string                                  `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision  *string                                  `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	Name         *string                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	AllDay       *int32                                   `json:"AllDay,omitempty" xml:"AllDay,omitempty"`
	TimeSections []*CreateTimeTemplateRequestTimeSections `json:"TimeSections,omitempty" xml:"TimeSections,omitempty" type:"Repeated"`
}

func (s CreateTimeTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTimeTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTimeTemplateRequest) SetApiProduct(v string) *CreateTimeTemplateRequest {
	s.ApiProduct = &v
	return s
}

func (s *CreateTimeTemplateRequest) SetApiRevision(v string) *CreateTimeTemplateRequest {
	s.ApiRevision = &v
	return s
}

func (s *CreateTimeTemplateRequest) SetName(v string) *CreateTimeTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateTimeTemplateRequest) SetAllDay(v int32) *CreateTimeTemplateRequest {
	s.AllDay = &v
	return s
}

func (s *CreateTimeTemplateRequest) SetTimeSections(v []*CreateTimeTemplateRequestTimeSections) *CreateTimeTemplateRequest {
	s.TimeSections = v
	return s
}

type CreateTimeTemplateRequestTimeSections struct {
	DayOfWeek *int32 `json:"DayOfWeek,omitempty" xml:"DayOfWeek,omitempty"`
	Begin     *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	End       *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s CreateTimeTemplateRequestTimeSections) String() string {
	return tea.Prettify(s)
}

func (s CreateTimeTemplateRequestTimeSections) GoString() string {
	return s.String()
}

func (s *CreateTimeTemplateRequestTimeSections) SetDayOfWeek(v int32) *CreateTimeTemplateRequestTimeSections {
	s.DayOfWeek = &v
	return s
}

func (s *CreateTimeTemplateRequestTimeSections) SetBegin(v int32) *CreateTimeTemplateRequestTimeSections {
	s.Begin = &v
	return s
}

func (s *CreateTimeTemplateRequestTimeSections) SetEnd(v int32) *CreateTimeTemplateRequestTimeSections {
	s.End = &v
	return s
}

type CreateTimeTemplateResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTimeTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTimeTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTimeTemplateResponseBody) SetRequestId(v string) *CreateTimeTemplateResponseBody {
	s.RequestId = &v
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

func (s *CreateTimeTemplateResponseBody) SetCode(v string) *CreateTimeTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTimeTemplateResponseBody) SetSuccess(v bool) *CreateTimeTemplateResponseBody {
	s.Success = &v
	return s
}

type CreateTimeTemplateResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTimeTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateTimeTemplateResponse) SetBody(v *CreateTimeTemplateResponseBody) *CreateTimeTemplateResponse {
	s.Body = v
	return s
}

type DeleteAlgorithmRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	AlgorithmId *int64  `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
}

func (s DeleteAlgorithmRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlgorithmRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlgorithmRequest) SetApiProduct(v string) *DeleteAlgorithmRequest {
	s.ApiProduct = &v
	return s
}

func (s *DeleteAlgorithmRequest) SetApiRevision(v string) *DeleteAlgorithmRequest {
	s.ApiRevision = &v
	return s
}

func (s *DeleteAlgorithmRequest) SetAlgorithmId(v int64) *DeleteAlgorithmRequest {
	s.AlgorithmId = &v
	return s
}

type DeleteAlgorithmResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAlgorithmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlgorithmResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlgorithmResponseBody) SetRequestId(v string) *DeleteAlgorithmResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAlgorithmResponseBody) SetErrorMessage(v string) *DeleteAlgorithmResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteAlgorithmResponseBody) SetCode(v string) *DeleteAlgorithmResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAlgorithmResponseBody) SetSuccess(v bool) *DeleteAlgorithmResponseBody {
	s.Success = &v
	return s
}

type DeleteAlgorithmResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAlgorithmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAlgorithmResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlgorithmResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlgorithmResponse) SetHeaders(v map[string]*string) *DeleteAlgorithmResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlgorithmResponse) SetBody(v *DeleteAlgorithmResponseBody) *DeleteAlgorithmResponse {
	s.Body = v
	return s
}

type DeleteEventRecordPlanRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	PlanId      *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s DeleteEventRecordPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventRecordPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventRecordPlanRequest) SetApiProduct(v string) *DeleteEventRecordPlanRequest {
	s.ApiProduct = &v
	return s
}

func (s *DeleteEventRecordPlanRequest) SetApiRevision(v string) *DeleteEventRecordPlanRequest {
	s.ApiRevision = &v
	return s
}

func (s *DeleteEventRecordPlanRequest) SetPlanId(v string) *DeleteEventRecordPlanRequest {
	s.PlanId = &v
	return s
}

type DeleteEventRecordPlanResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteEventRecordPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventRecordPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventRecordPlanResponseBody) SetRequestId(v string) *DeleteEventRecordPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEventRecordPlanResponseBody) SetErrorMessage(v string) *DeleteEventRecordPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteEventRecordPlanResponseBody) SetCode(v string) *DeleteEventRecordPlanResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEventRecordPlanResponseBody) SetSuccess(v bool) *DeleteEventRecordPlanResponseBody {
	s.Success = &v
	return s
}

type DeleteEventRecordPlanResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteEventRecordPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteEventRecordPlanResponse) SetBody(v *DeleteEventRecordPlanResponseBody) *DeleteEventRecordPlanResponse {
	s.Body = v
	return s
}

type DeleteEventRecordPlanDeviceRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId       *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	StreamType  *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s DeleteEventRecordPlanDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventRecordPlanDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventRecordPlanDeviceRequest) SetApiProduct(v string) *DeleteEventRecordPlanDeviceRequest {
	s.ApiProduct = &v
	return s
}

func (s *DeleteEventRecordPlanDeviceRequest) SetApiRevision(v string) *DeleteEventRecordPlanDeviceRequest {
	s.ApiRevision = &v
	return s
}

func (s *DeleteEventRecordPlanDeviceRequest) SetIotId(v string) *DeleteEventRecordPlanDeviceRequest {
	s.IotId = &v
	return s
}

func (s *DeleteEventRecordPlanDeviceRequest) SetStreamType(v int32) *DeleteEventRecordPlanDeviceRequest {
	s.StreamType = &v
	return s
}

type DeleteEventRecordPlanDeviceResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteEventRecordPlanDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventRecordPlanDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventRecordPlanDeviceResponseBody) SetRequestId(v string) *DeleteEventRecordPlanDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEventRecordPlanDeviceResponseBody) SetErrorMessage(v string) *DeleteEventRecordPlanDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteEventRecordPlanDeviceResponseBody) SetCode(v string) *DeleteEventRecordPlanDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEventRecordPlanDeviceResponseBody) SetSuccess(v bool) *DeleteEventRecordPlanDeviceResponseBody {
	s.Success = &v
	return s
}

type DeleteEventRecordPlanDeviceResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteEventRecordPlanDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteEventRecordPlanDeviceResponse) SetBody(v *DeleteEventRecordPlanDeviceResponseBody) *DeleteEventRecordPlanDeviceResponse {
	s.Body = v
	return s
}

type DeleteFaceDeviceGroupRequest struct {
	ApiProduct    *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId   *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
}

func (s DeleteFaceDeviceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceDeviceGroupRequest) SetApiProduct(v string) *DeleteFaceDeviceGroupRequest {
	s.ApiProduct = &v
	return s
}

func (s *DeleteFaceDeviceGroupRequest) SetApiRevision(v string) *DeleteFaceDeviceGroupRequest {
	s.ApiRevision = &v
	return s
}

func (s *DeleteFaceDeviceGroupRequest) SetIsolationId(v string) *DeleteFaceDeviceGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *DeleteFaceDeviceGroupRequest) SetDeviceGroupId(v string) *DeleteFaceDeviceGroupRequest {
	s.DeviceGroupId = &v
	return s
}

type DeleteFaceDeviceGroupResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFaceDeviceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceDeviceGroupResponseBody) SetRequestId(v string) *DeleteFaceDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFaceDeviceGroupResponseBody) SetErrorMessage(v string) *DeleteFaceDeviceGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFaceDeviceGroupResponseBody) SetCode(v string) *DeleteFaceDeviceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFaceDeviceGroupResponseBody) SetSuccess(v bool) *DeleteFaceDeviceGroupResponseBody {
	s.Success = &v
	return s
}

type DeleteFaceDeviceGroupResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFaceDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteFaceDeviceGroupResponse) SetBody(v *DeleteFaceDeviceGroupResponseBody) *DeleteFaceDeviceGroupResponse {
	s.Body = v
	return s
}

type DeleteFaceUserRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteFaceUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserRequest) SetApiProduct(v string) *DeleteFaceUserRequest {
	s.ApiProduct = &v
	return s
}

func (s *DeleteFaceUserRequest) SetApiRevision(v string) *DeleteFaceUserRequest {
	s.ApiRevision = &v
	return s
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
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFaceUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserResponseBody) SetRequestId(v string) *DeleteFaceUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFaceUserResponseBody) SetErrorMessage(v string) *DeleteFaceUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFaceUserResponseBody) SetCode(v string) *DeleteFaceUserResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFaceUserResponseBody) SetSuccess(v bool) *DeleteFaceUserResponseBody {
	s.Success = &v
	return s
}

type DeleteFaceUserResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFaceUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteFaceUserResponse) SetBody(v *DeleteFaceUserResponseBody) *DeleteFaceUserResponse {
	s.Body = v
	return s
}

type DeleteFaceUserGroupRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s DeleteFaceUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceUserGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserGroupRequest) SetApiProduct(v string) *DeleteFaceUserGroupRequest {
	s.ApiProduct = &v
	return s
}

func (s *DeleteFaceUserGroupRequest) SetApiRevision(v string) *DeleteFaceUserGroupRequest {
	s.ApiRevision = &v
	return s
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
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFaceUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserGroupResponseBody) SetRequestId(v string) *DeleteFaceUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFaceUserGroupResponseBody) SetErrorMessage(v string) *DeleteFaceUserGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFaceUserGroupResponseBody) SetCode(v string) *DeleteFaceUserGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFaceUserGroupResponseBody) SetSuccess(v bool) *DeleteFaceUserGroupResponseBody {
	s.Success = &v
	return s
}

type DeleteFaceUserGroupResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFaceUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteFaceUserGroupResponse) SetBody(v *DeleteFaceUserGroupResponseBody) *DeleteFaceUserGroupResponse {
	s.Body = v
	return s
}

type DeleteFaceUserGroupAndDeviceGroupRelationRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	ControlId   *string `json:"ControlId,omitempty" xml:"ControlId,omitempty"`
}

func (s DeleteFaceUserGroupAndDeviceGroupRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceUserGroupAndDeviceGroupRelationRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationRequest) SetApiProduct(v string) *DeleteFaceUserGroupAndDeviceGroupRelationRequest {
	s.ApiProduct = &v
	return s
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationRequest) SetApiRevision(v string) *DeleteFaceUserGroupAndDeviceGroupRelationRequest {
	s.ApiRevision = &v
	return s
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationRequest) SetIsolationId(v string) *DeleteFaceUserGroupAndDeviceGroupRelationRequest {
	s.IsolationId = &v
	return s
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationRequest) SetControlId(v string) *DeleteFaceUserGroupAndDeviceGroupRelationRequest {
	s.ControlId = &v
	return s
}

type DeleteFaceUserGroupAndDeviceGroupRelationResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFaceUserGroupAndDeviceGroupRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceUserGroupAndDeviceGroupRelationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationResponseBody) SetRequestId(v string) *DeleteFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationResponseBody) SetErrorMessage(v string) *DeleteFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationResponseBody) SetCode(v string) *DeleteFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationResponseBody) SetSuccess(v bool) *DeleteFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.Success = &v
	return s
}

type DeleteFaceUserGroupAndDeviceGroupRelationResponse struct {
	Headers map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFaceUserGroupAndDeviceGroupRelationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteFaceUserGroupAndDeviceGroupRelationResponse) SetBody(v *DeleteFaceUserGroupAndDeviceGroupRelationResponseBody) *DeleteFaceUserGroupAndDeviceGroupRelationResponse {
	s.Body = v
	return s
}

type DeleteFaceUserPictureRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	FacePicMd5  *string `json:"FacePicMd5,omitempty" xml:"FacePicMd5,omitempty"`
}

func (s DeleteFaceUserPictureRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceUserPictureRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserPictureRequest) SetApiProduct(v string) *DeleteFaceUserPictureRequest {
	s.ApiProduct = &v
	return s
}

func (s *DeleteFaceUserPictureRequest) SetApiRevision(v string) *DeleteFaceUserPictureRequest {
	s.ApiRevision = &v
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

func (s *DeleteFaceUserPictureRequest) SetFacePicMd5(v string) *DeleteFaceUserPictureRequest {
	s.FacePicMd5 = &v
	return s
}

type DeleteFaceUserPictureResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFaceUserPictureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceUserPictureResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserPictureResponseBody) SetRequestId(v string) *DeleteFaceUserPictureResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFaceUserPictureResponseBody) SetErrorMessage(v string) *DeleteFaceUserPictureResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFaceUserPictureResponseBody) SetCode(v string) *DeleteFaceUserPictureResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFaceUserPictureResponseBody) SetSuccess(v bool) *DeleteFaceUserPictureResponseBody {
	s.Success = &v
	return s
}

type DeleteFaceUserPictureResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFaceUserPictureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteFaceUserPictureResponse) SetBody(v *DeleteFaceUserPictureResponseBody) *DeleteFaceUserPictureResponse {
	s.Body = v
	return s
}

type DeleteModelRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	ModelId     *int64  `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
}

func (s DeleteModelRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelRequest) SetApiProduct(v string) *DeleteModelRequest {
	s.ApiProduct = &v
	return s
}

func (s *DeleteModelRequest) SetApiRevision(v string) *DeleteModelRequest {
	s.ApiRevision = &v
	return s
}

func (s *DeleteModelRequest) SetModelId(v int64) *DeleteModelRequest {
	s.ModelId = &v
	return s
}

type DeleteModelResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModelResponseBody) SetRequestId(v string) *DeleteModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteModelResponseBody) SetErrorMessage(v string) *DeleteModelResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteModelResponseBody) SetCode(v string) *DeleteModelResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteModelResponseBody) SetSuccess(v bool) *DeleteModelResponseBody {
	s.Success = &v
	return s
}

type DeleteModelResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteModelResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelResponse) GoString() string {
	return s.String()
}

func (s *DeleteModelResponse) SetHeaders(v map[string]*string) *DeleteModelResponse {
	s.Headers = v
	return s
}

func (s *DeleteModelResponse) SetBody(v *DeleteModelResponseBody) *DeleteModelResponse {
	s.Body = v
	return s
}

type DeleteRecordPlanRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	PlanId      *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s DeleteRecordPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecordPlanRequest) SetApiProduct(v string) *DeleteRecordPlanRequest {
	s.ApiProduct = &v
	return s
}

func (s *DeleteRecordPlanRequest) SetApiRevision(v string) *DeleteRecordPlanRequest {
	s.ApiRevision = &v
	return s
}

func (s *DeleteRecordPlanRequest) SetPlanId(v string) *DeleteRecordPlanRequest {
	s.PlanId = &v
	return s
}

type DeleteRecordPlanResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRecordPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRecordPlanResponseBody) SetRequestId(v string) *DeleteRecordPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRecordPlanResponseBody) SetErrorMessage(v string) *DeleteRecordPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRecordPlanResponseBody) SetCode(v string) *DeleteRecordPlanResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRecordPlanResponseBody) SetSuccess(v bool) *DeleteRecordPlanResponseBody {
	s.Success = &v
	return s
}

type DeleteRecordPlanResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRecordPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteRecordPlanResponse) SetBody(v *DeleteRecordPlanResponseBody) *DeleteRecordPlanResponse {
	s.Body = v
	return s
}

type DeleteRecordPlanDeviceRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId       *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	StreamType  *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s DeleteRecordPlanDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordPlanDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecordPlanDeviceRequest) SetApiProduct(v string) *DeleteRecordPlanDeviceRequest {
	s.ApiProduct = &v
	return s
}

func (s *DeleteRecordPlanDeviceRequest) SetApiRevision(v string) *DeleteRecordPlanDeviceRequest {
	s.ApiRevision = &v
	return s
}

func (s *DeleteRecordPlanDeviceRequest) SetIotId(v string) *DeleteRecordPlanDeviceRequest {
	s.IotId = &v
	return s
}

func (s *DeleteRecordPlanDeviceRequest) SetStreamType(v int32) *DeleteRecordPlanDeviceRequest {
	s.StreamType = &v
	return s
}

type DeleteRecordPlanDeviceResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRecordPlanDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordPlanDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRecordPlanDeviceResponseBody) SetRequestId(v string) *DeleteRecordPlanDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRecordPlanDeviceResponseBody) SetErrorMessage(v string) *DeleteRecordPlanDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRecordPlanDeviceResponseBody) SetCode(v string) *DeleteRecordPlanDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRecordPlanDeviceResponseBody) SetSuccess(v bool) *DeleteRecordPlanDeviceResponseBody {
	s.Success = &v
	return s
}

type DeleteRecordPlanDeviceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRecordPlanDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteRecordPlanDeviceResponse) SetBody(v *DeleteRecordPlanDeviceResponseBody) *DeleteRecordPlanDeviceResponse {
	s.Body = v
	return s
}

type DeleteTimeTemplateRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	TemplateId  *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteTimeTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTimeTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteTimeTemplateRequest) SetApiProduct(v string) *DeleteTimeTemplateRequest {
	s.ApiProduct = &v
	return s
}

func (s *DeleteTimeTemplateRequest) SetApiRevision(v string) *DeleteTimeTemplateRequest {
	s.ApiRevision = &v
	return s
}

func (s *DeleteTimeTemplateRequest) SetTemplateId(v string) *DeleteTimeTemplateRequest {
	s.TemplateId = &v
	return s
}

type DeleteTimeTemplateResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTimeTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTimeTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTimeTemplateResponseBody) SetRequestId(v string) *DeleteTimeTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTimeTemplateResponseBody) SetErrorMessage(v string) *DeleteTimeTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteTimeTemplateResponseBody) SetCode(v string) *DeleteTimeTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTimeTemplateResponseBody) SetSuccess(v bool) *DeleteTimeTemplateResponseBody {
	s.Success = &v
	return s
}

type DeleteTimeTemplateResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTimeTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteTimeTemplateResponse) SetBody(v *DeleteTimeTemplateResponseBody) *DeleteTimeTemplateResponse {
	s.Body = v
	return s
}

type DeployModelBatchRequest struct {
	ApiProduct  *string                              `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string                              `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	ModelId     *int64                               `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	DeviceList  []*DeployModelBatchRequestDeviceList `json:"DeviceList,omitempty" xml:"DeviceList,omitempty" type:"Repeated"`
}

func (s DeployModelBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployModelBatchRequest) GoString() string {
	return s.String()
}

func (s *DeployModelBatchRequest) SetApiProduct(v string) *DeployModelBatchRequest {
	s.ApiProduct = &v
	return s
}

func (s *DeployModelBatchRequest) SetApiRevision(v string) *DeployModelBatchRequest {
	s.ApiRevision = &v
	return s
}

func (s *DeployModelBatchRequest) SetModelId(v int64) *DeployModelBatchRequest {
	s.ModelId = &v
	return s
}

func (s *DeployModelBatchRequest) SetDeviceList(v []*DeployModelBatchRequestDeviceList) *DeployModelBatchRequest {
	s.DeviceList = v
	return s
}

type DeployModelBatchRequestDeviceList struct {
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
}

func (s DeployModelBatchRequestDeviceList) String() string {
	return tea.Prettify(s)
}

func (s DeployModelBatchRequestDeviceList) GoString() string {
	return s.String()
}

func (s *DeployModelBatchRequestDeviceList) SetProductKey(v string) *DeployModelBatchRequestDeviceList {
	s.ProductKey = &v
	return s
}

func (s *DeployModelBatchRequestDeviceList) SetDeviceName(v string) *DeployModelBatchRequestDeviceList {
	s.DeviceName = &v
	return s
}

func (s *DeployModelBatchRequestDeviceList) SetIotId(v string) *DeployModelBatchRequestDeviceList {
	s.IotId = &v
	return s
}

type DeployModelBatchResponseBody struct {
	RequestId    *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *DeployModelBatchResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeployModelBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployModelBatchResponseBody) GoString() string {
	return s.String()
}

func (s *DeployModelBatchResponseBody) SetRequestId(v string) *DeployModelBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployModelBatchResponseBody) SetData(v *DeployModelBatchResponseBodyData) *DeployModelBatchResponseBody {
	s.Data = v
	return s
}

func (s *DeployModelBatchResponseBody) SetErrorMessage(v string) *DeployModelBatchResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeployModelBatchResponseBody) SetCode(v string) *DeployModelBatchResponseBody {
	s.Code = &v
	return s
}

func (s *DeployModelBatchResponseBody) SetSuccess(v bool) *DeployModelBatchResponseBody {
	s.Success = &v
	return s
}

type DeployModelBatchResponseBodyData struct {
	DeployTaskResultVOList []map[string]interface{} `json:"DeployTaskResultVOList,omitempty" xml:"DeployTaskResultVOList,omitempty" type:"Repeated"`
}

func (s DeployModelBatchResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeployModelBatchResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeployModelBatchResponseBodyData) SetDeployTaskResultVOList(v []map[string]interface{}) *DeployModelBatchResponseBodyData {
	s.DeployTaskResultVOList = v
	return s
}

type DeployModelBatchResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeployModelBatchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeployModelBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployModelBatchResponse) GoString() string {
	return s.String()
}

func (s *DeployModelBatchResponse) SetHeaders(v map[string]*string) *DeployModelBatchResponse {
	s.Headers = v
	return s
}

func (s *DeployModelBatchResponse) SetBody(v *DeployModelBatchResponseBody) *DeployModelBatchResponse {
	s.Body = v
	return s
}

type DetectUserFaceByUrlRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	FacePicUrl  *string `json:"FacePicUrl,omitempty" xml:"FacePicUrl,omitempty"`
}

func (s DetectUserFaceByUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectUserFaceByUrlRequest) GoString() string {
	return s.String()
}

func (s *DetectUserFaceByUrlRequest) SetApiProduct(v string) *DetectUserFaceByUrlRequest {
	s.ApiProduct = &v
	return s
}

func (s *DetectUserFaceByUrlRequest) SetApiRevision(v string) *DetectUserFaceByUrlRequest {
	s.ApiRevision = &v
	return s
}

func (s *DetectUserFaceByUrlRequest) SetFacePicUrl(v string) *DetectUserFaceByUrlRequest {
	s.FacePicUrl = &v
	return s
}

type DetectUserFaceByUrlResponseBody struct {
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *DetectUserFaceByUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetectUserFaceByUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectUserFaceByUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DetectUserFaceByUrlResponseBody) SetRequestId(v string) *DetectUserFaceByUrlResponseBody {
	s.RequestId = &v
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

func (s *DetectUserFaceByUrlResponseBody) SetCode(v string) *DetectUserFaceByUrlResponseBody {
	s.Code = &v
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
	Gender          *int32                                            `json:"Gender,omitempty" xml:"Gender,omitempty"`
	FaceRects       *DetectUserFaceByUrlResponseBodyDataDataFaceRects `json:"FaceRects,omitempty" xml:"FaceRects,omitempty" type:"Struct"`
	Age             *int32                                            `json:"Age,omitempty" xml:"Age,omitempty"`
	Landmarks       *DetectUserFaceByUrlResponseBodyDataDataLandmarks `json:"Landmarks,omitempty" xml:"Landmarks,omitempty" type:"Struct"`
	FaceProbability *float32                                          `json:"FaceProbability,omitempty" xml:"FaceProbability,omitempty"`
}

func (s DetectUserFaceByUrlResponseBodyDataData) String() string {
	return tea.Prettify(s)
}

func (s DetectUserFaceByUrlResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *DetectUserFaceByUrlResponseBodyDataData) SetGender(v int32) *DetectUserFaceByUrlResponseBodyDataData {
	s.Gender = &v
	return s
}

func (s *DetectUserFaceByUrlResponseBodyDataData) SetFaceRects(v *DetectUserFaceByUrlResponseBodyDataDataFaceRects) *DetectUserFaceByUrlResponseBodyDataData {
	s.FaceRects = v
	return s
}

func (s *DetectUserFaceByUrlResponseBodyDataData) SetAge(v int32) *DetectUserFaceByUrlResponseBodyDataData {
	s.Age = &v
	return s
}

func (s *DetectUserFaceByUrlResponseBodyDataData) SetLandmarks(v *DetectUserFaceByUrlResponseBodyDataDataLandmarks) *DetectUserFaceByUrlResponseBodyDataData {
	s.Landmarks = v
	return s
}

func (s *DetectUserFaceByUrlResponseBodyDataData) SetFaceProbability(v float32) *DetectUserFaceByUrlResponseBodyDataData {
	s.FaceProbability = &v
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectUserFaceByUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetectUserFaceByUrlResponse) SetBody(v *DetectUserFaceByUrlResponseBody) *DetectUserFaceByUrlResponse {
	s.Body = v
	return s
}

type GetAIActionRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	ActionId    *string `json:"ActionId,omitempty" xml:"ActionId,omitempty"`
}

func (s GetAIActionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAIActionRequest) GoString() string {
	return s.String()
}

func (s *GetAIActionRequest) SetApiProduct(v string) *GetAIActionRequest {
	s.ApiProduct = &v
	return s
}

func (s *GetAIActionRequest) SetApiRevision(v string) *GetAIActionRequest {
	s.ApiRevision = &v
	return s
}

func (s *GetAIActionRequest) SetActionId(v string) *GetAIActionRequest {
	s.ActionId = &v
	return s
}

type GetAIActionResponseBody struct {
	RequestId    *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage *string                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *GetAIActionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetAIActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAIActionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAIActionResponseBody) SetRequestId(v string) *GetAIActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAIActionResponseBody) SetSuccess(v bool) *GetAIActionResponseBody {
	s.Success = &v
	return s
}

func (s *GetAIActionResponseBody) SetErrorMessage(v string) *GetAIActionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAIActionResponseBody) SetCode(v string) *GetAIActionResponseBody {
	s.Code = &v
	return s
}

func (s *GetAIActionResponseBody) SetData(v *GetAIActionResponseBodyData) *GetAIActionResponseBody {
	s.Data = v
	return s
}

type GetAIActionResponseBodyData struct {
	ActionId         *string `json:"ActionId,omitempty" xml:"ActionId,omitempty"`
	PlanId           *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	Action           *string `json:"Action,omitempty" xml:"Action,omitempty"`
	ActionTemplateId *string `json:"ActionTemplateId,omitempty" xml:"ActionTemplateId,omitempty"`
	ActionIndex      *int32  `json:"ActionIndex,omitempty" xml:"ActionIndex,omitempty"`
	Alog             *string `json:"Alog,omitempty" xml:"Alog,omitempty"`
	AlgoConfig       *string `json:"AlgoConfig,omitempty" xml:"AlgoConfig,omitempty"`
	ActionConfig     *string `json:"ActionConfig,omitempty" xml:"ActionConfig,omitempty"`
}

func (s GetAIActionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAIActionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAIActionResponseBodyData) SetActionId(v string) *GetAIActionResponseBodyData {
	s.ActionId = &v
	return s
}

func (s *GetAIActionResponseBodyData) SetPlanId(v string) *GetAIActionResponseBodyData {
	s.PlanId = &v
	return s
}

func (s *GetAIActionResponseBodyData) SetAction(v string) *GetAIActionResponseBodyData {
	s.Action = &v
	return s
}

func (s *GetAIActionResponseBodyData) SetActionTemplateId(v string) *GetAIActionResponseBodyData {
	s.ActionTemplateId = &v
	return s
}

func (s *GetAIActionResponseBodyData) SetActionIndex(v int32) *GetAIActionResponseBodyData {
	s.ActionIndex = &v
	return s
}

func (s *GetAIActionResponseBodyData) SetAlog(v string) *GetAIActionResponseBodyData {
	s.Alog = &v
	return s
}

func (s *GetAIActionResponseBodyData) SetAlgoConfig(v string) *GetAIActionResponseBodyData {
	s.AlgoConfig = &v
	return s
}

func (s *GetAIActionResponseBodyData) SetActionConfig(v string) *GetAIActionResponseBodyData {
	s.ActionConfig = &v
	return s
}

type GetAIActionResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAIActionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAIActionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAIActionResponse) GoString() string {
	return s.String()
}

func (s *GetAIActionResponse) SetHeaders(v map[string]*string) *GetAIActionResponse {
	s.Headers = v
	return s
}

func (s *GetAIActionResponse) SetBody(v *GetAIActionResponseBody) *GetAIActionResponse {
	s.Body = v
	return s
}

type GetAIActionConfigRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	Algo        *string `json:"Algo,omitempty" xml:"Algo,omitempty"`
	AlgoAction  *string `json:"AlgoAction,omitempty" xml:"AlgoAction,omitempty"`
}

func (s GetAIActionConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAIActionConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAIActionConfigRequest) SetApiProduct(v string) *GetAIActionConfigRequest {
	s.ApiProduct = &v
	return s
}

func (s *GetAIActionConfigRequest) SetApiRevision(v string) *GetAIActionConfigRequest {
	s.ApiRevision = &v
	return s
}

func (s *GetAIActionConfigRequest) SetAlgo(v string) *GetAIActionConfigRequest {
	s.Algo = &v
	return s
}

func (s *GetAIActionConfigRequest) SetAlgoAction(v string) *GetAIActionConfigRequest {
	s.AlgoAction = &v
	return s
}

type GetAIActionConfigResponseBody struct {
	RequestId    *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage *string                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *GetAIActionConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetAIActionConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAIActionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetAIActionConfigResponseBody) SetRequestId(v string) *GetAIActionConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAIActionConfigResponseBody) SetSuccess(v bool) *GetAIActionConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetAIActionConfigResponseBody) SetErrorMessage(v string) *GetAIActionConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAIActionConfigResponseBody) SetCode(v string) *GetAIActionConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetAIActionConfigResponseBody) SetData(v *GetAIActionConfigResponseBodyData) *GetAIActionConfigResponseBody {
	s.Data = v
	return s
}

type GetAIActionConfigResponseBodyData struct {
	AlgoAction      *string                                          `json:"AlgoAction,omitempty" xml:"AlgoAction,omitempty"`
	Des             *string                                          `json:"Des,omitempty" xml:"Des,omitempty"`
	NeedDevice      *bool                                            `json:"NeedDevice,omitempty" xml:"NeedDevice,omitempty"`
	Sync            *string                                          `json:"Sync,omitempty" xml:"Sync,omitempty"`
	AlgoConfigItems *string                                          `json:"AlgoConfigItems,omitempty" xml:"AlgoConfigItems,omitempty"`
	InParamList     []*GetAIActionConfigResponseBodyDataInParamList  `json:"InParamList,omitempty" xml:"InParamList,omitempty" type:"Repeated"`
	OutParamList    []*GetAIActionConfigResponseBodyDataOutParamList `json:"OutParamList,omitempty" xml:"OutParamList,omitempty" type:"Repeated"`
}

func (s GetAIActionConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAIActionConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAIActionConfigResponseBodyData) SetAlgoAction(v string) *GetAIActionConfigResponseBodyData {
	s.AlgoAction = &v
	return s
}

func (s *GetAIActionConfigResponseBodyData) SetDes(v string) *GetAIActionConfigResponseBodyData {
	s.Des = &v
	return s
}

func (s *GetAIActionConfigResponseBodyData) SetNeedDevice(v bool) *GetAIActionConfigResponseBodyData {
	s.NeedDevice = &v
	return s
}

func (s *GetAIActionConfigResponseBodyData) SetSync(v string) *GetAIActionConfigResponseBodyData {
	s.Sync = &v
	return s
}

func (s *GetAIActionConfigResponseBodyData) SetAlgoConfigItems(v string) *GetAIActionConfigResponseBodyData {
	s.AlgoConfigItems = &v
	return s
}

func (s *GetAIActionConfigResponseBodyData) SetInParamList(v []*GetAIActionConfigResponseBodyDataInParamList) *GetAIActionConfigResponseBodyData {
	s.InParamList = v
	return s
}

func (s *GetAIActionConfigResponseBodyData) SetOutParamList(v []*GetAIActionConfigResponseBodyDataOutParamList) *GetAIActionConfigResponseBodyData {
	s.OutParamList = v
	return s
}

type GetAIActionConfigResponseBodyDataInParamList struct {
	DataType    *string   `json:"DataType,omitempty" xml:"DataType,omitempty"`
	NeedConfig  *bool     `json:"NeedConfig,omitempty" xml:"NeedConfig,omitempty"`
	ConfigItems []*string `json:"ConfigItems,omitempty" xml:"ConfigItems,omitempty" type:"Repeated"`
}

func (s GetAIActionConfigResponseBodyDataInParamList) String() string {
	return tea.Prettify(s)
}

func (s GetAIActionConfigResponseBodyDataInParamList) GoString() string {
	return s.String()
}

func (s *GetAIActionConfigResponseBodyDataInParamList) SetDataType(v string) *GetAIActionConfigResponseBodyDataInParamList {
	s.DataType = &v
	return s
}

func (s *GetAIActionConfigResponseBodyDataInParamList) SetNeedConfig(v bool) *GetAIActionConfigResponseBodyDataInParamList {
	s.NeedConfig = &v
	return s
}

func (s *GetAIActionConfigResponseBodyDataInParamList) SetConfigItems(v []*string) *GetAIActionConfigResponseBodyDataInParamList {
	s.ConfigItems = v
	return s
}

type GetAIActionConfigResponseBodyDataOutParamList struct {
	DataType    *string   `json:"DataType,omitempty" xml:"DataType,omitempty"`
	NeedConfig  *bool     `json:"NeedConfig,omitempty" xml:"NeedConfig,omitempty"`
	ConfigItems []*string `json:"ConfigItems,omitempty" xml:"ConfigItems,omitempty" type:"Repeated"`
}

func (s GetAIActionConfigResponseBodyDataOutParamList) String() string {
	return tea.Prettify(s)
}

func (s GetAIActionConfigResponseBodyDataOutParamList) GoString() string {
	return s.String()
}

func (s *GetAIActionConfigResponseBodyDataOutParamList) SetDataType(v string) *GetAIActionConfigResponseBodyDataOutParamList {
	s.DataType = &v
	return s
}

func (s *GetAIActionConfigResponseBodyDataOutParamList) SetNeedConfig(v bool) *GetAIActionConfigResponseBodyDataOutParamList {
	s.NeedConfig = &v
	return s
}

func (s *GetAIActionConfigResponseBodyDataOutParamList) SetConfigItems(v []*string) *GetAIActionConfigResponseBodyDataOutParamList {
	s.ConfigItems = v
	return s
}

type GetAIActionConfigResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAIActionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAIActionConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAIActionConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAIActionConfigResponse) SetHeaders(v map[string]*string) *GetAIActionConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAIActionConfigResponse) SetBody(v *GetAIActionConfigResponseBody) *GetAIActionConfigResponse {
	s.Body = v
	return s
}

type GetAIAppRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetAIAppRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAIAppRequest) GoString() string {
	return s.String()
}

func (s *GetAIAppRequest) SetApiProduct(v string) *GetAIAppRequest {
	s.ApiProduct = &v
	return s
}

func (s *GetAIAppRequest) SetApiRevision(v string) *GetAIAppRequest {
	s.ApiRevision = &v
	return s
}

func (s *GetAIAppRequest) SetAppId(v string) *GetAIAppRequest {
	s.AppId = &v
	return s
}

type GetAIAppResponseBody struct {
	RequestId    *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                     `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage *string                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *GetAIAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetAIAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAIAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetAIAppResponseBody) SetRequestId(v string) *GetAIAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAIAppResponseBody) SetSuccess(v bool) *GetAIAppResponseBody {
	s.Success = &v
	return s
}

func (s *GetAIAppResponseBody) SetErrorMessage(v string) *GetAIAppResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAIAppResponseBody) SetCode(v string) *GetAIAppResponseBody {
	s.Code = &v
	return s
}

func (s *GetAIAppResponseBody) SetData(v *GetAIAppResponseBodyData) *GetAIAppResponseBody {
	s.Data = v
	return s
}

type GetAIAppResponseBodyData struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppTemplateId *string `json:"AppTemplateId,omitempty" xml:"AppTemplateId,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Level         *int32  `json:"Level,omitempty" xml:"Level,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s GetAIAppResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAIAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAIAppResponseBodyData) SetAppId(v string) *GetAIAppResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetAIAppResponseBodyData) SetAppTemplateId(v string) *GetAIAppResponseBodyData {
	s.AppTemplateId = &v
	return s
}

func (s *GetAIAppResponseBodyData) SetVersion(v string) *GetAIAppResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetAIAppResponseBodyData) SetLevel(v int32) *GetAIAppResponseBodyData {
	s.Level = &v
	return s
}

func (s *GetAIAppResponseBodyData) SetName(v string) *GetAIAppResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetAIAppResponseBodyData) SetDescription(v string) *GetAIAppResponseBodyData {
	s.Description = &v
	return s
}

type GetAIAppResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAIAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAIAppResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAIAppResponse) GoString() string {
	return s.String()
}

func (s *GetAIAppResponse) SetHeaders(v map[string]*string) *GetAIAppResponse {
	s.Headers = v
	return s
}

func (s *GetAIAppResponse) SetBody(v *GetAIAppResponseBody) *GetAIAppResponse {
	s.Body = v
	return s
}

type GetAIJobRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	JobId       *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetAIJobRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAIJobRequest) GoString() string {
	return s.String()
}

func (s *GetAIJobRequest) SetApiProduct(v string) *GetAIJobRequest {
	s.ApiProduct = &v
	return s
}

func (s *GetAIJobRequest) SetApiRevision(v string) *GetAIJobRequest {
	s.ApiRevision = &v
	return s
}

func (s *GetAIJobRequest) SetJobId(v string) *GetAIJobRequest {
	s.JobId = &v
	return s
}

type GetAIJobResponseBody struct {
	RequestId    *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                     `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage *string                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *GetAIJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetAIJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAIJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetAIJobResponseBody) SetRequestId(v string) *GetAIJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAIJobResponseBody) SetSuccess(v bool) *GetAIJobResponseBody {
	s.Success = &v
	return s
}

func (s *GetAIJobResponseBody) SetErrorMessage(v string) *GetAIJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAIJobResponseBody) SetCode(v string) *GetAIJobResponseBody {
	s.Code = &v
	return s
}

func (s *GetAIJobResponseBody) SetData(v *GetAIJobResponseBodyData) *GetAIJobResponseBody {
	s.Data = v
	return s
}

type GetAIJobResponseBodyData struct {
	DataDTOList  []*GetAIJobResponseBodyDataDataDTOList `json:"DataDTOList,omitempty" xml:"DataDTOList,omitempty" type:"Repeated"`
	ActionJobDTO *GetAIJobResponseBodyDataActionJobDTO  `json:"ActionJobDTO,omitempty" xml:"ActionJobDTO,omitempty" type:"Struct"`
}

func (s GetAIJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAIJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAIJobResponseBodyData) SetDataDTOList(v []*GetAIJobResponseBodyDataDataDTOList) *GetAIJobResponseBodyData {
	s.DataDTOList = v
	return s
}

func (s *GetAIJobResponseBodyData) SetActionJobDTO(v *GetAIJobResponseBodyDataActionJobDTO) *GetAIJobResponseBodyData {
	s.ActionJobDTO = v
	return s
}

type GetAIJobResponseBodyDataDataDTOList struct {
	DataId     *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	DataType   *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DataSource *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	AlgoData   *string `json:"AlgoData,omitempty" xml:"AlgoData,omitempty"`
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
}

func (s GetAIJobResponseBodyDataDataDTOList) String() string {
	return tea.Prettify(s)
}

func (s GetAIJobResponseBodyDataDataDTOList) GoString() string {
	return s.String()
}

func (s *GetAIJobResponseBodyDataDataDTOList) SetDataId(v string) *GetAIJobResponseBodyDataDataDTOList {
	s.DataId = &v
	return s
}

func (s *GetAIJobResponseBodyDataDataDTOList) SetDataType(v string) *GetAIJobResponseBodyDataDataDTOList {
	s.DataType = &v
	return s
}

func (s *GetAIJobResponseBodyDataDataDTOList) SetDataSource(v string) *GetAIJobResponseBodyDataDataDTOList {
	s.DataSource = &v
	return s
}

func (s *GetAIJobResponseBodyDataDataDTOList) SetAlgoData(v string) *GetAIJobResponseBodyDataDataDTOList {
	s.AlgoData = &v
	return s
}

func (s *GetAIJobResponseBodyDataDataDTOList) SetJobId(v string) *GetAIJobResponseBodyDataDataDTOList {
	s.JobId = &v
	return s
}

func (s *GetAIJobResponseBodyDataDataDTOList) SetIotId(v string) *GetAIJobResponseBodyDataDataDTOList {
	s.IotId = &v
	return s
}

type GetAIJobResponseBodyDataActionJobDTO struct {
	JobId    *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ActionId *string `json:"ActionId,omitempty" xml:"ActionId,omitempty"`
	Status   *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	IotId    *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
}

func (s GetAIJobResponseBodyDataActionJobDTO) String() string {
	return tea.Prettify(s)
}

func (s GetAIJobResponseBodyDataActionJobDTO) GoString() string {
	return s.String()
}

func (s *GetAIJobResponseBodyDataActionJobDTO) SetJobId(v string) *GetAIJobResponseBodyDataActionJobDTO {
	s.JobId = &v
	return s
}

func (s *GetAIJobResponseBodyDataActionJobDTO) SetActionId(v string) *GetAIJobResponseBodyDataActionJobDTO {
	s.ActionId = &v
	return s
}

func (s *GetAIJobResponseBodyDataActionJobDTO) SetStatus(v int32) *GetAIJobResponseBodyDataActionJobDTO {
	s.Status = &v
	return s
}

func (s *GetAIJobResponseBodyDataActionJobDTO) SetIotId(v string) *GetAIJobResponseBodyDataActionJobDTO {
	s.IotId = &v
	return s
}

type GetAIJobResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAIJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAIJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAIJobResponse) GoString() string {
	return s.String()
}

func (s *GetAIJobResponse) SetHeaders(v map[string]*string) *GetAIJobResponse {
	s.Headers = v
	return s
}

func (s *GetAIJobResponse) SetBody(v *GetAIJobResponseBody) *GetAIJobResponse {
	s.Body = v
	return s
}

type GetAIPlanRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	PlanId      *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s GetAIPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAIPlanRequest) GoString() string {
	return s.String()
}

func (s *GetAIPlanRequest) SetApiProduct(v string) *GetAIPlanRequest {
	s.ApiProduct = &v
	return s
}

func (s *GetAIPlanRequest) SetApiRevision(v string) *GetAIPlanRequest {
	s.ApiRevision = &v
	return s
}

func (s *GetAIPlanRequest) SetPlanId(v string) *GetAIPlanRequest {
	s.PlanId = &v
	return s
}

type GetAIPlanResponseBody struct {
	RequestId    *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage *string                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *GetAIPlanResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetAIPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAIPlanResponseBody) GoString() string {
	return s.String()
}

func (s *GetAIPlanResponseBody) SetRequestId(v string) *GetAIPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAIPlanResponseBody) SetSuccess(v bool) *GetAIPlanResponseBody {
	s.Success = &v
	return s
}

func (s *GetAIPlanResponseBody) SetErrorMessage(v string) *GetAIPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAIPlanResponseBody) SetCode(v string) *GetAIPlanResponseBody {
	s.Code = &v
	return s
}

func (s *GetAIPlanResponseBody) SetData(v *GetAIPlanResponseBodyData) *GetAIPlanResponseBody {
	s.Data = v
	return s
}

type GetAIPlanResponseBodyData struct {
	PlanId         *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PlanTemplateId *string `json:"PlanTemplateId,omitempty" xml:"PlanTemplateId,omitempty"`
	TriggerEnum    *int32  `json:"TriggerEnum,omitempty" xml:"TriggerEnum,omitempty"`
	IntervalTiming *int32  `json:"IntervalTiming,omitempty" xml:"IntervalTiming,omitempty"`
	PreTiming      *int64  `json:"PreTiming,omitempty" xml:"PreTiming,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s GetAIPlanResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAIPlanResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAIPlanResponseBodyData) SetPlanId(v string) *GetAIPlanResponseBodyData {
	s.PlanId = &v
	return s
}

func (s *GetAIPlanResponseBodyData) SetAppId(v string) *GetAIPlanResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetAIPlanResponseBodyData) SetPlanTemplateId(v string) *GetAIPlanResponseBodyData {
	s.PlanTemplateId = &v
	return s
}

func (s *GetAIPlanResponseBodyData) SetTriggerEnum(v int32) *GetAIPlanResponseBodyData {
	s.TriggerEnum = &v
	return s
}

func (s *GetAIPlanResponseBodyData) SetIntervalTiming(v int32) *GetAIPlanResponseBodyData {
	s.IntervalTiming = &v
	return s
}

func (s *GetAIPlanResponseBodyData) SetPreTiming(v int64) *GetAIPlanResponseBodyData {
	s.PreTiming = &v
	return s
}

func (s *GetAIPlanResponseBodyData) SetDescription(v string) *GetAIPlanResponseBodyData {
	s.Description = &v
	return s
}

type GetAIPlanResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAIPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAIPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAIPlanResponse) GoString() string {
	return s.String()
}

func (s *GetAIPlanResponse) SetHeaders(v map[string]*string) *GetAIPlanResponse {
	s.Headers = v
	return s
}

func (s *GetAIPlanResponse) SetBody(v *GetAIPlanResponseBody) *GetAIPlanResponse {
	s.Body = v
	return s
}

type GetAlgorithmDetailByIdRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	AlgorithmId *int64  `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
}

func (s GetAlgorithmDetailByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmDetailByIdRequest) GoString() string {
	return s.String()
}

func (s *GetAlgorithmDetailByIdRequest) SetApiProduct(v string) *GetAlgorithmDetailByIdRequest {
	s.ApiProduct = &v
	return s
}

func (s *GetAlgorithmDetailByIdRequest) SetApiRevision(v string) *GetAlgorithmDetailByIdRequest {
	s.ApiRevision = &v
	return s
}

func (s *GetAlgorithmDetailByIdRequest) SetAlgorithmId(v int64) *GetAlgorithmDetailByIdRequest {
	s.AlgorithmId = &v
	return s
}

type GetAlgorithmDetailByIdResponseBody struct {
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAlgorithmDetailByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmDetailByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlgorithmDetailByIdResponseBody) SetRequestId(v string) *GetAlgorithmDetailByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlgorithmDetailByIdResponseBody) SetData(v map[string]interface{}) *GetAlgorithmDetailByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetAlgorithmDetailByIdResponseBody) SetErrorMessage(v string) *GetAlgorithmDetailByIdResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAlgorithmDetailByIdResponseBody) SetCode(v string) *GetAlgorithmDetailByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetAlgorithmDetailByIdResponseBody) SetSuccess(v bool) *GetAlgorithmDetailByIdResponseBody {
	s.Success = &v
	return s
}

type GetAlgorithmDetailByIdResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAlgorithmDetailByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAlgorithmDetailByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmDetailByIdResponse) GoString() string {
	return s.String()
}

func (s *GetAlgorithmDetailByIdResponse) SetHeaders(v map[string]*string) *GetAlgorithmDetailByIdResponse {
	s.Headers = v
	return s
}

func (s *GetAlgorithmDetailByIdResponse) SetBody(v *GetAlgorithmDetailByIdResponseBody) *GetAlgorithmDetailByIdResponse {
	s.Body = v
	return s
}

type GetAlgorithmDetailByNameRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetAlgorithmDetailByNameRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmDetailByNameRequest) GoString() string {
	return s.String()
}

func (s *GetAlgorithmDetailByNameRequest) SetApiProduct(v string) *GetAlgorithmDetailByNameRequest {
	s.ApiProduct = &v
	return s
}

func (s *GetAlgorithmDetailByNameRequest) SetApiRevision(v string) *GetAlgorithmDetailByNameRequest {
	s.ApiRevision = &v
	return s
}

func (s *GetAlgorithmDetailByNameRequest) SetName(v string) *GetAlgorithmDetailByNameRequest {
	s.Name = &v
	return s
}

type GetAlgorithmDetailByNameResponseBody struct {
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAlgorithmDetailByNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmDetailByNameResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlgorithmDetailByNameResponseBody) SetRequestId(v string) *GetAlgorithmDetailByNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlgorithmDetailByNameResponseBody) SetData(v map[string]interface{}) *GetAlgorithmDetailByNameResponseBody {
	s.Data = v
	return s
}

func (s *GetAlgorithmDetailByNameResponseBody) SetErrorMessage(v string) *GetAlgorithmDetailByNameResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAlgorithmDetailByNameResponseBody) SetCode(v string) *GetAlgorithmDetailByNameResponseBody {
	s.Code = &v
	return s
}

func (s *GetAlgorithmDetailByNameResponseBody) SetSuccess(v bool) *GetAlgorithmDetailByNameResponseBody {
	s.Success = &v
	return s
}

type GetAlgorithmDetailByNameResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAlgorithmDetailByNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAlgorithmDetailByNameResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmDetailByNameResponse) GoString() string {
	return s.String()
}

func (s *GetAlgorithmDetailByNameResponse) SetHeaders(v map[string]*string) *GetAlgorithmDetailByNameResponse {
	s.Headers = v
	return s
}

func (s *GetAlgorithmDetailByNameResponse) SetBody(v *GetAlgorithmDetailByNameResponseBody) *GetAlgorithmDetailByNameResponse {
	s.Body = v
	return s
}

type GetDevicePurifyJobByJobIdRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	JobId       *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetDevicePurifyJobByJobIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDevicePurifyJobByJobIdRequest) GoString() string {
	return s.String()
}

func (s *GetDevicePurifyJobByJobIdRequest) SetApiProduct(v string) *GetDevicePurifyJobByJobIdRequest {
	s.ApiProduct = &v
	return s
}

func (s *GetDevicePurifyJobByJobIdRequest) SetApiRevision(v string) *GetDevicePurifyJobByJobIdRequest {
	s.ApiRevision = &v
	return s
}

func (s *GetDevicePurifyJobByJobIdRequest) SetJobId(v string) *GetDevicePurifyJobByJobIdRequest {
	s.JobId = &v
	return s
}

type GetDevicePurifyJobByJobIdResponseBody struct {
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *GetDevicePurifyJobByJobIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDevicePurifyJobByJobIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDevicePurifyJobByJobIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetDevicePurifyJobByJobIdResponseBody) SetRequestId(v string) *GetDevicePurifyJobByJobIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDevicePurifyJobByJobIdResponseBody) SetData(v *GetDevicePurifyJobByJobIdResponseBodyData) *GetDevicePurifyJobByJobIdResponseBody {
	s.Data = v
	return s
}

func (s *GetDevicePurifyJobByJobIdResponseBody) SetErrorMessage(v string) *GetDevicePurifyJobByJobIdResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDevicePurifyJobByJobIdResponseBody) SetCode(v string) *GetDevicePurifyJobByJobIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetDevicePurifyJobByJobIdResponseBody) SetSuccess(v bool) *GetDevicePurifyJobByJobIdResponseBody {
	s.Success = &v
	return s
}

type GetDevicePurifyJobByJobIdResponseBodyData struct {
	Status               *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	DeviceName           *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	UserId               *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	PlanId               *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	EndTime              *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime            *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	PurifyRecordIndexUrl *string `json:"PurifyRecordIndexUrl,omitempty" xml:"PurifyRecordIndexUrl,omitempty"`
	ProductKey           *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	PurifyRecordNameUrl  *string `json:"PurifyRecordNameUrl,omitempty" xml:"PurifyRecordNameUrl,omitempty"`
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	IotId                *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	TenantId             *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetDevicePurifyJobByJobIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDevicePurifyJobByJobIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDevicePurifyJobByJobIdResponseBodyData) SetStatus(v int32) *GetDevicePurifyJobByJobIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDevicePurifyJobByJobIdResponseBodyData) SetDeviceName(v string) *GetDevicePurifyJobByJobIdResponseBodyData {
	s.DeviceName = &v
	return s
}

func (s *GetDevicePurifyJobByJobIdResponseBodyData) SetUserId(v string) *GetDevicePurifyJobByJobIdResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetDevicePurifyJobByJobIdResponseBodyData) SetPlanId(v string) *GetDevicePurifyJobByJobIdResponseBodyData {
	s.PlanId = &v
	return s
}

func (s *GetDevicePurifyJobByJobIdResponseBodyData) SetEndTime(v int64) *GetDevicePurifyJobByJobIdResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetDevicePurifyJobByJobIdResponseBodyData) SetStartTime(v int64) *GetDevicePurifyJobByJobIdResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetDevicePurifyJobByJobIdResponseBodyData) SetPurifyRecordIndexUrl(v string) *GetDevicePurifyJobByJobIdResponseBodyData {
	s.PurifyRecordIndexUrl = &v
	return s
}

func (s *GetDevicePurifyJobByJobIdResponseBodyData) SetProductKey(v string) *GetDevicePurifyJobByJobIdResponseBodyData {
	s.ProductKey = &v
	return s
}

func (s *GetDevicePurifyJobByJobIdResponseBodyData) SetPurifyRecordNameUrl(v string) *GetDevicePurifyJobByJobIdResponseBodyData {
	s.PurifyRecordNameUrl = &v
	return s
}

func (s *GetDevicePurifyJobByJobIdResponseBodyData) SetJobId(v string) *GetDevicePurifyJobByJobIdResponseBodyData {
	s.JobId = &v
	return s
}

func (s *GetDevicePurifyJobByJobIdResponseBodyData) SetIotId(v string) *GetDevicePurifyJobByJobIdResponseBodyData {
	s.IotId = &v
	return s
}

func (s *GetDevicePurifyJobByJobIdResponseBodyData) SetTenantId(v string) *GetDevicePurifyJobByJobIdResponseBodyData {
	s.TenantId = &v
	return s
}

type GetDevicePurifyJobByJobIdResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDevicePurifyJobByJobIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDevicePurifyJobByJobIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDevicePurifyJobByJobIdResponse) GoString() string {
	return s.String()
}

func (s *GetDevicePurifyJobByJobIdResponse) SetHeaders(v map[string]*string) *GetDevicePurifyJobByJobIdResponse {
	s.Headers = v
	return s
}

func (s *GetDevicePurifyJobByJobIdResponse) SetBody(v *GetDevicePurifyJobByJobIdResponseBody) *GetDevicePurifyJobByJobIdResponse {
	s.Body = v
	return s
}

type GetModelDetailRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	AlgorithmId *int64  `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetModelDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetModelDetailRequest) GoString() string {
	return s.String()
}

func (s *GetModelDetailRequest) SetApiProduct(v string) *GetModelDetailRequest {
	s.ApiProduct = &v
	return s
}

func (s *GetModelDetailRequest) SetApiRevision(v string) *GetModelDetailRequest {
	s.ApiRevision = &v
	return s
}

func (s *GetModelDetailRequest) SetAlgorithmId(v int64) *GetModelDetailRequest {
	s.AlgorithmId = &v
	return s
}

func (s *GetModelDetailRequest) SetVersion(v string) *GetModelDetailRequest {
	s.Version = &v
	return s
}

type GetModelDetailResponseBody struct {
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetModelDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetModelDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelDetailResponseBody) SetRequestId(v string) *GetModelDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelDetailResponseBody) SetData(v map[string]interface{}) *GetModelDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetModelDetailResponseBody) SetErrorMessage(v string) *GetModelDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetModelDetailResponseBody) SetCode(v string) *GetModelDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetModelDetailResponseBody) SetSuccess(v bool) *GetModelDetailResponseBody {
	s.Success = &v
	return s
}

type GetModelDetailResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetModelDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetModelDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetModelDetailResponse) GoString() string {
	return s.String()
}

func (s *GetModelDetailResponse) SetHeaders(v map[string]*string) *GetModelDetailResponse {
	s.Headers = v
	return s
}

func (s *GetModelDetailResponse) SetBody(v *GetModelDetailResponseBody) *GetModelDetailResponse {
	s.Body = v
	return s
}

type GetModelDetailByIdRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	ModelId     *int64  `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
}

func (s GetModelDetailByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetModelDetailByIdRequest) GoString() string {
	return s.String()
}

func (s *GetModelDetailByIdRequest) SetApiProduct(v string) *GetModelDetailByIdRequest {
	s.ApiProduct = &v
	return s
}

func (s *GetModelDetailByIdRequest) SetApiRevision(v string) *GetModelDetailByIdRequest {
	s.ApiRevision = &v
	return s
}

func (s *GetModelDetailByIdRequest) SetModelId(v int64) *GetModelDetailByIdRequest {
	s.ModelId = &v
	return s
}

type GetModelDetailByIdResponseBody struct {
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetModelDetailByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetModelDetailByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelDetailByIdResponseBody) SetRequestId(v string) *GetModelDetailByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelDetailByIdResponseBody) SetData(v map[string]interface{}) *GetModelDetailByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetModelDetailByIdResponseBody) SetErrorMessage(v string) *GetModelDetailByIdResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetModelDetailByIdResponseBody) SetCode(v string) *GetModelDetailByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetModelDetailByIdResponseBody) SetSuccess(v bool) *GetModelDetailByIdResponseBody {
	s.Success = &v
	return s
}

type GetModelDetailByIdResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetModelDetailByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetModelDetailByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetModelDetailByIdResponse) GoString() string {
	return s.String()
}

func (s *GetModelDetailByIdResponse) SetHeaders(v map[string]*string) *GetModelDetailByIdResponse {
	s.Headers = v
	return s
}

func (s *GetModelDetailByIdResponse) SetBody(v *GetModelDetailByIdResponseBody) *GetModelDetailByIdResponse {
	s.Body = v
	return s
}

type GetModelOssPolicyRequest struct {
	ApiProduct           *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision          *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	AlgorithmId          *int64  `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	Hardware             *string `json:"Hardware,omitempty" xml:"Hardware,omitempty"`
	ModelPackageStandard *string `json:"ModelPackageStandard,omitempty" xml:"ModelPackageStandard,omitempty"`
}

func (s GetModelOssPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetModelOssPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetModelOssPolicyRequest) SetApiProduct(v string) *GetModelOssPolicyRequest {
	s.ApiProduct = &v
	return s
}

func (s *GetModelOssPolicyRequest) SetApiRevision(v string) *GetModelOssPolicyRequest {
	s.ApiRevision = &v
	return s
}

func (s *GetModelOssPolicyRequest) SetAlgorithmId(v int64) *GetModelOssPolicyRequest {
	s.AlgorithmId = &v
	return s
}

func (s *GetModelOssPolicyRequest) SetHardware(v string) *GetModelOssPolicyRequest {
	s.Hardware = &v
	return s
}

func (s *GetModelOssPolicyRequest) SetModelPackageStandard(v string) *GetModelOssPolicyRequest {
	s.ModelPackageStandard = &v
	return s
}

type GetModelOssPolicyResponseBody struct {
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetModelOssPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetModelOssPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelOssPolicyResponseBody) SetRequestId(v string) *GetModelOssPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelOssPolicyResponseBody) SetData(v map[string]interface{}) *GetModelOssPolicyResponseBody {
	s.Data = v
	return s
}

func (s *GetModelOssPolicyResponseBody) SetErrorMessage(v string) *GetModelOssPolicyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetModelOssPolicyResponseBody) SetCode(v string) *GetModelOssPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *GetModelOssPolicyResponseBody) SetSuccess(v bool) *GetModelOssPolicyResponseBody {
	s.Success = &v
	return s
}

type GetModelOssPolicyResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetModelOssPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetModelOssPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetModelOssPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetModelOssPolicyResponse) SetHeaders(v map[string]*string) *GetModelOssPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetModelOssPolicyResponse) SetBody(v *GetModelOssPolicyResponseBody) *GetModelOssPolicyResponse {
	s.Body = v
	return s
}

type GetPictureInfoWithVectorIdsRequest struct {
	ApiProduct   *string   `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision  *string   `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	VectorIdList []*string `json:"VectorIdList,omitempty" xml:"VectorIdList,omitempty" type:"Repeated"`
}

func (s GetPictureInfoWithVectorIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPictureInfoWithVectorIdsRequest) GoString() string {
	return s.String()
}

func (s *GetPictureInfoWithVectorIdsRequest) SetApiProduct(v string) *GetPictureInfoWithVectorIdsRequest {
	s.ApiProduct = &v
	return s
}

func (s *GetPictureInfoWithVectorIdsRequest) SetApiRevision(v string) *GetPictureInfoWithVectorIdsRequest {
	s.ApiRevision = &v
	return s
}

func (s *GetPictureInfoWithVectorIdsRequest) SetVectorIdList(v []*string) *GetPictureInfoWithVectorIdsRequest {
	s.VectorIdList = v
	return s
}

type GetPictureInfoWithVectorIdsResponseBody struct {
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         []*GetPictureInfoWithVectorIdsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMessage *string                                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPictureInfoWithVectorIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPictureInfoWithVectorIdsResponseBody) GoString() string {
	return s.String()
}

func (s *GetPictureInfoWithVectorIdsResponseBody) SetRequestId(v string) *GetPictureInfoWithVectorIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPictureInfoWithVectorIdsResponseBody) SetData(v []*GetPictureInfoWithVectorIdsResponseBodyData) *GetPictureInfoWithVectorIdsResponseBody {
	s.Data = v
	return s
}

func (s *GetPictureInfoWithVectorIdsResponseBody) SetErrorMessage(v string) *GetPictureInfoWithVectorIdsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPictureInfoWithVectorIdsResponseBody) SetCode(v string) *GetPictureInfoWithVectorIdsResponseBody {
	s.Code = &v
	return s
}

func (s *GetPictureInfoWithVectorIdsResponseBody) SetSuccess(v bool) *GetPictureInfoWithVectorIdsResponseBody {
	s.Success = &v
	return s
}

type GetPictureInfoWithVectorIdsResponseBodyData struct {
	PicUrl       *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	GatewayIotId *string `json:"GatewayIotId,omitempty" xml:"GatewayIotId,omitempty"`
	IotId        *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
}

func (s GetPictureInfoWithVectorIdsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPictureInfoWithVectorIdsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPictureInfoWithVectorIdsResponseBodyData) SetPicUrl(v string) *GetPictureInfoWithVectorIdsResponseBodyData {
	s.PicUrl = &v
	return s
}

func (s *GetPictureInfoWithVectorIdsResponseBodyData) SetGatewayIotId(v string) *GetPictureInfoWithVectorIdsResponseBodyData {
	s.GatewayIotId = &v
	return s
}

func (s *GetPictureInfoWithVectorIdsResponseBodyData) SetIotId(v string) *GetPictureInfoWithVectorIdsResponseBodyData {
	s.IotId = &v
	return s
}

type GetPictureInfoWithVectorIdsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPictureInfoWithVectorIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPictureInfoWithVectorIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPictureInfoWithVectorIdsResponse) GoString() string {
	return s.String()
}

func (s *GetPictureInfoWithVectorIdsResponse) SetHeaders(v map[string]*string) *GetPictureInfoWithVectorIdsResponse {
	s.Headers = v
	return s
}

func (s *GetPictureInfoWithVectorIdsResponse) SetBody(v *GetPictureInfoWithVectorIdsResponseBody) *GetPictureInfoWithVectorIdsResponse {
	s.Body = v
	return s
}

type GetPictureWithVectorIdRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	VectorId    *string `json:"VectorId,omitempty" xml:"VectorId,omitempty"`
}

func (s GetPictureWithVectorIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPictureWithVectorIdRequest) GoString() string {
	return s.String()
}

func (s *GetPictureWithVectorIdRequest) SetApiProduct(v string) *GetPictureWithVectorIdRequest {
	s.ApiProduct = &v
	return s
}

func (s *GetPictureWithVectorIdRequest) SetApiRevision(v string) *GetPictureWithVectorIdRequest {
	s.ApiRevision = &v
	return s
}

func (s *GetPictureWithVectorIdRequest) SetVectorId(v string) *GetPictureWithVectorIdRequest {
	s.VectorId = &v
	return s
}

type GetPictureWithVectorIdResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPictureWithVectorIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPictureWithVectorIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetPictureWithVectorIdResponseBody) SetRequestId(v string) *GetPictureWithVectorIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPictureWithVectorIdResponseBody) SetData(v string) *GetPictureWithVectorIdResponseBody {
	s.Data = &v
	return s
}

func (s *GetPictureWithVectorIdResponseBody) SetErrorMessage(v string) *GetPictureWithVectorIdResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPictureWithVectorIdResponseBody) SetCode(v string) *GetPictureWithVectorIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetPictureWithVectorIdResponseBody) SetSuccess(v bool) *GetPictureWithVectorIdResponseBody {
	s.Success = &v
	return s
}

type GetPictureWithVectorIdResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPictureWithVectorIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPictureWithVectorIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPictureWithVectorIdResponse) GoString() string {
	return s.String()
}

func (s *GetPictureWithVectorIdResponse) SetHeaders(v map[string]*string) *GetPictureWithVectorIdResponse {
	s.Headers = v
	return s
}

func (s *GetPictureWithVectorIdResponse) SetBody(v *GetPictureWithVectorIdResponseBody) *GetPictureWithVectorIdResponse {
	s.Body = v
	return s
}

type ListAlgorithmsByPageRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	NamePattern *string `json:"NamePattern,omitempty" xml:"NamePattern,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s ListAlgorithmsByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlgorithmsByPageRequest) GoString() string {
	return s.String()
}

func (s *ListAlgorithmsByPageRequest) SetApiProduct(v string) *ListAlgorithmsByPageRequest {
	s.ApiProduct = &v
	return s
}

func (s *ListAlgorithmsByPageRequest) SetApiRevision(v string) *ListAlgorithmsByPageRequest {
	s.ApiRevision = &v
	return s
}

func (s *ListAlgorithmsByPageRequest) SetNamePattern(v string) *ListAlgorithmsByPageRequest {
	s.NamePattern = &v
	return s
}

func (s *ListAlgorithmsByPageRequest) SetPageSize(v int32) *ListAlgorithmsByPageRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlgorithmsByPageRequest) SetCurrentPage(v int32) *ListAlgorithmsByPageRequest {
	s.CurrentPage = &v
	return s
}

type ListAlgorithmsByPageResponseBody struct {
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAlgorithmsByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlgorithmsByPageResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlgorithmsByPageResponseBody) SetRequestId(v string) *ListAlgorithmsByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlgorithmsByPageResponseBody) SetData(v map[string]interface{}) *ListAlgorithmsByPageResponseBody {
	s.Data = v
	return s
}

func (s *ListAlgorithmsByPageResponseBody) SetErrorMessage(v string) *ListAlgorithmsByPageResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListAlgorithmsByPageResponseBody) SetCode(v string) *ListAlgorithmsByPageResponseBody {
	s.Code = &v
	return s
}

func (s *ListAlgorithmsByPageResponseBody) SetSuccess(v bool) *ListAlgorithmsByPageResponseBody {
	s.Success = &v
	return s
}

type ListAlgorithmsByPageResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAlgorithmsByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAlgorithmsByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlgorithmsByPageResponse) GoString() string {
	return s.String()
}

func (s *ListAlgorithmsByPageResponse) SetHeaders(v map[string]*string) *ListAlgorithmsByPageResponse {
	s.Headers = v
	return s
}

func (s *ListAlgorithmsByPageResponse) SetBody(v *ListAlgorithmsByPageResponseBody) *ListAlgorithmsByPageResponse {
	s.Body = v
	return s
}

type ListDeployTaskByModelIdAndDevicesRequest struct {
	ApiProduct  *string                                               `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string                                               `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	ModelId     *int32                                                `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	DeviceList  []*ListDeployTaskByModelIdAndDevicesRequestDeviceList `json:"DeviceList,omitempty" xml:"DeviceList,omitempty" type:"Repeated"`
}

func (s ListDeployTaskByModelIdAndDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeployTaskByModelIdAndDevicesRequest) GoString() string {
	return s.String()
}

func (s *ListDeployTaskByModelIdAndDevicesRequest) SetApiProduct(v string) *ListDeployTaskByModelIdAndDevicesRequest {
	s.ApiProduct = &v
	return s
}

func (s *ListDeployTaskByModelIdAndDevicesRequest) SetApiRevision(v string) *ListDeployTaskByModelIdAndDevicesRequest {
	s.ApiRevision = &v
	return s
}

func (s *ListDeployTaskByModelIdAndDevicesRequest) SetModelId(v int32) *ListDeployTaskByModelIdAndDevicesRequest {
	s.ModelId = &v
	return s
}

func (s *ListDeployTaskByModelIdAndDevicesRequest) SetDeviceList(v []*ListDeployTaskByModelIdAndDevicesRequestDeviceList) *ListDeployTaskByModelIdAndDevicesRequest {
	s.DeviceList = v
	return s
}

type ListDeployTaskByModelIdAndDevicesRequestDeviceList struct {
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
}

func (s ListDeployTaskByModelIdAndDevicesRequestDeviceList) String() string {
	return tea.Prettify(s)
}

func (s ListDeployTaskByModelIdAndDevicesRequestDeviceList) GoString() string {
	return s.String()
}

func (s *ListDeployTaskByModelIdAndDevicesRequestDeviceList) SetProductKey(v string) *ListDeployTaskByModelIdAndDevicesRequestDeviceList {
	s.ProductKey = &v
	return s
}

func (s *ListDeployTaskByModelIdAndDevicesRequestDeviceList) SetDeviceName(v string) *ListDeployTaskByModelIdAndDevicesRequestDeviceList {
	s.DeviceName = &v
	return s
}

func (s *ListDeployTaskByModelIdAndDevicesRequestDeviceList) SetIotId(v string) *ListDeployTaskByModelIdAndDevicesRequestDeviceList {
	s.IotId = &v
	return s
}

type ListDeployTaskByModelIdAndDevicesResponseBody struct {
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDeployTaskByModelIdAndDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeployTaskByModelIdAndDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeployTaskByModelIdAndDevicesResponseBody) SetRequestId(v string) *ListDeployTaskByModelIdAndDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeployTaskByModelIdAndDevicesResponseBody) SetData(v map[string]interface{}) *ListDeployTaskByModelIdAndDevicesResponseBody {
	s.Data = v
	return s
}

func (s *ListDeployTaskByModelIdAndDevicesResponseBody) SetErrorMessage(v string) *ListDeployTaskByModelIdAndDevicesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDeployTaskByModelIdAndDevicesResponseBody) SetCode(v string) *ListDeployTaskByModelIdAndDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListDeployTaskByModelIdAndDevicesResponseBody) SetSuccess(v bool) *ListDeployTaskByModelIdAndDevicesResponseBody {
	s.Success = &v
	return s
}

type ListDeployTaskByModelIdAndDevicesResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDeployTaskByModelIdAndDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeployTaskByModelIdAndDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeployTaskByModelIdAndDevicesResponse) GoString() string {
	return s.String()
}

func (s *ListDeployTaskByModelIdAndDevicesResponse) SetHeaders(v map[string]*string) *ListDeployTaskByModelIdAndDevicesResponse {
	s.Headers = v
	return s
}

func (s *ListDeployTaskByModelIdAndDevicesResponse) SetBody(v *ListDeployTaskByModelIdAndDevicesResponseBody) *ListDeployTaskByModelIdAndDevicesResponse {
	s.Body = v
	return s
}

type ListDeployTaskByPageRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s ListDeployTaskByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeployTaskByPageRequest) GoString() string {
	return s.String()
}

func (s *ListDeployTaskByPageRequest) SetApiProduct(v string) *ListDeployTaskByPageRequest {
	s.ApiProduct = &v
	return s
}

func (s *ListDeployTaskByPageRequest) SetApiRevision(v string) *ListDeployTaskByPageRequest {
	s.ApiRevision = &v
	return s
}

func (s *ListDeployTaskByPageRequest) SetPageSize(v int32) *ListDeployTaskByPageRequest {
	s.PageSize = &v
	return s
}

func (s *ListDeployTaskByPageRequest) SetCurrentPage(v int32) *ListDeployTaskByPageRequest {
	s.CurrentPage = &v
	return s
}

type ListDeployTaskByPageResponseBody struct {
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDeployTaskByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeployTaskByPageResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeployTaskByPageResponseBody) SetRequestId(v string) *ListDeployTaskByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeployTaskByPageResponseBody) SetData(v map[string]interface{}) *ListDeployTaskByPageResponseBody {
	s.Data = v
	return s
}

func (s *ListDeployTaskByPageResponseBody) SetErrorMessage(v string) *ListDeployTaskByPageResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDeployTaskByPageResponseBody) SetCode(v string) *ListDeployTaskByPageResponseBody {
	s.Code = &v
	return s
}

func (s *ListDeployTaskByPageResponseBody) SetSuccess(v bool) *ListDeployTaskByPageResponseBody {
	s.Success = &v
	return s
}

type ListDeployTaskByPageResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDeployTaskByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeployTaskByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeployTaskByPageResponse) GoString() string {
	return s.String()
}

func (s *ListDeployTaskByPageResponse) SetHeaders(v map[string]*string) *ListDeployTaskByPageResponse {
	s.Headers = v
	return s
}

func (s *ListDeployTaskByPageResponse) SetBody(v *ListDeployTaskByPageResponseBody) *ListDeployTaskByPageResponse {
	s.Body = v
	return s
}

type ListModelsByPageRequest struct {
	ApiProduct           *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision          *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	AlgorithmId          *int64  `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	SizeType             *string `json:"SizeType,omitempty" xml:"SizeType,omitempty"`
	ModelPackageStandard *string `json:"ModelPackageStandard,omitempty" xml:"ModelPackageStandard,omitempty"`
	Hardware             *string `json:"Hardware,omitempty" xml:"Hardware,omitempty"`
	NamePattern          *string `json:"NamePattern,omitempty" xml:"NamePattern,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage          *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s ListModelsByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s ListModelsByPageRequest) GoString() string {
	return s.String()
}

func (s *ListModelsByPageRequest) SetApiProduct(v string) *ListModelsByPageRequest {
	s.ApiProduct = &v
	return s
}

func (s *ListModelsByPageRequest) SetApiRevision(v string) *ListModelsByPageRequest {
	s.ApiRevision = &v
	return s
}

func (s *ListModelsByPageRequest) SetAlgorithmId(v int64) *ListModelsByPageRequest {
	s.AlgorithmId = &v
	return s
}

func (s *ListModelsByPageRequest) SetSizeType(v string) *ListModelsByPageRequest {
	s.SizeType = &v
	return s
}

func (s *ListModelsByPageRequest) SetModelPackageStandard(v string) *ListModelsByPageRequest {
	s.ModelPackageStandard = &v
	return s
}

func (s *ListModelsByPageRequest) SetHardware(v string) *ListModelsByPageRequest {
	s.Hardware = &v
	return s
}

func (s *ListModelsByPageRequest) SetNamePattern(v string) *ListModelsByPageRequest {
	s.NamePattern = &v
	return s
}

func (s *ListModelsByPageRequest) SetPageSize(v int32) *ListModelsByPageRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelsByPageRequest) SetCurrentPage(v int32) *ListModelsByPageRequest {
	s.CurrentPage = &v
	return s
}

type ListModelsByPageResponseBody struct {
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListModelsByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListModelsByPageResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelsByPageResponseBody) SetRequestId(v string) *ListModelsByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelsByPageResponseBody) SetData(v map[string]interface{}) *ListModelsByPageResponseBody {
	s.Data = v
	return s
}

func (s *ListModelsByPageResponseBody) SetErrorMessage(v string) *ListModelsByPageResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListModelsByPageResponseBody) SetCode(v string) *ListModelsByPageResponseBody {
	s.Code = &v
	return s
}

func (s *ListModelsByPageResponseBody) SetSuccess(v bool) *ListModelsByPageResponseBody {
	s.Success = &v
	return s
}

type ListModelsByPageResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListModelsByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListModelsByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s ListModelsByPageResponse) GoString() string {
	return s.String()
}

func (s *ListModelsByPageResponse) SetHeaders(v map[string]*string) *ListModelsByPageResponse {
	s.Headers = v
	return s
}

func (s *ListModelsByPageResponse) SetBody(v *ListModelsByPageResponseBody) *ListModelsByPageResponse {
	s.Body = v
	return s
}

type ListModelVersionsByPageRequest struct {
	ApiProduct           *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision          *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	AlgorithmName        *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	SizeType             *string `json:"SizeType,omitempty" xml:"SizeType,omitempty"`
	ModelPackageStandard *string `json:"ModelPackageStandard,omitempty" xml:"ModelPackageStandard,omitempty"`
	Hardware             *string `json:"Hardware,omitempty" xml:"Hardware,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage          *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s ListModelVersionsByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s ListModelVersionsByPageRequest) GoString() string {
	return s.String()
}

func (s *ListModelVersionsByPageRequest) SetApiProduct(v string) *ListModelVersionsByPageRequest {
	s.ApiProduct = &v
	return s
}

func (s *ListModelVersionsByPageRequest) SetApiRevision(v string) *ListModelVersionsByPageRequest {
	s.ApiRevision = &v
	return s
}

func (s *ListModelVersionsByPageRequest) SetAlgorithmName(v string) *ListModelVersionsByPageRequest {
	s.AlgorithmName = &v
	return s
}

func (s *ListModelVersionsByPageRequest) SetSizeType(v string) *ListModelVersionsByPageRequest {
	s.SizeType = &v
	return s
}

func (s *ListModelVersionsByPageRequest) SetModelPackageStandard(v string) *ListModelVersionsByPageRequest {
	s.ModelPackageStandard = &v
	return s
}

func (s *ListModelVersionsByPageRequest) SetHardware(v string) *ListModelVersionsByPageRequest {
	s.Hardware = &v
	return s
}

func (s *ListModelVersionsByPageRequest) SetPageSize(v int32) *ListModelVersionsByPageRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelVersionsByPageRequest) SetCurrentPage(v int32) *ListModelVersionsByPageRequest {
	s.CurrentPage = &v
	return s
}

type ListModelVersionsByPageResponseBody struct {
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListModelVersionsByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListModelVersionsByPageResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelVersionsByPageResponseBody) SetRequestId(v string) *ListModelVersionsByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelVersionsByPageResponseBody) SetData(v map[string]interface{}) *ListModelVersionsByPageResponseBody {
	s.Data = v
	return s
}

func (s *ListModelVersionsByPageResponseBody) SetErrorMessage(v string) *ListModelVersionsByPageResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListModelVersionsByPageResponseBody) SetCode(v string) *ListModelVersionsByPageResponseBody {
	s.Code = &v
	return s
}

func (s *ListModelVersionsByPageResponseBody) SetSuccess(v bool) *ListModelVersionsByPageResponseBody {
	s.Success = &v
	return s
}

type ListModelVersionsByPageResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListModelVersionsByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListModelVersionsByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s ListModelVersionsByPageResponse) GoString() string {
	return s.String()
}

func (s *ListModelVersionsByPageResponse) SetHeaders(v map[string]*string) *ListModelVersionsByPageResponse {
	s.Headers = v
	return s
}

func (s *ListModelVersionsByPageResponse) SetBody(v *ListModelVersionsByPageResponseBody) *ListModelVersionsByPageResponse {
	s.Body = v
	return s
}

type PictureSearchPictureRequest struct {
	ApiProduct    *string  `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string  `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	AppInstanceId *string  `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	PageSize      *int32   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage   *int32   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	SearchPicUrl  *string  `json:"SearchPicUrl,omitempty" xml:"SearchPicUrl,omitempty"`
	StartTime     *int64   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *int64   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Threshold     *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	ContainPicUrl *bool    `json:"ContainPicUrl,omitempty" xml:"ContainPicUrl,omitempty"`
}

func (s PictureSearchPictureRequest) String() string {
	return tea.Prettify(s)
}

func (s PictureSearchPictureRequest) GoString() string {
	return s.String()
}

func (s *PictureSearchPictureRequest) SetApiProduct(v string) *PictureSearchPictureRequest {
	s.ApiProduct = &v
	return s
}

func (s *PictureSearchPictureRequest) SetApiRevision(v string) *PictureSearchPictureRequest {
	s.ApiRevision = &v
	return s
}

func (s *PictureSearchPictureRequest) SetAppInstanceId(v string) *PictureSearchPictureRequest {
	s.AppInstanceId = &v
	return s
}

func (s *PictureSearchPictureRequest) SetPageSize(v int32) *PictureSearchPictureRequest {
	s.PageSize = &v
	return s
}

func (s *PictureSearchPictureRequest) SetCurrentPage(v int32) *PictureSearchPictureRequest {
	s.CurrentPage = &v
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

func (s *PictureSearchPictureRequest) SetEndTime(v int64) *PictureSearchPictureRequest {
	s.EndTime = &v
	return s
}

func (s *PictureSearchPictureRequest) SetThreshold(v float32) *PictureSearchPictureRequest {
	s.Threshold = &v
	return s
}

func (s *PictureSearchPictureRequest) SetContainPicUrl(v bool) *PictureSearchPictureRequest {
	s.ContainPicUrl = &v
	return s
}

type PictureSearchPictureResponseBody struct {
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *PictureSearchPictureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                               `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PictureSearchPictureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PictureSearchPictureResponseBody) GoString() string {
	return s.String()
}

func (s *PictureSearchPictureResponseBody) SetRequestId(v string) *PictureSearchPictureResponseBody {
	s.RequestId = &v
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

func (s *PictureSearchPictureResponseBody) SetCode(v string) *PictureSearchPictureResponseBody {
	s.Code = &v
	return s
}

func (s *PictureSearchPictureResponseBody) SetSuccess(v bool) *PictureSearchPictureResponseBody {
	s.Success = &v
	return s
}

type PictureSearchPictureResponseBodyData struct {
	CurrentPage *int32                                          `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageData    []*PictureSearchPictureResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	PageSize    *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total       *int32                                          `json:"Total,omitempty" xml:"Total,omitempty"`
	PageCount   *int32                                          `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
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

func (s *PictureSearchPictureResponseBodyData) SetPageCount(v int32) *PictureSearchPictureResponseBodyData {
	s.PageCount = &v
	return s
}

type PictureSearchPictureResponseBodyDataPageData struct {
	PicUrl       *string  `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	EventTime    *int64   `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	GatewayIotId *string  `json:"GatewayIotId,omitempty" xml:"GatewayIotId,omitempty"`
	VectorId     *string  `json:"VectorId,omitempty" xml:"VectorId,omitempty"`
	Threshold    *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	IotId        *string  `json:"IotId,omitempty" xml:"IotId,omitempty"`
}

func (s PictureSearchPictureResponseBodyDataPageData) String() string {
	return tea.Prettify(s)
}

func (s PictureSearchPictureResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *PictureSearchPictureResponseBodyDataPageData) SetPicUrl(v string) *PictureSearchPictureResponseBodyDataPageData {
	s.PicUrl = &v
	return s
}

func (s *PictureSearchPictureResponseBodyDataPageData) SetEventTime(v int64) *PictureSearchPictureResponseBodyDataPageData {
	s.EventTime = &v
	return s
}

func (s *PictureSearchPictureResponseBodyDataPageData) SetGatewayIotId(v string) *PictureSearchPictureResponseBodyDataPageData {
	s.GatewayIotId = &v
	return s
}

func (s *PictureSearchPictureResponseBodyDataPageData) SetVectorId(v string) *PictureSearchPictureResponseBodyDataPageData {
	s.VectorId = &v
	return s
}

func (s *PictureSearchPictureResponseBodyDataPageData) SetThreshold(v float32) *PictureSearchPictureResponseBodyDataPageData {
	s.Threshold = &v
	return s
}

func (s *PictureSearchPictureResponseBodyDataPageData) SetIotId(v string) *PictureSearchPictureResponseBodyDataPageData {
	s.IotId = &v
	return s
}

type PictureSearchPictureResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PictureSearchPictureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *PictureSearchPictureResponse) SetBody(v *PictureSearchPictureResponseBody) *PictureSearchPictureResponse {
	s.Body = v
	return s
}

type QueryAIActionRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	PlanId      *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s QueryAIActionRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAIActionRequest) GoString() string {
	return s.String()
}

func (s *QueryAIActionRequest) SetApiProduct(v string) *QueryAIActionRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryAIActionRequest) SetApiRevision(v string) *QueryAIActionRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryAIActionRequest) SetPlanId(v string) *QueryAIActionRequest {
	s.PlanId = &v
	return s
}

type QueryAIActionResponseBody struct {
	RequestId    *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage *string                          `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         []*QueryAIActionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s QueryAIActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAIActionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAIActionResponseBody) SetRequestId(v string) *QueryAIActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAIActionResponseBody) SetSuccess(v bool) *QueryAIActionResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAIActionResponseBody) SetErrorMessage(v string) *QueryAIActionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryAIActionResponseBody) SetCode(v string) *QueryAIActionResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAIActionResponseBody) SetData(v []*QueryAIActionResponseBodyData) *QueryAIActionResponseBody {
	s.Data = v
	return s
}

type QueryAIActionResponseBodyData struct {
	ActionId         *string `json:"ActionId,omitempty" xml:"ActionId,omitempty"`
	PlanId           *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	ActionTemplateId *string `json:"ActionTemplateId,omitempty" xml:"ActionTemplateId,omitempty"`
	ActionIndex      *int32  `json:"ActionIndex,omitempty" xml:"ActionIndex,omitempty"`
	Algo             *string `json:"Algo,omitempty" xml:"Algo,omitempty"`
	Action           *string `json:"Action,omitempty" xml:"Action,omitempty"`
	AlgoConfig       *string `json:"AlgoConfig,omitempty" xml:"AlgoConfig,omitempty"`
	ActionConfig     *string `json:"ActionConfig,omitempty" xml:"ActionConfig,omitempty"`
}

func (s QueryAIActionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAIActionResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAIActionResponseBodyData) SetActionId(v string) *QueryAIActionResponseBodyData {
	s.ActionId = &v
	return s
}

func (s *QueryAIActionResponseBodyData) SetPlanId(v string) *QueryAIActionResponseBodyData {
	s.PlanId = &v
	return s
}

func (s *QueryAIActionResponseBodyData) SetActionTemplateId(v string) *QueryAIActionResponseBodyData {
	s.ActionTemplateId = &v
	return s
}

func (s *QueryAIActionResponseBodyData) SetActionIndex(v int32) *QueryAIActionResponseBodyData {
	s.ActionIndex = &v
	return s
}

func (s *QueryAIActionResponseBodyData) SetAlgo(v string) *QueryAIActionResponseBodyData {
	s.Algo = &v
	return s
}

func (s *QueryAIActionResponseBodyData) SetAction(v string) *QueryAIActionResponseBodyData {
	s.Action = &v
	return s
}

func (s *QueryAIActionResponseBodyData) SetAlgoConfig(v string) *QueryAIActionResponseBodyData {
	s.AlgoConfig = &v
	return s
}

func (s *QueryAIActionResponseBodyData) SetActionConfig(v string) *QueryAIActionResponseBodyData {
	s.ActionConfig = &v
	return s
}

type QueryAIActionResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAIActionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAIActionResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAIActionResponse) GoString() string {
	return s.String()
}

func (s *QueryAIActionResponse) SetHeaders(v map[string]*string) *QueryAIActionResponse {
	s.Headers = v
	return s
}

func (s *QueryAIActionResponse) SetBody(v *QueryAIActionResponseBody) *QueryAIActionResponse {
	s.Body = v
	return s
}

type QueryAIAppRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s QueryAIAppRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAIAppRequest) GoString() string {
	return s.String()
}

func (s *QueryAIAppRequest) SetApiProduct(v string) *QueryAIAppRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryAIAppRequest) SetApiRevision(v string) *QueryAIAppRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryAIAppRequest) SetPageSize(v int32) *QueryAIAppRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAIAppRequest) SetCurrentPage(v int32) *QueryAIAppRequest {
	s.CurrentPage = &v
	return s
}

type QueryAIAppResponseBody struct {
	RequestId    *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage *string                     `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryAIAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryAIAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAIAppResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAIAppResponseBody) SetRequestId(v string) *QueryAIAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAIAppResponseBody) SetSuccess(v bool) *QueryAIAppResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAIAppResponseBody) SetErrorMessage(v string) *QueryAIAppResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryAIAppResponseBody) SetCode(v string) *QueryAIAppResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAIAppResponseBody) SetData(v *QueryAIAppResponseBodyData) *QueryAIAppResponseBody {
	s.Data = v
	return s
}

type QueryAIAppResponseBodyData struct {
	Total       *int64                            `json:"Total,omitempty" xml:"Total,omitempty"`
	PageCount   *int32                            `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	CurrentPage *int32                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	List        []*QueryAIAppResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s QueryAIAppResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAIAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAIAppResponseBodyData) SetTotal(v int64) *QueryAIAppResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryAIAppResponseBodyData) SetPageCount(v int32) *QueryAIAppResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryAIAppResponseBodyData) SetCurrentPage(v int32) *QueryAIAppResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *QueryAIAppResponseBodyData) SetPageSize(v int32) *QueryAIAppResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryAIAppResponseBodyData) SetList(v []*QueryAIAppResponseBodyDataList) *QueryAIAppResponseBodyData {
	s.List = v
	return s
}

type QueryAIAppResponseBodyDataList struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppTemplateId *string `json:"AppTemplateId,omitempty" xml:"AppTemplateId,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Level         *int32  `json:"Level,omitempty" xml:"Level,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s QueryAIAppResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryAIAppResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryAIAppResponseBodyDataList) SetAppId(v string) *QueryAIAppResponseBodyDataList {
	s.AppId = &v
	return s
}

func (s *QueryAIAppResponseBodyDataList) SetAppTemplateId(v string) *QueryAIAppResponseBodyDataList {
	s.AppTemplateId = &v
	return s
}

func (s *QueryAIAppResponseBodyDataList) SetVersion(v string) *QueryAIAppResponseBodyDataList {
	s.Version = &v
	return s
}

func (s *QueryAIAppResponseBodyDataList) SetLevel(v int32) *QueryAIAppResponseBodyDataList {
	s.Level = &v
	return s
}

func (s *QueryAIAppResponseBodyDataList) SetName(v string) *QueryAIAppResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *QueryAIAppResponseBodyDataList) SetDescription(v string) *QueryAIAppResponseBodyDataList {
	s.Description = &v
	return s
}

type QueryAIAppResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAIAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAIAppResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAIAppResponse) GoString() string {
	return s.String()
}

func (s *QueryAIAppResponse) SetHeaders(v map[string]*string) *QueryAIAppResponse {
	s.Headers = v
	return s
}

func (s *QueryAIAppResponse) SetBody(v *QueryAIAppResponseBody) *QueryAIAppResponse {
	s.Body = v
	return s
}

type QueryAIJobsRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	ActionId    *string `json:"ActionId,omitempty" xml:"ActionId,omitempty"`
	IotId       *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s QueryAIJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAIJobsRequest) GoString() string {
	return s.String()
}

func (s *QueryAIJobsRequest) SetApiProduct(v string) *QueryAIJobsRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryAIJobsRequest) SetApiRevision(v string) *QueryAIJobsRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryAIJobsRequest) SetActionId(v string) *QueryAIJobsRequest {
	s.ActionId = &v
	return s
}

func (s *QueryAIJobsRequest) SetIotId(v string) *QueryAIJobsRequest {
	s.IotId = &v
	return s
}

func (s *QueryAIJobsRequest) SetPageSize(v int32) *QueryAIJobsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAIJobsRequest) SetCurrentPage(v int32) *QueryAIJobsRequest {
	s.CurrentPage = &v
	return s
}

type QueryAIJobsResponseBody struct {
	RequestId    *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage *string                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryAIJobsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryAIJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAIJobsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAIJobsResponseBody) SetRequestId(v string) *QueryAIJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAIJobsResponseBody) SetSuccess(v bool) *QueryAIJobsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAIJobsResponseBody) SetErrorMessage(v string) *QueryAIJobsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryAIJobsResponseBody) SetCode(v string) *QueryAIJobsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAIJobsResponseBody) SetData(v *QueryAIJobsResponseBodyData) *QueryAIJobsResponseBody {
	s.Data = v
	return s
}

type QueryAIJobsResponseBodyData struct {
	Total       *int64                             `json:"Total,omitempty" xml:"Total,omitempty"`
	PageCount   *int32                             `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	CurrentPage *int32                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	List        []*QueryAIJobsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s QueryAIJobsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAIJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAIJobsResponseBodyData) SetTotal(v int64) *QueryAIJobsResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryAIJobsResponseBodyData) SetPageCount(v int32) *QueryAIJobsResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryAIJobsResponseBodyData) SetCurrentPage(v int32) *QueryAIJobsResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *QueryAIJobsResponseBodyData) SetPageSize(v int32) *QueryAIJobsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryAIJobsResponseBodyData) SetList(v []*QueryAIJobsResponseBodyDataList) *QueryAIJobsResponseBodyData {
	s.List = v
	return s
}

type QueryAIJobsResponseBodyDataList struct {
	JobId    *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ActionId *string `json:"ActionId,omitempty" xml:"ActionId,omitempty"`
	Status   *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	IotId    *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
}

func (s QueryAIJobsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryAIJobsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryAIJobsResponseBodyDataList) SetJobId(v string) *QueryAIJobsResponseBodyDataList {
	s.JobId = &v
	return s
}

func (s *QueryAIJobsResponseBodyDataList) SetActionId(v string) *QueryAIJobsResponseBodyDataList {
	s.ActionId = &v
	return s
}

func (s *QueryAIJobsResponseBodyDataList) SetStatus(v int32) *QueryAIJobsResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *QueryAIJobsResponseBodyDataList) SetIotId(v string) *QueryAIJobsResponseBodyDataList {
	s.IotId = &v
	return s
}

type QueryAIJobsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAIJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAIJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAIJobsResponse) GoString() string {
	return s.String()
}

func (s *QueryAIJobsResponse) SetHeaders(v map[string]*string) *QueryAIJobsResponse {
	s.Headers = v
	return s
}

func (s *QueryAIJobsResponse) SetBody(v *QueryAIJobsResponseBody) *QueryAIJobsResponse {
	s.Body = v
	return s
}

type QueryAIPlanRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s QueryAIPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAIPlanRequest) GoString() string {
	return s.String()
}

func (s *QueryAIPlanRequest) SetApiProduct(v string) *QueryAIPlanRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryAIPlanRequest) SetApiRevision(v string) *QueryAIPlanRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryAIPlanRequest) SetAppId(v string) *QueryAIPlanRequest {
	s.AppId = &v
	return s
}

func (s *QueryAIPlanRequest) SetPageSize(v int32) *QueryAIPlanRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAIPlanRequest) SetCurrentPage(v int32) *QueryAIPlanRequest {
	s.CurrentPage = &v
	return s
}

type QueryAIPlanResponseBody struct {
	RequestId    *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage *string                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryAIPlanResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryAIPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAIPlanResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAIPlanResponseBody) SetRequestId(v string) *QueryAIPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAIPlanResponseBody) SetSuccess(v bool) *QueryAIPlanResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAIPlanResponseBody) SetErrorMessage(v string) *QueryAIPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryAIPlanResponseBody) SetCode(v string) *QueryAIPlanResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAIPlanResponseBody) SetData(v *QueryAIPlanResponseBodyData) *QueryAIPlanResponseBody {
	s.Data = v
	return s
}

type QueryAIPlanResponseBodyData struct {
	Total       *int64                             `json:"Total,omitempty" xml:"Total,omitempty"`
	PageCount   *int32                             `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	CurrentPage *int32                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	List        []*QueryAIPlanResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s QueryAIPlanResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAIPlanResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAIPlanResponseBodyData) SetTotal(v int64) *QueryAIPlanResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryAIPlanResponseBodyData) SetPageCount(v int32) *QueryAIPlanResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryAIPlanResponseBodyData) SetCurrentPage(v int32) *QueryAIPlanResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *QueryAIPlanResponseBodyData) SetPageSize(v int32) *QueryAIPlanResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryAIPlanResponseBodyData) SetList(v []*QueryAIPlanResponseBodyDataList) *QueryAIPlanResponseBodyData {
	s.List = v
	return s
}

type QueryAIPlanResponseBodyDataList struct {
	PlanId         *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PlanTemplateId *string `json:"PlanTemplateId,omitempty" xml:"PlanTemplateId,omitempty"`
	TriggerEnum    *int32  `json:"TriggerEnum,omitempty" xml:"TriggerEnum,omitempty"`
	IntervalTiming *int32  `json:"IntervalTiming,omitempty" xml:"IntervalTiming,omitempty"`
	PreTiming      *int64  `json:"PreTiming,omitempty" xml:"PreTiming,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s QueryAIPlanResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryAIPlanResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryAIPlanResponseBodyDataList) SetPlanId(v string) *QueryAIPlanResponseBodyDataList {
	s.PlanId = &v
	return s
}

func (s *QueryAIPlanResponseBodyDataList) SetAppId(v string) *QueryAIPlanResponseBodyDataList {
	s.AppId = &v
	return s
}

func (s *QueryAIPlanResponseBodyDataList) SetPlanTemplateId(v string) *QueryAIPlanResponseBodyDataList {
	s.PlanTemplateId = &v
	return s
}

func (s *QueryAIPlanResponseBodyDataList) SetTriggerEnum(v int32) *QueryAIPlanResponseBodyDataList {
	s.TriggerEnum = &v
	return s
}

func (s *QueryAIPlanResponseBodyDataList) SetIntervalTiming(v int32) *QueryAIPlanResponseBodyDataList {
	s.IntervalTiming = &v
	return s
}

func (s *QueryAIPlanResponseBodyDataList) SetPreTiming(v int64) *QueryAIPlanResponseBodyDataList {
	s.PreTiming = &v
	return s
}

func (s *QueryAIPlanResponseBodyDataList) SetDescription(v string) *QueryAIPlanResponseBodyDataList {
	s.Description = &v
	return s
}

type QueryAIPlanResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAIPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAIPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAIPlanResponse) GoString() string {
	return s.String()
}

func (s *QueryAIPlanResponse) SetHeaders(v map[string]*string) *QueryAIPlanResponse {
	s.Headers = v
	return s
}

func (s *QueryAIPlanResponse) SetBody(v *QueryAIPlanResponseBody) *QueryAIPlanResponse {
	s.Body = v
	return s
}

type QueryAIPlanTemplatesRequest struct {
	ApiProduct    *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	AppTemplateId *string `json:"AppTemplateId,omitempty" xml:"AppTemplateId,omitempty"`
	AppVersion    *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
}

func (s QueryAIPlanTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAIPlanTemplatesRequest) GoString() string {
	return s.String()
}

func (s *QueryAIPlanTemplatesRequest) SetApiProduct(v string) *QueryAIPlanTemplatesRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryAIPlanTemplatesRequest) SetApiRevision(v string) *QueryAIPlanTemplatesRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryAIPlanTemplatesRequest) SetAppTemplateId(v string) *QueryAIPlanTemplatesRequest {
	s.AppTemplateId = &v
	return s
}

func (s *QueryAIPlanTemplatesRequest) SetAppVersion(v string) *QueryAIPlanTemplatesRequest {
	s.AppVersion = &v
	return s
}

type QueryAIPlanTemplatesResponseBody struct {
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage *string                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         []*QueryAIPlanTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s QueryAIPlanTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAIPlanTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAIPlanTemplatesResponseBody) SetRequestId(v string) *QueryAIPlanTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAIPlanTemplatesResponseBody) SetSuccess(v bool) *QueryAIPlanTemplatesResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAIPlanTemplatesResponseBody) SetErrorMessage(v string) *QueryAIPlanTemplatesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryAIPlanTemplatesResponseBody) SetCode(v string) *QueryAIPlanTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAIPlanTemplatesResponseBody) SetData(v []*QueryAIPlanTemplatesResponseBodyData) *QueryAIPlanTemplatesResponseBody {
	s.Data = v
	return s
}

type QueryAIPlanTemplatesResponseBodyData struct {
	PlanTemplateId *string `json:"PlanTemplateId,omitempty" xml:"PlanTemplateId,omitempty"`
	AppTemplateId  *string `json:"AppTemplateId,omitempty" xml:"AppTemplateId,omitempty"`
	AppVersion     *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	TriggerEnum    *int32  `json:"TriggerEnum,omitempty" xml:"TriggerEnum,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IntervalTiming *int32  `json:"IntervalTiming,omitempty" xml:"IntervalTiming,omitempty"`
}

func (s QueryAIPlanTemplatesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAIPlanTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAIPlanTemplatesResponseBodyData) SetPlanTemplateId(v string) *QueryAIPlanTemplatesResponseBodyData {
	s.PlanTemplateId = &v
	return s
}

func (s *QueryAIPlanTemplatesResponseBodyData) SetAppTemplateId(v string) *QueryAIPlanTemplatesResponseBodyData {
	s.AppTemplateId = &v
	return s
}

func (s *QueryAIPlanTemplatesResponseBodyData) SetAppVersion(v string) *QueryAIPlanTemplatesResponseBodyData {
	s.AppVersion = &v
	return s
}

func (s *QueryAIPlanTemplatesResponseBodyData) SetTriggerEnum(v int32) *QueryAIPlanTemplatesResponseBodyData {
	s.TriggerEnum = &v
	return s
}

func (s *QueryAIPlanTemplatesResponseBodyData) SetDescription(v string) *QueryAIPlanTemplatesResponseBodyData {
	s.Description = &v
	return s
}

func (s *QueryAIPlanTemplatesResponseBodyData) SetIntervalTiming(v int32) *QueryAIPlanTemplatesResponseBodyData {
	s.IntervalTiming = &v
	return s
}

type QueryAIPlanTemplatesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAIPlanTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAIPlanTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAIPlanTemplatesResponse) GoString() string {
	return s.String()
}

func (s *QueryAIPlanTemplatesResponse) SetHeaders(v map[string]*string) *QueryAIPlanTemplatesResponse {
	s.Headers = v
	return s
}

func (s *QueryAIPlanTemplatesResponse) SetBody(v *QueryAIPlanTemplatesResponseBody) *QueryAIPlanTemplatesResponse {
	s.Body = v
	return s
}

type QueryDeviceEventRequest struct {
	ApiProduct    *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	EventType     *int32  `json:"EventType,omitempty" xml:"EventType,omitempty"`
	BeginTime     *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime       *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
}

func (s QueryDeviceEventRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceEventRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventRequest) SetApiProduct(v string) *QueryDeviceEventRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryDeviceEventRequest) SetApiRevision(v string) *QueryDeviceEventRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryDeviceEventRequest) SetIotId(v string) *QueryDeviceEventRequest {
	s.IotId = &v
	return s
}

func (s *QueryDeviceEventRequest) SetEventType(v int32) *QueryDeviceEventRequest {
	s.EventType = &v
	return s
}

func (s *QueryDeviceEventRequest) SetBeginTime(v int64) *QueryDeviceEventRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryDeviceEventRequest) SetEndTime(v int64) *QueryDeviceEventRequest {
	s.EndTime = &v
	return s
}

func (s *QueryDeviceEventRequest) SetCurrentPage(v int32) *QueryDeviceEventRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryDeviceEventRequest) SetPageSize(v int32) *QueryDeviceEventRequest {
	s.PageSize = &v
	return s
}

func (s *QueryDeviceEventRequest) SetIotInstanceId(v string) *QueryDeviceEventRequest {
	s.IotInstanceId = &v
	return s
}

type QueryDeviceEventResponseBody struct {
	RequestId    *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryDeviceEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDeviceEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceEventResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventResponseBody) SetRequestId(v string) *QueryDeviceEventResponseBody {
	s.RequestId = &v
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

func (s *QueryDeviceEventResponseBody) SetCode(v string) *QueryDeviceEventResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDeviceEventResponseBody) SetSuccess(v bool) *QueryDeviceEventResponseBody {
	s.Success = &v
	return s
}

type QueryDeviceEventResponseBodyData struct {
	List      []*QueryDeviceEventResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageSize  *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageCount *int32                                  `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	Total     *int32                                  `json:"Total,omitempty" xml:"Total,omitempty"`
	Page      *int32                                  `json:"Page,omitempty" xml:"Page,omitempty"`
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

func (s *QueryDeviceEventResponseBodyData) SetPageSize(v int32) *QueryDeviceEventResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryDeviceEventResponseBodyData) SetPageCount(v int32) *QueryDeviceEventResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryDeviceEventResponseBodyData) SetTotal(v int32) *QueryDeviceEventResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryDeviceEventResponseBodyData) SetPage(v int32) *QueryDeviceEventResponseBodyData {
	s.Page = &v
	return s
}

type QueryDeviceEventResponseBodyDataList struct {
	EventId    *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	EventTime  *string `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	EventType  *int32  `json:"EventType,omitempty" xml:"EventType,omitempty"`
	EventPicId *string `json:"EventPicId,omitempty" xml:"EventPicId,omitempty"`
	EventDesc  *string `json:"EventDesc,omitempty" xml:"EventDesc,omitempty"`
	EventData  *string `json:"EventData,omitempty" xml:"EventData,omitempty"`
}

func (s QueryDeviceEventResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceEventResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventResponseBodyDataList) SetEventId(v string) *QueryDeviceEventResponseBodyDataList {
	s.EventId = &v
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

func (s *QueryDeviceEventResponseBodyDataList) SetEventPicId(v string) *QueryDeviceEventResponseBodyDataList {
	s.EventPicId = &v
	return s
}

func (s *QueryDeviceEventResponseBodyDataList) SetEventDesc(v string) *QueryDeviceEventResponseBodyDataList {
	s.EventDesc = &v
	return s
}

func (s *QueryDeviceEventResponseBodyDataList) SetEventData(v string) *QueryDeviceEventResponseBodyDataList {
	s.EventData = &v
	return s
}

type QueryDeviceEventResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDeviceEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryDeviceEventResponse) SetBody(v *QueryDeviceEventResponseBody) *QueryDeviceEventResponse {
	s.Body = v
	return s
}

type QueryDeviceEventPictureRequest struct {
	ApiProduct    *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	EventId       *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
}

func (s QueryDeviceEventPictureRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceEventPictureRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventPictureRequest) SetApiProduct(v string) *QueryDeviceEventPictureRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryDeviceEventPictureRequest) SetApiRevision(v string) *QueryDeviceEventPictureRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryDeviceEventPictureRequest) SetIotId(v string) *QueryDeviceEventPictureRequest {
	s.IotId = &v
	return s
}

func (s *QueryDeviceEventPictureRequest) SetEventId(v string) *QueryDeviceEventPictureRequest {
	s.EventId = &v
	return s
}

func (s *QueryDeviceEventPictureRequest) SetIotInstanceId(v string) *QueryDeviceEventPictureRequest {
	s.IotInstanceId = &v
	return s
}

type QueryDeviceEventPictureResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDeviceEventPictureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceEventPictureResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventPictureResponseBody) SetRequestId(v string) *QueryDeviceEventPictureResponseBody {
	s.RequestId = &v
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

func (s *QueryDeviceEventPictureResponseBody) SetCode(v int32) *QueryDeviceEventPictureResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDeviceEventPictureResponseBody) SetSuccess(v bool) *QueryDeviceEventPictureResponseBody {
	s.Success = &v
	return s
}

type QueryDeviceEventPictureResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDeviceEventPictureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryDeviceEventPictureResponse) SetBody(v *QueryDeviceEventPictureResponseBody) *QueryDeviceEventPictureResponse {
	s.Body = v
	return s
}

type QueryDeviceEventRecordRequest struct {
	ApiProduct    *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	EventId       *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
}

func (s QueryDeviceEventRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceEventRecordRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventRecordRequest) SetApiProduct(v string) *QueryDeviceEventRecordRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryDeviceEventRecordRequest) SetApiRevision(v string) *QueryDeviceEventRecordRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryDeviceEventRecordRequest) SetIotId(v string) *QueryDeviceEventRecordRequest {
	s.IotId = &v
	return s
}

func (s *QueryDeviceEventRecordRequest) SetEventId(v string) *QueryDeviceEventRecordRequest {
	s.EventId = &v
	return s
}

func (s *QueryDeviceEventRecordRequest) SetIotInstanceId(v string) *QueryDeviceEventRecordRequest {
	s.IotInstanceId = &v
	return s
}

type QueryDeviceEventRecordResponseBody struct {
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         []*QueryDeviceEventRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMessage *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *int32                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDeviceEventRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceEventRecordResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventRecordResponseBody) SetRequestId(v string) *QueryDeviceEventRecordResponseBody {
	s.RequestId = &v
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

func (s *QueryDeviceEventRecordResponseBody) SetCode(v int32) *QueryDeviceEventRecordResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDeviceEventRecordResponseBody) SetSuccess(v bool) *QueryDeviceEventRecordResponseBody {
	s.Success = &v
	return s
}

type QueryDeviceEventRecordResponseBodyData struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	FileName  *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	VodUrl    *string `json:"VodUrl,omitempty" xml:"VodUrl,omitempty"`
}

func (s QueryDeviceEventRecordResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceEventRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventRecordResponseBodyData) SetEndTime(v string) *QueryDeviceEventRecordResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *QueryDeviceEventRecordResponseBodyData) SetBeginTime(v string) *QueryDeviceEventRecordResponseBodyData {
	s.BeginTime = &v
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDeviceEventRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryDeviceEventRecordResponse) SetBody(v *QueryDeviceEventRecordResponseBody) *QueryDeviceEventRecordResponse {
	s.Body = v
	return s
}

type QueryDeviceFileVodRequest struct {
	ApiProduct    *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	FileName      *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	ShouldEncrypt *bool   `json:"ShouldEncrypt,omitempty" xml:"ShouldEncrypt,omitempty"`
	EncryptType   *int32  `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
}

func (s QueryDeviceFileVodRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceFileVodRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceFileVodRequest) SetApiProduct(v string) *QueryDeviceFileVodRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryDeviceFileVodRequest) SetApiRevision(v string) *QueryDeviceFileVodRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryDeviceFileVodRequest) SetIotId(v string) *QueryDeviceFileVodRequest {
	s.IotId = &v
	return s
}

func (s *QueryDeviceFileVodRequest) SetFileName(v string) *QueryDeviceFileVodRequest {
	s.FileName = &v
	return s
}

func (s *QueryDeviceFileVodRequest) SetShouldEncrypt(v bool) *QueryDeviceFileVodRequest {
	s.ShouldEncrypt = &v
	return s
}

func (s *QueryDeviceFileVodRequest) SetEncryptType(v int32) *QueryDeviceFileVodRequest {
	s.EncryptType = &v
	return s
}

type QueryDeviceFileVodResponseBody struct {
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DecryptKey   *string                             `json:"DecryptKey,omitempty" xml:"DecryptKey,omitempty"`
	Data         *QueryDeviceFileVodResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDeviceFileVodResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceFileVodResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceFileVodResponseBody) SetRequestId(v string) *QueryDeviceFileVodResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDeviceFileVodResponseBody) SetDecryptKey(v string) *QueryDeviceFileVodResponseBody {
	s.DecryptKey = &v
	return s
}

func (s *QueryDeviceFileVodResponseBody) SetData(v *QueryDeviceFileVodResponseBodyData) *QueryDeviceFileVodResponseBody {
	s.Data = v
	return s
}

func (s *QueryDeviceFileVodResponseBody) SetErrorMessage(v string) *QueryDeviceFileVodResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryDeviceFileVodResponseBody) SetCode(v string) *QueryDeviceFileVodResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDeviceFileVodResponseBody) SetSuccess(v bool) *QueryDeviceFileVodResponseBody {
	s.Success = &v
	return s
}

type QueryDeviceFileVodResponseBodyData struct {
	VodUrl *string `json:"VodUrl,omitempty" xml:"VodUrl,omitempty"`
}

func (s QueryDeviceFileVodResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceFileVodResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDeviceFileVodResponseBodyData) SetVodUrl(v string) *QueryDeviceFileVodResponseBodyData {
	s.VodUrl = &v
	return s
}

type QueryDeviceFileVodResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDeviceFileVodResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDeviceFileVodResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceFileVodResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceFileVodResponse) SetHeaders(v map[string]*string) *QueryDeviceFileVodResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceFileVodResponse) SetBody(v *QueryDeviceFileVodResponseBody) *QueryDeviceFileVodResponse {
	s.Body = v
	return s
}

type QueryDevicePictureFileRequest struct {
	ApiProduct    *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	CaptureId     *string `json:"CaptureId,omitempty" xml:"CaptureId,omitempty"`
	PictureType   *int32  `json:"PictureType,omitempty" xml:"PictureType,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
}

func (s QueryDevicePictureFileRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePictureFileRequest) GoString() string {
	return s.String()
}

func (s *QueryDevicePictureFileRequest) SetApiProduct(v string) *QueryDevicePictureFileRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryDevicePictureFileRequest) SetApiRevision(v string) *QueryDevicePictureFileRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryDevicePictureFileRequest) SetIotId(v string) *QueryDevicePictureFileRequest {
	s.IotId = &v
	return s
}

func (s *QueryDevicePictureFileRequest) SetCaptureId(v string) *QueryDevicePictureFileRequest {
	s.CaptureId = &v
	return s
}

func (s *QueryDevicePictureFileRequest) SetPictureType(v int32) *QueryDevicePictureFileRequest {
	s.PictureType = &v
	return s
}

func (s *QueryDevicePictureFileRequest) SetIotInstanceId(v string) *QueryDevicePictureFileRequest {
	s.IotInstanceId = &v
	return s
}

type QueryDevicePictureFileResponseBody struct {
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryDevicePictureFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDevicePictureFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePictureFileResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDevicePictureFileResponseBody) SetRequestId(v string) *QueryDevicePictureFileResponseBody {
	s.RequestId = &v
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

func (s *QueryDevicePictureFileResponseBody) SetCode(v string) *QueryDevicePictureFileResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDevicePictureFileResponseBody) SetSuccess(v bool) *QueryDevicePictureFileResponseBody {
	s.Success = &v
	return s
}

type QueryDevicePictureFileResponseBodyData struct {
	PicUrl        *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	PicCreateTime *int64  `json:"PicCreateTime,omitempty" xml:"PicCreateTime,omitempty"`
	PicId         *string `json:"PicId,omitempty" xml:"PicId,omitempty"`
	ThumbUrl      *string `json:"ThumbUrl,omitempty" xml:"ThumbUrl,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
}

func (s QueryDevicePictureFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePictureFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDevicePictureFileResponseBodyData) SetPicUrl(v string) *QueryDevicePictureFileResponseBodyData {
	s.PicUrl = &v
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

func (s *QueryDevicePictureFileResponseBodyData) SetThumbUrl(v string) *QueryDevicePictureFileResponseBodyData {
	s.ThumbUrl = &v
	return s
}

func (s *QueryDevicePictureFileResponseBodyData) SetIotId(v string) *QueryDevicePictureFileResponseBodyData {
	s.IotId = &v
	return s
}

type QueryDevicePictureFileResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDevicePictureFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryDevicePictureFileResponse) SetBody(v *QueryDevicePictureFileResponseBody) *QueryDevicePictureFileResponse {
	s.Body = v
	return s
}

type QueryDevicePictureLifeCycleRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId       *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
}

func (s QueryDevicePictureLifeCycleRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePictureLifeCycleRequest) GoString() string {
	return s.String()
}

func (s *QueryDevicePictureLifeCycleRequest) SetApiProduct(v string) *QueryDevicePictureLifeCycleRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryDevicePictureLifeCycleRequest) SetApiRevision(v string) *QueryDevicePictureLifeCycleRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryDevicePictureLifeCycleRequest) SetIotId(v string) *QueryDevicePictureLifeCycleRequest {
	s.IotId = &v
	return s
}

type QueryDevicePictureLifeCycleResponseBody struct {
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryDevicePictureLifeCycleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDevicePictureLifeCycleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePictureLifeCycleResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDevicePictureLifeCycleResponseBody) SetRequestId(v string) *QueryDevicePictureLifeCycleResponseBody {
	s.RequestId = &v
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

func (s *QueryDevicePictureLifeCycleResponseBody) SetCode(v string) *QueryDevicePictureLifeCycleResponseBody {
	s.Code = &v
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
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDevicePictureLifeCycleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryDevicePictureLifeCycleResponse) SetBody(v *QueryDevicePictureLifeCycleResponseBody) *QueryDevicePictureLifeCycleResponse {
	s.Body = v
	return s
}

type QueryDevicePurifyJobsRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId       *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s QueryDevicePurifyJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePurifyJobsRequest) GoString() string {
	return s.String()
}

func (s *QueryDevicePurifyJobsRequest) SetApiProduct(v string) *QueryDevicePurifyJobsRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryDevicePurifyJobsRequest) SetApiRevision(v string) *QueryDevicePurifyJobsRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryDevicePurifyJobsRequest) SetIotId(v string) *QueryDevicePurifyJobsRequest {
	s.IotId = &v
	return s
}

func (s *QueryDevicePurifyJobsRequest) SetPageSize(v int32) *QueryDevicePurifyJobsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryDevicePurifyJobsRequest) SetCurrentPage(v int32) *QueryDevicePurifyJobsRequest {
	s.CurrentPage = &v
	return s
}

type QueryDevicePurifyJobsResponseBody struct {
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryDevicePurifyJobsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDevicePurifyJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePurifyJobsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDevicePurifyJobsResponseBody) SetRequestId(v string) *QueryDevicePurifyJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDevicePurifyJobsResponseBody) SetData(v *QueryDevicePurifyJobsResponseBodyData) *QueryDevicePurifyJobsResponseBody {
	s.Data = v
	return s
}

func (s *QueryDevicePurifyJobsResponseBody) SetErrorMessage(v string) *QueryDevicePurifyJobsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryDevicePurifyJobsResponseBody) SetCode(v string) *QueryDevicePurifyJobsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDevicePurifyJobsResponseBody) SetSuccess(v bool) *QueryDevicePurifyJobsResponseBody {
	s.Success = &v
	return s
}

type QueryDevicePurifyJobsResponseBodyData struct {
	CurrentPage *int32                                       `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	List        []*QueryDevicePurifyJobsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageSize    *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total       *int32                                       `json:"Total,omitempty" xml:"Total,omitempty"`
	PageCount   *int32                                       `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
}

func (s QueryDevicePurifyJobsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePurifyJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDevicePurifyJobsResponseBodyData) SetCurrentPage(v int32) *QueryDevicePurifyJobsResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *QueryDevicePurifyJobsResponseBodyData) SetList(v []*QueryDevicePurifyJobsResponseBodyDataList) *QueryDevicePurifyJobsResponseBodyData {
	s.List = v
	return s
}

func (s *QueryDevicePurifyJobsResponseBodyData) SetPageSize(v int32) *QueryDevicePurifyJobsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryDevicePurifyJobsResponseBodyData) SetTotal(v int32) *QueryDevicePurifyJobsResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryDevicePurifyJobsResponseBodyData) SetPageCount(v int32) *QueryDevicePurifyJobsResponseBodyData {
	s.PageCount = &v
	return s
}

type QueryDevicePurifyJobsResponseBodyDataList struct {
	Status               *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	DeviceName           *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	UserId               *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	PlanId               *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	EndTime              *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime            *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	PurifyRecordIndexUrl *string `json:"PurifyRecordIndexUrl,omitempty" xml:"PurifyRecordIndexUrl,omitempty"`
	ProductKey           *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	PurifyRecordNameUrl  *string `json:"PurifyRecordNameUrl,omitempty" xml:"PurifyRecordNameUrl,omitempty"`
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	IotId                *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	TenantId             *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryDevicePurifyJobsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePurifyJobsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryDevicePurifyJobsResponseBodyDataList) SetStatus(v int32) *QueryDevicePurifyJobsResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *QueryDevicePurifyJobsResponseBodyDataList) SetDeviceName(v string) *QueryDevicePurifyJobsResponseBodyDataList {
	s.DeviceName = &v
	return s
}

func (s *QueryDevicePurifyJobsResponseBodyDataList) SetUserId(v string) *QueryDevicePurifyJobsResponseBodyDataList {
	s.UserId = &v
	return s
}

func (s *QueryDevicePurifyJobsResponseBodyDataList) SetPlanId(v string) *QueryDevicePurifyJobsResponseBodyDataList {
	s.PlanId = &v
	return s
}

func (s *QueryDevicePurifyJobsResponseBodyDataList) SetEndTime(v int64) *QueryDevicePurifyJobsResponseBodyDataList {
	s.EndTime = &v
	return s
}

func (s *QueryDevicePurifyJobsResponseBodyDataList) SetStartTime(v int64) *QueryDevicePurifyJobsResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *QueryDevicePurifyJobsResponseBodyDataList) SetPurifyRecordIndexUrl(v string) *QueryDevicePurifyJobsResponseBodyDataList {
	s.PurifyRecordIndexUrl = &v
	return s
}

func (s *QueryDevicePurifyJobsResponseBodyDataList) SetProductKey(v string) *QueryDevicePurifyJobsResponseBodyDataList {
	s.ProductKey = &v
	return s
}

func (s *QueryDevicePurifyJobsResponseBodyDataList) SetPurifyRecordNameUrl(v string) *QueryDevicePurifyJobsResponseBodyDataList {
	s.PurifyRecordNameUrl = &v
	return s
}

func (s *QueryDevicePurifyJobsResponseBodyDataList) SetJobId(v string) *QueryDevicePurifyJobsResponseBodyDataList {
	s.JobId = &v
	return s
}

func (s *QueryDevicePurifyJobsResponseBodyDataList) SetIotId(v string) *QueryDevicePurifyJobsResponseBodyDataList {
	s.IotId = &v
	return s
}

func (s *QueryDevicePurifyJobsResponseBodyDataList) SetTenantId(v string) *QueryDevicePurifyJobsResponseBodyDataList {
	s.TenantId = &v
	return s
}

type QueryDevicePurifyJobsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDevicePurifyJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDevicePurifyJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePurifyJobsResponse) GoString() string {
	return s.String()
}

func (s *QueryDevicePurifyJobsResponse) SetHeaders(v map[string]*string) *QueryDevicePurifyJobsResponse {
	s.Headers = v
	return s
}

func (s *QueryDevicePurifyJobsResponse) SetBody(v *QueryDevicePurifyJobsResponseBody) *QueryDevicePurifyJobsResponse {
	s.Body = v
	return s
}

type QueryDevicePurifyPlanByPlanIdRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	PlanId      *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s QueryDevicePurifyPlanByPlanIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePurifyPlanByPlanIdRequest) GoString() string {
	return s.String()
}

func (s *QueryDevicePurifyPlanByPlanIdRequest) SetApiProduct(v string) *QueryDevicePurifyPlanByPlanIdRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryDevicePurifyPlanByPlanIdRequest) SetApiRevision(v string) *QueryDevicePurifyPlanByPlanIdRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryDevicePurifyPlanByPlanIdRequest) SetPlanId(v string) *QueryDevicePurifyPlanByPlanIdRequest {
	s.PlanId = &v
	return s
}

type QueryDevicePurifyPlanByPlanIdResponseBody struct {
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryDevicePurifyPlanByPlanIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDevicePurifyPlanByPlanIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePurifyPlanByPlanIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDevicePurifyPlanByPlanIdResponseBody) SetRequestId(v string) *QueryDevicePurifyPlanByPlanIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDevicePurifyPlanByPlanIdResponseBody) SetData(v *QueryDevicePurifyPlanByPlanIdResponseBodyData) *QueryDevicePurifyPlanByPlanIdResponseBody {
	s.Data = v
	return s
}

func (s *QueryDevicePurifyPlanByPlanIdResponseBody) SetErrorMessage(v string) *QueryDevicePurifyPlanByPlanIdResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryDevicePurifyPlanByPlanIdResponseBody) SetCode(v string) *QueryDevicePurifyPlanByPlanIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDevicePurifyPlanByPlanIdResponseBody) SetSuccess(v bool) *QueryDevicePurifyPlanByPlanIdResponseBody {
	s.Success = &v
	return s
}

type QueryDevicePurifyPlanByPlanIdResponseBodyData struct {
	EndTime    *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime  *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	PlanId     *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	TenantId   *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
}

func (s QueryDevicePurifyPlanByPlanIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePurifyPlanByPlanIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDevicePurifyPlanByPlanIdResponseBodyData) SetEndTime(v int32) *QueryDevicePurifyPlanByPlanIdResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *QueryDevicePurifyPlanByPlanIdResponseBodyData) SetStartTime(v int32) *QueryDevicePurifyPlanByPlanIdResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *QueryDevicePurifyPlanByPlanIdResponseBodyData) SetProductKey(v string) *QueryDevicePurifyPlanByPlanIdResponseBodyData {
	s.ProductKey = &v
	return s
}

func (s *QueryDevicePurifyPlanByPlanIdResponseBodyData) SetDeviceName(v string) *QueryDevicePurifyPlanByPlanIdResponseBodyData {
	s.DeviceName = &v
	return s
}

func (s *QueryDevicePurifyPlanByPlanIdResponseBodyData) SetUserId(v string) *QueryDevicePurifyPlanByPlanIdResponseBodyData {
	s.UserId = &v
	return s
}

func (s *QueryDevicePurifyPlanByPlanIdResponseBodyData) SetPlanId(v string) *QueryDevicePurifyPlanByPlanIdResponseBodyData {
	s.PlanId = &v
	return s
}

func (s *QueryDevicePurifyPlanByPlanIdResponseBodyData) SetTenantId(v string) *QueryDevicePurifyPlanByPlanIdResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *QueryDevicePurifyPlanByPlanIdResponseBodyData) SetIotId(v string) *QueryDevicePurifyPlanByPlanIdResponseBodyData {
	s.IotId = &v
	return s
}

type QueryDevicePurifyPlanByPlanIdResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDevicePurifyPlanByPlanIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDevicePurifyPlanByPlanIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePurifyPlanByPlanIdResponse) GoString() string {
	return s.String()
}

func (s *QueryDevicePurifyPlanByPlanIdResponse) SetHeaders(v map[string]*string) *QueryDevicePurifyPlanByPlanIdResponse {
	s.Headers = v
	return s
}

func (s *QueryDevicePurifyPlanByPlanIdResponse) SetBody(v *QueryDevicePurifyPlanByPlanIdResponseBody) *QueryDevicePurifyPlanByPlanIdResponse {
	s.Body = v
	return s
}

type QueryDevicePurifyPlansRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId       *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s QueryDevicePurifyPlansRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePurifyPlansRequest) GoString() string {
	return s.String()
}

func (s *QueryDevicePurifyPlansRequest) SetApiProduct(v string) *QueryDevicePurifyPlansRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryDevicePurifyPlansRequest) SetApiRevision(v string) *QueryDevicePurifyPlansRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryDevicePurifyPlansRequest) SetIotId(v string) *QueryDevicePurifyPlansRequest {
	s.IotId = &v
	return s
}

func (s *QueryDevicePurifyPlansRequest) SetPageSize(v int32) *QueryDevicePurifyPlansRequest {
	s.PageSize = &v
	return s
}

func (s *QueryDevicePurifyPlansRequest) SetCurrentPage(v int32) *QueryDevicePurifyPlansRequest {
	s.CurrentPage = &v
	return s
}

type QueryDevicePurifyPlansResponseBody struct {
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryDevicePurifyPlansResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDevicePurifyPlansResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePurifyPlansResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDevicePurifyPlansResponseBody) SetRequestId(v string) *QueryDevicePurifyPlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDevicePurifyPlansResponseBody) SetData(v *QueryDevicePurifyPlansResponseBodyData) *QueryDevicePurifyPlansResponseBody {
	s.Data = v
	return s
}

func (s *QueryDevicePurifyPlansResponseBody) SetErrorMessage(v string) *QueryDevicePurifyPlansResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryDevicePurifyPlansResponseBody) SetCode(v string) *QueryDevicePurifyPlansResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDevicePurifyPlansResponseBody) SetSuccess(v bool) *QueryDevicePurifyPlansResponseBody {
	s.Success = &v
	return s
}

type QueryDevicePurifyPlansResponseBodyData struct {
	CurrentPage *int32                                        `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	List        []*QueryDevicePurifyPlansResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageSize    *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total       *int32                                        `json:"Total,omitempty" xml:"Total,omitempty"`
	PageCount   *int32                                        `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
}

func (s QueryDevicePurifyPlansResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePurifyPlansResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDevicePurifyPlansResponseBodyData) SetCurrentPage(v int32) *QueryDevicePurifyPlansResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *QueryDevicePurifyPlansResponseBodyData) SetList(v []*QueryDevicePurifyPlansResponseBodyDataList) *QueryDevicePurifyPlansResponseBodyData {
	s.List = v
	return s
}

func (s *QueryDevicePurifyPlansResponseBodyData) SetPageSize(v int32) *QueryDevicePurifyPlansResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryDevicePurifyPlansResponseBodyData) SetTotal(v int32) *QueryDevicePurifyPlansResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryDevicePurifyPlansResponseBodyData) SetPageCount(v int32) *QueryDevicePurifyPlansResponseBodyData {
	s.PageCount = &v
	return s
}

type QueryDevicePurifyPlansResponseBodyDataList struct {
	EndTime    *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime  *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	PlanId     *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	TenantId   *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
}

func (s QueryDevicePurifyPlansResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePurifyPlansResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryDevicePurifyPlansResponseBodyDataList) SetEndTime(v int32) *QueryDevicePurifyPlansResponseBodyDataList {
	s.EndTime = &v
	return s
}

func (s *QueryDevicePurifyPlansResponseBodyDataList) SetStartTime(v int32) *QueryDevicePurifyPlansResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *QueryDevicePurifyPlansResponseBodyDataList) SetProductKey(v string) *QueryDevicePurifyPlansResponseBodyDataList {
	s.ProductKey = &v
	return s
}

func (s *QueryDevicePurifyPlansResponseBodyDataList) SetDeviceName(v string) *QueryDevicePurifyPlansResponseBodyDataList {
	s.DeviceName = &v
	return s
}

func (s *QueryDevicePurifyPlansResponseBodyDataList) SetUserId(v string) *QueryDevicePurifyPlansResponseBodyDataList {
	s.UserId = &v
	return s
}

func (s *QueryDevicePurifyPlansResponseBodyDataList) SetPlanId(v string) *QueryDevicePurifyPlansResponseBodyDataList {
	s.PlanId = &v
	return s
}

func (s *QueryDevicePurifyPlansResponseBodyDataList) SetTenantId(v string) *QueryDevicePurifyPlansResponseBodyDataList {
	s.TenantId = &v
	return s
}

func (s *QueryDevicePurifyPlansResponseBodyDataList) SetIotId(v string) *QueryDevicePurifyPlansResponseBodyDataList {
	s.IotId = &v
	return s
}

type QueryDevicePurifyPlansResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDevicePurifyPlansResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDevicePurifyPlansResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePurifyPlansResponse) GoString() string {
	return s.String()
}

func (s *QueryDevicePurifyPlansResponse) SetHeaders(v map[string]*string) *QueryDevicePurifyPlansResponse {
	s.Headers = v
	return s
}

func (s *QueryDevicePurifyPlansResponse) SetBody(v *QueryDevicePurifyPlansResponseBody) *QueryDevicePurifyPlansResponse {
	s.Body = v
	return s
}

type QueryDeviceRecordLifeCycleRequest struct {
	ApiProduct  *string   `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string   `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	DeviceList  []*string `json:"DeviceList,omitempty" xml:"DeviceList,omitempty" type:"Repeated"`
}

func (s QueryDeviceRecordLifeCycleRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceRecordLifeCycleRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceRecordLifeCycleRequest) SetApiProduct(v string) *QueryDeviceRecordLifeCycleRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryDeviceRecordLifeCycleRequest) SetApiRevision(v string) *QueryDeviceRecordLifeCycleRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryDeviceRecordLifeCycleRequest) SetDeviceList(v []*string) *QueryDeviceRecordLifeCycleRequest {
	s.DeviceList = v
	return s
}

type QueryDeviceRecordLifeCycleResponseBody struct {
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         []*QueryDeviceRecordLifeCycleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMessage *string                                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *int32                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDeviceRecordLifeCycleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceRecordLifeCycleResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceRecordLifeCycleResponseBody) SetRequestId(v string) *QueryDeviceRecordLifeCycleResponseBody {
	s.RequestId = &v
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

func (s *QueryDeviceRecordLifeCycleResponseBody) SetCode(v int32) *QueryDeviceRecordLifeCycleResponseBody {
	s.Code = &v
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
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDeviceRecordLifeCycleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryDeviceRecordLifeCycleResponse) SetBody(v *QueryDeviceRecordLifeCycleResponseBody) *QueryDeviceRecordLifeCycleResponse {
	s.Body = v
	return s
}

type QueryDeviceVodUrlRequest struct {
	ApiProduct    *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	FileName      *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	ShouldEncrypt *bool   `json:"ShouldEncrypt,omitempty" xml:"ShouldEncrypt,omitempty"`
	EncryptType   *int32  `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	Scheme        *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
	SeekTime      *int32  `json:"SeekTime,omitempty" xml:"SeekTime,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
}

func (s QueryDeviceVodUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceVodUrlRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceVodUrlRequest) SetApiProduct(v string) *QueryDeviceVodUrlRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) SetApiRevision(v string) *QueryDeviceVodUrlRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) SetIotId(v string) *QueryDeviceVodUrlRequest {
	s.IotId = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) SetFileName(v string) *QueryDeviceVodUrlRequest {
	s.FileName = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) SetShouldEncrypt(v bool) *QueryDeviceVodUrlRequest {
	s.ShouldEncrypt = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) SetEncryptType(v int32) *QueryDeviceVodUrlRequest {
	s.EncryptType = &v
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

func (s *QueryDeviceVodUrlRequest) SetIotInstanceId(v string) *QueryDeviceVodUrlRequest {
	s.IotInstanceId = &v
	return s
}

type QueryDeviceVodUrlResponseBody struct {
	RequestId    *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DecryptKey   *string                            `json:"DecryptKey,omitempty" xml:"DecryptKey,omitempty"`
	Data         *QueryDeviceVodUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDeviceVodUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceVodUrlResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceVodUrlResponseBody) SetRequestId(v string) *QueryDeviceVodUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDeviceVodUrlResponseBody) SetDecryptKey(v string) *QueryDeviceVodUrlResponseBody {
	s.DecryptKey = &v
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

func (s *QueryDeviceVodUrlResponseBody) SetCode(v string) *QueryDeviceVodUrlResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDeviceVodUrlResponseBody) SetSuccess(v bool) *QueryDeviceVodUrlResponseBody {
	s.Success = &v
	return s
}

type QueryDeviceVodUrlResponseBodyData struct {
	VodUrl *string `json:"VodUrl,omitempty" xml:"VodUrl,omitempty"`
}

func (s QueryDeviceVodUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceVodUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDeviceVodUrlResponseBodyData) SetVodUrl(v string) *QueryDeviceVodUrlResponseBodyData {
	s.VodUrl = &v
	return s
}

type QueryDeviceVodUrlResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDeviceVodUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryDeviceVodUrlResponse) SetBody(v *QueryDeviceVodUrlResponseBody) *QueryDeviceVodUrlResponse {
	s.Body = v
	return s
}

type QueryEventRecordPlanDetailRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	PlanId      *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s QueryEventRecordPlanDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDetailRequest) SetApiProduct(v string) *QueryEventRecordPlanDetailRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryEventRecordPlanDetailRequest) SetApiRevision(v string) *QueryEventRecordPlanDetailRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryEventRecordPlanDetailRequest) SetPlanId(v string) *QueryEventRecordPlanDetailRequest {
	s.PlanId = &v
	return s
}

type QueryEventRecordPlanDetailResponseBody struct {
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryEventRecordPlanDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                     `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryEventRecordPlanDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDetailResponseBody) SetRequestId(v string) *QueryEventRecordPlanDetailResponseBody {
	s.RequestId = &v
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

func (s *QueryEventRecordPlanDetailResponseBody) SetCode(v string) *QueryEventRecordPlanDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryEventRecordPlanDetailResponseBody) SetSuccess(v bool) *QueryEventRecordPlanDetailResponseBody {
	s.Success = &v
	return s
}

type QueryEventRecordPlanDetailResponseBodyData struct {
	RecordDuration    *int32                                                  `json:"RecordDuration,omitempty" xml:"RecordDuration,omitempty"`
	PreRecordDuration *int32                                                  `json:"PreRecordDuration,omitempty" xml:"PreRecordDuration,omitempty"`
	Name              *string                                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	PlanId            *string                                                 `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	TemplateInfo      *QueryEventRecordPlanDetailResponseBodyDataTemplateInfo `json:"TemplateInfo,omitempty" xml:"TemplateInfo,omitempty" type:"Struct"`
	TemplateId        *string                                                 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s QueryEventRecordPlanDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDetailResponseBodyData) SetRecordDuration(v int32) *QueryEventRecordPlanDetailResponseBodyData {
	s.RecordDuration = &v
	return s
}

func (s *QueryEventRecordPlanDetailResponseBodyData) SetPreRecordDuration(v int32) *QueryEventRecordPlanDetailResponseBodyData {
	s.PreRecordDuration = &v
	return s
}

func (s *QueryEventRecordPlanDetailResponseBodyData) SetName(v string) *QueryEventRecordPlanDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryEventRecordPlanDetailResponseBodyData) SetPlanId(v string) *QueryEventRecordPlanDetailResponseBodyData {
	s.PlanId = &v
	return s
}

func (s *QueryEventRecordPlanDetailResponseBodyData) SetTemplateInfo(v *QueryEventRecordPlanDetailResponseBodyDataTemplateInfo) *QueryEventRecordPlanDetailResponseBodyData {
	s.TemplateInfo = v
	return s
}

func (s *QueryEventRecordPlanDetailResponseBodyData) SetTemplateId(v string) *QueryEventRecordPlanDetailResponseBodyData {
	s.TemplateId = &v
	return s
}

type QueryEventRecordPlanDetailResponseBodyDataTemplateInfo struct {
	TimeSectionList []*QueryEventRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList `json:"TimeSectionList,omitempty" xml:"TimeSectionList,omitempty" type:"Repeated"`
	AllDay          *int32                                                                   `json:"AllDay,omitempty" xml:"AllDay,omitempty"`
	Default         *int32                                                                   `json:"Default,omitempty" xml:"Default,omitempty"`
	Name            *string                                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateId      *string                                                                  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s QueryEventRecordPlanDetailResponseBodyDataTemplateInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDetailResponseBodyDataTemplateInfo) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDetailResponseBodyDataTemplateInfo) SetTimeSectionList(v []*QueryEventRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList) *QueryEventRecordPlanDetailResponseBodyDataTemplateInfo {
	s.TimeSectionList = v
	return s
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

type QueryEventRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList struct {
	DayOfWeek *int32 `json:"DayOfWeek,omitempty" xml:"DayOfWeek,omitempty"`
	Begin     *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	End       *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s QueryEventRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList) SetDayOfWeek(v int32) *QueryEventRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList {
	s.DayOfWeek = &v
	return s
}

func (s *QueryEventRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList) SetBegin(v int32) *QueryEventRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList {
	s.Begin = &v
	return s
}

func (s *QueryEventRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList) SetEnd(v int32) *QueryEventRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList {
	s.End = &v
	return s
}

type QueryEventRecordPlanDetailResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryEventRecordPlanDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryEventRecordPlanDetailResponse) SetBody(v *QueryEventRecordPlanDetailResponseBody) *QueryEventRecordPlanDetailResponse {
	s.Body = v
	return s
}

type QueryEventRecordPlanDeviceByDeviceRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId       *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	StreamType  *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s QueryEventRecordPlanDeviceByDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDeviceByDeviceRequest) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDeviceByDeviceRequest) SetApiProduct(v string) *QueryEventRecordPlanDeviceByDeviceRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceRequest) SetApiRevision(v string) *QueryEventRecordPlanDeviceByDeviceRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceRequest) SetIotId(v string) *QueryEventRecordPlanDeviceByDeviceRequest {
	s.IotId = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceRequest) SetStreamType(v int32) *QueryEventRecordPlanDeviceByDeviceRequest {
	s.StreamType = &v
	return s
}

type QueryEventRecordPlanDeviceByDeviceResponseBody struct {
	RequestId    *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryEventRecordPlanDeviceByDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryEventRecordPlanDeviceByDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDeviceByDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBody) SetRequestId(v string) *QueryEventRecordPlanDeviceByDeviceResponseBody {
	s.RequestId = &v
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

func (s *QueryEventRecordPlanDeviceByDeviceResponseBody) SetCode(v string) *QueryEventRecordPlanDeviceByDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBody) SetSuccess(v bool) *QueryEventRecordPlanDeviceByDeviceResponseBody {
	s.Success = &v
	return s
}

type QueryEventRecordPlanDeviceByDeviceResponseBodyData struct {
	RecordDuration    *int32                                                          `json:"RecordDuration,omitempty" xml:"RecordDuration,omitempty"`
	PreRecordDuration *int32                                                          `json:"PreRecordDuration,omitempty" xml:"PreRecordDuration,omitempty"`
	Name              *string                                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	PlanId            *string                                                         `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	TemplateInfo      *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo `json:"TemplateInfo,omitempty" xml:"TemplateInfo,omitempty" type:"Struct"`
	TemplateId        *string                                                         `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s QueryEventRecordPlanDeviceByDeviceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDeviceByDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBodyData) SetRecordDuration(v int32) *QueryEventRecordPlanDeviceByDeviceResponseBodyData {
	s.RecordDuration = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBodyData) SetPreRecordDuration(v int32) *QueryEventRecordPlanDeviceByDeviceResponseBodyData {
	s.PreRecordDuration = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBodyData) SetName(v string) *QueryEventRecordPlanDeviceByDeviceResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBodyData) SetPlanId(v string) *QueryEventRecordPlanDeviceByDeviceResponseBodyData {
	s.PlanId = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBodyData) SetTemplateInfo(v *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo) *QueryEventRecordPlanDeviceByDeviceResponseBodyData {
	s.TemplateInfo = v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBodyData) SetTemplateId(v string) *QueryEventRecordPlanDeviceByDeviceResponseBodyData {
	s.TemplateId = &v
	return s
}

type QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo struct {
	TimeSectionList []*QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList `json:"TimeSectionList,omitempty" xml:"TimeSectionList,omitempty" type:"Repeated"`
	AllDay          *int32                                                                           `json:"AllDay,omitempty" xml:"AllDay,omitempty"`
	Default         *int32                                                                           `json:"Default,omitempty" xml:"Default,omitempty"`
	Name            *string                                                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateId      *string                                                                          `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo) SetTimeSectionList(v []*QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList) *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo {
	s.TimeSectionList = v
	return s
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

type QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList struct {
	DayOfWeek *int32 `json:"DayOfWeek,omitempty" xml:"DayOfWeek,omitempty"`
	Begin     *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	End       *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList) SetDayOfWeek(v int32) *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList {
	s.DayOfWeek = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList) SetBegin(v int32) *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList {
	s.Begin = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList) SetEnd(v int32) *QueryEventRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList {
	s.End = &v
	return s
}

type QueryEventRecordPlanDeviceByDeviceResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryEventRecordPlanDeviceByDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryEventRecordPlanDeviceByDeviceResponse) SetBody(v *QueryEventRecordPlanDeviceByDeviceResponseBody) *QueryEventRecordPlanDeviceByDeviceResponse {
	s.Body = v
	return s
}

type QueryEventRecordPlanDeviceByPlanRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	PlanId      *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryEventRecordPlanDeviceByPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDeviceByPlanRequest) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDeviceByPlanRequest) SetApiProduct(v string) *QueryEventRecordPlanDeviceByPlanRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanRequest) SetApiRevision(v string) *QueryEventRecordPlanDeviceByPlanRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanRequest) SetPlanId(v string) *QueryEventRecordPlanDeviceByPlanRequest {
	s.PlanId = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanRequest) SetCurrentPage(v int32) *QueryEventRecordPlanDeviceByPlanRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanRequest) SetPageSize(v int32) *QueryEventRecordPlanDeviceByPlanRequest {
	s.PageSize = &v
	return s
}

type QueryEventRecordPlanDeviceByPlanResponseBody struct {
	RequestId    *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryEventRecordPlanDeviceByPlanResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryEventRecordPlanDeviceByPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDeviceByPlanResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDeviceByPlanResponseBody) SetRequestId(v string) *QueryEventRecordPlanDeviceByPlanResponseBody {
	s.RequestId = &v
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

func (s *QueryEventRecordPlanDeviceByPlanResponseBody) SetCode(v string) *QueryEventRecordPlanDeviceByPlanResponseBody {
	s.Code = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanResponseBody) SetSuccess(v bool) *QueryEventRecordPlanDeviceByPlanResponseBody {
	s.Success = &v
	return s
}

type QueryEventRecordPlanDeviceByPlanResponseBodyData struct {
	List      []*QueryEventRecordPlanDeviceByPlanResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageSize  *int32                                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int32                                                  `json:"Total,omitempty" xml:"Total,omitempty"`
	PageCount *int32                                                  `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	Page      *int32                                                  `json:"Page,omitempty" xml:"Page,omitempty"`
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

func (s *QueryEventRecordPlanDeviceByPlanResponseBodyData) SetPageSize(v int32) *QueryEventRecordPlanDeviceByPlanResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanResponseBodyData) SetTotal(v int32) *QueryEventRecordPlanDeviceByPlanResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanResponseBodyData) SetPageCount(v int32) *QueryEventRecordPlanDeviceByPlanResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanResponseBodyData) SetPage(v int32) *QueryEventRecordPlanDeviceByPlanResponseBodyData {
	s.Page = &v
	return s
}

type QueryEventRecordPlanDeviceByPlanResponseBodyDataList struct {
	StreamType *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
}

func (s QueryEventRecordPlanDeviceByPlanResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlanDeviceByPlanResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDeviceByPlanResponseBodyDataList) SetStreamType(v int32) *QueryEventRecordPlanDeviceByPlanResponseBodyDataList {
	s.StreamType = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanResponseBodyDataList) SetIotId(v string) *QueryEventRecordPlanDeviceByPlanResponseBodyDataList {
	s.IotId = &v
	return s
}

type QueryEventRecordPlanDeviceByPlanResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryEventRecordPlanDeviceByPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryEventRecordPlanDeviceByPlanResponse) SetBody(v *QueryEventRecordPlanDeviceByPlanResponseBody) *QueryEventRecordPlanDeviceByPlanResponse {
	s.Body = v
	return s
}

type QueryEventRecordPlansRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryEventRecordPlansRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlansRequest) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlansRequest) SetApiProduct(v string) *QueryEventRecordPlansRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryEventRecordPlansRequest) SetApiRevision(v string) *QueryEventRecordPlansRequest {
	s.ApiRevision = &v
	return s
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
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryEventRecordPlansResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryEventRecordPlansResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRecordPlansResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlansResponseBody) SetRequestId(v string) *QueryEventRecordPlansResponseBody {
	s.RequestId = &v
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

func (s *QueryEventRecordPlansResponseBody) SetCode(v string) *QueryEventRecordPlansResponseBody {
	s.Code = &v
	return s
}

func (s *QueryEventRecordPlansResponseBody) SetSuccess(v bool) *QueryEventRecordPlansResponseBody {
	s.Success = &v
	return s
}

type QueryEventRecordPlansResponseBodyData struct {
	List      []*QueryEventRecordPlansResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageSize  *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageCount *int32                                       `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	Total     *int32                                       `json:"Total,omitempty" xml:"Total,omitempty"`
	Page      *int32                                       `json:"Page,omitempty" xml:"Page,omitempty"`
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

func (s *QueryEventRecordPlansResponseBodyData) SetPageSize(v int32) *QueryEventRecordPlansResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryEventRecordPlansResponseBodyData) SetPageCount(v int32) *QueryEventRecordPlansResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryEventRecordPlansResponseBodyData) SetTotal(v int32) *QueryEventRecordPlansResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryEventRecordPlansResponseBodyData) SetPage(v int32) *QueryEventRecordPlansResponseBodyData {
	s.Page = &v
	return s
}

type QueryEventRecordPlansResponseBodyDataList struct {
	EventType         *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	RecordDuration    *int32  `json:"RecordDuration,omitempty" xml:"RecordDuration,omitempty"`
	PreRecordDuration *int32  `json:"PreRecordDuration,omitempty" xml:"PreRecordDuration,omitempty"`
	PlanId            *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

func (s *QueryEventRecordPlansResponseBodyDataList) SetRecordDuration(v int32) *QueryEventRecordPlansResponseBodyDataList {
	s.RecordDuration = &v
	return s
}

func (s *QueryEventRecordPlansResponseBodyDataList) SetPreRecordDuration(v int32) *QueryEventRecordPlansResponseBodyDataList {
	s.PreRecordDuration = &v
	return s
}

func (s *QueryEventRecordPlansResponseBodyDataList) SetPlanId(v string) *QueryEventRecordPlansResponseBodyDataList {
	s.PlanId = &v
	return s
}

func (s *QueryEventRecordPlansResponseBodyDataList) SetName(v string) *QueryEventRecordPlansResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *QueryEventRecordPlansResponseBodyDataList) SetTemplateId(v string) *QueryEventRecordPlansResponseBodyDataList {
	s.TemplateId = &v
	return s
}

type QueryEventRecordPlansResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryEventRecordPlansResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryEventRecordPlansResponse) SetBody(v *QueryEventRecordPlansResponseBody) *QueryEventRecordPlansResponse {
	s.Body = v
	return s
}

type QueryFaceAllDeviceGroupRequest struct {
	ApiProduct    *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId   *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNo        *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
}

func (s QueryFaceAllDeviceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceAllDeviceGroupRequest) SetApiProduct(v string) *QueryFaceAllDeviceGroupRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryFaceAllDeviceGroupRequest) SetApiRevision(v string) *QueryFaceAllDeviceGroupRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryFaceAllDeviceGroupRequest) SetIsolationId(v string) *QueryFaceAllDeviceGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceAllDeviceGroupRequest) SetIotInstanceId(v string) *QueryFaceAllDeviceGroupRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryFaceAllDeviceGroupRequest) SetPageSize(v int32) *QueryFaceAllDeviceGroupRequest {
	s.PageSize = &v
	return s
}

func (s *QueryFaceAllDeviceGroupRequest) SetPageNo(v int32) *QueryFaceAllDeviceGroupRequest {
	s.PageNo = &v
	return s
}

type QueryFaceAllDeviceGroupResponseBody struct {
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryFaceAllDeviceGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceAllDeviceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceAllDeviceGroupResponseBody) SetRequestId(v string) *QueryFaceAllDeviceGroupResponseBody {
	s.RequestId = &v
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

func (s *QueryFaceAllDeviceGroupResponseBody) SetCode(v string) *QueryFaceAllDeviceGroupResponseBody {
	s.Code = &v
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
	ModifiedTime    *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	DeviceGroupName *string `json:"DeviceGroupName,omitempty" xml:"DeviceGroupName,omitempty"`
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

func (s *QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList) SetModifiedTime(v string) *QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList {
	s.ModifiedTime = &v
	return s
}

func (s *QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList) SetDeviceGroupName(v string) *QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList {
	s.DeviceGroupName = &v
	return s
}

type QueryFaceAllDeviceGroupResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryFaceAllDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryFaceAllDeviceGroupResponse) SetBody(v *QueryFaceAllDeviceGroupResponseBody) *QueryFaceAllDeviceGroupResponse {
	s.Body = v
	return s
}

type QueryFaceAllUserGroupRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNo      *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
}

func (s QueryFaceAllUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllUserGroupRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserGroupRequest) SetApiProduct(v string) *QueryFaceAllUserGroupRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryFaceAllUserGroupRequest) SetApiRevision(v string) *QueryFaceAllUserGroupRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryFaceAllUserGroupRequest) SetIsolationId(v string) *QueryFaceAllUserGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceAllUserGroupRequest) SetPageSize(v int32) *QueryFaceAllUserGroupRequest {
	s.PageSize = &v
	return s
}

func (s *QueryFaceAllUserGroupRequest) SetPageNo(v int32) *QueryFaceAllUserGroupRequest {
	s.PageNo = &v
	return s
}

type QueryFaceAllUserGroupResponseBody struct {
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryFaceAllUserGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceAllUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserGroupResponseBody) SetRequestId(v string) *QueryFaceAllUserGroupResponseBody {
	s.RequestId = &v
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

func (s *QueryFaceAllUserGroupResponseBody) SetCode(v string) *QueryFaceAllUserGroupResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceAllUserGroupResponseBody) SetSuccess(v bool) *QueryFaceAllUserGroupResponseBody {
	s.Success = &v
	return s
}

type QueryFaceAllUserGroupResponseBodyData struct {
	PageNo        *int32                                                `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	UserGroupList []*QueryFaceAllUserGroupResponseBodyDataUserGroupList `json:"UserGroupList,omitempty" xml:"UserGroupList,omitempty" type:"Repeated"`
	PageSize      *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total         *int32                                                `json:"Total,omitempty" xml:"Total,omitempty"`
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

func (s *QueryFaceAllUserGroupResponseBodyData) SetUserGroupList(v []*QueryFaceAllUserGroupResponseBodyDataUserGroupList) *QueryFaceAllUserGroupResponseBodyData {
	s.UserGroupList = v
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryFaceAllUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryFaceAllUserGroupResponse) SetBody(v *QueryFaceAllUserGroupResponseBody) *QueryFaceAllUserGroupResponse {
	s.Body = v
	return s
}

type QueryFaceAllUserGroupAndDeviceGroupRelationRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNo      *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
}

func (s QueryFaceAllUserGroupAndDeviceGroupRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllUserGroupAndDeviceGroupRelationRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationRequest) SetApiProduct(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationRequest) SetApiRevision(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationRequest) SetIsolationId(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationRequest) SetPageSize(v int32) *QueryFaceAllUserGroupAndDeviceGroupRelationRequest {
	s.PageSize = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationRequest) SetPageNo(v int32) *QueryFaceAllUserGroupAndDeviceGroupRelationRequest {
	s.PageNo = &v
	return s
}

type QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody struct {
	RequestId    *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) SetRequestId(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody {
	s.RequestId = &v
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

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) SetCode(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) SetSuccess(v bool) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody {
	s.Success = &v
	return s
}

type QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData struct {
	List     []*QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageSize *int32                                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int32                                                             `json:"Total,omitempty" xml:"Total,omitempty"`
	Page     *int32                                                             `json:"Page,omitempty" xml:"Page,omitempty"`
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

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData) SetPageSize(v int32) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData) SetTotal(v int32) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData) SetPage(v int32) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData {
	s.Page = &v
	return s
}

type QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList struct {
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	ModifiedTime  *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ControlType   *string `json:"ControlType,omitempty" xml:"ControlType,omitempty"`
	UserGroupId   *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	ControlId     *string `json:"ControlId,omitempty" xml:"ControlId,omitempty"`
}

func (s QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) SetDeviceGroupId(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList {
	s.DeviceGroupId = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) SetModifiedTime(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList {
	s.ModifiedTime = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) SetControlType(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList {
	s.ControlType = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) SetUserGroupId(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList {
	s.UserGroupId = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) SetControlId(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList {
	s.ControlId = &v
	return s
}

type QueryFaceAllUserGroupAndDeviceGroupRelationResponse struct {
	Headers map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponse) SetBody(v *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) *QueryFaceAllUserGroupAndDeviceGroupRelationResponse {
	s.Body = v
	return s
}

type QueryFaceAllUserIdsByGroupIdRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNo      *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
}

func (s QueryFaceAllUserIdsByGroupIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllUserIdsByGroupIdRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserIdsByGroupIdRequest) SetApiProduct(v string) *QueryFaceAllUserIdsByGroupIdRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdRequest) SetApiRevision(v string) *QueryFaceAllUserIdsByGroupIdRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdRequest) SetIsolationId(v string) *QueryFaceAllUserIdsByGroupIdRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdRequest) SetUserGroupId(v string) *QueryFaceAllUserIdsByGroupIdRequest {
	s.UserGroupId = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdRequest) SetPageSize(v int32) *QueryFaceAllUserIdsByGroupIdRequest {
	s.PageSize = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdRequest) SetPageNo(v int32) *QueryFaceAllUserIdsByGroupIdRequest {
	s.PageNo = &v
	return s
}

type QueryFaceAllUserIdsByGroupIdResponseBody struct {
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryFaceAllUserIdsByGroupIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceAllUserIdsByGroupIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllUserIdsByGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBody) SetRequestId(v string) *QueryFaceAllUserIdsByGroupIdResponseBody {
	s.RequestId = &v
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

func (s *QueryFaceAllUserIdsByGroupIdResponseBody) SetCode(v string) *QueryFaceAllUserIdsByGroupIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBody) SetSuccess(v bool) *QueryFaceAllUserIdsByGroupIdResponseBody {
	s.Success = &v
	return s
}

type QueryFaceAllUserIdsByGroupIdResponseBodyData struct {
	List     []*QueryFaceAllUserIdsByGroupIdResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageSize *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int32                                              `json:"Total,omitempty" xml:"Total,omitempty"`
	Page     *int32                                              `json:"Page,omitempty" xml:"Page,omitempty"`
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

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyData) SetPageSize(v int32) *QueryFaceAllUserIdsByGroupIdResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyData) SetTotal(v int32) *QueryFaceAllUserIdsByGroupIdResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyData) SetPage(v int32) *QueryFaceAllUserIdsByGroupIdResponseBodyData {
	s.Page = &v
	return s
}

type QueryFaceAllUserIdsByGroupIdResponseBodyDataList struct {
	Params       *string `json:"Params,omitempty" xml:"Params,omitempty"`
	CustomUserId *string `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s QueryFaceAllUserIdsByGroupIdResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceAllUserIdsByGroupIdResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyDataList) SetParams(v string) *QueryFaceAllUserIdsByGroupIdResponseBodyDataList {
	s.Params = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyDataList) SetCustomUserId(v string) *QueryFaceAllUserIdsByGroupIdResponseBodyDataList {
	s.CustomUserId = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyDataList) SetUserId(v string) *QueryFaceAllUserIdsByGroupIdResponseBodyDataList {
	s.UserId = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyDataList) SetName(v string) *QueryFaceAllUserIdsByGroupIdResponseBodyDataList {
	s.Name = &v
	return s
}

type QueryFaceAllUserIdsByGroupIdResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryFaceAllUserIdsByGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryFaceAllUserIdsByGroupIdResponse) SetBody(v *QueryFaceAllUserIdsByGroupIdResponseBody) *QueryFaceAllUserIdsByGroupIdResponse {
	s.Body = v
	return s
}

type QueryFaceCustomUserIdByUserIdRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryFaceCustomUserIdByUserIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceCustomUserIdByUserIdRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceCustomUserIdByUserIdRequest) SetApiProduct(v string) *QueryFaceCustomUserIdByUserIdRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryFaceCustomUserIdByUserIdRequest) SetApiRevision(v string) *QueryFaceCustomUserIdByUserIdRequest {
	s.ApiRevision = &v
	return s
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
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceCustomUserIdByUserIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceCustomUserIdByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceCustomUserIdByUserIdResponseBody) SetRequestId(v string) *QueryFaceCustomUserIdByUserIdResponseBody {
	s.RequestId = &v
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

func (s *QueryFaceCustomUserIdByUserIdResponseBody) SetCode(v string) *QueryFaceCustomUserIdByUserIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceCustomUserIdByUserIdResponseBody) SetSuccess(v bool) *QueryFaceCustomUserIdByUserIdResponseBody {
	s.Success = &v
	return s
}

type QueryFaceCustomUserIdByUserIdResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryFaceCustomUserIdByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryFaceCustomUserIdByUserIdResponse) SetBody(v *QueryFaceCustomUserIdByUserIdResponseBody) *QueryFaceCustomUserIdByUserIdResponse {
	s.Body = v
	return s
}

type QueryFaceDeviceGroupsByDeviceRequest struct {
	ApiProduct    *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId   *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNo        *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
}

func (s QueryFaceDeviceGroupsByDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceDeviceGroupsByDeviceRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceDeviceGroupsByDeviceRequest) SetApiProduct(v string) *QueryFaceDeviceGroupsByDeviceRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceRequest) SetApiRevision(v string) *QueryFaceDeviceGroupsByDeviceRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceRequest) SetIsolationId(v string) *QueryFaceDeviceGroupsByDeviceRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceRequest) SetIotInstanceId(v string) *QueryFaceDeviceGroupsByDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceRequest) SetProductKey(v string) *QueryFaceDeviceGroupsByDeviceRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceRequest) SetDeviceName(v string) *QueryFaceDeviceGroupsByDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceRequest) SetPageSize(v int32) *QueryFaceDeviceGroupsByDeviceRequest {
	s.PageSize = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceRequest) SetPageNo(v int32) *QueryFaceDeviceGroupsByDeviceRequest {
	s.PageNo = &v
	return s
}

type QueryFaceDeviceGroupsByDeviceResponseBody struct {
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryFaceDeviceGroupsByDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceDeviceGroupsByDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceDeviceGroupsByDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBody) SetRequestId(v string) *QueryFaceDeviceGroupsByDeviceResponseBody {
	s.RequestId = &v
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

func (s *QueryFaceDeviceGroupsByDeviceResponseBody) SetCode(v string) *QueryFaceDeviceGroupsByDeviceResponseBody {
	s.Code = &v
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
	ModifiedTime    *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	DeviceGroupName *string `json:"DeviceGroupName,omitempty" xml:"DeviceGroupName,omitempty"`
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

func (s *QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList) SetModifiedTime(v string) *QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList {
	s.ModifiedTime = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList) SetDeviceGroupName(v string) *QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList {
	s.DeviceGroupName = &v
	return s
}

type QueryFaceDeviceGroupsByDeviceResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryFaceDeviceGroupsByDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryFaceDeviceGroupsByDeviceResponse) SetBody(v *QueryFaceDeviceGroupsByDeviceResponseBody) *QueryFaceDeviceGroupsByDeviceResponse {
	s.Body = v
	return s
}

type QueryFaceUserRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryFaceUserRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceUserRequest) SetApiProduct(v string) *QueryFaceUserRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryFaceUserRequest) SetApiRevision(v string) *QueryFaceUserRequest {
	s.ApiRevision = &v
	return s
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
	RequestId    *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryFaceUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceUserResponseBody) SetRequestId(v string) *QueryFaceUserResponseBody {
	s.RequestId = &v
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

func (s *QueryFaceUserResponseBody) SetCode(v string) *QueryFaceUserResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceUserResponseBody) SetSuccess(v bool) *QueryFaceUserResponseBody {
	s.Success = &v
	return s
}

type QueryFaceUserResponseBodyData struct {
	Params       *string                                     `json:"Params,omitempty" xml:"Params,omitempty"`
	CustomUserId *string                                     `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	UserId       *string                                     `json:"UserId,omitempty" xml:"UserId,omitempty"`
	FacePicList  []*QueryFaceUserResponseBodyDataFacePicList `json:"FacePicList,omitempty" xml:"FacePicList,omitempty" type:"Repeated"`
	Name         *string                                     `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s QueryFaceUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceUserResponseBodyData) SetParams(v string) *QueryFaceUserResponseBodyData {
	s.Params = &v
	return s
}

func (s *QueryFaceUserResponseBodyData) SetCustomUserId(v string) *QueryFaceUserResponseBodyData {
	s.CustomUserId = &v
	return s
}

func (s *QueryFaceUserResponseBodyData) SetUserId(v string) *QueryFaceUserResponseBodyData {
	s.UserId = &v
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

type QueryFaceUserResponseBodyDataFacePicList struct {
	FaceUrl        *string                                                   `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	FeatureDTOList []*QueryFaceUserResponseBodyDataFacePicListFeatureDTOList `json:"FeatureDTOList,omitempty" xml:"FeatureDTOList,omitempty" type:"Repeated"`
	FaceMd5        *string                                                   `json:"FaceMd5,omitempty" xml:"FaceMd5,omitempty"`
}

func (s QueryFaceUserResponseBodyDataFacePicList) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserResponseBodyDataFacePicList) GoString() string {
	return s.String()
}

func (s *QueryFaceUserResponseBodyDataFacePicList) SetFaceUrl(v string) *QueryFaceUserResponseBodyDataFacePicList {
	s.FaceUrl = &v
	return s
}

func (s *QueryFaceUserResponseBodyDataFacePicList) SetFeatureDTOList(v []*QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) *QueryFaceUserResponseBodyDataFacePicList {
	s.FeatureDTOList = v
	return s
}

func (s *QueryFaceUserResponseBodyDataFacePicList) SetFaceMd5(v string) *QueryFaceUserResponseBodyDataFacePicList {
	s.FaceMd5 = &v
	return s
}

type QueryFaceUserResponseBodyDataFacePicListFeatureDTOList struct {
	AlgorithmName     *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	AlgorithmVersion  *string `json:"AlgorithmVersion,omitempty" xml:"AlgorithmVersion,omitempty"`
	AlgorithmProvider *string `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	ErrorMessage      *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode         *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
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

func (s *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) SetAlgorithmVersion(v string) *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList {
	s.AlgorithmVersion = &v
	return s
}

func (s *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) SetAlgorithmProvider(v string) *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList {
	s.AlgorithmProvider = &v
	return s
}

func (s *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) SetErrorMessage(v string) *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) SetErrorCode(v string) *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList {
	s.ErrorCode = &v
	return s
}

func (s *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) SetFaceMd5(v string) *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList {
	s.FaceMd5 = &v
	return s
}

type QueryFaceUserResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryFaceUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryFaceUserResponse) SetBody(v *QueryFaceUserResponseBody) *QueryFaceUserResponse {
	s.Body = v
	return s
}

type QueryFaceUserGroupRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNo      *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
}

func (s QueryFaceUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserGroupRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceUserGroupRequest) SetApiProduct(v string) *QueryFaceUserGroupRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryFaceUserGroupRequest) SetApiRevision(v string) *QueryFaceUserGroupRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryFaceUserGroupRequest) SetIsolationId(v string) *QueryFaceUserGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceUserGroupRequest) SetUserId(v string) *QueryFaceUserGroupRequest {
	s.UserId = &v
	return s
}

func (s *QueryFaceUserGroupRequest) SetPageSize(v int32) *QueryFaceUserGroupRequest {
	s.PageSize = &v
	return s
}

func (s *QueryFaceUserGroupRequest) SetPageNo(v int32) *QueryFaceUserGroupRequest {
	s.PageNo = &v
	return s
}

type QueryFaceUserGroupResponseBody struct {
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryFaceUserGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceUserGroupResponseBody) SetRequestId(v string) *QueryFaceUserGroupResponseBody {
	s.RequestId = &v
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

func (s *QueryFaceUserGroupResponseBody) SetCode(v string) *QueryFaceUserGroupResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceUserGroupResponseBody) SetSuccess(v bool) *QueryFaceUserGroupResponseBody {
	s.Success = &v
	return s
}

type QueryFaceUserGroupResponseBodyData struct {
	PageNo        *int32                                             `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	UserGroupList []*QueryFaceUserGroupResponseBodyDataUserGroupList `json:"UserGroupList,omitempty" xml:"UserGroupList,omitempty" type:"Repeated"`
	PageSize      *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total         *int32                                             `json:"Total,omitempty" xml:"Total,omitempty"`
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

func (s *QueryFaceUserGroupResponseBodyData) SetUserGroupList(v []*QueryFaceUserGroupResponseBodyDataUserGroupList) *QueryFaceUserGroupResponseBodyData {
	s.UserGroupList = v
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryFaceUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryFaceUserGroupResponse) SetBody(v *QueryFaceUserGroupResponseBody) *QueryFaceUserGroupResponse {
	s.Body = v
	return s
}

type QueryFaceUserGroupAndDeviceGroupRelationRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	ControlId   *string `json:"ControlId,omitempty" xml:"ControlId,omitempty"`
}

func (s QueryFaceUserGroupAndDeviceGroupRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserGroupAndDeviceGroupRelationRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationRequest) SetApiProduct(v string) *QueryFaceUserGroupAndDeviceGroupRelationRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationRequest) SetApiRevision(v string) *QueryFaceUserGroupAndDeviceGroupRelationRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationRequest) SetIsolationId(v string) *QueryFaceUserGroupAndDeviceGroupRelationRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationRequest) SetControlId(v string) *QueryFaceUserGroupAndDeviceGroupRelationRequest {
	s.ControlId = &v
	return s
}

type QueryFaceUserGroupAndDeviceGroupRelationResponseBody struct {
	RequestId    *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceUserGroupAndDeviceGroupRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserGroupAndDeviceGroupRelationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBody) SetRequestId(v string) *QueryFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.RequestId = &v
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

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBody) SetCode(v string) *QueryFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBody) SetSuccess(v bool) *QueryFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.Success = &v
	return s
}

type QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData struct {
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	ModifiedTime  *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ControlType   *string `json:"ControlType,omitempty" xml:"ControlType,omitempty"`
	UserGroupId   *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	ControlId     *string `json:"ControlId,omitempty" xml:"ControlId,omitempty"`
}

func (s QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) SetDeviceGroupId(v string) *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData {
	s.DeviceGroupId = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) SetModifiedTime(v string) *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) SetControlType(v string) *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData {
	s.ControlType = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) SetUserGroupId(v string) *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData {
	s.UserGroupId = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) SetControlId(v string) *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData {
	s.ControlId = &v
	return s
}

type QueryFaceUserGroupAndDeviceGroupRelationResponse struct {
	Headers map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryFaceUserGroupAndDeviceGroupRelationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponse) SetBody(v *QueryFaceUserGroupAndDeviceGroupRelationResponseBody) *QueryFaceUserGroupAndDeviceGroupRelationResponse {
	s.Body = v
	return s
}

type QueryFaceUserIdByCustomUserIdRequest struct {
	ApiProduct   *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision  *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId  *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	CustomUserId *string `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
}

func (s QueryFaceUserIdByCustomUserIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserIdByCustomUserIdRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceUserIdByCustomUserIdRequest) SetApiProduct(v string) *QueryFaceUserIdByCustomUserIdRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdRequest) SetApiRevision(v string) *QueryFaceUserIdByCustomUserIdRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdRequest) SetIsolationId(v string) *QueryFaceUserIdByCustomUserIdRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdRequest) SetCustomUserId(v string) *QueryFaceUserIdByCustomUserIdRequest {
	s.CustomUserId = &v
	return s
}

type QueryFaceUserIdByCustomUserIdResponseBody struct {
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryFaceUserIdByCustomUserIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceUserIdByCustomUserIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserIdByCustomUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceUserIdByCustomUserIdResponseBody) SetRequestId(v string) *QueryFaceUserIdByCustomUserIdResponseBody {
	s.RequestId = &v
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

func (s *QueryFaceUserIdByCustomUserIdResponseBody) SetCode(v string) *QueryFaceUserIdByCustomUserIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdResponseBody) SetSuccess(v bool) *QueryFaceUserIdByCustomUserIdResponseBody {
	s.Success = &v
	return s
}

type QueryFaceUserIdByCustomUserIdResponseBodyData struct {
	Params       *string `json:"Params,omitempty" xml:"Params,omitempty"`
	CustomUserId *string `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s QueryFaceUserIdByCustomUserIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceUserIdByCustomUserIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceUserIdByCustomUserIdResponseBodyData) SetParams(v string) *QueryFaceUserIdByCustomUserIdResponseBodyData {
	s.Params = &v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdResponseBodyData) SetCustomUserId(v string) *QueryFaceUserIdByCustomUserIdResponseBodyData {
	s.CustomUserId = &v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdResponseBodyData) SetUserId(v string) *QueryFaceUserIdByCustomUserIdResponseBodyData {
	s.UserId = &v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdResponseBodyData) SetName(v string) *QueryFaceUserIdByCustomUserIdResponseBodyData {
	s.Name = &v
	return s
}

type QueryFaceUserIdByCustomUserIdResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryFaceUserIdByCustomUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryFaceUserIdByCustomUserIdResponse) SetBody(v *QueryFaceUserIdByCustomUserIdResponseBody) *QueryFaceUserIdByCustomUserIdResponse {
	s.Body = v
	return s
}

type QueryIotIdsByAIPlanRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	PlanId      *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s QueryIotIdsByAIPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryIotIdsByAIPlanRequest) GoString() string {
	return s.String()
}

func (s *QueryIotIdsByAIPlanRequest) SetApiProduct(v string) *QueryIotIdsByAIPlanRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryIotIdsByAIPlanRequest) SetApiRevision(v string) *QueryIotIdsByAIPlanRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryIotIdsByAIPlanRequest) SetPlanId(v string) *QueryIotIdsByAIPlanRequest {
	s.PlanId = &v
	return s
}

func (s *QueryIotIdsByAIPlanRequest) SetPageSize(v int32) *QueryIotIdsByAIPlanRequest {
	s.PageSize = &v
	return s
}

func (s *QueryIotIdsByAIPlanRequest) SetCurrentPage(v int32) *QueryIotIdsByAIPlanRequest {
	s.CurrentPage = &v
	return s
}

type QueryIotIdsByAIPlanResponseBody struct {
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage *string                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryIotIdsByAIPlanResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryIotIdsByAIPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryIotIdsByAIPlanResponseBody) GoString() string {
	return s.String()
}

func (s *QueryIotIdsByAIPlanResponseBody) SetRequestId(v string) *QueryIotIdsByAIPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryIotIdsByAIPlanResponseBody) SetSuccess(v bool) *QueryIotIdsByAIPlanResponseBody {
	s.Success = &v
	return s
}

func (s *QueryIotIdsByAIPlanResponseBody) SetErrorMessage(v string) *QueryIotIdsByAIPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryIotIdsByAIPlanResponseBody) SetCode(v string) *QueryIotIdsByAIPlanResponseBody {
	s.Code = &v
	return s
}

func (s *QueryIotIdsByAIPlanResponseBody) SetData(v *QueryIotIdsByAIPlanResponseBodyData) *QueryIotIdsByAIPlanResponseBody {
	s.Data = v
	return s
}

type QueryIotIdsByAIPlanResponseBodyData struct {
	Total       *int64    `json:"Total,omitempty" xml:"Total,omitempty"`
	PageCount   *int32    `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	CurrentPage *int32    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	List        []*string `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s QueryIotIdsByAIPlanResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryIotIdsByAIPlanResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryIotIdsByAIPlanResponseBodyData) SetTotal(v int64) *QueryIotIdsByAIPlanResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryIotIdsByAIPlanResponseBodyData) SetPageCount(v int32) *QueryIotIdsByAIPlanResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryIotIdsByAIPlanResponseBodyData) SetCurrentPage(v int32) *QueryIotIdsByAIPlanResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *QueryIotIdsByAIPlanResponseBodyData) SetPageSize(v int32) *QueryIotIdsByAIPlanResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryIotIdsByAIPlanResponseBodyData) SetList(v []*string) *QueryIotIdsByAIPlanResponseBodyData {
	s.List = v
	return s
}

type QueryIotIdsByAIPlanResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryIotIdsByAIPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryIotIdsByAIPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryIotIdsByAIPlanResponse) GoString() string {
	return s.String()
}

func (s *QueryIotIdsByAIPlanResponse) SetHeaders(v map[string]*string) *QueryIotIdsByAIPlanResponse {
	s.Headers = v
	return s
}

func (s *QueryIotIdsByAIPlanResponse) SetBody(v *QueryIotIdsByAIPlanResponseBody) *QueryIotIdsByAIPlanResponse {
	s.Body = v
	return s
}

type QueryLiveStreamingRequest struct {
	ApiProduct    *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	StreamType    *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	ShouldEncrypt *bool   `json:"ShouldEncrypt,omitempty" xml:"ShouldEncrypt,omitempty"`
	EncryptType   *int32  `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	Scheme        *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
}

func (s QueryLiveStreamingRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveStreamingRequest) GoString() string {
	return s.String()
}

func (s *QueryLiveStreamingRequest) SetApiProduct(v string) *QueryLiveStreamingRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetApiRevision(v string) *QueryLiveStreamingRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetIotId(v string) *QueryLiveStreamingRequest {
	s.IotId = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetStreamType(v int32) *QueryLiveStreamingRequest {
	s.StreamType = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetShouldEncrypt(v bool) *QueryLiveStreamingRequest {
	s.ShouldEncrypt = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetEncryptType(v int32) *QueryLiveStreamingRequest {
	s.EncryptType = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetScheme(v string) *QueryLiveStreamingRequest {
	s.Scheme = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetIotInstanceId(v string) *QueryLiveStreamingRequest {
	s.IotInstanceId = &v
	return s
}

type QueryLiveStreamingResponseBody struct {
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryLiveStreamingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryLiveStreamingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveStreamingResponseBody) GoString() string {
	return s.String()
}

func (s *QueryLiveStreamingResponseBody) SetRequestId(v string) *QueryLiveStreamingResponseBody {
	s.RequestId = &v
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

func (s *QueryLiveStreamingResponseBody) SetCode(v string) *QueryLiveStreamingResponseBody {
	s.Code = &v
	return s
}

func (s *QueryLiveStreamingResponseBody) SetSuccess(v bool) *QueryLiveStreamingResponseBody {
	s.Success = &v
	return s
}

type QueryLiveStreamingResponseBodyData struct {
	Path       *string `json:"Path,omitempty" xml:"Path,omitempty"`
	DecryptKey *string `json:"DecryptKey,omitempty" xml:"DecryptKey,omitempty"`
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

func (s *QueryLiveStreamingResponseBodyData) SetDecryptKey(v string) *QueryLiveStreamingResponseBodyData {
	s.DecryptKey = &v
	return s
}

type QueryLiveStreamingResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryLiveStreamingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryLiveStreamingResponse) SetBody(v *QueryLiveStreamingResponseBody) *QueryLiveStreamingResponse {
	s.Body = v
	return s
}

type QueryMonthRecordRequest struct {
	ApiProduct    *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	Month         *string `json:"Month,omitempty" xml:"Month,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
}

func (s QueryMonthRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMonthRecordRequest) GoString() string {
	return s.String()
}

func (s *QueryMonthRecordRequest) SetApiProduct(v string) *QueryMonthRecordRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryMonthRecordRequest) SetApiRevision(v string) *QueryMonthRecordRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryMonthRecordRequest) SetIotId(v string) *QueryMonthRecordRequest {
	s.IotId = &v
	return s
}

func (s *QueryMonthRecordRequest) SetMonth(v string) *QueryMonthRecordRequest {
	s.Month = &v
	return s
}

func (s *QueryMonthRecordRequest) SetIotInstanceId(v string) *QueryMonthRecordRequest {
	s.IotInstanceId = &v
	return s
}

type QueryMonthRecordResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMonthRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMonthRecordResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMonthRecordResponseBody) SetRequestId(v string) *QueryMonthRecordResponseBody {
	s.RequestId = &v
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

func (s *QueryMonthRecordResponseBody) SetCode(v string) *QueryMonthRecordResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMonthRecordResponseBody) SetSuccess(v bool) *QueryMonthRecordResponseBody {
	s.Success = &v
	return s
}

type QueryMonthRecordResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMonthRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryMonthRecordResponse) SetBody(v *QueryMonthRecordResponseBody) *QueryMonthRecordResponse {
	s.Body = v
	return s
}

type QueryPictureFilesRequest struct {
	ApiProduct    *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	BeginTime     *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime       *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PictureType   *int32  `json:"PictureType,omitempty" xml:"PictureType,omitempty"`
	PictureSource *int32  `json:"PictureSource,omitempty" xml:"PictureSource,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
}

func (s QueryPictureFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureFilesRequest) GoString() string {
	return s.String()
}

func (s *QueryPictureFilesRequest) SetApiProduct(v string) *QueryPictureFilesRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryPictureFilesRequest) SetApiRevision(v string) *QueryPictureFilesRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryPictureFilesRequest) SetIotId(v string) *QueryPictureFilesRequest {
	s.IotId = &v
	return s
}

func (s *QueryPictureFilesRequest) SetBeginTime(v int64) *QueryPictureFilesRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryPictureFilesRequest) SetEndTime(v int64) *QueryPictureFilesRequest {
	s.EndTime = &v
	return s
}

func (s *QueryPictureFilesRequest) SetCurrentPage(v int32) *QueryPictureFilesRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryPictureFilesRequest) SetPageSize(v int32) *QueryPictureFilesRequest {
	s.PageSize = &v
	return s
}

func (s *QueryPictureFilesRequest) SetPictureType(v int32) *QueryPictureFilesRequest {
	s.PictureType = &v
	return s
}

func (s *QueryPictureFilesRequest) SetPictureSource(v int32) *QueryPictureFilesRequest {
	s.PictureSource = &v
	return s
}

func (s *QueryPictureFilesRequest) SetIotInstanceId(v string) *QueryPictureFilesRequest {
	s.IotInstanceId = &v
	return s
}

type QueryPictureFilesResponseBody struct {
	RequestId    *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryPictureFilesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryPictureFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureFilesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPictureFilesResponseBody) SetRequestId(v string) *QueryPictureFilesResponseBody {
	s.RequestId = &v
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

func (s *QueryPictureFilesResponseBody) SetCode(v string) *QueryPictureFilesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPictureFilesResponseBody) SetSuccess(v bool) *QueryPictureFilesResponseBody {
	s.Success = &v
	return s
}

type QueryPictureFilesResponseBodyData struct {
	List     []*QueryPictureFilesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageSize *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Page     *int32                                   `json:"Page,omitempty" xml:"Page,omitempty"`
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

func (s *QueryPictureFilesResponseBodyData) SetPageSize(v int32) *QueryPictureFilesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryPictureFilesResponseBodyData) SetPage(v int32) *QueryPictureFilesResponseBodyData {
	s.Page = &v
	return s
}

type QueryPictureFilesResponseBodyDataList struct {
	PicUrl        *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	PicCreateTime *int64  `json:"PicCreateTime,omitempty" xml:"PicCreateTime,omitempty"`
	PicId         *string `json:"PicId,omitempty" xml:"PicId,omitempty"`
	ThumbUrl      *string `json:"ThumbUrl,omitempty" xml:"ThumbUrl,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
}

func (s QueryPictureFilesResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureFilesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryPictureFilesResponseBodyDataList) SetPicUrl(v string) *QueryPictureFilesResponseBodyDataList {
	s.PicUrl = &v
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

func (s *QueryPictureFilesResponseBodyDataList) SetThumbUrl(v string) *QueryPictureFilesResponseBodyDataList {
	s.ThumbUrl = &v
	return s
}

func (s *QueryPictureFilesResponseBodyDataList) SetIotId(v string) *QueryPictureFilesResponseBodyDataList {
	s.IotId = &v
	return s
}

type QueryPictureFilesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryPictureFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryPictureFilesResponse) SetBody(v *QueryPictureFilesResponseBody) *QueryPictureFilesResponse {
	s.Body = v
	return s
}

type QueryPictureSearchAiboxesRequest struct {
	ApiProduct    *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s QueryPictureSearchAiboxesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchAiboxesRequest) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchAiboxesRequest) SetApiProduct(v string) *QueryPictureSearchAiboxesRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryPictureSearchAiboxesRequest) SetApiRevision(v string) *QueryPictureSearchAiboxesRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryPictureSearchAiboxesRequest) SetAppInstanceId(v string) *QueryPictureSearchAiboxesRequest {
	s.AppInstanceId = &v
	return s
}

func (s *QueryPictureSearchAiboxesRequest) SetPageSize(v int32) *QueryPictureSearchAiboxesRequest {
	s.PageSize = &v
	return s
}

func (s *QueryPictureSearchAiboxesRequest) SetCurrentPage(v int32) *QueryPictureSearchAiboxesRequest {
	s.CurrentPage = &v
	return s
}

type QueryPictureSearchAiboxesResponseBody struct {
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryPictureSearchAiboxesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryPictureSearchAiboxesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchAiboxesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchAiboxesResponseBody) SetRequestId(v string) *QueryPictureSearchAiboxesResponseBody {
	s.RequestId = &v
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

func (s *QueryPictureSearchAiboxesResponseBody) SetCode(v string) *QueryPictureSearchAiboxesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPictureSearchAiboxesResponseBody) SetSuccess(v bool) *QueryPictureSearchAiboxesResponseBody {
	s.Success = &v
	return s
}

type QueryPictureSearchAiboxesResponseBodyData struct {
	CurrentPage *int32                                               `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageData    []*QueryPictureSearchAiboxesResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	PageSize    *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total       *int32                                               `json:"Total,omitempty" xml:"Total,omitempty"`
	PageCount   *int32                                               `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
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

func (s *QueryPictureSearchAiboxesResponseBodyData) SetPageCount(v int32) *QueryPictureSearchAiboxesResponseBodyData {
	s.PageCount = &v
	return s
}

type QueryPictureSearchAiboxesResponseBodyDataPageData struct {
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	IotId    *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
}

func (s QueryPictureSearchAiboxesResponseBodyDataPageData) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchAiboxesResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchAiboxesResponseBodyDataPageData) SetNickName(v string) *QueryPictureSearchAiboxesResponseBodyDataPageData {
	s.NickName = &v
	return s
}

func (s *QueryPictureSearchAiboxesResponseBodyDataPageData) SetIotId(v string) *QueryPictureSearchAiboxesResponseBodyDataPageData {
	s.IotId = &v
	return s
}

type QueryPictureSearchAiboxesResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryPictureSearchAiboxesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryPictureSearchAiboxesResponse) SetBody(v *QueryPictureSearchAiboxesResponseBody) *QueryPictureSearchAiboxesResponse {
	s.Body = v
	return s
}

type QueryPictureSearchAppRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
}

func (s QueryPictureSearchAppRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchAppRequest) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchAppRequest) SetApiProduct(v string) *QueryPictureSearchAppRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryPictureSearchAppRequest) SetApiRevision(v string) *QueryPictureSearchAppRequest {
	s.ApiRevision = &v
	return s
}

type QueryPictureSearchAppResponseBody struct {
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryPictureSearchAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryPictureSearchAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchAppResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchAppResponseBody) SetRequestId(v string) *QueryPictureSearchAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPictureSearchAppResponseBody) SetData(v *QueryPictureSearchAppResponseBodyData) *QueryPictureSearchAppResponseBody {
	s.Data = v
	return s
}

func (s *QueryPictureSearchAppResponseBody) SetErrorMessage(v string) *QueryPictureSearchAppResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryPictureSearchAppResponseBody) SetCode(v string) *QueryPictureSearchAppResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPictureSearchAppResponseBody) SetSuccess(v bool) *QueryPictureSearchAppResponseBody {
	s.Success = &v
	return s
}

type QueryPictureSearchAppResponseBodyData struct {
	Data []*QueryPictureSearchAppResponseBodyDataData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s QueryPictureSearchAppResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchAppResponseBodyData) SetData(v []*QueryPictureSearchAppResponseBodyDataData) *QueryPictureSearchAppResponseBodyData {
	s.Data = v
	return s
}

type QueryPictureSearchAppResponseBodyDataData struct {
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	ModifiedTime  *int64  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	CreateTime    *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	AppTemplateId *string `json:"AppTemplateId,omitempty" xml:"AppTemplateId,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Level         *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s QueryPictureSearchAppResponseBodyDataData) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchAppResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchAppResponseBodyDataData) SetAppInstanceId(v string) *QueryPictureSearchAppResponseBodyDataData {
	s.AppInstanceId = &v
	return s
}

func (s *QueryPictureSearchAppResponseBodyDataData) SetModifiedTime(v int64) *QueryPictureSearchAppResponseBodyDataData {
	s.ModifiedTime = &v
	return s
}

func (s *QueryPictureSearchAppResponseBodyDataData) SetVersion(v string) *QueryPictureSearchAppResponseBodyDataData {
	s.Version = &v
	return s
}

func (s *QueryPictureSearchAppResponseBodyDataData) SetCreateTime(v int64) *QueryPictureSearchAppResponseBodyDataData {
	s.CreateTime = &v
	return s
}

func (s *QueryPictureSearchAppResponseBodyDataData) SetAppTemplateId(v string) *QueryPictureSearchAppResponseBodyDataData {
	s.AppTemplateId = &v
	return s
}

func (s *QueryPictureSearchAppResponseBodyDataData) SetName(v string) *QueryPictureSearchAppResponseBodyDataData {
	s.Name = &v
	return s
}

func (s *QueryPictureSearchAppResponseBodyDataData) SetLevel(v string) *QueryPictureSearchAppResponseBodyDataData {
	s.Level = &v
	return s
}

type QueryPictureSearchAppResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryPictureSearchAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPictureSearchAppResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchAppResponse) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchAppResponse) SetHeaders(v map[string]*string) *QueryPictureSearchAppResponse {
	s.Headers = v
	return s
}

func (s *QueryPictureSearchAppResponse) SetBody(v *QueryPictureSearchAppResponseBody) *QueryPictureSearchAppResponse {
	s.Body = v
	return s
}

type QueryPictureSearchDevicesRequest struct {
	ApiProduct    *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s QueryPictureSearchDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchDevicesRequest) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchDevicesRequest) SetApiProduct(v string) *QueryPictureSearchDevicesRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryPictureSearchDevicesRequest) SetApiRevision(v string) *QueryPictureSearchDevicesRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryPictureSearchDevicesRequest) SetAppInstanceId(v string) *QueryPictureSearchDevicesRequest {
	s.AppInstanceId = &v
	return s
}

func (s *QueryPictureSearchDevicesRequest) SetPageSize(v int32) *QueryPictureSearchDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *QueryPictureSearchDevicesRequest) SetCurrentPage(v int32) *QueryPictureSearchDevicesRequest {
	s.CurrentPage = &v
	return s
}

type QueryPictureSearchDevicesResponseBody struct {
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryPictureSearchDevicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryPictureSearchDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchDevicesResponseBody) SetRequestId(v string) *QueryPictureSearchDevicesResponseBody {
	s.RequestId = &v
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

func (s *QueryPictureSearchDevicesResponseBody) SetCode(v string) *QueryPictureSearchDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPictureSearchDevicesResponseBody) SetSuccess(v bool) *QueryPictureSearchDevicesResponseBody {
	s.Success = &v
	return s
}

type QueryPictureSearchDevicesResponseBodyData struct {
	CurrentPage *int32                                               `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageData    []*QueryPictureSearchDevicesResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	PageSize    *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total       *int32                                               `json:"Total,omitempty" xml:"Total,omitempty"`
	PageCount   *int32                                               `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
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

func (s *QueryPictureSearchDevicesResponseBodyData) SetPageCount(v int32) *QueryPictureSearchDevicesResponseBodyData {
	s.PageCount = &v
	return s
}

type QueryPictureSearchDevicesResponseBodyDataPageData struct {
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	IotId    *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
}

func (s QueryPictureSearchDevicesResponseBodyDataPageData) String() string {
	return tea.Prettify(s)
}

func (s QueryPictureSearchDevicesResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchDevicesResponseBodyDataPageData) SetNickName(v string) *QueryPictureSearchDevicesResponseBodyDataPageData {
	s.NickName = &v
	return s
}

func (s *QueryPictureSearchDevicesResponseBodyDataPageData) SetIotId(v string) *QueryPictureSearchDevicesResponseBodyDataPageData {
	s.IotId = &v
	return s
}

type QueryPictureSearchDevicesResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryPictureSearchDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryPictureSearchDevicesResponse) SetBody(v *QueryPictureSearchDevicesResponseBody) *QueryPictureSearchDevicesResponse {
	s.Body = v
	return s
}

type QueryRecordRequest struct {
	ApiProduct    *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	StreamType    *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	BeginTime     *int32  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime       *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RecordType    *int32  `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	NeedSnapshot  *bool   `json:"NeedSnapshot,omitempty" xml:"NeedSnapshot,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
}

func (s QueryRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordRequest) SetApiProduct(v string) *QueryRecordRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryRecordRequest) SetApiRevision(v string) *QueryRecordRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryRecordRequest) SetIotId(v string) *QueryRecordRequest {
	s.IotId = &v
	return s
}

func (s *QueryRecordRequest) SetStreamType(v int32) *QueryRecordRequest {
	s.StreamType = &v
	return s
}

func (s *QueryRecordRequest) SetBeginTime(v int32) *QueryRecordRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryRecordRequest) SetEndTime(v int32) *QueryRecordRequest {
	s.EndTime = &v
	return s
}

func (s *QueryRecordRequest) SetRecordType(v int32) *QueryRecordRequest {
	s.RecordType = &v
	return s
}

func (s *QueryRecordRequest) SetCurrentPage(v int32) *QueryRecordRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryRecordRequest) SetPageSize(v int32) *QueryRecordRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRecordRequest) SetNeedSnapshot(v bool) *QueryRecordRequest {
	s.NeedSnapshot = &v
	return s
}

func (s *QueryRecordRequest) SetIotInstanceId(v string) *QueryRecordRequest {
	s.IotInstanceId = &v
	return s
}

type QueryRecordResponseBody struct {
	RequestId    *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecordResponseBody) SetRequestId(v string) *QueryRecordResponseBody {
	s.RequestId = &v
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

func (s *QueryRecordResponseBody) SetCode(v string) *QueryRecordResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRecordResponseBody) SetSuccess(v bool) *QueryRecordResponseBody {
	s.Success = &v
	return s
}

type QueryRecordResponseBodyData struct {
	List     []*QueryRecordResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageSize *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Page     *int32                             `json:"Page,omitempty" xml:"Page,omitempty"`
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

func (s *QueryRecordResponseBodyData) SetPageSize(v int32) *QueryRecordResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryRecordResponseBodyData) SetPage(v int32) *QueryRecordResponseBodyData {
	s.Page = &v
	return s
}

type QueryRecordResponseBodyDataList struct {
	SnapshotUrl      *string `json:"SnapshotUrl,omitempty" xml:"SnapshotUrl,omitempty"`
	EndTime          *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RecordType       *int32  `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	StreamType       *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	BeginTime        *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	FileName         *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	VideoFrameNumber *int32  `json:"VideoFrameNumber,omitempty" xml:"VideoFrameNumber,omitempty"`
	FileSize         *int32  `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
}

func (s QueryRecordResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryRecordResponseBodyDataList) SetSnapshotUrl(v string) *QueryRecordResponseBodyDataList {
	s.SnapshotUrl = &v
	return s
}

func (s *QueryRecordResponseBodyDataList) SetEndTime(v string) *QueryRecordResponseBodyDataList {
	s.EndTime = &v
	return s
}

func (s *QueryRecordResponseBodyDataList) SetRecordType(v int32) *QueryRecordResponseBodyDataList {
	s.RecordType = &v
	return s
}

func (s *QueryRecordResponseBodyDataList) SetStreamType(v int32) *QueryRecordResponseBodyDataList {
	s.StreamType = &v
	return s
}

func (s *QueryRecordResponseBodyDataList) SetBeginTime(v string) *QueryRecordResponseBodyDataList {
	s.BeginTime = &v
	return s
}

func (s *QueryRecordResponseBodyDataList) SetFileName(v string) *QueryRecordResponseBodyDataList {
	s.FileName = &v
	return s
}

func (s *QueryRecordResponseBodyDataList) SetVideoFrameNumber(v int32) *QueryRecordResponseBodyDataList {
	s.VideoFrameNumber = &v
	return s
}

func (s *QueryRecordResponseBodyDataList) SetFileSize(v int32) *QueryRecordResponseBodyDataList {
	s.FileSize = &v
	return s
}

type QueryRecordResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryRecordResponse) SetBody(v *QueryRecordResponseBody) *QueryRecordResponse {
	s.Body = v
	return s
}

type QueryRecordByRecordIdRequest struct {
	ApiProduct    *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	RecordId      *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
}

func (s QueryRecordByRecordIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordByRecordIdRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordByRecordIdRequest) SetApiProduct(v string) *QueryRecordByRecordIdRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryRecordByRecordIdRequest) SetApiRevision(v string) *QueryRecordByRecordIdRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryRecordByRecordIdRequest) SetIotId(v string) *QueryRecordByRecordIdRequest {
	s.IotId = &v
	return s
}

func (s *QueryRecordByRecordIdRequest) SetRecordId(v string) *QueryRecordByRecordIdRequest {
	s.RecordId = &v
	return s
}

func (s *QueryRecordByRecordIdRequest) SetIotInstanceId(v string) *QueryRecordByRecordIdRequest {
	s.IotInstanceId = &v
	return s
}

type QueryRecordByRecordIdResponseBody struct {
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         []*QueryRecordByRecordIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMessage *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRecordByRecordIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordByRecordIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecordByRecordIdResponseBody) SetRequestId(v string) *QueryRecordByRecordIdResponseBody {
	s.RequestId = &v
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

func (s *QueryRecordByRecordIdResponseBody) SetCode(v string) *QueryRecordByRecordIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRecordByRecordIdResponseBody) SetSuccess(v bool) *QueryRecordByRecordIdResponseBody {
	s.Success = &v
	return s
}

type QueryRecordByRecordIdResponseBodyData struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	FileName  *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	VodUrl    *string `json:"VodUrl,omitempty" xml:"VodUrl,omitempty"`
}

func (s QueryRecordByRecordIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordByRecordIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRecordByRecordIdResponseBodyData) SetEndTime(v string) *QueryRecordByRecordIdResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *QueryRecordByRecordIdResponseBodyData) SetBeginTime(v string) *QueryRecordByRecordIdResponseBodyData {
	s.BeginTime = &v
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRecordByRecordIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryRecordByRecordIdResponse) SetBody(v *QueryRecordByRecordIdResponseBody) *QueryRecordByRecordIdResponse {
	s.Body = v
	return s
}

type QueryRecordPlanDetailRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	PlanId      *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s QueryRecordPlanDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDetailRequest) SetApiProduct(v string) *QueryRecordPlanDetailRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryRecordPlanDetailRequest) SetApiRevision(v string) *QueryRecordPlanDetailRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryRecordPlanDetailRequest) SetPlanId(v string) *QueryRecordPlanDetailRequest {
	s.PlanId = &v
	return s
}

type QueryRecordPlanDetailResponseBody struct {
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryRecordPlanDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRecordPlanDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDetailResponseBody) SetRequestId(v string) *QueryRecordPlanDetailResponseBody {
	s.RequestId = &v
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

func (s *QueryRecordPlanDetailResponseBody) SetCode(v string) *QueryRecordPlanDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRecordPlanDetailResponseBody) SetSuccess(v bool) *QueryRecordPlanDetailResponseBody {
	s.Success = &v
	return s
}

type QueryRecordPlanDetailResponseBodyData struct {
	PlanId       *string                                            `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	Name         *string                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateInfo *QueryRecordPlanDetailResponseBodyDataTemplateInfo `json:"TemplateInfo,omitempty" xml:"TemplateInfo,omitempty" type:"Struct"`
	TemplateId   *string                                            `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s QueryRecordPlanDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDetailResponseBodyData) SetPlanId(v string) *QueryRecordPlanDetailResponseBodyData {
	s.PlanId = &v
	return s
}

func (s *QueryRecordPlanDetailResponseBodyData) SetName(v string) *QueryRecordPlanDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryRecordPlanDetailResponseBodyData) SetTemplateInfo(v *QueryRecordPlanDetailResponseBodyDataTemplateInfo) *QueryRecordPlanDetailResponseBodyData {
	s.TemplateInfo = v
	return s
}

func (s *QueryRecordPlanDetailResponseBodyData) SetTemplateId(v string) *QueryRecordPlanDetailResponseBodyData {
	s.TemplateId = &v
	return s
}

type QueryRecordPlanDetailResponseBodyDataTemplateInfo struct {
	TimeSectionList []*QueryRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList `json:"TimeSectionList,omitempty" xml:"TimeSectionList,omitempty" type:"Repeated"`
	AllDay          *int32                                                              `json:"AllDay,omitempty" xml:"AllDay,omitempty"`
	Default         *int32                                                              `json:"Default,omitempty" xml:"Default,omitempty"`
	Name            *string                                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateId      *string                                                             `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s QueryRecordPlanDetailResponseBodyDataTemplateInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDetailResponseBodyDataTemplateInfo) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDetailResponseBodyDataTemplateInfo) SetTimeSectionList(v []*QueryRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList) *QueryRecordPlanDetailResponseBodyDataTemplateInfo {
	s.TimeSectionList = v
	return s
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

type QueryRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList struct {
	DayOfWeek *int32 `json:"DayOfWeek,omitempty" xml:"DayOfWeek,omitempty"`
	Begin     *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	End       *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s QueryRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList) SetDayOfWeek(v int32) *QueryRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList {
	s.DayOfWeek = &v
	return s
}

func (s *QueryRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList) SetBegin(v int32) *QueryRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList {
	s.Begin = &v
	return s
}

func (s *QueryRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList) SetEnd(v int32) *QueryRecordPlanDetailResponseBodyDataTemplateInfoTimeSectionList {
	s.End = &v
	return s
}

type QueryRecordPlanDetailResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRecordPlanDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryRecordPlanDetailResponse) SetBody(v *QueryRecordPlanDetailResponseBody) *QueryRecordPlanDetailResponse {
	s.Body = v
	return s
}

type QueryRecordPlanDeviceByDeviceRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId       *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	StreamType  *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s QueryRecordPlanDeviceByDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDeviceByDeviceRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDeviceByDeviceRequest) SetApiProduct(v string) *QueryRecordPlanDeviceByDeviceRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceRequest) SetApiRevision(v string) *QueryRecordPlanDeviceByDeviceRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceRequest) SetIotId(v string) *QueryRecordPlanDeviceByDeviceRequest {
	s.IotId = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceRequest) SetStreamType(v int32) *QueryRecordPlanDeviceByDeviceRequest {
	s.StreamType = &v
	return s
}

type QueryRecordPlanDeviceByDeviceResponseBody struct {
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryRecordPlanDeviceByDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRecordPlanDeviceByDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDeviceByDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDeviceByDeviceResponseBody) SetRequestId(v string) *QueryRecordPlanDeviceByDeviceResponseBody {
	s.RequestId = &v
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

func (s *QueryRecordPlanDeviceByDeviceResponseBody) SetCode(v string) *QueryRecordPlanDeviceByDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceResponseBody) SetSuccess(v bool) *QueryRecordPlanDeviceByDeviceResponseBody {
	s.Success = &v
	return s
}

type QueryRecordPlanDeviceByDeviceResponseBodyData struct {
	PlanId       *string                                                    `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	Name         *string                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateInfo *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo `json:"TemplateInfo,omitempty" xml:"TemplateInfo,omitempty" type:"Struct"`
	TemplateId   *string                                                    `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s QueryRecordPlanDeviceByDeviceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDeviceByDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDeviceByDeviceResponseBodyData) SetPlanId(v string) *QueryRecordPlanDeviceByDeviceResponseBodyData {
	s.PlanId = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceResponseBodyData) SetName(v string) *QueryRecordPlanDeviceByDeviceResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceResponseBodyData) SetTemplateInfo(v *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo) *QueryRecordPlanDeviceByDeviceResponseBodyData {
	s.TemplateInfo = v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceResponseBodyData) SetTemplateId(v string) *QueryRecordPlanDeviceByDeviceResponseBodyData {
	s.TemplateId = &v
	return s
}

type QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo struct {
	TimeSectionList []*QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList `json:"TimeSectionList,omitempty" xml:"TimeSectionList,omitempty" type:"Repeated"`
	AllDay          *int32                                                                      `json:"AllDay,omitempty" xml:"AllDay,omitempty"`
	Default         *int32                                                                      `json:"Default,omitempty" xml:"Default,omitempty"`
	Name            *string                                                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateId      *string                                                                     `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo) SetTimeSectionList(v []*QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList) *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfo {
	s.TimeSectionList = v
	return s
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

type QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList struct {
	DayOfWeek *int32 `json:"DayOfWeek,omitempty" xml:"DayOfWeek,omitempty"`
	Begin     *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	End       *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList) SetDayOfWeek(v int32) *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList {
	s.DayOfWeek = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList) SetBegin(v int32) *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList {
	s.Begin = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList) SetEnd(v int32) *QueryRecordPlanDeviceByDeviceResponseBodyDataTemplateInfoTimeSectionList {
	s.End = &v
	return s
}

type QueryRecordPlanDeviceByDeviceResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRecordPlanDeviceByDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryRecordPlanDeviceByDeviceResponse) SetBody(v *QueryRecordPlanDeviceByDeviceResponseBody) *QueryRecordPlanDeviceByDeviceResponse {
	s.Body = v
	return s
}

type QueryRecordPlanDeviceByPlanRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	PlanId      *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryRecordPlanDeviceByPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDeviceByPlanRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDeviceByPlanRequest) SetApiProduct(v string) *QueryRecordPlanDeviceByPlanRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryRecordPlanDeviceByPlanRequest) SetApiRevision(v string) *QueryRecordPlanDeviceByPlanRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryRecordPlanDeviceByPlanRequest) SetPlanId(v string) *QueryRecordPlanDeviceByPlanRequest {
	s.PlanId = &v
	return s
}

func (s *QueryRecordPlanDeviceByPlanRequest) SetCurrentPage(v int32) *QueryRecordPlanDeviceByPlanRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryRecordPlanDeviceByPlanRequest) SetPageSize(v int32) *QueryRecordPlanDeviceByPlanRequest {
	s.PageSize = &v
	return s
}

type QueryRecordPlanDeviceByPlanResponseBody struct {
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryRecordPlanDeviceByPlanResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRecordPlanDeviceByPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDeviceByPlanResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDeviceByPlanResponseBody) SetRequestId(v string) *QueryRecordPlanDeviceByPlanResponseBody {
	s.RequestId = &v
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

func (s *QueryRecordPlanDeviceByPlanResponseBody) SetCode(v string) *QueryRecordPlanDeviceByPlanResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRecordPlanDeviceByPlanResponseBody) SetSuccess(v bool) *QueryRecordPlanDeviceByPlanResponseBody {
	s.Success = &v
	return s
}

type QueryRecordPlanDeviceByPlanResponseBodyData struct {
	List      []*QueryRecordPlanDeviceByPlanResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageSize  *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int32                                             `json:"Total,omitempty" xml:"Total,omitempty"`
	PageCount *int32                                             `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	Page      *int32                                             `json:"Page,omitempty" xml:"Page,omitempty"`
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

func (s *QueryRecordPlanDeviceByPlanResponseBodyData) SetPageSize(v int32) *QueryRecordPlanDeviceByPlanResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryRecordPlanDeviceByPlanResponseBodyData) SetTotal(v int32) *QueryRecordPlanDeviceByPlanResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryRecordPlanDeviceByPlanResponseBodyData) SetPageCount(v int32) *QueryRecordPlanDeviceByPlanResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryRecordPlanDeviceByPlanResponseBodyData) SetPage(v int32) *QueryRecordPlanDeviceByPlanResponseBodyData {
	s.Page = &v
	return s
}

type QueryRecordPlanDeviceByPlanResponseBodyDataList struct {
	StreamType *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
}

func (s QueryRecordPlanDeviceByPlanResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlanDeviceByPlanResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDeviceByPlanResponseBodyDataList) SetStreamType(v int32) *QueryRecordPlanDeviceByPlanResponseBodyDataList {
	s.StreamType = &v
	return s
}

func (s *QueryRecordPlanDeviceByPlanResponseBodyDataList) SetIotId(v string) *QueryRecordPlanDeviceByPlanResponseBodyDataList {
	s.IotId = &v
	return s
}

type QueryRecordPlanDeviceByPlanResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRecordPlanDeviceByPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryRecordPlanDeviceByPlanResponse) SetBody(v *QueryRecordPlanDeviceByPlanResponseBody) *QueryRecordPlanDeviceByPlanResponse {
	s.Body = v
	return s
}

type QueryRecordPlansRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryRecordPlansRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlansRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordPlansRequest) SetApiProduct(v string) *QueryRecordPlansRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryRecordPlansRequest) SetApiRevision(v string) *QueryRecordPlansRequest {
	s.ApiRevision = &v
	return s
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
	RequestId    *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryRecordPlansResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRecordPlansResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlansResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecordPlansResponseBody) SetRequestId(v string) *QueryRecordPlansResponseBody {
	s.RequestId = &v
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

func (s *QueryRecordPlansResponseBody) SetCode(v string) *QueryRecordPlansResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRecordPlansResponseBody) SetSuccess(v bool) *QueryRecordPlansResponseBody {
	s.Success = &v
	return s
}

type QueryRecordPlansResponseBodyData struct {
	List      []*QueryRecordPlansResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageSize  *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageCount *int32                                  `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	Total     *int32                                  `json:"Total,omitempty" xml:"Total,omitempty"`
	Page      *int32                                  `json:"Page,omitempty" xml:"Page,omitempty"`
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

func (s *QueryRecordPlansResponseBodyData) SetPageSize(v int32) *QueryRecordPlansResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryRecordPlansResponseBodyData) SetPageCount(v int32) *QueryRecordPlansResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryRecordPlansResponseBodyData) SetTotal(v int32) *QueryRecordPlansResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryRecordPlansResponseBodyData) SetPage(v int32) *QueryRecordPlansResponseBodyData {
	s.Page = &v
	return s
}

type QueryRecordPlansResponseBodyDataList struct {
	PlanId     *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s QueryRecordPlansResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordPlansResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryRecordPlansResponseBodyDataList) SetPlanId(v string) *QueryRecordPlansResponseBodyDataList {
	s.PlanId = &v
	return s
}

func (s *QueryRecordPlansResponseBodyDataList) SetName(v string) *QueryRecordPlansResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *QueryRecordPlansResponseBodyDataList) SetTemplateId(v string) *QueryRecordPlansResponseBodyDataList {
	s.TemplateId = &v
	return s
}

type QueryRecordPlansResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRecordPlansResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryRecordPlansResponse) SetBody(v *QueryRecordPlansResponseBody) *QueryRecordPlansResponse {
	s.Body = v
	return s
}

type QueryRecordUrlRequest struct {
	ApiProduct    *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	FileName      *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
}

func (s QueryRecordUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordUrlRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordUrlRequest) SetApiProduct(v string) *QueryRecordUrlRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryRecordUrlRequest) SetApiRevision(v string) *QueryRecordUrlRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryRecordUrlRequest) SetIotId(v string) *QueryRecordUrlRequest {
	s.IotId = &v
	return s
}

func (s *QueryRecordUrlRequest) SetFileName(v string) *QueryRecordUrlRequest {
	s.FileName = &v
	return s
}

func (s *QueryRecordUrlRequest) SetIotInstanceId(v string) *QueryRecordUrlRequest {
	s.IotInstanceId = &v
	return s
}

type QueryRecordUrlResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRecordUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordUrlResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecordUrlResponseBody) SetRequestId(v string) *QueryRecordUrlResponseBody {
	s.RequestId = &v
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

func (s *QueryRecordUrlResponseBody) SetCode(v string) *QueryRecordUrlResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRecordUrlResponseBody) SetSuccess(v bool) *QueryRecordUrlResponseBody {
	s.Success = &v
	return s
}

type QueryRecordUrlResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRecordUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryRecordUrlResponse) SetBody(v *QueryRecordUrlResponseBody) *QueryRecordUrlResponse {
	s.Body = v
	return s
}

type QueryStandardAIAppTemplatesRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s QueryStandardAIAppTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryStandardAIAppTemplatesRequest) GoString() string {
	return s.String()
}

func (s *QueryStandardAIAppTemplatesRequest) SetApiProduct(v string) *QueryStandardAIAppTemplatesRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryStandardAIAppTemplatesRequest) SetApiRevision(v string) *QueryStandardAIAppTemplatesRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryStandardAIAppTemplatesRequest) SetPageSize(v int32) *QueryStandardAIAppTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *QueryStandardAIAppTemplatesRequest) SetCurrentPage(v int32) *QueryStandardAIAppTemplatesRequest {
	s.CurrentPage = &v
	return s
}

type QueryStandardAIAppTemplatesResponseBody struct {
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage *string                                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryStandardAIAppTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryStandardAIAppTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryStandardAIAppTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryStandardAIAppTemplatesResponseBody) SetRequestId(v string) *QueryStandardAIAppTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryStandardAIAppTemplatesResponseBody) SetSuccess(v bool) *QueryStandardAIAppTemplatesResponseBody {
	s.Success = &v
	return s
}

func (s *QueryStandardAIAppTemplatesResponseBody) SetErrorMessage(v string) *QueryStandardAIAppTemplatesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryStandardAIAppTemplatesResponseBody) SetCode(v string) *QueryStandardAIAppTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryStandardAIAppTemplatesResponseBody) SetData(v *QueryStandardAIAppTemplatesResponseBodyData) *QueryStandardAIAppTemplatesResponseBody {
	s.Data = v
	return s
}

type QueryStandardAIAppTemplatesResponseBodyData struct {
	Total       *int64                                             `json:"Total,omitempty" xml:"Total,omitempty"`
	PageCount   *int32                                             `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	CurrentPage *int32                                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	List        []*QueryStandardAIAppTemplatesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s QueryStandardAIAppTemplatesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryStandardAIAppTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryStandardAIAppTemplatesResponseBodyData) SetTotal(v int64) *QueryStandardAIAppTemplatesResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryStandardAIAppTemplatesResponseBodyData) SetPageCount(v int32) *QueryStandardAIAppTemplatesResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryStandardAIAppTemplatesResponseBodyData) SetCurrentPage(v int32) *QueryStandardAIAppTemplatesResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *QueryStandardAIAppTemplatesResponseBodyData) SetPageSize(v int32) *QueryStandardAIAppTemplatesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryStandardAIAppTemplatesResponseBodyData) SetList(v []*QueryStandardAIAppTemplatesResponseBodyDataList) *QueryStandardAIAppTemplatesResponseBodyData {
	s.List = v
	return s
}

type QueryStandardAIAppTemplatesResponseBodyDataList struct {
	AppTemplateId *string `json:"AppTemplateId,omitempty" xml:"AppTemplateId,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s QueryStandardAIAppTemplatesResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryStandardAIAppTemplatesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryStandardAIAppTemplatesResponseBodyDataList) SetAppTemplateId(v string) *QueryStandardAIAppTemplatesResponseBodyDataList {
	s.AppTemplateId = &v
	return s
}

func (s *QueryStandardAIAppTemplatesResponseBodyDataList) SetVersion(v string) *QueryStandardAIAppTemplatesResponseBodyDataList {
	s.Version = &v
	return s
}

func (s *QueryStandardAIAppTemplatesResponseBodyDataList) SetName(v string) *QueryStandardAIAppTemplatesResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *QueryStandardAIAppTemplatesResponseBodyDataList) SetDescription(v string) *QueryStandardAIAppTemplatesResponseBodyDataList {
	s.Description = &v
	return s
}

type QueryStandardAIAppTemplatesResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryStandardAIAppTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryStandardAIAppTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryStandardAIAppTemplatesResponse) GoString() string {
	return s.String()
}

func (s *QueryStandardAIAppTemplatesResponse) SetHeaders(v map[string]*string) *QueryStandardAIAppTemplatesResponse {
	s.Headers = v
	return s
}

func (s *QueryStandardAIAppTemplatesResponse) SetBody(v *QueryStandardAIAppTemplatesResponseBody) *QueryStandardAIAppTemplatesResponse {
	s.Body = v
	return s
}

type QueryTimeTemplateRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s QueryTimeTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTimeTemplateRequest) GoString() string {
	return s.String()
}

func (s *QueryTimeTemplateRequest) SetApiProduct(v string) *QueryTimeTemplateRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryTimeTemplateRequest) SetApiRevision(v string) *QueryTimeTemplateRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryTimeTemplateRequest) SetPageSize(v int32) *QueryTimeTemplateRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTimeTemplateRequest) SetCurrentPage(v int32) *QueryTimeTemplateRequest {
	s.CurrentPage = &v
	return s
}

type QueryTimeTemplateResponseBody struct {
	RequestId    *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryTimeTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTimeTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTimeTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTimeTemplateResponseBody) SetRequestId(v string) *QueryTimeTemplateResponseBody {
	s.RequestId = &v
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

func (s *QueryTimeTemplateResponseBody) SetCode(v string) *QueryTimeTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTimeTemplateResponseBody) SetSuccess(v bool) *QueryTimeTemplateResponseBody {
	s.Success = &v
	return s
}

type QueryTimeTemplateResponseBodyData struct {
	List      []*QueryTimeTemplateResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageSize  *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int32                                   `json:"Total,omitempty" xml:"Total,omitempty"`
	PageCount *int32                                   `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	Page      *int32                                   `json:"Page,omitempty" xml:"Page,omitempty"`
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

func (s *QueryTimeTemplateResponseBodyData) SetPageSize(v int32) *QueryTimeTemplateResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryTimeTemplateResponseBodyData) SetTotal(v int32) *QueryTimeTemplateResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryTimeTemplateResponseBodyData) SetPageCount(v int32) *QueryTimeTemplateResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryTimeTemplateResponseBodyData) SetPage(v int32) *QueryTimeTemplateResponseBodyData {
	s.Page = &v
	return s
}

type QueryTimeTemplateResponseBodyDataList struct {
	TimeSectionList []*QueryTimeTemplateResponseBodyDataListTimeSectionList `json:"TimeSectionList,omitempty" xml:"TimeSectionList,omitempty" type:"Repeated"`
	AllDay          *int32                                                  `json:"AllDay,omitempty" xml:"AllDay,omitempty"`
	Default         *int32                                                  `json:"Default,omitempty" xml:"Default,omitempty"`
	Name            *string                                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateId      *string                                                 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s QueryTimeTemplateResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryTimeTemplateResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryTimeTemplateResponseBodyDataList) SetTimeSectionList(v []*QueryTimeTemplateResponseBodyDataListTimeSectionList) *QueryTimeTemplateResponseBodyDataList {
	s.TimeSectionList = v
	return s
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

type QueryTimeTemplateResponseBodyDataListTimeSectionList struct {
	DayOfWeek *int32 `json:"DayOfWeek,omitempty" xml:"DayOfWeek,omitempty"`
	Begin     *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	End       *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s QueryTimeTemplateResponseBodyDataListTimeSectionList) String() string {
	return tea.Prettify(s)
}

func (s QueryTimeTemplateResponseBodyDataListTimeSectionList) GoString() string {
	return s.String()
}

func (s *QueryTimeTemplateResponseBodyDataListTimeSectionList) SetDayOfWeek(v int32) *QueryTimeTemplateResponseBodyDataListTimeSectionList {
	s.DayOfWeek = &v
	return s
}

func (s *QueryTimeTemplateResponseBodyDataListTimeSectionList) SetBegin(v int32) *QueryTimeTemplateResponseBodyDataListTimeSectionList {
	s.Begin = &v
	return s
}

func (s *QueryTimeTemplateResponseBodyDataListTimeSectionList) SetEnd(v int32) *QueryTimeTemplateResponseBodyDataListTimeSectionList {
	s.End = &v
	return s
}

type QueryTimeTemplateResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTimeTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryTimeTemplateResponse) SetBody(v *QueryTimeTemplateResponseBody) *QueryTimeTemplateResponse {
	s.Body = v
	return s
}

type QueryTimeTemplateDetailRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	TemplateId  *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s QueryTimeTemplateDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTimeTemplateDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryTimeTemplateDetailRequest) SetApiProduct(v string) *QueryTimeTemplateDetailRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryTimeTemplateDetailRequest) SetApiRevision(v string) *QueryTimeTemplateDetailRequest {
	s.ApiRevision = &v
	return s
}

func (s *QueryTimeTemplateDetailRequest) SetTemplateId(v string) *QueryTimeTemplateDetailRequest {
	s.TemplateId = &v
	return s
}

type QueryTimeTemplateDetailResponseBody struct {
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryTimeTemplateDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTimeTemplateDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTimeTemplateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTimeTemplateDetailResponseBody) SetRequestId(v string) *QueryTimeTemplateDetailResponseBody {
	s.RequestId = &v
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

func (s *QueryTimeTemplateDetailResponseBody) SetCode(v string) *QueryTimeTemplateDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTimeTemplateDetailResponseBody) SetSuccess(v bool) *QueryTimeTemplateDetailResponseBody {
	s.Success = &v
	return s
}

type QueryTimeTemplateDetailResponseBodyData struct {
	TimeSectionList []*QueryTimeTemplateDetailResponseBodyDataTimeSectionList `json:"TimeSectionList,omitempty" xml:"TimeSectionList,omitempty" type:"Repeated"`
	AllDay          *int32                                                    `json:"AllDay,omitempty" xml:"AllDay,omitempty"`
	Default         *int32                                                    `json:"Default,omitempty" xml:"Default,omitempty"`
	Name            *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateId      *string                                                   `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s QueryTimeTemplateDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTimeTemplateDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTimeTemplateDetailResponseBodyData) SetTimeSectionList(v []*QueryTimeTemplateDetailResponseBodyDataTimeSectionList) *QueryTimeTemplateDetailResponseBodyData {
	s.TimeSectionList = v
	return s
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

type QueryTimeTemplateDetailResponseBodyDataTimeSectionList struct {
	DayOfWeek *int32 `json:"DayOfWeek,omitempty" xml:"DayOfWeek,omitempty"`
	Begin     *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	End       *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s QueryTimeTemplateDetailResponseBodyDataTimeSectionList) String() string {
	return tea.Prettify(s)
}

func (s QueryTimeTemplateDetailResponseBodyDataTimeSectionList) GoString() string {
	return s.String()
}

func (s *QueryTimeTemplateDetailResponseBodyDataTimeSectionList) SetDayOfWeek(v int32) *QueryTimeTemplateDetailResponseBodyDataTimeSectionList {
	s.DayOfWeek = &v
	return s
}

func (s *QueryTimeTemplateDetailResponseBodyDataTimeSectionList) SetBegin(v int32) *QueryTimeTemplateDetailResponseBodyDataTimeSectionList {
	s.Begin = &v
	return s
}

func (s *QueryTimeTemplateDetailResponseBodyDataTimeSectionList) SetEnd(v int32) *QueryTimeTemplateDetailResponseBodyDataTimeSectionList {
	s.End = &v
	return s
}

type QueryTimeTemplateDetailResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTimeTemplateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryTimeTemplateDetailResponse) SetBody(v *QueryTimeTemplateDetailResponseBody) *QueryTimeTemplateDetailResponse {
	s.Body = v
	return s
}

type QueryVoiceIntercomRequest struct {
	ApiProduct    *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
}

func (s QueryVoiceIntercomRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryVoiceIntercomRequest) GoString() string {
	return s.String()
}

func (s *QueryVoiceIntercomRequest) SetApiProduct(v string) *QueryVoiceIntercomRequest {
	s.ApiProduct = &v
	return s
}

func (s *QueryVoiceIntercomRequest) SetApiRevision(v string) *QueryVoiceIntercomRequest {
	s.ApiRevision = &v
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

type QueryVoiceIntercomResponseBody struct {
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *QueryVoiceIntercomResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryVoiceIntercomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryVoiceIntercomResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVoiceIntercomResponseBody) SetRequestId(v string) *QueryVoiceIntercomResponseBody {
	s.RequestId = &v
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

func (s *QueryVoiceIntercomResponseBody) SetCode(v string) *QueryVoiceIntercomResponseBody {
	s.Code = &v
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
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Iv  *string `json:"Iv,omitempty" xml:"Iv,omitempty"`
}

func (s QueryVoiceIntercomResponseBodyDataCryptoKey) String() string {
	return tea.Prettify(s)
}

func (s QueryVoiceIntercomResponseBodyDataCryptoKey) GoString() string {
	return s.String()
}

func (s *QueryVoiceIntercomResponseBodyDataCryptoKey) SetKey(v string) *QueryVoiceIntercomResponseBodyDataCryptoKey {
	s.Key = &v
	return s
}

func (s *QueryVoiceIntercomResponseBodyDataCryptoKey) SetIv(v string) *QueryVoiceIntercomResponseBodyDataCryptoKey {
	s.Iv = &v
	return s
}

type QueryVoiceIntercomResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryVoiceIntercomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryVoiceIntercomResponse) SetBody(v *QueryVoiceIntercomResponseBody) *QueryVoiceIntercomResponse {
	s.Body = v
	return s
}

type RemoveAIAppRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s RemoveAIAppRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveAIAppRequest) GoString() string {
	return s.String()
}

func (s *RemoveAIAppRequest) SetApiProduct(v string) *RemoveAIAppRequest {
	s.ApiProduct = &v
	return s
}

func (s *RemoveAIAppRequest) SetApiRevision(v string) *RemoveAIAppRequest {
	s.ApiRevision = &v
	return s
}

func (s *RemoveAIAppRequest) SetAppId(v string) *RemoveAIAppRequest {
	s.AppId = &v
	return s
}

type RemoveAIAppResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s RemoveAIAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveAIAppResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveAIAppResponseBody) SetRequestId(v string) *RemoveAIAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveAIAppResponseBody) SetSuccess(v bool) *RemoveAIAppResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveAIAppResponseBody) SetErrorMessage(v string) *RemoveAIAppResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RemoveAIAppResponseBody) SetCode(v string) *RemoveAIAppResponseBody {
	s.Code = &v
	return s
}

type RemoveAIAppResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveAIAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveAIAppResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveAIAppResponse) GoString() string {
	return s.String()
}

func (s *RemoveAIAppResponse) SetHeaders(v map[string]*string) *RemoveAIAppResponse {
	s.Headers = v
	return s
}

func (s *RemoveAIAppResponse) SetBody(v *RemoveAIAppResponseBody) *RemoveAIAppResponse {
	s.Body = v
	return s
}

type RemoveAIPlanRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	PlanId      *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s RemoveAIPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveAIPlanRequest) GoString() string {
	return s.String()
}

func (s *RemoveAIPlanRequest) SetApiProduct(v string) *RemoveAIPlanRequest {
	s.ApiProduct = &v
	return s
}

func (s *RemoveAIPlanRequest) SetApiRevision(v string) *RemoveAIPlanRequest {
	s.ApiRevision = &v
	return s
}

func (s *RemoveAIPlanRequest) SetPlanId(v string) *RemoveAIPlanRequest {
	s.PlanId = &v
	return s
}

type RemoveAIPlanResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s RemoveAIPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveAIPlanResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveAIPlanResponseBody) SetRequestId(v string) *RemoveAIPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveAIPlanResponseBody) SetSuccess(v bool) *RemoveAIPlanResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveAIPlanResponseBody) SetErrorMessage(v string) *RemoveAIPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RemoveAIPlanResponseBody) SetCode(v string) *RemoveAIPlanResponseBody {
	s.Code = &v
	return s
}

type RemoveAIPlanResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveAIPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveAIPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveAIPlanResponse) GoString() string {
	return s.String()
}

func (s *RemoveAIPlanResponse) SetHeaders(v map[string]*string) *RemoveAIPlanResponse {
	s.Headers = v
	return s
}

func (s *RemoveAIPlanResponse) SetBody(v *RemoveAIPlanResponseBody) *RemoveAIPlanResponse {
	s.Body = v
	return s
}

type RemoveDevicePurifyPlanRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	PlanId      *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s RemoveDevicePurifyPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveDevicePurifyPlanRequest) GoString() string {
	return s.String()
}

func (s *RemoveDevicePurifyPlanRequest) SetApiProduct(v string) *RemoveDevicePurifyPlanRequest {
	s.ApiProduct = &v
	return s
}

func (s *RemoveDevicePurifyPlanRequest) SetApiRevision(v string) *RemoveDevicePurifyPlanRequest {
	s.ApiRevision = &v
	return s
}

func (s *RemoveDevicePurifyPlanRequest) SetPlanId(v string) *RemoveDevicePurifyPlanRequest {
	s.PlanId = &v
	return s
}

type RemoveDevicePurifyPlanResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveDevicePurifyPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveDevicePurifyPlanResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDevicePurifyPlanResponseBody) SetRequestId(v string) *RemoveDevicePurifyPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveDevicePurifyPlanResponseBody) SetErrorMessage(v string) *RemoveDevicePurifyPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RemoveDevicePurifyPlanResponseBody) SetCode(v string) *RemoveDevicePurifyPlanResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveDevicePurifyPlanResponseBody) SetSuccess(v bool) *RemoveDevicePurifyPlanResponseBody {
	s.Success = &v
	return s
}

type RemoveDevicePurifyPlanResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveDevicePurifyPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveDevicePurifyPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveDevicePurifyPlanResponse) GoString() string {
	return s.String()
}

func (s *RemoveDevicePurifyPlanResponse) SetHeaders(v map[string]*string) *RemoveDevicePurifyPlanResponse {
	s.Headers = v
	return s
}

func (s *RemoveDevicePurifyPlanResponse) SetBody(v *RemoveDevicePurifyPlanResponseBody) *RemoveDevicePurifyPlanResponse {
	s.Body = v
	return s
}

type RemoveFaceDeviceFromDeviceGroupRequest struct {
	ApiProduct    *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId   *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
}

func (s RemoveFaceDeviceFromDeviceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveFaceDeviceFromDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveFaceDeviceFromDeviceGroupRequest) SetApiProduct(v string) *RemoveFaceDeviceFromDeviceGroupRequest {
	s.ApiProduct = &v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupRequest) SetApiRevision(v string) *RemoveFaceDeviceFromDeviceGroupRequest {
	s.ApiRevision = &v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupRequest) SetIsolationId(v string) *RemoveFaceDeviceFromDeviceGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupRequest) SetIotInstanceId(v string) *RemoveFaceDeviceFromDeviceGroupRequest {
	s.IotInstanceId = &v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupRequest) SetProductKey(v string) *RemoveFaceDeviceFromDeviceGroupRequest {
	s.ProductKey = &v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupRequest) SetDeviceName(v string) *RemoveFaceDeviceFromDeviceGroupRequest {
	s.DeviceName = &v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupRequest) SetDeviceGroupId(v string) *RemoveFaceDeviceFromDeviceGroupRequest {
	s.DeviceGroupId = &v
	return s
}

type RemoveFaceDeviceFromDeviceGroupResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveFaceDeviceFromDeviceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveFaceDeviceFromDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveFaceDeviceFromDeviceGroupResponseBody) SetRequestId(v string) *RemoveFaceDeviceFromDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupResponseBody) SetErrorMessage(v string) *RemoveFaceDeviceFromDeviceGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupResponseBody) SetCode(v string) *RemoveFaceDeviceFromDeviceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupResponseBody) SetSuccess(v bool) *RemoveFaceDeviceFromDeviceGroupResponseBody {
	s.Success = &v
	return s
}

type RemoveFaceDeviceFromDeviceGroupResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveFaceDeviceFromDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RemoveFaceDeviceFromDeviceGroupResponse) SetBody(v *RemoveFaceDeviceFromDeviceGroupResponseBody) *RemoveFaceDeviceFromDeviceGroupResponse {
	s.Body = v
	return s
}

type RemoveFaceUserFromUserGroupRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s RemoveFaceUserFromUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveFaceUserFromUserGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveFaceUserFromUserGroupRequest) SetApiProduct(v string) *RemoveFaceUserFromUserGroupRequest {
	s.ApiProduct = &v
	return s
}

func (s *RemoveFaceUserFromUserGroupRequest) SetApiRevision(v string) *RemoveFaceUserFromUserGroupRequest {
	s.ApiRevision = &v
	return s
}

func (s *RemoveFaceUserFromUserGroupRequest) SetIsolationId(v string) *RemoveFaceUserFromUserGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *RemoveFaceUserFromUserGroupRequest) SetUserId(v string) *RemoveFaceUserFromUserGroupRequest {
	s.UserId = &v
	return s
}

func (s *RemoveFaceUserFromUserGroupRequest) SetUserGroupId(v string) *RemoveFaceUserFromUserGroupRequest {
	s.UserGroupId = &v
	return s
}

type RemoveFaceUserFromUserGroupResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveFaceUserFromUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveFaceUserFromUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveFaceUserFromUserGroupResponseBody) SetRequestId(v string) *RemoveFaceUserFromUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveFaceUserFromUserGroupResponseBody) SetErrorMessage(v string) *RemoveFaceUserFromUserGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RemoveFaceUserFromUserGroupResponseBody) SetCode(v string) *RemoveFaceUserFromUserGroupResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveFaceUserFromUserGroupResponseBody) SetSuccess(v bool) *RemoveFaceUserFromUserGroupResponseBody {
	s.Success = &v
	return s
}

type RemoveFaceUserFromUserGroupResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveFaceUserFromUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RemoveFaceUserFromUserGroupResponse) SetBody(v *RemoveFaceUserFromUserGroupResponseBody) *RemoveFaceUserFromUserGroupResponse {
	s.Body = v
	return s
}

type SetDevicePictureLifeCycleRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId       *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	Day         *int32  `json:"Day,omitempty" xml:"Day,omitempty"`
}

func (s SetDevicePictureLifeCycleRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDevicePictureLifeCycleRequest) GoString() string {
	return s.String()
}

func (s *SetDevicePictureLifeCycleRequest) SetApiProduct(v string) *SetDevicePictureLifeCycleRequest {
	s.ApiProduct = &v
	return s
}

func (s *SetDevicePictureLifeCycleRequest) SetApiRevision(v string) *SetDevicePictureLifeCycleRequest {
	s.ApiRevision = &v
	return s
}

func (s *SetDevicePictureLifeCycleRequest) SetIotId(v string) *SetDevicePictureLifeCycleRequest {
	s.IotId = &v
	return s
}

func (s *SetDevicePictureLifeCycleRequest) SetDay(v int32) *SetDevicePictureLifeCycleRequest {
	s.Day = &v
	return s
}

type SetDevicePictureLifeCycleResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetDevicePictureLifeCycleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDevicePictureLifeCycleResponseBody) GoString() string {
	return s.String()
}

func (s *SetDevicePictureLifeCycleResponseBody) SetRequestId(v string) *SetDevicePictureLifeCycleResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDevicePictureLifeCycleResponseBody) SetErrorMessage(v string) *SetDevicePictureLifeCycleResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SetDevicePictureLifeCycleResponseBody) SetCode(v string) *SetDevicePictureLifeCycleResponseBody {
	s.Code = &v
	return s
}

func (s *SetDevicePictureLifeCycleResponseBody) SetSuccess(v bool) *SetDevicePictureLifeCycleResponseBody {
	s.Success = &v
	return s
}

type SetDevicePictureLifeCycleResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDevicePictureLifeCycleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SetDevicePictureLifeCycleResponse) SetBody(v *SetDevicePictureLifeCycleResponseBody) *SetDevicePictureLifeCycleResponse {
	s.Body = v
	return s
}

type SetDeviceRecordLifeCycleRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId       *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	Day         *int32  `json:"Day,omitempty" xml:"Day,omitempty"`
}

func (s SetDeviceRecordLifeCycleRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDeviceRecordLifeCycleRequest) GoString() string {
	return s.String()
}

func (s *SetDeviceRecordLifeCycleRequest) SetApiProduct(v string) *SetDeviceRecordLifeCycleRequest {
	s.ApiProduct = &v
	return s
}

func (s *SetDeviceRecordLifeCycleRequest) SetApiRevision(v string) *SetDeviceRecordLifeCycleRequest {
	s.ApiRevision = &v
	return s
}

func (s *SetDeviceRecordLifeCycleRequest) SetIotId(v string) *SetDeviceRecordLifeCycleRequest {
	s.IotId = &v
	return s
}

func (s *SetDeviceRecordLifeCycleRequest) SetDay(v int32) *SetDeviceRecordLifeCycleRequest {
	s.Day = &v
	return s
}

type SetDeviceRecordLifeCycleResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetDeviceRecordLifeCycleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDeviceRecordLifeCycleResponseBody) GoString() string {
	return s.String()
}

func (s *SetDeviceRecordLifeCycleResponseBody) SetRequestId(v string) *SetDeviceRecordLifeCycleResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDeviceRecordLifeCycleResponseBody) SetErrorMessage(v string) *SetDeviceRecordLifeCycleResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SetDeviceRecordLifeCycleResponseBody) SetCode(v string) *SetDeviceRecordLifeCycleResponseBody {
	s.Code = &v
	return s
}

func (s *SetDeviceRecordLifeCycleResponseBody) SetSuccess(v bool) *SetDeviceRecordLifeCycleResponseBody {
	s.Success = &v
	return s
}

type SetDeviceRecordLifeCycleResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDeviceRecordLifeCycleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SetDeviceRecordLifeCycleResponse) SetBody(v *SetDeviceRecordLifeCycleResponseBody) *SetDeviceRecordLifeCycleResponse {
	s.Body = v
	return s
}

type StopLiveStreamingRequest struct {
	ApiProduct    *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	StreamType    *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
}

func (s StopLiveStreamingRequest) String() string {
	return tea.Prettify(s)
}

func (s StopLiveStreamingRequest) GoString() string {
	return s.String()
}

func (s *StopLiveStreamingRequest) SetApiProduct(v string) *StopLiveStreamingRequest {
	s.ApiProduct = &v
	return s
}

func (s *StopLiveStreamingRequest) SetApiRevision(v string) *StopLiveStreamingRequest {
	s.ApiRevision = &v
	return s
}

func (s *StopLiveStreamingRequest) SetIotId(v string) *StopLiveStreamingRequest {
	s.IotId = &v
	return s
}

func (s *StopLiveStreamingRequest) SetStreamType(v int32) *StopLiveStreamingRequest {
	s.StreamType = &v
	return s
}

func (s *StopLiveStreamingRequest) SetIotInstanceId(v string) *StopLiveStreamingRequest {
	s.IotInstanceId = &v
	return s
}

type StopLiveStreamingResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopLiveStreamingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopLiveStreamingResponseBody) GoString() string {
	return s.String()
}

func (s *StopLiveStreamingResponseBody) SetRequestId(v string) *StopLiveStreamingResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopLiveStreamingResponseBody) SetErrorMessage(v string) *StopLiveStreamingResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StopLiveStreamingResponseBody) SetCode(v string) *StopLiveStreamingResponseBody {
	s.Code = &v
	return s
}

func (s *StopLiveStreamingResponseBody) SetSuccess(v bool) *StopLiveStreamingResponseBody {
	s.Success = &v
	return s
}

type StopLiveStreamingResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopLiveStreamingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *StopLiveStreamingResponse) SetBody(v *StopLiveStreamingResponseBody) *StopLiveStreamingResponse {
	s.Body = v
	return s
}

type StopTriggeredRecordRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId       *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	RecordId    *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s StopTriggeredRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s StopTriggeredRecordRequest) GoString() string {
	return s.String()
}

func (s *StopTriggeredRecordRequest) SetApiProduct(v string) *StopTriggeredRecordRequest {
	s.ApiProduct = &v
	return s
}

func (s *StopTriggeredRecordRequest) SetApiRevision(v string) *StopTriggeredRecordRequest {
	s.ApiRevision = &v
	return s
}

func (s *StopTriggeredRecordRequest) SetIotId(v string) *StopTriggeredRecordRequest {
	s.IotId = &v
	return s
}

func (s *StopTriggeredRecordRequest) SetRecordId(v string) *StopTriggeredRecordRequest {
	s.RecordId = &v
	return s
}

type StopTriggeredRecordResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopTriggeredRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopTriggeredRecordResponseBody) GoString() string {
	return s.String()
}

func (s *StopTriggeredRecordResponseBody) SetRequestId(v string) *StopTriggeredRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopTriggeredRecordResponseBody) SetErrorMessage(v string) *StopTriggeredRecordResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StopTriggeredRecordResponseBody) SetCode(v string) *StopTriggeredRecordResponseBody {
	s.Code = &v
	return s
}

func (s *StopTriggeredRecordResponseBody) SetSuccess(v bool) *StopTriggeredRecordResponseBody {
	s.Success = &v
	return s
}

type StopTriggeredRecordResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopTriggeredRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *StopTriggeredRecordResponse) SetBody(v *StopTriggeredRecordResponseBody) *StopTriggeredRecordResponse {
	s.Body = v
	return s
}

type TriggerCapturePictureRequest struct {
	ApiProduct    *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
}

func (s TriggerCapturePictureRequest) String() string {
	return tea.Prettify(s)
}

func (s TriggerCapturePictureRequest) GoString() string {
	return s.String()
}

func (s *TriggerCapturePictureRequest) SetApiProduct(v string) *TriggerCapturePictureRequest {
	s.ApiProduct = &v
	return s
}

func (s *TriggerCapturePictureRequest) SetApiRevision(v string) *TriggerCapturePictureRequest {
	s.ApiRevision = &v
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

type TriggerCapturePictureResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TriggerCapturePictureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TriggerCapturePictureResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerCapturePictureResponseBody) SetRequestId(v string) *TriggerCapturePictureResponseBody {
	s.RequestId = &v
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

func (s *TriggerCapturePictureResponseBody) SetCode(v string) *TriggerCapturePictureResponseBody {
	s.Code = &v
	return s
}

func (s *TriggerCapturePictureResponseBody) SetSuccess(v bool) *TriggerCapturePictureResponseBody {
	s.Success = &v
	return s
}

type TriggerCapturePictureResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TriggerCapturePictureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TriggerCapturePictureResponse) SetBody(v *TriggerCapturePictureResponseBody) *TriggerCapturePictureResponse {
	s.Body = v
	return s
}

type TriggerRecordRequest struct {
	ApiProduct        *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision       *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IotId             *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	StreamType        *int32  `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	PreRecordDuration *int32  `json:"PreRecordDuration,omitempty" xml:"PreRecordDuration,omitempty"`
	RecordDuration    *int32  `json:"RecordDuration,omitempty" xml:"RecordDuration,omitempty"`
	IotInstanceId     *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
}

func (s TriggerRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s TriggerRecordRequest) GoString() string {
	return s.String()
}

func (s *TriggerRecordRequest) SetApiProduct(v string) *TriggerRecordRequest {
	s.ApiProduct = &v
	return s
}

func (s *TriggerRecordRequest) SetApiRevision(v string) *TriggerRecordRequest {
	s.ApiRevision = &v
	return s
}

func (s *TriggerRecordRequest) SetIotId(v string) *TriggerRecordRequest {
	s.IotId = &v
	return s
}

func (s *TriggerRecordRequest) SetStreamType(v int32) *TriggerRecordRequest {
	s.StreamType = &v
	return s
}

func (s *TriggerRecordRequest) SetPreRecordDuration(v int32) *TriggerRecordRequest {
	s.PreRecordDuration = &v
	return s
}

func (s *TriggerRecordRequest) SetRecordDuration(v int32) *TriggerRecordRequest {
	s.RecordDuration = &v
	return s
}

func (s *TriggerRecordRequest) SetIotInstanceId(v string) *TriggerRecordRequest {
	s.IotInstanceId = &v
	return s
}

type TriggerRecordResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TriggerRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TriggerRecordResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerRecordResponseBody) SetRequestId(v string) *TriggerRecordResponseBody {
	s.RequestId = &v
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

func (s *TriggerRecordResponseBody) SetCode(v string) *TriggerRecordResponseBody {
	s.Code = &v
	return s
}

func (s *TriggerRecordResponseBody) SetSuccess(v bool) *TriggerRecordResponseBody {
	s.Success = &v
	return s
}

type TriggerRecordResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TriggerRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TriggerRecordResponse) SetBody(v *TriggerRecordResponseBody) *TriggerRecordResponse {
	s.Body = v
	return s
}

type UnbindAIPlanWithDevicesRequest struct {
	ApiProduct  *string   `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string   `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	PlanId      *string   `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	IotIdList   []*string `json:"IotIdList,omitempty" xml:"IotIdList,omitempty" type:"Repeated"`
}

func (s UnbindAIPlanWithDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindAIPlanWithDevicesRequest) GoString() string {
	return s.String()
}

func (s *UnbindAIPlanWithDevicesRequest) SetApiProduct(v string) *UnbindAIPlanWithDevicesRequest {
	s.ApiProduct = &v
	return s
}

func (s *UnbindAIPlanWithDevicesRequest) SetApiRevision(v string) *UnbindAIPlanWithDevicesRequest {
	s.ApiRevision = &v
	return s
}

func (s *UnbindAIPlanWithDevicesRequest) SetPlanId(v string) *UnbindAIPlanWithDevicesRequest {
	s.PlanId = &v
	return s
}

func (s *UnbindAIPlanWithDevicesRequest) SetIotIdList(v []*string) *UnbindAIPlanWithDevicesRequest {
	s.IotIdList = v
	return s
}

type UnbindAIPlanWithDevicesResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UnbindAIPlanWithDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindAIPlanWithDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindAIPlanWithDevicesResponseBody) SetRequestId(v string) *UnbindAIPlanWithDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindAIPlanWithDevicesResponseBody) SetSuccess(v bool) *UnbindAIPlanWithDevicesResponseBody {
	s.Success = &v
	return s
}

func (s *UnbindAIPlanWithDevicesResponseBody) SetErrorMessage(v string) *UnbindAIPlanWithDevicesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UnbindAIPlanWithDevicesResponseBody) SetCode(v string) *UnbindAIPlanWithDevicesResponseBody {
	s.Code = &v
	return s
}

type UnbindAIPlanWithDevicesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnbindAIPlanWithDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindAIPlanWithDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindAIPlanWithDevicesResponse) GoString() string {
	return s.String()
}

func (s *UnbindAIPlanWithDevicesResponse) SetHeaders(v map[string]*string) *UnbindAIPlanWithDevicesResponse {
	s.Headers = v
	return s
}

func (s *UnbindAIPlanWithDevicesResponse) SetBody(v *UnbindAIPlanWithDevicesResponseBody) *UnbindAIPlanWithDevicesResponse {
	s.Body = v
	return s
}

type UnbindPictureSearchAppWithDevicesRequest struct {
	ApiProduct    *string   `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision   *string   `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	AppInstanceId *string   `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	DeviceIotIds  []*string `json:"DeviceIotIds,omitempty" xml:"DeviceIotIds,omitempty" type:"Repeated"`
}

func (s UnbindPictureSearchAppWithDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindPictureSearchAppWithDevicesRequest) GoString() string {
	return s.String()
}

func (s *UnbindPictureSearchAppWithDevicesRequest) SetApiProduct(v string) *UnbindPictureSearchAppWithDevicesRequest {
	s.ApiProduct = &v
	return s
}

func (s *UnbindPictureSearchAppWithDevicesRequest) SetApiRevision(v string) *UnbindPictureSearchAppWithDevicesRequest {
	s.ApiRevision = &v
	return s
}

func (s *UnbindPictureSearchAppWithDevicesRequest) SetAppInstanceId(v string) *UnbindPictureSearchAppWithDevicesRequest {
	s.AppInstanceId = &v
	return s
}

func (s *UnbindPictureSearchAppWithDevicesRequest) SetDeviceIotIds(v []*string) *UnbindPictureSearchAppWithDevicesRequest {
	s.DeviceIotIds = v
	return s
}

type UnbindPictureSearchAppWithDevicesResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnbindPictureSearchAppWithDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindPictureSearchAppWithDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindPictureSearchAppWithDevicesResponseBody) SetRequestId(v string) *UnbindPictureSearchAppWithDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindPictureSearchAppWithDevicesResponseBody) SetErrorMessage(v string) *UnbindPictureSearchAppWithDevicesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UnbindPictureSearchAppWithDevicesResponseBody) SetCode(v string) *UnbindPictureSearchAppWithDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindPictureSearchAppWithDevicesResponseBody) SetSuccess(v bool) *UnbindPictureSearchAppWithDevicesResponseBody {
	s.Success = &v
	return s
}

type UnbindPictureSearchAppWithDevicesResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnbindPictureSearchAppWithDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UnbindPictureSearchAppWithDevicesResponse) SetBody(v *UnbindPictureSearchAppWithDevicesResponseBody) *UnbindPictureSearchAppWithDevicesResponse {
	s.Body = v
	return s
}

type UpdateAIAppRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Level       *int32  `json:"Level,omitempty" xml:"Level,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s UpdateAIAppRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAIAppRequest) GoString() string {
	return s.String()
}

func (s *UpdateAIAppRequest) SetApiProduct(v string) *UpdateAIAppRequest {
	s.ApiProduct = &v
	return s
}

func (s *UpdateAIAppRequest) SetApiRevision(v string) *UpdateAIAppRequest {
	s.ApiRevision = &v
	return s
}

func (s *UpdateAIAppRequest) SetAppId(v string) *UpdateAIAppRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAIAppRequest) SetLevel(v int32) *UpdateAIAppRequest {
	s.Level = &v
	return s
}

func (s *UpdateAIAppRequest) SetName(v string) *UpdateAIAppRequest {
	s.Name = &v
	return s
}

func (s *UpdateAIAppRequest) SetDescription(v string) *UpdateAIAppRequest {
	s.Description = &v
	return s
}

type UpdateAIAppResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UpdateAIAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAIAppResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAIAppResponseBody) SetRequestId(v string) *UpdateAIAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAIAppResponseBody) SetSuccess(v bool) *UpdateAIAppResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAIAppResponseBody) SetErrorMessage(v string) *UpdateAIAppResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateAIAppResponseBody) SetCode(v string) *UpdateAIAppResponseBody {
	s.Code = &v
	return s
}

type UpdateAIAppResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAIAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAIAppResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAIAppResponse) GoString() string {
	return s.String()
}

func (s *UpdateAIAppResponse) SetHeaders(v map[string]*string) *UpdateAIAppResponse {
	s.Headers = v
	return s
}

func (s *UpdateAIAppResponse) SetBody(v *UpdateAIAppResponseBody) *UpdateAIAppResponse {
	s.Body = v
	return s
}

type UpdateAIPlanRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	PlanId      *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s UpdateAIPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAIPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateAIPlanRequest) SetApiProduct(v string) *UpdateAIPlanRequest {
	s.ApiProduct = &v
	return s
}

func (s *UpdateAIPlanRequest) SetApiRevision(v string) *UpdateAIPlanRequest {
	s.ApiRevision = &v
	return s
}

func (s *UpdateAIPlanRequest) SetPlanId(v string) *UpdateAIPlanRequest {
	s.PlanId = &v
	return s
}

func (s *UpdateAIPlanRequest) SetDescription(v string) *UpdateAIPlanRequest {
	s.Description = &v
	return s
}

type UpdateAIPlanResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UpdateAIPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAIPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAIPlanResponseBody) SetRequestId(v string) *UpdateAIPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAIPlanResponseBody) SetSuccess(v bool) *UpdateAIPlanResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAIPlanResponseBody) SetErrorMessage(v string) *UpdateAIPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateAIPlanResponseBody) SetCode(v string) *UpdateAIPlanResponseBody {
	s.Code = &v
	return s
}

type UpdateAIPlanResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAIPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAIPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAIPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateAIPlanResponse) SetHeaders(v map[string]*string) *UpdateAIPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateAIPlanResponse) SetBody(v *UpdateAIPlanResponseBody) *UpdateAIPlanResponse {
	s.Body = v
	return s
}

type UpdateDevicePurifyPlanRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	PlanId      *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	StartTime   *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s UpdateDevicePurifyPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDevicePurifyPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateDevicePurifyPlanRequest) SetApiProduct(v string) *UpdateDevicePurifyPlanRequest {
	s.ApiProduct = &v
	return s
}

func (s *UpdateDevicePurifyPlanRequest) SetApiRevision(v string) *UpdateDevicePurifyPlanRequest {
	s.ApiRevision = &v
	return s
}

func (s *UpdateDevicePurifyPlanRequest) SetPlanId(v string) *UpdateDevicePurifyPlanRequest {
	s.PlanId = &v
	return s
}

func (s *UpdateDevicePurifyPlanRequest) SetStartTime(v int32) *UpdateDevicePurifyPlanRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateDevicePurifyPlanRequest) SetEndTime(v int32) *UpdateDevicePurifyPlanRequest {
	s.EndTime = &v
	return s
}

type UpdateDevicePurifyPlanResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDevicePurifyPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDevicePurifyPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDevicePurifyPlanResponseBody) SetRequestId(v string) *UpdateDevicePurifyPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDevicePurifyPlanResponseBody) SetErrorMessage(v string) *UpdateDevicePurifyPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateDevicePurifyPlanResponseBody) SetCode(v string) *UpdateDevicePurifyPlanResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDevicePurifyPlanResponseBody) SetSuccess(v bool) *UpdateDevicePurifyPlanResponseBody {
	s.Success = &v
	return s
}

type UpdateDevicePurifyPlanResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDevicePurifyPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDevicePurifyPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDevicePurifyPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateDevicePurifyPlanResponse) SetHeaders(v map[string]*string) *UpdateDevicePurifyPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateDevicePurifyPlanResponse) SetBody(v *UpdateDevicePurifyPlanResponseBody) *UpdateDevicePurifyPlanResponse {
	s.Body = v
	return s
}

type UpdateEventRecordPlanRequest struct {
	ApiProduct        *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision       *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	PlanId            *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	EventTypes        *string `json:"EventTypes,omitempty" xml:"EventTypes,omitempty"`
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

func (s *UpdateEventRecordPlanRequest) SetApiProduct(v string) *UpdateEventRecordPlanRequest {
	s.ApiProduct = &v
	return s
}

func (s *UpdateEventRecordPlanRequest) SetApiRevision(v string) *UpdateEventRecordPlanRequest {
	s.ApiRevision = &v
	return s
}

func (s *UpdateEventRecordPlanRequest) SetPlanId(v string) *UpdateEventRecordPlanRequest {
	s.PlanId = &v
	return s
}

func (s *UpdateEventRecordPlanRequest) SetName(v string) *UpdateEventRecordPlanRequest {
	s.Name = &v
	return s
}

func (s *UpdateEventRecordPlanRequest) SetEventTypes(v string) *UpdateEventRecordPlanRequest {
	s.EventTypes = &v
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
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateEventRecordPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventRecordPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEventRecordPlanResponseBody) SetRequestId(v string) *UpdateEventRecordPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEventRecordPlanResponseBody) SetErrorMessage(v string) *UpdateEventRecordPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateEventRecordPlanResponseBody) SetCode(v string) *UpdateEventRecordPlanResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEventRecordPlanResponseBody) SetSuccess(v bool) *UpdateEventRecordPlanResponseBody {
	s.Success = &v
	return s
}

type UpdateEventRecordPlanResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateEventRecordPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateEventRecordPlanResponse) SetBody(v *UpdateEventRecordPlanResponseBody) *UpdateEventRecordPlanResponse {
	s.Body = v
	return s
}

type UpdateFaceUserRequest struct {
	ApiProduct   *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision  *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId  *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Params       *string `json:"Params,omitempty" xml:"Params,omitempty"`
	FacePicUrl   *string `json:"FacePicUrl,omitempty" xml:"FacePicUrl,omitempty"`
	CustomUserId *string `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
}

func (s UpdateFaceUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateFaceUserRequest) SetApiProduct(v string) *UpdateFaceUserRequest {
	s.ApiProduct = &v
	return s
}

func (s *UpdateFaceUserRequest) SetApiRevision(v string) *UpdateFaceUserRequest {
	s.ApiRevision = &v
	return s
}

func (s *UpdateFaceUserRequest) SetIsolationId(v string) *UpdateFaceUserRequest {
	s.IsolationId = &v
	return s
}

func (s *UpdateFaceUserRequest) SetUserId(v string) *UpdateFaceUserRequest {
	s.UserId = &v
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

func (s *UpdateFaceUserRequest) SetFacePicUrl(v string) *UpdateFaceUserRequest {
	s.FacePicUrl = &v
	return s
}

func (s *UpdateFaceUserRequest) SetCustomUserId(v string) *UpdateFaceUserRequest {
	s.CustomUserId = &v
	return s
}

type UpdateFaceUserResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateFaceUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFaceUserResponseBody) SetRequestId(v string) *UpdateFaceUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFaceUserResponseBody) SetErrorMessage(v string) *UpdateFaceUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateFaceUserResponseBody) SetCode(v string) *UpdateFaceUserResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFaceUserResponseBody) SetSuccess(v bool) *UpdateFaceUserResponseBody {
	s.Success = &v
	return s
}

type UpdateFaceUserResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFaceUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateFaceUserResponse) SetBody(v *UpdateFaceUserResponseBody) *UpdateFaceUserResponse {
	s.Body = v
	return s
}

type UpdateFaceUserGroupAndDeviceGroupRelationRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	ControlId   *string `json:"ControlId,omitempty" xml:"ControlId,omitempty"`
	Relation    *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
}

func (s UpdateFaceUserGroupAndDeviceGroupRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceUserGroupAndDeviceGroupRelationRequest) GoString() string {
	return s.String()
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationRequest) SetApiProduct(v string) *UpdateFaceUserGroupAndDeviceGroupRelationRequest {
	s.ApiProduct = &v
	return s
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationRequest) SetApiRevision(v string) *UpdateFaceUserGroupAndDeviceGroupRelationRequest {
	s.ApiRevision = &v
	return s
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationRequest) SetIsolationId(v string) *UpdateFaceUserGroupAndDeviceGroupRelationRequest {
	s.IsolationId = &v
	return s
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationRequest) SetControlId(v string) *UpdateFaceUserGroupAndDeviceGroupRelationRequest {
	s.ControlId = &v
	return s
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationRequest) SetRelation(v string) *UpdateFaceUserGroupAndDeviceGroupRelationRequest {
	s.Relation = &v
	return s
}

type UpdateFaceUserGroupAndDeviceGroupRelationResponseBody struct {
	RequestId    *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *UpdateFaceUserGroupAndDeviceGroupRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateFaceUserGroupAndDeviceGroupRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceUserGroupAndDeviceGroupRelationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationResponseBody) SetRequestId(v string) *UpdateFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.RequestId = &v
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

func (s *UpdateFaceUserGroupAndDeviceGroupRelationResponseBody) SetCode(v string) *UpdateFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationResponseBody) SetSuccess(v bool) *UpdateFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.Success = &v
	return s
}

type UpdateFaceUserGroupAndDeviceGroupRelationResponseBodyData struct {
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ControlId    *string `json:"ControlId,omitempty" xml:"ControlId,omitempty"`
}

func (s UpdateFaceUserGroupAndDeviceGroupRelationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceUserGroupAndDeviceGroupRelationResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationResponseBodyData) SetModifiedTime(v string) *UpdateFaceUserGroupAndDeviceGroupRelationResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationResponseBodyData) SetControlId(v string) *UpdateFaceUserGroupAndDeviceGroupRelationResponseBodyData {
	s.ControlId = &v
	return s
}

type UpdateFaceUserGroupAndDeviceGroupRelationResponse struct {
	Headers map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFaceUserGroupAndDeviceGroupRelationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateFaceUserGroupAndDeviceGroupRelationResponse) SetBody(v *UpdateFaceUserGroupAndDeviceGroupRelationResponseBody) *UpdateFaceUserGroupAndDeviceGroupRelationResponse {
	s.Body = v
	return s
}

type UpdateModelRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	ModelId     *int64  `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Hardware    *string `json:"Hardware,omitempty" xml:"Hardware,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s UpdateModelRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateModelRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelRequest) SetApiProduct(v string) *UpdateModelRequest {
	s.ApiProduct = &v
	return s
}

func (s *UpdateModelRequest) SetApiRevision(v string) *UpdateModelRequest {
	s.ApiRevision = &v
	return s
}

func (s *UpdateModelRequest) SetModelId(v int64) *UpdateModelRequest {
	s.ModelId = &v
	return s
}

func (s *UpdateModelRequest) SetName(v string) *UpdateModelRequest {
	s.Name = &v
	return s
}

func (s *UpdateModelRequest) SetHardware(v string) *UpdateModelRequest {
	s.Hardware = &v
	return s
}

func (s *UpdateModelRequest) SetDescription(v string) *UpdateModelRequest {
	s.Description = &v
	return s
}

type UpdateModelResponseBody struct {
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateModelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateModelResponseBody) SetRequestId(v string) *UpdateModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateModelResponseBody) SetData(v map[string]interface{}) *UpdateModelResponseBody {
	s.Data = v
	return s
}

func (s *UpdateModelResponseBody) SetErrorMessage(v string) *UpdateModelResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateModelResponseBody) SetCode(v string) *UpdateModelResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateModelResponseBody) SetSuccess(v bool) *UpdateModelResponseBody {
	s.Success = &v
	return s
}

type UpdateModelResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateModelResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateModelResponse) GoString() string {
	return s.String()
}

func (s *UpdateModelResponse) SetHeaders(v map[string]*string) *UpdateModelResponse {
	s.Headers = v
	return s
}

func (s *UpdateModelResponse) SetBody(v *UpdateModelResponseBody) *UpdateModelResponse {
	s.Body = v
	return s
}

type UpdateRecordPlanRequest struct {
	ApiProduct  *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	PlanId      *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateId  *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateRecordPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecordPlanRequest) SetApiProduct(v string) *UpdateRecordPlanRequest {
	s.ApiProduct = &v
	return s
}

func (s *UpdateRecordPlanRequest) SetApiRevision(v string) *UpdateRecordPlanRequest {
	s.ApiRevision = &v
	return s
}

func (s *UpdateRecordPlanRequest) SetPlanId(v string) *UpdateRecordPlanRequest {
	s.PlanId = &v
	return s
}

func (s *UpdateRecordPlanRequest) SetName(v string) *UpdateRecordPlanRequest {
	s.Name = &v
	return s
}

func (s *UpdateRecordPlanRequest) SetTemplateId(v string) *UpdateRecordPlanRequest {
	s.TemplateId = &v
	return s
}

type UpdateRecordPlanResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRecordPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRecordPlanResponseBody) SetRequestId(v string) *UpdateRecordPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRecordPlanResponseBody) SetErrorMessage(v string) *UpdateRecordPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateRecordPlanResponseBody) SetCode(v string) *UpdateRecordPlanResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRecordPlanResponseBody) SetSuccess(v bool) *UpdateRecordPlanResponseBody {
	s.Success = &v
	return s
}

type UpdateRecordPlanResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateRecordPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateRecordPlanResponse) SetBody(v *UpdateRecordPlanResponseBody) *UpdateRecordPlanResponse {
	s.Body = v
	return s
}

type UpdateTimeTemplateRequest struct {
	ApiProduct   *string                                  `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision  *string                                  `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	TemplateId   *string                                  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Name         *string                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	AllDay       *int32                                   `json:"AllDay,omitempty" xml:"AllDay,omitempty"`
	TimeSections []*UpdateTimeTemplateRequestTimeSections `json:"TimeSections,omitempty" xml:"TimeSections,omitempty" type:"Repeated"`
}

func (s UpdateTimeTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTimeTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateTimeTemplateRequest) SetApiProduct(v string) *UpdateTimeTemplateRequest {
	s.ApiProduct = &v
	return s
}

func (s *UpdateTimeTemplateRequest) SetApiRevision(v string) *UpdateTimeTemplateRequest {
	s.ApiRevision = &v
	return s
}

func (s *UpdateTimeTemplateRequest) SetTemplateId(v string) *UpdateTimeTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateTimeTemplateRequest) SetName(v string) *UpdateTimeTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdateTimeTemplateRequest) SetAllDay(v int32) *UpdateTimeTemplateRequest {
	s.AllDay = &v
	return s
}

func (s *UpdateTimeTemplateRequest) SetTimeSections(v []*UpdateTimeTemplateRequestTimeSections) *UpdateTimeTemplateRequest {
	s.TimeSections = v
	return s
}

type UpdateTimeTemplateRequestTimeSections struct {
	DayOfWeek *int32 `json:"DayOfWeek,omitempty" xml:"DayOfWeek,omitempty"`
	Begin     *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	End       *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s UpdateTimeTemplateRequestTimeSections) String() string {
	return tea.Prettify(s)
}

func (s UpdateTimeTemplateRequestTimeSections) GoString() string {
	return s.String()
}

func (s *UpdateTimeTemplateRequestTimeSections) SetDayOfWeek(v int32) *UpdateTimeTemplateRequestTimeSections {
	s.DayOfWeek = &v
	return s
}

func (s *UpdateTimeTemplateRequestTimeSections) SetBegin(v int32) *UpdateTimeTemplateRequestTimeSections {
	s.Begin = &v
	return s
}

func (s *UpdateTimeTemplateRequestTimeSections) SetEnd(v int32) *UpdateTimeTemplateRequestTimeSections {
	s.End = &v
	return s
}

type UpdateTimeTemplateResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTimeTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTimeTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTimeTemplateResponseBody) SetRequestId(v string) *UpdateTimeTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTimeTemplateResponseBody) SetErrorMessage(v string) *UpdateTimeTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateTimeTemplateResponseBody) SetCode(v string) *UpdateTimeTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTimeTemplateResponseBody) SetSuccess(v bool) *UpdateTimeTemplateResponseBody {
	s.Success = &v
	return s
}

type UpdateTimeTemplateResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTimeTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddEventRecordPlanDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddEventRecordPlanDevice"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddFaceDeviceGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddFaceDeviceGroup"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddFaceDeviceToDeviceGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddFaceDeviceToDeviceGroup"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddFaceUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddFaceUser"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddFaceUserGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddFaceUserGroup"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddFaceUserGroupAndDeviceGroupRelationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddFaceUserGroupAndDeviceGroupRelation"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddFaceUserPictureResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddFaceUserPicture"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddFaceUserToUserGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddFaceUserToUserGroup"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddRecordPlanDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddRecordPlanDevice"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) BindAIPlanWithDevicesWithOptions(request *BindAIPlanWithDevicesRequest, runtime *util.RuntimeOptions) (_result *BindAIPlanWithDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BindAIPlanWithDevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BindAIPlanWithDevices"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindAIPlanWithDevices(request *BindAIPlanWithDevicesRequest) (_result *BindAIPlanWithDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindAIPlanWithDevicesResponse{}
	_body, _err := client.BindAIPlanWithDevicesWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BindPictureSearchAppWithDevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BindPictureSearchAppWithDevices"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckFaceUserDoExistOnDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckFaceUserDoExistOnDevice"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ClearFaceDeviceDBResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ClearFaceDeviceDB"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ConfigAIActionWithOptions(request *ConfigAIActionRequest, runtime *util.RuntimeOptions) (_result *ConfigAIActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigAIActionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigAIAction"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigAIAction(request *ConfigAIActionRequest) (_result *ConfigAIActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigAIActionResponse{}
	_body, _err := client.ConfigAIActionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAIAppWithOptions(request *CreateAIAppRequest, runtime *util.RuntimeOptions) (_result *CreateAIAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAIAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAIApp"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAIApp(request *CreateAIAppRequest) (_result *CreateAIAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAIAppResponse{}
	_body, _err := client.CreateAIAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAIPlanWithOptions(request *CreateAIPlanRequest, runtime *util.RuntimeOptions) (_result *CreateAIPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAIPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAIPlan"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAIPlan(request *CreateAIPlanRequest) (_result *CreateAIPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAIPlanResponse{}
	_body, _err := client.CreateAIPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAlgorithmWithOptions(request *CreateAlgorithmRequest, runtime *util.RuntimeOptions) (_result *CreateAlgorithmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAlgorithmResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAlgorithm"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAlgorithm(request *CreateAlgorithmRequest) (_result *CreateAlgorithmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAlgorithmResponse{}
	_body, _err := client.CreateAlgorithmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDevicePurifyJobWithOptions(request *CreateDevicePurifyJobRequest, runtime *util.RuntimeOptions) (_result *CreateDevicePurifyJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDevicePurifyJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDevicePurifyJob"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDevicePurifyJob(request *CreateDevicePurifyJobRequest) (_result *CreateDevicePurifyJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDevicePurifyJobResponse{}
	_body, _err := client.CreateDevicePurifyJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDevicePurifyJobByPlanIdWithOptions(request *CreateDevicePurifyJobByPlanIdRequest, runtime *util.RuntimeOptions) (_result *CreateDevicePurifyJobByPlanIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDevicePurifyJobByPlanIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDevicePurifyJobByPlanId"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDevicePurifyJobByPlanId(request *CreateDevicePurifyJobByPlanIdRequest) (_result *CreateDevicePurifyJobByPlanIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDevicePurifyJobByPlanIdResponse{}
	_body, _err := client.CreateDevicePurifyJobByPlanIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDevicePurifyPlanWithOptions(request *CreateDevicePurifyPlanRequest, runtime *util.RuntimeOptions) (_result *CreateDevicePurifyPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDevicePurifyPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDevicePurifyPlan"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDevicePurifyPlan(request *CreateDevicePurifyPlanRequest) (_result *CreateDevicePurifyPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDevicePurifyPlanResponse{}
	_body, _err := client.CreateDevicePurifyPlanWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateEventRecordPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateEventRecordPlan"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateModelWithOptions(request *CreateModelRequest, runtime *util.RuntimeOptions) (_result *CreateModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateModelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateModel"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateModel(request *CreateModelRequest) (_result *CreateModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateModelResponse{}
	_body, _err := client.CreateModelWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreatePictureSearchAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreatePictureSearchApp"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateRecordPlanWithOptions(request *CreateRecordPlanRequest, runtime *util.RuntimeOptions) (_result *CreateRecordPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateRecordPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateRecordPlan"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateTimeTemplateWithOptions(request *CreateTimeTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateTimeTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateTimeTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateTimeTemplate"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeleteAlgorithmWithOptions(request *DeleteAlgorithmRequest, runtime *util.RuntimeOptions) (_result *DeleteAlgorithmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAlgorithmResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAlgorithm"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAlgorithm(request *DeleteAlgorithmRequest) (_result *DeleteAlgorithmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAlgorithmResponse{}
	_body, _err := client.DeleteAlgorithmWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteEventRecordPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteEventRecordPlan"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteEventRecordPlanDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteEventRecordPlanDevice"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFaceDeviceGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFaceDeviceGroup"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFaceUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFaceUser"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFaceUserGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFaceUserGroup"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFaceUserGroupAndDeviceGroupRelationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFaceUserGroupAndDeviceGroupRelation"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFaceUserPictureResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFaceUserPicture"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeleteModelWithOptions(request *DeleteModelRequest, runtime *util.RuntimeOptions) (_result *DeleteModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteModelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteModel"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteModel(request *DeleteModelRequest) (_result *DeleteModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteModelResponse{}
	_body, _err := client.DeleteModelWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteRecordPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteRecordPlan"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteRecordPlanDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteRecordPlanDevice"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeleteTimeTemplateWithOptions(request *DeleteTimeTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteTimeTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteTimeTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteTimeTemplate"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeployModelBatchWithOptions(request *DeployModelBatchRequest, runtime *util.RuntimeOptions) (_result *DeployModelBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeployModelBatchResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeployModelBatch"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeployModelBatch(request *DeployModelBatchRequest) (_result *DeployModelBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeployModelBatchResponse{}
	_body, _err := client.DeployModelBatchWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectUserFaceByUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectUserFaceByUrl"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetAIActionWithOptions(request *GetAIActionRequest, runtime *util.RuntimeOptions) (_result *GetAIActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAIActionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAIAction"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAIAction(request *GetAIActionRequest) (_result *GetAIActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAIActionResponse{}
	_body, _err := client.GetAIActionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAIActionConfigWithOptions(request *GetAIActionConfigRequest, runtime *util.RuntimeOptions) (_result *GetAIActionConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAIActionConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAIActionConfig"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAIActionConfig(request *GetAIActionConfigRequest) (_result *GetAIActionConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAIActionConfigResponse{}
	_body, _err := client.GetAIActionConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAIAppWithOptions(request *GetAIAppRequest, runtime *util.RuntimeOptions) (_result *GetAIAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAIAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAIApp"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAIApp(request *GetAIAppRequest) (_result *GetAIAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAIAppResponse{}
	_body, _err := client.GetAIAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAIJobWithOptions(request *GetAIJobRequest, runtime *util.RuntimeOptions) (_result *GetAIJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAIJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAIJob"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAIJob(request *GetAIJobRequest) (_result *GetAIJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAIJobResponse{}
	_body, _err := client.GetAIJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAIPlanWithOptions(request *GetAIPlanRequest, runtime *util.RuntimeOptions) (_result *GetAIPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAIPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAIPlan"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAIPlan(request *GetAIPlanRequest) (_result *GetAIPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAIPlanResponse{}
	_body, _err := client.GetAIPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAlgorithmDetailByIdWithOptions(request *GetAlgorithmDetailByIdRequest, runtime *util.RuntimeOptions) (_result *GetAlgorithmDetailByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAlgorithmDetailByIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAlgorithmDetailById"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAlgorithmDetailById(request *GetAlgorithmDetailByIdRequest) (_result *GetAlgorithmDetailByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAlgorithmDetailByIdResponse{}
	_body, _err := client.GetAlgorithmDetailByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAlgorithmDetailByNameWithOptions(request *GetAlgorithmDetailByNameRequest, runtime *util.RuntimeOptions) (_result *GetAlgorithmDetailByNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAlgorithmDetailByNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAlgorithmDetailByName"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAlgorithmDetailByName(request *GetAlgorithmDetailByNameRequest) (_result *GetAlgorithmDetailByNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAlgorithmDetailByNameResponse{}
	_body, _err := client.GetAlgorithmDetailByNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDevicePurifyJobByJobIdWithOptions(request *GetDevicePurifyJobByJobIdRequest, runtime *util.RuntimeOptions) (_result *GetDevicePurifyJobByJobIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDevicePurifyJobByJobIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDevicePurifyJobByJobId"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDevicePurifyJobByJobId(request *GetDevicePurifyJobByJobIdRequest) (_result *GetDevicePurifyJobByJobIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDevicePurifyJobByJobIdResponse{}
	_body, _err := client.GetDevicePurifyJobByJobIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetModelDetailWithOptions(request *GetModelDetailRequest, runtime *util.RuntimeOptions) (_result *GetModelDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetModelDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetModelDetail"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetModelDetail(request *GetModelDetailRequest) (_result *GetModelDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetModelDetailResponse{}
	_body, _err := client.GetModelDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetModelDetailByIdWithOptions(request *GetModelDetailByIdRequest, runtime *util.RuntimeOptions) (_result *GetModelDetailByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetModelDetailByIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetModelDetailById"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetModelDetailById(request *GetModelDetailByIdRequest) (_result *GetModelDetailByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetModelDetailByIdResponse{}
	_body, _err := client.GetModelDetailByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetModelOssPolicyWithOptions(request *GetModelOssPolicyRequest, runtime *util.RuntimeOptions) (_result *GetModelOssPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetModelOssPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetModelOssPolicy"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetModelOssPolicy(request *GetModelOssPolicyRequest) (_result *GetModelOssPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetModelOssPolicyResponse{}
	_body, _err := client.GetModelOssPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPictureInfoWithVectorIdsWithOptions(request *GetPictureInfoWithVectorIdsRequest, runtime *util.RuntimeOptions) (_result *GetPictureInfoWithVectorIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPictureInfoWithVectorIdsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPictureInfoWithVectorIds"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPictureInfoWithVectorIds(request *GetPictureInfoWithVectorIdsRequest) (_result *GetPictureInfoWithVectorIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPictureInfoWithVectorIdsResponse{}
	_body, _err := client.GetPictureInfoWithVectorIdsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPictureWithVectorIdWithOptions(request *GetPictureWithVectorIdRequest, runtime *util.RuntimeOptions) (_result *GetPictureWithVectorIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPictureWithVectorIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPictureWithVectorId"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPictureWithVectorId(request *GetPictureWithVectorIdRequest) (_result *GetPictureWithVectorIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPictureWithVectorIdResponse{}
	_body, _err := client.GetPictureWithVectorIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAlgorithmsByPageWithOptions(request *ListAlgorithmsByPageRequest, runtime *util.RuntimeOptions) (_result *ListAlgorithmsByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAlgorithmsByPageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAlgorithmsByPage"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAlgorithmsByPage(request *ListAlgorithmsByPageRequest) (_result *ListAlgorithmsByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAlgorithmsByPageResponse{}
	_body, _err := client.ListAlgorithmsByPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeployTaskByModelIdAndDevicesWithOptions(request *ListDeployTaskByModelIdAndDevicesRequest, runtime *util.RuntimeOptions) (_result *ListDeployTaskByModelIdAndDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDeployTaskByModelIdAndDevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDeployTaskByModelIdAndDevices"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeployTaskByModelIdAndDevices(request *ListDeployTaskByModelIdAndDevicesRequest) (_result *ListDeployTaskByModelIdAndDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeployTaskByModelIdAndDevicesResponse{}
	_body, _err := client.ListDeployTaskByModelIdAndDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeployTaskByPageWithOptions(request *ListDeployTaskByPageRequest, runtime *util.RuntimeOptions) (_result *ListDeployTaskByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDeployTaskByPageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDeployTaskByPage"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeployTaskByPage(request *ListDeployTaskByPageRequest) (_result *ListDeployTaskByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeployTaskByPageResponse{}
	_body, _err := client.ListDeployTaskByPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListModelsByPageWithOptions(request *ListModelsByPageRequest, runtime *util.RuntimeOptions) (_result *ListModelsByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListModelsByPageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListModelsByPage"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListModelsByPage(request *ListModelsByPageRequest) (_result *ListModelsByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListModelsByPageResponse{}
	_body, _err := client.ListModelsByPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListModelVersionsByPageWithOptions(request *ListModelVersionsByPageRequest, runtime *util.RuntimeOptions) (_result *ListModelVersionsByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListModelVersionsByPageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListModelVersionsByPage"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListModelVersionsByPage(request *ListModelVersionsByPageRequest) (_result *ListModelVersionsByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListModelVersionsByPageResponse{}
	_body, _err := client.ListModelVersionsByPageWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PictureSearchPictureResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PictureSearchPicture"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) QueryAIActionWithOptions(request *QueryAIActionRequest, runtime *util.RuntimeOptions) (_result *QueryAIActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryAIActionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryAIAction"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAIAction(request *QueryAIActionRequest) (_result *QueryAIActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAIActionResponse{}
	_body, _err := client.QueryAIActionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAIAppWithOptions(request *QueryAIAppRequest, runtime *util.RuntimeOptions) (_result *QueryAIAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryAIAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryAIApp"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAIApp(request *QueryAIAppRequest) (_result *QueryAIAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAIAppResponse{}
	_body, _err := client.QueryAIAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAIJobsWithOptions(request *QueryAIJobsRequest, runtime *util.RuntimeOptions) (_result *QueryAIJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryAIJobsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryAIJobs"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAIJobs(request *QueryAIJobsRequest) (_result *QueryAIJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAIJobsResponse{}
	_body, _err := client.QueryAIJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAIPlanWithOptions(request *QueryAIPlanRequest, runtime *util.RuntimeOptions) (_result *QueryAIPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryAIPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryAIPlan"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAIPlan(request *QueryAIPlanRequest) (_result *QueryAIPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAIPlanResponse{}
	_body, _err := client.QueryAIPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAIPlanTemplatesWithOptions(request *QueryAIPlanTemplatesRequest, runtime *util.RuntimeOptions) (_result *QueryAIPlanTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryAIPlanTemplatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryAIPlanTemplates"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAIPlanTemplates(request *QueryAIPlanTemplatesRequest) (_result *QueryAIPlanTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAIPlanTemplatesResponse{}
	_body, _err := client.QueryAIPlanTemplatesWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryDeviceEventResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryDeviceEvent"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryDeviceEventPictureResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryDeviceEventPicture"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryDeviceEventRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryDeviceEventRecord"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) QueryDeviceFileVodWithOptions(request *QueryDeviceFileVodRequest, runtime *util.RuntimeOptions) (_result *QueryDeviceFileVodResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryDeviceFileVodResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryDeviceFileVod"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDeviceFileVod(request *QueryDeviceFileVodRequest) (_result *QueryDeviceFileVodResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDeviceFileVodResponse{}
	_body, _err := client.QueryDeviceFileVodWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryDevicePictureFileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryDevicePictureFile"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryDevicePictureLifeCycleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryDevicePictureLifeCycle"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) QueryDevicePurifyJobsWithOptions(request *QueryDevicePurifyJobsRequest, runtime *util.RuntimeOptions) (_result *QueryDevicePurifyJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryDevicePurifyJobsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryDevicePurifyJobs"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDevicePurifyJobs(request *QueryDevicePurifyJobsRequest) (_result *QueryDevicePurifyJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDevicePurifyJobsResponse{}
	_body, _err := client.QueryDevicePurifyJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDevicePurifyPlanByPlanIdWithOptions(request *QueryDevicePurifyPlanByPlanIdRequest, runtime *util.RuntimeOptions) (_result *QueryDevicePurifyPlanByPlanIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryDevicePurifyPlanByPlanIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryDevicePurifyPlanByPlanId"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDevicePurifyPlanByPlanId(request *QueryDevicePurifyPlanByPlanIdRequest) (_result *QueryDevicePurifyPlanByPlanIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDevicePurifyPlanByPlanIdResponse{}
	_body, _err := client.QueryDevicePurifyPlanByPlanIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDevicePurifyPlansWithOptions(request *QueryDevicePurifyPlansRequest, runtime *util.RuntimeOptions) (_result *QueryDevicePurifyPlansResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryDevicePurifyPlansResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryDevicePurifyPlans"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDevicePurifyPlans(request *QueryDevicePurifyPlansRequest) (_result *QueryDevicePurifyPlansResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDevicePurifyPlansResponse{}
	_body, _err := client.QueryDevicePurifyPlansWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryDeviceRecordLifeCycleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryDeviceRecordLifeCycle"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryDeviceVodUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryDeviceVodUrl"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) QueryEventRecordPlanDetailWithOptions(request *QueryEventRecordPlanDetailRequest, runtime *util.RuntimeOptions) (_result *QueryEventRecordPlanDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryEventRecordPlanDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryEventRecordPlanDetail"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryEventRecordPlanDeviceByDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryEventRecordPlanDeviceByDevice"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryEventRecordPlanDeviceByPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryEventRecordPlanDeviceByPlan"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryEventRecordPlansResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryEventRecordPlans"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryFaceAllDeviceGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryFaceAllDeviceGroup"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryFaceAllUserGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryFaceAllUserGroup"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryFaceAllUserGroupAndDeviceGroupRelationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryFaceAllUserGroupAndDeviceGroupRelation"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryFaceAllUserIdsByGroupIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryFaceAllUserIdsByGroupId"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryFaceCustomUserIdByUserIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryFaceCustomUserIdByUserId"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryFaceDeviceGroupsByDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryFaceDeviceGroupsByDevice"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryFaceUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryFaceUser"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) QueryFaceUserGroupWithOptions(request *QueryFaceUserGroupRequest, runtime *util.RuntimeOptions) (_result *QueryFaceUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryFaceUserGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryFaceUserGroup"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryFaceUserGroupAndDeviceGroupRelationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryFaceUserGroupAndDeviceGroupRelation"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryFaceUserIdByCustomUserIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryFaceUserIdByCustomUserId"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) QueryIotIdsByAIPlanWithOptions(request *QueryIotIdsByAIPlanRequest, runtime *util.RuntimeOptions) (_result *QueryIotIdsByAIPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryIotIdsByAIPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryIotIdsByAIPlan"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryIotIdsByAIPlan(request *QueryIotIdsByAIPlanRequest) (_result *QueryIotIdsByAIPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryIotIdsByAIPlanResponse{}
	_body, _err := client.QueryIotIdsByAIPlanWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryLiveStreamingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryLiveStreaming"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) QueryMonthRecordWithOptions(request *QueryMonthRecordRequest, runtime *util.RuntimeOptions) (_result *QueryMonthRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryMonthRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryMonthRecord"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryPictureFilesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryPictureFiles"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryPictureSearchAiboxesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryPictureSearchAiboxes"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) QueryPictureSearchAppWithOptions(request *QueryPictureSearchAppRequest, runtime *util.RuntimeOptions) (_result *QueryPictureSearchAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryPictureSearchAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryPictureSearchApp"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPictureSearchApp(request *QueryPictureSearchAppRequest) (_result *QueryPictureSearchAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPictureSearchAppResponse{}
	_body, _err := client.QueryPictureSearchAppWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryPictureSearchDevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryPictureSearchDevices"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) QueryRecordWithOptions(request *QueryRecordRequest, runtime *util.RuntimeOptions) (_result *QueryRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryRecord"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryRecordByRecordIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryRecordByRecordId"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) QueryRecordPlanDetailWithOptions(request *QueryRecordPlanDetailRequest, runtime *util.RuntimeOptions) (_result *QueryRecordPlanDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryRecordPlanDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryRecordPlanDetail"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryRecordPlanDeviceByDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryRecordPlanDeviceByDevice"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryRecordPlanDeviceByPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryRecordPlanDeviceByPlan"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryRecordPlansResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryRecordPlans"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryRecordUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryRecordUrl"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) QueryStandardAIAppTemplatesWithOptions(request *QueryStandardAIAppTemplatesRequest, runtime *util.RuntimeOptions) (_result *QueryStandardAIAppTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryStandardAIAppTemplatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryStandardAIAppTemplates"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryStandardAIAppTemplates(request *QueryStandardAIAppTemplatesRequest) (_result *QueryStandardAIAppTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryStandardAIAppTemplatesResponse{}
	_body, _err := client.QueryStandardAIAppTemplatesWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryTimeTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryTimeTemplate"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryTimeTemplateDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryTimeTemplateDetail"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) QueryVoiceIntercomWithOptions(request *QueryVoiceIntercomRequest, runtime *util.RuntimeOptions) (_result *QueryVoiceIntercomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryVoiceIntercomResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryVoiceIntercom"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RemoveAIAppWithOptions(request *RemoveAIAppRequest, runtime *util.RuntimeOptions) (_result *RemoveAIAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveAIAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveAIApp"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveAIApp(request *RemoveAIAppRequest) (_result *RemoveAIAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveAIAppResponse{}
	_body, _err := client.RemoveAIAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveAIPlanWithOptions(request *RemoveAIPlanRequest, runtime *util.RuntimeOptions) (_result *RemoveAIPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveAIPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveAIPlan"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveAIPlan(request *RemoveAIPlanRequest) (_result *RemoveAIPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveAIPlanResponse{}
	_body, _err := client.RemoveAIPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveDevicePurifyPlanWithOptions(request *RemoveDevicePurifyPlanRequest, runtime *util.RuntimeOptions) (_result *RemoveDevicePurifyPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveDevicePurifyPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveDevicePurifyPlan"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveDevicePurifyPlan(request *RemoveDevicePurifyPlanRequest) (_result *RemoveDevicePurifyPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveDevicePurifyPlanResponse{}
	_body, _err := client.RemoveDevicePurifyPlanWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveFaceDeviceFromDeviceGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveFaceDeviceFromDeviceGroup"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveFaceUserFromUserGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveFaceUserFromUserGroup"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetDevicePictureLifeCycleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetDevicePictureLifeCycle"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetDeviceRecordLifeCycleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetDeviceRecordLifeCycle"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopLiveStreamingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopLiveStreaming"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopTriggeredRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopTriggeredRecord"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) TriggerCapturePictureWithOptions(request *TriggerCapturePictureRequest, runtime *util.RuntimeOptions) (_result *TriggerCapturePictureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TriggerCapturePictureResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TriggerCapturePicture"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TriggerRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TriggerRecord"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UnbindAIPlanWithDevicesWithOptions(request *UnbindAIPlanWithDevicesRequest, runtime *util.RuntimeOptions) (_result *UnbindAIPlanWithDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnbindAIPlanWithDevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnbindAIPlanWithDevices"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindAIPlanWithDevices(request *UnbindAIPlanWithDevicesRequest) (_result *UnbindAIPlanWithDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindAIPlanWithDevicesResponse{}
	_body, _err := client.UnbindAIPlanWithDevicesWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnbindPictureSearchAppWithDevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnbindPictureSearchAppWithDevices"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UpdateAIAppWithOptions(request *UpdateAIAppRequest, runtime *util.RuntimeOptions) (_result *UpdateAIAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAIAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAIApp"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAIApp(request *UpdateAIAppRequest) (_result *UpdateAIAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAIAppResponse{}
	_body, _err := client.UpdateAIAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAIPlanWithOptions(request *UpdateAIPlanRequest, runtime *util.RuntimeOptions) (_result *UpdateAIPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAIPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAIPlan"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAIPlan(request *UpdateAIPlanRequest) (_result *UpdateAIPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAIPlanResponse{}
	_body, _err := client.UpdateAIPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDevicePurifyPlanWithOptions(request *UpdateDevicePurifyPlanRequest, runtime *util.RuntimeOptions) (_result *UpdateDevicePurifyPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDevicePurifyPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDevicePurifyPlan"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDevicePurifyPlan(request *UpdateDevicePurifyPlanRequest) (_result *UpdateDevicePurifyPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDevicePurifyPlanResponse{}
	_body, _err := client.UpdateDevicePurifyPlanWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateEventRecordPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateEventRecordPlan"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateFaceUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateFaceUser"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateFaceUserGroupAndDeviceGroupRelationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateFaceUserGroupAndDeviceGroupRelation"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UpdateModelWithOptions(request *UpdateModelRequest, runtime *util.RuntimeOptions) (_result *UpdateModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateModelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateModel"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateModel(request *UpdateModelRequest) (_result *UpdateModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateModelResponse{}
	_body, _err := client.UpdateModelWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateRecordPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateRecordPlan"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UpdateTimeTemplateWithOptions(request *UpdateTimeTemplateRequest, runtime *util.RuntimeOptions) (_result *UpdateTimeTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateTimeTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateTimeTemplate"), tea.String("2018-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
