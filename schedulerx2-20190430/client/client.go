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

type BatchDeleteJobsRequest struct {
	GroupId         *string  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	JobIdList       []*int64 `json:"JobIdList,omitempty" xml:"JobIdList,omitempty" type:"Repeated"`
	Namespace       *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceSource *string  `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	RegionId        *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s BatchDeleteJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteJobsRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteJobsRequest) SetGroupId(v string) *BatchDeleteJobsRequest {
	s.GroupId = &v
	return s
}

func (s *BatchDeleteJobsRequest) SetJobIdList(v []*int64) *BatchDeleteJobsRequest {
	s.JobIdList = v
	return s
}

func (s *BatchDeleteJobsRequest) SetNamespace(v string) *BatchDeleteJobsRequest {
	s.Namespace = &v
	return s
}

func (s *BatchDeleteJobsRequest) SetNamespaceSource(v string) *BatchDeleteJobsRequest {
	s.NamespaceSource = &v
	return s
}

func (s *BatchDeleteJobsRequest) SetRegionId(v string) *BatchDeleteJobsRequest {
	s.RegionId = &v
	return s
}

type BatchDeleteJobsResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchDeleteJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteJobsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteJobsResponseBody) SetCode(v int32) *BatchDeleteJobsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchDeleteJobsResponseBody) SetMessage(v string) *BatchDeleteJobsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchDeleteJobsResponseBody) SetRequestId(v string) *BatchDeleteJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteJobsResponseBody) SetSuccess(v bool) *BatchDeleteJobsResponseBody {
	s.Success = &v
	return s
}

type BatchDeleteJobsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchDeleteJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchDeleteJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteJobsResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteJobsResponse) SetHeaders(v map[string]*string) *BatchDeleteJobsResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteJobsResponse) SetBody(v *BatchDeleteJobsResponseBody) *BatchDeleteJobsResponse {
	s.Body = v
	return s
}

type BatchDisableJobsRequest struct {
	GroupId         *string  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	JobIdList       []*int64 `json:"JobIdList,omitempty" xml:"JobIdList,omitempty" type:"Repeated"`
	Namespace       *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceSource *string  `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	RegionId        *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s BatchDisableJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDisableJobsRequest) GoString() string {
	return s.String()
}

func (s *BatchDisableJobsRequest) SetGroupId(v string) *BatchDisableJobsRequest {
	s.GroupId = &v
	return s
}

func (s *BatchDisableJobsRequest) SetJobIdList(v []*int64) *BatchDisableJobsRequest {
	s.JobIdList = v
	return s
}

func (s *BatchDisableJobsRequest) SetNamespace(v string) *BatchDisableJobsRequest {
	s.Namespace = &v
	return s
}

func (s *BatchDisableJobsRequest) SetNamespaceSource(v string) *BatchDisableJobsRequest {
	s.NamespaceSource = &v
	return s
}

func (s *BatchDisableJobsRequest) SetRegionId(v string) *BatchDisableJobsRequest {
	s.RegionId = &v
	return s
}

type BatchDisableJobsResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchDisableJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchDisableJobsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDisableJobsResponseBody) SetCode(v int32) *BatchDisableJobsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchDisableJobsResponseBody) SetMessage(v string) *BatchDisableJobsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchDisableJobsResponseBody) SetRequestId(v string) *BatchDisableJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDisableJobsResponseBody) SetSuccess(v bool) *BatchDisableJobsResponseBody {
	s.Success = &v
	return s
}

type BatchDisableJobsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchDisableJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchDisableJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDisableJobsResponse) GoString() string {
	return s.String()
}

func (s *BatchDisableJobsResponse) SetHeaders(v map[string]*string) *BatchDisableJobsResponse {
	s.Headers = v
	return s
}

func (s *BatchDisableJobsResponse) SetBody(v *BatchDisableJobsResponseBody) *BatchDisableJobsResponse {
	s.Body = v
	return s
}

type BatchEnableJobsRequest struct {
	GroupId         *string  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	JobIdList       []*int64 `json:"JobIdList,omitempty" xml:"JobIdList,omitempty" type:"Repeated"`
	Namespace       *string  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceSource *string  `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	RegionId        *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s BatchEnableJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchEnableJobsRequest) GoString() string {
	return s.String()
}

func (s *BatchEnableJobsRequest) SetGroupId(v string) *BatchEnableJobsRequest {
	s.GroupId = &v
	return s
}

func (s *BatchEnableJobsRequest) SetJobIdList(v []*int64) *BatchEnableJobsRequest {
	s.JobIdList = v
	return s
}

func (s *BatchEnableJobsRequest) SetNamespace(v string) *BatchEnableJobsRequest {
	s.Namespace = &v
	return s
}

func (s *BatchEnableJobsRequest) SetNamespaceSource(v string) *BatchEnableJobsRequest {
	s.NamespaceSource = &v
	return s
}

func (s *BatchEnableJobsRequest) SetRegionId(v string) *BatchEnableJobsRequest {
	s.RegionId = &v
	return s
}

type BatchEnableJobsResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchEnableJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchEnableJobsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchEnableJobsResponseBody) SetCode(v int32) *BatchEnableJobsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchEnableJobsResponseBody) SetMessage(v string) *BatchEnableJobsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchEnableJobsResponseBody) SetRequestId(v string) *BatchEnableJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchEnableJobsResponseBody) SetSuccess(v bool) *BatchEnableJobsResponseBody {
	s.Success = &v
	return s
}

type BatchEnableJobsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchEnableJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchEnableJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchEnableJobsResponse) GoString() string {
	return s.String()
}

func (s *BatchEnableJobsResponse) SetHeaders(v map[string]*string) *BatchEnableJobsResponse {
	s.Headers = v
	return s
}

func (s *BatchEnableJobsResponse) SetBody(v *BatchEnableJobsResponseBody) *BatchEnableJobsResponse {
	s.Body = v
	return s
}

type CreateAppGroupRequest struct {
	AlarmJson       *string `json:"AlarmJson,omitempty" xml:"AlarmJson,omitempty"`
	AppName         *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	MaxJobs         *int32  `json:"MaxJobs,omitempty" xml:"MaxJobs,omitempty"`
	Namespace       *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceName   *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateAppGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAppGroupRequest) SetAlarmJson(v string) *CreateAppGroupRequest {
	s.AlarmJson = &v
	return s
}

func (s *CreateAppGroupRequest) SetAppName(v string) *CreateAppGroupRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppGroupRequest) SetDescription(v string) *CreateAppGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateAppGroupRequest) SetGroupId(v string) *CreateAppGroupRequest {
	s.GroupId = &v
	return s
}

func (s *CreateAppGroupRequest) SetMaxJobs(v int32) *CreateAppGroupRequest {
	s.MaxJobs = &v
	return s
}

func (s *CreateAppGroupRequest) SetNamespace(v string) *CreateAppGroupRequest {
	s.Namespace = &v
	return s
}

func (s *CreateAppGroupRequest) SetNamespaceName(v string) *CreateAppGroupRequest {
	s.NamespaceName = &v
	return s
}

func (s *CreateAppGroupRequest) SetNamespaceSource(v string) *CreateAppGroupRequest {
	s.NamespaceSource = &v
	return s
}

func (s *CreateAppGroupRequest) SetRegionId(v string) *CreateAppGroupRequest {
	s.RegionId = &v
	return s
}

type CreateAppGroupResponseBody struct {
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateAppGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAppGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppGroupResponseBody) SetCode(v int32) *CreateAppGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAppGroupResponseBody) SetData(v *CreateAppGroupResponseBodyData) *CreateAppGroupResponseBody {
	s.Data = v
	return s
}

func (s *CreateAppGroupResponseBody) SetMessage(v string) *CreateAppGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAppGroupResponseBody) SetRequestId(v string) *CreateAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppGroupResponseBody) SetSuccess(v bool) *CreateAppGroupResponseBody {
	s.Success = &v
	return s
}

type CreateAppGroupResponseBodyData struct {
	AppGroupId *int64 `json:"AppGroupId,omitempty" xml:"AppGroupId,omitempty"`
}

func (s CreateAppGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAppGroupResponseBodyData) SetAppGroupId(v int64) *CreateAppGroupResponseBodyData {
	s.AppGroupId = &v
	return s
}

type CreateAppGroupResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateAppGroupResponse) SetHeaders(v map[string]*string) *CreateAppGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateAppGroupResponse) SetBody(v *CreateAppGroupResponseBody) *CreateAppGroupResponse {
	s.Body = v
	return s
}

type CreateJobRequest struct {
	AttemptInterval     *int32                         `json:"AttemptInterval,omitempty" xml:"AttemptInterval,omitempty"`
	Calendar            *string                        `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	ClassName           *string                        `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	ConsumerSize        *int32                         `json:"ConsumerSize,omitempty" xml:"ConsumerSize,omitempty"`
	ContactInfo         []*CreateJobRequestContactInfo `json:"ContactInfo,omitempty" xml:"ContactInfo,omitempty" type:"Repeated"`
	Content             *string                        `json:"Content,omitempty" xml:"Content,omitempty"`
	DataOffset          *int32                         `json:"DataOffset,omitempty" xml:"DataOffset,omitempty"`
	Description         *string                        `json:"Description,omitempty" xml:"Description,omitempty"`
	DispatcherSize      *int32                         `json:"DispatcherSize,omitempty" xml:"DispatcherSize,omitempty"`
	ExecuteMode         *string                        `json:"ExecuteMode,omitempty" xml:"ExecuteMode,omitempty"`
	FailEnable          *bool                          `json:"FailEnable,omitempty" xml:"FailEnable,omitempty"`
	GroupId             *string                        `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	JarUrl              *string                        `json:"JarUrl,omitempty" xml:"JarUrl,omitempty"`
	JobType             *string                        `json:"JobType,omitempty" xml:"JobType,omitempty"`
	MaxAttempt          *int32                         `json:"MaxAttempt,omitempty" xml:"MaxAttempt,omitempty"`
	MaxConcurrency      *int32                         `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	MissWorkerEnable    *bool                          `json:"MissWorkerEnable,omitempty" xml:"MissWorkerEnable,omitempty"`
	Name                *string                        `json:"Name,omitempty" xml:"Name,omitempty"`
	Namespace           *string                        `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceSource     *string                        `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	PageSize            *int32                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Parameters          *string                        `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	QueueSize           *int32                         `json:"QueueSize,omitempty" xml:"QueueSize,omitempty"`
	RegionId            *string                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SendChannel         *string                        `json:"SendChannel,omitempty" xml:"SendChannel,omitempty"`
	TaskAttemptInterval *int32                         `json:"TaskAttemptInterval,omitempty" xml:"TaskAttemptInterval,omitempty"`
	TaskMaxAttempt      *int32                         `json:"TaskMaxAttempt,omitempty" xml:"TaskMaxAttempt,omitempty"`
	TimeExpression      *string                        `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	TimeType            *int32                         `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	Timeout             *int64                         `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	TimeoutEnable       *bool                          `json:"TimeoutEnable,omitempty" xml:"TimeoutEnable,omitempty"`
	TimeoutKillEnable   *bool                          `json:"TimeoutKillEnable,omitempty" xml:"TimeoutKillEnable,omitempty"`
}

func (s CreateJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequest) GoString() string {
	return s.String()
}

func (s *CreateJobRequest) SetAttemptInterval(v int32) *CreateJobRequest {
	s.AttemptInterval = &v
	return s
}

func (s *CreateJobRequest) SetCalendar(v string) *CreateJobRequest {
	s.Calendar = &v
	return s
}

func (s *CreateJobRequest) SetClassName(v string) *CreateJobRequest {
	s.ClassName = &v
	return s
}

func (s *CreateJobRequest) SetConsumerSize(v int32) *CreateJobRequest {
	s.ConsumerSize = &v
	return s
}

func (s *CreateJobRequest) SetContactInfo(v []*CreateJobRequestContactInfo) *CreateJobRequest {
	s.ContactInfo = v
	return s
}

func (s *CreateJobRequest) SetContent(v string) *CreateJobRequest {
	s.Content = &v
	return s
}

func (s *CreateJobRequest) SetDataOffset(v int32) *CreateJobRequest {
	s.DataOffset = &v
	return s
}

func (s *CreateJobRequest) SetDescription(v string) *CreateJobRequest {
	s.Description = &v
	return s
}

func (s *CreateJobRequest) SetDispatcherSize(v int32) *CreateJobRequest {
	s.DispatcherSize = &v
	return s
}

func (s *CreateJobRequest) SetExecuteMode(v string) *CreateJobRequest {
	s.ExecuteMode = &v
	return s
}

func (s *CreateJobRequest) SetFailEnable(v bool) *CreateJobRequest {
	s.FailEnable = &v
	return s
}

func (s *CreateJobRequest) SetGroupId(v string) *CreateJobRequest {
	s.GroupId = &v
	return s
}

func (s *CreateJobRequest) SetJarUrl(v string) *CreateJobRequest {
	s.JarUrl = &v
	return s
}

func (s *CreateJobRequest) SetJobType(v string) *CreateJobRequest {
	s.JobType = &v
	return s
}

func (s *CreateJobRequest) SetMaxAttempt(v int32) *CreateJobRequest {
	s.MaxAttempt = &v
	return s
}

func (s *CreateJobRequest) SetMaxConcurrency(v int32) *CreateJobRequest {
	s.MaxConcurrency = &v
	return s
}

func (s *CreateJobRequest) SetMissWorkerEnable(v bool) *CreateJobRequest {
	s.MissWorkerEnable = &v
	return s
}

func (s *CreateJobRequest) SetName(v string) *CreateJobRequest {
	s.Name = &v
	return s
}

func (s *CreateJobRequest) SetNamespace(v string) *CreateJobRequest {
	s.Namespace = &v
	return s
}

func (s *CreateJobRequest) SetNamespaceSource(v string) *CreateJobRequest {
	s.NamespaceSource = &v
	return s
}

func (s *CreateJobRequest) SetPageSize(v int32) *CreateJobRequest {
	s.PageSize = &v
	return s
}

func (s *CreateJobRequest) SetParameters(v string) *CreateJobRequest {
	s.Parameters = &v
	return s
}

func (s *CreateJobRequest) SetQueueSize(v int32) *CreateJobRequest {
	s.QueueSize = &v
	return s
}

func (s *CreateJobRequest) SetRegionId(v string) *CreateJobRequest {
	s.RegionId = &v
	return s
}

func (s *CreateJobRequest) SetSendChannel(v string) *CreateJobRequest {
	s.SendChannel = &v
	return s
}

func (s *CreateJobRequest) SetTaskAttemptInterval(v int32) *CreateJobRequest {
	s.TaskAttemptInterval = &v
	return s
}

func (s *CreateJobRequest) SetTaskMaxAttempt(v int32) *CreateJobRequest {
	s.TaskMaxAttempt = &v
	return s
}

func (s *CreateJobRequest) SetTimeExpression(v string) *CreateJobRequest {
	s.TimeExpression = &v
	return s
}

func (s *CreateJobRequest) SetTimeType(v int32) *CreateJobRequest {
	s.TimeType = &v
	return s
}

func (s *CreateJobRequest) SetTimeout(v int64) *CreateJobRequest {
	s.Timeout = &v
	return s
}

func (s *CreateJobRequest) SetTimeoutEnable(v bool) *CreateJobRequest {
	s.TimeoutEnable = &v
	return s
}

func (s *CreateJobRequest) SetTimeoutKillEnable(v bool) *CreateJobRequest {
	s.TimeoutKillEnable = &v
	return s
}

type CreateJobRequestContactInfo struct {
	Ding      *string `json:"Ding,omitempty" xml:"Ding,omitempty"`
	UserMail  *string `json:"UserMail,omitempty" xml:"UserMail,omitempty"`
	UserName  *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserPhone *string `json:"UserPhone,omitempty" xml:"UserPhone,omitempty"`
}

func (s CreateJobRequestContactInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestContactInfo) GoString() string {
	return s.String()
}

func (s *CreateJobRequestContactInfo) SetDing(v string) *CreateJobRequestContactInfo {
	s.Ding = &v
	return s
}

func (s *CreateJobRequestContactInfo) SetUserMail(v string) *CreateJobRequestContactInfo {
	s.UserMail = &v
	return s
}

func (s *CreateJobRequestContactInfo) SetUserName(v string) *CreateJobRequestContactInfo {
	s.UserName = &v
	return s
}

func (s *CreateJobRequestContactInfo) SetUserPhone(v string) *CreateJobRequestContactInfo {
	s.UserPhone = &v
	return s
}

type CreateJobResponseBody struct {
	Code      *int32                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobResponseBody) SetCode(v int32) *CreateJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateJobResponseBody) SetData(v *CreateJobResponseBodyData) *CreateJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateJobResponseBody) SetMessage(v string) *CreateJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreateJobResponseBody) SetRequestId(v string) *CreateJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateJobResponseBody) SetSuccess(v bool) *CreateJobResponseBody {
	s.Success = &v
	return s
}

type CreateJobResponseBodyData struct {
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CreateJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateJobResponseBodyData) SetJobId(v int64) *CreateJobResponseBodyData {
	s.JobId = &v
	return s
}

type CreateJobResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateJobResponse) GoString() string {
	return s.String()
}

func (s *CreateJobResponse) SetHeaders(v map[string]*string) *CreateJobResponse {
	s.Headers = v
	return s
}

func (s *CreateJobResponse) SetBody(v *CreateJobResponseBody) *CreateJobResponse {
	s.Body = v
	return s
}

type DeleteJobRequest struct {
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	JobId           *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Namespace       *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobRequest) SetGroupId(v string) *DeleteJobRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteJobRequest) SetJobId(v int64) *DeleteJobRequest {
	s.JobId = &v
	return s
}

func (s *DeleteJobRequest) SetNamespace(v string) *DeleteJobRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteJobRequest) SetNamespaceSource(v string) *DeleteJobRequest {
	s.NamespaceSource = &v
	return s
}

func (s *DeleteJobRequest) SetRegionId(v string) *DeleteJobRequest {
	s.RegionId = &v
	return s
}

type DeleteJobResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobResponseBody) SetCode(v int32) *DeleteJobResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteJobResponseBody) SetMessage(v string) *DeleteJobResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteJobResponseBody) SetRequestId(v string) *DeleteJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteJobResponseBody) SetSuccess(v bool) *DeleteJobResponseBody {
	s.Success = &v
	return s
}

type DeleteJobResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteJobResponse) SetHeaders(v map[string]*string) *DeleteJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteJobResponse) SetBody(v *DeleteJobResponseBody) *DeleteJobResponse {
	s.Body = v
	return s
}

type DeleteWorkflowRequest struct {
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Namespace       *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	WorkflowId      *int64  `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s DeleteWorkflowRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkflowRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowRequest) SetGroupId(v string) *DeleteWorkflowRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteWorkflowRequest) SetNamespace(v string) *DeleteWorkflowRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteWorkflowRequest) SetNamespaceSource(v string) *DeleteWorkflowRequest {
	s.NamespaceSource = &v
	return s
}

