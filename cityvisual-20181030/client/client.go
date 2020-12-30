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

type AttachStreamRequest struct {
	JobGroupId *string                       `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	RegionId   *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string                       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Streams    []*AttachStreamRequestStreams `json:"Streams,omitempty" xml:"Streams,omitempty" type:"Repeated"`
}

func (s AttachStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachStreamRequest) GoString() string {
	return s.String()
}

func (s *AttachStreamRequest) SetJobGroupId(v string) *AttachStreamRequest {
	s.JobGroupId = &v
	return s
}

func (s *AttachStreamRequest) SetRegionId(v string) *AttachStreamRequest {
	s.RegionId = &v
	return s
}

func (s *AttachStreamRequest) SetInstanceId(v string) *AttachStreamRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachStreamRequest) SetStreams(v []*AttachStreamRequestStreams) *AttachStreamRequest {
	s.Streams = v
	return s
}

type AttachStreamRequestStreams struct {
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	StreamURL  *string `json:"StreamURL,omitempty" xml:"StreamURL,omitempty"`
}

func (s AttachStreamRequestStreams) String() string {
	return tea.Prettify(s)
}

func (s AttachStreamRequestStreams) GoString() string {
	return s.String()
}

func (s *AttachStreamRequestStreams) SetStreamName(v string) *AttachStreamRequestStreams {
	s.StreamName = &v
	return s
}

func (s *AttachStreamRequestStreams) SetStreamURL(v string) *AttachStreamRequestStreams {
	s.StreamURL = &v
	return s
}

type AttachStreamResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachStreamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachStreamResponseBody) GoString() string {
	return s.String()
}

func (s *AttachStreamResponseBody) SetRequestId(v string) *AttachStreamResponseBody {
	s.RequestId = &v
	return s
}

type AttachStreamResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachStreamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachStreamResponse) GoString() string {
	return s.String()
}

func (s *AttachStreamResponse) SetHeaders(v map[string]*string) *AttachStreamResponse {
	s.Headers = v
	return s
}

func (s *AttachStreamResponse) SetBody(v *AttachStreamResponseBody) *AttachStreamResponse {
	s.Body = v
	return s
}

type BatchModifyCameraStatusRequest struct {
	StreamStatus *string `json:"StreamStatus,omitempty" xml:"StreamStatus,omitempty"`
	CameraIds    *string `json:"CameraIds,omitempty" xml:"CameraIds,omitempty"`
	WorkGroupId  *string `json:"WorkGroupId,omitempty" xml:"WorkGroupId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s BatchModifyCameraStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchModifyCameraStatusRequest) GoString() string {
	return s.String()
}

func (s *BatchModifyCameraStatusRequest) SetStreamStatus(v string) *BatchModifyCameraStatusRequest {
	s.StreamStatus = &v
	return s
}

func (s *BatchModifyCameraStatusRequest) SetCameraIds(v string) *BatchModifyCameraStatusRequest {
	s.CameraIds = &v
	return s
}

func (s *BatchModifyCameraStatusRequest) SetWorkGroupId(v string) *BatchModifyCameraStatusRequest {
	s.WorkGroupId = &v
	return s
}

func (s *BatchModifyCameraStatusRequest) SetRegionId(v string) *BatchModifyCameraStatusRequest {
	s.RegionId = &v
	return s
}

type BatchModifyCameraStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchModifyCameraStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchModifyCameraStatusResponseBody) GoString() string {
	return s.String()
}

func (s *BatchModifyCameraStatusResponseBody) SetRequestId(v string) *BatchModifyCameraStatusResponseBody {
	s.RequestId = &v
	return s
}

type BatchModifyCameraStatusResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchModifyCameraStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchModifyCameraStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchModifyCameraStatusResponse) GoString() string {
	return s.String()
}

func (s *BatchModifyCameraStatusResponse) SetHeaders(v map[string]*string) *BatchModifyCameraStatusResponse {
	s.Headers = v
	return s
}

func (s *BatchModifyCameraStatusResponse) SetBody(v *BatchModifyCameraStatusResponseBody) *BatchModifyCameraStatusResponse {
	s.Body = v
	return s
}

type CreateAlgoLibRequest struct {
	AlgoLibName        *string `json:"AlgoLibName,omitempty" xml:"AlgoLibName,omitempty"`
	AlgoLibVersion     *string `json:"AlgoLibVersion,omitempty" xml:"AlgoLibVersion,omitempty"`
	Capacity           *int32  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	OssEndpoint        *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	OssAccessKeyId     *string `json:"OssAccessKeyId,omitempty" xml:"OssAccessKeyId,omitempty"`
	OssBucket          *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssPath            *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	JsonText           *string `json:"JsonText,omitempty" xml:"JsonText,omitempty"`
	CapabilityNames    *string `json:"CapabilityNames,omitempty" xml:"CapabilityNames,omitempty"`
	OssAccessKeySecret *string `json:"OssAccessKeySecret,omitempty" xml:"OssAccessKeySecret,omitempty"`
	Text               *string `json:"Text,omitempty" xml:"Text,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateAlgoLibRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAlgoLibRequest) GoString() string {
	return s.String()
}

func (s *CreateAlgoLibRequest) SetAlgoLibName(v string) *CreateAlgoLibRequest {
	s.AlgoLibName = &v
	return s
}

func (s *CreateAlgoLibRequest) SetAlgoLibVersion(v string) *CreateAlgoLibRequest {
	s.AlgoLibVersion = &v
	return s
}

func (s *CreateAlgoLibRequest) SetCapacity(v int32) *CreateAlgoLibRequest {
	s.Capacity = &v
	return s
}

func (s *CreateAlgoLibRequest) SetOssEndpoint(v string) *CreateAlgoLibRequest {
	s.OssEndpoint = &v
	return s
}

func (s *CreateAlgoLibRequest) SetOssAccessKeyId(v string) *CreateAlgoLibRequest {
	s.OssAccessKeyId = &v
	return s
}

func (s *CreateAlgoLibRequest) SetOssBucket(v string) *CreateAlgoLibRequest {
	s.OssBucket = &v
	return s
}

func (s *CreateAlgoLibRequest) SetOssPath(v string) *CreateAlgoLibRequest {
	s.OssPath = &v
	return s
}

func (s *CreateAlgoLibRequest) SetJsonText(v string) *CreateAlgoLibRequest {
	s.JsonText = &v
	return s
}

func (s *CreateAlgoLibRequest) SetCapabilityNames(v string) *CreateAlgoLibRequest {
	s.CapabilityNames = &v
	return s
}

func (s *CreateAlgoLibRequest) SetOssAccessKeySecret(v string) *CreateAlgoLibRequest {
	s.OssAccessKeySecret = &v
	return s
}

func (s *CreateAlgoLibRequest) SetText(v string) *CreateAlgoLibRequest {
	s.Text = &v
	return s
}

func (s *CreateAlgoLibRequest) SetInstanceId(v string) *CreateAlgoLibRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAlgoLibRequest) SetRegionId(v string) *CreateAlgoLibRequest {
	s.RegionId = &v
	return s
}

type CreateAlgoLibResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AlgoLibId *string `json:"AlgoLibId,omitempty" xml:"AlgoLibId,omitempty"`
}

func (s CreateAlgoLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAlgoLibResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlgoLibResponseBody) SetRequestId(v string) *CreateAlgoLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAlgoLibResponseBody) SetAlgoLibId(v string) *CreateAlgoLibResponseBody {
	s.AlgoLibId = &v
	return s
}

type CreateAlgoLibResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAlgoLibResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAlgoLibResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAlgoLibResponse) GoString() string {
	return s.String()
}

func (s *CreateAlgoLibResponse) SetHeaders(v map[string]*string) *CreateAlgoLibResponse {
	s.Headers = v
	return s
}

func (s *CreateAlgoLibResponse) SetBody(v *CreateAlgoLibResponseBody) *CreateAlgoLibResponse {
	s.Body = v
	return s
}

type CreateCameraRequest struct {
	CameraName   *string `json:"CameraName,omitempty" xml:"CameraName,omitempty"`
	WorkGroupId  *string `json:"WorkGroupId,omitempty" xml:"WorkGroupId,omitempty"`
	CameraId     *string `json:"CameraId,omitempty" xml:"CameraId,omitempty"`
	LocationInfo *string `json:"LocationInfo,omitempty" xml:"LocationInfo,omitempty"`
	InviteUri    *string `json:"InviteUri,omitempty" xml:"InviteUri,omitempty"`
	OssPath      *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateCameraRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCameraRequest) GoString() string {
	return s.String()
}

func (s *CreateCameraRequest) SetCameraName(v string) *CreateCameraRequest {
	s.CameraName = &v
	return s
}

func (s *CreateCameraRequest) SetWorkGroupId(v string) *CreateCameraRequest {
	s.WorkGroupId = &v
	return s
}

func (s *CreateCameraRequest) SetCameraId(v string) *CreateCameraRequest {
	s.CameraId = &v
	return s
}

func (s *CreateCameraRequest) SetLocationInfo(v string) *CreateCameraRequest {
	s.LocationInfo = &v
	return s
}

func (s *CreateCameraRequest) SetInviteUri(v string) *CreateCameraRequest {
	s.InviteUri = &v
	return s
}

func (s *CreateCameraRequest) SetOssPath(v string) *CreateCameraRequest {
	s.OssPath = &v
	return s
}

func (s *CreateCameraRequest) SetInstanceId(v string) *CreateCameraRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCameraRequest) SetRegionId(v string) *CreateCameraRequest {
	s.RegionId = &v
	return s
}

type CreateCameraResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CameraId  *string `json:"CameraId,omitempty" xml:"CameraId,omitempty"`
}

func (s CreateCameraResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCameraResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCameraResponseBody) SetRequestId(v string) *CreateCameraResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCameraResponseBody) SetCameraId(v string) *CreateCameraResponseBody {
	s.CameraId = &v
	return s
}

type CreateCameraResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCameraResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCameraResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCameraResponse) GoString() string {
	return s.String()
}

func (s *CreateCameraResponse) SetHeaders(v map[string]*string) *CreateCameraResponse {
	s.Headers = v
	return s
}

func (s *CreateCameraResponse) SetBody(v *CreateCameraResponseBody) *CreateCameraResponse {
	s.Body = v
	return s
}

type CreateCapabilityRequest struct {
	CapabilityName *string `json:"CapabilityName,omitempty" xml:"CapabilityName,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateCapabilityRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCapabilityRequest) GoString() string {
	return s.String()
}

func (s *CreateCapabilityRequest) SetCapabilityName(v string) *CreateCapabilityRequest {
	s.CapabilityName = &v
	return s
}

func (s *CreateCapabilityRequest) SetInstanceId(v string) *CreateCapabilityRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCapabilityRequest) SetRegionId(v string) *CreateCapabilityRequest {
	s.RegionId = &v
	return s
}

type CreateCapabilityResponseBody struct {
	CapabilityId *string `json:"CapabilityId,omitempty" xml:"CapabilityId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCapabilityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCapabilityResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCapabilityResponseBody) SetCapabilityId(v string) *CreateCapabilityResponseBody {
	s.CapabilityId = &v
	return s
}

func (s *CreateCapabilityResponseBody) SetRequestId(v string) *CreateCapabilityResponseBody {
	s.RequestId = &v
	return s
}

type CreateCapabilityResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCapabilityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCapabilityResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCapabilityResponse) GoString() string {
	return s.String()
}

func (s *CreateCapabilityResponse) SetHeaders(v map[string]*string) *CreateCapabilityResponse {
	s.Headers = v
	return s
}

func (s *CreateCapabilityResponse) SetBody(v *CreateCapabilityResponseBody) *CreateCapabilityResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	InstanceName     *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceCapacity *int32  `json:"InstanceCapacity,omitempty" xml:"InstanceCapacity,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetDescription(v string) *CreateInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceCapacity(v int32) *CreateInstanceRequest {
	s.InstanceCapacity = &v
	return s
}

func (s *CreateInstanceRequest) SetStatus(v string) *CreateInstanceRequest {
	s.Status = &v
	return s
}

func (s *CreateInstanceRequest) SetRegionId(v string) *CreateInstanceRequest {
	s.RegionId = &v
	return s
}

type CreateInstanceResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetInstanceId(v string) *CreateInstanceResponseBody {
	s.InstanceId = &v
	return s
}

type CreateInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponse) SetHeaders(v map[string]*string) *CreateInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
	s.Body = v
	return s
}

type CreateJobGroupRequest struct {
	JobGroupName      *string                                `json:"JobGroupName,omitempty" xml:"JobGroupName,omitempty"`
	ResourceProfileId *string                                `json:"ResourceProfileId,omitempty" xml:"ResourceProfileId,omitempty"`
	AlgoLibId         *string                                `json:"AlgoLibId,omitempty" xml:"AlgoLibId,omitempty"`
	PlanedJobCount    *int32                                 `json:"PlanedJobCount,omitempty" xml:"PlanedJobCount,omitempty"`
	StreamPerJob      *int32                                 `json:"StreamPerJob,omitempty" xml:"StreamPerJob,omitempty"`
	JobCount          *int32                                 `json:"JobCount,omitempty" xml:"JobCount,omitempty"`
	InstanceId        *string                                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId          *string                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	JobGroupParams    []*CreateJobGroupRequestJobGroupParams `json:"JobGroupParams,omitempty" xml:"JobGroupParams,omitempty" type:"Repeated"`
}

func (s CreateJobGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateJobGroupRequest) SetJobGroupName(v string) *CreateJobGroupRequest {
	s.JobGroupName = &v
	return s
}

func (s *CreateJobGroupRequest) SetResourceProfileId(v string) *CreateJobGroupRequest {
	s.ResourceProfileId = &v
	return s
}

func (s *CreateJobGroupRequest) SetAlgoLibId(v string) *CreateJobGroupRequest {
	s.AlgoLibId = &v
	return s
}

func (s *CreateJobGroupRequest) SetPlanedJobCount(v int32) *CreateJobGroupRequest {
	s.PlanedJobCount = &v
	return s
}

func (s *CreateJobGroupRequest) SetStreamPerJob(v int32) *CreateJobGroupRequest {
	s.StreamPerJob = &v
	return s
}

func (s *CreateJobGroupRequest) SetJobCount(v int32) *CreateJobGroupRequest {
	s.JobCount = &v
	return s
}

func (s *CreateJobGroupRequest) SetInstanceId(v string) *CreateJobGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateJobGroupRequest) SetRegionId(v string) *CreateJobGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateJobGroupRequest) SetJobGroupParams(v []*CreateJobGroupRequestJobGroupParams) *CreateJobGroupRequest {
	s.JobGroupParams = v
	return s
}

type CreateJobGroupRequestJobGroupParams struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateJobGroupRequestJobGroupParams) String() string {
	return tea.Prettify(s)
}

func (s CreateJobGroupRequestJobGroupParams) GoString() string {
	return s.String()
}

func (s *CreateJobGroupRequestJobGroupParams) SetKey(v string) *CreateJobGroupRequestJobGroupParams {
	s.Key = &v
	return s
}

func (s *CreateJobGroupRequestJobGroupParams) SetValue(v string) *CreateJobGroupRequestJobGroupParams {
	s.Value = &v
	return s
}

type CreateJobGroupResponseBody struct {
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateJobGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateJobGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobGroupResponseBody) SetJobGroupId(v string) *CreateJobGroupResponseBody {
	s.JobGroupId = &v
	return s
}

func (s *CreateJobGroupResponseBody) SetRequestId(v string) *CreateJobGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateJobGroupResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateJobGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateJobGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateJobGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateJobGroupResponse) SetHeaders(v map[string]*string) *CreateJobGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateJobGroupResponse) SetBody(v *CreateJobGroupResponseBody) *CreateJobGroupResponse {
	s.Body = v
	return s
}

type CreateResourceProfileRequest struct {
	ResourceProfileName   *string                                              `json:"ResourceProfileName,omitempty" xml:"ResourceProfileName,omitempty"`
	InstanceId            *string                                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId              *string                                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceProfileParams []*CreateResourceProfileRequestResourceProfileParams `json:"ResourceProfileParams,omitempty" xml:"ResourceProfileParams,omitempty" type:"Repeated"`
}

func (s CreateResourceProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceProfileRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceProfileRequest) SetResourceProfileName(v string) *CreateResourceProfileRequest {
	s.ResourceProfileName = &v
	return s
}

func (s *CreateResourceProfileRequest) SetInstanceId(v string) *CreateResourceProfileRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateResourceProfileRequest) SetRegionId(v string) *CreateResourceProfileRequest {
	s.RegionId = &v
	return s
}

func (s *CreateResourceProfileRequest) SetResourceProfileParams(v []*CreateResourceProfileRequestResourceProfileParams) *CreateResourceProfileRequest {
	s.ResourceProfileParams = v
	return s
}

type CreateResourceProfileRequestResourceProfileParams struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateResourceProfileRequestResourceProfileParams) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceProfileRequestResourceProfileParams) GoString() string {
	return s.String()
}

func (s *CreateResourceProfileRequestResourceProfileParams) SetKey(v string) *CreateResourceProfileRequestResourceProfileParams {
	s.Key = &v
	return s
}

func (s *CreateResourceProfileRequestResourceProfileParams) SetValue(v string) *CreateResourceProfileRequestResourceProfileParams {
	s.Value = &v
	return s
}

type CreateResourceProfileResponseBody struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceProfileId *string `json:"ResourceProfileId,omitempty" xml:"ResourceProfileId,omitempty"`
}

func (s CreateResourceProfileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceProfileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceProfileResponseBody) SetRequestId(v string) *CreateResourceProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceProfileResponseBody) SetResourceProfileId(v string) *CreateResourceProfileResponseBody {
	s.ResourceProfileId = &v
	return s
}

type CreateResourceProfileResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateResourceProfileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateResourceProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceProfileResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceProfileResponse) SetHeaders(v map[string]*string) *CreateResourceProfileResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceProfileResponse) SetBody(v *CreateResourceProfileResponseBody) *CreateResourceProfileResponse {
	s.Body = v
	return s
}

