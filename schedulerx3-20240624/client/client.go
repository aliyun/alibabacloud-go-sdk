// This file is auto-generated, don't edit it. Thanks.
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-pop/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateAppRequest struct {
	// example:
	//
	// ltk1ZXHv6LvibZypFkPHzRA
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// true
	EnableLog *bool `json:"EnableLog,omitempty" xml:"EnableLog,omitempty"`
	// example:
	//
	// 10
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// This parameter is required.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRequest) SetAccessToken(v string) *CreateAppRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateAppRequest) SetAppName(v string) *CreateAppRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppRequest) SetClusterId(v string) *CreateAppRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateAppRequest) SetEnableLog(v bool) *CreateAppRequest {
	s.EnableLog = &v
	return s
}

func (s *CreateAppRequest) SetMaxConcurrency(v int32) *CreateAppRequest {
	s.MaxConcurrency = &v
	return s
}

func (s *CreateAppRequest) SetTitle(v string) *CreateAppRequest {
	s.Title = &v
	return s
}

type CreateAppResponseBody struct {
	// example:
	//
	// 200
	Code *int32                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3835AA29-2298-5434-BC53-9CC377CDFD2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBody) SetCode(v int32) *CreateAppResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAppResponseBody) SetData(v *CreateAppResponseBodyData) *CreateAppResponseBody {
	s.Data = v
	return s
}

func (s *CreateAppResponseBody) SetMessage(v string) *CreateAppResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAppResponseBody) SetRequestId(v string) *CreateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppResponseBody) SetSuccess(v bool) *CreateAppResponseBody {
	s.Success = &v
	return s
}

type CreateAppResponseBodyData struct {
	// example:
	//
	// 4a0fae835cd741f3b12376d8a5a8e549v3
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// example:
	//
	// 10
	AppGroupId *int64 `json:"AppGroupId,omitempty" xml:"AppGroupId,omitempty"`
}

func (s CreateAppResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyData) SetAccessToken(v string) *CreateAppResponseBodyData {
	s.AccessToken = &v
	return s
}

func (s *CreateAppResponseBodyData) SetAppGroupId(v int64) *CreateAppResponseBodyData {
	s.AppGroupId = &v
	return s
}

type CreateAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponse) GoString() string {
	return s.String()
}

func (s *CreateAppResponse) SetHeaders(v map[string]*string) *CreateAppResponse {
	s.Headers = v
	return s
}

func (s *CreateAppResponse) SetStatusCode(v int32) *CreateAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppResponse) SetBody(v *CreateAppResponseBody) *CreateAppResponse {
	s.Body = v
	return s
}

type CreateClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// qianxi-test-0812
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// scx.dev.x1
	ClusterSpec *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// This parameter is required.
	VSwitches []*CreateClusterRequestVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// VPC id
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-aa1a18236n90rqhuhhnhh
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) SetClusterName(v string) *CreateClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateClusterRequest) SetClusterSpec(v string) *CreateClusterRequest {
	s.ClusterSpec = &v
	return s
}

func (s *CreateClusterRequest) SetEngineType(v string) *CreateClusterRequest {
	s.EngineType = &v
	return s
}

func (s *CreateClusterRequest) SetVSwitches(v []*CreateClusterRequestVSwitches) *CreateClusterRequest {
	s.VSwitches = v
	return s
}

func (s *CreateClusterRequest) SetVpcId(v string) *CreateClusterRequest {
	s.VpcId = &v
	return s
}

type CreateClusterRequestVSwitches struct {
	// This parameter is required.
	//
	// example:
	//
	// vsw-2ze745n3r2sfqtahhubpl
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateClusterRequestVSwitches) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestVSwitches) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestVSwitches) SetVSwitchId(v string) *CreateClusterRequestVSwitches {
	s.VSwitchId = &v
	return s
}

func (s *CreateClusterRequestVSwitches) SetZoneId(v string) *CreateClusterRequestVSwitches {
	s.ZoneId = &v
	return s
}

type CreateClusterShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// qianxi-test-0812
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// scx.dev.x1
	ClusterSpec *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// This parameter is required.
	VSwitchesShrink *string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty"`
	// VPC id
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-aa1a18236n90rqhuhhnhh
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateClusterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterShrinkRequest) SetClusterName(v string) *CreateClusterShrinkRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClusterSpec(v string) *CreateClusterShrinkRequest {
	s.ClusterSpec = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetEngineType(v string) *CreateClusterShrinkRequest {
	s.EngineType = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetVSwitchesShrink(v string) *CreateClusterShrinkRequest {
	s.VSwitchesShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetVpcId(v string) *CreateClusterShrinkRequest {
	s.VpcId = &v
	return s
}

type CreateClusterResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *CreateClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// INVALID_PARAMETER
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// B8733786-C045-59F1-8D79-99A52863F62D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) SetCode(v int32) *CreateClusterResponseBody {
	s.Code = &v
	return s
}

func (s *CreateClusterResponseBody) SetData(v *CreateClusterResponseBodyData) *CreateClusterResponseBody {
	s.Data = v
	return s
}

func (s *CreateClusterResponseBody) SetErrorCode(v string) *CreateClusterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateClusterResponseBody) SetMessage(v string) *CreateClusterResponseBody {
	s.Message = &v
	return s
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClusterResponseBody) SetSuccess(v bool) *CreateClusterResponseBody {
	s.Success = &v
	return s
}

type CreateClusterResponseBodyData struct {
	// example:
	//
	// xxljob-b21969c2309
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 229317760970086
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateClusterResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBodyData) SetClusterId(v string) *CreateClusterResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterResponseBodyData) SetOrderId(v int64) *CreateClusterResponseBodyData {
	s.OrderId = &v
	return s
}

type CreateClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterResponse) SetHeaders(v map[string]*string) *CreateClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateClusterResponse) SetStatusCode(v int32) *CreateClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClusterResponse) SetBody(v *CreateClusterResponseBody) *CreateClusterResponse {
	s.Body = v
	return s
}

type CreateJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 3
	AttemptInterval *int32 `json:"AttemptInterval,omitempty" xml:"AttemptInterval,omitempty"`
	// example:
	//
	// workday
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1
	ExecutorBlockStrategy *int32 `json:"ExecutorBlockStrategy,omitempty" xml:"ExecutorBlockStrategy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testJobVoidHandler
	JobHandler *string `json:"JobHandler,omitempty" xml:"JobHandler,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// example:
	//
	// 3
	MaxAttempt *int32 `json:"MaxAttempt,omitempty" xml:"MaxAttempt,omitempty"`
	// example:
	//
	// 1
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-job
	Name           *string                           `json:"Name,omitempty" xml:"Name,omitempty"`
	NoticeConfig   *CreateJobRequestNoticeConfig     `json:"NoticeConfig,omitempty" xml:"NoticeConfig,omitempty" type:"Struct"`
	NoticeContacts []*CreateJobRequestNoticeContacts `json:"NoticeContacts,omitempty" xml:"NoticeContacts,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 1
	RouteStrategy *int32 `json:"RouteStrategy,omitempty" xml:"RouteStrategy,omitempty"`
	// example:
	//
	// 1701310327000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 0 0 4 ? 	- Mon/1
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// -1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// example:
	//
	// GMT+8
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
}

func (s CreateJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequest) GoString() string {
	return s.String()
}

func (s *CreateJobRequest) SetAppName(v string) *CreateJobRequest {
	s.AppName = &v
	return s
}

func (s *CreateJobRequest) SetAttemptInterval(v int32) *CreateJobRequest {
	s.AttemptInterval = &v
	return s
}

func (s *CreateJobRequest) SetCalendar(v string) *CreateJobRequest {
	s.Calendar = &v
	return s
}

func (s *CreateJobRequest) SetClusterId(v string) *CreateJobRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateJobRequest) SetDescription(v string) *CreateJobRequest {
	s.Description = &v
	return s
}

func (s *CreateJobRequest) SetExecutorBlockStrategy(v int32) *CreateJobRequest {
	s.ExecutorBlockStrategy = &v
	return s
}

func (s *CreateJobRequest) SetJobHandler(v string) *CreateJobRequest {
	s.JobHandler = &v
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

func (s *CreateJobRequest) SetName(v string) *CreateJobRequest {
	s.Name = &v
	return s
}

func (s *CreateJobRequest) SetNoticeConfig(v *CreateJobRequestNoticeConfig) *CreateJobRequest {
	s.NoticeConfig = v
	return s
}

func (s *CreateJobRequest) SetNoticeContacts(v []*CreateJobRequestNoticeContacts) *CreateJobRequest {
	s.NoticeContacts = v
	return s
}

func (s *CreateJobRequest) SetParameters(v string) *CreateJobRequest {
	s.Parameters = &v
	return s
}

func (s *CreateJobRequest) SetPriority(v int32) *CreateJobRequest {
	s.Priority = &v
	return s
}

func (s *CreateJobRequest) SetRouteStrategy(v int32) *CreateJobRequest {
	s.RouteStrategy = &v
	return s
}

func (s *CreateJobRequest) SetStartTime(v int64) *CreateJobRequest {
	s.StartTime = &v
	return s
}

func (s *CreateJobRequest) SetStatus(v int32) *CreateJobRequest {
	s.Status = &v
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

func (s *CreateJobRequest) SetTimezone(v string) *CreateJobRequest {
	s.Timezone = &v
	return s
}

type CreateJobRequestNoticeConfig struct {
	// example:
	//
	// true
	FailEnable *bool `json:"FailEnable,omitempty" xml:"FailEnable,omitempty"`
	// example:
	//
	// 1
	FailLimitTimes *int32 `json:"FailLimitTimes,omitempty" xml:"FailLimitTimes,omitempty"`
	// example:
	//
	// true
	MissWorkerEnable *bool `json:"MissWorkerEnable,omitempty" xml:"MissWorkerEnable,omitempty"`
	// example:
	//
	// mail
	SendChannel *string `json:"SendChannel,omitempty" xml:"SendChannel,omitempty"`
	// example:
	//
	// true
	SuccessNotice *bool `json:"SuccessNotice,omitempty" xml:"SuccessNotice,omitempty"`
	// example:
	//
	// 30
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// example:
	//
	// true
	TimeoutEnable *bool `json:"TimeoutEnable,omitempty" xml:"TimeoutEnable,omitempty"`
	// example:
	//
	// true
	TimeoutKillEnable *bool `json:"TimeoutKillEnable,omitempty" xml:"TimeoutKillEnable,omitempty"`
}

func (s CreateJobRequestNoticeConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestNoticeConfig) GoString() string {
	return s.String()
}

func (s *CreateJobRequestNoticeConfig) SetFailEnable(v bool) *CreateJobRequestNoticeConfig {
	s.FailEnable = &v
	return s
}

func (s *CreateJobRequestNoticeConfig) SetFailLimitTimes(v int32) *CreateJobRequestNoticeConfig {
	s.FailLimitTimes = &v
	return s
}

func (s *CreateJobRequestNoticeConfig) SetMissWorkerEnable(v bool) *CreateJobRequestNoticeConfig {
	s.MissWorkerEnable = &v
	return s
}

func (s *CreateJobRequestNoticeConfig) SetSendChannel(v string) *CreateJobRequestNoticeConfig {
	s.SendChannel = &v
	return s
}

func (s *CreateJobRequestNoticeConfig) SetSuccessNotice(v bool) *CreateJobRequestNoticeConfig {
	s.SuccessNotice = &v
	return s
}

func (s *CreateJobRequestNoticeConfig) SetTimeout(v int64) *CreateJobRequestNoticeConfig {
	s.Timeout = &v
	return s
}

func (s *CreateJobRequestNoticeConfig) SetTimeoutEnable(v bool) *CreateJobRequestNoticeConfig {
	s.TimeoutEnable = &v
	return s
}

func (s *CreateJobRequestNoticeConfig) SetTimeoutKillEnable(v bool) *CreateJobRequestNoticeConfig {
	s.TimeoutKillEnable = &v
	return s
}

type CreateJobRequestNoticeContacts struct {
	// example:
	//
	// 1
	ContactType *int32 `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
	// example:
	//
	// xiaoming
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateJobRequestNoticeContacts) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestNoticeContacts) GoString() string {
	return s.String()
}

func (s *CreateJobRequestNoticeContacts) SetContactType(v int32) *CreateJobRequestNoticeContacts {
	s.ContactType = &v
	return s
}

func (s *CreateJobRequestNoticeContacts) SetName(v string) *CreateJobRequestNoticeContacts {
	s.Name = &v
	return s
}

type CreateJobShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 3
	AttemptInterval *int32 `json:"AttemptInterval,omitempty" xml:"AttemptInterval,omitempty"`
	// example:
	//
	// workday
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1
	ExecutorBlockStrategy *int32 `json:"ExecutorBlockStrategy,omitempty" xml:"ExecutorBlockStrategy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testJobVoidHandler
	JobHandler *string `json:"JobHandler,omitempty" xml:"JobHandler,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// example:
	//
	// 3
	MaxAttempt *int32 `json:"MaxAttempt,omitempty" xml:"MaxAttempt,omitempty"`
	// example:
	//
	// 1
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-job
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NoticeConfigShrink   *string `json:"NoticeConfig,omitempty" xml:"NoticeConfig,omitempty"`
	NoticeContactsShrink *string `json:"NoticeContacts,omitempty" xml:"NoticeContacts,omitempty"`
	// example:
	//
	// test
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 1
	RouteStrategy *int32 `json:"RouteStrategy,omitempty" xml:"RouteStrategy,omitempty"`
	// example:
	//
	// 1701310327000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 0 0 4 ? 	- Mon/1
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// -1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// example:
	//
	// GMT+8
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
}

func (s CreateJobShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateJobShrinkRequest) SetAppName(v string) *CreateJobShrinkRequest {
	s.AppName = &v
	return s
}

func (s *CreateJobShrinkRequest) SetAttemptInterval(v int32) *CreateJobShrinkRequest {
	s.AttemptInterval = &v
	return s
}

func (s *CreateJobShrinkRequest) SetCalendar(v string) *CreateJobShrinkRequest {
	s.Calendar = &v
	return s
}

func (s *CreateJobShrinkRequest) SetClusterId(v string) *CreateJobShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateJobShrinkRequest) SetDescription(v string) *CreateJobShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateJobShrinkRequest) SetExecutorBlockStrategy(v int32) *CreateJobShrinkRequest {
	s.ExecutorBlockStrategy = &v
	return s
}

func (s *CreateJobShrinkRequest) SetJobHandler(v string) *CreateJobShrinkRequest {
	s.JobHandler = &v
	return s
}

func (s *CreateJobShrinkRequest) SetJobType(v string) *CreateJobShrinkRequest {
	s.JobType = &v
	return s
}

func (s *CreateJobShrinkRequest) SetMaxAttempt(v int32) *CreateJobShrinkRequest {
	s.MaxAttempt = &v
	return s
}

func (s *CreateJobShrinkRequest) SetMaxConcurrency(v int32) *CreateJobShrinkRequest {
	s.MaxConcurrency = &v
	return s
}

func (s *CreateJobShrinkRequest) SetName(v string) *CreateJobShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateJobShrinkRequest) SetNoticeConfigShrink(v string) *CreateJobShrinkRequest {
	s.NoticeConfigShrink = &v
	return s
}

func (s *CreateJobShrinkRequest) SetNoticeContactsShrink(v string) *CreateJobShrinkRequest {
	s.NoticeContactsShrink = &v
	return s
}

func (s *CreateJobShrinkRequest) SetParameters(v string) *CreateJobShrinkRequest {
	s.Parameters = &v
	return s
}

func (s *CreateJobShrinkRequest) SetPriority(v int32) *CreateJobShrinkRequest {
	s.Priority = &v
	return s
}

func (s *CreateJobShrinkRequest) SetRouteStrategy(v int32) *CreateJobShrinkRequest {
	s.RouteStrategy = &v
	return s
}

func (s *CreateJobShrinkRequest) SetStartTime(v int64) *CreateJobShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *CreateJobShrinkRequest) SetStatus(v int32) *CreateJobShrinkRequest {
	s.Status = &v
	return s
}

func (s *CreateJobShrinkRequest) SetTimeExpression(v string) *CreateJobShrinkRequest {
	s.TimeExpression = &v
	return s
}

func (s *CreateJobShrinkRequest) SetTimeType(v int32) *CreateJobShrinkRequest {
	s.TimeType = &v
	return s
}

func (s *CreateJobShrinkRequest) SetTimezone(v string) *CreateJobShrinkRequest {
	s.Timezone = &v
	return s
}

type CreateJobResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *CreateJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3835AA29-2298-5434-BC53-9CC377CDFD2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// 74
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateJobResponse) SetStatusCode(v int32) *CreateJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateJobResponse) SetBody(v *CreateJobResponseBody) *CreateJobResponse {
	s.Body = v
	return s
}

type DeleteAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DeleteAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppRequest) SetAppName(v string) *DeleteAppRequest {
	s.AppName = &v
	return s
}

func (s *DeleteAppRequest) SetClusterId(v string) *DeleteAppRequest {
	s.ClusterId = &v
	return s
}

type DeleteAppResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CF99C381-C8F6-5A8D-8C24-57F46B706D2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppResponseBody) SetCode(v int32) *DeleteAppResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAppResponseBody) SetMessage(v string) *DeleteAppResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAppResponseBody) SetRequestId(v string) *DeleteAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppResponseBody) SetSuccess(v bool) *DeleteAppResponseBody {
	s.Success = &v
	return s
}

type DeleteAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppResponse) SetHeaders(v map[string]*string) *DeleteAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppResponse) SetStatusCode(v int32) *DeleteAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppResponse) SetBody(v *DeleteAppResponseBody) *DeleteAppResponse {
	s.Body = v
	return s
}

type DeleteClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxljob-a1804a3226d
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DeleteClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequest) SetClusterId(v string) *DeleteClusterRequest {
	s.ClusterId = &v
	return s
}

type DeleteClusterResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F131C3E0-3FAA-5FA4-A6F3-E974D69EF3C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponseBody) SetCode(v int32) *DeleteClusterResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteClusterResponseBody) SetMessage(v string) *DeleteClusterResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteClusterResponseBody) SetRequestId(v string) *DeleteClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteClusterResponseBody) SetSuccess(v bool) *DeleteClusterResponseBody {
	s.Success = &v
	return s
}

type DeleteClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponse) SetHeaders(v map[string]*string) *DeleteClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteClusterResponse) SetStatusCode(v int32) *DeleteClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClusterResponse) SetBody(v *DeleteClusterResponseBody) *DeleteClusterResponse {
	s.Body = v
	return s
}

type DeleteJobsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// -
	JobIds []*int64 `json:"JobIds,omitempty" xml:"JobIds,omitempty" type:"Repeated"`
}

func (s DeleteJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobsRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobsRequest) SetAppName(v string) *DeleteJobsRequest {
	s.AppName = &v
	return s
}

func (s *DeleteJobsRequest) SetClusterId(v string) *DeleteJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteJobsRequest) SetJobIds(v []*int64) *DeleteJobsRequest {
	s.JobIds = v
	return s
}

type DeleteJobsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// -
	JobIdsShrink *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
}

func (s DeleteJobsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobsShrinkRequest) SetAppName(v string) *DeleteJobsShrinkRequest {
	s.AppName = &v
	return s
}

func (s *DeleteJobsShrinkRequest) SetClusterId(v string) *DeleteJobsShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteJobsShrinkRequest) SetJobIdsShrink(v string) *DeleteJobsShrinkRequest {
	s.JobIdsShrink = &v
	return s
}

type DeleteJobsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 91143E1D-E235-5BE0-9364-C2EE28FFB5A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobsResponseBody) SetCode(v int32) *DeleteJobsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteJobsResponseBody) SetMessage(v string) *DeleteJobsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteJobsResponseBody) SetRequestId(v string) *DeleteJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteJobsResponseBody) SetSuccess(v bool) *DeleteJobsResponseBody {
	s.Success = &v
	return s
}

type DeleteJobsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobsResponse) GoString() string {
	return s.String()
}

func (s *DeleteJobsResponse) SetHeaders(v map[string]*string) *DeleteJobsResponse {
	s.Headers = v
	return s
}

func (s *DeleteJobsResponse) SetStatusCode(v int32) *DeleteJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteJobsResponse) SetBody(v *DeleteJobsResponseBody) *DeleteJobsResponse {
	s.Body = v
	return s
}

type ExportJobsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 1
	ExportJobType *int32 `json:"ExportJobType,omitempty" xml:"ExportJobType,omitempty"`
	// -
	JobIds []*int64 `json:"JobIds,omitempty" xml:"JobIds,omitempty" type:"Repeated"`
}

func (s ExportJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportJobsRequest) GoString() string {
	return s.String()
}

func (s *ExportJobsRequest) SetAppName(v string) *ExportJobsRequest {
	s.AppName = &v
	return s
}

func (s *ExportJobsRequest) SetClusterId(v string) *ExportJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *ExportJobsRequest) SetExportJobType(v int32) *ExportJobsRequest {
	s.ExportJobType = &v
	return s
}

func (s *ExportJobsRequest) SetJobIds(v []*int64) *ExportJobsRequest {
	s.JobIds = v
	return s
}

type ExportJobsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 1
	ExportJobType *int32 `json:"ExportJobType,omitempty" xml:"ExportJobType,omitempty"`
	// -
	JobIdsShrink *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
}

func (s ExportJobsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ExportJobsShrinkRequest) SetAppName(v string) *ExportJobsShrinkRequest {
	s.AppName = &v
	return s
}

func (s *ExportJobsShrinkRequest) SetClusterId(v string) *ExportJobsShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ExportJobsShrinkRequest) SetExportJobType(v int32) *ExportJobsShrinkRequest {
	s.ExportJobType = &v
	return s
}

func (s *ExportJobsShrinkRequest) SetJobIdsShrink(v string) *ExportJobsShrinkRequest {
	s.JobIdsShrink = &v
	return s
}

type ExportJobsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []byte             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportJobsResponse) GoString() string {
	return s.String()
}

func (s *ExportJobsResponse) SetHeaders(v map[string]*string) *ExportJobsResponse {
	s.Headers = v
	return s
}

func (s *ExportJobsResponse) SetStatusCode(v int32) *ExportJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportJobsResponse) SetBody(v []byte) *ExportJobsResponse {
	s.Body = v
	return s
}

type GetClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxljob-d6a5243b6fa
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s GetClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRequest) GoString() string {
	return s.String()
}

func (s *GetClusterRequest) SetClusterId(v string) *GetClusterRequest {
	s.ClusterId = &v
	return s
}

type GetClusterResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *GetClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D0DE9C33-992A-580B-89C4-B609A292748D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBody) SetCode(v int32) *GetClusterResponseBody {
	s.Code = &v
	return s
}

func (s *GetClusterResponseBody) SetData(v *GetClusterResponseBodyData) *GetClusterResponseBody {
	s.Data = v
	return s
}

func (s *GetClusterResponseBody) SetMessage(v string) *GetClusterResponseBody {
	s.Message = &v
	return s
}

func (s *GetClusterResponseBody) SetRequestId(v string) *GetClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterResponseBody) SetSuccess(v bool) *GetClusterResponseBody {
	s.Success = &v
	return s
}

type GetClusterResponseBodyData struct {
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// xxljob-e0d018c6df8
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// xxl-job-test-1730427575152
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// scx.small.x2
	ClusterSpec *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	// example:
	//
	// 2024-10-29 15:56:36
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2024-10-29 15:56:36
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// xxljob
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// example:
	//
	// 2.0.0
	EngineVersion  *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	InternetDomain *string `json:"InternetDomain,omitempty" xml:"InternetDomain,omitempty"`
	// example:
	//
	// http://xxljob-xxxxxx.schedulerx.mse.aliyuncs.com
	IntranetDomain *string `json:"IntranetDomain,omitempty" xml:"IntranetDomain,omitempty"`
	// example:
	//
	// 100
	JobNum *int32 `json:"JobNum,omitempty" xml:"JobNum,omitempty"`
	// example:
	//
	// {
	//
	//   "vSwitchIdList": [
	//
	//     "xxx",
	//
	//     "xxx"
	//
	//   ],
	//
	//   "cpu": xxx,
	//
	//   "cpuUnit": "xxx",
	//
	//   "diskCapacity": xxx,
	//
	//   "memoryCapacity": xxx,
	//
	//   "zoneIds": [
	//
	//     "xxx",
	//
	//     "xxx"
	//
	//   ],
	//
	//   "securityGroupList": [
	//
	//     "xxx"
	//
	//   ],
	//
	//   "eniCrossZone": "xxx",
	//
	//   "regionId": "xxx",
	//
	//   "instanceCount": xxx,
	//
	//   "vpcId": "xxx",
	//
	//   "memoryUnit": "xxx",
	//
	//   "diskType": "xxx",
	//
	//   "appClusterId": "xxx"
	//
	// }
	KubeConfig *string `json:"KubeConfig,omitempty" xml:"KubeConfig,omitempty"`
	// example:
	//
	// 1000
	MaxJobNum *int32 `json:"MaxJobNum,omitempty" xml:"MaxJobNum,omitempty"`
	// example:
	//
	// 2
	ProductType *int32 `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// 10
	Spm *int32 `json:"Spm,omitempty" xml:"Spm,omitempty"`
	// example:
	//
	// 2
	Status    *int32                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	VSwitches []*GetClusterResponseBodyDataVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// VPC ID
	//
	// example:
	//
	// vpc-bp1fiz967u39lt8yuxcs0
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// 10
	WorkerNum *int32    `json:"WorkerNum,omitempty" xml:"WorkerNum,omitempty"`
	Zones     []*string `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s GetClusterResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyData) SetChargeType(v string) *GetClusterResponseBodyData {
	s.ChargeType = &v
	return s
}

func (s *GetClusterResponseBodyData) SetClusterId(v string) *GetClusterResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *GetClusterResponseBodyData) SetClusterName(v string) *GetClusterResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *GetClusterResponseBodyData) SetClusterSpec(v string) *GetClusterResponseBodyData {
	s.ClusterSpec = &v
	return s
}

func (s *GetClusterResponseBodyData) SetCreateTime(v string) *GetClusterResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetClusterResponseBodyData) SetEndTime(v string) *GetClusterResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetClusterResponseBodyData) SetEngineType(v string) *GetClusterResponseBodyData {
	s.EngineType = &v
	return s
}

func (s *GetClusterResponseBodyData) SetEngineVersion(v string) *GetClusterResponseBodyData {
	s.EngineVersion = &v
	return s
}

func (s *GetClusterResponseBodyData) SetInternetDomain(v string) *GetClusterResponseBodyData {
	s.InternetDomain = &v
	return s
}

func (s *GetClusterResponseBodyData) SetIntranetDomain(v string) *GetClusterResponseBodyData {
	s.IntranetDomain = &v
	return s
}

func (s *GetClusterResponseBodyData) SetJobNum(v int32) *GetClusterResponseBodyData {
	s.JobNum = &v
	return s
}

func (s *GetClusterResponseBodyData) SetKubeConfig(v string) *GetClusterResponseBodyData {
	s.KubeConfig = &v
	return s
}

func (s *GetClusterResponseBodyData) SetMaxJobNum(v int32) *GetClusterResponseBodyData {
	s.MaxJobNum = &v
	return s
}

func (s *GetClusterResponseBodyData) SetProductType(v int32) *GetClusterResponseBodyData {
	s.ProductType = &v
	return s
}

func (s *GetClusterResponseBodyData) SetSpm(v int32) *GetClusterResponseBodyData {
	s.Spm = &v
	return s
}

func (s *GetClusterResponseBodyData) SetStatus(v int32) *GetClusterResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetClusterResponseBodyData) SetVSwitches(v []*GetClusterResponseBodyDataVSwitches) *GetClusterResponseBodyData {
	s.VSwitches = v
	return s
}

func (s *GetClusterResponseBodyData) SetVpcId(v string) *GetClusterResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *GetClusterResponseBodyData) SetWorkerNum(v int32) *GetClusterResponseBodyData {
	s.WorkerNum = &v
	return s
}

func (s *GetClusterResponseBodyData) SetZones(v []*string) *GetClusterResponseBodyData {
	s.Zones = v
	return s
}

type GetClusterResponseBodyDataVSwitches struct {
	// example:
	//
	// vsw-8vbf1n216nshvfjdyff8a
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetClusterResponseBodyDataVSwitches) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBodyDataVSwitches) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyDataVSwitches) SetVSwitchId(v string) *GetClusterResponseBodyDataVSwitches {
	s.VSwitchId = &v
	return s
}

func (s *GetClusterResponseBodyDataVSwitches) SetZoneId(v string) *GetClusterResponseBodyDataVSwitches {
	s.ZoneId = &v
	return s
}

type GetClusterResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponse) GoString() string {
	return s.String()
}

func (s *GetClusterResponse) SetHeaders(v map[string]*string) *GetClusterResponse {
	s.Headers = v
	return s
}

func (s *GetClusterResponse) SetStatusCode(v int32) *GetClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterResponse) SetBody(v *GetClusterResponseBody) *GetClusterResponse {
	s.Body = v
	return s
}

type GetDesigateInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 74
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetDesigateInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDesigateInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDesigateInfoRequest) SetAppName(v string) *GetDesigateInfoRequest {
	s.AppName = &v
	return s
}

func (s *GetDesigateInfoRequest) SetClusterId(v string) *GetDesigateInfoRequest {
	s.ClusterId = &v
	return s
}

func (s *GetDesigateInfoRequest) SetJobId(v int64) *GetDesigateInfoRequest {
	s.JobId = &v
	return s
}

type GetDesigateInfoResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *GetDesigateInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1DF6732E-15D8-5E1F-95E3-C10077F556B5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDesigateInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDesigateInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDesigateInfoResponseBody) SetCode(v int32) *GetDesigateInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetDesigateInfoResponseBody) SetData(v *GetDesigateInfoResponseBodyData) *GetDesigateInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetDesigateInfoResponseBody) SetMessage(v string) *GetDesigateInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetDesigateInfoResponseBody) SetRequestId(v string) *GetDesigateInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDesigateInfoResponseBody) SetSuccess(v bool) *GetDesigateInfoResponseBody {
	s.Success = &v
	return s
}

type GetDesigateInfoResponseBodyData struct {
	// example:
	//
	// 2
	DesignateType *int32 `json:"DesignateType,omitempty" xml:"DesignateType,omitempty"`
	// example:
	//
	// true
	Transferable *bool `json:"Transferable,omitempty" xml:"Transferable,omitempty"`
}

func (s GetDesigateInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDesigateInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDesigateInfoResponseBodyData) SetDesignateType(v int32) *GetDesigateInfoResponseBodyData {
	s.DesignateType = &v
	return s
}

func (s *GetDesigateInfoResponseBodyData) SetTransferable(v bool) *GetDesigateInfoResponseBodyData {
	s.Transferable = &v
	return s
}

type GetDesigateInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDesigateInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDesigateInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDesigateInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDesigateInfoResponse) SetHeaders(v map[string]*string) *GetDesigateInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDesigateInfoResponse) SetStatusCode(v int32) *GetDesigateInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDesigateInfoResponse) SetBody(v *GetDesigateInfoResponseBody) *GetDesigateInfoResponse {
	s.Body = v
	return s
}

type GetJobExecutionProgressRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1310630367761285120
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
}

func (s GetJobExecutionProgressRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobExecutionProgressRequest) GoString() string {
	return s.String()
}

func (s *GetJobExecutionProgressRequest) SetAppName(v string) *GetJobExecutionProgressRequest {
	s.AppName = &v
	return s
}

func (s *GetJobExecutionProgressRequest) SetClusterId(v string) *GetJobExecutionProgressRequest {
	s.ClusterId = &v
	return s
}

func (s *GetJobExecutionProgressRequest) SetJobExecutionId(v string) *GetJobExecutionProgressRequest {
	s.JobExecutionId = &v
	return s
}

type GetJobExecutionProgressResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *GetJobExecutionProgressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter format error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9B57FDD7-ABBE-5030-B348-86EB9943DB59
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetJobExecutionProgressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobExecutionProgressResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobExecutionProgressResponseBody) SetCode(v int32) *GetJobExecutionProgressResponseBody {
	s.Code = &v
	return s
}

func (s *GetJobExecutionProgressResponseBody) SetData(v *GetJobExecutionProgressResponseBodyData) *GetJobExecutionProgressResponseBody {
	s.Data = v
	return s
}

func (s *GetJobExecutionProgressResponseBody) SetMessage(v string) *GetJobExecutionProgressResponseBody {
	s.Message = &v
	return s
}

func (s *GetJobExecutionProgressResponseBody) SetRequestId(v string) *GetJobExecutionProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobExecutionProgressResponseBody) SetSuccess(v bool) *GetJobExecutionProgressResponseBody {
	s.Success = &v
	return s
}

type GetJobExecutionProgressResponseBodyData struct {
	JobDescription   *string                                                    `json:"JobDescription,omitempty" xml:"JobDescription,omitempty"`
	RootProgress     *GetJobExecutionProgressResponseBodyDataRootProgress       `json:"RootProgress,omitempty" xml:"RootProgress,omitempty" type:"Struct"`
	ShardingProgress []*GetJobExecutionProgressResponseBodyDataShardingProgress `json:"ShardingProgress,omitempty" xml:"ShardingProgress,omitempty" type:"Repeated"`
	TaskProgress     []*GetJobExecutionProgressResponseBodyDataTaskProgress     `json:"TaskProgress,omitempty" xml:"TaskProgress,omitempty" type:"Repeated"`
	TotalProgress    *GetJobExecutionProgressResponseBodyDataTotalProgress      `json:"TotalProgress,omitempty" xml:"TotalProgress,omitempty" type:"Struct"`
	WorkerProgress   []*GetJobExecutionProgressResponseBodyDataWorkerProgress   `json:"WorkerProgress,omitempty" xml:"WorkerProgress,omitempty" type:"Repeated"`
}

func (s GetJobExecutionProgressResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetJobExecutionProgressResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetJobExecutionProgressResponseBodyData) SetJobDescription(v string) *GetJobExecutionProgressResponseBodyData {
	s.JobDescription = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyData) SetRootProgress(v *GetJobExecutionProgressResponseBodyDataRootProgress) *GetJobExecutionProgressResponseBodyData {
	s.RootProgress = v
	return s
}

func (s *GetJobExecutionProgressResponseBodyData) SetShardingProgress(v []*GetJobExecutionProgressResponseBodyDataShardingProgress) *GetJobExecutionProgressResponseBodyData {
	s.ShardingProgress = v
	return s
}

func (s *GetJobExecutionProgressResponseBodyData) SetTaskProgress(v []*GetJobExecutionProgressResponseBodyDataTaskProgress) *GetJobExecutionProgressResponseBodyData {
	s.TaskProgress = v
	return s
}

func (s *GetJobExecutionProgressResponseBodyData) SetTotalProgress(v *GetJobExecutionProgressResponseBodyDataTotalProgress) *GetJobExecutionProgressResponseBodyData {
	s.TotalProgress = v
	return s
}

func (s *GetJobExecutionProgressResponseBodyData) SetWorkerProgress(v []*GetJobExecutionProgressResponseBodyDataWorkerProgress) *GetJobExecutionProgressResponseBodyData {
	s.WorkerProgress = v
	return s
}

type GetJobExecutionProgressResponseBodyDataRootProgress struct {
	// example:
	//
	// 2
	Finished *int64 `json:"Finished,omitempty" xml:"Finished,omitempty"`
	// example:
	//
	// 2
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetJobExecutionProgressResponseBodyDataRootProgress) String() string {
	return tea.Prettify(s)
}

func (s GetJobExecutionProgressResponseBodyDataRootProgress) GoString() string {
	return s.String()
}

func (s *GetJobExecutionProgressResponseBodyDataRootProgress) SetFinished(v int64) *GetJobExecutionProgressResponseBodyDataRootProgress {
	s.Finished = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataRootProgress) SetTotal(v int64) *GetJobExecutionProgressResponseBodyDataRootProgress {
	s.Total = &v
	return s
}

type GetJobExecutionProgressResponseBodyDataShardingProgress struct {
	// id
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1306189481388277762
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
	// example:
	//
	// 2,4,6,8,10
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 5
	Status     *int32                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusType *GetJobExecutionProgressResponseBodyDataShardingProgressStatusType `json:"StatusType,omitempty" xml:"StatusType,omitempty" type:"Struct"`
	// example:
	//
	// http://192.168.1.9:9999/
	WorkerAddr *string `json:"WorkerAddr,omitempty" xml:"WorkerAddr,omitempty"`
}

func (s GetJobExecutionProgressResponseBodyDataShardingProgress) String() string {
	return tea.Prettify(s)
}

func (s GetJobExecutionProgressResponseBodyDataShardingProgress) GoString() string {
	return s.String()
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgress) SetId(v int64) *GetJobExecutionProgressResponseBodyDataShardingProgress {
	s.Id = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgress) SetJobExecutionId(v string) *GetJobExecutionProgressResponseBodyDataShardingProgress {
	s.JobExecutionId = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgress) SetResult(v string) *GetJobExecutionProgressResponseBodyDataShardingProgress {
	s.Result = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgress) SetStatus(v int32) *GetJobExecutionProgressResponseBodyDataShardingProgress {
	s.Status = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgress) SetStatusType(v *GetJobExecutionProgressResponseBodyDataShardingProgressStatusType) *GetJobExecutionProgressResponseBodyDataShardingProgress {
	s.StatusType = v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgress) SetWorkerAddr(v string) *GetJobExecutionProgressResponseBodyDataShardingProgress {
	s.WorkerAddr = &v
	return s
}

type GetJobExecutionProgressResponseBodyDataShardingProgressStatusType struct {
	// example:
	//
	// 5
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// TaskStatus.FAILED
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// -
	Tips map[string]*string `json:"Tips,omitempty" xml:"Tips,omitempty"`
}

func (s GetJobExecutionProgressResponseBodyDataShardingProgressStatusType) String() string {
	return tea.Prettify(s)
}

func (s GetJobExecutionProgressResponseBodyDataShardingProgressStatusType) GoString() string {
	return s.String()
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgressStatusType) SetCode(v string) *GetJobExecutionProgressResponseBodyDataShardingProgressStatusType {
	s.Code = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgressStatusType) SetName(v string) *GetJobExecutionProgressResponseBodyDataShardingProgressStatusType {
	s.Name = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgressStatusType) SetTips(v map[string]*string) *GetJobExecutionProgressResponseBodyDataShardingProgressStatusType {
	s.Tips = v
	return s
}

type GetJobExecutionProgressResponseBodyDataTaskProgress struct {
	// example:
	//
	// 100
	Failed *int32 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// example:
	//
	// calendar_test_2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 10
	Pulled *int32 `json:"Pulled,omitempty" xml:"Pulled,omitempty"`
	// example:
	//
	// 100
	Queue *int32 `json:"Queue,omitempty" xml:"Queue,omitempty"`
	// example:
	//
	// 1
	Running *int32 `json:"Running,omitempty" xml:"Running,omitempty"`
	// example:
	//
	// 100
	Success *int32 `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1000
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetJobExecutionProgressResponseBodyDataTaskProgress) String() string {
	return tea.Prettify(s)
}

func (s GetJobExecutionProgressResponseBodyDataTaskProgress) GoString() string {
	return s.String()
}

func (s *GetJobExecutionProgressResponseBodyDataTaskProgress) SetFailed(v int32) *GetJobExecutionProgressResponseBodyDataTaskProgress {
	s.Failed = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataTaskProgress) SetName(v string) *GetJobExecutionProgressResponseBodyDataTaskProgress {
	s.Name = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataTaskProgress) SetPulled(v int32) *GetJobExecutionProgressResponseBodyDataTaskProgress {
	s.Pulled = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataTaskProgress) SetQueue(v int32) *GetJobExecutionProgressResponseBodyDataTaskProgress {
	s.Queue = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataTaskProgress) SetRunning(v int32) *GetJobExecutionProgressResponseBodyDataTaskProgress {
	s.Running = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataTaskProgress) SetSuccess(v int32) *GetJobExecutionProgressResponseBodyDataTaskProgress {
	s.Success = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataTaskProgress) SetTotal(v int32) *GetJobExecutionProgressResponseBodyDataTaskProgress {
	s.Total = &v
	return s
}

type GetJobExecutionProgressResponseBodyDataTotalProgress struct {
	// example:
	//
	// 15
	Finished *int64 `json:"Finished,omitempty" xml:"Finished,omitempty"`
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetJobExecutionProgressResponseBodyDataTotalProgress) String() string {
	return tea.Prettify(s)
}

func (s GetJobExecutionProgressResponseBodyDataTotalProgress) GoString() string {
	return s.String()
}

func (s *GetJobExecutionProgressResponseBodyDataTotalProgress) SetFinished(v int64) *GetJobExecutionProgressResponseBodyDataTotalProgress {
	s.Finished = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataTotalProgress) SetTotal(v int64) *GetJobExecutionProgressResponseBodyDataTotalProgress {
	s.Total = &v
	return s
}

type GetJobExecutionProgressResponseBodyDataWorkerProgress struct {
	// example:
	//
	// 20
	Failed *int32 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// example:
	//
	// 20
	Pulled *int32 `json:"Pulled,omitempty" xml:"Pulled,omitempty"`
	// example:
	//
	// 20
	Queue *int32 `json:"Queue,omitempty" xml:"Queue,omitempty"`
	// example:
	//
	// 20
	Running *int32 `json:"Running,omitempty" xml:"Running,omitempty"`
	// example:
	//
	// 20
	Success *int32 `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// example:
	//
	// 1a0e97fb17244665327205402dbd6d
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	// example:
	//
	// 10.10.116.53:61941
	WorkerAddr *string `json:"WorkerAddr,omitempty" xml:"WorkerAddr,omitempty"`
}