func (s *DeleteWorkflowRequest) SetRegionId(v string) *DeleteWorkflowRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteWorkflowRequest) SetWorkflowId(v int64) *DeleteWorkflowRequest {
	s.WorkflowId = &v
	return s
}

type DeleteWorkflowResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteWorkflowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowResponseBody) SetCode(v int32) *DeleteWorkflowResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteWorkflowResponseBody) SetMessage(v string) *DeleteWorkflowResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteWorkflowResponseBody) SetRequestId(v string) *DeleteWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkflowResponseBody) SetSuccess(v bool) *DeleteWorkflowResponseBody {
	s.Success = &v
	return s
}

type DeleteWorkflowResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWorkflowResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkflowResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowResponse) SetHeaders(v map[string]*string) *DeleteWorkflowResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkflowResponse) SetBody(v *DeleteWorkflowResponseBody) *DeleteWorkflowResponse {
	s.Body = v
	return s
}

type DescribeRegionsResponseBody struct {
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	Regions   []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetCode(v int32) *DescribeRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetMessage(v string) *DescribeRegionsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetSuccess(v bool) *DescribeRegionsResponseBody {
	s.Success = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
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

type DesignateWorkersRequest struct {
	// 指定机器的类型
	DesignateType *int32 `json:"DesignateType,omitempty" xml:"DesignateType,omitempty"`
	// 应用分组ID
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// 任务ID
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 指定label列表json格式
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// 命名空间UID
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// 命名空间来源
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 是否故障转移
	Transferable *bool `json:"Transferable,omitempty" xml:"Transferable,omitempty"`
	// 指定机器列表json格式
	Workers *string `json:"Workers,omitempty" xml:"Workers,omitempty"`
}

func (s DesignateWorkersRequest) String() string {
	return tea.Prettify(s)
}

func (s DesignateWorkersRequest) GoString() string {
	return s.String()
}

func (s *DesignateWorkersRequest) SetDesignateType(v int32) *DesignateWorkersRequest {
	s.DesignateType = &v
	return s
}

func (s *DesignateWorkersRequest) SetGroupId(v string) *DesignateWorkersRequest {
	s.GroupId = &v
	return s
}

func (s *DesignateWorkersRequest) SetJobId(v int64) *DesignateWorkersRequest {
	s.JobId = &v
	return s
}

func (s *DesignateWorkersRequest) SetLabels(v string) *DesignateWorkersRequest {
	s.Labels = &v
	return s
}

func (s *DesignateWorkersRequest) SetNamespace(v string) *DesignateWorkersRequest {
	s.Namespace = &v
	return s
}

func (s *DesignateWorkersRequest) SetNamespaceSource(v string) *DesignateWorkersRequest {
	s.NamespaceSource = &v
	return s
}

func (s *DesignateWorkersRequest) SetRegionId(v string) *DesignateWorkersRequest {
	s.RegionId = &v
	return s
}

func (s *DesignateWorkersRequest) SetTransferable(v bool) *DesignateWorkersRequest {
	s.Transferable = &v
	return s
}

func (s *DesignateWorkersRequest) SetWorkers(v string) *DesignateWorkersRequest {
	s.Workers = &v
	return s
}

type DesignateWorkersResponseBody struct {
	// 错误码
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DesignateWorkersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DesignateWorkersResponseBody) GoString() string {
	return s.String()
}

func (s *DesignateWorkersResponseBody) SetCode(v int32) *DesignateWorkersResponseBody {
	s.Code = &v
	return s
}

func (s *DesignateWorkersResponseBody) SetMessage(v string) *DesignateWorkersResponseBody {
	s.Message = &v
	return s
}

func (s *DesignateWorkersResponseBody) SetRequestId(v string) *DesignateWorkersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DesignateWorkersResponseBody) SetSuccess(v bool) *DesignateWorkersResponseBody {
	s.Success = &v
	return s
}

type DesignateWorkersResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DesignateWorkersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DesignateWorkersResponse) String() string {
	return tea.Prettify(s)
}

func (s DesignateWorkersResponse) GoString() string {
	return s.String()
}

func (s *DesignateWorkersResponse) SetHeaders(v map[string]*string) *DesignateWorkersResponse {
	s.Headers = v
	return s
}

func (s *DesignateWorkersResponse) SetBody(v *DesignateWorkersResponseBody) *DesignateWorkersResponse {
	s.Body = v
	return s
}

type DisableJobRequest struct {
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	JobId           *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Namespace       *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableJobRequest) GoString() string {
	return s.String()
}

func (s *DisableJobRequest) SetGroupId(v string) *DisableJobRequest {
	s.GroupId = &v
	return s
}

func (s *DisableJobRequest) SetJobId(v int64) *DisableJobRequest {
	s.JobId = &v
	return s
}

func (s *DisableJobRequest) SetNamespace(v string) *DisableJobRequest {
	s.Namespace = &v
	return s
}

func (s *DisableJobRequest) SetNamespaceSource(v string) *DisableJobRequest {
	s.NamespaceSource = &v
	return s
}

func (s *DisableJobRequest) SetRegionId(v string) *DisableJobRequest {
	s.RegionId = &v
	return s
}

type DisableJobResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableJobResponseBody) GoString() string {
	return s.String()
}

func (s *DisableJobResponseBody) SetCode(v int32) *DisableJobResponseBody {
	s.Code = &v
	return s
}

func (s *DisableJobResponseBody) SetMessage(v string) *DisableJobResponseBody {
	s.Message = &v
	return s
}

func (s *DisableJobResponseBody) SetRequestId(v string) *DisableJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableJobResponseBody) SetSuccess(v bool) *DisableJobResponseBody {
	s.Success = &v
	return s
}

type DisableJobResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableJobResponse) GoString() string {
	return s.String()
}

func (s *DisableJobResponse) SetHeaders(v map[string]*string) *DisableJobResponse {
	s.Headers = v
	return s
}

func (s *DisableJobResponse) SetBody(v *DisableJobResponseBody) *DisableJobResponse {
	s.Body = v
	return s
}

type DisableWorkflowRequest struct {
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Namespace       *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	WorkflowId      *int64  `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s DisableWorkflowRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableWorkflowRequest) GoString() string {
	return s.String()
}

func (s *DisableWorkflowRequest) SetGroupId(v string) *DisableWorkflowRequest {
	s.GroupId = &v
	return s
}

func (s *DisableWorkflowRequest) SetNamespace(v string) *DisableWorkflowRequest {
	s.Namespace = &v
	return s
}

func (s *DisableWorkflowRequest) SetNamespaceSource(v string) *DisableWorkflowRequest {
	s.NamespaceSource = &v
	return s
}

func (s *DisableWorkflowRequest) SetRegionId(v string) *DisableWorkflowRequest {
	s.RegionId = &v
	return s
}

func (s *DisableWorkflowRequest) SetWorkflowId(v int64) *DisableWorkflowRequest {
	s.WorkflowId = &v
	return s
}

type DisableWorkflowResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableWorkflowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *DisableWorkflowResponseBody) SetCode(v int32) *DisableWorkflowResponseBody {
	s.Code = &v
	return s
}

func (s *DisableWorkflowResponseBody) SetMessage(v string) *DisableWorkflowResponseBody {
	s.Message = &v
	return s
}

func (s *DisableWorkflowResponseBody) SetRequestId(v string) *DisableWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableWorkflowResponseBody) SetSuccess(v bool) *DisableWorkflowResponseBody {
	s.Success = &v
	return s
}

type DisableWorkflowResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableWorkflowResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableWorkflowResponse) GoString() string {
	return s.String()
}

func (s *DisableWorkflowResponse) SetHeaders(v map[string]*string) *DisableWorkflowResponse {
	s.Headers = v
	return s
}

func (s *DisableWorkflowResponse) SetBody(v *DisableWorkflowResponseBody) *DisableWorkflowResponse {
	s.Body = v
	return s
}

type EnableJobRequest struct {
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	JobId           *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Namespace       *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableJobRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableJobRequest) GoString() string {
	return s.String()
}

func (s *EnableJobRequest) SetGroupId(v string) *EnableJobRequest {
	s.GroupId = &v
	return s
}

func (s *EnableJobRequest) SetJobId(v int64) *EnableJobRequest {
	s.JobId = &v
	return s
}

func (s *EnableJobRequest) SetNamespace(v string) *EnableJobRequest {
	s.Namespace = &v
	return s
}

func (s *EnableJobRequest) SetNamespaceSource(v string) *EnableJobRequest {
	s.NamespaceSource = &v
	return s
}

func (s *EnableJobRequest) SetRegionId(v string) *EnableJobRequest {
	s.RegionId = &v
	return s
}

type EnableJobResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableJobResponseBody) GoString() string {
	return s.String()
}

func (s *EnableJobResponseBody) SetCode(v int32) *EnableJobResponseBody {
	s.Code = &v
	return s
}

func (s *EnableJobResponseBody) SetMessage(v string) *EnableJobResponseBody {
	s.Message = &v
	return s
}

func (s *EnableJobResponseBody) SetRequestId(v string) *EnableJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableJobResponseBody) SetSuccess(v bool) *EnableJobResponseBody {
	s.Success = &v
	return s
}

type EnableJobResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableJobResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableJobResponse) GoString() string {
	return s.String()
}

func (s *EnableJobResponse) SetHeaders(v map[string]*string) *EnableJobResponse {
	s.Headers = v
	return s
}

func (s *EnableJobResponse) SetBody(v *EnableJobResponseBody) *EnableJobResponse {
	s.Body = v
	return s
}

type EnableWorkflowRequest struct {
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Namespace       *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	WorkflowId      *int64  `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s EnableWorkflowRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableWorkflowRequest) GoString() string {
	return s.String()
}