type CreateWorkGroupRequest struct {
	WorkGroupName   *string `json:"WorkGroupName,omitempty" xml:"WorkGroupName,omitempty"`
	Protocol        *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	TheoryCutStatus *string `json:"TheoryCutStatus,omitempty" xml:"TheoryCutStatus,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	JobNum          *int32  `json:"JobNum,omitempty" xml:"JobNum,omitempty"`
	TopicNum        *int32  `json:"TopicNum,omitempty" xml:"TopicNum,omitempty"`
	TopicPrefix     *string `json:"TopicPrefix,omitempty" xml:"TopicPrefix,omitempty"`
	Account         *string `json:"Account,omitempty" xml:"Account,omitempty"`
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Ip              *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Port            *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateWorkGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkGroupRequest) SetWorkGroupName(v string) *CreateWorkGroupRequest {
	s.WorkGroupName = &v
	return s
}

func (s *CreateWorkGroupRequest) SetProtocol(v string) *CreateWorkGroupRequest {
	s.Protocol = &v
	return s
}

func (s *CreateWorkGroupRequest) SetTheoryCutStatus(v string) *CreateWorkGroupRequest {
	s.TheoryCutStatus = &v
	return s
}

func (s *CreateWorkGroupRequest) SetDescription(v string) *CreateWorkGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkGroupRequest) SetJobNum(v int32) *CreateWorkGroupRequest {
	s.JobNum = &v
	return s
}

func (s *CreateWorkGroupRequest) SetTopicNum(v int32) *CreateWorkGroupRequest {
	s.TopicNum = &v
	return s
}

func (s *CreateWorkGroupRequest) SetTopicPrefix(v string) *CreateWorkGroupRequest {
	s.TopicPrefix = &v
	return s
}

func (s *CreateWorkGroupRequest) SetAccount(v string) *CreateWorkGroupRequest {
	s.Account = &v
	return s
}

func (s *CreateWorkGroupRequest) SetPassword(v string) *CreateWorkGroupRequest {
	s.Password = &v
	return s
}

func (s *CreateWorkGroupRequest) SetIp(v string) *CreateWorkGroupRequest {
	s.Ip = &v
	return s
}

func (s *CreateWorkGroupRequest) SetPort(v int32) *CreateWorkGroupRequest {
	s.Port = &v
	return s
}

func (s *CreateWorkGroupRequest) SetInstanceId(v string) *CreateWorkGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateWorkGroupRequest) SetRegionId(v string) *CreateWorkGroupRequest {
	s.RegionId = &v
	return s
}

type CreateWorkGroupResponseBody struct {
	WorkGroupId *string `json:"WorkGroupId,omitempty" xml:"WorkGroupId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWorkGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkGroupResponseBody) SetWorkGroupId(v string) *CreateWorkGroupResponseBody {
	s.WorkGroupId = &v
	return s
}

func (s *CreateWorkGroupResponseBody) SetRequestId(v string) *CreateWorkGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateWorkGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateWorkGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWorkGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkGroupResponse) SetHeaders(v map[string]*string) *CreateWorkGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkGroupResponse) SetBody(v *CreateWorkGroupResponseBody) *CreateWorkGroupResponse {
	s.Body = v
	return s
}

type DeleteAlgoLibRequest struct {
	AlgoLibId *string `json:"AlgoLibId,omitempty" xml:"AlgoLibId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAlgoLibRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlgoLibRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlgoLibRequest) SetAlgoLibId(v string) *DeleteAlgoLibRequest {
	s.AlgoLibId = &v
	return s
}

func (s *DeleteAlgoLibRequest) SetRegionId(v string) *DeleteAlgoLibRequest {
	s.RegionId = &v
	return s
}

type DeleteAlgoLibResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAlgoLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlgoLibResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlgoLibResponseBody) SetRequestId(v string) *DeleteAlgoLibResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAlgoLibResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAlgoLibResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAlgoLibResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlgoLibResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlgoLibResponse) SetHeaders(v map[string]*string) *DeleteAlgoLibResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlgoLibResponse) SetBody(v *DeleteAlgoLibResponseBody) *DeleteAlgoLibResponse {
	s.Body = v
	return s
}

type DeleteCameraRequest struct {
	CameraId *string `json:"CameraId,omitempty" xml:"CameraId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCameraRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCameraRequest) GoString() string {
	return s.String()
}

func (s *DeleteCameraRequest) SetCameraId(v string) *DeleteCameraRequest {
	s.CameraId = &v
	return s
}

func (s *DeleteCameraRequest) SetRegionId(v string) *DeleteCameraRequest {
	s.RegionId = &v
	return s
}

type DeleteCameraResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCameraResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCameraResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCameraResponseBody) SetRequestId(v string) *DeleteCameraResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCameraResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCameraResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCameraResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCameraResponse) GoString() string {
	return s.String()
}

func (s *DeleteCameraResponse) SetHeaders(v map[string]*string) *DeleteCameraResponse {
	s.Headers = v
	return s
}

func (s *DeleteCameraResponse) SetBody(v *DeleteCameraResponseBody) *DeleteCameraResponse {
	s.Body = v
	return s
}

type DeleteCapabilityRequest struct {
	CapabilityId *string `json:"CapabilityId,omitempty" xml:"CapabilityId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCapabilityRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCapabilityRequest) GoString() string {
	return s.String()
}

func (s *DeleteCapabilityRequest) SetCapabilityId(v string) *DeleteCapabilityRequest {
	s.CapabilityId = &v
	return s
}

func (s *DeleteCapabilityRequest) SetRegionId(v string) *DeleteCapabilityRequest {
	s.RegionId = &v
	return s
}

type DeleteCapabilityResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCapabilityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCapabilityResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCapabilityResponseBody) SetRequestId(v string) *DeleteCapabilityResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCapabilityResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCapabilityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCapabilityResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCapabilityResponse) GoString() string {
	return s.String()
}

func (s *DeleteCapabilityResponse) SetHeaders(v map[string]*string) *DeleteCapabilityResponse {
	s.Headers = v
	return s
}

func (s *DeleteCapabilityResponse) SetBody(v *DeleteCapabilityResponseBody) *DeleteCapabilityResponse {
	s.Body = v
	return s
}

type DeleteInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) SetInstanceId(v string) *DeleteInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceRequest) SetRegionId(v string) *DeleteInstanceRequest {
	s.RegionId = &v
	return s
}

type DeleteInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponse) SetHeaders(v map[string]*string) *DeleteInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceResponse) SetBody(v *DeleteInstanceResponseBody) *DeleteInstanceResponse {
	s.Body = v
	return s
}

type DeleteJobGroupRequest struct {
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteJobGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobGroupRequest) SetJobGroupId(v string) *DeleteJobGroupRequest {
	s.JobGroupId = &v
	return s
}

func (s *DeleteJobGroupRequest) SetRegionId(v string) *DeleteJobGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteJobGroupRequest) SetInstanceId(v string) *DeleteJobGroupRequest {
	s.InstanceId = &v
	return s
}

type DeleteJobGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteJobGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobGroupResponseBody) SetRequestId(v string) *DeleteJobGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteJobGroupResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteJobGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteJobGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteJobGroupResponse) SetHeaders(v map[string]*string) *DeleteJobGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteJobGroupResponse) SetBody(v *DeleteJobGroupResponseBody) *DeleteJobGroupResponse {
	s.Body = v
	return s
}

type DeleteResourceProfileRequest struct {
	ResourceProfileId *string `json:"ResourceProfileId,omitempty" xml:"ResourceProfileId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteResourceProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceProfileRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceProfileRequest) SetResourceProfileId(v string) *DeleteResourceProfileRequest {
	s.ResourceProfileId = &v
	return s
}

func (s *DeleteResourceProfileRequest) SetRegionId(v string) *DeleteResourceProfileRequest {
	s.RegionId = &v
	return s
}

type DeleteResourceProfileResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceProfileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceProfileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceProfileResponseBody) SetRequestId(v string) *DeleteResourceProfileResponseBody {
	s.RequestId = &v
	return s
}

type DeleteResourceProfileResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteResourceProfileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteResourceProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceProfileResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceProfileResponse) SetHeaders(v map[string]*string) *DeleteResourceProfileResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceProfileResponse) SetBody(v *DeleteResourceProfileResponseBody) *DeleteResourceProfileResponse {
	s.Body = v
	return s
}

type DeleteWorkGroupRequest struct {
	WorkGroupId *string `json:"WorkGroupId,omitempty" xml:"WorkGroupId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteWorkGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkGroupRequest) SetWorkGroupId(v string) *DeleteWorkGroupRequest {
	s.WorkGroupId = &v
	return s
}

func (s *DeleteWorkGroupRequest) SetRegionId(v string) *DeleteWorkGroupRequest {
	s.RegionId = &v
	return s
}

type DeleteWorkGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWorkGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkGroupResponseBody) SetRequestId(v string) *DeleteWorkGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWorkGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteWorkGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWorkGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkGroupResponse) SetHeaders(v map[string]*string) *DeleteWorkGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkGroupResponse) SetBody(v *DeleteWorkGroupResponseBody) *DeleteWorkGroupResponse {
	s.Body = v
	return s
}

type DescribeAlgoLibsRequest struct {
	AlgoLibName    *string `json:"AlgoLibName,omitempty" xml:"AlgoLibName,omitempty"`
	CapabilityName *string `json:"CapabilityName,omitempty" xml:"CapabilityName,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AlgoLibId      *string `json:"AlgoLibId,omitempty" xml:"AlgoLibId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAlgoLibsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlgoLibsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlgoLibsRequest) SetAlgoLibName(v string) *DescribeAlgoLibsRequest {
	s.AlgoLibName = &v
	return s
}

func (s *DescribeAlgoLibsRequest) SetCapabilityName(v string) *DescribeAlgoLibsRequest {
	s.CapabilityName = &v
	return s
}

func (s *DescribeAlgoLibsRequest) SetPageNumber(v int32) *DescribeAlgoLibsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAlgoLibsRequest) SetPageSize(v int32) *DescribeAlgoLibsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAlgoLibsRequest) SetInstanceId(v string) *DescribeAlgoLibsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAlgoLibsRequest) SetAlgoLibId(v string) *DescribeAlgoLibsRequest {
	s.AlgoLibId = &v
	return s
}

func (s *DescribeAlgoLibsRequest) SetRegionId(v string) *DescribeAlgoLibsRequest {
	s.RegionId = &v
	return s
}

type DescribeAlgoLibsResponseBody struct {
	AlgoLibs   *DescribeAlgoLibsResponseBodyAlgoLibs `json:"AlgoLibs,omitempty" xml:"AlgoLibs,omitempty" type:"Struct"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeAlgoLibsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlgoLibsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlgoLibsResponseBody) SetAlgoLibs(v *DescribeAlgoLibsResponseBodyAlgoLibs) *DescribeAlgoLibsResponseBody {
	s.AlgoLibs = v
	return s
}

func (s *DescribeAlgoLibsResponseBody) SetTotalCount(v int32) *DescribeAlgoLibsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAlgoLibsResponseBody) SetPageSize(v int32) *DescribeAlgoLibsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAlgoLibsResponseBody) SetRequestId(v string) *DescribeAlgoLibsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlgoLibsResponseBody) SetPageNumber(v int32) *DescribeAlgoLibsResponseBody {
	s.PageNumber = &v
	return s
}

type DescribeAlgoLibsResponseBodyAlgoLibs struct {
	AlgoLib []*DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib `json:"AlgoLib,omitempty" xml:"AlgoLib,omitempty" type:"Repeated"`
}

func (s DescribeAlgoLibsResponseBodyAlgoLibs) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlgoLibsResponseBodyAlgoLibs) GoString() string {
	return s.String()
}

func (s *DescribeAlgoLibsResponseBodyAlgoLibs) SetAlgoLib(v []*DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib) *DescribeAlgoLibsResponseBodyAlgoLibs {
	s.AlgoLib = v
	return s
}

type DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib struct {
	OssAccessKeyId  *string `json:"OssAccessKeyId,omitempty" xml:"OssAccessKeyId,omitempty"`
	Capacity        *string `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	AlgoLibId       *string `json:"AlgoLibId,omitempty" xml:"AlgoLibId,omitempty"`
	JsonText        *string `json:"JsonText,omitempty" xml:"JsonText,omitempty"`
	AlgoLibVersion  *string `json:"AlgoLibVersion,omitempty" xml:"AlgoLibVersion,omitempty"`
	AlgoLibName     *string `json:"AlgoLibName,omitempty" xml:"AlgoLibName,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
	OssBucket       *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	CapabilityNames *string `json:"CapabilityNames,omitempty" xml:"CapabilityNames,omitempty"`
	OssPath         *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	OssEndpoint     *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	UploadTime      *string `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
}

func (s DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib) GoString() string {
	return s.String()
}

func (s *DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib) SetOssAccessKeyId(v string) *DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib {
	s.OssAccessKeyId = &v
	return s
}

func (s *DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib) SetCapacity(v string) *DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib {
	s.Capacity = &v
	return s
}

func (s *DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib) SetAlgoLibId(v string) *DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib {
	s.AlgoLibId = &v
	return s
}

func (s *DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib) SetJsonText(v string) *DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib {
	s.JsonText = &v
	return s
}

func (s *DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib) SetAlgoLibVersion(v string) *DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib {
	s.AlgoLibVersion = &v
	return s
}

func (s *DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib) SetAlgoLibName(v string) *DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib {
	s.AlgoLibName = &v
	return s
}

func (s *DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib) SetText(v string) *DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib {
	s.Text = &v
	return s
}

func (s *DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib) SetOssBucket(v string) *DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib {
	s.OssBucket = &v
	return s
}

func (s *DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib) SetCapabilityNames(v string) *DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib {
	s.CapabilityNames = &v
	return s
}

func (s *DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib) SetOssPath(v string) *DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib {
	s.OssPath = &v
	return s
}

func (s *DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib) SetOssEndpoint(v string) *DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib {
	s.OssEndpoint = &v
	return s
}

func (s *DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib) SetUploadTime(v string) *DescribeAlgoLibsResponseBodyAlgoLibsAlgoLib {
	s.UploadTime = &v
	return s
}

type DescribeAlgoLibsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAlgoLibsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAlgoLibsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlgoLibsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlgoLibsResponse) SetHeaders(v map[string]*string) *DescribeAlgoLibsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlgoLibsResponse) SetBody(v *DescribeAlgoLibsResponseBody) *DescribeAlgoLibsResponse {
	s.Body = v
	return s
}

type DescribeCamerasRequest struct {
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CameraId     *string `json:"CameraId,omitempty" xml:"CameraId,omitempty"`
	CameraName   *string `json:"CameraName,omitempty" xml:"CameraName,omitempty"`
	StreamStatus *string `json:"StreamStatus,omitempty" xml:"StreamStatus,omitempty"`
	WorkGroupId  *string `json:"WorkGroupId,omitempty" xml:"WorkGroupId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCamerasRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCamerasRequest) GoString() string {
	return s.String()
}

func (s *DescribeCamerasRequest) SetPageNumber(v int32) *DescribeCamerasRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCamerasRequest) SetPageSize(v int32) *DescribeCamerasRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCamerasRequest) SetCameraId(v string) *DescribeCamerasRequest {
	s.CameraId = &v
	return s
}

func (s *DescribeCamerasRequest) SetCameraName(v string) *DescribeCamerasRequest {
	s.CameraName = &v
	return s
}

func (s *DescribeCamerasRequest) SetStreamStatus(v string) *DescribeCamerasRequest {
	s.StreamStatus = &v
	return s
}

func (s *DescribeCamerasRequest) SetWorkGroupId(v string) *DescribeCamerasRequest {
	s.WorkGroupId = &v
	return s
}

func (s *DescribeCamerasRequest) SetInstanceId(v string) *DescribeCamerasRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCamerasRequest) SetRegionId(v string) *DescribeCamerasRequest {
	s.RegionId = &v
	return s
}

type DescribeCamerasResponseBody struct {
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Cameras    *DescribeCamerasResponseBodyCameras `json:"Cameras,omitempty" xml:"Cameras,omitempty" type:"Struct"`
}

func (s DescribeCamerasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCamerasResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCamerasResponseBody) SetTotalCount(v int32) *DescribeCamerasResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCamerasResponseBody) SetRequestId(v string) *DescribeCamerasResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCamerasResponseBody) SetPageSize(v int32) *DescribeCamerasResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCamerasResponseBody) SetPageNumber(v int32) *DescribeCamerasResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCamerasResponseBody) SetCameras(v *DescribeCamerasResponseBodyCameras) *DescribeCamerasResponseBody {
	s.Cameras = v
	return s
}

type DescribeCamerasResponseBodyCameras struct {
	Camera []*DescribeCamerasResponseBodyCamerasCamera `json:"Camera,omitempty" xml:"Camera,omitempty" type:"Repeated"`
}

func (s DescribeCamerasResponseBodyCameras) String() string {
	return tea.Prettify(s)
}

func (s DescribeCamerasResponseBodyCameras) GoString() string {
	return s.String()
}

func (s *DescribeCamerasResponseBodyCameras) SetCamera(v []*DescribeCamerasResponseBodyCamerasCamera) *DescribeCamerasResponseBodyCameras {
	s.Camera = v
	return s
}

type DescribeCamerasResponseBodyCamerasCamera struct {
	StreamStatus *string `json:"StreamStatus,omitempty" xml:"StreamStatus,omitempty"`
	UpdateTime   *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	CameraId     *string `json:"CameraId,omitempty" xml:"CameraId,omitempty"`
	Conf         *string `json:"Conf,omitempty" xml:"Conf,omitempty"`
	RtmpPath     *string `json:"RtmpPath,omitempty" xml:"RtmpPath,omitempty"`
	InviteUri    *string `json:"InviteUri,omitempty" xml:"InviteUri,omitempty"`
	CameraName   *string `json:"CameraName,omitempty" xml:"CameraName,omitempty"`
	WorkGroupId  *string `json:"WorkGroupId,omitempty" xml:"WorkGroupId,omitempty"`
	Location     *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s DescribeCamerasResponseBodyCamerasCamera) String() string {
	return tea.Prettify(s)
}

func (s DescribeCamerasResponseBodyCamerasCamera) GoString() string {
	return s.String()
}

func (s *DescribeCamerasResponseBodyCamerasCamera) SetStreamStatus(v string) *DescribeCamerasResponseBodyCamerasCamera {
	s.StreamStatus = &v
	return s
}

func (s *DescribeCamerasResponseBodyCamerasCamera) SetUpdateTime(v string) *DescribeCamerasResponseBodyCamerasCamera {
	s.UpdateTime = &v
	return s
}

func (s *DescribeCamerasResponseBodyCamerasCamera) SetCameraId(v string) *DescribeCamerasResponseBodyCamerasCamera {
	s.CameraId = &v
	return s
}

func (s *DescribeCamerasResponseBodyCamerasCamera) SetConf(v string) *DescribeCamerasResponseBodyCamerasCamera {
	s.Conf = &v
	return s
}

func (s *DescribeCamerasResponseBodyCamerasCamera) SetRtmpPath(v string) *DescribeCamerasResponseBodyCamerasCamera {
	s.RtmpPath = &v
	return s
}

func (s *DescribeCamerasResponseBodyCamerasCamera) SetInviteUri(v string) *DescribeCamerasResponseBodyCamerasCamera {
	s.InviteUri = &v
	return s
}

func (s *DescribeCamerasResponseBodyCamerasCamera) SetCameraName(v string) *DescribeCamerasResponseBodyCamerasCamera {
	s.CameraName = &v
	return s
}

func (s *DescribeCamerasResponseBodyCamerasCamera) SetWorkGroupId(v string) *DescribeCamerasResponseBodyCamerasCamera {
	s.WorkGroupId = &v
	return s
}

func (s *DescribeCamerasResponseBodyCamerasCamera) SetLocation(v string) *DescribeCamerasResponseBodyCamerasCamera {
	s.Location = &v
	return s
}

type DescribeCamerasResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCamerasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCamerasResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCamerasResponse) GoString() string {
	return s.String()
}

func (s *DescribeCamerasResponse) SetHeaders(v map[string]*string) *DescribeCamerasResponse {
	s.Headers = v
	return s
}

func (s *DescribeCamerasResponse) SetBody(v *DescribeCamerasResponseBody) *DescribeCamerasResponse {
	s.Body = v
	return s
}

type DescribeCapabilitiesRequest struct {
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCapabilitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCapabilitiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCapabilitiesRequest) SetPageNumber(v int32) *DescribeCapabilitiesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCapabilitiesRequest) SetPageSize(v int32) *DescribeCapabilitiesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCapabilitiesRequest) SetInstanceId(v string) *DescribeCapabilitiesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCapabilitiesRequest) SetRegionId(v string) *DescribeCapabilitiesRequest {
	s.RegionId = &v
	return s
}

type DescribeCapabilitiesResponseBody struct {
	TotalCount   *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize     *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber   *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Capabilities *DescribeCapabilitiesResponseBodyCapabilities `json:"Capabilities,omitempty" xml:"Capabilities,omitempty" type:"Struct"`
}

func (s DescribeCapabilitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCapabilitiesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCapabilitiesResponseBody) SetTotalCount(v int32) *DescribeCapabilitiesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCapabilitiesResponseBody) SetPageSize(v int32) *DescribeCapabilitiesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCapabilitiesResponseBody) SetRequestId(v string) *DescribeCapabilitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCapabilitiesResponseBody) SetPageNumber(v int32) *DescribeCapabilitiesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCapabilitiesResponseBody) SetCapabilities(v *DescribeCapabilitiesResponseBodyCapabilities) *DescribeCapabilitiesResponseBody {
	s.Capabilities = v
	return s
}

type DescribeCapabilitiesResponseBodyCapabilities struct {
	Capability []*DescribeCapabilitiesResponseBodyCapabilitiesCapability `json:"Capability,omitempty" xml:"Capability,omitempty" type:"Repeated"`
}

func (s DescribeCapabilitiesResponseBodyCapabilities) String() string {
	return tea.Prettify(s)
}

func (s DescribeCapabilitiesResponseBodyCapabilities) GoString() string {
	return s.String()
}

func (s *DescribeCapabilitiesResponseBodyCapabilities) SetCapability(v []*DescribeCapabilitiesResponseBodyCapabilitiesCapability) *DescribeCapabilitiesResponseBodyCapabilities {
	s.Capability = v
	return s
}

type DescribeCapabilitiesResponseBodyCapabilitiesCapability struct {
	CapabilityId   *string `json:"CapabilityId,omitempty" xml:"CapabilityId,omitempty"`
	CapabilityName *string `json:"CapabilityName,omitempty" xml:"CapabilityName,omitempty"`
}

func (s DescribeCapabilitiesResponseBodyCapabilitiesCapability) String() string {
	return tea.Prettify(s)
}

func (s DescribeCapabilitiesResponseBodyCapabilitiesCapability) GoString() string {
	return s.String()
}

func (s *DescribeCapabilitiesResponseBodyCapabilitiesCapability) SetCapabilityId(v string) *DescribeCapabilitiesResponseBodyCapabilitiesCapability {
	s.CapabilityId = &v
	return s
}

func (s *DescribeCapabilitiesResponseBodyCapabilitiesCapability) SetCapabilityName(v string) *DescribeCapabilitiesResponseBodyCapabilitiesCapability {
	s.CapabilityName = &v
	return s
}

type DescribeCapabilitiesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCapabilitiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCapabilitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCapabilitiesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCapabilitiesResponse) SetHeaders(v map[string]*string) *DescribeCapabilitiesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCapabilitiesResponse) SetBody(v *DescribeCapabilitiesResponseBody) *DescribeCapabilitiesResponse {
	s.Body = v
	return s
}

type DescribeInstancesRequest struct {
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceIds  *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) SetInstanceName(v string) *DescribeInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesRequest) SetRegionId(v string) *DescribeInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceIds(v string) *DescribeInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageNumber(v int32) *DescribeInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageSize(v int32) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

type DescribeInstancesResponseBody struct {
	Instances  *DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) SetInstances(v *DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int32) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageSize(v int32) *DescribeInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageNumber(v int32) *DescribeInstancesResponseBody {
	s.PageNumber = &v
	return s
}

type DescribeInstancesResponseBodyInstances struct {
	Instance []*DescribeInstancesResponseBodyInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) SetInstance(v []*DescribeInstancesResponseBodyInstancesInstance) *DescribeInstancesResponseBodyInstances {
	s.Instance = v
	return s
}

type DescribeInstancesResponseBodyInstancesInstance struct {
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	InstanceCapacity *int32  `json:"InstanceCapacity,omitempty" xml:"InstanceCapacity,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	InstanceName     *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstance) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetStatus(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.Status = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInstanceCapacity(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.InstanceCapacity = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetDescription(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.Description = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCreateTime(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInstanceName(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInstanceId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetRegionId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.RegionId = &v
	return s
}

type DescribeInstancesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponse) SetHeaders(v map[string]*string) *DescribeInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstancesResponse) SetBody(v *DescribeInstancesResponseBody) *DescribeInstancesResponse {
	s.Body = v
	return s
}

type DescribeJobGroupsRequest struct {
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeJobGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobGroupsRequest) SetPageNumber(v int32) *DescribeJobGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeJobGroupsRequest) SetPageSize(v int32) *DescribeJobGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeJobGroupsRequest) SetInstanceId(v string) *DescribeJobGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeJobGroupsRequest) SetJobGroupId(v string) *DescribeJobGroupsRequest {
	s.JobGroupId = &v
	return s
}

func (s *DescribeJobGroupsRequest) SetRegionId(v string) *DescribeJobGroupsRequest {
	s.RegionId = &v
	return s
}

type DescribeJobGroupsResponseBody struct {
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	JobGroups  *DescribeJobGroupsResponseBodyJobGroups `json:"JobGroups,omitempty" xml:"JobGroups,omitempty" type:"Struct"`
}

func (s DescribeJobGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJobGroupsResponseBody) SetTotalCount(v int32) *DescribeJobGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeJobGroupsResponseBody) SetRequestId(v string) *DescribeJobGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeJobGroupsResponseBody) SetPageSize(v int32) *DescribeJobGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeJobGroupsResponseBody) SetPageNumber(v int32) *DescribeJobGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeJobGroupsResponseBody) SetJobGroups(v *DescribeJobGroupsResponseBodyJobGroups) *DescribeJobGroupsResponseBody {
	s.JobGroups = v
	return s
}

type DescribeJobGroupsResponseBodyJobGroups struct {
	JobGroup []*DescribeJobGroupsResponseBodyJobGroupsJobGroup `json:"JobGroup,omitempty" xml:"JobGroup,omitempty" type:"Repeated"`
}

func (s DescribeJobGroupsResponseBodyJobGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobGroupsResponseBodyJobGroups) GoString() string {
	return s.String()
}

func (s *DescribeJobGroupsResponseBodyJobGroups) SetJobGroup(v []*DescribeJobGroupsResponseBodyJobGroupsJobGroup) *DescribeJobGroupsResponseBodyJobGroups {
	s.JobGroup = v
	return s
}

type DescribeJobGroupsResponseBodyJobGroupsJobGroup struct {
	Status            *string                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	JobGroupId        *string                                                       `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	AlgoLibId         *string                                                       `json:"AlgoLibId,omitempty" xml:"AlgoLibId,omitempty"`
	JobCount          *int32                                                        `json:"JobCount,omitempty" xml:"JobCount,omitempty"`
	JobGroupName      *string                                                       `json:"JobGroupName,omitempty" xml:"JobGroupName,omitempty"`
	JobGroupParams    *DescribeJobGroupsResponseBodyJobGroupsJobGroupJobGroupParams `json:"JobGroupParams,omitempty" xml:"JobGroupParams,omitempty" type:"Struct"`
	ResourceProfileId *string                                                       `json:"ResourceProfileId,omitempty" xml:"ResourceProfileId,omitempty"`
	StreamPerJob      *int32                                                        `json:"StreamPerJob,omitempty" xml:"StreamPerJob,omitempty"`
	PlanedJobCount    *int32                                                        `json:"PlanedJobCount,omitempty" xml:"PlanedJobCount,omitempty"`
	ComputeJobs       *DescribeJobGroupsResponseBodyJobGroupsJobGroupComputeJobs    `json:"ComputeJobs,omitempty" xml:"ComputeJobs,omitempty" type:"Struct"`
}

func (s DescribeJobGroupsResponseBodyJobGroupsJobGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobGroupsResponseBodyJobGroupsJobGroup) GoString() string {
	return s.String()
}

func (s *DescribeJobGroupsResponseBodyJobGroupsJobGroup) SetStatus(v string) *DescribeJobGroupsResponseBodyJobGroupsJobGroup {
	s.Status = &v
	return s
}

func (s *DescribeJobGroupsResponseBodyJobGroupsJobGroup) SetJobGroupId(v string) *DescribeJobGroupsResponseBodyJobGroupsJobGroup {
	s.JobGroupId = &v
	return s
}

func (s *DescribeJobGroupsResponseBodyJobGroupsJobGroup) SetAlgoLibId(v string) *DescribeJobGroupsResponseBodyJobGroupsJobGroup {
	s.AlgoLibId = &v
	return s
}

func (s *DescribeJobGroupsResponseBodyJobGroupsJobGroup) SetJobCount(v int32) *DescribeJobGroupsResponseBodyJobGroupsJobGroup {
	s.JobCount = &v
	return s
}

func (s *DescribeJobGroupsResponseBodyJobGroupsJobGroup) SetJobGroupName(v string) *DescribeJobGroupsResponseBodyJobGroupsJobGroup {
	s.JobGroupName = &v
	return s
}

func (s *DescribeJobGroupsResponseBodyJobGroupsJobGroup) SetJobGroupParams(v *DescribeJobGroupsResponseBodyJobGroupsJobGroupJobGroupParams) *DescribeJobGroupsResponseBodyJobGroupsJobGroup {
	s.JobGroupParams = v
	return s
}

func (s *DescribeJobGroupsResponseBodyJobGroupsJobGroup) SetResourceProfileId(v string) *DescribeJobGroupsResponseBodyJobGroupsJobGroup {
	s.ResourceProfileId = &v
	return s
}

func (s *DescribeJobGroupsResponseBodyJobGroupsJobGroup) SetStreamPerJob(v int32) *DescribeJobGroupsResponseBodyJobGroupsJobGroup {
	s.StreamPerJob = &v
	return s
}

func (s *DescribeJobGroupsResponseBodyJobGroupsJobGroup) SetPlanedJobCount(v int32) *DescribeJobGroupsResponseBodyJobGroupsJobGroup {
	s.PlanedJobCount = &v
	return s
}

func (s *DescribeJobGroupsResponseBodyJobGroupsJobGroup) SetComputeJobs(v *DescribeJobGroupsResponseBodyJobGroupsJobGroupComputeJobs) *DescribeJobGroupsResponseBodyJobGroupsJobGroup {
	s.ComputeJobs = v
	return s
}

type DescribeJobGroupsResponseBodyJobGroupsJobGroupJobGroupParams struct {
	JobGroupParam []*DescribeJobGroupsResponseBodyJobGroupsJobGroupJobGroupParamsJobGroupParam `json:"JobGroupParam,omitempty" xml:"JobGroupParam,omitempty" type:"Repeated"`
}

func (s DescribeJobGroupsResponseBodyJobGroupsJobGroupJobGroupParams) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobGroupsResponseBodyJobGroupsJobGroupJobGroupParams) GoString() string {
	return s.String()
}

func (s *DescribeJobGroupsResponseBodyJobGroupsJobGroupJobGroupParams) SetJobGroupParam(v []*DescribeJobGroupsResponseBodyJobGroupsJobGroupJobGroupParamsJobGroupParam) *DescribeJobGroupsResponseBodyJobGroupsJobGroupJobGroupParams {
	s.JobGroupParam = v
	return s
}

type DescribeJobGroupsResponseBodyJobGroupsJobGroupJobGroupParamsJobGroupParam struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeJobGroupsResponseBodyJobGroupsJobGroupJobGroupParamsJobGroupParam) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobGroupsResponseBodyJobGroupsJobGroupJobGroupParamsJobGroupParam) GoString() string {
	return s.String()
}

func (s *DescribeJobGroupsResponseBodyJobGroupsJobGroupJobGroupParamsJobGroupParam) SetKey(v string) *DescribeJobGroupsResponseBodyJobGroupsJobGroupJobGroupParamsJobGroupParam {
	s.Key = &v
	return s
}

func (s *DescribeJobGroupsResponseBodyJobGroupsJobGroupJobGroupParamsJobGroupParam) SetValue(v string) *DescribeJobGroupsResponseBodyJobGroupsJobGroupJobGroupParamsJobGroupParam {
	s.Value = &v
	return s
}

type DescribeJobGroupsResponseBodyJobGroupsJobGroupComputeJobs struct {
	ComputeJob []*DescribeJobGroupsResponseBodyJobGroupsJobGroupComputeJobsComputeJob `json:"ComputeJob,omitempty" xml:"ComputeJob,omitempty" type:"Repeated"`
}

func (s DescribeJobGroupsResponseBodyJobGroupsJobGroupComputeJobs) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobGroupsResponseBodyJobGroupsJobGroupComputeJobs) GoString() string {
	return s.String()
}

func (s *DescribeJobGroupsResponseBodyJobGroupsJobGroupComputeJobs) SetComputeJob(v []*DescribeJobGroupsResponseBodyJobGroupsJobGroupComputeJobsComputeJob) *DescribeJobGroupsResponseBodyJobGroupsJobGroupComputeJobs {
	s.ComputeJob = v
	return s
}

type DescribeJobGroupsResponseBodyJobGroupsJobGroupComputeJobsComputeJob struct {
	EndTime          *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime        *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ComputeJobName   *string `json:"ComputeJobName,omitempty" xml:"ComputeJobName,omitempty"`
	ComputeJobId     *string `json:"ComputeJobId,omitempty" xml:"ComputeJobId,omitempty"`
	ComputeJobStatus *string `json:"ComputeJobStatus,omitempty" xml:"ComputeJobStatus,omitempty"`
	Duration         *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s DescribeJobGroupsResponseBodyJobGroupsJobGroupComputeJobsComputeJob) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobGroupsResponseBodyJobGroupsJobGroupComputeJobsComputeJob) GoString() string {
	return s.String()
}

func (s *DescribeJobGroupsResponseBodyJobGroupsJobGroupComputeJobsComputeJob) SetEndTime(v string) *DescribeJobGroupsResponseBodyJobGroupsJobGroupComputeJobsComputeJob {
	s.EndTime = &v
	return s
}

func (s *DescribeJobGroupsResponseBodyJobGroupsJobGroupComputeJobsComputeJob) SetStartTime(v string) *DescribeJobGroupsResponseBodyJobGroupsJobGroupComputeJobsComputeJob {
	s.StartTime = &v
	return s
}

func (s *DescribeJobGroupsResponseBodyJobGroupsJobGroupComputeJobsComputeJob) SetComputeJobName(v string) *DescribeJobGroupsResponseBodyJobGroupsJobGroupComputeJobsComputeJob {
	s.ComputeJobName = &v
	return s
}

func (s *DescribeJobGroupsResponseBodyJobGroupsJobGroupComputeJobsComputeJob) SetComputeJobId(v string) *DescribeJobGroupsResponseBodyJobGroupsJobGroupComputeJobsComputeJob {
	s.ComputeJobId = &v
	return s
}

func (s *DescribeJobGroupsResponseBodyJobGroupsJobGroupComputeJobsComputeJob) SetComputeJobStatus(v string) *DescribeJobGroupsResponseBodyJobGroupsJobGroupComputeJobsComputeJob {
	s.ComputeJobStatus = &v
	return s
}

func (s *DescribeJobGroupsResponseBodyJobGroupsJobGroupComputeJobsComputeJob) SetDuration(v int32) *DescribeJobGroupsResponseBodyJobGroupsJobGroupComputeJobsComputeJob {
	s.Duration = &v
	return s
}

type DescribeJobGroupsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeJobGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeJobGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeJobGroupsResponse) SetHeaders(v map[string]*string) *DescribeJobGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeJobGroupsResponse) SetBody(v *DescribeJobGroupsResponseBody) *DescribeJobGroupsResponse {
	s.Body = v
	return s
}

type DescribeProtocolsRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeProtocolsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtocolsRequest) GoString() string {
	return s.String()
}

func (s *DescribeProtocolsRequest) SetRegionId(v string) *DescribeProtocolsRequest {
	s.RegionId = &v
	return s
}

type DescribeProtocolsResponseBody struct {
	Protocols map[string]interface{} `json:"Protocols,omitempty" xml:"Protocols,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProtocolsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtocolsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProtocolsResponseBody) SetProtocols(v map[string]interface{}) *DescribeProtocolsResponseBody {
	s.Protocols = v
	return s
}

func (s *DescribeProtocolsResponseBody) SetRequestId(v string) *DescribeProtocolsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeProtocolsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeProtocolsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeProtocolsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtocolsResponse) GoString() string {
	return s.String()
}

func (s *DescribeProtocolsResponse) SetHeaders(v map[string]*string) *DescribeProtocolsResponse {
	s.Headers = v
	return s
}

func (s *DescribeProtocolsResponse) SetBody(v *DescribeProtocolsResponseBody) *DescribeProtocolsResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

type DescribeRegionsResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeResourceProfilesRequest struct {
	PageNumber        *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceProfileId *string `json:"ResourceProfileId,omitempty" xml:"ResourceProfileId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeResourceProfilesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceProfilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceProfilesRequest) SetPageNumber(v int32) *DescribeResourceProfilesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeResourceProfilesRequest) SetPageSize(v int32) *DescribeResourceProfilesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeResourceProfilesRequest) SetInstanceId(v string) *DescribeResourceProfilesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeResourceProfilesRequest) SetResourceProfileId(v string) *DescribeResourceProfilesRequest {
	s.ResourceProfileId = &v
	return s
}

func (s *DescribeResourceProfilesRequest) SetRegionId(v string) *DescribeResourceProfilesRequest {
	s.RegionId = &v
	return s
}

type DescribeResourceProfilesResponseBody struct {
	TotalCount       *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId        *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize         *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber       *int32                                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ResourceProfiles *DescribeResourceProfilesResponseBodyResourceProfiles `json:"ResourceProfiles,omitempty" xml:"ResourceProfiles,omitempty" type:"Struct"`
}

func (s DescribeResourceProfilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceProfilesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceProfilesResponseBody) SetTotalCount(v int32) *DescribeResourceProfilesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeResourceProfilesResponseBody) SetRequestId(v string) *DescribeResourceProfilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceProfilesResponseBody) SetPageSize(v int32) *DescribeResourceProfilesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeResourceProfilesResponseBody) SetPageNumber(v int32) *DescribeResourceProfilesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeResourceProfilesResponseBody) SetResourceProfiles(v *DescribeResourceProfilesResponseBodyResourceProfiles) *DescribeResourceProfilesResponseBody {
	s.ResourceProfiles = v
	return s
}

type DescribeResourceProfilesResponseBodyResourceProfiles struct {
	ResourceProfile []*DescribeResourceProfilesResponseBodyResourceProfilesResourceProfile `json:"ResourceProfile,omitempty" xml:"ResourceProfile,omitempty" type:"Repeated"`
}

func (s DescribeResourceProfilesResponseBodyResourceProfiles) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceProfilesResponseBodyResourceProfiles) GoString() string {
	return s.String()
}

func (s *DescribeResourceProfilesResponseBodyResourceProfiles) SetResourceProfile(v []*DescribeResourceProfilesResponseBodyResourceProfilesResourceProfile) *DescribeResourceProfilesResponseBodyResourceProfiles {
	s.ResourceProfile = v
	return s
}

type DescribeResourceProfilesResponseBodyResourceProfilesResourceProfile struct {
	ResourceProfileName   *string                                                                                   `json:"ResourceProfileName,omitempty" xml:"ResourceProfileName,omitempty"`
	ResourceProfileId     *string                                                                                   `json:"ResourceProfileId,omitempty" xml:"ResourceProfileId,omitempty"`
	ResourceProfileParams *DescribeResourceProfilesResponseBodyResourceProfilesResourceProfileResourceProfileParams `json:"ResourceProfileParams,omitempty" xml:"ResourceProfileParams,omitempty" type:"Struct"`
}

func (s DescribeResourceProfilesResponseBodyResourceProfilesResourceProfile) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceProfilesResponseBodyResourceProfilesResourceProfile) GoString() string {
	return s.String()
}

func (s *DescribeResourceProfilesResponseBodyResourceProfilesResourceProfile) SetResourceProfileName(v string) *DescribeResourceProfilesResponseBodyResourceProfilesResourceProfile {
	s.ResourceProfileName = &v
	return s
}

func (s *DescribeResourceProfilesResponseBodyResourceProfilesResourceProfile) SetResourceProfileId(v string) *DescribeResourceProfilesResponseBodyResourceProfilesResourceProfile {
	s.ResourceProfileId = &v
	return s
}

func (s *DescribeResourceProfilesResponseBodyResourceProfilesResourceProfile) SetResourceProfileParams(v *DescribeResourceProfilesResponseBodyResourceProfilesResourceProfileResourceProfileParams) *DescribeResourceProfilesResponseBodyResourceProfilesResourceProfile {
	s.ResourceProfileParams = v
	return s
}

type DescribeResourceProfilesResponseBodyResourceProfilesResourceProfileResourceProfileParams struct {
	ResourceProfileParam []*DescribeResourceProfilesResponseBodyResourceProfilesResourceProfileResourceProfileParamsResourceProfileParam `json:"ResourceProfileParam,omitempty" xml:"ResourceProfileParam,omitempty" type:"Repeated"`
}

func (s DescribeResourceProfilesResponseBodyResourceProfilesResourceProfileResourceProfileParams) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceProfilesResponseBodyResourceProfilesResourceProfileResourceProfileParams) GoString() string {
	return s.String()
}

func (s *DescribeResourceProfilesResponseBodyResourceProfilesResourceProfileResourceProfileParams) SetResourceProfileParam(v []*DescribeResourceProfilesResponseBodyResourceProfilesResourceProfileResourceProfileParamsResourceProfileParam) *DescribeResourceProfilesResponseBodyResourceProfilesResourceProfileResourceProfileParams {
	s.ResourceProfileParam = v
	return s
}

type DescribeResourceProfilesResponseBodyResourceProfilesResourceProfileResourceProfileParamsResourceProfileParam struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeResourceProfilesResponseBodyResourceProfilesResourceProfileResourceProfileParamsResourceProfileParam) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceProfilesResponseBodyResourceProfilesResourceProfileResourceProfileParamsResourceProfileParam) GoString() string {
	return s.String()
}

func (s *DescribeResourceProfilesResponseBodyResourceProfilesResourceProfileResourceProfileParamsResourceProfileParam) SetKey(v string) *DescribeResourceProfilesResponseBodyResourceProfilesResourceProfileResourceProfileParamsResourceProfileParam {
	s.Key = &v
	return s
}

func (s *DescribeResourceProfilesResponseBodyResourceProfilesResourceProfileResourceProfileParamsResourceProfileParam) SetValue(v string) *DescribeResourceProfilesResponseBodyResourceProfilesResourceProfileResourceProfileParamsResourceProfileParam {
	s.Value = &v
	return s
}

type DescribeResourceProfilesResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeResourceProfilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeResourceProfilesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceProfilesResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceProfilesResponse) SetHeaders(v map[string]*string) *DescribeResourceProfilesResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceProfilesResponse) SetBody(v *DescribeResourceProfilesResponseBody) *DescribeResourceProfilesResponse {
	s.Body = v
	return s
}

type DescribeStreamsRequest struct {
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeStreamsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStreamsRequest) GoString() string {
	return s.String()
}

func (s *DescribeStreamsRequest) SetJobGroupId(v string) *DescribeStreamsRequest {
	s.JobGroupId = &v
	return s
}

func (s *DescribeStreamsRequest) SetPageNumber(v int32) *DescribeStreamsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeStreamsRequest) SetPageSize(v int32) *DescribeStreamsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeStreamsRequest) SetRegionId(v string) *DescribeStreamsRequest {
	s.RegionId = &v
	return s
}

type DescribeStreamsResponseBody struct {
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Streams    *DescribeStreamsResponseBodyStreams `json:"Streams,omitempty" xml:"Streams,omitempty" type:"Struct"`
}

func (s DescribeStreamsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStreamsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStreamsResponseBody) SetTotalCount(v int32) *DescribeStreamsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeStreamsResponseBody) SetPageSize(v int32) *DescribeStreamsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeStreamsResponseBody) SetRequestId(v string) *DescribeStreamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStreamsResponseBody) SetPageNumber(v int32) *DescribeStreamsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeStreamsResponseBody) SetStreams(v *DescribeStreamsResponseBodyStreams) *DescribeStreamsResponseBody {
	s.Streams = v
	return s
}

type DescribeStreamsResponseBodyStreams struct {
	Stream []*DescribeStreamsResponseBodyStreamsStream `json:"Stream,omitempty" xml:"Stream,omitempty" type:"Repeated"`
}

func (s DescribeStreamsResponseBodyStreams) String() string {
	return tea.Prettify(s)
}

func (s DescribeStreamsResponseBodyStreams) GoString() string {
	return s.String()
}

func (s *DescribeStreamsResponseBodyStreams) SetStream(v []*DescribeStreamsResponseBodyStreamsStream) *DescribeStreamsResponseBodyStreams {
	s.Stream = v
	return s
}

type DescribeStreamsResponseBodyStreamsStream struct {
	JobGroupId   *int32  `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	StreamName   *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	ComputeJobId *string `json:"ComputeJobId,omitempty" xml:"ComputeJobId,omitempty"`
	ObjCount     *int32  `json:"ObjCount,omitempty" xml:"ObjCount,omitempty"`
	DelayMS      *int32  `json:"DelayMS,omitempty" xml:"DelayMS,omitempty"`
	StreamURL    *string `json:"StreamURL,omitempty" xml:"StreamURL,omitempty"`
}

func (s DescribeStreamsResponseBodyStreamsStream) String() string {
	return tea.Prettify(s)
}

func (s DescribeStreamsResponseBodyStreamsStream) GoString() string {
	return s.String()
}

func (s *DescribeStreamsResponseBodyStreamsStream) SetJobGroupId(v int32) *DescribeStreamsResponseBodyStreamsStream {
	s.JobGroupId = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreamsStream) SetStreamName(v string) *DescribeStreamsResponseBodyStreamsStream {
	s.StreamName = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreamsStream) SetComputeJobId(v string) *DescribeStreamsResponseBodyStreamsStream {
	s.ComputeJobId = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreamsStream) SetObjCount(v int32) *DescribeStreamsResponseBodyStreamsStream {
	s.ObjCount = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreamsStream) SetDelayMS(v int32) *DescribeStreamsResponseBodyStreamsStream {
	s.DelayMS = &v
	return s
}

func (s *DescribeStreamsResponseBodyStreamsStream) SetStreamURL(v string) *DescribeStreamsResponseBodyStreamsStream {
	s.StreamURL = &v
	return s
}

type DescribeStreamsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeStreamsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeStreamsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStreamsResponse) GoString() string {
	return s.String()
}

func (s *DescribeStreamsResponse) SetHeaders(v map[string]*string) *DescribeStreamsResponse {
	s.Headers = v
	return s
}

func (s *DescribeStreamsResponse) SetBody(v *DescribeStreamsResponseBody) *DescribeStreamsResponse {
	s.Body = v
	return s
}

type DescribeWorkGroupsRequest struct {
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	WorkGroupId *string `json:"WorkGroupId,omitempty" xml:"WorkGroupId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeWorkGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWorkGroupsRequest) SetPageNumber(v int32) *DescribeWorkGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeWorkGroupsRequest) SetPageSize(v int32) *DescribeWorkGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWorkGroupsRequest) SetInstanceId(v string) *DescribeWorkGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeWorkGroupsRequest) SetWorkGroupId(v string) *DescribeWorkGroupsRequest {
	s.WorkGroupId = &v
	return s
}

func (s *DescribeWorkGroupsRequest) SetRegionId(v string) *DescribeWorkGroupsRequest {
	s.RegionId = &v
	return s
}

type DescribeWorkGroupsResponseBody struct {
	WorkGroups *DescribeWorkGroupsResponseBodyWorkGroups `json:"WorkGroups,omitempty" xml:"WorkGroups,omitempty" type:"Struct"`
	TotalCount *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeWorkGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWorkGroupsResponseBody) SetWorkGroups(v *DescribeWorkGroupsResponseBodyWorkGroups) *DescribeWorkGroupsResponseBody {
	s.WorkGroups = v
	return s
}

func (s *DescribeWorkGroupsResponseBody) SetTotalCount(v int32) *DescribeWorkGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWorkGroupsResponseBody) SetRequestId(v string) *DescribeWorkGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWorkGroupsResponseBody) SetPageSize(v int32) *DescribeWorkGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeWorkGroupsResponseBody) SetPageNumber(v int32) *DescribeWorkGroupsResponseBody {
	s.PageNumber = &v
	return s
}

type DescribeWorkGroupsResponseBodyWorkGroups struct {
	WorkGroup []*DescribeWorkGroupsResponseBodyWorkGroupsWorkGroup `json:"WorkGroup,omitempty" xml:"WorkGroup,omitempty" type:"Repeated"`
}

func (s DescribeWorkGroupsResponseBodyWorkGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkGroupsResponseBodyWorkGroups) GoString() string {
	return s.String()
}

func (s *DescribeWorkGroupsResponseBodyWorkGroups) SetWorkGroup(v []*DescribeWorkGroupsResponseBodyWorkGroupsWorkGroup) *DescribeWorkGroupsResponseBodyWorkGroups {
	s.WorkGroup = v
	return s
}

type DescribeWorkGroupsResponseBodyWorkGroupsWorkGroup struct {
	UpdateTime      *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Protocol        *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TheoryCutStatus *string `json:"TheoryCutStatus,omitempty" xml:"TheoryCutStatus,omitempty"`
	WorkGroupId     *string `json:"WorkGroupId,omitempty" xml:"WorkGroupId,omitempty"`
	WorkGroupName   *string `json:"WorkGroupName,omitempty" xml:"WorkGroupName,omitempty"`
}

func (s DescribeWorkGroupsResponseBodyWorkGroupsWorkGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkGroupsResponseBodyWorkGroupsWorkGroup) GoString() string {
	return s.String()
}

func (s *DescribeWorkGroupsResponseBodyWorkGroupsWorkGroup) SetUpdateTime(v string) *DescribeWorkGroupsResponseBodyWorkGroupsWorkGroup {
	s.UpdateTime = &v
	return s
}

func (s *DescribeWorkGroupsResponseBodyWorkGroupsWorkGroup) SetDescription(v string) *DescribeWorkGroupsResponseBodyWorkGroupsWorkGroup {
	s.Description = &v
	return s
}

func (s *DescribeWorkGroupsResponseBodyWorkGroupsWorkGroup) SetProtocol(v string) *DescribeWorkGroupsResponseBodyWorkGroupsWorkGroup {
	s.Protocol = &v
	return s
}

func (s *DescribeWorkGroupsResponseBodyWorkGroupsWorkGroup) SetCreateTime(v string) *DescribeWorkGroupsResponseBodyWorkGroupsWorkGroup {
	s.CreateTime = &v
	return s
}

func (s *DescribeWorkGroupsResponseBodyWorkGroupsWorkGroup) SetInstanceId(v string) *DescribeWorkGroupsResponseBodyWorkGroupsWorkGroup {
	s.InstanceId = &v
	return s
}

func (s *DescribeWorkGroupsResponseBodyWorkGroupsWorkGroup) SetTheoryCutStatus(v string) *DescribeWorkGroupsResponseBodyWorkGroupsWorkGroup {
	s.TheoryCutStatus = &v
	return s
}

func (s *DescribeWorkGroupsResponseBodyWorkGroupsWorkGroup) SetWorkGroupId(v string) *DescribeWorkGroupsResponseBodyWorkGroupsWorkGroup {
	s.WorkGroupId = &v
	return s
}

func (s *DescribeWorkGroupsResponseBodyWorkGroupsWorkGroup) SetWorkGroupName(v string) *DescribeWorkGroupsResponseBodyWorkGroupsWorkGroup {
	s.WorkGroupName = &v
	return s
}

type DescribeWorkGroupsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeWorkGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWorkGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWorkGroupsResponse) SetHeaders(v map[string]*string) *DescribeWorkGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWorkGroupsResponse) SetBody(v *DescribeWorkGroupsResponseBody) *DescribeWorkGroupsResponse {
	s.Body = v
	return s
}

type DetachStreamRequest struct {
	JobGroupId *string                       `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	RegionId   *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Streams    []*DetachStreamRequestStreams `json:"Streams,omitempty" xml:"Streams,omitempty" type:"Repeated"`
}

func (s DetachStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachStreamRequest) GoString() string {
	return s.String()
}

func (s *DetachStreamRequest) SetJobGroupId(v string) *DetachStreamRequest {
	s.JobGroupId = &v
	return s
}

func (s *DetachStreamRequest) SetRegionId(v string) *DetachStreamRequest {
	s.RegionId = &v
	return s
}

func (s *DetachStreamRequest) SetStreams(v []*DetachStreamRequestStreams) *DetachStreamRequest {
	s.Streams = v
	return s
}

type DetachStreamRequestStreams struct {
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DetachStreamRequestStreams) String() string {
	return tea.Prettify(s)
}

func (s DetachStreamRequestStreams) GoString() string {
	return s.String()
}

func (s *DetachStreamRequestStreams) SetStreamName(v string) *DetachStreamRequestStreams {
	s.StreamName = &v
	return s
}

type DetachStreamResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachStreamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachStreamResponseBody) GoString() string {
	return s.String()
}

func (s *DetachStreamResponseBody) SetRequestId(v string) *DetachStreamResponseBody {
	s.RequestId = &v
	return s
}

type DetachStreamResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetachStreamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachStreamResponse) GoString() string {
	return s.String()
}

func (s *DetachStreamResponse) SetHeaders(v map[string]*string) *DetachStreamResponse {
	s.Headers = v
	return s
}

func (s *DetachStreamResponse) SetBody(v *DetachStreamResponseBody) *DetachStreamResponse {
	s.Body = v
	return s
}

type GetCameraConfForCameraRequest struct {
	CameraId   *string `json:"CameraId,omitempty" xml:"CameraId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetCameraConfForCameraRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCameraConfForCameraRequest) GoString() string {
	return s.String()
}

func (s *GetCameraConfForCameraRequest) SetCameraId(v string) *GetCameraConfForCameraRequest {
	s.CameraId = &v
	return s
}

func (s *GetCameraConfForCameraRequest) SetInstanceId(v string) *GetCameraConfForCameraRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCameraConfForCameraRequest) SetRegionId(v string) *GetCameraConfForCameraRequest {
	s.RegionId = &v
	return s
}

type GetCameraConfForCameraResponseBody struct {
	CameraConf *GetCameraConfForCameraResponseBodyCameraConf `json:"CameraConf,omitempty" xml:"CameraConf,omitempty" type:"Struct"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCameraConfForCameraResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCameraConfForCameraResponseBody) GoString() string {
	return s.String()
}

func (s *GetCameraConfForCameraResponseBody) SetCameraConf(v *GetCameraConfForCameraResponseBodyCameraConf) *GetCameraConfForCameraResponseBody {
	s.CameraConf = v
	return s
}

func (s *GetCameraConfForCameraResponseBody) SetRequestId(v string) *GetCameraConfForCameraResponseBody {
	s.RequestId = &v
	return s
}

type GetCameraConfForCameraResponseBodyCameraConf struct {
	Context    *string `json:"Context,omitempty" xml:"Context,omitempty"`
	CameraId   *string `json:"CameraId,omitempty" xml:"CameraId,omitempty"`
	CameraName *string `json:"CameraName,omitempty" xml:"CameraName,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetCameraConfForCameraResponseBodyCameraConf) String() string {
	return tea.Prettify(s)
}

func (s GetCameraConfForCameraResponseBodyCameraConf) GoString() string {
	return s.String()
}

func (s *GetCameraConfForCameraResponseBodyCameraConf) SetContext(v string) *GetCameraConfForCameraResponseBodyCameraConf {
	s.Context = &v
	return s
}

func (s *GetCameraConfForCameraResponseBodyCameraConf) SetCameraId(v string) *GetCameraConfForCameraResponseBodyCameraConf {
	s.CameraId = &v
	return s
}

func (s *GetCameraConfForCameraResponseBodyCameraConf) SetCameraName(v string) *GetCameraConfForCameraResponseBodyCameraConf {
	s.CameraName = &v
	return s
}

func (s *GetCameraConfForCameraResponseBodyCameraConf) SetId(v string) *GetCameraConfForCameraResponseBodyCameraConf {
	s.Id = &v
	return s
}

type GetCameraConfForCameraResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCameraConfForCameraResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCameraConfForCameraResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCameraConfForCameraResponse) GoString() string {
	return s.String()
}

func (s *GetCameraConfForCameraResponse) SetHeaders(v map[string]*string) *GetCameraConfForCameraResponse {
	s.Headers = v
	return s
}

func (s *GetCameraConfForCameraResponse) SetBody(v *GetCameraConfForCameraResponseBody) *GetCameraConfForCameraResponse {
	s.Body = v
	return s
}

type GetComputeJobLogRequest struct {
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobLogName *string `json:"JobLogName,omitempty" xml:"JobLogName,omitempty"`
	Size       *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetComputeJobLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetComputeJobLogRequest) GoString() string {
	return s.String()
}

func (s *GetComputeJobLogRequest) SetJobGroupId(v string) *GetComputeJobLogRequest {
	s.JobGroupId = &v
	return s
}

func (s *GetComputeJobLogRequest) SetJobId(v string) *GetComputeJobLogRequest {
	s.JobId = &v
	return s
}

func (s *GetComputeJobLogRequest) SetJobLogName(v string) *GetComputeJobLogRequest {
	s.JobLogName = &v
	return s
}

func (s *GetComputeJobLogRequest) SetSize(v int32) *GetComputeJobLogRequest {
	s.Size = &v
	return s
}

func (s *GetComputeJobLogRequest) SetRegionId(v string) *GetComputeJobLogRequest {
	s.RegionId = &v
	return s
}

type GetComputeJobLogResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	JobLog    *string `json:"JobLog,omitempty" xml:"JobLog,omitempty"`
}

func (s GetComputeJobLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetComputeJobLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetComputeJobLogResponseBody) SetRequestId(v string) *GetComputeJobLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetComputeJobLogResponseBody) SetJobLog(v string) *GetComputeJobLogResponseBody {
	s.JobLog = &v
	return s
}

type GetComputeJobLogResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetComputeJobLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetComputeJobLogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetComputeJobLogResponse) GoString() string {
	return s.String()
}

func (s *GetComputeJobLogResponse) SetHeaders(v map[string]*string) *GetComputeJobLogResponse {
	s.Headers = v
	return s
}

func (s *GetComputeJobLogResponse) SetBody(v *GetComputeJobLogResponseBody) *GetComputeJobLogResponse {
	s.Body = v
	return s
}

type GetPlayUrlForCameraRequest struct {
	CameraId   *string `json:"CameraId,omitempty" xml:"CameraId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetPlayUrlForCameraRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPlayUrlForCameraRequest) GoString() string {
	return s.String()
}