func (s GetJobExecutionProgressResponseBodyDataWorkerProgress) String() string {
	return tea.Prettify(s)
}

func (s GetJobExecutionProgressResponseBodyDataWorkerProgress) GoString() string {
	return s.String()
}

func (s *GetJobExecutionProgressResponseBodyDataWorkerProgress) SetFailed(v int32) *GetJobExecutionProgressResponseBodyDataWorkerProgress {
	s.Failed = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataWorkerProgress) SetPulled(v int32) *GetJobExecutionProgressResponseBodyDataWorkerProgress {
	s.Pulled = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataWorkerProgress) SetQueue(v int32) *GetJobExecutionProgressResponseBodyDataWorkerProgress {
	s.Queue = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataWorkerProgress) SetRunning(v int32) *GetJobExecutionProgressResponseBodyDataWorkerProgress {
	s.Running = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataWorkerProgress) SetSuccess(v int32) *GetJobExecutionProgressResponseBodyDataWorkerProgress {
	s.Success = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataWorkerProgress) SetTotal(v int32) *GetJobExecutionProgressResponseBodyDataWorkerProgress {
	s.Total = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataWorkerProgress) SetTraceId(v string) *GetJobExecutionProgressResponseBodyDataWorkerProgress {
	s.TraceId = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataWorkerProgress) SetWorkerAddr(v string) *GetJobExecutionProgressResponseBodyDataWorkerProgress {
	s.WorkerAddr = &v
	return s
}

type GetJobExecutionProgressResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobExecutionProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobExecutionProgressResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobExecutionProgressResponse) GoString() string {
	return s.String()
}

func (s *GetJobExecutionProgressResponse) SetHeaders(v map[string]*string) *GetJobExecutionProgressResponse {
	s.Headers = v
	return s
}

func (s *GetJobExecutionProgressResponse) SetStatusCode(v int32) *GetJobExecutionProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobExecutionProgressResponse) SetBody(v *GetJobExecutionProgressResponseBody) *GetJobExecutionProgressResponse {
	s.Body = v
	return s
}

type GetLogRequest struct {
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 1721636220
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1310630367761285120
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
	// example:
	//
	// hello word
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// INFO
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// LineNum
	//
	// example:
	//
	// 2
	LineNum *int32 `json:"LineNum,omitempty" xml:"LineNum,omitempty"`
	// example:
	//
	// 344008
	LogId *int64 `json:"LogId,omitempty" xml:"LogId,omitempty"`
	// example:
	//
	// 0
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// example:
	//
	// true
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// example:
	//
	// 1721636220
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLogRequest) GoString() string {
	return s.String()
}

func (s *GetLogRequest) SetAppName(v string) *GetLogRequest {
	s.AppName = &v
	return s
}

func (s *GetLogRequest) SetClusterId(v string) *GetLogRequest {
	s.ClusterId = &v
	return s
}

func (s *GetLogRequest) SetEndTime(v int64) *GetLogRequest {
	s.EndTime = &v
	return s
}

func (s *GetLogRequest) SetJobExecutionId(v string) *GetLogRequest {
	s.JobExecutionId = &v
	return s
}

func (s *GetLogRequest) SetKeyword(v string) *GetLogRequest {
	s.Keyword = &v
	return s
}

func (s *GetLogRequest) SetLevel(v string) *GetLogRequest {
	s.Level = &v
	return s
}

func (s *GetLogRequest) SetLineNum(v int32) *GetLogRequest {
	s.LineNum = &v
	return s
}

func (s *GetLogRequest) SetLogId(v int64) *GetLogRequest {
	s.LogId = &v
	return s
}

func (s *GetLogRequest) SetOffset(v int32) *GetLogRequest {
	s.Offset = &v
	return s
}

func (s *GetLogRequest) SetReverse(v bool) *GetLogRequest {
	s.Reverse = &v
	return s
}

func (s *GetLogRequest) SetStartTime(v int64) *GetLogRequest {
	s.StartTime = &v
	return s
}

type GetLogResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// C78E2AD2-5985-515B-BAD2-31A248AFC263
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetLogResponseBody) SetCode(v int32) *GetLogResponseBody {
	s.Code = &v
	return s
}

func (s *GetLogResponseBody) SetData(v []*string) *GetLogResponseBody {
	s.Data = v
	return s
}

func (s *GetLogResponseBody) SetMessage(v string) *GetLogResponseBody {
	s.Message = &v
	return s
}

func (s *GetLogResponseBody) SetRequestId(v string) *GetLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLogResponseBody) SetSuccess(v bool) *GetLogResponseBody {
	s.Success = &v
	return s
}

type GetLogResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLogResponse) GoString() string {
	return s.String()
}

func (s *GetLogResponse) SetHeaders(v map[string]*string) *GetLogResponse {
	s.Headers = v
	return s
}

func (s *GetLogResponse) SetStatusCode(v int32) *GetLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLogResponse) SetBody(v *GetLogResponseBody) *GetLogResponse {
	s.Body = v
	return s
}

type ImportCalendarRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"month":1,"days":[3,4,5,6,9,10,11,12,13,16,17,18,19,20,28,29,30,31]},{"month":2,"days":[1,2,3,6,7,8,9,10,13,14,15,16,17,20,21,22,23,24,27,28]},{"month":3,"days":[1,2,3,6,7,8,9,10,13,14,15,16,17,20,21,22,23,24,27,28,29,30,31]},{"month":4,"days":[3,4,6,7,10,11,12,13,14,17,18,19,20,21,23,24,25,26,27,28]},{"month":5,"days":[4,5,6,8,9,10,11,12,15,16,17,18,19,22,23,24,25,26,29,30,31]},{"month":6,"days":[1,2,5,6,7,8,9,12,13,14,15,16,19,20,21,25,26,27,28,29,30]},{"month":7,"days":[3,4,5,6,7,10,11,12,13,14,17,18,19,20,21,24,25,26,27,28,31]},{"month":8,"days":[1,2,3,4,7,8,9,10,11,14,15,16,17,18,21,22,23,24,25,28,29,30,31]},{"month":9,"days":[1,4,5,6,7,8,11,12,13,14,15,18,19,20,21,22,25,26,27,28]},{"month":10,"days":[7,8,9,10,11,12,13,16,17,18,19,20,23,24,25,26,27,30,31]},{"month":11,"days":[1,2,3,6,7,8,9,10,13,14,15,16,17,20,21,22,23,24,27,28,29,30]},{"month":12,"days":[1,4,5,6,7,8,11,12,13,14,15,18,19,20,21,22,25,26,27,28,29]}]
	Months *string `json:"Months,omitempty" xml:"Months,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// workday
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024
	Year *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s ImportCalendarRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportCalendarRequest) GoString() string {
	return s.String()
}

func (s *ImportCalendarRequest) SetClusterId(v string) *ImportCalendarRequest {
	s.ClusterId = &v
	return s
}

func (s *ImportCalendarRequest) SetMonths(v string) *ImportCalendarRequest {
	s.Months = &v
	return s
}

func (s *ImportCalendarRequest) SetName(v string) *ImportCalendarRequest {
	s.Name = &v
	return s
}

func (s *ImportCalendarRequest) SetYear(v int32) *ImportCalendarRequest {
	s.Year = &v
	return s
}

type ImportCalendarResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2ECA6FC9-7557-5576-AF5F-FC3E7BCC9C21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportCalendarResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportCalendarResponseBody) GoString() string {
	return s.String()
}

func (s *ImportCalendarResponseBody) SetCode(v int32) *ImportCalendarResponseBody {
	s.Code = &v
	return s
}

func (s *ImportCalendarResponseBody) SetData(v []*string) *ImportCalendarResponseBody {
	s.Data = v
	return s
}

func (s *ImportCalendarResponseBody) SetMessage(v string) *ImportCalendarResponseBody {
	s.Message = &v
	return s
}

func (s *ImportCalendarResponseBody) SetRequestId(v string) *ImportCalendarResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportCalendarResponseBody) SetSuccess(v bool) *ImportCalendarResponseBody {
	s.Success = &v
	return s
}

type ImportCalendarResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportCalendarResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportCalendarResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportCalendarResponse) GoString() string {
	return s.String()
}

func (s *ImportCalendarResponse) SetHeaders(v map[string]*string) *ImportCalendarResponse {
	s.Headers = v
	return s
}

func (s *ImportCalendarResponse) SetStatusCode(v int32) *ImportCalendarResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportCalendarResponse) SetBody(v *ImportCalendarResponseBody) *ImportCalendarResponse {
	s.Body = v
	return s
}

type ImportJobsRequest struct {
	// example:
	//
	// true
	AutoCreateApp *bool `json:"AutoCreateApp,omitempty" xml:"AutoCreateApp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// {
	//
	//   "kind": "SchedulerXJobs",
	//
	//   "type": "JSON",
	//
	//   "version": "2.0",
	//
	//   "content": [
	//
	//     {
	//
	//       "appName": "xxl-job-executor-perf-test-xx",
	//
	//       "groupId": "xxl-job-executor-perf-test-xx",
	//
	//       "description": "xxl-job-executor-xx",
	//
	//       "jobConfigInfo": [
	//
	//         {
	//
	//           "jobHandler": "testJobVoidHandler",
	//
	//           "dataOffset": 0,
	//
	//           "executeMode": "standalone",
	//
	//           "monitorConfigInfo": {
	//
	//             "alarmType": "CustomContacts",
	//
	//             "failLimitTimes": 1,
	//
	//             "failEnable": true,
	//
	//             "failRate": 100,
	//
	//             "timeoutKillEnable": false,
	//
	//             "missWorkerEnable": false,
	//
	//             "sendChannel": "webhook",
	//
	//             "timeoutEnable": true,
	//
	//             "timeout": 7200,
	//
	//             "daysOfDeadline": 0,
	//
	//             "successNotice": false
	//
	//           },
	//
	//           "attemptInterval": 30,
	//
	//           "cleanMode": "{\\"cleanMode\\":\\"NUM_ONLY\\",\\"totalRemain\\":300}",
	//
	//           "description": "",
	//
	//           "routeStrategy": 1,
	//
	//           "userName": "xx",
	//
	//           "userId": "xx",
	//
	//           "content": "{\\"jobHandler\\":\\"testJobVoidHandler\\"}",
	//
	//           "maxConcurrency": 1,
	//
	//           "maxAttempt": 0,
	//
	//           "name": "perf_auto_test_0",
	//
	//           "xattrs": "",
	//
	//           "jobType": "xxljob",
	//
	//           "contentType": 1,
	//
	//           "parameters": "success-withMsg",
	//
	//           "timeConfig": {
	//
	//             "calendar": "",
	//
	//             "dataOffset": 0,
	//
	//             "timeType": 1,
	//
	//             "paramMap": {},
	//
	//             "timeExpression": "	- 	- 	- 	- 	- ?"
	//
	//           },
	//
	//           "contactInfoList": [],
	//
	//           "status": 0
	//
	//         }
	//
	//       ]
	//
	//     }
	//
	//   ]
	//
	// }
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// true
	Overwrite *bool `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
}

func (s ImportJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportJobsRequest) GoString() string {
	return s.String()
}

func (s *ImportJobsRequest) SetAutoCreateApp(v bool) *ImportJobsRequest {
	s.AutoCreateApp = &v
	return s
}

func (s *ImportJobsRequest) SetClusterId(v string) *ImportJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *ImportJobsRequest) SetContent(v string) *ImportJobsRequest {
	s.Content = &v
	return s
}

func (s *ImportJobsRequest) SetOverwrite(v bool) *ImportJobsRequest {
	s.Overwrite = &v
	return s
}

type ImportJobsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9A48E22F-F30A-5CE5-AC7A-E0FED1B6942E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ImportJobsResponseBody) SetCode(v int32) *ImportJobsResponseBody {
	s.Code = &v
	return s
}

func (s *ImportJobsResponseBody) SetMessage(v string) *ImportJobsResponseBody {
	s.Message = &v
	return s
}

func (s *ImportJobsResponseBody) SetRequestId(v string) *ImportJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportJobsResponseBody) SetSuccess(v bool) *ImportJobsResponseBody {
	s.Success = &v
	return s
}

type ImportJobsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportJobsResponse) GoString() string {
	return s.String()
}

func (s *ImportJobsResponse) SetHeaders(v map[string]*string) *ImportJobsResponse {
	s.Headers = v
	return s
}

func (s *ImportJobsResponse) SetStatusCode(v int32) *ImportJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportJobsResponse) SetBody(v *ImportJobsResponseBody) *ImportJobsResponse {
	s.Body = v
	return s
}

type ListAlarmEventRequest struct {
	// example:
	//
	// webhook
	AlarmChannel *string `json:"AlarmChannel,omitempty" xml:"AlarmChannel,omitempty"`
	// example:
	//
	// true
	AlarmStatus *string `json:"AlarmStatus,omitempty" xml:"AlarmStatus,omitempty"`
	// example:
	//
	// schedulerx3_fail_alarm
	AlarmType *string `json:"AlarmType,omitempty" xml:"AlarmType,omitempty"`
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 1731636011558
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// test-job
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// 1
	PageNum *string `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// true
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// example:
	//
	// 1690419316000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListAlarmEventRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmEventRequest) GoString() string {
	return s.String()
}

func (s *ListAlarmEventRequest) SetAlarmChannel(v string) *ListAlarmEventRequest {
	s.AlarmChannel = &v
	return s
}

func (s *ListAlarmEventRequest) SetAlarmStatus(v string) *ListAlarmEventRequest {
	s.AlarmStatus = &v
	return s
}

func (s *ListAlarmEventRequest) SetAlarmType(v string) *ListAlarmEventRequest {
	s.AlarmType = &v
	return s
}

func (s *ListAlarmEventRequest) SetAppName(v string) *ListAlarmEventRequest {
	s.AppName = &v
	return s
}

func (s *ListAlarmEventRequest) SetClusterId(v string) *ListAlarmEventRequest {
	s.ClusterId = &v
	return s
}

func (s *ListAlarmEventRequest) SetEndTime(v int64) *ListAlarmEventRequest {
	s.EndTime = &v
	return s
}

func (s *ListAlarmEventRequest) SetJobName(v string) *ListAlarmEventRequest {
	s.JobName = &v
	return s
}

func (s *ListAlarmEventRequest) SetPageNum(v string) *ListAlarmEventRequest {
	s.PageNum = &v
	return s
}

func (s *ListAlarmEventRequest) SetPageSize(v string) *ListAlarmEventRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlarmEventRequest) SetReverse(v bool) *ListAlarmEventRequest {
	s.Reverse = &v
	return s
}

func (s *ListAlarmEventRequest) SetStartTime(v int64) *ListAlarmEventRequest {
	s.StartTime = &v
	return s
}

type ListAlarmEventResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *ListAlarmEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 27B1345D-5F71-5972-8E4C-AABA6C6232F0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAlarmEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmEventResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlarmEventResponseBody) SetCode(v int32) *ListAlarmEventResponseBody {
	s.Code = &v
	return s
}

func (s *ListAlarmEventResponseBody) SetData(v *ListAlarmEventResponseBodyData) *ListAlarmEventResponseBody {
	s.Data = v
	return s
}

func (s *ListAlarmEventResponseBody) SetMessage(v string) *ListAlarmEventResponseBody {
	s.Message = &v
	return s
}

func (s *ListAlarmEventResponseBody) SetRequestId(v string) *ListAlarmEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlarmEventResponseBody) SetSuccess(v bool) *ListAlarmEventResponseBody {
	s.Success = &v
	return s
}

type ListAlarmEventResponseBodyData struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Records  []*ListAlarmEventResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// example:
	//
	// 64
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListAlarmEventResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAlarmEventResponseBodyData) SetPageNumber(v int32) *ListAlarmEventResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListAlarmEventResponseBodyData) SetPageSize(v int32) *ListAlarmEventResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListAlarmEventResponseBodyData) SetRecords(v []*ListAlarmEventResponseBodyDataRecords) *ListAlarmEventResponseBodyData {
	s.Records = v
	return s
}

func (s *ListAlarmEventResponseBodyData) SetTotal(v int64) *ListAlarmEventResponseBodyData {
	s.Total = &v
	return s
}

type ListAlarmEventResponseBodyDataRecords struct {
	// example:
	//
	// webhook
	AlarmChannel *string `json:"AlarmChannel,omitempty" xml:"AlarmChannel,omitempty"`
	// example:
	//
	// zhangsan
	AlarmContacts *string `json:"AlarmContacts,omitempty" xml:"AlarmContacts,omitempty"`
	AlarmMessage  *string `json:"AlarmMessage,omitempty" xml:"AlarmMessage,omitempty"`
	// example:
	//
	// true
	AlarmStatus *string `json:"AlarmStatus,omitempty" xml:"AlarmStatus,omitempty"`
	// example:
	//
	// schedulerx3_fail_alarm
	AlarmType *string `json:"AlarmType,omitempty" xml:"AlarmType,omitempty"`
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// test-job
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// 2024-10-31 16:43:51
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s ListAlarmEventResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmEventResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListAlarmEventResponseBodyDataRecords) SetAlarmChannel(v string) *ListAlarmEventResponseBodyDataRecords {
	s.AlarmChannel = &v
	return s
}

func (s *ListAlarmEventResponseBodyDataRecords) SetAlarmContacts(v string) *ListAlarmEventResponseBodyDataRecords {
	s.AlarmContacts = &v
	return s
}

func (s *ListAlarmEventResponseBodyDataRecords) SetAlarmMessage(v string) *ListAlarmEventResponseBodyDataRecords {
	s.AlarmMessage = &v
	return s
}

func (s *ListAlarmEventResponseBodyDataRecords) SetAlarmStatus(v string) *ListAlarmEventResponseBodyDataRecords {
	s.AlarmStatus = &v
	return s
}

func (s *ListAlarmEventResponseBodyDataRecords) SetAlarmType(v string) *ListAlarmEventResponseBodyDataRecords {
	s.AlarmType = &v
	return s
}

func (s *ListAlarmEventResponseBodyDataRecords) SetAppName(v string) *ListAlarmEventResponseBodyDataRecords {
	s.AppName = &v
	return s
}

func (s *ListAlarmEventResponseBodyDataRecords) SetJobName(v string) *ListAlarmEventResponseBodyDataRecords {
	s.JobName = &v
	return s
}

func (s *ListAlarmEventResponseBodyDataRecords) SetTime(v string) *ListAlarmEventResponseBodyDataRecords {
	s.Time = &v
	return s
}

type ListAlarmEventResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlarmEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlarmEventResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmEventResponse) GoString() string {
	return s.String()
}

func (s *ListAlarmEventResponse) SetHeaders(v map[string]*string) *ListAlarmEventResponse {
	s.Headers = v
	return s
}

func (s *ListAlarmEventResponse) SetStatusCode(v int32) *ListAlarmEventResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlarmEventResponse) SetBody(v *ListAlarmEventResponseBody) *ListAlarmEventResponse {
	s.Body = v
	return s
}

type ListAppNamesRequest struct {
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ListAppNamesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppNamesRequest) GoString() string {
	return s.String()
}