func (s *EnableWorkflowRequest) SetGroupId(v string) *EnableWorkflowRequest {
	s.GroupId = &v
	return s
}

func (s *EnableWorkflowRequest) SetNamespace(v string) *EnableWorkflowRequest {
	s.Namespace = &v
	return s
}

func (s *EnableWorkflowRequest) SetNamespaceSource(v string) *EnableWorkflowRequest {
	s.NamespaceSource = &v
	return s
}

func (s *EnableWorkflowRequest) SetRegionId(v string) *EnableWorkflowRequest {
	s.RegionId = &v
	return s
}

func (s *EnableWorkflowRequest) SetWorkflowId(v int64) *EnableWorkflowRequest {
	s.WorkflowId = &v
	return s
}

type EnableWorkflowResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableWorkflowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *EnableWorkflowResponseBody) SetCode(v int32) *EnableWorkflowResponseBody {
	s.Code = &v
	return s
}

func (s *EnableWorkflowResponseBody) SetMessage(v string) *EnableWorkflowResponseBody {
	s.Message = &v
	return s
}

func (s *EnableWorkflowResponseBody) SetRequestId(v string) *EnableWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableWorkflowResponseBody) SetSuccess(v bool) *EnableWorkflowResponseBody {
	s.Success = &v
	return s
}

type EnableWorkflowResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableWorkflowResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableWorkflowResponse) GoString() string {
	return s.String()
}

func (s *EnableWorkflowResponse) SetHeaders(v map[string]*string) *EnableWorkflowResponse {
	s.Headers = v
	return s
}

func (s *EnableWorkflowResponse) SetBody(v *EnableWorkflowResponseBody) *EnableWorkflowResponse {
	s.Body = v
	return s
}

type ExecuteJobRequest struct {
	CheckJobStatus *bool `json:"CheckJobStatus,omitempty" xml:"CheckJobStatus,omitempty"`
	// 指定机器类型：1.workerAddr; 2. label
	DesignateType      *int32  `json:"DesignateType,omitempty" xml:"DesignateType,omitempty"`
	GroupId            *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceParameters *string `json:"InstanceParameters,omitempty" xml:"InstanceParameters,omitempty"`
	JobId              *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 指定机器的标签
	Label           *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Namespace       *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 指定机器的workerAddr
	Worker *string `json:"Worker,omitempty" xml:"Worker,omitempty"`
}

func (s ExecuteJobRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteJobRequest) GoString() string {
	return s.String()
}

func (s *ExecuteJobRequest) SetCheckJobStatus(v bool) *ExecuteJobRequest {
	s.CheckJobStatus = &v
	return s
}

func (s *ExecuteJobRequest) SetDesignateType(v int32) *ExecuteJobRequest {
	s.DesignateType = &v
	return s
}

func (s *ExecuteJobRequest) SetGroupId(v string) *ExecuteJobRequest {
	s.GroupId = &v
	return s
}

func (s *ExecuteJobRequest) SetInstanceParameters(v string) *ExecuteJobRequest {
	s.InstanceParameters = &v
	return s
}

func (s *ExecuteJobRequest) SetJobId(v int64) *ExecuteJobRequest {
	s.JobId = &v
	return s
}

func (s *ExecuteJobRequest) SetLabel(v string) *ExecuteJobRequest {
	s.Label = &v
	return s
}

func (s *ExecuteJobRequest) SetNamespace(v string) *ExecuteJobRequest {
	s.Namespace = &v
	return s
}

func (s *ExecuteJobRequest) SetNamespaceSource(v string) *ExecuteJobRequest {
	s.NamespaceSource = &v
	return s
}

func (s *ExecuteJobRequest) SetRegionId(v string) *ExecuteJobRequest {
	s.RegionId = &v
	return s
}

func (s *ExecuteJobRequest) SetWorker(v string) *ExecuteJobRequest {
	s.Worker = &v
	return s
}

type ExecuteJobResponseBody struct {
	Code      *int32                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ExecuteJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecuteJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteJobResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteJobResponseBody) SetCode(v int32) *ExecuteJobResponseBody {
	s.Code = &v
	return s
}

func (s *ExecuteJobResponseBody) SetData(v *ExecuteJobResponseBodyData) *ExecuteJobResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteJobResponseBody) SetMessage(v string) *ExecuteJobResponseBody {
	s.Message = &v
	return s
}

func (s *ExecuteJobResponseBody) SetRequestId(v string) *ExecuteJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteJobResponseBody) SetSuccess(v bool) *ExecuteJobResponseBody {
	s.Success = &v
	return s
}

type ExecuteJobResponseBodyData struct {
	JobInstanceId *int64 `json:"JobInstanceId,omitempty" xml:"JobInstanceId,omitempty"`
}

func (s ExecuteJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExecuteJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExecuteJobResponseBodyData) SetJobInstanceId(v int64) *ExecuteJobResponseBodyData {
	s.JobInstanceId = &v
	return s
}

type ExecuteJobResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExecuteJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteJobResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteJobResponse) GoString() string {
	return s.String()
}

func (s *ExecuteJobResponse) SetHeaders(v map[string]*string) *ExecuteJobResponse {
	s.Headers = v
	return s
}

func (s *ExecuteJobResponse) SetBody(v *ExecuteJobResponseBody) *ExecuteJobResponse {
	s.Body = v
	return s
}

type ExecuteWorkflowRequest struct {
	GroupId            *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceParameters *string `json:"InstanceParameters,omitempty" xml:"InstanceParameters,omitempty"`
	Namespace          *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceSource    *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	WorkflowId         *int64  `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s ExecuteWorkflowRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteWorkflowRequest) GoString() string {
	return s.String()
}

func (s *ExecuteWorkflowRequest) SetGroupId(v string) *ExecuteWorkflowRequest {
	s.GroupId = &v
	return s
}

func (s *ExecuteWorkflowRequest) SetInstanceParameters(v string) *ExecuteWorkflowRequest {
	s.InstanceParameters = &v
	return s
}

func (s *ExecuteWorkflowRequest) SetNamespace(v string) *ExecuteWorkflowRequest {
	s.Namespace = &v
	return s
}

func (s *ExecuteWorkflowRequest) SetNamespaceSource(v string) *ExecuteWorkflowRequest {
	s.NamespaceSource = &v
	return s
}

func (s *ExecuteWorkflowRequest) SetWorkflowId(v int64) *ExecuteWorkflowRequest {
	s.WorkflowId = &v
	return s
}

type ExecuteWorkflowResponseBody struct {
	Code      *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ExecuteWorkflowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecuteWorkflowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteWorkflowResponseBody) SetCode(v int32) *ExecuteWorkflowResponseBody {
	s.Code = &v
	return s
}

func (s *ExecuteWorkflowResponseBody) SetData(v *ExecuteWorkflowResponseBodyData) *ExecuteWorkflowResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteWorkflowResponseBody) SetMessage(v string) *ExecuteWorkflowResponseBody {
	s.Message = &v
	return s
}

func (s *ExecuteWorkflowResponseBody) SetRequestId(v string) *ExecuteWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteWorkflowResponseBody) SetSuccess(v bool) *ExecuteWorkflowResponseBody {
	s.Success = &v
	return s
}

type ExecuteWorkflowResponseBodyData struct {
	WfInstanceId *int64 `json:"WfInstanceId,omitempty" xml:"WfInstanceId,omitempty"`
}

func (s ExecuteWorkflowResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExecuteWorkflowResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExecuteWorkflowResponseBodyData) SetWfInstanceId(v int64) *ExecuteWorkflowResponseBodyData {
	s.WfInstanceId = &v
	return s
}

type ExecuteWorkflowResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExecuteWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteWorkflowResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteWorkflowResponse) GoString() string {
	return s.String()
}

func (s *ExecuteWorkflowResponse) SetHeaders(v map[string]*string) *ExecuteWorkflowResponse {
	s.Headers = v
	return s
}

func (s *ExecuteWorkflowResponse) SetBody(v *ExecuteWorkflowResponseBody) *ExecuteWorkflowResponse {
	s.Body = v
	return s
}

type GetJobInfoRequest struct {
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	JobId           *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Namespace       *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetJobInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobInfoRequest) GoString() string {
	return s.String()
}

func (s *GetJobInfoRequest) SetGroupId(v string) *GetJobInfoRequest {
	s.GroupId = &v
	return s
}

func (s *GetJobInfoRequest) SetJobId(v int64) *GetJobInfoRequest {
	s.JobId = &v
	return s
}

func (s *GetJobInfoRequest) SetNamespace(v string) *GetJobInfoRequest {
	s.Namespace = &v
	return s
}

func (s *GetJobInfoRequest) SetNamespaceSource(v string) *GetJobInfoRequest {
	s.NamespaceSource = &v
	return s
}

func (s *GetJobInfoRequest) SetRegionId(v string) *GetJobInfoRequest {
	s.RegionId = &v
	return s
}

type GetJobInfoResponseBody struct {
	Code      *int32                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetJobInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetJobInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponseBody) SetCode(v int32) *GetJobInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetJobInfoResponseBody) SetData(v *GetJobInfoResponseBodyData) *GetJobInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetJobInfoResponseBody) SetMessage(v string) *GetJobInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetJobInfoResponseBody) SetRequestId(v string) *GetJobInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobInfoResponseBody) SetSuccess(v bool) *GetJobInfoResponseBody {
	s.Success = &v
	return s
}

type GetJobInfoResponseBodyData struct {
	JobConfigInfo *GetJobInfoResponseBodyDataJobConfigInfo `json:"JobConfigInfo,omitempty" xml:"JobConfigInfo,omitempty" type:"Struct"`
}

func (s GetJobInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetJobInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponseBodyData) SetJobConfigInfo(v *GetJobInfoResponseBodyDataJobConfigInfo) *GetJobInfoResponseBodyData {
	s.JobConfigInfo = v
	return s
}

type GetJobInfoResponseBodyDataJobConfigInfo struct {
	AttemptInterval *int32                                                 `json:"AttemptInterval,omitempty" xml:"AttemptInterval,omitempty"`
	ClassName       *string                                                `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	Content         *string                                                `json:"Content,omitempty" xml:"Content,omitempty"`
	Description     *string                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	ExecuteMode     *string                                                `json:"ExecuteMode,omitempty" xml:"ExecuteMode,omitempty"`
	JarUrl          *string                                                `json:"JarUrl,omitempty" xml:"JarUrl,omitempty"`
	JobMonitorInfo  *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfo `json:"JobMonitorInfo,omitempty" xml:"JobMonitorInfo,omitempty" type:"Struct"`
	MapTaskXAttrs   *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs  `json:"MapTaskXAttrs,omitempty" xml:"MapTaskXAttrs,omitempty" type:"Struct"`
	MaxAttempt      *int32                                                 `json:"MaxAttempt,omitempty" xml:"MaxAttempt,omitempty"`
	MaxConcurrency  *string                                                `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	Name            *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	Parameters      *string                                                `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	Status          *int32                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	TimeConfig      *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig     `json:"TimeConfig,omitempty" xml:"TimeConfig,omitempty" type:"Struct"`
}

func (s GetJobInfoResponseBodyDataJobConfigInfo) String() string {
	return tea.Prettify(s)
}

func (s GetJobInfoResponseBodyDataJobConfigInfo) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetAttemptInterval(v int32) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.AttemptInterval = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetClassName(v string) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.ClassName = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetContent(v string) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.Content = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetDescription(v string) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.Description = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetExecuteMode(v string) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.ExecuteMode = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetJarUrl(v string) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.JarUrl = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetJobMonitorInfo(v *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfo) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.JobMonitorInfo = v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetMapTaskXAttrs(v *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.MapTaskXAttrs = v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetMaxAttempt(v int32) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.MaxAttempt = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetMaxConcurrency(v string) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.MaxConcurrency = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetName(v string) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.Name = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetParameters(v string) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.Parameters = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetStatus(v int32) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.Status = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetTimeConfig(v *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.TimeConfig = v
	return s
}

type GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfo struct {
	ContactInfo   []*GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo `json:"ContactInfo,omitempty" xml:"ContactInfo,omitempty" type:"Repeated"`
	MonitorConfig *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig `json:"MonitorConfig,omitempty" xml:"MonitorConfig,omitempty" type:"Struct"`
}

func (s GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfo) String() string {
	return tea.Prettify(s)
}

func (s GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfo) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfo) SetContactInfo(v []*GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo) *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfo {
	s.ContactInfo = v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfo) SetMonitorConfig(v *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig) *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfo {
	s.MonitorConfig = v
	return s
}

type GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo struct {
	Ding      *string `json:"Ding,omitempty" xml:"Ding,omitempty"`
	UserMail  *string `json:"UserMail,omitempty" xml:"UserMail,omitempty"`
	UserName  *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserPhone *string `json:"UserPhone,omitempty" xml:"UserPhone,omitempty"`
}

func (s GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo) String() string {
	return tea.Prettify(s)
}

func (s GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo) SetDing(v string) *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo {
	s.Ding = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo) SetUserMail(v string) *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo {
	s.UserMail = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo) SetUserName(v string) *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo {
	s.UserName = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo) SetUserPhone(v string) *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo {
	s.UserPhone = &v
	return s
}

type GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig struct {
	FailEnable        *bool   `json:"FailEnable,omitempty" xml:"FailEnable,omitempty"`
	MissWorkerEnable  *bool   `json:"MissWorkerEnable,omitempty" xml:"MissWorkerEnable,omitempty"`
	SendChannel       *string `json:"SendChannel,omitempty" xml:"SendChannel,omitempty"`
	Timeout           *int64  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	TimeoutEnable     *bool   `json:"TimeoutEnable,omitempty" xml:"TimeoutEnable,omitempty"`
	TimeoutKillEnable *bool   `json:"TimeoutKillEnable,omitempty" xml:"TimeoutKillEnable,omitempty"`
}

func (s GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig) String() string {
	return tea.Prettify(s)
}

func (s GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig) SetFailEnable(v bool) *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig {
	s.FailEnable = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig) SetMissWorkerEnable(v bool) *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig {
	s.MissWorkerEnable = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig) SetSendChannel(v string) *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig {
	s.SendChannel = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig) SetTimeout(v int64) *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig {
	s.Timeout = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig) SetTimeoutEnable(v bool) *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig {
	s.TimeoutEnable = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig) SetTimeoutKillEnable(v bool) *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig {
	s.TimeoutKillEnable = &v
	return s
}

type GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs struct {
	ConsumerSize        *int32 `json:"ConsumerSize,omitempty" xml:"ConsumerSize,omitempty"`
	DispatcherSize      *int32 `json:"DispatcherSize,omitempty" xml:"DispatcherSize,omitempty"`
	PageSize            *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueueSize           *int32 `json:"QueueSize,omitempty" xml:"QueueSize,omitempty"`
	TaskAttemptInterval *int32 `json:"TaskAttemptInterval,omitempty" xml:"TaskAttemptInterval,omitempty"`
	TaskMaxAttempt      *int32 `json:"TaskMaxAttempt,omitempty" xml:"TaskMaxAttempt,omitempty"`
}

func (s GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs) String() string {
	return tea.Prettify(s)
}

func (s GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs) SetConsumerSize(v int32) *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs {
	s.ConsumerSize = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs) SetDispatcherSize(v int32) *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs {
	s.DispatcherSize = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs) SetPageSize(v int32) *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs {
	s.PageSize = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs) SetQueueSize(v int32) *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs {
	s.QueueSize = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs) SetTaskAttemptInterval(v int32) *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs {
	s.TaskAttemptInterval = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs) SetTaskMaxAttempt(v int32) *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs {
	s.TaskMaxAttempt = &v
	return s
}

type GetJobInfoResponseBodyDataJobConfigInfoTimeConfig struct {
	Calendar       *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	DataOffset     *int32  `json:"DataOffset,omitempty" xml:"DataOffset,omitempty"`
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	TimeType       *int32  `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
}

func (s GetJobInfoResponseBodyDataJobConfigInfoTimeConfig) String() string {
	return tea.Prettify(s)
}

func (s GetJobInfoResponseBodyDataJobConfigInfoTimeConfig) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig) SetCalendar(v string) *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig {
	s.Calendar = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig) SetDataOffset(v int32) *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig {
	s.DataOffset = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig) SetTimeExpression(v string) *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig {
	s.TimeExpression = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig) SetTimeType(v int32) *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig {
	s.TimeType = &v
	return s
}

type GetJobInfoResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetJobInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobInfoResponse) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponse) SetHeaders(v map[string]*string) *GetJobInfoResponse {
	s.Headers = v
	return s
}

func (s *GetJobInfoResponse) SetBody(v *GetJobInfoResponseBody) *GetJobInfoResponse {
	s.Body = v
	return s
}

type GetJobInstanceRequest struct {
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	JobId           *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobInstanceId   *int64  `json:"JobInstanceId,omitempty" xml:"JobInstanceId,omitempty"`
	Namespace       *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
}

func (s GetJobInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetJobInstanceRequest) SetGroupId(v string) *GetJobInstanceRequest {
	s.GroupId = &v
	return s
}

func (s *GetJobInstanceRequest) SetJobId(v int64) *GetJobInstanceRequest {
	s.JobId = &v
	return s
}

func (s *GetJobInstanceRequest) SetJobInstanceId(v int64) *GetJobInstanceRequest {
	s.JobInstanceId = &v
	return s
}

func (s *GetJobInstanceRequest) SetNamespace(v string) *GetJobInstanceRequest {
	s.Namespace = &v
	return s
}

func (s *GetJobInstanceRequest) SetNamespaceSource(v string) *GetJobInstanceRequest {
	s.NamespaceSource = &v
	return s
}

type GetJobInstanceResponseBody struct {
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetJobInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetJobInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobInstanceResponseBody) SetCode(v int32) *GetJobInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetJobInstanceResponseBody) SetData(v *GetJobInstanceResponseBodyData) *GetJobInstanceResponseBody {
	s.Data = v
	return s
}

func (s *GetJobInstanceResponseBody) SetMessage(v string) *GetJobInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetJobInstanceResponseBody) SetRequestId(v string) *GetJobInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobInstanceResponseBody) SetSuccess(v bool) *GetJobInstanceResponseBody {
	s.Success = &v
	return s
}

type GetJobInstanceResponseBodyData struct {
	JobInstanceDetail *GetJobInstanceResponseBodyDataJobInstanceDetail `json:"JobInstanceDetail,omitempty" xml:"JobInstanceDetail,omitempty" type:"Struct"`
}

func (s GetJobInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetJobInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetJobInstanceResponseBodyData) SetJobInstanceDetail(v *GetJobInstanceResponseBodyDataJobInstanceDetail) *GetJobInstanceResponseBodyData {
	s.JobInstanceDetail = v
	return s
}

type GetJobInstanceResponseBodyDataJobInstanceDetail struct {
	DataTime     *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Executor     *string `json:"Executor,omitempty" xml:"Executor,omitempty"`
	InstanceId   *int64  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId        *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Result       *string `json:"Result,omitempty" xml:"Result,omitempty"`
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status       *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TimeType     *int32  `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	TriggerType  *int32  `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	WorkAddr     *string `json:"WorkAddr,omitempty" xml:"WorkAddr,omitempty"`
}

func (s GetJobInstanceResponseBodyDataJobInstanceDetail) String() string {
	return tea.Prettify(s)
}

func (s GetJobInstanceResponseBodyDataJobInstanceDetail) GoString() string {
	return s.String()
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetDataTime(v string) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.DataTime = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetEndTime(v string) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.EndTime = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetExecutor(v string) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.Executor = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetInstanceId(v int64) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.InstanceId = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetJobId(v int64) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.JobId = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetProgress(v string) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.Progress = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetResult(v string) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.Result = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetScheduleTime(v string) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.ScheduleTime = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetStartTime(v string) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.StartTime = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetStatus(v int32) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.Status = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetTimeType(v int32) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.TimeType = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetTriggerType(v int32) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.TriggerType = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetWorkAddr(v string) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.WorkAddr = &v
	return s
}

type GetJobInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetJobInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetJobInstanceResponse) SetHeaders(v map[string]*string) *GetJobInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetJobInstanceResponse) SetBody(v *GetJobInstanceResponseBody) *GetJobInstanceResponse {
	s.Body = v
	return s
}

type GetJobInstanceListRequest struct {
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	JobId           *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Namespace       *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetJobInstanceListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobInstanceListRequest) GoString() string {
	return s.String()
}

func (s *GetJobInstanceListRequest) SetGroupId(v string) *GetJobInstanceListRequest {
	s.GroupId = &v
	return s
}

func (s *GetJobInstanceListRequest) SetJobId(v int64) *GetJobInstanceListRequest {
	s.JobId = &v
	return s
}

func (s *GetJobInstanceListRequest) SetNamespace(v string) *GetJobInstanceListRequest {
	s.Namespace = &v
	return s
}

func (s *GetJobInstanceListRequest) SetNamespaceSource(v string) *GetJobInstanceListRequest {
	s.NamespaceSource = &v
	return s
}

func (s *GetJobInstanceListRequest) SetRegionId(v string) *GetJobInstanceListRequest {
	s.RegionId = &v
	return s
}

type GetJobInstanceListResponseBody struct {
	Code      *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetJobInstanceListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetJobInstanceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobInstanceListResponseBody) SetCode(v int32) *GetJobInstanceListResponseBody {
	s.Code = &v
	return s
}

func (s *GetJobInstanceListResponseBody) SetData(v *GetJobInstanceListResponseBodyData) *GetJobInstanceListResponseBody {
	s.Data = v
	return s
}

func (s *GetJobInstanceListResponseBody) SetMessage(v string) *GetJobInstanceListResponseBody {
	s.Message = &v
	return s
}

func (s *GetJobInstanceListResponseBody) SetRequestId(v string) *GetJobInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobInstanceListResponseBody) SetSuccess(v bool) *GetJobInstanceListResponseBody {
	s.Success = &v
	return s
}

type GetJobInstanceListResponseBodyData struct {
	JobInstanceDetails []*GetJobInstanceListResponseBodyDataJobInstanceDetails `json:"JobInstanceDetails,omitempty" xml:"JobInstanceDetails,omitempty" type:"Repeated"`
}

func (s GetJobInstanceListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetJobInstanceListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetJobInstanceListResponseBodyData) SetJobInstanceDetails(v []*GetJobInstanceListResponseBodyDataJobInstanceDetails) *GetJobInstanceListResponseBodyData {
	s.JobInstanceDetails = v
	return s
}

type GetJobInstanceListResponseBodyDataJobInstanceDetails struct {
	DataTime     *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Executor     *string `json:"Executor,omitempty" xml:"Executor,omitempty"`
	InstanceId   *int64  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId        *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Result       *string `json:"Result,omitempty" xml:"Result,omitempty"`
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status       *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TimeType     *int32  `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	TriggerType  *int32  `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	WorkAddr     *string `json:"WorkAddr,omitempty" xml:"WorkAddr,omitempty"`
}

func (s GetJobInstanceListResponseBodyDataJobInstanceDetails) String() string {
	return tea.Prettify(s)
}

func (s GetJobInstanceListResponseBodyDataJobInstanceDetails) GoString() string {
	return s.String()
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) SetDataTime(v string) *GetJobInstanceListResponseBodyDataJobInstanceDetails {
	s.DataTime = &v
	return s
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) SetEndTime(v string) *GetJobInstanceListResponseBodyDataJobInstanceDetails {
	s.EndTime = &v
	return s
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) SetExecutor(v string) *GetJobInstanceListResponseBodyDataJobInstanceDetails {
	s.Executor = &v
	return s
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) SetInstanceId(v int64) *GetJobInstanceListResponseBodyDataJobInstanceDetails {
	s.InstanceId = &v
	return s
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) SetJobId(v int64) *GetJobInstanceListResponseBodyDataJobInstanceDetails {
	s.JobId = &v
	return s
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) SetProgress(v string) *GetJobInstanceListResponseBodyDataJobInstanceDetails {
	s.Progress = &v
	return s
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) SetResult(v string) *GetJobInstanceListResponseBodyDataJobInstanceDetails {
	s.Result = &v
	return s
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) SetScheduleTime(v string) *GetJobInstanceListResponseBodyDataJobInstanceDetails {
	s.ScheduleTime = &v
	return s
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) SetStartTime(v string) *GetJobInstanceListResponseBodyDataJobInstanceDetails {
	s.StartTime = &v
	return s
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) SetStatus(v int32) *GetJobInstanceListResponseBodyDataJobInstanceDetails {
	s.Status = &v
	return s
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) SetTimeType(v int32) *GetJobInstanceListResponseBodyDataJobInstanceDetails {
	s.TimeType = &v
	return s
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) SetTriggerType(v int32) *GetJobInstanceListResponseBodyDataJobInstanceDetails {
	s.TriggerType = &v
	return s
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) SetWorkAddr(v string) *GetJobInstanceListResponseBodyDataJobInstanceDetails {
	s.WorkAddr = &v
	return s
}

type GetJobInstanceListResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetJobInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobInstanceListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobInstanceListResponse) GoString() string {
	return s.String()
}

func (s *GetJobInstanceListResponse) SetHeaders(v map[string]*string) *GetJobInstanceListResponse {
	s.Headers = v
	return s
}

func (s *GetJobInstanceListResponse) SetBody(v *GetJobInstanceListResponseBody) *GetJobInstanceListResponse {
	s.Body = v
	return s
}

type GetWorkFlowRequest struct {
	// 应用分组ID
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// 命名空间uid
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// 命名空间来源
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 工作流ID
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s GetWorkFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWorkFlowRequest) GoString() string {
	return s.String()
}

func (s *GetWorkFlowRequest) SetGroupId(v string) *GetWorkFlowRequest {
	s.GroupId = &v
	return s
}

func (s *GetWorkFlowRequest) SetNamespace(v string) *GetWorkFlowRequest {
	s.Namespace = &v
	return s
}

func (s *GetWorkFlowRequest) SetNamespaceSource(v string) *GetWorkFlowRequest {
	s.NamespaceSource = &v
	return s
}

func (s *GetWorkFlowRequest) SetRegionId(v string) *GetWorkFlowRequest {
	s.RegionId = &v
	return s
}

func (s *GetWorkFlowRequest) SetWorkflowId(v int64) *GetWorkFlowRequest {
	s.WorkflowId = &v
	return s
}

type GetWorkFlowResponseBody struct {
	// 错误码
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 工作流的数据
	Data *GetWorkFlowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 会否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWorkFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkFlowResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkFlowResponseBody) SetCode(v int32) *GetWorkFlowResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkFlowResponseBody) SetData(v *GetWorkFlowResponseBodyData) *GetWorkFlowResponseBody {
	s.Data = v
	return s
}

func (s *GetWorkFlowResponseBody) SetMessage(v string) *GetWorkFlowResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkFlowResponseBody) SetRequestId(v string) *GetWorkFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkFlowResponseBody) SetSuccess(v bool) *GetWorkFlowResponseBody {
	s.Success = &v
	return s
}

type GetWorkFlowResponseBodyData struct {
	// 工作流基本信息
	WorkFlowInfo *GetWorkFlowResponseBodyDataWorkFlowInfo `json:"WorkFlowInfo,omitempty" xml:"WorkFlowInfo,omitempty" type:"Struct"`
	// 工作流节点信息
	WorkFlowNodeInfo *GetWorkFlowResponseBodyDataWorkFlowNodeInfo `json:"WorkFlowNodeInfo,omitempty" xml:"WorkFlowNodeInfo,omitempty" type:"Struct"`
}

func (s GetWorkFlowResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetWorkFlowResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWorkFlowResponseBodyData) SetWorkFlowInfo(v *GetWorkFlowResponseBodyDataWorkFlowInfo) *GetWorkFlowResponseBodyData {
	s.WorkFlowInfo = v
	return s
}

func (s *GetWorkFlowResponseBodyData) SetWorkFlowNodeInfo(v *GetWorkFlowResponseBodyDataWorkFlowNodeInfo) *GetWorkFlowResponseBodyData {
	s.WorkFlowNodeInfo = v
	return s
}

type GetWorkFlowResponseBodyDataWorkFlowInfo struct {
	// 工作流描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 工作流名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 工作流状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 工作流时间表达式
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// 工作流时间类型
	TimeType *string `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// 工作流ID
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s GetWorkFlowResponseBodyDataWorkFlowInfo) String() string {
	return tea.Prettify(s)
}