func (s *GetPlayUrlForCameraRequest) SetCameraId(v string) *GetPlayUrlForCameraRequest {
	s.CameraId = &v
	return s
}

func (s *GetPlayUrlForCameraRequest) SetInstanceId(v string) *GetPlayUrlForCameraRequest {
	s.InstanceId = &v
	return s
}

func (s *GetPlayUrlForCameraRequest) SetRegionId(v string) *GetPlayUrlForCameraRequest {
	s.RegionId = &v
	return s
}

type GetPlayUrlForCameraResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PlayUrl   *string `json:"PlayUrl,omitempty" xml:"PlayUrl,omitempty"`
}

func (s GetPlayUrlForCameraResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPlayUrlForCameraResponseBody) GoString() string {
	return s.String()
}

func (s *GetPlayUrlForCameraResponseBody) SetRequestId(v string) *GetPlayUrlForCameraResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPlayUrlForCameraResponseBody) SetPlayUrl(v string) *GetPlayUrlForCameraResponseBody {
	s.PlayUrl = &v
	return s
}

type GetPlayUrlForCameraResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPlayUrlForCameraResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPlayUrlForCameraResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPlayUrlForCameraResponse) GoString() string {
	return s.String()
}

func (s *GetPlayUrlForCameraResponse) SetHeaders(v map[string]*string) *GetPlayUrlForCameraResponse {
	s.Headers = v
	return s
}

func (s *GetPlayUrlForCameraResponse) SetBody(v *GetPlayUrlForCameraResponseBody) *GetPlayUrlForCameraResponse {
	s.Body = v
	return s
}

type ListComputeJobLogsRequest struct {
	JobGroupId   *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	ComputeJobId *string `json:"ComputeJobId,omitempty" xml:"ComputeJobId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListComputeJobLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListComputeJobLogsRequest) GoString() string {
	return s.String()
}

func (s *ListComputeJobLogsRequest) SetJobGroupId(v string) *ListComputeJobLogsRequest {
	s.JobGroupId = &v
	return s
}

func (s *ListComputeJobLogsRequest) SetComputeJobId(v string) *ListComputeJobLogsRequest {
	s.ComputeJobId = &v
	return s
}

func (s *ListComputeJobLogsRequest) SetRegionId(v string) *ListComputeJobLogsRequest {
	s.RegionId = &v
	return s
}

type ListComputeJobLogsResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	JobLogs   *ListComputeJobLogsResponseBodyJobLogs `json:"JobLogs,omitempty" xml:"JobLogs,omitempty" type:"Struct"`
}

func (s ListComputeJobLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListComputeJobLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListComputeJobLogsResponseBody) SetRequestId(v string) *ListComputeJobLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListComputeJobLogsResponseBody) SetJobLogs(v *ListComputeJobLogsResponseBodyJobLogs) *ListComputeJobLogsResponseBody {
	s.JobLogs = v
	return s
}

type ListComputeJobLogsResponseBodyJobLogs struct {
	JobLog []*string `json:"JobLog,omitempty" xml:"JobLog,omitempty" type:"Repeated"`
}

func (s ListComputeJobLogsResponseBodyJobLogs) String() string {
	return tea.Prettify(s)
}

func (s ListComputeJobLogsResponseBodyJobLogs) GoString() string {
	return s.String()
}

func (s *ListComputeJobLogsResponseBodyJobLogs) SetJobLog(v []*string) *ListComputeJobLogsResponseBodyJobLogs {
	s.JobLog = v
	return s
}

type ListComputeJobLogsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListComputeJobLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListComputeJobLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListComputeJobLogsResponse) GoString() string {
	return s.String()
}

func (s *ListComputeJobLogsResponse) SetHeaders(v map[string]*string) *ListComputeJobLogsResponse {
	s.Headers = v
	return s
}

func (s *ListComputeJobLogsResponse) SetBody(v *ListComputeJobLogsResponseBody) *ListComputeJobLogsResponse {
	s.Body = v
	return s
}

type ListParkingResultsRequest struct {
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	CameraIds    *string `json:"CameraIds,omitempty" xml:"CameraIds,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VehiclePlate *string `json:"VehiclePlate,omitempty" xml:"VehiclePlate,omitempty"`
}

func (s ListParkingResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListParkingResultsRequest) GoString() string {
	return s.String()
}

func (s *ListParkingResultsRequest) SetStartTime(v string) *ListParkingResultsRequest {
	s.StartTime = &v
	return s
}

func (s *ListParkingResultsRequest) SetEndTime(v string) *ListParkingResultsRequest {
	s.EndTime = &v
	return s
}

func (s *ListParkingResultsRequest) SetCameraIds(v string) *ListParkingResultsRequest {
	s.CameraIds = &v
	return s
}

func (s *ListParkingResultsRequest) SetInstanceId(v string) *ListParkingResultsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListParkingResultsRequest) SetRegionId(v string) *ListParkingResultsRequest {
	s.RegionId = &v
	return s
}

func (s *ListParkingResultsRequest) SetVehiclePlate(v string) *ListParkingResultsRequest {
	s.VehiclePlate = &v
	return s
}

type ListParkingResultsResponseBody struct {
	TotalCount *int64                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int64                                   `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Results    []*ListParkingResultsResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s ListParkingResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListParkingResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListParkingResultsResponseBody) SetTotalCount(v int64) *ListParkingResultsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListParkingResultsResponseBody) SetTotalPage(v int64) *ListParkingResultsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListParkingResultsResponseBody) SetPageSize(v int32) *ListParkingResultsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListParkingResultsResponseBody) SetRequestId(v string) *ListParkingResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListParkingResultsResponseBody) SetPageNumber(v int32) *ListParkingResultsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListParkingResultsResponseBody) SetResults(v []*ListParkingResultsResponseBodyResults) *ListParkingResultsResponseBody {
	s.Results = v
	return s
}

type ListParkingResultsResponseBodyResults struct {
	VehicleType  *string `json:"VehicleType,omitempty" xml:"VehicleType,omitempty"`
	CameraId     *string `json:"CameraId,omitempty" xml:"CameraId,omitempty"`
	ObjRight     *int32  `json:"ObjRight,omitempty" xml:"ObjRight,omitempty"`
	ObjType      *string `json:"ObjType,omitempty" xml:"ObjType,omitempty"`
	Feature      *string `json:"Feature,omitempty" xml:"Feature,omitempty"`
	ObjLeft      *int32  `json:"ObjLeft,omitempty" xml:"ObjLeft,omitempty"`
	CropImage    *string `json:"CropImage,omitempty" xml:"CropImage,omitempty"`
	ObjTop       *int32  `json:"ObjTop,omitempty" xml:"ObjTop,omitempty"`
	Label        *string `json:"Label,omitempty" xml:"Label,omitempty"`
	LeaveTime    *string `json:"LeaveTime,omitempty" xml:"LeaveTime,omitempty"`
	ObjBottom    *int32  `json:"ObjBottom,omitempty" xml:"ObjBottom,omitempty"`
	FunctionType *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	OrigImage    *string `json:"OrigImage,omitempty" xml:"OrigImage,omitempty"`
	TimeStamp    *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	VehiclePlate *string `json:"VehiclePlate,omitempty" xml:"VehiclePlate,omitempty"`
	EntryTime    *string `json:"EntryTime,omitempty" xml:"EntryTime,omitempty"`
	ImageId      *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s ListParkingResultsResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s ListParkingResultsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *ListParkingResultsResponseBodyResults) SetVehicleType(v string) *ListParkingResultsResponseBodyResults {
	s.VehicleType = &v
	return s
}

func (s *ListParkingResultsResponseBodyResults) SetCameraId(v string) *ListParkingResultsResponseBodyResults {
	s.CameraId = &v
	return s
}

func (s *ListParkingResultsResponseBodyResults) SetObjRight(v int32) *ListParkingResultsResponseBodyResults {
	s.ObjRight = &v
	return s
}

func (s *ListParkingResultsResponseBodyResults) SetObjType(v string) *ListParkingResultsResponseBodyResults {
	s.ObjType = &v
	return s
}

func (s *ListParkingResultsResponseBodyResults) SetFeature(v string) *ListParkingResultsResponseBodyResults {
	s.Feature = &v
	return s
}

func (s *ListParkingResultsResponseBodyResults) SetObjLeft(v int32) *ListParkingResultsResponseBodyResults {
	s.ObjLeft = &v
	return s
}

func (s *ListParkingResultsResponseBodyResults) SetCropImage(v string) *ListParkingResultsResponseBodyResults {
	s.CropImage = &v
	return s
}

func (s *ListParkingResultsResponseBodyResults) SetObjTop(v int32) *ListParkingResultsResponseBodyResults {
	s.ObjTop = &v
	return s
}

func (s *ListParkingResultsResponseBodyResults) SetLabel(v string) *ListParkingResultsResponseBodyResults {
	s.Label = &v
	return s
}

func (s *ListParkingResultsResponseBodyResults) SetLeaveTime(v string) *ListParkingResultsResponseBodyResults {
	s.LeaveTime = &v
	return s
}

func (s *ListParkingResultsResponseBodyResults) SetObjBottom(v int32) *ListParkingResultsResponseBodyResults {
	s.ObjBottom = &v
	return s
}

func (s *ListParkingResultsResponseBodyResults) SetFunctionType(v string) *ListParkingResultsResponseBodyResults {
	s.FunctionType = &v
	return s
}

func (s *ListParkingResultsResponseBodyResults) SetOrigImage(v string) *ListParkingResultsResponseBodyResults {
	s.OrigImage = &v
	return s
}

func (s *ListParkingResultsResponseBodyResults) SetTimeStamp(v string) *ListParkingResultsResponseBodyResults {
	s.TimeStamp = &v
	return s
}

func (s *ListParkingResultsResponseBodyResults) SetVehiclePlate(v string) *ListParkingResultsResponseBodyResults {
	s.VehiclePlate = &v
	return s
}

func (s *ListParkingResultsResponseBodyResults) SetEntryTime(v string) *ListParkingResultsResponseBodyResults {
	s.EntryTime = &v
	return s
}

func (s *ListParkingResultsResponseBodyResults) SetImageId(v string) *ListParkingResultsResponseBodyResults {
	s.ImageId = &v
	return s
}

type ListParkingResultsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListParkingResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListParkingResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListParkingResultsResponse) GoString() string {
	return s.String()
}

func (s *ListParkingResultsResponse) SetHeaders(v map[string]*string) *ListParkingResultsResponse {
	s.Headers = v
	return s
}

func (s *ListParkingResultsResponse) SetBody(v *ListParkingResultsResponseBody) *ListParkingResultsResponse {
	s.Body = v
	return s
}

type ListSafetyHelmetResultsRequest struct {
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	CameraIds  *string `json:"CameraIds,omitempty" xml:"CameraIds,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListSafetyHelmetResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSafetyHelmetResultsRequest) GoString() string {
	return s.String()
}

func (s *ListSafetyHelmetResultsRequest) SetStartTime(v string) *ListSafetyHelmetResultsRequest {
	s.StartTime = &v
	return s
}

func (s *ListSafetyHelmetResultsRequest) SetEndTime(v string) *ListSafetyHelmetResultsRequest {
	s.EndTime = &v
	return s
}

func (s *ListSafetyHelmetResultsRequest) SetCameraIds(v string) *ListSafetyHelmetResultsRequest {
	s.CameraIds = &v
	return s
}

func (s *ListSafetyHelmetResultsRequest) SetInstanceId(v string) *ListSafetyHelmetResultsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSafetyHelmetResultsRequest) SetRegionId(v string) *ListSafetyHelmetResultsRequest {
	s.RegionId = &v
	return s
}

type ListSafetyHelmetResultsResponseBody struct {
	TotalCount *int64                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int64                                        `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Results    []*ListSafetyHelmetResultsResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	PageIndex  *int32                                        `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
}

func (s ListSafetyHelmetResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSafetyHelmetResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSafetyHelmetResultsResponseBody) SetTotalCount(v int64) *ListSafetyHelmetResultsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSafetyHelmetResultsResponseBody) SetTotalPage(v int64) *ListSafetyHelmetResultsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListSafetyHelmetResultsResponseBody) SetRequestId(v string) *ListSafetyHelmetResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSafetyHelmetResultsResponseBody) SetPageNumber(v int32) *ListSafetyHelmetResultsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSafetyHelmetResultsResponseBody) SetResults(v []*ListSafetyHelmetResultsResponseBodyResults) *ListSafetyHelmetResultsResponseBody {
	s.Results = v
	return s
}

func (s *ListSafetyHelmetResultsResponseBody) SetPageIndex(v int32) *ListSafetyHelmetResultsResponseBody {
	s.PageIndex = &v
	return s
}

type ListSafetyHelmetResultsResponseBodyResults struct {
	CameraId  *string `json:"CameraId,omitempty" xml:"CameraId,omitempty"`
	ObjType   *string `json:"ObjType,omitempty" xml:"ObjType,omitempty"`
	ObjRight  *int32  `json:"ObjRight,omitempty" xml:"ObjRight,omitempty"`
	Feature   *string `json:"Feature,omitempty" xml:"Feature,omitempty"`
	ObjLeft   *int32  `json:"ObjLeft,omitempty" xml:"ObjLeft,omitempty"`
	CropImage *string `json:"CropImage,omitempty" xml:"CropImage,omitempty"`
	ObjTop    *int32  `json:"ObjTop,omitempty" xml:"ObjTop,omitempty"`
	Label     *string `json:"Label,omitempty" xml:"Label,omitempty"`
	LeaveTime *string `json:"LeaveTime,omitempty" xml:"LeaveTime,omitempty"`
	ObjBottom *int32  `json:"ObjBottom,omitempty" xml:"ObjBottom,omitempty"`
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	OrigImage *string `json:"OrigImage,omitempty" xml:"OrigImage,omitempty"`
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	EntryTime *string `json:"EntryTime,omitempty" xml:"EntryTime,omitempty"`
}

func (s ListSafetyHelmetResultsResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s ListSafetyHelmetResultsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *ListSafetyHelmetResultsResponseBodyResults) SetCameraId(v string) *ListSafetyHelmetResultsResponseBodyResults {
	s.CameraId = &v
	return s
}

func (s *ListSafetyHelmetResultsResponseBodyResults) SetObjType(v string) *ListSafetyHelmetResultsResponseBodyResults {
	s.ObjType = &v
	return s
}

func (s *ListSafetyHelmetResultsResponseBodyResults) SetObjRight(v int32) *ListSafetyHelmetResultsResponseBodyResults {
	s.ObjRight = &v
	return s
}

func (s *ListSafetyHelmetResultsResponseBodyResults) SetFeature(v string) *ListSafetyHelmetResultsResponseBodyResults {
	s.Feature = &v
	return s
}

func (s *ListSafetyHelmetResultsResponseBodyResults) SetObjLeft(v int32) *ListSafetyHelmetResultsResponseBodyResults {
	s.ObjLeft = &v
	return s
}

func (s *ListSafetyHelmetResultsResponseBodyResults) SetCropImage(v string) *ListSafetyHelmetResultsResponseBodyResults {
	s.CropImage = &v
	return s
}

func (s *ListSafetyHelmetResultsResponseBodyResults) SetObjTop(v int32) *ListSafetyHelmetResultsResponseBodyResults {
	s.ObjTop = &v
	return s
}

func (s *ListSafetyHelmetResultsResponseBodyResults) SetLabel(v string) *ListSafetyHelmetResultsResponseBodyResults {
	s.Label = &v
	return s
}

func (s *ListSafetyHelmetResultsResponseBodyResults) SetLeaveTime(v string) *ListSafetyHelmetResultsResponseBodyResults {
	s.LeaveTime = &v
	return s
}

func (s *ListSafetyHelmetResultsResponseBodyResults) SetObjBottom(v int32) *ListSafetyHelmetResultsResponseBodyResults {
	s.ObjBottom = &v
	return s
}

func (s *ListSafetyHelmetResultsResponseBodyResults) SetTimeStamp(v string) *ListSafetyHelmetResultsResponseBodyResults {
	s.TimeStamp = &v
	return s
}

func (s *ListSafetyHelmetResultsResponseBodyResults) SetOrigImage(v string) *ListSafetyHelmetResultsResponseBodyResults {
	s.OrigImage = &v
	return s
}

func (s *ListSafetyHelmetResultsResponseBodyResults) SetImageId(v string) *ListSafetyHelmetResultsResponseBodyResults {
	s.ImageId = &v
	return s
}

func (s *ListSafetyHelmetResultsResponseBodyResults) SetEntryTime(v string) *ListSafetyHelmetResultsResponseBodyResults {
	s.EntryTime = &v
	return s
}

type ListSafetyHelmetResultsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSafetyHelmetResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSafetyHelmetResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSafetyHelmetResultsResponse) GoString() string {
	return s.String()
}

func (s *ListSafetyHelmetResultsResponse) SetHeaders(v map[string]*string) *ListSafetyHelmetResultsResponse {
	s.Headers = v
	return s
}

func (s *ListSafetyHelmetResultsResponse) SetBody(v *ListSafetyHelmetResultsResponseBody) *ListSafetyHelmetResultsResponse {
	s.Body = v
	return s
}