func (s *ListAppNamesRequest) SetAppName(v string) *ListAppNamesRequest {
	s.AppName = &v
	return s
}

func (s *ListAppNamesRequest) SetClusterId(v string) *ListAppNamesRequest {
	s.ClusterId = &v
	return s
}

type ListAppNamesResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// .
	Data []*ListAppNamesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3881C59F-59F1-5B2E-8110-7D689CA9B207
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAppNamesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppNamesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppNamesResponseBody) SetCode(v int32) *ListAppNamesResponseBody {
	s.Code = &v
	return s
}

func (s *ListAppNamesResponseBody) SetData(v []*ListAppNamesResponseBodyData) *ListAppNamesResponseBody {
	s.Data = v
	return s
}

func (s *ListAppNamesResponseBody) SetMessage(v string) *ListAppNamesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAppNamesResponseBody) SetRequestId(v string) *ListAppNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppNamesResponseBody) SetSuccess(v bool) *ListAppNamesResponseBody {
	s.Success = &v
	return s
}

type ListAppNamesResponseBodyData struct {
	// example:
	//
	// test-app
	AppGroupId *string `json:"AppGroupId,omitempty" xml:"AppGroupId,omitempty"`
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 15
	Id    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListAppNamesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAppNamesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAppNamesResponseBodyData) SetAppGroupId(v string) *ListAppNamesResponseBodyData {
	s.AppGroupId = &v
	return s
}

func (s *ListAppNamesResponseBodyData) SetAppName(v string) *ListAppNamesResponseBodyData {
	s.AppName = &v
	return s
}

func (s *ListAppNamesResponseBodyData) SetId(v int64) *ListAppNamesResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListAppNamesResponseBodyData) SetTitle(v string) *ListAppNamesResponseBodyData {
	s.Title = &v
	return s
}

type ListAppNamesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppNamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppNamesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppNamesResponse) GoString() string {
	return s.String()
}

func (s *ListAppNamesResponse) SetHeaders(v map[string]*string) *ListAppNamesResponse {
	s.Headers = v
	return s
}

func (s *ListAppNamesResponse) SetStatusCode(v int32) *ListAppNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppNamesResponse) SetBody(v *ListAppNamesResponseBody) *ListAppNamesResponse {
	s.Body = v
	return s
}

type ListAppsRequest struct {
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Title    *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppsRequest) GoString() string {
	return s.String()
}

func (s *ListAppsRequest) SetAppName(v string) *ListAppsRequest {
	s.AppName = &v
	return s
}

func (s *ListAppsRequest) SetClusterId(v string) *ListAppsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListAppsRequest) SetPageNum(v int32) *ListAppsRequest {
	s.PageNum = &v
	return s
}

func (s *ListAppsRequest) SetPageSize(v int32) *ListAppsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppsRequest) SetTitle(v string) *ListAppsRequest {
	s.Title = &v
	return s
}

type ListAppsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *ListAppsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2C3E52FF-CBE9-5C0E-8252-37ACFF1F5EFF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBody) SetCode(v int32) *ListAppsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAppsResponseBody) SetData(v *ListAppsResponseBodyData) *ListAppsResponseBody {
	s.Data = v
	return s
}

func (s *ListAppsResponseBody) SetMessage(v string) *ListAppsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAppsResponseBody) SetRequestId(v string) *ListAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppsResponseBody) SetSuccess(v bool) *ListAppsResponseBody {
	s.Success = &v
	return s
}

type ListAppsResponseBodyData struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// -
	Records []*ListAppsResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListAppsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBodyData) SetPageNumber(v int32) *ListAppsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListAppsResponseBodyData) SetPageSize(v int32) *ListAppsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListAppsResponseBodyData) SetRecords(v []*ListAppsResponseBodyDataRecords) *ListAppsResponseBodyData {
	s.Records = v
	return s
}

func (s *ListAppsResponseBodyData) SetTotal(v int32) *ListAppsResponseBodyData {
	s.Total = &v
	return s
}

type ListAppsResponseBodyDataRecords struct {
	// AccessToken
	//
	// example:
	//
	// 2f4ddeab8e344ed68e0402cf9b8502ffv3
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 1827811800555555
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// true
	EnableLog *bool `json:"EnableLog,omitempty" xml:"EnableLog,omitempty"`
	// example:
	//
	// 1
	ExecutorNum *int64 `json:"ExecutorNum,omitempty" xml:"ExecutorNum,omitempty"`
	// example:
	//
	// 43885
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 10
	JobNum *int32 `json:"JobNum,omitempty" xml:"JobNum,omitempty"`
	// example:
	//
	// http://28.5.128.3:80
	Leader *string `json:"Leader,omitempty" xml:"Leader,omitempty"`
	// example:
	//
	// 100
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// example:
	//
	// 1000
	MaxJobs *int32  `json:"MaxJobs,omitempty" xml:"MaxJobs,omitempty"`
	Title   *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 1827811800555555
	Updater *string `json:"Updater,omitempty" xml:"Updater,omitempty"`
}

func (s ListAppsResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBodyDataRecords) SetAccessToken(v string) *ListAppsResponseBodyDataRecords {
	s.AccessToken = &v
	return s
}

func (s *ListAppsResponseBodyDataRecords) SetAppName(v string) *ListAppsResponseBodyDataRecords {
	s.AppName = &v
	return s
}

func (s *ListAppsResponseBodyDataRecords) SetCreator(v string) *ListAppsResponseBodyDataRecords {
	s.Creator = &v
	return s
}

func (s *ListAppsResponseBodyDataRecords) SetEnableLog(v bool) *ListAppsResponseBodyDataRecords {
	s.EnableLog = &v
	return s
}

func (s *ListAppsResponseBodyDataRecords) SetExecutorNum(v int64) *ListAppsResponseBodyDataRecords {
	s.ExecutorNum = &v
	return s
}

func (s *ListAppsResponseBodyDataRecords) SetId(v int64) *ListAppsResponseBodyDataRecords {
	s.Id = &v
	return s
}

func (s *ListAppsResponseBodyDataRecords) SetJobNum(v int32) *ListAppsResponseBodyDataRecords {
	s.JobNum = &v
	return s
}

func (s *ListAppsResponseBodyDataRecords) SetLeader(v string) *ListAppsResponseBodyDataRecords {
	s.Leader = &v
	return s
}

func (s *ListAppsResponseBodyDataRecords) SetMaxConcurrency(v int32) *ListAppsResponseBodyDataRecords {
	s.MaxConcurrency = &v
	return s
}

func (s *ListAppsResponseBodyDataRecords) SetMaxJobs(v int32) *ListAppsResponseBodyDataRecords {
	s.MaxJobs = &v
	return s
}

func (s *ListAppsResponseBodyDataRecords) SetTitle(v string) *ListAppsResponseBodyDataRecords {
	s.Title = &v
	return s
}

func (s *ListAppsResponseBodyDataRecords) SetUpdater(v string) *ListAppsResponseBodyDataRecords {
	s.Updater = &v
	return s
}

type ListAppsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponse) GoString() string {
	return s.String()
}

func (s *ListAppsResponse) SetHeaders(v map[string]*string) *ListAppsResponse {
	s.Headers = v
	return s
}

func (s *ListAppsResponse) SetStatusCode(v int32) *ListAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppsResponse) SetBody(v *ListAppsResponseBody) *ListAppsResponse {
	s.Body = v
	return s
}

type ListCalendarNamesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ListCalendarNamesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCalendarNamesRequest) GoString() string {
	return s.String()
}

func (s *ListCalendarNamesRequest) SetClusterId(v string) *ListCalendarNamesRequest {
	s.ClusterId = &v
	return s
}

type ListCalendarNamesResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AA3538A0-FBE6-5E31-AD88-A02C6FF0DACC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCalendarNamesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCalendarNamesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCalendarNamesResponseBody) SetCode(v int32) *ListCalendarNamesResponseBody {
	s.Code = &v
	return s
}

func (s *ListCalendarNamesResponseBody) SetData(v []*string) *ListCalendarNamesResponseBody {
	s.Data = v
	return s
}

func (s *ListCalendarNamesResponseBody) SetMessage(v string) *ListCalendarNamesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCalendarNamesResponseBody) SetRequestId(v string) *ListCalendarNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCalendarNamesResponseBody) SetSuccess(v bool) *ListCalendarNamesResponseBody {
	s.Success = &v
	return s
}

type ListCalendarNamesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCalendarNamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCalendarNamesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCalendarNamesResponse) GoString() string {
	return s.String()
}

func (s *ListCalendarNamesResponse) SetHeaders(v map[string]*string) *ListCalendarNamesResponse {
	s.Headers = v
	return s
}

func (s *ListCalendarNamesResponse) SetStatusCode(v int32) *ListCalendarNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCalendarNamesResponse) SetBody(v *ListCalendarNamesResponseBody) *ListCalendarNamesResponse {
	s.Body = v
	return s
}

type ListClustersRequest struct {
	// example:
	//
	// xxljob-d6a5243b6fa
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// cluster-test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// 5
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClustersRequest) GoString() string {
	return s.String()
}

func (s *ListClustersRequest) SetClusterId(v string) *ListClustersRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClustersRequest) SetClusterName(v string) *ListClustersRequest {
	s.ClusterName = &v
	return s
}

func (s *ListClustersRequest) SetPageNum(v int32) *ListClustersRequest {
	s.PageNum = &v
	return s
}

func (s *ListClustersRequest) SetPageSize(v int32) *ListClustersRequest {
	s.PageSize = &v
	return s
}

type ListClustersResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *ListClustersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 39938688-0BAB-5AD8-BF02-F4910FAC7589
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBody) SetCode(v int32) *ListClustersResponseBody {
	s.Code = &v
	return s
}

func (s *ListClustersResponseBody) SetData(v *ListClustersResponseBodyData) *ListClustersResponseBody {
	s.Data = v
	return s
}

func (s *ListClustersResponseBody) SetMessage(v string) *ListClustersResponseBody {
	s.Message = &v
	return s
}

func (s *ListClustersResponseBody) SetRequestId(v string) *ListClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClustersResponseBody) SetSuccess(v bool) *ListClustersResponseBody {
	s.Success = &v
	return s
}

type ListClustersResponseBodyData struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// -
	Records []*ListClustersResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// example:
	//
	// 30
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListClustersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyData) SetPageNumber(v int32) *ListClustersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListClustersResponseBodyData) SetPageSize(v int32) *ListClustersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListClustersResponseBodyData) SetRecords(v []*ListClustersResponseBodyDataRecords) *ListClustersResponseBodyData {
	s.Records = v
	return s
}

func (s *ListClustersResponseBodyData) SetTotal(v int32) *ListClustersResponseBodyData {
	s.Total = &v
	return s
}

type ListClustersResponseBodyDataRecords struct {
	// example:
	//
	// PREPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// xxljob-c20f7ec9a78
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// xxl-job-test-1730427510169
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// scx.small.x2
	ClusterSpec *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	// example:
	//
	// 2024-10-29 15:56:36
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2024-10-29 15:56:36
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// xxljob
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// example:
	//
	// 2.0.0
	EngineVersion  *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	InternetDomain *string `json:"InternetDomain,omitempty" xml:"InternetDomain,omitempty"`
	// example:
	//
	// http://xxljob-b9e19e46c4e.schedulerx.mse.aliyuncs.com
	IntranetDomain *string `json:"IntranetDomain,omitempty" xml:"IntranetDomain,omitempty"`
	// example:
	//
	// 1
	ProductType  *int32  `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SpInstanceId *string `json:"SpInstanceId,omitempty" xml:"SpInstanceId,omitempty"`
	// example:
	//
	// 1
	Status    *int32                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	VSwitches []*ListClustersResponseBodyDataRecordsVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// VPC ID
	//
	// example:
	//
	// vpc-bp1fxort6ag5h9752i305
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListClustersResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyDataRecords) SetChargeType(v string) *ListClustersResponseBodyDataRecords {
	s.ChargeType = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetClusterId(v string) *ListClustersResponseBodyDataRecords {
	s.ClusterId = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetClusterName(v string) *ListClustersResponseBodyDataRecords {
	s.ClusterName = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetClusterSpec(v string) *ListClustersResponseBodyDataRecords {
	s.ClusterSpec = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetCreateTime(v string) *ListClustersResponseBodyDataRecords {
	s.CreateTime = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetEndTime(v string) *ListClustersResponseBodyDataRecords {
	s.EndTime = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetEngineType(v string) *ListClustersResponseBodyDataRecords {
	s.EngineType = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetEngineVersion(v string) *ListClustersResponseBodyDataRecords {
	s.EngineVersion = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetInternetDomain(v string) *ListClustersResponseBodyDataRecords {
	s.InternetDomain = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetIntranetDomain(v string) *ListClustersResponseBodyDataRecords {
	s.IntranetDomain = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetProductType(v int32) *ListClustersResponseBodyDataRecords {
	s.ProductType = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetSpInstanceId(v string) *ListClustersResponseBodyDataRecords {
	s.SpInstanceId = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetStatus(v int32) *ListClustersResponseBodyDataRecords {
	s.Status = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetVSwitches(v []*ListClustersResponseBodyDataRecordsVSwitches) *ListClustersResponseBodyDataRecords {
	s.VSwitches = v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetVpcId(v string) *ListClustersResponseBodyDataRecords {
	s.VpcId = &v
	return s
}

type ListClustersResponseBodyDataRecordsVSwitches struct {
	// example:
	//
	// vsw-8vbl54xzux86usy61r5zm
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListClustersResponseBodyDataRecordsVSwitches) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyDataRecordsVSwitches) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyDataRecordsVSwitches) SetVSwitchId(v string) *ListClustersResponseBodyDataRecordsVSwitches {
	s.VSwitchId = &v
	return s
}

func (s *ListClustersResponseBodyDataRecordsVSwitches) SetZoneId(v string) *ListClustersResponseBodyDataRecordsVSwitches {
	s.ZoneId = &v
	return s
}

type ListClustersResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponse) GoString() string {
	return s.String()
}

func (s *ListClustersResponse) SetHeaders(v map[string]*string) *ListClustersResponse {
	s.Headers = v
	return s
}

func (s *ListClustersResponse) SetStatusCode(v int32) *ListClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClustersResponse) SetBody(v *ListClustersResponseBody) *ListClustersResponse {
	s.Body = v
	return s
}

type ListExecutorsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 74
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// gray
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s ListExecutorsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExecutorsRequest) GoString() string {
	return s.String()
}

func (s *ListExecutorsRequest) SetAppName(v string) *ListExecutorsRequest {
	s.AppName = &v
	return s
}

func (s *ListExecutorsRequest) SetClusterId(v string) *ListExecutorsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListExecutorsRequest) SetJobId(v int64) *ListExecutorsRequest {
	s.JobId = &v
	return s
}

func (s *ListExecutorsRequest) SetLabel(v string) *ListExecutorsRequest {
	s.Label = &v
	return s
}

type ListExecutorsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data []*ListExecutorsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 5EF879D0-3B43-5AD1-9BF7-52418F9C5E73
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListExecutorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExecutorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExecutorsResponseBody) SetCode(v int32) *ListExecutorsResponseBody {
	s.Code = &v
	return s
}

func (s *ListExecutorsResponseBody) SetData(v []*ListExecutorsResponseBodyData) *ListExecutorsResponseBody {
	s.Data = v
	return s
}

func (s *ListExecutorsResponseBody) SetMessage(v string) *ListExecutorsResponseBody {
	s.Message = &v
	return s
}

func (s *ListExecutorsResponseBody) SetRequestId(v string) *ListExecutorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExecutorsResponseBody) SetSuccess(v bool) *ListExecutorsResponseBody {
	s.Success = &v
	return s
}

type ListExecutorsResponseBodyData struct {
	// example:
	//
	// http://192.168.1.10:9999/
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// 192.168.1.10
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// true
	IsDesignated *bool `json:"IsDesignated,omitempty" xml:"IsDesignated,omitempty"`
	// example:
	//
	// gray
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// true
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// example:
	//
	// 9999
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// 2.0.2
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListExecutorsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListExecutorsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListExecutorsResponseBodyData) SetAddress(v string) *ListExecutorsResponseBodyData {
	s.Address = &v
	return s
}

func (s *ListExecutorsResponseBodyData) SetIp(v string) *ListExecutorsResponseBodyData {
	s.Ip = &v
	return s
}

func (s *ListExecutorsResponseBodyData) SetIsDesignated(v bool) *ListExecutorsResponseBodyData {
	s.IsDesignated = &v
	return s
}

func (s *ListExecutorsResponseBodyData) SetLabel(v string) *ListExecutorsResponseBodyData {
	s.Label = &v
	return s
}

func (s *ListExecutorsResponseBodyData) SetOnline(v bool) *ListExecutorsResponseBodyData {
	s.Online = &v
	return s
}

func (s *ListExecutorsResponseBodyData) SetPort(v int32) *ListExecutorsResponseBodyData {
	s.Port = &v
	return s
}

func (s *ListExecutorsResponseBodyData) SetVersion(v string) *ListExecutorsResponseBodyData {
	s.Version = &v
	return s
}

type ListExecutorsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExecutorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExecutorsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExecutorsResponse) GoString() string {
	return s.String()
}

func (s *ListExecutorsResponse) SetHeaders(v map[string]*string) *ListExecutorsResponse {
	s.Headers = v
	return s
}

func (s *ListExecutorsResponse) SetStatusCode(v int32) *ListExecutorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExecutorsResponse) SetBody(v *ListExecutorsResponseBody) *ListExecutorsResponse {
	s.Body = v
	return s
}

type ListJobExecutionsRequest struct {
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 2024-11-12 20:50:56
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1310630367761285120
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
	// example:
	//
	// 74
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// test-job
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2024-11-12 20:50:55
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 4
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListJobExecutionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobExecutionsRequest) GoString() string {
	return s.String()
}

func (s *ListJobExecutionsRequest) SetAppName(v string) *ListJobExecutionsRequest {
	s.AppName = &v
	return s
}

func (s *ListJobExecutionsRequest) SetClusterId(v string) *ListJobExecutionsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListJobExecutionsRequest) SetEndTime(v string) *ListJobExecutionsRequest {
	s.EndTime = &v
	return s
}

func (s *ListJobExecutionsRequest) SetJobExecutionId(v string) *ListJobExecutionsRequest {
	s.JobExecutionId = &v
	return s
}

func (s *ListJobExecutionsRequest) SetJobId(v int64) *ListJobExecutionsRequest {
	s.JobId = &v
	return s
}

func (s *ListJobExecutionsRequest) SetJobName(v string) *ListJobExecutionsRequest {
	s.JobName = &v
	return s
}

func (s *ListJobExecutionsRequest) SetPageNum(v int32) *ListJobExecutionsRequest {
	s.PageNum = &v
	return s
}

func (s *ListJobExecutionsRequest) SetPageSize(v int32) *ListJobExecutionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobExecutionsRequest) SetStartTime(v string) *ListJobExecutionsRequest {
	s.StartTime = &v
	return s
}

func (s *ListJobExecutionsRequest) SetStatus(v int32) *ListJobExecutionsRequest {
	s.Status = &v
	return s
}

type ListJobExecutionsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *ListJobExecutionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6BCE89B3-E882-511D-9A75-D452A56EC4B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListJobExecutionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobExecutionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobExecutionsResponseBody) SetCode(v int32) *ListJobExecutionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListJobExecutionsResponseBody) SetData(v *ListJobExecutionsResponseBodyData) *ListJobExecutionsResponseBody {
	s.Data = v
	return s
}

func (s *ListJobExecutionsResponseBody) SetMessage(v string) *ListJobExecutionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListJobExecutionsResponseBody) SetRequestId(v string) *ListJobExecutionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobExecutionsResponseBody) SetSuccess(v bool) *ListJobExecutionsResponseBody {
	s.Success = &v
	return s
}

type ListJobExecutionsResponseBodyData struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// -
	Records []*ListJobExecutionsResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListJobExecutionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListJobExecutionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListJobExecutionsResponseBodyData) SetPageNumber(v int32) *ListJobExecutionsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListJobExecutionsResponseBodyData) SetPageSize(v int32) *ListJobExecutionsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListJobExecutionsResponseBodyData) SetRecords(v []*ListJobExecutionsResponseBodyDataRecords) *ListJobExecutionsResponseBodyData {
	s.Records = v
	return s
}

func (s *ListJobExecutionsResponseBodyData) SetTotal(v int32) *ListJobExecutionsResponseBodyData {
	s.Total = &v
	return s
}

type ListJobExecutionsResponseBodyDataRecords struct {
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 1
	Attempt *int32 `json:"Attempt,omitempty" xml:"Attempt,omitempty"`
	// example:
	//
	// 2024-11-12 14:52:42
	DataTime *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	// example:
	//
	// 10
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 2024-11-12 14:52:42
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1827811800526000
	Executor *string `json:"Executor,omitempty" xml:"Executor,omitempty"`
	// example:
	//
	// 1310630367761285120
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
	// example:
	//
	// 74
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// test-job
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// xxljob
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// example:
	//
	// name=zhangsan
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 1
	RouteStrategy *int32 `json:"RouteStrategy,omitempty" xml:"RouteStrategy,omitempty"`
	// example:
	//
	// 2024-11-12 14:52:42
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// example:
	//
	// 28.0.168.46
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// example:
	//
	// 4
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// example:
	//
	// 1
	TriggerType *int32 `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// example:
	//
	// http://192.168.1.9:9999/
	WorkAddr *string `json:"WorkAddr,omitempty" xml:"WorkAddr,omitempty"`
}