func (s GetWorkFlowResponseBodyDataWorkFlowInfo) GoString() string {
	return s.String()
}

func (s *GetWorkFlowResponseBodyDataWorkFlowInfo) SetDescription(v string) *GetWorkFlowResponseBodyDataWorkFlowInfo {
	s.Description = &v
	return s
}

func (s *GetWorkFlowResponseBodyDataWorkFlowInfo) SetName(v string) *GetWorkFlowResponseBodyDataWorkFlowInfo {
	s.Name = &v
	return s
}

func (s *GetWorkFlowResponseBodyDataWorkFlowInfo) SetStatus(v string) *GetWorkFlowResponseBodyDataWorkFlowInfo {
	s.Status = &v
	return s
}

func (s *GetWorkFlowResponseBodyDataWorkFlowInfo) SetTimeExpression(v string) *GetWorkFlowResponseBodyDataWorkFlowInfo {
	s.TimeExpression = &v
	return s
}

func (s *GetWorkFlowResponseBodyDataWorkFlowInfo) SetTimeType(v string) *GetWorkFlowResponseBodyDataWorkFlowInfo {
	s.TimeType = &v
	return s
}

func (s *GetWorkFlowResponseBodyDataWorkFlowInfo) SetWorkflowId(v int64) *GetWorkFlowResponseBodyDataWorkFlowInfo {
	s.WorkflowId = &v
	return s
}

type GetWorkFlowResponseBodyDataWorkFlowNodeInfo struct {
	// 工作流边列表
	Edges []*GetWorkFlowResponseBodyDataWorkFlowNodeInfoEdges `json:"Edges,omitempty" xml:"Edges,omitempty" type:"Repeated"`
	// 工作流节点列表
	Nodes []*GetWorkFlowResponseBodyDataWorkFlowNodeInfoNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s GetWorkFlowResponseBodyDataWorkFlowNodeInfo) String() string {
	return tea.Prettify(s)
}

func (s GetWorkFlowResponseBodyDataWorkFlowNodeInfo) GoString() string {
	return s.String()
}

func (s *GetWorkFlowResponseBodyDataWorkFlowNodeInfo) SetEdges(v []*GetWorkFlowResponseBodyDataWorkFlowNodeInfoEdges) *GetWorkFlowResponseBodyDataWorkFlowNodeInfo {
	s.Edges = v
	return s
}

func (s *GetWorkFlowResponseBodyDataWorkFlowNodeInfo) SetNodes(v []*GetWorkFlowResponseBodyDataWorkFlowNodeInfoNodes) *GetWorkFlowResponseBodyDataWorkFlowNodeInfo {
	s.Nodes = v
	return s
}

type GetWorkFlowResponseBodyDataWorkFlowNodeInfoEdges struct {
	// 起始任务ID
	Source *int64 `json:"Source,omitempty" xml:"Source,omitempty"`
	// 目的任务ID
	Target *int64 `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s GetWorkFlowResponseBodyDataWorkFlowNodeInfoEdges) String() string {
	return tea.Prettify(s)
}

func (s GetWorkFlowResponseBodyDataWorkFlowNodeInfoEdges) GoString() string {
	return s.String()
}

func (s *GetWorkFlowResponseBodyDataWorkFlowNodeInfoEdges) SetSource(v int64) *GetWorkFlowResponseBodyDataWorkFlowNodeInfoEdges {
	s.Source = &v
	return s
}

func (s *GetWorkFlowResponseBodyDataWorkFlowNodeInfoEdges) SetTarget(v int64) *GetWorkFlowResponseBodyDataWorkFlowNodeInfoEdges {
	s.Target = &v
	return s
}

type GetWorkFlowResponseBodyDataWorkFlowNodeInfoNodes struct {
	// 任务ID
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 任务名称
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// 任务状态
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetWorkFlowResponseBodyDataWorkFlowNodeInfoNodes) String() string {
	return tea.Prettify(s)
}

func (s GetWorkFlowResponseBodyDataWorkFlowNodeInfoNodes) GoString() string {
	return s.String()
}

func (s *GetWorkFlowResponseBodyDataWorkFlowNodeInfoNodes) SetId(v int64) *GetWorkFlowResponseBodyDataWorkFlowNodeInfoNodes {
	s.Id = &v
	return s
}

func (s *GetWorkFlowResponseBodyDataWorkFlowNodeInfoNodes) SetLabel(v string) *GetWorkFlowResponseBodyDataWorkFlowNodeInfoNodes {
	s.Label = &v
	return s
}

func (s *GetWorkFlowResponseBodyDataWorkFlowNodeInfoNodes) SetStatus(v int32) *GetWorkFlowResponseBodyDataWorkFlowNodeInfoNodes {
	s.Status = &v
	return s
}

type GetWorkFlowResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWorkFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWorkFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkFlowResponse) GoString() string {
	return s.String()
}

func (s *GetWorkFlowResponse) SetHeaders(v map[string]*string) *GetWorkFlowResponse {
	s.Headers = v
	return s
}

func (s *GetWorkFlowResponse) SetBody(v *GetWorkFlowResponseBody) *GetWorkFlowResponse {
	s.Body = v
	return s
}

type GetWorkerListRequest struct {
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Namespace       *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetWorkerListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWorkerListRequest) GoString() string {
	return s.String()
}

func (s *GetWorkerListRequest) SetGroupId(v string) *GetWorkerListRequest {
	s.GroupId = &v
	return s
}

func (s *GetWorkerListRequest) SetNamespace(v string) *GetWorkerListRequest {
	s.Namespace = &v
	return s
}

func (s *GetWorkerListRequest) SetNamespaceSource(v string) *GetWorkerListRequest {
	s.NamespaceSource = &v
	return s
}

func (s *GetWorkerListRequest) SetRegionId(v string) *GetWorkerListRequest {
	s.RegionId = &v
	return s
}

type GetWorkerListResponseBody struct {
	Code    *int32                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *GetWorkerListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWorkerListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkerListResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkerListResponseBody) SetCode(v int32) *GetWorkerListResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkerListResponseBody) SetData(v *GetWorkerListResponseBodyData) *GetWorkerListResponseBody {
	s.Data = v
	return s
}

func (s *GetWorkerListResponseBody) SetMessage(v string) *GetWorkerListResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkerListResponseBody) SetRequestId(v string) *GetWorkerListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkerListResponseBody) SetSuccess(v bool) *GetWorkerListResponseBody {
	s.Success = &v
	return s
}

type GetWorkerListResponseBodyData struct {
	WorkerInfos []*GetWorkerListResponseBodyDataWorkerInfos `json:"WorkerInfos,omitempty" xml:"WorkerInfos,omitempty" type:"Repeated"`
}

func (s GetWorkerListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetWorkerListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWorkerListResponseBodyData) SetWorkerInfos(v []*GetWorkerListResponseBodyDataWorkerInfos) *GetWorkerListResponseBodyData {
	s.WorkerInfos = v
	return s
}

type GetWorkerListResponseBodyDataWorkerInfos struct {
	Ip            *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Label         *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Port          *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Starter       *string `json:"Starter,omitempty" xml:"Starter,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	WorkerAddress *string `json:"WorkerAddress,omitempty" xml:"WorkerAddress,omitempty"`
}

func (s GetWorkerListResponseBodyDataWorkerInfos) String() string {
	return tea.Prettify(s)
}

func (s GetWorkerListResponseBodyDataWorkerInfos) GoString() string {
	return s.String()
}

func (s *GetWorkerListResponseBodyDataWorkerInfos) SetIp(v string) *GetWorkerListResponseBodyDataWorkerInfos {
	s.Ip = &v
	return s
}

func (s *GetWorkerListResponseBodyDataWorkerInfos) SetLabel(v string) *GetWorkerListResponseBodyDataWorkerInfos {
	s.Label = &v
	return s
}

func (s *GetWorkerListResponseBodyDataWorkerInfos) SetPort(v int32) *GetWorkerListResponseBodyDataWorkerInfos {
	s.Port = &v
	return s
}

func (s *GetWorkerListResponseBodyDataWorkerInfos) SetStarter(v string) *GetWorkerListResponseBodyDataWorkerInfos {
	s.Starter = &v
	return s
}

func (s *GetWorkerListResponseBodyDataWorkerInfos) SetVersion(v string) *GetWorkerListResponseBodyDataWorkerInfos {
	s.Version = &v
	return s
}

func (s *GetWorkerListResponseBodyDataWorkerInfos) SetWorkerAddress(v string) *GetWorkerListResponseBodyDataWorkerInfos {
	s.WorkerAddress = &v
	return s
}

type GetWorkerListResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWorkerListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWorkerListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkerListResponse) GoString() string {
	return s.String()
}

func (s *GetWorkerListResponse) SetHeaders(v map[string]*string) *GetWorkerListResponse {
	s.Headers = v
	return s
}

func (s *GetWorkerListResponse) SetBody(v *GetWorkerListResponseBody) *GetWorkerListResponse {
	s.Body = v
	return s
}

type GrantPermissionRequest struct {
	GrantOption     *bool   `json:"GrantOption,omitempty" xml:"GrantOption,omitempty"`
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Namespace       *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserId          *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName        *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GrantPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantPermissionRequest) GoString() string {
	return s.String()
}

func (s *GrantPermissionRequest) SetGrantOption(v bool) *GrantPermissionRequest {
	s.GrantOption = &v
	return s
}

func (s *GrantPermissionRequest) SetGroupId(v string) *GrantPermissionRequest {
	s.GroupId = &v
	return s
}

func (s *GrantPermissionRequest) SetNamespace(v string) *GrantPermissionRequest {
	s.Namespace = &v
	return s
}

func (s *GrantPermissionRequest) SetNamespaceSource(v string) *GrantPermissionRequest {
	s.NamespaceSource = &v
	return s
}

func (s *GrantPermissionRequest) SetRegionId(v string) *GrantPermissionRequest {
	s.RegionId = &v
	return s
}

func (s *GrantPermissionRequest) SetUserId(v string) *GrantPermissionRequest {
	s.UserId = &v
	return s
}

func (s *GrantPermissionRequest) SetUserName(v string) *GrantPermissionRequest {
	s.UserName = &v
	return s
}

type GrantPermissionResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GrantPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *GrantPermissionResponseBody) SetCode(v int32) *GrantPermissionResponseBody {
	s.Code = &v
	return s
}

func (s *GrantPermissionResponseBody) SetMessage(v string) *GrantPermissionResponseBody {
	s.Message = &v
	return s
}

func (s *GrantPermissionResponseBody) SetRequestId(v string) *GrantPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantPermissionResponseBody) SetSuccess(v bool) *GrantPermissionResponseBody {
	s.Success = &v
	return s
}

type GrantPermissionResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GrantPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GrantPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantPermissionResponse) GoString() string {
	return s.String()
}

func (s *GrantPermissionResponse) SetHeaders(v map[string]*string) *GrantPermissionResponse {
	s.Headers = v
	return s
}

func (s *GrantPermissionResponse) SetBody(v *GrantPermissionResponseBody) *GrantPermissionResponse {
	s.Body = v
	return s
}

type ListGroupsRequest struct {
	Namespace       *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsRequest) SetNamespace(v string) *ListGroupsRequest {
	s.Namespace = &v
	return s
}