type ListStreamsForCamerasRequest struct {
	CameraIds  *string `json:"CameraIds,omitempty" xml:"CameraIds,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListStreamsForCamerasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStreamsForCamerasRequest) GoString() string {
	return s.String()
}

func (s *ListStreamsForCamerasRequest) SetCameraIds(v string) *ListStreamsForCamerasRequest {
	s.CameraIds = &v
	return s
}

func (s *ListStreamsForCamerasRequest) SetRegionId(v string) *ListStreamsForCamerasRequest {
	s.RegionId = &v
	return s
}

func (s *ListStreamsForCamerasRequest) SetInstanceId(v string) *ListStreamsForCamerasRequest {
	s.InstanceId = &v
	return s
}

type ListStreamsForCamerasResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Streams   *ListStreamsForCamerasResponseBodyStreams `json:"Streams,omitempty" xml:"Streams,omitempty" type:"Struct"`
}

func (s ListStreamsForCamerasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStreamsForCamerasResponseBody) GoString() string {
	return s.String()
}

func (s *ListStreamsForCamerasResponseBody) SetRequestId(v string) *ListStreamsForCamerasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStreamsForCamerasResponseBody) SetStreams(v *ListStreamsForCamerasResponseBodyStreams) *ListStreamsForCamerasResponseBody {
	s.Streams = v
	return s
}

type ListStreamsForCamerasResponseBodyStreams struct {
	Stream []*string `json:"Stream,omitempty" xml:"Stream,omitempty" type:"Repeated"`
}

func (s ListStreamsForCamerasResponseBodyStreams) String() string {
	return tea.Prettify(s)
}

func (s ListStreamsForCamerasResponseBodyStreams) GoString() string {
	return s.String()
}

func (s *ListStreamsForCamerasResponseBodyStreams) SetStream(v []*string) *ListStreamsForCamerasResponseBodyStreams {
	s.Stream = v
	return s
}

type ListStreamsForCamerasResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListStreamsForCamerasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStreamsForCamerasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStreamsForCamerasResponse) GoString() string {
	return s.String()
}

func (s *ListStreamsForCamerasResponse) SetHeaders(v map[string]*string) *ListStreamsForCamerasResponse {
	s.Headers = v
	return s
}

func (s *ListStreamsForCamerasResponse) SetBody(v *ListStreamsForCamerasResponseBody) *ListStreamsForCamerasResponse {
	s.Body = v
	return s
}

type ListVehicleEntryResultsRequest struct {
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	CameraIds    *string `json:"CameraIds,omitempty" xml:"CameraIds,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VehiclePlate *string `json:"VehiclePlate,omitempty" xml:"VehiclePlate,omitempty"`
}

func (s ListVehicleEntryResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVehicleEntryResultsRequest) GoString() string {
	return s.String()
}

func (s *ListVehicleEntryResultsRequest) SetStartTime(v string) *ListVehicleEntryResultsRequest {
	s.StartTime = &v
	return s
}

func (s *ListVehicleEntryResultsRequest) SetEndTime(v string) *ListVehicleEntryResultsRequest {
	s.EndTime = &v
	return s
}

func (s *ListVehicleEntryResultsRequest) SetCameraIds(v string) *ListVehicleEntryResultsRequest {
	s.CameraIds = &v
	return s
}

func (s *ListVehicleEntryResultsRequest) SetInstanceId(v string) *ListVehicleEntryResultsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListVehicleEntryResultsRequest) SetRegionId(v string) *ListVehicleEntryResultsRequest {
	s.RegionId = &v
	return s
}

func (s *ListVehicleEntryResultsRequest) SetVehiclePlate(v string) *ListVehicleEntryResultsRequest {
	s.VehiclePlate = &v
	return s
}

type ListVehicleEntryResultsResponseBody struct {
	TotalCount *int64                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int64                                        `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	PageSize   *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Results    []*ListVehicleEntryResultsResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s ListVehicleEntryResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVehicleEntryResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVehicleEntryResultsResponseBody) SetTotalCount(v int64) *ListVehicleEntryResultsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVehicleEntryResultsResponseBody) SetTotalPage(v int64) *ListVehicleEntryResultsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListVehicleEntryResultsResponseBody) SetPageSize(v int32) *ListVehicleEntryResultsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListVehicleEntryResultsResponseBody) SetRequestId(v string) *ListVehicleEntryResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVehicleEntryResultsResponseBody) SetPageNumber(v int32) *ListVehicleEntryResultsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListVehicleEntryResultsResponseBody) SetResults(v []*ListVehicleEntryResultsResponseBodyResults) *ListVehicleEntryResultsResponseBody {
	s.Results = v
	return s
}

type ListVehicleEntryResultsResponseBodyResults struct {
	CameraId     *string  `json:"CameraId,omitempty" xml:"CameraId,omitempty"`
	ObjRight     *int32   `json:"ObjRight,omitempty" xml:"ObjRight,omitempty"`
	ObjType      *string  `json:"ObjType,omitempty" xml:"ObjType,omitempty"`
	Feature      *string  `json:"Feature,omitempty" xml:"Feature,omitempty"`
	PlateLeft    *int32   `json:"PlateLeft,omitempty" xml:"PlateLeft,omitempty"`
	ObjLeft      *int32   `json:"ObjLeft,omitempty" xml:"ObjLeft,omitempty"`
	Score        *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	CropImage    *string  `json:"CropImage,omitempty" xml:"CropImage,omitempty"`
	ObjTop       *int32   `json:"ObjTop,omitempty" xml:"ObjTop,omitempty"`
	Label        *string  `json:"Label,omitempty" xml:"Label,omitempty"`
	PlateBottom  *int32   `json:"PlateBottom,omitempty" xml:"PlateBottom,omitempty"`
	LeaveTime    *string  `json:"LeaveTime,omitempty" xml:"LeaveTime,omitempty"`
	PlateTop     *int32   `json:"PlateTop,omitempty" xml:"PlateTop,omitempty"`
	ObjBottom    *int32   `json:"ObjBottom,omitempty" xml:"ObjBottom,omitempty"`
	PlateRight   *int32   `json:"PlateRight,omitempty" xml:"PlateRight,omitempty"`
	OrigImage    *string  `json:"OrigImage,omitempty" xml:"OrigImage,omitempty"`
	TimeStamp    *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	VehiclePlate *string  `json:"VehiclePlate,omitempty" xml:"VehiclePlate,omitempty"`
	EntryTime    *string  `json:"EntryTime,omitempty" xml:"EntryTime,omitempty"`
	ImageId      *string  `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s ListVehicleEntryResultsResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s ListVehicleEntryResultsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *ListVehicleEntryResultsResponseBodyResults) SetCameraId(v string) *ListVehicleEntryResultsResponseBodyResults {
	s.CameraId = &v
	return s
}

func (s *ListVehicleEntryResultsResponseBodyResults) SetObjRight(v int32) *ListVehicleEntryResultsResponseBodyResults {
	s.ObjRight = &v
	return s
}

func (s *ListVehicleEntryResultsResponseBodyResults) SetObjType(v string) *ListVehicleEntryResultsResponseBodyResults {
	s.ObjType = &v
	return s
}

func (s *ListVehicleEntryResultsResponseBodyResults) SetFeature(v string) *ListVehicleEntryResultsResponseBodyResults {
	s.Feature = &v
	return s
}

func (s *ListVehicleEntryResultsResponseBodyResults) SetPlateLeft(v int32) *ListVehicleEntryResultsResponseBodyResults {
	s.PlateLeft = &v
	return s
}

func (s *ListVehicleEntryResultsResponseBodyResults) SetObjLeft(v int32) *ListVehicleEntryResultsResponseBodyResults {
	s.ObjLeft = &v
	return s
}

func (s *ListVehicleEntryResultsResponseBodyResults) SetScore(v float32) *ListVehicleEntryResultsResponseBodyResults {
	s.Score = &v
	return s
}

func (s *ListVehicleEntryResultsResponseBodyResults) SetCropImage(v string) *ListVehicleEntryResultsResponseBodyResults {
	s.CropImage = &v
	return s
}

func (s *ListVehicleEntryResultsResponseBodyResults) SetObjTop(v int32) *ListVehicleEntryResultsResponseBodyResults {
	s.ObjTop = &v
	return s
}

func (s *ListVehicleEntryResultsResponseBodyResults) SetLabel(v string) *ListVehicleEntryResultsResponseBodyResults {
	s.Label = &v
	return s
}

func (s *ListVehicleEntryResultsResponseBodyResults) SetPlateBottom(v int32) *ListVehicleEntryResultsResponseBodyResults {
	s.PlateBottom = &v
	return s
}

func (s *ListVehicleEntryResultsResponseBodyResults) SetLeaveTime(v string) *ListVehicleEntryResultsResponseBodyResults {
	s.LeaveTime = &v
	return s
}

func (s *ListVehicleEntryResultsResponseBodyResults) SetPlateTop(v int32) *ListVehicleEntryResultsResponseBodyResults {
	s.PlateTop = &v
	return s
}

func (s *ListVehicleEntryResultsResponseBodyResults) SetObjBottom(v int32) *ListVehicleEntryResultsResponseBodyResults {
	s.ObjBottom = &v
	return s
}

func (s *ListVehicleEntryResultsResponseBodyResults) SetPlateRight(v int32) *ListVehicleEntryResultsResponseBodyResults {
	s.PlateRight = &v
	return s
}

func (s *ListVehicleEntryResultsResponseBodyResults) SetOrigImage(v string) *ListVehicleEntryResultsResponseBodyResults {
	s.OrigImage = &v
	return s
}

func (s *ListVehicleEntryResultsResponseBodyResults) SetTimeStamp(v string) *ListVehicleEntryResultsResponseBodyResults {
	s.TimeStamp = &v
	return s
}

func (s *ListVehicleEntryResultsResponseBodyResults) SetVehiclePlate(v string) *ListVehicleEntryResultsResponseBodyResults {
	s.VehiclePlate = &v
	return s
}

func (s *ListVehicleEntryResultsResponseBodyResults) SetEntryTime(v string) *ListVehicleEntryResultsResponseBodyResults {
	s.EntryTime = &v
	return s
}

func (s *ListVehicleEntryResultsResponseBodyResults) SetImageId(v string) *ListVehicleEntryResultsResponseBodyResults {
	s.ImageId = &v
	return s
}

type ListVehicleEntryResultsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVehicleEntryResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVehicleEntryResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVehicleEntryResultsResponse) GoString() string {
	return s.String()
}

func (s *ListVehicleEntryResultsResponse) SetHeaders(v map[string]*string) *ListVehicleEntryResultsResponse {
	s.Headers = v
	return s
}

func (s *ListVehicleEntryResultsResponse) SetBody(v *ListVehicleEntryResultsResponseBody) *ListVehicleEntryResultsResponse {
	s.Body = v
	return s
}

type ModifyAlgoLibRequest struct {
	AlgoLibName        *string `json:"AlgoLibName,omitempty" xml:"AlgoLibName,omitempty"`
	AlgoLibVersion     *string `json:"AlgoLibVersion,omitempty" xml:"AlgoLibVersion,omitempty"`
	Capacity           *int32  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	OssEndpoint        *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	OssAccessKeyId     *string `json:"OssAccessKeyId,omitempty" xml:"OssAccessKeyId,omitempty"`
	OssBucket          *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssPath            *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	JsonText           *string `json:"JsonText,omitempty" xml:"JsonText,omitempty"`
	CapabilityNames    *string `json:"CapabilityNames,omitempty" xml:"CapabilityNames,omitempty"`
	AlgoLibId          *string `json:"AlgoLibId,omitempty" xml:"AlgoLibId,omitempty"`
	OssAccessKeySecret *string `json:"OssAccessKeySecret,omitempty" xml:"OssAccessKeySecret,omitempty"`
	Text               *string `json:"Text,omitempty" xml:"Text,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ModifyAlgoLibRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAlgoLibRequest) GoString() string {
	return s.String()
}

func (s *ModifyAlgoLibRequest) SetAlgoLibName(v string) *ModifyAlgoLibRequest {
	s.AlgoLibName = &v
	return s
}

func (s *ModifyAlgoLibRequest) SetAlgoLibVersion(v string) *ModifyAlgoLibRequest {
	s.AlgoLibVersion = &v
	return s
}

func (s *ModifyAlgoLibRequest) SetCapacity(v int32) *ModifyAlgoLibRequest {
	s.Capacity = &v
	return s
}

func (s *ModifyAlgoLibRequest) SetOssEndpoint(v string) *ModifyAlgoLibRequest {
	s.OssEndpoint = &v
	return s
}

func (s *ModifyAlgoLibRequest) SetOssAccessKeyId(v string) *ModifyAlgoLibRequest {
	s.OssAccessKeyId = &v
	return s
}

func (s *ModifyAlgoLibRequest) SetOssBucket(v string) *ModifyAlgoLibRequest {
	s.OssBucket = &v
	return s
}

func (s *ModifyAlgoLibRequest) SetOssPath(v string) *ModifyAlgoLibRequest {
	s.OssPath = &v
	return s
}

func (s *ModifyAlgoLibRequest) SetJsonText(v string) *ModifyAlgoLibRequest {
	s.JsonText = &v
	return s
}

func (s *ModifyAlgoLibRequest) SetCapabilityNames(v string) *ModifyAlgoLibRequest {
	s.CapabilityNames = &v
	return s
}

func (s *ModifyAlgoLibRequest) SetAlgoLibId(v string) *ModifyAlgoLibRequest {
	s.AlgoLibId = &v
	return s
}

func (s *ModifyAlgoLibRequest) SetOssAccessKeySecret(v string) *ModifyAlgoLibRequest {
	s.OssAccessKeySecret = &v
	return s
}

func (s *ModifyAlgoLibRequest) SetText(v string) *ModifyAlgoLibRequest {
	s.Text = &v
	return s
}

func (s *ModifyAlgoLibRequest) SetRegionId(v string) *ModifyAlgoLibRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAlgoLibRequest) SetInstanceId(v string) *ModifyAlgoLibRequest {
	s.InstanceId = &v
	return s
}

type ModifyAlgoLibResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAlgoLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAlgoLibResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAlgoLibResponseBody) SetRequestId(v string) *ModifyAlgoLibResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAlgoLibResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAlgoLibResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAlgoLibResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAlgoLibResponse) GoString() string {
	return s.String()
}

func (s *ModifyAlgoLibResponse) SetHeaders(v map[string]*string) *ModifyAlgoLibResponse {
	s.Headers = v
	return s
}

func (s *ModifyAlgoLibResponse) SetBody(v *ModifyAlgoLibResponseBody) *ModifyAlgoLibResponse {
	s.Body = v
	return s
}

type ModifyCameraRequest struct {
	CameraName   *string `json:"CameraName,omitempty" xml:"CameraName,omitempty"`
	CameraId     *string `json:"CameraId,omitempty" xml:"CameraId,omitempty"`
	LocationInfo *string `json:"LocationInfo,omitempty" xml:"LocationInfo,omitempty"`
	InviteUri    *string `json:"InviteUri,omitempty" xml:"InviteUri,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StreamStatus *string `json:"StreamStatus,omitempty" xml:"StreamStatus,omitempty"`
}

func (s ModifyCameraRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCameraRequest) GoString() string {
	return s.String()
}

func (s *ModifyCameraRequest) SetCameraName(v string) *ModifyCameraRequest {
	s.CameraName = &v
	return s
}

func (s *ModifyCameraRequest) SetCameraId(v string) *ModifyCameraRequest {
	s.CameraId = &v
	return s
}

func (s *ModifyCameraRequest) SetLocationInfo(v string) *ModifyCameraRequest {
	s.LocationInfo = &v
	return s
}

func (s *ModifyCameraRequest) SetInviteUri(v string) *ModifyCameraRequest {
	s.InviteUri = &v
	return s
}

func (s *ModifyCameraRequest) SetRegionId(v string) *ModifyCameraRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCameraRequest) SetStreamStatus(v string) *ModifyCameraRequest {
	s.StreamStatus = &v
	return s
}

type ModifyCameraResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCameraResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyCameraResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCameraResponseBody) SetRequestId(v string) *ModifyCameraResponseBody {
	s.RequestId = &v
	return s
}

type ModifyCameraResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyCameraResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyCameraResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyCameraResponse) GoString() string {
	return s.String()
}

func (s *ModifyCameraResponse) SetHeaders(v map[string]*string) *ModifyCameraResponse {
	s.Headers = v
	return s
}

func (s *ModifyCameraResponse) SetBody(v *ModifyCameraResponseBody) *ModifyCameraResponse {
	s.Body = v
	return s
}

type ModifyCapabilityRequest struct {
	CapabilityName *string `json:"CapabilityName,omitempty" xml:"CapabilityName,omitempty"`
	CapabilityId   *string `json:"CapabilityId,omitempty" xml:"CapabilityId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ModifyCapabilityRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCapabilityRequest) GoString() string {
	return s.String()
}

func (s *ModifyCapabilityRequest) SetCapabilityName(v string) *ModifyCapabilityRequest {
	s.CapabilityName = &v
	return s
}

func (s *ModifyCapabilityRequest) SetCapabilityId(v string) *ModifyCapabilityRequest {
	s.CapabilityId = &v
	return s
}

func (s *ModifyCapabilityRequest) SetRegionId(v string) *ModifyCapabilityRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCapabilityRequest) SetInstanceId(v string) *ModifyCapabilityRequest {
	s.InstanceId = &v
	return s
}

type ModifyCapabilityResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCapabilityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyCapabilityResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCapabilityResponseBody) SetRequestId(v string) *ModifyCapabilityResponseBody {
	s.RequestId = &v
	return s
}

type ModifyCapabilityResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyCapabilityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyCapabilityResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyCapabilityResponse) GoString() string {
	return s.String()
}

func (s *ModifyCapabilityResponse) SetHeaders(v map[string]*string) *ModifyCapabilityResponse {
	s.Headers = v
	return s
}

func (s *ModifyCapabilityResponse) SetBody(v *ModifyCapabilityResponseBody) *ModifyCapabilityResponse {
	s.Body = v
	return s
}

type ModifyInstanceRequest struct {
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRequest) SetInstanceName(v string) *ModifyInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceRequest) SetDescription(v string) *ModifyInstanceRequest {
	s.Description = &v
	return s
}

func (s *ModifyInstanceRequest) SetStatus(v string) *ModifyInstanceRequest {
	s.Status = &v
	return s
}

func (s *ModifyInstanceRequest) SetInstanceId(v string) *ModifyInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceRequest) SetRegionId(v string) *ModifyInstanceRequest {
	s.RegionId = &v
	return s
}

type ModifyInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceResponseBody) SetRequestId(v string) *ModifyInstanceResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceResponse) SetHeaders(v map[string]*string) *ModifyInstanceResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceResponse) SetBody(v *ModifyInstanceResponseBody) *ModifyInstanceResponse {
	s.Body = v
	return s
}

type ModifyJobGroupRequest struct {
	JobGroupName      *string                                `json:"JobGroupName,omitempty" xml:"JobGroupName,omitempty"`
	ResourceProfileId *string                                `json:"ResourceProfileId,omitempty" xml:"ResourceProfileId,omitempty"`
	AlgoLibId         *string                                `json:"AlgoLibId,omitempty" xml:"AlgoLibId,omitempty"`
	PlanedJobCount    *int32                                 `json:"PlanedJobCount,omitempty" xml:"PlanedJobCount,omitempty"`
	StreamPerJob      *int32                                 `json:"StreamPerJob,omitempty" xml:"StreamPerJob,omitempty"`
	JobCount          *int32                                 `json:"JobCount,omitempty" xml:"JobCount,omitempty"`
	JobGroupId        *string                                `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	Status            *string                                `json:"Status,omitempty" xml:"Status,omitempty"`
	RegionId          *string                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId        *string                                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobGroupParams    []*ModifyJobGroupRequestJobGroupParams `json:"JobGroupParams,omitempty" xml:"JobGroupParams,omitempty" type:"Repeated"`
}