func (s ListJobExecutionsResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s ListJobExecutionsResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetAppName(v string) *ListJobExecutionsResponseBodyDataRecords {
	s.AppName = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetAttempt(v int32) *ListJobExecutionsResponseBodyDataRecords {
	s.Attempt = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetDataTime(v string) *ListJobExecutionsResponseBodyDataRecords {
	s.DataTime = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetDuration(v int64) *ListJobExecutionsResponseBodyDataRecords {
	s.Duration = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetEndTime(v string) *ListJobExecutionsResponseBodyDataRecords {
	s.EndTime = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetExecutor(v string) *ListJobExecutionsResponseBodyDataRecords {
	s.Executor = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetJobExecutionId(v string) *ListJobExecutionsResponseBodyDataRecords {
	s.JobExecutionId = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetJobId(v int64) *ListJobExecutionsResponseBodyDataRecords {
	s.JobId = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetJobName(v string) *ListJobExecutionsResponseBodyDataRecords {
	s.JobName = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetJobType(v string) *ListJobExecutionsResponseBodyDataRecords {
	s.JobType = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetParameters(v string) *ListJobExecutionsResponseBodyDataRecords {
	s.Parameters = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetResult(v string) *ListJobExecutionsResponseBodyDataRecords {
	s.Result = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetRouteStrategy(v int32) *ListJobExecutionsResponseBodyDataRecords {
	s.RouteStrategy = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetScheduleTime(v string) *ListJobExecutionsResponseBodyDataRecords {
	s.ScheduleTime = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetServerIp(v string) *ListJobExecutionsResponseBodyDataRecords {
	s.ServerIp = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetStatus(v int32) *ListJobExecutionsResponseBodyDataRecords {
	s.Status = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetTimeType(v int32) *ListJobExecutionsResponseBodyDataRecords {
	s.TimeType = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetTriggerType(v int32) *ListJobExecutionsResponseBodyDataRecords {
	s.TriggerType = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetWorkAddr(v string) *ListJobExecutionsResponseBodyDataRecords {
	s.WorkAddr = &v
	return s
}

type ListJobExecutionsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobExecutionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobExecutionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobExecutionsResponse) GoString() string {
	return s.String()
}

func (s *ListJobExecutionsResponse) SetHeaders(v map[string]*string) *ListJobExecutionsResponse {
	s.Headers = v
	return s
}

func (s *ListJobExecutionsResponse) SetStatusCode(v int32) *ListJobExecutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobExecutionsResponse) SetBody(v *ListJobExecutionsResponseBody) *ListJobExecutionsResponse {
	s.Body = v
	return s
}

type ListJobsRequest struct {
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// jobDemoHandler
	JobHandler *string `json:"JobHandler,omitempty" xml:"JobHandler,omitempty"`
	// example:
	//
	// 10
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// job01
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 0
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) SetAppName(v string) *ListJobsRequest {
	s.AppName = &v
	return s
}

func (s *ListJobsRequest) SetClusterId(v string) *ListJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListJobsRequest) SetDescription(v string) *ListJobsRequest {
	s.Description = &v
	return s
}

func (s *ListJobsRequest) SetJobHandler(v string) *ListJobsRequest {
	s.JobHandler = &v
	return s
}

func (s *ListJobsRequest) SetJobId(v int64) *ListJobsRequest {
	s.JobId = &v
	return s
}

func (s *ListJobsRequest) SetJobName(v string) *ListJobsRequest {
	s.JobName = &v
	return s
}

func (s *ListJobsRequest) SetPageNum(v int32) *ListJobsRequest {
	s.PageNum = &v
	return s
}

func (s *ListJobsRequest) SetPageSize(v int32) *ListJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsRequest) SetStatus(v string) *ListJobsRequest {
	s.Status = &v
	return s
}

type ListJobsResponseBody struct {
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
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
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// -
	Records []*ListJobsResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// example:
	//
	// 65
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListJobsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyData) SetPageNumber(v int32) *ListJobsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListJobsResponseBodyData) SetPageSize(v int32) *ListJobsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListJobsResponseBodyData) SetRecords(v []*ListJobsResponseBodyDataRecords) *ListJobsResponseBodyData {
	s.Records = v
	return s
}

func (s *ListJobsResponseBodyData) SetTotal(v int32) *ListJobsResponseBodyData {
	s.Total = &v
	return s
}

type ListJobsResponseBodyDataRecords struct {
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 30
	AttemptInterval *int32 `json:"AttemptInterval,omitempty" xml:"AttemptInterval,omitempty"`
	// example:
	//
	// work-day
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// example:
	//
	// {"cleanMode":"NUM_ONLY","totalRemain":300}
	CleanMode *string `json:"CleanMode,omitempty" xml:"CleanMode,omitempty"`
	// example:
	//
	// 1963096506470832
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 3
	CurrentExecuteStatus *int32 `json:"CurrentExecuteStatus,omitempty" xml:"CurrentExecuteStatus,omitempty"`
	// example:
	//
	// 3
	DataOffset            *int32  `json:"DataOffset,omitempty" xml:"DataOffset,omitempty"`
	Description           *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExecutorBlockStrategy *string `json:"ExecutorBlockStrategy,omitempty" xml:"ExecutorBlockStrategy,omitempty"`
	// example:
	//
	// jobDemoHandler
	JobHandler *string `json:"JobHandler,omitempty" xml:"JobHandler,omitempty"`
	// example:
	//
	// 74
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// xxljob
	JobType            *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	LastExecuteEndTime *string `json:"LastExecuteEndTime,omitempty" xml:"LastExecuteEndTime,omitempty"`
	LastExecuteStatus  *int32  `json:"LastExecuteStatus,omitempty" xml:"LastExecuteStatus,omitempty"`
	// example:
	//
	// 5
	MaxAttempt *int32 `json:"MaxAttempt,omitempty" xml:"MaxAttempt,omitempty"`
	// example:
	//
	// 100
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// example:
	//
	// job01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {"failLimitTimes":1,"failEnable":true,"timeoutKillEnable":false,"missWorkerEnable":true,"timeoutEnable":true,"sendChannel":"","timeout":300,"successNotice":false}
	NoticeConfig   *string `json:"NoticeConfig,omitempty" xml:"NoticeConfig,omitempty"`
	NoticeContacts *string `json:"NoticeContacts,omitempty" xml:"NoticeContacts,omitempty"`
	// example:
	//
	// name=10
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 1
	RouteStrategy *int32 `json:"RouteStrategy,omitempty" xml:"RouteStrategy,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 0 0 12 	- 	- ?
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// example:
	//
	// HangKong
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	// example:
	//
	// HangKong
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// example:
	//
	// 1963096506470832
	Updater *string `json:"Updater,omitempty" xml:"Updater,omitempty"`
	Xattrs  *string `json:"Xattrs,omitempty" xml:"Xattrs,omitempty"`
}

func (s ListJobsResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyDataRecords) SetAppName(v string) *ListJobsResponseBodyDataRecords {
	s.AppName = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetAttemptInterval(v int32) *ListJobsResponseBodyDataRecords {
	s.AttemptInterval = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetCalendar(v string) *ListJobsResponseBodyDataRecords {
	s.Calendar = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetCleanMode(v string) *ListJobsResponseBodyDataRecords {
	s.CleanMode = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetCreator(v string) *ListJobsResponseBodyDataRecords {
	s.Creator = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetCurrentExecuteStatus(v int32) *ListJobsResponseBodyDataRecords {
	s.CurrentExecuteStatus = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetDataOffset(v int32) *ListJobsResponseBodyDataRecords {
	s.DataOffset = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetDescription(v string) *ListJobsResponseBodyDataRecords {
	s.Description = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetExecutorBlockStrategy(v string) *ListJobsResponseBodyDataRecords {
	s.ExecutorBlockStrategy = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetJobHandler(v string) *ListJobsResponseBodyDataRecords {
	s.JobHandler = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetJobId(v int64) *ListJobsResponseBodyDataRecords {
	s.JobId = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetJobType(v string) *ListJobsResponseBodyDataRecords {
	s.JobType = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetLastExecuteEndTime(v string) *ListJobsResponseBodyDataRecords {
	s.LastExecuteEndTime = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetLastExecuteStatus(v int32) *ListJobsResponseBodyDataRecords {
	s.LastExecuteStatus = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetMaxAttempt(v int32) *ListJobsResponseBodyDataRecords {
	s.MaxAttempt = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetMaxConcurrency(v int32) *ListJobsResponseBodyDataRecords {
	s.MaxConcurrency = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetName(v string) *ListJobsResponseBodyDataRecords {
	s.Name = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetNoticeConfig(v string) *ListJobsResponseBodyDataRecords {
	s.NoticeConfig = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetNoticeContacts(v string) *ListJobsResponseBodyDataRecords {
	s.NoticeContacts = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetParameters(v string) *ListJobsResponseBodyDataRecords {
	s.Parameters = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetPriority(v int32) *ListJobsResponseBodyDataRecords {
	s.Priority = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetRouteStrategy(v int32) *ListJobsResponseBodyDataRecords {
	s.RouteStrategy = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetStatus(v int32) *ListJobsResponseBodyDataRecords {
	s.Status = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetTimeExpression(v string) *ListJobsResponseBodyDataRecords {
	s.TimeExpression = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetTimeType(v int32) *ListJobsResponseBodyDataRecords {
	s.TimeType = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetTimeZone(v string) *ListJobsResponseBodyDataRecords {
	s.TimeZone = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetTimezone(v string) *ListJobsResponseBodyDataRecords {
	s.Timezone = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetUpdater(v string) *ListJobsResponseBodyDataRecords {
	s.Updater = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetXattrs(v string) *ListJobsResponseBodyDataRecords {
	s.Xattrs = &v
	return s
}

type ListJobsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListJobsResponse) SetStatusCode(v int32) *ListJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobsResponse) SetBody(v *ListJobsResponseBody) *ListJobsResponse {
	s.Body = v
	return s
}

type ListLablesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 15
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s ListLablesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLablesRequest) GoString() string {
	return s.String()
}

func (s *ListLablesRequest) SetAppName(v string) *ListLablesRequest {
	s.AppName = &v
	return s
}

func (s *ListLablesRequest) SetClusterId(v string) *ListLablesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListLablesRequest) SetJobId(v int64) *ListLablesRequest {
	s.JobId = &v
	return s
}

type ListLablesResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data []*ListLablesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9079A828-9138-50F1-801E-F2BC3D222A06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListLablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLablesResponseBody) SetCode(v int32) *ListLablesResponseBody {
	s.Code = &v
	return s
}

func (s *ListLablesResponseBody) SetData(v []*ListLablesResponseBodyData) *ListLablesResponseBody {
	s.Data = v
	return s
}

func (s *ListLablesResponseBody) SetMessage(v string) *ListLablesResponseBody {
	s.Message = &v
	return s
}

func (s *ListLablesResponseBody) SetRequestId(v string) *ListLablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLablesResponseBody) SetSuccess(v bool) *ListLablesResponseBody {
	s.Success = &v
	return s
}

type ListLablesResponseBodyData struct {
	// example:
	//
	// true
	IsDesignated *bool `json:"IsDesignated,omitempty" xml:"IsDesignated,omitempty"`
	// example:
	//
	// gray
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// true
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// example:
	//
	// 2
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListLablesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListLablesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLablesResponseBodyData) SetIsDesignated(v bool) *ListLablesResponseBodyData {
	s.IsDesignated = &v
	return s
}

func (s *ListLablesResponseBodyData) SetLabel(v string) *ListLablesResponseBodyData {
	s.Label = &v
	return s
}

func (s *ListLablesResponseBodyData) SetOnline(v bool) *ListLablesResponseBodyData {
	s.Online = &v
	return s
}

func (s *ListLablesResponseBodyData) SetSize(v int32) *ListLablesResponseBodyData {
	s.Size = &v
	return s
}

type ListLablesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLablesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLablesResponse) GoString() string {
	return s.String()
}

func (s *ListLablesResponse) SetHeaders(v map[string]*string) *ListLablesResponse {
	s.Headers = v
	return s
}

func (s *ListLablesResponse) SetStatusCode(v int32) *ListLablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLablesResponse) SetBody(v *ListLablesResponseBody) *ListLablesResponse {
	s.Body = v
	return s
}

type ListRegionZoneResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data []*ListRegionZoneResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// IllegalRequest
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 438737AC-760A-57D9-B646-B7EF79426243
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRegionZoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRegionZoneResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionZoneResponseBody) SetCode(v int32) *ListRegionZoneResponseBody {
	s.Code = &v
	return s
}

func (s *ListRegionZoneResponseBody) SetData(v []*ListRegionZoneResponseBodyData) *ListRegionZoneResponseBody {
	s.Data = v
	return s
}

func (s *ListRegionZoneResponseBody) SetErrorCode(v string) *ListRegionZoneResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRegionZoneResponseBody) SetMessage(v string) *ListRegionZoneResponseBody {
	s.Message = &v
	return s
}

func (s *ListRegionZoneResponseBody) SetRequestId(v string) *ListRegionZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRegionZoneResponseBody) SetSuccess(v bool) *ListRegionZoneResponseBody {
	s.Success = &v
	return s
}

type ListRegionZoneResponseBodyData struct {
	// example:
	//
	// E
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// zone id
	//
	// example:
	//
	// cn-beijing-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListRegionZoneResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListRegionZoneResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRegionZoneResponseBodyData) SetLocalName(v string) *ListRegionZoneResponseBodyData {
	s.LocalName = &v
	return s
}

func (s *ListRegionZoneResponseBodyData) SetZoneId(v string) *ListRegionZoneResponseBodyData {
	s.ZoneId = &v
	return s
}

type ListRegionZoneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRegionZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRegionZoneResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRegionZoneResponse) GoString() string {
	return s.String()
}

func (s *ListRegionZoneResponse) SetHeaders(v map[string]*string) *ListRegionZoneResponse {
	s.Headers = v
	return s
}

func (s *ListRegionZoneResponse) SetStatusCode(v int32) *ListRegionZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRegionZoneResponse) SetBody(v *ListRegionZoneResponseBody) *ListRegionZoneResponse {
	s.Body = v
	return s
}

type ListRegionsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// -
	Regions []*ListRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// example:
	//
	// AFD5B166-4A7D-50DF-91BF-EFAFD41F7335
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) SetCode(v int32) *ListRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListRegionsResponseBody) SetMessage(v string) *ListRegionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListRegionsResponseBody) SetRegions(v []*ListRegionsResponseBodyRegions) *ListRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRegionsResponseBody) SetSuccess(v bool) *ListRegionsResponseBody {
	s.Success = &v
	return s
}

type ListRegionsResponseBodyRegions struct {
	// example:
	//
	// China (Hangzhou)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// endpoint
	//
	// example:
	//
	// schedulerx3.cn-hangzhou.aliyuncs.com
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyRegions) SetLocalName(v string) *ListRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *ListRegionsResponseBodyRegions) SetRegionEndpoint(v string) *ListRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *ListRegionsResponseBodyRegions) SetRegionId(v string) *ListRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type ListRegionsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListRegionsResponse) SetHeaders(v map[string]*string) *ListRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListRegionsResponse) SetStatusCode(v int32) *ListRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRegionsResponse) SetBody(v *ListRegionsResponseBody) *ListRegionsResponse {
	s.Body = v
	return s
}

type ListScheduleEventRequest struct {
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 1728872796295
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// INFO
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 1310630367761285120
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
	// example:
	//
	// test-job
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// hello word
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// true
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// example:
	//
	// 1581317873000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListScheduleEventRequest) String() string {
	return tea.Prettify(s)
}

func (s ListScheduleEventRequest) GoString() string {
	return s.String()
}

func (s *ListScheduleEventRequest) SetAppName(v string) *ListScheduleEventRequest {
	s.AppName = &v
	return s
}

func (s *ListScheduleEventRequest) SetClusterId(v string) *ListScheduleEventRequest {
	s.ClusterId = &v
	return s
}

func (s *ListScheduleEventRequest) SetEndTime(v int64) *ListScheduleEventRequest {
	s.EndTime = &v
	return s
}

func (s *ListScheduleEventRequest) SetEvent(v string) *ListScheduleEventRequest {
	s.Event = &v
	return s
}

func (s *ListScheduleEventRequest) SetJobExecutionId(v string) *ListScheduleEventRequest {
	s.JobExecutionId = &v
	return s
}

func (s *ListScheduleEventRequest) SetJobName(v string) *ListScheduleEventRequest {
	s.JobName = &v
	return s
}

func (s *ListScheduleEventRequest) SetKeyword(v string) *ListScheduleEventRequest {
	s.Keyword = &v
	return s
}

func (s *ListScheduleEventRequest) SetPageNum(v int32) *ListScheduleEventRequest {
	s.PageNum = &v
	return s
}

func (s *ListScheduleEventRequest) SetPageSize(v int32) *ListScheduleEventRequest {
	s.PageSize = &v
	return s
}

func (s *ListScheduleEventRequest) SetReverse(v bool) *ListScheduleEventRequest {
	s.Reverse = &v
	return s
}

func (s *ListScheduleEventRequest) SetStartTime(v int64) *ListScheduleEventRequest {
	s.StartTime = &v
	return s
}

type ListScheduleEventResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *ListScheduleEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// B880122A-B0E4-52E8-8F54-87DB7779EB74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListScheduleEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListScheduleEventResponseBody) GoString() string {
	return s.String()
}

func (s *ListScheduleEventResponseBody) SetCode(v int32) *ListScheduleEventResponseBody {
	s.Code = &v
	return s
}

func (s *ListScheduleEventResponseBody) SetData(v *ListScheduleEventResponseBodyData) *ListScheduleEventResponseBody {
	s.Data = v
	return s
}

func (s *ListScheduleEventResponseBody) SetMessage(v string) *ListScheduleEventResponseBody {
	s.Message = &v
	return s
}

func (s *ListScheduleEventResponseBody) SetRequestId(v string) *ListScheduleEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScheduleEventResponseBody) SetSuccess(v bool) *ListScheduleEventResponseBody {
	s.Success = &v
	return s
}

type ListScheduleEventResponseBodyData struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Records  []*ListScheduleEventResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListScheduleEventResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListScheduleEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListScheduleEventResponseBodyData) SetPageNumber(v int32) *ListScheduleEventResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListScheduleEventResponseBodyData) SetPageSize(v int32) *ListScheduleEventResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListScheduleEventResponseBodyData) SetRecords(v []*ListScheduleEventResponseBodyDataRecords) *ListScheduleEventResponseBodyData {
	s.Records = v
	return s
}

func (s *ListScheduleEventResponseBodyData) SetTotal(v int64) *ListScheduleEventResponseBodyData {
	s.Total = &v
	return s
}

type ListScheduleEventResponseBodyDataRecords struct {
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// hello word
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// INFO
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// 130
	//
	// example:
	//
	// 1310630367761285120
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
	// example:
	//
	// test-job
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// 2024-10-31 16:43:51
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// example:
	//
	// http://192.168.1.5:9999/
	WorkerAddr *string `json:"WorkerAddr,omitempty" xml:"WorkerAddr,omitempty"`
}

func (s ListScheduleEventResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s ListScheduleEventResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListScheduleEventResponseBodyDataRecords) SetAppName(v string) *ListScheduleEventResponseBodyDataRecords {
	s.AppName = &v
	return s
}

func (s *ListScheduleEventResponseBodyDataRecords) SetContent(v string) *ListScheduleEventResponseBodyDataRecords {
	s.Content = &v
	return s
}

func (s *ListScheduleEventResponseBodyDataRecords) SetEvent(v string) *ListScheduleEventResponseBodyDataRecords {
	s.Event = &v
	return s
}

func (s *ListScheduleEventResponseBodyDataRecords) SetJobExecutionId(v string) *ListScheduleEventResponseBodyDataRecords {
	s.JobExecutionId = &v
	return s
}

func (s *ListScheduleEventResponseBodyDataRecords) SetJobName(v string) *ListScheduleEventResponseBodyDataRecords {
	s.JobName = &v
	return s
}

func (s *ListScheduleEventResponseBodyDataRecords) SetTime(v string) *ListScheduleEventResponseBodyDataRecords {
	s.Time = &v
	return s
}

func (s *ListScheduleEventResponseBodyDataRecords) SetWorkerAddr(v string) *ListScheduleEventResponseBodyDataRecords {
	s.WorkerAddr = &v
	return s
}

type ListScheduleEventResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScheduleEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScheduleEventResponse) String() string {
	return tea.Prettify(s)
}

func (s ListScheduleEventResponse) GoString() string {
	return s.String()
}

func (s *ListScheduleEventResponse) SetHeaders(v map[string]*string) *ListScheduleEventResponse {
	s.Headers = v
	return s
}

func (s *ListScheduleEventResponse) SetStatusCode(v int32) *ListScheduleEventResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScheduleEventResponse) SetBody(v *ListScheduleEventResponseBody) *ListScheduleEventResponse {
	s.Body = v
	return s
}

type ListScheduleTimesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// workday
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0 0/10 	- 	- 	- ?
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// example:
	//
	// Asia/Beijing
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s ListScheduleTimesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListScheduleTimesRequest) GoString() string {
	return s.String()
}

func (s *ListScheduleTimesRequest) SetAppName(v string) *ListScheduleTimesRequest {
	s.AppName = &v
	return s
}

func (s *ListScheduleTimesRequest) SetCalendar(v string) *ListScheduleTimesRequest {
	s.Calendar = &v
	return s
}

func (s *ListScheduleTimesRequest) SetClusterId(v string) *ListScheduleTimesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListScheduleTimesRequest) SetTimeExpression(v string) *ListScheduleTimesRequest {
	s.TimeExpression = &v
	return s
}

func (s *ListScheduleTimesRequest) SetTimeType(v int32) *ListScheduleTimesRequest {
	s.TimeType = &v
	return s
}

func (s *ListScheduleTimesRequest) SetTimeZone(v string) *ListScheduleTimesRequest {
	s.TimeZone = &v
	return s
}

type ListScheduleTimesResponseBody struct {
	// example:
	//
	// 200
	Code *int32    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9A48E22F-F30A-5CE5-AC7A-E0FED1B6942E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListScheduleTimesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListScheduleTimesResponseBody) GoString() string {
	return s.String()
}

func (s *ListScheduleTimesResponseBody) SetCode(v int32) *ListScheduleTimesResponseBody {
	s.Code = &v
	return s
}

func (s *ListScheduleTimesResponseBody) SetData(v []*string) *ListScheduleTimesResponseBody {
	s.Data = v
	return s
}

func (s *ListScheduleTimesResponseBody) SetMessage(v string) *ListScheduleTimesResponseBody {
	s.Message = &v
	return s
}

func (s *ListScheduleTimesResponseBody) SetRequestId(v string) *ListScheduleTimesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScheduleTimesResponseBody) SetSuccess(v bool) *ListScheduleTimesResponseBody {
	s.Success = &v
	return s
}

type ListScheduleTimesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScheduleTimesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScheduleTimesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListScheduleTimesResponse) GoString() string {
	return s.String()
}

func (s *ListScheduleTimesResponse) SetHeaders(v map[string]*string) *ListScheduleTimesResponse {
	s.Headers = v
	return s
}

func (s *ListScheduleTimesResponse) SetStatusCode(v int32) *ListScheduleTimesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScheduleTimesResponse) SetBody(v *ListScheduleTimesResponseBody) *ListScheduleTimesResponse {
	s.Body = v
	return s
}

type OperateDesignateExecutorsRequest struct {
	// This parameter is required.
	AddressList []*string `json:"AddressList,omitempty" xml:"AddressList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	DesignateType *int32 `json:"DesignateType,omitempty" xml:"DesignateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 74
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// true
	Transferable *bool `json:"Transferable,omitempty" xml:"Transferable,omitempty"`
}

func (s OperateDesignateExecutorsRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateDesignateExecutorsRequest) GoString() string {
	return s.String()
}

func (s *OperateDesignateExecutorsRequest) SetAddressList(v []*string) *OperateDesignateExecutorsRequest {
	s.AddressList = v
	return s
}

func (s *OperateDesignateExecutorsRequest) SetAppName(v string) *OperateDesignateExecutorsRequest {
	s.AppName = &v
	return s
}

func (s *OperateDesignateExecutorsRequest) SetClusterId(v string) *OperateDesignateExecutorsRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateDesignateExecutorsRequest) SetDesignateType(v int32) *OperateDesignateExecutorsRequest {
	s.DesignateType = &v
	return s
}

func (s *OperateDesignateExecutorsRequest) SetJobId(v int64) *OperateDesignateExecutorsRequest {
	s.JobId = &v
	return s
}

func (s *OperateDesignateExecutorsRequest) SetTransferable(v bool) *OperateDesignateExecutorsRequest {
	s.Transferable = &v
	return s
}

type OperateDesignateExecutorsShrinkRequest struct {
	// This parameter is required.
	AddressListShrink *string `json:"AddressList,omitempty" xml:"AddressList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	DesignateType *int32 `json:"DesignateType,omitempty" xml:"DesignateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 74
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// true
	Transferable *bool `json:"Transferable,omitempty" xml:"Transferable,omitempty"`
}

func (s OperateDesignateExecutorsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateDesignateExecutorsShrinkRequest) GoString() string {
	return s.String()
}

func (s *OperateDesignateExecutorsShrinkRequest) SetAddressListShrink(v string) *OperateDesignateExecutorsShrinkRequest {
	s.AddressListShrink = &v
	return s
}

func (s *OperateDesignateExecutorsShrinkRequest) SetAppName(v string) *OperateDesignateExecutorsShrinkRequest {
	s.AppName = &v
	return s
}

func (s *OperateDesignateExecutorsShrinkRequest) SetClusterId(v string) *OperateDesignateExecutorsShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateDesignateExecutorsShrinkRequest) SetDesignateType(v int32) *OperateDesignateExecutorsShrinkRequest {
	s.DesignateType = &v
	return s
}

func (s *OperateDesignateExecutorsShrinkRequest) SetJobId(v int64) *OperateDesignateExecutorsShrinkRequest {
	s.JobId = &v
	return s
}

func (s *OperateDesignateExecutorsShrinkRequest) SetTransferable(v bool) *OperateDesignateExecutorsShrinkRequest {
	s.Transferable = &v
	return s
}

type OperateDesignateExecutorsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AFD5B166-4A7D-50DF-91BF-EFAFD41F7335
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateDesignateExecutorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperateDesignateExecutorsResponseBody) GoString() string {
	return s.String()
}

func (s *OperateDesignateExecutorsResponseBody) SetCode(v int32) *OperateDesignateExecutorsResponseBody {
	s.Code = &v
	return s
}

func (s *OperateDesignateExecutorsResponseBody) SetMessage(v string) *OperateDesignateExecutorsResponseBody {
	s.Message = &v
	return s
}

func (s *OperateDesignateExecutorsResponseBody) SetRequestId(v string) *OperateDesignateExecutorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateDesignateExecutorsResponseBody) SetSuccess(v bool) *OperateDesignateExecutorsResponseBody {
	s.Success = &v
	return s
}

type OperateDesignateExecutorsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateDesignateExecutorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateDesignateExecutorsResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateDesignateExecutorsResponse) GoString() string {
	return s.String()
}

func (s *OperateDesignateExecutorsResponse) SetHeaders(v map[string]*string) *OperateDesignateExecutorsResponse {
	s.Headers = v
	return s
}

func (s *OperateDesignateExecutorsResponse) SetStatusCode(v int32) *OperateDesignateExecutorsResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateDesignateExecutorsResponse) SetBody(v *OperateDesignateExecutorsResponseBody) *OperateDesignateExecutorsResponse {
	s.Body = v
	return s
}

type OperateDisableJobsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// -
	JobIds []*int64 `json:"JobIds,omitempty" xml:"JobIds,omitempty" type:"Repeated"`
}

func (s OperateDisableJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateDisableJobsRequest) GoString() string {
	return s.String()
}

func (s *OperateDisableJobsRequest) SetAppName(v string) *OperateDisableJobsRequest {
	s.AppName = &v
	return s
}

func (s *OperateDisableJobsRequest) SetClusterId(v string) *OperateDisableJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateDisableJobsRequest) SetJobIds(v []*int64) *OperateDisableJobsRequest {
	s.JobIds = v
	return s
}

type OperateDisableJobsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// -
	JobIdsShrink *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
}

func (s OperateDisableJobsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateDisableJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *OperateDisableJobsShrinkRequest) SetAppName(v string) *OperateDisableJobsShrinkRequest {
	s.AppName = &v
	return s
}

func (s *OperateDisableJobsShrinkRequest) SetClusterId(v string) *OperateDisableJobsShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateDisableJobsShrinkRequest) SetJobIdsShrink(v string) *OperateDisableJobsShrinkRequest {
	s.JobIdsShrink = &v
	return s
}

type OperateDisableJobsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 29ED6209-5DE6-5E1D-89B0-B7B1D823A1BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateDisableJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperateDisableJobsResponseBody) GoString() string {
	return s.String()
}

func (s *OperateDisableJobsResponseBody) SetCode(v int32) *OperateDisableJobsResponseBody {
	s.Code = &v
	return s
}

func (s *OperateDisableJobsResponseBody) SetMessage(v string) *OperateDisableJobsResponseBody {
	s.Message = &v
	return s
}

func (s *OperateDisableJobsResponseBody) SetRequestId(v string) *OperateDisableJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateDisableJobsResponseBody) SetSuccess(v bool) *OperateDisableJobsResponseBody {
	s.Success = &v
	return s
}

type OperateDisableJobsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateDisableJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateDisableJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateDisableJobsResponse) GoString() string {
	return s.String()
}