func (s *ListGroupsRequest) SetNamespaceSource(v string) *ListGroupsRequest {
	s.NamespaceSource = &v
	return s
}

func (s *ListGroupsRequest) SetRegionId(v string) *ListGroupsRequest {
	s.RegionId = &v
	return s
}

type ListGroupsResponseBody struct {
	Code      *int32                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBody) SetCode(v int32) *ListGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListGroupsResponseBody) SetData(v *ListGroupsResponseBodyData) *ListGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListGroupsResponseBody) SetMessage(v string) *ListGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListGroupsResponseBody) SetRequestId(v string) *ListGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupsResponseBody) SetSuccess(v bool) *ListGroupsResponseBody {
	s.Success = &v
	return s
}

type ListGroupsResponseBodyData struct {
	AppGroups []*ListGroupsResponseBodyDataAppGroups `json:"AppGroups,omitempty" xml:"AppGroups,omitempty" type:"Repeated"`
}

func (s ListGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBodyData) SetAppGroups(v []*ListGroupsResponseBodyDataAppGroups) *ListGroupsResponseBodyData {
	s.AppGroups = v
	return s
}

type ListGroupsResponseBodyDataAppGroups struct {
	AppKey      *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s ListGroupsResponseBodyDataAppGroups) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponseBodyDataAppGroups) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBodyDataAppGroups) SetAppKey(v string) *ListGroupsResponseBodyDataAppGroups {
	s.AppKey = &v
	return s
}

func (s *ListGroupsResponseBodyDataAppGroups) SetAppName(v string) *ListGroupsResponseBodyDataAppGroups {
	s.AppName = &v
	return s
}

func (s *ListGroupsResponseBodyDataAppGroups) SetDescription(v string) *ListGroupsResponseBodyDataAppGroups {
	s.Description = &v
	return s
}

func (s *ListGroupsResponseBodyDataAppGroups) SetGroupId(v string) *ListGroupsResponseBodyDataAppGroups {
	s.GroupId = &v
	return s
}

type ListGroupsResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListGroupsResponse) SetHeaders(v map[string]*string) *ListGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListGroupsResponse) SetBody(v *ListGroupsResponseBody) *ListGroupsResponse {
	s.Body = v
	return s
}

type ListJobsRequest struct {
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	JobName         *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	Namespace       *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) SetGroupId(v string) *ListJobsRequest {
	s.GroupId = &v
	return s
}

func (s *ListJobsRequest) SetJobName(v string) *ListJobsRequest {
	s.JobName = &v
	return s
}

func (s *ListJobsRequest) SetNamespace(v string) *ListJobsRequest {
	s.Namespace = &v
	return s
}

func (s *ListJobsRequest) SetNamespaceSource(v string) *ListJobsRequest {
	s.NamespaceSource = &v
	return s
}

func (s *ListJobsRequest) SetRegionId(v string) *ListJobsRequest {
	s.RegionId = &v
	return s
}

func (s *ListJobsRequest) SetStatus(v string) *ListJobsRequest {
	s.Status = &v
	return s
}

type ListJobsResponseBody struct {
	Code      *int32                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListJobsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBody) SetCode(v int32) *ListJobsResponseBody {
	s.Code = &v
	return s
}

func (s *ListJobsResponseBody) SetData(v *ListJobsResponseBodyData) *ListJobsResponseBody {
	s.Data = v
	return s
}

func (s *ListJobsResponseBody) SetMessage(v string) *ListJobsResponseBody {
	s.Message = &v
	return s
}

func (s *ListJobsResponseBody) SetRequestId(v string) *ListJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobsResponseBody) SetSuccess(v bool) *ListJobsResponseBody {
	s.Success = &v
	return s
}

type ListJobsResponseBodyData struct {
	Jobs []*ListJobsResponseBodyDataJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
}

func (s ListJobsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyData) SetJobs(v []*ListJobsResponseBodyDataJobs) *ListJobsResponseBodyData {
	s.Jobs = v
	return s
}

type ListJobsResponseBodyDataJobs struct {
	AttemptInterval *int32                                      `json:"AttemptInterval,omitempty" xml:"AttemptInterval,omitempty"`
	ClassName       *string                                     `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	Content         *string                                     `json:"Content,omitempty" xml:"Content,omitempty"`
	Description     *string                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	ExecuteMode     *string                                     `json:"ExecuteMode,omitempty" xml:"ExecuteMode,omitempty"`
	JarUrl          *string                                     `json:"JarUrl,omitempty" xml:"JarUrl,omitempty"`
	JobId           *int64                                      `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobMonitorInfo  *ListJobsResponseBodyDataJobsJobMonitorInfo `json:"JobMonitorInfo,omitempty" xml:"JobMonitorInfo,omitempty" type:"Struct"`
	MapTaskXAttrs   *ListJobsResponseBodyDataJobsMapTaskXAttrs  `json:"MapTaskXAttrs,omitempty" xml:"MapTaskXAttrs,omitempty" type:"Struct"`
	MaxAttempt      *int32                                      `json:"MaxAttempt,omitempty" xml:"MaxAttempt,omitempty"`
	MaxConcurrency  *string                                     `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	Name            *string                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	Parameters      *string                                     `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	Status          *int32                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	TimeConfig      *ListJobsResponseBodyDataJobsTimeConfig     `json:"TimeConfig,omitempty" xml:"TimeConfig,omitempty" type:"Struct"`
}

func (s ListJobsResponseBodyDataJobs) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyDataJobs) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyDataJobs) SetAttemptInterval(v int32) *ListJobsResponseBodyDataJobs {
	s.AttemptInterval = &v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetClassName(v string) *ListJobsResponseBodyDataJobs {
	s.ClassName = &v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetContent(v string) *ListJobsResponseBodyDataJobs {
	s.Content = &v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetDescription(v string) *ListJobsResponseBodyDataJobs {
	s.Description = &v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetExecuteMode(v string) *ListJobsResponseBodyDataJobs {
	s.ExecuteMode = &v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetJarUrl(v string) *ListJobsResponseBodyDataJobs {
	s.JarUrl = &v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetJobId(v int64) *ListJobsResponseBodyDataJobs {
	s.JobId = &v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetJobMonitorInfo(v *ListJobsResponseBodyDataJobsJobMonitorInfo) *ListJobsResponseBodyDataJobs {
	s.JobMonitorInfo = v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetMapTaskXAttrs(v *ListJobsResponseBodyDataJobsMapTaskXAttrs) *ListJobsResponseBodyDataJobs {
	s.MapTaskXAttrs = v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetMaxAttempt(v int32) *ListJobsResponseBodyDataJobs {
	s.MaxAttempt = &v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetMaxConcurrency(v string) *ListJobsResponseBodyDataJobs {
	s.MaxConcurrency = &v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetName(v string) *ListJobsResponseBodyDataJobs {
	s.Name = &v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetParameters(v string) *ListJobsResponseBodyDataJobs {
	s.Parameters = &v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetStatus(v int32) *ListJobsResponseBodyDataJobs {
	s.Status = &v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetTimeConfig(v *ListJobsResponseBodyDataJobsTimeConfig) *ListJobsResponseBodyDataJobs {
	s.TimeConfig = v
	return s
}

type ListJobsResponseBodyDataJobsJobMonitorInfo struct {
	ContactInfo   []*ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo `json:"ContactInfo,omitempty" xml:"ContactInfo,omitempty" type:"Repeated"`
	MonitorConfig *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig `json:"MonitorConfig,omitempty" xml:"MonitorConfig,omitempty" type:"Struct"`
}

func (s ListJobsResponseBodyDataJobsJobMonitorInfo) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyDataJobsJobMonitorInfo) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfo) SetContactInfo(v []*ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo) *ListJobsResponseBodyDataJobsJobMonitorInfo {
	s.ContactInfo = v
	return s
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfo) SetMonitorConfig(v *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig) *ListJobsResponseBodyDataJobsJobMonitorInfo {
	s.MonitorConfig = v
	return s
}

type ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo struct {
	Ding      *string `json:"Ding,omitempty" xml:"Ding,omitempty"`
	UserMail  *string `json:"UserMail,omitempty" xml:"UserMail,omitempty"`
	UserName  *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserPhone *string `json:"UserPhone,omitempty" xml:"UserPhone,omitempty"`
}

func (s ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo) SetDing(v string) *ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo {
	s.Ding = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo) SetUserMail(v string) *ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo {
	s.UserMail = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo) SetUserName(v string) *ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo {
	s.UserName = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo) SetUserPhone(v string) *ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo {
	s.UserPhone = &v
	return s
}

type ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig struct {
	FailEnable        *bool   `json:"FailEnable,omitempty" xml:"FailEnable,omitempty"`
	MissWorkerEnable  *bool   `json:"MissWorkerEnable,omitempty" xml:"MissWorkerEnable,omitempty"`
	SendChannel       *string `json:"SendChannel,omitempty" xml:"SendChannel,omitempty"`
	Timeout           *int64  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	TimeoutEnable     *bool   `json:"TimeoutEnable,omitempty" xml:"TimeoutEnable,omitempty"`
	TimeoutKillEnable *bool   `json:"TimeoutKillEnable,omitempty" xml:"TimeoutKillEnable,omitempty"`
}

func (s ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig) SetFailEnable(v bool) *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig {
	s.FailEnable = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig) SetMissWorkerEnable(v bool) *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig {
	s.MissWorkerEnable = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig) SetSendChannel(v string) *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig {
	s.SendChannel = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig) SetTimeout(v int64) *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig {
	s.Timeout = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig) SetTimeoutEnable(v bool) *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig {
	s.TimeoutEnable = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig) SetTimeoutKillEnable(v bool) *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig {
	s.TimeoutKillEnable = &v
	return s
}

type ListJobsResponseBodyDataJobsMapTaskXAttrs struct {
	ConsumerSize        *int32 `json:"ConsumerSize,omitempty" xml:"ConsumerSize,omitempty"`
	DispatcherSize      *int32 `json:"DispatcherSize,omitempty" xml:"DispatcherSize,omitempty"`
	PageSize            *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueueSize           *int32 `json:"QueueSize,omitempty" xml:"QueueSize,omitempty"`
	TaskAttemptInterval *int32 `json:"TaskAttemptInterval,omitempty" xml:"TaskAttemptInterval,omitempty"`
	TaskMaxAttempt      *int32 `json:"TaskMaxAttempt,omitempty" xml:"TaskMaxAttempt,omitempty"`
}

func (s ListJobsResponseBodyDataJobsMapTaskXAttrs) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyDataJobsMapTaskXAttrs) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyDataJobsMapTaskXAttrs) SetConsumerSize(v int32) *ListJobsResponseBodyDataJobsMapTaskXAttrs {
	s.ConsumerSize = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsMapTaskXAttrs) SetDispatcherSize(v int32) *ListJobsResponseBodyDataJobsMapTaskXAttrs {
	s.DispatcherSize = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsMapTaskXAttrs) SetPageSize(v int32) *ListJobsResponseBodyDataJobsMapTaskXAttrs {
	s.PageSize = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsMapTaskXAttrs) SetQueueSize(v int32) *ListJobsResponseBodyDataJobsMapTaskXAttrs {
	s.QueueSize = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsMapTaskXAttrs) SetTaskAttemptInterval(v int32) *ListJobsResponseBodyDataJobsMapTaskXAttrs {
	s.TaskAttemptInterval = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsMapTaskXAttrs) SetTaskMaxAttempt(v int32) *ListJobsResponseBodyDataJobsMapTaskXAttrs {
	s.TaskMaxAttempt = &v
	return s
}

type ListJobsResponseBodyDataJobsTimeConfig struct {
	Calendar       *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	DataOffset     *int32  `json:"DataOffset,omitempty" xml:"DataOffset,omitempty"`
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	TimeType       *int32  `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
}

func (s ListJobsResponseBodyDataJobsTimeConfig) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyDataJobsTimeConfig) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyDataJobsTimeConfig) SetCalendar(v string) *ListJobsResponseBodyDataJobsTimeConfig {
	s.Calendar = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsTimeConfig) SetDataOffset(v int32) *ListJobsResponseBodyDataJobsTimeConfig {
	s.DataOffset = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsTimeConfig) SetTimeExpression(v string) *ListJobsResponseBodyDataJobsTimeConfig {
	s.TimeExpression = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsTimeConfig) SetTimeType(v int32) *ListJobsResponseBodyDataJobsTimeConfig {
	s.TimeType = &v
	return s
}

type ListJobsResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponse) GoString() string {
	return s.String()
}

func (s *ListJobsResponse) SetHeaders(v map[string]*string) *ListJobsResponse {
	s.Headers = v
	return s
}

func (s *ListJobsResponse) SetBody(v *ListJobsResponseBody) *ListJobsResponse {
	s.Body = v
	return s
}

type ListNamespacesRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListNamespacesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNamespacesRequest) GoString() string {
	return s.String()
}

func (s *ListNamespacesRequest) SetRegionId(v string) *ListNamespacesRequest {
	s.RegionId = &v
	return s
}

type ListNamespacesResponseBody struct {
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListNamespacesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListNamespacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponseBody) SetCode(v int32) *ListNamespacesResponseBody {
	s.Code = &v
	return s
}

func (s *ListNamespacesResponseBody) SetData(v *ListNamespacesResponseBodyData) *ListNamespacesResponseBody {
	s.Data = v
	return s
}

func (s *ListNamespacesResponseBody) SetMessage(v string) *ListNamespacesResponseBody {
	s.Message = &v
	return s
}

func (s *ListNamespacesResponseBody) SetRequestId(v string) *ListNamespacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNamespacesResponseBody) SetSuccess(v bool) *ListNamespacesResponseBody {
	s.Success = &v
	return s
}

type ListNamespacesResponseBodyData struct {
	Namespaces []*ListNamespacesResponseBodyDataNamespaces `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
}

func (s ListNamespacesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListNamespacesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponseBodyData) SetNamespaces(v []*ListNamespacesResponseBodyDataNamespaces) *ListNamespacesResponseBodyData {
	s.Namespaces = v
	return s
}

type ListNamespacesResponseBodyDataNamespaces struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	UId         *string `json:"UId,omitempty" xml:"UId,omitempty"`
}

func (s ListNamespacesResponseBodyDataNamespaces) String() string {
	return tea.Prettify(s)
}

func (s ListNamespacesResponseBodyDataNamespaces) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponseBodyDataNamespaces) SetDescription(v string) *ListNamespacesResponseBodyDataNamespaces {
	s.Description = &v
	return s
}

func (s *ListNamespacesResponseBodyDataNamespaces) SetName(v string) *ListNamespacesResponseBodyDataNamespaces {
	s.Name = &v
	return s
}

func (s *ListNamespacesResponseBodyDataNamespaces) SetUId(v string) *ListNamespacesResponseBodyDataNamespaces {
	s.UId = &v
	return s
}

type ListNamespacesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListNamespacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNamespacesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNamespacesResponse) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponse) SetHeaders(v map[string]*string) *ListNamespacesResponse {
	s.Headers = v
	return s
}

func (s *ListNamespacesResponse) SetBody(v *ListNamespacesResponseBody) *ListNamespacesResponse {
	s.Body = v
	return s
}

type RevokePermissionRequest struct {
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Namespace       *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserId          *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RevokePermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokePermissionRequest) GoString() string {
	return s.String()
}

func (s *RevokePermissionRequest) SetGroupId(v string) *RevokePermissionRequest {
	s.GroupId = &v
	return s
}

func (s *RevokePermissionRequest) SetNamespace(v string) *RevokePermissionRequest {
	s.Namespace = &v
	return s
}

func (s *RevokePermissionRequest) SetNamespaceSource(v string) *RevokePermissionRequest {
	s.NamespaceSource = &v
	return s
}

func (s *RevokePermissionRequest) SetRegionId(v string) *RevokePermissionRequest {
	s.RegionId = &v
	return s
}

func (s *RevokePermissionRequest) SetUserId(v string) *RevokePermissionRequest {
	s.UserId = &v
	return s
}

type RevokePermissionResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RevokePermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *RevokePermissionResponseBody) SetCode(v int32) *RevokePermissionResponseBody {
	s.Code = &v
	return s
}

func (s *RevokePermissionResponseBody) SetMessage(v string) *RevokePermissionResponseBody {
	s.Message = &v
	return s
}

func (s *RevokePermissionResponseBody) SetRequestId(v string) *RevokePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokePermissionResponseBody) SetSuccess(v bool) *RevokePermissionResponseBody {
	s.Success = &v
	return s
}

type RevokePermissionResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RevokePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RevokePermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokePermissionResponse) GoString() string {
	return s.String()
}

func (s *RevokePermissionResponse) SetHeaders(v map[string]*string) *RevokePermissionResponse {
	s.Headers = v
	return s
}

func (s *RevokePermissionResponse) SetBody(v *RevokePermissionResponseBody) *RevokePermissionResponse {
	s.Body = v
	return s
}

type StopInstanceRequest struct {
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId      *int64  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId           *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Namespace       *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopInstanceRequest) SetGroupId(v string) *StopInstanceRequest {
	s.GroupId = &v
	return s
}

func (s *StopInstanceRequest) SetInstanceId(v int64) *StopInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StopInstanceRequest) SetJobId(v int64) *StopInstanceRequest {
	s.JobId = &v
	return s
}

func (s *StopInstanceRequest) SetNamespace(v string) *StopInstanceRequest {
	s.Namespace = &v
	return s
}

func (s *StopInstanceRequest) SetNamespaceSource(v string) *StopInstanceRequest {
	s.NamespaceSource = &v
	return s
}

func (s *StopInstanceRequest) SetRegionId(v string) *StopInstanceRequest {
	s.RegionId = &v
	return s
}

type StopInstanceResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopInstanceResponseBody) SetCode(v int32) *StopInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *StopInstanceResponseBody) SetMessage(v string) *StopInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *StopInstanceResponseBody) SetRequestId(v string) *StopInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopInstanceResponseBody) SetSuccess(v bool) *StopInstanceResponseBody {
	s.Success = &v
	return s
}

type StopInstanceResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopInstanceResponse) SetHeaders(v map[string]*string) *StopInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopInstanceResponse) SetBody(v *StopInstanceResponseBody) *StopInstanceResponse {
	s.Body = v
	return s
}

type UpdateJobRequest struct {
	AttemptInterval     *int32                         `json:"AttemptInterval,omitempty" xml:"AttemptInterval,omitempty"`
	Calendar            *string                        `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	ClassName           *string                        `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	ConsumerSize        *int32                         `json:"ConsumerSize,omitempty" xml:"ConsumerSize,omitempty"`
	ContactInfo         []*UpdateJobRequestContactInfo `json:"ContactInfo,omitempty" xml:"ContactInfo,omitempty" type:"Repeated"`
	Content             *string                        `json:"Content,omitempty" xml:"Content,omitempty"`
	DataOffset          *int32                         `json:"DataOffset,omitempty" xml:"DataOffset,omitempty"`
	Description         *string                        `json:"Description,omitempty" xml:"Description,omitempty"`
	DispatcherSize      *int32                         `json:"DispatcherSize,omitempty" xml:"DispatcherSize,omitempty"`
	ExecuteMode         *string                        `json:"ExecuteMode,omitempty" xml:"ExecuteMode,omitempty"`
	FailEnable          *bool                          `json:"FailEnable,omitempty" xml:"FailEnable,omitempty"`
	GroupId             *string                        `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	JarUrl              *string                        `json:"JarUrl,omitempty" xml:"JarUrl,omitempty"`
	JobId               *int64                         `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MaxAttempt          *int32                         `json:"MaxAttempt,omitempty" xml:"MaxAttempt,omitempty"`
	MaxConcurrency      *int32                         `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	MissWorkerEnable    *bool                          `json:"MissWorkerEnable,omitempty" xml:"MissWorkerEnable,omitempty"`
	Name                *string                        `json:"Name,omitempty" xml:"Name,omitempty"`
	Namespace           *string                        `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceSource     *string                        `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	PageSize            *int32                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Parameters          *string                        `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	QueueSize           *int32                         `json:"QueueSize,omitempty" xml:"QueueSize,omitempty"`
	RegionId            *string                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SendChannel         *string                        `json:"SendChannel,omitempty" xml:"SendChannel,omitempty"`
	TaskAttemptInterval *int32                         `json:"TaskAttemptInterval,omitempty" xml:"TaskAttemptInterval,omitempty"`
	TaskMaxAttempt      *int32                         `json:"TaskMaxAttempt,omitempty" xml:"TaskMaxAttempt,omitempty"`
	TimeExpression      *string                        `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	TimeType            *int32                         `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	Timeout             *int64                         `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	TimeoutEnable       *bool                          `json:"TimeoutEnable,omitempty" xml:"TimeoutEnable,omitempty"`
	TimeoutKillEnable   *bool                          `json:"TimeoutKillEnable,omitempty" xml:"TimeoutKillEnable,omitempty"`
}

func (s UpdateJobRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobRequest) SetAttemptInterval(v int32) *UpdateJobRequest {
	s.AttemptInterval = &v
	return s
}

func (s *UpdateJobRequest) SetCalendar(v string) *UpdateJobRequest {
	s.Calendar = &v
	return s
}

func (s *UpdateJobRequest) SetClassName(v string) *UpdateJobRequest {
	s.ClassName = &v
	return s
}

func (s *UpdateJobRequest) SetConsumerSize(v int32) *UpdateJobRequest {
	s.ConsumerSize = &v
	return s
}

func (s *UpdateJobRequest) SetContactInfo(v []*UpdateJobRequestContactInfo) *UpdateJobRequest {
	s.ContactInfo = v
	return s
}

func (s *UpdateJobRequest) SetContent(v string) *UpdateJobRequest {
	s.Content = &v
	return s
}

func (s *UpdateJobRequest) SetDataOffset(v int32) *UpdateJobRequest {
	s.DataOffset = &v
	return s
}

func (s *UpdateJobRequest) SetDescription(v string) *UpdateJobRequest {
	s.Description = &v
	return s
}

func (s *UpdateJobRequest) SetDispatcherSize(v int32) *UpdateJobRequest {
	s.DispatcherSize = &v
	return s
}

func (s *UpdateJobRequest) SetExecuteMode(v string) *UpdateJobRequest {
	s.ExecuteMode = &v
	return s
}

func (s *UpdateJobRequest) SetFailEnable(v bool) *UpdateJobRequest {
	s.FailEnable = &v
	return s
}

func (s *UpdateJobRequest) SetGroupId(v string) *UpdateJobRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateJobRequest) SetJarUrl(v string) *UpdateJobRequest {
	s.JarUrl = &v
	return s
}

func (s *UpdateJobRequest) SetJobId(v int64) *UpdateJobRequest {
	s.JobId = &v
	return s
}

func (s *UpdateJobRequest) SetMaxAttempt(v int32) *UpdateJobRequest {
	s.MaxAttempt = &v
	return s
}

func (s *UpdateJobRequest) SetMaxConcurrency(v int32) *UpdateJobRequest {
	s.MaxConcurrency = &v
	return s
}

func (s *UpdateJobRequest) SetMissWorkerEnable(v bool) *UpdateJobRequest {
	s.MissWorkerEnable = &v
	return s
}

func (s *UpdateJobRequest) SetName(v string) *UpdateJobRequest {
	s.Name = &v
	return s
}

func (s *UpdateJobRequest) SetNamespace(v string) *UpdateJobRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateJobRequest) SetNamespaceSource(v string) *UpdateJobRequest {
	s.NamespaceSource = &v
	return s
}

func (s *UpdateJobRequest) SetPageSize(v int32) *UpdateJobRequest {
	s.PageSize = &v
	return s
}

func (s *UpdateJobRequest) SetParameters(v string) *UpdateJobRequest {
	s.Parameters = &v
	return s
}

func (s *UpdateJobRequest) SetQueueSize(v int32) *UpdateJobRequest {
	s.QueueSize = &v
	return s
}

func (s *UpdateJobRequest) SetRegionId(v string) *UpdateJobRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateJobRequest) SetSendChannel(v string) *UpdateJobRequest {
	s.SendChannel = &v
	return s
}

func (s *UpdateJobRequest) SetTaskAttemptInterval(v int32) *UpdateJobRequest {
	s.TaskAttemptInterval = &v
	return s
}

func (s *UpdateJobRequest) SetTaskMaxAttempt(v int32) *UpdateJobRequest {
	s.TaskMaxAttempt = &v
	return s
}

func (s *UpdateJobRequest) SetTimeExpression(v string) *UpdateJobRequest {
	s.TimeExpression = &v
	return s
}

func (s *UpdateJobRequest) SetTimeType(v int32) *UpdateJobRequest {
	s.TimeType = &v
	return s
}

func (s *UpdateJobRequest) SetTimeout(v int64) *UpdateJobRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateJobRequest) SetTimeoutEnable(v bool) *UpdateJobRequest {
	s.TimeoutEnable = &v
	return s
}

func (s *UpdateJobRequest) SetTimeoutKillEnable(v bool) *UpdateJobRequest {
	s.TimeoutKillEnable = &v
	return s
}

type UpdateJobRequestContactInfo struct {
	Ding      *string `json:"Ding,omitempty" xml:"Ding,omitempty"`
	UserMail  *string `json:"UserMail,omitempty" xml:"UserMail,omitempty"`
	UserName  *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserPhone *string `json:"UserPhone,omitempty" xml:"UserPhone,omitempty"`
}

func (s UpdateJobRequestContactInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobRequestContactInfo) GoString() string {
	return s.String()
}

func (s *UpdateJobRequestContactInfo) SetDing(v string) *UpdateJobRequestContactInfo {
	s.Ding = &v
	return s
}

func (s *UpdateJobRequestContactInfo) SetUserMail(v string) *UpdateJobRequestContactInfo {
	s.UserMail = &v
	return s
}

func (s *UpdateJobRequestContactInfo) SetUserName(v string) *UpdateJobRequestContactInfo {
	s.UserName = &v
	return s
}

func (s *UpdateJobRequestContactInfo) SetUserPhone(v string) *UpdateJobRequestContactInfo {
	s.UserPhone = &v
	return s
}

type UpdateJobResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateJobResponseBody) SetCode(v int32) *UpdateJobResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateJobResponseBody) SetMessage(v string) *UpdateJobResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateJobResponseBody) SetRequestId(v string) *UpdateJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateJobResponseBody) SetSuccess(v bool) *UpdateJobResponseBody {
	s.Success = &v
	return s
}

type UpdateJobResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateJobResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobResponse) GoString() string {
	return s.String()
}

func (s *UpdateJobResponse) SetHeaders(v map[string]*string) *UpdateJobResponse {
	s.Headers = v
	return s
}

func (s *UpdateJobResponse) SetBody(v *UpdateJobResponseBody) *UpdateJobResponse {
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
		"cn-beijing":  tea.String("schedulerx.cn-beijing.aliyuncs.com"),
		"cn-hangzhou": tea.String("schedulerx.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai": tea.String("schedulerx.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen": tea.String("schedulerx.cn-shenzhen.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("schedulerx2"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) BatchDeleteJobsWithOptions(request *BatchDeleteJobsRequest, runtime *util.RuntimeOptions) (_result *BatchDeleteJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceSource)) {
		query["NamespaceSource"] = request.NamespaceSource
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobIdList)) {
		body["JobIdList"] = request.JobIdList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchDeleteJobs"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchDeleteJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchDeleteJobs(request *BatchDeleteJobsRequest) (_result *BatchDeleteJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchDeleteJobsResponse{}
	_body, _err := client.BatchDeleteJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchDisableJobsWithOptions(request *BatchDisableJobsRequest, runtime *util.RuntimeOptions) (_result *BatchDisableJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceSource)) {
		query["NamespaceSource"] = request.NamespaceSource
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobIdList)) {
		body["JobIdList"] = request.JobIdList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchDisableJobs"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchDisableJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchDisableJobs(request *BatchDisableJobsRequest) (_result *BatchDisableJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchDisableJobsResponse{}
	_body, _err := client.BatchDisableJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchEnableJobsWithOptions(request *BatchEnableJobsRequest, runtime *util.RuntimeOptions) (_result *BatchEnableJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceSource)) {
		query["NamespaceSource"] = request.NamespaceSource
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobIdList)) {
		body["JobIdList"] = request.JobIdList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchEnableJobs"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchEnableJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchEnableJobs(request *BatchEnableJobsRequest) (_result *BatchEnableJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchEnableJobsResponse{}
	_body, _err := client.BatchEnableJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppGroupWithOptions(request *CreateAppGroupRequest, runtime *util.RuntimeOptions) (_result *CreateAppGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppGroup"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAppGroup(request *CreateAppGroupRequest) (_result *CreateAppGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppGroupResponse{}
	_body, _err := client.CreateAppGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateJobWithOptions(request *CreateJobRequest, runtime *util.RuntimeOptions) (_result *CreateJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttemptInterval)) {
		body["AttemptInterval"] = request.AttemptInterval
	}

	if !tea.BoolValue(util.IsUnset(request.Calendar)) {
		body["Calendar"] = request.Calendar
	}

	if !tea.BoolValue(util.IsUnset(request.ClassName)) {
		body["ClassName"] = request.ClassName
	}

	if !tea.BoolValue(util.IsUnset(request.ConsumerSize)) {
		body["ConsumerSize"] = request.ConsumerSize
	}

	if !tea.BoolValue(util.IsUnset(request.ContactInfo)) {
		body["ContactInfo"] = request.ContactInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.DataOffset)) {
		body["DataOffset"] = request.DataOffset
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DispatcherSize)) {
		body["DispatcherSize"] = request.DispatcherSize
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteMode)) {
		body["ExecuteMode"] = request.ExecuteMode
	}

	if !tea.BoolValue(util.IsUnset(request.FailEnable)) {
		body["FailEnable"] = request.FailEnable
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.JarUrl)) {
		body["JarUrl"] = request.JarUrl
	}

	if !tea.BoolValue(util.IsUnset(request.JobType)) {
		body["JobType"] = request.JobType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxAttempt)) {
		body["MaxAttempt"] = request.MaxAttempt
	}

	if !tea.BoolValue(util.IsUnset(request.MaxConcurrency)) {
		body["MaxConcurrency"] = request.MaxConcurrency
	}

	if !tea.BoolValue(util.IsUnset(request.MissWorkerEnable)) {
		body["MissWorkerEnable"] = request.MissWorkerEnable
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		body["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceSource)) {
		body["NamespaceSource"] = request.NamespaceSource
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		body["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.QueueSize)) {
		body["QueueSize"] = request.QueueSize
	}

	if !tea.BoolValue(util.IsUnset(request.SendChannel)) {
		body["SendChannel"] = request.SendChannel
	}

	if !tea.BoolValue(util.IsUnset(request.TaskAttemptInterval)) {
		body["TaskAttemptInterval"] = request.TaskAttemptInterval
	}

	if !tea.BoolValue(util.IsUnset(request.TaskMaxAttempt)) {
		body["TaskMaxAttempt"] = request.TaskMaxAttempt
	}

	if !tea.BoolValue(util.IsUnset(request.TimeExpression)) {
		body["TimeExpression"] = request.TimeExpression
	}

	if !tea.BoolValue(util.IsUnset(request.TimeType)) {
		body["TimeType"] = request.TimeType
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		body["Timeout"] = request.Timeout
	}

	if !tea.BoolValue(util.IsUnset(request.TimeoutEnable)) {
		body["TimeoutEnable"] = request.TimeoutEnable
	}

	if !tea.BoolValue(util.IsUnset(request.TimeoutKillEnable)) {
		body["TimeoutKillEnable"] = request.TimeoutKillEnable
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateJob"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateJob(request *CreateJobRequest) (_result *CreateJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateJobResponse{}
	_body, _err := client.CreateJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteJobWithOptions(request *DeleteJobRequest, runtime *util.RuntimeOptions) (_result *DeleteJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteJob"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteJob(request *DeleteJobRequest) (_result *DeleteJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteJobResponse{}
	_body, _err := client.DeleteJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWorkflowWithOptions(request *DeleteWorkflowRequest, runtime *util.RuntimeOptions) (_result *DeleteWorkflowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWorkflow"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWorkflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWorkflow(request *DeleteWorkflowRequest) (_result *DeleteWorkflowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWorkflowResponse{}
	_body, _err := client.DeleteWorkflowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions() (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DesignateWorkersWithOptions(request *DesignateWorkersRequest, runtime *util.RuntimeOptions) (_result *DesignateWorkersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DesignateWorkers"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DesignateWorkersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DesignateWorkers(request *DesignateWorkersRequest) (_result *DesignateWorkersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DesignateWorkersResponse{}
	_body, _err := client.DesignateWorkersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableJobWithOptions(request *DisableJobRequest, runtime *util.RuntimeOptions) (_result *DisableJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableJob"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableJob(request *DisableJobRequest) (_result *DisableJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableJobResponse{}
	_body, _err := client.DisableJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableWorkflowWithOptions(request *DisableWorkflowRequest, runtime *util.RuntimeOptions) (_result *DisableWorkflowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableWorkflow"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableWorkflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableWorkflow(request *DisableWorkflowRequest) (_result *DisableWorkflowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableWorkflowResponse{}
	_body, _err := client.DisableWorkflowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableJobWithOptions(request *EnableJobRequest, runtime *util.RuntimeOptions) (_result *EnableJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableJob"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableJob(request *EnableJobRequest) (_result *EnableJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableJobResponse{}
	_body, _err := client.EnableJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableWorkflowWithOptions(request *EnableWorkflowRequest, runtime *util.RuntimeOptions) (_result *EnableWorkflowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableWorkflow"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableWorkflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableWorkflow(request *EnableWorkflowRequest) (_result *EnableWorkflowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableWorkflowResponse{}
	_body, _err := client.EnableWorkflowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteJobWithOptions(request *ExecuteJobRequest, runtime *util.RuntimeOptions) (_result *ExecuteJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteJob"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteJob(request *ExecuteJobRequest) (_result *ExecuteJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteJobResponse{}
	_body, _err := client.ExecuteJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteWorkflowWithOptions(request *ExecuteWorkflowRequest, runtime *util.RuntimeOptions) (_result *ExecuteWorkflowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteWorkflow"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteWorkflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteWorkflow(request *ExecuteWorkflowRequest) (_result *ExecuteWorkflowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteWorkflowResponse{}
	_body, _err := client.ExecuteWorkflowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobInfoWithOptions(request *GetJobInfoRequest, runtime *util.RuntimeOptions) (_result *GetJobInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJobInfo"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobInfo(request *GetJobInfoRequest) (_result *GetJobInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobInfoResponse{}
	_body, _err := client.GetJobInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobInstanceWithOptions(request *GetJobInstanceRequest, runtime *util.RuntimeOptions) (_result *GetJobInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJobInstance"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobInstance(request *GetJobInstanceRequest) (_result *GetJobInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobInstanceResponse{}
	_body, _err := client.GetJobInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobInstanceListWithOptions(request *GetJobInstanceListRequest, runtime *util.RuntimeOptions) (_result *GetJobInstanceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJobInstanceList"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobInstanceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobInstanceList(request *GetJobInstanceListRequest) (_result *GetJobInstanceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobInstanceListResponse{}
	_body, _err := client.GetJobInstanceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWorkFlowWithOptions(request *GetWorkFlowRequest, runtime *util.RuntimeOptions) (_result *GetWorkFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWorkFlow"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWorkFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWorkFlow(request *GetWorkFlowRequest) (_result *GetWorkFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWorkFlowResponse{}
	_body, _err := client.GetWorkFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWorkerListWithOptions(request *GetWorkerListRequest, runtime *util.RuntimeOptions) (_result *GetWorkerListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWorkerList"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWorkerListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWorkerList(request *GetWorkerListRequest) (_result *GetWorkerListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWorkerListResponse{}
	_body, _err := client.GetWorkerListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GrantPermissionWithOptions(request *GrantPermissionRequest, runtime *util.RuntimeOptions) (_result *GrantPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GrantOption)) {
		query["GrantOption"] = request.GrantOption
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceSource)) {
		query["NamespaceSource"] = request.NamespaceSource
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GrantPermission"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GrantPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GrantPermission(request *GrantPermissionRequest) (_result *GrantPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GrantPermissionResponse{}
	_body, _err := client.GrantPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGroupsWithOptions(request *ListGroupsRequest, runtime *util.RuntimeOptions) (_result *ListGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGroups"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGroups(request *ListGroupsRequest) (_result *ListGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGroupsResponse{}
	_body, _err := client.ListGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListJobsWithOptions(request *ListJobsRequest, runtime *util.RuntimeOptions) (_result *ListJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListJobs"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListJobs(request *ListJobsRequest) (_result *ListJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListJobsResponse{}
	_body, _err := client.ListJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNamespacesWithOptions(request *ListNamespacesRequest, runtime *util.RuntimeOptions) (_result *ListNamespacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNamespaces"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNamespacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNamespaces(request *ListNamespacesRequest) (_result *ListNamespacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNamespacesResponse{}
	_body, _err := client.ListNamespacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RevokePermissionWithOptions(request *RevokePermissionRequest, runtime *util.RuntimeOptions) (_result *RevokePermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceSource)) {
		query["NamespaceSource"] = request.NamespaceSource
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RevokePermission"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RevokePermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RevokePermission(request *RevokePermissionRequest) (_result *RevokePermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevokePermissionResponse{}
	_body, _err := client.RevokePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopInstanceWithOptions(request *StopInstanceRequest, runtime *util.RuntimeOptions) (_result *StopInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopInstance"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopInstance(request *StopInstanceRequest) (_result *StopInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopInstanceResponse{}
	_body, _err := client.StopInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateJobWithOptions(request *UpdateJobRequest, runtime *util.RuntimeOptions) (_result *UpdateJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttemptInterval)) {
		body["AttemptInterval"] = request.AttemptInterval
	}

	if !tea.BoolValue(util.IsUnset(request.Calendar)) {
		body["Calendar"] = request.Calendar
	}

	if !tea.BoolValue(util.IsUnset(request.ClassName)) {
		body["ClassName"] = request.ClassName
	}

	if !tea.BoolValue(util.IsUnset(request.ConsumerSize)) {
		body["ConsumerSize"] = request.ConsumerSize
	}

	if !tea.BoolValue(util.IsUnset(request.ContactInfo)) {
		body["ContactInfo"] = request.ContactInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.DataOffset)) {
		body["DataOffset"] = request.DataOffset
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DispatcherSize)) {
		body["DispatcherSize"] = request.DispatcherSize
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteMode)) {
		body["ExecuteMode"] = request.ExecuteMode
	}

	if !tea.BoolValue(util.IsUnset(request.FailEnable)) {
		body["FailEnable"] = request.FailEnable
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.JarUrl)) {
		body["JarUrl"] = request.JarUrl
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxAttempt)) {
		body["MaxAttempt"] = request.MaxAttempt
	}

	if !tea.BoolValue(util.IsUnset(request.MaxConcurrency)) {
		body["MaxConcurrency"] = request.MaxConcurrency
	}

	if !tea.BoolValue(util.IsUnset(request.MissWorkerEnable)) {
		body["MissWorkerEnable"] = request.MissWorkerEnable
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		body["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceSource)) {
		body["NamespaceSource"] = request.NamespaceSource
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		body["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.QueueSize)) {
		body["QueueSize"] = request.QueueSize
	}

	if !tea.BoolValue(util.IsUnset(request.SendChannel)) {
		body["SendChannel"] = request.SendChannel
	}

	if !tea.BoolValue(util.IsUnset(request.TaskAttemptInterval)) {
		body["TaskAttemptInterval"] = request.TaskAttemptInterval
	}

	if !tea.BoolValue(util.IsUnset(request.TaskMaxAttempt)) {
		body["TaskMaxAttempt"] = request.TaskMaxAttempt
	}

	if !tea.BoolValue(util.IsUnset(request.TimeExpression)) {
		body["TimeExpression"] = request.TimeExpression
	}

	if !tea.BoolValue(util.IsUnset(request.TimeType)) {
		body["TimeType"] = request.TimeType
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		body["Timeout"] = request.Timeout
	}

	if !tea.BoolValue(util.IsUnset(request.TimeoutEnable)) {
		body["TimeoutEnable"] = request.TimeoutEnable
	}

	if !tea.BoolValue(util.IsUnset(request.TimeoutKillEnable)) {
		body["TimeoutKillEnable"] = request.TimeoutKillEnable
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateJob"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateJob(request *UpdateJobRequest) (_result *UpdateJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateJobResponse{}
	_body, _err := client.UpdateJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