func (s ModifyJobGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyJobGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyJobGroupRequest) SetJobGroupName(v string) *ModifyJobGroupRequest {
	s.JobGroupName = &v
	return s
}

func (s *ModifyJobGroupRequest) SetResourceProfileId(v string) *ModifyJobGroupRequest {
	s.ResourceProfileId = &v
	return s
}

func (s *ModifyJobGroupRequest) SetAlgoLibId(v string) *ModifyJobGroupRequest {
	s.AlgoLibId = &v
	return s
}

func (s *ModifyJobGroupRequest) SetPlanedJobCount(v int32) *ModifyJobGroupRequest {
	s.PlanedJobCount = &v
	return s
}

func (s *ModifyJobGroupRequest) SetStreamPerJob(v int32) *ModifyJobGroupRequest {
	s.StreamPerJob = &v
	return s
}

func (s *ModifyJobGroupRequest) SetJobCount(v int32) *ModifyJobGroupRequest {
	s.JobCount = &v
	return s
}

func (s *ModifyJobGroupRequest) SetJobGroupId(v string) *ModifyJobGroupRequest {
	s.JobGroupId = &v
	return s
}

func (s *ModifyJobGroupRequest) SetStatus(v string) *ModifyJobGroupRequest {
	s.Status = &v
	return s
}

func (s *ModifyJobGroupRequest) SetRegionId(v string) *ModifyJobGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyJobGroupRequest) SetInstanceId(v string) *ModifyJobGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyJobGroupRequest) SetJobGroupParams(v []*ModifyJobGroupRequestJobGroupParams) *ModifyJobGroupRequest {
	s.JobGroupParams = v
	return s
}

type ModifyJobGroupRequestJobGroupParams struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyJobGroupRequestJobGroupParams) String() string {
	return tea.Prettify(s)
}

func (s ModifyJobGroupRequestJobGroupParams) GoString() string {
	return s.String()
}

func (s *ModifyJobGroupRequestJobGroupParams) SetKey(v string) *ModifyJobGroupRequestJobGroupParams {
	s.Key = &v
	return s
}

func (s *ModifyJobGroupRequestJobGroupParams) SetValue(v string) *ModifyJobGroupRequestJobGroupParams {
	s.Value = &v
	return s
}

type ModifyJobGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyJobGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyJobGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyJobGroupResponseBody) SetRequestId(v string) *ModifyJobGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyJobGroupResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyJobGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyJobGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyJobGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyJobGroupResponse) SetHeaders(v map[string]*string) *ModifyJobGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyJobGroupResponse) SetBody(v *ModifyJobGroupResponseBody) *ModifyJobGroupResponse {
	s.Body = v
	return s
}

type ModifyResourceProfileRequest struct {
	ResourceProfileName   *string                                              `json:"ResourceProfileName,omitempty" xml:"ResourceProfileName,omitempty"`
	ResourceProfileId     *string                                              `json:"ResourceProfileId,omitempty" xml:"ResourceProfileId,omitempty"`
	RegionId              *string                                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId            *string                                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceProfileParams []*ModifyResourceProfileRequestResourceProfileParams `json:"ResourceProfileParams,omitempty" xml:"ResourceProfileParams,omitempty" type:"Repeated"`
}

func (s ModifyResourceProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyResourceProfileRequest) GoString() string {
	return s.String()
}

func (s *ModifyResourceProfileRequest) SetResourceProfileName(v string) *ModifyResourceProfileRequest {
	s.ResourceProfileName = &v
	return s
}

func (s *ModifyResourceProfileRequest) SetResourceProfileId(v string) *ModifyResourceProfileRequest {
	s.ResourceProfileId = &v
	return s
}

func (s *ModifyResourceProfileRequest) SetRegionId(v string) *ModifyResourceProfileRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyResourceProfileRequest) SetInstanceId(v string) *ModifyResourceProfileRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyResourceProfileRequest) SetResourceProfileParams(v []*ModifyResourceProfileRequestResourceProfileParams) *ModifyResourceProfileRequest {
	s.ResourceProfileParams = v
	return s
}

type ModifyResourceProfileRequestResourceProfileParams struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyResourceProfileRequestResourceProfileParams) String() string {
	return tea.Prettify(s)
}

func (s ModifyResourceProfileRequestResourceProfileParams) GoString() string {
	return s.String()
}

func (s *ModifyResourceProfileRequestResourceProfileParams) SetKey(v string) *ModifyResourceProfileRequestResourceProfileParams {
	s.Key = &v
	return s
}

func (s *ModifyResourceProfileRequestResourceProfileParams) SetValue(v string) *ModifyResourceProfileRequestResourceProfileParams {
	s.Value = &v
	return s
}

type ModifyResourceProfileResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyResourceProfileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyResourceProfileResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyResourceProfileResponseBody) SetRequestId(v string) *ModifyResourceProfileResponseBody {
	s.RequestId = &v
	return s
}

type ModifyResourceProfileResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyResourceProfileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyResourceProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyResourceProfileResponse) GoString() string {
	return s.String()
}

func (s *ModifyResourceProfileResponse) SetHeaders(v map[string]*string) *ModifyResourceProfileResponse {
	s.Headers = v
	return s
}

func (s *ModifyResourceProfileResponse) SetBody(v *ModifyResourceProfileResponseBody) *ModifyResourceProfileResponse {
	s.Body = v
	return s
}

type ModifyWorkGroupRequest struct {
	WorkGroupId *string `json:"WorkGroupId,omitempty" xml:"WorkGroupId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyWorkGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWorkGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyWorkGroupRequest) SetWorkGroupId(v string) *ModifyWorkGroupRequest {
	s.WorkGroupId = &v
	return s
}

func (s *ModifyWorkGroupRequest) SetDescription(v string) *ModifyWorkGroupRequest {
	s.Description = &v
	return s
}

func (s *ModifyWorkGroupRequest) SetRegionId(v string) *ModifyWorkGroupRequest {
	s.RegionId = &v
	return s
}

type ModifyWorkGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWorkGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyWorkGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWorkGroupResponseBody) SetRequestId(v string) *ModifyWorkGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyWorkGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyWorkGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyWorkGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyWorkGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyWorkGroupResponse) SetHeaders(v map[string]*string) *ModifyWorkGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyWorkGroupResponse) SetBody(v *ModifyWorkGroupResponseBody) *ModifyWorkGroupResponse {
	s.Body = v
	return s
}

type PutCameraConfForCameraRequest struct {
	CameraId   *string `json:"CameraId,omitempty" xml:"CameraId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	CameraName *string `json:"CameraName,omitempty" xml:"CameraName,omitempty"`
	Context    *string `json:"Context,omitempty" xml:"Context,omitempty"`
}

func (s PutCameraConfForCameraRequest) String() string {
	return tea.Prettify(s)
}

func (s PutCameraConfForCameraRequest) GoString() string {
	return s.String()
}

func (s *PutCameraConfForCameraRequest) SetCameraId(v string) *PutCameraConfForCameraRequest {
	s.CameraId = &v
	return s
}

func (s *PutCameraConfForCameraRequest) SetInstanceId(v string) *PutCameraConfForCameraRequest {
	s.InstanceId = &v
	return s
}

func (s *PutCameraConfForCameraRequest) SetRegionId(v string) *PutCameraConfForCameraRequest {
	s.RegionId = &v
	return s
}

func (s *PutCameraConfForCameraRequest) SetCameraName(v string) *PutCameraConfForCameraRequest {
	s.CameraName = &v
	return s
}

func (s *PutCameraConfForCameraRequest) SetContext(v string) *PutCameraConfForCameraRequest {
	s.Context = &v
	return s
}

type PutCameraConfForCameraResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutCameraConfForCameraResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutCameraConfForCameraResponseBody) GoString() string {
	return s.String()
}

func (s *PutCameraConfForCameraResponseBody) SetRequestId(v string) *PutCameraConfForCameraResponseBody {
	s.RequestId = &v
	return s
}

type PutCameraConfForCameraResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutCameraConfForCameraResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutCameraConfForCameraResponse) String() string {
	return tea.Prettify(s)
}

func (s PutCameraConfForCameraResponse) GoString() string {
	return s.String()
}

func (s *PutCameraConfForCameraResponse) SetHeaders(v map[string]*string) *PutCameraConfForCameraResponse {
	s.Headers = v
	return s
}

func (s *PutCameraConfForCameraResponse) SetBody(v *PutCameraConfForCameraResponseBody) *PutCameraConfForCameraResponse {
	s.Body = v
	return s
}

type SearchImagesRequest struct {
	From       *int32                          `json:"From,omitempty" xml:"From,omitempty"`
	Size       *int32                          `json:"Size,omitempty" xml:"Size,omitempty"`
	Type       *string                         `json:"Type,omitempty" xml:"Type,omitempty"`
	Contents   *string                         `json:"Contents,omitempty" xml:"Contents,omitempty"`
	StartTime  *string                         `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string                         `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ImageIds   *string                         `json:"ImageIds,omitempty" xml:"ImageIds,omitempty"`
	CameraIds  *string                         `json:"CameraIds,omitempty" xml:"CameraIds,omitempty"`
	InstanceId *string                         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NoFeature  *string                         `json:"NoFeature,omitempty" xml:"NoFeature,omitempty"`
	RegionId   *string                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Attribute  []*SearchImagesRequestAttribute `json:"Attribute,omitempty" xml:"Attribute,omitempty" type:"Repeated"`
}

func (s SearchImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchImagesRequest) GoString() string {
	return s.String()
}

func (s *SearchImagesRequest) SetFrom(v int32) *SearchImagesRequest {
	s.From = &v
	return s
}

func (s *SearchImagesRequest) SetSize(v int32) *SearchImagesRequest {
	s.Size = &v
	return s
}

func (s *SearchImagesRequest) SetType(v string) *SearchImagesRequest {
	s.Type = &v
	return s
}

func (s *SearchImagesRequest) SetContents(v string) *SearchImagesRequest {
	s.Contents = &v
	return s
}

func (s *SearchImagesRequest) SetStartTime(v string) *SearchImagesRequest {
	s.StartTime = &v
	return s
}

func (s *SearchImagesRequest) SetEndTime(v string) *SearchImagesRequest {
	s.EndTime = &v
	return s
}

func (s *SearchImagesRequest) SetImageIds(v string) *SearchImagesRequest {
	s.ImageIds = &v
	return s
}

func (s *SearchImagesRequest) SetCameraIds(v string) *SearchImagesRequest {
	s.CameraIds = &v
	return s
}

func (s *SearchImagesRequest) SetInstanceId(v string) *SearchImagesRequest {
	s.InstanceId = &v
	return s
}

func (s *SearchImagesRequest) SetNoFeature(v string) *SearchImagesRequest {
	s.NoFeature = &v
	return s
}

func (s *SearchImagesRequest) SetRegionId(v string) *SearchImagesRequest {
	s.RegionId = &v
	return s
}

func (s *SearchImagesRequest) SetAttribute(v []*SearchImagesRequestAttribute) *SearchImagesRequest {
	s.Attribute = v
	return s
}

type SearchImagesRequestAttribute struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchImagesRequestAttribute) String() string {
	return tea.Prettify(s)
}

func (s SearchImagesRequestAttribute) GoString() string {
	return s.String()
}

func (s *SearchImagesRequestAttribute) SetKey(v string) *SearchImagesRequestAttribute {
	s.Key = &v
	return s
}

func (s *SearchImagesRequestAttribute) SetValue(v string) *SearchImagesRequestAttribute {
	s.Value = &v
	return s
}

type SearchImagesResponseBody struct {
	TotalCount *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Images     *SearchImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Struct"`
}

func (s SearchImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchImagesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchImagesResponseBody) SetTotalCount(v int32) *SearchImagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchImagesResponseBody) SetRequestId(v string) *SearchImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchImagesResponseBody) SetImages(v *SearchImagesResponseBodyImages) *SearchImagesResponseBody {
	s.Images = v
	return s
}

type SearchImagesResponseBodyImages struct {
	Image []*SearchImagesResponseBodyImagesImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Repeated"`
}

func (s SearchImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s SearchImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *SearchImagesResponseBodyImages) SetImage(v []*SearchImagesResponseBodyImagesImage) *SearchImagesResponseBodyImages {
	s.Image = v
	return s
}