func (s *OperateDisableJobsResponse) SetHeaders(v map[string]*string) *OperateDisableJobsResponse {
	s.Headers = v
	return s
}

func (s *OperateDisableJobsResponse) SetStatusCode(v int32) *OperateDisableJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateDisableJobsResponse) SetBody(v *OperateDisableJobsResponseBody) *OperateDisableJobsResponse {
	s.Body = v
	return s
}

type OperateEnableJobsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// -
	JobIds []*int64 `json:"JobIds,omitempty" xml:"JobIds,omitempty" type:"Repeated"`
}

func (s OperateEnableJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateEnableJobsRequest) GoString() string {
	return s.String()
}

func (s *OperateEnableJobsRequest) SetAppName(v string) *OperateEnableJobsRequest {
	s.AppName = &v
	return s
}

func (s *OperateEnableJobsRequest) SetClusterId(v string) *OperateEnableJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateEnableJobsRequest) SetJobIds(v []*int64) *OperateEnableJobsRequest {
	s.JobIds = v
	return s
}

type OperateEnableJobsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// -
	JobIdsShrink *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
}

func (s OperateEnableJobsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateEnableJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *OperateEnableJobsShrinkRequest) SetAppName(v string) *OperateEnableJobsShrinkRequest {
	s.AppName = &v
	return s
}

func (s *OperateEnableJobsShrinkRequest) SetClusterId(v string) *OperateEnableJobsShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateEnableJobsShrinkRequest) SetJobIdsShrink(v string) *OperateEnableJobsShrinkRequest {
	s.JobIdsShrink = &v
	return s
}

type OperateEnableJobsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4CC4132F-B798-5D6E-9F06-D44B33E417E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateEnableJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperateEnableJobsResponseBody) GoString() string {
	return s.String()
}

func (s *OperateEnableJobsResponseBody) SetCode(v int32) *OperateEnableJobsResponseBody {
	s.Code = &v
	return s
}

func (s *OperateEnableJobsResponseBody) SetMessage(v string) *OperateEnableJobsResponseBody {
	s.Message = &v
	return s
}

func (s *OperateEnableJobsResponseBody) SetRequestId(v string) *OperateEnableJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateEnableJobsResponseBody) SetSuccess(v bool) *OperateEnableJobsResponseBody {
	s.Success = &v
	return s
}

type OperateEnableJobsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateEnableJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateEnableJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateEnableJobsResponse) GoString() string {
	return s.String()
}

func (s *OperateEnableJobsResponse) SetHeaders(v map[string]*string) *OperateEnableJobsResponse {
	s.Headers = v
	return s
}

func (s *OperateEnableJobsResponse) SetStatusCode(v int32) *OperateEnableJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateEnableJobsResponse) SetBody(v *OperateEnableJobsResponseBody) *OperateEnableJobsResponse {
	s.Body = v
	return s
}

type OperateExecuteJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// name=zhangsan
	InstanceParameters *string `json:"InstanceParameters,omitempty" xml:"InstanceParameters,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 74
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// gray
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// http://192.168.1.5:9999/
	Worker *string `json:"Worker,omitempty" xml:"Worker,omitempty"`
}

func (s OperateExecuteJobRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateExecuteJobRequest) GoString() string {
	return s.String()
}

func (s *OperateExecuteJobRequest) SetAppName(v string) *OperateExecuteJobRequest {
	s.AppName = &v
	return s
}

func (s *OperateExecuteJobRequest) SetClusterId(v string) *OperateExecuteJobRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateExecuteJobRequest) SetInstanceParameters(v string) *OperateExecuteJobRequest {
	s.InstanceParameters = &v
	return s
}

func (s *OperateExecuteJobRequest) SetJobId(v int64) *OperateExecuteJobRequest {
	s.JobId = &v
	return s
}

func (s *OperateExecuteJobRequest) SetLabel(v string) *OperateExecuteJobRequest {
	s.Label = &v
	return s
}

func (s *OperateExecuteJobRequest) SetWorker(v string) *OperateExecuteJobRequest {
	s.Worker = &v
	return s
}

type OperateExecuteJobResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *OperateExecuteJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6305893D-517D-5131-A767-644EDA81CEC3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateExecuteJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperateExecuteJobResponseBody) GoString() string {
	return s.String()
}

func (s *OperateExecuteJobResponseBody) SetCode(v int32) *OperateExecuteJobResponseBody {
	s.Code = &v
	return s
}

func (s *OperateExecuteJobResponseBody) SetData(v *OperateExecuteJobResponseBodyData) *OperateExecuteJobResponseBody {
	s.Data = v
	return s
}

func (s *OperateExecuteJobResponseBody) SetMessage(v string) *OperateExecuteJobResponseBody {
	s.Message = &v
	return s
}

func (s *OperateExecuteJobResponseBody) SetRequestId(v string) *OperateExecuteJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateExecuteJobResponseBody) SetSuccess(v bool) *OperateExecuteJobResponseBody {
	s.Success = &v
	return s
}

type OperateExecuteJobResponseBodyData struct {
	// example:
	//
	// 1310630367761285120
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
}

func (s OperateExecuteJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OperateExecuteJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *OperateExecuteJobResponseBodyData) SetJobExecutionId(v string) *OperateExecuteJobResponseBodyData {
	s.JobExecutionId = &v
	return s
}

type OperateExecuteJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateExecuteJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateExecuteJobResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateExecuteJobResponse) GoString() string {
	return s.String()
}

func (s *OperateExecuteJobResponse) SetHeaders(v map[string]*string) *OperateExecuteJobResponse {
	s.Headers = v
	return s
}

func (s *OperateExecuteJobResponse) SetStatusCode(v int32) *OperateExecuteJobResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateExecuteJobResponse) SetBody(v *OperateExecuteJobResponseBody) *OperateExecuteJobResponse {
	s.Body = v
	return s
}

type OperateRerunJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 14:11:10
	DataTime *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1698458024000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 74
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1698458024000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s OperateRerunJobRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateRerunJobRequest) GoString() string {
	return s.String()
}

func (s *OperateRerunJobRequest) SetAppName(v string) *OperateRerunJobRequest {
	s.AppName = &v
	return s
}

func (s *OperateRerunJobRequest) SetClusterId(v string) *OperateRerunJobRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateRerunJobRequest) SetDataTime(v string) *OperateRerunJobRequest {
	s.DataTime = &v
	return s
}

func (s *OperateRerunJobRequest) SetEndDate(v int64) *OperateRerunJobRequest {
	s.EndDate = &v
	return s
}

func (s *OperateRerunJobRequest) SetJobId(v int64) *OperateRerunJobRequest {
	s.JobId = &v
	return s
}

func (s *OperateRerunJobRequest) SetStartDate(v int64) *OperateRerunJobRequest {
	s.StartDate = &v
	return s
}

type OperateRerunJobResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// BCDF4006-C8A1-5F83-9368-588347D3EE84
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateRerunJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperateRerunJobResponseBody) GoString() string {
	return s.String()
}

func (s *OperateRerunJobResponseBody) SetCode(v int32) *OperateRerunJobResponseBody {
	s.Code = &v
	return s
}

func (s *OperateRerunJobResponseBody) SetMessage(v string) *OperateRerunJobResponseBody {
	s.Message = &v
	return s
}

func (s *OperateRerunJobResponseBody) SetRequestId(v string) *OperateRerunJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateRerunJobResponseBody) SetSuccess(v bool) *OperateRerunJobResponseBody {
	s.Success = &v
	return s
}

type OperateRerunJobResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateRerunJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateRerunJobResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateRerunJobResponse) GoString() string {
	return s.String()
}

func (s *OperateRerunJobResponse) SetHeaders(v map[string]*string) *OperateRerunJobResponse {
	s.Headers = v
	return s
}

func (s *OperateRerunJobResponse) SetStatusCode(v int32) *OperateRerunJobResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateRerunJobResponse) SetBody(v *OperateRerunJobResponseBody) *OperateRerunJobResponse {
	s.Body = v
	return s
}

type OperateRetryJobExecutionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1310630367761285120
	JobExecutionId *string   `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
	TaskList       []*string `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
}

func (s OperateRetryJobExecutionRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateRetryJobExecutionRequest) GoString() string {
	return s.String()
}

func (s *OperateRetryJobExecutionRequest) SetAppName(v string) *OperateRetryJobExecutionRequest {
	s.AppName = &v
	return s
}

func (s *OperateRetryJobExecutionRequest) SetClusterId(v string) *OperateRetryJobExecutionRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateRetryJobExecutionRequest) SetJobExecutionId(v string) *OperateRetryJobExecutionRequest {
	s.JobExecutionId = &v
	return s
}

func (s *OperateRetryJobExecutionRequest) SetTaskList(v []*string) *OperateRetryJobExecutionRequest {
	s.TaskList = v
	return s
}

type OperateRetryJobExecutionShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1310630367761285120
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
	TaskListShrink *string `json:"TaskList,omitempty" xml:"TaskList,omitempty"`
}

func (s OperateRetryJobExecutionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateRetryJobExecutionShrinkRequest) GoString() string {
	return s.String()
}

func (s *OperateRetryJobExecutionShrinkRequest) SetAppName(v string) *OperateRetryJobExecutionShrinkRequest {
	s.AppName = &v
	return s
}

func (s *OperateRetryJobExecutionShrinkRequest) SetClusterId(v string) *OperateRetryJobExecutionShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateRetryJobExecutionShrinkRequest) SetJobExecutionId(v string) *OperateRetryJobExecutionShrinkRequest {
	s.JobExecutionId = &v
	return s
}

func (s *OperateRetryJobExecutionShrinkRequest) SetTaskListShrink(v string) *OperateRetryJobExecutionShrinkRequest {
	s.TaskListShrink = &v
	return s
}

type OperateRetryJobExecutionResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 438737AC-760A-57D9-B646-B7EF79426243
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateRetryJobExecutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperateRetryJobExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *OperateRetryJobExecutionResponseBody) SetCode(v int32) *OperateRetryJobExecutionResponseBody {
	s.Code = &v
	return s
}

func (s *OperateRetryJobExecutionResponseBody) SetMessage(v string) *OperateRetryJobExecutionResponseBody {
	s.Message = &v
	return s
}

func (s *OperateRetryJobExecutionResponseBody) SetRequestId(v string) *OperateRetryJobExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateRetryJobExecutionResponseBody) SetSuccess(v bool) *OperateRetryJobExecutionResponseBody {
	s.Success = &v
	return s
}

type OperateRetryJobExecutionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateRetryJobExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateRetryJobExecutionResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateRetryJobExecutionResponse) GoString() string {
	return s.String()
}

func (s *OperateRetryJobExecutionResponse) SetHeaders(v map[string]*string) *OperateRetryJobExecutionResponse {
	s.Headers = v
	return s
}

func (s *OperateRetryJobExecutionResponse) SetStatusCode(v int32) *OperateRetryJobExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateRetryJobExecutionResponse) SetBody(v *OperateRetryJobExecutionResponseBody) *OperateRetryJobExecutionResponse {
	s.Body = v
	return s
}

type OperateStopJobExecutionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1310630367761285120
	JobExecutionId *string   `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
	TaskList       []*string `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
}

func (s OperateStopJobExecutionRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateStopJobExecutionRequest) GoString() string {
	return s.String()
}

func (s *OperateStopJobExecutionRequest) SetAppName(v string) *OperateStopJobExecutionRequest {
	s.AppName = &v
	return s
}

func (s *OperateStopJobExecutionRequest) SetClusterId(v string) *OperateStopJobExecutionRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateStopJobExecutionRequest) SetJobExecutionId(v string) *OperateStopJobExecutionRequest {
	s.JobExecutionId = &v
	return s
}

func (s *OperateStopJobExecutionRequest) SetTaskList(v []*string) *OperateStopJobExecutionRequest {
	s.TaskList = v
	return s
}

type OperateStopJobExecutionShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1310630367761285120
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
	TaskListShrink *string `json:"TaskList,omitempty" xml:"TaskList,omitempty"`
}

func (s OperateStopJobExecutionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateStopJobExecutionShrinkRequest) GoString() string {
	return s.String()
}

func (s *OperateStopJobExecutionShrinkRequest) SetAppName(v string) *OperateStopJobExecutionShrinkRequest {
	s.AppName = &v
	return s
}

func (s *OperateStopJobExecutionShrinkRequest) SetClusterId(v string) *OperateStopJobExecutionShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateStopJobExecutionShrinkRequest) SetJobExecutionId(v string) *OperateStopJobExecutionShrinkRequest {
	s.JobExecutionId = &v
	return s
}

func (s *OperateStopJobExecutionShrinkRequest) SetTaskListShrink(v string) *OperateStopJobExecutionShrinkRequest {
	s.TaskListShrink = &v
	return s
}

type OperateStopJobExecutionResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E82D8B33-204D-58E1-8F56-909F6B48F3D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateStopJobExecutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperateStopJobExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *OperateStopJobExecutionResponseBody) SetCode(v int32) *OperateStopJobExecutionResponseBody {
	s.Code = &v
	return s
}

func (s *OperateStopJobExecutionResponseBody) SetMessage(v string) *OperateStopJobExecutionResponseBody {
	s.Message = &v
	return s
}

func (s *OperateStopJobExecutionResponseBody) SetRequestId(v string) *OperateStopJobExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateStopJobExecutionResponseBody) SetSuccess(v bool) *OperateStopJobExecutionResponseBody {
	s.Success = &v
	return s
}

type OperateStopJobExecutionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateStopJobExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateStopJobExecutionResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateStopJobExecutionResponse) GoString() string {
	return s.String()
}

func (s *OperateStopJobExecutionResponse) SetHeaders(v map[string]*string) *OperateStopJobExecutionResponse {
	s.Headers = v
	return s
}

func (s *OperateStopJobExecutionResponse) SetStatusCode(v int32) *OperateStopJobExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateStopJobExecutionResponse) SetBody(v *OperateStopJobExecutionResponseBody) *OperateStopJobExecutionResponse {
	s.Body = v
	return s
}

type UpdateAppRequest struct {
	// example:
	//
	// f312159702f4469585586ed5a6904163v3
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// true
	EnableLog *bool `json:"EnableLog,omitempty" xml:"EnableLog,omitempty"`
	// example:
	//
	// 10
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// This parameter is required.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppRequest) SetAccessToken(v string) *UpdateAppRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateAppRequest) SetAppName(v string) *UpdateAppRequest {
	s.AppName = &v
	return s
}

func (s *UpdateAppRequest) SetClusterId(v string) *UpdateAppRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateAppRequest) SetEnableLog(v bool) *UpdateAppRequest {
	s.EnableLog = &v
	return s
}

func (s *UpdateAppRequest) SetMaxConcurrency(v int32) *UpdateAppRequest {
	s.MaxConcurrency = &v
	return s
}

func (s *UpdateAppRequest) SetTitle(v string) *UpdateAppRequest {
	s.Title = &v
	return s
}

type UpdateAppResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 39AA91C1-7BB7-5934-B15B-FD8E706D76C4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppResponseBody) SetCode(v int32) *UpdateAppResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAppResponseBody) SetMessage(v string) *UpdateAppResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAppResponseBody) SetRequestId(v string) *UpdateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppResponseBody) SetSuccess(v bool) *UpdateAppResponseBody {
	s.Success = &v
	return s
}

type UpdateAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppResponse) SetHeaders(v map[string]*string) *UpdateAppResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppResponse) SetStatusCode(v int32) *UpdateAppResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppResponse) SetBody(v *UpdateAppResponseBody) *UpdateAppResponse {
	s.Body = v
	return s
}

type UpdateClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxljob-c20f7ec9a78
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xx-test-1107
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
}

func (s UpdateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterRequest) GoString() string {
	return s.String()
}

func (s *UpdateClusterRequest) SetClusterId(v string) *UpdateClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateClusterRequest) SetClusterName(v string) *UpdateClusterRequest {
	s.ClusterName = &v
	return s
}

type UpdateClusterResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EEF9AF15-AEEF-5E59-BF7B-BCBB119DC53F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClusterResponseBody) SetCode(v int32) *UpdateClusterResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateClusterResponseBody) SetMessage(v string) *UpdateClusterResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateClusterResponseBody) SetRequestId(v string) *UpdateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateClusterResponseBody) SetSuccess(v bool) *UpdateClusterResponseBody {
	s.Success = &v
	return s
}

type UpdateClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterResponse) GoString() string {
	return s.String()
}

func (s *UpdateClusterResponse) SetHeaders(v map[string]*string) *UpdateClusterResponse {
	s.Headers = v
	return s
}

func (s *UpdateClusterResponse) SetStatusCode(v int32) *UpdateClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateClusterResponse) SetBody(v *UpdateClusterResponseBody) *UpdateClusterResponse {
	s.Body = v
	return s
}

type UpdateJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 3
	AttemptInterval *int32 `json:"AttemptInterval,omitempty" xml:"AttemptInterval,omitempty"`
	// example:
	//
	// workday
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1
	ExecutorBlockStrategy *int32 `json:"ExecutorBlockStrategy,omitempty" xml:"ExecutorBlockStrategy,omitempty"`
	// example:
	//
	// testJobVoidHandler
	JobHandler *string `json:"JobHandler,omitempty" xml:"JobHandler,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 74
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 3
	MaxAttempt *int32 `json:"MaxAttempt,omitempty" xml:"MaxAttempt,omitempty"`
	// example:
	//
	// 1
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// example:
	//
	// test-job
	Name           *string                           `json:"Name,omitempty" xml:"Name,omitempty"`
	NoticeConfig   *UpdateJobRequestNoticeConfig     `json:"NoticeConfig,omitempty" xml:"NoticeConfig,omitempty" type:"Struct"`
	NoticeContacts []*UpdateJobRequestNoticeContacts `json:"NoticeContacts,omitempty" xml:"NoticeContacts,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 1
	RouteStrategy *int32 `json:"RouteStrategy,omitempty" xml:"RouteStrategy,omitempty"`
	// example:
	//
	// 1716902187
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 0 0 4 ? 	- Mon/1
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// example:
	//
	// Asia/Beijing
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
}

func (s UpdateJobRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobRequest) SetAppName(v string) *UpdateJobRequest {
	s.AppName = &v
	return s
}

func (s *UpdateJobRequest) SetAttemptInterval(v int32) *UpdateJobRequest {
	s.AttemptInterval = &v
	return s
}

func (s *UpdateJobRequest) SetCalendar(v string) *UpdateJobRequest {
	s.Calendar = &v
	return s
}

func (s *UpdateJobRequest) SetClusterId(v string) *UpdateJobRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateJobRequest) SetDescription(v string) *UpdateJobRequest {
	s.Description = &v
	return s
}

func (s *UpdateJobRequest) SetExecutorBlockStrategy(v int32) *UpdateJobRequest {
	s.ExecutorBlockStrategy = &v
	return s
}

func (s *UpdateJobRequest) SetJobHandler(v string) *UpdateJobRequest {
	s.JobHandler = &v
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

func (s *UpdateJobRequest) SetName(v string) *UpdateJobRequest {
	s.Name = &v
	return s
}

func (s *UpdateJobRequest) SetNoticeConfig(v *UpdateJobRequestNoticeConfig) *UpdateJobRequest {
	s.NoticeConfig = v
	return s
}

func (s *UpdateJobRequest) SetNoticeContacts(v []*UpdateJobRequestNoticeContacts) *UpdateJobRequest {
	s.NoticeContacts = v
	return s
}

func (s *UpdateJobRequest) SetParameters(v string) *UpdateJobRequest {
	s.Parameters = &v
	return s
}

func (s *UpdateJobRequest) SetPriority(v int32) *UpdateJobRequest {
	s.Priority = &v
	return s
}

func (s *UpdateJobRequest) SetRouteStrategy(v int32) *UpdateJobRequest {
	s.RouteStrategy = &v
	return s
}

func (s *UpdateJobRequest) SetStartTime(v int64) *UpdateJobRequest {
	s.StartTime = &v
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

func (s *UpdateJobRequest) SetTimezone(v string) *UpdateJobRequest {
	s.Timezone = &v
	return s
}

type UpdateJobRequestNoticeConfig struct {
	// example:
	//
	// true
	FailEnable *bool `json:"FailEnable,omitempty" xml:"FailEnable,omitempty"`
	// example:
	//
	// true
	FailLimitTimes *int32 `json:"FailLimitTimes,omitempty" xml:"FailLimitTimes,omitempty"`
	// example:
	//
	// true
	MissWorkerEnable *bool `json:"MissWorkerEnable,omitempty" xml:"MissWorkerEnable,omitempty"`
	// example:
	//
	// webhook,sms,mail,phone
	SendChannel *string `json:"SendChannel,omitempty" xml:"SendChannel,omitempty"`
	// example:
	//
	// true
	SuccessNotice *bool `json:"SuccessNotice,omitempty" xml:"SuccessNotice,omitempty"`
	// example:
	//
	// 90
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// example:
	//
	// true
	TimeoutEnable *bool `json:"TimeoutEnable,omitempty" xml:"TimeoutEnable,omitempty"`
	// example:
	//
	// true
	TimeoutKillEnable *bool `json:"TimeoutKillEnable,omitempty" xml:"TimeoutKillEnable,omitempty"`
}

func (s UpdateJobRequestNoticeConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobRequestNoticeConfig) GoString() string {
	return s.String()
}

func (s *UpdateJobRequestNoticeConfig) SetFailEnable(v bool) *UpdateJobRequestNoticeConfig {
	s.FailEnable = &v
	return s
}

func (s *UpdateJobRequestNoticeConfig) SetFailLimitTimes(v int32) *UpdateJobRequestNoticeConfig {
	s.FailLimitTimes = &v
	return s
}

func (s *UpdateJobRequestNoticeConfig) SetMissWorkerEnable(v bool) *UpdateJobRequestNoticeConfig {
	s.MissWorkerEnable = &v
	return s
}

func (s *UpdateJobRequestNoticeConfig) SetSendChannel(v string) *UpdateJobRequestNoticeConfig {
	s.SendChannel = &v
	return s
}

func (s *UpdateJobRequestNoticeConfig) SetSuccessNotice(v bool) *UpdateJobRequestNoticeConfig {
	s.SuccessNotice = &v
	return s
}

func (s *UpdateJobRequestNoticeConfig) SetTimeout(v int64) *UpdateJobRequestNoticeConfig {
	s.Timeout = &v
	return s
}

func (s *UpdateJobRequestNoticeConfig) SetTimeoutEnable(v bool) *UpdateJobRequestNoticeConfig {
	s.TimeoutEnable = &v
	return s
}

func (s *UpdateJobRequestNoticeConfig) SetTimeoutKillEnable(v bool) *UpdateJobRequestNoticeConfig {
	s.TimeoutKillEnable = &v
	return s
}

type UpdateJobRequestNoticeContacts struct {
	// example:
	//
	// 1
	ContactType *int32 `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
	// example:
	//
	// xiaoming
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateJobRequestNoticeContacts) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobRequestNoticeContacts) GoString() string {
	return s.String()
}

func (s *UpdateJobRequestNoticeContacts) SetContactType(v int32) *UpdateJobRequestNoticeContacts {
	s.ContactType = &v
	return s
}

func (s *UpdateJobRequestNoticeContacts) SetName(v string) *UpdateJobRequestNoticeContacts {
	s.Name = &v
	return s
}

type UpdateJobShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 3
	AttemptInterval *int32 `json:"AttemptInterval,omitempty" xml:"AttemptInterval,omitempty"`
	// example:
	//
	// workday
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1
	ExecutorBlockStrategy *int32 `json:"ExecutorBlockStrategy,omitempty" xml:"ExecutorBlockStrategy,omitempty"`
	// example:
	//
	// testJobVoidHandler
	JobHandler *string `json:"JobHandler,omitempty" xml:"JobHandler,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 74
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 3
	MaxAttempt *int32 `json:"MaxAttempt,omitempty" xml:"MaxAttempt,omitempty"`
	// example:
	//
	// 1
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// example:
	//
	// test-job
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NoticeConfigShrink   *string `json:"NoticeConfig,omitempty" xml:"NoticeConfig,omitempty"`
	NoticeContactsShrink *string `json:"NoticeContacts,omitempty" xml:"NoticeContacts,omitempty"`
	// example:
	//
	// test
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 1
	RouteStrategy *int32 `json:"RouteStrategy,omitempty" xml:"RouteStrategy,omitempty"`
	// example:
	//
	// 1716902187
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 0 0 4 ? 	- Mon/1
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// example:
	//
	// Asia/Beijing
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
}

func (s UpdateJobShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobShrinkRequest) SetAppName(v string) *UpdateJobShrinkRequest {
	s.AppName = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetAttemptInterval(v int32) *UpdateJobShrinkRequest {
	s.AttemptInterval = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetCalendar(v string) *UpdateJobShrinkRequest {
	s.Calendar = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetClusterId(v string) *UpdateJobShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetDescription(v string) *UpdateJobShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetExecutorBlockStrategy(v int32) *UpdateJobShrinkRequest {
	s.ExecutorBlockStrategy = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetJobHandler(v string) *UpdateJobShrinkRequest {
	s.JobHandler = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetJobId(v int64) *UpdateJobShrinkRequest {
	s.JobId = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetMaxAttempt(v int32) *UpdateJobShrinkRequest {
	s.MaxAttempt = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetMaxConcurrency(v int32) *UpdateJobShrinkRequest {
	s.MaxConcurrency = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetName(v string) *UpdateJobShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetNoticeConfigShrink(v string) *UpdateJobShrinkRequest {
	s.NoticeConfigShrink = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetNoticeContactsShrink(v string) *UpdateJobShrinkRequest {
	s.NoticeContactsShrink = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetParameters(v string) *UpdateJobShrinkRequest {
	s.Parameters = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetPriority(v int32) *UpdateJobShrinkRequest {
	s.Priority = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetRouteStrategy(v int32) *UpdateJobShrinkRequest {
	s.RouteStrategy = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetStartTime(v int64) *UpdateJobShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetTimeExpression(v string) *UpdateJobShrinkRequest {
	s.TimeExpression = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetTimeType(v int32) *UpdateJobShrinkRequest {
	s.TimeType = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetTimezone(v string) *UpdateJobShrinkRequest {
	s.Timezone = &v
	return s
}

type UpdateJobResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3808cf26-dde2-4286-8503-b0a2cd4065a7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateJobResponse) SetStatusCode(v int32) *UpdateJobResponse {
	s.StatusCode = &v
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
	client.ProductId = tea.String("SchedulerX3")
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("schedulerx3"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 创建应用
//
// @param request - CreateAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppResponse
func (client *Client) CreateAppWithOptions(request *CreateAppRequest, runtime *util.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		body["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EnableLog)) {
		body["EnableLog"] = request.EnableLog
	}

	if !tea.BoolValue(util.IsUnset(request.MaxConcurrency)) {
		body["MaxConcurrency"] = request.MaxConcurrency
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApp"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateAppResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateAppResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建应用
//
// @param request - CreateAppRequest
//
// @return CreateAppResponse
func (client *Client) CreateApp(request *CreateAppRequest) (_result *CreateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppResponse{}
	_body, _err := client.CreateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建集群
//
// @param tmpReq - CreateClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterResponse
func (client *Client) CreateClusterWithOptions(tmpReq *CreateClusterRequest, runtime *util.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.VSwitches)) {
		request.VSwitchesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VSwitches, tea.String("VSwitches"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		body["ClusterName"] = request.ClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterSpec)) {
		body["ClusterSpec"] = request.ClusterSpec
	}

	if !tea.BoolValue(util.IsUnset(request.EngineType)) {
		body["EngineType"] = request.EngineType
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchesShrink)) {
		body["VSwitches"] = request.VSwitchesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCluster"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateClusterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateClusterResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建集群
//
// @param request - CreateClusterRequest
//
// @return CreateClusterResponse
func (client *Client) CreateCluster(request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建任务
//
// @param tmpReq - CreateJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateJobResponse
func (client *Client) CreateJobWithOptions(tmpReq *CreateJobRequest, runtime *util.RuntimeOptions) (_result *CreateJobResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NoticeConfig)) {
		request.NoticeConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NoticeConfig, tea.String("NoticeConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.NoticeContacts)) {
		request.NoticeContactsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NoticeContacts, tea.String("NoticeContacts"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AttemptInterval)) {
		body["AttemptInterval"] = request.AttemptInterval
	}

	if !tea.BoolValue(util.IsUnset(request.Calendar)) {
		body["Calendar"] = request.Calendar
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorBlockStrategy)) {
		body["ExecutorBlockStrategy"] = request.ExecutorBlockStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.JobHandler)) {
		body["JobHandler"] = request.JobHandler
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

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NoticeConfigShrink)) {
		body["NoticeConfig"] = request.NoticeConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NoticeContactsShrink)) {
		body["NoticeContacts"] = request.NoticeContactsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		body["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.RouteStrategy)) {
		body["RouteStrategy"] = request.RouteStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TimeExpression)) {
		body["TimeExpression"] = request.TimeExpression
	}

	if !tea.BoolValue(util.IsUnset(request.TimeType)) {
		body["TimeType"] = request.TimeType
	}

	if !tea.BoolValue(util.IsUnset(request.Timezone)) {
		body["Timezone"] = request.Timezone
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateJob"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建任务
//
// @param request - CreateJobRequest
//
// @return CreateJobResponse
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

// Summary:
//
// 删除应用分组
//
// @param request - DeleteAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppResponse
func (client *Client) DeleteAppWithOptions(request *DeleteAppRequest, runtime *util.RuntimeOptions) (_result *DeleteAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApp"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteAppResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteAppResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除应用分组
//
// @param request - DeleteAppRequest
//
// @return DeleteAppResponse
func (client *Client) DeleteApp(request *DeleteAppRequest) (_result *DeleteAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppResponse{}
	_body, _err := client.DeleteAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 释放删除集群
//
// @param request - DeleteClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClusterResponse
func (client *Client) DeleteClusterWithOptions(request *DeleteClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCluster"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteClusterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteClusterResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 释放删除集群
//
// @param request - DeleteClusterRequest
//
// @return DeleteClusterResponse
func (client *Client) DeleteCluster(request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DeleteClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量删除任务
//
// @param tmpReq - DeleteJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteJobsResponse
func (client *Client) DeleteJobsWithOptions(tmpReq *DeleteJobsRequest, runtime *util.RuntimeOptions) (_result *DeleteJobsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteJobsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.JobIds)) {
		request.JobIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobIds, tea.String("JobIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.JobIdsShrink)) {
		body["JobIds"] = request.JobIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteJobs"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteJobsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteJobsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 批量删除任务
//
// @param request - DeleteJobsRequest
//
// @return DeleteJobsResponse
func (client *Client) DeleteJobs(request *DeleteJobsRequest) (_result *DeleteJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteJobsResponse{}
	_body, _err := client.DeleteJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量导出任务信息
//
// @param tmpReq - ExportJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportJobsResponse
func (client *Client) ExportJobsWithOptions(tmpReq *ExportJobsRequest, runtime *util.RuntimeOptions) (_result *ExportJobsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ExportJobsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.JobIds)) {
		request.JobIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobIds, tea.String("JobIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ExportJobType)) {
		body["ExportJobType"] = request.ExportJobType
	}

	if !tea.BoolValue(util.IsUnset(request.JobIdsShrink)) {
		body["JobIds"] = request.JobIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportJobs"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("byte"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ExportJobsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ExportJobsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 批量导出任务信息
//
// @param request - ExportJobsRequest
//
// @return ExportJobsResponse
func (client *Client) ExportJobs(request *ExportJobsRequest) (_result *ExportJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportJobsResponse{}
	_body, _err := client.ExportJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取集群详细信息
//
// @param request - GetClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterResponse
func (client *Client) GetClusterWithOptions(request *GetClusterRequest, runtime *util.RuntimeOptions) (_result *GetClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCluster"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetClusterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetClusterResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取集群详细信息
//
// @param request - GetClusterRequest
//
// @return GetClusterResponse
func (client *Client) GetCluster(request *GetClusterRequest) (_result *GetClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetClusterResponse{}
	_body, _err := client.GetClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定机器信息
//
// @param request - GetDesigateInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDesigateInfoResponse
func (client *Client) GetDesigateInfoWithOptions(request *GetDesigateInfoRequest, runtime *util.RuntimeOptions) (_result *GetDesigateInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDesigateInfo"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetDesigateInfoResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetDesigateInfoResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取指定机器信息
//
// @param request - GetDesigateInfoRequest
//
// @return GetDesigateInfoResponse
func (client *Client) GetDesigateInfo(request *GetDesigateInfoRequest) (_result *GetDesigateInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDesigateInfoResponse{}
	_body, _err := client.GetDesigateInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取任务执行的详情
//
// @param request - GetJobExecutionProgressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobExecutionProgressResponse
func (client *Client) GetJobExecutionProgressWithOptions(request *GetJobExecutionProgressRequest, runtime *util.RuntimeOptions) (_result *GetJobExecutionProgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJobExecutionProgress"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetJobExecutionProgressResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetJobExecutionProgressResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取任务执行的详情
//
// @param request - GetJobExecutionProgressRequest
//
// @return GetJobExecutionProgressResponse
func (client *Client) GetJobExecutionProgress(request *GetJobExecutionProgressRequest) (_result *GetJobExecutionProgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobExecutionProgressResponse{}
	_body, _err := client.GetJobExecutionProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询日志
//
// @param request - GetLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLogResponse
func (client *Client) GetLogWithOptions(request *GetLogRequest, runtime *util.RuntimeOptions) (_result *GetLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLog"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetLogResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetLogResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询日志
//
// @param request - GetLogRequest
//
// @return GetLogResponse
func (client *Client) GetLog(request *GetLogRequest) (_result *GetLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLogResponse{}
	_body, _err := client.GetLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导入日历
//
// @param request - ImportCalendarRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportCalendarResponse
func (client *Client) ImportCalendarWithOptions(request *ImportCalendarRequest, runtime *util.RuntimeOptions) (_result *ImportCalendarResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Months)) {
		body["Months"] = request.Months
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Year)) {
		body["Year"] = request.Year
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportCalendar"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ImportCalendarResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ImportCalendarResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 导入日历
//
// @param request - ImportCalendarRequest
//
// @return ImportCalendarResponse
func (client *Client) ImportCalendar(request *ImportCalendarRequest) (_result *ImportCalendarResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportCalendarResponse{}
	_body, _err := client.ImportCalendarWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量导入任务
//
// @param request - ImportJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportJobsResponse
func (client *Client) ImportJobsWithOptions(request *ImportJobsRequest, runtime *util.RuntimeOptions) (_result *ImportJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoCreateApp)) {
		body["AutoCreateApp"] = request.AutoCreateApp
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Overwrite)) {
		body["Overwrite"] = request.Overwrite
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportJobs"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ImportJobsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ImportJobsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 批量导入任务
//
// @param request - ImportJobsRequest
//
// @return ImportJobsResponse
func (client *Client) ImportJobs(request *ImportJobsRequest) (_result *ImportJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportJobsResponse{}
	_body, _err := client.ImportJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取报警事件
//
// @param request - ListAlarmEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlarmEventResponse
func (client *Client) ListAlarmEventWithOptions(request *ListAlarmEventRequest, runtime *util.RuntimeOptions) (_result *ListAlarmEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAlarmEvent"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListAlarmEventResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListAlarmEventResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取报警事件
//
// @param request - ListAlarmEventRequest
//
// @return ListAlarmEventResponse
func (client *Client) ListAlarmEvent(request *ListAlarmEventRequest) (_result *ListAlarmEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAlarmEventResponse{}
	_body, _err := client.ListAlarmEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取应用名字列表
//
// @param request - ListAppNamesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppNamesResponse
func (client *Client) ListAppNamesWithOptions(request *ListAppNamesRequest, runtime *util.RuntimeOptions) (_result *ListAppNamesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAppNames"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListAppNamesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListAppNamesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取应用名字列表
//
// @param request - ListAppNamesRequest
//
// @return ListAppNamesResponse
func (client *Client) ListAppNames(request *ListAppNamesRequest) (_result *ListAppNamesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppNamesResponse{}
	_body, _err := client.ListAppNamesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppsResponse
func (client *Client) ListAppsWithOptions(request *ListAppsRequest, runtime *util.RuntimeOptions) (_result *ListAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApps"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListAppsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListAppsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ListAppsRequest
//
// @return ListAppsResponse
func (client *Client) ListApps(request *ListAppsRequest) (_result *ListAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppsResponse{}
	_body, _err := client.ListAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取日历名字列表
//
// @param request - ListCalendarNamesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCalendarNamesResponse
func (client *Client) ListCalendarNamesWithOptions(request *ListCalendarNamesRequest, runtime *util.RuntimeOptions) (_result *ListCalendarNamesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCalendarNames"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListCalendarNamesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListCalendarNamesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取日历名字列表
//
// @param request - ListCalendarNamesRequest
//
// @return ListCalendarNamesResponse
func (client *Client) ListCalendarNames(request *ListCalendarNamesRequest) (_result *ListCalendarNamesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCalendarNamesResponse{}
	_body, _err := client.ListCalendarNamesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询实例列表
//
// @param request - ListClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClustersResponse
func (client *Client) ListClustersWithOptions(request *ListClustersRequest, runtime *util.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusters"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListClustersResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListClustersResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询实例列表
//
// @param request - ListClustersRequest
//
// @return ListClustersResponse
func (client *Client) ListClusters(request *ListClustersRequest) (_result *ListClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClustersResponse{}
	_body, _err := client.ListClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询Executor列表
//
// @param request - ListExecutorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExecutorsResponse
func (client *Client) ListExecutorsWithOptions(request *ListExecutorsRequest, runtime *util.RuntimeOptions) (_result *ListExecutorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListExecutors"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListExecutorsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListExecutorsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询Executor列表
//
// @param request - ListExecutorsRequest
//
// @return ListExecutorsResponse
func (client *Client) ListExecutors(request *ListExecutorsRequest) (_result *ListExecutorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListExecutorsResponse{}
	_body, _err := client.ListExecutorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取任务实例列表
//
// @param request - ListJobExecutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobExecutionsResponse
func (client *Client) ListJobExecutionsWithOptions(request *ListJobExecutionsRequest, runtime *util.RuntimeOptions) (_result *ListJobExecutionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListJobExecutions"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListJobExecutionsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListJobExecutionsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取任务实例列表
//
// @param request - ListJobExecutionsRequest
//
// @return ListJobExecutionsResponse
func (client *Client) ListJobExecutions(request *ListJobExecutionsRequest) (_result *ListJobExecutionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListJobExecutionsResponse{}
	_body, _err := client.ListJobExecutionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobsResponse
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
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListJobsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListJobsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ListJobsRequest
//
// @return ListJobsResponse
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

// Summary:
//
// 获取executor的label列表
//
// @param request - ListLablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLablesResponse
func (client *Client) ListLablesWithOptions(request *ListLablesRequest, runtime *util.RuntimeOptions) (_result *ListLablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLables"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListLablesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListLablesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取executor的label列表
//
// @param request - ListLablesRequest
//
// @return ListLablesResponse
func (client *Client) ListLables(request *ListLablesRequest) (_result *ListLablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLablesResponse{}
	_body, _err := client.ListLablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取可用区列表
//
// @param request - ListRegionZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegionZoneResponse
func (client *Client) ListRegionZoneWithOptions(runtime *util.RuntimeOptions) (_result *ListRegionZoneResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListRegionZone"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListRegionZoneResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListRegionZoneResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取可用区列表
//
// @return ListRegionZoneResponse
func (client *Client) ListRegionZone() (_result *ListRegionZoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRegionZoneResponse{}
	_body, _err := client.ListRegionZoneWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取所有region列表
//
// @param request - ListRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegionsResponse
func (client *Client) ListRegionsWithOptions(runtime *util.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListRegions"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListRegionsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListRegionsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取所有region列表
//
// @return ListRegionsResponse
func (client *Client) ListRegions() (_result *ListRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询调度事件
//
// @param request - ListScheduleEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScheduleEventResponse
func (client *Client) ListScheduleEventWithOptions(request *ListScheduleEventRequest, runtime *util.RuntimeOptions) (_result *ListScheduleEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListScheduleEvent"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListScheduleEventResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListScheduleEventResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询调度事件
//
// @param request - ListScheduleEventRequest
//
// @return ListScheduleEventResponse
func (client *Client) ListScheduleEvent(request *ListScheduleEventRequest) (_result *ListScheduleEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListScheduleEventResponse{}
	_body, _err := client.ListScheduleEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定时间类型和表达式未来5次调度时间
//
// @param request - ListScheduleTimesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScheduleTimesResponse
func (client *Client) ListScheduleTimesWithOptions(request *ListScheduleTimesRequest, runtime *util.RuntimeOptions) (_result *ListScheduleTimesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListScheduleTimes"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListScheduleTimesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListScheduleTimesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取指定时间类型和表达式未来5次调度时间
//
// @param request - ListScheduleTimesRequest
//
// @return ListScheduleTimesResponse
func (client *Client) ListScheduleTimes(request *ListScheduleTimesRequest) (_result *ListScheduleTimesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListScheduleTimesResponse{}
	_body, _err := client.ListScheduleTimesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 指定执行器
//
// @param tmpReq - OperateDesignateExecutorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateDesignateExecutorsResponse
func (client *Client) OperateDesignateExecutorsWithOptions(tmpReq *OperateDesignateExecutorsRequest, runtime *util.RuntimeOptions) (_result *OperateDesignateExecutorsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &OperateDesignateExecutorsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AddressList)) {
		request.AddressListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddressList, tea.String("AddressList"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressListShrink)) {
		body["AddressList"] = request.AddressListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DesignateType)) {
		body["DesignateType"] = request.DesignateType
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.Transferable)) {
		body["Transferable"] = request.Transferable
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("OperateDesignateExecutors"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OperateDesignateExecutorsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OperateDesignateExecutorsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 指定执行器
//
// @param request - OperateDesignateExecutorsRequest
//
// @return OperateDesignateExecutorsResponse
func (client *Client) OperateDesignateExecutors(request *OperateDesignateExecutorsRequest) (_result *OperateDesignateExecutorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OperateDesignateExecutorsResponse{}
	_body, _err := client.OperateDesignateExecutorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量禁用任务
//
// @param tmpReq - OperateDisableJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateDisableJobsResponse
func (client *Client) OperateDisableJobsWithOptions(tmpReq *OperateDisableJobsRequest, runtime *util.RuntimeOptions) (_result *OperateDisableJobsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &OperateDisableJobsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.JobIds)) {
		request.JobIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobIds, tea.String("JobIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.JobIdsShrink)) {
		body["JobIds"] = request.JobIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("OperateDisableJobs"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OperateDisableJobsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OperateDisableJobsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 批量禁用任务
//
// @param request - OperateDisableJobsRequest
//
// @return OperateDisableJobsResponse
func (client *Client) OperateDisableJobs(request *OperateDisableJobsRequest) (_result *OperateDisableJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OperateDisableJobsResponse{}
	_body, _err := client.OperateDisableJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量启用任务
//
// @param tmpReq - OperateEnableJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateEnableJobsResponse
func (client *Client) OperateEnableJobsWithOptions(tmpReq *OperateEnableJobsRequest, runtime *util.RuntimeOptions) (_result *OperateEnableJobsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &OperateEnableJobsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.JobIds)) {
		request.JobIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobIds, tea.String("JobIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.JobIdsShrink)) {
		body["JobIds"] = request.JobIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("OperateEnableJobs"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OperateEnableJobsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OperateEnableJobsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 批量启用任务
//
// @param request - OperateEnableJobsRequest
//
// @return OperateEnableJobsResponse
func (client *Client) OperateEnableJobs(request *OperateEnableJobsRequest) (_result *OperateEnableJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OperateEnableJobsResponse{}
	_body, _err := client.OperateEnableJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 运行一次任务
//
// @param request - OperateExecuteJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateExecuteJobResponse
func (client *Client) OperateExecuteJobWithOptions(request *OperateExecuteJobRequest, runtime *util.RuntimeOptions) (_result *OperateExecuteJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceParameters)) {
		body["InstanceParameters"] = request.InstanceParameters
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.Label)) {
		body["Label"] = request.Label
	}

	if !tea.BoolValue(util.IsUnset(request.Worker)) {
		body["Worker"] = request.Worker
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("OperateExecuteJob"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OperateExecuteJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OperateExecuteJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 运行一次任务
//
// @param request - OperateExecuteJobRequest
//
// @return OperateExecuteJobResponse
func (client *Client) OperateExecuteJob(request *OperateExecuteJobRequest) (_result *OperateExecuteJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OperateExecuteJobResponse{}
	_body, _err := client.OperateExecuteJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重刷任务历史数据
//
// @param request - OperateRerunJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateRerunJobResponse
func (client *Client) OperateRerunJobWithOptions(request *OperateRerunJobRequest, runtime *util.RuntimeOptions) (_result *OperateRerunJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DataTime)) {
		query["DataTime"] = request.DataTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OperateRerunJob"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OperateRerunJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OperateRerunJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 重刷任务历史数据
//
// @param request - OperateRerunJobRequest
//
// @return OperateRerunJobResponse
func (client *Client) OperateRerunJob(request *OperateRerunJobRequest) (_result *OperateRerunJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OperateRerunJobResponse{}
	_body, _err := client.OperateRerunJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重跑失败的任务实例
//
// @param tmpReq - OperateRetryJobExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateRetryJobExecutionResponse
func (client *Client) OperateRetryJobExecutionWithOptions(tmpReq *OperateRetryJobExecutionRequest, runtime *util.RuntimeOptions) (_result *OperateRetryJobExecutionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &OperateRetryJobExecutionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TaskList)) {
		request.TaskListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskList, tea.String("TaskList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.JobExecutionId)) {
		query["JobExecutionId"] = request.JobExecutionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskListShrink)) {
		query["TaskList"] = request.TaskListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OperateRetryJobExecution"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OperateRetryJobExecutionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OperateRetryJobExecutionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 重跑失败的任务实例
//
// @param request - OperateRetryJobExecutionRequest
//
// @return OperateRetryJobExecutionResponse
func (client *Client) OperateRetryJobExecution(request *OperateRetryJobExecutionRequest) (_result *OperateRetryJobExecutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OperateRetryJobExecutionResponse{}
	_body, _err := client.OperateRetryJobExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止正在运行的任务实例
//
// @param tmpReq - OperateStopJobExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateStopJobExecutionResponse
func (client *Client) OperateStopJobExecutionWithOptions(tmpReq *OperateStopJobExecutionRequest, runtime *util.RuntimeOptions) (_result *OperateStopJobExecutionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &OperateStopJobExecutionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TaskList)) {
		request.TaskListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskList, tea.String("TaskList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.JobExecutionId)) {
		query["JobExecutionId"] = request.JobExecutionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskListShrink)) {
		query["TaskList"] = request.TaskListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OperateStopJobExecution"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OperateStopJobExecutionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OperateStopJobExecutionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 停止正在运行的任务实例
//
// @param request - OperateStopJobExecutionRequest
//
// @return OperateStopJobExecutionResponse
func (client *Client) OperateStopJobExecution(request *OperateStopJobExecutionRequest) (_result *OperateStopJobExecutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OperateStopJobExecutionResponse{}
	_body, _err := client.OperateStopJobExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新应用分组
//
// @param request - UpdateAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAppResponse
func (client *Client) UpdateAppWithOptions(request *UpdateAppRequest, runtime *util.RuntimeOptions) (_result *UpdateAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		body["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EnableLog)) {
		body["EnableLog"] = request.EnableLog
	}

	if !tea.BoolValue(util.IsUnset(request.MaxConcurrency)) {
		body["MaxConcurrency"] = request.MaxConcurrency
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateApp"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateAppResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateAppResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新应用分组
//
// @param request - UpdateAppRequest
//
// @return UpdateAppResponse
func (client *Client) UpdateApp(request *UpdateAppRequest) (_result *UpdateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAppResponse{}
	_body, _err := client.UpdateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新集群
//
// @param request - UpdateClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateClusterResponse
func (client *Client) UpdateClusterWithOptions(request *UpdateClusterRequest, runtime *util.RuntimeOptions) (_result *UpdateClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		query["ClusterName"] = request.ClusterName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCluster"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateClusterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateClusterResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新集群
//
// @param request - UpdateClusterRequest
//
// @return UpdateClusterResponse
func (client *Client) UpdateCluster(request *UpdateClusterRequest) (_result *UpdateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateClusterResponse{}
	_body, _err := client.UpdateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新任务信息
//
// @param tmpReq - UpdateJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateJobResponse
func (client *Client) UpdateJobWithOptions(tmpReq *UpdateJobRequest, runtime *util.RuntimeOptions) (_result *UpdateJobResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NoticeConfig)) {
		request.NoticeConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NoticeConfig, tea.String("NoticeConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.NoticeContacts)) {
		request.NoticeContactsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NoticeContacts, tea.String("NoticeContacts"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AttemptInterval)) {
		body["AttemptInterval"] = request.AttemptInterval
	}

	if !tea.BoolValue(util.IsUnset(request.Calendar)) {
		body["Calendar"] = request.Calendar
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorBlockStrategy)) {
		body["ExecutorBlockStrategy"] = request.ExecutorBlockStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.JobHandler)) {
		body["JobHandler"] = request.JobHandler
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

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NoticeConfigShrink)) {
		body["NoticeConfig"] = request.NoticeConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NoticeContactsShrink)) {
		body["NoticeContacts"] = request.NoticeContactsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		body["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.RouteStrategy)) {
		body["RouteStrategy"] = request.RouteStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TimeExpression)) {
		body["TimeExpression"] = request.TimeExpression
	}

	if !tea.BoolValue(util.IsUnset(request.TimeType)) {
		body["TimeType"] = request.TimeType
	}

	if !tea.BoolValue(util.IsUnset(request.Timezone)) {
		body["Timezone"] = request.Timezone
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateJob"),
		Version:     tea.String("2024-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新任务信息
//
// @param request - UpdateJobRequest
//
// @return UpdateJobResponse
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