type SearchImagesResponseBodyImagesImage struct {
	CameraId            *string  `json:"CameraId,omitempty" xml:"CameraId,omitempty"`
	TrouserTypeScore    *float32 `json:"TrouserTypeScore,omitempty" xml:"TrouserTypeScore,omitempty"`
	ObjType             *string  `json:"ObjType,omitempty" xml:"ObjType,omitempty"`
	TrouserColor        *string  `json:"TrouserColor,omitempty" xml:"TrouserColor,omitempty"`
	ClothTypeScore      *float32 `json:"ClothTypeScore,omitempty" xml:"ClothTypeScore,omitempty"`
	Brand               *string  `json:"Brand,omitempty" xml:"Brand,omitempty"`
	PoseType            *string  `json:"PoseType,omitempty" xml:"PoseType,omitempty"`
	VehicleColor        *string  `json:"VehicleColor,omitempty" xml:"VehicleColor,omitempty"`
	ObjLeft             *int32   `json:"ObjLeft,omitempty" xml:"ObjLeft,omitempty"`
	Score               *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	HeadWearScore       *float32 `json:"HeadWearScore,omitempty" xml:"HeadWearScore,omitempty"`
	AgeTypeScore        *float32 `json:"AgeTypeScore,omitempty" xml:"AgeTypeScore,omitempty"`
	SexTypeScore        *float32 `json:"SexTypeScore,omitempty" xml:"SexTypeScore,omitempty"`
	NonVehicleType      *string  `json:"NonVehicleType,omitempty" xml:"NonVehicleType,omitempty"`
	ObjBottom           *int32   `json:"ObjBottom,omitempty" xml:"ObjBottom,omitempty"`
	PlateNumber         *string  `json:"PlateNumber,omitempty" xml:"PlateNumber,omitempty"`
	ClothType           *string  `json:"ClothType,omitempty" xml:"ClothType,omitempty"`
	TimeStamp           *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	OrigImage           *string  `json:"OrigImage,omitempty" xml:"OrigImage,omitempty"`
	VehicleTypeScore    *float32 `json:"VehicleTypeScore,omitempty" xml:"VehicleTypeScore,omitempty"`
	TrouserColorScore   *float32 `json:"TrouserColorScore,omitempty" xml:"TrouserColorScore,omitempty"`
	ImageId             *string  `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	VehicleColorScore   *float32 `json:"VehicleColorScore,omitempty" xml:"VehicleColorScore,omitempty"`
	HairTypeScore       *float32 `json:"HairTypeScore,omitempty" xml:"HairTypeScore,omitempty"`
	HairType            *string  `json:"HairType,omitempty" xml:"HairType,omitempty"`
	NonVehicleTypeScore *float32 `json:"NonVehicleTypeScore,omitempty" xml:"NonVehicleTypeScore,omitempty"`
	HeadWear            *string  `json:"HeadWear,omitempty" xml:"HeadWear,omitempty"`
	VehicleType         *string  `json:"VehicleType,omitempty" xml:"VehicleType,omitempty"`
	SexType             *string  `json:"SexType,omitempty" xml:"SexType,omitempty"`
	PoseTypeScore       *float32 `json:"PoseTypeScore,omitempty" xml:"PoseTypeScore,omitempty"`
	ObjRight            *int32   `json:"ObjRight,omitempty" xml:"ObjRight,omitempty"`
	Feature             *string  `json:"Feature,omitempty" xml:"Feature,omitempty"`
	ClothColorScore     *float32 `json:"ClothColorScore,omitempty" xml:"ClothColorScore,omitempty"`
	CropImage           *string  `json:"CropImage,omitempty" xml:"CropImage,omitempty"`
	ObjTop              *int32   `json:"ObjTop,omitempty" xml:"ObjTop,omitempty"`
	BrandScore          *float32 `json:"BrandScore,omitempty" xml:"BrandScore,omitempty"`
	ClothColor          *string  `json:"ClothColor,omitempty" xml:"ClothColor,omitempty"`
	AgeType             *string  `json:"AgeType,omitempty" xml:"AgeType,omitempty"`
	LeaveTime           *string  `json:"LeaveTime,omitempty" xml:"LeaveTime,omitempty"`
	TrouserType         *string  `json:"TrouserType,omitempty" xml:"TrouserType,omitempty"`
	EntryTime           *string  `json:"EntryTime,omitempty" xml:"EntryTime,omitempty"`
}

func (s SearchImagesResponseBodyImagesImage) String() string {
	return tea.Prettify(s)
}

func (s SearchImagesResponseBodyImagesImage) GoString() string {
	return s.String()
}

func (s *SearchImagesResponseBodyImagesImage) SetCameraId(v string) *SearchImagesResponseBodyImagesImage {
	s.CameraId = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetTrouserTypeScore(v float32) *SearchImagesResponseBodyImagesImage {
	s.TrouserTypeScore = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetObjType(v string) *SearchImagesResponseBodyImagesImage {
	s.ObjType = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetTrouserColor(v string) *SearchImagesResponseBodyImagesImage {
	s.TrouserColor = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetClothTypeScore(v float32) *SearchImagesResponseBodyImagesImage {
	s.ClothTypeScore = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetBrand(v string) *SearchImagesResponseBodyImagesImage {
	s.Brand = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetPoseType(v string) *SearchImagesResponseBodyImagesImage {
	s.PoseType = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetVehicleColor(v string) *SearchImagesResponseBodyImagesImage {
	s.VehicleColor = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetObjLeft(v int32) *SearchImagesResponseBodyImagesImage {
	s.ObjLeft = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetScore(v float32) *SearchImagesResponseBodyImagesImage {
	s.Score = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetHeadWearScore(v float32) *SearchImagesResponseBodyImagesImage {
	s.HeadWearScore = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetAgeTypeScore(v float32) *SearchImagesResponseBodyImagesImage {
	s.AgeTypeScore = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetSexTypeScore(v float32) *SearchImagesResponseBodyImagesImage {
	s.SexTypeScore = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetNonVehicleType(v string) *SearchImagesResponseBodyImagesImage {
	s.NonVehicleType = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetObjBottom(v int32) *SearchImagesResponseBodyImagesImage {
	s.ObjBottom = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetPlateNumber(v string) *SearchImagesResponseBodyImagesImage {
	s.PlateNumber = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetClothType(v string) *SearchImagesResponseBodyImagesImage {
	s.ClothType = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetTimeStamp(v string) *SearchImagesResponseBodyImagesImage {
	s.TimeStamp = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetOrigImage(v string) *SearchImagesResponseBodyImagesImage {
	s.OrigImage = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetVehicleTypeScore(v float32) *SearchImagesResponseBodyImagesImage {
	s.VehicleTypeScore = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetTrouserColorScore(v float32) *SearchImagesResponseBodyImagesImage {
	s.TrouserColorScore = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetImageId(v string) *SearchImagesResponseBodyImagesImage {
	s.ImageId = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetVehicleColorScore(v float32) *SearchImagesResponseBodyImagesImage {
	s.VehicleColorScore = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetHairTypeScore(v float32) *SearchImagesResponseBodyImagesImage {
	s.HairTypeScore = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetHairType(v string) *SearchImagesResponseBodyImagesImage {
	s.HairType = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetNonVehicleTypeScore(v float32) *SearchImagesResponseBodyImagesImage {
	s.NonVehicleTypeScore = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetHeadWear(v string) *SearchImagesResponseBodyImagesImage {
	s.HeadWear = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetVehicleType(v string) *SearchImagesResponseBodyImagesImage {
	s.VehicleType = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetSexType(v string) *SearchImagesResponseBodyImagesImage {
	s.SexType = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetPoseTypeScore(v float32) *SearchImagesResponseBodyImagesImage {
	s.PoseTypeScore = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetObjRight(v int32) *SearchImagesResponseBodyImagesImage {
	s.ObjRight = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetFeature(v string) *SearchImagesResponseBodyImagesImage {
	s.Feature = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetClothColorScore(v float32) *SearchImagesResponseBodyImagesImage {
	s.ClothColorScore = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetCropImage(v string) *SearchImagesResponseBodyImagesImage {
	s.CropImage = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetObjTop(v int32) *SearchImagesResponseBodyImagesImage {
	s.ObjTop = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetBrandScore(v float32) *SearchImagesResponseBodyImagesImage {
	s.BrandScore = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetClothColor(v string) *SearchImagesResponseBodyImagesImage {
	s.ClothColor = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetAgeType(v string) *SearchImagesResponseBodyImagesImage {
	s.AgeType = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetLeaveTime(v string) *SearchImagesResponseBodyImagesImage {
	s.LeaveTime = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetTrouserType(v string) *SearchImagesResponseBodyImagesImage {
	s.TrouserType = &v
	return s
}

func (s *SearchImagesResponseBodyImagesImage) SetEntryTime(v string) *SearchImagesResponseBodyImagesImage {
	s.EntryTime = &v
	return s
}

type SearchImagesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchImagesResponse) GoString() string {
	return s.String()
}

func (s *SearchImagesResponse) SetHeaders(v map[string]*string) *SearchImagesResponse {
	s.Headers = v
	return s
}

func (s *SearchImagesResponse) SetBody(v *SearchImagesResponseBody) *SearchImagesResponse {
	s.Body = v
	return s
}

type StartJobGroupRequest struct {
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StartJobGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s StartJobGroupRequest) GoString() string {
	return s.String()
}

func (s *StartJobGroupRequest) SetJobGroupId(v string) *StartJobGroupRequest {
	s.JobGroupId = &v
	return s
}

func (s *StartJobGroupRequest) SetRegionId(v string) *StartJobGroupRequest {
	s.RegionId = &v
	return s
}

func (s *StartJobGroupRequest) SetInstanceId(v string) *StartJobGroupRequest {
	s.InstanceId = &v
	return s
}

type StartJobGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartJobGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartJobGroupResponseBody) GoString() string {
	return s.String()
}

func (s *StartJobGroupResponseBody) SetRequestId(v string) *StartJobGroupResponseBody {
	s.RequestId = &v
	return s
}

type StartJobGroupResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartJobGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartJobGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s StartJobGroupResponse) GoString() string {
	return s.String()
}

func (s *StartJobGroupResponse) SetHeaders(v map[string]*string) *StartJobGroupResponse {
	s.Headers = v
	return s
}

func (s *StartJobGroupResponse) SetBody(v *StartJobGroupResponseBody) *StartJobGroupResponse {
	s.Body = v
	return s
}

type StopJobGroupRequest struct {
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StopJobGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s StopJobGroupRequest) GoString() string {
	return s.String()
}

func (s *StopJobGroupRequest) SetJobGroupId(v string) *StopJobGroupRequest {
	s.JobGroupId = &v
	return s
}

func (s *StopJobGroupRequest) SetRegionId(v string) *StopJobGroupRequest {
	s.RegionId = &v
	return s
}

func (s *StopJobGroupRequest) SetInstanceId(v string) *StopJobGroupRequest {
	s.InstanceId = &v
	return s
}

type StopJobGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopJobGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopJobGroupResponseBody) GoString() string {
	return s.String()
}

func (s *StopJobGroupResponseBody) SetRequestId(v string) *StopJobGroupResponseBody {
	s.RequestId = &v
	return s
}

type StopJobGroupResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopJobGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopJobGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s StopJobGroupResponse) GoString() string {
	return s.String()
}

func (s *StopJobGroupResponse) SetHeaders(v map[string]*string) *StopJobGroupResponse {
	s.Headers = v
	return s
}

func (s *StopJobGroupResponse) SetBody(v *StopJobGroupResponseBody) *StopJobGroupResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("cityvisual"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AttachStreamWithOptions(request *AttachStreamRequest, runtime *util.RuntimeOptions) (_result *AttachStreamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AttachStreamResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AttachStream"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachStream(request *AttachStreamRequest) (_result *AttachStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachStreamResponse{}
	_body, _err := client.AttachStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchModifyCameraStatusWithOptions(request *BatchModifyCameraStatusRequest, runtime *util.RuntimeOptions) (_result *BatchModifyCameraStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BatchModifyCameraStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BatchModifyCameraStatus"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchModifyCameraStatus(request *BatchModifyCameraStatusRequest) (_result *BatchModifyCameraStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchModifyCameraStatusResponse{}
	_body, _err := client.BatchModifyCameraStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAlgoLibWithOptions(request *CreateAlgoLibRequest, runtime *util.RuntimeOptions) (_result *CreateAlgoLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAlgoLibResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAlgoLib"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAlgoLib(request *CreateAlgoLibRequest) (_result *CreateAlgoLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAlgoLibResponse{}
	_body, _err := client.CreateAlgoLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCameraWithOptions(request *CreateCameraRequest, runtime *util.RuntimeOptions) (_result *CreateCameraResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCameraResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCamera"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCamera(request *CreateCameraRequest) (_result *CreateCameraResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCameraResponse{}
	_body, _err := client.CreateCameraWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCapabilityWithOptions(request *CreateCapabilityRequest, runtime *util.RuntimeOptions) (_result *CreateCapabilityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCapabilityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCapability"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCapability(request *CreateCapabilityRequest) (_result *CreateCapabilityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCapabilityResponse{}
	_body, _err := client.CreateCapabilityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateInstance"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateJobGroupWithOptions(request *CreateJobGroupRequest, runtime *util.RuntimeOptions) (_result *CreateJobGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateJobGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateJobGroup"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateJobGroup(request *CreateJobGroupRequest) (_result *CreateJobGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateJobGroupResponse{}
	_body, _err := client.CreateJobGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateResourceProfileWithOptions(request *CreateResourceProfileRequest, runtime *util.RuntimeOptions) (_result *CreateResourceProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateResourceProfileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateResourceProfile"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateResourceProfile(request *CreateResourceProfileRequest) (_result *CreateResourceProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateResourceProfileResponse{}
	_body, _err := client.CreateResourceProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWorkGroupWithOptions(request *CreateWorkGroupRequest, runtime *util.RuntimeOptions) (_result *CreateWorkGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateWorkGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateWorkGroup"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWorkGroup(request *CreateWorkGroupRequest) (_result *CreateWorkGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWorkGroupResponse{}
	_body, _err := client.CreateWorkGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAlgoLibWithOptions(request *DeleteAlgoLibRequest, runtime *util.RuntimeOptions) (_result *DeleteAlgoLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAlgoLibResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAlgoLib"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAlgoLib(request *DeleteAlgoLibRequest) (_result *DeleteAlgoLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAlgoLibResponse{}
	_body, _err := client.DeleteAlgoLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCameraWithOptions(request *DeleteCameraRequest, runtime *util.RuntimeOptions) (_result *DeleteCameraResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteCameraResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCamera"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCamera(request *DeleteCameraRequest) (_result *DeleteCameraResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCameraResponse{}
	_body, _err := client.DeleteCameraWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCapabilityWithOptions(request *DeleteCapabilityRequest, runtime *util.RuntimeOptions) (_result *DeleteCapabilityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteCapabilityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCapability"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCapability(request *DeleteCapabilityRequest) (_result *DeleteCapabilityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCapabilityResponse{}
	_body, _err := client.DeleteCapabilityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteInstance"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInstance(request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteJobGroupWithOptions(request *DeleteJobGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteJobGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteJobGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteJobGroup"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteJobGroup(request *DeleteJobGroupRequest) (_result *DeleteJobGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteJobGroupResponse{}
	_body, _err := client.DeleteJobGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteResourceProfileWithOptions(request *DeleteResourceProfileRequest, runtime *util.RuntimeOptions) (_result *DeleteResourceProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteResourceProfileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteResourceProfile"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteResourceProfile(request *DeleteResourceProfileRequest) (_result *DeleteResourceProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteResourceProfileResponse{}
	_body, _err := client.DeleteResourceProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWorkGroupWithOptions(request *DeleteWorkGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteWorkGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteWorkGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteWorkGroup"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWorkGroup(request *DeleteWorkGroupRequest) (_result *DeleteWorkGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWorkGroupResponse{}
	_body, _err := client.DeleteWorkGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAlgoLibsWithOptions(request *DescribeAlgoLibsRequest, runtime *util.RuntimeOptions) (_result *DescribeAlgoLibsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAlgoLibsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAlgoLibs"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAlgoLibs(request *DescribeAlgoLibsRequest) (_result *DescribeAlgoLibsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlgoLibsResponse{}
	_body, _err := client.DescribeAlgoLibsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCamerasWithOptions(request *DescribeCamerasRequest, runtime *util.RuntimeOptions) (_result *DescribeCamerasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCamerasResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCameras"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCameras(request *DescribeCamerasRequest) (_result *DescribeCamerasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCamerasResponse{}
	_body, _err := client.DescribeCamerasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCapabilitiesWithOptions(request *DescribeCapabilitiesRequest, runtime *util.RuntimeOptions) (_result *DescribeCapabilitiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCapabilitiesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCapabilities"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCapabilities(request *DescribeCapabilitiesRequest) (_result *DescribeCapabilitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCapabilitiesResponse{}
	_body, _err := client.DescribeCapabilitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstances"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstances(request *DescribeInstancesRequest) (_result *DescribeInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DescribeInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeJobGroupsWithOptions(request *DescribeJobGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeJobGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeJobGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeJobGroups"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeJobGroups(request *DescribeJobGroupsRequest) (_result *DescribeJobGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeJobGroupsResponse{}
	_body, _err := client.DescribeJobGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeProtocolsWithOptions(request *DescribeProtocolsRequest, runtime *util.RuntimeOptions) (_result *DescribeProtocolsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeProtocolsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeProtocols"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeProtocols(request *DescribeProtocolsRequest) (_result *DescribeProtocolsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProtocolsResponse{}
	_body, _err := client.DescribeProtocolsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResourceProfilesWithOptions(request *DescribeResourceProfilesRequest, runtime *util.RuntimeOptions) (_result *DescribeResourceProfilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeResourceProfilesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeResourceProfiles"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResourceProfiles(request *DescribeResourceProfilesRequest) (_result *DescribeResourceProfilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResourceProfilesResponse{}
	_body, _err := client.DescribeResourceProfilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeStreamsWithOptions(request *DescribeStreamsRequest, runtime *util.RuntimeOptions) (_result *DescribeStreamsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeStreamsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeStreams"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeStreams(request *DescribeStreamsRequest) (_result *DescribeStreamsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeStreamsResponse{}
	_body, _err := client.DescribeStreamsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWorkGroupsWithOptions(request *DescribeWorkGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeWorkGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeWorkGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeWorkGroups"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWorkGroups(request *DescribeWorkGroupsRequest) (_result *DescribeWorkGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWorkGroupsResponse{}
	_body, _err := client.DescribeWorkGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachStreamWithOptions(request *DetachStreamRequest, runtime *util.RuntimeOptions) (_result *DetachStreamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetachStreamResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetachStream"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachStream(request *DetachStreamRequest) (_result *DetachStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachStreamResponse{}
	_body, _err := client.DetachStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCameraConfForCameraWithOptions(request *GetCameraConfForCameraRequest, runtime *util.RuntimeOptions) (_result *GetCameraConfForCameraResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetCameraConfForCameraResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetCameraConfForCamera"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCameraConfForCamera(request *GetCameraConfForCameraRequest) (_result *GetCameraConfForCameraResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCameraConfForCameraResponse{}
	_body, _err := client.GetCameraConfForCameraWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetComputeJobLogWithOptions(request *GetComputeJobLogRequest, runtime *util.RuntimeOptions) (_result *GetComputeJobLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetComputeJobLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetComputeJobLog"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetComputeJobLog(request *GetComputeJobLogRequest) (_result *GetComputeJobLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetComputeJobLogResponse{}
	_body, _err := client.GetComputeJobLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPlayUrlForCameraWithOptions(request *GetPlayUrlForCameraRequest, runtime *util.RuntimeOptions) (_result *GetPlayUrlForCameraResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPlayUrlForCameraResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPlayUrlForCamera"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPlayUrlForCamera(request *GetPlayUrlForCameraRequest) (_result *GetPlayUrlForCameraResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPlayUrlForCameraResponse{}
	_body, _err := client.GetPlayUrlForCameraWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListComputeJobLogsWithOptions(request *ListComputeJobLogsRequest, runtime *util.RuntimeOptions) (_result *ListComputeJobLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListComputeJobLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListComputeJobLogs"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListComputeJobLogs(request *ListComputeJobLogsRequest) (_result *ListComputeJobLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListComputeJobLogsResponse{}
	_body, _err := client.ListComputeJobLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListParkingResultsWithOptions(request *ListParkingResultsRequest, runtime *util.RuntimeOptions) (_result *ListParkingResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListParkingResultsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListParkingResults"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListParkingResults(request *ListParkingResultsRequest) (_result *ListParkingResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListParkingResultsResponse{}
	_body, _err := client.ListParkingResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSafetyHelmetResultsWithOptions(request *ListSafetyHelmetResultsRequest, runtime *util.RuntimeOptions) (_result *ListSafetyHelmetResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSafetyHelmetResultsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSafetyHelmetResults"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSafetyHelmetResults(request *ListSafetyHelmetResultsRequest) (_result *ListSafetyHelmetResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSafetyHelmetResultsResponse{}
	_body, _err := client.ListSafetyHelmetResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListStreamsForCamerasWithOptions(request *ListStreamsForCamerasRequest, runtime *util.RuntimeOptions) (_result *ListStreamsForCamerasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListStreamsForCamerasResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListStreamsForCameras"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListStreamsForCameras(request *ListStreamsForCamerasRequest) (_result *ListStreamsForCamerasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListStreamsForCamerasResponse{}
	_body, _err := client.ListStreamsForCamerasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVehicleEntryResultsWithOptions(request *ListVehicleEntryResultsRequest, runtime *util.RuntimeOptions) (_result *ListVehicleEntryResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListVehicleEntryResultsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListVehicleEntryResults"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVehicleEntryResults(request *ListVehicleEntryResultsRequest) (_result *ListVehicleEntryResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVehicleEntryResultsResponse{}
	_body, _err := client.ListVehicleEntryResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAlgoLibWithOptions(request *ModifyAlgoLibRequest, runtime *util.RuntimeOptions) (_result *ModifyAlgoLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyAlgoLibResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyAlgoLib"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAlgoLib(request *ModifyAlgoLibRequest) (_result *ModifyAlgoLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAlgoLibResponse{}
	_body, _err := client.ModifyAlgoLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyCameraWithOptions(request *ModifyCameraRequest, runtime *util.RuntimeOptions) (_result *ModifyCameraResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyCameraResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyCamera"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyCamera(request *ModifyCameraRequest) (_result *ModifyCameraResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyCameraResponse{}
	_body, _err := client.ModifyCameraWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyCapabilityWithOptions(request *ModifyCapabilityRequest, runtime *util.RuntimeOptions) (_result *ModifyCapabilityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyCapabilityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyCapability"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyCapability(request *ModifyCapabilityRequest) (_result *ModifyCapabilityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyCapabilityResponse{}
	_body, _err := client.ModifyCapabilityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceWithOptions(request *ModifyInstanceRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstance"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstance(request *ModifyInstanceRequest) (_result *ModifyInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceResponse{}
	_body, _err := client.ModifyInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyJobGroupWithOptions(request *ModifyJobGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyJobGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyJobGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyJobGroup"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyJobGroup(request *ModifyJobGroupRequest) (_result *ModifyJobGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyJobGroupResponse{}
	_body, _err := client.ModifyJobGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyResourceProfileWithOptions(request *ModifyResourceProfileRequest, runtime *util.RuntimeOptions) (_result *ModifyResourceProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyResourceProfileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyResourceProfile"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyResourceProfile(request *ModifyResourceProfileRequest) (_result *ModifyResourceProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyResourceProfileResponse{}
	_body, _err := client.ModifyResourceProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyWorkGroupWithOptions(request *ModifyWorkGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyWorkGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyWorkGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyWorkGroup"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyWorkGroup(request *ModifyWorkGroupRequest) (_result *ModifyWorkGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyWorkGroupResponse{}
	_body, _err := client.ModifyWorkGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutCameraConfForCameraWithOptions(request *PutCameraConfForCameraRequest, runtime *util.RuntimeOptions) (_result *PutCameraConfForCameraResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PutCameraConfForCameraResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PutCameraConfForCamera"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutCameraConfForCamera(request *PutCameraConfForCameraRequest) (_result *PutCameraConfForCameraResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutCameraConfForCameraResponse{}
	_body, _err := client.PutCameraConfForCameraWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchImagesWithOptions(request *SearchImagesRequest, runtime *util.RuntimeOptions) (_result *SearchImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchImagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchImages"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchImages(request *SearchImagesRequest) (_result *SearchImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchImagesResponse{}
	_body, _err := client.SearchImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartJobGroupWithOptions(request *StartJobGroupRequest, runtime *util.RuntimeOptions) (_result *StartJobGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartJobGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartJobGroup"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartJobGroup(request *StartJobGroupRequest) (_result *StartJobGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartJobGroupResponse{}
	_body, _err := client.StartJobGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopJobGroupWithOptions(request *StopJobGroupRequest, runtime *util.RuntimeOptions) (_result *StopJobGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopJobGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopJobGroup"), tea.String("2018-10-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopJobGroup(request *StopJobGroupRequest) (_result *StopJobGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopJobGroupResponse{}
	_body, _err := client.StopJobGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
